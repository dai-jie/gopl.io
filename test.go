package main

import "gopl.io/ch2/tempconv"
import "fmt"

func main(){
	k := tempconv.CToK(tempconv.BoilingC)
	fmt.Println(k)
}