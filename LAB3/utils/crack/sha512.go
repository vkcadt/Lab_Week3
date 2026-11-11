package crack

import (
	"encoding/hex"
	"crypto/sha512"
	
)

func CrackSha512(text string)string{
	hashGuess:= sha512.Sum512([]byte(text))
	hashGuessString:=hex.EncodeToString(hashGuess[:])
	return hashGuessString

}