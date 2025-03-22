package model

type GroupMembersResponse struct {
	OdataContext  string  `json:"@odata.context"`
	OdataNextLink *string `json:"@odata.nextLink"`
	Value         []struct {
		OdataType string `json:"@odata.type"`
		Mail      string `json:"mail"`
	} `json:"value"`
}
