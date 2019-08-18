output "ip_master" {
  value = google_compute_instance.vm_master.network_interface.0.access_config.0.nat_ip
}


output "ip_worker1" {
  value = google_compute_instance.vm_worker1.network_interface.0.access_config.0.nat_ip
}
