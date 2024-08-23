package router

/*import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// MessageRouter handles the routing for messages
func MessageRouter(r *mux.Router) {
	r.HandleFunc("/messages", getAllMessages).Methods(http.MethodGet)
	r.HandleFunc("/messages/{id}", getMessageByID).Methods(http.MethodGet)
	r.HandleFunc("/messages", createMessage).Methods(http.MethodPost)
	r.HandleFunc("/messages/{id}", updateMessage).Methods(http.MethodPut)
	r.HandleFunc("/messages/{id}", deleteMessage).Methods(http.MethodDelete)
}

func getAllMessages(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement logic to get all messages from the database
	messages := getAllMessagesFromDatabase()
	jsonResponse, err := json.Marshal(messages)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func getMessageByID(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement logic to get a message by its ID from the database
	// Example implementation:
	// id := mux.Vars(r)["id"]
	// message := getMessageByIDFromDatabase(id)
	// if message == nil {
	//     http.NotFound(w, r)
	//     return
	// }
	// jsonResponse, err := json.Marshal(message)
	// if err != nil {
	//     http.Error(w, err.Error(), http.StatusInternalServerError)
	//     return
	// }
	// w.Header().Set("Content-Type", "application/json")
	// w.Write(jsonResponse)
}

func createMessage(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement logic to create a new message in the database
	// Example implementation:
	// var message Message
	// err := json.NewDecoder(r.Body).Decode(&message)
	// if err != nil {
	//     http.Error(w, err.Error(), http.StatusBadRequest)
	//     return
	// }
	// createdMessage := createMessageInDatabase(message)
	// jsonResponse, err := json.Marshal(createdMessage)
	// if err != nil {
	//     http.Error(w, err.Error(), http.StatusInternalServerError)
	//     return
	// }
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusCreated)
	// w.Write(jsonResponse)
}

func updateMessage(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement logic to update a message in the database
	// Example implementation:
	// id := mux.Vars(r)["id"]
	// var updatedMessage Message
	// err := json.NewDecoder(r.Body).Decode(&updatedMessage)
	// if err != nil {
	//     http.Error(w, err.Error(), http.StatusBadRequest)
	//     return
	// }
	// updatedMessage.ID = id
	// updatedMessage := updateMessageInDatabase(updatedMessage)
	// jsonResponse, err := json.Marshal(updatedMessage)
	// if err != nil {
	//     http.Error(w, err.Error(), http.StatusInternalServerError)
	//     return
	// }
	// w.Header().Set("Content-Type", "application/json")
	// w.Write(jsonResponse)
}

func deleteMessage(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement logic to delete a message from the database
	// Example implementation:
	id := mux.Vars(r)["id"]
	deleteMessageFromDatabase(id)
	w.WriteHeader(http.StatusNoContent)
}
*/
