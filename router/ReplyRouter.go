package router

import (
	"encoding/json"
	"net/http"
	"strconv"

	"AiPetBack/db"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var replyCRUD = db.ReplyCRUD{}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/replies", createReply).Methods("POST")
    r.HandleFunc("/replies/{id}", getReply).Methods("GET")
    r.HandleFunc("/replies", getAllReplies).Methods("GET")
    r.HandleFunc("/replies/{id}", updateReply).Methods("PUT")
    r.HandleFunc("/replies/{id}", deleteReply).Methods("DELETE")

    http.ListenAndServe(":8080", r)
}

func createReply(w http.ResponseWriter, r *http.Request) {
    var reply db.Reply
    if err := json.NewDecoder(r.Body).Decode(&reply); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := replyCRUD.CreateByObject(&reply); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(reply)
}

func getReply(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        http.Error(w, "Invalid reply ID", http.StatusBadRequest)
        return
    }

    reply, err := replyCRUD.FindById(uint(id))
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            http.Error(w, "Reply not found", http.StatusNotFound)
        } else {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
    }

    json.NewEncoder(w).Encode(reply)
}

func getAllReplies(w http.ResponseWriter, r *http.Request) {
    replies, err := replyCRUD.FindAll()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(replies)
}

func updateReply(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        http.Error(w, "Invalid reply ID", http.StatusBadRequest)
        return
    }

    var reply db.Reply
    if err := json.NewDecoder(r.Body).Decode(&reply); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    reply.ID = uint(id)
    if err := replyCRUD.UpdateByObject(&reply); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(reply)
}

func deleteReply(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        http.Error(w, "Invalid reply ID", http.StatusBadRequest)
        return
    }

    if err := replyCRUD.SafeDeleteById(uint(id)); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}