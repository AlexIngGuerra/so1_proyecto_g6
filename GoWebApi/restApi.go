package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

type Saludo struct {
	Codigo int    `json:"Codigo"`
	Saludo string `json:"Saludo,omitempty"`
}

var (
	ipdestino = flag.String("ipdest", "192.168.0.8:33000", "La ip del destino de Redis")
)

type RespuestaPais struct {
	Codigo    int    `json:"Codigo"`
	Paises    Paises `json:"Paises"`
	Resultado string `json:"Resultado"`
}

type Paises struct {
	Paises []string `json:"Paises"`
}

type RespuestaFase struct {
	Codigo    int    `json:"Codigo"`
	Fase      Fase   `json: "Fase"`
	Resultado string `json:"Resultado"`
}

type Fase struct {
	Pais    string   `json:"Pais"`
	Predics []Predic `json:"Predics"`
}

type Predic struct {
	Punteo    string `json:"Punteo"`
	Votos     int    `json:"Votos"`
	Resultado string `json:"Resultado"`
}

type Vacio struct {
	Codigo    int    `json:"Codigo"`
	Resultado string `json:"Resultado"`
}

func Hola(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var Jason Saludo
	Jason.Saludo = "Que tals"
	Jason.Codigo = 200
	json.NewEncoder(response).Encode(Jason)
}

func GetPaises(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	redisIp := strings.Split(*ipdestino, "'")[1]
	var ctx = context.Background()
	rbd := redis.NewClient(&redis.Options{
		//Addr:     "172.17.0.2:6379",
		Addr:     redisIp,
		Password: "",
		DB:       0,
	})

	RecPaises, err := rbd.Get(ctx, "Paises").Result()
	if err == redis.Nil {
		fmt.Println("No hay países")
		var Respuesta RespuestaPais
		Respuesta.Codigo = 301
		Respuesta.Resultado = "No existe llave"
		json.NewEncoder(response).Encode(Respuesta)
		return
	} else if err != nil {
		fmt.Println("Hubo un error", err)
		var Respuesta RespuestaPais
		Respuesta.Codigo = 302
		Respuesta.Resultado = "Hubo un error"
		json.NewEncoder(response).Encode(Respuesta)
		return
	} else {
		fmt.Println(RecPaises)
		var Paises Paises
		json.Unmarshal([]byte(RecPaises), &Paises)
		var Respuesta RespuestaPais
		Respuesta.Codigo = 200
		Respuesta.Resultado = "Todo correcto"
		Respuesta.Paises = Paises
		json.NewEncoder(response).Encode(Respuesta)
		return
	}
}

type PaisFase struct {
	Pais string `json:"Pais"`
	Fase int    `json:"Fase"`
}

func GetPaisFase(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	redisIp := strings.Split(*ipdestino, "'")[1]
	var ctx = context.Background()
	rbd := redis.NewClient(&redis.Options{
		//Addr:     "172.17.0.2:6379",
		Addr:     redisIp,
		Password: "",
		DB:       0,
	})
	var entrada PaisFase
	json.NewDecoder((request.Body)).Decode(&entrada)
	Llave := entrada.Pais + "," + string(entrada.Fase)
	Predicciones, err := rbd.Get(ctx, Llave).Result()
	if err == redis.Nil {
		fmt.Println("No hay datos con ese valor")
		var Resultado RespuestaFase
		Resultado.Codigo = 201
		Resultado.Resultado = "No hay datos con esta fase para ese país"
		json.NewEncoder(response).Encode(Resultado)
		return
	} else if err != nil {
		fmt.Println("Error inesperado ", err)
		var Resultado RespuestaFase
		Resultado.Codigo = 301
		Resultado.Resultado = "Error inesperado ha ocurrido"
		json.NewEncoder(response).Encode(Resultado)
		return
	} else {
		fmt.Println(Predicciones)
		var Resultado RespuestaFase
		var Fase Fase
		json.Unmarshal([]byte(Predicciones), &Fase)
		Resultado.Codigo = 200
		Resultado.Resultado = "Estas soon las predicciones"
		Resultado.Fase = Fase
		json.NewEncoder(response).Encode(Resultado)
	}
}

func enableCORS(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
	}).Methods(http.MethodOptions)
	router.Use(middlewareCors)
}

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			// Just put some headers to allow CORS...
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			// and call next handler!
			next.ServeHTTP(w, req)
		})
}

func main() {
	router := mux.NewRouter()
	enableCORS(router)
	fmt.Println("Server on Port ", 8000)

	router.HandleFunc("/Hola", Hola).Methods("GET")
	router.HandleFunc("/GetPaises", GetPaises).Methods("POST")
	router.HandleFunc("/GetPaisFase", GetPaisFase).Methods("POST")
	http.ListenAndServe(":8000", router)
	fmt.Println("Este es el main")
}
