# Change Log

## 2.6.0 

**Release date:** 2021-03-05

![AppVersion: 0.7.0](https://img.shields.io/static/v1?label=AppVersion&message=0.7.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-nats-exporter] Bump prometheus-nats-expoter app version (#722) 

### Default value changes

```diff
diff --git a/charts/prometheus-nats-exporter/values.yaml b/charts/prometheus-nats-exporter/values.yaml
index 7de0f86..7ab7e94 100644
--- a/charts/prometheus-nats-exporter/values.yaml
+++ b/charts/prometheus-nats-exporter/values.yaml
@@ -5,8 +5,8 @@
 replicaCount: 1
 
 image:
-  repository: synadia/prometheus-nats-exporter
-  tag: 0.6.2
+  repository: natsio/prometheus-nats-exporter
+  tag: 0.7.0
   pullPolicy: IfNotPresent
 
 service:
```

## 2.5.1 

**Release date:** 2020-08-20

![AppVersion: 0.6.2](https://img.shields.io/static/v1?label=AppVersion&message=0.6.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Prep initial charts indexing (#14) 

### Default value changes

```diff
diff --git a/charts/prometheus-nats-exporter/values.yaml b/charts/prometheus-nats-exporter/values.yaml
index bd26c22..7de0f86 100644
--- a/charts/prometheus-nats-exporter/values.yaml
+++ b/charts/prometheus-nats-exporter/values.yaml
@@ -35,6 +35,7 @@ resources: {}
 
 config:
   nats:
+    # See https://github.com/helm/charts/blob/master/stable/nats/templates/monitoring-svc.yaml
     service: nats-nats-monitoring
     namespace: default
     port: 8222
```

## 2.5.0 

**Release date:** 2020-04-29

![AppVersion: 0.6.2](https://img.shields.io/static/v1?label=AppVersion&message=0.6.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-nats-exporter] Update prometheus-nats-exporter version (#22178) 

### Default value changes

```diff
diff --git a/charts/prometheus-nats-exporter/values.yaml b/charts/prometheus-nats-exporter/values.yaml
index 98325ea..bd26c22 100644
--- a/charts/prometheus-nats-exporter/values.yaml
+++ b/charts/prometheus-nats-exporter/values.yaml
@@ -6,7 +6,7 @@ replicaCount: 1
 
 image:
   repository: synadia/prometheus-nats-exporter
-  tag: 0.6.0
+  tag: 0.6.2
   pullPolicy: IfNotPresent
 
 service:
@@ -45,6 +45,7 @@ config:
     routez: true
     serverz: true
     subz: true
+    gatewayz: true
 
 nodeSelector: {}
 
```

## 2.4.0 

**Release date:** 2020-02-28

![AppVersion: 0.6.0](https://img.shields.io/static/v1?label=AppVersion&message=0.6.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-nats-exporter]: add prometheus-operator servicemonitor support (#21088) 

### Default value changes

```diff
diff --git a/charts/prometheus-nats-exporter/values.yaml b/charts/prometheus-nats-exporter/values.yaml
index 3f486e4..98325ea 100644
--- a/charts/prometheus-nats-exporter/values.yaml
+++ b/charts/prometheus-nats-exporter/values.yaml
@@ -14,6 +14,13 @@ service:
   port: 80
   targetPort: 7777
 
+serviceMonitor:
+  enabled: false
+  additionalLabels: {}
+  namespace:
+  interval:
+  scrapeTimeout:
+
 resources: {}
   # We usually recommend not to specify default resources and to leave this as a conscious
   # choice for the user. This also increases chances charts run on environments with little
```

## 2.3.0 

**Release date:** 2019-10-12

![AppVersion: 0.6.0](https://img.shields.io/static/v1?label=AppVersion&message=0.6.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-nats-exporter]: upgrade (#17085) 

### Default value changes

```diff
diff --git a/charts/prometheus-nats-exporter/values.yaml b/charts/prometheus-nats-exporter/values.yaml
index 4f3f7b1..3f486e4 100644
--- a/charts/prometheus-nats-exporter/values.yaml
+++ b/charts/prometheus-nats-exporter/values.yaml
@@ -6,7 +6,7 @@ replicaCount: 1
 
 image:
   repository: synadia/prometheus-nats-exporter
-  tag: 0.5.0
+  tag: 0.6.0
   pullPolicy: IfNotPresent
 
 service:
```

## 2.2.1 

**Release date:** 2019-08-04

![AppVersion: 0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=0.5.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* bump chart version; fix hardcoded cluster dns name (#16075) 

### Default value changes

```diff
# No changes in this release
```

## 2.2.0 

**Release date:** 2019-07-27

![AppVersion: 0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=0.5.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Updated prometheus-nats-exporter to 0.5.0 (#15921) 

### Default value changes

```diff
diff --git a/charts/prometheus-nats-exporter/values.yaml b/charts/prometheus-nats-exporter/values.yaml
index a0ef8ed..4f3f7b1 100644
--- a/charts/prometheus-nats-exporter/values.yaml
+++ b/charts/prometheus-nats-exporter/values.yaml
@@ -6,7 +6,7 @@ replicaCount: 1
 
 image:
   repository: synadia/prometheus-nats-exporter
-  tag: 0.4.0
+  tag: 0.5.0
   pullPolicy: IfNotPresent
 
 service:
```

## 2.1.0 

**Release date:** 2019-06-25

![AppVersion: 0.4.0](https://img.shields.io/static/v1?label=AppVersion&message=0.4.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Update prometheus-nats-exporter to version 0.4.0 (#15048) 

### Default value changes

```diff
diff --git a/charts/prometheus-nats-exporter/values.yaml b/charts/prometheus-nats-exporter/values.yaml
index d4a0f91..a0ef8ed 100644
--- a/charts/prometheus-nats-exporter/values.yaml
+++ b/charts/prometheus-nats-exporter/values.yaml
@@ -6,7 +6,7 @@ replicaCount: 1
 
 image:
   repository: synadia/prometheus-nats-exporter
-  tag: 0.3.0
+  tag: 0.4.0
   pullPolicy: IfNotPresent
 
 service:
```

## 2.0.0 

**Release date:** 2019-06-09

![AppVersion: 0.3.0](https://img.shields.io/static/v1?label=AppVersion&message=0.3.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-nats-exporter] upgrade to official image (#14528) 

### Default value changes

```diff
diff --git a/charts/prometheus-nats-exporter/values.yaml b/charts/prometheus-nats-exporter/values.yaml
index 90e4a8f..d4a0f91 100644
--- a/charts/prometheus-nats-exporter/values.yaml
+++ b/charts/prometheus-nats-exporter/values.yaml
@@ -5,14 +5,14 @@
 replicaCount: 1
 
 image:
-  repository: appcelerator/prometheus-nats-exporter
-  tag: 0.17.0
+  repository: synadia/prometheus-nats-exporter
+  tag: 0.3.0
   pullPolicy: IfNotPresent
 
 service:
   type: ClusterIP
   port: 80
-  targetPort: 8222
+  targetPort: 7777
 
 resources: {}
   # We usually recommend not to specify default resources and to leave this as a conscious
@@ -31,6 +31,13 @@ config:
     service: nats-nats-monitoring
     namespace: default
     port: 8222
+  metrics:
+    varz: true
+    channelz: true
+    connz: true
+    routez: true
+    serverz: true
+    subz: true
 
 nodeSelector: {}
 
```

## 1.0.1 

**Release date:** 2019-06-05

![AppVersion: 0.17.0](https://img.shields.io/static/v1?label=AppVersion&message=0.17.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* added me to prometheus-nats-exporter/OWNERS (#14525) 

### Default value changes

```diff
# No changes in this release
```

## 1.0.0 

**Release date:** 2019-03-14

![AppVersion: 0.17.0](https://img.shields.io/static/v1?label=AppVersion&message=0.17.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Prometheus NATS exporter (#10742) 

### Default value changes

```diff
# Default values for prometheus-nats-exporter.
# This is a YAML-formatted file.
# Declare variables to be passed into your templates.

replicaCount: 1

image:
  repository: appcelerator/prometheus-nats-exporter
  tag: 0.17.0
  pullPolicy: IfNotPresent

service:
  type: ClusterIP
  port: 80
  targetPort: 8222

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

config:
  nats:
    service: nats-nats-monitoring
    namespace: default
    port: 8222

nodeSelector: {}

tolerations: []

affinity: {}

annotations: {}

extraContainers: |

extraVolumes: |
```

---
Autogenerated from Helm Chart and git history using [helm-changelog](https://github.com/mogensen/helm-changelog)
