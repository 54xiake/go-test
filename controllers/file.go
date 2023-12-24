package test

import (
	"encoding/base64"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"math/rand"
	"strconv"
	"time"
)

type FileController struct {
	BaseController
}

func (c *FileController) UploadExcel() {
	fmt.Println("================UploadExcel ========")
	uploadType, _ := c.GetInt("type", 0)
	if uploadType <= 0 || uploadType > 3 {
		c.ErrorJson(500, "上传类型错误")
	}

	// 文件类型判断
	fileExt := c.GetString("suffix", "")
	if fileExt != "XLS" && fileExt != "XLSX" {
		c.ErrorJson(500, "上传文件格式错误")
	}

	// 随机一个文件名
	rand.Seed(time.Now().Unix())
	fileName := "/tmp/" + time.Now().Format("20060102150405") + strconv.Itoa(rand.Intn(10000)) + "." + fileExt
	// 读取原始文件内容
	fileBase64 := c.GetString("file", "")
	// base64转化
	decodeBytes, err := base64.StdEncoding.DecodeString(fileBase64)
	if err != nil {
		c.ErrorJson(500, err.Error())
	}
	// 存文件
	//func WriteFile(filename string, data []byte, perm os.FileMode) error
	err = ioutil.WriteFile(fileName, []byte(decodeBytes), 0666)
	if err != nil {
		c.ErrorJson(500, err.Error())
	}

	// 	用excel打开文件
	f_excel, err := excelize.OpenFile(fileName)
	if err != nil {
		logs.Info("excelize.OpenFile,err:" + err.Error())
		c.ErrorJson(500, err.Error())
	}

	// 用excel读取文件 func (f *File) GetSheetList() []string
	// 得到第一个工作表的名字
	sheetList := f_excel.GetSheetMap()
	// 得到第一个工作表的全部单元格
	rows, err := f_excel.GetRows(sheetList[0])
	if err != nil {
		logs.Info("excel.GetRows.err:" + err.Error())
		c.ErrorJson(500, err.Error())
	}
	fmt.Println(rows)
}
