// QRコード生成モジュール
package M_Generator

import (
	"fmt"
	"image"
	"os"
	"path/filepath"
	"time"
)

func init() {
	// Get Project's Root Directory
	ProjectRoot, err1 := os.Getwd()
	if err1 != nil {
		return fmt.Errorf("faild to get project-root directory: %w", err1)
	}
}

// func G_index(projectRoot,) {}
func G_index() {
	// fmt.Println("This is Generator")
	// ime := Create_name()  test OK
	// fmt.Println(ime)      test OK

	/*
		本番用のハンドラー
		code := Generate()
		fileName := CreateName()
		err := SaveStr(projectRoot, fileName, s)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

	*/
}

// QRコードを生成する
func Generate() {

}

// QRコードを保存する
func SaveImage(projectRoot, fileName string, qrImage image.Image) error {
	imageSaveDir := filepath.Join(projectRoot, "QR_image")

	// ファイルを作成、pngとして保存するようにするが、その際にファイル名まで指定したい
	fullPath := filepath.Join(imageSaveDir, fileName, qrImage)
}

// QRコード生成元の文字列も保存　※名前の対応付けが必要
func SaveStr(projectRoot, fileName, s string) error {
	strSaveDir := filepath.Join(projectRoot, "QR_string")

	strFileName := ""
	strFileName = fileName + ".txt"

	fullPath := filepath.Join(strSaveDir, strFileName)
	err2 := os.WriteFile(fullPath, []byte(s), 0644)
	if err2 != nil {
		return err2
	}
	return nil
}

// 生成物に付ける名前を付ける
func CreateName() string {
	now := time.Now()

	// 20060102150405 はGoの仕様で決まっている日時
	FormattedTime := now.Format("20060102150405")
	return FormattedTime
}
