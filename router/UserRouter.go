package router

import (
	"encoding/json"
	"net/http"
	"strconv"

	"AiPetBack/db"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var userCRUD = db.UserCRUD{}

func RegisterUserRoutes(r *mux.Router) {
    r.HandleFunc("/users", createUser).Methods("POST")
    r.HandleFunc("/users/{name}", getUserByName).Methods("GET")
    r.HandleFunc("/users", getAllUsers).Methods("GET")
    r.HandleFunc("/users/{name}", updateUser).Methods("PUT")
    r.HandleFunc("/users/{name}", deleteUserByName).Methods("DELETE")
    r.HandleFunc("/users/location", getUsersByLocation).Methods("GET")
}

func createUser(w http.ResponseWriter, r *http.Request) {
    var user db.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := userCRUD.CreateByObject(&user); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

func getUserByName(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    name := params["name"]

    user, err := userCRUD.GetUserByName(name)
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            http.Error(w, "User not found", http.StatusNotFound)
        } else {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
    }

    json.NewEncoder(w).Encode(user)
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
    users, err := userCRUD.GetAllUser()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(users)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    name := params["name"]

    var user db.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    user.UserName = name
    if err := userCRUD.UpdateByObject(&user); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(user)
}

func deleteUserByName(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    name := params["name"]

    if err := userCRUD.DeleteUserbyName(name); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

func getUsersByLocation(w http.ResponseWriter, r *http.Request) {
    lat, err := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
    if err != nil {
        http.Error(w, "Invalid latitude", http.StatusBadRequest)
        return
    }

    long, err := strconv.ParseFloat(r.URL.Query().Get("long"), 64)
    if err != nil {
        http.Error(w, "Invalid longitude", http.StatusBadRequest)
        return
    }

    radius, err := strconv.ParseFloat(r.URL.Query().Get("radius"), 64)
    if err != nil {
        http.Error(w, "Invalid radius", http.StatusBadRequest)
        return
    }

    users, err := userCRUD.GetUserByLocation(lat, long, radius)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(users)
}
