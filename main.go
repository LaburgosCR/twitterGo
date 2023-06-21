package main

import (
	//el paquete os me permite leer las variables de entorno
	"context"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/twitterGo/awsgo"
	"github.com/twitterGo/bd"
	"github.com/twitterGo/handlers"
	"github.com/twitterGo/models"
	"github.com/twitterGo/secretmanager"
)

func main() {
	//llamada y el comienzo de la lambda
	lambda.Start(EjecutoLambda)
}

//funcion principal

func EjecutoLambda(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var respuesta *events.APIGatewayProxyResponse
	awsgo.InicializoAWS()

	if !ValidoParametros() {
		respuesta = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error en las variables de entorno deben incluir 'SecretName', 'BucketName', 'UrlPrefix'",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return respuesta, nil
	}

	SecretModel, err := secretmanager.GetSecret(os.Getenv("SecretName"))
	if err != nil {
		respuesta = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error en la lectura de Secret " + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return respuesta, nil
	}

	path := strings.Replace(request.PathParameters["twitterGo"], os.Getenv("UrlPrefix"), "", -1)

	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("method"), path)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("method"), request.HTTPMethod)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("user"), SecretModel.UserName)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("password"), SecretModel.Password)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("host"), SecretModel.Host)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("database"), SecretModel.Database)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("jwtsing"), SecretModel.JWTSing)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("body"), request.Body)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("bucketName"), os.Getenv("BucketName"))

	//chequeo conexion a la bd o conecto a la bd

	err = bd.ConectarBD(awsgo.Ctx)
	if err != nil {
		respuesta = &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error conectando a la base de datos " + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return respuesta, nil
	}

	respAPI := handlers.Manejadores(awsgo.Ctx, request)
	if respAPI.CustomResp == nil {
		respuesta = &events.APIGatewayProxyResponse{
			StatusCode: respAPI.Status,
			Body:       respAPI.Message,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return respuesta, nil
	} else {
		return respAPI.CustomResp, nil
	}
}

func ValidoParametros() bool {
	// es obligatorio que mi lambda reciba 3 variables de entorno sino no funciona

	_, traeParametro := os.LookupEnv("SecretName")
	if !traeParametro {
		return traeParametro
	}
	_, traeParametro = os.LookupEnv("BucketName")
	if !traeParametro {
		return traeParametro
	}
	_, traeParametro = os.LookupEnv("UrlPrefix")
	if !traeParametro {
		return traeParametro
	}

	return traeParametro

}
