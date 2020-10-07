CURRENT_NAMESPACE=$(oc config view --minify --output=jsonpath='{..namespace}')
oc delete route/${CURRENT_NAMESPACE} --namespace istio-system
oc expose service/istio-ingressgateway --name=${CURRENT_NAMESPACE} --port=http2 --namespace istio-system
MICROSERVICE_URL=$(oc get route/${CURRENT_NAMESPACE} --output=jsonpath='{.spec.host}' --namespace istio-system)
oc apply --filename istio-way.yaml
oc patch gateway/microservice --type=merge --patch '"spec": {"servers": [{"hosts": ["'${MICROSERVICE_URL}'"],"port": {"name": "http","number": 80,"protocol": "HTTP"}}]}'
oc patch virtualservice/microservice --type=merge --patch '{"spec":{"hosts":["'${MICROSERVICE_URL}'"]}}'
echo "\n>> Service available on: curl http://${MICROSERVICE_URL}"
