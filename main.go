package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	//"github.com/gorilla/handlers"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var candidatos []CANDIDATE
var votantes []VOTANTE

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

//Para el candidato
func GetCandidatos(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	json.NewEncoder(w).Encode(candidatos)

}

//CREATE
func CreateCandidato(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var NewCandidato CANDIDATE
	reqbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Invalid Data")
	}

	json.Unmarshal(reqbody, &NewCandidato)

	NewCandidato.ID = len(candidatos) + 1
	candidatos = append(candidatos, NewCandidato)

	json.NewEncoder(w).Encode(candidatos)
	fmt.Fprintf(w, "Candidato added")

}

//READ
func ReadCandidato(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	params := mux.Vars(r)
	CandidatoID, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	for _, _candidato := range candidatos {
		if _candidato.ID == CandidatoID {
			json.NewEncoder(w).Encode(_candidato)
		}
	}
}

//UPDATE
func UpdateCandidato(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	params := mux.Vars(r)
	CandidatoID, err := strconv.Atoi(params["id"])
	var updateCandidato CANDIDATE
	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Invalid Data")

	}
	json.Unmarshal(reqBody, &updateCandidato)
	for i, candidato := range candidatos {
		if candidato.ID == CandidatoID {
			candidatos = append(candidatos[:i], candidatos[i+1:]...)
			updateCandidato.ID = CandidatoID
			candidatos = append(candidatos, updateCandidato)

			fmt.Fprintf(w, "votante Updated %v", CandidatoID)
		}
	}

}

//DELETED
func DeletedCandidato(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	params := mux.Vars(r)
	CandidatoID, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	for index, _candidato := range candidatos {
		if _candidato.ID == CandidatoID {
			candidatos = append(candidatos[:index], candidatos[index+1:]...)
			fmt.Fprintf(w, "votante fue elimano:  %v", CandidatoID)
			break
		}
	}

}

//Para el votante
//Para el candidato
func GetVotantes(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	json.NewEncoder(w).Encode(votantes)
}

//CREATE
func CreateVotante(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var NewVotante VOTANTE
	reqbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Invalid Data")
	}

	json.Unmarshal(reqbody, &NewVotante)

	NewVotante.ID = len(votantes) + 1
	votantes = append(votantes, NewVotante)

	json.NewEncoder(w).Encode(votantes)
	fmt.Fprintf(w, "Candidato added")

}

//READ
func ReadVotante(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	params := mux.Vars(r)
	votanteID, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	for _, _votante := range votantes {
		if _votante.ID == votanteID {
			json.NewEncoder(w).Encode(_votante)
		}
	}
}

//UPDATE
func UpdateVotante(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	params := mux.Vars(r)
	votanteID, err := strconv.Atoi(params["id"])
	var updateVotante VOTANTE
	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Invalid Data")

	}
	json.Unmarshal(reqBody, &updateVotante)
	for i, votante := range votantes {
		if votante.ID == votanteID {
			votantes = append(votantes[:i], votantes[i+1:]...)
			updateVotante.ID = votanteID
			votantes = append(votantes, updateVotante)

			fmt.Fprintf(w, "votante Updated %v", votanteID)
		}
	}

}

//DELETED
func DeletedVotante(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	params := mux.Vars(r)
	votanteID, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	for index, _votante := range votantes {
		if _votante.ID == votanteID {
			votantes = append(votantes[:index], votantes[index+1:]...)
			fmt.Fprintf(w, "votante Deleted %v", votanteID)
			break
		}
	}

}

type CANDIDATE struct {
	ID              int    `json:"id,omitempty"`
	Nombre          string `json:"nombre,omitempty"`
	Apellido        string `json:"apellido,omitempty"`
	DNI             string `json:"dni,omitempty"`
	Numero_votacion string `json:"numero_votacion,omitempty"`
}

//CREA votante
//READ botante
//UP datos
//delete
type VOTANTE struct {
	ID             int    `json:"id,omitempty"`
	DNI            string `json:"dni,omitempty"`
	Nombre         string `json:"nombre,omitempty"`
	Apellido       string `json:"apellido,omitempty"`
	Lugar_Votacion string `json:"lugar_votacion,omitempty"`
	Candidato_Voto string `json:"candidato_voto,omitempty"`
}

func main() {
	router := mux.NewRouter()

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	router.HandleFunc("/candidatos", GetCandidatos).Methods(http.MethodGet)
	router.HandleFunc("/candidatos/create", CreateCandidato).Methods(http.MethodPost)
	router.HandleFunc("/candidatos/find/{id}", ReadCandidato).Methods(http.MethodGet)
	router.HandleFunc("/candidatos/update/{id}", UpdateCandidato).Methods(http.MethodPut)
	router.HandleFunc("/candidatos/delete/{id}", DeletedCandidato).Methods(http.MethodDelete)

	router.HandleFunc("/votantes", GetVotantes).Methods("GET")
	router.HandleFunc("/votantes/create", CreateVotante).Methods("POST")
	router.HandleFunc("/votantes/find/{id}", ReadVotante).Methods("GET")
	router.HandleFunc("/votantes/update/{id}", UpdateVotante).Methods("PUT")
	router.HandleFunc("/votantes/delete/{id}", DeletedVotante).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router)))

}
