resource "google_compute_network" "peering_network" {
  name = "peering-network"
}

resource "google_compute_subnetwork" "peering_subnet" {
  name          = "peering-network-subnet"
  region        = "europe-west1"
  network       = google_compute_network.peering_network.self_link
  ip_cidr_range = "10.2.0.0/28"
}

resource "google_vpc_access_connector" "connector" {
  name          = "peering-tool"
  subnet {
    name = google_compute_subnetwork.peering_subnet.name
  }
  region        = var.region
}

resource "google_compute_route" "default_route" {
  name       = "traffic-conector"
  dest_range = "0.0.0.0/0"
  network    = google_compute_network.peering_network.self_link

  next_hop_gateway = "global/gateways/default-internet-gateway"
  priority         = 1000
}

resource "google_compute_global_address" "private_ip_alloc" {
  name          = "private-ip-alloc"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 16
  network       = google_compute_network.peering_network.self_link
}

resource "google_service_networking_connection" "default" {
  network                 = google_compute_network.peering_network.self_link
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.private_ip_alloc.name]
}
