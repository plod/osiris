package main

import (
	"context"

	endpoints "github.com/plod/osiris/pkg/endpoints/controller"
	"github.com/plod/osiris/pkg/kubernetes"
	"github.com/plod/osiris/pkg/version"
	"github.com/golang/glog"
)

func runEndpointsController(ctx context.Context) {
	glog.Infof(
		"Starting Osiris Endpoints Controller -- version %s -- commit %s",
		version.Version(),
		version.Commit(),
	)

	client, err := kubernetes.Client()
	if err != nil {
		glog.Fatalf("Error building kubernetes clientset: %s", err)
	}

	controllerCfg, err := endpoints.GetConfigFromEnvironment()
	if err != nil {
		glog.Fatalf(
			"Error retrieving endpoints controller configuration: %s",
			err,
		)
	}

	// Run the controller
	endpoints.NewController(controllerCfg, client).Run(ctx)
}
