# stakeway-backend
A Golang-based backend service for managing validator creation requests, integrating Redis for storage, and supporting asynchronous processing. The project includes RESTful API endpoints, background task execution, and monitoring capabilities for scalable blockchain infrastructure.


# Steps: 
- Install Redis from: `https://redis.io/docs/latest/operate/oss_and_stack/install/install-redis/`
- Run the Server (in 1st terminal): 
```bash
redis-server
```
- Run the Backend Server (in 2nd terminal): 
```bash
cd backend
go run .
```
- API Call: (To create and store the Validators in DB)
```bash
curl -X POST "http://localhost:8080/validators" -H "Content-Type: application/json" -d '{
  "num_validators": 5,
  "fee_recipient": "0x1234567890abcdef1234567890abcdef12345678"
}'
```
- API Call: (To get the Validator Status)
```bash
curl -X GET "http://localhost:8080/validators/0aa163c4-dc51-427a-acf4-24eed8c76b16" | jq
```