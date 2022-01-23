// TODO: 김성모 짤 생성기 https://sungmo.jjong.co.kr/api/download?first=%EA%B9%80%EC%B0%AC%EC%9B%85%EC%9D%B4&second=%ED%95%98%ED%83%9C%EC%9B%85%EC%9D%84%20%EC%A3%BD%EC%97%AC?

package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/tnsgud/go_myammi/cmd"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "검색",
			Description: "검색머를 구글에 검색하며 미미지 림크를 반환합니다.",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "keyword",
					Description: "검색머를 밈력해",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "limit",
					Description: "갯수 밈력해",
					Required:    true,
				},
			},
		},
	}
	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"검색": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			cmd.GoogleImageController(s, i)
		},
	}
)

func main() {
	envErr := godotenv.Load(".env")

	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	discord, err := discordgo.New("Bot " + os.Getenv("TOKEN"))

	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	discord.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	err = discord.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	for _, v := range commands {
		for _, guild := range discord.State.Guilds {
			fmt.Printf("this is guild id : %s\n", guild.ID)
			_, err := discord.ApplicationCommandCreate(discord.State.User.ID, guild.ID, v)
			if err != nil {
				log.Panicf("Cannot create '%v' command: %v", v.Name, err)
			}
		}
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	err = discord.Close()
}
