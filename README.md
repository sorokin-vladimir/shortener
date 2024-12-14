# URL Shortener

URL Shortener is a simple link-shortening service designed as a learning project. It features a web interface and a Telegram bot, both sharing the same backend logic and database.

## Features
- Shorten URLs via:
  - **Web interface**: Easily shorten links through a user-friendly website.
  - **Telegram bot**: Interact with the service directly from Telegram.
- One shared Redis database for all operations.
- Easy deployment using Docker.
- URL statistics and analytics (optional, later).

## Technologies Used
- **Programming Language**: Go
- **Database**: Redis
- **Telegram Bot API**: [go-telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api)
- **Web Framework**: Built-in `net/http` (migrate to another soon)
- **Containerization**: Docker

## Installation and Setup

### Prerequisites
- [Go](https://go.dev/dl/) installed on your system.
- [Docker](https://www.docker.com/) and [Docker Compose](https://docs.docker.com/compose/).
- [mise-en-place](https://mise.jdx.dev/installing-mise.html) tasks runner

### Running Locally
1.  Clone the repository:
    ```bash
    git clone git@github.com:sorokin-vladimir/shortener.git url-shortener
    cd url-shortener
    ```
2.  Copy an `.env` file from `.env.example` in the root directory and set your Telegram API key
3.	Start the application locally:
    ```bash
    mise dev
    ```
4.	Access the web interface at http://localhost:8080
5.	Interact with the Telegram bot using the bot token you configured in .env.

### Running Tests
```bash
# Terminal 1
mise dev

# Terminal 2
mise test
```

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.
