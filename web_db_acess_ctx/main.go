package main

import (
	"context"
	"fmt"
	"time"
)

// このコードの動作

// 	1.	親コンテキスト:
// 	•	全体のリクエストに 5 秒のタイムアウトを設定。
// 	•	親コンテキストがキャンセルされると、すべての子コンテキストもキャンセルされます。
// 	2.	子コンテキスト:
// 	•	データベースクエリ用の子コンテキストには 2 秒のタイムアウトを設定。
// 	•	外部API呼び出し用の子コンテキストには 3 秒のタイムアウトを設定。
// 	3.	タスクの動作:
// 	•	データベースクエリは 1 秒で終了するため成功。
// 	•	外部API呼び出しは 4 秒かかるため、3 秒のタイムアウトで失敗。
// 	4.	リクエスト全体の終了:
// 	•	親コンテキストの 5 秒のタイムアウトが経過、もしくは子タスクの結果をすべて受け取った時点でリクエストが終了。

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
	case <-time.After(4 * time.Second): // 擬似的な外部API（4秒かかる）
		return nil
	}
}