package cfg

import (
	"encoding/json"

	"Dbot/console"
	"Dbot/src/util/http"
)

var Groups = map[string]string{}

func PopulateGroups(shouldLogJson bool) {
	var target *map[string]string
	response, err := http.RestyClient.R().SetResult(&target).Get("https://api.jxson.tech/get/groups")
	if err != nil {
		console.Warn("Failed to populate groups with error %v. This may cause some issues", err)
	} else if response.IsError() {
		console.Warn("Failed to populate groups with HTTP error %v. This may cause some issues", response.StatusCode())
	}
	Groups = *target
	console.Info("Groups populated with %d children", len(Groups))
	if shouldLogJson {
		console.Info("Logging json...")
		groupsJSON, err := json.Marshal(Groups)
		if err != nil {
			console.Warn("failed to marshal groups with error %v", err)
		} else {
			console.Info(string(groupsJSON))
		}
	}
}
