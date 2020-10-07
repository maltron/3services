INGRESS_GATEWAY_URL=$(oc get route/istio-ingressgateway --output=jsonpath='http://{.spec.host}' --namespace istio-system)
while true; do curl ${INGRESS_GATEWAY_URL}; echo; sleep 1.5s; done
