/*
*Experiment
*Read file and create a new file with a \n between to words.
* knew bugs: split  words but with tabulation.
*/
package main

import (
	"os"
	"bufio"
	"strings"
	//"fmt"
)

func main () {

	args := os.Args[1:]
	
	if len(args) > 0 {	
		for _, arg := range args {

			f, err := os.Open(arg)
			if err != nil {
				continue
			}
			readFile(f, arg)
			f.Close()
	}
	}
}

func readFile (f *os.File, name string){
	
	newFile, err := os.Create(name+"_new.txt") //Create New file
	

	if err != nil { //Handle error
		panic(err)
	}

	write := bufio.NewWriter(newFile)

	read := bufio.NewScanner(f)
	
	for read.Scan(){

		words := strings.Split(read.Text(), " ")
		
		for _, w := range words {

			//for string(w[0])==" " {
			//	w= w[1:]
			//}
			
			if _, err := write.WriteString(w+"\n"); err !=nil {
				panic(err)
			}
		}
	}

	if err = write.Flush(); err != nil {
		panic(err)
	}

	newFile.Close()
}
