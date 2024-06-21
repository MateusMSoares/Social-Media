package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
	Group    int    `json:"grupo"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// explicação dos makes: cria mapas para armazenar clientes e canais para transmissão de mensagens.
var broadcast = make(chan Message)
var groups = make(map[int]map[*websocket.Conn]bool)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		caminho := strings.TrimPrefix(r.URL.Path, "/")
		if caminho == "" {
			caminho = "index.html"
		} else {
			caminho = "index.html"
		}
		http.ServeFile(w, r, caminho)
	})
	http.HandleFunc("/ws", handleConnections)

	go handleMessages()

	fmt.Println("Servidor iniciado na porta :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Erro ao iniciar servidor: %v\n", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Erro ao conectar websocket: %v\n", err)
		return
	}
	defer conn.Close()

	// Recebe a identificação do grupo da URL
	grupoIdStr := r.URL.Query().Get("grupo")
	var grupoId int
	if grupoIdStr != "" {
		// Converte a string grupoIdStr em um inteiro e armazena em grupoId
		fmt.Sscanf(grupoIdStr, "%d", &grupoId)
	}

	// Se o grupo não existir, crie-o
	if groups[grupoId] == nil {
		groups[grupoId] = make(map[*websocket.Conn]bool)
	}
	groups[grupoId][conn] = true

	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			fmt.Printf("Erro ao ler mensagem: %v\n", err)
			delete(groups[grupoId], conn)
			break
		}
		msg.Group = grupoId
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		msg := <-broadcast
		for client := range groups[msg.Group] {
			err := client.WriteJSON(msg)
			if err != nil {
				fmt.Printf("Erro ao enviar mensagem: %v\n", err)
				client.Close()
				delete(groups[msg.Group], client)
			}
		}
	}
}
