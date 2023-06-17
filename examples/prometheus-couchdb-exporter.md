# Change Log

## 1.0.0

**Release date:** 2023-05-11

![AppVersion: 1.0](https://img.shields.io/static/v1?label=AppVersion&message=1.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-couchdb-exporter] fix Chart.yaml - remove engine: gotpl (#3319)

### Default value changes

```diff
# No changes in this release
```

## 0.2.1

**Release date:** 2022-10-16

![AppVersion: 1.0](https://img.shields.io/static/v1?label=AppVersion&message=1.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-couchdb-exporter] Fix PSP deprecation after k8s 1.25+ (#2568)

### Default value changes

```diff
# No changes in this release
```

## 0.2.0

**Release date:** 2020-11-18

![AppVersion: 1.0](https://img.shields.io/static/v1?label=AppVersion&message=1.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* feat(prometheus-couchdb-exporter): add configurable apiVersion (#384)

### Default value changes

```diff
# No changes in this release
```

## 0.1.3

**Release date:** 2020-11-06

![AppVersion: 1.0](https://img.shields.io/static/v1?label=AppVersion&message=1.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* charts/prometheus-couchdb-exporter: fix typos (#316)

### Default value changes

```diff
diff --git a/charts/prometheus-couchdb-exporter/values.yaml b/charts/prometheus-couchdb-exporter/values.yaml
index 773a7ddf..9a575e97 100644
--- a/charts/prometheus-couchdb-exporter/values.yaml
+++ b/charts/prometheus-couchdb-exporter/values.yaml
@@ -56,11 +56,11 @@ tolerations: []
 
 affinity: {}
 
-## CouchDB exporter configuratons
+## CouchDB exporter configurations
 couchdb:
-  ## URI ofthe couchdb instance
+  ## URI of the couchdb instance
   uri: http://couchdb.default.svc:5984
-  ## Specify the list of databases to get the disk usage stats as comma seperates like "db-1,db-2"
+  ## Specify the list of databases to get the disk usage stats as comma separates like "db-1,db-2"
   ## or to get stats for every database, please use "_all_dbs"
   databases: _all_dbs
   ## CouchDB username

```

## 0.1.2

**Release date:** 2020-08-20

![AppVersion: 1.0](https://img.shields.io/static/v1?label=AppVersion&message=1.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Prep initial charts indexing (#14)

### Default value changes

```diff
# No changes in this release
```

## 0.1.1

**Release date:** 2019-06-18

![AppVersion: 1.0](https://img.shields.io/static/v1?label=AppVersion&message=1.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Run couchdb-exporter with username and password (#14704)

### Default value changes

```diff
# No changes in this release
```

## 0.1.0

**Release date:** 2018-10-04

![AppVersion: 1.0](https://img.shields.io/static/v1?label=AppVersion&message=1.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* adding chart for prometheus-couchdb-exporters (#7494)

### Default value changes

```diff
# Default values for prometheus-couchdb-exporter.
# This is a YAML-formatted file.
# Declare variables to be passed into your templates.

rbac:
  # Specifies whether RBAC resources should be created
  create: true
  # Specifies whether a PodSecurityPolicy should be created
  pspEnabled: true
serviceAccount:
  # Specifies whether a ServiceAccount should be created
  create: true
  # The name of the ServiceAccount to use.
  # If not set and create is true, a name is generated using the fullname template
  name:

replicaCount: 1

image:
  repository: gesellix/couchdb-prometheus-exporter
  tag: 16
  pullPolicy: IfNotPresent

service:
  type: ClusterIP
  port: 9984

ingress:
  enabled: false
  annotations: {}
    # kubernetes.io/ingress.class: nginx
    # kubernetes.io/tls-acme: "true"
  path: /
  hosts:
    # - chart-example.local
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
  #  cpu: 100m
  #  memory: 128Mi
  # requests:
  #  cpu: 100m
  #  memory: 128Mi

nodeSelector: {}

tolerations: []

affinity: {}

## CouchDB exporter configuratons
couchdb:
  ## URI ofthe couchdb instance
  uri: http://couchdb.default.svc:5984
  ## Specify the list of databases to get the disk usage stats as comma seperates like "db-1,db-2"
  ## or to get stats for every database, please use "_all_dbs"
  databases: _all_dbs
  ## CouchDB username
  # username:
  ## CouchDB Password
  # password:

```

---
Autogenerated from Helm Chart and git history using [helm-changelog](https://github.com/mogensen/helm-changelog)
