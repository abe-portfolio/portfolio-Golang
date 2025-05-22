// QRコード読み込みモジュール
package M_Reader

import (
	"fmt"
	"runtime"
)

func R_index() {
	fmt.Println("This is Reader") // test code
}

// QRコードから内容を取得
func Read_QR() {

}

// 保存した画像から読み込むケース
func Open_dialog() {

}

// 内蔵カメラの起動
func Launch_the_camera() {
	// プログラム起動元のOSを取得
	currentOS := runtime.GOOS

	switch currentOS {
	case "windows":
		fmt.Println("現在実行中のOSは Windows です。")
	case "linux":
		fmt.Println("現在実行中のOSは Linux です。")
	case "darwin":
		fmt.Println("現在実行中のOSは macOS です。")
	default:
		fmt.Println("お使いのOSではこの機能は使用できません。")
		// どこに戻すか？
	}
}

// 読み込んだ結果を保存？ or 保存先を指定？
func Save_data() {

}
