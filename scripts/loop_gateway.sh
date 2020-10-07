MICROSERVICE_URL=$(oc get route/$(oc config view --minify --output=jsonpath='{..namespace}') --output=jsonpath='http://{.spec.host}' --namespace istio-system)
while true; do curl ${MICROSERVICE_URL}; echo; sleep 1.5s; done
