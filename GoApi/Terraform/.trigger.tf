resource "google_cloudbuild_trigger" "filename-trigger" {
    name = "ASTEROID"
    location = var.region
    project = var.project
    github {
        owner = "AlejandroPintosAlcarazo"
        name = "Asteroid.api"
        push {
            branch = "^main$"   
        }
    }
    filename = "cloudbuild.yaml"
}

