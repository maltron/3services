while true; do curl $(oc get route/microservice-a --output jsonpath='{.spec.port.targetPort}://{.spec.host}'); echo; sleep 1.5s; done