package utils

import (
	"fmt"
	"github.com/dgryski/dgoogauth"
	"testing"
	"time"
)

func TestGenTwoFaOTP(t *testing.T) {
	secret := GenerateTwoFaSecret()
	fmt.Println(secret)

	if len(secret) != TwoFASecretLength {
		t.Fail()
	}
}

func TestValidate2FaOTP(t *testing.T) {
	t0 := int64(time.Now().UTC().Unix() / 30)
	otp := dgoogauth.ComputeCode("WGFXLRRVVUSRNNEV", t0)

	isValid, err := Validate2FaOTP("WGFXLRRVVUSRNNEV", fmt.Sprintf("%v", otp))
	if err != nil {
		t.Fatal(err)
	}
	if !isValid {
		t.Fail()
	}
}
