package util

import (
	"os"
	"path"
)

// IsExist 判断文件是否存在，如果存在返回true，否则返回false
func IsExist(name string) bool {
	_, err := os.Stat(name)
	return os.IsExist(err)
}

// IsDir判断文件是否为文件夹，如果是返回true，如果不是或者文件不存在则返回false
func IsDir(name string) bool {
	stat, err := os.Stat(name)
	if os.IsNotExist(err){
		return false
	}
	return stat.IsDir()
}

// isExcel判断文件是否为excel文件
func isExcel(name string) bool {
	suffix := path.Base(name)
	return suffix == "xlsx" || suffix == "xls"
}

// isTxt判断文件是否为txt文件
func isTxt(name string) bool {
	suffix := path.Base(name)
	return suffix == "txt"
}