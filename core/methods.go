package core

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func ProcessCurl(curl Curl) string {

	req, err := http.NewRequest(curl.Method, curl.Url, bytes.NewBuffer([]byte(curl.Body)))

	req.Header = curl.Headers

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		println("ProcessCurl ERROR: ", err)
		return ""
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}

func GetFileParsed(filename string, replace map[string]string) ([]byte, error) {
	file, err := ioutil.ReadFile(filename)
	if replace != nil {
		for k, v := range replace {
			file = bytes.Replace(file, []byte(k), []byte(v), -1)
		}
	}
	return file, err
}

func SaveFileLine(item string, dir string) {
	Log(item)
	f, err := os.OpenFile(dir, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := f.Write([]byte(item)); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func PrepareFile(fileName string) {
	if _, err := os.Stat(fileName); err == nil {
		err := os.Remove(fileName)

		if err != nil {
			log.Fatal(err)
			return
		}
	}
	Log(fileName + ": CREATED!")
}

func Log(text string) {
	fmt.Println(text)
}

func ReadCSV(fileName string) *csv.Reader {
	println("ReadCSV -file:" + fileName)

	csvFile, _ := os.Open(fileName)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Comma = ';'
	return reader
}
