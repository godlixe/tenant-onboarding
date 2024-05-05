
module "main" {
	source = "../../main"
  
	tenant_name = "bd363d21-9fa5-49d9-ac91-530bfa59f623"
	tenant_password = "123root987"
  
  }
  
  terraform {
	backend "gcs"{
	  credentials = "/home/alex/project/golang/tenant-onboarding/creds/static-booster-418207-f03f2fb4355d.json"
	  bucket = "terraform-dep"
	  prefix = "bd363d21-9fa5-49d9-ac91-530bfa59f623"
	}
  }
