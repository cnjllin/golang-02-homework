package store

import (
	"io/ioutil"
	"log"
	"encoding/json"
	"github.com/bitly/go-simplejson"
)

// dump and load json from file

var DataFile = "/tmp/dump.db"

func Dump(jsonData map[string]interface{}) {
	buf, err := json.Marshal(jsonData)
	if err != nil {
		log.Print(err)
	}

	err = ioutil.WriteFile(DataFile, buf, 0644)
	if err != nil {
		log.Print(err)
	}
}

func Load() (map[string]interface{}) {
	fileData, err := ioutil.ReadFile(DataFile)

	if err != nil {
		//log.Fatal(err)
		emptyJson := "{}"
		ioutil.WriteFile(DataFile, []byte(emptyJson), 0644)
		//log.Fatalf("Init an empty file: %s\n", DataFile) // exit 1
		log.Printf("Init an empty file: %s\n", DataFile)
		result := make(map[string]interface{}, 0)
		return result
	}

	//fmt.Println(fileData)

	resultJson, err := simplejson.NewJson(fileData)
	if err != nil {
		log.Print(err)
	}

	//fmt.Printf("%T %v \n", resultJson, resultJson)
	result, err := resultJson.Map()

	return result
}