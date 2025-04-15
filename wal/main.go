package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

type WAL struct {
	filePath string
	mu       sync.Mutex
	file     *os.File
	store    map[string]string
}

const (
	logFormat = "%s\t%s\t%s\n" // operation, key, value
)

func NewWAL(filePath string) (*WAL, error) {
	// 创建或打开日志文件
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	wal := &WAL{
		filePath: filePath,
		file:     file,
		store:    make(map[string]string),
	}

	// 尝试恢复数据
	if err := wal.recover(); err != nil {
		return nil, err
	}

	return wal, nil
}

func (w *WAL) Put(key, value string) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	// 1. 先写入日志
	entry := fmt.Sprintf(logFormat, "PUT", key, value)
	if _, err := fmt.Fprint(w.file, entry); err != nil {
		return err
	}

	// 2. 立即同步到磁盘（重要！）
	if err := w.file.Sync(); err != nil {
		return err
	}

	// 3. 更新内存存储
	w.store[key] = value
	return nil
}

func (w *WAL) Get(key string) (string, bool) {
	w.mu.Lock()
	defer w.mu.Unlock()
	value, exists := w.store[key]
	return value, exists
}

func (w *WAL) recover() error {
	// 重新打开文件以读取模式
	file, err := os.Open(w.filePath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		var op, k, v string
		_, err := fmt.Sscanf(line, logFormat, &op, &k, &v)
		if err != nil {
			continue // 跳过无效行
		}

		if op == "PUT" {
			w.store[k] = v
		}
	}

	return scanner.Err()
}

func (w *WAL) writeEndMarker() error {
	if _, err := fmt.Fprint(w.file, "\n"); err != nil {
		return err
	}
	return w.file.Sync()
}

func (w *WAL) Close() error {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.writeEndMarker()
}

func main() {
	wal, err := NewWAL("wal.log")
	if err != nil {
		panic(err)
	}
	defer func(wal *WAL) {
		err := wal.Close()
		if err != nil {
			fmt.Printf("Error closing WAL: %s\n", err)
		}
	}(wal)

	// 正常写入
	if err := wal.Put("name", "Alice"); err != nil {
		panic(err)
	}
	if err := wal.Put("age", "30"); err != nil {
		panic(err)
	}

	// 模拟程序崩溃后重启
	fmt.Println("Oh!!! Crashing...")
	wal.store = nil
	fmt.Println("Current store:", wal.store)

	// 新实例会自动恢复数据
	newWal, err := NewWAL("wal.log")
	if err != nil {
		panic(err)
	}
	defer func(newWal *WAL) {
		err := newWal.Close()
		if err != nil {
			fmt.Printf("Error closing WAL: %s\n", err)
		}
	}(newWal)

	fmt.Println("Recovered store:", newWal.store)
}
