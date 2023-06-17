# Change Log

## 0.4.3

**Release date:** 2023-04-27

![AppVersion: v0.9.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.9.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-smartctl-exporter] Add README file (#3290)

### Default value changes

```diff
# No changes in this release
```

## 0.4.2

**Release date:** 2023-04-26

![AppVersion: v0.9.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.9.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-smartctl-exporter] Add a maintainer (#3277)

### Default value changes

```diff
# No changes in this release
```

## 0.4.1

**Release date:** 2023-04-21

![AppVersion: v0.9.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.9.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-smartctl-exporter]: allow annotation to daemonset  (#3265)

### Default value changes

```diff
diff --git a/charts/prometheus-smartctl-exporter/values.yaml b/charts/prometheus-smartctl-exporter/values.yaml
index c2d0bccd..571a403a 100644
--- a/charts/prometheus-smartctl-exporter/values.yaml
+++ b/charts/prometheus-smartctl-exporter/values.yaml
@@ -65,6 +65,8 @@ resources: {}
   #  cpu: 100m
   #  memory: 128Mi
 
+annotations: {}
+
 nodeSelector: {}
 
 tolerations:

```

## 0.4.0

**Release date:** 2023-04-15

![AppVersion: v0.9.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.9.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-smartctl-exporter]: expose serviceMonitor `interval` and `scrapeTimeout` (#3197)

### Default value changes

```diff
diff --git a/charts/prometheus-smartctl-exporter/values.yaml b/charts/prometheus-smartctl-exporter/values.yaml
index 2c6f6222..c2d0bccd 100644
--- a/charts/prometheus-smartctl-exporter/values.yaml
+++ b/charts/prometheus-smartctl-exporter/values.yaml
@@ -23,6 +23,8 @@ serviceMonitor:
   # Add Extra labels if needed. Prometeus operator may need them to find it.
   extraLabels: {}
   # release: prometheus-operator
+  interval: 60s
+  scrapeTimeout: 30s
 
 prometheusRules:
   enabled: false

```

## 0.3.2

**Release date:** 2023-04-06

![AppVersion: v0.9.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.9.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Upgrade to smartctl exporter v0.9.1 (#3186)

### Default value changes

```diff
# No changes in this release
```

## 0.3.1

**Release date:** 2022-10-12

![AppVersion: v0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.8.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Fix quoting issue in smartctl exporter (#2547)

### Default value changes

```diff
# No changes in this release
```

## 0.3.0

**Release date:** 2022-10-12

![AppVersion: v0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.8.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-smartctl-exporter] Quick mapping of chart to new 0.8.0 config and updated to that version. (#2530)

### Default value changes

```diff
diff --git a/charts/prometheus-smartctl-exporter/values.yaml b/charts/prometheus-smartctl-exporter/values.yaml
index b9aa14d4..2c6f6222 100644
--- a/charts/prometheus-smartctl-exporter/values.yaml
+++ b/charts/prometheus-smartctl-exporter/values.yaml
@@ -33,10 +33,7 @@ prometheusRules:
   # release: prometheus-operator
 
 image:
-  # future location
-  # repository: quay.io/prometheus/smartctl-exporter
-  # temporary location until an offical build exists.
-  repository: prometheuscommunity/smartctl-exporter
+  repository: quay.io/prometheuscommunity/smartctl-exporter
   pullPolicy: IfNotPresent
   # Overrides the image tag whose default is the chart appVersion.
   tag: ""

```

## 0.2.0

**Release date:** 2022-10-11

![AppVersion: v0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.8.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Upgrade to smartctl exporter v0.8.0 (#2542)

### Default value changes

```diff
# No changes in this release
```

## 0.1.1

**Release date:** 2022-08-16

![AppVersion: v0.7.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.7.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Fix smartctl image and servicemonitor. (#2365)

### Default value changes

```diff
diff --git a/charts/prometheus-smartctl-exporter/values.yaml b/charts/prometheus-smartctl-exporter/values.yaml
index adadf3dc..b9aa14d4 100644
--- a/charts/prometheus-smartctl-exporter/values.yaml
+++ b/charts/prometheus-smartctl-exporter/values.yaml
@@ -33,7 +33,10 @@ prometheusRules:
   # release: prometheus-operator
 
 image:
-  repository: quay.io/prometheus/smartctl-exporter
+  # future location
+  # repository: quay.io/prometheus/smartctl-exporter
+  # temporary location until an offical build exists.
+  repository: prometheuscommunity/smartctl-exporter
   pullPolicy: IfNotPresent
   # Overrides the image tag whose default is the chart appVersion.
   tag: ""

```

## 0.1.0

**Release date:** 2022-08-12

![AppVersion: v0.7.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.7.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-smartctl-exporter] Initial stab at a prometheus-smartctl-exporter chart (#2357)

### Default value changes

```diff
config: {}
#  devices:
#  - /dev/sda

extraInstances: []
# - config:
#     devices:
#     - /dev/nvme0n1
#   nodeSelector:
#     type: other

common:
  config:
    bind_to: "0.0.0.0:9633"
    url_path: "/metrics"
    smartctl_location: /usr/sbin/smartctl
    collect_not_more_than_period: 120s

serviceMonitor:
  enabled: false
  # Specify namespace to load the monitor if not in the same namespace
  # namespace: prometheus-operator
  # Add Extra labels if needed. Prometeus operator may need them to find it.
  extraLabels: {}
  # release: prometheus-operator

prometheusRules:
  enabled: false
  # Specify namespace to load the monitor if not in the same namespace
  # namespace: prometheus-operator
  # Add Extra labels if needed. Prometeus operator may need them to find it.
  extraLabels: {}
  # release: prometheus-operator

image:
  repository: quay.io/prometheus/smartctl-exporter
  pullPolicy: IfNotPresent
  # Overrides the image tag whose default is the chart appVersion.
  tag: ""

serviceAccount:
  # Specifies whether a service account should be created
  create: true
  # Annotations to add to the service account
  annotations: {}
  # The name of the service account to use.
  # If not set and create is true, a name is generated using the fullname template
  name: ""

rbac:
  create: true
  podSecurityPolicy: unrestricted-psp

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

tolerations:
  - effect: NoSchedule
    operator: Exists

affinity: {}

service:
  type: ClusterIP
  port: 80

```

---
Autogenerated from Helm Chart and git history using [helm-changelog](https://github.com/mogensen/helm-changelog)
