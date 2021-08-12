package solanarpc

import (
	"encoding/json"
)

type VersionResult struct {
	FeatureSet uint32 `json:"feature-set"`
	SolanaCore string `json:"solana-core"`
}

func (r SolanaRpc) RpcVersion() (*VersionResult, error) {
	resp, err := r.rpcCall(NewPostData("getVersion"))
	if err != nil {
		return nil, err
	}

	obj, err := r.toObjMap(resp)
	if err != nil {
		return nil, err
	}

	bytes, err := obj["result"].MarshalJSON()
	result := new(VersionResult)

	err = json.Unmarshal(bytes, result)
	return result, err
}
