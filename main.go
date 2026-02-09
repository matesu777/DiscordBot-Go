package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main(){
	godotenv.Load()
	rand.Seed(time.Now().UnixNano())
	TOKEN := os.Getenv("TOKEN")
	dg, err := discordgo.New("Bot " + TOKEN)
	if err != nil{
		fmt.Println("Erro ao criar session com discord: ", err)
		return
	}

	dg.AddHandler(messageCreate)
	dg.Identify.Intents = discordgo.IntentsGuildMessages
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()

}

func readHentai(dir string)(string, error){
	entries, err := os.ReadDir(dir)
	if err != nil {
		return "", err
	}

	var files []string
	for _, entry := range entries{
		files = append(files, entry.Name())
	}

	randomFile := files[rand.Intn(len(files))]

	return dir + "/" + randomFile, nil
}
func uptime()(time.Duration, error){
	data, err := os.ReadFile("/proc/uptime")
	if err != nil{
		return 0, err
	}
	var uptimeSeconds float64
	reader := bytes.NewReader(data)

	if _, err := fmt.Fscan(reader, &uptimeSeconds); err != nil {
		return 0, err
	}

	return time.Duration(uptimeSeconds * float64(time.Second)), nil
}
func covertorUptime(d time.Duration)(string){
	seconds := int(d.Seconds())

	days := seconds / 86400
	seconds %= 86400

	hours := seconds / 3600
	seconds %= 3600

	minutes := seconds / 60
	seconds %= 60

	return fmt.Sprintf("%dd %dh %dm %ds",days, hours, minutes, seconds,)
}


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
		hentai, err := readHentai("Images")
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
	if m.Content == botprefix + "uptime"{
		updatetime, err := uptime()
		
		if err != nil{
			s.ChannelMessageSend(m.ChannelID, "Erro ao pegar uptime!")
            return
		}
		
		_, err = s.ChannelMessageSend(m.ChannelID, covertorUptime(updatetime))
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