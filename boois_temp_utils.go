package boois_temp_utils

import (
	"io/ioutil"
	"net/http"
	"fmt"
)

var TempList = make(map[string] string)

func GetTemp(file_name string,TempCached bool,OSS_Mode bool,OSS_URL string) string{
	if !TempCached {
		delete(TempList,file_name)
	}
	if file_content,ok := TempList[file_name];!ok{
		if OSS_Mode { //从oss中获取
			resp,err := http.Get(OSS_URL+file_name)
			if err != nil {
				fmt.Println("无法获取"+OSS_URL+file_name)
			}else{
				defer resp.Body.Close()
				body,_ := ioutil.ReadAll(resp.Body)
				file_content = string(body)
			}
		}else{ //从文件系统中获取
			f,file_err := ioutil.ReadFile("template/"+file_name)
			if file_err == nil{
				file_content = string(f)
			}
		}
		TempList[file_name] = file_content
		if OSS_Mode{
			fmt.Println("没有命中缓存,直接从OSS中读取模板文件!")
		}else{
			fmt.Println("没有命中缓存,直接从本地文件系统中读取模板文件!")
		}
	}
	return TempList[file_name]
}

func ClearTempCache(paths ...string)  {
	if len(paths) == 0 {
		for k,_:= range TempList{
			delete(TempList,k)
		}
	}else{
		for v := range paths{
			delete(TempList,paths[v])
		}
	}
}
