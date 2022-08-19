package pinata

import (
	"fmt"
	"io"
	"net/http"
)

type Params struct {
	Key         string
	Value       interface{}
	SecondValue *interface{}
	Operator    string
}

type QueryOneParams struct {
	Value    interface{} `json:"value"`
	Operator string      `json:"op"`
}

type QueryTwoParams struct {
	Value       interface{} `json:"value"`
	SecondValue interface{} `json:"secondValue"`
	Operator    string      `json:"op"`
}

func (pinata *Pinata) queryPinata(query *string) ([]byte, error) {
	var result []byte = nil
	clientRequest := &http.Client{}

	url := string(QUERYFILES) + "?" + *query

	fmt.Println(url)

	req, errReq := http.NewRequest(string(GET), url, nil)

	if errReq != nil {
		return nil, errReq
	}

	req.Header.Add("Authorization", "Bearer "+pinata.authentication)

	res, errRes := clientRequest.Do(req)

	if errRes != nil {
		return nil, errRes
	}

	defer res.Body.Close()

	result, errReadBody := io.ReadAll(res.Body)

	if errReadBody != nil {
		return nil, errReadBody
	}

	return result, nil
}
