package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Usuario struct {
	ID              primitive.ObjectID `bson:"_id, omitempty" json:"id"`
	Nombre          string             `bson:"nombre" json:"nombre,omnitempty"`
	Apellidos       string             `bson:"apellidos" json:"apellidos,omnitempty"`
	FechaNacimiento time.Time          `bson:"fechaNacimiento" json:"fechaNacimiento,omnitempty"`
	Email           string             `bson:"email" json:"email"`
	Password        string             `bson:"password" json:"password,omnitempty"`
	Avatar          string             `bson:"avatar" json:"avatar,omnitempty"`
	Banner          string             `bson:"banner" json:"banner,omnitempty"`
	Biografia       string             `bson:"biografia" json:"biografia,omnitempty"`
	Ubicacion       string             `bson:"ubicacion" json:"ubicacion,omnitempty"`
	SitioWeb        string             `bson:"sitioWeb" json:"sitioWeb,omnitempty"`
}
