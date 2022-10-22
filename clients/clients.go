package clients

import (
	"capture-sre/bart/common"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	ClientId       int    `json:"clientId"`
	Name           string `json:"name"`
	AccountId      string `json:"accountId"`
	Status         string `json:"status"`
	San            string `json:"san"`
	EncryptionMode string `json:"encryptionMode"`
}

type GetClientByIdResponse struct {
	IsSuccess bool   `json:"success"`
	Client    Client `json:"result"`
}

type GetClientByIdQuery struct {
	ClientId string
	Region   string
	Tier     string
}

func GetClientById(getClientByIdQuery GetClientByIdQuery) (Client, error) {

	cert := common.GetCertificateForRegion(getClientByIdQuery.Region)
	api := common.GetApiUrlForApp("prva", getClientByIdQuery.Region, getClientByIdQuery.Tier)
	url := fmt.Sprintf("%s/clients/%s", api, getClientByIdQuery.ClientId)
	client := resty.New()
	client.SetCertificates((cert))
	result := &GetClientByIdResponse{}
	_, err := client.R().SetResult(result).Get(url)
	if err != nil {
		error := fmt.Errorf("could not fetch client id %s. error=%v", getClientByIdQuery.ClientId, err)
		return Client{}, error

	}

	return result.Client, nil

}
