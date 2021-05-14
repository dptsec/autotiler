package main

import (
	"log"
	"fmt"
	"context"
	sway "github.com/joshuarubin/go-sway"
)

type SwayHandler struct {
	sway.EventHandler
	Client sway.Client
}

func (h SwayHandler) Window(ctx context.Context, event sway.WindowEvent) {
	if event.Change != "focus" {
		return
	}

	processFocus(ctx, h.Client, event.Container.FocusedNode())
}

func processFocus(ctx context.Context, client sway.Client, node *sway.Node) {
	var newLayout string

	if node == nil || node.AppID == nil {
		return
	}

	if node.Type != "con" {
		return
	}

	if node.Rect.Height > node.Rect.Width {
		newLayout = "splitv"
	} else {
		newLayout = "splith"
	}

	client.RunCommand(ctx, newLayout)
}

func main() {
	ctx := context.Background()

	client, err := sway.New(ctx)
	if err != nil {
		log.Fatal(err)
	}

	sHandler := SwayHandler{
		Client: client,
		EventHandler: sway.NoOpEventHandler(),
	}

	node, err := client.GetTree(ctx)
	if err != nil {
		log.Fatal(err)
	}

	processFocus(ctx, client, node.FocusedNode())
	err = sway.Subscribe(ctx, sHandler, sway.EventTypeWindow)
}
