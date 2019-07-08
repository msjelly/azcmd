# azcmd

Go implementation equivalent to Azure CLI commands:
   * az network traffic-manager endpoint create --profile &lt;profile-name> -g &lt;resource-group> -n &lt;ep-name> --type externalEndpoints --endpoint-location &lt;location> --taresource-groupet &lt;FQDN|IP>
   * az network traffic-manager endpoint delete --profile &lt;profile-name> -g &lt;resource-group> -n &lt;ep-name> --type externalEndpoints 

To build:
* standalone app - make azcmd
* containers -  make docker-build && make docker-push
* See Makefile.

To run on k8s cluster, see samples.
   * kubectl apply -f samples/aztmCreate.yaml
   * kubectl apply -f samples/aztmDelete.yaml
   * For authentication/authorization, see set aadpodidbinding.  For details, see [AAD Pod Identity](https://github.com/Azure/aad-pod-identity).

Future work:
Currently supports Profile with Performance routing method.  Add support for Profile with Weighted routing method.  For Endpoint in a Profile with Weighted routing method, Location is not a valid arg and Weight can be specified.