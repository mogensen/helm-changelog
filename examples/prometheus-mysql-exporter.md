# Change Log

## 1.14.0

**Release date:** 2023-04-26

![AppVersion: v0.14.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.14.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-mysql-exporter] Add option for extraArgs in cloudsqlproxy (#3284)

### Default value changes

```diff
diff --git a/charts/prometheus-mysql-exporter/values.yaml b/charts/prometheus-mysql-exporter/values.yaml
index 36866c9e..503099fe 100644
--- a/charts/prometheus-mysql-exporter/values.yaml
+++ b/charts/prometheus-mysql-exporter/values.yaml
@@ -172,3 +172,4 @@ cloudsqlproxy:
   workloadIdentity:
     enabled: false
     serviceAccountEmail: ""
+  extraArgs: ""

```

## 1.13.0

**Release date:** 2023-03-06

![AppVersion: v0.14.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.14.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Allow overriding of image registry and add quay.io default (#3099)

### Default value changes

```diff
diff --git a/charts/prometheus-mysql-exporter/values.yaml b/charts/prometheus-mysql-exporter/values.yaml
index 77f234c0..36866c9e 100644
--- a/charts/prometheus-mysql-exporter/values.yaml
+++ b/charts/prometheus-mysql-exporter/values.yaml
@@ -8,7 +8,8 @@ fullnameOverride: ""
 replicaCount: 1
 
 image:
-  repository: "prom/mysqld-exporter"
+  registry: quay.io
+  repository: prometheus/mysqld-exporter
   ## if not set charts appVersion var is used
   tag: ""
   pullPolicy: "IfNotPresent"

```

## 1.12.1

**Release date:** 2023-01-11

![AppVersion: v0.14.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.14.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-mysql-exporter] add imagePullSecrets and fullname override (#2905)

### Default value changes

```diff
diff --git a/charts/prometheus-mysql-exporter/values.yaml b/charts/prometheus-mysql-exporter/values.yaml
index 559b8ad5..77f234c0 100644
--- a/charts/prometheus-mysql-exporter/values.yaml
+++ b/charts/prometheus-mysql-exporter/values.yaml
@@ -1,14 +1,22 @@
-# Default values for prometheus-mysql-exporter.
-# This is a YAML-formatted file.
-# Declare variables to be passed into your templates.
+## Default values for prometheus-mysql-exporter.
+## This is a YAML-formatted file.
+## Declare variables to be passed into your templates.
+
+## override release name
+fullnameOverride: ""
 
 replicaCount: 1
 
 image:
   repository: "prom/mysqld-exporter"
-  tag: "v0.14.0"
+  ## if not set charts appVersion var is used
+  tag: ""
   pullPolicy: "IfNotPresent"
 
+# imagePullSecrets:
+# - name: secret-name
+imagePullSecrets: []
+
 service:
   labels: {}
   annotations: {}

```

## 1.11.1

**Release date:** 2022-12-01

![AppVersion: v0.14.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.14.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-mysql-exporter] fix error in extraVolumes (#2770)

### Default value changes

```diff
# No changes in this release
```

## 1.11.0

**Release date:** 2022-12-01

![AppVersion: v0.14.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.14.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-mysql-exporter] add extraVolumes and extraVolumeMounts to chart (#2764)

### Default value changes

```diff
diff --git a/charts/prometheus-mysql-exporter/values.yaml b/charts/prometheus-mysql-exporter/values.yaml
index efd628da..559b8ad5 100644
--- a/charts/prometheus-mysql-exporter/values.yaml
+++ b/charts/prometheus-mysql-exporter/values.yaml
@@ -62,6 +62,17 @@ affinity: {}
 
 podLabels: {}
 
+# Extra Volume Mounts for the mysql exporter container
+extraVolumeMounts: []
+# - name: example
+#   mountPath: /example
+
+# Extra Volumes for the pod
+extraVolumes: []
+# - name: example
+#   configMap:
+#     name: example
+
 podSecurityContext: {}
   # fsGroup: 65534
 

```

## 1.10.0

**Release date:** 2022-11-08

![AppVersion: v0.14.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.14.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-mysql-exporter] Support Workload Identity (#2645)

### Default value changes

```diff
diff --git a/charts/prometheus-mysql-exporter/values.yaml b/charts/prometheus-mysql-exporter/values.yaml
index c4c77375..efd628da 100644
--- a/charts/prometheus-mysql-exporter/values.yaml
+++ b/charts/prometheus-mysql-exporter/values.yaml
@@ -141,7 +141,7 @@ cloudsqlproxy:
   enabled: false
   image:
     repo: "gcr.io/cloudsql-docker/gce-proxy"
-    tag: "1.30.1-alpine"
+    tag: "1.33.0-alpine"
     pullPolicy: "IfNotPresent"
   instanceConnectionName: "project:us-central1:dbname"
   ipAddressTypes: ""
@@ -149,3 +149,6 @@ cloudsqlproxy:
   credentialsSecret: ""
   # service account json
   credentials: ""
+  workloadIdentity:
+    enabled: false
+    serviceAccountEmail: ""

```

## 1.9.1

**Release date:** 2022-10-19

![AppVersion: v0.14.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.14.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-mysql-exporter] Fix typo in README.md (#2587)

### Default value changes

```diff
# No changes in this release
```

## 1.9.0

**Release date:** 2022-08-03

![AppVersion: v0.14.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.14.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-mysql-exporter] Add option for ipAddressTypes (#2339)

### Default value changes

```diff
diff --git a/charts/prometheus-mysql-exporter/values.yaml b/charts/prometheus-mysql-exporter/values.yaml
index 861a6c26..c4c77375 100644
--- a/charts/prometheus-mysql-exporter/values.yaml
+++ b/charts/prometheus-mysql-exporter/values.yaml
@@ -144,6 +144,7 @@ cloudsqlproxy:
     tag: "1.30.1-alpine"
     pullPolicy: "IfNotPresent"
   instanceConnectionName: "project:us-central1:dbname"
+  ipAddressTypes: ""
   port: "3306"
   credentialsSecret: ""
   # service account json

```

## 1.8.1

**Release date:** 2022-06-01

![AppVersion: v0.14.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.14.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-mysql-exporter] add missing collect parameter (#2103)

### Default value changes

```diff
# No changes in this release
```

## 1.8.0

**Release date:** 2022-05-31

![AppVersion: v0.14.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.14.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-mysql-exporter] add logLevel and logFormat options (#2099)

### Default value changes

```diff
diff --git a/charts/prometheus-mysql-exporter/values.yaml b/charts/prometheus-mysql-exporter/values.yaml
index 645c094a..861a6c26 100644
--- a/charts/prometheus-mysql-exporter/values.yaml
+++ b/charts/prometheus-mysql-exporter/values.yaml
@@ -6,7 +6,7 @@ replicaCount: 1
 
 image:
   repository: "prom/mysqld-exporter"
-  tag: "v0.12.1"
+  tag: "v0.14.0"
   pullPolicy: "IfNotPresent"
 
 service:
@@ -73,12 +73,16 @@ securityContext: {}
   # runAsNonRoot: true
   # runAsUser: 65534
 
-
 annotations:
   prometheus.io/scrape: "true"
   prometheus.io/path: "/metrics"
   prometheus.io/port: "9104"
 
+config: {}
+  # Allow to set specifc options on the exporter
+  # logLevel: info
+  # logFormat: "logger:stderr"
+
 collectors: {}
   # auto_increment.columns: false
   # binlog_size: false
@@ -137,21 +141,10 @@ cloudsqlproxy:
   enabled: false
   image:
     repo: "gcr.io/cloudsql-docker/gce-proxy"
-    tag: "1.19.1-alpine"
+    tag: "1.30.1-alpine"
     pullPolicy: "IfNotPresent"
   instanceConnectionName: "project:us-central1:dbname"
   port: "3306"
   credentialsSecret: ""
+  # service account json
   credentials: ""
-    # {
-    #   "type": "service_account",
-    #   "project_id": "project",
-    #   "private_key_id": "KEYID1",
-    #   "private_key": "-----BEGIN PRIVATE KEY-----\sdajsdnasd\n-----END PRIVATE KEY-----\n",
-    #   "client_email": "user@project.iam.gserviceaccount.com",
-    #   "client_id": "111111111",
-    #   "auth_uri": "https://accounts.google.com/o/oauth2/auth",
-    #   "token_uri": "https://accounts.google.com/o/oauth2/token",
-    #   "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
-    #   "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/user%40project.iam.gserviceaccount.com"
-    # }

```

## 1.7.0

**Release date:** 2022-03-21

![AppVersion: v0.12.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.12.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-mysql-exporter] Add relabelings to servicemonitor (#1891)

### Default value changes

```diff
diff --git a/charts/prometheus-mysql-exporter/values.yaml b/charts/prometheus-mysql-exporter/values.yaml
index 7cedf176..645c094a 100644
--- a/charts/prometheus-mysql-exporter/values.yaml
+++ b/charts/prometheus-mysql-exporter/values.yaml
@@ -24,13 +24,15 @@ serviceMonitor:
   # interval: 30s
   # scrapeTimeout is the timeout after which the scrape is ended
   # scrapeTimeout: 10s
-  # additionalLabels is the set of additional labels to add to the ServiceMonitor
   # namespace: monitoring
+  # additionalLabels is the set of additional labels to add to the ServiceMonitor
   additionalLabels: {}
   jobLabel: ""
   targetLabels: []
   podTargetLabels: []
   metricRelabelings: []
+  # Set relabel_configs as per https://prometheus.io/docs/prometheus/latest/configuration/configuration/#relabel_config
+  relabelings: []
 
 serviceAccount:
   # Specifies whether a ServiceAccount should be created

```

## 1.6.0

**Release date:** 2022-02-14

![AppVersion: v0.12.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.12.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-mysql-exporter] Add missing namespaceSelector to servicemonitor (#1786)

### Default value changes

```diff
# No changes in this release
```

## 1.5.0

**Release date:** 2021-12-30

![AppVersion: v0.12.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.12.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-mysql-exporter] add annotations to service account (#1649)

### Default value changes

```diff
diff --git a/charts/prometheus-mysql-exporter/values.yaml b/charts/prometheus-mysql-exporter/values.yaml
index 2b16ae46..7cedf176 100644
--- a/charts/prometheus-mysql-exporter/values.yaml
+++ b/charts/prometheus-mysql-exporter/values.yaml
@@ -38,6 +38,7 @@ serviceAccount:
   # The name of the ServiceAccount to use.
   # If not set and create is true, a name is generated using the fullname template
   name:
+  annotations: {}
 
 resources: {}
   # We usually recommend not to specify default resources and to leave this as a conscious

```

## 1.4.0

**Release date:** 2021-12-23

![AppVersion: v0.12.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.12.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-mysql-exporter] Add namespace support to service monitor of MySQL exporter (#1630)

### Default value changes

```diff
diff --git a/charts/prometheus-mysql-exporter/values.yaml b/charts/prometheus-mysql-exporter/values.yaml
index 20678b52..2b16ae46 100644
--- a/charts/prometheus-mysql-exporter/values.yaml
+++ b/charts/prometheus-mysql-exporter/values.yaml
@@ -25,6 +25,7 @@ serviceMonitor:
   # scrapeTimeout is the timeout after which the scrape is ended
   # scrapeTimeout: 10s
   # additionalLabels is the set of additional labels to add to the ServiceMonitor
+  # namespace: monitoring
   additionalLabels: {}
   jobLabel: ""
   targetLabels: []

```

## 1.3.0

**Release date:** 2021-10-30

![AppVersion: v0.12.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.12.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-mysql-exporter] Add proxysql credentials as secret (#1469)

### Default value changes

```diff
diff --git a/charts/prometheus-mysql-exporter/values.yaml b/charts/prometheus-mysql-exporter/values.yaml
index c735d6dc..20678b52 100644
--- a/charts/prometheus-mysql-exporter/values.yaml
+++ b/charts/prometheus-mysql-exporter/values.yaml
@@ -137,6 +137,7 @@ cloudsqlproxy:
     pullPolicy: "IfNotPresent"
   instanceConnectionName: "project:us-central1:dbname"
   port: "3306"
+  credentialsSecret: ""
   credentials: ""
     # {
     #   "type": "service_account",

```

## 1.2.2

**Release date:** 2021-08-23

![AppVersion: v0.12.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.12.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-mysql-exporter] Restart pod on password change (#1262)

### Default value changes

```diff
# No changes in this release
```

## 1.2.1

**Release date:** 2021-06-29

![AppVersion: v0.12.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.12.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-mysql-exporter] Add ServiceAccount (#1087)

### Default value changes

```diff
diff --git a/charts/prometheus-mysql-exporter/values.yaml b/charts/prometheus-mysql-exporter/values.yaml
index be5bfdab..c735d6dc 100644
--- a/charts/prometheus-mysql-exporter/values.yaml
+++ b/charts/prometheus-mysql-exporter/values.yaml
@@ -31,6 +31,13 @@ serviceMonitor:
   podTargetLabels: []
   metricRelabelings: []
 
+serviceAccount:
+  # Specifies whether a ServiceAccount should be created
+  create: false
+  # The name of the ServiceAccount to use.
+  # If not set and create is true, a name is generated using the fullname template
+  name:
+
 resources: {}
   # We usually recommend not to specify default resources and to leave this as a conscious
   # choice for the user. This also increases chances charts run on environments with little

```

## 1.2.0

**Release date:** 2021-04-26

![AppVersion: v0.12.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.12.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* make existing secret config more flexible (#705)

### Default value changes

```diff
diff --git a/charts/prometheus-mysql-exporter/values.yaml b/charts/prometheus-mysql-exporter/values.yaml
index c1a3cd77..be5bfdab 100644
--- a/charts/prometheus-mysql-exporter/values.yaml
+++ b/charts/prometheus-mysql-exporter/values.yaml
@@ -114,7 +114,12 @@ mysql:
   port: 3306
   protocol: ""
   user: "exporter"
-  existingSecret: false
+  # secret with full DATA_SOURCE_NAME env var as stringdata
+  existingSecret: ""
+  # secret only containing the password
+  existingPasswordSecret:
+    name: ""
+    key: ""
 
 # cloudsqlproxy https://cloud.google.com/sql/docs/mysql/sql-proxy
 cloudsqlproxy:

```

## 1.1.0

**Release date:** 2021-02-17

![AppVersion: v0.12.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.12.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-mysql-exporter] Add securityContext and podSecurityContext (#662)

### Default value changes

```diff
diff --git a/charts/prometheus-mysql-exporter/values.yaml b/charts/prometheus-mysql-exporter/values.yaml
index f77038b7..c1a3cd77 100644
--- a/charts/prometheus-mysql-exporter/values.yaml
+++ b/charts/prometheus-mysql-exporter/values.yaml
@@ -51,6 +51,18 @@ affinity: {}
 
 podLabels: {}
 
+podSecurityContext: {}
+  # fsGroup: 65534
+
+securityContext: {}
+  # capabilities:
+  #   drop:
+  #   - ALL
+  # readOnlyRootFilesystem: true
+  # runAsNonRoot: true
+  # runAsUser: 65534
+
+
 annotations:
   prometheus.io/scrape: "true"
   prometheus.io/path: "/metrics"

```

## 1.0.2

**Release date:** 2021-02-08

![AppVersion: v0.12.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.12.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-mysql-exporter] fixed link to Prometheus MySQL Exporter repo (#630)

### Default value changes

```diff
# No changes in this release
```

## 1.0.1

**Release date:** 2020-12-08

![AppVersion: v0.12.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.12.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* fix: use alpine cloudsqlproxy image (#464)

### Default value changes

```diff
diff --git a/charts/prometheus-mysql-exporter/values.yaml b/charts/prometheus-mysql-exporter/values.yaml
index 5c0b107a..f77038b7 100644
--- a/charts/prometheus-mysql-exporter/values.yaml
+++ b/charts/prometheus-mysql-exporter/values.yaml
@@ -109,7 +109,7 @@ cloudsqlproxy:
   enabled: false
   image:
     repo: "gcr.io/cloudsql-docker/gce-proxy"
-    tag: "1.16"
+    tag: "1.19.1-alpine"
     pullPolicy: "IfNotPresent"
   instanceConnectionName: "project:us-central1:dbname"
   port: "3306"

```

## 1.0.0

**Release date:** 2020-10-11

![AppVersion: v0.12.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.12.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-mysql-exporter] update docker image, readme links & use default labels (#68)

### Default value changes

```diff
diff --git a/charts/prometheus-mysql-exporter/values.yaml b/charts/prometheus-mysql-exporter/values.yaml
index b67a8297..5c0b107a 100644
--- a/charts/prometheus-mysql-exporter/values.yaml
+++ b/charts/prometheus-mysql-exporter/values.yaml
@@ -6,7 +6,7 @@ replicaCount: 1
 
 image:
   repository: "prom/mysqld-exporter"
-  tag: "v0.11.0"
+  tag: "v0.12.1"
   pullPolicy: "IfNotPresent"
 
 service:
@@ -109,20 +109,20 @@ cloudsqlproxy:
   enabled: false
   image:
     repo: "gcr.io/cloudsql-docker/gce-proxy"
-    tag: "1.14"
+    tag: "1.16"
     pullPolicy: "IfNotPresent"
   instanceConnectionName: "project:us-central1:dbname"
   port: "3306"
-  credentials:
-    '{
-      "type": "service_account",
-      "project_id": "project",
-      "private_key_id": "KEYID1",
-      "private_key": "-----BEGIN PRIVATE KEY-----\sdajsdnasd\n-----END PRIVATE KEY-----\n",
-      "client_email": "user@project.iam.gserviceaccount.com",
-      "client_id": "111111111",
-      "auth_uri": "https://accounts.google.com/o/oauth2/auth",
-      "token_uri": "https://accounts.google.com/o/oauth2/token",
-      "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
-      "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/user%40project.iam.gserviceaccount.com"
-    }'
+  credentials: ""
+    # {
+    #   "type": "service_account",
+    #   "project_id": "project",
+    #   "private_key_id": "KEYID1",
+    #   "private_key": "-----BEGIN PRIVATE KEY-----\sdajsdnasd\n-----END PRIVATE KEY-----\n",
+    #   "client_email": "user@project.iam.gserviceaccount.com",
+    #   "client_id": "111111111",
+    #   "auth_uri": "https://accounts.google.com/o/oauth2/auth",
+    #   "token_uri": "https://accounts.google.com/o/oauth2/token",
+    #   "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
+    #   "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/user%40project.iam.gserviceaccount.com"
+    # }

```

## 0.7.1

**Release date:** 2020-08-20

![AppVersion: v0.11.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.11.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Prep initial charts indexing (#14)

### Default value changes

```diff
# No changes in this release
```

## 0.7.0

**Release date:** 2020-07-27

![AppVersion: v0.11.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.11.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-mysql-exporter] Add ServiceMonitor metricRelabeling (#23352)

### Default value changes

```diff
diff --git a/charts/prometheus-mysql-exporter/values.yaml b/charts/prometheus-mysql-exporter/values.yaml
index d308647d..b67a8297 100644
--- a/charts/prometheus-mysql-exporter/values.yaml
+++ b/charts/prometheus-mysql-exporter/values.yaml
@@ -29,6 +29,7 @@ serviceMonitor:
   jobLabel: ""
   targetLabels: []
   podTargetLabels: []
+  metricRelabelings: []
 
 resources: {}
   # We usually recommend not to specify default resources and to leave this as a conscious

```

## 0.6.0

**Release date:** 2020-06-29

![AppVersion: v0.11.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.11.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-mysql-exporter]:  Add service labels and annotations (#22986)

### Default value changes

```diff
diff --git a/charts/prometheus-mysql-exporter/values.yaml b/charts/prometheus-mysql-exporter/values.yaml
index 9b8afd9f..d308647d 100644
--- a/charts/prometheus-mysql-exporter/values.yaml
+++ b/charts/prometheus-mysql-exporter/values.yaml
@@ -10,6 +10,8 @@ image:
   pullPolicy: "IfNotPresent"
 
 service:
+  labels: {}
+  annotations: {}
   name: mysql-exporter
   type: ClusterIP
   externalPort: 9104

```

## 0.5.3

**Release date:** 2020-06-05

![AppVersion: v0.11.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.11.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-mysql-exporter]: fix docs typo (#22674)

### Default value changes

```diff
# No changes in this release
```

## 0.5.2

**Release date:** 2019-10-25

![AppVersion: v0.11.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.11.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [mysqld-exporter] Improve Secrets (#18314)

### Default value changes

```diff
diff --git a/charts/prometheus-mysql-exporter/values.yaml b/charts/prometheus-mysql-exporter/values.yaml
index 1f2309f6..9b8afd9f 100644
--- a/charts/prometheus-mysql-exporter/values.yaml
+++ b/charts/prometheus-mysql-exporter/values.yaml
@@ -99,6 +99,7 @@ mysql:
   port: 3306
   protocol: ""
   user: "exporter"
+  existingSecret: false
 
 # cloudsqlproxy https://cloud.google.com/sql/docs/mysql/sql-proxy
 cloudsqlproxy:

```

## 0.5.1

**Release date:** 2019-07-09

![AppVersion: v0.11.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.11.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Mysql exporter managing labels (#15349)

### Default value changes

```diff
diff --git a/charts/prometheus-mysql-exporter/values.yaml b/charts/prometheus-mysql-exporter/values.yaml
index d1e48a4c..1f2309f6 100644
--- a/charts/prometheus-mysql-exporter/values.yaml
+++ b/charts/prometheus-mysql-exporter/values.yaml
@@ -24,6 +24,9 @@ serviceMonitor:
   # scrapeTimeout: 10s
   # additionalLabels is the set of additional labels to add to the ServiceMonitor
   additionalLabels: {}
+  jobLabel: ""
+  targetLabels: []
+  podTargetLabels: []
 
 resources: {}
   # We usually recommend not to specify default resources and to leave this as a conscious

```

## 0.5.0

**Release date:** 2019-06-23

![AppVersion: v0.11.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.11.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* feat(stable/prometheus-mysql-exporter): add optional ServiceMonitor (#14600)

### Default value changes

```diff
diff --git a/charts/prometheus-mysql-exporter/values.yaml b/charts/prometheus-mysql-exporter/values.yaml
index c0441410..d1e48a4c 100644
--- a/charts/prometheus-mysql-exporter/values.yaml
+++ b/charts/prometheus-mysql-exporter/values.yaml
@@ -15,6 +15,16 @@ service:
   externalPort: 9104
   internalPort: 9104
 
+serviceMonitor:
+  # enabled should be set to true to enable prometheus-operator discovery of this service
+  enabled: false
+  # interval is the interval at which metrics should be scraped
+  # interval: 30s
+  # scrapeTimeout is the timeout after which the scrape is ended
+  # scrapeTimeout: 10s
+  # additionalLabels is the set of additional labels to add to the ServiceMonitor
+  additionalLabels: {}
+
 resources: {}
   # We usually recommend not to specify default resources and to leave this as a conscious
   # choice for the user. This also increases chances charts run on environments with little

```

## 0.4.0

**Release date:** 2019-06-20

![AppVersion: v0.11.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.11.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* feat(stable/prometheus-mysql-exporter): store MySQL password in a k8s Secret (#14712)

### Default value changes

```diff
# No changes in this release
```

## 0.3.4

**Release date:** 2019-06-16

![AppVersion: v0.11.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.11.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-mysql-exporter] updated gce-proxy to 1.14 & added myself to maintainers & owners (#14776)

### Default value changes

```diff
diff --git a/charts/prometheus-mysql-exporter/values.yaml b/charts/prometheus-mysql-exporter/values.yaml
index 7f4828f5..c0441410 100644
--- a/charts/prometheus-mysql-exporter/values.yaml
+++ b/charts/prometheus-mysql-exporter/values.yaml
@@ -92,7 +92,7 @@ cloudsqlproxy:
   enabled: false
   image:
     repo: "gcr.io/cloudsql-docker/gce-proxy"
-    tag: "1.11"
+    tag: "1.14"
     pullPolicy: "IfNotPresent"
   instanceConnectionName: "project:us-central1:dbname"
   port: "3306"

```

## 0.3.3

**Release date:** 2019-06-11

![AppVersion: v0.11.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.11.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-mysql-exporter] Fix invalid YAML with non-boolean collectors (#14701)

### Default value changes

```diff
# No changes in this release
```

## 0.3.2

**Release date:** 2019-03-26

![AppVersion: v0.11.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.11.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* fixed readme (#11660)

### Default value changes

```diff
# No changes in this release
```

## 0.3.1

**Release date:** 2019-03-25

![AppVersion: v0.11.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.11.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-mysql-importer] Add optional labels to each pod  (#11380)

### Default value changes

```diff
diff --git a/charts/prometheus-mysql-exporter/values.yaml b/charts/prometheus-mysql-exporter/values.yaml
index 7fd4de24..7f4828f5 100644
--- a/charts/prometheus-mysql-exporter/values.yaml
+++ b/charts/prometheus-mysql-exporter/values.yaml
@@ -33,6 +33,8 @@ tolerations: []
 
 affinity: {}
 
+podLabels: {}
+
 annotations:
   prometheus.io/scrape: "true"
   prometheus.io/path: "/metrics"

```

## 0.3.0

**Release date:** 2019-03-05

![AppVersion: v0.11.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.11.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-mysql-exporter] Collector flags config (#11569)

### Default value changes

```diff
diff --git a/charts/prometheus-mysql-exporter/values.yaml b/charts/prometheus-mysql-exporter/values.yaml
index 77fac30a..7fd4de24 100644
--- a/charts/prometheus-mysql-exporter/values.yaml
+++ b/charts/prometheus-mysql-exporter/values.yaml
@@ -38,6 +38,43 @@ annotations:
   prometheus.io/path: "/metrics"
   prometheus.io/port: "9104"
 
+collectors: {}
+  # auto_increment.columns: false
+  # binlog_size: false
+  # engine_innodb_status: false
+  # engine_tokudb_status: false
+  # global_status: true
+  # global_variables: true
+  # info_schema.clientstats: false
+  # info_schema.innodb_metrics: false
+  # info_schema.innodb_tablespaces: false
+  # info_schema.innodb_cmp: false
+  # info_schema.innodb_cmpmem: false
+  # info_schema.processlist: false
+  # info_schema.processlist.min_time: 0
+  # info_schema.query_response_time: false
+  # info_schema.tables: true
+  # info_schema.tables.databases: '*'
+  # info_schema.tablestats: false
+  # info_schema.schemastats: false
+  # info_schema.userstats: false
+  # perf_schema.eventsstatements: false
+  # perf_schema.eventsstatements.digest_text_limit: 120
+  # perf_schema.eventsstatements.limit: false
+  # perf_schema.eventsstatements.timelimit: 86400
+  # perf_schema.eventswaits: false
+  # perf_schema.file_events: false
+  # perf_schema.file_instances: false
+  # perf_schema.indexiowaits: false
+  # perf_schema.tableiowaits: false
+  # perf_schema.tablelocks: false
+  # perf_schema.replication_group_member_stats: false
+  # slave_status: true
+  # slave_hosts: false
+  # heartbeat: false
+  # heartbeat.database: heartbeat
+  # heartbeat.table: heartbeat
+
 # mysql connection params which build the DATA_SOURCE_NAME env var of the docker container
 mysql:
   db: ""

```

## 0.2.1

**Release date:** 2018-10-23

![AppVersion: v0.11.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.11.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* fixed annotations (#8431)

### Default value changes

```diff
diff --git a/charts/prometheus-mysql-exporter/values.yaml b/charts/prometheus-mysql-exporter/values.yaml
index 58387d5d..77fac30a 100644
--- a/charts/prometheus-mysql-exporter/values.yaml
+++ b/charts/prometheus-mysql-exporter/values.yaml
@@ -36,7 +36,7 @@ affinity: {}
 annotations:
   prometheus.io/scrape: "true"
   prometheus.io/path: "/metrics"
-  prometheus.io/port: 9104
+  prometheus.io/port: "9104"
 
 # mysql connection params which build the DATA_SOURCE_NAME env var of the docker container
 mysql:

```

## 0.2.0

**Release date:** 2018-10-13

![AppVersion: v0.11.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.11.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-mysql-exporter] split datasource in separate mysql connection vars (#8000)

### Default value changes

```diff
diff --git a/charts/prometheus-mysql-exporter/values.yaml b/charts/prometheus-mysql-exporter/values.yaml
index fa1d9c99..58387d5d 100644
--- a/charts/prometheus-mysql-exporter/values.yaml
+++ b/charts/prometheus-mysql-exporter/values.yaml
@@ -6,7 +6,7 @@ replicaCount: 1
 
 image:
   repository: "prom/mysqld-exporter"
-  tag: "v0.10.0"
+  tag: "v0.11.0"
   pullPolicy: "IfNotPresent"
 
 service:
@@ -33,13 +33,20 @@ tolerations: []
 
 affinity: {}
 
-annotations: {}
-#  prometheus.io/scrape: "true"
-#  prometheus.io/path: "/metrics"
-#  prometheus.io/port: "{{ .Values.service.internalPort }}"
+annotations:
+  prometheus.io/scrape: "true"
+  prometheus.io/path: "/metrics"
+  prometheus.io/port: 9104
 
-# Scraper datasource
-datasource: "username:password@(localhost:3306)/"
+# mysql connection params which build the DATA_SOURCE_NAME env var of the docker container
+mysql:
+  db: ""
+  host: "localhost"
+  param: ""
+  pass: "password"
+  port: 3306
+  protocol: ""
+  user: "exporter"
 
 # cloudsqlproxy https://cloud.google.com/sql/docs/mysql/sql-proxy
 cloudsqlproxy:

```

## 0.1.0

**Release date:** 2018-05-17

![AppVersion: v0.10.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.10.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Added prometheus mysql exporter chart (#5373)

### Default value changes

```diff
# Default values for prometheus-mysql-exporter.
# This is a YAML-formatted file.
# Declare variables to be passed into your templates.

replicaCount: 1

image:
  repository: "prom/mysqld-exporter"
  tag: "v0.10.0"
  pullPolicy: "IfNotPresent"

service:
  name: mysql-exporter
  type: ClusterIP
  externalPort: 9104
  internalPort: 9104

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

annotations: {}
#  prometheus.io/scrape: "true"
#  prometheus.io/path: "/metrics"
#  prometheus.io/port: "{{ .Values.service.internalPort }}"

# Scraper datasource
datasource: "username:password@(localhost:3306)/"

# cloudsqlproxy https://cloud.google.com/sql/docs/mysql/sql-proxy
cloudsqlproxy:
  enabled: false
  image:
    repo: "gcr.io/cloudsql-docker/gce-proxy"
    tag: "1.11"
    pullPolicy: "IfNotPresent"
  instanceConnectionName: "project:us-central1:dbname"
  port: "3306"
  credentials:
    '{
      "type": "service_account",
      "project_id": "project",
      "private_key_id": "KEYID1",
      "private_key": "-----BEGIN PRIVATE KEY-----\sdajsdnasd\n-----END PRIVATE KEY-----\n",
      "client_email": "user@project.iam.gserviceaccount.com",
      "client_id": "111111111",
      "auth_uri": "https://accounts.google.com/o/oauth2/auth",
      "token_uri": "https://accounts.google.com/o/oauth2/token",
      "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
      "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/user%40project.iam.gserviceaccount.com"
    }'

```

---
Autogenerated from Helm Chart and git history using [helm-changelog](https://github.com/mogensen/helm-changelog)
