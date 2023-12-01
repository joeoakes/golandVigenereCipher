package main

import "fmt"

// Vigenere Cipher Encryption Function
func vigenereEncrypt(plainText string, key string) string {
	cipherText := ""
	keyIndex := 0
	keyLength := len(key)

	for _, char := range plainText {
		if char >= 'A' && char <= 'Z' {
			shift := key[keyIndex] - 'A'
			cipherText += string((char+rune(shift)-'A')%26 + 'A')
			keyIndex = (keyIndex + 1) % keyLength
		} else if char >= 'a' && char <= 'z' {
			shift := key[keyIndex] - 'a'
			cipherText += string((char+rune(shift)-'a')%26 + 'a')
			keyIndex = (keyIndex + 1) % keyLength
		} else {
			cipherText += string(char)
		}
	}

	return cipherText
}

// Vigenere Cipher Decryption Function
func vigenereDecrypt(cipherText string, key string) string {
	plainText := ""
	keyIndex := 0

	for _, char := range cipherText {
		if char >= 'A' && char <= 'Z' {
			shift := int(key[keyIndex] - 'A')
			plainText += string((char-'A'-rune(shift)+26)%26 + 'A')
			keyIndex = (keyIndex + 1) % len(key)
		} else if char >= 'a' && char <= 'z' {
			shift := int(key[keyIndex] - 'a')
			plainText += string((char-'a'-rune(shift)+26)%26 + 'a')
			keyIndex = (keyIndex + 1) % len(key)
		} else {
			plainText += string(char)
		}
	}

	return plainText
}

func main() {
	plainText := "HELLO WORLD"
	key := "SECRET"

	cipherText := vigenereEncrypt(plainText, key)

	fmt.Printf("Plaintext: %s\n", plainText)
	fmt.Printf("Key: %s\n", key)
	fmt.Printf("Ciphertext: %s\n", cipherText)

	plainText = vigenereDecrypt(cipherText, key)
	fmt.Printf("Plaintext: %s\n", plainText)
}
