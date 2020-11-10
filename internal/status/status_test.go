package status

import (
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestGetSquareStatus(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Exact URL match
	httpmock.RegisterResponder("GET", "https://www.issquareup.com/api/v2/status.json",
		httpmock.NewStringResponder(200, `{
			"page": {
				"id": "hxwnhyfglktf",
				"name": "Square US",
				"url": "https://www.issquareup.com",
				"time_zone": "America/Los_Angeles",
				"updated_at": "2020-11-09T09:23:39.202-08:00"
			},
			"status": {
				"indicator": "none",
				"description": "All Systems Operational"
			}
		}`))

	status, err := GetSquareStatus()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if status != "✅ All Services Are Healthy" {
		t.Errorf("Unexpected status: got %s", status)
	}

}
