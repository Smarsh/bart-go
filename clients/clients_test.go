package clients

import (
	"testing"
)

func TestGetClientById(t *testing.T) {
	c, err := GetClientById(GetClientByIdQuery{"200082", "us-east-1", "production"})
	if err != nil {
		t.Errorf("got an error %v", err)
	}
	wantId := 200082
	wantName := "MFS Investment Management"
	if c.ClientId != wantId {
		t.Errorf("got client with ID %v instead of %v", c.ClientId, wantId)
	}

	if c.Name != wantName {
		t.Errorf("got client with Name %v instead of %v", c.Name, wantName)
	}

}
