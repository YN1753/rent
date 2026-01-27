package model

type UploadParams struct {
	Id        string `json:"id"`
	Policy    string `json:"policy"`
	Signature string `json:"signature"`
	KeyPrefix string `json:"keyPrefix"`
	Host      string `json:"host"`
}
