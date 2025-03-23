# stakeway-backend

A **Golang-based backend service** for managing validator creation requests, integrating **Redis** for in-memory storage, and supporting **asynchronous processing**. This project includes:

- **RESTful API Endpoints** for validator management
- **Background task execution** for processing validator requests
- **Monitoring capabilities** via **Prometheus & Grafana**
- **Docker & Kubernetes** integration for scalability

---

## 🛠️ Setup & Installation

### **Section 1: Local Development**

#### **1️⃣ Install Redis**  
Follow the installation guide: [Redis Installation](https://redis.io/docs/latest/operate/oss_and_stack/install/install-redis/)

#### **2️⃣ Start Redis Server**  
Run in the first terminal:
```bash
redis-server
```

#### **3️⃣ Start the Backend Server**  
Run in the second terminal:
```bash
cd backend
go mod tidy
go run .
```

#### **4️⃣ Test API Endpoints**  
- **Create and Store Validators**
```bash
curl -X POST "http://localhost:8080/validators" -H "Content-Type: application/json" -d '{
  "num_validators": 5,
  "fee_recipient": "0x1234567890abcdef1234567890abcdef12345678"
}'
```
- **Get Validator Status**
```bash
curl -X GET "http://localhost:8080/validators/0aa163c4-dc51-427a-acf4-24eed8c76b16" | jq
```

---

### **Section 2: Docker & Kubernetes Deployment**

#### **1️⃣ Start Services with Docker**  
- Ensure Docker is installed and running, then run:
```bash
cd backend
docker-compose up --build
```

#### **2️⃣ Test API Calls (Same as Section 1)**  
- **Create and Store Validators**
```bash
curl -X POST "http://localhost:8080/validators" -H "Content-Type: application/json" -d '{
  "num_validators": 5,
  "fee_recipient": "0x1234567890abcdef1234567890abcdef12345678"
}'
```
- **Get Validator Status**
```bash
curl -X GET "http://localhost:8080/validators/0aa163c4-dc51-427a-acf4-24eed8c76b16" | jq
```

#### **3️⃣ Kubernetes (K8s) Integration**  
- **Navigate to k8s folder**
```bash
cd k8s
```
- **Install Minikube (Local K8s Cluster)**  
```bash
brew install minikube  # For macOS
minikube version
```

- **Start Minikube Cluster**
```bash
# Before restarting the cluster, ensure you stop and delete it using the CLI first.
minikube stop
minikube delete
minikube start
minikube status
```

- **Deploy Backend & Redis to Minikube**
```bash
eval $(minikube -p minikube docker-env)
cd ..
docker build -t stakeway-backend .
docker pull redis:alpine
```

- **Apply Kubernetes YAML Files**
```bash
cd k8s/
kubectl apply -f ./  
```

- Wait a minute for all services to start up. 

- **Verify Deployment**
```bash
kubectl get pods
kubectl get services
minikube service backend-service --url # To get POD URL
minikube ip # To get exposed IP
```

- **Port Forwarding to Access Services from client side**
```bash
kubectl port-forward svc/backend-service 8080:80
kubectl port-forward svc/prometheus-service 9090:9090
kubectl port-forward svc/grafana-service 3000:3000
```

- **Test API Calls (Same as Section 1)**  
- **Create and Store Validators**
```bash
curl -X POST "http://localhost:8080/validators" -H "Content-Type: application/json" -d '{
  "num_validators": 5,
  "fee_recipient": "0x1234567890abcdef1234567890abcdef12345678"
}'
```
- **Get Validator Status**
```bash
curl -X GET "http://localhost:8080/validators/778aa6c7-56a3-4c76-a498-28b41d5d249f" | jq
```

#### **4️⃣ Access Monitoring UIs**  
- **Prometheus:** [http://localhost:9090/](http://localhost:9090/)
  - Search for `http_requests_total` in Prometheus UI to get the metrics data. 
- **Grafana:** [http://localhost:3000/](http://localhost:3000/)  
  - Username: `admin`
  - Password: `admin`

#### **5️⃣ Set Up Grafana Dashboard**  
- **Step 1: Connect to Prometheus**
  - Go to **Connections → Data Sources**
  - Click **Add Data Source** → Select **Prometheus**
  - Set **URL** to: `http://prometheus-service.default.svc.cluster.local:9090`
  - Click **Save & Test** → `Grafana is connected to Prometheus!`

- **Step 2: Create API Metrics Dashboard**
  - Navigate to **Dashboards → Create Dashboard → Add Visualization**
  - Select **Prometheus** as the **Data Source**
  - Under `Query` section, choose `__name__` in **Label filters**
  - Enter: `http_requests_total` in **Select Value**
  - Click **Run Queries → Save** the dashboard
  - Switch to **Graph Mode** → Save as `API Metrics Dashboard`

✅ **Now, API request counts will update in real-time!**

---

### **Section 3: Holesky Network Integration**

#### **1️⃣ Set Private Key in Session**  
```bash
export ETH_PRIVATE_KEY="<Your Private Key>"
```

#### **2️⃣ Fund Your Address**  
Use the [Holesky Faucet](https://holesky-faucet.pk910.de/) to get at least **33 ETH**.

#### **3️⃣ Verify ETH Balance**  
```bash
cast balance <Account Address> --rpc-url https://ethereum-holesky.publicnode.com
```

#### **4️⃣ Run the Script**  
```bash
cd holesky-integration
go mod tidy
go run .
```

#### **5️⃣ Verify Transaction**  
- **Transaction Hash:** `0x54e6231a2b37353cfc9f6715282f3f553adc8d72d17af47666e563f6496b626c`
- **Explorer:** [View on Holesky Etherscan](https://holesky.etherscan.io/tx/0x54e6231a2b37353cfc9f6715282f3f553adc8d72d17af47666e563f6496b626c)
- **Filter transactions from address:** `0x2fA0e97dc0bc4A5C65A700dc9e364f158853B1A4`

---

## 🎯 **Final Notes**
- Follow each section carefully for a smooth setup and seamless interactions.
- Reach out if you encounter any issues! 🚀
