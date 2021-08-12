package solanarpc

import (
	"bytes"
	"encoding/json"
)

type PostData struct {
	JsonRpc string      `json:"jsonrpc"`
	ID      int         `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
}

// 1 - method name, 2 - params
func NewPostData(fields ...interface{}) *bytes.Buffer {
	data := PostData{
		JsonRpc: "2.0",
		ID:      1,
		Method:  fields[0].(string),
	}

	if len(fields) > 1 {
		data.Params = fields[1]
	}

	post, _ := json.Marshal(data) //TODO err?

	return bytes.NewBuffer(post)
}
