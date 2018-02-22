package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/Tomoka64/final/model"
)

const DBPath = "config/data.json"

type History struct {
	datas []model.Result
}

func newHistory(items ...string) (another, error) {
	datas, err := ListAll()
	if err != nil {
		return &History{}, err
	}
	return &History{
		datas: datas,
	}, nil
}
func (r *History) Run() error {
	for _, data := range r.datas {
		fmt.Printf("%v\t%v\t%d\t%v\n", data.Filename, data.Keyword, data.Line, data.Detail)
	}
	return nil
}

func FileGetContents(filename string) []byte {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}
	return contents
}

//ListAll shows all the search history in terminal
func ListAll() ([]model.Result, error) {
	var datas []model.Result
	contents := FileGetContents(DBPath)
	dec := json.NewDecoder(bytes.NewReader(contents))

	for {
		var data model.Result
		if err := dec.Decode(&data); err == io.EOF {
			break
		} else if err != nil {
			return []model.Result{}, err
		}
		datas = append(datas, data)
	}
	return datas, nil
}

//SaveToFile saves all the searched results into config/data.json
func SaveToFile(r *model.Result) error {
	f, err := os.OpenFile(DBPath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	b, err := json.Marshal(&r)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	if _, err = f.WriteString(string(b[:])); err != nil {
		log.Fatalln(err)
	}
	return err
}
