# Change Log

## 5.2.0

**Release date:** 2023-06-12

![AppVersion: 1.5.0](https://img.shields.io/static/v1?label=AppVersion&message=1.5.0&color=success)
![Kubernetes: >=1.10.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.10.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-elasticsearch-exporter] Add init containers (#3482)

### Default value changes

```diff
diff --git a/charts/prometheus-elasticsearch-exporter/values.yaml b/charts/prometheus-elasticsearch-exporter/values.yaml
index 0e28b7b1..720070e3 100644
--- a/charts/prometheus-elasticsearch-exporter/values.yaml
+++ b/charts/prometheus-elasticsearch-exporter/values.yaml
@@ -67,6 +67,8 @@ podLabels: {}
 
 affinity: {}
 
+initContainers: []
+
 service:
   type: ClusterIP
   httpPort: 9108

```

## 5.1.3

**Release date:** 2023-06-12

![AppVersion: 1.5.0](https://img.shields.io/static/v1?label=AppVersion&message=1.5.0&color=success)
![Kubernetes: >=1.10.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.10.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-elasticsearch-exporter] fix preStop command (#3466)

### Default value changes

```diff
# No changes in this release
```

## 5.1.2

**Release date:** 2023-06-12

![AppVersion: 1.5.0](https://img.shields.io/static/v1?label=AppVersion&message=1.5.0&color=success)
![Kubernetes: >=1.10.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.10.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-elasticsearch-exporter] Add a chart maintainer (#3483)

### Default value changes

```diff
# No changes in this release
```

## 5.1.1

**Release date:** 2023-04-08

![AppVersion: 1.5.0](https://img.shields.io/static/v1?label=AppVersion&message=1.5.0&color=success)
![Kubernetes: >=1.10.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.10.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Set Pod container name to Service port name (#3157)

### Default value changes

```diff
# No changes in this release
```

## 5.1.0

**Release date:** 2023-03-30

![AppVersion: 1.5.0](https://img.shields.io/static/v1?label=AppVersion&message=1.5.0&color=success)
![Kubernetes: >=1.10.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.10.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-elasticsearch-exporter] add setters for missed out exporters (#3109)

### Default value changes

```diff
diff --git a/charts/prometheus-elasticsearch-exporter/values.yaml b/charts/prometheus-elasticsearch-exporter/values.yaml
index 4821a2e8..0e28b7b1 100644
--- a/charts/prometheus-elasticsearch-exporter/values.yaml
+++ b/charts/prometheus-elasticsearch-exporter/values.yaml
@@ -155,6 +155,10 @@ es:
   ##
   indices_mappings: true
 
+  ## If true, query stats for aliases.
+  ##
+  aliases: false
+
   ## If true, query stats for shards in the cluster.
   ##
   shards: true
@@ -167,6 +171,14 @@ es:
   ##
   cluster_settings: false
 
+  ## If true, query stats for SLM snapshots.
+  ##
+  slm: false
+
+  ## If true, query stats for data streams.
+  ##
+  data_stream: false
+
   ## Timeout for trying to get stats from Elasticsearch. (ex: 20s)
   ##
   timeout: 30s

```

## 5.0.0

**Release date:** 2022-12-01

![AppVersion: 1.5.0](https://img.shields.io/static/v1?label=AppVersion&message=1.5.0&color=success)
![Kubernetes: >=1.10.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.10.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-elasticsearch-exporter] Refactor securityContext configuration (#2694)

### Default value changes

```diff
diff --git a/charts/prometheus-elasticsearch-exporter/values.yaml b/charts/prometheus-elasticsearch-exporter/values.yaml
index 3a2c918a..4821a2e8 100644
--- a/charts/prometheus-elasticsearch-exporter/values.yaml
+++ b/charts/prometheus-elasticsearch-exporter/values.yaml
@@ -18,13 +18,18 @@ image:
   pullPolicy: IfNotPresent
   pullSecret: ""
 
-## Set enabled to false if you don't want securityContext
-## in your Deployment.
-## The below values are the default for kubernetes.
-## Openshift won't deploy with runAsUser: 1000 without additional permissions.
-securityContext:
-  enabled: true  # Should be set to false when running on OpenShift
+podSecurityContext:
+  runAsNonRoot: true
   runAsUser: 1000
+  seccompProfile:
+    type: "RuntimeDefault"
+
+securityContext:
+  allowPrivilegeEscalation: false
+  capabilities:
+    drop:
+      - ALL
+  readOnlyRootFilesystem: true
 
 # Custom DNS configuration to be added to prometheus-elasticsearch-exporter pods
 dnsConfig: {}

```

## 4.15.1

**Release date:** 2022-10-16

![AppVersion: 1.5.0](https://img.shields.io/static/v1?label=AppVersion&message=1.5.0&color=success)
![Kubernetes: >=1.10.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.10.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-elasticsearch-exporter] Fix PSP deprecation after k8s 1.25+ (#2569)

### Default value changes

```diff
# No changes in this release
```

## 4.15.0

**Release date:** 2022-09-21

![AppVersion: 1.5.0](https://img.shields.io/static/v1?label=AppVersion&message=1.5.0&color=success)
![Kubernetes: >=1.10.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.10.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-elasticsearch-exporter] allow configuration of deployment labels (#2085)

### Default value changes

```diff
diff --git a/charts/prometheus-elasticsearch-exporter/values.yaml b/charts/prometheus-elasticsearch-exporter/values.yaml
index d0da9a11..3a2c918a 100644
--- a/charts/prometheus-elasticsearch-exporter/values.yaml
+++ b/charts/prometheus-elasticsearch-exporter/values.yaml
@@ -72,6 +72,7 @@ service:
 
 deployment:
   annotations: {}
+  labels: {}
 
 ## Extra environment variables that will be passed into the exporter pod
 ## example:

```

## 4.14.1

**Release date:** 2022-09-13

![AppVersion: 1.5.0](https://img.shields.io/static/v1?label=AppVersion&message=1.5.0&color=success)
![Kubernetes: >=1.10.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.10.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* feature(prometheus-elasticsearch-exporter): Added option to use tpl to define secret name. (#2423)

### Default value changes

```diff
# No changes in this release
```

## 4.14.0

**Release date:** 2022-08-04

![AppVersion: 1.5.0](https://img.shields.io/static/v1?label=AppVersion&message=1.5.0&color=success)
![Kubernetes: >=1.10.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.10.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-elasticsearch-exporter] Bump exporter to 1.5.0 (#2342)

### Default value changes

```diff
diff --git a/charts/prometheus-elasticsearch-exporter/values.yaml b/charts/prometheus-elasticsearch-exporter/values.yaml
index 2a32b714..d0da9a11 100644
--- a/charts/prometheus-elasticsearch-exporter/values.yaml
+++ b/charts/prometheus-elasticsearch-exporter/values.yaml
@@ -14,7 +14,7 @@ restartPolicy: Always
 
 image:
   repository: quay.io/prometheuscommunity/elasticsearch-exporter
-  tag: v1.3.0
+  tag: v1.5.0
   pullPolicy: IfNotPresent
   pullSecret: ""
 

```

## 4.13.0

**Release date:** 2022-06-13

![AppVersion: 1.3.0](https://img.shields.io/static/v1?label=AppVersion&message=1.3.0&color=success)
![Kubernetes: >=1.10.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.10.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-elasticsearch-exporter] Allow to set global imagePullSecrets (#2039)

### Default value changes

```diff
diff --git a/charts/prometheus-elasticsearch-exporter/values.yaml b/charts/prometheus-elasticsearch-exporter/values.yaml
index 9bf17c0f..2a32b714 100644
--- a/charts/prometheus-elasticsearch-exporter/values.yaml
+++ b/charts/prometheus-elasticsearch-exporter/values.yaml
@@ -1,3 +1,9 @@
+## Global settings for all charts in a tree of dependent charts
+## This allows to set the same values for the imagePullSecrets, obviating the need
+## to set the same values for each chart (and image) separately
+global:
+  imagePullSecrets: []
+
 ## number of exporter instances
 ##
 replicaCount: 1

```

## 4.12.0

**Release date:** 2022-05-19

![AppVersion: 1.3.0](https://img.shields.io/static/v1?label=AppVersion&message=1.3.0&color=success)
![Kubernetes: >=1.10.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.10.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-elasticsearch-exporter] Allow jobLabel configuration (#2017)

### Default value changes

```diff
diff --git a/charts/prometheus-elasticsearch-exporter/values.yaml b/charts/prometheus-elasticsearch-exporter/values.yaml
index 7f61fc1e..9bf17c0f 100644
--- a/charts/prometheus-elasticsearch-exporter/values.yaml
+++ b/charts/prometheus-elasticsearch-exporter/values.yaml
@@ -74,7 +74,6 @@ deployment:
 ##   KEY_2: value2
 env: {}
 
-
 ## Extra arguments to pass to the container's executable
 ## example:
 # extraArgs:
@@ -165,7 +164,6 @@ es:
   ##
   sslSkipVerify: false
 
-
   ssl:
     ## If true, a secure connection to ES cluster is used
     ##
@@ -176,7 +174,6 @@ es:
     useExistingSecrets: false
 
     ca:
-
       ## PEM that contains trusted CAs used for setting up secure Elasticsearch connection
       ##
       # pem:
@@ -214,6 +211,7 @@ serviceMonitor:
   #  namespace: monitoring
   labels: {}
   interval: 10s
+  jobLabel: ""
   scrapeTimeout: 10s
   scheme: http
   relabelings: []

```

## 4.11.2

**Release date:** 2022-05-10

![AppVersion: 1.3.0](https://img.shields.io/static/v1?label=AppVersion&message=1.3.0&color=success)
![Kubernetes: >=1.10.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.10.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-elasticsearch-exporter] remove me from elasticsearch chart (#2040)

### Default value changes

```diff
# No changes in this release
```

## 4.11.1

**Release date:** 2022-03-27

![AppVersion: 1.3.0](https://img.shields.io/static/v1?label=AppVersion&message=1.3.0&color=success)
![Kubernetes: >=1.10.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.10.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* fixed prometheus-elasticsearch-exporter pre-stop lifecycle hook (#1857)

### Default value changes

```diff
# No changes in this release
```

## 4.11.0

**Release date:** 2022-01-02

![AppVersion: 1.3.0](https://img.shields.io/static/v1?label=AppVersion&message=1.3.0&color=success)
![Kubernetes: >=1.10.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.10.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-elasticsearch-exporter] Support extra arguments to the pod's container (#1547)

### Default value changes

```diff
diff --git a/charts/prometheus-elasticsearch-exporter/values.yaml b/charts/prometheus-elasticsearch-exporter/values.yaml
index 6727a0ee..7f61fc1e 100644
--- a/charts/prometheus-elasticsearch-exporter/values.yaml
+++ b/charts/prometheus-elasticsearch-exporter/values.yaml
@@ -74,6 +74,13 @@ deployment:
 ##   KEY_2: value2
 env: {}
 
+
+## Extra arguments to pass to the container's executable
+## example:
+# extraArgs:
+#   - --es.indices_mappings
+extraArgs: []
+
 ## The name of a secret in the same kubernetes namespace which contain values to be added to the environment
 ## This can be useful for auth tokens, etc
 envFromSecret: ""

```

## 4.10.0

**Release date:** 2021-12-20

![AppVersion: 1.3.0](https://img.shields.io/static/v1?label=AppVersion&message=1.3.0&color=success)
![Kubernetes: >=1.10.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.10.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-elasticsearch-exporter] Add support for indices mappings (#1610)

### Default value changes

```diff
diff --git a/charts/prometheus-elasticsearch-exporter/values.yaml b/charts/prometheus-elasticsearch-exporter/values.yaml
index fafd9460..6727a0ee 100644
--- a/charts/prometheus-elasticsearch-exporter/values.yaml
+++ b/charts/prometheus-elasticsearch-exporter/values.yaml
@@ -133,6 +133,10 @@ es:
   ##
   indices_settings: true
 
+  ## If true, query mapping stats for all indices in the cluster.
+  ##
+  indices_mappings: true
+
   ## If true, query stats for shards in the cluster.
   ##
   shards: true

```

## 4.9.0

**Release date:** 2021-11-30

![AppVersion: 1.3.0](https://img.shields.io/static/v1?label=AppVersion&message=1.3.0&color=success)
![Kubernetes: >=1.10.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.10.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-elasticsearch-exporter] Bump exporter to 1.3.0 (#1546)

### Default value changes

```diff
diff --git a/charts/prometheus-elasticsearch-exporter/values.yaml b/charts/prometheus-elasticsearch-exporter/values.yaml
index 6050825f..fafd9460 100644
--- a/charts/prometheus-elasticsearch-exporter/values.yaml
+++ b/charts/prometheus-elasticsearch-exporter/values.yaml
@@ -8,7 +8,7 @@ restartPolicy: Always
 
 image:
   repository: quay.io/prometheuscommunity/elasticsearch-exporter
-  tag: v1.2.1
+  tag: v1.3.0
   pullPolicy: IfNotPresent
   pullSecret: ""
 

```

## 4.8.0

**Release date:** 2021-11-27

![AppVersion: 1.2.1](https://img.shields.io/static/v1?label=AppVersion&message=1.2.1&color=success)
![Kubernetes: >=1.10.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.10.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* add support for serviceaccount annotations (#1529)

### Default value changes

```diff
diff --git a/charts/prometheus-elasticsearch-exporter/values.yaml b/charts/prometheus-elasticsearch-exporter/values.yaml
index e2af0005..6050825f 100644
--- a/charts/prometheus-elasticsearch-exporter/values.yaml
+++ b/charts/prometheus-elasticsearch-exporter/values.yaml
@@ -251,6 +251,7 @@ serviceAccount:
   create: false
   name: default
   automountServiceAccountToken: true
+  annotations: {}
 
 # Creates a PodSecurityPolicy and the role/rolebinding
 # allowing the serviceaccount to use it

```

## 4.7.0

**Release date:** 2021-09-24

![AppVersion: 1.2.1](https://img.shields.io/static/v1?label=AppVersion&message=1.2.1&color=success)
![Kubernetes: >=1.10.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.10.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-elasticsearch-exporter] Use appropriate apiVersion for rbac (#1146)

### Default value changes

```diff
# No changes in this release
```

## 4.6.1

**Release date:** 2021-09-12

![AppVersion: 1.2.1](https://img.shields.io/static/v1?label=AppVersion&message=1.2.1&color=success)
![Kubernetes: >=1.10.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.10.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Fixes: #1226 by using tpl funcion at es.uri injection (#1227)
* [prometheus-elasticsearch-exporter] add optional annotations to deployment (#1224)

### Default value changes

```diff
diff --git a/charts/prometheus-elasticsearch-exporter/values.yaml b/charts/prometheus-elasticsearch-exporter/values.yaml
index fb42ecef..e2af0005 100644
--- a/charts/prometheus-elasticsearch-exporter/values.yaml
+++ b/charts/prometheus-elasticsearch-exporter/values.yaml
@@ -64,6 +64,9 @@ service:
   annotations: {}
   labels: {}
 
+deployment:
+  annotations: {}
+
 ## Extra environment variables that will be passed into the exporter pod
 ## example:
 ## env:

```

## 4.6.0

**Release date:** 2021-08-11

![AppVersion: 1.2.1](https://img.shields.io/static/v1?label=AppVersion&message=1.2.1&color=success)
![Kubernetes: >=1.10.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.10.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-elasticsearch-exporter] Add ability to use values from secrets for basic-auth (#1182)

### Default value changes

```diff
# No changes in this release
```

## 4.5.0

**Release date:** 2021-07-20

![AppVersion: 1.2.1](https://img.shields.io/static/v1?label=AppVersion&message=1.2.1&color=success)
![Kubernetes: >=1.10.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.10.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-elasticsearch-exporter] update image version (#1126)

### Default value changes

```diff
diff --git a/charts/prometheus-elasticsearch-exporter/values.yaml b/charts/prometheus-elasticsearch-exporter/values.yaml
index 62f98ad6..fb42ecef 100644
--- a/charts/prometheus-elasticsearch-exporter/values.yaml
+++ b/charts/prometheus-elasticsearch-exporter/values.yaml
@@ -7,8 +7,8 @@ replicaCount: 1
 restartPolicy: Always
 
 image:
-  repository: justwatch/elasticsearch_exporter
-  tag: 1.1.0
+  repository: quay.io/prometheuscommunity/elasticsearch-exporter
+  tag: v1.2.1
   pullPolicy: IfNotPresent
   pullSecret: ""
 

```

## 4.4.0

**Release date:** 2021-04-02

![AppVersion: 1.1.0](https://img.shields.io/static/v1?label=AppVersion&message=1.1.0&color=success)
![Kubernetes: >=1.10.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.10.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Allows to set automountServiceAccountToken on ServiceAccount (#808)

### Default value changes

```diff
diff --git a/charts/prometheus-elasticsearch-exporter/values.yaml b/charts/prometheus-elasticsearch-exporter/values.yaml
index 0da12320..62f98ad6 100644
--- a/charts/prometheus-elasticsearch-exporter/values.yaml
+++ b/charts/prometheus-elasticsearch-exporter/values.yaml
@@ -247,6 +247,7 @@ prometheusRule:
 serviceAccount:
   create: false
   name: default
+  automountServiceAccountToken: true
 
 # Creates a PodSecurityPolicy and the role/rolebinding
 # allowing the serviceaccount to use it

```

## 4.3.0

**Release date:** 2021-02-23

![AppVersion: 1.1.0](https://img.shields.io/static/v1?label=AppVersion&message=1.1.0&color=success)
![Kubernetes: >=1.10.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.10.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* feat(extraVolume\Mounts): introduction (#698)

### Default value changes

```diff
diff --git a/charts/prometheus-elasticsearch-exporter/values.yaml b/charts/prometheus-elasticsearch-exporter/values.yaml
index 00a22034..0da12320 100644
--- a/charts/prometheus-elasticsearch-exporter/values.yaml
+++ b/charts/prometheus-elasticsearch-exporter/values.yaml
@@ -91,6 +91,24 @@ secretMounts: []
 #    secretName: elastic-certs
 #    path: /ssl
 
+# A list of additional Volume to add to the deployment
+# this is useful if the volume you need is not a secret (csi volume etc.)
+extraVolumes: []
+#  - name: csi-volume
+#    csi:
+#      driver: secrets-store.csi.k8s.io
+#      readOnly: true
+#      volumeAttributes:
+#        secretProviderClass: my-spc
+
+#  A list of additional VolumeMounts to add to the deployment
+#  this is useful for mounting any other needed resource into
+#  the elasticsearch-exporter pod
+extraVolumeMounts: []
+#  - name: csi-volume
+#    mountPath: /csi/volume
+#    readOnly: true
+
 es:
   ## Address (host and port) of the Elasticsearch node we should connect to.
   ## This could be a local node (localhost:9200, for instance), or the address

```

## 4.2.0

**Release date:** 2021-02-09

![AppVersion: 1.1.0](https://img.shields.io/static/v1?label=AppVersion&message=1.1.0&color=success)
![Kubernetes: >=1.10.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.10.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Ability for custom dnsConfig in prometheus-elasticsearch-exporter (#641)

### Default value changes

```diff
diff --git a/charts/prometheus-elasticsearch-exporter/values.yaml b/charts/prometheus-elasticsearch-exporter/values.yaml
index 00caa62e..00a22034 100644
--- a/charts/prometheus-elasticsearch-exporter/values.yaml
+++ b/charts/prometheus-elasticsearch-exporter/values.yaml
@@ -20,6 +20,18 @@ securityContext:
   enabled: true  # Should be set to false when running on OpenShift
   runAsUser: 1000
 
+# Custom DNS configuration to be added to prometheus-elasticsearch-exporter pods
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
 log:
   format: logfmt
   level: info

```

## 4.1.0

**Release date:** 2021-01-22

![AppVersion: 1.1.0](https://img.shields.io/static/v1?label=AppVersion&message=1.1.0&color=success)
![Kubernetes: >=1.10.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.10.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-elasticsearch-exporter] Add additional pod labels support (#578)

### Default value changes

```diff
diff --git a/charts/prometheus-elasticsearch-exporter/values.yaml b/charts/prometheus-elasticsearch-exporter/values.yaml
index 82a5f1e2..00caa62e 100644
--- a/charts/prometheus-elasticsearch-exporter/values.yaml
+++ b/charts/prometheus-elasticsearch-exporter/values.yaml
@@ -40,6 +40,8 @@ tolerations: []
 
 podAnnotations: {}
 
+podLabels: {}
+
 affinity: {}
 
 service:

```

## 4.0.1

**Release date:** 2020-12-10

![AppVersion: 1.1.0](https://img.shields.io/static/v1?label=AppVersion&message=1.1.0&color=success)
![Kubernetes: >=1.10.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.10.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-elasticsearch-exporter] change tolerations from map to list (#476)
* Update charts/prometheus-elasticsearch-exporter/README.md
* used fixed name for container

### Default value changes

```diff
diff --git a/charts/prometheus-elasticsearch-exporter/values.yaml b/charts/prometheus-elasticsearch-exporter/values.yaml
index 385ff4e1..82a5f1e2 100644
--- a/charts/prometheus-elasticsearch-exporter/values.yaml
+++ b/charts/prometheus-elasticsearch-exporter/values.yaml
@@ -36,7 +36,7 @@ priorityClassName: ""
 
 nodeSelector: {}
 
-tolerations: {}
+tolerations: []
 
 podAnnotations: {}
 

```

## 4.0.0

**Release date:** 2020-10-23

![AppVersion: 1.1.0](https://img.shields.io/static/v1?label=AppVersion&message=1.1.0&color=success)
![Kubernetes: >=1.10.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.10.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Added upgrade instructions to migrate from stable chart
* Update charts/prometheus-elasticsearch-exporter/README.md
* Fixed copy & paste mistakes

### Default value changes

```diff
# No changes in this release
```

## 3.8.0

**Release date:** 2020-10-06

![AppVersion: 1.1.0](https://img.shields.io/static/v1?label=AppVersion&message=1.1.0&color=success)
![Kubernetes: >=1.10.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.10.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Bump Chart versions, adjust README

### Default value changes

```diff
# No changes in this release
```

## 3.7.0

**Release date:** 2020-08-08

![AppVersion: 1.1.0](https://img.shields.io/static/v1?label=AppVersion&message=1.1.0&color=success)
![Kubernetes: >=1.10.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.10.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/elasticsearch-exporter] Make sampleLimit adjustable (#23454)

### Default value changes

```diff
diff --git a/charts/prometheus-elasticsearch-exporter/values.yaml b/charts/prometheus-elasticsearch-exporter/values.yaml
index 9b8d524c..385ff4e1 100644
--- a/charts/prometheus-elasticsearch-exporter/values.yaml
+++ b/charts/prometheus-elasticsearch-exporter/values.yaml
@@ -173,6 +173,7 @@ serviceMonitor:
   relabelings: []
   targetLabels: []
   metricRelabelings: []
+  sampleLimit: 0
 
 prometheusRule:
   ## If true, a PrometheusRule CRD is created for a prometheus operator

```

## 3.6.0

**Release date:** 2020-07-22

![AppVersion: 1.1.0](https://img.shields.io/static/v1?label=AppVersion&message=1.1.0&color=success)
![Kubernetes: >=1.10.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.10.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/elasticsearch-exporter] Add ServiceMonitor metricRelabeling (#23291)

### Default value changes

```diff
diff --git a/charts/prometheus-elasticsearch-exporter/values.yaml b/charts/prometheus-elasticsearch-exporter/values.yaml
index 453deef2..9b8d524c 100644
--- a/charts/prometheus-elasticsearch-exporter/values.yaml
+++ b/charts/prometheus-elasticsearch-exporter/values.yaml
@@ -172,6 +172,7 @@ serviceMonitor:
   scheme: http
   relabelings: []
   targetLabels: []
+  metricRelabelings: []
 
 prometheusRule:
   ## If true, a PrometheusRule CRD is created for a prometheus operator

```

## 3.5.0

**Release date:** 2020-07-17

![AppVersion: 1.1.0](https://img.shields.io/static/v1?label=AppVersion&message=1.1.0&color=success)
![Kubernetes: >=1.10.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.10.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/elasticsearch-exporter]: customize log flags (#23216)

### Default value changes

```diff
diff --git a/charts/prometheus-elasticsearch-exporter/values.yaml b/charts/prometheus-elasticsearch-exporter/values.yaml
index 579c8198..453deef2 100644
--- a/charts/prometheus-elasticsearch-exporter/values.yaml
+++ b/charts/prometheus-elasticsearch-exporter/values.yaml
@@ -20,6 +20,10 @@ securityContext:
   enabled: true  # Should be set to false when running on OpenShift
   runAsUser: 1000
 
+log:
+  format: logfmt
+  level: info
+
 resources: {}
   # requests:
   #   cpu: 100m

```

## 3.4.0

**Release date:** 2020-06-26

![AppVersion: 1.1.0](https://img.shields.io/static/v1?label=AppVersion&message=1.1.0&color=success)
![Kubernetes: >=1.10.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.10.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/elasticsearch-exporter] Add serviceMonitor targetLabels (#22951)

### Default value changes

```diff
diff --git a/charts/prometheus-elasticsearch-exporter/values.yaml b/charts/prometheus-elasticsearch-exporter/values.yaml
index 0fd2a632..579c8198 100644
--- a/charts/prometheus-elasticsearch-exporter/values.yaml
+++ b/charts/prometheus-elasticsearch-exporter/values.yaml
@@ -167,6 +167,7 @@ serviceMonitor:
   scrapeTimeout: 10s
   scheme: http
   relabelings: []
+  targetLabels: []
 
 prometheusRule:
   ## If true, a PrometheusRule CRD is created for a prometheus operator

```

## 3.3.0

**Release date:** 2020-05-04

![AppVersion: 1.1.0](https://img.shields.io/static/v1?label=AppVersion&message=1.1.0&color=success)
![Kubernetes: >=1.10.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.10.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/elasticsearch-exporter] Add options to set relabel configuration for metrics (#21636)

### Default value changes

```diff
diff --git a/charts/prometheus-elasticsearch-exporter/values.yaml b/charts/prometheus-elasticsearch-exporter/values.yaml
index bf92a297..0fd2a632 100644
--- a/charts/prometheus-elasticsearch-exporter/values.yaml
+++ b/charts/prometheus-elasticsearch-exporter/values.yaml
@@ -166,6 +166,7 @@ serviceMonitor:
   interval: 10s
   scrapeTimeout: 10s
   scheme: http
+  relabelings: []
 
 prometheusRule:
   ## If true, a PrometheusRule CRD is created for a prometheus operator

```

## 3.2.0

**Release date:** 2020-05-02

![AppVersion: 1.1.0](https://img.shields.io/static/v1?label=AppVersion&message=1.1.0&color=success)
![Kubernetes: >=1.10.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.10.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/elasticsearch-exporter] Add option to disable client cert auth in Elasticsearch exporter (#21900)

### Default value changes

```diff
diff --git a/charts/prometheus-elasticsearch-exporter/values.yaml b/charts/prometheus-elasticsearch-exporter/values.yaml
index 441b05d3..bf92a297 100644
--- a/charts/prometheus-elasticsearch-exporter/values.yaml
+++ b/charts/prometheus-elasticsearch-exporter/values.yaml
@@ -117,7 +117,7 @@ es:
 
 
   ssl:
-    ## If true, a secure connection to ES cluster is used (requires SSL certs below)
+    ## If true, a secure connection to ES cluster is used
     ##
     enabled: false
 
@@ -132,22 +132,25 @@ es:
       # pem:
 
       # Path of ca pem file which should match a secretMount path
-      # path: /ssl/ca.pem
+      path: /ssl/ca.pem
     client:
+      ## if true, client SSL certificate is used for authentication
+      ##
+      enabled: true
 
       ## PEM that contains the client cert to connect to Elasticsearch.
       ##
       # pem:
 
       # Path of client pem file which should match a secretMount path
-      # pemPath: /ssl/client.pem
+      pemPath: /ssl/client.pem
 
       ## Private key for client auth when connecting to Elasticsearch
       ##
       # key:
 
       # Path of client key file which should match a secretMount path
-      # keyPath: /ssl/client.key
+      keyPath: /ssl/client.key
 web:
   ## Path under which to expose metrics.
   ##

```

## 3.1.0

**Release date:** 2020-04-28

![AppVersion: 1.1.0](https://img.shields.io/static/v1?label=AppVersion&message=1.1.0&color=success)
![Kubernetes: >=1.10.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.10.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/elasticsearch-exporter] add extraEnvSecrets (#22067)

### Default value changes

```diff
diff --git a/charts/prometheus-elasticsearch-exporter/values.yaml b/charts/prometheus-elasticsearch-exporter/values.yaml
index 4afdcf86..441b05d3 100644
--- a/charts/prometheus-elasticsearch-exporter/values.yaml
+++ b/charts/prometheus-elasticsearch-exporter/values.yaml
@@ -57,6 +57,15 @@ env: {}
 ## This can be useful for auth tokens, etc
 envFromSecret: ""
 
+## A list of environment variables from secret refs that will be passed into the exporter pod
+## example:
+## This will set ${ES_PASSWORD} to the 'password' key from the 'my-secret' secret
+## extraEnvSecrets:
+##   ES_PASSWORD:
+##     secret: my-secret
+##     key: password
+extraEnvSecrets: {}
+
 # A list of secrets and their paths to mount inside the pod
 # This is useful for mounting certificates for security
 secretMounts: []

```

## 3.0.0

**Release date:** 2020-02-12

![AppVersion: 1.1.0](https://img.shields.io/static/v1?label=AppVersion&message=1.1.0&color=success)
![Kubernetes: >=1.10.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.10.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/elasticsearch-exporter] Allow PrometheusRules to be templated. (#20322)

### Default value changes

```diff
diff --git a/charts/prometheus-elasticsearch-exporter/values.yaml b/charts/prometheus-elasticsearch-exporter/values.yaml
index 4c5ca779..4afdcf86 100644
--- a/charts/prometheus-elasticsearch-exporter/values.yaml
+++ b/charts/prometheus-elasticsearch-exporter/values.yaml
@@ -159,32 +159,35 @@ prometheusRule:
   ## If true, a PrometheusRule CRD is created for a prometheus operator
   ## https://github.com/coreos/prometheus-operator
   ##
+  ## The rules will be processed as Helm template, allowing to set variables in them.
   enabled: false
   #  namespace: monitoring
   labels: {}
   rules: []
     # - record: elasticsearch_filesystem_data_used_percent
-    #   expr: 100 * (elasticsearch_filesystem_data_size_bytes - elasticsearch_filesystem_data_free_bytes)
-    #     / elasticsearch_filesystem_data_size_bytes
+    #   expr: |
+    #     100 * (elasticsearch_filesystem_data_size_bytes{service="{{ template "elasticsearch-exporter.fullname" . }}"} - elasticsearch_filesystem_data_free_bytes{service="{{ template "elasticsearch-exporter.fullname" . }}"})
+    #     / elasticsearch_filesystem_data_size_bytes{service="{{ template "elasticsearch-exporter.fullname" . }}"}
     # - record: elasticsearch_filesystem_data_free_percent
-    #   expr: 100 - elasticsearch_filesystem_data_used_percent
+    #   expr: 100 - elasticsearch_filesystem_data_used_percent{service="{{ template "elasticsearch-exporter.fullname" . }}"}
     # - alert: ElasticsearchTooFewNodesRunning
-    #   expr: elasticsearch_cluster_health_number_of_nodes < 3
+    #   expr: elasticsearch_cluster_health_number_of_nodes{service="{{ template "elasticsearch-exporter.fullname" . }}"} < 3
     #   for: 5m
     #   labels:
     #     severity: critical
     #   annotations:
-    #     description: There are only {{$value}} < 3 ElasticSearch nodes running
+    #     description: There are only {{ "{{ $value }}" }} < 3 ElasticSearch nodes running
     #     summary: ElasticSearch running on less than 3 nodes
     # - alert: ElasticsearchHeapTooHigh
-    #   expr: elasticsearch_jvm_memory_used_bytes{area="heap"} / elasticsearch_jvm_memory_max_bytes{area="heap"}
+    #   expr: |
+    #     elasticsearch_jvm_memory_used_bytes{service="{{ template "elasticsearch-exporter.fullname" . }}", area="heap"} / elasticsearch_jvm_memory_max_bytes{service="{{ template "elasticsearch-exporter.fullname" . }}", area="heap"}
     #     > 0.9
     #   for: 15m
     #   labels:
     #     severity: critical
     #   annotations:
     #     description: The heap usage is over 90% for 15m
-    #     summary: ElasticSearch node {{$labels.node}} heap usage is high
+    #     summary: ElasticSearch node {{ "{{ $labels.node }}" }} heap usage is high
 
 # Create a service account
 # To use a service account not handled by the chart, set the name here

```

## 2.3.0

**Release date:** 2020-02-07

![AppVersion: 1.1.0](https://img.shields.io/static/v1?label=AppVersion&message=1.1.0&color=success)
![Kubernetes: >=1.10.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.10.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/elasticsearch-exporter] Feature flap the flag es.uri (#20589)

### Default value changes

```diff
# No changes in this release
```

## 2.2.0

**Release date:** 2020-01-15

![AppVersion: 1.1.0](https://img.shields.io/static/v1?label=AppVersion&message=1.1.0&color=success)
![Kubernetes: >=1.10.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.10.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/elasticsearch-exporter] Enabling environment from secrets (#19115)

### Default value changes

```diff
diff --git a/charts/prometheus-elasticsearch-exporter/values.yaml b/charts/prometheus-elasticsearch-exporter/values.yaml
index 2ddce73d..4c5ca779 100644
--- a/charts/prometheus-elasticsearch-exporter/values.yaml
+++ b/charts/prometheus-elasticsearch-exporter/values.yaml
@@ -51,7 +51,11 @@ service:
 ## env:
 ##   KEY_1: value1
 ##   KEY_2: value2
-# env:
+env: {}
+
+## The name of a secret in the same kubernetes namespace which contain values to be added to the environment
+## This can be useful for auth tokens, etc
+envFromSecret: ""
 
 # A list of secrets and their paths to mount inside the pod
 # This is useful for mounting certificates for security

```

## 2.1.1

**Release date:** 2019-11-15

![AppVersion: 1.1.0](https://img.shields.io/static/v1?label=AppVersion&message=1.1.0&color=success)
![Kubernetes: >=1.10.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.10.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Add option es.indices_settings (#18909)

### Default value changes

```diff
diff --git a/charts/prometheus-elasticsearch-exporter/values.yaml b/charts/prometheus-elasticsearch-exporter/values.yaml
index 69bca7f2..2ddce73d 100644
--- a/charts/prometheus-elasticsearch-exporter/values.yaml
+++ b/charts/prometheus-elasticsearch-exporter/values.yaml
@@ -77,6 +77,10 @@ es:
   ##
   indices: true
 
+  ## If true, query settings stats for all indices in the cluster.
+  ##
+  indices_settings: true
+
   ## If true, query stats for shards in the cluster.
   ##
   shards: true

```

## 2.1.0

**Release date:** 2019-10-31

![AppVersion: 1.1.0](https://img.shields.io/static/v1?label=AppVersion&message=1.1.0&color=success)
![Kubernetes: >=1.10.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.10.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/elasticsearch-exporter]: improved healthchecks (#16808)

### Default value changes

```diff
# No changes in this release
```

## 2.0.0

**Release date:** 2019-10-15

![AppVersion: 1.1.0](https://img.shields.io/static/v1?label=AppVersion&message=1.1.0&color=success)
![Kubernetes: >=1.10.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.10.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/elasticsearch-exporter] Update api versions to support k8s 1.16 (#17644)

### Default value changes

```diff
# No changes in this release
```

## 1.11.0

**Release date:** 2019-10-09

![AppVersion: 1.1.0](https://img.shields.io/static/v1?label=AppVersion&message=1.1.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Add support for PodSecurityPolicies, ServiceAccounts (#16361)

### Default value changes

```diff
diff --git a/charts/prometheus-elasticsearch-exporter/values.yaml b/charts/prometheus-elasticsearch-exporter/values.yaml
index 1afdf51c..69bca7f2 100644
--- a/charts/prometheus-elasticsearch-exporter/values.yaml
+++ b/charts/prometheus-elasticsearch-exporter/values.yaml
@@ -177,3 +177,15 @@ prometheusRule:
     #   annotations:
     #     description: The heap usage is over 90% for 15m
     #     summary: ElasticSearch node {{$labels.node}} heap usage is high
+
+# Create a service account
+# To use a service account not handled by the chart, set the name here
+# and set create to false
+serviceAccount:
+  create: false
+  name: default
+
+# Creates a PodSecurityPolicy and the role/rolebinding
+# allowing the serviceaccount to use it
+podSecurityPolicies:
+  enabled: false

```

## 1.10.1

**Release date:** 2019-09-12

![AppVersion: 1.1.0](https://img.shields.io/static/v1?label=AppVersion&message=1.1.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Fix space issue in volume mount (#17087)

### Default value changes

```diff
diff --git a/charts/prometheus-elasticsearch-exporter/values.yaml b/charts/prometheus-elasticsearch-exporter/values.yaml
index 1abe10cd..1afdf51c 100644
--- a/charts/prometheus-elasticsearch-exporter/values.yaml
+++ b/charts/prometheus-elasticsearch-exporter/values.yaml
@@ -56,15 +56,9 @@ service:
 # A list of secrets and their paths to mount inside the pod
 # This is useful for mounting certificates for security
 secretMounts: []
-#  - name: ca
-#    secretName: elastic-ca
-#    path: /ssl/ca.pem
-#  - name: elastic-client-pem
-#    secretName: elastic-client-pem
-#    path: /ssl/client.pem
-#  - name: elastic-client-key
-#    secretName: elastic-client-key
-#    path: /ssl/client.key
+#  - name: elastic-certs
+#    secretName: elastic-certs
+#    path: /ssl
 
 es:
   ## Address (host and port) of the Elasticsearch node we should connect to.
@@ -121,7 +115,7 @@ es:
       # pem:
 
       # Path of ca pem file which should match a secretMount path
-      # pemPath: /ssl/ca.pem
+      # path: /ssl/ca.pem
     client:
 
       ## PEM that contains the client cert to connect to Elasticsearch.

```

## 1.10.0

**Release date:** 2019-09-04

![AppVersion: 1.1.0](https://img.shields.io/static/v1?label=AppVersion&message=1.1.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Add ability to use existing secrets for SSL certs (#16817)

### Default value changes

```diff
diff --git a/charts/prometheus-elasticsearch-exporter/values.yaml b/charts/prometheus-elasticsearch-exporter/values.yaml
index 3f156d7b..1abe10cd 100644
--- a/charts/prometheus-elasticsearch-exporter/values.yaml
+++ b/charts/prometheus-elasticsearch-exporter/values.yaml
@@ -53,6 +53,19 @@ service:
 ##   KEY_2: value2
 # env:
 
+# A list of secrets and their paths to mount inside the pod
+# This is useful for mounting certificates for security
+secretMounts: []
+#  - name: ca
+#    secretName: elastic-ca
+#    path: /ssl/ca.pem
+#  - name: elastic-client-pem
+#    secretName: elastic-client-pem
+#    path: /ssl/client.pem
+#  - name: elastic-client-key
+#    secretName: elastic-client-key
+#    path: /ssl/client.key
+
 es:
   ## Address (host and port) of the Elasticsearch node we should connect to.
   ## This could be a local node (localhost:9200, for instance), or the address
@@ -97,22 +110,33 @@ es:
     ##
     enabled: false
 
+    ## If true, certs from secretMounts will be need to be referenced instead of certs below
+    ##
+    useExistingSecrets: false
+
     ca:
 
       ## PEM that contains trusted CAs used for setting up secure Elasticsearch connection
       ##
       # pem:
 
+      # Path of ca pem file which should match a secretMount path
+      # pemPath: /ssl/ca.pem
     client:
 
       ## PEM that contains the client cert to connect to Elasticsearch.
       ##
       # pem:
 
+      # Path of client pem file which should match a secretMount path
+      # pemPath: /ssl/client.pem
+
       ## Private key for client auth when connecting to Elasticsearch
       ##
       # key:
 
+      # Path of client key file which should match a secretMount path
+      # keyPath: /ssl/client.key
 web:
   ## Path under which to expose metrics.
   ##

```

## 1.9.0

**Release date:** 2019-08-23

![AppVersion: 1.1.0](https://img.shields.io/static/v1?label=AppVersion&message=1.1.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/elasticsearch-exporter]: adding affinity support (#16571)

### Default value changes

```diff
diff --git a/charts/prometheus-elasticsearch-exporter/values.yaml b/charts/prometheus-elasticsearch-exporter/values.yaml
index 3231d088..3f156d7b 100644
--- a/charts/prometheus-elasticsearch-exporter/values.yaml
+++ b/charts/prometheus-elasticsearch-exporter/values.yaml
@@ -36,6 +36,8 @@ tolerations: {}
 
 podAnnotations: {}
 
+affinity: {}
+
 service:
   type: ClusterIP
   httpPort: 9108

```

## 1.8.1

**Release date:** 2019-08-21

![AppVersion: 1.1.0](https://img.shields.io/static/v1?label=AppVersion&message=1.1.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/elasticsearch-exporter]: adding myself as a maintainer (#16456)

### Default value changes

```diff
# No changes in this release
```

## 1.8.0

**Release date:** 2019-08-20

![AppVersion: 1.1.0](https://img.shields.io/static/v1?label=AppVersion&message=1.1.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/elasticsearch-exporter]: shards, snapshots, update to 1.1.0 (#16409)

### Default value changes

```diff
diff --git a/charts/prometheus-elasticsearch-exporter/values.yaml b/charts/prometheus-elasticsearch-exporter/values.yaml
index 4a1ebdca..3231d088 100644
--- a/charts/prometheus-elasticsearch-exporter/values.yaml
+++ b/charts/prometheus-elasticsearch-exporter/values.yaml
@@ -8,7 +8,7 @@ restartPolicy: Always
 
 image:
   repository: justwatch/elasticsearch_exporter
-  tag: 1.0.2
+  tag: 1.1.0
   pullPolicy: IfNotPresent
   pullSecret: ""
 
@@ -68,6 +68,18 @@ es:
   ##
   indices: true
 
+  ## If true, query stats for shards in the cluster.
+  ##
+  shards: true
+
+  ## If true, query stats for snapshots in the cluster.
+  ##
+  snapshots: true
+
+  ## If true, query stats for cluster settings.
+  ##
+  cluster_settings: false
+
   ## Timeout for trying to get stats from Elasticsearch. (ex: 20s)
   ##
   timeout: 30s

```

## 1.7.0

**Release date:** 2019-07-19

![AppVersion: 1.0.2](https://img.shields.io/static/v1?label=AppVersion&message=1.0.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Add possibility to provide a PrometheusRule resource (#15698)

### Default value changes

```diff
diff --git a/charts/prometheus-elasticsearch-exporter/values.yaml b/charts/prometheus-elasticsearch-exporter/values.yaml
index 74b93fa3..4a1ebdca 100644
--- a/charts/prometheus-elasticsearch-exporter/values.yaml
+++ b/charts/prometheus-elasticsearch-exporter/values.yaml
@@ -114,3 +114,34 @@ serviceMonitor:
   interval: 10s
   scrapeTimeout: 10s
   scheme: http
+
+prometheusRule:
+  ## If true, a PrometheusRule CRD is created for a prometheus operator
+  ## https://github.com/coreos/prometheus-operator
+  ##
+  enabled: false
+  #  namespace: monitoring
+  labels: {}
+  rules: []
+    # - record: elasticsearch_filesystem_data_used_percent
+    #   expr: 100 * (elasticsearch_filesystem_data_size_bytes - elasticsearch_filesystem_data_free_bytes)
+    #     / elasticsearch_filesystem_data_size_bytes
+    # - record: elasticsearch_filesystem_data_free_percent
+    #   expr: 100 - elasticsearch_filesystem_data_used_percent
+    # - alert: ElasticsearchTooFewNodesRunning
+    #   expr: elasticsearch_cluster_health_number_of_nodes < 3
+    #   for: 5m
+    #   labels:
+    #     severity: critical
+    #   annotations:
+    #     description: There are only {{$value}} < 3 ElasticSearch nodes running
+    #     summary: ElasticSearch running on less than 3 nodes
+    # - alert: ElasticsearchHeapTooHigh
+    #   expr: elasticsearch_jvm_memory_used_bytes{area="heap"} / elasticsearch_jvm_memory_max_bytes{area="heap"}
+    #     > 0.9
+    #   for: 15m
+    #   labels:
+    #     severity: critical
+    #   annotations:
+    #     description: The heap usage is over 90% for 15m
+    #     summary: ElasticSearch node {{$labels.node}} heap usage is high

```

## 1.6.0

**Release date:** 2019-07-12

![AppVersion: 1.0.2](https://img.shields.io/static/v1?label=AppVersion&message=1.0.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/elasticsearch-exporter] Add envs to deployment template (#13981)

### Default value changes

```diff
diff --git a/charts/prometheus-elasticsearch-exporter/values.yaml b/charts/prometheus-elasticsearch-exporter/values.yaml
index ff6c48bb..74b93fa3 100644
--- a/charts/prometheus-elasticsearch-exporter/values.yaml
+++ b/charts/prometheus-elasticsearch-exporter/values.yaml
@@ -44,6 +44,13 @@ service:
   annotations: {}
   labels: {}
 
+## Extra environment variables that will be passed into the exporter pod
+## example:
+## env:
+##   KEY_1: value1
+##   KEY_2: value2
+# env:
+
 es:
   ## Address (host and port) of the Elasticsearch node we should connect to.
   ## This could be a local node (localhost:9200, for instance), or the address

```

## 1.5.0

**Release date:** 2019-06-27

![AppVersion: 1.0.2](https://img.shields.io/static/v1?label=AppVersion&message=1.0.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/elasticsearch-exporter] Add imagePullSecret support (#14887)

### Default value changes

```diff
diff --git a/charts/prometheus-elasticsearch-exporter/values.yaml b/charts/prometheus-elasticsearch-exporter/values.yaml
index 03ac2531..ff6c48bb 100644
--- a/charts/prometheus-elasticsearch-exporter/values.yaml
+++ b/charts/prometheus-elasticsearch-exporter/values.yaml
@@ -10,6 +10,7 @@ image:
   repository: justwatch/elasticsearch_exporter
   tag: 1.0.2
   pullPolicy: IfNotPresent
+  pullSecret: ""
 
 ## Set enabled to false if you don't want securityContext
 ## in your Deployment.

```

## 1.4.1

**Release date:** 2019-06-20

![AppVersion: 1.0.2](https://img.shields.io/static/v1?label=AppVersion&message=1.0.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/elastic-exporter] Fix wrong default value (#13996)

### Default value changes

```diff
diff --git a/charts/prometheus-elasticsearch-exporter/values.yaml b/charts/prometheus-elasticsearch-exporter/values.yaml
index 8ace3c3d..03ac2531 100644
--- a/charts/prometheus-elasticsearch-exporter/values.yaml
+++ b/charts/prometheus-elasticsearch-exporter/values.yaml
@@ -49,7 +49,7 @@ es:
   ## of a remote Elasticsearch server. When basic auth is needed,
   ## specify as: <proto>://<user>:<password>@<host>:<port>. e.g., http://admin:pass@localhost:9200.
   ##
-  uri: localhost:9200
+  uri: http://localhost:9200
 
   ## If true, query stats for all nodes in the cluster, rather than just the
   ## node we connect to.

```

## 1.4.0

**Release date:** 2019-06-14

![AppVersion: 1.0.2](https://img.shields.io/static/v1?label=AppVersion&message=1.0.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [elasticsearch-exporter] Support ssl-skip-verify option (and make securitycontext configurable) (#14748)

### Default value changes

```diff
diff --git a/charts/prometheus-elasticsearch-exporter/values.yaml b/charts/prometheus-elasticsearch-exporter/values.yaml
index 0f4cc052..8ace3c3d 100644
--- a/charts/prometheus-elasticsearch-exporter/values.yaml
+++ b/charts/prometheus-elasticsearch-exporter/values.yaml
@@ -11,6 +11,14 @@ image:
   tag: 1.0.2
   pullPolicy: IfNotPresent
 
+## Set enabled to false if you don't want securityContext
+## in your Deployment.
+## The below values are the default for kubernetes.
+## Openshift won't deploy with runAsUser: 1000 without additional permissions.
+securityContext:
+  enabled: true  # Should be set to false when running on OpenShift
+  runAsUser: 1000
+
 resources: {}
   # requests:
   #   cpu: 100m
@@ -56,6 +64,12 @@ es:
   ##
   timeout: 30s
 
+  ## Skip SSL verification when connecting to Elasticsearch
+  ## (only available if image.tag >= 1.0.4rc1)
+  ##
+  sslSkipVerify: false
+
+
   ssl:
     ## If true, a secure connection to ES cluster is used (requires SSL certs below)
     ##

```

## 1.3.1

**Release date:** 2019-06-05

![AppVersion: 1.0.2](https://img.shields.io/static/v1?label=AppVersion&message=1.0.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* stable/elasticsearch-exporter: fix bug in template (#14486)

### Default value changes

```diff
# No changes in this release
```

## 1.3.0

**Release date:** 2019-06-03

![AppVersion: 1.0.2](https://img.shields.io/static/v1?label=AppVersion&message=1.0.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stabled/elasticsearch-exporter] add custom labels to the service and set service http port name (#14395)

### Default value changes

```diff
diff --git a/charts/prometheus-elasticsearch-exporter/values.yaml b/charts/prometheus-elasticsearch-exporter/values.yaml
index 310fc607..0f4cc052 100644
--- a/charts/prometheus-elasticsearch-exporter/values.yaml
+++ b/charts/prometheus-elasticsearch-exporter/values.yaml
@@ -30,7 +30,10 @@ podAnnotations: {}
 service:
   type: ClusterIP
   httpPort: 9108
+  metricsPort:
+    name: http
   annotations: {}
+  labels: {}
 
 es:
   ## Address (host and port) of the Elasticsearch node we should connect to.

```

## 1.2.0

**Release date:** 2019-04-18

![AppVersion: 1.0.2](https://img.shields.io/static/v1?label=AppVersion&message=1.0.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/elasticsearch-exporter] Extending Prometheus ServiceMonitor configuration (#12711)

### Default value changes

```diff
diff --git a/charts/prometheus-elasticsearch-exporter/values.yaml b/charts/prometheus-elasticsearch-exporter/values.yaml
index 8245f87e..310fc607 100644
--- a/charts/prometheus-elasticsearch-exporter/values.yaml
+++ b/charts/prometheus-elasticsearch-exporter/values.yaml
@@ -84,4 +84,8 @@ serviceMonitor:
   ## https://github.com/coreos/prometheus-operator
   ##
   enabled: false
+  #  namespace: monitoring
   labels: {}
+  interval: 10s
+  scrapeTimeout: 10s
+  scheme: http

```

## 1.1.3

**Release date:** 2019-03-11

![AppVersion: 1.0.2](https://img.shields.io/static/v1?label=AppVersion&message=1.0.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/elasticsearch-exporter] Fix ServiceMonitor to use service port name instead of service port number. (#11609)

### Default value changes

```diff
# No changes in this release
```

## 1.1.2

**Release date:** 2019-02-20

![AppVersion: 1.0.2](https://img.shields.io/static/v1?label=AppVersion&message=1.0.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/elasticsearch-exporter] Add scheme to the ServiceMonitor. Cannot fetch metrics from exporter if prometheus has a different default scheme. (#11421)

### Default value changes

```diff
# No changes in this release
```

## 1.1.1

**Release date:** 2019-02-11

![AppVersion: 1.0.2](https://img.shields.io/static/v1?label=AppVersion&message=1.0.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/elasticsearch-exporter] ServiceMonitor: set namespace to monitor + Add OWNER and add myself to owners. (#11323)

### Default value changes

```diff
# No changes in this release
```

## 1.1.0

**Release date:** 2019-01-26

![AppVersion: 1.0.2](https://img.shields.io/static/v1?label=AppVersion&message=1.0.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [elasticsearch-exporter] add ServiceMonitor (#10759)

### Default value changes

```diff
diff --git a/charts/prometheus-elasticsearch-exporter/values.yaml b/charts/prometheus-elasticsearch-exporter/values.yaml
index 16a74f7d..8245f87e 100644
--- a/charts/prometheus-elasticsearch-exporter/values.yaml
+++ b/charts/prometheus-elasticsearch-exporter/values.yaml
@@ -78,3 +78,10 @@ web:
   ## Path under which to expose metrics.
   ##
   path: /metrics
+
+serviceMonitor:
+  ## If true, a ServiceMonitor CRD is created for a prometheus operator
+  ## https://github.com/coreos/prometheus-operator
+  ##
+  enabled: false
+  labels: {}

```

## 1.0.0

**Release date:** 2019-01-18

![AppVersion: 1.0.2](https://img.shields.io/static/v1?label=AppVersion&message=1.0.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* fix fullname (#10752)

### Default value changes

```diff
# No changes in this release
```

## 0.4.1

**Release date:** 2018-10-30

![AppVersion: 1.0.2](https://img.shields.io/static/v1?label=AppVersion&message=1.0.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/elasticsearch-exporter] Fix typo - Add Readiness probe to exporter deployment (#8799)

### Default value changes

```diff
# No changes in this release
```

## 0.4.0

**Release date:** 2018-09-29

![AppVersion: 1.0.2](https://img.shields.io/static/v1?label=AppVersion&message=1.0.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Add support for deployment tolerations and update helpers (#7706)

### Default value changes

```diff
diff --git a/charts/prometheus-elasticsearch-exporter/values.yaml b/charts/prometheus-elasticsearch-exporter/values.yaml
index 949b4415..16a74f7d 100644
--- a/charts/prometheus-elasticsearch-exporter/values.yaml
+++ b/charts/prometheus-elasticsearch-exporter/values.yaml
@@ -23,6 +23,8 @@ priorityClassName: ""
 
 nodeSelector: {}
 
+tolerations: {}
+
 podAnnotations: {}
 
 service:

```

## 0.3.0

**Release date:** 2018-09-25

![AppVersion: 1.0.2](https://img.shields.io/static/v1?label=AppVersion&message=1.0.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/elasticsearch-exporter] Adding pod annotations (#7960)

### Default value changes

```diff
diff --git a/charts/prometheus-elasticsearch-exporter/values.yaml b/charts/prometheus-elasticsearch-exporter/values.yaml
index eaa3d442..949b4415 100644
--- a/charts/prometheus-elasticsearch-exporter/values.yaml
+++ b/charts/prometheus-elasticsearch-exporter/values.yaml
@@ -23,6 +23,8 @@ priorityClassName: ""
 
 nodeSelector: {}
 
+podAnnotations: {}
+
 service:
   type: ClusterIP
   httpPort: 9108

```

## 0.2.2

**Release date:** 2018-08-31

![AppVersion: 1.0.2](https://img.shields.io/static/v1?label=AppVersion&message=1.0.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Fix nodeSelector indent (#7458)

### Default value changes

```diff
# No changes in this release
```

## 0.2.1

**Release date:** 2018-08-31

![AppVersion: 1.0.2](https://img.shields.io/static/v1?label=AppVersion&message=1.0.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Add nodeSelector to make pods run on specific nodes (#7456)

### Default value changes

```diff
diff --git a/charts/prometheus-elasticsearch-exporter/values.yaml b/charts/prometheus-elasticsearch-exporter/values.yaml
index 25ec4e69..eaa3d442 100644
--- a/charts/prometheus-elasticsearch-exporter/values.yaml
+++ b/charts/prometheus-elasticsearch-exporter/values.yaml
@@ -21,6 +21,8 @@ resources: {}
 
 priorityClassName: ""
 
+nodeSelector: {}
+
 service:
   type: ClusterIP
   httpPort: 9108

```

## 0.2.0

**Release date:** 2018-07-15

![AppVersion: 1.0.2](https://img.shields.io/static/v1?label=AppVersion&message=1.0.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/elasticsearch-exporter] add support for priorityClasses (#6584)

### Default value changes

```diff
diff --git a/charts/prometheus-elasticsearch-exporter/values.yaml b/charts/prometheus-elasticsearch-exporter/values.yaml
index f6572cc9..25ec4e69 100644
--- a/charts/prometheus-elasticsearch-exporter/values.yaml
+++ b/charts/prometheus-elasticsearch-exporter/values.yaml
@@ -19,6 +19,8 @@ resources: {}
   #   cpu: 100m
   #   memory: 128Mi
 
+priorityClassName: ""
+
 service:
   type: ClusterIP
   httpPort: 9108

```

## 0.1.4

**Release date:** 2018-04-19

![AppVersion: 1.0.2](https://img.shields.io/static/v1?label=AppVersion&message=1.0.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Typo fix in elasticsearch-exporter/README.md: a->an (#4988)

### Default value changes

```diff
# No changes in this release
```

## 0.1.3

**Release date:** 2018-04-07

![AppVersion: 1.0.2](https://img.shields.io/static/v1?label=AppVersion&message=1.0.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Add home key for elasticsearch-exporter (#4756)

### Default value changes

```diff
# No changes in this release
```

## 0.1.2

**Release date:** 2018-03-01

![AppVersion: 1.0.2](https://img.shields.io/static/v1?label=AppVersion&message=1.0.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* prometheus service annotations for elastisearch-exporter (#3842)

### Default value changes

```diff
diff --git a/charts/prometheus-elasticsearch-exporter/values.yaml b/charts/prometheus-elasticsearch-exporter/values.yaml
index b603bca1..f6572cc9 100644
--- a/charts/prometheus-elasticsearch-exporter/values.yaml
+++ b/charts/prometheus-elasticsearch-exporter/values.yaml
@@ -22,6 +22,7 @@ resources: {}
 service:
   type: ClusterIP
   httpPort: 9108
+  annotations: {}
 
 es:
   ## Address (host and port) of the Elasticsearch node we should connect to.

```

## 0.1.1

**Release date:** 2018-02-10

![AppVersion: 1.0.2](https://img.shields.io/static/v1?label=AppVersion&message=1.0.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Adjust README to use stable chart repo (#3533)

### Default value changes

```diff
# No changes in this release
```

## 0.1.0

**Release date:** 2018-02-02

![AppVersion: 1.0.2](https://img.shields.io/static/v1?label=AppVersion&message=1.0.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Add chart elasticsearch-exporter (#2525)

### Default value changes

```diff
## number of exporter instances
##
replicaCount: 1

## restart policy for all containers
##
restartPolicy: Always

image:
  repository: justwatch/elasticsearch_exporter
  tag: 1.0.2
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
  httpPort: 9108

es:
  ## Address (host and port) of the Elasticsearch node we should connect to.
  ## This could be a local node (localhost:9200, for instance), or the address
  ## of a remote Elasticsearch server. When basic auth is needed,
  ## specify as: <proto>://<user>:<password>@<host>:<port>. e.g., http://admin:pass@localhost:9200.
  ##
  uri: localhost:9200

  ## If true, query stats for all nodes in the cluster, rather than just the
  ## node we connect to.
  ##
  all: true

  ## If true, query stats for all indices in the cluster.
  ##
  indices: true

  ## Timeout for trying to get stats from Elasticsearch. (ex: 20s)
  ##
  timeout: 30s

  ssl:
    ## If true, a secure connection to ES cluster is used (requires SSL certs below)
    ##
    enabled: false

    ca:

      ## PEM that contains trusted CAs used for setting up secure Elasticsearch connection
      ##
      # pem:

    client:

      ## PEM that contains the client cert to connect to Elasticsearch.
      ##
      # pem:

      ## Private key for client auth when connecting to Elasticsearch
      ##
      # key:

web:
  ## Path under which to expose metrics.
  ##
  path: /metrics

```

---
Autogenerated from Helm Chart and git history using [helm-changelog](https://github.com/mogensen/helm-changelog)
