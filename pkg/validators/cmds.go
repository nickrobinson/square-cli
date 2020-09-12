package validators

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// ExactArgs is a validator for commands to print an error when the number provided
// is different than the arguments passed in
func ExactArgs(num int) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		argument := "argument"
		if num > 1 {
			argument = "arguments"
		}

		errorMessage := fmt.Sprintf(
			"%s only takes %d %s. See `square %s --help` for supported flags and usage",
			cmd.Name(),
			num,
			argument,
			cmd.Name(),
		)

		if len(args) != num {
			return errors.New(errorMessage)
		}
		return nil
	}
}
