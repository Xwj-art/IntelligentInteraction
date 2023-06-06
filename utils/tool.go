package utils

import "fmt"

// 处理可能会发生的错误

func HandleErr(err error, msg string) {
	if err != nil {
		fmt.Printf("\033[33m======================================\033[0m\n")
		fmt.Printf("\033[34m=============%s=============\033[0m\n", msg)
		fmt.Printf("\033[33m======================================\033[0m\n")
		fmt.Println(err.Error())
		fmt.Println(err)
		panic(err)
	}
}
