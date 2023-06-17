# Change Log

## 0.33.1

**Release date:** 2023-06-15

![AppVersion: v0.25.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.25.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [alertmanager] Fix alertmanager's configmapReload schema (#3495)

### Default value changes

```diff
diff --git a/charts/alertmanager/values.yaml b/charts/alertmanager/values.yaml
index 09c5bfc2..27c96b5c 100644
--- a/charts/alertmanager/values.yaml
+++ b/charts/alertmanager/values.yaml
@@ -1,3 +1,4 @@
+# yaml-language-server: $schema=values.schema.json
 # Default values for alertmanager.
 # This is a YAML-formatted file.
 # Declare variables to be passed into your templates.

```

## 0.33.0

**Release date:** 2023-06-12

![AppVersion: v0.25.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.25.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [alertmanager] Support overriding namespace (#3475)

### Default value changes

```diff
diff --git a/charts/alertmanager/values.yaml b/charts/alertmanager/values.yaml
index 68db8634..09c5bfc2 100644
--- a/charts/alertmanager/values.yaml
+++ b/charts/alertmanager/values.yaml
@@ -24,6 +24,8 @@ extraSecretMounts: []
 imagePullSecrets: []
 nameOverride: ""
 fullnameOverride: ""
+## namespaceOverride overrides the namespace which the resources will be deployed in
+namespaceOverride: ""
 
 automountServiceAccountToken: true
 

```

## 0.32.1

**Release date:** 2023-06-07

![AppVersion: v0.25.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.25.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* handle special case to group_by all labels (#3469)

### Default value changes

```diff
# No changes in this release
```

## 0.32.0

**Release date:** 2023-06-07

![AppVersion: v0.25.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.25.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-alertmanager] add schedulerName value (#3463)

### Default value changes

```diff
diff --git a/charts/alertmanager/values.yaml b/charts/alertmanager/values.yaml
index ef380c86..68db8634 100644
--- a/charts/alertmanager/values.yaml
+++ b/charts/alertmanager/values.yaml
@@ -39,6 +39,9 @@ serviceAccount:
 # Sets priorityClassName in alertmanager pod
 priorityClassName: ""
 
+# Sets schedulerName in alertmanager pod
+schedulerName: ""
+
 podSecurityContext:
   fsGroup: 65534
 dnsConfig: {}

```

## 0.31.1

**Release date:** 2023-05-25

![AppVersion: v0.25.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.25.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [alertmanager] Values json schema (#3397)

### Default value changes

```diff
diff --git a/charts/alertmanager/values.yaml b/charts/alertmanager/values.yaml
index 7709c7e6..ef380c86 100644
--- a/charts/alertmanager/values.yaml
+++ b/charts/alertmanager/values.yaml
@@ -34,7 +34,7 @@ serviceAccount:
   annotations: {}
   # The name of the service account to use.
   # If not set and create is true, a name is generated using the fullname template
-  name:
+  name: ""
 
 # Sets priorityClassName in alertmanager pod
 priorityClassName: ""
@@ -51,7 +51,7 @@ dnsConfig: {}
   #   - name: ndots
   #     value: "2"
   #   - name: edns0
-hostAliases: {}
+hostAliases: []
   # - ip: "127.0.0.1"
   #   hostnames:
   #   - "foo.local"
@@ -78,7 +78,6 @@ extraInitContainers: []
 ## Additional containers to add to the stateful set. This will allow to setup sidecarContainers like a proxy to integrate
 ## alertmanager with an external tool like teams that has not direct integration.
 ##
-
 extraContainers: []
 
 livenessProbe:

```

## 0.31.0

**Release date:** 2023-05-23

![AppVersion: v0.25.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.25.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Add support to the helm chart to add sidecard containers to the alertmanager statefulset. Currently the only extracontainer allowed is the config reloader, but somtimes we may need add other containers like a proxy to integrate alertmanager with external tool because there is no direct integration (#3388)

### Default value changes

```diff
diff --git a/charts/alertmanager/values.yaml b/charts/alertmanager/values.yaml
index a4fe559e..7709c7e6 100644
--- a/charts/alertmanager/values.yaml
+++ b/charts/alertmanager/values.yaml
@@ -75,6 +75,12 @@ additionalPeers: []
 ##
 extraInitContainers: []
 
+## Additional containers to add to the stateful set. This will allow to setup sidecarContainers like a proxy to integrate
+## alertmanager with an external tool like teams that has not direct integration.
+##
+
+extraContainers: []
+
 livenessProbe:
   httpGet:
     path: /

```

## 0.30.1

**Release date:** 2023-05-03

![AppVersion: v0.25.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.25.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [alertmanager] Adjust configmap-reload port (#3270)

### Default value changes

```diff
# No changes in this release
```

## 0.30.0

**Release date:** 2023-05-01

![AppVersion: v0.25.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.25.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [alertmanager] - 3286 Enahance test connection template (#3294)

### Default value changes

```diff
diff --git a/charts/alertmanager/values.yaml b/charts/alertmanager/values.yaml
index f3c1f4e9..a4fe559e 100644
--- a/charts/alertmanager/values.yaml
+++ b/charts/alertmanager/values.yaml
@@ -251,3 +251,9 @@ extraVolumes: []
 extraEnv: []
   # - name: FOO
   #   value: BAR
+
+testFramework:
+  enabled: false
+  annotations:
+    "helm.sh/hook": test-success
+    # "helm.sh/hook-delete-policy": "before-hook-creation,hook-succeeded"

```

## 0.29.0

**Release date:** 2023-04-13

![AppVersion: v0.25.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.25.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [alertmanager] Adding hostAliases to Alertmanager chart. (#3209)

### Default value changes

```diff
diff --git a/charts/alertmanager/values.yaml b/charts/alertmanager/values.yaml
index 2c3aa7df..f3c1f4e9 100644
--- a/charts/alertmanager/values.yaml
+++ b/charts/alertmanager/values.yaml
@@ -51,6 +51,15 @@ dnsConfig: {}
   #   - name: ndots
   #     value: "2"
   #   - name: edns0
+hostAliases: {}
+  # - ip: "127.0.0.1"
+  #   hostnames:
+  #   - "foo.local"
+  #   - "bar.local"
+  # - ip: "10.1.2.3"
+  #   hostnames:
+  #   - "foo.remote"
+  #   - "bar.remote"
 securityContext:
   # capabilities:
   #   drop:

```

## 0.28.0

**Release date:** 2023-04-12

![AppVersion: v0.25.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.25.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [alertmanager] Add annotations for ConfigMap (#3203)

### Default value changes

```diff
diff --git a/charts/alertmanager/values.yaml b/charts/alertmanager/values.yaml
index a167f30a..2c3aa7df 100644
--- a/charts/alertmanager/values.yaml
+++ b/charts/alertmanager/values.yaml
@@ -168,6 +168,17 @@ persistence:
     - ReadWriteOnce
   size: 50Mi
 
+configAnnotations: {}
+  ## For example if you want to provide private data from a secret vault
+  ## https://github.com/banzaicloud/bank-vaults/tree/main/charts/vault-secrets-webhook
+  ## P.s.: Add option `configMapMutation: true` for vault-secrets-webhook
+  # vault.security.banzaicloud.io/vault-role: "admin"
+  # vault.security.banzaicloud.io/vault-addr: "https://vault.vault.svc.cluster.local:8200"
+  # vault.security.banzaicloud.io/vault-skip-verify: "true"
+  # vault.security.banzaicloud.io/vault-path: "kubernetes"
+  ## Example for inject secret
+  # slack_api_url: '${vault:secret/data/slack-hook-alerts#URL}'
+
 config:
   global: {}
     # slack_api_url: ''

```

## 0.27.0

**Release date:** 2023-03-23

![AppVersion: v0.25.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.25.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [alertmanager] Disable automountServiceAccountToken (#3147)

### Default value changes

```diff
diff --git a/charts/alertmanager/values.yaml b/charts/alertmanager/values.yaml
index de03023a..a167f30a 100644
--- a/charts/alertmanager/values.yaml
+++ b/charts/alertmanager/values.yaml
@@ -25,6 +25,8 @@ imagePullSecrets: []
 nameOverride: ""
 fullnameOverride: ""
 
+automountServiceAccountToken: true
+
 serviceAccount:
   # Specifies whether a service account should be created
   create: true

```

## 0.26.1

**Release date:** 2023-03-20

![AppVersion: v0.25.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.25.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* exclude `.Values.Config` from the conditions to enabled reloader (#3136)

### Default value changes

```diff
# No changes in this release
```

## 0.26.0

**Release date:** 2023-02-14

![AppVersion: v0.25.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.25.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [alertmanager] Add support for setting priority class name (#3025)

### Default value changes

```diff
diff --git a/charts/alertmanager/values.yaml b/charts/alertmanager/values.yaml
index 688fbc2f..de03023a 100644
--- a/charts/alertmanager/values.yaml
+++ b/charts/alertmanager/values.yaml
@@ -34,6 +34,9 @@ serviceAccount:
   # If not set and create is true, a name is generated using the fullname template
   name:
 
+# Sets priorityClassName in alertmanager pod
+priorityClassName: ""
+
 podSecurityContext:
   fsGroup: 65534
 dnsConfig: {}

```

## 0.25.0

**Release date:** 2023-01-13

![AppVersion: v0.25.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.25.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [alertmanager] Allow to configure init containers (#2917)

### Default value changes

```diff
diff --git a/charts/alertmanager/values.yaml b/charts/alertmanager/values.yaml
index c7809487..688fbc2f 100644
--- a/charts/alertmanager/values.yaml
+++ b/charts/alertmanager/values.yaml
@@ -57,6 +57,10 @@ securityContext:
 
 additionalPeers: []
 
+## Additional InitContainers to initialize the pod
+##
+extraInitContainers: []
+
 livenessProbe:
   httpGet:
     path: /

```

## 0.24.1

**Release date:** 2022-12-24

![AppVersion: v0.25.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.25.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* cleanup readme (#2857)

### Default value changes

```diff
# No changes in this release
```

## 0.24.0

**Release date:** 2022-12-23

![AppVersion: v0.25.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.25.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [alertmanager] updates alertmanager image to v0.25.0 (#2846)

### Default value changes

```diff
# No changes in this release
```

## 0.23.0

**Release date:** 2022-12-16

![AppVersion: v0.24.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.24.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [alertmanager] feat: support for custom volumes and environment variables (#2797)

### Default value changes

```diff
diff --git a/charts/alertmanager/values.yaml b/charts/alertmanager/values.yaml
index 9592b07a..c7809487 100644
--- a/charts/alertmanager/values.yaml
+++ b/charts/alertmanager/values.yaml
@@ -206,3 +206,19 @@ configmapReload:
 
 templates: {}
 #   alertmanager.tmpl: |-
+
+## Optionally specify extra list of additional volumeMounts
+extraVolumeMounts: []
+  # - name: extras
+  #   mountPath: /usr/share/extras
+  #   readOnly: true
+
+## Optionally specify extra list of additional volumes
+extraVolumes: []
+  # - name: extras
+  #   emptyDir: {}
+
+## Optionally specify extra environment variables to add to alertmanager container
+extraEnv: []
+  # - name: FOO
+  #   value: BAR

```

## 0.22.2

**Release date:** 2022-12-07

![AppVersion: v0.24.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.24.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [alertmanager] fix broken pdb due to missing line break (#2796)

### Default value changes

```diff
# No changes in this release
```

## 0.22.1

**Release date:** 2022-12-05

![AppVersion: v0.24.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.24.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [alertmanager] Refactor (#2781)

### Default value changes

```diff
# No changes in this release
```

## 0.22.0

**Release date:** 2022-11-04

![AppVersion: v0.24.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.24.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [alertmanager] update to current release (#2644)

### Default value changes

```diff
diff --git a/charts/alertmanager/values.yaml b/charts/alertmanager/values.yaml
index 0dadfdf1..9592b07a 100644
--- a/charts/alertmanager/values.yaml
+++ b/charts/alertmanager/values.yaml
@@ -194,7 +194,7 @@ configmapReload:
   ##
   image:
     repository: jimmidyson/configmap-reload
-    tag: v0.5.0
+    tag: v0.8.0
     pullPolicy: IfNotPresent
 
   # containerPort: 9533

```

## 0.21.0

**Release date:** 2022-09-29

![AppVersion: v0.23.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.23.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [alertmanager] add loadBalancerIP and loadBalancerSourceRanges (#2502)

### Default value changes

```diff
diff --git a/charts/alertmanager/values.yaml b/charts/alertmanager/values.yaml
index f20dea18..0dadfdf1 100644
--- a/charts/alertmanager/values.yaml
+++ b/charts/alertmanager/values.yaml
@@ -72,6 +72,8 @@ service:
   type: ClusterIP
   port: 9093
   clusterPort: 9094
+  loadBalancerIP: ""  # Assign ext IP when Service type is LoadBalancer
+  loadBalancerSourceRanges: []  # Only allow access to loadBalancerIP from these IPs
   # if you want to force a specific nodePort. Must be use with service.type=NodePort
   # nodePort:
 

```

## 0.20.1

**Release date:** 2022-09-15

![AppVersion: v0.23.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.23.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Fix all values being dumped for affinity rules (#2459)

### Default value changes

```diff
# No changes in this release
```

## 0.20.0

**Release date:** 2022-09-12

![AppVersion: v0.23.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.23.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [alertmanager] Added podAntiAffinity functionality for sts (#2428)

### Default value changes

```diff
diff --git a/charts/alertmanager/values.yaml b/charts/alertmanager/values.yaml
index d9ce2249..f20dea18 100644
--- a/charts/alertmanager/values.yaml
+++ b/charts/alertmanager/values.yaml
@@ -109,6 +109,18 @@ tolerations: []
 
 affinity: {}
 
+## Pod anti-affinity can prevent the scheduler from placing Alertmanager replicas on the same node.
+## The default value "soft" means that the scheduler should *prefer* to not schedule two replica pods onto the same node but no guarantee is provided.
+## The value "hard" means that the scheduler is *required* to not schedule two replica pods onto the same node.
+## The value "" will disable pod anti-affinity so that no anti-affinity rules will be configured.
+##
+podAntiAffinity: ""
+
+## If anti-affinity is enabled sets the topologyKey to use for anti-affinity.
+## This can be changed to, for example, failure-domain.beta.kubernetes.io/zone
+##
+podAntiAffinityTopologyKey: kubernetes.io/hostname
+
 ## Topology spread constraints rely on node labels to identify the topology domain(s) that each Node is in.
 ## Ref: https://kubernetes.io/docs/concepts/workloads/pods/pod-topology-spread-constraints/
 topologySpreadConstraints: []

```

## 0.19.0

**Release date:** 2022-07-21

![AppVersion: v0.23.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.23.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [alertmanager] Add optional topologySpreadConstraint to statefulset (#2292)

### Default value changes

```diff
diff --git a/charts/alertmanager/values.yaml b/charts/alertmanager/values.yaml
index 9ee76e15..d9ce2249 100644
--- a/charts/alertmanager/values.yaml
+++ b/charts/alertmanager/values.yaml
@@ -109,6 +109,16 @@ tolerations: []
 
 affinity: {}
 
+## Topology spread constraints rely on node labels to identify the topology domain(s) that each Node is in.
+## Ref: https://kubernetes.io/docs/concepts/workloads/pods/pod-topology-spread-constraints/
+topologySpreadConstraints: []
+  # - maxSkew: 1
+  #   topologyKey: failure-domain.beta.kubernetes.io/zone
+  #   whenUnsatisfiable: DoNotSchedule
+  #   labelSelector:
+  #     matchLabels:
+  #       app.kubernetes.io/instance: alertmanager
+
 statefulSet:
   annotations: {}
 

```

## 0.18.3

**Release date:** 2022-07-05

![AppVersion: v0.23.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.23.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* fix(alertmanager): 🚑 resolve replicaCount comparision type issue (#2231)

### Default value changes

```diff
# No changes in this release
```

## 0.18.2

**Release date:** 2022-06-17

![AppVersion: v0.23.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.23.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Upgrade pdb version to apiVersion: policy/v1 (#2165)

### Default value changes

```diff
# No changes in this release
```

## 0.18.1

**Release date:** 2022-05-26

![AppVersion: v0.23.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.23.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [AlertManager] Allow IPv6 support for main AlertManager StatefulSet (#2094)

### Default value changes

```diff
# No changes in this release
```

## 0.18.0

**Release date:** 2022-05-11

![AppVersion: v0.23.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.23.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [alertmanager] Add support change cluster port (#2043)

### Default value changes

```diff
diff --git a/charts/alertmanager/values.yaml b/charts/alertmanager/values.yaml
index 22eaa31d..9ee76e15 100644
--- a/charts/alertmanager/values.yaml
+++ b/charts/alertmanager/values.yaml
@@ -71,6 +71,7 @@ service:
   annotations: {}
   type: ClusterIP
   port: 9093
+  clusterPort: 9094
   # if you want to force a specific nodePort. Must be use with service.type=NodePort
   # nodePort:
 

```

## 0.17.1

**Release date:** 2022-05-04

![AppVersion: v0.23.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.23.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* trigger alertmanager build (#2028)

### Default value changes

```diff
# No changes in this release
```

## 0.17.0

**Release date:** 2022-03-28

![AppVersion: v0.23.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.23.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [alertmanager] Allow exposing configmap reloaders port (#1915)

### Default value changes

```diff
diff --git a/charts/alertmanager/values.yaml b/charts/alertmanager/values.yaml
index 16964266..22eaa31d 100644
--- a/charts/alertmanager/values.yaml
+++ b/charts/alertmanager/values.yaml
@@ -172,6 +172,8 @@ configmapReload:
     tag: v0.5.0
     pullPolicy: IfNotPresent
 
+  # containerPort: 9533
+
   ## configmap-reload resource requests and limits
   ## Ref: http://kubernetes.io/docs/user-guide/compute-resources/
   ##

```

## 0.16.0

**Release date:** 2022-02-09

![AppVersion: v0.23.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.23.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [alertmanager] Add alertmanager configurable livenessProbe/readinessProbe (#1783)

### Default value changes

```diff
diff --git a/charts/alertmanager/values.yaml b/charts/alertmanager/values.yaml
index 465ad01a..16964266 100644
--- a/charts/alertmanager/values.yaml
+++ b/charts/alertmanager/values.yaml
@@ -57,6 +57,16 @@ securityContext:
 
 additionalPeers: []
 
+livenessProbe:
+  httpGet:
+    path: /
+    port: http
+
+readinessProbe:
+  httpGet:
+    path: /
+    port: http
+
 service:
   annotations: {}
   type: ClusterIP

```

## 0.15.0

**Release date:** 2022-02-07

![AppVersion: v0.23.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.23.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [alertmanager] Add alertmanager extraSecretMounts (#1772)

### Default value changes

```diff
diff --git a/charts/alertmanager/values.yaml b/charts/alertmanager/values.yaml
index 6a912c71..465ad01a 100644
--- a/charts/alertmanager/values.yaml
+++ b/charts/alertmanager/values.yaml
@@ -12,6 +12,15 @@ image:
 
 extraArgs: {}
 
+## Additional Alertmanager Secret mounts
+# Defines additional mounts with secrets. Secrets must be manually created in the namespace.
+extraSecretMounts: []
+  # - name: secret-files
+  #   mountPath: /etc/secrets
+  #   subPath: ""
+  #   secretName: alertmanager-secret-files
+  #   readOnly: true
+
 imagePullSecrets: []
 nameOverride: ""
 fullnameOverride: ""

```

## 0.14.0

**Release date:** 2021-10-16

![AppVersion: v0.23.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.23.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [alertmanager] add podLabels value (#1429)

### Default value changes

```diff
diff --git a/charts/alertmanager/values.yaml b/charts/alertmanager/values.yaml
index f27cf76f..6a912c71 100644
--- a/charts/alertmanager/values.yaml
+++ b/charts/alertmanager/values.yaml
@@ -93,6 +93,7 @@ statefulSet:
   annotations: {}
 
 podAnnotations: {}
+podLabels: {}
 
 # Ref: https://kubernetes.io/docs/tasks/run-application/configure-pdb/
 podDisruptionBudget: {}

```

## 0.13.0

**Release date:** 2021-10-03

![AppVersion: v0.23.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.23.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [alertmanager] Add support for ingressClassName (#1371)

### Default value changes

```diff
diff --git a/charts/alertmanager/values.yaml b/charts/alertmanager/values.yaml
index 642fd574..f27cf76f 100644
--- a/charts/alertmanager/values.yaml
+++ b/charts/alertmanager/values.yaml
@@ -57,10 +57,15 @@ service:
 
 ingress:
   enabled: false
+  className: ""
   annotations: {}
+    # kubernetes.io/ingress.class: nginx
+    # kubernetes.io/tls-acme: "true"
   hosts:
     - host: alertmanager.domain.com
-      paths: []
+      paths:
+        - path: /
+          pathType: ImplementationSpecific
   tls: []
   #  - secretName: chart-example-tls
   #    hosts:

```

## 0.12.2

**Release date:** 2021-07-20

![AppVersion: v0.22.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.22.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [alertmanager] Unit test sample for ingress version (#1149)

### Default value changes

```diff
# No changes in this release
```

## 0.12.1

**Release date:** 2021-06-25

![AppVersion: v0.22.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.22.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [alertmanager] forse restart if configmapReload disable (#1094)

### Default value changes

```diff
# No changes in this release
```

## 0.12.0

**Release date:** 2021-06-16

![AppVersion: v0.22.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.22.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Add PDF to Alertmanager chart (#1080)

### Default value changes

```diff
diff --git a/charts/alertmanager/values.yaml b/charts/alertmanager/values.yaml
index 3b385b88..642fd574 100644
--- a/charts/alertmanager/values.yaml
+++ b/charts/alertmanager/values.yaml
@@ -89,6 +89,11 @@ statefulSet:
 
 podAnnotations: {}
 
+# Ref: https://kubernetes.io/docs/tasks/run-application/configure-pdb/
+podDisruptionBudget: {}
+  # maxUnavailable: 1
+  # minAvailable: 1
+
 command: []
 
 persistence:

```

## 0.11.0

**Release date:** 2021-06-02

![AppVersion: v0.22.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.22.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [alertmanager] Upgrade version to 0.22.1 (#1031)

### Default value changes

```diff
# No changes in this release
```

## 0.10.2

**Release date:** 2021-05-18

![AppVersion: v0.21.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.21.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Move tolerations and nodeSelector to not disrupt volumes (#970)

### Default value changes

```diff
# No changes in this release
```

## 0.10.1

**Release date:** 2021-05-12

![AppVersion: v0.21.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.21.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* AlertManager: Fixing indentation for node selector keys in case persistence is enabled (#963)

### Default value changes

```diff
# No changes in this release
```

## 0.10.0

**Release date:** 2021-05-05

![AppVersion: v0.21.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.21.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [alertmanager] Use quay.io image instead of docker hub (#928)

### Default value changes

```diff
diff --git a/charts/alertmanager/values.yaml b/charts/alertmanager/values.yaml
index bcc193b5..3b385b88 100644
--- a/charts/alertmanager/values.yaml
+++ b/charts/alertmanager/values.yaml
@@ -5,7 +5,7 @@
 replicaCount: 1
 
 image:
-  repository: prom/alertmanager
+  repository: quay.io/prometheus/alertmanager
   pullPolicy: IfNotPresent
   # Overrides the image tag whose default is the chart appVersion.
   tag: ""

```

## 0.9.3

**Release date:** 2021-04-23

![AppVersion: v0.21.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.21.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* fix, cluster-peer needs to be name+service (#886)

### Default value changes

```diff
# No changes in this release
```

## 0.9.2

**Release date:** 2021-04-23

![AppVersion: v0.21.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.21.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [alertmanager]fix cluster-peer name follow statefullSet name  (#881)

### Default value changes

```diff
# No changes in this release
```

## 0.9.1

**Release date:** 2021-04-22

![AppVersion: v0.21.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.21.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* fix cluster-peer names within cluster (#878)

### Default value changes

```diff
# No changes in this release
```

## 0.9.0

**Release date:** 2021-03-29

![AppVersion: v0.21.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.21.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [alertmanager] add NodePort support (#790)

### Default value changes

```diff
diff --git a/charts/alertmanager/values.yaml b/charts/alertmanager/values.yaml
index 12f277f5..bcc193b5 100644
--- a/charts/alertmanager/values.yaml
+++ b/charts/alertmanager/values.yaml
@@ -52,6 +52,8 @@ service:
   annotations: {}
   type: ClusterIP
   port: 9093
+  # if you want to force a specific nodePort. Must be use with service.type=NodePort
+  # nodePort:
 
 ingress:
   enabled: false

```

## 0.8.0

**Release date:** 2021-03-19

![AppVersion: v0.21.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.21.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [Alertmanager] command podannotations (#781)

### Default value changes

```diff
diff --git a/charts/alertmanager/values.yaml b/charts/alertmanager/values.yaml
index aa695a47..12f277f5 100644
--- a/charts/alertmanager/values.yaml
+++ b/charts/alertmanager/values.yaml
@@ -85,6 +85,10 @@ affinity: {}
 statefulSet:
   annotations: {}
 
+podAnnotations: {}
+
+command: []
+
 persistence:
   enabled: true
   ## Persistent Volume Storage Class

```

## 0.7.1

**Release date:** 2021-03-19

![AppVersion: v0.21.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.21.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [Alertmanager]: Fix indent for service tpl annotation (#777)

### Default value changes

```diff
# No changes in this release
```

## 0.7.0

**Release date:** 2021-03-18

![AppVersion: v0.21.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.21.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [Alertmanager]: Enable annotations for service tpl (#767)

### Default value changes

```diff
diff --git a/charts/alertmanager/values.yaml b/charts/alertmanager/values.yaml
index 0e07e8f4..aa695a47 100644
--- a/charts/alertmanager/values.yaml
+++ b/charts/alertmanager/values.yaml
@@ -49,6 +49,7 @@ securityContext:
 additionalPeers: []
 
 service:
+  annotations: {}
   type: ClusterIP
   port: 9093
 

```

## 0.6.0

**Release date:** 2021-02-09

![AppVersion: v0.21.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.21.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Ability for custom dnsConfig in alertmanager (#642)

### Default value changes

```diff
diff --git a/charts/alertmanager/values.yaml b/charts/alertmanager/values.yaml
index b8acca9a..0e07e8f4 100644
--- a/charts/alertmanager/values.yaml
+++ b/charts/alertmanager/values.yaml
@@ -27,7 +27,16 @@ serviceAccount:
 
 podSecurityContext:
   fsGroup: 65534
-
+dnsConfig: {}
+  # nameservers:
+  #   - 1.2.3.4
+  # searches:
+  #   - ns1.svc.cluster-domain.example
+  #   - my.dns.search.suffix
+  # options:
+  #   - name: ndots
+  #     value: "2"
+  #   - name: edns0
 securityContext:
   # capabilities:
   #   drop:

```

## 0.5.1

**Release date:** 2021-02-07

![AppVersion: v0.21.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.21.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* chore: bump configmap reloader (#645)

### Default value changes

```diff
diff --git a/charts/alertmanager/values.yaml b/charts/alertmanager/values.yaml
index 782a531c..b8acca9a 100644
--- a/charts/alertmanager/values.yaml
+++ b/charts/alertmanager/values.yaml
@@ -123,7 +123,7 @@ configmapReload:
   ##
   image:
     repository: jimmidyson/configmap-reload
-    tag: v0.4.0
+    tag: v0.5.0
     pullPolicy: IfNotPresent
 
   ## configmap-reload resource requests and limits

```

## 0.5.0

**Release date:** 2021-01-23

![AppVersion: v0.21.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.21.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [alertmanager] allow possibility to use alertmanager in HA (#609)

### Default value changes

```diff
diff --git a/charts/alertmanager/values.yaml b/charts/alertmanager/values.yaml
index eb6188d9..782a531c 100644
--- a/charts/alertmanager/values.yaml
+++ b/charts/alertmanager/values.yaml
@@ -37,6 +37,8 @@ securityContext:
   runAsNonRoot: true
   runAsGroup: 65534
 
+additionalPeers: []
+
 service:
   type: ClusterIP
   port: 9093

```

## 0.4.0

**Release date:** 2021-01-23

![AppVersion: v0.21.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.21.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [alertmanager] Add optional config reloader (#348)

### Default value changes

```diff
diff --git a/charts/alertmanager/values.yaml b/charts/alertmanager/values.yaml
index de6954ad..eb6188d9 100644
--- a/charts/alertmanager/values.yaml
+++ b/charts/alertmanager/values.yaml
@@ -105,5 +105,29 @@ config:
     receiver: default-receiver
     repeat_interval: 3h
 
+## Monitors ConfigMap changes and POSTs to a URL
+## Ref: https://github.com/jimmidyson/configmap-reload
+##
+configmapReload:
+  ## If false, the configmap-reload container will not be deployed
+  ##
+  enabled: false
+
+  ## configmap-reload container name
+  ##
+  name: configmap-reload
+
+  ## configmap-reload container image
+  ##
+  image:
+    repository: jimmidyson/configmap-reload
+    tag: v0.4.0
+    pullPolicy: IfNotPresent
+
+  ## configmap-reload resource requests and limits
+  ## Ref: http://kubernetes.io/docs/user-guide/compute-resources/
+  ##
+  resources: {}
+
 templates: {}
 #   alertmanager.tmpl: |-

```

## 0.3.0

**Release date:** 2020-12-12

![AppVersion: v0.21.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.21.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [alertmanager] Allow passing notification templates (#484)

### Default value changes

```diff
diff --git a/charts/alertmanager/values.yaml b/charts/alertmanager/values.yaml
index acbd7380..de6954ad 100644
--- a/charts/alertmanager/values.yaml
+++ b/charts/alertmanager/values.yaml
@@ -90,6 +90,9 @@ config:
   global: {}
     # slack_api_url: ''
 
+  templates:
+    - '/etc/alertmanager/*.tmpl'
+
   receivers:
     - name: default-receiver
       # slack_configs:
@@ -101,3 +104,6 @@ config:
     group_interval: 5m
     receiver: default-receiver
     repeat_interval: 3h
+
+templates: {}
+#   alertmanager.tmpl: |-

```

## 0.2.0

**Release date:** 2020-11-25

![AppVersion: v0.21.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.21.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [alertmanager] add statefulSet annotations (#414)

### Default value changes

```diff
diff --git a/charts/alertmanager/values.yaml b/charts/alertmanager/values.yaml
index 70e9a7de..acbd7380 100644
--- a/charts/alertmanager/values.yaml
+++ b/charts/alertmanager/values.yaml
@@ -70,6 +70,9 @@ tolerations: []
 
 affinity: {}
 
+statefulSet:
+  annotations: {}
+
 persistence:
   enabled: true
   ## Persistent Volume Storage Class

```

## 0.1.4

**Release date:** 2020-10-28

![AppVersion: v0.21.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.21.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* alertmanager: add maintainer (#276)

### Default value changes

```diff
# No changes in this release
```

## 0.1.3

**Release date:** 2020-10-21

![AppVersion: v0.21.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.21.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [alertmanager] Add support change default config (#240)

### Default value changes

```diff
diff --git a/charts/alertmanager/values.yaml b/charts/alertmanager/values.yaml
index 6f4a487d..70e9a7de 100644
--- a/charts/alertmanager/values.yaml
+++ b/charts/alertmanager/values.yaml
@@ -10,6 +10,8 @@ image:
   # Overrides the image tag whose default is the chart appVersion.
   tag: ""
 
+extraArgs: {}
+
 imagePullSecrets: []
 nameOverride: ""
 fullnameOverride: ""

```

## 0.1.2

**Release date:** 2020-10-20

![AppVersion: v0.21.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.21.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Fix alertmanager statefulset template (#236)

### Default value changes

```diff
# No changes in this release
```

## 0.1.1

**Release date:** 2020-10-02

![AppVersion: v0.21.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.21.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [alertmanager ]fix persistence (#174)

### Default value changes

```diff
diff --git a/charts/alertmanager/values.yaml b/charts/alertmanager/values.yaml
index b8dac665..6f4a487d 100644
--- a/charts/alertmanager/values.yaml
+++ b/charts/alertmanager/values.yaml
@@ -69,7 +69,7 @@ tolerations: []
 affinity: {}
 
 persistence:
-  enabled: false
+  enabled: true
   ## Persistent Volume Storage Class
   ## If defined, storageClassName: <storageClass>
   ## If set to "-", storageClassName: "", which disables dynamic provisioning
@@ -79,7 +79,7 @@ persistence:
   # storageClass: "-"
   accessModes:
     - ReadWriteOnce
-  size: 1Gi
+  size: 50Mi
 
 config:
   global: {}

```

## 0.1.0

**Release date:** 2020-09-27

![AppVersion: v0.21.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.21.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [alertmanager] add new chart (#64)

### Default value changes

```diff
# Default values for alertmanager.
# This is a YAML-formatted file.
# Declare variables to be passed into your templates.

replicaCount: 1

image:
  repository: prom/alertmanager
  pullPolicy: IfNotPresent
  # Overrides the image tag whose default is the chart appVersion.
  tag: ""

imagePullSecrets: []
nameOverride: ""
fullnameOverride: ""

serviceAccount:
  # Specifies whether a service account should be created
  create: true
  # Annotations to add to the service account
  annotations: {}
  # The name of the service account to use.
  # If not set and create is true, a name is generated using the fullname template
  name:

podSecurityContext:
  fsGroup: 65534

securityContext:
  # capabilities:
  #   drop:
  #   - ALL
  # readOnlyRootFilesystem: true
  runAsUser: 65534
  runAsNonRoot: true
  runAsGroup: 65534

service:
  type: ClusterIP
  port: 9093

ingress:
  enabled: false
  annotations: {}
  hosts:
    - host: alertmanager.domain.com
      paths: []
  tls: []
  #  - secretName: chart-example-tls
  #    hosts:
  #      - alertmanager.domain.com

resources: {}
  # We usually recommend not to specify default resources and to leave this as a conscious
  # choice for the user. This also increases chances charts run on environments with little
  # resources, such as Minikube. If you do want to specify resources, uncomment the following
  # lines, adjust them as necessary, and remove the curly braces after 'resources:'.
  # limits:
  #   cpu: 100m
  #   memory: 128Mi
  # requests:
  #   cpu: 10m
  #   memory: 32Mi

nodeSelector: {}

tolerations: []

affinity: {}

persistence:
  enabled: false
  ## Persistent Volume Storage Class
  ## If defined, storageClassName: <storageClass>
  ## If set to "-", storageClassName: "", which disables dynamic provisioning
  ## If undefined (the default) or set to null, no storageClassName spec is
  ## set, choosing the default provisioner.
  ##
  # storageClass: "-"
  accessModes:
    - ReadWriteOnce
  size: 1Gi

config:
  global: {}
    # slack_api_url: ''

  receivers:
    - name: default-receiver
      # slack_configs:
      #  - channel: '@you'
      #    send_resolved: true

  route:
    group_wait: 10s
    group_interval: 5m
    receiver: default-receiver
    repeat_interval: 3h

```

---
Autogenerated from Helm Chart and git history using [helm-changelog](https://github.com/mogensen/helm-changelog)
