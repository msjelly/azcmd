/*
Copyright 2019 Microsoft.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package trafficmanager

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/services/trafficmanager/mgmt/2018-04-01/trafficmanager"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/msjelly/azcmd/pkg/azservice"
	"log"
)

// EndpointWorker handles trafficmanager endpoint commands
type EndpointWorker struct {
	subcmd string
	config map[EndpointConfigKey]string
	client *trafficmanager.EndpointsClient
}

// Work handles TrafficManager Endpoint requests for azcmd.
func (worker EndpointWorker) Work(args []string) azservice.ExitCode {

	exitCode := azservice.NoError
	if err := worker.parseArgs(args); err != nil {
		log.Printf("Error. Failed to parse cmdline args. %s", err.Error())
		return azservice.UsageError
	}

	if err := worker.getEndpointsClient(); err != nil {
		log.Printf("Error. Failed to get client. %s", err.Error())
		return azservice.ClientSetupError
	}

	var err error
	switch worker.subcmd {
	case "create":
		if err = worker.createOrUpdateEndpoint(); err != nil {
			exitCode = azservice.CreateError
		}
	case "delete":
		if err = worker.deleteEndpoint(); err != nil {
			exitCode = azservice.DeleteError
		}
	default:
		{
			err = fmt.Errorf("Invalid subcmd %s", worker.subcmd)
			exitCode = azservice.UsageError
		}
	}

	if err != nil {
		log.Printf("Azcmd failed. subcmd %s. %s", worker.subcmd, err.Error())
	}
	return exitCode
}

func (worker *EndpointWorker) getEndpointsClient() error {

	authorizer, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		log.Printf("Failed NewAuthorizerFromEnvironment. authorizer %+v. %s", authorizer, err.Error())
		return err
	}

	client := trafficmanager.NewEndpointsClient(worker.config[EpConfigSubscriptionID])
	client.Authorizer = authorizer
	worker.client = &client

	return nil
}

func (worker *EndpointWorker) createOrUpdateEndpoint() error {

	tmepProperties := &trafficmanager.EndpointProperties{
		Target:           to.StringPtr(worker.config[EpConfigEndpointIPAddress]),
		EndpointLocation: to.StringPtr(worker.config[EpConfigLocation]),
	}

	result, err := worker.client.CreateOrUpdate(context.TODO(),
		worker.config[EpConfigResourceGroup], worker.config[EpConfigProfileName],
		worker.config[EpConfigEndpointType], worker.config[EpConfigEndpointName],
		trafficmanager.Endpoint{
			EndpointProperties: tmepProperties,
		},
	)
	if err == nil {
		log.Printf("Created endpoint. TrafficManagerEndpoint %+v", result)
	} else {
		log.Printf("Failed to CreatOrUpdate endpoint. %s", err.Error())
	}
	return err
}

func (worker *EndpointWorker) deleteEndpoint() error {

	result, err := worker.client.Delete(context.TODO(),
		worker.config[EpConfigResourceGroup], worker.config[EpConfigProfileName],
		worker.config[EpConfigEndpointType], worker.config[EpConfigEndpointName])
	if err == nil {
		log.Printf("Deleted endpoint. TrafficManagerEndpoint %+v", result)
	} else {
		log.Printf("Failed to delete endpoint. %s. %+v", err.Error(), result)
	}
	return err
}
