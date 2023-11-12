# Disc Golf App
Disc golf application that scrapes the web and saves disc golf manufacturer, and their discs to a postgres db. and then are served in RESTful API.
### Prerequisites
1. [Docker](https://www.docker.com/get-started/)
1. [Postgres-sql (14.1)](https://www.postgresql.org/)
1. [GoLang (1.18)](https://go.dev/learn/)

## Local Setup
1. local Server
```bash
# Startup local db and server
docker-compose up -d
```
2. Running the scraper
```bash
    go scraper/scraper.go
```
