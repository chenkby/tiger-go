package model

type Column struct {
	Name      string `json:"name"`
	Default   string `json:"default"`
	IsNull    string `json:"is_null"`
	MaxLength int64  `json:"max_length"`
	Precision int64  `json:"precision"`
	Type      string `json:"type"`
	Key       string `json:"key"`
	Extra     string `json:"extra"`
	Comment   string `json:"comment"`
}
