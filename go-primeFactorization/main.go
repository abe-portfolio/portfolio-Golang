// 素因数分解を行う関数の作成
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func PrimeFactorization(x int) [][2]int {
	var primes [][2]int

	for i := 2; i*i <= x; i++ {
		if x%i != 0 {
			continue
		}

		var e int

		for x%i == 0 {
			e++
			x /= i
		}

		primes = append(primes, [2]int{i, e})
	}

	if x != 1 {
		primes = append(primes, [2]int{x, 1})
	}

	return primes
}

func StrStdin() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main() {
INPUT:
	fmt.Print("Input x:")
	s := StrStdin()
	x, _ := strconv.Atoi(s)

	if x <= 0 {
		fmt.Println("正の数で入力してください。")
		goto INPUT
	}
	result := PrimeFactorization(x)
	fmt.Println(result)
}
