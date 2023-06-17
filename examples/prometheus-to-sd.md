# Change Log

## 0.4.2

**Release date:** 2022-11-11

![AppVersion: 0.5.2](https://img.shields.io/static/v1?label=AppVersion&message=0.5.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Bump chart to test CI (#2676)

### Default value changes

```diff
# No changes in this release
```

## 0.4.1

**Release date:** 2022-11-03

![AppVersion: 0.5.2](https://img.shields.io/static/v1?label=AppVersion&message=0.5.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-to-sd] Remove Helm 2 info from README (#2641)

### Default value changes

```diff
# No changes in this release
```

## 0.4.0

**Release date:** 2021-01-20

![AppVersion: 0.5.2](https://img.shields.io/static/v1?label=AppVersion&message=0.5.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-to-sd]add optional monitoredResourceTypes (#591)

### Default value changes

```diff
diff --git a/charts/prometheus-to-sd/values.yaml b/charts/prometheus-to-sd/values.yaml
index 187a48b1..c90b44d4 100644
--- a/charts/prometheus-to-sd/values.yaml
+++ b/charts/prometheus-to-sd/values.yaml
@@ -7,5 +7,7 @@ resources: {}
 port: 6060
 metricsSources:
   kube-state-metrics: http://kube-state-metrics:8080
+# The monitored resource types to use, either the legacy 'gke_container', or the new 'k8s'
+monitoredResourceTypes: gke_container
 nodeSelector: {}
 tolerations: []

```

## 0.3.1

**Release date:** 2020-08-20

![AppVersion: 0.5.2](https://img.shields.io/static/v1?label=AppVersion&message=0.5.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Prep initial charts indexing (#14)

### Default value changes

```diff
# No changes in this release
```

## 0.3.0

**Release date:** 2019-10-02

![AppVersion: 0.5.2](https://img.shields.io/static/v1?label=AppVersion&message=0.5.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-to-sd] Add tolerations support for prometheus-to-sd (#17615)

### Default value changes

```diff
diff --git a/charts/prometheus-to-sd/values.yaml b/charts/prometheus-to-sd/values.yaml
index 364bfd1c..187a48b1 100644
--- a/charts/prometheus-to-sd/values.yaml
+++ b/charts/prometheus-to-sd/values.yaml
@@ -8,3 +8,4 @@ port: 6060
 metricsSources:
   kube-state-metrics: http://kube-state-metrics:8080
 nodeSelector: {}
+tolerations: []

```

## 0.2.0

**Release date:** 2019-07-15

![AppVersion: 0.5.2](https://img.shields.io/static/v1?label=AppVersion&message=0.5.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Adding acondrat to prometheus-to-sd owners (#14827)

### Default value changes

```diff
diff --git a/charts/prometheus-to-sd/values.yaml b/charts/prometheus-to-sd/values.yaml
index 5ba9bfe5..364bfd1c 100644
--- a/charts/prometheus-to-sd/values.yaml
+++ b/charts/prometheus-to-sd/values.yaml
@@ -1,9 +1,10 @@
 replicaCount: 1
 image:
   repository: gcr.io/google-containers/prometheus-to-sd
-  tag: v0.2.2
+  tag: v0.5.2
   pullPolicy: IfNotPresent
 resources: {}
 port: 6060
-metricsSources: {}
+metricsSources:
+  kube-state-metrics: http://kube-state-metrics:8080
 nodeSelector: {}

```

## 0.1.1

**Release date:** 2018-06-21

![AppVersion: 0.2.2](https://img.shields.io/static/v1?label=AppVersion&message=0.2.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-to-sd] typo fix: tables lists->table lists (#5719)

### Default value changes

```diff
# No changes in this release
```

## 0.1.0

**Release date:** 2017-11-24

![AppVersion: 0.2.2](https://img.shields.io/static/v1?label=AppVersion&message=0.2.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Adding prometheus-to-sd chart (#2363)

### Default value changes

```diff
replicaCount: 1
image:
  repository: gcr.io/google-containers/prometheus-to-sd
  tag: v0.2.2
  pullPolicy: IfNotPresent
resources: {}
port: 6060
metricsSources: {}
nodeSelector: {}

```

---
Autogenerated from Helm Chart and git history using [helm-changelog](https://github.com/mogensen/helm-changelog)
