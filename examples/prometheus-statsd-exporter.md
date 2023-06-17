# Change Log

## 0.8.0

**Release date:** 2023-03-23

![AppVersion: v0.22.8](https://img.shields.io/static/v1?label=AppVersion&message=v0.22.8&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-statsd-exporter] add podLabels to add extra labels to pods (#3090)

### Default value changes

```diff
diff --git a/charts/prometheus-statsd-exporter/values.yaml b/charts/prometheus-statsd-exporter/values.yaml
index 8e1d51af..279bac8f 100644
--- a/charts/prometheus-statsd-exporter/values.yaml
+++ b/charts/prometheus-statsd-exporter/values.yaml
@@ -106,6 +106,9 @@ serviceAccount:
 
 podAnnotations: {}
 
+# Extra labels to be added to pods
+podLabels: {}
+
 podSecurityContext: {}
   # fsGroup: 2000
 

```

## 0.7.0

**Release date:** 2022-11-07

![AppVersion: v0.22.8](https://img.shields.io/static/v1?label=AppVersion&message=v0.22.8&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-statsd-exporter] add PodMonitoring for Google Managed Prometheus (#2596)

### Default value changes

```diff
diff --git a/charts/prometheus-statsd-exporter/values.yaml b/charts/prometheus-statsd-exporter/values.yaml
index f5705dcc..8e1d51af 100644
--- a/charts/prometheus-statsd-exporter/values.yaml
+++ b/charts/prometheus-statsd-exporter/values.yaml
@@ -9,7 +9,7 @@ image:
   repository: prom/statsd-exporter
   pullPolicy: IfNotPresent
   # Overrides the image tag whose default is the chart appVersion.
-  tag: v0.22.7
+  tag: ""
 
 imagePullSecrets: []
 nameOverride: ""
@@ -63,6 +63,19 @@ readinessProbe:
     path: /health
     port: http
 
+prometheus:
+  monitor:
+    enabled: false
+    additionalLabels: {}
+    namespace: ""
+    interval: 30s
+    # ApiVersion for the podMonitor Resource(defaults to "monitoring.googleapis.com/v1")
+    apiVersion: ""
+    ## resource kind(defaults to "PodMonitoring")
+    kind: ""
+    # metrics exposing endpoint
+    endpoint: "/metrics"
+
 serviceMonitor:
   enabled: false
   interval: 30s
@@ -70,11 +83,17 @@ serviceMonitor:
   namespace: monitoring
   honorLabels: false
   additionalLabels: {}
+  ## resource kind(defaults to "ServiceMonitor")
+  kind: ""
   # MetricRelabelConfigs to apply to samples before ingestion
   metricRelabelings: []
   # RelabelConfigs to apply to samples before scraping.
   # More info https://prometheus.io/docs/prometheus/latest/configuration/configuration/#relabel_config
   relabelings: []
+  # ApiVersion for the podMonitor Resource(defaults to "monitoring.coreos.com/v1")
+  apiVersion: ""
+  # metrics exposing endpoint
+  endpoint: "/metrics"
 
 serviceAccount:
   # Specifies whether a service account should be created

```

## 0.6.2

**Release date:** 2022-10-09

![AppVersion: 0.22.7](https://img.shields.io/static/v1?label=AppVersion&message=0.22.7&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-statsd-exporter] override key in mapping ConfigMap (#2532)

### Default value changes

```diff
diff --git a/charts/prometheus-statsd-exporter/values.yaml b/charts/prometheus-statsd-exporter/values.yaml
index c2386a76..f5705dcc 100644
--- a/charts/prometheus-statsd-exporter/values.yaml
+++ b/charts/prometheus-statsd-exporter/values.yaml
@@ -47,6 +47,9 @@ statsd:
   # Metric mapping ConfigMap
 #  mappingConfigMapName: ""
 
+  # Name of the key inside Metric mapping ConfigMap. Default is "statsd.mappingConf".
+#  mappingConfigMapKey: ""
+
   # Metric mapping configuration
 #  mappingConfig: |-
 

```

## 0.6.1

**Release date:** 2022-10-09

![AppVersion: 0.22.7](https://img.shields.io/static/v1?label=AppVersion&message=0.22.7&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-statsd-exporter] support relabelings and metricsRelabelings in ServiceMonitor (#2531)

### Default value changes

```diff
diff --git a/charts/prometheus-statsd-exporter/values.yaml b/charts/prometheus-statsd-exporter/values.yaml
index 5fe85fea..c2386a76 100644
--- a/charts/prometheus-statsd-exporter/values.yaml
+++ b/charts/prometheus-statsd-exporter/values.yaml
@@ -67,6 +67,11 @@ serviceMonitor:
   namespace: monitoring
   honorLabels: false
   additionalLabels: {}
+  # MetricRelabelConfigs to apply to samples before ingestion
+  metricRelabelings: []
+  # RelabelConfigs to apply to samples before scraping.
+  # More info https://prometheus.io/docs/prometheus/latest/configuration/configuration/#relabel_config
+  relabelings: []
 
 serviceAccount:
   # Specifies whether a service account should be created

```

## 0.6.0

**Release date:** 2022-08-10

![AppVersion: 0.22.7](https://img.shields.io/static/v1?label=AppVersion&message=0.22.7&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-statsd-exporter]: parameterize readiness and liveness pro… (#2355)

### Default value changes

```diff
diff --git a/charts/prometheus-statsd-exporter/values.yaml b/charts/prometheus-statsd-exporter/values.yaml
index 56d78442..5fe85fea 100644
--- a/charts/prometheus-statsd-exporter/values.yaml
+++ b/charts/prometheus-statsd-exporter/values.yaml
@@ -50,6 +50,16 @@ statsd:
   # Metric mapping configuration
 #  mappingConfig: |-
 
+livenessProbe:
+  httpGet:
+    path: /health
+    port: http
+
+readinessProbe:
+  httpGet:
+    path: /health
+    port: http
+
 serviceMonitor:
   enabled: false
   interval: 30s

```

## 0.5.0

**Release date:** 2022-07-14

![AppVersion: 0.22.7](https://img.shields.io/static/v1?label=AppVersion&message=0.22.7&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-statsd-exporter] Add extraEnv and extraArgs (#2253)

### Default value changes

```diff
diff --git a/charts/prometheus-statsd-exporter/values.yaml b/charts/prometheus-statsd-exporter/values.yaml
index 93b6dc3f..56d78442 100644
--- a/charts/prometheus-statsd-exporter/values.yaml
+++ b/charts/prometheus-statsd-exporter/values.yaml
@@ -9,12 +9,20 @@ image:
   repository: prom/statsd-exporter
   pullPolicy: IfNotPresent
   # Overrides the image tag whose default is the chart appVersion.
-  tag: v0.22.1
+  tag: v0.22.7
 
 imagePullSecrets: []
 nameOverride: ""
 fullnameOverride: ""
 
+# Additional container environment variables
+extraEnv: {}
+#   HTTP_PROXY: "http://superproxy.com:3128"
+#   NO_PROXY: "localhost,127.0.0.1"
+
+# Additional arguments in the statsd_exporter command line
+extraArgs: []
+  # - --history.limit=1000
 
 statsd:
   # The UDP port on which to receive statsd metric lines.

```

## 0.4.2

**Release date:** 2021-11-15

![AppVersion: 0.22.1](https://img.shields.io/static/v1?label=AppVersion&message=0.22.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Fixed a bug that does not add mapping config when configuring a config map using mappingConfigMapName field (#1508)

### Default value changes

```diff
# No changes in this release
```

## 0.4.1

**Release date:** 2021-10-05

![AppVersion: 0.22.1](https://img.shields.io/static/v1?label=AppVersion&message=0.22.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* fix statsd exporter ingress for k8s >= 1.19 & mappingConfigMapName (#1355)

### Default value changes

```diff
diff --git a/charts/prometheus-statsd-exporter/values.yaml b/charts/prometheus-statsd-exporter/values.yaml
index b84b9417..93b6dc3f 100644
--- a/charts/prometheus-statsd-exporter/values.yaml
+++ b/charts/prometheus-statsd-exporter/values.yaml
@@ -82,13 +82,15 @@ service:
 
 ingress:
   enabled: false
+  className: ""
   annotations: {}
     # kubernetes.io/ingress.class: nginx
     # kubernetes.io/tls-acme: "true"
   hosts:
     - host: chart-example.local
-      #  pathType: ImplementationSpecific # only Kubernetes v1.19+
-      paths: []
+      paths:
+        - path: /
+          pathType: ImplementationSpecific
   tls: []
   #  - secretName: chart-example-tls
   #    hosts:

```

## 0.4.0

**Release date:** 2021-09-20

![AppVersion: 0.22.1](https://img.shields.io/static/v1?label=AppVersion&message=0.22.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-statsd-exporter] Bump statsd-exporter to version to 0.22.1 (#1303)

### Default value changes

```diff
diff --git a/charts/prometheus-statsd-exporter/values.yaml b/charts/prometheus-statsd-exporter/values.yaml
index 26a99a0d..b84b9417 100644
--- a/charts/prometheus-statsd-exporter/values.yaml
+++ b/charts/prometheus-statsd-exporter/values.yaml
@@ -9,7 +9,7 @@ image:
   repository: prom/statsd-exporter
   pullPolicy: IfNotPresent
   # Overrides the image tag whose default is the chart appVersion.
-  tag: v0.20.0
+  tag: v0.22.1
 
 imagePullSecrets: []
 nameOverride: ""

```

## 0.3.1

**Release date:** 2021-03-10

![AppVersion: 0.20.0](https://img.shields.io/static/v1?label=AppVersion&message=0.20.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* prometheus-statsd-exporter - Fix statsd service monitor config (#732)

### Default value changes

```diff
# No changes in this release
```

## 0.3.0

**Release date:** 2021-02-09

![AppVersion: 0.20.0](https://img.shields.io/static/v1?label=AppVersion&message=0.20.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* update statsd-exporter version (#658)

### Default value changes

```diff
diff --git a/charts/prometheus-statsd-exporter/values.yaml b/charts/prometheus-statsd-exporter/values.yaml
index a7fc28f0..26a99a0d 100644
--- a/charts/prometheus-statsd-exporter/values.yaml
+++ b/charts/prometheus-statsd-exporter/values.yaml
@@ -9,7 +9,7 @@ image:
   repository: prom/statsd-exporter
   pullPolicy: IfNotPresent
   # Overrides the image tag whose default is the chart appVersion.
-  tag: v0.18.0
+  tag: v0.20.0
 
 imagePullSecrets: []
 nameOverride: ""

```

## 0.2.0

**Release date:** 2021-01-09

![AppVersion: 0.18.0](https://img.shields.io/static/v1?label=AppVersion&message=0.18.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-statsd-exporter] Adding `honorLabels` to `statsd-exporter` `ServiceMonitor` (#534)

### Default value changes

```diff
diff --git a/charts/prometheus-statsd-exporter/values.yaml b/charts/prometheus-statsd-exporter/values.yaml
index d10d3d10..a7fc28f0 100644
--- a/charts/prometheus-statsd-exporter/values.yaml
+++ b/charts/prometheus-statsd-exporter/values.yaml
@@ -47,6 +47,7 @@ serviceMonitor:
   interval: 30s
   scrapeTimeout: 10s
   namespace: monitoring
+  honorLabels: false
   additionalLabels: {}
 
 serviceAccount:

```

## 0.1.0

**Release date:** 2020-12-08

![AppVersion: 0.18.0](https://img.shields.io/static/v1?label=AppVersion&message=0.18.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-statsd-exporter] add first release (#449)

### Default value changes

```diff
# Default values for prometheus-statsd-exporter.
# This is a YAML-formatted file.
# Declare variables to be passed into your templates.

replicaCount: 1
# deploymentRevisionHistoryLimit: 10

image:
  repository: prom/statsd-exporter
  pullPolicy: IfNotPresent
  # Overrides the image tag whose default is the chart appVersion.
  tag: v0.18.0

imagePullSecrets: []
nameOverride: ""
fullnameOverride: ""


statsd:
  # The UDP port on which to receive statsd metric lines.
  udpPort: 9125

  # The TCP port on which to receive statsd metric lines.
  tcpPort: 9125

  # Maximum size of your metric mapping cache.
  # Relies on least recently used replacement policy if max size is reached.
  cacheSize: 1000

  # Size of internal queue for processing events.
  eventQueueSize: 10000

  # Number of events to hold in queue before flushing.
  eventFlushThreshold: 1000

  # Time interval before flushing events in queue.
  eventFlushInterval: 200ms

  # Metric mapping ConfigMap
#  mappingConfigMapName: ""

  # Metric mapping configuration
#  mappingConfig: |-

serviceMonitor:
  enabled: false
  interval: 30s
  scrapeTimeout: 10s
  namespace: monitoring
  additionalLabels: {}

serviceAccount:
  # Specifies whether a service account should be created
  create: true
  # Annotations to add to the service account
  annotations: {}
  # The name of the service account to use.
  # If not set and create is true, a name is generated using the fullname template
  name: ""

podAnnotations: {}

podSecurityContext: {}
  # fsGroup: 2000

securityContext: {}
  # capabilities:
  #   drop:
  #   - ALL
  # readOnlyRootFilesystem: true
  # runAsNonRoot: true
  # runAsUser: 1000

service:
  type: ClusterIP
  # The address on which to expose the web interface and generated Prometheus metrics.
  port: 9102
  # Path under which to expose metrics.
  path: /metrics
  annotations: {}

ingress:
  enabled: false
  annotations: {}
    # kubernetes.io/ingress.class: nginx
    # kubernetes.io/tls-acme: "true"
  hosts:
    - host: chart-example.local
      #  pathType: ImplementationSpecific # only Kubernetes v1.19+
      paths: []
  tls: []
  #  - secretName: chart-example-tls
  #    hosts:
  #      - chart-example.local

resources: {}
  # We usually recommend not to specify default resources and to leave this as a conscious
  # choice for the user. This also increases chances charts run on environments with little
  # resources, such as Minikube. If you do want to specify resources, uncomment the following
  # lines, adjust them as necessary, and remove the curly braces after 'resources:'.
  # limits:
  #   cpu: 100m
  #   memory: 128Mi
  # requests:
  #   cpu: 100m
  #   memory: 128Mi

autoscaling:
  enabled: false
  minReplicas: 1
  maxReplicas: 100
  targetCPUUtilizationPercentage: 80
  # targetMemoryUtilizationPercentage: 80

nodeSelector: {}

tolerations: []

affinity: {}

```

---
Autogenerated from Helm Chart and git history using [helm-changelog](https://github.com/mogensen/helm-changelog)
