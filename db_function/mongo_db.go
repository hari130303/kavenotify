package dbfunction

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LogUserStruct struct {
	Uid      string    `bson:"_id"`
	UserName string    `bson:"user_name"`
	LogType  string    `bson:"entry_type"`
	LogTime  time.Time `bson:"entry_time"`
}

func UserLogger(userName string, entryType string) {
	Data := bson.M{"user_name": userName, "entry_type": entryType, "entry_time": time.Now()}
	insertManyResult, err := lmdb.InsertOne(context.TODO(), Data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted document with IDs:", insertManyResult.InsertedID)
}

func UserLoginLog(searchTerm string, start int, limit int, orderBy int) (any, error) {

	// ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	// defer cancel()

	var tableData TableData

	searchFilter := bson.M{"user_name": bson.M{"$regex": searchTerm}}

	// Define find options with sorting
	orderColumn := map[int]string{
		1: "user_name",
	}
	if orderBy == 0 || orderBy > len(orderColumn) {
		orderBy = 1
	}

	column, ok := orderColumn[orderBy]
	if !ok {
		column = "user_name"
	}

	findOptions := options.Find().SetSkip(int64(start)).SetLimit(int64(limit)).SetSort(bson.D{{column, 1}}) // Sort by age ascending, then name descending

	cursor, err := lmdb.Find(context.TODO(), searchFilter, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())

	results := []LogUserStruct{}
	if err := cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}

	opts := options.Count().SetHint("_id_")
	count, err := lmdb.CountDocuments(context.TODO(), searchFilter, opts)
	if err != nil {
		panic(err)
	}
	tableData.Count = int(count)
	tableData.Data = results
	tableData.Start = start
	tableData.Limit = limit
	return tableData, nil

}
