package validators

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestExactArgs(t *testing.T) {
	c := &cobra.Command{Use: "c"}
	args := []string{"foo"}

	result := ExactArgs(1)(c, args)
	assert.Nil(t, result)
}

func TestExactArgsTooMany(t *testing.T) {
	c := &cobra.Command{Use: "c"}
	args := []string{"foo", "bar"}

	result := ExactArgs(1)(c, args)
	assert.EqualError(t, result, "c only takes 1 argument. See `square c --help` for supported flags and usage")
}

func TestExactArgsTooManyMoreThan1(t *testing.T) {
	c := &cobra.Command{Use: "c"}
	args := []string{"foo", "bar", "baz"}

	result := ExactArgs(2)(c, args)
	assert.EqualError(t, result, "c only takes 2 arguments. See `square c --help` for supported flags and usage")
}
