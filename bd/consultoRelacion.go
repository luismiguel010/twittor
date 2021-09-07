package bd

import (
	"context"
	"fmt"
	"github/luismiguel010/twittor/models"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// ConsultoRelacion consulta la relación entre 2 usuarios
func ConsultoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database(os.Getenv("MONGO_DBNAME"))
	col := db.Collection("relacion")

	condicion := bson.M{
		"usuarioid":         t.UsuarioID,
		"usuariorelacionid": t.UsuarioRelacionID,
	}

	var resultado models.Relacion
	fmt.Println(resultado)
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	return true, nil
}
