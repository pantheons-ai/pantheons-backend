package rest

type CreateIPReq struct {
	// the name of IP universal
	Name string `json:"name"`
	// address of contract
	ContractAddress string `json:"contract_address"` 
	// address of founder
	CreatorAddress string  `json:"creator_address"`
}

type CreateIPResp struct {
	FolderCID string `json:"folder_cid"`
}

type AddUserReq struct {
	ContractAddress string `json:"contract_address"` 
	UserAddress string `json:"user_address"`
}

type PublishContentReq struct {
	ContractAddress string
	UserAddress string
	Bytes         int    `json:"bytes"`
	FileName      string `json:"filename"`
}

type PublishContentResp struct {
	Cid string `json:"cid"`
}