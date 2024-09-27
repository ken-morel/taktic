package taktik


import (
	"fmt"
	"log"
	"os"

	"context"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/sbook-cm/backend/backend"
)

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func GetDBClient() *mongo.Client {
	var url := os.Getenv("MONGO_URL")
	fmt.Println("connecting to " + url)
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(url).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func _main() {
	client := connectDatabase()
	backend.SetDatabase(client.Database("test"))
	defer client.Disconnect(context.TODO())

	http.Handle("/", backend.Route())

	log.Println("Server is running on http://localhost:1236")
	log.Fatal(http.ListenAndServe(":1236", nil))
}
