package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gonum.org/v1/gonum/fourier"
)

// FFTRequest はAPIで受け取るリクエストデータの形式
type FFTRequest struct {
	Real []float64 `json:"real"` // 実部の配列
	Imag []float64 `json:"imag"` // 虚部の配列
}

// FFTResponse はAPIが返すレスポンスデータの形式
type FFTResponse struct {
	Real []float64 `json:"real"` // 結果の実部
	Imag []float64 `json:"imag"` // 結果の虚部
}

func performFFT(req FFTRequest) (FFTResponse, error) {
	// 入力データの長さ確認
	if len(req.Real) != len(req.Imag) {
		return FFTResponse{}, fmt.Errorf("Real and Imaginary parts must have the same length")
	}

	// 複素数配列の作成
	input := make([]complex128, len(req.Real))
	for i := range req.Real {
		input[i] = complex(req.Real[i], req.Imag[i])
	}

	// FFTの計算
	fft := fourier.NewFFT(len(input))
	output := fft.Coefficients(nil, input)

	// 結果を分解
	real := make([]float64, len(output))
	imag := make([]float64, len(output))
	for i, v := range output {
		real[i] = real(v)
		imag[i] = imag(v)
	}

	return FFTResponse{Real: real, Imag: imag}, nil
}

func fftHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is supported", http.StatusMethodNotAllowed)
		return
	}

	var req FFTRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	result, err := performFFT(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/fft", fftHandler)

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
