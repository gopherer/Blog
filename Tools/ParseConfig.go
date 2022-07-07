package Tools

import (
	"blog/Model"
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

var Cfg *Model.Config

func ParseConfig(filePath string) error {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0644) //文件所有者对该文件有读写权限，用户组和其他人只有读权限，没有执行权限
	if err != nil {
		fmt.Println(err) //后续会专门处理
		os.Exit(1)
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()
	reader := bufio.NewReader(file)    //NewDecoder创建一个从r读取并解码json对象的*Decoder，解码器有自己的缓冲，并可能超前读取部分json数据。
	decoder := json.NewDecoder(reader) //Decode从输入流读取下一个json编码值并保存在v指向的值里，参见Unmarshal函数的文档获取细节信息。
	err = decoder.Decode(&Cfg)
	if err != nil {
		fmt.Println(err)
		return err
	} else {
		return nil
	}

}
