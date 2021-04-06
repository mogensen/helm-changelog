# Change Log

## 1.0.0 

**Release date:** 2021-04-06

![AppVersion: v0.29.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.29.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-rabbitmq-exporter] Polish up Prometheus rules (#724) 

### Default value changes

```diff
# No changes in this release
```

## 0.7.0 

**Release date:** 2021-03-19

![AppVersion: v0.29.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.29.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-rabbitmq-exporter] update default rabbitmq secret key (#778) 
* [prometheus-rabbitmq-exporter] Make it possible to add podLabels and priorityClass  (#513) 

### Default value changes

```diff
diff --git a/charts/prometheus-rabbitmq-exporter/values.yaml b/charts/prometheus-rabbitmq-exporter/values.yaml
index b2b7a0f..30b7b8f 100644
--- a/charts/prometheus-rabbitmq-exporter/values.yaml
+++ b/charts/prometheus-rabbitmq-exporter/values.yaml
@@ -37,6 +37,7 @@ rabbitmq:
   password: guest
   # If existingPasswordSecret is set then password is ignored
   existingPasswordSecret: ~
+  existingPasswordSecretKey: password
   capabilities: bert,no_sort
   include_queues: ".*"
   include_vhost: ".*"
@@ -48,6 +49,8 @@ rabbitmq:
   timeout: 30
   max_queues: 0
 
+podLabels: {}
+
 annotations: {}
 #  prometheus.io/scrape: "true"
 #  prometheus.io/path: "/metrics"
```

## 0.6.0 

**Release date:** 2021-01-16

![AppVersion: v0.29.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.29.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-rabbitmq-exporter] Add the possibility to specify a priority class (#416) 

### Default value changes

```diff
diff --git a/charts/prometheus-rabbitmq-exporter/values.yaml b/charts/prometheus-rabbitmq-exporter/values.yaml
index 33ea475..b2b7a0f 100644
--- a/charts/prometheus-rabbitmq-exporter/values.yaml
+++ b/charts/prometheus-rabbitmq-exporter/values.yaml
@@ -22,6 +22,8 @@ resources: {}
   #  cpu: 100m
   #  memory: 128Mi
 
+priorityClassName: ""
+
 nodeSelector: {}
 
 tolerations: []
```

## 0.5.8 

**Release date:** 2021-01-15

![AppVersion: v0.29.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.29.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-rabbitmq-exporter] Add desaintmartin & monotek as maintainers (#571) 

### Default value changes

```diff
# No changes in this release
```

## 0.5.7 

**Release date:** 2021-01-13

![AppVersion: v0.29.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.29.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-rabbitmq-exporter] Updated rabbitmq exporter maintainer list (#282) 

### Default value changes

```diff
# No changes in this release
```

## 0.5.6 

**Release date:** 2020-08-20

![AppVersion: v0.29.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.29.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Prep initial charts indexing (#14) 

### Default value changes

```diff
# No changes in this release
```

## 0.5.5 

**Release date:** 2019-12-02

![AppVersion: v0.29.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.29.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-rabbitmq-exporter] Fix ServiceMonitor + adds alerts (#19276) 

### Default value changes

```diff
diff --git a/charts/prometheus-rabbitmq-exporter/values.yaml b/charts/prometheus-rabbitmq-exporter/values.yaml
index f68e5f7..33ea475 100644
--- a/charts/prometheus-rabbitmq-exporter/values.yaml
+++ b/charts/prometheus-rabbitmq-exporter/values.yaml
@@ -57,3 +57,6 @@ prometheus:
     additionalLabels: {}
     interval: 15s
     namespace: []
+  rules:
+    enabled: false
+    additionalLabels: {}
```

## 0.5.3 

**Release date:** 2019-09-15

![AppVersion: v0.29.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.29.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-rabbitmq-exporter] ServiceMonitor Port Bugfix (#15801) 

### Default value changes

```diff
# No changes in this release
```

## 0.5.2 

**Release date:** 2019-08-13

![AppVersion: v0.29.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.29.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Rabbitmq prometheus exporter should use the same name for the port, otherwise servicemonitor will not scrape successfully (#16134) 

### Default value changes

```diff
# No changes in this release
```

## 0.5.1 

**Release date:** 2019-06-22

![AppVersion: v0.29.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.29.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Added support for existing rabbitmq password secret (#14958) 

### Default value changes

```diff
diff --git a/charts/prometheus-rabbitmq-exporter/values.yaml b/charts/prometheus-rabbitmq-exporter/values.yaml
index 596c45b..f68e5f7 100644
--- a/charts/prometheus-rabbitmq-exporter/values.yaml
+++ b/charts/prometheus-rabbitmq-exporter/values.yaml
@@ -33,6 +33,8 @@ rabbitmq:
   url: http://myrabbit:15672
   user: guest
   password: guest
+  # If existingPasswordSecret is set then password is ignored
+  existingPasswordSecret: ~
   capabilities: bert,no_sort
   include_queues: ".*"
   include_vhost: ".*"
```

## 0.5.0 

**Release date:** 2019-06-04

![AppVersion: v0.29.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.29.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-rabbitmq-exporter] Adding Prometheus Operator Service Monitor (#12723) 

### Default value changes

```diff
diff --git a/charts/prometheus-rabbitmq-exporter/values.yaml b/charts/prometheus-rabbitmq-exporter/values.yaml
index 5429c20..596c45b 100644
--- a/charts/prometheus-rabbitmq-exporter/values.yaml
+++ b/charts/prometheus-rabbitmq-exporter/values.yaml
@@ -47,4 +47,11 @@ rabbitmq:
 annotations: {}
 #  prometheus.io/scrape: "true"
 #  prometheus.io/path: "/metrics"
-#  prometheus.io/port: "9419"
+#  prometheus.io/port: 9419
+
+prometheus:
+  monitor:
+    enabled: false
+    additionalLabels: {}
+    interval: 15s
+    namespace: []
```

## 0.4.1 

**Release date:** 2019-04-29

![AppVersion: v0.29.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.29.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-rabbitmq-exporter] Fix default parameters in readme  (#12858) 

### Default value changes

```diff
diff --git a/charts/prometheus-rabbitmq-exporter/values.yaml b/charts/prometheus-rabbitmq-exporter/values.yaml
index ad98034..5429c20 100644
--- a/charts/prometheus-rabbitmq-exporter/values.yaml
+++ b/charts/prometheus-rabbitmq-exporter/values.yaml
@@ -47,4 +47,4 @@ rabbitmq:
 annotations: {}
 #  prometheus.io/scrape: "true"
 #  prometheus.io/path: "/metrics"
-#  prometheus.io/port: 9419
+#  prometheus.io/port: "9419"
```

## 0.4.0 

**Release date:** 2019-02-08

![AppVersion: v0.29.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.29.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-rabbitmq-export] Add skip_verify to rabbitmq-exporter (#11107) 

### Default value changes

```diff
diff --git a/charts/prometheus-rabbitmq-exporter/values.yaml b/charts/prometheus-rabbitmq-exporter/values.yaml
index 9eb82a5..ad98034 100644
--- a/charts/prometheus-rabbitmq-exporter/values.yaml
+++ b/charts/prometheus-rabbitmq-exporter/values.yaml
@@ -37,6 +37,7 @@ rabbitmq:
   include_queues: ".*"
   include_vhost: ".*"
   skip_queues: "^$"
+  skip_verify: "false"
   skip_vhost: "^$"
   exporters: "exchange,node,overview,queue"
   output_format: "TTY"
```

## 0.3.0 

**Release date:** 2019-01-09

![AppVersion: v0.29.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.29.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Rabbitmq exporter env vars (#9244) 

### Default value changes

```diff
diff --git a/charts/prometheus-rabbitmq-exporter/values.yaml b/charts/prometheus-rabbitmq-exporter/values.yaml
index d75baf7..9eb82a5 100644
--- a/charts/prometheus-rabbitmq-exporter/values.yaml
+++ b/charts/prometheus-rabbitmq-exporter/values.yaml
@@ -38,6 +38,10 @@ rabbitmq:
   include_vhost: ".*"
   skip_queues: "^$"
   skip_vhost: "^$"
+  exporters: "exchange,node,overview,queue"
+  output_format: "TTY"
+  timeout: 30
+  max_queues: 0
 
 annotations: {}
 #  prometheus.io/scrape: "true"
```

## 0.2.0 

**Release date:** 2019-01-04

![AppVersion: v0.29.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.29.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-rabbitmq-exporter] Bump to v0.29.0 (#10273) 

### Default value changes

```diff
diff --git a/charts/prometheus-rabbitmq-exporter/values.yaml b/charts/prometheus-rabbitmq-exporter/values.yaml
index f7bf6ce..d75baf7 100644
--- a/charts/prometheus-rabbitmq-exporter/values.yaml
+++ b/charts/prometheus-rabbitmq-exporter/values.yaml
@@ -4,7 +4,7 @@
 replicaCount: 1
 image:
   repository: kbudde/rabbitmq-exporter
-  tag: v0.28.0
+  tag: v0.29.0
   pullPolicy: IfNotPresent
 service:
   type: ClusterIP
@@ -35,7 +35,9 @@ rabbitmq:
   password: guest
   capabilities: bert,no_sort
   include_queues: ".*"
+  include_vhost: ".*"
   skip_queues: "^$"
+  skip_vhost: "^$"
 
 annotations: {}
 #  prometheus.io/scrape: "true"
```

## 0.1.4 

**Release date:** 2018-08-15

![AppVersion: v0.28.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.28.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-rabbitmq-exporter] fix values (#7078) 

### Default value changes

```diff
diff --git a/charts/prometheus-rabbitmq-exporter/values.yaml b/charts/prometheus-rabbitmq-exporter/values.yaml
index e18c062..f7bf6ce 100644
--- a/charts/prometheus-rabbitmq-exporter/values.yaml
+++ b/charts/prometheus-rabbitmq-exporter/values.yaml
@@ -37,7 +37,7 @@ rabbitmq:
   include_queues: ".*"
   skip_queues: "^$"
 
-annotation: {}
+annotations: {}
 #  prometheus.io/scrape: "true"
 #  prometheus.io/path: "/metrics"
 #  prometheus.io/port: 9419
```

## 0.1.3 

**Release date:** 2018-05-25

![AppVersion: v0.28.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.28.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-rabbitmq-exporter] Typo fix: tables lists->table lists (#5727) 

### Default value changes

```diff
# No changes in this release
```

## 0.1.2 

**Release date:** 2018-05-23

![AppVersion: v0.28.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.28.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* fix typo in prometheus-rabbitmq-exporter's readme (#5666) 

### Default value changes

```diff
# No changes in this release
```

## 0.1.1 

**Release date:** 2018-05-16

![AppVersion: v0.28.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.28.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Added prometheus-rabbitmq-exporter helm chart (#5311) 

### Default value changes

```diff
# Default values for prometheus-rabbitmq-exporter.
# This is a YAML-formatted file.
# Declare variables to be passed into your templates.
replicaCount: 1
image:
  repository: kbudde/rabbitmq-exporter
  tag: v0.28.0
  pullPolicy: IfNotPresent
service:
  type: ClusterIP
  externalPort: 9419
  internalPort: 9419
resources: {}
  # We usually recommend not to specify default resources and to leave this as a conscious
  # choice for the user. This also increases chances charts run on environments with little
  # resources, such as Minikube. If you do want to specify resources, uncomment the following
  # lines, adjust them as necessary, and remove the curly braces after 'resources:'.
  # limits:
  #  cpu: 100m
  #  memory: 128Mi
  # requests:
  #  cpu: 100m
  #  memory: 128Mi

nodeSelector: {}

tolerations: []

affinity: {}

loglevel: info
rabbitmq:
  url: http://myrabbit:15672
  user: guest
  password: guest
  capabilities: bert,no_sort
  include_queues: ".*"
  skip_queues: "^$"

annotation: {}
#  prometheus.io/scrape: "true"
#  prometheus.io/path: "/metrics"
#  prometheus.io/port: 9419
```

---
Autogenerated from Helm Chart and git history using [helm-changelog](https://github.com/mogensen/helm-changelog)
