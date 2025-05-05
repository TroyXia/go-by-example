package main

import (
	"encoding/json"
	"fmt"
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
	FromAccount string `json:"fromAccount"`
	ToAccount   string `json:"toAccount"`
	Amount      int    `json:"amount"`
}

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

	mu.Lock()
	defer mu.Unlock()

	fromBalance, exists := accounts[req.FromAccount]
	if !exists || fromBalance < req.Amount {
		http.Error(w, "Insufficient balance", http.StatusBadRequest)
		return
	}

	// 执行转账（非幂等操作：重复请求会多次扣款！）
	accounts[req.FromAccount] -= req.Amount
	accounts[req.ToAccount] += req.Amount

	// 返回成功响应（未记录请求唯一性）
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Transfer success. New balances: A:%d, B:%d",
		accounts["A"], accounts["B"])
}

func main() {
	http.HandleFunc("/transfer", transferHandler)
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
