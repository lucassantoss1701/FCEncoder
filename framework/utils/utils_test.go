package utils_test

import (
	"encoder/framework/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsJson(t *testing.T) {
	json := `{
		"id": "3424233-234234234-23423423-42332423",
		"file_path": "convite.mp4",
		"status": "pending"
	}`

	err := utils.IsJson(json)
	require.Nil(t, err)

	json = `oba`

	err = utils.IsJson(json)
	require.Error(t, err)
}
