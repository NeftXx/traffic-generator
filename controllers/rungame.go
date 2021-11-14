package controllers

import (
	"fmt"
	"strconv"
	"strings"
	"traffic-generator/helpers"
	"traffic-generator/models"
)

func Rungame(urlNipIo string, gamename string, players int, rungames int, concurrence int, timeout int) {
	var gamenames []models.GameName
	gamenamesStr := strings.Split(gamename, "|")
	for _, gamename := range gamenamesStr {
		gamename = strings.TrimSpace(gamename)
		gameArr := strings.Split(gamename, ";")
		idStr := gameArr[0]
		name := gameArr[1]
		id, _ := strconv.Atoi(idStr)
		game := models.GameName{Id: id, Name: name}
		gamenames = append(gamenames, game)
	}
	urls := helpers.CreateUrls(urlNipIo, gamenames)
	length := len(urls)
	var requestPerThread int = rungames / concurrence

	channels := make(chan string, concurrence)
	for i := 0; i < concurrence; i++ {
		go makePostRequest(urls, length, players, requestPerThread, i+1, channels)
	}

	count := 0
	for elem := range channels {
		if count == concurrence-1 {
			close(channels)
		}
		count++
		fmt.Println(elem)
	}
}

func makePostRequest(urls []string, maxUrls int, players int, requestPerThread int, noThread int, channels chan string) {
	for i := 0; i < requestPerThread; i++ {
		positionRandom := helpers.RandomNumber(0, maxUrls-1)
		var url strings.Builder
		url.WriteString(urls[positionRandom])
		url.WriteString(strconv.Itoa(helpers.RandomNumber(1, players)))
	}
	chain := fmt.Sprintf("Thread: %d", noThread)
	channels <- chain
}
