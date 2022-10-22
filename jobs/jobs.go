package jobs

import (
	"capture-sre/bart/common"
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type JobSettings struct {
	Protocol       string `json:"protocol"`
	FeedId         int    `json:"feedId"`
	ClientId       int    `json:"clientId"`
	ExternalFeedId string `json:"externalFeedId"`
}
type Job struct {
	JobId       int         `json:"jobId"`
	Type        string      `json:"type"`
	Schedule    string      `json:"schedule"`
	LastRunDate string      `json:"lastRunDate"`
	Enabled     bool        `json:"enabled"`
	Settings    JobSettings `json:"settings"`
}

type GetJobByIdResult struct {
	Success bool `json:"success"`
	Result  Job  `json:"result"`
}

type GetJobQuery struct {
	JobId         string
	Region        string
	Tier          string
	SchedulerType string
}

func GetJobById(getJobQuery GetJobQuery) (Job, error) {

	cert := common.GetCertificateForRegion(getJobQuery.Region)
	schedulerApi := common.GetApiUrlForApp(getJobQuery.SchedulerType, getJobQuery.Region, getJobQuery.Tier)
	url := fmt.Sprintf("%s/jobs/%s", schedulerApi, getJobQuery.JobId)
	client := resty.New()
	client.SetCertificates((cert))
	result := &GetJobByIdResult{}
	_, err := client.R().SetResult(result).Get(url)
	if err != nil {
		error := fmt.Errorf("could not fetch job id %v", err)
		return Job{}, error

	}

	if result.Result.JobId == 0 {
		return Job{}, errors.New("Could not get job id")
	}

	return result.Result, nil

}
