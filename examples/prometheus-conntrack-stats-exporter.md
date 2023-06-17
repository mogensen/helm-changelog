# Change Log

## 0.5.6

**Release date:** 2023-05-26

![AppVersion: v0.4.15](https://img.shields.io/static/v1?label=AppVersion&message=v0.4.15&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-conntrack-stats-exporter] update image to v0.4.15 (#3437)

### Default value changes

```diff
# No changes in this release
```

## 0.5.5

**Release date:** 2023-01-16

![AppVersion: v0.4.11](https://img.shields.io/static/v1?label=AppVersion&message=v0.4.11&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* update image to v0.4.11 (#2920)

### Default value changes

```diff
# No changes in this release
```

## 0.5.4

**Release date:** 2022-12-24

![AppVersion: v0.4.10](https://img.shields.io/static/v1?label=AppVersion&message=v0.4.10&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* update image to v0.4.10 (#2858)

### Default value changes

```diff
# No changes in this release
```

## 0.5.3

**Release date:** 2022-11-05

![AppVersion: v0.4.7](https://img.shields.io/static/v1?label=AppVersion&message=v0.4.7&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* update image to v0.4.7 (#2649)

### Default value changes

```diff
# No changes in this release
```

## 0.5.2

**Release date:** 2022-10-12

![AppVersion: v0.4.6](https://img.shields.io/static/v1?label=AppVersion&message=v0.4.6&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* raise version to v0.4.6 (#2550)

### Default value changes

```diff
# No changes in this release
```

## 0.5.1

**Release date:** 2022-10-11

![AppVersion: v0.4.4](https://img.shields.io/static/v1?label=AppVersion&message=v0.4.4&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* fix livenessprobe (#2541)

### Default value changes

```diff
# No changes in this release
```

## 0.5.0

**Release date:** 2022-10-11

![AppVersion: v0.4.4](https://img.shields.io/static/v1?label=AppVersion&message=v0.4.4&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* add podlabels (#2538)

### Default value changes

```diff
diff --git a/charts/prometheus-conntrack-stats-exporter/values.yaml b/charts/prometheus-conntrack-stats-exporter/values.yaml
index c4f89f82..3dcd7a51 100644
--- a/charts/prometheus-conntrack-stats-exporter/values.yaml
+++ b/charts/prometheus-conntrack-stats-exporter/values.yaml
@@ -18,6 +18,8 @@ podAnnotations:
 
 commonLabels: {}
 
+podLabels: {}
+
 podMonitor:
   enabled: false
   honorLabels: true

```

## 0.4.0

**Release date:** 2022-10-11

![AppVersion: v0.4.4](https://img.shields.io/static/v1?label=AppVersion&message=v0.4.4&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* add commonlabels (#2537)

### Default value changes

```diff
diff --git a/charts/prometheus-conntrack-stats-exporter/values.yaml b/charts/prometheus-conntrack-stats-exporter/values.yaml
index b65361ce..c4f89f82 100644
--- a/charts/prometheus-conntrack-stats-exporter/values.yaml
+++ b/charts/prometheus-conntrack-stats-exporter/values.yaml
@@ -16,6 +16,8 @@ podAnnotations:
   prometheus.io/scrape: "true"
   prometheus.io/port: "9371"
 
+commonLabels: {}
+
 podMonitor:
   enabled: false
   honorLabels: true

```

## 0.3.0

**Release date:** 2022-09-16

![AppVersion: v0.4.4](https://img.shields.io/static/v1?label=AppVersion&message=v0.4.4&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* update image to v0.4.4 (#2460)

### Default value changes

```diff
# No changes in this release
```

## 0.2.1

**Release date:** 2022-08-03

![AppVersion: v0.3.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.3.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* fix conntrack stats exporter image (#2337)

### Default value changes

```diff
diff --git a/charts/prometheus-conntrack-stats-exporter/values.yaml b/charts/prometheus-conntrack-stats-exporter/values.yaml
index 5034b078..b65361ce 100644
--- a/charts/prometheus-conntrack-stats-exporter/values.yaml
+++ b/charts/prometheus-conntrack-stats-exporter/values.yaml
@@ -3,7 +3,8 @@ fullnameOverride: ""
 image:
   repository: jwkohnen/conntrack-stats-exporter
   pullPolicy: IfNotPresent
-  tag: "v0.2.2"
+  # if not set appVersion field from Chart.yaml is used
+  tag: ""
 
 imagePullSecrets: []
 

```

## 0.2.0

**Release date:** 2022-08-03

![AppVersion: 0.3.0](https://img.shields.io/static/v1?label=AppVersion&message=0.3.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* update docker image (#2336)

### Default value changes

```diff
# No changes in this release
```

## 0.1.0

**Release date:** 2022-05-05

![AppVersion: 0.2.2](https://img.shields.io/static/v1?label=AppVersion&message=0.2.2&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-conntrack-stats-exporter] added prometheus-conntrack-stats-exporter chart (#2032)

### Default value changes

```diff
fullnameOverride: ""

image:
  repository: jwkohnen/conntrack-stats-exporter
  pullPolicy: IfNotPresent
  tag: "v0.2.2"

imagePullSecrets: []

nameOverride: ""

nodeSelector: {}

podAnnotations:
  prometheus.io/scrape: "true"
  prometheus.io/port: "9371"

podMonitor:
  enabled: false
  honorLabels: true
  interval: 10s
  labels: {}
  metricRelabelings: []
  relabelings: []
  selectorOverride: {}
  scrapeTimeout: 10s

podSecurityContext: {}
  # fsGroup: 2000

resources: {}
  # requests:
  #   cpu: 10m
  #   memory: 8Mi
  # limits:
  #   cpu: 20m
  #   memory: 15Mi

securityContext:
  privileged: true

serviceAccount:
  # Specifies whether a service account should be created
  create: true
  # Annotations to add to the service account
  annotations: {}
  # The name of the service account to use.
  # If not set and create is true, a name is generated using the fullname template
  name: ""

tolerations: []

```

---
Autogenerated from Helm Chart and git history using [helm-changelog](https://github.com/mogensen/helm-changelog)
