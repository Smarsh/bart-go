package usecases

import (
	"capture-sre/bart/jobs"
	"fmt"
	"strconv"
)

// func GetJobInfo(jobId string, region string, tier string) map[string]string {
// 	m := make(map[string]string)

// 	j, err := jobs.GetJobById(jobs.GetJobQuery{JobId: jobId, Region: region, Tier: tier, SchedulerType: "shda"})
// 	if err != nil {
// 		return m

// 	}

// 	c, _ := clients.GetClientById(clients.GetClientByIdQuery{ClientId: strconv.Itoa(j.Settings.ClientId), Region: region, Tier: tier})

// 	m["Job ID"] = strconv.Itoa(j.JobId)
// 	m["Type"] = j.Type
// 	m["Enabled"] = strconv.FormatBool(j.Enabled)
// 	m["Last run"] = j.LastRunDate
// 	m["Schedule"] = j.Schedule
// 	m["External Feed ID"] = j.Settings.ExternalFeedId
// 	m["Feed ID"] = strconv.Itoa(j.Settings.FeedId)
// 	m["External Feed ID"] = j.Settings.ExternalFeedId
// 	m["Client ID"] = strconv.Itoa(j.Settings.ClientId)
// 	m["Client Name"] = c.Name

// 	return m

// }

func GetJobInfoAsSlackFormattedString(jobId string, region string, tier string) string {
	fmsgAdditional := ""
	j, err := jobs.GetJobById(jobs.GetJobQuery{JobId: jobId, Region: region, Tier: tier, SchedulerType: "shda"})
	if err != nil {

		return fmt.Sprintf("no job found.error=%s", err)

	}

	// c, err := clients.GetClientById(clients.GetClientByIdQuery{ClientId: j.Settings["clientId"], Region: region, Tier: tier})
	if err != nil {

		return fmt.Sprintf("no client found.error=%s", err)

	}

	fmsg := fmt.Sprintf("*Job Info:*\n"+
		"\t`Job Id: %s`\n"+
		"\t`Job Type: %s`\n"+
		"\t`Enabled: %s`\n"+
		"\t`Last run: %s`\n"+
		"\t`Schedule: %s`\n"+
		"\t`Additional: %s`\n",
		strconv.Itoa(j.JobId),
		j.Type,
		strconv.FormatBool(j.Enabled),
		j.LastRunDate,
		j.Schedule,
		fmsgAdditional,
	)

	if j.Type == "purge" {
		settings := jobs.GetPurgeJobSettings(j)
		fmsgAdditional := fmt.Sprintf("%s", settings.JobType)
		fmsg = fmsg + fmsgAdditional

	}

	if j.Type == "export-recon" {
		settings := jobs.GetExportReconJobSettings(j)
		fmsgAdditional := fmt.Sprintf("*Job Info:*\n"+
			"\t`Application Id: %s`\n"+
			"\t`Client Id: %s`\n"+
			"\t`Feed Id: %s`\n"+
			"\t`Destination Email Address: %s`\n",
			strconv.Itoa(settings.ApplicationId),
			settings.ClientId,
			settings.FeedId,
			settings.DestinationEmailAddress[:],
		)

		fmsg = fmsg + fmsgAdditional

	}

	return fmsg

}
