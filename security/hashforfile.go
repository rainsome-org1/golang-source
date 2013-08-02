package main

import (
	"crypto/sha1"
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func main() {
	testFile := "src.hash"
	infile, ferr := os.Open(testFile)

	if ferr != nil {
		fmt.Fprintf(os.Stderr, "File open failed")
		os.Exit(1)
	}

	md5h := md5.New()
	io.Copy(md5h, infile)
	fmt.Printf("%x\n%s\n", md5h.Sum([]byte("")), testFile)

	sha1h := sha1.New()
	io.Copy(sha1h, infile)
	fmt.Printf("%x\n%s\n", sha1h.Sum([]byte("")), testFile)
}
