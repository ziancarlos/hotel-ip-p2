package helper

import (
	"crypto/sha512"
	"fmt"
)

func ValidateMidtransSignature(orderID, statusCode, grossAmount, signatureKey string) bool {
	serverKey := AppConfig.GetMidtransServerKey()

	signatureString := orderID + statusCode + grossAmount + serverKey

	hash := sha512.New()
	hash.Write([]byte(signatureString))
	calculatedSignature := fmt.Sprintf("%x", hash.Sum(nil))

	return calculatedSignature == signatureKey
}
