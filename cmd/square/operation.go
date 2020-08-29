package main

import (
	"encoding/json"
	"fmt"
	"github.com/nickrobinson/square-cli/pkg/square"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

type OperationCmd struct {
	Name      string
	HTTPVerb  string
	Path      string
	URLParams []string

	stringFlags map[string]*string

	data []string
}

func (oc *OperationCmd) runOperationCmd(cmd *cobra.Command, args []string) error {
	client := &http.Client{}
	req, err := http.NewRequest(oc.HTTPVerb, "https://connect.squareupsandbox.com"+formatURL(oc.Path, args), nil)
	req.Header.Add("Authorization", "Bearer EAAAEF6-EJlxXrA9ifXKL0mswhcBue62xUiwQJx77KWSmlzQ6u5IMvY8fML8yh0L")
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	var data map[string]interface{}
	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &data)
	fmt.Println(data)
	return nil
}

func buildOperationCommand(sq *square.Square, name, path, httpVerb string, propFlags map[string]string) *cobra.Command {
	urlParams := extractURLParams(path)
	httpVerb = strings.ToUpper(httpVerb)
	operationCmd := &OperationCmd{
		Name:      name,
		HTTPVerb:  httpVerb,
		Path:      path,
		URLParams: urlParams,

		stringFlags: make(map[string]*string),
	}
	cmd := &cobra.Command{
		Use:         name,
		Annotations: make(map[string]string),
		RunE:        operationCmd.runOperationCmd,
		Args:        cobra.ExactArgs(len(urlParams)),
	}

	for prop := range propFlags {
		flagName := strings.ReplaceAll(prop, "_", "-")
		operationCmd.stringFlags[flagName] = cmd.Flags().String(flagName, "", "")
		cmd.Flags().SetAnnotation(flagName, "request", []string{"true"})
	}

	return cmd
}

//
// Private functions
//

func extractURLParams(path string) []string {
	re := regexp.MustCompile(`{\w+}`)
	return re.FindAllString(path, -1)
}

func formatURL(path string, urlParams []string) string {
	s := make([]interface{}, len(urlParams))
	for i, v := range urlParams {
		s[i] = v
	}

	re := regexp.MustCompile(`{\w+}`)
	format := re.ReplaceAllString(path, "%s")

	return fmt.Sprintf(format, s...)
}
