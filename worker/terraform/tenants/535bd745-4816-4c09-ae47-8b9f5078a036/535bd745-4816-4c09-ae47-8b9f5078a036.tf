
module "main" {
	source = "../../main"
  
	tenant_name = "535bd745-4816-4c09-ae47-8b9f5078a036"
	tenant_password = "123root987"
  
  }
  
  terraform {
	backend "gcs"{
	  credentials = "/home/alex/project/golang/tenant-onboarding/creds/static-booster-418207-f03f2fb4355d.json"
	  bucket = "terraform-dep"
	  prefix = "535bd745-4816-4c09-ae47-8b9f5078a036"
	}
  }
