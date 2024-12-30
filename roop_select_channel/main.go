package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	quit := make(chan bool)

	go func() {
		time.Sleep(1 * time.Second)
		ch <- "データ送信1"
		time.Sleep(1 * time.Second)
		ch <- "データ送信2"
		time.Sleep(1 * time.Second)
		close(ch) // チャネルを閉じる
	}()

	go func() {
		time.Sleep(3 * time.Second)
		quit <- true
	}()

LOOP: // 外側のループにラベルを付ける
	for {
		select {
		case v, ok := <-ch:
			if ok {
				fmt.Println(v)
			} else {
				fmt.Println("ch が閉じられました")
				break LOOP // ループ全体を終了
			}
		case <-quit:
			fmt.Println("終了信号を受信しました")
			break LOOP // 外側のループを終了
		default:
			fmt.Println("データ受信待ち")
			time.Sleep(500 * time.Millisecond) // 待機
		}
	}
}
