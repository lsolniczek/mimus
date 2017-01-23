package models

type Headers map[string]string
type JSON map[string]interface{}

type APICase struct {
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

func APICaseTemplate() APICase {
	return APICase{
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
