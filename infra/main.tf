resource "google_compute_instance" "vm_instance" {
  name = "instance1"
  machine_type = "f1-micro"
  tags = [
    "web",
    "dev"
  ]

  boot_disk {
    initialize_params {
      image = "ubuntu-1804-lts"
    }
  }


  network_interface {
    network = "default"
    access_config {}
  }
}
