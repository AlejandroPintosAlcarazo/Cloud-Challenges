
resource "google_dns_managed_zone" "peering_zone" {
  name        = "private-zone"
  dns_name    = "example.internal."
  description = "Private DNS zone"
  visibility  = "private"

  private_visibility_config {
    networks {
      network_url = google_compute_network.peering_network.self_link
    }
  }
}

resource "google_dns_record_set" "private_sql" {
  depends_on = [google_dns_managed_zone.peering_zone]
  name         = google_dns_managed_zone.peering_zone.dns_name
  type         = "A"
  ttl          = 300
  managed_zone = google_dns_managed_zone.peering_zone.name
  rrdatas      = [google_sql_database_instance.my-database.private_ip_address]
}