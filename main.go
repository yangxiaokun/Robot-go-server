package main

import (
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/ws", handleWebSocket)
    http.HandleFunc("/send", handleSendMessage)
    http.HandleFunc("/delete", handleDeleteFriends)
    http.HandleFunc("/addGroup", handleAddGroup)
    log.Println("Server started on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
