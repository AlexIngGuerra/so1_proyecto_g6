package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"strings"
	"github.com/go-redis/redis/v9"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "grcp/proyecto/protoc"
)

var (
	port      = flag.Int("port", 50051, "Puerto del servidor--")
	ipdestino = flag.String("ipdest", "192.168.0.8:33000", "La ip del destino de Redis")
)

type Paises struct {
	Paises []string `json:"Paises"`
}

type Fase struct {
	Pais    string   `json:"Pais"`
	Predics []Predic `json:"Predics"`
}

type Predic struct {
	Punteo string `json:"Punteo"`
	Votos  int    `json:"Votos"`
}

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) IngresoDatos(ctx context.Context, in *pb.IngresoSolicitud) (*pb.Respuesta, error) {
	log.Printf("Se registra información. \nteam1: %v\nteam2: %v\nscore: %v\nphase: %v", in.Team1, in.Team2, in.Score, in.Phase)
	Almacenar(in.Team1, in.Team2, in.Score, in.Phase)
	return &pb.Respuesta{Codigo: "200", Mensaje: "Se registró información"}, nil

}

func Almacenar(T1 string, T2 string, Scr string, Ph string) {
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
		//No hay datos registrados
		fmt.Println("No hay paises")
		JsonPaises := `{"Paises":["` + T1 + "-" + T2 + `"]}`
		var paisanos Paises
		paisanos.Paises = append(paisanos.Paises, T1+"-"+T2)
		err2 := rbd.Set(ctx, "Paises", JsonPaises, 0).Err()
		if err2 != nil {
			panic(err2)
		}
		//Registramos la nueva fase
		var unaFase Fase
		//Fase.Pais = "Pais-Pais,Fase
		unaFase.Pais = T1 + "-" + T2 + "," + Ph
		var Prediccion Predic
		Prediccion.Punteo = Scr
		Prediccion.Votos = 1
		unaFase.Predics = append(unaFase.Predics, Prediccion)
		SetFase, err := json.Marshal(unaFase)
		if err != nil {
			panic(err)
		}
		err3 := rbd.Set(ctx, unaFase.Pais, SetFase, 0).Err()
		if err3 != nil {
			panic(err3)
		}
		return
	} else if err != nil {
		panic(err)
	} else {
		Partido := T1 + "-" + T2
		Partidoaux := T2 + "-" + T1
		fmt.Println("key2: ", RecPaises)
		var Paises Paises
		json.Unmarshal([]byte(RecPaises), &Paises)
		fmt.Printf("Partidos registrados: %+v", Paises.Paises)
		for _, pais := range Paises.Paises {
			fmt.Println(pais)
			if Partido == pais {
				//Se encontró
				//Se encontró pero orden invertido
				//Pais-pais,fase
				Llave := Partido + "," + Ph
				//Obtenemos data en JSON {[]}
				Predicciones, err := rbd.Get(ctx, Llave).Result()
				//No existe registro de esa fase
				if err == redis.Nil {
					//Registramos la nueva fase
					var unaFase Fase
					//Fase.Pais = "Pais-Pais,Fase
					unaFase.Pais = Llave
					var Prediccion Predic
					Prediccion.Punteo = Scr
					Prediccion.Votos = 1
					unaFase.Predics = append(unaFase.Predics, Prediccion)
					SetFase, err := json.Marshal(unaFase)
					if err != nil {
						panic(err)
					}
					err3 := rbd.Set(ctx, unaFase.Pais, SetFase, 0).Err()
					if err3 != nil {
						panic(err3)
					}
					return
				}
				if err != nil {
					panic(err)
				}
				var unaFase Fase
				json.Unmarshal([]byte(Predicciones), &unaFase)
				for i, pred := range unaFase.Predics {
					if Scr == pred.Punteo {
						//Existe ya una predicción identica
						//Actualizamos +1 los votos
						unaFase.Predics[i].Votos = pred.Votos + 1
						setFase, err := json.Marshal(unaFase)
						if err != nil {
							panic(err)
						}
						//Guardamos el cambio
						err3 := rbd.Set(ctx, Llave, setFase, 0).Err()
						if err3 != nil {
							panic(err3)
						}
						return
					}
				}
				//Se crea la nueva predicción
				var unaPred Predic
				unaPred.Punteo = Scr
				unaPred.Votos = 1
				//Se añade a la lista
				unaFase.Predics = append(unaFase.Predics, unaPred)
				//pais-pais,fase
				setFase, err := json.Marshal(unaFase)
				if err != nil {
					panic(err)
				}
				//Se registra el cambio
				err3 := rbd.Set(ctx, Llave, setFase, 0).Err()
				if err3 != nil {
					panic(err3)
				}
				return
			} else if Partidoaux == pais {
				//Se encontró pero orden invertido
				//Pais-pais,fase
				Llave := Partidoaux + "," + Ph
				//Obtenemos data en JSON {[]}
				Predicciones, err := rbd.Get(ctx, Llave).Result()

				//No existe registro de esa fase
				if err == redis.Nil {
					//Registramos la nueva fase
					var unaFase Fase
					//Fase.Pais = "Pais-Pais,Fase
					unaFase.Pais = Llave
					var Prediccion Predic
					Prediccion.Punteo = Scr
					Prediccion.Votos = 1
					unaFase.Predics = append(unaFase.Predics, Prediccion)
					SetFase, err := json.Marshal(unaFase)
					if err != nil {
						panic(err)
					}
					err3 := rbd.Set(ctx, unaFase.Pais, SetFase, 0).Err()
					if err3 != nil {
						panic(err3)
					}
					return
				}
				if err != nil {
					panic(err)
				}
				var unaFase Fase
				json.Unmarshal([]byte(Predicciones), &unaFase)
				for i, pred := range unaFase.Predics {
					if Scr == pred.Punteo {
						//Existe ya una predicción identica
						//Actualizamos +1 los votos
						unaFase.Predics[i].Votos = pred.Votos + 1
						setFase, err := json.Marshal(unaFase)
						if err != nil {
							panic(err)
						}
						//Guardamos el cambio
						err3 := rbd.Set(ctx, Llave, setFase, 0).Err()
						if err3 != nil {
							panic(err3)
						}
						return
					}
				}
				//Se crea la nueva predicción
				var unaPred Predic
				unaPred.Punteo = Scr
				unaPred.Votos = 1
				//Se añade a la lista
				unaFase.Predics = append(unaFase.Predics, unaPred)
				//pais-pais,fase
				setFase, err := json.Marshal(unaFase)
				if err != nil {
					panic(err)
				}
				//Se registra el cambio
				err3 := rbd.Set(ctx, Llave, setFase, 0).Err()
				if err3 != nil {
					panic(err3)
				}
				return
			}
		}
		//No encontró
		Paises.Paises = append(Paises.Paises, Partido)

		SetPaises, err := json.Marshal(Paises)
		if err != nil {
			fmt.Println(err)
			panic("Falló convertir a JSON ")
			return
		}
		//Agregamos el país
		err2 := rbd.Set(ctx, "Paises", SetPaises, 0).Err()
		if err2 != nil {
			panic(err2)
		}
		var unaFase Fase
		//Fase.Pais = "Pais-Pais,Fase
		unaFase.Pais = Partido + "," + Ph
		var Prediccion Predic
		Prediccion.Punteo = Scr
		Prediccion.Votos = 1
		unaFase.Predics = append(unaFase.Predics, Prediccion)
		SetFase, err := json.Marshal(unaFase)
		if err != nil {
			panic(err)
		}
		err3 := rbd.Set(ctx, unaFase.Pais, SetFase, 0).Err()
		if err3 != nil {
			panic(err3)
		}
		return
	}
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("la ip destino: ", *ipdestino)

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
