package pinata

import (
	"io/ioutil"
	"net/http"
)

func (pinata *Pinata) listPinByCid() ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(string(GET), "https://api.pinata.cloud/pinning/pinJobs?status=retrieving&sort=ASC", nil)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+pinata.authentication)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
