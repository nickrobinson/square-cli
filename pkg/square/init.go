package square

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

func (s *Square) InitConfig() error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Profile Name (default): ")

	profileName, _ := reader.ReadString('\n')
	profileName = strings.TrimSuffix(profileName, "\n")
	if profileName == "" {
		profileName = "default"
	}
	s.Config.ProfileName = profileName

	fmt.Print("Enter Sandbox Access Token: ")
	accessTokenBytes, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return err
	}
	accessToken := string(accessTokenBytes)
	s.Config.Profile.SandboxAccessToken = accessToken

	fmt.Print("\nEnter Production Access Token: ")
	accessTokenBytes, err = terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return err
	}
	accessToken = string(accessTokenBytes)
	s.Config.Profile.ProductionAccessToken = accessToken

	s.Config.WriteProfile()

	return nil
}
