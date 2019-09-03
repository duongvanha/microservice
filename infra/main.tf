resource "google_container_cluster" "primary" {
  provider = "google-beta"
  name = "my-gke-cluster"
  location = var.region

  remove_default_node_pool = true
  initial_node_count = 1

  network_policy {
    enabled = true
  }

  addons_config {
    istio_config {
      disabled = false
    }

    kubernetes_dashboard {
      disabled = false
    }
  }

  logging_service = "none"

  monitoring_service = "none"

}

resource "google_container_node_pool" "node-1" {
  project = var.project
  name = "node-1"
  location = var.region
  cluster = google_container_cluster.primary.name
  initial_node_count = 1

  node_config {
    preemptible = true
    machine_type = "n1-standard-2"

    metadata = {
      disable-legacy-endpoints = "true"
    }

  }
}
