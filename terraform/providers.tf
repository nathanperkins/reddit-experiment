
terraform {
    required_version= ">= 0.12"
    required_providers {
        google = {
            source = "hashicorp/google"
        }
    }
}

provider "google" {
    region = var.region
    project = var.project_id
}
