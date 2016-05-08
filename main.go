package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"bytes"
	"io"
	"time"
)

// APIJSON 注意，转换JSON的时候首字母必须大写，否则转换不成功
type APIJSON struct {
	Err     bool           `json:"error"`
	Results []resultObject `json:"results"`
}

type resultObject struct {
	ID          string `json:"_id"`
	CreateAt    string `json:"createdAt"`
	Desc        string `json:"desc"`
	PublishedAt string `json:"publishedAt"`
	Source      string `json:"source"`
	ImgType     string `json:"type"`
	URL         string `json:"url"`
	Used        bool   `json:"used"`
	Who         string `json:"who"`
}

const (
	baseURL = "http://gank.io/api/data/%E7%A6%8F%E5%88%A9/5/1"
	// baseURL = "http://7xltko.com1.z0.glb.clouddn.com/gank.io"
	// returnCount = 10
	// pageNum = 1
)

var (
	FilePath string // Set Download Path
)

func init() {
	if len(os.Args) < 2 {
		fmt.Printf("%s", "Please set up the download path:")
		fmt.Scanf("%s", &FilePath)
	} else {
		FilePath = os.Args[1]
	}
}

func main() {
	// init Var
	t0 := time.Now()
	Schedule := make(chan byte, 10)

	// Get the API Data
	res, err := http.Get(baseURL)
	defer res.Body.Close()
	if err != nil {
		fmt.Println("Read API Address Error --> ", baseURL)
		fmt.Println("Error --> ", err.Error())
		return
	}
	body, _ := ioutil.ReadAll(res.Body)

	// JSON 2 Struct
	var apiResult APIJSON
	if err1 := json.Unmarshal(body, &apiResult); err1 != nil {
		fmt.Println("JSON Data Translate Error --> ", err1.Error())
	}

	// init Channel
	// make sure the Channel have enough element
	for i := 0; i < len(apiResult.Results); i++ {
		Schedule <- 0
	}
	// download images
	for i := 0; i < len(apiResult.Results); i++ {
		fmt.Println(i, " --> ", apiResult.Results[i].URL)
		if isExist(FilePath + apiResult.Results[i].ID + ".jpg") {
			fmt.Println(apiResult.Results[i].ID+".jpg", " Has been download!")
			// out!
			<-Schedule
			continue
		}
		go SaveImage(apiResult.Results[i].URL, apiResult.Results[i].ID, Schedule)
	}

	for {
		fmt.Println("Len --> ", len(Schedule))
		if len(Schedule) == 0 {
			break
		}
	}
	t1 := time.Now()
	fmt.Println("used time --> ", t1.Sub(t0).String())
	fmt.Println("Done!")
}

// SaveImage 传入URL地址，获取网络图片
func SaveImage(url, filename string, sche <-chan byte) (n int64, err error) {
	DirExists(FilePath)
	out, err := os.Create(FilePath + filename + ".jpg")
	if err != nil {
		fmt.Printf("%s File Create Failed!\n", FilePath+filename+".jpg")
		return
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("=== Error ===")
		fmt.Println(err)
		fmt.Println("=== Error ===")
	}
	defer resp.Body.Close()
	pix, err := ioutil.ReadAll(resp.Body)
	n, err = io.Copy(out, bytes.NewReader(pix))
	// out!
	<-sche
	return
}

// isExist Check the File has been Existed
func isExist(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil || os.IsExist(err)
}

// DirExists Check the Dir has been Existed
// and Create it if not Existed
func DirExists(path string) bool {
	p, err := os.Stat(path)
	if err != nil {
		_ = os.Mkdir(path, 0755)
		return true
	}
	return p.IsDir()
}

// CreateDir Mean Mkdir 0755
// func CreateDir(path string) bool {
// 	err := os.Mkdir(path, 0755)
// 	if err != nil {
// 		return os.IsExist(err)
// 	}
// 	return true
// }
