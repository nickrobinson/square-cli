package requests

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/nickrobinson/square-cli/internal/ansi"
	"github.com/nickrobinson/square-cli/pkg/config"
	"github.com/nickrobinson/square-cli/pkg/validators"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// RequestParameters captures the structure of the parameters that can be sent to Square
type RequestParameters struct {
	Data        []string
	Idempotency string
	Limit       string
	Version     string
}

func (r *RequestParameters) AppendData(data []string) {
	r.Data = append(r.Data, data...)
}

// Base does stuff
type Base struct {
	Cmd *cobra.Command

	Method string
	Config *config.Config

	Parameters RequestParameters

	// SuppressOutput is used by `trigger` to hide output
	SuppressOutput bool

	APIBaseURL string

	AutoConfirm bool
	ShowHeaders bool
}

var confirmationCommands = map[string]bool{http.MethodDelete: true}

// RunRequestsCmd is the interface exposed for the CLI to run network requests through
func (rb *Base) RunRequestsCmd(cmd *cobra.Command, args []string) error {
	if len(args) > 1 {
		return fmt.Errorf("this command only supports one argument. Run with the --help flag to see usage and examples")
	}

	confirmed, err := rb.confirmCommand()
	if err != nil {
		return err
	} else if !confirmed {
		fmt.Println("Exiting without execution. User did not confirm the command.")
		return nil
	}

	accessToken, err := rb.Config.GetAccessToken()
	if err != nil {
		return err
	}

	path := normalizePath(args[0])

	_, err = rb.MakeRequest(accessToken, path, &rb.Parameters)

	return err
}

// InitFlags initialize shared flags for all requests commands
func (rb *Base) InitFlags() {
	dataUsage := "Data to pass for the API request"
	if rb.Method == http.MethodPut || rb.Method == http.MethodPost {
		dataUsage = "JSON data to pass in API request body"
	}
	if rb.Method == http.MethodPost {
		rb.Cmd.Flags().StringVarP(&rb.Parameters.Idempotency, "idempotency", "i", "", "Sets the idempotency key for your request, preventing replaying the same requests within a 24 hour period")
	}
	rb.Cmd.Flags().StringArrayVarP(&rb.Parameters.Data, "data", "d", []string{}, dataUsage)
	rb.Cmd.Flags().BoolVarP(&rb.ShowHeaders, "show-headers", "s", false, "Show headers on responses to GET, POST, and DELETE requests")
	rb.Cmd.Flags().BoolVarP(&rb.AutoConfirm, "confirm", "c", false, "Automatically confirm the command being entered. WARNING: This will result in NOT being prompted for confirmation for certain commands")
	// Conditionally add flags for GET requests. I'm doing it here to keep `limit`, `start_after` and `ending_before` unexported
	if rb.Method == http.MethodGet {
		rb.Cmd.Flags().StringVarP(&rb.Parameters.Limit, "limit", "l", "", "A limit on the number of objects to be returned, between 1 and 100 (default is 10)")
	}

	rb.Cmd.Flags().StringVarP(&rb.Parameters.Version, "api-version", "v", "", "Square API Version to use for request")

	// Hidden configuration flags, useful for dev/debugging
	rb.Cmd.Flags().StringVar(&rb.APIBaseURL, "base-url", "", "Sets the API base URL")
	rb.Cmd.Flags().MarkHidden("base-url")
}

// MakeRequest will make a request to the Square API with the specific variables given to it
func (rb *Base) MakeRequest(accessToken, path string, params *RequestParameters) ([]byte, error) {
	parsedBaseURL, err := url.Parse(rb.getURL())
	if err != nil {
		log.Error(err)
		return []byte{}, err
	}

	client := Client{
		BaseURL:     parsedBaseURL,
		AccessToken: accessToken,
		Verbose:     rb.ShowHeaders,
	}

	data := ""
	if rb.Method == http.MethodGet || rb.Method == http.MethodDelete {
		data, err = rb.buildDataForRequest(params)
		if err != nil {
			log.Error(err)
			return []byte{}, err
		}
	} else {
		data = params.Data[0]
		if validators.IsValidJSON(data) == false {
			return []byte{}, errors.New("Invalid JSON data provided")
		}
	}

	configureReq := func(req *http.Request) {
		rb.setVersionHeader(req, params)
		rb.setUserAgent(req, params)
	}

	resp, err := client.PerformRequest(rb.Method, path, data, configureReq)
	if err != nil {
		log.Error(err)
		return []byte{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if !rb.SuppressOutput {
		if err != nil {
			log.Error(err)
			return []byte{}, err
		}

		result := ansi.ColorizeJSON(string(body), os.Stdout)
		fmt.Print(result)
	}

	return body, nil
}

// Note: We converted to using two arrays to track keys and values, with our own
// implementation of Go's url.Values Encode function due to our query parameters being
// order sensitive for API requests involving arrays like `items` for `/v1/orders`.
// Go's url.Values uses Go's map, which jumbles the key ordering, and their Encode
// implementation sorts keys by alphabetical order, but this doesn't work for us since
// some API endpoints have required parameter ordering. Yes, this is hacky, but it works.
func (rb *Base) buildDataForRequest(params *RequestParameters) (string, error) {
	keys := []string{}
	values := []string{}

	if len(params.Data) > 0 {
		for _, datum := range params.Data {
			splitDatum := strings.SplitN(datum, "=", 2)

			if len(splitDatum) < 2 {
				return "", fmt.Errorf("Invalid data argument: %s", datum)
			}

			keys = append(keys, splitDatum[0])
			values = append(values, splitDatum[1])
		}
	}

	if rb.Method == http.MethodGet {
		if params.Limit != "" {
			keys = append(keys, "limit")
			values = append(values, params.Limit)
		}
	}

	return encode(keys, values), nil
}

// encode creates a url encoded string with the request parameters
func encode(keys []string, values []string) string {
	var buf strings.Builder
	for i := range keys {
		key := keys[i]
		value := values[i]

		keyEscaped := url.QueryEscape(key)

		// Don't use strict form encoding by changing the square bracket
		// control characters back to their literals. This is fine by the
		// server, and makes these parameter strings easier to read.
		keyEscaped = strings.ReplaceAll(keyEscaped, "%5B", "[")
		keyEscaped = strings.ReplaceAll(keyEscaped, "%5D", "]")

		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(keyEscaped)
		buf.WriteByte('=')
		buf.WriteString(url.QueryEscape(value))
	}
	return buf.String()
}

func (rb *Base) setVersionHeader(request *http.Request, params *RequestParameters) {
	if params.Version != "" {
		request.Header.Set("Square-Version", params.Version)
	}
}

func (rb *Base) setUserAgent(request *http.Request, params *RequestParameters) {
	if rb.Cmd != nil {
		request.Header.Set("User-Agent", fmt.Sprintf("square-cli/%s", rb.Cmd.Root().Version))
	} else {
		request.Header.Set("User-Agent", "square-cli")
	}
}

func (rb *Base) confirmCommand() (bool, error) {
	reader := bufio.NewReader(os.Stdin)
	return rb.getUserConfirmation(reader)
}

func (rb *Base) getUserConfirmation(reader *bufio.Reader) (bool, error) {
	if _, needsConfirmation := confirmationCommands[rb.Method]; needsConfirmation && !rb.AutoConfirm {
		confirmationPrompt := fmt.Sprintf("Are you sure you want to perform the command: %s?\nEnter 'yes' to confirm: ", rb.Method)
		fmt.Print(confirmationPrompt)

		input, err := reader.ReadString('\n')
		if err != nil {
			return false, err
		}

		return strings.Compare(strings.ToLower(input), "yes\n") == 0, nil
	}

	// Always confirm the command if it does not require explicit user confirmation
	return true, nil
}

func normalizePath(path string) string {
	if strings.HasPrefix(path, "/v2/") {
		return path
	}
	if strings.HasPrefix(path, "v2/") {
		return "/" + path
	}
	if strings.HasPrefix(path, "/") {
		return "/v2" + path
	}
	return "/v2/" + path
}

func (rb *Base) getURL() string {
	if rb.APIBaseURL != "" {
		return rb.APIBaseURL
	}

	return rb.Config.GetBaseURL()
}
