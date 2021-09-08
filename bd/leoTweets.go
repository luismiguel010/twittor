package bd

import (
	"context"
	"github/luismiguel010/twittor/models"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// LeoTweetsSeguidores lee los tweets de mis seguidores
func LeoTweetsSeguidores(ID string, pagina int) ([]models.DevuelvoTweetsSeguidores, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(os.Getenv("MONGO_DBNAME"))
	col := db.Collection("relacion")

	skip := (pagina - 1) * 20

	condiciones := make([]bson.M, 0)
	condiciones = append(condiciones, bson.M{"$match": bson.M{"usuarioid": ID}})
	condiciones = append(condiciones, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "usuariorelacionid",
			"foreignField": "userid",
			"as":           "tweet",
		}})
	condiciones = append(condiciones, bson.M{"$unwind": "$tweet"})
	condiciones = append(condiciones, bson.M{"$sort": bson.M{"fecha": -1}})
	condiciones = append(condiciones, bson.M{"$skip": skip})
	condiciones = append(condiciones, bson.M{"$limit": 20})

	cursor, err := col.Aggregate(ctx, condiciones)
	if err != nil {
		return nil, false
	}
	var result []models.DevuelvoTweetsSeguidores
	err = cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}
	return result, true
}
