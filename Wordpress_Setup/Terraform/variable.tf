variable "region" {
    type = string
    default = "europe-west1"
}
variable "project" {
    type = string
    default = "fluted-oath-420013"
}
#variable "project" {
#    type = string
#    default = "my-second-project-420603"
#}
variable "account" {
    type = string
    default = "alejandropintosdev@gmail.com"
}
variable "cloud_run_service_account" {
  description = "The service account used by Cloud Run to interact with other GCP services."
}