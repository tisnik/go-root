package main

import (
	"github.com/mingrammer/cfmt"
)

func main() {
	cfmt.Success("Success message\n")
	cfmt.Info("Info message\n")
	cfmt.Warning("Warning message\n")
	cfmt.Error("Error message\n")
}
