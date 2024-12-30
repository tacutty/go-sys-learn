package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("start func()")

	// タイムアウト3秒のコンテキストを作成
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel() // 必ずキャンセルしてリソースを解放

	// ゴルーチンで処理を実行
	go func() {
		fmt.Println("start goroutine()")
		time.Sleep(2 * time.Second) // 2秒間の処理
		fmt.Println("end goroutine()")
	}()

	// タイムアウトまたはキャンセルを待つ
	<-ctx.Done()

	// タイムアウトまたはキャンセルの理由を出力
	fmt.Println("end func()", ctx.Err()) // タイムアウトかキャンセルか
}