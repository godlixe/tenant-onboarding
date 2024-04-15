provider "google" {
  credentials = file("../../static-booster-418207-f03f2fb4355d.json")
  project     = "static-booster-418207"
  region      = "asia-southeast2"
}

module "storage" {
  source = "../modules/storage"
  
  storage_instance_name = var.tenant_name
  storage_user_password = var.tenant_password
}

module "compute" {
  source = "../modules/compute"

  compute_name = var.tenant_name
  storage_host = module.storage.host
  storage_port = module.storage.port
  storage_name = module.storage.name
  storage_user = module.storage.user
  storage_password = var.tenant_password
}
