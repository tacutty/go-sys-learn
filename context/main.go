package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 親コンテキストを作成（リクエスト全体のタイムアウトを設定）
	parentCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // 最後に親コンテキストをキャンセル

	// 子コンテキストを使ってDBアクセスの処理を実行
	go func(ctx context.Context) {
		childCtx, childCancel := context.WithTimeout(ctx, 2*time.Second) // 子コンテキスト: DBアクセス用
		defer childCancel()

		if err := queryDatabase(childCtx); err != nil {
			fmt.Println("DBクエリ失敗:", err)
		} else {
			fmt.Println("DBクエリ成功")
		}
	}(parentCtx)

	// 子コンテキストを使って外部API呼び出しの処理を実行
	go func(ctx context.Context) {
		childCtx, childCancel := context.WithTimeout(ctx, 3*time.Second) // 子コンテキスト: 外部API用
		defer childCancel()

		if err := callExternalAPI(childCtx); err != nil {
			fmt.Println("外部API失敗:", err)
		} else {
			fmt.Println("外部API成功")
		}
	}(parentCtx)

	// 親コンテキストの終了を待機
	<-parentCtx.Done()
	fmt.Println("リクエスト全体が終了:", parentCtx.Err())
}

// 擬似的なDBクエリ処理
func queryDatabase(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err() // タイムアウトまたはキャンセルを検知
	case <-time.After(1 * time.Second): // 擬似的なDBクエリ（1秒かかる）
		return nil
	}
}

// 擬似的な外部API呼び出し処理
func callExternalAPI(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err() // タイムアウトまたはキャンセルを検知
	case <-time.After(2 * time.Second): // 擬似的な外部API（4秒かかる）
		return nil
	}
}