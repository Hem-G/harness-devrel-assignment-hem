# 🚀 Harness DevRel Assignment — Multi-Environment Microservice Deployment

This repository implements the **Harness DevRel Exercise**: deploying a simple microservice across **Dev**, **QA**, and **Production** Kubernetes environments using a **fully automated Harness CD pipeline** with approval gates and a Blue/Green deployment strategy.

## ✅ Summary
This repository demonstrates a complete, production-ready Harness CD setup:
    - Go microservice deployed across Dev, QA, and Prod environments.
    - Kustomize-based Kubernetes configuration.
    - Blue/Green deployment strategy with manual approval gates.
    - All components integrated and verified with Harness connectors and delegate.

## 👀 Reviewer Tip
Each file in this repo is self-documented with plain-English comments.
You can browse in order (app/ → k8s/ → design-doc/) to see how the code, Kubernetes manifests, and Harness pipeline fit together.

---

## 🎯 Objectives of the Assignment

- Create a Harness CD pipeline to deploy a **microservice** into **Dev**, **QA**, and **Prod** environments.
- Include **approval gates** before Production deployments.
- Implement a configurable deployment strategy (**Blue/Green** chosen here).
- Demonstrate:
  - Pipeline design and execution
  - Running pod verification
  - Deployment strategy reasoning
  - Feedback on Harness documentation and flow

---

## 🏗️ Repository Structure

| Path | Purpose |
|------|----------|
| **`app/service/`** | Contains the Go-based microservice and its Dockerfile. A lightweight web server exposing `/`, `/health`, and `/version` routes. |
| **`k8s/base/`** | The *shared base* Kubernetes manifests (Deployment + Service) used by all environments. |
| **`k8s/overlays/dev`**, **`qa`**, **`prod`** | Environment-specific overlays that reuse the base manifests and override only what’s needed — like replica counts and resource prefixes. |
| **`design-doc/`** | Contains documentation artifacts describing the Harness setup. Includes the conceptual pipeline file **`pipeline-design-concept.yaml`**, which mirrors the structure and logic of the real pipeline created in Harness. |

---

## ⚙️ How It All Works Together

1. **Go Microservice (`app/service`)**
   - Minimal web service written in Go.
   - Endpoints:
     - `/` → main route returning `"hello from myservice"`.
     - `/health` → used by Kubernetes readiness and liveness probes.
     - `/version` → returns `"version 0.1.0"` for traceable build info.
   - Simple, stateless design — ideal for demonstrating CI/CD flows.

2. **Docker Image**
   - Built using a minimal **multi-stage Dockerfile** for small image size.
   - Image is published to **GitHub Container Registry (GHCR)**.
   - Example tag: `ghcr.io/hem-g/myservice:0.1.0`

3. **Kubernetes Manifests (Kustomize)**
   - `base/` defines common deployment and service resources.
   - Each environment (`dev`, `qa`, `prod`) extends this base using overlays.
   - The overlays:
     - Add environment-specific prefixes (`dev-`, `qa-`, `prod-`).
     - Modify the replica count (1 for dev, 2 for QA, 3 for prod).
     - Reuse the same Docker image and service configuration.
   - `harness.io/color` labels are added to the Service for **Blue/Green deployments**.

4. **Harness CD Pipeline**
   - Defined and executed in the Harness UI.
   - The conceptual structure is documented in [`design-doc/pipeline-design-concept.yaml`](design-doc/pipeline-design-concept.yaml).
   - The pipeline contains four stages:
     1. **Deploy to Dev** → Rolling deployment.
     2. **Deploy to QA** → Canary rollout for controlled testing.
     3. **Approve Prod Deployment** → Manual approval gate.
     4. **Deploy to Prod** → Blue/Green deployment with live traffic swap.
   - Each stage references its matching **Harness Environment** (Dev, QA, Prod).
   - The pipeline integrates all connectors:
     - GitHub connector (for this repo)
     - Docker registry connector (GHCR)
     - Kubernetes cluster connector (manual certificate-based setup)

---

## 🧪 Local Test Instructions

### 🧱 Build & Push the Docker Image
```bash
docker build -t ghcr.io/<your-username>/myservice:0.1.0 ./app/service
docker push ghcr.io/<your-username>/myservice:0.1.0
```

### 🚀 Deploy Locally to Dev

```bash
Copy code
kubectl apply -k k8s/overlays/dev
kubectl get pods
kubectl port-forward svc/dev-myservice 8080:80
curl http://localhost:8080/health
```

### 🔍 Verify

```bash
Copy code
kubectl get pods -A           # View running pods
curl http://localhost:8080/   # Returns "hello from myservice"
curl http://localhost:8080/version  # Returns "version 0.1.0"
```

## 🧩 Harness Pipeline Overview

| Stage                       | Environment | Strategy   | Description                                         |
| --------------------------- | ----------- | ---------- | --------------------------------------------------- |
| **Deploy to Dev**           | Dev         | Rolling    | Deploy baseline version for functional testing      |
| **Deploy to QA**            | QA          | Canary     | Deploy a single canary pod for validation           |
| **Approve Prod Deployment** | —           | Manual     | Human review before production promotion            |
| **Deploy to Prod**          | Prod        | Blue/Green | Deploy new version alongside old, then swap traffic |

Each stage deploys the same manifests but uses its own Kustomize overlay — ensuring isolated, reproducible environments.

## 🧠 Design Rationale & Strategy
Chosen Deployment Strategy:
    - Blue/Green for Production — selected for:
        - Zero-downtime user experience.
        - Easy rollback (switch traffic back to the old version instantly).
        - Simple integration with Harness via harness.io/color label swapping.

    - Why Kustomize (and not Helm):
        - Lightweight and declarative.
        - No template engine — purely YAML-based.
        - Ideal for GitOps and Harness CD pipelines.
        - Makes it clear exactly what changes between environments.

    - Why This Structure:
        - One base → three overlays = consistency + flexibility.
        - Harness can directly map its environments to these overlays.
        - Enables clean promotion flows (Dev → QA → Prod).

## 🩺 Verification Checklist

✅ Pipeline executed successfully across all stages (Dev, QA, Prod)
✅ Prod Blue/Green deployment logs verified in Harness
✅ Running pods confirmed via kubectl get pods
✅ /health and /version endpoints verified in all environments
✅ All connectors (Git, GHCR, K8s) tested successfully

## 💬 Feedback on Harness Documentation

The official Harness “Deploy Your Own App” guide is an excellent starting point, but it could be improved by:
- Adding troubleshooting steps for delegate-based K8s connectors.
- Including examples for local Kind cluster integrations.
- Explaining label requirements for Blue/Green (harness.io/color).
- Providing sample manifests for simple apps like this one.

