package configs

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort            string
	Dsn                   string
	AppSecret             string
	TwilioAccountSid      string
	TwilioAuthToken       string
	TwilioFromPhoneNumber string
	StripeSecret          string
	SuccessUrl            string
	CancelUrl             string
	PubKey                string
}

func SetupEnv() (cfg AppConfig, err error) {
	if os.Getenv("APP_ENV") == "dev" {
		godotenv.Load()
	}

	// godotenv.Load()

	//check env port
	httpPort := os.Getenv("HTTP_PORT")
	if len(httpPort) < 1 {
		return AppConfig{}, errors.New("env variables not found")
	}

	// check env dsn
	// Dsn := os.Getenv("DSN")
	Dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	if len(Dsn) < 1 {
		return AppConfig{}, errors.New("env variables not found")
	}

	appSecret := os.Getenv("APP_SECRET")
	if len(Dsn) < 1 {
		return AppConfig{}, errors.New("app secret not found")
	}

	return AppConfig{ServerPort: httpPort, Dsn: Dsn, AppSecret: appSecret,
		TwilioAccountSid:      os.Getenv("TWILIO_ACCOUNT_SID"),
		TwilioAuthToken:       os.Getenv("TWILIO_AUTH_TOKEN"),
		TwilioFromPhoneNumber: os.Getenv("TWILIO_FROM_PHONE_NUMBER"),
		StripeSecret:          os.Getenv("STRIPE_SECRET"),
		SuccessUrl:            os.Getenv("SUCCESS_URL"),
		CancelUrl:             os.Getenv("CANCEL_URL"),
		PubKey:                os.Getenv("STRIPE_PUBLIC_KEY"),
	}, nil
}
