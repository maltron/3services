MICROSERVICE_URL=$(oc get route/microservice-a --output=jsonpath='http://{.spec.host}')
while true; do curl ${MICROSERVICE_URL}; echo; sleep 1.5s; done
