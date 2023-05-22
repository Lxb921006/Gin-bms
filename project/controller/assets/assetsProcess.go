package assets

import (
	"encoding/json"
)

type ProcessData struct {
	Ip          string `json:"ip"`
	ProcessName string `json:"name"`
}

func (p *ProcessData) GetJsonData(data string) (err error) {

	if err = json.Unmarshal([]byte(data), p); err != nil {
		return
	}

	return
}
