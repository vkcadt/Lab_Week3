package main


import (
	"bufio"
	"fmt"
	"os"
	"main.go/utils/crack"
)

func main(){

	hashTarget :=""

	//fmt.Print("This Function Take The Password From nord-vpn.txt")
	fmt.Print("Enter Yout Target Hashing : ")
	fmt.Scan(&hashTarget)

	file, err := os.Open("nord_vpn.txt")
	if err != nil{
		fmt.Print("This Open File Error!\n",err)
		return 
	 }
	defer file.Close()
	fmt.Print("+this file open successfully\n")

	logFile, err := os.Create("logFile.txt")

	if err !=nil{
		fmt.Print("Create File Error!",err)
		return
	} 
	defer logFile.Close()
	fmt.Print("+create file successfully\n")
	
	scanner := bufio.NewScanner(file)

	found := false
	attempts := 0

	fmt.Print("\nString SHA1 password search...\n")

	for scanner.Scan(){

		text := scanner.Text()
		//fmt.Print("Output:\n",text)

		guessHash := crack.CrackSha1(text)

		attempts ++

		if guessHash == hashTarget{
			fmt.Print("+Password Match With:",text)
			fmt.Fprint(logFile,"trying:",text,"-> hash:", guessHash,"-> match:\n")
			found =true
			break
		}else{
			fmt.Fprint(logFile,"trying:",text,"-> hash:", guessHash,"-> not match:\n")
		}

		if err:= scanner.Err(); err !=nil {
			fmt.Print("Error reading wordlist:",err)
		}		
	}	

	if !found {
		fmt.Println("\n+Password not found in wordlist.\n")
		fmt.Fprintln(logFile, "Password not found in wordlist.")
	}

	fmt.Print("\n+Total attempts: ", attempts)
	fmt.Print(" Search completd.")
	

}