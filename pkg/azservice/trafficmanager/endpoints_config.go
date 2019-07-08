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
	"fmt"
	"log"
)

// EndpointConfigKey is a label for Endpoint config keys
type EndpointConfigKey string

// These are valid values for EndpointConfigKey
const (
	EpConfigSubscriptionID    EndpointConfigKey = "SubscriptionID"
	EpConfigResourceGroup     EndpointConfigKey = "ResourceGroup"
	EpConfigProfileName       EndpointConfigKey = "ProfileName"
	EpConfigEndpointType      EndpointConfigKey = "EndpoinType"
	EpConfigEndpointName      EndpointConfigKey = "EndpointName"
	EpConfigEndpointIPAddress EndpointConfigKey = "EndpointIPAddress"
	EpConfigLocation          EndpointConfigKey = "Location"
)

// EndpointConfigKeyList is a list of keys for TrafficManagerEndpoing config
// Note EpConfigLocation is valid and required arg when the Profile routing method is Performance.
// FIXME Add option for Profile with either Performance or Weighed routing methods.
var EndpointConfigKeyList = []EndpointConfigKey{
	EpConfigSubscriptionID,
	EpConfigResourceGroup,
	EpConfigProfileName,
	EpConfigEndpointName,
	EpConfigEndpointIPAddress,
	EpConfigLocation,
}

var mandatoryCreateEpArgs = []EndpointConfigKey{
	EpConfigSubscriptionID,
	EpConfigResourceGroup,
	EpConfigProfileName,
	EpConfigEndpointName,
	EpConfigEndpointIPAddress,
	EpConfigLocation,
}

var deleteEpArgs = []EndpointConfigKey{
	EpConfigSubscriptionID,
	EpConfigResourceGroup,
	EpConfigProfileName,
	EpConfigEndpointName,
}

func (worker *EndpointWorker) parseArgs(cmdargs []string) error {

	if len(cmdargs) < 1 {
		err := fmt.Errorf("Usage - missing subcmd")
		return err
	}

	worker.subcmd = cmdargs[0]
	args := cmdargs[1:]
	nargs := len(args)
	switch worker.subcmd {
	case "create":
		if nargs < len(mandatoryCreateEpArgs) {
			err := fmt.Errorf("Usage - missing mandatory args %#v", mandatoryCreateEpArgs)
			return err
		}
	case "delete":
		if nargs != len(deleteEpArgs) {
			err := fmt.Errorf("Usage - expected args %#v", deleteEpArgs)
			return err
		}
	}

	worker.config = make(map[EndpointConfigKey]string)
	for i, key := range EndpointConfigKeyList {
		if nargs >= (i + 1) {
			worker.config[key] = args[i]
			log.Printf("Parsed cmdline [%d]%s:%s", i, key, worker.config[key])
		}
	}

	worker.config[EpConfigEndpointType] = "externalEndpoints"
	return nil
}
