package util

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"log"
	"strconv"
)

// ReadExcel读取一个excel文件，返回一个map，key为姓名，value为学号
// 注:部分人有不写学号的习惯，所以用学号做key会影响后续处理
func ReadExcel(name string) map[string]string {
	f, err := excelize.OpenFile(name)
	if err != nil {
		log.Fatalf("打开excel文件失败:%s", err.Error())
	}
	userList := make(map[string]string)
	// 获取工作表中指定单元格的值
	var i = 1
	for {
		s := strconv.Itoa(i)
		number := f.GetCellValue("Sheet1", "A"+s)
		name := f.GetCellValue("Sheet1", "B"+s)
		// 如果name为空,这里学号是可选的,可以只写姓名不写学号
		if name == "" {
			break
		}
		userList[name] = number
		i++
	}
	return userList
}
