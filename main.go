package main

import (
	"capture-sre/bart/usecases"
	"log"

	"flag"
	"fmt"
	"os"
)

const (
	RESET   = "\033[0m"
	ERROR   = "\033[31m"
	SUCCESS = "\033[32m"
	WARNING = "\033[33m"
)

func printer(m map[string]string) {
	for k, v := range m {
		fmt.Printf("%s -> %s\n", k, v)
	}
}

func main() {

	getJobCmd := flag.NewFlagSet("get-job", flag.ExitOnError)
	jobId := getJobCmd.String("id", "", "Job ID of the job to search for.")
	region := getJobCmd.String("region", "", "Cloud Capture Region.")
	tier := getJobCmd.String("tier", "", "Cloud Capture Tier.")
	// schedulerType := getJobCmd.String("type", "", "Cloud Capture Scheduler Type. Valid values are shda and shda-spo")

	if len(os.Args) < 2 {
		fmt.Printf("expected 'get-job' or 'help' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "get-job":
		getJobCmd.Parse(os.Args[2:])
		jInfo := executeGetJobCMd(jobId, region, tier)
		printer(jInfo)

	case "get-feed":
		log.Fatal("Not implemented")

	case "help":
		// TODO: handle help for sub-commands
		fmt.Printf("Bart is helpful. Bart knows how to find a lot of useful information about Cloud Capture.\n\n")
		fmt.Printf("%sUsage:%s", WARNING, RESET)
		fmt.Printf("\n\n")
		fmt.Printf("\tbart <command> [arguments]\n\n")
		fmt.Printf("%sCommands:%s\n\n", WARNING, RESET)
		fmt.Printf("\tget-job\t Get details of a scheduler job.\n")

	default:
		fmt.Printf("%sERROR!%s Expected 'get-job' or 'help' subcommands.\n Type 'jarvis help;' for help", ERROR, RESET)
		os.Exit(1)
	}

}

func executeGetJobCMd(jobId *string, region *string, tier *string) map[string]string {
	// s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	// s.Start()
	// fmt.Println("Searching for job with ID:", jobId)
	j := usecases.GetJobInfo(*jobId, *region, *tier)
	return j

	// s.Stop()
	// if err != nil {
	// 	fmt.Printf("%sERROR:%s could not find job with ID: %s", ERROR, RESET, *jobId)
	// } else {
	// 	fid := j.Settings.ExternalFeedId
	// 	fregion := query.Region
	// 	ftier := query.Tier
	// 	fquery := feeds.GetFeedByIdQuery{FeedId: fid, Region: fregion, Tier: ftier}
	// 	f, err := feeds.GetFeedById(fquery)
	// 	if err != nil {
	// 		fmt.Printf("%sERROR:%s could not find feed with ID: %s", ERROR, RESET, fid)
	// 	}

	// 	cid := f.ClientId
	// 	cidAsStr := strconv.Itoa(cid)
	// 	c, err := clients.GetClientById(clients.GetClientByIdQuery{ClientId: cidAsStr, Region: fregion, Tier: ftier})
	// 	fmt.Printf("%sSUCCESS:%s\n", SUCCESS, RESET)
	// 	fmt.Printf("\t job ID:\t%d\n", j.JobId)
	// 	fmt.Printf("\t type:\t\t%s\n", j.Type)
	// 	fmt.Printf("\t schedule:\t%s\n", j.Schedule)
	// 	fmt.Printf("\t last run on:\t%s\n", j.LastRunDate)
	// 	fmt.Printf("\t is enabled:\t%v\n", j.Enabled)
	// 	fmt.Printf("\t client ID:\t%v\n", j.Settings.ClientId)
	// 	fmt.Printf("\t feedID:\t%v\n", j.Settings.FeedId)
	// 	fmt.Printf("\t efeedID:\t%v\n", j.Settings.ExternalFeedId)
	// 	fmt.Printf("\t ClientID:\t%v\n", f.ClientId)
	// 	fmt.Printf("\t Client:\t%v\n", c.Name)

	// client, _ := clients.GetClientById(strconv.Itoa(j.Settings.ClientId), query.Region, query.Tier)
	// feed := feeds.GetFeedById(query.Region, strconv.Itoa(j.Settings.FeedId))
	// fmt.Printf("\t Client:\t%v\n", client.Name)
	// fmt.Printf("\t FeedType:\t%v\n", feed.FeedType)

	// }
}
