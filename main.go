package main

import (
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const Otp = "0123456789"

func HashToken(Token string) (string, error) {
	data, err := bcrypt.GenerateFromPassword([]byte(Token), 14)
	return string(data), err

}

func CheckHashToken(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil

}

func main() {
	token, _ := GenerateOTP(6)
	hashToken, _ := HashToken(token)
	fmt.Println(token)
	fmt.Println(hashToken)
	match := CheckHashToken(token, hashToken)
	fmt.Println("match:", match)
}

func GenerateOTP(length int) (string, error) {
	buffer := make([]byte, length)
	res, err := rand.Reader.Read(
		buffer,
	)
	if err != nil {
		return "", err
	}
	fmt.Println(res)

	otpLength := len(Otp)
	for i := 0; i < length; i++ {
		buffer[i] = Otp[int(buffer[i])%otpLength]
	}

	return string(buffer), nil
}
