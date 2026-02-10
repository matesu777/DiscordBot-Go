package bot

import (
	"go_bot/media"

	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate){
	var botprefix string = "%"
	var responsesTxt = []string{
		"oiee me chamou? kkkk",
		"a eu não gosto desse cara",
		"esse é foda, admito, até que gosto",
		"oi?",
		"vc tem problema comigo pedro?",
	}
	var comands = []string{
		"bom dia pedro",
		"oq ue vc acha do pedro?",
		"e do mateus?",
		"hentaiii",
		"eu não gosto dele",
		"hentai",
	}
	if m.Author.ID == s.State.User.ID{
		return
	}
	if m.Content == botprefix + comands[5]{
		hentai, err := media.RandomImage("media/Images")
		fmt.Println("imagem enviada: ", hentai)
		if err != nil{
			s.ChannelMessageSend(m.ChannelID, "Erro ao escolher a imagem!")
            return
		}
		file, err := os.Open(hentai)
        if err != nil {
            s.ChannelMessageSend(m.ChannelID, "Erro ao abrir a imagem!")
            return
        }
        defer file.Close()
		_, err = s.ChannelFileSend(
			m.ChannelID,
			"Delicia.jpg",
			file,
		)
		fmt.Println("Arquivo enviado!")
		if err != nil {
            fmt.Println(err)
        }
        return
	}
	
	mention1 := "<@" + s.State.User.ID + ">"
	if m.Content == mention1  {
		fmt.Println("Bot mencionado")
		s.ChannelMessageSend(m.ChannelID, "fala buceta desgraçada fudida fala ai porra")
	}
	if m.Content == comands[0] {
		fmt.Println(responsesTxt[0])
		s.ChannelMessageSend(m.ChannelID, responsesTxt[0])
	}
	if m.Content == comands[1] {
		fmt.Println(responsesTxt[1])
		s.ChannelMessageSend(m.ChannelID, responsesTxt[1])
	}
	if m.Content == comands[2] {
		fmt.Println(responsesTxt[2])
		s.ChannelMessageSend(m.ChannelID, responsesTxt[2])
	}
	if m.Content == comands[3] {
		fmt.Println(responsesTxt[3])
		s.ChannelMessageSend(m.ChannelID, responsesTxt[3])
	}
	if m.Content == comands[4] {
		fmt.Println(responsesTxt[4])
		s.ChannelMessageSend(m.ChannelID, responsesTxt[4])
	}
}