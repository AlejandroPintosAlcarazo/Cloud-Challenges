variable "project_id" {
  description = "The ID of the GCP project"
  type        = string
}

variable "region" {
  description = "The region where resources will be deployed"
  type        = string
  default     = "us-central1"
}


variable "access_token" {
  description = "Access token for Google APIs"
  type        = string
}

variable "dataform_repository" {
  description = "The name of the Dataform repository"
  type        = string
}
variable "connector_image" {
  description = "The Docker image for the Cloud Run connector"
  type        = string
}

