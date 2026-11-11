package crack

import (
	"crypto/md5"
	"encoding/hex"
)

func CrackMd5(text string) string{

	// text = 
	// hashTarget
	// fmt.Print("this is fucntion for cracking")
	// fmt.Print("Enter The Target Hash!")
	// fmt.Scan(&hashTarget)
	// fmt.Print("Enter Your Text!")
	// fmt.Scan(&text)


	hashGuess := md5.Sum([]byte(text))
	hashGuessString := hex.EncodeToString(hashGuess[:])

	//fmt.Println("Output Guess Hass :",hashGuessString)

	// if (hashGuessString==hashTarget){
	// 	fmt.Print("This Guessing Is Match !")
	// }else{
	// 	fmt.Print("This Guessing Is Not Match !")
	// }

	return (hashGuessString);
}