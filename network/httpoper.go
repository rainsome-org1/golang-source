package main

import (
	"net/http"
	"os"
	"fmt"
	"io"
)

func main() {

	resp, err := http.Get("http://localhost")
	if err != nil {
		fmt.Println("Error...")
		return
	}

	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)

<<<<<<< HEAD
	resp2, err := http.Head("http://www.google.com")
=======
	resp2, err := http.Head("http://localhost")
>>>>>>> 85275d63728b8d7e02531eb0686f4237a8db2c97
	if err != nil {
		fmt.Println("Error...")
		return
	}

<<<<<<< HEAD
	fmt.Printf("\n\n\n\n\n\n\n")
=======
>>>>>>> 85275d63728b8d7e02531eb0686f4237a8db2c97
	defer resp2.Body.Close()
	fmt.Println(resp2.Header)
}
