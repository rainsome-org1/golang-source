package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"./mlib"
	"./mp"
)

var lib *library.MusicManager
var id int = 1
var ctr1, signal chan int

func handleLibCommands(token []string) {
	switch token[1] {
	case "list":
		for i := 0; i < lib.Len(); i++ {
			e, _ := lib.Get(i)
			fmt.Println(i+1, ":", e.Name, e.Artist,
			    e.source, e.Type)
		    }
	case "add":
		if len(token) == 6 {
			id++
			lib.Add(&mlib.MusicEntry{strconv.Itoa(id),
				token[2], token[3], token[4], token[5]})
		} else {
			fmt.Println("Usage: lib add <name><artist><source><type>")
		}
	case "remove":
		if len(token) == 3 {
			lib.RemoveByName(token[2])
		} else {
			fmt.Println("Usage: lib remove <name>")
		}
	default:
		fmt.Println("Unrecognized lib command: ", token[1])
	}
}

func handlePlayCommand(token []string) {
	if len(token) != 2 {
		fmt.Println("Usage: play <name>")
		return
	}

	e := lib.Find(token[1])
	if e == nil {
		fmt.Println("The music", token[1], "dosen't exisit")
		return
	}
	mp.Play(e.Source, e.Type, ctrl, signal)
}

func main() {
	fmt.Println(`
	    Enter following commands to control the player:
	    lib list -- view the existing music lib
	    lib add <name><artist><source><type> -- add a record to lib
	    lib remove <name> -- remove the specified music from the lib
	    play <name> -- play the specified music
	    `)
	lib = mlib.NewMusicMnanager()

	r := bufio.NewReader(os.Stin)

	for {
		fmt.Print("Enter command--> ")
		rawLine, _, _ := r.ReadLine()
		line := string(rawLine)

		if line == "q" || line == "e" {
			break
		}

		token := strings.Split(line, " ")
		if token[0] == "lib" {
			handleLibCommands(token)
		} else if token[0] == "play" {
			handleLibCommand(token)
		} else {
			fmt.Println("Uncognized command:", token[0])
		}
	}
}

