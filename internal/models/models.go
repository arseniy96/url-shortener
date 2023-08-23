package models

// RequestCreateLink – структура запроса на роут `/api/shorten`
type RequestCreateLink struct {
	// URL – ссылка, которую необходимо сократить
	URL string `json:"url"`
}

// ResponseCreateLink – структура ответа на роут `/api/shorten`
type ResponseCreateLink struct {
	// Result – сокращённая ссылка
	Result string `json:"result"`
}

// RequestCreateLinksBatch – структура запроса на роут `/api/shorten/batch`
type RequestCreateLinksBatch []RequestLinks

// RequestLinks – структура содержимого RequestCreateLinksBatch
type RequestLinks struct {
	// CorrelationID – ID ссылки
	CorrelationID string `json:"correlation_id"`
	// OriginalURL – ссылка, которую необходимо сократить
	OriginalURL string `json:"original_url"`
}

// ResponseCreateLinksBatch – структура ответа на роут `/api/shorten/batch`
type ResponseCreateLinksBatch []ResponseLinks

// ResponseLinks – структура содержимого ResponseCreateLinksBatch
type ResponseLinks struct {
	// CorrelationID – ID ссылки
	CorrelationID string `json:"correlation_id"`
	// ShortURL – сокращённая ссылка
	ShortURL string `json:"short_url"`
}

// ResponseUserURLS – структура запроса на роут `/api/user/urls`
type ResponseUserURLS []ResponseUserURL

// ResponseUserURL – структура содержимого ResponseUserURLS
type ResponseUserURL struct {
	// ShortURL – сокращённая ссылка
	ShortURL string `json:"short_url"`
	// ShortURL – оригинальная ссылка
	OriginalURL string `json:"original_url"`
}

// RequestDeleteUserURLS – структура запроса на /api/user/urls
// содержит ссылки, которые были удалены
type RequestDeleteUserURLS []string
