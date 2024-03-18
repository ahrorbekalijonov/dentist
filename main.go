package main

import (
	"fmt"
	"time"
)

func main(){
	vaqt := time.Now()
	newtime := vaqt.AddDate(0, 0, 0)
	fmt.Println(vaqt)
	fmt.Println(newtime)
}