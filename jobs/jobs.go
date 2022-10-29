package jobs

import (
	"capture-sre/bart/common"
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/mitchellh/mapstructure"
)

type PurgeJobSettings struct {
	RetentionDays int
	JobType       string
}

type ExportReconJobSettings struct {
	EventType               string
	ApplicationId           int
	DestinationEmailAddress []string
	Timezone                string
	TemplateId              string
	StartDate               string
	Enddate                 string
	SenderEmailAddress      string
	ClientId                string
	FeedId                  string
}
type JobSettings struct {
	Protocol       string `json:"protocol"`
	FeedId         int    `json:"feedId"`
	ClientId       int    `json:"clientId"`
	ExternalFeedId string `json:"externalFeedId"`
}
type Job struct {
	JobId       int                    `json:"jobId"`
	Type        string                 `json:"type"`
	Schedule    string                 `json:"schedule"`
	LastRunDate string                 `json:"lastRunDate"`
	Enabled     bool                   `json:"enabled"`
	Settings    map[string]interface{} `json:"settings"`
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
	resp, err := client.R().SetResult(result).Get(url)
	fmt.Println(resp)
	if err != nil {
		error := fmt.Errorf("could not fetch job id %s", err)
		return Job{}, error

	}

	if result.Result.JobId == 0 {
		return Job{}, errors.New("Could not get job id")
	}

	return result.Result, nil

}

func GetPurgeJobSettings(j Job) PurgeJobSettings {
	purgeJobSettings := &PurgeJobSettings{}
	mapstructure.Decode(j.Settings, &purgeJobSettings)
	return *purgeJobSettings
}

func GetExportReconJobSettings(j Job) ExportReconJobSettings {
	exportReconJobSettings := &ExportReconJobSettings{}
	mapstructure.Decode(j.Settings, &exportReconJobSettings)
	return *exportReconJobSettings
}
