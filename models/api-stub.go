package models

type APIStub struct {
	Request  Request  `json:"request"`
	Response Response `json:"response"`
}

type Request struct {
	Method string `json:"method"`
	URL    string `json:"url"`
}

type Response struct {
	Status   int    `json:"status"`
	BodyJSON string `json:"bodyJSON"`
}

func APIStubTemplate() APIStub {
	return APIStub{
		Request{"GET", "/api-stub/template"},
		Response{200, `{"name": "Adam","age": 23}`},
	}
}
