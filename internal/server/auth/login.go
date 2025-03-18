package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"kaffino/internal/database"
)

const lockoutDuration = 5 * time.Minute

type response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

var (
	failedLogins = make(map[string]loginAttempt)
)

type loginAttempt struct {
	Attempts    int
	Lockout     time.Time
	LastAttempt time.Time
}

type loginRequest struct {
	Email string `json:"email"`
	OTP   string `json:"otp"`
}

func jsonResponse(w http.ResponseWriter, status int, resp response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
	}
}

func VerifyOTPHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, response{Success: false, Error: err.Error()})
		return
	}

	if r.Method != http.MethodPost {
		jsonResponse(w, http.StatusMethodNotAllowed, response{Success: false, Error: "Method not allowed"})
		return
	}

	var req loginRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		jsonResponse(w, http.StatusBadRequest, response{Success: false, Error: "Error decoding JSON body"})
		return
	}
	
	email := req.Email
	otp := req.OTP

	attempt, ok := failedLogins[email]
	if ok && attempt.Lockout.After(time.Now()) {
		remaining := time.Until(attempt.Lockout).String()
		jsonResponse(w, http.StatusTooManyRequests, response{Success: false, Error: fmt.Sprintf("Too many failed attempts. Please try again in %s", remaining)})
		return
	}

	storedOTP := RetrieveOTP(email)

	if storedOTP == "" {
		jsonResponse(w, http.StatusBadRequest, response{Success: false, Error: "Invalid OTP"})
		return
	}
	log.Println(storedOTP)

	if otp != storedOTP {
		attempt.Attempts++
		attempt.LastAttempt = time.Now()
		if attempt.Attempts >= 5 {
			attempt.Lockout = time.Now().Add(lockoutDuration)
			jsonResponse(w, http.StatusTooManyRequests, response{Success: false, Error: "Too many failed attempts. Account locked for 5 minutes."})
			failedLogins[email] = attempt
			return
		}
		failedLogins[email] = attempt
		jsonResponse(w, http.StatusBadRequest, response{Success: false, Error: "Invalid OTP"})
		return
	}

	// Reset failed attempts on successful login
	delete(failedLogins, email)

	db := database.NewDB()
	userID, err := db.GetUserID(email)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, response{Success: false, Error: err.Error()})
		return
	}
	session.Values["userID"] = userID
	session.Values["username"] = email // Store the user ID in the session
	err = session.Save(r, w)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, response{Success: false, Error: err.Error()})
		return
	}

	// Return success response
	jsonResponse(w, http.StatusOK, response{Success: true, Message: "Login successful", Data: map[string]interface{}{"userID": userID}})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, response{Success: false, Error: err.Error()})
		return
	}
	id, ok := session.Values["userID"]
	if ok && !strings.HasPrefix(id.(string), "schrödinger-") {
		jsonResponse(w, http.StatusOK, response{Success: true, Message: "Already logged in", Data: map[string]string{"userID": id.(string)}})
		return
	}

	if r.Method != http.MethodPost {
		jsonResponse(w, http.StatusMethodNotAllowed, response{Success: false, Error: "Method not allowed"})
		return
	}

	var req loginRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		jsonResponse(w, http.StatusBadRequest, response{Success: false, Error: "Error decoding JSON body"})
		return
	}

	email := req.Email

	attempt, ok := failedLogins[email]
	if ok && attempt.Lockout.After(time.Now()) {
		remaining := time.Until(attempt.Lockout).String()
		jsonResponse(w, http.StatusTooManyRequests, response{Success: false, Error: fmt.Sprintf("Too many failed attempts. Please try again in %s", remaining)})
		return
	}

	otp, err := GenerateOTP(6) // Generate a 6-digit OTP
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, response{Success: false, Error: "Failed to generate OTP"})
		return
	}

	StoreOTP(email, otp)

	subject := "Your OTP for Login"
	body := fmt.Sprintf("Your OTP is: %s", otp)

	err = SendEmail(email, subject, body)
	if err != nil {
		log.Printf("Error sending email: %v", err)
		jsonResponse(w, http.StatusInternalServerError, response{Success: false, Error: "Failed to send email, try again later."})
		return
	}

	fmt.Println("OTP sent to:", email)

	// Return success response
	jsonResponse(w, http.StatusOK, response{Success: true, Message: "OTP sent successfully", Data: map[string]string{"email": email}})
}
