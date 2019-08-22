resource "google_compute_instance" "vm_master" {
  name = "master"
  machine_type = "n1-highcpu-2"
  tags = [
    "master",
    "k8"
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

  metadata = {
    ssh-keys = "duongha:${file("../env/id_rsa.pub")}"
  }

}

resource "google_compute_firewall" "access_k8s" {
  name = "allow-access-k8s"
  network = "default"

  allow {
    protocol = "tcp"
    ports = ["6443"]
  }

  source_ranges = ["0.0.0.0/0"]
  target_tags = ["master"]

}

resource "google_compute_instance" "vm_worker1" {
  name = "worker1"
  machine_type = "f1-micro"
  tags = [
    "worker",
    "dev"
  ]

  boot_disk {
    initialize_params {
      image = "ubuntu-1804-lts"
    }
  }

  metadata = {
    ssh-keys = "duongha:${file("../env/id_rsa.pub")}"
  }

  network_interface {
    network = "default"
    access_config {}
  }
}
