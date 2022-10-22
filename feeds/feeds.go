package feeds

import (
	"capture-sre/bart/common"
	"fmt"

	"github.com/go-resty/resty/v2"
)

var ProvisioningApi = "https://prva-perf.cc.smarsh.cloud"

var FeedTypes = map[int]string{
	5:  "LinkedIn",
	13: "Slack",
	16: "MSTeams",
	17: "Zoom",
	18: "Salesforce Chatter",
	20: "Facebook Workplace",
	21: "Sharepoint",
}

type Client struct {
	Name     string `json:"name"`
	ClientId int    `json:"clientId"`
}

// structs required
type GetFeedByIdQuery struct {
	FeedId string
	Region string
	Tier   string
}

type Feed struct {
	FeedId     int `json:"feedId"`
	FeedTypeId int `json:"feedTypeId"`
	ClientId   int `json:"clientId"`
	FeedType   string
}

type GetFeedByIdResult struct {
	Success bool `json:"success"`
	Result  Feed `json:"result"`
}

type ClientSearchResult struct {
	Client *Client `json:"result"`
}

// tokens = authorization_status.get('tokens')
type AuthorizationToken struct {
	OauthId         int    `json:"oauthId"`
	TokenExpiryTime string `json:"tokenExpiryTime"`
	UpdateTime      string `json:"updateTime"`
	UserId          string `json:"userId"`
	Active          bool   `json:"active"`
}
type AuthorizationStatus struct {
	FeedId           int                   `json:"feedId"`
	FeedTypeId       int                   `json:"feedTypeId"`
	ExternalId       string                `json:"externalId"`
	Status           string                `json:"status"`
	AuthorizationUrl string                `json:"authorizationUrl"`
	ResourceType     string                `json:"resourceType"`
	Tokens           []*AuthorizationToken `json:"tokens"`
}
type AuthorizationStatuses struct {
	Result []*AuthorizationStatus `json:"result"`
}

// func CheckAll(region string) {

// 	// TODO: function to pick the correct certificate based on region.

// 	result := fetchAllAuthorizedFeedsFromCloudCapture(region)

// 	totalFeeds := 0

// 	for _, authStatus := range result.Result {
// 		if isProvisioned(authStatus) {
// 			feedId := authStatus.FeedId
// 			feedIdAsStr := strconv.Itoa(feedId)
// 			feed := GetFeedById(region, feedIdAsStr)
// 			clientIdAsStr := strconv.Itoa(feed.ClientId)
// 			client := getClientById(region, clientIdAsStr)
// 			clientId := client.ClientId
// 			if !isClientTestClient(clientId) {
// 				totalFeeds++
// 				feedTypeId := authStatus.FeedTypeId
// 				feedType := FeedTypes[feedTypeId]
// 				status := authStatus.Status
// 				resource := authStatus.ResourceType
// 				if len(authStatus.Tokens) > 0 {
// 					for _, t := range authStatus.Tokens {
// 						tokenActive := t.Active
// 						tokenId := t.OauthId
// 						outputString := fmt.Sprintf("Checking client: %s with FeedID : %d of type: %s and Resource: %s => Status: %s. Token => [%d, Active:%v]", client.Name, feed.FeedId, feedType, resource, status, tokenId, tokenActive)
// 						fmt.Println(outputString)

// 					}
// 				}

// 				outputString := fmt.Sprintf("Checking client: %s with FeedID : %d of type: %s and Resource: %s => Status: %s", client.Name, feed.FeedId, feedType, resource, status)
// 				fmt.Println(outputString)

// 			}

// 		}

// 	}
// 	fmt.Println("Total feeds ", totalFeeds)

// }

// func isProvisioned(auth *AuthorizationStatus) bool {
// 	return auth.Status == "successful"
// }

// func isClientTestClient(clientId int) bool {
// 	testClientIds := [14]int{300058, 200285, 200275, 201078, 200070, 200285, 200070, 200289, 200181, 200285, 200181, 201078, 21839}
// 	for _, c := range testClientIds {
// 		if clientId == c {
// 			return true
// 		}
// 	}
// 	return false

// }

// func getClientById(region string, clientId string) *Client {
// 	cert := getCertificateForRegion(region)
// 	provisioningApi := getProvisioningApiUrlForRegion(region)

// 	client := resty.New()
// 	client.SetCertificates((cert))

// 	result := &ClientSearchResult{}
// 	url := fmt.Sprintf("%s/clients/%s", provisioningApi, clientId)

// 	resp, err := client.R().
// 		SetResult(result).
// 		Get(url)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	if resp.StatusCode() != 200 {
// 		log.Fatalln("Could not fetch results from Cloud capture.", resp.StatusCode())

// 	}
// 	return result.Client

// }

func GetFeedById(getFeedByIdQuery GetFeedByIdQuery) (Feed, error) {

	cert := common.GetCertificateForRegion(getFeedByIdQuery.Region)
	provisioningApi := common.GetApiUrlForApp("prva", getFeedByIdQuery.Region, getFeedByIdQuery.Tier)
	client := resty.New()
	client.SetCertificates((cert))
	result := &GetFeedByIdResult{}
	url := fmt.Sprintf("%s/feeds/%s", provisioningApi, getFeedByIdQuery.FeedId)

	_, err := client.R().SetResult(result).Get(url)

	if err != nil {
		error := fmt.Errorf("could not fetch feed id %v", err)
		return Feed{}, error
	}
	return result.Result, nil
}

// func fetchAllAuthorizedFeedsFromCloudCapture(region string) *AuthorizationStatuses {
// 	cert := getCertificateForRegion(region)
// 	provisioningApi := getProvisioningApiUrlForRegion(region)

// 	result := &AuthorizationStatuses{}
// 	client := resty.New()
// 	client.SetCertificates((cert))
// 	url := fmt.Sprintf("%s/provisioning/auth/statuses", provisioningApi)

// 	resp, err := client.R().
// 		SetResult(result).
// 		Get(url)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	if resp.StatusCode() != 200 {
// 		log.Fatalln("Could not fetch results from Cloud capture.", resp.StatusCode())

// 	}
// 	return result
// }
