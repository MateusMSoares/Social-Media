package handlers

// import "fmt"

// func handleMessages() {
// 	for {
// 		// Pegar próxima mensagem do canal de broadcast
// 		msg := <-broadcast
// 		// Enviar mensagem para todos os clientes conectados
// 		for client := range clients {
// 			err := client.WriteJSON(msg)
// 			if err != nil {
// 				fmt.Printf("Erro ao enviar mensagem: %v\n", err)
// 				client.Close()
// 				delete(clients, client)
// 			}
// 		}
// 	}
// }
