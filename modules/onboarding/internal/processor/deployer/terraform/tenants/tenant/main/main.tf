provider "google" {
  credentials = file(var.service_account_file_path)
  project     = var.project_id
  region      = var.region
}

module "storage" {
  source = "../modules/storage"

  storage_instance_name = var.tenant_name
  storage_user_password = var.tenant_password
}

module "compute" {
  source = "../modules/compute"

  compute_name     = var.tenant_name
  storage_host     = module.storage.host
  storage_port     = module.storage.port
  storage_name     = module.storage.name
  storage_user     = module.storage.user
  storage_password = var.tenant_password
}
