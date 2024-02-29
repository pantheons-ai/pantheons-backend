package rest

type CreateIPReq struct {
	Name string // the name of IP universal
	ContractAddress string  // address of contract
	FounderAddress string  // address of founder
}

type CreateIPResp struct {

}

type AddUserReq struct {
	ContractAddress string // address of user
	UserAddress string
}

type AddUserResp struct {

}