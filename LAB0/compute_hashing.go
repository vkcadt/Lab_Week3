package main

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"crypto/sha1"
	"golang.org/x/crypto/sha3"
	"encoding/hex"
)

func main (){

	var A string
	var B string

	fmt.Print("Enter You Two String!\n")
	fmt.Print("String A : ")
	fmt.Scan(&A)
	fmt.Print("String B : ")
	fmt.Scan(&B)

	computeAsMd5(A , B)
	computeAsSha1(A, B)
	computeAsSha256(A,B)
	computeAsSha512(A,B)
	computeAsSha3(A, B)

	// hash1:= MD5.sum([]byte(A))
	// hash2:= MD5.sum([]byte(B))

	// hash1String:= hex.EncodeToString(hash1[:])
	// hash2String:= hex.EncodeToString(hash2[:])


	// fmt.Println("hash1(MD5) =\n ",hash1String)
	// fmt.Println("hash2(MD5) =\n ",hash2String)
}

func computeAsMd5(A string, B string){

	hash1:= md5.Sum([]byte(A))
	hash2:= md5.Sum([]byte(B))

	hash1String:= hex.EncodeToString(hash1[:])
	hash2String:= hex.EncodeToString(hash2[:])
	print("+MD5\n")

	fmt.Println("Output A = ",hash1String)
	fmt.Println("Output B = ",hash2String)
	
	if (hash1String==hash2String){
		fmt.Println("It match")
	}else{
		fmt.Println("It not match")
	}
}

func computeAsSha1(A string, B string){
	hash1:= sha1.Sum([]byte(A))
	hash2:= sha1.Sum([]byte(B))

	hash1String:= hex.EncodeToString(hash1[:])
	hash2String:= hex.EncodeToString(hash2[:])
	print("+SHA1\n")

	fmt.Println("Output A =",hash1String)
	fmt.Println("Output B =",hash2String)

	if (hash1String==hash2String){
		fmt.Println("It match")
	}else{
		fmt.Println("It not match")
	}
}


func computeAsSha256(A string, B string){
	hash1:= sha256.Sum256([]byte(A))
	hash2:= sha256.Sum256([]byte(B))

	hash1String:= hex.EncodeToString(hash1[:])
	hash2String:= hex.EncodeToString(hash2[:])
	print("+SHA256\n")

	fmt.Println("Output A =",hash1String)
	fmt.Println("Output B =",hash2String)

	if (hash1String==hash2String){
		fmt.Println("It match")
	}else{
		fmt.Println("It not match")
	}
}

func computeAsSha512(A string, B string){
	hash1:= sha512.Sum512([]byte(A))
	hash2:= sha512.Sum512([]byte(B))

	hash1String:= hex.EncodeToString(hash1[:])
	hash2String:= hex.EncodeToString(hash2[:])
	print("+SHA512\n")

	fmt.Println("Output A =",hash1String)
	fmt.Println("Output B =",hash2String)

	if (hash1String==hash2String){
		fmt.Println("It match")
	}else{
		fmt.Println("It not match")
	}
}

func computeAsSha3(A string, B string){
	hash1:= sha3.Sum224([]byte(A))
	hash2:= sha3.Sum224([]byte(B))

	
	hash1String:= hex.EncodeToString(hash1[:])
	hash2String:= hex.EncodeToString(hash2[:])
	print("+SHA1\n")

	fmt.Println("Output A =",hash1String)
	fmt.Println("Output B =",hash2String)

	if (hash1String==hash2String){
		fmt.Println("It match")
	}else{
		fmt.Println("It not match")
	}
}
