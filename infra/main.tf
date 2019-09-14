resource "google_container_cluster" "primary" {
  provider = "google-beta"
  name = "my-gke-cluster"
  location = var.region

  initial_node_count = 1

  network_policy {
    enabled = true
  }

  node_config {
    preemptible = true
    machine_type = "n1-standard-4"

    metadata = {
      disable-legacy-endpoints = "true"
    }
  }

  //  addons_config {
  //    istio_config {
  //      disabled = false
  //    }
  //
  //    kubernetes_dashboard {
  //      disabled = false
  //    }
  //  }

  logging_service = "none"

  monitoring_service = "none"

}

//resource "google_container_node_pool" "node-1" {
//  project = var.project
//  name = "node-1"
//  location = "us-central1"
//  cluster = google_container_cluster.primary.name
//  initial_node_count = 1
//
//  node_config {
//    preemptible = true
//    machine_type = "n1-standard-1"
//
//    metadata = {
//      disable-legacy-endpoints = "true"
//    }
//
//  }
//}

resource "null_resource" "apply" {

  depends_on = [
    //    google_container_node_pool.node-1,
    google_container_cluster.primary
  ]

  provisioner "local-exec" {
    command = <<EOF
      gcloud beta container clusters get-credentials ${google_container_cluster.primary.name} --region ${google_container_cluster.primary.region} --project ${google_container_cluster.primary.project}
      kubectl apply -f tiller-rbac.yaml
      kubectl create clusterrolebinding cluster-admin-binding --clusterrole=cluster-admin --user=$(gcloud config get-value core/account)
      rm -rf ~/.helm
      helm init --service-account haduong
      cd istio
      kubectl create namespace istio-system
      helm template install/kubernetes/helm/istio-init --name istio-init --namespace istio-system | kubectl apply -f -
      ./../await.sh
      helm template install/kubernetes/helm/istio --name istio --namespace istio-system | kubectl apply -f -
      kubectl apply -f ./../services/node-grafana.yaml
      kubectl apply -f ./../services/elk
      kubectl create namespace microservice
      kubectl label namespace microservice istio-injection=enabled
    EOF
  }
}
