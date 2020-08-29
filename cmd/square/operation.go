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

func buildOperationCommand(sq *square.Square, name, path, httpVerb string, propFlags map[string]string) *cobra.Command {
	urlParams := extractURLParams(path)
	httpVerb = strings.ToUpper(httpVerb)
	cmd := &cobra.Command{
		Use:         name,
		Annotations: make(map[string]string),
		RunE: func(cmd *cobra.Command, args []string) error {
			client := &http.Client{}
			req, err := http.NewRequest(httpVerb, "https://connect.squareupsandbox.com"+formatURL(path, args), nil)
			req.Header.Add("Authorization", "Bearer "+sq.AccessKey)
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
		},
		Args: cobra.ExactArgs(len(urlParams)),
	}

	for prop := range propFlags {
		flagName := strings.ReplaceAll(prop, "_", "-")
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
