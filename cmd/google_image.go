package cmd

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func GoogleImageController(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "잠시만 기다려",
		},
	})

	keyword := i.ApplicationCommandData().Options[0].StringValue()
	limit := i.ApplicationCommandData().Options[1].IntValue()
	googleSearch := exec.Command(os.Getenv("PYTHON_PATH"), "./py/google_image.py", keyword, strconv.FormatInt(limit, 10))
	googleSearchOutput, googleSearchError := googleSearch.Output()

	if googleSearchError != nil {
		return
	}

	arr := strings.Split(string(googleSearchOutput), "\n")

	s.ChannelMessageSend(i.ChannelID, fmt.Sprintf("%s를 검색한 결과 줌 %d를 가져뫘다.", keyword, limit))

	for _, result := range arr {
		_, err := s.ChannelMessageSend(i.ChannelID, result)
		if err != nil {
			return
		}
	}
}
