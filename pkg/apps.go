package splunk

import (
	"fmt"
	"net/url"
)

// InstallApp installs splunk apps
func (conn Connection) InstallApp(path string, update bool) (string, error) {
	data := make(url.Values)
	data.Add("name", path)

	updateApp := "false"
	if update == true {
		updateApp = "true"
	}

	data.Add("update", updateApp)
	response, err := conn.httpPost(fmt.Sprintf("%s/services/apps/local", conn.BaseURL), &data)
	return response, err
}

// AppInspect inspects queried app's details and properties.
func (conn Connection) AppInspect(path string) (string, error) {
	data := make(url.Values)
	data.Add("output_mode", "json")
	response, err := conn.httpGet(fmt.Sprintf("%s/services/apps/local/%s", conn.BaseURL, path), &data)
	return response, err
}
