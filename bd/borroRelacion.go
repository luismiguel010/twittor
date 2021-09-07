package bd

import (
	"context"
	"github/luismiguel010/twittor/models"
	"os"
	"time"
)

// BorroRelacion borra la relacion en la BD
func BorroRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database(os.Getenv("MONGO_DBNAME"))
	col := db.Collection("relacion")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
