module dc20clerk/backend/identity

go 1.25.1

require (
	dc20clerk/common v0.0.0
	github.com/gorilla/mux v1.8.1
	github.com/joho/godotenv v1.5.1
)

replace dc20clerk/common => ../common

require github.com/golang-jwt/jwt/v5 v5.3.0
