package main

import (
    "github.com/gorilla/websocket"
	// "github.com/google/uuid"
	"net/http"
    "log"
    "sync"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

type Client struct {
    conn *websocket.Conn
    ID   string
}

var clients = make(map[string]*Client)
var mu sync.Mutex

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("Error while upgrading connection:", err)
        return
    }

    // 生成一个唯一的客户端 ID（可以使用 UUID 或其他方法）
    // clientID := uuid.New().String() 
	clientID := "yangxiaokun"
    client := &Client{conn: conn, ID: clientID}
	log.Println("Client connected with ID:", clientID)

    mu.Lock()
    clients[clientID] = client
    mu.Unlock()

    defer func() {
        mu.Lock()
        delete(clients, clientID)
        mu.Unlock()
        conn.Close()
    }()

    for {
        _, _, err := conn.ReadMessage()
        if err != nil {
            log.Println("Error while reading message:", err)
            break
        }
    }
}

func broadcastSendMessage(clientID string, action string, name string, content *string) {
    mu.Lock()
    defer mu.Unlock()
    if client, ok := clients[clientID]; ok {
		sendMessage := Message{Action: action, Name: name, Content: content}
        err := client.conn.WriteJSON(sendMessage)
        if err != nil {
            log.Println("Error while sending message:", err)
            client.conn.Close()
            delete(clients, clientID)
        }
    } else {
        log.Printf("Client %s not found\n", clientID)
    }
}

func broadcastDeleteFriends(clientID string, action string, customerList []string) {
	mu.Lock()
	defer mu.Unlock()
	if client, ok := clients[clientID]; ok {
		sendMessage := DeleteFriendsMessage{Action: action, CustomerList: customerList}
        err := client.conn.WriteJSON(sendMessage)
		if err != nil {
			log.Println("Error while sending message:", err)
			client.conn.Close()
			delete(clients, clientID)
		}
	} else {
		log.Printf("Client %s not found\n", clientID)
	}
}

func broadcastAddGroup(clientID string, action string, name string, staffList []string) {
	mu.Lock()
	defer mu.Unlock()
	if client, ok := clients[clientID]; ok {
		sendMessage := AddGroupMessage{Action: action, Name: name, StaffList: staffList}
        err := client.conn.WriteJSON(sendMessage)
		if err != nil {
			log.Println("Error while sending message:", err)
			client.conn.Close()
			delete(clients, clientID)
		}
	}
}
