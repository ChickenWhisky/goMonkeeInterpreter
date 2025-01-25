package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/ChickenWhisky/monkeyInterpreter/repl"
)

func main(){
	user,err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s !! Welcome to Monkeeeee Programing Language\n",user.Username)
	fmt.Printf("Type in commands\n")
	repl.Start(os.Stdin,os.Stdout)
}