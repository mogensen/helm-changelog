# Change Log

## 4.3.0

**Release date:** 2023-05-08

![AppVersion: 0.13.0](https://img.shields.io/static/v1?label=AppVersion&message=0.13.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-stackdriver-exporter] Add prometheusRule support (#3115)

### Default value changes

```diff
diff --git a/charts/prometheus-stackdriver-exporter/values.yaml b/charts/prometheus-stackdriver-exporter/values.yaml
index 07aa8a13..b6f06365 100644
--- a/charts/prometheus-stackdriver-exporter/values.yaml
+++ b/charts/prometheus-stackdriver-exporter/values.yaml
@@ -146,3 +146,11 @@ serviceMonitor:
   metricRelabelings: []
   # RelabelConfigs to apply to samples before scraping. https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
   relabelings: []
+
+## Custom PrometheusRules to be defined
+## ref: https://github.com/coreos/prometheus-operator#customresourcedefinitions
+prometheusRule:
+  enabled: false
+  additionalLabels: {}
+  namespace: ""
+  rules: []

```

## 4.2.0

**Release date:** 2023-02-04

![AppVersion: 0.13.0](https://img.shields.io/static/v1?label=AppVersion&message=0.13.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-stackdriver-exporter] Adding new params: aggregate-deltas and aggregate-deltas-ttl (#2987)

### Default value changes

```diff
diff --git a/charts/prometheus-stackdriver-exporter/values.yaml b/charts/prometheus-stackdriver-exporter/values.yaml
index 774c5d05..07aa8a13 100644
--- a/charts/prometheus-stackdriver-exporter/values.yaml
+++ b/charts/prometheus-stackdriver-exporter/values.yaml
@@ -12,7 +12,7 @@ restartPolicy: Always
 
 image:
   repository: prometheuscommunity/stackdriver-exporter
-  tag: v0.12.0
+  tag: v0.13.0
   pullPolicy: IfNotPresent
 
   ## Optionally specify an array of imagePullSecrets.
@@ -81,6 +81,10 @@ stackdriver:
     offset: '0s'
     # Offset for the Google Stackdriver Monitoring Metrics interval into the past by the ingest delay from the metric's metadata.
     ingestDelay: false
+    # If enabled will treat all DELTA metrics as an in-memory counter instead of a gauge.
+    aggregateDeltas: false
+    # How long should a delta metric continue to be exported after GCP stops producing a metric
+    aggregateDeltasTTL: '30m'
 
 web:
   # Port to listen on

```

## 4.1.0

**Release date:** 2022-09-26

![AppVersion: 0.12.0](https://img.shields.io/static/v1?label=AppVersion&message=0.12.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-stackdriver-exporter] Adding support for container-level security context settings. (#2486)

### Default value changes

```diff
diff --git a/charts/prometheus-stackdriver-exporter/values.yaml b/charts/prometheus-stackdriver-exporter/values.yaml
index 7681ac26..774c5d05 100644
--- a/charts/prometheus-stackdriver-exporter/values.yaml
+++ b/charts/prometheus-stackdriver-exporter/values.yaml
@@ -32,6 +32,8 @@ resources: {}
 
 securityContext: {}
 
+containerSecurityContext: {}
+
 service:
   type: ClusterIP
   httpPort: 9255

```

## 4.0.1

**Release date:** 2022-08-18

![AppVersion: 0.12.0](https://img.shields.io/static/v1?label=AppVersion&message=0.12.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-stackdriver-exporter] Fix spec.serviceAccountName (#2378)

### Default value changes

```diff
# No changes in this release
```

## 4.0.0

**Release date:** 2022-08-01

![AppVersion: 0.12.0](https://img.shields.io/static/v1?label=AppVersion&message=0.12.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-stackdriver] Allow passing multiple metric filter values. (#2330)

### Default value changes

```diff
diff --git a/charts/prometheus-stackdriver-exporter/values.yaml b/charts/prometheus-stackdriver-exporter/values.yaml
index e7e9175d..7681ac26 100644
--- a/charts/prometheus-stackdriver-exporter/values.yaml
+++ b/charts/prometheus-stackdriver-exporter/values.yaml
@@ -68,8 +68,11 @@ stackdriver:
   metrics:
     # The prefixes to gather metrics for, we default to just CPU metrics.
     typePrefixes: 'compute.googleapis.com/instance/cpu'
-    # The monitoring filters to specify resources, metric types, group definitions, and time series.
-    filters: ""
+    # The filters to refine the metrics query by using Filter objects that Google provides.
+    # Filter objects: project, group.id, resource.type, resource.labels.[KEY], metric.type, metric.labels.[KEY]
+    # https://cloud.google.com/monitoring/api/v3/filters
+    filters: []
+      # - 'pubsub.googleapis.com/subscription:resource.labels.subscription_id=monitoring.regex.full_match("us-west4.*my-team.*")'
     # The frequency to request
     interval: '5m'
     # How far into the past to offset

```

## 3.0.2

**Release date:** 2022-07-28

![AppVersion: 0.12.0](https://img.shields.io/static/v1?label=AppVersion&message=0.12.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Stackdriver exporter: Add monitoring filters (#2314)

### Default value changes

```diff
diff --git a/charts/prometheus-stackdriver-exporter/values.yaml b/charts/prometheus-stackdriver-exporter/values.yaml
index bd0c754d..e7e9175d 100644
--- a/charts/prometheus-stackdriver-exporter/values.yaml
+++ b/charts/prometheus-stackdriver-exporter/values.yaml
@@ -68,6 +68,8 @@ stackdriver:
   metrics:
     # The prefixes to gather metrics for, we default to just CPU metrics.
     typePrefixes: 'compute.googleapis.com/instance/cpu'
+    # The monitoring filters to specify resources, metric types, group definitions, and time series.
+    filters: ""
     # The frequency to request
     interval: '5m'
     # How far into the past to offset

```

## 3.0.1

**Release date:** 2022-06-13

![AppVersion: 0.12.0](https://img.shields.io/static/v1?label=AppVersion&message=0.12.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-stackdriver-exporter] Add: README instructions to upgrade 2.x to 3.x (#2146)

### Default value changes

```diff
# No changes in this release
```

## 3.0.0

**Release date:** 2022-06-10

![AppVersion: 0.12.0](https://img.shields.io/static/v1?label=AppVersion&message=0.12.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-stackdriver-exporter] Add: Kubernetes recommended labels and custom labels (#2125)

### Default value changes

```diff
diff --git a/charts/prometheus-stackdriver-exporter/values.yaml b/charts/prometheus-stackdriver-exporter/values.yaml
index eae5db4e..bd0c754d 100644
--- a/charts/prometheus-stackdriver-exporter/values.yaml
+++ b/charts/prometheus-stackdriver-exporter/values.yaml
@@ -37,6 +37,10 @@ service:
   httpPort: 9255
   annotations: {}
 
+## Additional labels to add to all resources
+customLabels: {}
+  # app: prometheus-stackdriver-exporter
+
 secret:
   labels: {}
 

```

## 2.2.0

**Release date:** 2022-03-03

![AppVersion: 0.12.0](https://img.shields.io/static/v1?label=AppVersion&message=0.12.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-stackdriver-exporter] Add option to enable metrics ingest delay (#1844)

### Default value changes

```diff
diff --git a/charts/prometheus-stackdriver-exporter/values.yaml b/charts/prometheus-stackdriver-exporter/values.yaml
index f2c5de62..eae5db4e 100644
--- a/charts/prometheus-stackdriver-exporter/values.yaml
+++ b/charts/prometheus-stackdriver-exporter/values.yaml
@@ -68,6 +68,8 @@ stackdriver:
     interval: '5m'
     # How far into the past to offset
     offset: '0s'
+    # Offset for the Google Stackdriver Monitoring Metrics interval into the past by the ingest delay from the metric's metadata.
+    ingestDelay: false
 
 web:
   # Port to listen on

```

## 2.1.1

**Release date:** 2022-03-02

![AppVersion: 0.12.0](https://img.shields.io/static/v1?label=AppVersion&message=0.12.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-stackdriver-exporter] ServiceMonitor: allow to specify scrapeTimeout independently (#1841)

### Default value changes

```diff
# No changes in this release
```

## 2.1.0

**Release date:** 2022-02-15

![AppVersion: 0.12.0](https://img.shields.io/static/v1?label=AppVersion&message=0.12.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-stackdriver-exporter] Add support for imagePullSecrets (#1794)

### Default value changes

```diff
diff --git a/charts/prometheus-stackdriver-exporter/values.yaml b/charts/prometheus-stackdriver-exporter/values.yaml
index 7bb17758..f2c5de62 100644
--- a/charts/prometheus-stackdriver-exporter/values.yaml
+++ b/charts/prometheus-stackdriver-exporter/values.yaml
@@ -15,6 +15,13 @@ image:
   tag: v0.12.0
   pullPolicy: IfNotPresent
 
+  ## Optionally specify an array of imagePullSecrets.
+  ## Secrets must be manually created in the namespace.
+  ## ref: https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/
+  ##
+  # pullSecrets:
+  #   - myDockerConfigJsonSecretName
+
 resources: {}
   # requests:
   #   cpu: 100m

```

## 2.0.0

**Release date:** 2022-02-15

![AppVersion: 0.12.0](https://img.shields.io/static/v1?label=AppVersion&message=0.12.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-stackdriver-exporter] Switch exporter config to agrs instead of env variables. (#1790)

### Default value changes

```diff
diff --git a/charts/prometheus-stackdriver-exporter/values.yaml b/charts/prometheus-stackdriver-exporter/values.yaml
index 4c220197..7bb17758 100644
--- a/charts/prometheus-stackdriver-exporter/values.yaml
+++ b/charts/prometheus-stackdriver-exporter/values.yaml
@@ -12,7 +12,7 @@ restartPolicy: Always
 
 image:
   repository: prometheuscommunity/stackdriver-exporter
-  tag: v0.11.0
+  tag: v0.12.0
   pullPolicy: IfNotPresent
 
 resources: {}

```

## 1.12.2

**Release date:** 2022-02-08

![AppVersion: 0.11.0](https://img.shields.io/static/v1?label=AppVersion&message=0.11.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-stackdriver-exporter] Fix servicemonitor name (#1777)

### Default value changes

```diff
# No changes in this release
```

## 1.12.1

**Release date:** 2021-12-07

![AppVersion: 0.11.0](https://img.shields.io/static/v1?label=AppVersion&message=0.11.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-stackdriver-exporter] fix chart name in source url. (#1569)

### Default value changes

```diff
# No changes in this release
```

## 1.12.0

**Release date:** 2021-12-02

![AppVersion: 0.11.0](https://img.shields.io/static/v1?label=AppVersion&message=0.11.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-stackdriver-exporter] add support for secret labels (#1553)

### Default value changes

```diff
diff --git a/charts/prometheus-stackdriver-exporter/values.yaml b/charts/prometheus-stackdriver-exporter/values.yaml
index 20116abf..4c220197 100644
--- a/charts/prometheus-stackdriver-exporter/values.yaml
+++ b/charts/prometheus-stackdriver-exporter/values.yaml
@@ -30,6 +30,9 @@ service:
   httpPort: 9255
   annotations: {}
 
+secret:
+  labels: {}
+
 stackdriver:
   # The Google Project ID to gather metrics for
   projectId: "FALSE"

```

## 1.11.0

**Release date:** 2021-10-07

![AppVersion: 0.11.0](https://img.shields.io/static/v1?label=AppVersion&message=0.11.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* allow using securityContext (#1402)

### Default value changes

```diff
diff --git a/charts/prometheus-stackdriver-exporter/values.yaml b/charts/prometheus-stackdriver-exporter/values.yaml
index 9f4be034..20116abf 100644
--- a/charts/prometheus-stackdriver-exporter/values.yaml
+++ b/charts/prometheus-stackdriver-exporter/values.yaml
@@ -23,6 +23,8 @@ resources: {}
   #   cpu: 100m
   #   memory: 128Mi
 
+securityContext: {}
+
 service:
   type: ClusterIP
   httpPort: 9255

```

## 1.10.0

**Release date:** 2021-09-07

![AppVersion: 0.11.0](https://img.shields.io/static/v1?label=AppVersion&message=0.11.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Add honorTimestamps config to stackdriver-exporter serviceMonitor (#1321)

### Default value changes

```diff
diff --git a/charts/prometheus-stackdriver-exporter/values.yaml b/charts/prometheus-stackdriver-exporter/values.yaml
index 90efe1ac..9f4be034 100644
--- a/charts/prometheus-stackdriver-exporter/values.yaml
+++ b/charts/prometheus-stackdriver-exporter/values.yaml
@@ -111,6 +111,8 @@ serviceMonitor:
   # interval: 10s
   ## Defaults to what's used if you follow CoreOS [Prometheus Install Instructions](https://github.com/helm/charts/tree/master/stable/prometheus-operator#tldr)
   honorLabels: true
+  ## Whether Prometheus should use the timestamps of the metrics exposed by stackdriver-exporter
+  honorTimestamps: true
   # MetricRelabelConfigs to apply to samples before ingestion https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
   metricRelabelings: []
   # RelabelConfigs to apply to samples before scraping. https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig

```

## 1.9.0

**Release date:** 2021-06-09

![AppVersion: 0.11.0](https://img.shields.io/static/v1?label=AppVersion&message=0.11.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-stackdriver-exporter] Add possiblity to pass extra args to the exporter (#1051)

### Default value changes

```diff
diff --git a/charts/prometheus-stackdriver-exporter/values.yaml b/charts/prometheus-stackdriver-exporter/values.yaml
index 944b4c06..90efe1ac 100644
--- a/charts/prometheus-stackdriver-exporter/values.yaml
+++ b/charts/prometheus-stackdriver-exporter/values.yaml
@@ -69,6 +69,10 @@ affinity: {}
 
 annotations: {}
 
+## Pod extra arguments
+##
+extraArgs: {}
+
 ## Node labels for stackdriver-exporter pod assignment
 ## Ref: https://kubernetes.io/docs/user-guide/node-selection/
 ##

```

## 1.8.3

**Release date:** 2021-05-29

![AppVersion: 0.11.0](https://img.shields.io/static/v1?label=AppVersion&message=0.11.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* add nameOverride and fullnameOverride to values.yaml (#1011)

### Default value changes

```diff
diff --git a/charts/prometheus-stackdriver-exporter/values.yaml b/charts/prometheus-stackdriver-exporter/values.yaml
index d615e3c7..944b4c06 100644
--- a/charts/prometheus-stackdriver-exporter/values.yaml
+++ b/charts/prometheus-stackdriver-exporter/values.yaml
@@ -1,3 +1,9 @@
+# Provide a name in place of prometheus-stackdriver-exporter for `app:` labels
+nameOverride: ""
+
+# Provide a name to substitute for the full names of resources
+fullnameOverride: ""
+
 # Number of exporters to run
 replicaCount: 1
 

```

## 1.8.2

**Release date:** 2021-03-04

![AppVersion: 0.11.0](https://img.shields.io/static/v1?label=AppVersion&message=0.11.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Moving scrapeTimeout to the correct location. (#728)

### Default value changes

```diff
diff --git a/charts/prometheus-stackdriver-exporter/values.yaml b/charts/prometheus-stackdriver-exporter/values.yaml
index a956dcba..d615e3c7 100644
--- a/charts/prometheus-stackdriver-exporter/values.yaml
+++ b/charts/prometheus-stackdriver-exporter/values.yaml
@@ -48,8 +48,6 @@ stackdriver:
     typePrefixes: 'compute.googleapis.com/instance/cpu'
     # The frequency to request
     interval: '5m'
-    # How long until a scrape request times out.
-    scrapeTimeout: '10s'
     # How far into the past to offset
     offset: '0s'
 
@@ -97,6 +95,8 @@ serviceMonitor:
   namespace: monitoring
   # additionalLabels is the set of additional labels to add to the ServiceMonitor
   additionalLabels: {}
+  # How long until a scrape request times out.
+  scrapeTimeout: '10s'
   # fallback to the prometheus default unless specified
   # interval: 10s
   ## Defaults to what's used if you follow CoreOS [Prometheus Install Instructions](https://github.com/helm/charts/tree/master/stable/prometheus-operator#tldr)

```

## 1.8.1

**Release date:** 2021-03-03

![AppVersion: 0.11.0](https://img.shields.io/static/v1?label=AppVersion&message=0.11.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Add feature for custom key. (#723)

### Default value changes

```diff
diff --git a/charts/prometheus-stackdriver-exporter/values.yaml b/charts/prometheus-stackdriver-exporter/values.yaml
index e974512b..a956dcba 100644
--- a/charts/prometheus-stackdriver-exporter/values.yaml
+++ b/charts/prometheus-stackdriver-exporter/values.yaml
@@ -27,6 +27,8 @@ stackdriver:
   projectId: "FALSE"
   # An existing secret which contains credentials.json
   serviceAccountSecret: ""
+  # Provide custom key for the existing secret to load credentials.json from
+  serviceAccountSecretKey: ""
   # A service account key JSON file. Must be provided when no existing secret is used, in this case a new secret will be created holding this service account
   serviceAccountKey: ""
   # Max number of retries that should be attempted on 503 errors from Stackdriver

```

## 1.8.0

**Release date:** 2021-01-28

![AppVersion: 0.11.0](https://img.shields.io/static/v1?label=AppVersion&message=0.11.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Using prometheus community docker image instead of deprecated version (#618)

### Default value changes

```diff
diff --git a/charts/prometheus-stackdriver-exporter/values.yaml b/charts/prometheus-stackdriver-exporter/values.yaml
index 5bcd02ca..e974512b 100644
--- a/charts/prometheus-stackdriver-exporter/values.yaml
+++ b/charts/prometheus-stackdriver-exporter/values.yaml
@@ -5,8 +5,8 @@ replicaCount: 1
 restartPolicy: Always
 
 image:
-  repository: frodenas/stackdriver-exporter
-  tag: v0.6.0
+  repository: prometheuscommunity/stackdriver-exporter
+  tag: v0.11.0
   pullPolicy: IfNotPresent
 
 resources: {}

```

## 1.7.0

**Release date:** 2021-01-22

![AppVersion: 0.6.0](https://img.shields.io/static/v1?label=AppVersion&message=0.6.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stackdriver-exporter] feat: Adding relabeling config to serviceMonitor (#605)

### Default value changes

```diff
diff --git a/charts/prometheus-stackdriver-exporter/values.yaml b/charts/prometheus-stackdriver-exporter/values.yaml
index 917b4882..5bcd02ca 100644
--- a/charts/prometheus-stackdriver-exporter/values.yaml
+++ b/charts/prometheus-stackdriver-exporter/values.yaml
@@ -99,3 +99,7 @@ serviceMonitor:
   # interval: 10s
   ## Defaults to what's used if you follow CoreOS [Prometheus Install Instructions](https://github.com/helm/charts/tree/master/stable/prometheus-operator#tldr)
   honorLabels: true
+  # MetricRelabelConfigs to apply to samples before ingestion https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
+  metricRelabelings: []
+  # RelabelConfigs to apply to samples before scraping. https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
+  relabelings: []

```

## 1.6.1

**Release date:** 2021-01-09

![AppVersion: 0.6.0](https://img.shields.io/static/v1?label=AppVersion&message=0.6.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* fix: scrapeTimeout value (#527)

### Default value changes

```diff
diff --git a/charts/prometheus-stackdriver-exporter/values.yaml b/charts/prometheus-stackdriver-exporter/values.yaml
index 55b2985d..917b4882 100644
--- a/charts/prometheus-stackdriver-exporter/values.yaml
+++ b/charts/prometheus-stackdriver-exporter/values.yaml
@@ -47,7 +47,7 @@ stackdriver:
     # The frequency to request
     interval: '5m'
     # How long until a scrape request times out.
-    scrapeTimeout: 10s
+    scrapeTimeout: '10s'
     # How far into the past to offset
     offset: '0s'
 

```

## 1.6.0

**Release date:** 2020-12-29

![AppVersion: 0.6.0](https://img.shields.io/static/v1?label=AppVersion&message=0.6.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stackdriver-exporter] add scrapeTimeout parameter (#522)

### Default value changes

```diff
diff --git a/charts/prometheus-stackdriver-exporter/values.yaml b/charts/prometheus-stackdriver-exporter/values.yaml
index f8eb9171..55b2985d 100644
--- a/charts/prometheus-stackdriver-exporter/values.yaml
+++ b/charts/prometheus-stackdriver-exporter/values.yaml
@@ -46,6 +46,8 @@ stackdriver:
     typePrefixes: 'compute.googleapis.com/instance/cpu'
     # The frequency to request
     interval: '5m'
+    # How long until a scrape request times out.
+    scrapeTimeout: 10s
     # How far into the past to offset
     offset: '0s'
 

```

## 1.5.0

**Release date:** 2020-11-10

![AppVersion: 0.6.0](https://img.shields.io/static/v1?label=AppVersion&message=0.6.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-stackdriver-exporter] Add dropDelegatedProjects to stackdriver-exporter (#296)
* Fix CI + add rpahli as maintainer
* Update charts/prometheus-stackdriver-exporter/README.md

### Default value changes

```diff
diff --git a/charts/prometheus-stackdriver-exporter/values.yaml b/charts/prometheus-stackdriver-exporter/values.yaml
index 601c0cdc..f8eb9171 100644
--- a/charts/prometheus-stackdriver-exporter/values.yaml
+++ b/charts/prometheus-stackdriver-exporter/values.yaml
@@ -39,6 +39,8 @@ stackdriver:
   backoffJitter: 1s
   # The HTTP statuses that should trigger a retry
   retryStatuses: 503
+  # Drop metrics from attached projects and fetch `project_id` only
+  dropDelegatedProjects: false
   metrics:
     # The prefixes to gather metrics for, we default to just CPU metrics.
     typePrefixes: 'compute.googleapis.com/instance/cpu'

```

## 1.4.0

**Release date:** 2020-10-06

![AppVersion: 0.6.0](https://img.shields.io/static/v1?label=AppVersion&message=0.6.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Bump Chart versions, adjust README

### Default value changes

```diff
# No changes in this release
```

## 1.3.0

**Release date:** 2020-08-23

![AppVersion: 0.6.0](https://img.shields.io/static/v1?label=AppVersion&message=0.6.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/stackdriver-exporter] Allow to specify service account annotations (#23182)

### Default value changes

```diff
diff --git a/charts/prometheus-stackdriver-exporter/values.yaml b/charts/prometheus-stackdriver-exporter/values.yaml
index 668ba02e..601c0cdc 100644
--- a/charts/prometheus-stackdriver-exporter/values.yaml
+++ b/charts/prometheus-stackdriver-exporter/values.yaml
@@ -83,6 +83,7 @@ serviceAccount:
   # If not set and create is false, 'default' is used
   # If not set and create is true, a name is generated using the fullname template
   name:
+  annotations: {}
 
 # Enable this if you're using https://github.com/coreos/prometheus-operator
 serviceMonitor:

```

## 1.2.3

**Release date:** 2020-06-06

![AppVersion: 0.6.0](https://img.shields.io/static/v1?label=AppVersion&message=0.6.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/stackdriver-exporter] update k8s deployement api version (#21720)

### Default value changes

```diff
# No changes in this release
```

## 1.2.2

**Release date:** 2019-12-30

![AppVersion: 0.6.0](https://img.shields.io/static/v1?label=AppVersion&message=0.6.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/stackdriver-exporter] - Add prometheus rule additionalLabels in servicemonitor.yaml  (#19802)

### Default value changes

```diff
diff --git a/charts/prometheus-stackdriver-exporter/values.yaml b/charts/prometheus-stackdriver-exporter/values.yaml
index 75d6d261..668ba02e 100644
--- a/charts/prometheus-stackdriver-exporter/values.yaml
+++ b/charts/prometheus-stackdriver-exporter/values.yaml
@@ -88,6 +88,8 @@ serviceAccount:
 serviceMonitor:
   enabled: false
   namespace: monitoring
+  # additionalLabels is the set of additional labels to add to the ServiceMonitor
+  additionalLabels: {}
   # fallback to the prometheus default unless specified
   # interval: 10s
   ## Defaults to what's used if you follow CoreOS [Prometheus Install Instructions](https://github.com/helm/charts/tree/master/stable/prometheus-operator#tldr)

```

## 1.2.1

**Release date:** 2019-09-04

![AppVersion: 0.6.0](https://img.shields.io/static/v1?label=AppVersion&message=0.6.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/stackdriver-exporter] Correct servicemonitor name (#16750)

### Default value changes

```diff
# No changes in this release
```

## 1.2.0

**Release date:** 2019-09-01

![AppVersion: 0.6.0](https://img.shields.io/static/v1?label=AppVersion&message=0.6.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/stackdriver-exporter] Add support for Prometheus Operator ServiceMonitor (#16372)

### Default value changes

```diff
diff --git a/charts/prometheus-stackdriver-exporter/values.yaml b/charts/prometheus-stackdriver-exporter/values.yaml
index cb91ccf8..75d6d261 100644
--- a/charts/prometheus-stackdriver-exporter/values.yaml
+++ b/charts/prometheus-stackdriver-exporter/values.yaml
@@ -83,3 +83,12 @@ serviceAccount:
   # If not set and create is false, 'default' is used
   # If not set and create is true, a name is generated using the fullname template
   name:
+
+# Enable this if you're using https://github.com/coreos/prometheus-operator
+serviceMonitor:
+  enabled: false
+  namespace: monitoring
+  # fallback to the prometheus default unless specified
+  # interval: 10s
+  ## Defaults to what's used if you follow CoreOS [Prometheus Install Instructions](https://github.com/helm/charts/tree/master/stable/prometheus-operator#tldr)
+  honorLabels: true

```

## 1.1.1

**Release date:** 2019-07-29

![AppVersion: 0.6.0](https://img.shields.io/static/v1?label=AppVersion&message=0.6.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* stable/stackdriver-exporter: Allow setting of pod ServiceAccount name (#15931)

### Default value changes

```diff
diff --git a/charts/prometheus-stackdriver-exporter/values.yaml b/charts/prometheus-stackdriver-exporter/values.yaml
index ec3b8d9a..cb91ccf8 100644
--- a/charts/prometheus-stackdriver-exporter/values.yaml
+++ b/charts/prometheus-stackdriver-exporter/values.yaml
@@ -72,3 +72,14 @@ tolerations: []
   #   operator: "Equal|Exists"
   #   value: "value"
   #   effect: "NoSchedule|PreferNoSchedule|NoExecute(1.6 only)"
+
+
+## Service Account
+##
+serviceAccount:
+  # Specifies whether a ServiceAccount should be created
+  create: false
+  # The name of the ServiceAccount to use.
+  # If not set and create is false, 'default' is used
+  # If not set and create is true, a name is generated using the fullname template
+  name:

```

## 1.1.0

**Release date:** 2019-04-26

![AppVersion: 0.6.0](https://img.shields.io/static/v1?label=AppVersion&message=0.6.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* stackdriver-exporter: allow google service account (#13214)

### Default value changes

```diff
diff --git a/charts/prometheus-stackdriver-exporter/values.yaml b/charts/prometheus-stackdriver-exporter/values.yaml
index 06612e83..ec3b8d9a 100644
--- a/charts/prometheus-stackdriver-exporter/values.yaml
+++ b/charts/prometheus-stackdriver-exporter/values.yaml
@@ -25,6 +25,10 @@ service:
 stackdriver:
   # The Google Project ID to gather metrics for
   projectId: "FALSE"
+  # An existing secret which contains credentials.json
+  serviceAccountSecret: ""
+  # A service account key JSON file. Must be provided when no existing secret is used, in this case a new secret will be created holding this service account
+  serviceAccountKey: ""
   # Max number of retries that should be attempted on 503 errors from Stackdriver
   maxRetries: 0
   # How long should Stackdriver_exporter wait for a result from the Stackdriver API

```

## 1.0.0

**Release date:** 2019-03-05

![AppVersion: 0.6.0](https://img.shields.io/static/v1?label=AppVersion&message=0.6.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* add tolerations, affinity and nodeSelector to stackdriver-exporter (#11854)

### Default value changes

```diff
diff --git a/charts/prometheus-stackdriver-exporter/values.yaml b/charts/prometheus-stackdriver-exporter/values.yaml
index 931b02aa..06612e83 100644
--- a/charts/prometheus-stackdriver-exporter/values.yaml
+++ b/charts/prometheus-stackdriver-exporter/values.yaml
@@ -49,4 +49,22 @@ web:
   # Path under which to expose metrics.
   path: /metrics
 
+## Pod affinity
+##
+affinity: {}
+
 annotations: {}
+
+## Node labels for stackdriver-exporter pod assignment
+## Ref: https://kubernetes.io/docs/user-guide/node-selection/
+##
+nodeSelector: {}
+
+## Node tolerations for stackdriver-exporter scheduling to nodes with taints
+## Ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
+##
+tolerations: []
+  # - key: "key"
+  #   operator: "Equal|Exists"
+  #   value: "value"
+  #   effect: "NoSchedule|PreferNoSchedule|NoExecute(1.6 only)"

```

## 0.0.6

**Release date:** 2019-01-02

![AppVersion: 0.6.0](https://img.shields.io/static/v1?label=AppVersion&message=0.6.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/stackdriver-exporter] fix duplicated 'livenessProbe', add 'readinessProbe' (#10321)

### Default value changes

```diff
diff --git a/charts/prometheus-stackdriver-exporter/values.yaml b/charts/prometheus-stackdriver-exporter/values.yaml
index 07cc92ee..931b02aa 100644
--- a/charts/prometheus-stackdriver-exporter/values.yaml
+++ b/charts/prometheus-stackdriver-exporter/values.yaml
@@ -25,13 +25,13 @@ service:
 stackdriver:
   # The Google Project ID to gather metrics for
   projectId: "FALSE"
-  # Max number of retries that should be attempted on 503 errors from stackdriver
+  # Max number of retries that should be attempted on 503 errors from Stackdriver
   maxRetries: 0
-  # How long should stackdriver_exporter wait for a result from the Stackdriver API
+  # How long should Stackdriver_exporter wait for a result from the Stackdriver API
   httpTimeout: 10s
   # Max time between each request in an exp backoff scenario
   maxBackoff: 5s
-  # The amount of jitter to introduce in a exp backoff scenario
+  # The amount of jitter to introduce in an exp backoff scenario
   backoffJitter: 1s
   # The HTTP statuses that should trigger a retry
   retryStatuses: 503

```

## 0.0.5

**Release date:** 2018-12-19

![AppVersion: 0.6.0](https://img.shields.io/static/v1?label=AppVersion&message=0.6.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* update stackdriver-exporter to 0.6.0 with new options (#10142)

### Default value changes

```diff
diff --git a/charts/prometheus-stackdriver-exporter/values.yaml b/charts/prometheus-stackdriver-exporter/values.yaml
index 9e75bd3a..07cc92ee 100644
--- a/charts/prometheus-stackdriver-exporter/values.yaml
+++ b/charts/prometheus-stackdriver-exporter/values.yaml
@@ -6,7 +6,7 @@ restartPolicy: Always
 
 image:
   repository: frodenas/stackdriver-exporter
-  tag: v0.5.1
+  tag: v0.6.0
   pullPolicy: IfNotPresent
 
 resources: {}
@@ -25,6 +25,16 @@ service:
 stackdriver:
   # The Google Project ID to gather metrics for
   projectId: "FALSE"
+  # Max number of retries that should be attempted on 503 errors from stackdriver
+  maxRetries: 0
+  # How long should stackdriver_exporter wait for a result from the Stackdriver API
+  httpTimeout: 10s
+  # Max time between each request in an exp backoff scenario
+  maxBackoff: 5s
+  # The amount of jitter to introduce in a exp backoff scenario
+  backoffJitter: 1s
+  # The HTTP statuses that should trigger a retry
+  retryStatuses: 503
   metrics:
     # The prefixes to gather metrics for, we default to just CPU metrics.
     typePrefixes: 'compute.googleapis.com/instance/cpu'

```

## 0.0.4

**Release date:** 2018-04-27

![AppVersion: 0.5.1](https://img.shields.io/static/v1?label=AppVersion&message=0.5.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* adding annotations (#5170)

### Default value changes

```diff
diff --git a/charts/prometheus-stackdriver-exporter/values.yaml b/charts/prometheus-stackdriver-exporter/values.yaml
index 61c906ea..9e75bd3a 100644
--- a/charts/prometheus-stackdriver-exporter/values.yaml
+++ b/charts/prometheus-stackdriver-exporter/values.yaml
@@ -38,3 +38,5 @@ web:
   listenAddress: ':9255'
   # Path under which to expose metrics.
   path: /metrics
+
+annotations: {}

```

## 0.0.3

**Release date:** 2018-04-18

![AppVersion: 0.5.1](https://img.shields.io/static/v1?label=AppVersion&message=0.5.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Typo fix in stackdriver-exporter/README.md: "for this to exporter to work"->"for this exporter to work" (#4940)

### Default value changes

```diff
# No changes in this release
```

## 0.0.2

**Release date:** 2018-04-05

![AppVersion: 0.5.1](https://img.shields.io/static/v1?label=AppVersion&message=0.5.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Add home key for stackdriver-exporter (#4626)

### Default value changes

```diff
# No changes in this release
```

## 0.0.1

**Release date:** 2018-03-27

![AppVersion: 0.5.1](https://img.shields.io/static/v1?label=AppVersion&message=0.5.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Add stackdriver-exporter (#4167)

### Default value changes

```diff
# Number of exporters to run
replicaCount: 1

# Restart policy for container
restartPolicy: Always

image:
  repository: frodenas/stackdriver-exporter
  tag: v0.5.1
  pullPolicy: IfNotPresent

resources: {}
  # requests:
  #   cpu: 100m
  #   memory: 128Mi
  # limits:
  #   cpu: 100m
  #   memory: 128Mi

service:
  type: ClusterIP
  httpPort: 9255
  annotations: {}

stackdriver:
  # The Google Project ID to gather metrics for
  projectId: "FALSE"
  metrics:
    # The prefixes to gather metrics for, we default to just CPU metrics.
    typePrefixes: 'compute.googleapis.com/instance/cpu'
    # The frequency to request
    interval: '5m'
    # How far into the past to offset
    offset: '0s'

web:
  # Port to listen on
  listenAddress: ':9255'
  # Path under which to expose metrics.
  path: /metrics

```

---
Autogenerated from Helm Chart and git history using [helm-changelog](https://github.com/mogensen/helm-changelog)
