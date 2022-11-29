package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const PFAD = "database/units.json"

// "database/units.json"
type Unit struct {
	Id   int    `json:"id"`
	Name string `json:"unit"`
}

func (m *Unit) Load() []Unit {
	jsonFile, err := os.Open(PFAD)
	if err != nil {
		fmt.Println("Read Json File", err)
	}
	defer jsonFile.Close()
	jsonByte, _ := ioutil.ReadAll(jsonFile)

	var result map[string][]Unit
	json.Unmarshal([]byte(jsonByte), &result)

	units := result["units"]
	return units
}

func (m *Unit) String() string {
	return fmt.Sprintf("Id: %d, Unit: %s", m.Id, m.Name)
}
