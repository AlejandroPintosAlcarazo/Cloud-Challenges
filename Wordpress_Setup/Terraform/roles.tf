resource "google_service_account" "service_account" {
  account_id   = "service-account"
  display_name = "My Service Account"
  project      = var.project
}

resource "google_project_iam_member" "cloud_sql_admin" {
  project = var.project
  role    = "roles/cloudsql.admin"
  member  = "serviceAccount:${google_service_account.service_account.email}"
}