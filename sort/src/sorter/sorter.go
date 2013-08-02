package main

import (
	"flag"
	"fmt"
	"bufio"
	"io"
	"os"
	"strconv"
	"time"
)

import (
	"algorithms/bubblesort"
	"algorithms/qsort"
)

var infile *string = flag.String("i", "infile", "File contains values for sort")
var outfile *string = flag.String("o", "outfile", "File contains sorted data")
var algorithms *string = flag.String("a", "qsort", "Sort algorithms")

func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Failed create the out file: ", outfile)
		return err
	}
	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}

	return nil
}

func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Failed open the input file: ", infile)
		return
	}
	defer file.Close()

	br := bufio.NewReader(file)
	values = make([]int, 0)

	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}
		if isPrefix {
			fmt.Println("A tool long line, seems unexpected.")
			return
		}
		str := string(line)// make the array to string
		value, err1 := strconv.Atoi(str)//make the str to int type
		if err1 != nil {
			err = err1
			return
		}
		values = append(values, value)
	}

	return
}



func main() {
	flag.Parse()

	if infile != nil {
		fmt.Println("infile:", *infile, " outfile:", *outfile,
			" algorithm:", *algorithms)
	}
	values, err := readValues(*infile)
	if err == nil {
		t1 := time.Now()
		switch *algorithms {
		case "qsort":
			qsort.QuikSort(values)
		case "bubblesort":
			bubblesort.BubbleSort(values)
		default:
			fmt.Println("Sort alogrithms", *algorithms,
			    "is either unknow or unsupported.")
		}
		t2 := time.Now()

		fmt.Println("The sort process cost", t2.Sub(t1),"to complete.")
		writeValues(values, *outfile)
	} else {
		fmt.Println(err)
	}
}
