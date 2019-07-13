package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/djimenez/iconv-go"
)

func Now(deli string) string {
	if deli == "" {
		deli = "/"
	}
	return time.Now().Format("01" + deli + "02")
}

func Scrape() string {
	url := "http://www.nnh.to/" + Now("")
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return enc2utf8(string(body))
}

func enc2utf8(s string) string {
	u, err := iconv.ConvertString(s, "shift-jis", "utf-8")
	if err != nil {
		panic(err)
	}
	return u
}

func createResDir(dirname string) {
	var e error
	if _, err := os.Stat(dirname); os.IsNotExist(err) {
		e = os.Mkdir(dirname, os.ModePerm)
	}
	if e != nil {
		panic(e)
	}
}

func Save(filename string, value string) string {
	createResDir("resource")

	path := "resource/" + filename + ".html"
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString(value)
	if err != nil {
		panic(err)
	}
	return path
}

func main() {
	fmt.Println(Save(Now("-"), Scrape()))
}
