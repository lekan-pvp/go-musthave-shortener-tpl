package cookie_handler

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/go-musthave-shortener-tpl/internal/generate_random"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var key = []byte("secret key")

func CreateCookie() *http.Cookie {
	gen, err := generate_random.GenerateRandom(16)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	data, err := hex.DecodeString(fmt.Sprintf("%x", gen))
	if err != nil {
		log.Fatal(err)
		return nil
	}

	uid := binary.BigEndian.Uint32(data[:4])
	log.Printf("%d\n", uid)

	id:= strconv.Itoa(int(uid))

	h := hmac.New(sha256.New, key)
	h.Write([]byte(id))
	dst := h.Sum(nil)

	cookie := &http.Cookie{
		Name: "uid",
		Value: fmt.Sprintf("%s:%x", id, dst),
		Path: "/",
	}
	log.Printf("%s, %x", id, dst)
	return cookie
}


func CheckCookie(cookie *http.Cookie) bool {
	values := strings.Split(cookie.Value, ":")
	data, err := hex.DecodeString(values[1])
	if err != nil {
		log.Fatal(err)
		return false
	}
	//uuid := binary.BigEndian.Uint32(data[:4])
	//id := values[0]
	h := hmac.New(sha256.New, key)
	h.Write(data[:4])
	sign := h.Sum(nil)

	if hmac.Equal(sign, data[:4]) {
		return true
	} else {
		return false
	}
}
