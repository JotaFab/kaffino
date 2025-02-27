package auth

import (
	"crypto/rand"
	"math/big"
	"sync"
	"time"
)

// GenerateOTP generates a random OTP of the specified length.
func GenerateOTP(length int) (string, error) {
	const digits = "0123456789"
	max := big.NewInt(int64(len(digits)))
	otp := ""

	for i := 0; i < length; i++ {
		index, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}
		otp += string(digits[index.Int64()])
	}

	return otp, nil
}

// OTPData stores the OTP and the creation time.
type OTPData struct {
	OTP       string
	CreatedAt time.Time
}

var (
	otpStorage    = make(map[string]OTPData)
	otpStorageMu  sync.Mutex
	otpExpiration = 5 * time.Minute // OTP expires after 5 minutes
)

// StoreOTP stores the OTP for the given email.
func StoreOTP(email, otp string) {
	otpStorageMu.Lock()
	defer otpStorageMu.Unlock()
	otpStorage[email] = OTPData{OTP: otp, CreatedAt: time.Now()}
}

// RetrieveOTP retrieves the OTP for the given email.  Returns an empty string if not found or expired.
func RetrieveOTP(email string) string {
	otpStorageMu.Lock()
	defer otpStorageMu.Unlock()
	data, ok := otpStorage[email]
	if !ok || time.Since(data.CreatedAt) > otpExpiration {
		delete(otpStorage, email) // Remove expired OTP
		return ""
	}
	return data.OTP
}
