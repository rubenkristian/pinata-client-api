package pinata

import (
	"encoding/json"
	"fmt"
	"sync"
)

type Method string

type Url string

const (
	DELETE Method = "DELETE"
	GET    Method = "GET"
	POST   Method = "POST"
	PUT    Method = "PUT"
)

const (
	PINFILE    Url = "https://api.pinata.cloud/pinning/pinFileToIPFS"
	PINJSON    Url = "https://api.pinata.cloud/pinning/pinJSONToIPFS"
	PINBYCID   Url = "https://api.pinata.cloud/pinning/pinByHash"
	LISTPIN    Url = "https://api.pinata.cloud/pinning/pinJobs"
	UPDATEMETA Url = "https://api.pinata.cloud/pinning/hashMetadata"
	UNPIN      Url = "https://api.pinata.cloud/pinning/unpin/"
	DATAUSAGE  Url = "https://api.pinata.cloud/data/userPinnedDataTotal"
	QUERYFILES Url = "https://api.pinata.cloud/data/pinList"
)

type Pinata struct {
	authentication string
	pinataOptions  *PinataOptions
	Loading        bool
}

type PinataOptions struct {
	CidVersion        int8 `json:"cidVersion"`
	WrapWithDirectory bool `json:"wrapWithDirectory"`
}

type PinataMetadata struct {
	Name      string             `json:"name"`
	KeyValues *map[string]string `json:"keyvalues"`
}

type PinataRegion struct {
	CurrentReplicationCount int64  `json:"currentReplicationCount"`
	DesiredReplicationCount int64  `json:"desiredReplicationCount"`
	RegionId                string `json:"regionId"`
}

type PinataRow struct {
	Id           string         `json:"id"`
	IpfsPinHash  string         `json:"ipfs_pin_hash"`
	Size         int64          `json:"size"`
	UserId       string         `json:"user_id"`
	DatePinned   string         `json:"date_pinned"`
	DateUnpinned string         `json:"date_unpinned"`
	MetaData     PinataMetadata `json:"metadata"`
	Regions      []PinataRegion `json:"regions"`
}

type PinataBody struct {
	Count int64       `json:"count"`
	Rows  []PinataRow `json:"rows"`
}

type PinataUsageData struct {
	PinCount                     int `json:"pin_count"`
	PinSizeTotal                 int `json:"pin_size_total"`
	PinSizeWithReplicationsTotal int `json:"pin_size_with_replications_total"`
}

func CreatePinata(auth string, cidVersion int8, wrapWithDirectory bool) *Pinata {
	return &Pinata{
		authentication: auth,
		pinataOptions:  &PinataOptions{CidVersion: cidVersion, WrapWithDirectory: wrapWithDirectory},
		Loading:        false,
	}
}

func (pinata *Pinata) PinFile(fileLoc string, name string, keyvalues *map[string]string) string {
	_, err := pinata.uploadPinFile(fileLoc, name, keyvalues)

	if err != nil {
		return "Error Pin File"
	}

	return "Successful Pin File"
}

func (pinata *Pinata) PinJSON() {

}

func (pinata *Pinata) PinByCID() {

}

func (pinata *Pinata) LastPinByCID() {

}

func (pinata *Pinata) UpdateMetaData() {

}

func (pinata *Pinata) RemoveByHash(hash string) {
	pinata.removeFile(hash)
	fmt.Println("Finished remove file with hash " + hash)
}

func (pinata *Pinata) RemoveFiles(rows []PinataRow) {
	pinata.Loading = true
	var wg sync.WaitGroup
	wg.Add(len(rows))
	for _, row := range rows {
		go func(hash string) {
			defer wg.Done()
			pinata.removeFile(hash)
			fmt.Println("Remove file with hash: " + hash)
		}(row.IpfsPinHash)
	}

	wg.Wait()
	fmt.Println("Finished remove files")
	pinata.Loading = false
}

func (pinata *Pinata) DataUsage() PinataUsageData {
	bodyPinataUsage := &PinataUsageData{}
	body, err := pinata.statUsage()

	if err != nil {
		fmt.Println("Error")
	}

	json.Unmarshal(body, bodyPinataUsage)

	return *bodyPinataUsage
}

func (pinata *Pinata) QueryFiles(query string) PinataBody {
	bodyPinata := &PinataBody{}
	body, err := pinata.queryPinata(&query)

	if err != nil {
		fmt.Println("Error")
	}

	json.Unmarshal(body, bodyPinata)

	return *bodyPinata
}
