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

package main

import (
	"fmt"
	"github.com/msjelly/azcmd/pkg/azservice"
	"github.com/msjelly/azcmd/pkg/azservice/trafficmanager"
	"log"
	"os"
)

func main() {

	exitCode := azservice.NoError

	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)

	if len(os.Args) < 2 {
		err := fmt.Errorf("Missing service name")
		log.Printf("Error. Insufficient args. %s", err.Error())
	}

	service := os.Args[1]

	switch service {
	case "trafficmanager-endpoint":
		worker := &trafficmanager.EndpointWorker{}
		exitCode = worker.Work(os.Args[2:])
	default:
		exitCode = azservice.UsageError
	}

	os.Exit(int(exitCode))
}
