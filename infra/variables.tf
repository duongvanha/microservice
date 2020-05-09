variable "project" {
  default = "mythical-height-252821"
}

variable "name_cluster" {
  default = "my-gke-cluster"
  type = string
}

variable "region" {
  default = "us-east1"
  type = string
}

variable "zone" {
  default = "us-east1-b"
  type = string
}
