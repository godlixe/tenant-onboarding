package worker

import (
	"fmt"
	"os"

	"github.com/hashicorp/terraform-exec/tfexec"
)

func deploy() {
	tf, err := tfexec.NewTerraform("../terraform", "/usr/bin/terraform")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(tf)
}
