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

## 🧭 Reproducibility Notes

- This repository is fully self-contained for building and testing the microservice locally.  
- Harness pipelines are defined and executed in a private Harness account and is not available publicly.  
- Pipeline structure and execution logs are included in the companion Google Doc.

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
kubectl apply -k k8s/overlays/dev
kubectl get pods
kubectl port-forward svc/dev-myservice 8080:80
curl http://localhost:8080/health
```

### 🔍 Verify
```bash
kubectl get pods -A           # View running pods
curl http://localhost:8080/   # Returns "hello from myservice"
curl http://localhost:8080/health  # Returns "ok"
```

### ⚠️ Note on Updating Deployments
If you modify labels or selectors in your Kubernetes manifests (for example, when adding harness.io/color: blue for Blue/Green deployments), you may encounter this error when reapplying resources locally with Kustomize: `The Deployment "myservice" is invalid: spec.selector: Invalid value: {"matchLabels":{"app":"myservice","harness.io/color":"blue"}}: field is immutable`

This happens because Kubernetes does not allow changing spec.selector on an existing Deployment — the selector determines which Pods the Deployment manages, and changing it could orphan or duplicate Pods.

### ✅ How to fix it
Delete the existing Deployment first, then reapply:

```bash
kubectl delete deployment <deployment-name>
kubectl apply -k k8s/overlays/<environment>
```

### Example:

```bash
kubectl delete deployment prod-myservice
kubectl apply -k k8s/overlays/prod
```

---

## 🧩 Harness Pipeline Overview

| Stage                       | Environment | Strategy   | Description                                         |
| --------------------------- | ----------- | ---------- | --------------------------------------------------- |
| **Deploy to Dev**           | Dev         | Rolling    | Deploy baseline version for functional testing      |
| **Deploy to QA**            | QA          | Canary     | Deploy a single canary pod for validation           |
| **Approve Prod Deployment** | —           | Manual     | Human review before production promotion            |
| **Deploy to Prod**          | Prod        | Blue/Green | Deploy new version alongside old, then swap traffic |

Each stage deploys the same manifests but uses its own Kustomize overlay — ensuring isolated, reproducible environments.

---

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

---

## 🩺 Verification Checklist

✅ Pipeline executed successfully across all stages (Dev, QA, Prod)
✅ Prod Blue/Green deployment logs verified in Harness
✅ Running pods confirmed via kubectl get pods
✅ /health and /version endpoints verified in all environments
✅ All connectors (Git, GHCR, K8s) tested successfully

---

## 🗣️ Detailed Feedback on Harness Documentation & User Experience

As a new user with no prior DevOps experience, I found the Harness documentation very hard to navigate.  
While it contains a lot of information, it lacks a clear, beginner-friendly flow from *sign-up* to *successful deployment*.  
Below is a summary of key issues and improvement suggestions based on my real experience completing this assignment.

### 🔹 1. No Complete "Start-to-Finish" Flow
The documentation does not provide a single linear path that takes a new user from:
**sign up → create project → set up connectors → delegate → secrets → services → environments → pipeline → successful deployment.**

Each topic is explained in isolation. As a result, it’s unclear how the pieces connect.  
Even the official “Deploy Your Own App” tutorial doesn’t follow a sequential, beginner-oriented flow.

**Suggested Fix:**  
Create a *single master tutorial* that follows the real onboarding journey:
1. Sign up or log in to Harness
2. Create a project (show UI)
3. Add connectors (Git, Docker, Kubernetes)
4. Set up the delegate (with local/Kind example)
5. Create a service (attach manifests)
6. Create environments (Dev/QA/Prod)
7. Build and run a basic pipeline  
Each step should have a **clear goal, expected result, and screenshot.**

### 🔹 2. UI vs YAML Confusion
Many guides say *“set this up in the UI”*, but then show only YAML examples.  
Some pages use screenshots; others jump straight into YAML editing without context.  
This is confusing for new users who are not yet comfortable with Harness-as-Code.

**Suggested Fix:**  
Every tutorial should:
- Show **both** the UI path *and* the YAML equivalent.
- Be clear upfront: “This section assumes you are using the visual builder” or “This uses YAML editing mode.”
- Include consistent screenshots for the main stages — connectors, delegates, pipeline creation.

### 🔹 3. Missing Reference Example (No Code or Sample App)
The official docs explain *how Harness works*, but don’t give a working example repository or app.  
For someone new to CI/CD, this means there’s no baseline to compare your setup with.

**Suggested Fix:**  
Provide a public sample repo with:
- A tiny microservice (e.g., Go, Node.js, or Python)
- Dockerfile and Kubernetes manifests
- Example pipeline YAML or screenshots  
Users should be able to fork it, follow the docs, and get a successful deployment within 30–45 minutes.

### 🔹 4. Inconsistent and AI-Generated Style
It’s clear that some doc pages were auto-generated or bulk-written by AI without proper editing.  
Random words like **what** or **which** are in bold.  
Tone and formatting vary between pages, and sometimes the same concept is explained differently in separate places.

**Suggested Fix:**  
Have a consistent technical writing style guide for docs:
- Use imperative tone for instructions (“Click”, “Select”, “Run”)
- Avoid unnecessary bolding or filler words
- Keep all headings task-oriented (e.g., “Create a Delegate” instead of “About Delegates”)
- Run human QA for readability and flow

### 🔹 5. No Troubleshooting Guidance for Local or Kind Clusters
Running a local Kind or Minikube cluster is common for testing.  
However, the Harness docs assume everyone uses cloud clusters (EKS, GKE, AKS).  
This made connecting the delegate very difficult.  
Errors like *“Cannot invoke java.util.Map.getOrDefault”* or *“Unauthorized”* were undocumented.

**Suggested Fix:**  
Add a dedicated “Local Setup Guide” that explains:
- How to connect a Docker-based delegate to a Kind/Minikube cluster
- How to mount `~/.kube/config`
- Common failure cases and how to debug them  
(For example, localhost vs host.docker.internal connectivity issues.)

### 🔹 6. Documentation Tone — Concept-Heavy, Not Task-Driven
The docs spend more time describing *what* a connector or stage is, rather than *how* to set one up.  
As a new user, I didn’t want definitions — I wanted a working example first, then explanations later.

**Suggested Fix:**  
Lead with practical steps (“do this”) and follow with explanations (“here’s why this works”).  
In documentation design, this is called the **“show, then tell”** model — it keeps users engaged and helps them learn by doing.

### 🔹 7. Missing FAQ or Troubleshooting Section
During this project, I faced multiple issues:
- Delegate connection failures
- Kubeconfig authentication errors
- YAML validation errors in connectors
- Harness rejecting extra fields like `spec` under `InheritFromDelegate`
None of these were documented or easy to search for.

**Suggested Fix:**  
Create a **FAQ / Troubleshooting Guide** for common errors, e.g.:
- “Delegate cannot connect to K8s cluster”
- “Unauthorized when testing connector”
- “Cannot invoke java.util.Map.getOrDefault”
Each FAQ should include cause, resolution steps, and a working example.

### 🔹 8. Documentation Should Be Task-Oriented
Right now, the docs feel like a product encyclopedia.  
But developers approach docs with a *goal*, not to learn the product architecture.

**Suggested Fix:**  
Reorganize the docs by **tasks**, not by entities:
| Goal | Link |
|------|------|
| Deploy my first app | Beginner tutorial |
| Connect to Kubernetes | Connector setup |
| Add an approval stage | Pipeline examples |
| Debug delegate issues | Troubleshooting guide |

### 🔹 9. Suggestion — Create a Real “Example Journey” Document
We recommend that Harness publishes a **Google Doc or interactive guide** that walks users through:
- Signing up
- Creating the project (auto-created but explain the defaults)
- Setting up all connectors in sequence
- Deploying a working example app  
It should follow the actual user journey and serve as a *realistic onboarding flow* — something new users can complete in under 1 hour.

### 🧩 Summary of Improvements

| Problem | Suggested Fix |
|----------|----------------|
| No linear start-to-finish guide | Add one complete tutorial |
| YAML/UI inconsistency | Always show both methods |
| No working repo | Provide a sample project |
| Inconsistent style | Follow a doc style guide |
| Missing local setup docs | Add Kind/Minikube examples |
| Concept-heavy tone | Use “show then tell” writing |
| No troubleshooting | Add FAQ section |
| Scattered flow | Reorganize docs by tasks |

### 🧠 Final Reflection
Harness is a powerful platform — but its documentation currently assumes DevOps familiarity.  
For developers, students, or DevRel users who are just getting started, the learning curve is steep.  
Good docs should make *the first win easy* — right now, that win takes too long.

With small structural and stylistic improvements, Harness can make its docs **dramatically more approachable, task-oriented, and rewarding for new users**.

---