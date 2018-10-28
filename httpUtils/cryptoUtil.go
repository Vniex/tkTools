package httpUtils

import (
	"crypto/sha256"
	"fmt"
)

func GetSHA256(str string)string{
	h := sha256.New()
	h.Write([]byte(str))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x",bs)
}