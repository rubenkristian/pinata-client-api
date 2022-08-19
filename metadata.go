package pinata

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type PinataUpdateMeta struct {
	IpfsPinHash string             `json:"ipfsPinHash"`
	Name        string             `json:"name"`
	KeyValues   *map[string]string `json:"keyvalues"`
}

func (pinata *Pinata) updateMetadata(cid string, name string, keyValues *map[string]string) ([]byte, error) {
	method := "PUT"

	metaData := &PinataUpdateMeta{
		IpfsPinHash: cid,
		Name:        name,
		KeyValues:   keyValues,
	}

	metaDataString, err := json.MarshalIndent(metaData, "", "\t")

	if err != nil {
		return nil, err
	}
	payload := strings.NewReader(string(metaDataString))

	client := &http.Client{}
	req, err := http.NewRequest(method, string(UPDATEMETA), payload)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+pinata.authentication)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
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
