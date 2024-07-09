resource "google_cloud_run_service" "connector" {
  name     = "connector"
  location = var.region
  project  = var.project_id

  template {
    spec {
      containers {
        image = var.connector_image
        ports {
          container_port = 8080
        }
      }
    }
  }
}

# IAM binding for Cloud Run to allow Cloud Scheduler to invoke it
resource "google_cloud_run_service_iam_binding" "invoker" {
  location    = google_cloud_run_service.connector.location
  project     = google_cloud_run_service.connector.project
  service     = google_cloud_run_service.connector.name
  role        = "roles/run.invoker"
  members     = [
    "serviceAccount:${google_service_account.sa.email}"
  ]
}

# Output for the Cloud Run service URL
output "connector_url" {
  value = google_cloud_run_service.connector.status[0].url
}

