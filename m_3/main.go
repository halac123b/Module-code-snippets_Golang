package m_3

type Response struct {
	Status     string
	StatusCode int
	Method     string
	Body       []byte
}

type ZenQuotes struct {
	Quote  string `json:"q"`
	Author string `json:"a"`
}
