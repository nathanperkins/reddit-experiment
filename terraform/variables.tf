variable "project_id" {
    description = "project id"
}

variable "region" {
    description = "region"
}

variable "domain" {
    description = "domain to assign to cluster"
}

variable "gke_location" {
    description = "gke cluster location"
}

variable "gke_username" {
    default = ""
    description = "gke username"
}

variable "gke_password" {
    default = ""
    description = "gke password"
}

variable "gke_num_nodes" {
    default = 1
    description = "gke number of nodes in cluster"
}
