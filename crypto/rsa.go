package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func Cripting(pass string, privKey rsa.PrivateKey) string {
	//privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	privateKey := privKey
	//CheckError(err)
	publicKey := privateKey.PublicKey
	secretMessage := pass

	encryptedMessage := RSA_OAEP_Encrypt(secretMessage, publicKey)

	//fmt.Println("Cipher Text:", encryptedMessage)

	//RSA_OAEP_Decrypt(encryptedMessage, privateKey)
	return encryptedMessage
}

func CheckError(e error) {
	if e != nil {
		fmt.Println(e.Error)
	}
}

func RSA_OAEP_Encrypt(secretMessage string, key rsa.PublicKey) string {
	label := []byte("OAEP Encrypted")
	rng := rand.Reader
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rng, &key,
		[]byte(secretMessage), label)
	CheckError(err)
	return base64.StdEncoding.EncodeToString(ciphertext)
}

func RSA_OAEP_Decrypt(cipherText string, privKey rsa.PrivateKey) string {
	ct, _ := base64.StdEncoding.DecodeString(cipherText)
	label := []byte("OAEP Encrypted")
	rng := rand.Reader
	plaintext, err := rsa.DecryptOAEP(sha256.New(), rng, &privKey, ct, label)
	CheckError(err)
	//fmt.Println("Your data is:", string(plaintext))
	return string(plaintext)
}
