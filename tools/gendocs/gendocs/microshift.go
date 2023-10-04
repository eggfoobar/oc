package gendocs

import (
	"k8s.io/apimachinery/pkg/util/sets"
)

// a set of commands excluded from microshift
var microshiftCommands = sets.NewString(
	"oc adm build-chain",
	"oc adm certificate approve",
	"oc adm certificate deny",
	"oc adm cordon",
	"oc adm create-bootstrap-project-template",
	"oc adm create-error-template",
	"oc adm create-login-template",
	"oc adm create-provider-selection-template",
	"oc adm drain",
	"oc adm groups add-users",
	"oc adm groups new",
	"oc adm groups prune",
	"oc adm groups remove-users",
	"oc adm groups sync",
	"oc adm migrate template-instances",
	"oc adm must-gather",
	"oc adm new-project",
	"oc adm node-logs",
	"oc adm pod-network isolate-projects",
	"oc adm pod-network join-projects",
	"oc adm pod-network make-projects-global",
	"oc adm policy add-role-to-user",
	"oc adm policy add-scc-to-group",
	"oc adm policy add-scc-to-user",
	"oc adm policy scc-review",
	"oc adm policy scc-subject-review",
	"oc adm prune builds",
	"oc adm prune deployments",
	"oc adm prune groups",
	"oc adm prune images",
	"oc adm release new",
	"oc adm top images",
	"oc adm top imagestreams",
	"oc adm top node",
	"oc adm uncordon",
	"oc adm upgrade",
	"oc adm verify-image-signature",
	"oc autoscale",
	"oc cancel-build",
	"oc create build",
	"oc create clusterresourcequota",
	"oc create deploymentconfig",
	"oc create identity",
	"oc create imagestream",
	"oc create imagestreamtag",
	"oc create user",
	"oc create useridentitymapping",
	"oc debug",
	"oc idle",
	"oc image mirror",
	"oc import-image",
	"oc login",
	"oc logout",
	"oc new-app",
	"oc new-build",
	"oc new-project",
	"oc process",
	"oc project",
	"oc projects",
	"oc registry info",
	"oc registry login",
	"oc replace",
	"oc rollback",
	"oc rollout cancel",
	"oc rollout history",
	"oc rollout latest",
	"oc rollout pause",
	"oc rollout restart",
	"oc rollout resume",
	"oc rollout retry",
	"oc rollout status",
	"oc rollout undo",
	"oc set build-hook",
	"oc set build-secret",
	"oc set deployment-hook",
	"oc set image-lookup",
	"oc set triggers",
	"oc start-build",
	"oc status",
	"oc tag",
	"oc whoami",
)