resource "google_composer_environment" "composer_environment" {
  name   = "composer-environment"
  region = var.region
  config {
    software_config {
      image_version = var.composer_image_version
    }
  }
}

