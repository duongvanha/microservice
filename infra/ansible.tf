data "template_file" "dev_hosts" {
  template = file("ansible/hosts.tpl")

  vars = {
    master_ip = google_compute_instance.vm_master.network_interface.0.access_config.0.nat_ip
    worker1_ip = google_compute_instance.vm_worker1.network_interface.0.access_config.0.nat_ip
  }
}

resource "null_resource" "dev-hosts" {
  triggers = {
    template_rendered = data.template_file.dev_hosts.rendered
  }

  provisioner "local-exec" {
    // https://buddy.works/blog/new-feature-passing-parameters
    command = "echo \"${data.template_file.dev_hosts.rendered}\" >> hosts"
  }

}
