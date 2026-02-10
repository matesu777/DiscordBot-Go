package bot

import (
	"go_bot/media"

	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

var botprefix string = "$"

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate){
	if m.Author.ID == s.State.User.ID{
		return
	}
	switch m.Content{
		case botprefix + "hentai":{

			rand.Seed(time.Now().UnixNano())

			hentai, err := media.RandomImage("media/Images")
			file, err := os.Open(hentai)

			
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, "Erro ao abrir a imagem!")
				return
			}
			
			defer file.Close()
			
			fmt.Println("imagem enviada: ", hentai)
			
			_, err = s.ChannelFileSend(m.ChannelID, "hentai.jpg" ,file,)
			
			fmt.Println("Arquivo enviado!")

			if err != nil {
            	fmt.Println(err)
        	}

			return
		}
	}
}