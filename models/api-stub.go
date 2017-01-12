package models

type APIStub struct {
	Request  Request  `json:"request"`
	Response Response `json:"response"`
}

type Request struct {
	Method  string  `json:"method"`
	URL     string  `json:"url"`
	Headers Headers `json:"headers"`
}

type Response struct {
	Status   int     `json:"status"`
	BodyJSON string  `json:"bodyJSON"`
	Headers  Headers `json:"headers"`
}

type Headers map[string]string

func APIStubTemplate() APIStub {
	return APIStub{
		Request{
			"GET",
			"/api-stub/template",
			Headers{
				"Content-Type": "application/json",
			},
		},
		Response{
			200,
			`{"name": "Adam","age": 23}`,
			Headers{
				"Content-Type": "application/json",
			},
		},
	}
}
