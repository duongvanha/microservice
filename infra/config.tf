provider "google" {
  credentials = file("../env/microservices.json")

  project = var.project
  region = var.region
  zone = var.zone

}

terraform {
  backend "remote" {
    organization = "microservices"

    workspaces {
      name = "Workspace1"
    }
  }
}
