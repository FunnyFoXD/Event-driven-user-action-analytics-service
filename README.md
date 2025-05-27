# Event-driven-user-action-analytics-service
Project, where I just want impove my stack base

## What is it?
Product can save events (events type in progress, now it is lines in database) and show it in browser. I will be improving this product until perfection.

## Tech stack

- 🟦 [Go](https://go.dev/)
- 🌐 [Chi](https://github.com/go-chi/chi)
- 🗄️ [GORM](https://gorm.io/)
- ⚡ [Zap](https://uber-go.github.io/zap/)
- ⚙️ [Viper](https://github.com/spf13/viper)
- 🐘 [Postgres](https://www.postgresql.org/)
- 🔄 [Kafka](https://kafka.apache.org/)
- 📊 [Grafana](https://grafana.com/)
- 🐳 [Docker](https://www.docker.com/)
- 📦 [Docker Compose](https://docs.docker.com/compose/)
- 🇫 [HTML/CSS/JS](https://www.w3schools.com/)

## Getting Start
Follow these steps to run the project locally:

### 1. Clone the repository
```bash
git clone https://github.com/FunnyFoXD/Event-driven-user-action-analytics-service.git
cd Event-driven-user-action-analytics-service
```

### 2. Add .env in root directory and configure it (follow for .env.example)
### 3. Start frontend using python:
```bash
# in /frontend/
python3 -m http.server 8081
```

### 4. Start frontend using docker-compose
```bash
# in /
docker-compose up --build -d
```

## Let's get fun
```bash
curl -X POST http://localhost:8080/event \
     -H "Content-Type: application/json" \
     -d '{"user_id": "admin", "action": "login"}'
```

Go to the localhost:8081/logs.html and you see table with your evet

## *More comfortable event manager will be in future updates
