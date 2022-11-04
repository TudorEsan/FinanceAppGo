package helpers_test

import (
	"os"
	"testing"

	"github.com/TudorEsan/FinanceAppGo/BrokerService/helpers"
	"github.com/stretchr/testify/assert"
)

func TestEncrypt(t *testing.T) {
	os.Setenv("ENCRYPTION_KEY", "1234567891234567")
	os.Setenv("TESTING", "true")
	textToEncrypt := "Hello World"

	encryptedText, err := helpers.Encrypt(textToEncrypt)
	assert.Nil(t, err)
	assert.NotEqual(t, encryptedText, textToEncrypt)
	assert.NotEqual(t, encryptedText, "")
}

func TestDecrypt(t *testing.T) {
	os.Setenv("ENCRYPTION_KEY", "1234567891234567")
	os.Setenv("TESTING", "true")
	textToEncrypt := "Hello World"

	encryptedText, err := helpers.Encrypt(textToEncrypt)
	assert.Nil(t, err)

	decryptedText, err := helpers.Decrypt(encryptedText)
	assert.Nil(t, err)
	t.Log(decryptedText)
	t.Log(textToEncrypt)

	assert.Equal(t, decryptedText, textToEncrypt)

}
