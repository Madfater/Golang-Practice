package main

import (
	"fmt"
	"time"
)

// 定義一個函數，該函數作為一個協程運行
func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Second) // 模擬耗時操作
	}
}

func printStrings() {
	var arr = [5]string{"a", "b", "c", "d", "e"}
	for _, v := range arr {
		fmt.Println(v)
		time.Sleep(2*time.Second)
	}
}

func main() {
	// 啟動一個新的協程，並在其中運行 printNumbers 函數
	go printNumbers()
	go printStrings()

	// 主協程繼續執行
	fmt.Println("Main goroutine continues...")

	// 主協程等待一段時間，以確保 printNumbers 協程有足夠的時間運行
	time.Sleep(2 * time.Second)

	// 主協程完成
	fmt.Println("Main goroutine exits.")
}
