package dbfunction

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

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

	// searchFilter := bson.M{"user_name": bson.M{"$regex": searchTerm}}
	// // Define find options with sorting

	// if orderBy == 0 {
	// 	orderBy = 1
	// }

	// orderColumn := map[int]string{
	// 	1: "user_name",
	// }
	// findOptions := options.Find().SetLimit(int64(limit)).SetSort(bson.D{{orderColumn[orderBy], 1}}) // Sort by age ascending, then name descending

	return tableData, nil

}
