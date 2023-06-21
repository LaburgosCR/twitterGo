package models

//creamos una estructura
type Secret struct {
	//aqui vamos a agregar los 5 campos que incluimos en nuestro secret
	Host     string `json:"host"` //ALT + 96
	UserName string `json:"username"`
	Password string `json:"password"`
	JWTSing  string `json:"jwtsign"`
	Database string `json:"database"`
}
