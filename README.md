# stakeway-backend
A Golang-based backend service for managing validator creation requests, integrating Redis for storage, and supporting asynchronous processing. The project includes RESTful API endpoints, background task execution, and monitoring capabilities for scalable blockchain infrastructure.


# Steps: 
## Section 1: 
- Install Redis from: `https://redis.io/docs/latest/operate/oss_and_stack/install/install-redis/`
- Run the Server (in 1st terminal): 
```bash
redis-server
```
- Run the Backend Server (in 2nd terminal): 
```bash
cd backend
go mod tidy
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

## Section 3: 
- Add Private Key in current session: 
```bash
export ETH_PRIVATE_KEY="<Your Private Key>"
```
- Fund the address from Faucet: `https://holesky-faucet.pk910.de/` (Make sure to mine atleast 32 ETH for successful transaction)
- Verify the balance after mining: 
```bash 
cast balance <Account Address> --rpc-url https://ethereum-holesky.publicnode.com
```
- Execute the script: 
```bash
cd holesky-integration
go mod tidy
go run . 
```