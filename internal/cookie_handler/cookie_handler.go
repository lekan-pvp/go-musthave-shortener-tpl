package cookie_handler

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
	"strings"
)

var key = []byte("secret key")


func CreateCookie() *http.Cookie {

	id := uuid.NewString()

	log.Println("IN CreateCookie:", id)

	h := hmac.New(sha256.New, key)
	h.Write([]byte(id))
	dst := h.Sum(nil)

	cookie := &http.Cookie{
		Name: "token",
		Value: fmt.Sprintf("%s:%x", id, dst),
		Path: "/",
	}
	return cookie
}


func CheckCookie(cookie *http.Cookie) bool {
	values := strings.Split(cookie.Value, ":")

	data, err := hex.DecodeString(values[1])
	if err != nil {
		log.Fatal(err)
		return false
	}

	id := values[0]

	h := hmac.New(sha256.New, key)
	h.Write([]byte(id))
	sign := h.Sum(nil)

	if hmac.Equal(sign, []byte(data)) {
		return true
	} else {
		return false
	}
}

