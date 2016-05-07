package main

import (
    "fmt"
    "net/http"
    "os"
    "io/ioutil"
    "encoding/json"
	// "path/filepath"
    "bytes"
    "io"
)

// 注意，转换JSON的时候首字母必须大写，否则转换不成功
type APIJSON struct {
    Err bool `json:"error"`
    Results []resultObject `json:"results"`
}

type resultObject struct {
    ID string `json:"_id"`
    CreateAt string `json:"createdAt"`
    Desc string `json:"desc"`
    PublishedAt string `json:"publishedAt"`
    Source string `json:"source"`
    ImgType string `json:"type"`
    URL string `json:"url"`
    Used bool `json:"used"`
    Who string `json:"who"`
}

const (
    baseURL = "http://gank.io/api/data/%E7%A6%8F%E5%88%A9/10/1"
    // returnCount = 10
    // pageNum = 1
)

func main()  {
    
    res, err := http.Get(baseURL)
    defer res.Body.Close()
    if err != nil {
        fmt.Println("API地址读取出错 --> ", baseURL)
        fmt.Println("Error --> ", err.Error())
        return        
    }
    body, _ := ioutil.ReadAll(res.Body)
    // fmt.Println("body --> ", string(body))
    
    var apiResult APIJSON
    if err1 := json.Unmarshal(body, &apiResult); err1 != nil {
        fmt.Println("JSON数据转换出错 --> ", err.Error())
    }
    
    // for i := 0; i < len(apiResult.Results); i++ {
    //     fmt.Println(i , " --> ", apiResult.Results[i].URL)
    //     SaveImage(apiResult.Results[i].URL, apiResult.Results[i].ID)
    // }
    SaveImage(apiResult.Results[1].URL, apiResult.Results[1].ID)
}

// SaveImage 传入URL地址，获取网络图片（后期可能改成并发的）
func SaveImage(url, filename string)(n int64, err error)  {
    out, err := os.Create(filename+".jpg")
    defer out.Close()
    
    resp, err := http.Get(url)
    // 这里出现了一个错误：
    // Get http://ww1.sinaimg.cn/large/7a8aed7bjw1f3k9dp8r9qj20dw0jljtd.jpg: write tcp 10.15.32.248:63418->112.90.6.238:80: wsasend: An operation was attempted on something that is not a socket.
    if err != nil {
        fmt.Println("=== Error ===")
        fmt.Println(err)
        fmt.Println("=== Error ===")        
    }
    defer resp.Body.Close()
    pix, err := ioutil.ReadAll(resp.Body)
    n, err = io.Copy(out, bytes.NewReader(pix))
    return
    // if err != nil {
    //     // fmt.Println("保存", filename, "图片出错！ URL --> ", url)
    //     fmt.Println("Error --> ", err)
    // }
    // fmt.Println(resp.Body)
    // // 下载并保存
    // //根据URL文件名创建文件
    // // filename = filepath.Base()
    // dst, err := os.Create(filename)
    // if err != nil {
    //     fmt.Printf("%s HTTP ERROR:%s\n", filename, err)
    //     return
    // }
    // // 写入文件
    // io.Copy(dst, res.Body)
}