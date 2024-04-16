package worker

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
	  bucket = "terraform-dep"
	  prefix = "%v"
	}
  }
`

func Deploy(ctx context.Context, tenantData *entity.Tenant) {

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
			os.Exit(1)
		}
	}

	// create file with templated content for tenant
	f, err := os.Create(filepath.Join(tenantDirPath, fmt.Sprintf("%v.tf", tenantId)))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf(tfTemplate, tenantId, tenantId))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	tf, err := tfexec.NewTerraform(tenantDirPath, "/usr/bin/terraform")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = tf.Init(ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = tf.Apply(ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
