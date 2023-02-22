package useDB

import (
	"log"
)

func GetAnswer(roomId string) string {
	ctx, client, _err := connectDB()
	if _err != nil {
		log.Fatalf("failed to connect to database: %v", _err)
	}

	iter, err := client.Collection("Room").Doc(roomId).Get(ctx)
	if err != nil {
		log.Fatalf("error getting Room documents: %v\n", err)
	}
	answer := iter.Data()["Answer"].(string)
	defer client.Close()
	return answer

}