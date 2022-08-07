package main

import (
	"fmt"
	"os"
	"testing"
)

// Examplesという機能。標準出力をそのままテストできる
func ExampleHi() {
	fmt.Println("Hi!")
	// Output: Hi!
}

// Go 1.7からUnordered outputがサポート
func ExampleUnorderedHi() {
	// mapをイテレートして出力する処理。結果は順不同

	// 期待する出力が1,2,3の順でなくても通る
	// Unordered output:
	// 1
	// 2
	// 3
}

//   > go test -v // 関数名がTestでなくても認識される

// Benchmarkで始まる関数はベンチマーク用
func BenchmarkXX(b *testing.B) {
	// 中でベンチマーク実行用の関数を呼ぶ。引数にbを渡す。
	//benchHelper(b, {回数とか}, {対象の関数とか})
	// 実行用の関数の中で b.ReportAllocs() を呼ぶと、ループ回数や所要時間、
	// アロケートされたバイト数や回数を出力してくれる。

	// なおGo 1.17からこの関数内に複数ベンチマークを書けるサブベンチマークが加わった。
}

//  > go test -bench . // コマンド実行時

func TestMain(m *testing.M) {
	setup()             // DBのインサートとか前処理が書ける
	exitCode := m.Run() // 同じファイル内のTest～メソッドを全部実行
	shutdown()          // DBのクリア処理とか後処理
	os.Exit(exitCode)
}

// +buid integ

func TestForIntegration(t *testing.T) {
	// ...
}

//   > go test // 実行されない
//   > go test -tags=integ // 実行される。結合テストのみとか
