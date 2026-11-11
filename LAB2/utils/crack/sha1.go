package crack

import (
	"encoding/hex"
	"crypto/sha1"
	
)

func CrackSha1(text string)string{
	hashGuess:= sha1.Sum([]byte(text))
	hashGuessString:=hex.EncodeToString(hashGuess[:])
	return hashGuessString

}