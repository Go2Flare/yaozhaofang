package main

import (
	"fmt"
	"github.com/tedcy/fdfs_client"
	"os"
	"path/filepath"
)

func main(){
	//	初始化客户端，配置文件
	workDir, _ := os.Getwd()
	client, err := fdfs_client.NewClientWithConfig(filepath.Join(workDir+"/fdfs.conf"))
	if err != nil {
		fmt.Println("初始化客户端错误",err)
		return
	}

	//上传文件，尝试文件名上传，传入服务器storage
	resp, err := client.UploadByFilename(filepath.Join(workDir+"/画画猫猫.jpg"))
	fmt.Println(resp, err)
}

//func TestUpload(t *testing.T) {
//	client, err := NewClientWithConfig("fdfs.conf")
//	defer client.Destory()
//	if err != nil {
//		fmt.Println(err.Error())
//		return
//	}
//	fileId, err := client.UploadByFilename("client_test.go")
//	if err != nil {
//		fmt.Println(err.Error())
//		return
//	}
//	fmt.Println(fileId)
//	if err := fdfs_client.DownloadToFile(fileId, "tempFile", 0, 0); err != nil {
//		fmt.Println(err.Error())
//		return
//	}
//	if buffer, err := client.DownloadToBuffer(fileId, 0, 19); err != nil {
//		fmt.Println(err.Error())
//	} else {
//		fmt.Println(string(buffer))
//	}
//	if err := client.DeleteFile(fileId); err != nil {
//		fmt.Println(err.Error())
//		return
//	}
//}