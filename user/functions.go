package user

import (
	"crypto/rand"
	"encoding/base64"
)

/*
This function returns a URL-safe, base64 encoded securely generated random string
*/
func generateRandomString(size int) (string, error) {
	byteData, err := generateRandomBytes(size)
	return base64.URLEncoding.EncodeToString(byteData), err
}

/*
This function is used to securely generated random bytes
*/

func generateRandomBytes(size int) ([]byte, error) {
	token := make([]byte, size)
	_, err := rand.Read(token)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}
	return token, nil
}
