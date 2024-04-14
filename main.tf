provider "google" {
    project = var.project
    region = var.region
}

resource "google_sql_database_instance" "my-database" {
  name             = "my-database"
  database_version = "MYSQL_5_7"
  region           = "us-central1"
  deletion_protection = false

  settings {
    tier = "db-f1-micro"
  }
}

module "vpc" {
    source  = "terraform-google-modules/network/google"
    version = "~> 9.0"

    project_id   = var.project
    network_name = "example"
    routing_mode = "GLOBAL"

    subnets = [
        {
            subnet_name           = "subnet-01"
            subnet_ip             = "10.10.10.0/24"
            subnet_region         = "europe-west1"
        }
    ]
}

resource "google_integration_connectors_connection" "zendeskconnection" {
  name     = "test-zendesk"
  location = "us-central1"
  service_account = "${data.google_project.test_project.number}-compute@developer.gserviceaccount.com"
  connector_version = "projects/${data.google_project.test_project.project_id}/locations/global/providers/zendesk/connectors/zendesk/versions/1"
}

resource "google_cloudbuild_trigger" "filename-trigger" {
    name = "my-launcher"
    location = var.region
    project = var.project
    github {
        owner = "AlejandroPintosAlcarazo"
        name = "Cloud-Challenges"
        push {
            branch = "^main$"   
        }
    }
    filename = "cloudbuild.yaml"
}