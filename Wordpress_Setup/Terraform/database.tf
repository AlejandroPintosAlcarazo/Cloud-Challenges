resource "google_sql_database_instance" "my-database" {
  name             = "my-database"
  database_version = "MYSQL_8_0"
  region           = "europe-west1"
  deletion_protection = false
  settings {
    tier = "db-f1-micro"
    ip_configuration {
        ipv4_enabled = false
        private_network   =  google_compute_network.peering_network.self_link # self_link
    }
  }
  depends_on = [google_vpc_access_connector.connector]
}

resource "google_sql_database" "database" {
  name     = "my-database"
  instance = google_sql_database_instance.my-database.name
}

resource "google_sql_user" "sql-users" {
  name     = "user"
  instance = google_sql_database_instance.my-database.name
  password = "password"
}