package deployer

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"tenant-onboarding/internal/domain/users/entity"

	"github.com/hashicorp/terraform-exec/tfexec"
)

const tfTemplate string = `
module "main" {
	source = "../../main"
  
	tenant_name = "%v"
	tenant_password = "123root987"
  
  }
  
  terraform {
	backend "gcs"{
	  credentials = "/home/alex/project/golang/tenant-onboarding/creds/static-booster-418207-f03f2fb4355d.json"
	  bucket = "terraform-dep"
	  prefix = "%v"
	}
  }
`

func Deploy(ctx context.Context, tenantData *entity.Tenant) error {
	var err error

	tenantId := tenantData.ID.String()
	tenantDirPath := filepath.Join(
		os.Getenv("TF_ABSOLUTE_PATH"),
		"tenants",
		tenantId,
	)
	// create dir for new tenant
	if _, err := os.Stat(tenantId); os.IsNotExist(err) {
		err := os.Mkdir(tenantDirPath, fs.ModePerm)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	// create file with templated content for tenant
	f, err := os.Create(filepath.Join(tenantDirPath, fmt.Sprintf("%v.tf", tenantId)))
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf(tfTemplate, tenantId, tenantId))
	if err != nil {
		fmt.Println(err)
		return err
	}

	tf, err := tfexec.NewTerraform(tenantDirPath, "/usr/bin/terraform")
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = tf.Init(ctx)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("tf-apply")
	err = tf.Apply(ctx)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
