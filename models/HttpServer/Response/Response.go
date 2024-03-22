package Response

type DevIsValidResp struct {
	IsValid        bool   `json:"isValid"`
}

type DevIsBindResp struct {
	BindCount        uint   `json:"bindCount"`
}
