package often

import (
	"encoding/csv"
	"fmt"
	"os"
)

// LogEngine oop
type LogEngine struct {
	writeObj *csv.Writer
}

//NewLog return a new blank Engine instance
func NewLog(logPath string) *LogEngine {
	var writeObj *csv.Writer
	_, er := os.Open(logPath)
	if er != nil && os.IsNotExist(er) {
		file, er := os.Create(logPath)
		if er != nil {
			panic(er)
		}
		//defer file.Close()
		// 写入字段标题
		writeObj = csv.NewWriter(file) //创建一个新的写入文件流

	} else {
		// 如果文件存在，直接加在末尾
		txt, err := os.OpenFile(logPath, os.O_APPEND|os.O_RDWR, 0666)
		//defer txt.Close()
		if err != nil {
			panic(err)
		}
		writeObj = csv.NewWriter(txt) //创建一个新的写入文件流
	}

	engine := &LogEngine{writeObj}
	return engine
}

//Info push string info into log file
func (engine *LogEngine) Info(data []string) {
	engine.writeObj.Write(data)
	engine.writeObj.Flush()
	// 这里必须刷新，才能将数据写入文件。
}

// Log use to log some info
func Log(logPath string) *LogEngine {
	engine := NewLog(logPath)
	//engine.info()
	return engine
}
func beautyPrint(obj map[string]interface{}) {
	fmt.Println("\033[35m{\033[0m")
	for key, value := range obj {
		v, isString := value.(string)
		if isString {
			fmt.Printf("\033[32m   \"%s\"\033[0m : \033[33m\"%v\"\033[0m\n", key, v)
		} else {
			fmt.Printf("\033[32m   \"%s\"\033[0m : \033[33m%v\033[0m\n", key, value)
		}
	}
	fmt.Println("\033[35m}\033[0m")
}
