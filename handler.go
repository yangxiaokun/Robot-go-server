package main

import (
    "encoding/json"
    "net/http"
    "log"
)

type SendMessageRequest struct {
    ClientID string `json:"client_id"`
    Action   string `json:"action"`
    Name     string `json:"name"`
    Content  *string `json:"content,omitempty"`
}

type DeleteFriendsRequest struct {
    ClientID string `json:"client_id"`
    Action   string `json:"action"`
    CustomerList     []string `json:"customer_list"`
}

type AddGroupRequest struct {
    ClientID string `json:"client_id"`
    Action   string `json:"action"`
    Name     string `json:"name"`
    StaffList []string `json:"staff_list"`
}

func handleSendMessage(w http.ResponseWriter, r *http.Request) {
    var req SendMessageRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        log.Println("Error decoding request:", err)
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }
    log.Println("Received message from client:", req.ClientID, req.Action,req.Name, req.Content)
    broadcastSendMessage(req.ClientID, req.Action, req.Name, req.Content)
    w.WriteHeader(http.StatusOK)
}

func handleDeleteFriends(w http.ResponseWriter, r *http.Request) {
    var req DeleteFriendsRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        log.Println("Error decoding request:", err)
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }
    log.Println("Received delete friends request:", req.ClientID, req.Action, req.CustomerList)
    broadcastDeleteFriends(req.ClientID, req.Action, req.CustomerList)
    w.WriteHeader(http.StatusOK)
}

func handleAddGroup(w http.ResponseWriter, r *http.Request) {
    var req AddGroupRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        log.Println("Error decoding request:", err)
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }
    log.Println("Received add group request:", req.ClientID, req.Action, req.Name, req.StaffList)
    broadcastAddGroup(req.ClientID, req.Action, req.Name, req.StaffList)
    w.WriteHeader(http.StatusOK)
}