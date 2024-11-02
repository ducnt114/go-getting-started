package utils

import (
	"strings"

	"github.com/dgryski/dgoogauth"
)

const (
	TwoFASecretLength = 16
)

func GenerateTwoFaSecret() string {
	randomString := RandomStringWithLength(TwoFASecretLength)
	randomString = strings.ToUpper(randomString)
	return randomString
}

func Validate2FaOTP(secret, otp string) (bool, error) {
	otpConf := &dgoogauth.OTPConfig{
		Secret:     secret,
		WindowSize: 3,
		UTC:        true,
	}
	return otpConf.Authenticate(otp)
}
