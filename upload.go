package pinata

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func (pinata *Pinata) uploadPinFile(fileLoc string, name string, keyValues *map[string]string) ([]byte, error) {
	var body []byte = nil
	payload := &bytes.Buffer{}

	writer := multipart.NewWriter(payload)

	file, errFile := os.Open(fileLoc)

	if errFile != nil {
		err := fmt.Errorf("cannot open file, some error equaried %q", errFile.Error())
		return body, err
	}

	defer file.Close()

	copyFile, _ := writer.CreateFormFile("file", filepath.Base(fileLoc))
	_, errFileCopy := io.Copy(copyFile, file)

	if errFileCopy != nil {
		err := fmt.Errorf("some error equaried %q", errFileCopy.Error())
		return body, err
	}

	pinataOptionsJson, errParse := json.Marshal(pinata.pinataOptions)

	if errParse != nil {
		err := fmt.Errorf("some error equaried %q", errParse.Error())
		return body, err
	}

	pinataMetadata := &PinataMetadata{Name: name, KeyValues: keyValues}

	pinataMetadataJson, errParse := json.Marshal(pinataMetadata)

	if errParse != nil {
		err := fmt.Errorf("some error equaried %q", errParse.Error())
		return body, err
	}

	_ = writer.WriteField("pinataOptions", string(pinataOptionsJson))
	_ = writer.WriteField("pinataMetadata", string(pinataMetadataJson))

	errWriter := writer.Close()

	if errWriter != nil {
		err := fmt.Errorf("some error equaried %q", errWriter.Error())
		return body, err
	}

	client := &http.Client{}

	req, errReq := http.NewRequest(string(POST), string(PINFILE), payload)

	if errReq != nil {
		err := fmt.Errorf("some error equaried %q", errReq.Error())
		return body, err
	}

	req.Header.Add("Authorization", "Bearer "+pinata.authentication)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	res, errRes := client.Do(req)

	if errRes != nil {
		err := fmt.Errorf("some error equaried %q", errRes.Error())
		return body, err
	}

	defer res.Body.Close()

	body, errBody := io.ReadAll(res.Body)
	if errBody != nil {
		err := fmt.Errorf("some error equaried %q", errBody.Error())
		return body, err
	}

	return body, nil
}
