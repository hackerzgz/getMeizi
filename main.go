package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"sync"
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
	// baseURL = "http://gank.io/api/data/%E7%A6%8F%E5%88%A9/10/1"
	baseURL = "http://gank.io/api/data/%E7%A6%8F%E5%88%A9/10/1"
	// MaxGORO The Max Goroutine
	MaxGORO = 3
)

var (
	// FilePath Set Download Path
	FilePath string
	wg       sync.WaitGroup
)

func init() {
	if len(os.Args) < 2 {
		fmt.Printf("Please set up the Download Path:")
		fmt.Scanf("%s", &FilePath)
	} else {
		FilePath = os.Args[1]
	}
}

func main() {

	if FilePath == "" {
		os.Exit(-1)
	}
	// init Var
	t0 := time.Now()

	// Must Know How many Picture you Download!
	u, err := url.Parse(baseURL)
	if err != nil {
		fmt.Println("Your BaseURL is Wrong! Fix it!")
	}
	// Get The Number of Download Picture.
	reg := regexp.MustCompile(`[^\D]+`)
	count, err := strconv.Atoi(reg.FindAllString(u.Path, -1)[0])
	if err != nil {
		fmt.Println("Your BaseURL is Wrong! I don't konw how many Picture you want!")
	}
	fmt.Println(count, "Picture Download Now!")
	wg.Add(count)

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

	// Set the MaxGo
	var MaxGo int
	if MaxGORO == -1 {
		MaxGo = count
	} else {
		MaxGo = MaxGORO
	}
	Schedule := make(chan byte, MaxGo)
	// init Channel
	// make sure the Channel have enough element
	for i := 0; i < cap(Schedule); i++ {
		Schedule <- 0
	}

	// download images
	for i := 0; i < len(apiResult.Results); i++ {
		select {
		case <-Schedule:
			HandleDown(i, apiResult.Results[i], Schedule)
		case <-time.After(2 * time.Second):
			fmt.Println("Time out!")
			wg.Done()
		}
	}

	wg.Wait()
	// time.Sleep(10*time.Second)
	t1 := time.Now()
	fmt.Println("used time --> ", t1.Sub(t0).String())
	fmt.Println("Done!")
}

// HandleDown Handle Download in one Func
func HandleDown(i int, result resultObject, Schedule chan<- byte) {
	fmt.Println(i, " --> ", result.URL)

	// Check Dir Path Vaild
	FilePath = CheckDirPathVaild(FilePath)
	fmt.Println("FilePath -->", FilePath)

	if isExist(FilePath + result.ID + ".jpg") {
		fmt.Println(result.ID+".jpg", " Has been download!")
		// out!
		Schedule <- 0
		wg.Done()
		return
	}
	go SaveImage(result.URL, result.ID, Schedule)
}

// SaveImage Passing URL Location, Get Network Picture.
func SaveImage(url, filename string, sche chan<- byte) (n int64, err error) {
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
	// Here I can Get the images ContentLength,
	// But the io.Copy have a private methon,
	// I can not get the Buffer Progress.
	fmt.Printf("ContextLength --> %d\n", resp.ContentLength)
	defer resp.Body.Close()
	pix, err := ioutil.ReadAll(resp.Body)
	n, err = io.Copy(out, bytes.NewReader(pix))
	// out!
	sche <- 0
	wg.Done()
	return
}

// isExist Check the File has been Existed
func isExist(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil || os.IsExist(err)
}

// DirExists Check the Dir has been Existed
// and Create it if not Existed
// @param path string
// @return status bool
func DirExists(path string) bool {
	p, err := os.Stat(path)
	if err != nil {
		_ = os.Mkdir(path, 0755)
		return true
	}
	return p.IsDir()
}

// CheckDirPathVaild Check typing File Path is Vaild.
// @param filepath string
// @return filepath + '/'
func CheckDirPathVaild(filepath string) string {
	fpByte := []byte(filepath)
	if fpByte[len(fpByte)-1] != '/' {
		fpByte = append(fpByte, '/')
	}
	return string(fpByte)
}
