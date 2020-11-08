resource "google_dns_managed_zone" "dns" {
    name = "${var.project_id}-dns-zone"
    dns_name = "${var.domain}."
}

resource "google_dns_record_set" "reddit" {
    name = "reddit.${google_dns_managed_zone.dns.dns_name}"
    type = "A"
    ttl = 300
    managed_zone = google_dns_managed_zone.dns.name

    rrdatas = [google_container_cluster.primary.endpoint]
}
