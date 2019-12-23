package Err

import "fmt"

func Handle(err error, text string){
	if err != nil{
		fmt.Println(text)
	}
}
