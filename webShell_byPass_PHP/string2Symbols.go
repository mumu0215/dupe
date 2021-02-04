package main

import (
	"flag"
	"fmt"
	"src/common"
	"strings"
)

var(
	inputString=flag.String("s","","input php string to be transformed to symbols")
)

func init() {
	flag.Parse()
}
func disPlay(a []string,b []string){
	fmt.Println("Result:")
	for i:=0;i<len(a);i++{
		fmt.Println(a[i]+"\t->  "+b[i])
	}
}
func main() {
	if *inputString==""{
		output:=common.TransForm(common.FuncWord)
		disPlay(common.FuncWord,output)
	}else {
		inputSlice:=strings.Split(*inputString," ")
		output:=common.TransForm(inputSlice)
		disPlay(inputSlice,output)
	}
}