package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"theia-stu/util"
	"time"
)

const (
	Version = `201012-Beta`
	tips = `===========================================
Theia
Author:BuTn<https://github.com/kimmosc2>
Versoion:`+Version+`
===========================================`
)

var (
	// 名单文件
	listFile string
	// 目标文件夹
	targetDir string
	// 已交作业名单
	completeList = make(map[string]string, 0)
	// 未交作业名单
	UnfinishedList = make(map[string]string, 0)
)

// TODO 考虑到受众群体，暂时先不加flag功能
func main() {
	logInit()
	fmt.Println(tips)
	fmt.Print("请输入名单文件:")
	// TODO 自动判断txt,excel
	_, err := fmt.Scanf("%s\n", &listFile)
	if err != nil {
		log.Fatalf("scanf listFile error:%s", err.Error())
	}

	fmt.Print("请输入作业文件夹:")
	_, err = fmt.Scanf("%s\n", &targetDir)
	if err != nil {
		log.Fatalf("scanf targetDir error:%s", err.Error())
	}
	fmt.Println("-------------------------------------------")

	// 从excel里读出的完整用户名单
	userList := util.ReadExcel(listFile)

	// 从目标文件夹里读出的文件列表
	fileList := GetFileNameList(targetDir)

	for _, fName := range fileList {
		// 找到标识
		if uid, ok := userList[fName]; ok {
			// 说明交作业了，不过没写别的信息，只写了个名字
			completeList[fName] = uid
			continue
		}

		found := false
		for k, v := range userList {
			if strings.Contains(fName, k) {
				completeList[k] = v
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("名为《%s》的文件未在名单内找到，请确认书写是否正确\n", fName)
		}
	}
	fmt.Println("-------------------------------------------")
	fmt.Printf("已交学生名单(共%d个)\n",len(completeList))
	for k, v := range completeList {
		fmt.Printf("名称:%s\t学号:%s\n", k, v)
	}
	fmt.Println("-------------------------------------------")
	for user, _ := range userList {
		if _, ok := completeList[user]; !ok {
			UnfinishedList[user] = userList[user]
		}
	}
	fmt.Printf("未完成学生名单(共(%d个)\n",len(UnfinishedList))
	for name, number := range UnfinishedList {
		fmt.Printf("姓名:%s\t学号:%s\n", name, number)
	}
	fmt.Println("-------------------------------------------")
	for {
		time.Sleep(time.Hour)
	}
}

// GetFileNameList返回文件基础名，不带后缀的
func GetFileNameList(dirname string) []string {
	nameList := make([]string, 0)
	if dir, err := ioutil.ReadDir(dirname); err == nil {
		for _, info := range dir {
			fileName := filepath.Base(info.Name())
			extName := filepath.Ext(info.Name())
			fileName = fileName[:len(fileName)-len(extName)]
			nameList = append(nameList, fileName)
		}
	}
	return nameList
}

func logInit() {
	log.SetPrefix("[Theia]")
	logName := "theia.log"

	// 判断日志文件是否存在
	if util.IsExist(logName) {
		file, _ := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
		log.SetOutput(file)
	} else {
		create, err := os.Create(logName)
		if err != nil {
			log.Fatalf("创建日志文件失败:%s", err.Error())
		}
		log.SetOutput(create)
	}

}
