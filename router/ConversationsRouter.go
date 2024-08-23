package router

import (
	"encoding/json"
	"net/http"
	"strconv"

	"AiPetBack/db" 

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var conversationCRUD = db.ConversationCRUD{}

func RegisterConversationRoutes(r *mux.Router) {
    r.HandleFunc("/conversations", createConversation).Methods("POST")
    r.HandleFunc("/conversations/{id}", getConversationByID).Methods("GET")
    r.HandleFunc("/conversations", getConversationsByUser1).Methods("GET").Queries("user1", "{user1}")
    r.HandleFunc("/conversations", getConversationsByUser2).Methods("GET").Queries("user2", "{user2}")
    r.HandleFunc("/conversations/{id}", updateConversation).Methods("PUT")
}

func createConversation(w http.ResponseWriter, r *http.Request) {
    var conversation db.Conversations
    if err := json.NewDecoder(r.Body).Decode(&conversation); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := conversationCRUD.CreateByObject(&conversation); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(conversation)
}

func getConversationByID(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        http.Error(w, "Invalid conversation ID", http.StatusBadRequest)
        return
    }

    conversation, err := conversationCRUD.GetConversationByID(uint(id))
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            http.Error(w, "Conversation not found", http.StatusNotFound)
        } else {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
    }

    json.NewEncoder(w).Encode(conversation)
}

func getConversationsByUser1(w http.ResponseWriter, r *http.Request) {
    user1 := r.URL.Query().Get("user1")
    conversations, err := conversationCRUD.GetConversationByUser1Name(user1)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(conversations)
}

func getConversationsByUser2(w http.ResponseWriter, r *http.Request) {
    user2 := r.URL.Query().Get("user2")
    conversations, err := conversationCRUD.GetConversationByUser2Name(user2)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(conversations)
}

func updateConversation(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        http.Error(w, "Invalid conversation ID", http.StatusBadRequest)
        return
    }

    var conversation db.Conversations
    if err := json.NewDecoder(r.Body).Decode(&conversation); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    conversation.ID = uint(id)
    if err := conversationCRUD.UpdateByObject(&conversation); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(conversation)
}
