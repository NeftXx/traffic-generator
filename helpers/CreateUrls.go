package helpers

import (
	"strconv"
	"strings"
	"traffic-generator/models"
)

func CreateUrls(baseUrl string, gamenames []models.GameName) []string {
	var result []string
	for _, gamename := range gamenames {
		var url strings.Builder
		url.WriteString(baseUrl)
		url.WriteString("/")
		url.WriteString("game")
		url.WriteString("/")
		url.WriteString(strconv.Itoa(gamename.Id))
		url.WriteString("/gamename/")
		url.WriteString(gamename.Name)
		url.WriteString("/players/")
		result = append(result, url.String())
	}
	return result
}
