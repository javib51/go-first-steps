/*
*go run args.go arg1 arg2 arg3
*/
package main

//import ("fmt";"os") //alternative

import (
	"fmt"
	//"os"
)
/*
func main() {
	
	//var v1,....,vn type
	var s, sep string 
	//No parentheses but with brackets
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)
}
*/

func main() {
	
	i:=0 
	//No parentheses but with brackets
	for true {
		if i >= 10 {
			break;
		}
		fmt.Println(i)		
		i++
	}


}
