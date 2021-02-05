package square

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh/terminal"
)

// InitConfig will add user provided data to
// new/existing Square CLI config
func (s *Square) InitConfig() error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Profile Name (default): ")

	profileName, _ := reader.ReadString('\n')
	profileName = strings.TrimSuffix(profileName, "\n")
	if profileName == "" {
		profileName = "default"
	}

	fmt.Print("Enter Sandbox Access Token: ")
	accessTokenBytes, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return err
	}
	sandboxAccessToken := string(accessTokenBytes)
	s.Config.SandboxAccessToken = sandboxAccessToken

	fmt.Print("\nEnter Production Access Token: ")
	accessTokenBytes, err = terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return err
	}
	productionAccessToken := string(accessTokenBytes)
	s.Config.ProductionAccessToken = productionAccessToken

	cfg := map[string]interface{}{
		"sandbox_access_token":    sandboxAccessToken,
		"production_access_token": productionAccessToken,
	}
	viper.Set(profileName, cfg)
	return viper.WriteConfig()
}
