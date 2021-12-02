package workspace

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestScanner_ToIntArray(t *testing.T) {
	t.Skip("meant to be run locally")

	scanner := NewScanner()
	arr := scanner.ToIntArray("<file-path>")
	require.NotNil(t, arr)
}