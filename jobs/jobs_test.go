package jobs

import (
	"testing"
)

func TestGetJobById(t *testing.T) {
	j, err := GetJobById(GetJobQuery{SchedulerType: "shda", Region: "us-east-1", Tier: "production", JobId: "1"})
	if err != nil {
		t.Errorf("got an error %v", err)
	}
	wantId := 1
	wantType := "purge"
	if j.JobId != wantId {
		t.Errorf("got job with ID %v instead of %v", j.JobId, wantId)
	}

	if j.Type != wantType {
		t.Errorf("got job with Type %v instead of %v", j.Type, wantType)
	}

}
