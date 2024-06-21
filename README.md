# Social Media Chat Application / Aplicativo de Chat de Mídia Social

Welcome to the Social Media Chat Application! This project aims to create a scalable and maintainable chat application where users can join different groups and communicate in real time.

## Introduction / Introdução

This Social Media Chat Application is designed to facilitate real-time communication among users. It supports group-based messaging, allowing users to join different chat groups and interact with each other.

O Aplicativo de Chat de Mídia Social foi projetado para facilitar a comunicação em tempo real entre os usuários. Ele suporta mensagens em grupos, permitindo que os usuários participem de diferentes grupos de chat e interajam entre si.

## Features / Funcionalidades

- Real-time messaging using WebSocket / Mensagens em tempo real usando WebSocket
- Group-based chat rooms / Salas de chat baseadas em grupos
- User-friendly web interface / Interface web amigável
- Scalable architecture / Arquitetura escalável
- Easy to configure and deploy / Fácil de configurar e implementar

## Architecture / Arquitetura

The application is structured to follow a clean and modular architecture, making it easy to maintain and extend. Below is an overview of the architecture:

A aplicação é estruturada seguindo uma arquitetura limpa e modular, facilitando a manutenção e a extensão. Abaixo está uma visão geral da arquitetura:

### Directory Structure / Estrutura de Diretórios

```
/chat-app
│
├── /cmd
│ └── /server
│ └── main.go
│
├── /pkg
│ └── /websocket
│ ├── hub.go
│ ├── client.go
│ └── message.go
│
├── /internal
│ ├── /handlers
│ │ ├── websocket_handler.go
│ │ └── http_handler.go
│ └── /config
│ └── config.go
│
└── /web
└── index.html
```

### Components / Componentes

1. **Configuration / Configuração (`internal/config`)**: Manages application configuration settings / Gerencia as configurações de aplicação.
2. **WebSocket Hub / Hub WebSocket (`pkg/websocket/hub.go`)**: Manages client connections and group-based messaging / Gerencia as conexões de clientes e mensagens baseadas em grupos.
3. **Client / Cliente (`pkg/websocket/client.go`)**: Represents individual WebSocket connections / Representa as conexões individuais de WebSocket.
4. **Message / Mensagem (`pkg/websocket/message.go`)**: Defines the structure of messages exchanged between clients / Define a estrutura das mensagens trocadas entre os clientes.
5. **Handlers / Manipuladores (`internal/handlers`)**: Defines HTTP and WebSocket handlers for routing and managing connections / Define manipuladores HTTP e WebSocket para roteamento e gerenciamento de conexões.

### Detailed Overview / Visão Geral Detalhada

- **Configuration / Configuração (`internal/config/config.go`)**: Centralizes all configuration settings, making it easy to manage and update application settings / Centraliza todas as configurações, facilitando a gestão e atualização das configurações da aplicação.
- **WebSocket Hub / Hub WebSocket (`pkg/websocket/hub.go`)**: Manages client connections and group-based messaging, ensuring efficient communication / Gerencia as conexões dos clientes e as mensagens baseadas em grupos, garantindo uma comunicação eficiente.
- **Client / Cliente (`pkg/websocket/client.go`)**: Handles individual WebSocket connections, processing incoming and outgoing messages / Gerencia conexões individuais de WebSocket, processando mensagens de entrada e saída.
- **Handlers / Manipuladores (`internal/handlers`)**: Implements HTTP and WebSocket handlers for routing requests and managing connections / Implementa manipuladores HTTP e WebSocket para rotear solicitações e gerenciar conexões.

## Installation / Instalação

To install and run the application, follow these steps:

Para instalar e executar a aplicação, siga estes passos:

1. Clone the repository: `[git clone https://github.com/yourusername/chat-app](https://github.com/MateusMSoares/Social-Media)`
   Clone o repositório: `[git clone https://github.com/yourusername/chat-app](https://github.com/MateusMSoares/Social-Media)`
   
2. Navigate to the project directory: `cd chat-app`
   Navegue até o diretório do projeto: `cd chat-app`
   
3. Install dependencies: `go mod tidy`
   Instale as dependências: `go mod tidy`
   
4. Build and run the application: `go run cmd/server/main.go`
   Compile e execute a aplicação: `go run cmd/server/main.go`

## Usage / Uso

Once the application is running, access the web interface at `http://localhost:8080` to start using the chat application.

Depois que a aplicação estiver em execução, acesse a interface web em `http://localhost:8080` para começar a usar o aplicativo de chat.

## Configuration / Configuração

The application can be configured using environment variables or by editing the `config/config.go` file.

A aplicação pode ser configurada utilizando variáveis de ambiente ou editando o arquivo `config/config.go`.

## Contributing / Contribuições

Contributions are welcome! Feel free to submit issues and pull requests.

Contribuições são bem-vindas! Sinta-se à vontade para enviar problemas e solicitações de pull.

## License / Licença

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

Este projeto está licenciado sob a Licença MIT. Consulte o arquivo [LICENSE](LICENSE) para mais detalhes.
