provider "google" {
    project = var.project
    region = var.region
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
    substitutions = {
    _VPC_CONNECTOR = google_vpc_access_connector.connector.name
    }
}