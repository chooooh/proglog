package log

import (
	"os"

	"github.com/stretchr/testify/require"
)

func TestIndex(t *testing) {
	f, err := os.CreateTemp("", "index_test")
	require.NoError(t, err)
	defer os.Remove(f.Name())

}
