package outbursts

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/bwmarrin/discordgo"
	"google.golang.org/api/iterator"
)

type OutburstHandler struct {
	outbursts []outburst
	client    *firestore.Client
	ctx       context.Context
}

func (oh *OutburstHandler) Init(client *firestore.Client, ctx context.Context) {
	oh.client = client
	oh.ctx = ctx
}

func (oh *OutburstHandler) Execute(s *discordgo.Session, cid string, key string) {
	for _, burst := range oh.outbursts {
		fmt.Println(key)
		fmt.Println(burst.Key)
		if key == burst.Key {
			burst.fire(s, cid)
		}
	}
}

func (oh *OutburstHandler) LoadOutbursts() {
	// FROM https://stackoverflow.com/questions/61423735/go-firestore-get-all-documents-from-collection
	outburstDocumentIterator := oh.client.Collection("Outbursts").Documents(oh.ctx)
	defer outburstDocumentIterator.Stop()

	for {
		doc, err := outburstDocumentIterator.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("Error when looping through Outburst documents: %s", err)
			return
		}
		var outburstToAdd outburst
		if err := doc.DataTo(&outburstToAdd); err != nil {
			log.Printf("Error when converting outburst document to outburst struct %s", err)
			return
		}
		oh.outbursts = append(oh.outbursts, outburstToAdd)
	}
}

func (oh *OutburstHandler) AddOutburst(newKey string, newMessages []string, newRandomMessages []string) {
	outburstToAdd := outburst{
		Key:            newKey,
		CallCount:      0,
		Messages:       newMessages,
		RandomMessages: newRandomMessages,
	}

	_, _, err := oh.client.Collection("Outbursts").Add(oh.ctx, outburstToAdd)
	if err != nil {
		fmt.Printf("Could not add outburst: %s", err)
	}
	oh.outbursts = append(oh.outbursts, outburstToAdd)
}

func (oh *OutburstHandler) DeleteOutburst() {
	// TODO
}
