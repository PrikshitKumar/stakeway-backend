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

## Section 2: 
- cd backend
- Make Sure to install the docker first. 
- Up the Services in Container after build
```bash
docker-compose up --build
```

- Test the API calls: 
```bash
curl -X POST "http://localhost:8080/validators" -H "Content-Type: application/json" -d '{
  "num_validators": 5,
  "fee_recipient": "0x1234567890abcdef1234567890abcdef12345678"
}'
```
```bash
curl -X GET "http://localhost:8080/validators/0aa163c4-dc51-427a-acf4-24eed8c76b16" | jq
```

K8s integration: 
- cd k8s
- Install Minikube (Local Kubernates Cluster): 
```bash
brew install minikube (for macOS)
minikube version
```
- Start Minikube Cluster
```bash
# If running the cluster again, make sure to use stop and delete CLI before starting the cluster
minikube stop 
minikube delete
minikube start
minikube status
```

- Deploy the Backend and Refis on Minikube (Light weight Cluster): 
```bash
eval $(minikube -p minikube docker-env)
cd .. 
docker build -t stakeway-backend .
docker pull redis:alpine
```
- Apply Kubernetes YAML Files
```bash
cd k8s/
kubectl apply -f ./  
```
- Wait for a minute to get all the services up. 
- Verify the Deployment of Pods and Services: 
```bash
kubectl get pods
kubectl get services
minikube service backend-service --url

minikube ip
```

- Route Forwading to communicate with Pods: 
```bash
kubectl port-forward svc/backend-service 8080:80
kubectl port-forward svc/prometheus-service 9090:9090
kubectl port-forward svc/grafana-service 3000:3000
```

- Test the APIs: 
```bash
curl -X POST "http://127.0.0.1:8080/validators" -H "Content-Type: application/json" -d '{
  "num_validators": 5,
  "fee_recipient": "0x1234567890abcdef1234567890abcdef12345678"
}'
```
```bash
curl -X GET "http://127.0.0.1:8080/validators/778aa6c7-56a3-4c76-a498-28b41d5d249f" | jq
```

UIs can be accessible throug: 
- Prometheus: http://localhost:9090/
- Grafna: http://localhost:3000/
Username: admin, Password: admin


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

- TxHash for my successful transaction: ``
- Verify the transaction from Explorer: ``
