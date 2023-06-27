package models

type RequestCreateLink struct {
	URL string `json:"url"`
}

type ResponseCreateLink struct {
	Result string `json:"result"`
}

type RequestCreateLinksBatch []RequestLinks

type RequestLinks struct {
	CorrelationID string `json:"correlation_id"`
	OriginalURL   string `json:"original_url"`
}

type ResponseCreateLinksBatch []ResponseLinks

type ResponseLinks struct {
	CorrelationID string `json:"correlation_id"`
	ShortURL      string `json:"short_url"`
}

type ResponseUserURLS []ResponseUserURL

type ResponseUserURL struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}
