package pinata

import (
	"fmt"
	"io"
	"net/http"
)

func (pinata *Pinata) statUsage() ([]byte, error) {
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, string(DATAUSAGE), nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+pinata.authentication)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return body, nil
}
