package main

import "testing"

func someFunc(a int, b int) int {
	return a + b
}

// 引数は固定、関数名先頭がTestで認識される
func TestSum(t *testing.T) {
	if someFunc(1, 2) != 3 {
		t.Fatal("テスト失敗時のメッセージ")
	}
}

// > go test {フルパスのパッケージ名}
// > go test -v ./calc // modulesで認識されていれば相対パスOK -v が詳細
// // そのパッケージの*_test.go が対象で、ファイル名先頭が _, . なら無視
// > go test -v run {関数名} // 関数単位で実行
// > go test -v run "some$" // 関数名前方一致
// // Go1.10からテスト結果がキャッシュ、ファイルの内容と環境変数未変更なら
// // 使われる。明示的クリアが以下。
// > go clean -testcache
