package status

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type SquareStatusResp struct {
	Status SquareStatusField `json:"status"`
}

type SquareStatusField struct {
	Indicator   string `json:"indicator"`
	Description string `json:"description"`
}

const SQUARE_STATUS_URL = "https://www.issquareup.com/api/v2/status.json"

func GetSquareStatus() (string, error) {
	squareStatusClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, SQUARE_STATUS_URL, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", "square-cli")

	res, getErr := squareStatusClient.Do(req)
	if getErr != nil {
		return "", getErr
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return "", readErr
	}

	resp := SquareStatusResp{}
	jsonErr := json.Unmarshal(body, &resp)
	if jsonErr != nil {
		return "", jsonErr
	}

	if resp.Status.Indicator == "none" {
		return "✅ All Services Are Healthy", nil
	} else {
		return fmt.Sprintf("🚨 Issue Found: %s", resp.Status.Description), nil
	}
}
