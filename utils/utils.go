package utils

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type FirebaseCredentials struct {
	ProjectID string `json:"project_id"`
}

type AppConfig struct {
	ServerPort int
	ClientPort int
}

// RegisterRequest struct for validation
type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoadConfig() (*AppConfig, error) {
	if err := godotenv.Load("P:/blog-web/.env"); err != nil {
		return nil, errors.New("error loading .env file")
	}

	// Get server port
	serverPortStr := os.Getenv("SERVER_PORT")
	serverPort, err := strconv.Atoi(serverPortStr)
	if err != nil || serverPort <= 0 {
		return nil, errors.New("invalid SERVER_PORT in .env file")
	}

	// Get client port
	clientPortStr := os.Getenv("CLIENT_PORT")
	clientPort, err := strconv.Atoi(clientPortStr)
	if err != nil || clientPort <= 0 {
		return nil, errors.New("invalid CLIENT_PORT in .env file")
	}

	return &AppConfig{
		ServerPort: serverPort,
		ClientPort: clientPort,
	}, nil
}

func GetFirebaseProjectID(credentialsPath string) (string, error) {
	// Read the credentials file
	data, err := ioutil.ReadFile(credentialsPath)
	if err != nil {
		return "", errors.New("failed to read Firebase credentials file")
	}

	// Parse the JSON
	var creds FirebaseCredentials
	if err := json.Unmarshal(data, &creds); err != nil {
		return "", errors.New("failed to parse Firebase credentials JSON")
	}

	if creds.ProjectID == "" {
		return "", errors.New("project_id not found in Firebase credentials")
	}

	return creds.ProjectID, nil
}

// ValidateRegistrationData - exported function for validation
func ValidateRegistrationData(req RegisterRequest) error {
	// Check if fields are empty
	if strings.TrimSpace(req.Username) == "" {
		return errors.New("username cannot be empty")
	}
	if strings.TrimSpace(req.Email) == "" {
		return errors.New("email cannot be empty")
	}
	if strings.TrimSpace(req.Password) == "" {
		return errors.New("password cannot be empty")
	}

	// Username validation (max 20 characters)
	if len(req.Username) > 20 {
		return errors.New("username must not exceed 20 characters")
	}

	// Email validation (only gmail.com allowed)
	if !strings.HasSuffix(strings.ToLower(req.Email), "@gmail.com") {
		return errors.New("only @gmail.com email addresses are allowed")
	}

	// Email format validation
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@gmail\.com$`)
	if !emailRegex.MatchString(req.Email) {
		return errors.New("invalid email format")
	}

	// Password validation (essential requirements)
	if err := validatePassword(req.Password); err != nil {
		return err
	}

	return nil
}

func validatePassword(password string) error {
	// Minimum length check
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	// Maximum length check
	if len(password) > 128 {
		return errors.New("password must not exceed 128 characters")
	}

	// Check for uppercase letter
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	if !hasUpper {
		return errors.New("password must contain at least one uppercase letter")
	}

	// Check for lowercase letter
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	if !hasLower {
		return errors.New("password must contain at least one lowercase letter")
	}

	// Check for digit
	hasDigit := regexp.MustCompile(`[0-9]`).MatchString(password)
	if !hasDigit {
		return errors.New("password must contain at least one digit")
	}

	// Check for special character
	hasSpecial := regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]`).MatchString(password)
	if !hasSpecial {
		return errors.New("password must contain at least one special character")
	}

	return nil
}

func IsValidImageType(contentType string) bool {
	validTypes := []string{
		"image/jpeg",
		"image/jpg",
		"image/png",
		"image/webp",
		"image/gif",
	}

	for _, validType := range validTypes {
		if strings.HasPrefix(contentType, validType) {
			return true
		}
	}
	return false
}	

func ConvertImageToBase64(imageBytes []byte) string {
	if len(imageBytes) == 0 {
		return ""
	}
	return base64.StdEncoding.EncodeToString(imageBytes)
}
