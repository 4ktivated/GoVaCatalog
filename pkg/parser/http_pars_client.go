package parser

import (
	"net/http"
	curl "net/url"

	"some_app/internal/repository"
)


//храним параметры опции запроса в базе 
//"https://api.hh.ru/vacancies"
// params.Add("text", lang)
// params.Add("page", page)
// params.Add("per_page", per_page)

type HTTPParseClient struct{
	title string
}

func NewHTTPParseClient(title string) *HTTPParseClient {
	return &HTTPParseClient{title}
}


func (h *HTTPParseClient) recvData(title string) (*http.Response, error) {
	
	data := repository.Data{}
	UP, err := data.GetOptUrl("hh")
	if err != nil {
		return nil, err
	}

	params := curl.Values{}
	for prm, val := range UP.Param {
		params.Add(prm, val)
	}

	path, err := curl.Parse(UP.Url)
	if err != nil {
		return nil, err
	}
	path.RawQuery = params.Encode()

	response, err := http.Get(path.String())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return response, nil
}

func (h *HTTPParseClient) SendData() (*http.Response, error) {//TODO: зачем я это сделал?
	var response *http.Response
	return response, nil
}


