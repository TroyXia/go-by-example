package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/xuri/excelize/v2"
)

func main() {
	// 设置路由
	http.HandleFunc("/export", exportExcelHandler)

	// 启动服务器
	fmt.Println("Server starting on http://127.0.0.1:8085")
	log.Fatal(http.ListenAndServe(":8085", nil))
}

func exportExcelHandler(w http.ResponseWriter, r *http.Request) {
	// 创建一个新的Excel文件
	f := excelize.NewFile()

	// 创建一个工作表
	index, err := f.NewSheet("Sheet1")
	if err != nil {
		log.Printf("创建工作表失败: %v", err)
		http.Error(w, "创建Excel失败", http.StatusInternalServerError)
		return
	}

	// 设置单元格值
	f.SetCellValue("Sheet1", "A1", "姓名")
	f.SetCellValue("Sheet1", "B1", "年龄")
	f.SetCellValue("Sheet1", "C1", "城市")

	f.SetCellValue("Sheet1", "A2", "张三")
	f.SetCellValue("Sheet1", "B2", 25)
	f.SetCellValue("Sheet1", "C2", "北京")

	f.SetCellValue("Sheet1", "A3", "李四")
	f.SetCellValue("Sheet1", "B3", 30)
	f.SetCellValue("Sheet1", "C3", "上海")

	f.SetCellValue("Sheet1", "A4", "王五")
	f.SetCellValue("Sheet1", "B4", 28)
	f.SetCellValue("Sheet1", "C4", "广州")

	// 设置活动工作表
	f.SetActiveSheet(index)

	// 设置响应头，使浏览器下载文件
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Set("Content-Disposition", "attachment; filename=export.xlsx")
	w.Header().Set("Cache-Control", "no-cache")

	// 将Excel文件写入响应体
	if err := f.Write(w); err != nil {
		log.Printf("写入Excel文件失败: %v", err)
		http.Error(w, "导出Excel失败", http.StatusInternalServerError)
		return
	}
}