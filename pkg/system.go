package splunk

import (
	"fmt"
	"net/url"
)

// RestartServer performs a reboot operation Splunk
func (conn Connection) RestartServer() (string, error) {
	response, err := conn.httpPost(fmt.Sprintf("%s/services/server/control/restart", conn.BaseURL), nil)
	return response, err
}

// InspectControl Lists actions that can be performed at this endpoint.
func (conn Connection) InspectControl(endpoint string) (string, error) {
	data := make(url.Values)
	data.Add("output_mode", "json")
	response, err := conn.httpGet(fmt.Sprintf("%s/services/server/control/%s", conn.BaseURL, endpoint), &data)
	return response, err
}
