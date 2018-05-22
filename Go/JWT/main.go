package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

type header struct {
	Algorithm string `json:"alg"`
	Type      string `json:"typ"`
}

type payload struct {
	Issuer         string `json:"iss"`
	Subject        string `json:"sub"`
	ExpirationTime int    `json:"exp"`
}

var payloadDecrypted payload

func base64Encode(b []byte) string {
	return strings.TrimRight(base64.URLEncoding.EncodeToString(b), "=")
}

func base64Decode(s string) ([]byte, error) {
	//corrige a remoção dos `=`s
	switch len(s) % 4 {
	case 2:
		s += "=="
	case 3:
		s += "="
	}
	return base64.URLEncoding.DecodeString(s)
}

func signature(message, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	return base64Encode(h.Sum(nil))
}

func verify(token, secret string) bool {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return false
	}

	rawHeader, err := base64Decode(parts[0])
	if err != nil {
		return false
	}
	var hdr header
	if json.Unmarshal([]byte(rawHeader), &hdr) != nil {
		return false
	}

	if hdr.Type != "JWT" {
		return false
	}

	if hdr.Algorithm != "HS256" {
		return false
	}

	payloadHash, _ := base64Decode(parts[1])

	err = json.Unmarshal([]byte(payloadHash), &payloadDecrypted)
	if err != nil {
		return false
	}

	return signature(parts[0]+"."+parts[1], secret) == parts[2]
}

func main() {
	
	//header
	h := header{ "HS256", "JWT" }
	hj, _ := json.Marshal(h)
	hb64 := base64Encode(hj)
	
	//payload encryp
	p := payload{ "spr", "1", 1 }
	pj, _ := json.Marshal(p)
	pb64 := base64Encode(pj)
	
	//hash create
	hashToEncrypt := hb64 + "." + pb64
	secret := "secret"
	signature := signature(hashToEncrypt, secret)
	
	hash := hb64 + "." + pb64 + "." + signature

	fmt.Println(hash)

	//decode
	if verify(hash, secret) {
		fmt.Println("Autenticated")
		fmt.Println(payloadDecrypted.Issuer)
		fmt.Println(payloadDecrypted.ExpirationTime)
		fmt.Println(payloadDecrypted.Subject)
	}
}
