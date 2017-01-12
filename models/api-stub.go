package models

type Headers map[string]string
type JSON map[string]interface{}

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
	BodyJSON JSON    `json:"bodyJSON"`
	Headers  Headers `json:"headers"`
}

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
			JSON{
				"name": "Adam",
				"age":  23,
			},
			Headers{
				"Content-Type": "application/json",
			},
		},
	}
}
