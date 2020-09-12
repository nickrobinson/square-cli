package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

type initCmd struct {
	Cmd *cobra.Command
}

func newInitCmd() *initCmd {
	ic := &initCmd{}

	ic.Cmd = &cobra.Command{
		Use:   "init",
		Args:  cobra.NoArgs,
		Short: "Initialize Square CLI config.",
		Long:  `Initialize Square CLI configuration file.`,

		RunE: ic.runInitCmd,
	}

	return ic
}

func (ic *initCmd) runInitCmd(cmd *cobra.Command, args []string) error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Profile Name (default): ")

	profileName, _ := reader.ReadString('\n')
	profileName = strings.TrimSuffix(profileName, "\n")
	if profileName == "" {
		profileName = "default"
	}

	fmt.Print("Enter Access Token: ")
	accessTokenBytes, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return err
	}
	accessToken := string(accessTokenBytes)
	Profile.ProfileName = profileName
	Profile.AccessToken = accessToken

	profileErr := Profile.CreateProfile()
	if profileErr != nil {
		return profileErr
	}

	return nil
}
