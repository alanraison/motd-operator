package controller

import (
	"github.com/alanraison/motd-operator/pkg/controller/motdsource"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, motdsource.Add)
}
