module "main" {
  source = "./main"

  service_account_file_path = var.service_account_file_path
  project_id                = var.project_id
  region                    = var.region
  tenant_name               = var.tenant_name
  tenant_subdomain          = var.tenant_subdomain
  tenant_password           = var.tenant_password

}

terraform {
  backend "gcs" {
    credentials = service_account_file_path
  }
}
