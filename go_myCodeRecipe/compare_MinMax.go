package main

import "fmt"

// C_maxBy: 比較関数を使って大きい方を選ぶ
func C_maxBy(a, b int, compare func(int, int) bool) int {
	if compare(a, b) {
		return a
	}
	return b
}

// C_minBy: 比較関数を使って小さい方を選ぶ
func C_minBy(a, b int, compare func(int, int) bool) int {
	if compare(a, b) {
		return b
	}
	return a
}

// C_min_0: 引数aには式を渡し、計算結果がマイナスの時は0を返す
func C_min_0(a int) int {
	if 0 > a {
		return 0
	}
	return a
}

// C_min_n: 引数aには式を渡し、計算結果がn未満の時はnを返す
func C_min_n(n, a int) int {
	if n > a {
		return 0
	}
	return n
}

func main() {
	// 通常の比較
	fmt.Println(C_maxBy(10, 20, func(a, b int) bool { return a > b })) // 出力: 20
	fmt.Println(C_minBy(10, 20, func(a, b int) bool { return a > b })) // 出力: 10

	// カスタム比較: 偶数を優先
	fmt.Println(C_maxBy(3, 4, func(a, b int) bool { return a%2 == 0 })) // 出力: 4
}
