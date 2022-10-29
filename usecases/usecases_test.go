package usecases

import (
	"fmt"
	"testing"
)

func TestGetJobInfoAsSlackString(t *testing.T) {
	jInfo := GetJobInfoAsSlackFormattedString("347", "us-east-1", "production")
	want := fmt.Sprintf("*Job Info:*\n" +
		"\t`Job Id: 1`\n" +
		"\t`Job Type: purge`\n" +
		"\t`Enabled: true`\n" +
		"\t`Last run: 2022-10-24T03:00:00`\n" +
		"\t`Schedule: 0 0 3/7 * * ?`\n" +
		"\t`Client: `\n",
	)

	if jInfo != want {
		t.Errorf("got different result %v than expected %v", jInfo, want)
	}
}
