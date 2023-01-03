package main

import (
	"fmt"
	"log"
	"os"
)

type myWrappedLogger struct {
	*log.Logger
	level string
}

func (ml *myWrappedLogger) Println(level string, v ...any) {
	fmt.Println("Fancy Logger")
	fmt.Println("------------")
	fmt.Printf("Level: %s\n", level)
	ml.Output(2, fmt.Sprintln(v...))
	fmt.Println("------------")
}

func main() {
	mwl := myWrappedLogger{log.New(os.Stdout, "LogPrefix: ", log.LstdFlags), ""}
	mwl.Println("error", "this is the log output")
}