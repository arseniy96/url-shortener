// Package models – содержит описание моделей
package models

// RequestCreateLink – структура запроса на роут `/api/shorten`.
type RequestCreateLink struct {
	// URL – ссылка, которую необходимо сократить.
	URL string `json:"url"`
}

// ResponseCreateLink – структура ответа на роут `/api/shorten`.
type ResponseCreateLink struct {
	// Result – сокращённая ссылка.
	Result string `json:"result" example:"http://localhost:8080/maIJa1"`
}

// RequestCreateLinksBatch – структура запроса на роут `/api/shorten/batch`.
type RequestCreateLinksBatch []RequestLinks

// RequestLinks – структура содержимого RequestCreateLinksBatch.
type RequestLinks struct {
	// CorrelationID – ID ссылки.
	CorrelationID string `json:"correlation_id" example:"58039b0a-480d-11ee-9ace-0e6250a0eb02" format:"uuid"`
	// OriginalURL – ссылка, которую необходимо сократить.
	OriginalURL string `json:"original_url" example:"https://ya.ru"`
}

// ResponseCreateLinksBatch – структура ответа на роут `/api/shorten/batch`.
type ResponseCreateLinksBatch []ResponseLinks

// ResponseLinks – структура содержимого ResponseCreateLinksBatch.
type ResponseLinks struct {
	// CorrelationID – ID ссылки.
	CorrelationID string `json:"correlation_id" example:"58039b0a-480d-11ee-9ace-0e6250a0eb02" format:"uuid"`
	// ShortURL – сокращённая ссылка.
	ShortURL string `json:"short_url" example:"http://localhost:8080/maIJa1"`
}

// ResponseUserURLS – структура запроса на роут `/api/user/urls`.
type ResponseUserURLS []ResponseUserURL

// ResponseUserURL – структура содержимого ResponseUserURLS.
type ResponseUserURL struct {
	// ShortURL – сокращённая ссылка.
	ShortURL string `json:"short_url"`
	// ShortURL – оригинальная ссылка.
	OriginalURL string `json:"original_url"`
}

// RequestDeleteUserURLS – структура запроса на /api/user/urls
// содержит ссылки, которые были удалены.
type RequestDeleteUserURLS []string
