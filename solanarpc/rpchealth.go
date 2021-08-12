package solanarpc

import (
	"io/ioutil"
	"net/http"
)

func (r SolanaRpc) HealthCheck() (bool, error) {
	str, err := r.RpcHealth()
	return str == "ok", err
}

func (r SolanaRpc) RpcHealth() (string, error) {
	resp, err := http.Get(r.address + "/health") //
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return "", err
	}

	return string(body), nil
}
