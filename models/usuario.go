package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Usuario struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Nombre          string             `bson:"nombre,omitempty" json:"nombre,omitempty"`
	Apellidos       string             `bson:"apellidos,omitempty" json:"apellidos,omitempty"`
	FechaNacimiento time.Time          `bson:"fechaNacimiento,omitempty" json:"fechaNacimiento,omitempty"`
	Email           string             `bson:"email" json:"email"`
	Password        string             `bson:"password,omitempty" json:"password,omitempty"`
	Avatar          string             `bson:"avatar,omitempty" json:"avatar,omitempty"`
	Banner          string             `bson:"banner,omitempty" json:"banner,omitempty"`
	Biografia       string             `bson:"biografia,omitempty" json:"biografia,omitempty"`
	Ubicacion       string             `bson:"ubicacion,omitempty" json:"ubicacion,omitempty"`
	SitioWeb        string             `bson:"sitioWeb,omitempty" json:"sitioWeb,omitempty"`
}
