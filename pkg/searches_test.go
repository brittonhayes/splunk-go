package splunk

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestICouldSearchSync(t *testing.T) {

	connection, _ := CreateConnectionFromEnvironment()
	require.NotNil(t, connection, "error on searching")

	sessionKey, _ := connection.Login()
	require.NotNil(t, sessionKey, "error on searching")

	searchString := "search index=_internal linecount=1"
	outputMode := "json"
	searchResponse, _ := connection.SearchSync(searchString, outputMode)
	require.NotNil(t, searchResponse, "error on searching")
}

func TestICouldSearchSyncLongQuery(t *testing.T) {

	connection, _ := CreateConnectionFromEnvironment()
	require.NotNil(t, connection, "error on searching")

	sessionKey, _ := connection.Login()
	require.NotNil(t, sessionKey, "error on searching")

	searchString := "search index=_internal|head 100"
	outputMode := "json"
	searchResponse, _ := connection.SearchSync(searchString, outputMode)
	assert.Contains(t, searchResponse, "linecount", "error on searching")
	require.NotNil(t, searchResponse, "error on searching")
}
