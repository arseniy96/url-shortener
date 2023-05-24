package models

type RequestCreateLink struct {
	URL string `json:"url"`
}

type ResponseCreateLink struct {
	Result string `json:"result"`
}
