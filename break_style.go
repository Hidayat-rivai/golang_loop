package main

import "fmt"

func main(){
	var i = 0
	for {
		fmt.Println(i)
		i++
		
		if i == 10{
			break
		}
	}
}