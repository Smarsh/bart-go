package feeds

import (
	"testing"
)

func TestGetFeedById(t *testing.T) {
	f, err := GetFeedById(GetFeedByIdQuery{"20063", "us-east-1", "production"})
	if err != nil {
		t.Errorf("got an error %v", err)
	}
	wantId := 20063
	wantType := 16
	if f.FeedId != wantId {
		t.Errorf("got feed with ID %v instead of %v", f.FeedId, wantId)
	}

	if f.FeedTypeId != wantType {
		t.Errorf("got feed with Type %v instead of %v", f.FeedTypeId, wantType)
	}

}
