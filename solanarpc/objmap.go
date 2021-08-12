package solanarpc

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Objmap map[string]json.RawMessage

func (r SolanaRpc) toObjMap(resp *http.Response) (Objmap, error) {
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	var objmap = Objmap{}
	err = json.Unmarshal(body, &objmap)
	return objmap, err
}
