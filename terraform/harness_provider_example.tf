provider "harness" {
  # set HARNESS_ACCOUNT, HARNESS_API_KEY as env vars
}

resource "harness_organization" "org" {
  identifier = "example-org"
  name       = "Example Org"
}

resource "harness_project" "proj" {
  identifier = "demo-project"
  name       = "Demo Project"
  org_id     = harness_organization.org.identifier
}

# Example: create Git connector
resource "harness_connector_github" "repo_conn" {
  identifier = "repo-connector"
  project_id = harness_project.proj.identifier
  org_id     = harness_organization.org.identifier
  url        = "https://github.com"
  connection_type = "Repo"
  token_ref = "account.GITHUB_TOKEN" # use secret manager
}
