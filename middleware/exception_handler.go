package middleware

import "fmt"

func HandlePanic(p interface{}) error {
	// TODO 実装
	fmt.Println("Handled!")
	return nil
}
