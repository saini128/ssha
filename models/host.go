package models

type Host struct {
	Alias        string `json:"alias"`
	Hostname     string `json:"hostname"`
	User         string `json:"user"`
	Password     string `json:"password"`
	Port         int    `json:"port"`
	PrivateKey   string `json:"privateKey"`
	SecurityType string `json:"securityType"`
}
