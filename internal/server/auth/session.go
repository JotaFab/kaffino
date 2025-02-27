package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
)

// Session Key (Keep this secret!)
var sessionKey = []byte(os.Getenv("SESSION_KEY")) // Get from environment
var store *sessions.CookieStore

func init() {
	if len(sessionKey) == 0 {
		sessionKey = []byte("super-secret-key") // Development fallback
		log.Println("Warning: Using default session key.  Set SESSION_KEY environment variable in production!")
	}
	store = sessions.NewCookieStore(sessionKey)
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600 * 2, // 8 hours
		HttpOnly: true,
		Secure:   false, // Set to true in production
		SameSite: http.SameSiteLaxMode,
	}
}

type contextKey string

const userContextKey contextKey = "user"

// GenerateGuestUserID generates a unique ID for guest users.
func GenerateGuestUserID() string {
	return "schrödinger-" + uuid.New().String()
}

// SessionMiddleware is middleware that checks for a session cookie and retrieves the user information.
func SessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, "session-name") // Get session, create if doesn't exist
		log.Println("Session:", session)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		userID := session.Values["userID"]
		if userID == nil {
			// No session, continue to the next handler
			guestUserID := GenerateGuestUserID()
			session.Values["userID"] = guestUserID
			session.Values["username"] = ""
			session.Values["loginFailTries"] = 0
			err = session.Save(r, w)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			userID = guestUserID
		}
		username := session.Values["username"]

		var disallowedGuestRoutes = []string{
			"/create-product",
			"/profile",
			"/update-product", // Add this route
		}

		if strings.HasPrefix(userID.(string), "schrödinger-") {
			for _, route := range disallowedGuestRoutes {
				if r.URL.Path == route {
					// Redirect guest users to the login page or an unauthorized page
					http.Redirect(w, r, "/login", http.StatusSeeOther) // Or "/unauthorized"
					return
				}
			}
		}

		// Session is valid, add the user information to the request context
		ctx := context.WithValue(r.Context(), "userID", userID.(string))
		ctx = context.WithValue(ctx, userContextKey, username.(string))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("username").(string)
	if !ok {
		// User is not authenticated, redirect to login
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	fmt.Fprintf(w, "<h1>Welcome, %s!</h1>", user) // Display the user's email
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Clear session values
	session.Values["userID"] = nil
	session.Options.MaxAge = -1 // Expire the cookie

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
