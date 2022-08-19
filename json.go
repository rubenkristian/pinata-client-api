package pinata

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type PinataPayload struct {
	PinataOptions string         `json:"pinataOptions`
	Metadata      PinataMetadata `json:"pinataMetadata"`
	PinataContent interface{}    `json:"pinataContent"`
}

func uploadJson(url string, auth string, fileLoc string, pinataOptions string, pinataMetaData PinataMetadata, pinataContent interface{}) ([]byte, error) {
	method := "POST"

	payloadPinata := &PinataPayload{
		PinataOptions: pinataOptions,
		Metadata:      pinataMetaData,
		PinataContent: pinataContent,
	}

	payloadPinataString, err := json.MarshalIndent(payloadPinata, "", "\t")

	if err != nil {
		err := fmt.Errorf("cannot open file, some error equaried %q", err.Error())
		return nil, err
	}

	payload := strings.NewReader(string(payloadPinataString))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		err := fmt.Errorf("cannot create new request %q", err.Error())
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+auth)

	res, err := client.Do(req)
	if err != nil {
		err := fmt.Errorf("error request %q", err.Error())
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		err := fmt.Errorf("cannot read body response %q", err.Error())
		return nil, err
	}

	return body, nil
}
