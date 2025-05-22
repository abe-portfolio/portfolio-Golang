package main

import (
	"fmt"
	"runtime"
)

func main() {
	yourOS := runtime.GOOS

	switch yourOS {
	case "windows":
		fmt.Println("現在実行中のOSは Windows です。")
	case "linux":
		fmt.Println("現在実行中のOSは Linux です。")
	case "darwin":
		fmt.Println("現在実行中のOSは macOS です。")
	default:
		fmt.Printf("現在実行中のOSは %s です。\n", yourOS)
	}
}
