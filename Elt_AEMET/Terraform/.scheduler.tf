resource "google_pubsub_topic" "orchestrator_topic" {
  name = "orchestrator-topic"
}

resource "google_cloud_scheduler_job" "daily_launcher" {
  name        = "daily-launcher"
  schedule    = "0 1 * * *"
  time_zone   = "Etc/UTC"
  pubsub_target {
    topic_name = google_pubsub_topic.orchestrator_topic.id
    data       = base64encode("Trigger Orchestrator")
  }
}

