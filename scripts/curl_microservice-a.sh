curl $(oc get route/microservice-a --output=jsonpath='http://{.spec.host}')
