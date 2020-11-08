resource "google_compute_network" "vpc" {
    name = "${var.project_id}-vpc"
    auto_create_subnetworks = false
}

resource "google_compute_firewall" "egress-web" {
    name = "${var.project_id}-egress-web"
    network = google_compute_network.vpc.name
    direction = "EGRESS"

    allow {
        protocol = "tcp"
        ports = [
            "443",
        ]
    }
}

resource "google_compute_subnetwork" "subnet" {
    name = "${var.project_id}-subnet"
    region = var.region
    network = google_compute_network.vpc.name
    ip_cidr_range = "10.10.0.0/24"
}
