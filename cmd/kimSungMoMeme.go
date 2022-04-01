package cmd

import (
	"bufio"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io"
	"net/http"
	"os"
)

func KimSungMoMemeController(s *discordgo.Session, i *discordgo.InteractionCreate) {

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "잠시만 기다려",
		},
	})
	if err != nil {
		return
	}

	first := i.ApplicationCommandData().Options[0].StringValue()
	second := i.ApplicationCommandData().Options[1].StringValue()

	fileUrl := fmt.Sprintf("https://sungmo.jjong.co.kr/api/download?first=%s&second=%s", first, second)
	err = DownloadFile("download.jpg", fileUrl)
	if err != nil {
		fmt.Printf("file download err : %v", err)
		return
	}

	file, err := os.Open("download.jpg")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	s.ChannelMessageSendComplex(i.ChannelID, &discordgo.MessageSend{Files: []*discordgo.File{
		{
			ContentType: "image/jpg",
			Name:        "download.jpg",
			Reader:      bufio.NewReader(file),
		},
	}})
}

func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return err
}
