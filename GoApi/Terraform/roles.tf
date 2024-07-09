resource "google_service_account" "service_account" {
  account_id   = "service-account"
  display_name = "My Service Account"
  project      = var.project
}
resource "google_project_iam_member" "allow_act_as_service_account" {
  project = "challenge-3-425717"
  role    = "roles/iam.serviceAccountUser"
  member  = "serviceAccount:530132527823@cloudbuild.gserviceaccount.com"
}
resource "google_project_iam_member" "run_admin" {
  project = "challenge-3-425717"
  role    = "roles/run.admin"
  member  = "serviceAccount:530132527823@cloudbuild.gserviceaccount.com"
}

