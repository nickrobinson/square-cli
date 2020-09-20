package ansi

import (
	"io"
	"os"

	"github.com/logrusorgru/aurora"
	"github.com/tidwall/pretty"
	"golang.org/x/crypto/ssh/terminal"
)

//
// Public variables
//

// ForceColors forces the use of colors and other ANSI sequences.
var ForceColors = false

// DisableColors disables all colors and other ANSI sequences.
var DisableColors = false

// EnvironmentOverrideColors overs coloring based on `CLICOLOR` and
// `CLICOLOR_FORCE`. Cf. https://bixense.com/clicolors/
var EnvironmentOverrideColors = true

//
// Public functions
//

// Color returns an aurora.Aurora instance with colors enabled or disabled
// depending on whether the writer supports colors.
func Color(w io.Writer) aurora.Aurora {
	return aurora.NewAurora(shouldUseColors(w))
}

// ColorizeJSON returns a colorized version of the input JSON, if the writer
// supports colors.
func ColorizeJSON(json string, w io.Writer) string {
	if !shouldUseColors(w) {
		return json
	}

	return string(pretty.Color(pretty.Pretty([]byte(json)), nil))
}

//
// Private functions
//
func checkIfTerminal(w io.Writer) bool {
	switch v := w.(type) {
	case *os.File:
		return terminal.IsTerminal(int(v.Fd()))
	default:
		return false
	}
}

func shouldUseColors(w io.Writer) bool {
	useColors := ForceColors || checkIfTerminal(w)

	if EnvironmentOverrideColors {
		force, ok := os.LookupEnv("CLICOLOR_FORCE")

		switch {
		case ok && force != "0":
			useColors = true
		case ok && force == "0":
			useColors = false
		case os.Getenv("CLICOLOR") == "0":
			useColors = false
		}
	}

	return useColors && !DisableColors
}
