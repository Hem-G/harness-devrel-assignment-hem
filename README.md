# Harness DevRel Assignment — Deploy microservice to Dev/QA/Prod

This repository contains a minimal microservice, Kubernetes manifests, and example Harness pipeline definitions and automation snippets to satisfy the DevRel exercise.

## Goals (assignment)
- Create a Harness pipeline to deploy a microservice to Dev, QA and Prod K8s environments.
- Prod deployment must have approval gates.
- Support configurable deployment strategy (Rolling, Canary, Blue/Green).
- Provide pipeline design, execution steps, running pod verification, deployment strategy explanation, automation approaches, and feedback.

## What’s in this repo
- `app/service/` — minimal Go service + Dockerfile
- `k8s/` — base and overlays (dev/qa/prod) using Kustomize
- `helm-chart/myservice` — optional Helm chart
- `harness/` — example Harness entities as YAML (pipeline, service, environments)
- `terraform/` — example harness terraform provider snippet to create a project / connector
- `docs/design.md` — design explanation, verification steps, automation doc

## Quick local test (build and run)
1. Build docker: `docker build -t ghcr.io/<your-org>/myservice:0.1.0 ./app/service`
2. Push to your registry.
3. Apply k8s dev overlay locally: `kubectl apply -k k8s/overlays/dev`

4. Verify pod:
                `kubectl get pods -l app=myservice -n dev`
                `kubectl logs -l app=myservice -n dev`
                `curl http://<service-cluster-ip-or-port-forward>/health`


## Harness usage (manual flow)
1. Create a Harness project (or use Terraform / CLI shown below).
2. Create connectors:
- Git connector pointing at this repo
- K8s cluster connector(s) for Dev/QA/Prod
- Docker registry connector
3. Create Harness Service: reference manifests (k8s/ or helm-chart)
4. Create Environments: Dev, QA, Prod (map them to respective clusters/namespaces)
5. Create CD Pipeline (see `harness/pipeline.yaml`) — includes stages:
- Dev stage (deploy)
- QA stage (deploy)
- Prod stage (Approval manual step → deploy)
6. Run the Pipeline from Harness UI or via API.

## Automation (examples)
- Use Harness Terraform provider to create Projects, Connectors, Services, Pipelines.
- Use Harness CLI (`harness` open-source CLI) to apply YAML manifests to your account.
- Use Harness REST APIs to create and trigger pipelines programmatically.
See `terraform/harness_provider_example.tf` and `harness/` folder for examples.

## Verification checklist to submit
- Screenshot of pipeline design (UI) — take from your Harness account.
- Pipeline execution logs — get from Harness UI or API.
- `kubectl get pods -n <env>` showing the running pod.
- Pod logs & `curl` output from the service endpoint.

## References
- Harness CD tutorial: Deploy your own microservice app. https://developer.harness.io/docs/continuous-delivery/get-started/cd-tutorials/ownapp/ :contentReference[oaicite:0]{index=0}
- Harness Terraform provider docs & quickstarts. :contentReference[oaicite:1]{index=1}
- Harness CLI docs & examples. :contentReference[oaicite:2]{index=2}

