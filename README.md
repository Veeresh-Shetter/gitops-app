# 🚀 Cloud-Native GitOps Application Lifecycle Pipeline

An enterprise-grade, fully automated GitOps Continuous Integration (CI) and Continuous Delivery (CD) pipeline deploying a high-availability Go web microservice into a local Kubernetes environment.

---

## 🏗️ Architecture & Tools Breakdown

*   **Application Backend:** High-performance HTTP web server using **Go (Golang)** featuring active endpoint logging and standardized `/healthz` liveness probes.
*   **Containerization:** Multi-stage **Docker** build optimized for minimal production footprint and quick container layer caching.
*   **Continuous Integration (CI):** Formulated a **GitHub Actions** orchestration pipeline triggering on code updates to dynamically build, tag, and publish application containers directly to the **GitHub Container Registry (GHCR)**.
*   **Orchestration & Infrastructure:** Configured declarative **Kubernetes** manifest layouts (`deployment.yaml`, `service.yaml`) running multiple decoupled app replicas behind an internal network load balancer via **Minikube**.
*   **Continuous Delivery (CD):** Implemented **Argo CD** to manage automated target state synchronization, declarative drift self-healing, and structural resource pruning following strict GitOps design principles.

---

## 🛠️ Step-by-Step Implementation Guide

### 📂 Phase 1: Application & Container Setup
1. **Develop the Go Microservice:** Create a `main.go` file establishing the web routes, logging middleware, and target port `8080`.
2. **Write the Multi-Stage Dockerfile:** Create a file named `Dockerfile` to handle the isolated compilation block and light Alpine runner image.
3. **Local Testing:** Build and run the container locally to verify functionality:
   ```bash
   docker build -t gitops-app:local .
   docker run -d -p 8080:8080 gitops-app:local
   ```

### 🤖 Phase 2: Automated CI Setup (GitHub Actions)
1. **Establish the Workflow Path:** Create the rigid hidden folder structure that GitHub requires for automation scanning:
   ```bash
   mkdir -p .github/workflows
   ```
2. **Create the CI Configuration:** Create a file inside that directory named `ci.yml`.
3. **Configure the Automation Script:** Write the job definitions specifying the `ghcr.io` login rules, short SHA version tagging, and push conditions (`on: push: branches: [main]`).
4. **Generate Access Token:** Go to GitHub Developer Settings, create a Personal Access Token (PAT) with `repo` and `workflow` permissions.
5. **Push Code to Remote:** Initialize Git, authenticate using the token, and push files to trigger the remote automation virtual runner.

### ☸️ Phase 3: Infrastructure Setup (Kubernetes)
1. **Spin Up the Cluster:** Boot your local production-grade single-node cluster environment:
   ```bash
   minikube start --driver=docker
   ```
2. **Author the Manifests:** Use terminal output strings to compose your declarative deployment configurations:
   * `deployment.yaml`: Defines target state, `imagePullPolicy: Always`, and 2 matching replicas.
   * `service.yaml`: Provisions internal network routing utilizing a static `NodePort (30080)`.
3. **Apply to Cluster Engine:** Instruct `kubectl` to deploy the resources:
   ```bash
   kubectl apply -f deployment.yaml
   kubectl apply -f service.yaml
   ```

### 🐙 Phase 4: GitOps Git Delivery Setup (Argo CD)
1. **Create the Isolated Namespace:** Set up a dedicated virtual room inside the cluster:
   ```bash
   kubectl create namespace argocd
   ```
2. **Deploy Argo CD Infrastructure:** Run the official resource manifests using a decoded base64 stream to establish the background controller pods.
3. **Retrieve Default Password:** Access the automatically generated secret layer using terminal decryption:
   ```bash
   kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d; echo ""
   ```
4. **Establish Dashboard Bridge:** Start port-forward networking to expose the dashboard on `localhost:8080`.
5. **Connect Git Repository:** Open the UI, click **New App**, and fill out the mapping form targeting your repository link, set the sync policy to **Automatic**, and select `self-heal` capabilities.
6. **Expose Globally:** Boot an edge-routing reverse proxy tunnel engine (`ngrok http 34219`) to securely expose the running cluster application to external public internet networks.
