// 配列操作のアウトプットとして8x8のリバーシを作成

package main

import "fmt"

const (
	Empty = 0 // 空白
	Black = 1 // 黒
	White = 2 // 白
)

// Board型の定義
type Board [8][8]int

// ボードの初期化
func initializeBoard() Board {
	board := Board{}
	// 初期配置
	board[3][3], board[4][4] = White, White
	board[3][4], board[4][3] = Black, Black
	return board
}

// ボードを表示
func printBoard(board Board) {
	fmt.Println("--------------------------------------------------------")
	fmt.Println("  0 1 2 3 4 5 6 7")
	for y := 0; y < 8; y++ {
		fmt.Print(y, " ")
		for x := 0; x < 8; x++ {
			switch board[y][x] {
			case Empty:
				fmt.Print(". ")
			case Black:
				fmt.Print("● ")
			case White:
				fmt.Print("○ ")
			}
		}
		fmt.Println()
	}
}

// 方向ベクトル
var directions = [][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

// コマを置けるか判定
func canPlace(board Board, x, y, color int) bool {
	if board[y][x] != Empty {
		return false
	}
	opponent := Black
	if color == Black {
		opponent = White
	}
	for _, dir := range directions {
		dx, dy := dir[0], dir[1]
		nx, ny := x+dx, y+dy
		foundOpponent := false
		for nx >= 0 && ny >= 0 && nx < 8 && ny < 8 {
			if board[ny][nx] == opponent {
				foundOpponent = true
			} else if board[ny][nx] == color && foundOpponent {
				return true
			} else {
				break
			}
			nx += dx
			ny += dy
		}
	}
	return false
}

// コマを置く
func placePiece(board *Board, x, y, color int) {
	board[y][x] = color
	opponent := Black
	if color == Black {
		opponent = White
	}
	for _, dir := range directions {
		dx, dy := dir[0], dir[1]
		nx, ny := x+dx, y+dy
		flipPositions := [][2]int{}
		for nx >= 0 && ny >= 0 && nx < 8 && ny < 8 {
			if board[ny][nx] == opponent {
				flipPositions = append(flipPositions, [2]int{nx, ny})
			} else if board[ny][nx] == color {
				for _, pos := range flipPositions {
					board[pos[1]][pos[0]] = color
				}
				break
			} else {
				break
			}
			nx += dx
			ny += dy
		}
	}
}

func playerName(color int) string {
	if color == Black {
		return "黒 (●)"
	}
	return "白 (○)"
}

// ゲームのメインループ
func main() {
	board := initializeBoard()
	currentPlayer := Black

	for {
		printBoard(board)
		fmt.Println("\"9 9\" と入力すると終了します。")
		fmt.Printf("プレイヤー %s のターンです(x y)： ", playerName(currentPlayer))
		var x, y int
		fmt.Scan(&x, &y)

		if x == 9 || y == 9 {
			break
		}

		if x < 0 || x >= 8 || y < 0 || y >= 8 || !canPlace(board, x, y, currentPlayer) {
			fmt.Println("その場所には置けません。")
			continue
		}

		placePiece(&board, x, y, currentPlayer)

		// 次のプレイヤーに交代
		if currentPlayer == Black {
			currentPlayer = White
		} else {
			currentPlayer = Black
		}
	}

	fmt.Println("ゲームを終了します。")
}
