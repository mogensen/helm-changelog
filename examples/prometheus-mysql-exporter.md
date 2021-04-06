# Change Log

## 1.1.0 

**Release date:** 2021-02-17

![AppVersion: v0.12.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.12.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-mysql-exporter] Add securityContext and podSecurityContext (#662) 

### Default value changes

```diff
diff --git a/charts/prometheus-mysql-exporter/values.yaml b/charts/prometheus-mysql-exporter/values.yaml
index f77038b..c1a3cd7 100644
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

![AppVersion: v0.12.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.12.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-mysql-exporter] fixed link to Prometheus MySQL Exporter repo (#630) 

### Default value changes

```diff
# No changes in this release
```

## 1.0.1 

**Release date:** 2020-12-08

![AppVersion: v0.12.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.12.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* fix: use alpine cloudsqlproxy image (#464) 

### Default value changes

```diff
diff --git a/charts/prometheus-mysql-exporter/values.yaml b/charts/prometheus-mysql-exporter/values.yaml
index 5c0b107..f77038b 100644
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

![AppVersion: v0.12.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.12.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-mysql-exporter] update docker image, readme links & use default labels (#68) 

### Default value changes

```diff
diff --git a/charts/prometheus-mysql-exporter/values.yaml b/charts/prometheus-mysql-exporter/values.yaml
index b67a829..5c0b107 100644
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

![AppVersion: v0.11.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.11.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Prep initial charts indexing (#14) 

### Default value changes

```diff
# No changes in this release
```

## 0.7.0 

**Release date:** 2020-07-28

![AppVersion: v0.11.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.11.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-mysql-exporter] Add ServiceMonitor metricRelabeling (#23352) 

### Default value changes

```diff
diff --git a/charts/prometheus-mysql-exporter/values.yaml b/charts/prometheus-mysql-exporter/values.yaml
index d308647..b67a829 100644
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

![AppVersion: v0.11.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.11.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-mysql-exporter]:  Add service labels and annotations (#22986) 

### Default value changes

```diff
diff --git a/charts/prometheus-mysql-exporter/values.yaml b/charts/prometheus-mysql-exporter/values.yaml
index 9b8afd9..d308647 100644
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

**Release date:** 2020-06-04

![AppVersion: v0.11.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.11.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-mysql-exporter]: fix docs typo (#22674) 

### Default value changes

```diff
# No changes in this release
```

## 0.5.2 

**Release date:** 2019-10-25

![AppVersion: v0.11.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.11.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [mysqld-exporter] Improve Secrets (#18314) 

### Default value changes

```diff
diff --git a/charts/prometheus-mysql-exporter/values.yaml b/charts/prometheus-mysql-exporter/values.yaml
index 1f2309f..9b8afd9 100644
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

![AppVersion: v0.11.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.11.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Mysql exporter managing labels (#15349) 

### Default value changes

```diff
diff --git a/charts/prometheus-mysql-exporter/values.yaml b/charts/prometheus-mysql-exporter/values.yaml
index d1e48a4..1f2309f 100644
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

![AppVersion: v0.11.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.11.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* feat(stable/prometheus-mysql-exporter): add optional ServiceMonitor (#14600) 

### Default value changes

```diff
diff --git a/charts/prometheus-mysql-exporter/values.yaml b/charts/prometheus-mysql-exporter/values.yaml
index c044141..d1e48a4 100644
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

![AppVersion: v0.11.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.11.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* feat(stable/prometheus-mysql-exporter): store MySQL password in a k8s Secret (#14712) 

### Default value changes

```diff
# No changes in this release
```

## 0.3.4 

**Release date:** 2019-06-16

![AppVersion: v0.11.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.11.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-mysql-exporter] updated gce-proxy to 1.14 & added myself to maintainers & owners (#14776) 

### Default value changes

```diff
diff --git a/charts/prometheus-mysql-exporter/values.yaml b/charts/prometheus-mysql-exporter/values.yaml
index 7f4828f..c044141 100644
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

![AppVersion: v0.11.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.11.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-mysql-exporter] Fix invalid YAML with non-boolean collectors (#14701) 

### Default value changes

```diff
# No changes in this release
```

## 0.3.2 

**Release date:** 2019-03-26

![AppVersion: v0.11.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.11.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* fixed readme (#11660) 

### Default value changes

```diff
# No changes in this release
```

## 0.3.1 

**Release date:** 2019-03-25

![AppVersion: v0.11.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.11.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-mysql-importer] Add optional labels to each pod  (#11380) 

### Default value changes

```diff
diff --git a/charts/prometheus-mysql-exporter/values.yaml b/charts/prometheus-mysql-exporter/values.yaml
index 7fd4de2..7f4828f 100644
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

**Release date:** 2019-03-06

![AppVersion: v0.11.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.11.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-mysql-exporter] Collector flags config (#11569) 

### Default value changes

```diff
diff --git a/charts/prometheus-mysql-exporter/values.yaml b/charts/prometheus-mysql-exporter/values.yaml
index 77fac30..7fd4de2 100644
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

![AppVersion: v0.11.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.11.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* fixed annotations (#8431) 

### Default value changes

```diff
diff --git a/charts/prometheus-mysql-exporter/values.yaml b/charts/prometheus-mysql-exporter/values.yaml
index 58387d5..77fac30 100644
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

![AppVersion: v0.11.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.11.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-mysql-exporter] split datasource in separate mysql connection vars (#8000) 

### Default value changes

```diff
diff --git a/charts/prometheus-mysql-exporter/values.yaml b/charts/prometheus-mysql-exporter/values.yaml
index fa1d9c9..58387d5 100644
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

![AppVersion: v0.10.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.10.0&color=success&logo=)
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
