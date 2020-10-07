oc new-project ${1}
oc delete limitrange/${1}-core-resource-limits
oc label namespace/${1} istio-injection=enabled
oc adm policy add-scc-to-group privileged system:serviceaccounts:${1}
oc adm policy add-scc-to-group anyuid system:serviceaccounts:${1}
cat <<EOF | oc -n ${1} create -f -
apiVersion: "k8s.cni.cncf.io/v1"
kind: NetworkAttachmentDefinition
metadata:
  name: istio-cni
EOF
