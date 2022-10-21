package helper

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	UserId  int
	Email   string
	Expired time.Time
}

func GenerateToken(payload *Token) (string, error) {
	expHourStr := os.Getenv("JWT_EXPIRE")
	expHour, err := strconv.Atoi(expHourStr)
	if err != nil {
		return "", err
	}
	exp := time.Now().Add(time.Hour * time.Duration(expHour))

	claims := jwt.MapClaims{
		"payload": payload,
		"iat":     time.Now().Unix(),
		"exp":     exp,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tok, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tok, nil
}

func ValidateToken(tokString string) (*Token, error) {
	errResponse := fmt.Errorf("need signin")
	tok, err := jwt.Parse(tokString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok && !tok.Valid {
		return nil, errResponse
	}

	payload, ok := claims["payload"]
	if !ok && !tok.Valid {
		return nil, errResponse
	}

	bytePayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	var token Token
	err = json.Unmarshal(bytePayload, &token)
	if err != nil {
		return nil, err
	}
	return &token, nil
}
