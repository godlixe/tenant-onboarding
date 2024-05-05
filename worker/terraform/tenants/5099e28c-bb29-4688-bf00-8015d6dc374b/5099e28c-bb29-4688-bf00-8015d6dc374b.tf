
module "main" {
	source = "../../main"
  
	tenant_name = "5099e28c-bb29-4688-bf00-8015d6dc374b"
	tenant_password = "123root987"
  
  }
  
  terraform {
	backend "gcs"{
	  credentials = "/home/alex/project/golang/tenant-onboarding/creds/static-booster-418207-f03f2fb4355d.json"
	  bucket = "terraform-dep"
	  prefix = "5099e28c-bb29-4688-bf00-8015d6dc374b"
	}
  }
