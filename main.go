package main

import (
	"fmt"
	"log"
	"os"
)

import "./wordcount"

func main() {

	dirpath := os.Args[1]

	filenames, err := wordcount.OSReadDir(dirpath)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(filenames)

	combined := wordcount.MapReducefunc(filenames)
	fmt.Println("Final count---------------------", combined)

	//fmt.Println(string(bytes))

}
