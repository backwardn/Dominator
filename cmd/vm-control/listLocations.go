package main

import (
	"fmt"
	"os"

	"github.com/Symantec/Dominator/lib/errors"
	"github.com/Symantec/Dominator/lib/log"
	"github.com/Symantec/Dominator/lib/srpc"
	proto "github.com/Symantec/Dominator/proto/fleetmanager"
)

func listLocationsSubcommand(args []string, logger log.DebugLogger) {
	var topLocation string
	if len(args) > 0 {
		topLocation = args[0]
	}
	if err := listLocations(topLocation, logger); err != nil {
		fmt.Fprintf(os.Stderr, "Error listing locations: %s\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}

func listLocations(topLocation string, logger log.DebugLogger) error {
	fleetManager := fmt.Sprintf("%s:%d",
		*fleetManagerHostname, *fleetManagerPortNum)
	client, err := srpc.DialHTTP("tcp", fleetManager, 0)
	if err != nil {
		return err
	}
	defer client.Close()
	request := proto.ListHypervisorLocationsRequest{topLocation}
	var reply proto.ListHypervisorLocationsResponse
	err = client.RequestReply("FleetManager.ListHypervisorLocations",
		request, &reply)
	if err != nil {
		return err
	}
	if err := errors.New(reply.Error); err != nil {
		return err
	}
	for _, location := range reply.Locations {
		if _, err := fmt.Println(location); err != nil {
			return err
		}
	}
	return nil
}