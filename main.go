package main

import (
	"path/filepath"
	"log"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {


	results, err := filepath.Glob("ladygrace_emails/*.eml") //returns the .eml files from the folder
	if err != nil{
		log.Fatalln(err)
	}

	for i := 0; i < len(results) ; i ++ {
		file, err := ioutil.ReadFile(results[i]) //reads all the files in the folder
		if err != nil {
			log.Fatal(err)
		}

		s := string(file)

		//fmt.Println(s) //prints out all the files
		emailLine := strings.Split(s,"\n") //split by newline

		for j := 0; j < len(emailLine); j++ {
			if strings.Contains(emailLine[j], "Email: "){
				fmt.Println(emailLine[j])
			}
		}

	}

}


