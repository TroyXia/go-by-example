package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"sync"
)

// 账户余额（模拟数据库）
var accounts = map[string]int{
	"A": 1000, // 初始余额 1000 元
	"B": 1000,
}

var mu sync.Mutex

// 转账请求结构体
type TransferRequest struct {
	RequestId   string `json:"requestId"`
	FromAccount string `json:"fromAccount"`
	ToAccount   string `json:"toAccount"`
	Amount      int    `json:"amount"`
}

var requestHandle = make(map[string]bool)

func transferHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 解析请求参数
	var req TransferRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if requestHandle[req.RequestId] {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "request %s already handled", req.RequestId)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	fromBalance, exists := accounts[req.FromAccount]
	if !exists || fromBalance < req.Amount {
		http.Error(w, "Insufficient balance", http.StatusBadRequest)
		return
	}

	// 执行转账（幂等操作：重复请求记录 request id 表示唯一性）
	accounts[req.FromAccount] -= req.Amount
	accounts[req.ToAccount] += req.Amount

	requestHandle[req.RequestId] = true

	// 返回成功响应（未记录请求唯一性）
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Transfer success. New balances: A:%d, B:%d",
		accounts["A"], accounts["B"])
}

func generateUUID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	requestId, err := uuid.NewUUID()
	if err != nil {
		http.Error(w, "failed to generate uuid", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "UUID generated: %s", requestId)
}

func main() {
	http.HandleFunc("/transfer", transferHandler)
	http.HandleFunc("/uuid", generateUUID)
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
