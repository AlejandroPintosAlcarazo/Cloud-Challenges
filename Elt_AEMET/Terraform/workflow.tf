resource "google_workflows_workflow" "orchestrator" {
  name     = "orchestrator"
  project  = var.project_id
  region   = var.region

  description = "Orchestrates the data pipeline"
  source_contents = <<-EOT
    main:
      steps:
      - step1:
          call: http.get
          args:
            url: "${google_cloud_run_service.connector.status[0].url}"
          result: response1
      - step2:
          switch:
            - condition: response1.status_code == 200
              steps:
                - call: http.post
                  args:
                    url: "https://dataform.googleapis.com/v1beta1/projects/${var.project_id}/locations/${var.region}/repositories/${var.dataform_repository}:run"
                    body:
                      run: {}
                    auth:
                      type: OIDC
                  result: response2
  EOT
}

