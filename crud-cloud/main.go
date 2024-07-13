package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/gorilla/handlers"
)

// Tarea representa una estructura para las tareas
type Tarea struct {
    ID     string `json:"id"`
    Nombre string `json:"nombre"`
}

var tareas []Tarea

func obtenerTareas(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(tareas)
}

func obtenerTarea(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    parametros := mux.Vars(r)
    for _, item := range tareas {
        if item.ID == parametros["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Tarea{})
}

func crearTarea(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var tarea Tarea
    _ = json.NewDecoder(r.Body).Decode(&tarea)
    tarea.ID = "1" // Esto debería ser un ID único en un caso real
    tareas = append(tareas, tarea)
    json.NewEncoder(w).Encode(tarea)
}

func actualizarTarea(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    parametros := mux.Vars(r)
    for indice, item := range tareas {
        if item.ID == parametros["id"] {
            tareas = append(tareas[:indice], tareas[indice+1:]...)
            var tarea Tarea
            _ = json.NewDecoder(r.Body).Decode(&tarea)
            tarea.ID = parametros["id"]
            tareas = append(tareas, tarea)
            json.NewEncoder(w).Encode(tarea)
            return
        }
    }
    json.NewEncoder(w).Encode(tareas)
}

func eliminarTarea(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    parametros := mux.Vars(r)
    for indice, item := range tareas {
        if item.ID == parametros["id"] {
            tareas = append(tareas[:indice], tareas[indice+1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(tareas)
}

func main() {
    enrutador := mux.NewRouter()

    // Rutas API
    enrutador.HandleFunc("/tareas", obtenerTareas).Methods("GET")
    enrutador.HandleFunc("/tareas/{id}", obtenerTarea).Methods("GET")
    enrutador.HandleFunc("/tareas", crearTarea).Methods("POST")
    enrutador.HandleFunc("/tareas/{id}", actualizarTarea).Methods("PUT")
    enrutador.HandleFunc("/tareas/{id}", eliminarTarea).Methods("DELETE")

    // Servir archivos estáticos
    enrutador.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

    // Manejo de CORS
    cabecerasOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
    metodosOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
    orígenesOk := handlers.AllowedOrigins([]string{"*"})

    log.Fatal(http.ListenAndServe(":8000", handlers.CORS(cabecerasOk, metodosOk, orígenesOk)(enrutador)))
}