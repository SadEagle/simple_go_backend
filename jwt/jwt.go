package jwt

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

const JWTExpirationTime = time.Hour * 24 * 7
const JWTSecretKeyLen = 32

var JWTSecretKey = func() []byte {
	var key = make([]byte, JWTSecretKeyLen)
	rand.Read(key)
	return key
}()

var jwtHeader = func() []byte {
	var headerInfo = map[string]string{
		"alg": "HS256",
		"typ": "JWT",
	}
	var headerBytes, err = json.Marshal(headerInfo)
	if err != nil {
		panic("Can't marshal jwt header data")
	}
	var headerBytesBase64 = make([]byte, base64.URLEncoding.EncodedLen(len(headerBytes)))
	base64.URLEncoding.Encode(headerBytesBase64, headerBytes)
	return headerBytesBase64
}()

type UserJWTPayloadRequest struct {
	UserID pgtype.UUID `json:"user_id"`
}

type UserJWTPayloadResponse struct {
	UserID pgtype.UUID `json:"user_id"`
	Exp    time.Time   `json:"exp"`
}

func getJWTPayload(payloadReq UserJWTPayloadRequest) ([]byte, error) {
	payloadResp := UserJWTPayloadResponse{
		UserID: payloadReq.UserID,
		Exp:    time.Now().UTC().Add(JWTExpirationTime),
	}
	var payloadBytes, err = json.Marshal(payloadResp)
	if err != nil {
		return nil, err
	}
	var payloadBU64 = make([]byte, base64.URLEncoding.EncodedLen(len(payloadBytes)))
	base64.URLEncoding.Encode(payloadBU64, payloadBytes)
	return payloadBU64, nil
}

func CreateToken(payload UserJWTPayloadRequest) (string, error) {
	jwtPayload, err := getJWTPayload(payload)
	if err != nil {
		return "", err
	}
	message := fmt.Sprintf("%s.%s", jwtHeader, jwtPayload)
	hash := hmac.New(sha256.New, JWTSecretKey)
	hash.Write([]byte(message))
	signature := hash.Sum(nil)
	token := fmt.Sprintf("%s.%s", message, signature)
	return token, nil
}

func VerifyToken(w any) {
}
