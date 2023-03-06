package readDB

import (
	connectDB "gamename-back-end/pkg/connect_db"
	"log"
)

func GetTheme(roomId string) string {
	ctx, client, _err := connectDB.ConnectDB()
	if _err != nil {
		log.Printf("An error has occurred: %s", _err)
	}

	iter, err := client.Collection("Room").Doc(roomId).Get(ctx)
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}
	theme := iter.Data()["Theme"].(string)
	defer client.Close()
	return theme

}
