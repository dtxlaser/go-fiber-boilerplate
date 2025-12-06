package config

import (
	"log"
	"os"
	"strconv"
)

var (
	IsProd              bool
	AppHost             string
	AppPort             int
	DBHost              string
	DBUser              string
	DBPassword          string
	DBName              string
	DBPort              int
	JWTSecret           string
	JWTAccessExp        int
	JWTRefreshExp       int
	JWTResetPasswordExp int
	JWTVerifyEmailExp   int
	SMTPHost            string
	SMTPPort            int
	SMTPUsername        string
	SMTPPassword        string
	EmailFrom           string
	GoogleClientID      string
	GoogleClientSecret  string
	RedirectURL         string
)

func init() {
	var err error

	// server configuration
	IsProd = os.Getenv("APP_ENV") == "prod"
	AppHost = os.Getenv("APP_HOST")
	AppPort, err = strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		log.Fatalf("Invalid APP_PORT: %v", err)
	}

	// database configuration
	DBHost = os.Getenv("DB_HOST")
	DBUser = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBName = os.Getenv("DB_NAME")
	DBPort, err = strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("Invalid DB_PORT: %v", err)
	}

	// jwt configuration
	JWTSecret = os.Getenv("JWT_SECRET")
	JWTAccessExp, _ = strconv.Atoi(os.Getenv("JWT_ACCESS_EXP_MINUTES"))
	JWTRefreshExp, _ = strconv.Atoi(os.Getenv("JWT_REFRESH_EXP_DAYS"))
	JWTResetPasswordExp, _ = strconv.Atoi(os.Getenv("JWT_RESET_PASSWORD_EXP_MINUTES"))
	JWTVerifyEmailExp, _ = strconv.Atoi(os.Getenv("JWT_VERIFY_EMAIL_EXP_MINUTES"))

	// SMTP configuration
	SMTPHost = os.Getenv("SMTP_HOST")
	SMTPPort, _ = strconv.Atoi(os.Getenv("SMTP_PORT"))
	SMTPUsername = os.Getenv("SMTP_USERNAME")
	SMTPPassword = os.Getenv("SMTP_PASSWORD")
	EmailFrom = os.Getenv("EMAIL_FROM")

	// oauth2 configuration
	GoogleClientID = os.Getenv("GOOGLE_CLIENT_ID")
	GoogleClientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")
	RedirectURL = os.Getenv("REDIRECT_URL")
}
