package outbursts

import (
	"context"
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
		if key == burst.getKey() {
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
			log.Println("Error when looping through Outburst documents. OutburstHandler slice empty.")
			return
		}
		var outburstToAdd outburst
		if err := doc.DataTo(&outburstToAdd); err != nil {
			log.Println("Error when converting outburst document to outburst struct. Several outbursts may not be loaded in.")
			return
		}
		oh.outbursts = append(oh.outbursts, outburstToAdd)
	}
}

func (oh *OutburstHandler) AddOutburst(newKey string, newMessages []string, newRandomMessages []string) {
	outburstCollectionRef := oh.client.Collection("Outbursts")
	outburstToAdd := outburst{
		key:            newKey,
		callCount:      0,
		messages:       newMessages,
		randomMessages: newRandomMessages,
	}
	_, _, err := outburstCollectionRef.Add(oh.ctx, outburstToAdd)
	if err != nil {
		log.Printf("An error occurred when adding this outburst: %s", err)
		return
	}
	oh.outbursts = append(oh.outbursts, outburstToAdd)
}

func (oh *OutburstHandler) DeleteOutburst() {
	// TODO
}
