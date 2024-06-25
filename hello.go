package main

import (
	"fmt"
	"log"
)

import "rsc.io/quote"

func main() {
	// 设置预定义Logger的属性，包括
	// 日志条目前缀和禁用打印的标志
	//  时间、源文件和行号.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	fmt.Println(quote.Go())
	names := []string{"Gladys", "Samantha", "Darrin"}
	messages, err := Hellos(names)

	// message, err := Hello("justin")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
