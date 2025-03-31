package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"learn.zone01dakar.sn/forum-rest-api/app/lib"
)

const SECRET_KEY = "0ada51d3-b3d9-4b0b-ae3d-cdb0ff79e453"

// The above code defines an interface named Payload that has a method ToJSON() which returns a byte
// slice and an error.
// @property ToJSON - ToJSON is a method that returns the JSON representation of the object
// implementing the Payload interface. It returns a byte slice containing the JSON data and an error if
// there was any issue during the conversion.
type Payload interface {
	ToJSON() ([]byte, error)
}

// The above code defines a struct type called "Header" with two string fields, "Alg" and "Typ", which
// are tagged for JSON serialization.
// @property {string} Alg - The "Alg" property in the Header struct represents the algorithm used for
// signing the JSON Web Token (JWT). It specifies the cryptographic algorithm that is used to secure
// the token.
// @property {string} Typ - The "Typ" property in the Header struct is a string that represents the
// type of the token. It is typically set to "JWT" (JSON Web Token) to indicate that the token is a
// JWT.
type Header struct {
	Alg  string `json:"alg"`
	Typ  string `json:"typ"`
	Addr string `json:"addr"`
}

type JWT struct {
	Header  Header
	Payload Payload
}

// The function Base64Encode encodes a byte slice into a base64 string and removes any trailing equal signs.
func (j JWT) Base64Encode(src []byte) string {
	return strings.TrimRight(base64.URLEncoding.EncodeToString(src), "=")
}

// The Sign function takes a data string and a secret key, and returns a base64-encoded
// HMAC-SHA256 signature of the data using the secret key.
// This signature is used to generate the third part of the JWT token by hashing the concatenated
// header and payload, along with the secret key, using the HMAC-SHA256 algorithm.

func (j JWT) Sign(data string) (string, error) {

	h := hmac.New(sha256.New, []byte(SECRET_KEY))
	_, err := h.Write([]byte(data))
	if err != nil {
		return "", err
	}

	return j.Base64Encode(h.Sum(nil)), nil
}

// The GenerateJWT function takes in a header, payload, and secret, encodes them into a JSON Web Token
// (JWT), signs the token using the secret, and returns the complete JWT.
func (j JWT) Generate() (string, error) {
	headerEncoded, err := json.Marshal(j.Header)
	if err != nil {
		return "", err
	}

	payloadJSON, err := json.Marshal(j.Payload)
	if err != nil {
		return "", err
	}

	encodedHeader := j.Base64Encode(headerEncoded)
	encodedPayload := j.Base64Encode(payloadJSON)
	unsignedToken := encodedHeader + "." + encodedPayload

	signature, err := j.Sign(unsignedToken)
	if err != nil {
		return "", err
	}

	return unsignedToken + "." + signature, nil
}

func (j JWT) Parse(token string, payload *lib.Payload) error {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return errors.New("invalid token format")
	}

	encodedPayload := parts[1]

	// Decode payload
	payloadBytes, err := base64.RawURLEncoding.DecodeString(encodedPayload)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(payloadBytes, payload); err != nil {
		return err
	}

	return nil
}

// VerifyJWT function verifies the validity of the JWT by checking the signature.
func (j JWT) Valid(r *http.Request, token, secret string) bool {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return false
	}

	encodedHeader := parts[0]
	encodedPayload := parts[1]
	signature := parts[2]

	// Decode header
	headerBytes, err := base64.RawURLEncoding.DecodeString(encodedHeader)
	if err != nil {
		return false
	}

	header := Header{}
	err = json.Unmarshal(headerBytes, &header)
	if err != nil {
		return false
	}

	// Calculate the signature
	unsignedToken := encodedHeader + "." + encodedPayload
	expectedSignature, err := j.Sign(unsignedToken)
	if err != nil {
		return false
	}

	// Compare signatures
	return signature == expectedSignature
}

func (j JWT) CheckingToken(r *http.Request, respone *lib.Response, token string) (lib.Payload, error) {

	payload := lib.Payload{}

	if !j.Valid(r, token, SECRET_KEY) {
		lib.ErrorWriter(respone, "You are not allowed to access to this ressource.Please login and try again.", http.StatusForbidden)
		return payload, errors.New("you are not allowed to access to this ressource.Please login and try again")
	}

	if err := j.Parse(token, &payload); err != nil {
		lib.ErrorWriter(respone, "Error getting your informations!.", http.StatusInternalServerError)
		return payload, errors.New("error getting your informations")
	}

	if payload.ExpirationDate.Before(time.Now()) {
		lib.ErrorWriter(respone, "Your token has been expired.", http.StatusForbidden)
		return payload, errors.New(respone.Message)
	}
	return payload, nil
}
