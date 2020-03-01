package main

import (
	TokenDaos "./lib"
	"fmt"
)

func insertNewToken() string {
	return TokenDaos.CreateToken()
}

func checkTokenExists(token string) bool {
	return TokenDaos.TokenExists(token)
}

func checkPathCount(path string) int {
	return TokenDaos.PathCount(path)
}

func main() {
	fmt.Println(insertNewToken())

	if checkTokenExists("1234567") {
		fmt.Println("count: ", checkPathCount("abc"))
	} else {
		fmt.Println("Unathorized")
	}

	if checkTokenExists("1") {
		fmt.Println("count: ", checkPathCount("abc"))
	} else {
		fmt.Println("Unathorized")
	}

	if checkTokenExists("1234567") {
		fmt.Println("count: ", checkPathCount("abcd"))
	} else {
		fmt.Println("Unathorized")
	}

}
