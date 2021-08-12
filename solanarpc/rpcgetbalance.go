package solanarpc

func (r SolanaRpc) RpcGetBalance() (*string, error) {
	resp, err := r.rpcCall(NewPostData("getBalance", []string{"Gbwg3HCbD9gzma2o6oTqYkDQnojeZ1z7Ygc95DPFA2me"}))
	if err != nil {
		return nil, err
	}

	obj, err := r.toObjMap(resp)
	if err != nil {
		return nil, err
	}

	r.logger.Info(obj)

	return nil, err
}
