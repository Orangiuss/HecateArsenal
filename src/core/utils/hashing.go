package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/md4"
)

// MD5 hashing
func HashMD5(data string) string {
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}

// SHA1 hashing
func HashSHA1(data string) string {
	hash := sha1.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}

// SHA256 hashing
func HashSHA256(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

// SHA512 hashing
func HashSHA512(data string) string {
	hash := sha512.Sum512([]byte(data))
	return hex.EncodeToString(hash[:])
}

// NTLM hashing
// NTLM uses a legacy hash system for Windows authentication.
func HashNTLM(data string) string {
	hash := md4Sum(utf16leEncode(data))
	return hex.EncodeToString(hash)
}

// Helper function to create an MD4 hash (used in NTLM)
func md4Sum(data []byte) []byte {
	// You need to import a third-party package for MD4 hashing
	// Install it with: go get golang.org/x/crypto/md4
	h := md4.New()
	h.Write(data)
	return h.Sum(nil)
}

// Encode data to UTF-16LE (for NTLM hashing)
func utf16leEncode(data string) []byte {
	var utf16le []byte
	for _, r := range data {
		utf16le = append(utf16le, byte(r), byte(r>>8))
	}
	return utf16le
}

// Bcrypt hashing
// Bcrypt is used for securely storing passwords.
func HashBcrypt(data string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
	return string(hash), err
}

// Base16 encoding (Hexadecimal)
func EncodeBase16(data string) string {
	return hex.EncodeToString([]byte(data))
}

// Base64 encoding
func EncodeBase64(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

// Base64 decoding
func DecodeBase64(data string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}

// URL encoding
func URLEncode(data string) string {
	return url.QueryEscape(data)
}

// URL decoding
func URLDecode(data string) (string, error) {
	decoded, err := url.QueryUnescape(data)
	if err != nil {
		return "", err
	}
	return decoded, nil
}

// Utility to handle multiple hashes in one function
func Hash(data, method string) (string, error) {
	method = strings.ToLower(method)
	switch method {
	case "md5":
		return HashMD5(data), nil
	case "sha1":
		return HashSHA1(data), nil
	case "sha256":
		return HashSHA256(data), nil
	case "sha512":
		return HashSHA512(data), nil
	case "bcrypt":
		return HashBcrypt(data)
	case "ntlm":
		return HashNTLM(data), nil
	default:
		return "", fmt.Errorf("unsupported hashing method: %s", method)
	}
}

func showsAllsHashingEncoding() {
	data := "example"
	fmt.Println("MD5:", HashMD5(data))
	fmt.Println("SHA256:", HashSHA256(data))
	fmt.Println("NTLM:", HashNTLM(data))

	bcryptHash, _ := HashBcrypt(data)
	fmt.Println("Bcrypt:", bcryptHash)

	base64Encoded := EncodeBase64(data)
	fmt.Println("Base64 Encoded:", base64Encoded)
	base64Decoded, _ := DecodeBase64(base64Encoded)
	fmt.Println("Base64 Decoded:", base64Decoded)

	urlEncoded := URLEncode(data)
	fmt.Println("URL Encoded:", urlEncoded)
	urlDecoded, _ := URLDecode(urlEncoded)
	fmt.Println("URL Decoded:", urlDecoded)
}
