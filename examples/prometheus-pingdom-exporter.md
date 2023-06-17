# Change Log

## 2.4.1

**Release date:** 2021-09-13

![AppVersion: 20190610-1](https://img.shields.io/static/v1?label=AppVersion&message=20190610-1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-pingdom-exporter] Minor tweaks to README (#1335)

### Default value changes

```diff
# No changes in this release
```

## 2.4.0

**Release date:** 2021-04-12

![AppVersion: 20190610-1](https://img.shields.io/static/v1?label=AppVersion&message=20190610-1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-pingdom-exporter] Add secret override for pingdom exporter (#741)

### Default value changes

```diff
diff --git a/charts/prometheus-pingdom-exporter/values.yaml b/charts/prometheus-pingdom-exporter/values.yaml
index 9be0fc8c..5e4bfb21 100644
--- a/charts/prometheus-pingdom-exporter/values.yaml
+++ b/charts/prometheus-pingdom-exporter/values.yaml
@@ -56,6 +56,9 @@ pod:
     # key: "true"
     # example: "false"
 
+existingSecret:
+  name: ""
+
 secret:
   annotations: {}
     # key: "true"

```

## 2.3.2

**Release date:** 2021-01-20

![AppVersion: 20190610-1](https://img.shields.io/static/v1?label=AppVersion&message=20190610-1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* update version (#561)

### Default value changes

```diff
# No changes in this release
```

## 2.3.1

**Release date:** 2020-11-06

![AppVersion: 20190610-1](https://img.shields.io/static/v1?label=AppVersion&message=20190610-1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* charts/prometheus-pingdom-exporter: fix typo (#317)
* fix yaml lint
* added artifact hub annotations

### Default value changes

```diff
# No changes in this release
```

## 2.3.0

**Release date:** 2020-10-07

![AppVersion: 20190610-1](https://img.shields.io/static/v1?label=AppVersion&message=20190610-1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* update pingdom exporter

### Default value changes

```diff
diff --git a/charts/prometheus-pingdom-exporter/values.yaml b/charts/prometheus-pingdom-exporter/values.yaml
index 7a9bdcc4..9be0fc8c 100644
--- a/charts/prometheus-pingdom-exporter/values.yaml
+++ b/charts/prometheus-pingdom-exporter/values.yaml
@@ -7,7 +7,7 @@ replicaCount: 1
 image:
   # we use camptocamp/prometheus-pingdom-exporter image as giantswarm did not publish recent versions after 0.1.1
   repository: camptocamp/prometheus-pingdom-exporter
-  tag: 20180821-1
+  tag: 20190610-1
   pullPolicy: IfNotPresent
 
 nameOverride: ""

```

## 2.2.0

**Release date:** 2019-09-23

![AppVersion: 20180821-1](https://img.shields.io/static/v1?label=AppVersion&message=20180821-1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* add annotations to secret (#187)

### Default value changes

```diff
diff --git a/charts/prometheus-pingdom-exporter/values.yaml b/charts/prometheus-pingdom-exporter/values.yaml
index 26fd9a0e..7a9bdcc4 100644
--- a/charts/prometheus-pingdom-exporter/values.yaml
+++ b/charts/prometheus-pingdom-exporter/values.yaml
@@ -54,4 +54,9 @@ pingdom:
 pod:
   annotations: {}
     # key: "true"
-  # example: "false"
+    # example: "false"
+
+secret:
+  annotations: {}
+    # key: "true"
+    # example: "false"

```

## 2.1.0

**Release date:** 2019-09-21

![AppVersion: 20180821-1](https://img.shields.io/static/v1?label=AppVersion&message=20180821-1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* add pod annotations (#186)

### Default value changes

```diff
diff --git a/charts/prometheus-pingdom-exporter/values.yaml b/charts/prometheus-pingdom-exporter/values.yaml
index 06b0c6e7..26fd9a0e 100644
--- a/charts/prometheus-pingdom-exporter/values.yaml
+++ b/charts/prometheus-pingdom-exporter/values.yaml
@@ -50,3 +50,8 @@ pingdom:
   accountEmail: somebodyorelse@invalid
   # time (in seconds) between accessing the Pingdom  API
   wait: 10
+
+pod:
+  annotations: {}
+    # key: "true"
+  # example: "false"

```

## 2.0.0

**Release date:** 2019-01-14

![AppVersion: 20180821-1](https://img.shields.io/static/v1?label=AppVersion&message=20180821-1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Prometheus pingdom exporter (#18)

### Default value changes

```diff
# Default values for prometheus-pingdom-exporter.
# This is a YAML-formatted file.
# Declare variables to be passed into your templates.

replicaCount: 1

image:
  # we use camptocamp/prometheus-pingdom-exporter image as giantswarm did not publish recent versions after 0.1.1
  repository: camptocamp/prometheus-pingdom-exporter
  tag: 20180821-1
  pullPolicy: IfNotPresent

nameOverride: ""
fullnameOverride: ""

service:
  type: ClusterIP
  port: 9100
  annotations: {}
    # prometheus.io/scrape: "true"
    # prometheus.io/port: "9100"

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

# configuration of the pingdom credentials
pingdom:
  # username of the pingdom account
  user: somebody@invalid
  # password of the pingdom account
  password: totallysecret
  # application id / api secret can be created on the pingdom website
  appId: alsototallysecret
  # account email of the account owner if using multiaccount / team accounts
  accountEmail: somebodyorelse@invalid
  # time (in seconds) between accessing the Pingdom  API
  wait: 10

```

---
Autogenerated from Helm Chart and git history using [helm-changelog](https://github.com/mogensen/helm-changelog)
