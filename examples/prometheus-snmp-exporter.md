# Change Log

## 1.5.0

**Release date:** 2023-05-26

![AppVersion: v0.21.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.21.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-snmp-exporter] Add extra init containers (#3402)

### Default value changes

```diff
diff --git a/charts/prometheus-snmp-exporter/values.yaml b/charts/prometheus-snmp-exporter/values.yaml
index 5a0bea0a..7ec8fba9 100644
--- a/charts/prometheus-snmp-exporter/values.yaml
+++ b/charts/prometheus-snmp-exporter/values.yaml
@@ -39,6 +39,13 @@ extraConfigmapMounts: []
   #   readOnly: true
   #   defaultMode: 420
 
+## Additional init containers
+# These will be added to the prometheus-snmp-exporter pod.
+extraInitContainers: []
+  # - name: init-myservice
+  #   image: busybox:1.28
+  #   command: [ 'sh', '-c', "sleep 10; done" ]
+
 ## Additional secret mounts
 # Defines additional mounts with secrets. Secrets must be manually created in the namespace.
 extraSecretMounts: []

```

## 1.4.1

**Release date:** 2023-05-25

![AppVersion: v0.21.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.21.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-snmp-exporter] Add xiu as a maintainer (#3405)

### Default value changes

```diff
# No changes in this release
```

## 1.4.0

**Release date:** 2023-02-26

![AppVersion: v0.21.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.21.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Correct ingress template (#3072)

### Default value changes

```diff
diff --git a/charts/prometheus-snmp-exporter/values.yaml b/charts/prometheus-snmp-exporter/values.yaml
index 28525db0..5a0bea0a 100644
--- a/charts/prometheus-snmp-exporter/values.yaml
+++ b/charts/prometheus-snmp-exporter/values.yaml
@@ -87,6 +87,10 @@ service:
 ## ref: https://kubernetes.io/docs/concepts/services-networking/ingress/
 ingress:
   enabled: false
+  ## Class name can be set since version 1.18.
+  className: ""
+  ## Path type is required since version 1.18. Default: ImplementationSpecific.
+  pathType: ""
   hosts: []
     # - chart-example.local
   annotations: {}

```

## 1.3.0

**Release date:** 2023-02-14

![AppVersion: v0.21.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.21.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-snmp-exporter] Fix service monitor relabelings (#2989)

### Default value changes

```diff
diff --git a/charts/prometheus-snmp-exporter/values.yaml b/charts/prometheus-snmp-exporter/values.yaml
index f7963d0e..28525db0 100644
--- a/charts/prometheus-snmp-exporter/values.yaml
+++ b/charts/prometheus-snmp-exporter/values.yaml
@@ -4,7 +4,8 @@ kind: Deployment
 
 image:
   repository: prom/snmp-exporter
-  tag: v0.19.0
+  # if not set appVersion field from Chart.yaml is used
+  tag: ""
   pullPolicy: IfNotPresent
 
 imagePullSecrets: []
@@ -123,37 +124,58 @@ configmapReload:
   ##
   resources: {}
 
-# Enable this if you're using https://github.com/coreos/prometheus-operator
+# Enable this if you're using https://github.com/prometheus-operator/prometheus-operator
+# A service monitor will be created per each item in serviceMonitor.params[]
 serviceMonitor:
   enabled: false
   namespace: monitoring
 
   path: /snmp
 
-  # fallback to the prometheus default unless specified
+  # Fall back to the prometheus default unless specified
   # interval: 10s
   scrapeTimeout: 10s
   module:
     - if_mib
-  # relabelings: []
+
+  # Relabelings dynamically rewrite the label set of a target before it gets scraped.
+  # Set if defined unless overriden by params.relabelings.
+  # https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#monitoring.coreos.com/v1.RelabelConfig
+  relabelings: []
+    # - sourceLabels: [__address__]
+    #   targetLabel: __param_target
+    # - sourceLabels: [__param_target]
+    #   targetLabel: instance
+
+  # Metric relabeling is applied to samples as the last step before ingestion.
+  # Set if defined unless overriden by params.additionalMetricsRelabels.
   additionalMetricsRelabels: {}
 
-  ## Defaults to what's used if you follow CoreOS [Prometheus Install Instructions](https://github.com/helm/charts/tree/master/stable/prometheus-operator#tldr)
-  ## [Prometheus Selector Label](https://github.com/helm/charts/tree/master/stable/prometheus-operator#prometheus-operator-1)
-  ## [Kube Prometheus Selector Label](https://github.com/helm/charts/tree/master/stable/prometheus-operator#exporters)
+  # Label for selecting service monitors as set in Prometheus CRD.
+  # https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#monitoring.coreos.com/v1.PrometheusSpec
   selector:
     prometheus: kube-prometheus
-  # Retain the job and instance labels of the metrics pushed to the snmp-exporter
-  # [Scraping SNMP-exporter](https://github.com/prometheus/snmp_exporter#configure-the-snmp_exporter-as-a-target-to-scrape)
+
+  # Retain the job and instance labels of the metrics retrieved by the snmp-exporter
+  # https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#monitoring.coreos.com/v1.Endpoint
   honorLabels: true
 
-  params:
-    # - name: localhost                 # Human readable URL that will appear in Prometheus / AlertManager
-    #   target: 127.0.0.1               # The target that snmp will scrape
-    #   module:                         # Module used for scraping. Overrides value set in `serviceMonitor.module`
+  params: []
+    # Human readable URL that will appear in Prometheus / AlertManager
+    # - name: localhost
+    # The target that snmp will scrape
+    #   target: 127.0.0.1
+    # Module used for scraping. Overrides value set in serviceMonitor.module
+    #   module:
     #     - if_mib
-    #   labels: {}                      # Map of labels for ServiceMonitor. Overrides value set in `serviceMonitor.selector`
-    #   interval: 30s                   # Scraping interval. Overrides value set in `serviceMonitor.interval`
-    #   scrapeTimeout: 30s              # Scrape timeout. Overrides value set in `serviceMonitor.scrapeTimeout`
-    #   relabelings: []                 # MetricRelabelConfigs to apply to samples before ingestion. Overrides value set in `serviceMonitor.relabelings`
-    #   additionalMetricsRelabels: {}   # Map of metric labels and values to add
+    # Map of labels for ServiceMonitor. Overrides value set in serviceMonitor.selector
+    #   labels: {}
+          # release: kube-prometheus-stack
+    # Scraping interval. Overrides value set in serviceMonitor.interval
+    #   interval: 30s
+    # Scrape timeout. Overrides value set in serviceMonitor.scrapeTimeout
+    #   scrapeTimeout: 30s
+    # Relabelings. Overrides value set in serviceMonitor.relabelings
+    #   relabelings: []
+    # Map of metric labels and values to add. Overrides value set in serviceMonitor.additionalMetricsRelabels
+    #   additionalMetricsRelabels: {}

```

## 1.2.1

**Release date:** 2022-09-03

![AppVersion: 0.19.0](https://img.shields.io/static/v1?label=AppVersion&message=0.19.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-snmp-exporter] update ingress api version template to support multiple kubernetes versions (#2395)

### Default value changes

```diff
# No changes in this release
```

## 1.2.0

**Release date:** 2022-08-24

![AppVersion: 0.19.0](https://img.shields.io/static/v1?label=AppVersion&message=0.19.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-snmp-exporter] make DaemonSet available (#2394)

### Default value changes

```diff
diff --git a/charts/prometheus-snmp-exporter/values.yaml b/charts/prometheus-snmp-exporter/values.yaml
index 4b9c482a..f7963d0e 100644
--- a/charts/prometheus-snmp-exporter/values.yaml
+++ b/charts/prometheus-snmp-exporter/values.yaml
@@ -1,5 +1,7 @@
 restartPolicy: Always
 
+kind: Deployment
+
 image:
   repository: prom/snmp-exporter
   tag: v0.19.0

```

## 1.1.0

**Release date:** 2022-04-22

![AppVersion: 0.19.0](https://img.shields.io/static/v1?label=AppVersion&message=0.19.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Add pod and container security context to snmp-exporter (#2002)

### Default value changes

```diff
diff --git a/charts/prometheus-snmp-exporter/values.yaml b/charts/prometheus-snmp-exporter/values.yaml
index 72917b0a..4b9c482a 100644
--- a/charts/prometheus-snmp-exporter/values.yaml
+++ b/charts/prometheus-snmp-exporter/values.yaml
@@ -10,6 +10,18 @@ nodeSelector: {}
 tolerations: []
 affinity: {}
 
+## Security context to be added to snmp-exporter pods
+securityContext: {}
+  # fsGroup: 1000
+  # runAsUser: 1000
+  # runAsNonRoot: true
+
+## Security context to be added to snmp-exporter containers
+containerSecurityContext:
+  runAsNonRoot: true
+  runAsUser: 1000
+  readOnlyRootFilesystem: true
+
 ## Additional labels to add to all resources
 customLabels: {}
   # app: snmp-exporter

```

## 1.0.1

**Release date:** 2022-03-31

![AppVersion: 0.19.0](https://img.shields.io/static/v1?label=AppVersion&message=0.19.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-snmp-exporter] Fix README.md (#1941)

### Default value changes

```diff
# No changes in this release
```

## 1.0.0

**Release date:** 2022-03-14

![AppVersion: 0.19.0](https://img.shields.io/static/v1?label=AppVersion&message=0.19.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-snmp-exporter] Support for multiple targets when using ServiceMonitor (#1870)

### Default value changes

```diff
diff --git a/charts/prometheus-snmp-exporter/values.yaml b/charts/prometheus-snmp-exporter/values.yaml
index f13f3d0a..72917b0a 100644
--- a/charts/prometheus-snmp-exporter/values.yaml
+++ b/charts/prometheus-snmp-exporter/values.yaml
@@ -73,7 +73,7 @@ service:
 ingress:
   enabled: false
   hosts: []
-     # - chart-example.local
+    # - chart-example.local
   annotations: {}
     # kubernetes.io/ingress.class: nginx
     # kubernetes.io/tls-acme: "true"
@@ -114,8 +114,15 @@ serviceMonitor:
   enabled: false
   namespace: monitoring
 
+  path: /snmp
+
   # fallback to the prometheus default unless specified
   # interval: 10s
+  scrapeTimeout: 10s
+  module:
+    - if_mib
+  # relabelings: []
+  additionalMetricsRelabels: {}
 
   ## Defaults to what's used if you follow CoreOS [Prometheus Install Instructions](https://github.com/helm/charts/tree/master/stable/prometheus-operator#tldr)
   ## [Prometheus Selector Label](https://github.com/helm/charts/tree/master/stable/prometheus-operator#prometheus-operator-1)
@@ -127,12 +134,12 @@ serviceMonitor:
   honorLabels: true
 
   params:
-    enabled: false
-    conf:
-      module:
-      - if_mib
-      target:
-      - 127.0.0.1
-
-  path: /snmp
-  scrapeTimeout: 10s
+    # - name: localhost                 # Human readable URL that will appear in Prometheus / AlertManager
+    #   target: 127.0.0.1               # The target that snmp will scrape
+    #   module:                         # Module used for scraping. Overrides value set in `serviceMonitor.module`
+    #     - if_mib
+    #   labels: {}                      # Map of labels for ServiceMonitor. Overrides value set in `serviceMonitor.selector`
+    #   interval: 30s                   # Scraping interval. Overrides value set in `serviceMonitor.interval`
+    #   scrapeTimeout: 30s              # Scrape timeout. Overrides value set in `serviceMonitor.scrapeTimeout`
+    #   relabelings: []                 # MetricRelabelConfigs to apply to samples before ingestion. Overrides value set in `serviceMonitor.relabelings`
+    #   additionalMetricsRelabels: {}   # Map of metric labels and values to add

```

## 0.2.0

**Release date:** 2022-01-31

![AppVersion: 0.19.0](https://img.shields.io/static/v1?label=AppVersion&message=0.19.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-snmp-exporter] Added missing imagePullSecrets property (#1756)

### Default value changes

```diff
diff --git a/charts/prometheus-snmp-exporter/values.yaml b/charts/prometheus-snmp-exporter/values.yaml
index 8f0001cb..f13f3d0a 100644
--- a/charts/prometheus-snmp-exporter/values.yaml
+++ b/charts/prometheus-snmp-exporter/values.yaml
@@ -5,6 +5,7 @@ image:
   tag: v0.19.0
   pullPolicy: IfNotPresent
 
+imagePullSecrets: []
 nodeSelector: {}
 tolerations: []
 affinity: {}

```

## 0.1.5

**Release date:** 2021-10-20

![AppVersion: 0.19.0](https://img.shields.io/static/v1?label=AppVersion&message=0.19.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Add: can customize probes (#1448)

### Default value changes

```diff
diff --git a/charts/prometheus-snmp-exporter/values.yaml b/charts/prometheus-snmp-exporter/values.yaml
index fc42cac2..8f0001cb 100644
--- a/charts/prometheus-snmp-exporter/values.yaml
+++ b/charts/prometheus-snmp-exporter/values.yaml
@@ -51,6 +51,15 @@ resources: {}
   # requests:
   #   memory: 50Mi
 
+livenessProbe:
+  httpGet:
+    path: /health
+    port: http
+readinessProbe:
+  httpGet:
+    path: /health
+    port: http
+
 service:
   annotations: {}
   type: ClusterIP

```

## 0.1.4

**Release date:** 2021-08-16

![AppVersion: 0.19.0](https://img.shields.io/static/v1?label=AppVersion&message=0.19.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* prometheus-snmp-exporter:getting error in  servicemonitor.yaml after enabling  serviceMonitor  (#1251)
* [prometheus-snmp-exporter] Kubernetes recommanded labels and custom labels (#1028)

### Default value changes

```diff
diff --git a/charts/prometheus-snmp-exporter/values.yaml b/charts/prometheus-snmp-exporter/values.yaml
index b067cd8f..fc42cac2 100644
--- a/charts/prometheus-snmp-exporter/values.yaml
+++ b/charts/prometheus-snmp-exporter/values.yaml
@@ -9,6 +9,10 @@ nodeSelector: {}
 tolerations: []
 affinity: {}
 
+## Additional labels to add to all resources
+customLabels: {}
+  # app: snmp-exporter
+
 # config:
 
 extraConfigmapMounts: []

```

## 0.1.3

**Release date:** 2021-07-15

![AppVersion: 0.19.0](https://img.shields.io/static/v1?label=AppVersion&message=0.19.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Merge duplicate security context in deployment (#1143)

### Default value changes

```diff
# No changes in this release
```

## 0.1.2

**Release date:** 2021-02-09

![AppVersion: 0.19.0](https://img.shields.io/static/v1?label=AppVersion&message=0.19.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* chore: bump configmap reloader (#646)

### Default value changes

```diff
diff --git a/charts/prometheus-snmp-exporter/values.yaml b/charts/prometheus-snmp-exporter/values.yaml
index 32a7bc7f..b067cd8f 100644
--- a/charts/prometheus-snmp-exporter/values.yaml
+++ b/charts/prometheus-snmp-exporter/values.yaml
@@ -87,7 +87,7 @@ configmapReload:
   ##
   image:
     repository: jimmidyson/configmap-reload
-    tag: v0.2.2
+    tag: v0.5.0
     pullPolicy: IfNotPresent
 
   ## configmap-reload resource requests and limits

```

## 0.1.1

**Release date:** 2020-11-26

![AppVersion: 0.19.0](https://img.shields.io/static/v1?label=AppVersion&message=0.19.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-snmp-exporter] Fixes #398 and adds some features (#399)

### Default value changes

```diff
# No changes in this release
```

## 0.1.0

**Release date:** 2020-11-21

![AppVersion: 0.19.0](https://img.shields.io/static/v1?label=AppVersion&message=0.19.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-snmp-exporter] Add support mount extra secret and configmap  (#218)

### Default value changes

```diff
diff --git a/charts/prometheus-snmp-exporter/values.yaml b/charts/prometheus-snmp-exporter/values.yaml
index 0695068a..32a7bc7f 100644
--- a/charts/prometheus-snmp-exporter/values.yaml
+++ b/charts/prometheus-snmp-exporter/values.yaml
@@ -2,7 +2,7 @@ restartPolicy: Always
 
 image:
   repository: prom/snmp-exporter
-  tag: v0.14.0
+  tag: v0.19.0
   pullPolicy: IfNotPresent
 
 nodeSelector: {}
@@ -11,6 +11,23 @@ affinity: {}
 
 # config:
 
+extraConfigmapMounts: []
+  # - name: snmp-exporter-configmap
+  #   mountPath: /run/secrets/snmp-exporter
+  #   subPath: snmp.yaml # (optional)
+  #   configMap: snmp-exporter-configmap-configmap
+  #   readOnly: true
+  #   defaultMode: 420
+
+## Additional secret mounts
+# Defines additional mounts with secrets. Secrets must be manually created in the namespace.
+extraSecretMounts: []
+  # - name: secret-files
+  #   mountPath: /run/secrets/snmp-exporter
+  #   secretName: snmp-exporter-secret-files
+  #   readOnly: true
+  #   defaultMode: 420
+
 ## For RBAC support:
 rbac:
   # Specifies whether RBAC resources should be created

```

## 0.0.6

**Release date:** 2020-08-20

![AppVersion: 0.14.0](https://img.shields.io/static/v1?label=AppVersion&message=0.14.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Prep initial charts indexing (#14)

### Default value changes

```diff
# No changes in this release
```

## 0.0.5

**Release date:** 2020-03-14

![AppVersion: 0.14.0](https://img.shields.io/static/v1?label=AppVersion&message=0.14.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-snmp-exporter] Fix serviceMonitor params (#19985)

### Default value changes

```diff
# No changes in this release
```

## 0.0.4

**Release date:** 2019-06-28

![AppVersion: 0.14.0](https://img.shields.io/static/v1?label=AppVersion&message=0.14.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-snmp-exporter] Don't run config through toYaml (#13884)

### Default value changes

```diff
# No changes in this release
```

## 0.0.3

**Release date:** 2019-06-11

![AppVersion: 0.14.0](https://img.shields.io/static/v1?label=AppVersion&message=0.14.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-snmp-exporter] add namespace metadata (#14328)

### Default value changes

```diff
# No changes in this release
```

## 0.0.2

**Release date:** 2019-03-19

![AppVersion: 0.14.0](https://img.shields.io/static/v1?label=AppVersion&message=0.14.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Add servicemonitor to prometheus-snmp-exporter chart (#12320)

### Default value changes

```diff
diff --git a/charts/prometheus-snmp-exporter/values.yaml b/charts/prometheus-snmp-exporter/values.yaml
index 38dee25c..0695068a 100644
--- a/charts/prometheus-snmp-exporter/values.yaml
+++ b/charts/prometheus-snmp-exporter/values.yaml
@@ -77,3 +77,31 @@ configmapReload:
   ## Ref: http://kubernetes.io/docs/user-guide/compute-resources/
   ##
   resources: {}
+
+# Enable this if you're using https://github.com/coreos/prometheus-operator
+serviceMonitor:
+  enabled: false
+  namespace: monitoring
+
+  # fallback to the prometheus default unless specified
+  # interval: 10s
+
+  ## Defaults to what's used if you follow CoreOS [Prometheus Install Instructions](https://github.com/helm/charts/tree/master/stable/prometheus-operator#tldr)
+  ## [Prometheus Selector Label](https://github.com/helm/charts/tree/master/stable/prometheus-operator#prometheus-operator-1)
+  ## [Kube Prometheus Selector Label](https://github.com/helm/charts/tree/master/stable/prometheus-operator#exporters)
+  selector:
+    prometheus: kube-prometheus
+  # Retain the job and instance labels of the metrics pushed to the snmp-exporter
+  # [Scraping SNMP-exporter](https://github.com/prometheus/snmp_exporter#configure-the-snmp_exporter-as-a-target-to-scrape)
+  honorLabels: true
+
+  params:
+    enabled: false
+    conf:
+      module:
+      - if_mib
+      target:
+      - 127.0.0.1
+
+  path: /snmp
+  scrapeTimeout: 10s

```

## 0.0.1

**Release date:** 2019-01-21

![AppVersion: 0.14.0](https://img.shields.io/static/v1?label=AppVersion&message=0.14.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Add prometheus-snmp-exporter chart (#10734)

### Default value changes

```diff
restartPolicy: Always

image:
  repository: prom/snmp-exporter
  tag: v0.14.0
  pullPolicy: IfNotPresent

nodeSelector: {}
tolerations: []
affinity: {}

# config:

## For RBAC support:
rbac:
  # Specifies whether RBAC resources should be created
  create: true

serviceAccount:
  # Specifies whether a ServiceAccount should be created
  create: true

  # The name of the ServiceAccount to use.
  # If not set and create is true, a name is generated using the fullname template
  name:

resources: {}
  # limits:
  #   memory: 300Mi
  # requests:
  #   memory: 50Mi

service:
  annotations: {}
  type: ClusterIP
  port: 9116

## An Ingress resource can provide name-based virtual hosting and TLS
## termination among other things for CouchDB deployments which are accessed
## from outside the Kubernetes cluster.
## ref: https://kubernetes.io/docs/concepts/services-networking/ingress/
ingress:
  enabled: false
  hosts: []
     # - chart-example.local
  annotations: {}
    # kubernetes.io/ingress.class: nginx
    # kubernetes.io/tls-acme: "true"
  tls: []
    # Secrets must be manually created in the namespace.
    # - secretName: chart-example-tls
    #   hosts:
    #     - chart-example.local

podAnnotations: {}

extraArgs: []
#  --history.limit=1000

replicas: 1
## Monitors ConfigMap changes and POSTs to a URL
## Ref: https://github.com/jimmidyson/configmap-reload
##
configmapReload:
  ## configmap-reload container name
  ##
  name: configmap-reload

  ## configmap-reload container image
  ##
  image:
    repository: jimmidyson/configmap-reload
    tag: v0.2.2
    pullPolicy: IfNotPresent

  ## configmap-reload resource requests and limits
  ## Ref: http://kubernetes.io/docs/user-guide/compute-resources/
  ##
  resources: {}

```

---
Autogenerated from Helm Chart and git history using [helm-changelog](https://github.com/mogensen/helm-changelog)
