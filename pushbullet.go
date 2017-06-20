package main

import (
	"fmt"
	"os"
        "strconv"
	"log"

	"github.com/mitsuse/pushbullet-go"
	"github.com/mitsuse/pushbullet-go/requests"
)

func sendPush(title, url, price string, total, limit int) {
        //Set token from Config
	pbtoken := config.PushBullet.Token
	log.Println("Sending PushBullet to: " + pbtoken)	

	// Create client for Pushbullet
	pb := pushbullet.New(pbtoken)

	//Create the push content
	pbnote := requests.NewNote()
	pbnote.Title = "NEWEGG-WATCHER | IN STOCK!"
	pbnote.Body = "Url: " + url + "\n\n" +
			"Title: " + title + "\n" +
			"Price: " + price + "\n" +
			"Limit: " + strconv.Itoa(limit) + "\n" +
			"Total: " + strconv.Itoa(total)
			
	//Push It!
	if _, err := pb.PostPushesNote(pbnote); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		return
	}
}
