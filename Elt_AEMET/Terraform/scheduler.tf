resource "google_cloud_scheduler_job" "job" {
  name        = "launcher"
  description = "Launches the orchestrator workflow"
  schedule    = "0 1 * * *"  # A la 1:00 AM todos los días
  time_zone   = "Europe/Madrid"
  project     = var.project_id
  region      = var.region  # Cambiado a "region"

  http_target {
    http_method = "POST"
    uri         = "https://workflowexecutions.googleapis.com/v1/projects/${var.project_id}/locations/${var.region}/workflows/${google_workflows_workflow.orchestrator.name}/executions"
    oidc_token {
      service_account_email = google_service_account.sa.email
    }
  }
}

# Service Account
resource "google_service_account" "sa" {
  account_id   = "cloud-scheduler-sa"
  display_name = "Cloud Scheduler Service Account"
  project      = var.project_id
}

