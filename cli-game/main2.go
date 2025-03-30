package main

import (
	"LeonFelipeCorder/beer-game/clitool/graph"
	"context"
	"log"
)

func main() {
	ctx := context.Background()
	//client := graph.NewBoardGraphClient()
	//object, err := client.CreateBoard(ctx, "test")
	//object, err := client.CreateBoard(ctx, uuid.New().String())
	//board, err := client.GetBoardByName(ctx, "test")
	//board, err := client.GetBoardByID(ctx, "0824b1d8-8879-48bc-8f86-65cf3644c472")

	client := graph.NewPlayerGraphClient()
	object, err := client.GetPlayerByID(ctx, "1aab5b44-d43a-4aa9-8d08-23bc951c9039")

	if err != nil {
		log.Println("Error creating board:", err)
	}
	log.Println(object)
	//client := graph.NewBoardGraphClient()
	//updatesChan, err := client.SubscribeToBoard(context.Background(), "3cc963d2-4596-4cb3-938f-fd63b69274d7")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//var wg sync.WaitGroup
	//wg.Add(1)
	//go func() {
	//	for message := range updatesChan {
	//		fmt.Println(json.Marshal(message))
	//	}
	//}()
	//wg.Wait()
}
