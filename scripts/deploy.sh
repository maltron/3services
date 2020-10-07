for deployment in A/deployment-A-blue.yaml A/deployment-A-green.yaml A/service.yaml B/deployment-B.yaml C/deployment-C.yaml; do oc apply --filename ${deployment}; done
