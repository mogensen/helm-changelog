# Change Log

## 3.1.3

**Release date:** 2023-05-02

![AppVersion: 0.31.0](https://img.shields.io/static/v1?label=AppVersion&message=0.31.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-mongodb-exporter] Add a maintainer (#3316)

### Default value changes

```diff
# No changes in this release
```

## 3.1.2

**Release date:** 2022-10-15

![AppVersion: 0.31.0](https://img.shields.io/static/v1?label=AppVersion&message=0.31.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-mongodb-exporter] Update ServiceMonitor documentation (#2563)

### Default value changes

```diff
# No changes in this release
```

## 3.1.1

**Release date:** 2022-08-26

![AppVersion: 0.31.0](https://img.shields.io/static/v1?label=AppVersion&message=0.31.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-mongodb-exporter] fixing default mongodb uri (#2379)

### Default value changes

```diff
diff --git a/charts/prometheus-mongodb-exporter/values.yaml b/charts/prometheus-mongodb-exporter/values.yaml
index 21561244..9ee7f8d3 100644
--- a/charts/prometheus-mongodb-exporter/values.yaml
+++ b/charts/prometheus-mongodb-exporter/values.yaml
@@ -22,7 +22,7 @@ livenessProbe:
 
 # [mongodb[+srv]://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]
 mongodb:
-  uri: "mongodb://monogdb:27017"
+  uri: "mongodb://mongodb:27017"
 
 # Name of an externally managed secret (in the same namespace) containing the connection uri as key `mongodb-uri`.
 # If this is provided, the value mongodb.uri is ignored.

```

## 3.1.0

**Release date:** 2022-06-18

![AppVersion: 0.31.0](https://img.shields.io/static/v1?label=AppVersion&message=0.31.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Add tls configuration option for service monitor and add an argument to rename service port name (#2170)

### Default value changes

```diff
diff --git a/charts/prometheus-mongodb-exporter/values.yaml b/charts/prometheus-mongodb-exporter/values.yaml
index b5d7272c..21561244 100644
--- a/charts/prometheus-mongodb-exporter/values.yaml
+++ b/charts/prometheus-mongodb-exporter/values.yaml
@@ -83,6 +83,7 @@ service:
   annotations: {}
   port: 9216
   type: ClusterIP
+  portName: metrics
 
 serviceAccount:
   create: true
@@ -98,5 +99,7 @@ serviceMonitor:
   additionalLabels: {}
   targetLabels: []
   metricRelabelings: []
+  scheme: ""
+  tlsConfig: {}
 
 tolerations: []

```

## 3.0.0

**Release date:** 2022-06-10

![AppVersion: 0.31.0](https://img.shields.io/static/v1?label=AppVersion&message=0.31.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* update mongo exporter (#2108)

### Default value changes

```diff
diff --git a/charts/prometheus-mongodb-exporter/values.yaml b/charts/prometheus-mongodb-exporter/values.yaml
index e7417846..b5d7272c 100644
--- a/charts/prometheus-mongodb-exporter/values.yaml
+++ b/charts/prometheus-mongodb-exporter/values.yaml
@@ -3,18 +3,14 @@ affinity: {}
 annotations: {}
 
 extraArgs:
-- --collect.collection
-- --collect.database
-- --collect.indexusage
-- --collect.topmetrics
-- --collect.connpoolstats
+  - --collect-all
 
 fullnameOverride: ""
 
 image:
   pullPolicy: IfNotPresent
-  repository: ssheehy/mongodb-exporter
-  tag: 0.11.0
+  repository: percona/mongodb_exporter
+  tag: ""
 
 imagePullSecrets: []
 
@@ -26,7 +22,7 @@ livenessProbe:
 
 # [mongodb[+srv]://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]
 mongodb:
-  uri: ""
+  uri: "mongodb://monogdb:27017"
 
 # Name of an externally managed secret (in the same namespace) containing the connection uri as key `mongodb-uri`.
 # If this is provided, the value mongodb.uri is ignored.
@@ -57,12 +53,12 @@ readinessProbe:
 replicas: 1
 
 resources: {}
-# limits:
-#   cpu: 250m
-#   memory: 192Mi
-# requests:
-#   cpu: 100m
-#   memory: 128Mi
+  # limits:
+  #   cpu: 250m
+  #   memory: 192Mi
+  # requests:
+  #   cpu: 100m
+  #   memory: 128Mi
 
 # Extra environment variables that will be passed into the exporter pod
 env: {}
@@ -95,7 +91,7 @@ serviceAccount:
   name:
 
 serviceMonitor:
-  enabled: true
+  enabled: false
   interval: 30s
   scrapeTimeout: 10s
   namespace:

```

## 2.10.0

**Release date:** 2022-06-02

![AppVersion: v0.11.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.11.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-mongodb-exporter] - Added podLabels (#2105)

### Default value changes

```diff
diff --git a/charts/prometheus-mongodb-exporter/values.yaml b/charts/prometheus-mongodb-exporter/values.yaml
index 0e245ad9..e7417846 100644
--- a/charts/prometheus-mongodb-exporter/values.yaml
+++ b/charts/prometheus-mongodb-exporter/values.yaml
@@ -14,7 +14,7 @@ fullnameOverride: ""
 image:
   pullPolicy: IfNotPresent
   repository: ssheehy/mongodb-exporter
-  tag: 0.10.0
+  tag: 0.11.0
 
 imagePullSecrets: []
 
@@ -42,6 +42,8 @@ podAnnotations: {}
 #  prometheus.io/scrape: "true"
 #  prometheus.io/port: "metrics"
 
+podLabels: {}
+
 port: "9216"
 
 priorityClassName: ""

```

## 2.9.0

**Release date:** 2022-01-12

![AppVersion: v0.10.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.10.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-mongodb-exporter] Add volumes support (#1706)

### Default value changes

```diff
diff --git a/charts/prometheus-mongodb-exporter/values.yaml b/charts/prometheus-mongodb-exporter/values.yaml
index 446e2a07..0e245ad9 100644
--- a/charts/prometheus-mongodb-exporter/values.yaml
+++ b/charts/prometheus-mongodb-exporter/values.yaml
@@ -65,6 +65,12 @@ resources: {}
 # Extra environment variables that will be passed into the exporter pod
 env: {}
 
+# Volumes that will be mounted into the exporter pod
+volumeMounts: []
+
+# Volumes that will be attached to the exporter deployment
+volumes: []
+
 securityContext:
   allowPrivilegeEscalation: false
   capabilities:

```

## 2.8.1

**Release date:** 2020-08-20

![AppVersion: v0.10.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.10.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Prep initial charts indexing (#14)

### Default value changes

```diff
# No changes in this release
```

## 2.8.0

**Release date:** 2020-07-27

![AppVersion: v0.10.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.10.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-mongodb-exporter] Add ServiceMonitor metricRelabelings (#23290)

### Default value changes

```diff
diff --git a/charts/prometheus-mongodb-exporter/values.yaml b/charts/prometheus-mongodb-exporter/values.yaml
index 32fc30a5..446e2a07 100644
--- a/charts/prometheus-mongodb-exporter/values.yaml
+++ b/charts/prometheus-mongodb-exporter/values.yaml
@@ -93,5 +93,6 @@ serviceMonitor:
   namespace:
   additionalLabels: {}
   targetLabels: []
+  metricRelabelings: []
 
 tolerations: []

```

## 2.7.0

**Release date:** 2020-07-20

![AppVersion: v0.10.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.10.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Possibility to also add a custom key in the secret (#23262)

### Default value changes

```diff
diff --git a/charts/prometheus-mongodb-exporter/values.yaml b/charts/prometheus-mongodb-exporter/values.yaml
index 90b7eb02..32fc30a5 100644
--- a/charts/prometheus-mongodb-exporter/values.yaml
+++ b/charts/prometheus-mongodb-exporter/values.yaml
@@ -32,6 +32,7 @@ mongodb:
 # If this is provided, the value mongodb.uri is ignored.
 existingSecret:
   name: ""
+  key: "mongodb-uri"
 
 nameOverride: ""
 

```

## 2.6.0

**Release date:** 2020-07-08

![AppVersion: v0.10.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.10.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-mongodb-exporter] Add serviceMonitor targetLabels and service labels (#22952)

### Default value changes

```diff
diff --git a/charts/prometheus-mongodb-exporter/values.yaml b/charts/prometheus-mongodb-exporter/values.yaml
index e7baf986..90b7eb02 100644
--- a/charts/prometheus-mongodb-exporter/values.yaml
+++ b/charts/prometheus-mongodb-exporter/values.yaml
@@ -74,6 +74,7 @@ securityContext:
   runAsUser: 10000
 
 service:
+  labels: {}
   annotations: {}
   port: 9216
   type: ClusterIP
@@ -90,5 +91,6 @@ serviceMonitor:
   scrapeTimeout: 10s
   namespace:
   additionalLabels: {}
+  targetLabels: []
 
 tolerations: []

```

## 2.5.0

**Release date:** 2020-04-25

![AppVersion: v0.10.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.10.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Provide the possibility for referring to an existing secret to use the MongoDB URI from (#21931)

### Default value changes

```diff
diff --git a/charts/prometheus-mongodb-exporter/values.yaml b/charts/prometheus-mongodb-exporter/values.yaml
index 5fbd3938..e7baf986 100644
--- a/charts/prometheus-mongodb-exporter/values.yaml
+++ b/charts/prometheus-mongodb-exporter/values.yaml
@@ -24,9 +24,14 @@ livenessProbe:
     port: metrics
   initialDelaySeconds: 10
 
-# [mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]
+# [mongodb[+srv]://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]
 mongodb:
-  uri:
+  uri: ""
+
+# Name of an externally managed secret (in the same namespace) containing the connection uri as key `mongodb-uri`.
+# If this is provided, the value mongodb.uri is ignored.
+existingSecret:
+  name: ""
 
 nameOverride: ""
 

```

## 2.4.0

**Release date:** 2019-11-12

![AppVersion: v0.10.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.10.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* chore(prometheus-mongodb-exporter): version bump percona mongdb exporter (#18388)

### Default value changes

```diff
diff --git a/charts/prometheus-mongodb-exporter/values.yaml b/charts/prometheus-mongodb-exporter/values.yaml
index 0f2910dd..5fbd3938 100644
--- a/charts/prometheus-mongodb-exporter/values.yaml
+++ b/charts/prometheus-mongodb-exporter/values.yaml
@@ -7,13 +7,14 @@ extraArgs:
 - --collect.database
 - --collect.indexusage
 - --collect.topmetrics
+- --collect.connpoolstats
 
 fullnameOverride: ""
 
 image:
   pullPolicy: IfNotPresent
   repository: ssheehy/mongodb-exporter
-  tag: 0.7.0
+  tag: 0.10.0
 
 imagePullSecrets: []
 

```

## 2.3.0

**Release date:** 2019-09-27

![AppVersion: v0.7.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.7.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/promethus-mongodb-exporter] Customize serviceAccountName for mongodb-exporter deployment (#17133)

### Default value changes

```diff
diff --git a/charts/prometheus-mongodb-exporter/values.yaml b/charts/prometheus-mongodb-exporter/values.yaml
index 431746be..0f2910dd 100644
--- a/charts/prometheus-mongodb-exporter/values.yaml
+++ b/charts/prometheus-mongodb-exporter/values.yaml
@@ -72,6 +72,12 @@ service:
   port: 9216
   type: ClusterIP
 
+serviceAccount:
+  create: true
+  # If create is true and name is not set, then a name is generated using the
+  # fullname template.
+  name:
+
 serviceMonitor:
   enabled: true
   interval: 30s

```

## 2.2.1

**Release date:** 2019-09-04

![AppVersion: v0.7.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.7.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Redeploy pods on config changes. (#16835)

### Default value changes

```diff
# No changes in this release
```

## 2.2.0

**Release date:** 2019-07-30

![AppVersion: v0.7.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.7.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-mongodb-exporter] Add Service object for ServiceMonitor to work properly (#15218)

### Default value changes

```diff
diff --git a/charts/prometheus-mongodb-exporter/values.yaml b/charts/prometheus-mongodb-exporter/values.yaml
index 5b58202a..431746be 100644
--- a/charts/prometheus-mongodb-exporter/values.yaml
+++ b/charts/prometheus-mongodb-exporter/values.yaml
@@ -67,6 +67,11 @@ securityContext:
   runAsNonRoot: true
   runAsUser: 10000
 
+service:
+  annotations: {}
+  port: 9216
+  type: ClusterIP
+
 serviceMonitor:
   enabled: true
   interval: 30s

```

## 2.1.0

**Release date:** 2019-05-17

![AppVersion: v0.7.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.7.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-mongodb-exporter] Add MONGODB_URI env variable secret (#13959)

### Default value changes

```diff
# No changes in this release
```

## 2.0.1

**Release date:** 2019-05-14

![AppVersion: v0.7.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.7.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-mongodb-exporter] Fix mongodb-exporter default values (#13756)

### Default value changes

```diff
diff --git a/charts/prometheus-mongodb-exporter/values.yaml b/charts/prometheus-mongodb-exporter/values.yaml
index d3449a5d..5b58202a 100644
--- a/charts/prometheus-mongodb-exporter/values.yaml
+++ b/charts/prometheus-mongodb-exporter/values.yaml
@@ -49,10 +49,10 @@ replicas: 1
 
 resources: {}
 # limits:
-#   cpu: 250mm
+#   cpu: 250m
 #   memory: 192Mi
 # requests:
-#   cpu: 100mm
+#   cpu: 100m
 #   memory: 128Mi
 
 # Extra environment variables that will be passed into the exporter pod

```

## 2.0.0

**Release date:** 2019-04-25

![AppVersion: v0.7.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.7.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-mongodb-exporter] Remove service & bump to v0.7.0 (#12796)

### Default value changes

```diff
diff --git a/charts/prometheus-mongodb-exporter/values.yaml b/charts/prometheus-mongodb-exporter/values.yaml
index eb85d061..d3449a5d 100644
--- a/charts/prometheus-mongodb-exporter/values.yaml
+++ b/charts/prometheus-mongodb-exporter/values.yaml
@@ -3,17 +3,17 @@ affinity: {}
 annotations: {}
 
 extraArgs:
-- -collect.collection=true
-- -collect.database=true
-- -collect.indexusage=true
-- -collect.topmetrics=true
+- --collect.collection
+- --collect.database
+- --collect.indexusage
+- --collect.topmetrics
 
 fullnameOverride: ""
 
 image:
   pullPolicy: IfNotPresent
-  repository: ssalaues/mongodb-exporter
-  tag: 0.6.1
+  repository: ssheehy/mongodb-exporter
+  tag: 0.7.0
 
 imagePullSecrets: []
 
@@ -23,7 +23,7 @@ livenessProbe:
     port: metrics
   initialDelaySeconds: 10
 
-# mongodb://metrics-user:password@mongodb:27017
+# [mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]
 mongodb:
   uri:
 
@@ -31,6 +31,12 @@ nameOverride: ""
 
 nodeSelector: {}
 
+podAnnotations: {}
+#  prometheus.io/scrape: "true"
+#  prometheus.io/port: "metrics"
+
+port: "9216"
+
 priorityClassName: ""
 
 readinessProbe:
@@ -57,16 +63,10 @@ securityContext:
   capabilities:
     drop: ["all"]
   readOnlyRootFilesystem: true
+  runAsGroup: 10000
   runAsNonRoot: true
   runAsUser: 10000
 
-service:
-  annotations: {}
-  #  prometheus.io/scrape: "true"
-  #  prometheus.io/port: "9216"
-  port: 9216
-  type: ClusterIP
-
 serviceMonitor:
   enabled: true
   interval: 30s

```

## 1.1.1

**Release date:** 2019-04-15

![AppVersion: v0.6.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.6.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* adds mongodb prometheus exporter scrape timeout config option (#13063)

### Default value changes

```diff
diff --git a/charts/prometheus-mongodb-exporter/values.yaml b/charts/prometheus-mongodb-exporter/values.yaml
index 7d5bd048..eb85d061 100644
--- a/charts/prometheus-mongodb-exporter/values.yaml
+++ b/charts/prometheus-mongodb-exporter/values.yaml
@@ -70,6 +70,7 @@ service:
 serviceMonitor:
   enabled: true
   interval: 30s
+  scrapeTimeout: 10s
   namespace:
   additionalLabels: {}
 

```

## 1.1.0

**Release date:** 2019-04-01

![AppVersion: v0.6.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.6.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-mongodb-exporter] Add envs to deployment template (#12615)

### Default value changes

```diff
diff --git a/charts/prometheus-mongodb-exporter/values.yaml b/charts/prometheus-mongodb-exporter/values.yaml
index f6b50602..7d5bd048 100644
--- a/charts/prometheus-mongodb-exporter/values.yaml
+++ b/charts/prometheus-mongodb-exporter/values.yaml
@@ -49,6 +49,9 @@ resources: {}
 #   cpu: 100mm
 #   memory: 128Mi
 
+# Extra environment variables that will be passed into the exporter pod
+env: {}
+
 securityContext:
   allowPrivilegeEscalation: false
   capabilities:

```

## 1.0.0

**Release date:** 2019-02-08

![AppVersion: v0.6.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.6.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-mongodb-exporter] Add MongoDB Exporter chart (#10979)

### Default value changes

```diff
affinity: {}

annotations: {}

extraArgs:
- -collect.collection=true
- -collect.database=true
- -collect.indexusage=true
- -collect.topmetrics=true

fullnameOverride: ""

image:
  pullPolicy: IfNotPresent
  repository: ssalaues/mongodb-exporter
  tag: 0.6.1

imagePullSecrets: []

livenessProbe:
  httpGet:
    path: /
    port: metrics
  initialDelaySeconds: 10

# mongodb://metrics-user:password@mongodb:27017
mongodb:
  uri:

nameOverride: ""

nodeSelector: {}

priorityClassName: ""

readinessProbe:
  httpGet:
    path: /
    port: metrics
  initialDelaySeconds: 10

replicas: 1

resources: {}
# limits:
#   cpu: 250mm
#   memory: 192Mi
# requests:
#   cpu: 100mm
#   memory: 128Mi

securityContext:
  allowPrivilegeEscalation: false
  capabilities:
    drop: ["all"]
  readOnlyRootFilesystem: true
  runAsNonRoot: true
  runAsUser: 10000

service:
  annotations: {}
  #  prometheus.io/scrape: "true"
  #  prometheus.io/port: "9216"
  port: 9216
  type: ClusterIP

serviceMonitor:
  enabled: true
  interval: 30s
  namespace:
  additionalLabels: {}

tolerations: []

```

---
Autogenerated from Helm Chart and git history using [helm-changelog](https://github.com/mogensen/helm-changelog)
