# 📈 Stocky Assignment

A minimal working Golang REST API using **Gin**, **Logrus**, and **PostgreSQL**.

## Features
- POST /reward → Record user stock reward
- GET /stats/:userId → Get today’s stock rewards & total portfolio INR value
- Hourly stock price updater (mock random prices)

## Setup
1. Update `.env` with your PostgreSQL credentials
2. Create DB:
```sql
CREATE DATABASE assignment;
\c assignment
CREATE TABLE rewards (
	id SERIAL PRIMARY KEY,
	user_id INT NOT NULL,
	stock VARCHAR(20) NOT NULL,
	shares NUMERIC(18,6),
	timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```
3. Run project:
```bash
go mod tidy
go run main.go
```
4. Test APIs in Postman
