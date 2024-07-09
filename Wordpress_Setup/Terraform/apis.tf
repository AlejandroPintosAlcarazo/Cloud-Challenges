resource "google_project_service" "compute_api" {
  project = "my-second-project-420603"
  service = "compute.googleapis.com"
}
