# Change Log

## 5.8.0

**Release date:** 2023-06-13

![AppVersion: 2.9.2](https://img.shields.io/static/v1?label=AppVersion&message=2.9.2&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Support CustomResourceState (#3470)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index d34b6317..c7d510c4 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -334,6 +334,12 @@ kubeconfig:
   # base64 encoded kube-config file
   secret:
 
+# Enabling support for customResourceState, will create a configMap including your config that will be read from kube-state-metrics
+customResourceState:
+  enabled: false
+  # Add (Cluster)Role permissions to list/watch the customResources defined in the config to rbac.extraRules
+  config: {}
+
 # Enable only the release namespace for collecting resources. By default all namespaces are collected.
 # If releaseNamespace and namespaces are both set a merged list will be collected.
 releaseNamespace: false

```

## 5.7.0

**Release date:** 2023-05-31

![AppVersion: 2.9.2](https://img.shields.io/static/v1?label=AppVersion&message=2.9.2&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Version bump to 2.9.2 (#3450)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index c02b8970..d34b6317 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -327,7 +327,6 @@ collectors:
   - storageclasses
   - validatingwebhookconfigurations
   - volumeattachments
-  # - verticalpodautoscalers # not a default resource, see also: https://github.com/kubernetes/kube-state-metrics#enabling-verticalpodautoscalers
 
 # Enabling kubeconfig will pass the --kubeconfig argument to the container
 kubeconfig:

```

## 5.6.4

**Release date:** 2023-05-24

![AppVersion: 2.8.2](https://img.shields.io/static/v1?label=AppVersion&message=2.8.2&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] add missing bearerToken for ServiceMonitor self endpoint (#3423)

### Default value changes

```diff
# No changes in this release
```

## 5.6.3

**Release date:** 2023-05-23

![AppVersion: 2.8.2](https://img.shields.io/static/v1?label=AppVersion&message=2.8.2&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Fix VPA when autosharding is enabled (#3420)

### Default value changes

```diff
# No changes in this release
```

## 5.6.2

**Release date:** 2023-05-11

![AppVersion: 2.8.2](https://img.shields.io/static/v1?label=AppVersion&message=2.8.2&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Fix body must be of type string bug for cilium network policy (#3363)

### Default value changes

```diff
# No changes in this release
```

## 5.6.1

**Release date:** 2023-05-03

![AppVersion: 2.8.2](https://img.shields.io/static/v1?label=AppVersion&message=2.8.2&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Add ingress in the CiliumNetworkPolicy (#3299)

### Default value changes

```diff
# No changes in this release
```

## 5.6.0

**Release date:** 2023-04-24

![AppVersion: 2.8.2](https://img.shields.io/static/v1?label=AppVersion&message=2.8.2&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Add cilium networkpolicy support (#3281)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 3a293d87..c02b8970 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -202,6 +202,18 @@ podSecurityPolicy:
 ## Configure network policy for kube-state-metrics
 networkPolicy:
   enabled: false
+  # networkPolicy.flavor -- Flavor of the network policy to use.
+  # Can be:
+  # * kubernetes for networking.k8s.io/v1/NetworkPolicy
+  # * cilium     for cilium.io/v2/CiliumNetworkPolicy
+  flavor: kubernetes
+
+  ## Configure the cilium network policy kube-apiserver selector
+  # cilium:
+    # kubeApiServerSelector:
+      # - toEntities:
+      #   - kube-apiserver
+
   # egress:
   # - {}
   # ingress:

```

## 5.5.0

**Release date:** 2023-04-15

![AppVersion: 2.8.2](https://img.shields.io/static/v1?label=AppVersion&message=2.8.2&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] set parameters for podsecurity restricted (#3235)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 9b932698..3a293d87 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -215,11 +215,18 @@ securityContext:
   runAsGroup: 65534
   runAsUser: 65534
   fsGroup: 65534
+  runAsNonRoot: true
+  seccompProfile:
+    type: RuntimeDefault
 
 ## Specify security settings for a Container
 ## Allows overrides and additional options compared to (Pod) securityContext
 ## Ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/#set-the-security-context-for-a-container
-containerSecurityContext: {}
+containerSecurityContext:
+  allowPrivilegeEscalation: false
+  capabilities:
+    drop:
+    - ALL
 
 ## Node labels for pod assignment
 ## Ref: https://kubernetes.io/docs/user-guide/node-selection/

```

## 5.4.2

**Release date:** 2023-04-14

![AppVersion: 2.8.2](https://img.shields.io/static/v1?label=AppVersion&message=2.8.2&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Fix ServiceMonitor annotations templating (#3238)

### Default value changes

```diff
# No changes in this release
```

## 5.4.1

**Release date:** 2023-04-13

![AppVersion: 2.8.2](https://img.shields.io/static/v1?label=AppVersion&message=2.8.2&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* fix(kube-state-metrics): revert #3194 (#3233)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index e168fd5f..9b932698 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -214,14 +214,7 @@ securityContext:
   enabled: true
   runAsGroup: 65534
   runAsUser: 65534
-  runAsNonRoot: true
   fsGroup: 65534
-  allowPrivilegeEscalation: false
-  capabilities:
-    drop:
-    - ALL
-  seccompProfile:
-    type: RuntimeDefault
 
 ## Specify security settings for a Container
 ## Allows overrides and additional options compared to (Pod) securityContext

```

## 5.4.0

**Release date:** 2023-04-11

![AppVersion: 2.8.2](https://img.shields.io/static/v1?label=AppVersion&message=2.8.2&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] set parameters for podsecurity restricted (#3194)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 9b932698..e168fd5f 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -214,7 +214,14 @@ securityContext:
   enabled: true
   runAsGroup: 65534
   runAsUser: 65534
+  runAsNonRoot: true
   fsGroup: 65534
+  allowPrivilegeEscalation: false
+  capabilities:
+    drop:
+    - ALL
+  seccompProfile:
+    type: RuntimeDefault
 
 ## Specify security settings for a Container
 ## Allows overrides and additional options compared to (Pod) securityContext

```

## 5.3.0

**Release date:** 2023-03-31

![AppVersion: 2.8.2](https://img.shields.io/static/v1?label=AppVersion&message=2.8.2&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Add annotations to servicemonitor (#3178)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 48490dd9..9b932698 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -139,6 +139,7 @@ serviceAccount:
 prometheus:
   monitor:
     enabled: false
+    annotations: {}
     additionalLabels: {}
     namespace: ""
     jobLabel: ""

```

## 5.2.0

**Release date:** 2023-03-28

![AppVersion: 2.8.2](https://img.shields.io/static/v1?label=AppVersion&message=2.8.2&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Support volume mounts in sidecar containers (#3159)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index bc2d4a1e..48490dd9 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -115,6 +115,13 @@ kubeRBACProxy:
     #  cpu: 10m
     #  memory: 32Mi
 
+  ## volumeMounts enables mounting custom volumes in rbac-proxy containers
+  ## Useful for TLS certificates and keys
+  volumeMounts: []
+    # - mountPath: /etc/tls
+    #   name: kube-rbac-proxy-tls
+    #   readOnly: true
+
 serviceAccount:
   # Specifies whether a ServiceAccount should be created, require rbac true
   create: true

```

## 5.1.0

**Release date:** 2023-03-27

![AppVersion: 2.8.2](https://img.shields.io/static/v1?label=AppVersion&message=2.8.2&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Support bearer token (#3153)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 14f66f62..bc2d4a1e 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -164,6 +164,14 @@ prometheus:
     metricRelabelings: []
     relabelings: []
     scheme: ""
+    ## File to read bearer token for scraping targets
+    bearerTokenFile: ""
+    ## Secret to mount to read bearer token for scraping targets. The secret needs
+    ## to be in the same namespace as the service monitor and accessible by the
+    ## Prometheus Operator
+    bearerTokenSecret: {}
+      # name: secret-name
+      # key:  key-name
     tlsConfig: {}
 
 ## Specify if a Pod Security Policy for kube-state-metrics must be created

```

## 5.0.1

**Release date:** 2023-03-20

![AppVersion: 2.8.2](https://img.shields.io/static/v1?label=AppVersion&message=2.8.2&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Bump to v2.8.2 (#3137)

### Default value changes

```diff
# No changes in this release
```

## 5.0.0

**Release date:** 2023-03-07

![AppVersion: 2.8.1](https://img.shields.io/static/v1?label=AppVersion&message=2.8.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] fix/split Registry and Repository Values in a proper way (#3076)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 13111541..14f66f62 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -1,7 +1,8 @@
 # Default values for kube-state-metrics.
 prometheusScrape: true
 image:
-  repository: registry.k8s.io/kube-state-metrics/kube-state-metrics
+  registry: registry.k8s.io
+  repository: kube-state-metrics/kube-state-metrics
   # If unset use v + .Charts.appVersion
   tag: ""
   sha: ""
@@ -86,7 +87,8 @@ rbac:
 kubeRBACProxy:
   enabled: false
   image:
-    repository: quay.io/brancz/kube-rbac-proxy
+    registry: quay.io
+    repository: brancz/kube-rbac-proxy
     tag: v0.14.0
     sha: ""
     pullPolicy: IfNotPresent

```

## 4.32.0

**Release date:** 2023-03-06

![AppVersion: 2.8.1](https://img.shields.io/static/v1?label=AppVersion&message=2.8.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Add network policy and bump to 2.8.1 (#3089)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index d5173e82..13111541 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -181,6 +181,17 @@ podSecurityPolicy:
 
   additionalVolumes: []
 
+## Configure network policy for kube-state-metrics
+networkPolicy:
+  enabled: false
+  # egress:
+  # - {}
+  # ingress:
+  # - {}
+  # podSelector:
+  #   matchLabels:
+  #     app.kubernetes.io/name: kube-state-metrics
+
 securityContext:
   enabled: true
   runAsGroup: 65534

```

## 4.31.0

**Release date:** 2023-02-20

![AppVersion: 2.8.0](https://img.shields.io/static/v1?label=AppVersion&message=2.8.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Allow parent charts to override registry hostname (#3043)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 0f185fbe..d5173e82 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -23,6 +23,9 @@ global:
   #   - pullSecret1
   #   - pullSecret2
   imagePullSecrets: []
+  #
+  # Allow parent charts to override registry hostname
+  imageRegistry: ""
 
 # If set to true, this will deploy kube-state-metrics as a StatefulSet and the data
 # will be automatically sharded across <.Values.replicas> pods using the built-in

```

## 4.30.0

**Release date:** 2023-02-14

![AppVersion: 2.8.0](https://img.shields.io/static/v1?label=AppVersion&message=2.8.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Version bump to 2.8.0 (#3033)

### Default value changes

```diff
# No changes in this release
```

## 4.29.0

**Release date:** 2023-01-19

![AppVersion: 2.7.0](https://img.shields.io/static/v1?label=AppVersion&message=2.7.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Use appVersion to define container image version (#2936)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 2691d6ee..0f185fbe 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -2,7 +2,8 @@
 prometheusScrape: true
 image:
   repository: registry.k8s.io/kube-state-metrics/kube-state-metrics
-  tag: v2.7.0
+  # If unset use v + .Charts.appVersion
+  tag: ""
   sha: ""
   pullPolicy: IfNotPresent
 

```

## 4.28.0

**Release date:** 2023-01-10

![AppVersion: 2.7.0](https://img.shields.io/static/v1?label=AppVersion&message=2.7.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Allow to enable kube-rbac-proxy for kube-state-metric (#2888)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index fbcefd95..2691d6ee 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -77,6 +77,38 @@ rbac:
   #   verbs: ["list", "watch"]
   extraRules: []
 
+# Configure kube-rbac-proxy. When enabled, creates one kube-rbac-proxy container per exposed HTTP endpoint (metrics and telemetry if enabled).
+# The requests are served through the same service but requests are then HTTPS.
+kubeRBACProxy:
+  enabled: false
+  image:
+    repository: quay.io/brancz/kube-rbac-proxy
+    tag: v0.14.0
+    sha: ""
+    pullPolicy: IfNotPresent
+
+  # List of additional cli arguments to configure kube-rbac-prxy
+  # for example: --tls-cipher-suites, --log-file, etc.
+  # all the possible args can be found here: https://github.com/brancz/kube-rbac-proxy#usage
+  extraArgs: []
+
+  ## Specify security settings for a Container
+  ## Allows overrides and additional options compared to (Pod) securityContext
+  ## Ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/#set-the-security-context-for-a-container
+  containerSecurityContext: {}
+
+  resources: {}
+    # We usually recommend not to specify default resources and to leave this as a conscious
+    # choice for the user. This also increases chances charts run on environments with little
+    # resources, such as Minikube. If you do want to specify resources, uncomment the following
+    # lines, adjust them as necessary, and remove the curly braces after 'resources:'.
+    # limits:
+    #  cpu: 100m
+    #  memory: 64Mi
+    # requests:
+    #  cpu: 10m
+    #  memory: 32Mi
+
 serviceAccount:
   # Specifies whether a ServiceAccount should be created, require rbac true
   create: true

```

## 4.27.0

**Release date:** 2023-01-09

![AppVersion: 2.7.0](https://img.shields.io/static/v1?label=AppVersion&message=2.7.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] add support for overriding selector labels (#2889)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 72bc2e2e..fbcefd95 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -52,6 +52,9 @@ service:
 customLabels: {}
   # app: kube-state-metrics
 
+## Override selector labels
+selectorOverride: {}
+
 ## set to true to add the release label so scraping of the servicemonitor with kube-prometheus-stack works out of the box
 releaseLabel: false
 

```

## 4.26.0

**Release date:** 2023-01-06

![AppVersion: 2.7.0](https://img.shields.io/static/v1?label=AppVersion&message=2.7.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Add support for global imagePullSecrets values (#2769)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 36c1a1d9..72bc2e2e 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -9,6 +9,20 @@ image:
 imagePullSecrets: []
 # - name: "image-pull-secret"
 
+global:
+  # To help compatibility with other charts which use global.imagePullSecrets.
+  # Allow either an array of {name: pullSecret} maps (k8s-style), or an array of strings (more common helm-style).
+  # global:
+  #   imagePullSecrets:
+  #   - name: pullSecret1
+  #   - name: pullSecret2
+  # or
+  # global:
+  #   imagePullSecrets:
+  #   - pullSecret1
+  #   - pullSecret2
+  imagePullSecrets: []
+
 # If set to true, this will deploy kube-state-metrics as a StatefulSet and the data
 # will be automatically sharded across <.Values.replicas> pods using the built-in
 # autodiscovery feature: https://github.com/kubernetes/kube-state-metrics#automated-sharding

```

## 4.25.0

**Release date:** 2022-12-20

![AppVersion: 2.7.0](https://img.shields.io/static/v1?label=AppVersion&message=2.7.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Manage 'targetLabels' and 'podTargetLabels' (#2832)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index d32b1232..36c1a1d9 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -80,6 +80,8 @@ prometheus:
     additionalLabels: {}
     namespace: ""
     jobLabel: ""
+    targetLabels: []
+    podTargetLabels: []
     interval: ""
     ## SampleLimit defines per-scrape limit on number of scraped samples that will be accepted.
     ##

```

## 4.24.0

**Release date:** 2022-11-25

![AppVersion: 2.7.0](https://img.shields.io/static/v1?label=AppVersion&message=2.7.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Bump kube-state-metrics version to v2.7.0 #2736 (#2737)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 931f3dfa..d32b1232 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -2,7 +2,7 @@
 prometheusScrape: true
 image:
   repository: registry.k8s.io/kube-state-metrics/kube-state-metrics
-  tag: v2.6.0
+  tag: v2.7.0
   sha: ""
   pullPolicy: IfNotPresent
 

```

## 4.23.0

**Release date:** 2022-11-10

![AppVersion: 2.6.0](https://img.shields.io/static/v1?label=AppVersion&message=2.6.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] add service load balancer IP ranges (#2672)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 67e8bca6..931f3dfa 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -29,6 +29,8 @@ service:
   type: ClusterIP
   nodePort: 0
   loadBalancerIP: ""
+  # Only allow access to the loadBalancerIP from these IPs
+  loadBalancerSourceRanges: []
   clusterIP: ""
   annotations: {}
 

```

## 4.22.3

**Release date:** 2022-11-08

![AppVersion: 2.6.0](https://img.shields.io/static/v1?label=AppVersion&message=2.6.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* fix(kube-state-metrics): updated maintainer email (#2659)

### Default value changes

```diff
# No changes in this release
```

## 4.22.2

**Release date:** 2022-11-07

![AppVersion: 2.6.0](https://img.shields.io/static/v1?label=AppVersion&message=2.6.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Make headless service for auto sharding (#2652)

### Default value changes

```diff
# No changes in this release
```

## 4.22.1

**Release date:** 2022-11-02

![AppVersion: 2.6.0](https://img.shields.io/static/v1?label=AppVersion&message=2.6.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* fix(kube-state-metrics): removed empty line introduced in #2610 (#2628)

### Default value changes

```diff
# No changes in this release
```

## 4.22.0

**Release date:** 2022-11-01

![AppVersion: 2.6.0](https://img.shields.io/static/v1?label=AppVersion&message=2.6.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] add servicemonitor scrape limits (#2610)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index edecda7e..67e8bca6 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -79,6 +79,25 @@ prometheus:
     namespace: ""
     jobLabel: ""
     interval: ""
+    ## SampleLimit defines per-scrape limit on number of scraped samples that will be accepted.
+    ##
+    sampleLimit: 0
+
+    ## TargetLimit defines a limit on the number of scraped targets that will be accepted.
+    ##
+    targetLimit: 0
+
+    ## Per-scrape limit on number of labels that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.
+    ##
+    labelLimit: 0
+
+    ## Per-scrape limit on length of labels name that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.
+    ##
+    labelNameLengthLimit: 0
+
+    ## Per-scrape limit on length of labels value that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.
+    ##
+    labelValueLengthLimit: 0
     scrapeTimeout: ""
     proxyUrl: ""
     selectorOverride: {}

```

## 4.21.1

**Release date:** 2022-10-31

![AppVersion: 2.6.0](https://img.shields.io/static/v1?label=AppVersion&message=2.6.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] ServiceMonitor selectorOverride fix (#2621)

### Default value changes

```diff
# No changes in this release
```

## 4.21.0

**Release date:** 2022-10-18

![AppVersion: 2.6.0](https://img.shields.io/static/v1?label=AppVersion&message=2.6.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] add leases to collectors (#2584)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index fb35dd72..edecda7e 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -184,6 +184,7 @@ collectors:
   - horizontalpodautoscalers
   - ingresses
   - jobs
+  - leases
   - limitranges
   - mutatingwebhookconfigurations
   - namespaces

```

## 4.20.3

**Release date:** 2022-10-16

![AppVersion: 2.6.0](https://img.shields.io/static/v1?label=AppVersion&message=2.6.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Fix PSP deprecation after k8s 1.25+ (#2564)

### Default value changes

```diff
# No changes in this release
```

## 4.20.2

**Release date:** 2022-09-30

![AppVersion: 2.6.0](https://img.shields.io/static/v1?label=AppVersion&message=2.6.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Allow namespaces to be defined as both string/list (#2507)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index fca069f8..fb35dd72 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -214,7 +214,7 @@ kubeconfig:
 # If releaseNamespace and namespaces are both set a merged list will be collected.
 releaseNamespace: false
 
-# Comma-separated list of namespaces to be enabled for collecting resources. By default all namespaces are collected.
+# Comma-separated list(string) or yaml list of namespaces to be enabled for collecting resources. By default all namespaces are collected.
 namespaces: ""
 
 # Comma-separated list of namespaces not to be enabled. If namespaces and namespaces-denylist are both set,

```

## 4.20.1

**Release date:** 2022-09-27

![AppVersion: 2.6.0](https://img.shields.io/static/v1?label=AppVersion&message=2.6.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] - Ensure namespaces uniqueness (#2484)

### Default value changes

```diff
# No changes in this release
```

## 4.20.0

**Release date:** 2022-09-20

![AppVersion: 2.6.0](https://img.shields.io/static/v1?label=AppVersion&message=2.6.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Add vpa support (#2469)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index cfeb5196..fca069f8 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -251,6 +251,26 @@ selfMonitor:
   # telemetryPort: 8081
   # telemetryNodePort: 0
 
+# Enable vertical pod autoscaler support for kube-state-metrics
+verticalPodAutoscaler:
+  enabled: false
+  # List of resources that the vertical pod autoscaler can control. Defaults to cpu and memory
+  controlledResources: []
+
+  # Define the max allowed resources for the pod
+  maxAllowed: {}
+  # cpu: 200m
+  # memory: 100Mi
+  # Define the min allowed resources for the pod
+  minAllowed: {}
+  # cpu: 200m
+  # memory: 100Mi
+
+  # updatePolicy:
+    # Specifies whether recommended updates are applied when a Pod is started and whether recommended updates
+    # are applied during the life of a Pod. Possible values are "Off", "Initial", "Recreate", and "Auto".
+    # updateMode: Auto
+
 # volumeMounts are used to add custom volume mounts to deployment.
 # See example below
 volumeMounts: []

```

## 4.19.0

**Release date:** 2022-09-17

![AppVersion: 2.6.0](https://img.shields.io/static/v1?label=AppVersion&message=2.6.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Add option to merge namespaces (#2457)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 1fe0660a..cfeb5196 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -211,7 +211,7 @@ kubeconfig:
   secret:
 
 # Enable only the release namespace for collecting resources. By default all namespaces are collected.
-# If releaseNamespace and namespaces are both set only releaseNamespace will be used.
+# If releaseNamespace and namespaces are both set a merged list will be collected.
 releaseNamespace: false
 
 # Comma-separated list of namespaces to be enabled for collecting resources. By default all namespaces are collected.

```

## 4.18.0

**Release date:** 2022-09-04

![AppVersion: 2.6.0](https://img.shields.io/static/v1?label=AppVersion&message=2.6.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Bump kube-state-metrics version to v2.6.0 (#2422)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 4582d80a..1fe0660a 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -2,7 +2,7 @@
 prometheusScrape: true
 image:
   repository: registry.k8s.io/kube-state-metrics/kube-state-metrics
-  tag: v2.5.0
+  tag: v2.6.0
   sha: ""
   pullPolicy: IfNotPresent
 

```

## 4.17.0

**Release date:** 2022-09-01

![AppVersion: 2.5.0](https://img.shields.io/static/v1?label=AppVersion&message=2.5.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Fix selfMonitor port (#2418)

### Default value changes

```diff
# No changes in this release
```

## 4.16.0

**Release date:** 2022-08-16

![AppVersion: 2.5.0](https://img.shields.io/static/v1?label=AppVersion&message=2.5.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* feat[kube-state-metrics]: add image sha support for image specified (#2370)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index cf74d67e..4582d80a 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -3,6 +3,7 @@ prometheusScrape: true
 image:
   repository: registry.k8s.io/kube-state-metrics/kube-state-metrics
   tag: v2.5.0
+  sha: ""
   pullPolicy: IfNotPresent
 
 imagePullSecrets: []

```

## 4.15.0

**Release date:** 2022-08-02

![AppVersion: 2.5.0](https://img.shields.io/static/v1?label=AppVersion&message=2.5.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Allow adding volumeMounts and volumes in deployment (#2296)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index d5962e2c..cf74d67e 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -249,3 +249,16 @@ selfMonitor:
   # telemetryHost: 0.0.0.0
   # telemetryPort: 8081
   # telemetryNodePort: 0
+
+# volumeMounts are used to add custom volume mounts to deployment.
+# See example below
+volumeMounts: []
+#  - mountPath: /etc/config
+#    name: config-volume
+
+# volumes are used to add custom volumes to deployment
+# See example below
+volumes: []
+#  - configMap:
+#      name: cm-for-volume
+#    name: config-volume

```

## 4.14.0

**Release date:** 2022-07-28

![AppVersion: 2.5.0](https://img.shields.io/static/v1?label=AppVersion&message=2.5.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] - Add option to only use the release namespace (#2298)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 72dbc2c9..d5962e2c 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -209,6 +209,10 @@ kubeconfig:
   # base64 encoded kube-config file
   secret:
 
+# Enable only the release namespace for collecting resources. By default all namespaces are collected.
+# If releaseNamespace and namespaces are both set only releaseNamespace will be used.
+releaseNamespace: false
+
 # Comma-separated list of namespaces to be enabled for collecting resources. By default all namespaces are collected.
 namespaces: ""
 

```

## 4.13.0

**Release date:** 2022-07-08

![AppVersion: 2.5.0](https://img.shields.io/static/v1?label=AppVersion&message=2.5.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Configurable extra rules for KSM ClusterRole (#2249)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 60e16b2e..72dbc2c9 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -50,6 +50,13 @@ rbac:
   # If set to false - Run without Cluteradmin privs needed - ONLY works if namespace is also set (if useExistingRole is set this name is used as ClusterRole or Role to bind to)
   useClusterRole: true
 
+  # Add permissions for CustomResources' apiGroups in Role/ClusterRole. Should be used in conjunction with Custom Resource State Metrics configuration
+  # Example:
+  # - apiGroups: ["monitoring.coreos.com"]
+  #   resources: ["prometheuses"]
+  #   verbs: ["list", "watch"]
+  extraRules: []
+
 serviceAccount:
   # Specifies whether a ServiceAccount should be created, require rbac true
   create: true

```

## 4.12.0

**Release date:** 2022-07-08

![AppVersion: 2.5.0](https://img.shields.io/static/v1?label=AppVersion&message=2.5.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] adds deployment annotations for kube-state-metrics (#2247)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 606fcaca..60e16b2e 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -124,6 +124,9 @@ tolerations: []
 ## Ref: https://kubernetes.io/docs/concepts/workloads/pods/pod-topology-spread-constraints/
 topologySpreadConstraints: []
 
+# Annotations to be added to the deployment/statefulset
+annotations: {}
+
 # Annotations to be added to the pod
 podAnnotations: {}
 

```

## 4.11.0

**Release date:** 2022-07-04

![AppVersion: 2.5.0](https://img.shields.io/static/v1?label=AppVersion&message=2.5.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Add selfMonitor.telemetryNodePort field to values (#2222)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index b85820bf..606fcaca 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -229,7 +229,9 @@ kubeTargetVersionOverride: ""
 
 # Enable self metrics configuration for service and Service Monitor
 # Default values for telemetry configuration can be overridden
+# If you set telemetryNodePort, you must also set service.type to NodePort
 selfMonitor:
   enabled: false
   # telemetryHost: 0.0.0.0
   # telemetryPort: 8081
+  # telemetryNodePort: 0

```

## 4.10.0

**Release date:** 2022-06-29

![AppVersion: 2.5.0](https://img.shields.io/static/v1?label=AppVersion&message=2.5.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Support for topology spread constraints (#2213)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 1acd8a1a..b85820bf 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -120,6 +120,10 @@ affinity: {}
 ## Ref: https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/
 tolerations: []
 
+## Topology spread constraints for pod assignment
+## Ref: https://kubernetes.io/docs/concepts/workloads/pods/pod-topology-spread-constraints/
+topologySpreadConstraints: []
+
 # Annotations to be added to the pod
 podAnnotations: {}
 

```

## 4.9.3

**Release date:** 2022-06-29

![AppVersion: 2.5.0](https://img.shields.io/static/v1?label=AppVersion&message=2.5.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Use template name for container (#2210)

### Default value changes

```diff
# No changes in this release
```

## 4.9.2

**Release date:** 2022-06-17

![AppVersion: 2.5.0](https://img.shields.io/static/v1?label=AppVersion&message=2.5.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* fix(kube-state-metrics): render extraArgs as YAML to allow multi-line strings in arguments (#2167)

### Default value changes

```diff
# No changes in this release
```

## 4.9.1

**Release date:** 2022-06-17

![AppVersion: 2.5.0](https://img.shields.io/static/v1?label=AppVersion&message=2.5.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Upgrade pdb version to apiVersion: policy/v1 (#2166)

### Default value changes

```diff
# No changes in this release
```

## 4.9.0

**Release date:** 2022-06-10

![AppVersion: 2.5.0](https://img.shields.io/static/v1?label=AppVersion&message=2.5.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Add scheme and tlsConfig arguments to serviceMonitor (#2145)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 0aecd7cb..1acd8a1a 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -77,6 +77,8 @@ prometheus:
     honorLabels: false
     metricRelabelings: []
     relabelings: []
+    scheme: ""
+    tlsConfig: {}
 
 ## Specify if a Pod Security Policy for kube-state-metrics must be created
 ## Ref: https://kubernetes.io/docs/concepts/policy/pod-security-policy/

```

## 4.8.1

**Release date:** 2022-06-07

![AppVersion: 2.5.0](https://img.shields.io/static/v1?label=AppVersion&message=2.5.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Allow complete pod securityContext to kube-state-metrics deploy chart (#2119)

### Default value changes

```diff
# No changes in this release
```

## 4.8.0

**Release date:** 2022-06-07

![AppVersion: 2.5.0](https://img.shields.io/static/v1?label=AppVersion&message=2.5.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-stat-metrics] Version bump to 2.5.0 (#2128)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 89e0da79..0aecd7cb 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -1,8 +1,8 @@
 # Default values for kube-state-metrics.
 prometheusScrape: true
 image:
-  repository: k8s.gcr.io/kube-state-metrics/kube-state-metrics
-  tag: v2.4.1
+  repository: registry.k8s.io/kube-state-metrics/kube-state-metrics
+  tag: v2.5.0
   pullPolicy: IfNotPresent
 
 imagePullSecrets: []

```

## 4.7.0

**Release date:** 2022-02-25

![AppVersion: 2.4.1](https://img.shields.io/static/v1?label=AppVersion&message=2.4.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Version bump to 2.4.1 (#1829)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index e2250b68..89e0da79 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -2,7 +2,7 @@
 prometheusScrape: true
 image:
   repository: k8s.gcr.io/kube-state-metrics/kube-state-metrics
-  tag: v2.3.0
+  tag: v2.4.1
   pullPolicy: IfNotPresent
 
 imagePullSecrets: []

```

## 4.6.0

**Release date:** 2022-02-25

![AppVersion: 2.3.0](https://img.shields.io/static/v1?label=AppVersion&message=2.3.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Allow specifying clusterIP (#1827)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 5d6c0872..e2250b68 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -28,6 +28,7 @@ service:
   type: ClusterIP
   nodePort: 0
   loadBalancerIP: ""
+  clusterIP: ""
   annotations: {}
 
 ## Additional labels to add to all resources

```

## 4.5.0

**Release date:** 2022-02-10

![AppVersion: 2.3.0](https://img.shields.io/static/v1?label=AppVersion&message=2.3.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Add namespacesDenylist (#1789)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index aced5cf4..5d6c0872 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -195,6 +195,10 @@ kubeconfig:
 # Comma-separated list of namespaces to be enabled for collecting resources. By default all namespaces are collected.
 namespaces: ""
 
+# Comma-separated list of namespaces not to be enabled. If namespaces and namespaces-denylist are both set,
+# only namespaces that are excluded in namespaces-denylist will be used.
+namespacesDenylist: ""
+
 ## Override the deployment namespace
 ##
 namespaceOverride: ""

```

## 4.4.3

**Release date:** 2022-02-02

![AppVersion: 2.3.0](https://img.shields.io/static/v1?label=AppVersion&message=2.3.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] fix serviceMonitor override indentation (#1763)

### Default value changes

```diff
# No changes in this release
```

## 4.4.2

**Release date:** 2022-01-31

![AppVersion: 2.3.0](https://img.shields.io/static/v1?label=AppVersion&message=2.3.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Fix an issue where cluster role would be generated multiple times when (#1748)

### Default value changes

```diff
# No changes in this release
```

## 4.4.1

**Release date:** 2022-01-14

![AppVersion: 2.3.0](https://img.shields.io/static/v1?label=AppVersion&message=2.3.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Minor change to label templating, to stop errors when it is a subchart and RoleBindings are toggled (#1712)

### Default value changes

```diff
# No changes in this release
```

## 4.4.0

**Release date:** 2022-01-12

![AppVersion: 2.3.0](https://img.shields.io/static/v1?label=AppVersion&message=2.3.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] reintroduce service monitor selectorOverride (#1698)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index e4384d88..aced5cf4 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -72,6 +72,7 @@ prometheus:
     interval: ""
     scrapeTimeout: ""
     proxyUrl: ""
+    selectorOverride: {}
     honorLabels: false
     metricRelabelings: []
     relabelings: []

```

## 4.3.0

**Release date:** 2022-01-07

![AppVersion: 2.3.0](https://img.shields.io/static/v1?label=AppVersion&message=2.3.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] added release label option (#1683)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 7a2cc6a1..e4384d88 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -34,6 +34,9 @@ service:
 customLabels: {}
   # app: kube-state-metrics
 
+## set to true to add the release label so scraping of the servicemonitor with kube-prometheus-stack works out of the box
+releaseLabel: false
+
 hostNetwork: false
 
 rbac:

```

## 4.2.2

**Release date:** 2022-01-05

![AppVersion: 2.3.0](https://img.shields.io/static/v1?label=AppVersion&message=2.3.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] added dotdc as a maintainer (#1675)

### Default value changes

```diff
# No changes in this release
```

## 4.2.1

**Release date:** 2022-01-04

![AppVersion: 2.3.0](https://img.shields.io/static/v1?label=AppVersion&message=2.3.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Fix code duplication (#1665)

### Default value changes

```diff
# No changes in this release
```

## 4.2.0

**Release date:** 2021-12-18

![AppVersion: 2.3.0](https://img.shields.io/static/v1?label=AppVersion&message=2.3.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Version bump to 2.3.0 (#1600)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 5986ff07..7a2cc6a1 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -2,7 +2,7 @@
 prometheusScrape: true
 image:
   repository: k8s.gcr.io/kube-state-metrics/kube-state-metrics
-  tag: v2.2.4
+  tag: v2.3.0
   pullPolicy: IfNotPresent
 
 imagePullSecrets: []

```

## 4.1.2

**Release date:** 2021-12-13

![AppVersion: 2.2.4](https://img.shields.io/static/v1?label=AppVersion&message=2.2.4&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Enable custom ServiceMonitor jobLabel (#1592)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index da09c892..5986ff07 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -65,6 +65,7 @@ prometheus:
     enabled: false
     additionalLabels: {}
     namespace: ""
+    jobLabel: ""
     interval: ""
     scrapeTimeout: ""
     proxyUrl: ""

```

## 4.1.1

**Release date:** 2021-12-08

![AppVersion: 2.2.4](https://img.shields.io/static/v1?label=AppVersion&message=2.2.4&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Add indentation for customLabels in role template (#1574)

### Default value changes

```diff
# No changes in this release
```

## 4.1.0

**Release date:** 2021-12-06

![AppVersion: 2.2.4](https://img.shields.io/static/v1?label=AppVersion&message=2.2.4&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Improve ServiceMonitor capabilities (#1556)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 70ba3568..da09c892 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -65,6 +65,9 @@ prometheus:
     enabled: false
     additionalLabels: {}
     namespace: ""
+    interval: ""
+    scrapeTimeout: ""
+    proxyUrl: ""
     honorLabels: false
     metricRelabelings: []
     relabelings: []

```

## 4.0.2

**Release date:** 2021-11-21

![AppVersion: 2.2.4](https://img.shields.io/static/v1?label=AppVersion&message=2.2.4&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Bump to v2.2.4 (#1499)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 53edffc9..70ba3568 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -2,7 +2,7 @@
 prometheusScrape: true
 image:
   repository: k8s.gcr.io/kube-state-metrics/kube-state-metrics
-  tag: v2.2.0
+  tag: v2.2.4
   pullPolicy: IfNotPresent
 
 imagePullSecrets: []

```

## 4.0.1

**Release date:** 2021-11-20

![AppVersion: 2.2.0](https://img.shields.io/static/v1?label=AppVersion&message=2.2.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Named deployment ports (#1514)

### Default value changes

```diff
# No changes in this release
```

## 4.0.0

**Release date:** 2021-10-25

![AppVersion: 2.2.0](https://img.shields.io/static/v1?label=AppVersion&message=2.2.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Add: Kubernetes recommanded labels and custom labels (#927)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 917db1d8..53edffc9 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -30,7 +30,9 @@ service:
   loadBalancerIP: ""
   annotations: {}
 
+## Additional labels to add to all resources
 customLabels: {}
+  # app: kube-state-metrics
 
 hostNetwork: false
 

```

## 3.5.2

**Release date:** 2021-09-27

![AppVersion: 2.2.0](https://img.shields.io/static/v1?label=AppVersion&message=2.2.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Add metricRelabelings and relabelings options (#1374)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 72cdc8e8..917db1d8 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -64,6 +64,8 @@ prometheus:
     additionalLabels: {}
     namespace: ""
     honorLabels: false
+    metricRelabelings: []
+    relabelings: []
 
 ## Specify if a Pod Security Policy for kube-state-metrics must be created
 ## Ref: https://kubernetes.io/docs/concepts/policy/pod-security-policy/

```

## 3.5.1

**Release date:** 2021-09-21

![AppVersion: 2.2.0](https://img.shields.io/static/v1?label=AppVersion&message=2.2.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Enable configuring kubernetes annotation metrics (#1362)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 86e3d77e..72cdc8e8 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -133,6 +133,15 @@ metricDenylist: []
 metricLabelsAllowlist: []
   # - namespaces=[k8s-label-1,k8s-label-n]
 
+# Comma-separated list of Kubernetes annotations keys that will be used in the resource'
+# labels metric. By default the metric contains only name and namespace labels.
+# To include additional annotations provide a list of resource names in their plural form and Kubernetes
+# annotation keys you would like to allow for them (Example: '=namespaces=[kubernetes.io/team,...],pods=[kubernetes.io/team],...)'.
+# A single '*' can be provided per resource instead to allow any annotations, but that has
+# severe performance implications (Example: '=pods=[*]').
+metricAnnotationsAllowList: []
+  # - pods=[k8s-annotation-1,k8s-annotation-n]
+
 # Available collectors for kube-state-metrics.
 # By default, all available resources are enabled, comment out to disable.
 collectors:

```

## 3.5.0

**Release date:** 2021-08-24

![AppVersion: 2.2.0](https://img.shields.io/static/v1?label=AppVersion&message=2.2.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Version bump to 2.2.0 (#1285)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index a09e47e3..86e3d77e 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -2,7 +2,7 @@
 prometheusScrape: true
 image:
   repository: k8s.gcr.io/kube-state-metrics/kube-state-metrics
-  tag: v2.1.1
+  tag: v2.2.0
   pullPolicy: IfNotPresent
 
 imagePullSecrets: []

```

## 3.4.2

**Release date:** 2021-08-04

![AppVersion: 2.1.1](https://img.shields.io/static/v1?label=AppVersion&message=2.1.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Version bump to 2.1.1 (#1221)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 3c34e413..a09e47e3 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -2,7 +2,7 @@
 prometheusScrape: true
 image:
   repository: k8s.gcr.io/kube-state-metrics/kube-state-metrics
-  tag: v2.1.0
+  tag: v2.1.1
   pullPolicy: IfNotPresent
 
 imagePullSecrets: []

```

## 3.4.1

**Release date:** 2021-07-20

![AppVersion: 2.1.0](https://img.shields.io/static/v1?label=AppVersion&message=2.1.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Disable VPA by default (#1181)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 6b794dac..3c34e413 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -162,8 +162,8 @@ collectors:
   - statefulsets
   - storageclasses
   - validatingwebhookconfigurations
-  - verticalpodautoscalers
   - volumeattachments
+  # - verticalpodautoscalers # not a default resource, see also: https://github.com/kubernetes/kube-state-metrics#enabling-verticalpodautoscalers
 
 # Enabling kubeconfig will pass the --kubeconfig argument to the container
 kubeconfig:

```

## 3.4.0

**Release date:** 2021-06-29

![AppVersion: 2.1.0](https://img.shields.io/static/v1?label=AppVersion&message=2.1.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Bump version to 2.1.0 (#1112)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index a9853bae..6b794dac 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -2,7 +2,7 @@
 prometheusScrape: true
 image:
   repository: k8s.gcr.io/kube-state-metrics/kube-state-metrics
-  tag: v2.0.0
+  tag: v2.1.0
   pullPolicy: IfNotPresent
 
 imagePullSecrets: []

```

## 3.3.1

**Release date:** 2021-06-24

![AppVersion: 2.0.0](https://img.shields.io/static/v1?label=AppVersion&message=2.0.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] allow for namespaces flag to be passed as a list (#1105)

### Default value changes

```diff
# No changes in this release
```

## 3.3.0

**Release date:** 2021-06-23

![AppVersion: 2.0.0](https://img.shields.io/static/v1?label=AppVersion&message=2.0.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Fix and improve template and include new features (#997)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index a02c1181..a9853bae 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -114,37 +114,56 @@ podAnnotations: {}
 # Ref: https://kubernetes.io/docs/tasks/run-application/configure-pdb/
 podDisruptionBudget: {}
 
-# Available collectors for kube-state-metrics. By default all available
-# resources are enabled.
+# Comma-separated list of metrics to be exposed.
+# This list comprises of exact metric names and/or regex patterns.
+# The allowlist and denylist are mutually exclusive.
+metricAllowlist: []
+
+# Comma-separated list of metrics not to be enabled.
+# This list comprises of exact metric names and/or regex patterns.
+# The allowlist and denylist are mutually exclusive.
+metricDenylist: []
+
+# Comma-separated list of additional Kubernetes label keys that will be used in the resource's
+# labels metric. By default the metric contains only name and namespace labels.
+# To include additional labels, provide a list of resource names in their plural form and Kubernetes
+# label keys you would like to allow for them (Example: '=namespaces=[k8s-label-1,k8s-label-n,...],pods=[app],...)'.
+# A single '*' can be provided per resource instead to allow any labels, but that has
+# severe performance implications (Example: '=pods=[*]').
+metricLabelsAllowlist: []
+  # - namespaces=[k8s-label-1,k8s-label-n]
+
+# Available collectors for kube-state-metrics.
+# By default, all available resources are enabled, comment out to disable.
 collectors:
-  certificatesigningrequests: true
-  configmaps: true
-  cronjobs: true
-  daemonsets: true
-  deployments: true
-  endpoints: true
-  horizontalpodautoscalers: true
-  ingresses: true
-  jobs: true
-  limitranges: true
-  mutatingwebhookconfigurations: true
-  namespaces: true
-  networkpolicies: true
-  nodes: true
-  persistentvolumeclaims: true
-  persistentvolumes: true
-  poddisruptionbudgets: true
-  pods: true
-  replicasets: true
-  replicationcontrollers: true
-  resourcequotas: true
-  secrets: true
-  services: true
-  statefulsets: true
-  storageclasses: true
-  validatingwebhookconfigurations: true
-  verticalpodautoscalers: false
-  volumeattachments: true
+  - certificatesigningrequests
+  - configmaps
+  - cronjobs
+  - daemonsets
+  - deployments
+  - endpoints
+  - horizontalpodautoscalers
+  - ingresses
+  - jobs
+  - limitranges
+  - mutatingwebhookconfigurations
+  - namespaces
+  - networkpolicies
+  - nodes
+  - persistentvolumeclaims
+  - persistentvolumes
+  - poddisruptionbudgets
+  - pods
+  - replicasets
+  - replicationcontrollers
+  - resourcequotas
+  - secrets
+  - services
+  - statefulsets
+  - storageclasses
+  - validatingwebhookconfigurations
+  - verticalpodautoscalers
+  - volumeattachments
 
 # Enabling kubeconfig will pass the --kubeconfig argument to the container
 kubeconfig:

```

## 3.2.2

**Release date:** 2021-06-18

![AppVersion: 2.0.0](https://img.shields.io/static/v1?label=AppVersion&message=2.0.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] allow for namespaces flag to be passed as a list (#1070)

### Default value changes

```diff
# No changes in this release
```

## 3.2.1

**Release date:** 2021-06-11

![AppVersion: 2.0.0](https://img.shields.io/static/v1?label=AppVersion&message=2.0.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Fix: Use serviceAccountName template helper where it should (#1063)

### Default value changes

```diff
# No changes in this release
```

## 3.2.0

**Release date:** 2021-06-06

![AppVersion: 2.0.0](https://img.shields.io/static/v1?label=AppVersion&message=2.0.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Add containerSecurityContext (#1017)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 62e178d2..a02c1181 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -88,6 +88,11 @@ securityContext:
   runAsUser: 65534
   fsGroup: 65534
 
+## Specify security settings for a Container
+## Allows overrides and additional options compared to (Pod) securityContext
+## Ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/#set-the-security-context-for-a-container
+containerSecurityContext: {}
+
 ## Node labels for pod assignment
 ## Ref: https://kubernetes.io/docs/user-guide/node-selection/
 nodeSelector: {}

```

## 3.1.1

**Release date:** 2021-05-31

![AppVersion: 2.0.0](https://img.shields.io/static/v1?label=AppVersion&message=2.0.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] fix documentation for helm show values (#1025)

### Default value changes

```diff
# No changes in this release
```

## 3.1.0

**Release date:** 2021-05-18

![AppVersion: 2.0.0](https://img.shields.io/static/v1?label=AppVersion&message=2.0.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Fix rbac.useClusterRole=false behavior (#953)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 042c91de..62e178d2 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -147,8 +147,8 @@ kubeconfig:
   # base64 encoded kube-config file
   secret:
 
-# Namespaces to be enabled for collecting resources. By default all namespaces are collected.
-# namespaces: ""
+# Comma-separated list of namespaces to be enabled for collecting resources. By default all namespaces are collected.
+namespaces: ""
 
 ## Override the deployment namespace
 ##

```

## 3.0.2

**Release date:** 2021-05-07

![AppVersion: 2.0.0](https://img.shields.io/static/v1?label=AppVersion&message=2.0.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Allow templating inside .Values.namespace (#904)

### Default value changes

```diff
# No changes in this release
```

## 3.0.1

**Release date:** 2021-05-04

![AppVersion: 2.0.0](https://img.shields.io/static/v1?label=AppVersion&message=2.0.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Ensure consistent telemetry configuration (#903)

### Default value changes

```diff
# No changes in this release
```

## 3.0.0

**Release date:** 2021-04-27

![AppVersion: 2.0.0](https://img.shields.io/static/v1?label=AppVersion&message=2.0.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Version bump to 2.0.0 (#891)
* [kube-state-metrics] Remove license
* [kube-state-metrics] Fix CI issues

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 9522cfe0..042c91de 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -2,7 +2,7 @@
 prometheusScrape: true
 image:
   repository: k8s.gcr.io/kube-state-metrics/kube-state-metrics
-  tag: v1.9.8
+  tag: v2.0.0
   pullPolicy: IfNotPresent
 
 imagePullSecrets: []
@@ -110,7 +110,7 @@ podAnnotations: {}
 podDisruptionBudget: {}
 
 # Available collectors for kube-state-metrics. By default all available
-# collectors are enabled.
+# resources are enabled.
 collectors:
   certificatesigningrequests: true
   configmaps: true
@@ -147,8 +147,8 @@ kubeconfig:
   # base64 encoded kube-config file
   secret:
 
-# Namespace to be enabled for collecting resources. By default all namespaces are collected.
-# namespace: ""
+# Namespaces to be enabled for collecting resources. By default all namespaces are collected.
+# namespaces: ""
 
 ## Override the deployment namespace
 ##
@@ -172,7 +172,7 @@ resources: {}
 kubeTargetVersionOverride: ""
 
 # Enable self metrics configuration for service and Service Monitor
-# Default values for telemetry configuration can be overriden
+# Default values for telemetry configuration can be overridden
 selfMonitor:
   enabled: false
   # telemetryHost: 0.0.0.0

```

## 2.13.2

**Release date:** 2021-04-07

![AppVersion: 1.9.8](https://img.shields.io/static/v1?label=AppVersion&message=1.9.8&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Fix role.yaml for helm2 compatibility

### Default value changes

```diff
# No changes in this release
```

## 2.13.1

**Release date:** 2021-03-25

![AppVersion: 1.9.8](https://img.shields.io/static/v1?label=AppVersion&message=1.9.8&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* update RBAC manifests to v1

### Default value changes

```diff
# No changes in this release
```

## 2.13.0

**Release date:** 2021-02-23

![AppVersion: 1.9.8](https://img.shields.io/static/v1?label=AppVersion&message=1.9.8&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* chart: Upgrade to multi-arch v1.9.8 image
* chart: Add AppVersion to labels

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 4f8a498e..9522cfe0 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -1,8 +1,8 @@
 # Default values for kube-state-metrics.
 prometheusScrape: true
 image:
-  repository: quay.io/coreos/kube-state-metrics
-  tag: v1.9.7
+  repository: k8s.gcr.io/kube-state-metrics/kube-state-metrics
+  tag: v1.9.8
   pullPolicy: IfNotPresent
 
 imagePullSecrets: []

```

## 2.12.2

**Release date:** 2021-02-11

![AppVersion: 1.9.7](https://img.shields.io/static/v1?label=AppVersion&message=1.9.7&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* chart: Add custom labels to pdb and servicemonitor

### Default value changes

```diff
# No changes in this release
```

## 2.12.1

**Release date:** 2021-02-03

![AppVersion: 1.9.7](https://img.shields.io/static/v1?label=AppVersion&message=1.9.7&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* allow for namespace flag to be passed as a list

### Default value changes

```diff
# No changes in this release
```

## 2.11.0

**Release date:** 2021-01-23

![AppVersion: 1.9.7](https://img.shields.io/static/v1?label=AppVersion&message=1.9.7&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* added: support for providing extraArgs to ksm
* Fix create role

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 66972fca..4f8a498e 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -17,6 +17,11 @@ autosharding:
 
 replicas: 1
 
+# List of additional cli arguments to configure kube-state-metrics
+# for example: --enable-gzip-encoding, --log-file, etc.
+# all the possible args can be found here: https://github.com/kubernetes/kube-state-metrics/blob/master/docs/cli-arguments.md
+extraArgs: []
+
 service:
   port: 8080
   # Default to clusterIP for backward compatibility

```

## 2.10.1

**Release date:** 2021-01-26

![AppVersion: 1.9.7](https://img.shields.io/static/v1?label=AppVersion&message=1.9.7&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Hardcoded telemetry internal port and bumped version

### Default value changes

```diff
# No changes in this release
```

## 2.10.0

**Release date:** 2021-01-26

![AppVersion: 1.9.7](https://img.shields.io/static/v1?label=AppVersion&message=1.9.7&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Add useExistingRole and useClusterrole options - allowing installation in specific namespaces only when user has no access to entire cluster (like in typical OpenShift environments).
* Fix spacing issue

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 5bc2ebea..66972fca 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -33,6 +33,12 @@ rbac:
   # If true, create & use RBAC resources
   create: true
 
+  # Set to a rolename to use existing role - skipping role creating - but still doing serviceaccount and rolebinding to it, rolename set here.
+  # useExistingRole: your-existing-role
+
+  # If set to false - Run without Cluteradmin privs needed - ONLY works if namespace is also set (if useExistingRole is set this name is used as ClusterRole or Role to bind to)
+  useClusterRole: true
+
 serviceAccount:
   # Specifies whether a ServiceAccount should be created, require rbac true
   create: true

```

## 2.9.8

**Release date:** 2021-01-11

![AppVersion: 1.9.7](https://img.shields.io/static/v1?label=AppVersion&message=1.9.7&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* bumped chart version

### Default value changes

```diff
# No changes in this release
```

## 2.9.7

**Release date:** 2021-01-10

![AppVersion: 1.9.7](https://img.shields.io/static/v1?label=AppVersion&message=1.9.7&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Release chart to gh-pages
* Update charts/kube-state-metrics/README.md
* charts: point out that the repo isn't set up yet
* Corrected var check

### Default value changes

```diff
# No changes in this release
```

## 2.9.6

**Release date:** 2021-01-05

![AppVersion: 1.9.7](https://img.shields.io/static/v1?label=AppVersion&message=1.9.7&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* bumped Helm version
* Added Helm Chart support for self monitor

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 3edd4c81..5bc2ebea 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -159,3 +159,10 @@ resources: {}
 ## For example: kubeTargetVersionOverride: 1.14.9
 ##
 kubeTargetVersionOverride: ""
+
+# Enable self metrics configuration for service and Service Monitor
+# Default values for telemetry configuration can be overriden
+selfMonitor:
+  enabled: false
+  # telemetryHost: 0.0.0.0
+  # telemetryPort: 8081

```

## 2.9.5

**Release date:** 2020-12-18

![AppVersion: 1.9.7](https://img.shields.io/static/v1?label=AppVersion&message=1.9.7&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Bump version to re-release chart after branch protection was removed

### Default value changes

```diff
# No changes in this release
```

## 2.9.4

**Release date:** 2020-12-17

![AppVersion: 1.9.7](https://img.shields.io/static/v1?label=AppVersion&message=1.9.7&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Bump chart PATCH version
* Remove fiunchinho per these comments:
* Retain license from helm/charts, where the previous commits were moved from.

### Default value changes

```diff
# No changes in this release
```

## 2.9.3

**Release date:** 2020-11-02

![AppVersion: 1.9.7](https://img.shields.io/static/v1?label=AppVersion&message=1.9.7&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] Update apiVersion for clusterrole (#23968)

### Default value changes

```diff
# No changes in this release
```

## 2.9.2

**Release date:** 2020-09-18

![AppVersion: 1.9.7](https://img.shields.io/static/v1?label=AppVersion&message=1.9.7&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Bump chart PATCH version
* Clean up chart README. To-do: remove Helm 2 command documentation after Nov 13 2020

### Default value changes

```diff
# No changes in this release
```

## 2.9.1

**Release date:** 2020-08-21

![AppVersion: 1.9.7](https://img.shields.io/static/v1?label=AppVersion&message=1.9.7&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] Add kubeconfig arg for kube-state-metrics  (#23380)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 92de8e22..3edd4c81 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -130,6 +130,12 @@ collectors:
   verticalpodautoscalers: false
   volumeattachments: true
 
+# Enabling kubeconfig will pass the --kubeconfig argument to the container
+kubeconfig:
+  enabled: false
+  # base64 encoded kube-config file
+  secret:
+
 # Namespace to be enabled for collecting resources. By default all namespaces are collected.
 # namespace: ""
 

```

## 2.8.14

**Release date:** 2020-08-02

![AppVersion: 1.9.7](https://img.shields.io/static/v1?label=AppVersion&message=1.9.7&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] Add annotations for serviceAccount (#23411)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 08d04356..92de8e22 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -42,6 +42,10 @@ serviceAccount:
   # Reference to one or more secrets to be used when pulling images
   # ref: https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/
   imagePullSecrets: []
+  # ServiceAccount annotations.
+  # Use case: AWS EKS IAM roles for service accounts
+  # ref: https://docs.aws.amazon.com/eks/latest/userguide/specify-service-account-role.html
+  annotations: {}
 
 prometheus:
   monitor:

```

## 2.8.13

**Release date:** 2020-07-20

![AppVersion: 1.9.7](https://img.shields.io/static/v1?label=AppVersion&message=1.9.7&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] added resource config to values.yaml (#22850)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 11f48fb7..08d04356 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -133,6 +133,18 @@ collectors:
 ##
 namespaceOverride: ""
 
+resources: {}
+  # We usually recommend not to specify default resources and to leave this as a conscious
+  # choice for the user. This also increases chances charts run on environments with little
+  # resources, such as Minikube. If you do want to specify resources, uncomment the following
+  # lines, adjust them as necessary, and remove the curly braces after 'resources:'.
+  # limits:
+  #  cpu: 100m
+  #  memory: 64Mi
+  # requests:
+  #  cpu: 10m
+  #  memory: 32Mi
+
 ## Provide a k8s version to define apiGroups for podSecurityPolicy Cluster Role.
 ## For example: kubeTargetVersionOverride: 1.14.9
 ##

```

## 2.8.12

**Release date:** 2020-07-18

![AppVersion: 1.9.7](https://img.shields.io/static/v1?label=AppVersion&message=1.9.7&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] Fix STS role for Autosharding (#23210)

### Default value changes

```diff
# No changes in this release
```

## 2.8.11

**Release date:** 2020-06-22

![AppVersion: 1.9.7](https://img.shields.io/static/v1?label=AppVersion&message=1.9.7&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] Document `kubeTargetVersionOverride` (#22792)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index acf5fdb6..11f48fb7 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -132,3 +132,8 @@ collectors:
 ## Override the deployment namespace
 ##
 namespaceOverride: ""
+
+## Provide a k8s version to define apiGroups for podSecurityPolicy Cluster Role.
+## For example: kubeTargetVersionOverride: 1.14.9
+##
+kubeTargetVersionOverride: ""

```

## 2.8.10

**Release date:** 2020-06-16

![AppVersion: 1.9.7](https://img.shields.io/static/v1?label=AppVersion&message=1.9.7&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] Version bump to 1.9.7 (#22649)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 5f547daf..acf5fdb6 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -2,7 +2,7 @@
 prometheusScrape: true
 image:
   repository: quay.io/coreos/kube-state-metrics
-  tag: v1.9.6
+  tag: v1.9.7
   pullPolicy: IfNotPresent
 
 imagePullSecrets: []

```

## 2.8.9

**Release date:** 2020-06-12

![AppVersion: 1.9.6](https://img.shields.io/static/v1?label=AppVersion&message=1.9.6&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] Support imagePullSecrets (#22295)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 75ddb9e4..5f547daf 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -5,6 +5,9 @@ image:
   tag: v1.9.6
   pullPolicy: IfNotPresent
 
+imagePullSecrets: []
+# - name: "image-pull-secret"
+
 # If set to true, this will deploy kube-state-metrics as a StatefulSet and the data
 # will be automatically sharded across <.Values.replicas> pods using the built-in
 # autodiscovery feature: https://github.com/kubernetes/kube-state-metrics#automated-sharding

```

## 2.8.8

**Release date:** 2020-06-02

![AppVersion: 1.9.6](https://img.shields.io/static/v1?label=AppVersion&message=1.9.6&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] Set runAsGroup (#22602)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 2ab34d7c..75ddb9e4 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -66,6 +66,7 @@ podSecurityPolicy:
 
 securityContext:
   enabled: true
+  runAsGroup: 65534
   runAsUser: 65534
   fsGroup: 65534
 

```

## 2.8.7

**Release date:** 2020-05-24

![AppVersion: 1.9.6](https://img.shields.io/static/v1?label=AppVersion&message=1.9.6&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] Allow specification of additional volume types in podSecurityPolicy of kube-state-metrics (#22122)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index fa210388..2ab34d7c 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -62,6 +62,7 @@ podSecurityPolicy:
     # seccomp.security.alpha.kubernetes.io/defaultProfileName: 'docker/default'
     # apparmor.security.beta.kubernetes.io/defaultProfileName: 'runtime/default'
 
+  additionalVolumes: []
 
 securityContext:
   enabled: true

```

## 2.8.6

**Release date:** 2020-05-24

![AppVersion: 1.9.6](https://img.shields.io/static/v1?label=AppVersion&message=1.9.6&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics]disable VPA collectors by default (#22495)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 934d7f79..fa210388 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -118,7 +118,7 @@ collectors:
   statefulsets: true
   storageclasses: true
   validatingwebhookconfigurations: true
-  verticalpodautoscalers: true
+  verticalpodautoscalers: false
   volumeattachments: true
 
 # Namespace to be enabled for collecting resources. By default all namespaces are collected.

```

## 2.8.5

**Release date:** 2020-05-16

![AppVersion: 1.9.6](https://img.shields.io/static/v1?label=AppVersion&message=1.9.6&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] Version bump to 1.9.6 (#22420)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 11f20fa3..934d7f79 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -2,7 +2,7 @@
 prometheusScrape: true
 image:
   repository: quay.io/coreos/kube-state-metrics
-  tag: v1.9.5
+  tag: v1.9.6
   pullPolicy: IfNotPresent
 
 # If set to true, this will deploy kube-state-metrics as a StatefulSet and the data

```

## 2.8.4

**Release date:** 2020-04-24

![AppVersion: 1.9.5](https://img.shields.io/static/v1?label=AppVersion&message=1.9.5&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] Use apiGroup policy rather than extension for k8s > 1.15 in ClusterRole psp-kube-state-metrics (#21563)

### Default value changes

```diff
# No changes in this release
```

## 2.8.3

**Release date:** 2020-04-24

![AppVersion: 1.9.5](https://img.shields.io/static/v1?label=AppVersion&message=1.9.5&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] use the correct rbac values parameter (#22089)

### Default value changes

```diff
# No changes in this release
```

## 2.8.2

**Release date:** 2020-03-14

![AppVersion: 1.9.5](https://img.shields.io/static/v1?label=AppVersion&message=1.9.5&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] Add PodDisruptionBudget (#21340)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index ebcebed5..11f20fa3 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -86,6 +86,9 @@ podAnnotations: {}
 ## Assign a PriorityClassName to pods if set
 # priorityClassName: ""
 
+# Ref: https://kubernetes.io/docs/tasks/run-application/configure-pdb/
+podDisruptionBudget: {}
+
 # Available collectors for kube-state-metrics. By default all available
 # collectors are enabled.
 collectors:

```

## 2.8.1

**Release date:** 2020-03-05

![AppVersion: 1.9.5](https://img.shields.io/static/v1?label=AppVersion&message=1.9.5&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] Only create PSP RBAC if rbac.create is true. (#19986)

### Default value changes

```diff
# No changes in this release
```

## 2.8.0

**Release date:** 2020-03-05

![AppVersion: 1.9.5](https://img.shields.io/static/v1?label=AppVersion&message=1.9.5&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] Enable all collectors by default (#21245)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 57c92553..ebcebed5 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -99,9 +99,9 @@ collectors:
   ingresses: true
   jobs: true
   limitranges: true
-  mutatingwebhookconfigurations: false
+  mutatingwebhookconfigurations: true
   namespaces: true
-  networkpolicies: false
+  networkpolicies: true
   nodes: true
   persistentvolumeclaims: true
   persistentvolumes: true
@@ -114,9 +114,9 @@ collectors:
   services: true
   statefulsets: true
   storageclasses: true
-  validatingwebhookconfigurations: false
-  verticalpodautoscalers: false
-  volumeattachments: false
+  validatingwebhookconfigurations: true
+  verticalpodautoscalers: true
+  volumeattachments: true
 
 # Namespace to be enabled for collecting resources. By default all namespaces are collected.
 # namespace: ""

```

## 2.7.2

**Release date:** 2020-03-04

![AppVersion: 1.9.5](https://img.shields.io/static/v1?label=AppVersion&message=1.9.5&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] Add missing permission for volumeattachments collector (#21079)

### Default value changes

```diff
# No changes in this release
```

## 2.7.1

**Release date:** 2020-02-20

![AppVersion: 1.9.5](https://img.shields.io/static/v1?label=AppVersion&message=1.9.5&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] Version bump to 1.9.5 (#20915)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index d8ceef52..57c92553 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -2,7 +2,7 @@
 prometheusScrape: true
 image:
   repository: quay.io/coreos/kube-state-metrics
-  tag: v1.9.4
+  tag: v1.9.5
   pullPolicy: IfNotPresent
 
 # If set to true, this will deploy kube-state-metrics as a StatefulSet and the data
@@ -116,7 +116,7 @@ collectors:
   storageclasses: true
   validatingwebhookconfigurations: false
   verticalpodautoscalers: false
-  volumeattachements: false
+  volumeattachments: false
 
 # Namespace to be enabled for collecting resources. By default all namespaces are collected.
 # namespace: ""

```

## 2.7.0

**Release date:** 2020-02-20

![AppVersion: 1.9.4](https://img.shields.io/static/v1?label=AppVersion&message=1.9.4&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] Support automatic sharding (#20791)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 5af53834..d8ceef52 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -5,6 +5,13 @@ image:
   tag: v1.9.4
   pullPolicy: IfNotPresent
 
+# If set to true, this will deploy kube-state-metrics as a StatefulSet and the data
+# will be automatically sharded across <.Values.replicas> pods using the built-in
+# autodiscovery feature: https://github.com/kubernetes/kube-state-metrics#automated-sharding
+# This is an experimental feature and there are no stability guarantees.
+autosharding:
+  enabled: false
+
 replicas: 1
 
 service:

```

## 2.6.4

**Release date:** 2020-02-04

![AppVersion: 1.9.4](https://img.shields.io/static/v1?label=AppVersion&message=1.9.4&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] Version bump to 1.9.4 (#20543)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index bf9453b8..5af53834 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -2,7 +2,7 @@
 prometheusScrape: true
 image:
   repository: quay.io/coreos/kube-state-metrics
-  tag: v1.9.3
+  tag: v1.9.4
   pullPolicy: IfNotPresent
 
 replicas: 1

```

## 2.6.3

**Release date:** 2020-01-22

![AppVersion: 1.9.3](https://img.shields.io/static/v1?label=AppVersion&message=1.9.3&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] Version bump to 1.9.3 (#20289)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 956c718c..bf9453b8 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -2,7 +2,7 @@
 prometheusScrape: true
 image:
   repository: quay.io/coreos/kube-state-metrics
-  tag: v1.9.2
+  tag: v1.9.3
   pullPolicy: IfNotPresent
 
 replicas: 1

```

## 2.6.2

**Release date:** 2020-01-14

![AppVersion: 1.9.2](https://img.shields.io/static/v1?label=AppVersion&message=1.9.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] Version bump to 1.9.2 (#20103)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 6ad6c28b..956c718c 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -2,7 +2,7 @@
 prometheusScrape: true
 image:
   repository: quay.io/coreos/kube-state-metrics
-  tag: v1.9.1
+  tag: v1.9.2
   pullPolicy: IfNotPresent
 
 replicas: 1

```

## 2.6.1

**Release date:** 2020-01-10

![AppVersion: 1.9.1](https://img.shields.io/static/v1?label=AppVersion&message=1.9.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] Version bump to 1.9.1 (#20028)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 747e5d8b..6ad6c28b 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -2,7 +2,7 @@
 prometheusScrape: true
 image:
   repository: quay.io/coreos/kube-state-metrics
-  tag: v1.9.0
+  tag: v1.9.1
   pullPolicy: IfNotPresent
 
 replicas: 1

```

## 2.6.0

**Release date:** 2020-01-02

![AppVersion: 1.9.0](https://img.shields.io/static/v1?label=AppVersion&message=1.9.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] Namespace override (#18990)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index a57a86da..747e5d8b 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -113,3 +113,7 @@ collectors:
 
 # Namespace to be enabled for collecting resources. By default all namespaces are collected.
 # namespace: ""
+
+## Override the deployment namespace
+##
+namespaceOverride: ""

```

## 2.5.0

**Release date:** 2019-12-20

![AppVersion: 1.9.0](https://img.shields.io/static/v1?label=AppVersion&message=1.9.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] Version bump to 1.9.0 (#19726)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 2bf5b308..a57a86da 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -2,7 +2,7 @@
 prometheusScrape: true
 image:
   repository: quay.io/coreos/kube-state-metrics
-  tag: v1.8.0
+  tag: v1.9.0
   pullPolicy: IfNotPresent
 
 replicas: 1
@@ -92,7 +92,9 @@ collectors:
   ingresses: true
   jobs: true
   limitranges: true
+  mutatingwebhookconfigurations: false
   namespaces: true
+  networkpolicies: false
   nodes: true
   persistentvolumeclaims: true
   persistentvolumes: true
@@ -105,7 +107,9 @@ collectors:
   services: true
   statefulsets: true
   storageclasses: true
+  validatingwebhookconfigurations: false
   verticalpodautoscalers: false
+  volumeattachements: false
 
 # Namespace to be enabled for collecting resources. By default all namespaces are collected.
 # namespace: ""

```

## 2.4.1

**Release date:** 2019-10-09

![AppVersion: 1.8.0](https://img.shields.io/static/v1?label=AppVersion&message=1.8.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] Avoid spec.template.spec.containers.0.resources with null value (#17817)

### Default value changes

```diff
# No changes in this release
```

## 2.4.0

**Release date:** 2019-10-01

![AppVersion: 1.8.0](https://img.shields.io/static/v1?label=AppVersion&message=1.8.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] Version bump to 1.8.0 / Add myself as maintainer (#17593)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index d5631b00..2bf5b308 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -2,7 +2,7 @@
 prometheusScrape: true
 image:
   repository: quay.io/coreos/kube-state-metrics
-  tag: v1.7.2
+  tag: v1.8.0
   pullPolicy: IfNotPresent
 
 replicas: 1

```

## 2.3.1

**Release date:** 2019-08-20

![AppVersion: 1.7.2](https://img.shields.io/static/v1?label=AppVersion&message=1.7.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] Add honorlabel parameter (#16401)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 2c32ccf8..d5631b00 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -38,6 +38,7 @@ prometheus:
     enabled: false
     additionalLabels: {}
     namespace: ""
+    honorLabels: false
 
 ## Specify if a Pod Security Policy for kube-state-metrics must be created
 ## Ref: https://kubernetes.io/docs/concepts/policy/pod-security-policy/

```

## 2.3.0

**Release date:** 2019-08-13

![AppVersion: 1.7.2](https://img.shields.io/static/v1?label=AppVersion&message=1.7.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] Update liveness- and readinessProbes (#16284)

### Default value changes

```diff
# No changes in this release
```

## 2.2.3

**Release date:** 2019-08-08

![AppVersion: 1.7.2](https://img.shields.io/static/v1?label=AppVersion&message=1.7.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] update rbac to use new ingress api group (#16188) (#16189)

### Default value changes

```diff
# No changes in this release
```

## 2.2.2

**Release date:** 2019-08-05

![AppVersion: 1.7.2](https://img.shields.io/static/v1?label=AppVersion&message=1.7.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] Version bump to 1.7.2 (#16093)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index a08b8f1f..2c32ccf8 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -2,7 +2,7 @@
 prometheusScrape: true
 image:
   repository: quay.io/coreos/kube-state-metrics
-  tag: v1.7.1
+  tag: v1.7.2
   pullPolicy: IfNotPresent
 
 replicas: 1

```

## 2.2.1

**Release date:** 2019-07-26

![AppVersion: 1.7.1](https://img.shields.io/static/v1?label=AppVersion&message=1.7.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] Add option for service annotations (#15864)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 6c099f6f..a08b8f1f 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -13,6 +13,7 @@ service:
   type: ClusterIP
   nodePort: 0
   loadBalancerIP: ""
+  annotations: {}
 
 customLabels: {}
 

```

## 2.1.1

**Release date:** 2019-07-25

![AppVersion: 1.7.1](https://img.shields.io/static/v1?label=AppVersion&message=1.7.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] add customLabels to service, deployment and pods (#15808)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 50da9b42..6c099f6f 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -14,6 +14,8 @@ service:
   nodePort: 0
   loadBalancerIP: ""
 
+customLabels: {}
+
 hostNetwork: false
 
 rbac:

```

## 2.1.0

**Release date:** 2019-07-24

![AppVersion: 1.7.1](https://img.shields.io/static/v1?label=AppVersion&message=1.7.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] update kube-state-metrics chart to use v1.7.1 (#15818)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 39686653..50da9b42 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -2,7 +2,7 @@
 prometheusScrape: true
 image:
   repository: quay.io/coreos/kube-state-metrics
-  tag: v1.6.0
+  tag: v1.7.1
   pullPolicy: IfNotPresent
 
 replicas: 1
@@ -100,6 +100,8 @@ collectors:
   secrets: true
   services: true
   statefulsets: true
+  storageclasses: true
+  verticalpodautoscalers: false
 
 # Namespace to be enabled for collecting resources. By default all namespaces are collected.
 # namespace: ""

```

## 2.0.0

**Release date:** 2019-07-05

![AppVersion: 1.6.0](https://img.shields.io/static/v1?label=AppVersion&message=1.6.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] Apply best practices for label naming (#15261)

### Default value changes

```diff
# No changes in this release
```

## 1.6.5

**Release date:** 2019-07-02

![AppVersion: 1.6.0](https://img.shields.io/static/v1?label=AppVersion&message=1.6.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] Rename ServiceMonitor template (#15214)

### Default value changes

```diff
# No changes in this release
```

## 1.6.4

**Release date:** 2019-06-05

![AppVersion: 1.6.0](https://img.shields.io/static/v1?label=AppVersion&message=1.6.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Add Kubernetes to keywords list (#14523)

### Default value changes

```diff
# No changes in this release
```

## 1.6.3

**Release date:** 2019-05-24

![AppVersion: 1.6.0](https://img.shields.io/static/v1?label=AppVersion&message=1.6.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] fix service monitor label nesting (#14128)

### Default value changes

```diff
# No changes in this release
```

## 1.6.2

**Release date:** 2019-05-22

![AppVersion: 1.6.0](https://img.shields.io/static/v1?label=AppVersion&message=1.6.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] Solve #13961 (#13995)

### Default value changes

```diff
# No changes in this release
```

## 1.6.1

**Release date:** 2019-05-15

![AppVersion: 1.6.0](https://img.shields.io/static/v1?label=AppVersion&message=1.6.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics]: Fix link to metric docu (#13647)

### Default value changes

```diff
# No changes in this release
```

## 1.6.0

**Release date:** 2019-05-15

![AppVersion: 1.6.0](https://img.shields.io/static/v1?label=AppVersion&message=1.6.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics]: Use the recommended apiGroup versions in kube-state-metrics (#13844)

### Default value changes

```diff
# No changes in this release
```

## 1.5.0

**Release date:** 2019-05-15

![AppVersion: 1.6.0](https://img.shields.io/static/v1?label=AppVersion&message=1.6.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] add pod affinity (#12819)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 05626e8a..39686653 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -61,6 +61,10 @@ securityContext:
 ## Ref: https://kubernetes.io/docs/user-guide/node-selection/
 nodeSelector: {}
 
+## Affinity settings for pod assignment
+## Ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
+affinity: {}
+
 ## Tolerations for pod assignment
 ## Ref: https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/
 tolerations: []

```

## 1.4.0

**Release date:** 2019-05-09

![AppVersion: 1.6.0](https://img.shields.io/static/v1?label=AppVersion&message=1.6.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] Allows deploying ServiceAccount without RBAC (#13495)

### Default value changes

```diff
# No changes in this release
```

## 1.3.0

**Release date:** 2019-05-07

![AppVersion: 1.6.0](https://img.shields.io/static/v1?label=AppVersion&message=1.6.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] upgrade to 1.6.0 (#13553)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index bf1e0ee6..05626e8a 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -1,8 +1,8 @@
 # Default values for kube-state-metrics.
 prometheusScrape: true
 image:
-  repository: k8s.gcr.io/kube-state-metrics
-  tag: v1.5.0
+  repository: quay.io/coreos/kube-state-metrics
+  tag: v1.6.0
   pullPolicy: IfNotPresent
 
 replicas: 1
@@ -74,12 +74,14 @@ podAnnotations: {}
 # Available collectors for kube-state-metrics. By default all available
 # collectors are enabled.
 collectors:
+  certificatesigningrequests: true
   configmaps: true
   cronjobs: true
   daemonsets: true
   deployments: true
   endpoints: true
   horizontalpodautoscalers: true
+  ingresses: true
   jobs: true
   limitranges: true
   namespaces: true

```

## 1.2.0

**Release date:** 2019-04-30

![AppVersion: 1.5.0](https://img.shields.io/static/v1?label=AppVersion&message=1.5.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics]: Add support for host networking (#13414)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index bc286193..bf1e0ee6 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -14,6 +14,8 @@ service:
   nodePort: 0
   loadBalancerIP: ""
 
+hostNetwork: false
+
 rbac:
   # If true, create & use RBAC resources
   create: true

```

## 1.1.0

**Release date:** 2019-04-23

![AppVersion: 1.5.0](https://img.shields.io/static/v1?label=AppVersion&message=1.5.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics ] Fix option to specify service account name (#13160)

### Default value changes

```diff
# No changes in this release
```

## 1.0.1

**Release date:** 2019-04-22

![AppVersion: 1.5.0](https://img.shields.io/static/v1?label=AppVersion&message=1.5.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Fixed invalid empty annotations in kube-state-metrics chart template (#12207)

### Default value changes

```diff
# No changes in this release
```

## 1.0.0

**Release date:** 2019-04-15

![AppVersion: 1.5.0](https://img.shields.io/static/v1?label=AppVersion&message=1.5.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] add optional prometheus service monitor (#12145)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 86aab9bf..bc286193 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -28,6 +28,12 @@ serviceAccount:
   # ref: https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/
   imagePullSecrets: []
 
+prometheus:
+  monitor:
+    enabled: false
+    additionalLabels: {}
+    namespace: ""
+
 ## Specify if a Pod Security Policy for kube-state-metrics must be created
 ## Ref: https://kubernetes.io/docs/concepts/policy/pod-security-policy/
 ##

```

## 0.16.0

**Release date:** 2019-03-25

![AppVersion: 1.5.0](https://img.shields.io/static/v1?label=AppVersion&message=1.5.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Adds replicas to kube-state-metrics deployment (#12267)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 611b0078..86aab9bf 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -5,6 +5,8 @@ image:
   tag: v1.5.0
   pullPolicy: IfNotPresent
 
+replicas: 1
+
 service:
   port: 8080
   # Default to clusterIP for backward compatibility

```

## 0.15.0

**Release date:** 2019-03-14

![AppVersion: 1.5.0](https://img.shields.io/static/v1?label=AppVersion&message=1.5.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] Add imagePullSecrets in serviceAccount (#11405)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 48ae8ac4..611b0078 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -4,17 +4,27 @@ image:
   repository: k8s.gcr.io/kube-state-metrics
   tag: v1.5.0
   pullPolicy: IfNotPresent
+
 service:
   port: 8080
   # Default to clusterIP for backward compatibility
   type: ClusterIP
   nodePort: 0
   loadBalancerIP: ""
+
 rbac:
   # If true, create & use RBAC resources
   create: true
-  # Ignored if rbac.create is true
-  serviceAccountName: default
+
+serviceAccount:
+  # Specifies whether a ServiceAccount should be created, require rbac true
+  create: true
+  # The name of the ServiceAccount to use.
+  # If not set and create is true, a name is generated using the fullname template
+  name:
+  # Reference to one or more secrets to be used when pulling images
+  # ref: https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/
+  imagePullSecrets: []
 
 ## Specify if a Pod Security Policy for kube-state-metrics must be created
 ## Ref: https://kubernetes.io/docs/concepts/policy/pod-security-policy/

```

## 0.14.1

**Release date:** 2019-02-15

![AppVersion: 1.5.0](https://img.shields.io/static/v1?label=AppVersion&message=1.5.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] update kube-state-metrics README.md (#11440)

### Default value changes

```diff
# No changes in this release
```

## 0.14.0

**Release date:** 2019-02-14

![AppVersion: 1.5.0](https://img.shields.io/static/v1?label=AppVersion&message=1.5.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* update kube-state-metrics chart to support kube-state-metrics 1.5.0 (#11316)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 7d644850..48ae8ac4 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -1,8 +1,8 @@
 # Default values for kube-state-metrics.
 prometheusScrape: true
 image:
-  repository: quay.io/coreos/kube-state-metrics
-  tag: v1.4.0
+  repository: k8s.gcr.io/kube-state-metrics
+  tag: v1.5.0
   pullPolicy: IfNotPresent
 service:
   port: 8080
@@ -66,6 +66,7 @@ collectors:
   nodes: true
   persistentvolumeclaims: true
   persistentvolumes: true
+  poddisruptionbudgets: true
   pods: true
   replicasets: true
   replicationcontrollers: true

```

## 0.13.1

**Release date:** 2019-02-03

![AppVersion: 1.4.0](https://img.shields.io/static/v1?label=AppVersion&message=1.4.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Use rbac.authorization.k8s.io/v1 apiVersion in RBAC resources (#11094)

### Default value changes

```diff
# No changes in this release
```

## 0.13.0

**Release date:** 2019-01-15

![AppVersion: 1.4.0](https://img.shields.io/static/v1?label=AppVersion&message=1.4.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Add PodSecurityPolicy support (#10480)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 626b17d9..7d644850 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -16,6 +16,22 @@ rbac:
   # Ignored if rbac.create is true
   serviceAccountName: default
 
+## Specify if a Pod Security Policy for kube-state-metrics must be created
+## Ref: https://kubernetes.io/docs/concepts/policy/pod-security-policy/
+##
+podSecurityPolicy:
+  enabled: false
+  annotations: {}
+    ## Specify pod annotations
+    ## Ref: https://kubernetes.io/docs/concepts/policy/pod-security-policy/#apparmor
+    ## Ref: https://kubernetes.io/docs/concepts/policy/pod-security-policy/#seccomp
+    ## Ref: https://kubernetes.io/docs/concepts/policy/pod-security-policy/#sysctl
+    ##
+    # seccomp.security.alpha.kubernetes.io/allowedProfileNames: '*'
+    # seccomp.security.alpha.kubernetes.io/defaultProfileName: 'docker/default'
+    # apparmor.security.beta.kubernetes.io/defaultProfileName: 'runtime/default'
+
+
 securityContext:
   enabled: true
   runAsUser: 65534

```

## 0.12.2

**Release date:** 2019-01-15

![AppVersion: 1.4.0](https://img.shields.io/static/v1?label=AppVersion&message=1.4.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* add configmap and secret collectors (#10665)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index bdc64fb6..626b17d9 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -38,6 +38,7 @@ podAnnotations: {}
 # Available collectors for kube-state-metrics. By default all available
 # collectors are enabled.
 collectors:
+  configmaps: true
   cronjobs: true
   daemonsets: true
   deployments: true
@@ -53,6 +54,7 @@ collectors:
   replicasets: true
   replicationcontrollers: true
   resourcequotas: true
+  secrets: true
   services: true
   statefulsets: true
 

```

## 0.12.1

**Release date:** 2018-11-28

![AppVersion: 1.4.0](https://img.shields.io/static/v1?label=AppVersion&message=1.4.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] Fix indentation of priorityClassName (#9553)

### Default value changes

```diff
# No changes in this release
```

## 0.12.0

**Release date:** 2018-11-26

![AppVersion: 1.4.0](https://img.shields.io/static/v1?label=AppVersion&message=1.4.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Add priorityClassName option to values (#9478)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index fbb376f7..bdc64fb6 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -32,6 +32,9 @@ tolerations: []
 # Annotations to be added to the pod
 podAnnotations: {}
 
+## Assign a PriorityClassName to pods if set
+# priorityClassName: ""
+
 # Available collectors for kube-state-metrics. By default all available
 # collectors are enabled.
 collectors:

```

## 0.11.0

**Release date:** 2018-11-05

![AppVersion: 1.4.0](https://img.shields.io/static/v1?label=AppVersion&message=1.4.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] Enable RBAC by default (needed for securitycontexts) (#9010)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index b7d830de..fbb376f7 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -12,7 +12,7 @@ service:
   loadBalancerIP: ""
 rbac:
   # If true, create & use RBAC resources
-  create: false
+  create: true
   # Ignored if rbac.create is true
   serviceAccountName: default
 

```

## 0.10.0

**Release date:** 2018-11-05

![AppVersion: 1.4.0](https://img.shields.io/static/v1?label=AppVersion&message=1.4.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Add securityContext (#9001)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index a9b60776..b7d830de 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -16,6 +16,11 @@ rbac:
   # Ignored if rbac.create is true
   serviceAccountName: default
 
+securityContext:
+  enabled: true
+  runAsUser: 65534
+  fsGroup: 65534
+
 ## Node labels for pod assignment
 ## Ref: https://kubernetes.io/docs/user-guide/node-selection/
 nodeSelector: {}

```

## 0.9.0

**Release date:** 2018-09-04

![AppVersion: 1.4.0](https://img.shields.io/static/v1?label=AppVersion&message=1.4.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* upgrade kube-state-metrics to v1.4.0 (#7515)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 9cfc6ef4..a9b60776 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -2,7 +2,7 @@
 prometheusScrape: true
 image:
   repository: quay.io/coreos/kube-state-metrics
-  tag: v1.3.1
+  tag: v1.4.0
   pullPolicy: IfNotPresent
 service:
   port: 8080

```

## 0.8.2

**Release date:** 2018-09-04

![AppVersion: 1.3.1](https://img.shields.io/static/v1?label=AppVersion&message=1.3.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-state-metrics] review readme (#6810)

### Default value changes

```diff
# No changes in this release
```

## 0.8.1

**Release date:** 2018-07-15

![AppVersion: 1.3.1](https://img.shields.io/static/v1?label=AppVersion&message=1.3.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] Use recommended repo (#6531)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 969bf94c..9cfc6ef4 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -1,7 +1,7 @@
 # Default values for kube-state-metrics.
 prometheusScrape: true
 image:
-  repository: k8s.gcr.io/kube-state-metrics
+  repository: quay.io/coreos/kube-state-metrics
   tag: v1.3.1
   pullPolicy: IfNotPresent
 service:

```

## 0.8.0

**Release date:** 2018-06-17

![AppVersion: 1.3.1](https://img.shields.io/static/v1?label=AppVersion&message=1.3.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* modify helpers to apply popular rule to decide fullname (#6126)

### Default value changes

```diff
# No changes in this release
```

## 0.7.2

**Release date:** 2018-06-17

![AppVersion: 1.3.1](https://img.shields.io/static/v1?label=AppVersion&message=1.3.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* kube-state-metrics: Update tag to 1.3.1 (#5973)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 09fc263c..969bf94c 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -2,7 +2,7 @@
 prometheusScrape: true
 image:
   repository: k8s.gcr.io/kube-state-metrics
-  tag: v1.2.0
+  tag: v1.3.1
   pullPolicy: IfNotPresent
 service:
   port: 8080

```

## 0.7.1

**Release date:** 2018-03-29

![AppVersion: 1.2.0](https://img.shields.io/static/v1?label=AppVersion&message=1.2.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] Update ClusterRole to match new collectors (#4467)

### Default value changes

```diff
# No changes in this release
```

## 0.7.0

**Release date:** 2018-03-27

![AppVersion: 1.2.0](https://img.shields.io/static/v1?label=AppVersion&message=1.2.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics]: support toleration configuration (#4412)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index eea7a4ad..09fc263c 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -20,6 +20,10 @@ rbac:
 ## Ref: https://kubernetes.io/docs/user-guide/node-selection/
 nodeSelector: {}
 
+## Tolerations for pod assignment
+## Ref: https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/
+tolerations: []
+
 # Annotations to be added to the pod
 podAnnotations: {}
 

```

## 0.6.0

**Release date:** 2018-03-23

![AppVersion: 1.2.0](https://img.shields.io/static/v1?label=AppVersion&message=1.2.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Upgrade kube-state-metrics to 1.2.0, add new collectors (#4146)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 45b96f6e..eea7a4ad 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -2,7 +2,7 @@
 prometheusScrape: true
 image:
   repository: k8s.gcr.io/kube-state-metrics
-  tag: v1.1.0
+  tag: v1.2.0
   pullPolicy: IfNotPresent
 service:
   port: 8080
@@ -26,19 +26,23 @@ podAnnotations: {}
 # Available collectors for kube-state-metrics. By default all available
 # collectors are enabled.
 collectors:
+  cronjobs: true
   daemonsets: true
   deployments: true
+  endpoints: true
+  horizontalpodautoscalers: true
+  jobs: true
   limitranges: true
+  namespaces: true
   nodes: true
+  persistentvolumeclaims: true
+  persistentvolumes: true
   pods: true
   replicasets: true
   replicationcontrollers: true
   resourcequotas: true
   services: true
-  jobs: true
-  cronjobs: true
   statefulsets: true
-  persistentvolumeclaims: true
 
 # Namespace to be enabled for collecting resources. By default all namespaces are collected.
 # namespace: ""

```

## 0.5.3

**Release date:** 2018-02-09

![AppVersion: 1.1.0](https://img.shields.io/static/v1?label=AppVersion&message=1.1.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] allow for namespace flag to be configured (#3595)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 47626491..45b96f6e 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -39,3 +39,6 @@ collectors:
   cronjobs: true
   statefulsets: true
   persistentvolumeclaims: true
+
+# Namespace to be enabled for collecting resources. By default all namespaces are collected.
+# namespace: ""

```

## 0.5.2

**Release date:** 2018-01-13

![AppVersion: 1.1.0](https://img.shields.io/static/v1?label=AppVersion&message=1.1.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] allow for nodeSelector to be configured (#3231)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index e07ad32f..47626491 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -16,6 +16,10 @@ rbac:
   # Ignored if rbac.create is true
   serviceAccountName: default
 
+## Node labels for pod assignment
+## Ref: https://kubernetes.io/docs/user-guide/node-selection/
+nodeSelector: {}
+
 # Annotations to be added to the pod
 podAnnotations: {}
 

```

## 0.5.1

**Release date:** 2017-12-28

![AppVersion: 1.1.0](https://img.shields.io/static/v1?label=AppVersion&message=1.1.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] Convert registry to k8s.gcr.io (#3163)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 980b0da3..e07ad32f 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -1,7 +1,7 @@
 # Default values for kube-state-metrics.
 prometheusScrape: true
 image:
-  repository: gcr.io/google_containers/kube-state-metrics
+  repository: k8s.gcr.io/kube-state-metrics
   tag: v1.1.0
   pullPolicy: IfNotPresent
 service:

```

## 0.5.0

**Release date:** 2017-12-07

![AppVersion: 1.1.0](https://img.shields.io/static/v1?label=AppVersion&message=1.1.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] Add support for pod annotations (#2849)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 5ce04254..980b0da3 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -16,6 +16,9 @@ rbac:
   # Ignored if rbac.create is true
   serviceAccountName: default
 
+# Annotations to be added to the pod
+podAnnotations: {}
+
 # Available collectors for kube-state-metrics. By default all available
 # collectors are enabled.
 collectors:

```

## 0.4.0

**Release date:** 2017-12-05

![AppVersion: 1.1.0](https://img.shields.io/static/v1?label=AppVersion&message=1.1.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] Update to v1.1.0 (#2822)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index e4211eaa..5ce04254 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -2,7 +2,7 @@
 prometheusScrape: true
 image:
   repository: gcr.io/google_containers/kube-state-metrics
-  tag: v1.0.1
+  tag: v1.1.0
   pullPolicy: IfNotPresent
 service:
   port: 8080

```

## 0.3.2

**Release date:** 2017-11-25

![AppVersion: 1.0.1](https://img.shields.io/static/v1?label=AppVersion&message=1.0.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] add readiness probe (#2718)
* [stable/kube-state-metrics] allow configuring collectors (#2124)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index cb89236d..e4211eaa 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -15,3 +15,20 @@ rbac:
   create: false
   # Ignored if rbac.create is true
   serviceAccountName: default
+
+# Available collectors for kube-state-metrics. By default all available
+# collectors are enabled.
+collectors:
+  daemonsets: true
+  deployments: true
+  limitranges: true
+  nodes: true
+  pods: true
+  replicasets: true
+  replicationcontrollers: true
+  resourcequotas: true
+  services: true
+  jobs: true
+  cronjobs: true
+  statefulsets: true
+  persistentvolumeclaims: true

```

## 0.3.1

**Release date:** 2017-10-14

![AppVersion: 1.0.1](https://img.shields.io/static/v1?label=AppVersion&message=1.0.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] Add service type (#2002)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index a6ccf103..cb89236d 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -6,6 +6,10 @@ image:
   pullPolicy: IfNotPresent
 service:
   port: 8080
+  # Default to clusterIP for backward compatibility
+  type: ClusterIP
+  nodePort: 0
+  loadBalancerIP: ""
 rbac:
   # If true, create & use RBAC resources
   create: false

```

## 0.3.0

**Release date:** 2017-09-19

![AppVersion: 1.0.1](https://img.shields.io/static/v1?label=AppVersion&message=1.0.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] #1785 namespace defined templates with chart name (#2133)

### Default value changes

```diff
# No changes in this release
```

## 0.2.4

**Release date:** 2017-09-04

![AppVersion: 1.0.1](https://img.shields.io/static/v1?label=AppVersion&message=1.0.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] Enable chart upgrades (#1865)

### Default value changes

```diff
# No changes in this release
```

## 0.2.3

**Release date:** 2017-09-04

![AppVersion: 1.0.1](https://img.shields.io/static/v1?label=AppVersion&message=1.0.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/kube-state-metrics] Add RBAC support (#1870)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 33f27378..a6ccf103 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -6,3 +6,8 @@ image:
   pullPolicy: IfNotPresent
 service:
   port: 8080
+rbac:
+  # If true, create & use RBAC resources
+  create: false
+  # Ignored if rbac.create is true
+  serviceAccountName: default

```

## 0.2.2

**Release date:** 2017-09-02

![AppVersion: 1.0.1](https://img.shields.io/static/v1?label=AppVersion&message=1.0.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* name kubeStateMetrics port as 'metrics' (#1885)

### Default value changes

```diff
# No changes in this release
```

## 0.2.1

**Release date:** 2017-09-02

![AppVersion: 1.0.1](https://img.shields.io/static/v1?label=AppVersion&message=1.0.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* kube-state-metrics image has been updated (#1877)
* [stable/kube-state-metrics] Update documented image version (#1710)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 6554cccd..33f27378 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -2,7 +2,7 @@
 prometheusScrape: true
 image:
   repository: gcr.io/google_containers/kube-state-metrics
-  tag: v1.0.0
+  tag: v1.0.1
   pullPolicy: IfNotPresent
 service:
   port: 8080

```

## 0.2.0

**Release date:** 2017-08-10

![AppVersion: 1.0.0](https://img.shields.io/static/v1?label=AppVersion&message=1.0.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Upgrade kube-state-metrics version (#1700)

### Default value changes

```diff
diff --git a/charts/kube-state-metrics/values.yaml b/charts/kube-state-metrics/values.yaml
index 0379f4fb..6554cccd 100644
--- a/charts/kube-state-metrics/values.yaml
+++ b/charts/kube-state-metrics/values.yaml
@@ -2,7 +2,7 @@
 prometheusScrape: true
 image:
   repository: gcr.io/google_containers/kube-state-metrics
-  tag: v0.4.1
+  tag: v1.0.0
   pullPolicy: IfNotPresent
 service:
   port: 8080

```

## 0.1.0

**Release date:** 2017-05-30

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Added kube-state-metrics (#662)

### Default value changes

```diff
# Default values for kube-state-metrics.
prometheusScrape: true
image:
  repository: gcr.io/google_containers/kube-state-metrics
  tag: v0.4.1
  pullPolicy: IfNotPresent
service:
  port: 8080

```

---
Autogenerated from Helm Chart and git history using [helm-changelog](https://github.com/mogensen/helm-changelog)
