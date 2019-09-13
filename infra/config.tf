provider "google-beta" {
  credentials = file("../env/microservice.json")

  project = var.project
  region = var.region
  zone = var.zone

}

//terraform {
//  backend "remote" {
//    organization = "microservices"
//
//    workspaces {
//      name = "microservice"
//    }
//  }
//}
