
package store

import (
	"github.com/a8m/djson"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"encoding/json"
	"os"
)

func init()  {
	_, err := os.Stat("dump.db")
	if err != nil {
		file, error := os.Create("dump.db")

		if error != nil {
			log.Errorf("Creating file error: %s", error)
			return
		}
		defer file.Close()
	}
}

func Dump(data map[string]interface{}) {
	j, err := json.Marshal(data)
	if err != nil {
		log.Errorf("Marshal json error: %s", err)
		return
	}
	ioutil.WriteFile("dump.db", j, 0644)
}

func Load() map[string]interface{}{
	res, err := ioutil.ReadFile("dump.db")
	if err != nil {
		log.Errorf("Load db error: %s", err)
		return nil
	}

	if len(res) == 0 {
		return make(map[string]interface{})
	}

	jsonRes, err := djson.DecodeObject(res)
	if err != nil {
		log.Errorf("Unmarshal error: %s", err)
		return nil
	}

	return jsonRes
}
