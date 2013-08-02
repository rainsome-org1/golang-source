package main

import (
	"crypto/sha1"
	"crypto/md5"
	"fmt"
)

func main() {
	TestString := "youngcy"

	Md5Inst := md5.New()
	Md5Inst.Write([]byte(TestString))
	fmt.Println("md5 result: ", Md5Inst.Sum([]byte("")))

	sha1Inst := sha1.New()
	sha1Inst.Write([]byte(TestString))
	fmt.Println("sha1 result: ", sha1Inst.Sum([]byte("")))

}
