# Change Log

## 1.0.0

**Release date:** 2023-02-07

![AppVersion: 0.4.0](https://img.shields.io/static/v1?label=AppVersion&message=0.4.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-consul-exporter] Increment Helm apiVersion (#2991)

### Default value changes

```diff
# No changes in this release
```

## 0.5.1

**Release date:** 2022-10-16

![AppVersion: 0.4.0](https://img.shields.io/static/v1?label=AppVersion&message=0.4.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-consul-exporter] Fix PSP deprecation after k8s 1.25+ (#2567)

### Default value changes

```diff
# No changes in this release
```

## 0.5.0

**Release date:** 2021-11-09

![AppVersion: 0.4.0](https://img.shields.io/static/v1?label=AppVersion&message=0.4.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-consul-exporter] Add Pod Annotations Value (#1416)

### Default value changes

```diff
diff --git a/charts/prometheus-consul-exporter/values.yaml b/charts/prometheus-consul-exporter/values.yaml
index 22cad251..d4ad38ae 100644
--- a/charts/prometheus-consul-exporter/values.yaml
+++ b/charts/prometheus-consul-exporter/values.yaml
@@ -88,6 +88,9 @@ affinity: {}
 # Extra environment variables
 extraEnv: []
 
+# Annotations for the pods
+podAnnotations: {}
+
 # Init Containers for Exporter Pod
 initContainers: []
 

```

## 0.4.0

**Release date:** 2020-12-02

![AppVersion: 0.4.0](https://img.shields.io/static/v1?label=AppVersion&message=0.4.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* add service monitor for consul exporter (#439)

### Default value changes

```diff
diff --git a/charts/prometheus-consul-exporter/values.yaml b/charts/prometheus-consul-exporter/values.yaml
index 868a7507..22cad251 100644
--- a/charts/prometheus-consul-exporter/values.yaml
+++ b/charts/prometheus-consul-exporter/values.yaml
@@ -62,6 +62,23 @@ resources: {}
   #  cpu: 100m
   #  memory: 128Mi
 
+serviceMonitor:
+  # When set true then use a ServiceMonitor to configure scraping
+  enabled: false
+  # Set the namespace the ServiceMonitor should be deployed
+  # namespace: monitoring
+  # Set how frequently Prometheus should scrape
+  # interval: 30s
+  # Set path to consul-exporter telemtery-path
+  # telemetryPath: /metrics
+  # Set labels for the ServiceMonitor, use this to define your scrape label for Prometheus Operator
+  # labels:
+  # Set timeout for scrape
+  # timeout: 10s
+  # Set of labels to transfer on the Kubernetes Service onto the target.
+  # targetLabels: []
+  # metricRelabelings: []
+
 nodeSelector: {}
 
 tolerations: []

```

## 0.3.0

**Release date:** 2020-11-18

![AppVersion: 0.4.0](https://img.shields.io/static/v1?label=AppVersion&message=0.4.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* feat(prometheus-consul-exporter): add configurable apiVersion (#383)

### Default value changes

```diff
# No changes in this release
```

## 0.2.0

**Release date:** 2020-10-04

![AppVersion: 0.4.0](https://img.shields.io/static/v1?label=AppVersion&message=0.4.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-consul-exporter] Support extra containers, environment variables and volumes (#162)

### Default value changes

```diff
diff --git a/charts/prometheus-consul-exporter/values.yaml b/charts/prometheus-consul-exporter/values.yaml
index 0a8939f5..868a7507 100644
--- a/charts/prometheus-consul-exporter/values.yaml
+++ b/charts/prometheus-consul-exporter/values.yaml
@@ -67,3 +67,23 @@ nodeSelector: {}
 tolerations: []
 
 affinity: {}
+
+# Extra environment variables
+extraEnv: []
+
+# Init Containers for Exporter Pod
+initContainers: []
+
+# Extra containers for the exporter pod
+extraContainers: []
+
+# Extra Volumes for the pod
+extraVolumes: []
+# - name: example
+#   configMap:
+#     name: example
+
+# Extra Volume Mounts for the exporter container
+extraVolumeMounts: []
+# - name: example
+#   mountPath: /example

```

## 0.1.7

**Release date:** 2020-09-07

![AppVersion: 0.4.0](https://img.shields.io/static/v1?label=AppVersion&message=0.4.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-consul-exporter] add gkarthiks as additional maintainers (#50)

### Default value changes

```diff
# No changes in this release
```

## 0.1.6

**Release date:** 2020-08-20

![AppVersion: 0.4.0](https://img.shields.io/static/v1?label=AppVersion&message=0.4.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Prep initial charts indexing (#14)

### Default value changes

```diff
# No changes in this release
```

## 0.1.5

**Release date:** 2020-07-18

![AppVersion: 0.4.0](https://img.shields.io/static/v1?label=AppVersion&message=0.4.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-consul-exporter] Fix PSP apiVersion (#22834)

### Default value changes

```diff
# No changes in this release
```

## 0.1.4

**Release date:** 2019-06-28

![AppVersion: 0.4.0](https://img.shields.io/static/v1?label=AppVersion&message=0.4.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-consul-exporter] 14589: Bugfix to support boolean flags in prometheus-consul-exporter (#14694)

### Default value changes

```diff
diff --git a/charts/prometheus-consul-exporter/values.yaml b/charts/prometheus-consul-exporter/values.yaml
index e4e87ac8..0a8939f5 100644
--- a/charts/prometheus-consul-exporter/values.yaml
+++ b/charts/prometheus-consul-exporter/values.yaml
@@ -29,6 +29,8 @@ consulServer: host:port
 
 # Flags - for a list visit https://github.com/prometheus/consul_exporter#flags
 options: {}
+  # no-consul.health-summary:
+  # kv.filter=foobar
 
 service:
   type: ClusterIP

```

## 0.1.3

**Release date:** 2019-03-18

![AppVersion: 0.4.0](https://img.shields.io/static/v1?label=AppVersion&message=0.4.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Add possibility to disable apparmor podsecuritypolicy annotations. (#11934)

### Default value changes

```diff
diff --git a/charts/prometheus-consul-exporter/values.yaml b/charts/prometheus-consul-exporter/values.yaml
index 72049c61..e4e87ac8 100644
--- a/charts/prometheus-consul-exporter/values.yaml
+++ b/charts/prometheus-consul-exporter/values.yaml
@@ -8,6 +8,7 @@ rbac:
   # Specifies whether RBAC resources should be created
   create: true
   pspEnabled: true
+  pspUseAppArmor: true
 serviceAccount:
   # Specifies whether a ServiceAccount should be created
   create: true

```

## 0.1.2

**Release date:** 2018-12-17

![AppVersion: 0.4.0](https://img.shields.io/static/v1?label=AppVersion&message=0.4.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-consul-exporter] Fix service selector (#9958)

### Default value changes

```diff
# No changes in this release
```

## 0.1.1

**Release date:** 2018-12-12

![AppVersion: 0.4.0](https://img.shields.io/static/v1?label=AppVersion&message=0.4.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Add service annotations for prometheus-consul-exporter (#9882)

### Default value changes

```diff
# No changes in this release
```

## 0.1.0

**Release date:** 2018-12-06

![AppVersion: 0.4.0](https://img.shields.io/static/v1?label=AppVersion&message=0.4.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Add chart for consul-exporter (#9686)

### Default value changes

```diff
# Default values for consul-exporter.
# This is a YAML-formatted file.
# Declare variables to be passed into your templates.

replicaCount: 1

rbac:
  # Specifies whether RBAC resources should be created
  create: true
  pspEnabled: true
serviceAccount:
  # Specifies whether a ServiceAccount should be created
  create: true
  # The name of the ServiceAccount to use.
  # If not set and create is true, a name is generated using the fullname template
  name:

image:
  repository: prom/consul-exporter
  tag: v0.4.0
  pullPolicy: IfNotPresent

nameOverride: ""
fullnameOverride: ""

# Add your consul server details here
consulServer: host:port

# Flags - for a list visit https://github.com/prometheus/consul_exporter#flags
options: {}

service:
  type: ClusterIP
  port: 9107
  annotations: {}

ingress:
  enabled: false
  annotations: {}
    # kubernetes.io/ingress.class: nginx
    # kubernetes.io/tls-acme: "true"
  path: /
  hosts:
    - chart-example.local
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

```

---
Autogenerated from Helm Chart and git history using [helm-changelog](https://github.com/mogensen/helm-changelog)
