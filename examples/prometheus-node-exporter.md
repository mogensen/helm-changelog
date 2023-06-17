# Change Log

## 4.17.5

**Release date:** 2023-06-06

![AppVersion: 1.5.0](https://img.shields.io/static/v1?label=AppVersion&message=1.5.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] remove prometheus scrape annotation (#3123)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index ca9b77ed..45acd897 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -71,8 +71,7 @@ service:
   nodePort:
   portName: metrics
   listenOnAllInterfaces: true
-  annotations:
-    prometheus.io/scrape: "true"
+  annotations: {}
 
 # Set a NetworkPolicy with:
 # ingress only on service.port

```

## 4.17.4

**Release date:** 2023-06-01

![AppVersion: 1.5.0](https://img.shields.io/static/v1?label=AppVersion&message=1.5.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] Patch nodeSelector on node-exporter daemonset (#3190)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 87deb273..ca9b77ed 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -341,9 +341,9 @@ dnsConfig: {}
 
 ## Assign a nodeSelector if operating a hybrid cluster
 ##
-nodeSelector: {}
-#   beta.kubernetes.io/arch: amd64
-#   beta.kubernetes.io/os: linux
+nodeSelector:
+  kubernetes.io/os: linux
+  #  kubernetes.io/arch: amd64
 
 tolerations:
   - effect: NoSchedule

```

## 4.17.3

**Release date:** 2023-05-23

![AppVersion: 1.5.0](https://img.shields.io/static/v1?label=AppVersion&message=1.5.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] added fullnameOverride to values.yaml (#3419)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 643b9c85..87deb273 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -11,6 +11,8 @@ image:
 
 imagePullSecrets: []
 # - name: "image-pull-secret"
+nameOverride: ""
+fullnameOverride: ""
 
 global:
   # To help compatibility with other charts which use global.imagePullSecrets.

```

## 4.17.2

**Release date:** 2023-05-07

![AppVersion: 1.5.0](https://img.shields.io/static/v1?label=AppVersion&message=1.5.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* fix node-exporter listen address with kubeRBACProxy enabled (#3336)

### Default value changes

```diff
# No changes in this release
```

## 4.17.1

**Release date:** 2023-05-03

![AppVersion: 1.5.0](https://img.shields.io/static/v1?label=AppVersion&message=1.5.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] Add a maintainer (#3314)

### Default value changes

```diff
# No changes in this release
```

## 4.17.0

**Release date:** 2023-04-21

![AppVersion: 1.5.0](https://img.shields.io/static/v1?label=AppVersion&message=1.5.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] set readOnlyRootFilesystem=true (#3272)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index ff1e53f8..643b9c85 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -264,7 +264,8 @@ securityContext:
   runAsNonRoot: true
   runAsUser: 65534
 
-containerSecurityContext: {}
+containerSecurityContext:
+  readOnlyRootFilesystem: true
   # capabilities:
   #   add:
   #   - SYS_TIME

```

## 4.16.0

**Release date:** 2023-04-12

![AppVersion: 1.5.0](https://img.shields.io/static/v1?label=AppVersion&message=1.5.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* add an optional `NetworkPolicy` (#3212)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 1236b9e1..ff1e53f8 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -72,6 +72,12 @@ service:
   annotations:
     prometheus.io/scrape: "true"
 
+# Set a NetworkPolicy with:
+# ingress only on service.port
+# no egress permitted
+networkPolicy:
+  enabled: false
+
 # Additional environment variables that will be passed to the daemonset
 env: {}
 ##  env:

```

## 4.15.1

**Release date:** 2023-04-11

![AppVersion: 1.5.0](https://img.shields.io/static/v1?label=AppVersion&message=1.5.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter]  Use digest instead of sha (#3130)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index b498a208..1236b9e1 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -7,7 +7,7 @@ image:
   # Overrides the image tag whose default is {{ printf "v%s" .Chart.AppVersion }}
   tag: ""
   pullPolicy: IfNotPresent
-  sha: ""
+  digest: ""
 
 imagePullSecrets: []
 # - name: "image-pull-secret"

```

## 4.15.0

**Release date:** 2023-04-06

![AppVersion: 1.5.0](https://img.shields.io/static/v1?label=AppVersion&message=1.5.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] add attachMetadata to ServiceMonitor (#3176)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index eb7dc6b1..b498a208 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -102,6 +102,11 @@ prometheus:
     ##
     selectorOverride: {}
 
+    ## Attach node metadata to discovered targets. Requires Prometheus v2.35.0 and above.
+    ##
+    attachMetadata:
+      node: false
+
     relabelings: []
     metricRelabelings: []
     interval: ""

```

## 4.14.0

**Release date:** 2023-02-22

![AppVersion: 1.5.0](https://img.shields.io/static/v1?label=AppVersion&message=1.5.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] Allow parent charts to override registry hostname (#3044)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index ae43c00d..eb7dc6b1 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -2,7 +2,8 @@
 # This is a YAML-formatted file.
 # Declare variables to be passed into your templates.
 image:
-  repository: quay.io/prometheus/node-exporter
+  registry: quay.io
+  repository: prometheus/node-exporter
   # Overrides the image tag whose default is {{ printf "v%s" .Chart.AppVersion }}
   tag: ""
   pullPolicy: IfNotPresent
@@ -24,13 +25,17 @@ global:
   #   - pullSecret1
   #   - pullSecret2
   imagePullSecrets: []
+  #
+  # Allow parent charts to override registry hostname
+  imageRegistry: ""
 
 # Configure kube-rbac-proxy. When enabled, creates a kube-rbac-proxy to protect the node-exporter http endpoint.
 # The requests are served through the same service but requests are HTTPS.
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

## 4.13.1

**Release date:** 2023-02-15

![AppVersion: 1.5.0](https://img.shields.io/static/v1?label=AppVersion&message=1.5.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] Update VPA containerPolicies container name to match DaemonSet (#2999)

### Default value changes

```diff
# No changes in this release
```

## 4.13.0

**Release date:** 2023-01-23

![AppVersion: 1.5.0](https://img.shields.io/static/v1?label=AppVersion&message=1.5.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] Add udev path arg (#2941)

### Default value changes

```diff
# No changes in this release
```

## 4.12.0

**Release date:** 2023-01-11

![AppVersion: 1.5.0](https://img.shields.io/static/v1?label=AppVersion&message=1.5.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] Allow to enable kube-rbac-proxy for node-exporter (#2892)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index db55144f..ae43c00d 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -25,6 +25,38 @@ global:
   #   - pullSecret2
   imagePullSecrets: []
 
+# Configure kube-rbac-proxy. When enabled, creates a kube-rbac-proxy to protect the node-exporter http endpoint.
+# The requests are served through the same service but requests are HTTPS.
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
+  # requests:
+  #  cpu: 10m
+  #  memory: 32Mi
+
 service:
   type: ClusterIP
   port: 9100

```

## 4.11.0

**Release date:** 2023-01-10

![AppVersion: 1.5.0](https://img.shields.io/static/v1?label=AppVersion&message=1.5.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] Add a pod monitor (#2861) (#2866)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 4df845bb..db55144f 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -92,6 +92,96 @@ prometheus:
     ##
     labelValueLengthLimit: 0
 
+  # PodMonitor defines monitoring for a set of pods.
+  # ref. https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#monitoring.coreos.com/v1.PodMonitor
+  # Using a PodMonitor may be preferred in some environments where there is very large number
+  # of Node Exporter endpoints (1000+) behind a single service.
+  # The PodMonitor is disabled by default. When switching from ServiceMonitor to PodMonitor,
+  # the time series resulting from the configuration through PodMonitor may have different labels.
+  # For instance, there will not be the service label any longer which might
+  # affect PromQL queries selecting that label.
+  podMonitor:
+    enabled: false
+    # Namespace in which to deploy the pod monitor. Defaults to the release namespace.
+    namespace: ""
+    # Additional labels, e.g. setting a label for pod monitor selector as set in prometheus
+    additionalLabels: {}
+    #  release: kube-prometheus-stack
+    # PodTargetLabels transfers labels of the Kubernetes Pod onto the target.
+    podTargetLabels: []
+    # apiVersion defaults to monitoring.coreos.com/v1.
+    apiVersion: ""
+    # Override pod selector to select pod objects.
+    selectorOverride: {}
+    # Attach node metadata to discovered targets. Requires Prometheus v2.35.0 and above.
+    attachMetadata:
+      node: false
+    # The label to use to retrieve the job name from. Defaults to label app.kubernetes.io/name.
+    jobLabel: ""
+
+    # Scheme/protocol to use for scraping.
+    scheme: "http"
+    # Path to scrape metrics at.
+    path: "/metrics"
+
+    # BasicAuth allow an endpoint to authenticate over basic authentication.
+    # More info: https://prometheus.io/docs/operating/configuration/#endpoint
+    basicAuth: {}
+    # Secret to mount to read bearer token for scraping targets.
+    # The secret needs to be in the same namespace as the pod monitor and accessible by the Prometheus Operator.
+    # https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.24/#secretkeyselector-v1-core
+    bearerTokenSecret: {}
+    # TLS configuration to use when scraping the endpoint.
+    tlsConfig: {}
+    # Authorization section for this endpoint.
+    # https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#monitoring.coreos.com/v1.SafeAuthorization
+    authorization: {}
+    # OAuth2 for the URL. Only valid in Prometheus versions 2.27.0 and newer.
+    # https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#monitoring.coreos.com/v1.OAuth2
+    oauth2: {}
+
+    # ProxyURL eg http://proxyserver:2195. Directs scrapes through proxy to this endpoint.
+    proxyUrl: ""
+    # Interval at which endpoints should be scraped. If not specified Prometheus’ global scrape interval is used.
+    interval: ""
+    # Timeout after which the scrape is ended. If not specified, the Prometheus global scrape interval is used.
+    scrapeTimeout: ""
+    # HonorTimestamps controls whether Prometheus respects the timestamps present in scraped data.
+    honorTimestamps: true
+    # HonorLabels chooses the metric’s labels on collisions with target labels.
+    honorLabels: true
+    # Whether to enable HTTP2. Default false.
+    enableHttp2: ""
+    # Drop pods that are not running. (Failed, Succeeded).
+    # Enabled by default. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#pod-phase
+    filterRunning: ""
+    # FollowRedirects configures whether scrape requests follow HTTP 3xx redirects. Default false.
+    followRedirects: ""
+    # Optional HTTP URL parameters
+    params: {}
+
+    # RelabelConfigs to apply to samples before scraping. Prometheus Operator automatically adds
+    # relabelings for a few standard Kubernetes fields. The original scrape job’s name
+    # is available via the __tmp_prometheus_job_name label.
+    # More info: https://prometheus.io/docs/prometheus/latest/configuration/configuration/#relabel_config
+    relabelings: []
+    # MetricRelabelConfigs to apply to samples before ingestion.
+    metricRelabelings: []
+
+    # SampleLimit defines per-scrape limit on number of scraped samples that will be accepted.
+    sampleLimit: 0
+    # TargetLimit defines a limit on the number of scraped targets that will be accepted.
+    targetLimit: 0
+    # Per-scrape limit on number of labels that will be accepted for a sample.
+    # Only valid in Prometheus versions 2.27.0 and newer.
+    labelLimit: 0
+    # Per-scrape limit on length of labels name that will be accepted for a sample.
+    # Only valid in Prometheus versions 2.27.0 and newer.
+    labelNameLengthLimit: 0
+    # Per-scrape limit on length of labels value that will be accepted for a sample.
+    # Only valid in Prometheus versions 2.27.0 and newer.
+    labelValueLengthLimit: 0
+
 ## Customize the updateStrategy if set
 updateStrategy:
   type: RollingUpdate

```

## 4.10.0

**Release date:** 2023-01-10

![AppVersion: 1.5.0](https://img.shields.io/static/v1?label=AppVersion&message=1.5.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] Adds podTargetLabels configuration to service monitor (#2785)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 3036bed7..4df845bb 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -48,6 +48,10 @@ prometheus:
 
     jobLabel: ""
 
+    # List of pod labels to add to node exporter metrics
+    # https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#servicemonitor
+    podTargetLabels: []
+
     scheme: http
     basicAuth: {}
     bearerTokenFile:

```

## 4.9.0

**Release date:** 2023-01-07

![AppVersion: 1.5.0](https://img.shields.io/static/v1?label=AppVersion&message=1.5.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] Add support for global imagePullSecrets values (#2772)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index f491bdfa..3036bed7 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -11,6 +11,20 @@ image:
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
 service:
   type: ClusterIP
   port: 9100

```

## 4.8.1

**Release date:** 2022-12-21

![AppVersion: 1.5.0](https://img.shields.io/static/v1?label=AppVersion&message=1.5.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] Fix variable scope for sidecar volume mounts (#2825)

### Default value changes

```diff
# No changes in this release
```

## 4.8.0

**Release date:** 2022-12-01

![AppVersion: 1.5.0](https://img.shields.io/static/v1?label=AppVersion&message=1.5.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] updates the Prometheus Node Exporter image to v1.5.0 (#2766)

### Default value changes

```diff
# No changes in this release
```

## 4.7.1

**Release date:** 2022-11-28

![AppVersion: 1.4.0](https://img.shields.io/static/v1?label=AppVersion&message=1.4.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] Correct indentation of environment variables (#2744)

### Default value changes

```diff
# No changes in this release
```

## 4.7.0

**Release date:** 2022-11-22

![AppVersion: 1.4.0](https://img.shields.io/static/v1?label=AppVersion&message=1.4.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] Update image to v1.4.0 (#2696)

### Default value changes

```diff
# No changes in this release
```

## 4.6.0

**Release date:** 2022-11-18

![AppVersion: 1.3.1](https://img.shields.io/static/v1?label=AppVersion&message=1.3.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] Add daemonset annotations (#2632) (#2701)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index c4f2b571..f491bdfa 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -162,6 +162,9 @@ podAnnotations:
 # Extra labels to be added to node exporter pods
 podLabels: {}
 
+# Annotations to be added to node exporter daemonset
+daemonsetAnnotations: {}
+
 ## set to true to add the release label so scraping of the servicemonitor with kube-prometheus-stack works out of the box
 releaseLabel: false
 

```

## 4.5.2

**Release date:** 2022-11-15

![AppVersion: 1.3.1](https://img.shields.io/static/v1?label=AppVersion&message=1.3.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] Line break issue in labels (#2682)

### Default value changes

```diff
# No changes in this release
```

## 4.5.1

**Release date:** 2022-11-11

![AppVersion: 1.3.1](https://img.shields.io/static/v1?label=AppVersion&message=1.3.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] Refactor (#2650)

### Default value changes

```diff
# No changes in this release
```

## 4.5.0

**Release date:** 2022-11-01

![AppVersion: 1.3.1](https://img.shields.io/static/v1?label=AppVersion&message=1.3.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] add servicemonitor/podmonitor scrape limits (#2609)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 3da0a2f4..c4f2b571 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -54,6 +54,26 @@ prometheus:
     ## prometheus.monitor.apiVersion ApiVersion for the serviceMonitor Resource(defaults to "monitoring.coreos.com/v1")
     apiVersion: ""
 
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
+
 ## Customize the updateStrategy if set
 updateStrategy:
   type: RollingUpdate

```

## 4.4.2

**Release date:** 2022-10-28

![AppVersion: 1.3.1](https://img.shields.io/static/v1?label=AppVersion&message=1.3.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] ServiceMonitor selectorOverride fix (#2618)

### Default value changes

```diff
# No changes in this release
```

## 4.4.1

**Release date:** 2022-10-16

![AppVersion: 1.3.1](https://img.shields.io/static/v1?label=AppVersion&message=1.3.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] Fix PSP deprecation after k8s 1.25+ (#2574)

### Default value changes

```diff
# No changes in this release
```

## 4.4.0

**Release date:** 2022-10-12

![AppVersion: 1.3.1](https://img.shields.io/static/v1?label=AppVersion&message=1.3.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] Add Vertical Pod Autoscaler support (#2553)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index a00ae4d3..3da0a2f4 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -246,3 +246,23 @@ readinessProbe:
   periodSeconds: 10
   successThreshold: 1
   timeoutSeconds: 1
+
+# Enable vertical pod autoscaler support for prometheus-node-exporter
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

```

## 4.3.1

**Release date:** 2022-10-12

![AppVersion: 1.3.1](https://img.shields.io/static/v1?label=AppVersion&message=1.3.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] use recent version of PSP API in K8S 1.25 (#2520)

### Default value changes

```diff
# No changes in this release
```

## 4.3.0

**Release date:** 2022-09-16

![AppVersion: 1.3.1](https://img.shields.io/static/v1?label=AppVersion&message=1.3.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] custom servicemonitor apiversion (#2407)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index ac22d28e..a00ae4d3 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -51,6 +51,8 @@ prometheus:
     metricRelabelings: []
     interval: ""
     scrapeTimeout: 10s
+    ## prometheus.monitor.apiVersion ApiVersion for the serviceMonitor Resource(defaults to "monitoring.coreos.com/v1")
+    apiVersion: ""
 
 ## Customize the updateStrategy if set
 updateStrategy:

```

## 4.2.0

**Release date:** 2022-09-13

![AppVersion: 1.3.1](https://img.shields.io/static/v1?label=AppVersion&message=1.3.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* add releaseLabel (#2453)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index a584f528..ac22d28e 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -140,6 +140,9 @@ podAnnotations:
 # Extra labels to be added to node exporter pods
 podLabels: {}
 
+## set to true to add the release label so scraping of the servicemonitor with kube-prometheus-stack works out of the box
+releaseLabel: false
+
 # Custom DNS configuration to be added to prometheus-node-exporter pods
 dnsConfig: {}
 # nameservers:

```

## 4.1.0

**Release date:** 2022-09-12

![AppVersion: 1.3.1](https://img.shields.io/static/v1?label=AppVersion&message=1.3.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] Add sidecarHostVolumeMounts (#2441)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index dba5750a..a584f528 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -171,7 +171,7 @@ extraArgs: []
 #   - --collector.diskstats.ignored-devices=^(ram|loop|fd|(h|s|v)d[a-z]|nvme\\d+n\\d+p)\\d+$
 #   - --collector.textfile.directory=/run/prometheus
 
-## Additional mounts from the host
+## Additional mounts from the host to node-exporter container
 ##
 extraHostVolumeMounts: []
 #  - name: <mountName>
@@ -205,6 +205,15 @@ sidecarVolumeMount: []
 ##    mountPath: /run/prometheus
 ##    readOnly: false
 
+## Additional mounts from the host to sidecar containers
+##
+sidecarHostVolumeMounts: []
+#  - name: <mountName>
+#    hostPath: <hostPath>
+#    mountPath: <mountPath>
+#    readOnly: true|false
+#    mountPropagation: None|HostToContainer|Bidirectional
+
 ## Additional InitContainers to initialize the pod
 ##
 extraInitContainers: []

```

## 4.0.0

**Release date:** 2022-08-22

![AppVersion: 1.3.1](https://img.shields.io/static/v1?label=AppVersion&message=1.3.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] Update labels to match best practices (#2212)

### Default value changes

```diff
# No changes in this release
```

## 3.4.0

**Release date:** 2022-08-17

![AppVersion: 1.3.1](https://img.shields.io/static/v1?label=AppVersion&message=1.3.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* feat[prometheus-node-exporter]: add image sha support for image specified (#2377)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 80b399a0..dba5750a 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -6,6 +6,7 @@ image:
   # Overrides the image tag whose default is {{ printf "v%s" .Chart.AppVersion }}
   tag: ""
   pullPolicy: IfNotPresent
+  sha: ""
 
 imagePullSecrets: []
 # - name: "image-pull-secret"

```

## 3.3.1

**Release date:** 2022-07-24

![AppVersion: 1.3.1](https://img.shields.io/static/v1?label=AppVersion&message=1.3.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Removing me as maintainer (#2301)

### Default value changes

```diff
# No changes in this release
```

## 3.3.0

**Release date:** 2022-05-26

![AppVersion: 1.3.1](https://img.shields.io/static/v1?label=AppVersion&message=1.3.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] Make liveness and readiness probes customi… (#2053)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index ef994353..80b399a0 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -207,3 +207,27 @@ sidecarVolumeMount: []
 ## Additional InitContainers to initialize the pod
 ##
 extraInitContainers: []
+
+## Liveness probe
+##
+livenessProbe:
+  failureThreshold: 3
+  httpGet:
+    httpHeaders: []
+    scheme: http
+  initialDelaySeconds: 0
+  periodSeconds: 10
+  successThreshold: 1
+  timeoutSeconds: 1
+
+## Readiness probe
+##
+readinessProbe:
+  failureThreshold: 3
+  httpGet:
+    httpHeaders: []
+    scheme: http
+  initialDelaySeconds: 0
+  periodSeconds: 10
+  successThreshold: 1
+  timeoutSeconds: 1

```

## 3.2.0

**Release date:** 2022-05-17

![AppVersion: 1.3.1](https://img.shields.io/static/v1?label=AppVersion&message=1.3.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] Add basicAuth configuration to ServiceMonitor (#2061)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 2f303747..ef994353 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -34,6 +34,7 @@ prometheus:
     jobLabel: ""
 
     scheme: http
+    basicAuth: {}
     bearerTokenFile:
     tlsConfig: {}
 

```

## 3.1.1

**Release date:** 2022-04-13

![AppVersion: 1.3.1](https://img.shields.io/static/v1?label=AppVersion&message=1.3.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] enables adding imagePullSecrets in node exporter chart. (#1982)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 7e02e211..2f303747 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -7,6 +7,9 @@ image:
   tag: ""
   pullPolicy: IfNotPresent
 
+imagePullSecrets: []
+# - name: "image-pull-secret"
+
 service:
   type: ClusterIP
   port: 9100

```

## 3.1.0

**Release date:** 2022-03-28

![AppVersion: 1.3.1](https://img.shields.io/static/v1?label=AppVersion&message=1.3.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] be able to set environment variables (#1920)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 714dc9bd..7e02e211 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -17,6 +17,11 @@ service:
   annotations:
     prometheus.io/scrape: "true"
 
+# Additional environment variables that will be passed to the daemonset
+env: {}
+##  env:
+##    VARIABLE: value
+
 prometheus:
   monitor:
     enabled: false

```

## 3.0.2

**Release date:** 2022-03-24

![AppVersion: 1.3.1](https://img.shields.io/static/v1?label=AppVersion&message=1.3.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] Bracket $HOST_IP to enable IPv6 support (#1906)

### Default value changes

```diff
# No changes in this release
```

## 3.0.1

**Release date:** 2022-03-01

![AppVersion: 1.3.1](https://img.shields.io/static/v1?label=AppVersion&message=1.3.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] Add zanhsieh as one of maintainers (#1823)

### Default value changes

```diff
# No changes in this release
```

## 3.0.0

**Release date:** 2022-02-10

![AppVersion: 1.3.1](https://img.shields.io/static/v1?label=AppVersion&message=1.3.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] Parameterise host root FS mount propagation (#1583)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index adcf6b59..714dc9bd 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -100,9 +100,15 @@ hostNetwork: true
 # Share the host process ID namespace
 hostPID: true
 
-## If true, node-exporter pods mounts host / at /host/root
-##
-hostRootFsMount: true
+# Mount the node's root file system (/) at /host/root in the container
+hostRootFsMount:
+  enabled: true
+  # Defines how new mounts in existing mounts on the node or in the container
+  # are propagated to the container or node, respectively. Possible values are
+  # None, HostToContainer, and Bidirectional. If this field is omitted, then
+  # None is used. More information on:
+  # https://kubernetes.io/docs/concepts/storage/volumes/#mount-propagation
+  mountPropagation: HostToContainer
 
 ## Assign a group of affinity scheduling rules
 ##

```

## 2.5.0

**Release date:** 2022-01-13

![AppVersion: 1.3.1](https://img.shields.io/static/v1?label=AppVersion&message=1.3.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] reintroduce service monitor selectorOverride (#1699)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 62a74738..adcf6b59 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -33,6 +33,10 @@ prometheus:
     ##
     proxyUrl: ""
 
+    ## Override serviceMonitor selector
+    ##
+    selectorOverride: {}
+
     relabelings: []
     metricRelabelings: []
     interval: ""

```

## 2.4.1

**Release date:** 2021-12-22

![AppVersion: 1.3.1](https://img.shields.io/static/v1?label=AppVersion&message=1.3.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] Fix yaml comments prometheus node exporter (#1612)

### Default value changes

```diff
# No changes in this release
```

## 2.4.0

**Release date:** 2021-12-17

![AppVersion: 1.3.1](https://img.shields.io/static/v1?label=AppVersion&message=1.3.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] Update image to v1.3.1 & enable custom ServiceMonitor jobLabel (#1591)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 6f70b8ee..62a74738 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -3,7 +3,8 @@
 # Declare variables to be passed into your templates.
 image:
   repository: quay.io/prometheus/node-exporter
-  tag: v1.2.2
+  # Overrides the image tag whose default is {{ printf "v%s" .Chart.AppVersion }}
+  tag: ""
   pullPolicy: IfNotPresent
 
 service:
@@ -21,6 +22,9 @@ prometheus:
     enabled: false
     additionalLabels: {}
     namespace: ""
+
+    jobLabel: ""
+
     scheme: http
     bearerTokenFile:
     tlsConfig: {}

```

## 2.3.0

**Release date:** 2021-12-09

![AppVersion: 1.2.2](https://img.shields.io/static/v1?label=AppVersion&message=1.2.2&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [promethus-node-exporter] Improve ServiceMonitor (#1566)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 0d2c9aa8..6f70b8ee 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -30,6 +30,8 @@ prometheus:
     proxyUrl: ""
 
     relabelings: []
+    metricRelabelings: []
+    interval: ""
     scrapeTimeout: 10s
 
 ## Customize the updateStrategy if set

```

## 2.2.2

**Release date:** 2021-12-06

![AppVersion: 1.2.2](https://img.shields.io/static/v1?label=AppVersion&message=1.2.2&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] Remove me as maintainer (#1563)

### Default value changes

```diff
# No changes in this release
```

## 2.2.1

**Release date:** 2021-11-27

![AppVersion: 1.2.2](https://img.shields.io/static/v1?label=AppVersion&message=1.2.2&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter]Add availability to deploy ServiceMonitor onto a specific namespace.  (#1526)

### Default value changes

```diff
# No changes in this release
```

## 2.2.0

**Release date:** 2021-10-07

![AppVersion: 1.2.2](https://img.shields.io/static/v1?label=AppVersion&message=1.2.2&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] Allow to set service port name (#1406)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 203020ce..0d2c9aa8 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -11,6 +11,7 @@ service:
   port: 9100
   targetPort: 9100
   nodePort:
+  portName: metrics
   listenOnAllInterfaces: true
   annotations:
     prometheus.io/scrape: "true"

```

## 2.1.0

**Release date:** 2021-09-17

![AppVersion: 1.2.2](https://img.shields.io/static/v1?label=AppVersion&message=1.2.2&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] Allow for specifying automountServiceAccountToken (#1350)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 27b87831..203020ce 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -57,6 +57,7 @@ serviceAccount:
   name:
   annotations: {}
   imagePullSecrets: []
+  automountServiceAccountToken: false
 
 securityContext:
   fsGroup: 65534

```

## 2.0.4

**Release date:** 2021-08-19

![AppVersion: 1.2.2](https://img.shields.io/static/v1?label=AppVersion&message=1.2.2&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] Version bump to 1.2.2 (#1258)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index ac6416d6..27b87831 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -3,7 +3,7 @@
 # Declare variables to be passed into your templates.
 image:
   repository: quay.io/prometheus/node-exporter
-  tag: v1.2.1
+  tag: v1.2.2
   pullPolicy: IfNotPresent
 
 service:

```

## 2.0.3

**Release date:** 2021-08-06

![AppVersion: 1.2.1](https://img.shields.io/static/v1?label=AppVersion&message=1.2.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] Version bump to 1.2.1 (#1230)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 8b5f1587..ac6416d6 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -3,7 +3,7 @@
 # Declare variables to be passed into your templates.
 image:
   repository: quay.io/prometheus/node-exporter
-  tag: v1.2.0
+  tag: v1.2.1
   pullPolicy: IfNotPresent
 
 service:

```

## 2.0.2

**Release date:** 2021-07-26

![AppVersion: 1.2.0](https://img.shields.io/static/v1?label=AppVersion&message=1.2.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] Add configurable host PID (#1202)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 8c88bb91..8b5f1587 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -85,6 +85,9 @@ endpoints: []
 # Expose the service to the host network
 hostNetwork: true
 
+# Share the host process ID namespace
+hostPID: true
+
 ## If true, node-exporter pods mounts host / at /host/root
 ##
 hostRootFsMount: true

```

## 2.0.1

**Release date:** 2021-07-21

![AppVersion: 1.2.0](https://img.shields.io/static/v1?label=AppVersion&message=1.2.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* node-exporter: do not automount service account token (#1167)

### Default value changes

```diff
# No changes in this release
```

## 2.0.0

**Release date:** 2021-07-19

![AppVersion: 1.2.0](https://img.shields.io/static/v1?label=AppVersion&message=1.2.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] Version bump to 1.2.0; drop helm2 support (#1176)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index f4e927fa..8c88bb91 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -3,7 +3,7 @@
 # Declare variables to be passed into your templates.
 image:
   repository: quay.io/prometheus/node-exporter
-  tag: v1.1.2
+  tag: v1.2.0
   pullPolicy: IfNotPresent
 
 service:

```

## 1.18.2

**Release date:** 2021-06-19

![AppVersion: 1.1.2](https://img.shields.io/static/v1?label=AppVersion&message=1.1.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] Add the ability to define a proxyUrl via values (#1091)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 400ce113..f4e927fa 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -24,6 +24,10 @@ prometheus:
     bearerTokenFile:
     tlsConfig: {}
 
+    ## proxyUrl: URL of a proxy that should be used for scraping.
+    ##
+    proxyUrl: ""
+
     relabelings: []
     scrapeTimeout: 10s
 

```

## 1.18.1

**Release date:** 2021-06-03

![AppVersion: 1.1.2](https://img.shields.io/static/v1?label=AppVersion&message=1.1.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Fix PSP annotation identation (#1030)

### Default value changes

```diff
# No changes in this release
```

## 1.18.0

**Release date:** 2021-04-14

![AppVersion: 1.1.2](https://img.shields.io/static/v1?label=AppVersion&message=1.1.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] - feat: added psp annotations (#848)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 5cb6981a..400ce113 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -72,6 +72,7 @@ rbac:
   ## If true, create & use Pod Security Policy resources
   ## https://kubernetes.io/docs/concepts/policy/pod-security-policy/
   pspEnabled: true
+  pspAnnotations: {}
 
 # for deployments that have node_exporter deployed outside of the cluster, list
 # their addresses here

```

## 1.17.0

**Release date:** 2021-04-14

![AppVersion: 1.1.2](https://img.shields.io/static/v1?label=AppVersion&message=1.1.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] 1.16.0 Allow to configure init containers (#748)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 93f8ef80..5cb6981a 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -168,3 +168,7 @@ sidecarVolumeMount: []
 ##  - name: collector-textfiles
 ##    mountPath: /run/prometheus
 ##    readOnly: false
+
+## Additional InitContainers to initialize the pod
+##
+extraInitContainers: []

```

## 1.16.2

**Release date:** 2021-03-18

![AppVersion: 1.1.2](https://img.shields.io/static/v1?label=AppVersion&message=1.1.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* prometheus-node-exporter - use custom service accounts (#755)

### Default value changes

```diff
# No changes in this release
```

## 1.16.1

**Release date:** 2021-03-18

![AppVersion: 1.1.2](https://img.shields.io/static/v1?label=AppVersion&message=1.1.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] Allow to opt-out node exporter rootfs mount (#757)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 63ff1364..93f8ef80 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -80,6 +80,10 @@ endpoints: []
 # Expose the service to the host network
 hostNetwork: true
 
+## If true, node-exporter pods mounts host / at /host/root
+##
+hostRootFsMount: true
+
 ## Assign a group of affinity scheduling rules
 ##
 affinity: {}

```

## 1.16.0

**Release date:** 2021-03-18

![AppVersion: 1.1.2](https://img.shields.io/static/v1?label=AppVersion&message=1.1.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] AppVersion bump to 1.1.2 (#774)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 966e4174..63ff1364 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -3,7 +3,7 @@
 # Declare variables to be passed into your templates.
 image:
   repository: quay.io/prometheus/node-exporter
-  tag: v1.0.1
+  tag: v1.1.2
   pullPolicy: IfNotPresent
 
 service:

```

## 1.15.0

**Release date:** 2021-02-10

![AppVersion: 1.0.1](https://img.shields.io/static/v1?label=AppVersion&message=1.0.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Ability for custom dnsConfig on prometheus-node-exporter (#644)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 4be3f9c2..966e4174 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -100,6 +100,18 @@ podAnnotations:
 # Extra labels to be added to node exporter pods
 podLabels: {}
 
+# Custom DNS configuration to be added to prometheus-node-exporter pods
+dnsConfig: {}
+# nameservers:
+#   - 1.2.3.4
+# searches:
+#   - ns1.svc.cluster-domain.example
+#   - my.dns.search.suffix
+# options:
+#   - name: ndots
+#     value: "2"
+#   - name: edns0
+
 ## Assign a nodeSelector if operating a hybrid cluster
 ##
 nodeSelector: {}

```

## 1.14.2

**Release date:** 2021-02-03

![AppVersion: 1.0.1](https://img.shields.io/static/v1?label=AppVersion&message=1.0.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* add containerSecurityContext to node-exporter (#544)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 2b2205ce..4be3f9c2 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -60,6 +60,11 @@ securityContext:
   runAsNonRoot: true
   runAsUser: 65534
 
+containerSecurityContext: {}
+  # capabilities:
+  #   add:
+  #   - SYS_TIME
+
 rbac:
   ## If true, create & use RBAC resources
   ##

```

## 1.14.1

**Release date:** 2021-02-01

![AppVersion: 1.0.1](https://img.shields.io/static/v1?label=AppVersion&message=1.0.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] Add bismarck as maintainer (#625)

### Default value changes

```diff
# No changes in this release
```

## 1.14.0

**Release date:** 2021-02-01

![AppVersion: 1.0.1](https://img.shields.io/static/v1?label=AppVersion&message=1.0.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] Fix very slow GKE cluster upgrades (#485)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 8ef77484..2b2205ce 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -88,7 +88,9 @@ affinity: {}
 #                 - target-host-name
 
 # Annotations to be added to node exporter pods
-podAnnotations: {}
+podAnnotations:
+  # Fix for very slow GKE cluster upgrades
+  cluster-autoscaler.kubernetes.io/safe-to-evict: "true"
 
 # Extra labels to be added to node exporter pods
 podLabels: {}

```

## 1.13.0

**Release date:** 2021-01-29

![AppVersion: 1.0.1](https://img.shields.io/static/v1?label=AppVersion&message=1.0.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Adding secret support for oauth psidecar (#486)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 7edd8936..8ef77484 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -20,6 +20,9 @@ prometheus:
     enabled: false
     additionalLabels: {}
     namespace: ""
+    scheme: http
+    bearerTokenFile:
+    tlsConfig: {}
 
     relabelings: []
     scrapeTimeout: 10s
@@ -48,6 +51,7 @@ serviceAccount:
   # The name of the ServiceAccount to use.
   # If not set and create is true, a name is generated using the fullname template
   name:
+  annotations: {}
   imagePullSecrets: []
 
 securityContext:
@@ -122,7 +126,9 @@ extraHostVolumeMounts: []
 configmaps: []
 # - name: <configMapName>
 #   mountPath: <mountPath>
-
+secrets: []
+# - name: <secretName>
+#   mountPath: <mountPatch>
 ## Override the deployment namespace
 ##
 namespaceOverride: ""

```

## 1.12.0

**Release date:** 2020-10-30

![AppVersion: 1.0.1](https://img.shields.io/static/v1?label=AppVersion&message=1.0.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] Add rootfs mount by default (#80)

### Default value changes

```diff
# No changes in this release
```

## 1.11.2

**Release date:** 2020-08-20

![AppVersion: 1.0.1](https://img.shields.io/static/v1?label=AppVersion&message=1.0.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Prep initial charts indexing (#14)

### Default value changes

```diff
# No changes in this release
```

## 1.11.1

**Release date:** 2020-07-15

![AppVersion: 1.0.1](https://img.shields.io/static/v1?label=AppVersion&message=1.0.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-node-exporter] Version bump to 1.0.1 (#23151)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 3e0acd07..7edd8936 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -3,7 +3,7 @@
 # Declare variables to be passed into your templates.
 image:
   repository: quay.io/prometheus/node-exporter
-  tag: v1.0.0
+  tag: v1.0.1
   pullPolicy: IfNotPresent
 
 service:

```

## 1.11.0

**Release date:** 2020-06-16

![AppVersion: 1.0.0](https://img.shields.io/static/v1?label=AppVersion&message=1.0.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Add relabelings config for node-exporter ServiceMonitor (#22764)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 5e4b2f2b..3e0acd07 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -21,6 +21,7 @@ prometheus:
     additionalLabels: {}
     namespace: ""
 
+    relabelings: []
     scrapeTimeout: 10s
 
 ## Customize the updateStrategy if set

```

## 1.10.0

**Release date:** 2020-05-31

![AppVersion: 1.0.0](https://img.shields.io/static/v1?label=AppVersion&message=1.0.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-node-exporter] Upgrade to 1.0.0 & set group to non-root (#22603)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 868bfcd7..5e4b2f2b 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -3,7 +3,7 @@
 # Declare variables to be passed into your templates.
 image:
   repository: quay.io/prometheus/node-exporter
-  tag: v0.18.1
+  tag: v1.0.0
   pullPolicy: IfNotPresent
 
 service:
@@ -50,6 +50,8 @@ serviceAccount:
   imagePullSecrets: []
 
 securityContext:
+  fsGroup: 65534
+  runAsGroup: 65534
   runAsNonRoot: true
   runAsUser: 65534
 

```

## 1.9.1

**Release date:** 2020-03-19

![AppVersion: 0.18.1](https://img.shields.io/static/v1?label=AppVersion&message=0.18.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* prometheus node exporter: Configure exposing IP (#21490)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 1567c872..868bfcd7 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -11,6 +11,7 @@ service:
   port: 9100
   targetPort: 9100
   nodePort:
+  listenOnAllInterfaces: true
   annotations:
     prometheus.io/scrape: "true"
 

```

## 1.9.0

**Release date:** 2020-02-22

![AppVersion: 0.18.1](https://img.shields.io/static/v1?label=AppVersion&message=0.18.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-node-exporter] node-exporter add sidecar (#20952)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 10746628..1567c872 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -102,6 +102,7 @@ tolerations:
 ##
 extraArgs: []
 #   - --collector.diskstats.ignored-devices=^(ram|loop|fd|(h|s|v)d[a-z]|nvme\\d+n\\d+p)\\d+$
+#   - --collector.textfile.directory=/run/prometheus
 
 ## Additional mounts from the host
 ##
@@ -121,3 +122,16 @@ configmaps: []
 ## Override the deployment namespace
 ##
 namespaceOverride: ""
+
+## Additional containers for export metrics to text file
+##
+sidecars: []
+##  - name: nvidia-dcgm-exporter
+##    image: nvidia/dcgm-exporter:1.4.3
+
+## Volume for sidecar containers
+##
+sidecarVolumeMount: []
+##  - name: collector-textfiles
+##    mountPath: /run/prometheus
+##    readOnly: false

```

## 1.8.2

**Release date:** 2020-02-19

![AppVersion: 0.18.1](https://img.shields.io/static/v1?label=AppVersion&message=0.18.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-node-exporter] Add updateStrategy to template (#20858)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index bdac9906..10746628 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -22,6 +22,12 @@ prometheus:
 
     scrapeTimeout: 10s
 
+## Customize the updateStrategy if set
+updateStrategy:
+  type: RollingUpdate
+  rollingUpdate:
+    maxUnavailable: 1
+
 resources: {}
   # We usually recommend not to specify default resources and to leave this as a conscious
   # choice for the user. This also increases chances charts run on environments with little

```

## 1.8.1

**Release date:** 2019-12-17

![AppVersion: 0.18.1](https://img.shields.io/static/v1?label=AppVersion&message=0.18.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] - add podAnnotations value (#17740)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index cf793673..bdac9906 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -73,6 +73,12 @@ affinity: {}
 #               values:
 #                 - target-host-name
 
+# Annotations to be added to node exporter pods
+podAnnotations: {}
+
+# Extra labels to be added to node exporter pods
+podLabels: {}
+
 ## Assign a nodeSelector if operating a hybrid cluster
 ##
 nodeSelector: {}

```

## 1.8.0

**Release date:** 2019-11-19

![AppVersion: 0.18.1](https://img.shields.io/static/v1?label=AppVersion&message=0.18.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Prometheus node exporter namespace override (#18982)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 900111cc..cf793673 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -105,3 +105,7 @@ extraHostVolumeMounts: []
 configmaps: []
 # - name: <configMapName>
 #   mountPath: <mountPath>
+
+## Override the deployment namespace
+##
+namespaceOverride: ""

```

## 1.7.3

**Release date:** 2019-11-04

![AppVersion: 0.18.1](https://img.shields.io/static/v1?label=AppVersion&message=0.18.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-node-exporter] Fix maps -> slices, bump version (#18527)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 9fa06863..900111cc 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -3,7 +3,7 @@
 # Declare variables to be passed into your templates.
 image:
   repository: quay.io/prometheus/node-exporter
-  tag: v0.18.0
+  tag: v0.18.1
   pullPolicy: IfNotPresent
 
 service:
@@ -88,12 +88,12 @@ tolerations:
 
 ## Additional container arguments
 ##
-extraArgs: {}
+extraArgs: []
 #   - --collector.diskstats.ignored-devices=^(ram|loop|fd|(h|s|v)d[a-z]|nvme\\d+n\\d+p)\\d+$
 
 ## Additional mounts from the host
 ##
-extraHostVolumeMounts: {}
+extraHostVolumeMounts: []
 #  - name: <mountName>
 #    hostPath: <hostPath>
 #    mountPath: <mountPath>
@@ -102,6 +102,6 @@ extraHostVolumeMounts: {}
 
 ## Additional configmaps to be mounted.
 ##
-configmaps: {}
+configmaps: []
 # - name: <configMapName>
 #   mountPath: <mountPath>

```

## 1.7.2

**Release date:** 2019-10-31

![AppVersion: 0.18.0](https://img.shields.io/static/v1?label=AppVersion&message=0.18.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-node-exporter] Add vsliouniaev as owner (#18461)

### Default value changes

```diff
# No changes in this release
```

## 1.7.1

**Release date:** 2019-10-27

![AppVersion: 0.18.0](https://img.shields.io/static/v1?label=AppVersion&message=0.18.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-node-exporter] update NOTES.txt to show correct port (#17931)

### Default value changes

```diff
# No changes in this release
```

## 1.7.0

**Release date:** 2019-10-10

![AppVersion: 0.18.0](https://img.shields.io/static/v1?label=AppVersion&message=0.18.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] Use latest API versions (#17785)

### Default value changes

```diff
# No changes in this release
```

## 1.6.0

**Release date:** 2019-07-28

![AppVersion: 0.18.0](https://img.shields.io/static/v1?label=AppVersion&message=0.18.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* allow mounting configmaps (#15356)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index b73b75c5..9fa06863 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -99,3 +99,9 @@ extraHostVolumeMounts: {}
 #    mountPath: <mountPath>
 #    readOnly: true|false
 #    mountPropagation: None|HostToContainer|Bidirectional
+
+## Additional configmaps to be mounted.
+##
+configmaps: {}
+# - name: <configMapName>
+#   mountPath: <mountPath>

```

## 1.5.2

**Release date:** 2019-07-26

![AppVersion: 0.18.0](https://img.shields.io/static/v1?label=AppVersion&message=0.18.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* removed errant space in commented line (#15661)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 2da6571c..b73b75c5 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -29,7 +29,7 @@ resources: {}
   # lines, adjust them as necessary, and remove the curly braces after 'resources:'.
   # limits:
   #   cpu: 200m
-  #    memory: 50Mi
+  #   memory: 50Mi
   # requests:
   #   cpu: 100m
   #   memory: 30Mi

```

## 1.5.1

**Release date:** 2019-06-28

![AppVersion: 0.18.0](https://img.shields.io/static/v1?label=AppVersion&message=0.18.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-node-exporter] add scrapeTimeout for the service monitor (#14293)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 7ca723ff..2da6571c 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -20,6 +20,8 @@ prometheus:
     additionalLabels: {}
     namespace: ""
 
+    scrapeTimeout: 10s
+
 resources: {}
   # We usually recommend not to specify default resources and to leave this as a conscious
   # choice for the user. This also increases chances charts run on environments with little

```

## 1.5.0

**Release date:** 2019-05-20

![AppVersion: 0.18.0](https://img.shields.io/static/v1?label=AppVersion&message=0.18.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-node-exporter] Upgrade to 0.18.0 (#13739)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 09261c4f..7ca723ff 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -3,7 +3,7 @@
 # Declare variables to be passed into your templates.
 image:
   repository: quay.io/prometheus/node-exporter
-  tag: v0.17.0
+  tag: v0.18.0
   pullPolicy: IfNotPresent
 
 service:

```

## 1.4.2

**Release date:** 2019-04-03

![AppVersion: 0.17.0](https://img.shields.io/static/v1?label=AppVersion&message=0.17.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-node-exporter] configurable host network (#12747)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 14bcfc53..09261c4f 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -56,6 +56,9 @@ rbac:
 # their addresses here
 endpoints: []
 
+# Expose the service to the host network
+hostNetwork: true
+
 ## Assign a group of affinity scheduling rules
 ##
 affinity: {}

```

## 1.3.2

**Release date:** 2019-03-07

![AppVersion: 0.17.0](https://img.shields.io/static/v1?label=AppVersion&message=0.17.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* add prometheus service monitor to prometheus node exporter (#12004)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index fcd7acbd..14bcfc53 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -14,6 +14,12 @@ service:
   annotations:
     prometheus.io/scrape: "true"
 
+prometheus:
+  monitor:
+    enabled: false
+    additionalLabels: {}
+    namespace: ""
+
 resources: {}
   # We usually recommend not to specify default resources and to leave this as a conscious
   # choice for the user. This also increases chances charts run on environments with little

```

## 1.3.1

**Release date:** 2019-03-07

![AppVersion: 0.17.0](https://img.shields.io/static/v1?label=AppVersion&message=0.17.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-node-exporter] Fix imagePullSecrets value path (#11222)

### Default value changes

```diff
# No changes in this release
```

## 1.3.0

**Release date:** 2019-02-18

![AppVersion: 0.17.0](https://img.shields.io/static/v1?label=AppVersion&message=0.17.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-node-exporter] Add options to change volume mount propagation  (#11194)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 5f4298c6..fcd7acbd 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -87,3 +87,4 @@ extraHostVolumeMounts: {}
 #    hostPath: <hostPath>
 #    mountPath: <mountPath>
 #    readOnly: true|false
+#    mountPropagation: None|HostToContainer|Bidirectional

```

## 1.2.0

**Release date:** 2019-01-31

![AppVersion: 0.17.0](https://img.shields.io/static/v1?label=AppVersion&message=0.17.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-node-exporter] Add nodeSelector to DaemonSet (#10952)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 6df24368..5f4298c6 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -50,6 +50,24 @@ rbac:
 # their addresses here
 endpoints: []
 
+## Assign a group of affinity scheduling rules
+##
+affinity: {}
+#   nodeAffinity:
+#     requiredDuringSchedulingIgnoredDuringExecution:
+#       nodeSelectorTerms:
+#         - matchFields:
+#             - key: metadata.name
+#               operator: In
+#               values:
+#                 - target-host-name
+
+## Assign a nodeSelector if operating a hybrid cluster
+##
+nodeSelector: {}
+#   beta.kubernetes.io/arch: amd64
+#   beta.kubernetes.io/os: linux
+
 tolerations:
   - effect: NoSchedule
     operator: Exists

```

## 1.1.0

**Release date:** 2018-12-31

![AppVersion: 0.17.0](https://img.shields.io/static/v1?label=AppVersion&message=0.17.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-node-exporter] Node exporter endpoints (#10243)

### Default value changes

```diff
# No changes in this release
```

## 1.0.1

**Release date:** 2018-12-20

![AppVersion: 0.17.0](https://img.shields.io/static/v1?label=AppVersion&message=0.17.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] - Bump app version to 0.17.0 ADDENDUM (#10125)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 31694243..6df24368 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -3,7 +3,7 @@
 # Declare variables to be passed into your templates.
 image:
   repository: quay.io/prometheus/node-exporter
-  tag: v0.16.0
+  tag: v0.17.0
   pullPolicy: IfNotPresent
 
 service:

```

## 1.0.0

**Release date:** 2018-12-17

![AppVersion: 0.17.0](https://img.shields.io/static/v1?label=AppVersion&message=0.17.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-node-exporter]Add support for nodePort (#9754)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 7d3fc76a..31694243 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -10,6 +10,7 @@ service:
   type: ClusterIP
   port: 9100
   targetPort: 9100
+  nodePort:
   annotations:
     prometheus.io/scrape: "true"
 

```

## 0.6.2

**Release date:** 2018-12-17

![AppVersion: 0.17.0](https://img.shields.io/static/v1?label=AppVersion&message=0.17.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] - Bump app version to 0.17.0 (#10058)

### Default value changes

```diff
# No changes in this release
```

## 0.6.1

**Release date:** 2018-12-02

![AppVersion: 0.16.0](https://img.shields.io/static/v1?label=AppVersion&message=0.16.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-node-exporter] add securityContext value (#9618)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 38a70597..7d3fc76a 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -33,6 +33,10 @@ serviceAccount:
   name:
   imagePullSecrets: []
 
+securityContext:
+  runAsNonRoot: true
+  runAsUser: 65534
+
 rbac:
   ## If true, create & use RBAC resources
   ##

```

## 0.6.0

**Release date:** 2018-11-22

![AppVersion: 0.16.0](https://img.shields.io/static/v1?label=AppVersion&message=0.16.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Add priorityClassName option to values (#9483)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 6fac1b1b..38a70597 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -49,6 +49,9 @@ tolerations:
   - effect: NoSchedule
     operator: Exists
 
+## Assign a PriorityClassName to pods if set
+# priorityClassName: ""
+
 ## Additional container arguments
 ##
 extraArgs: {}

```

## 0.5.1

**Release date:** 2018-11-05

![AppVersion: 0.16.0](https://img.shields.io/static/v1?label=AppVersion&message=0.16.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Fix typo causing RBAC resources not being created (#8991)

### Default value changes

```diff
# No changes in this release
```

## 0.5.0

**Release date:** 2018-10-04

![AppVersion: 0.16.0](https://img.shields.io/static/v1?label=AppVersion&message=0.16.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] allow setting additional host mounts (#8117)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index effe0b42..6fac1b1b 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -53,3 +53,11 @@ tolerations:
 ##
 extraArgs: {}
 #   - --collector.diskstats.ignored-devices=^(ram|loop|fd|(h|s|v)d[a-z]|nvme\\d+n\\d+p)\\d+$
+
+## Additional mounts from the host
+##
+extraHostVolumeMounts: {}
+#  - name: <mountName>
+#    hostPath: <hostPath>
+#    mountPath: <mountPath>
+#    readOnly: true|false

```

## 0.4.0

**Release date:** 2018-10-02

![AppVersion: 0.16.0](https://img.shields.io/static/v1?label=AppVersion&message=0.16.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* added service annotations with the default value (#8086)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 18a6693c..effe0b42 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -3,13 +3,15 @@
 # Declare variables to be passed into your templates.
 image:
   repository: quay.io/prometheus/node-exporter
-  tag: v0.15.2
+  tag: v0.16.0
   pullPolicy: IfNotPresent
 
 service:
   type: ClusterIP
   port: 9100
   targetPort: 9100
+  annotations:
+    prometheus.io/scrape: "true"
 
 resources: {}
   # We usually recommend not to specify default resources and to leave this as a conscious

```

## 0.3.0

**Release date:** 2018-07-26

![AppVersion: 0.15.2](https://img.shields.io/static/v1?label=AppVersion&message=0.15.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-node-exporter] Add ability to customize labels and endpoints, Pod security policy (#6742)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 30db7c53..18a6693c 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -29,6 +29,19 @@ serviceAccount:
   # The name of the ServiceAccount to use.
   # If not set and create is true, a name is generated using the fullname template
   name:
+  imagePullSecrets: []
+
+rbac:
+  ## If true, create & use RBAC resources
+  ##
+  create: true
+  ## If true, create & use Pod Security Policy resources
+  ## https://kubernetes.io/docs/concepts/policy/pod-security-policy/
+  pspEnabled: true
+
+# for deployments that have node_exporter deployed outside of the cluster, list
+# their addresses here
+endpoints: []
 
 tolerations:
   - effect: NoSchedule
@@ -36,5 +49,5 @@ tolerations:
 
 ## Additional container arguments
 ##
-# extraArgs:
+extraArgs: {}
 #   - --collector.diskstats.ignored-devices=^(ram|loop|fd|(h|s|v)d[a-z]|nvme\\d+n\\d+p)\\d+$

```

## 0.2.0

**Release date:** 2018-05-23

![AppVersion: 0.15.2](https://img.shields.io/static/v1?label=AppVersion&message=0.15.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [Node-exporter] Add extra flags arguments and rollingUpdate strategy (#5699)

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 0a635369..30db7c53 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -33,3 +33,8 @@ serviceAccount:
 tolerations:
   - effect: NoSchedule
     operator: Exists
+
+## Additional container arguments
+##
+# extraArgs:
+#   - --collector.diskstats.ignored-devices=^(ram|loop|fd|(h|s|v)d[a-z]|nvme\\d+n\\d+p)\\d+$

```

## 0.1.2

**Release date:** 2018-04-27

![AppVersion: 0.15.2](https://img.shields.io/static/v1?label=AppVersion&message=0.15.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/promehteus-node-exporter] Fix health checks (#5295)

### Default value changes

```diff
# No changes in this release
```

## 0.1.1

**Release date:** 2018-04-03

![AppVersion: 0.15.2](https://img.shields.io/static/v1?label=AppVersion&message=0.15.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* stable/prometheus-node-exporter/README.md (#4550)

### Default value changes

```diff
# No changes in this release
```

## 0.1.0

**Release date:** 2018-03-30

![AppVersion: 0.15.2](https://img.shields.io/static/v1?label=AppVersion&message=0.15.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* add prometheus node exporter chart (#4260)

### Default value changes

```diff
# Default values for prometheus-node-exporter.
# This is a YAML-formatted file.
# Declare variables to be passed into your templates.
image:
  repository: quay.io/prometheus/node-exporter
  tag: v0.15.2
  pullPolicy: IfNotPresent

service:
  type: ClusterIP
  port: 9100
  targetPort: 9100

resources: {}
  # We usually recommend not to specify default resources and to leave this as a conscious
  # choice for the user. This also increases chances charts run on environments with little
  # resources, such as Minikube. If you do want to specify resources, uncomment the following
  # lines, adjust them as necessary, and remove the curly braces after 'resources:'.
  # limits:
  #   cpu: 200m
  #    memory: 50Mi
  # requests:
  #   cpu: 100m
  #   memory: 30Mi

serviceAccount:
  # Specifies whether a ServiceAccount should be created
  create: true
  # The name of the ServiceAccount to use.
  # If not set and create is true, a name is generated using the fullname template
  name:

tolerations:
  - effect: NoSchedule
    operator: Exists

```

---
Autogenerated from Helm Chart and git history using [helm-changelog](https://github.com/mogensen/helm-changelog)
