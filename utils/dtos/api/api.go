package api

type InsertReqBody struct {
	Dimensions []Dimension `json:"dim"`
	Metrics    []Metric    `json:"metrics"`
}

type InsertRespBody struct {
	Message string `json:"message"`
}

type QueryReqBody struct {
	Dimensions []Dimension `json:"dim"`
}

type QueryRespBody struct {
	Dimensions []Dimension `json:"dim"`
	Metrics    []Metric    `json:"metrics"`
}

type Dimension struct {
	Key string `json:"key"`
	Val string `json:"val"`
}

type Metric struct {
	Key string `json:"key"`
	Val int64  `json:"val"`
}
