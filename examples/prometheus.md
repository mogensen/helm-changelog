# Change Log

## 13.6.0 

**Release date:** 2021-03-10

![AppVersion: 2.24.0](https://img.shields.io/static/v1?label=AppVersion&message=2.24.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus] make path.rootfs conditional (#737) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index b14d844..9f16ddd 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -442,6 +442,10 @@ nodeExporter:
   ##
   hostPID: true
 
+  ## If true, node-exporter pods mounts host / at /host/root
+  ##
+  hostRootfs: true
+
   ## node-exporter container name
   ##
   name: node-exporter
```

## 13.5.0 

**Release date:** 2021-03-08

![AppVersion: 2.24.0](https://img.shields.io/static/v1?label=AppVersion&message=2.24.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus] Allow to enable hostNetwork (#733) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index dc4fe31..b14d844 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -946,6 +946,14 @@ server:
     #   cpu: 500m
     #   memory: 512Mi
 
+  # Required for use in managed kubernetes clusters (such as AWS EKS) with custom CNI (such as calico),
+  # because control-plane managed by AWS cannot communicate with pods' IP CIDR and admission webhooks are not working
+  ##
+  hostNetwork: false
+
+  # When hostNetwork is enabled, you probably want to set this to ClusterFirstWithHostNet
+  dnsPolicy: ClusterFirst
+
   ## Vertical Pod Autoscaler config
   ## Ref: https://github.com/kubernetes/autoscaler/tree/master/vertical-pod-autoscaler
   verticalAutoscaler:
```

## 13.4.0 

**Release date:** 2021-02-27

![AppVersion: 2.24.0](https://img.shields.io/static/v1?label=AppVersion&message=2.24.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Update Kube State Metrics chart to 2.13.* to include latest patch release and ability to provide extra arguments (#714) 

### Default value changes

```diff
# No changes in this release
```

## 13.3.3 

**Release date:** 2021-02-20

![AppVersion: 2.24.0](https://img.shields.io/static/v1?label=AppVersion&message=2.24.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus] fix imagePullSecrets indentation (#691) 

### Default value changes

```diff
# No changes in this release
```

## 13.3.2 

**Release date:** 2021-02-10

![AppVersion: 2.24.0](https://img.shields.io/static/v1?label=AppVersion&message=2.24.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* chore: fix and bump the configmap reloader for real (#664) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 75c97f3..dc4fe31 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -355,7 +355,7 @@ configmapReload:
     ##
     image:
       repository: jimmidyson/configmap-reload
-      tag: v0.4.0
+      tag: v0.5.0
       pullPolicy: IfNotPresent
 
     ## Additional configmap-reload container arguments
@@ -393,7 +393,7 @@ configmapReload:
     ##
     image:
       repository: jimmidyson/configmap-reload
-      tag: v0.4.0
+      tag: v0.5.0
       pullPolicy: IfNotPresent
 
     ## Additional configmap-reload container arguments
```

## 13.3.1 

**Release date:** 2021-02-09

![AppVersion: 2.24.0](https://img.shields.io/static/v1?label=AppVersion&message=2.24.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* update chart version (#636) 

### Default value changes

```diff
# No changes in this release
```

## 13.3.0 

**Release date:** 2021-02-09

![AppVersion: 2.24.0](https://img.shields.io/static/v1?label=AppVersion&message=2.24.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Ability for custom dnsConfig to prometheus (#643) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 891d899..75c97f3 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -297,6 +297,18 @@ alertmanager:
     #   cpu: 10m
     #   memory: 32Mi
 
+  # Custom DNS configuration to be added to alertmanager pods
+  dnsConfig: {}
+    # nameservers:
+    #   - 1.2.3.4
+    # searches:
+    #   - ns1.svc.cluster-domain.example
+    #   - my.dns.search.suffix
+    # options:
+    #   - name: ndots
+    #     value: "2"
+  #   - name: edns0
+
   ## Security context to be added to alertmanager pods
   ##
   securityContext:
@@ -528,6 +540,18 @@ nodeExporter:
     #   cpu: 100m
     #   memory: 30Mi
 
+  # Custom DNS configuration to be added to node-exporter pods
+  dnsConfig: {}
+    # nameservers:
+    #   - 1.2.3.4
+    # searches:
+    #   - ns1.svc.cluster-domain.example
+    #   - my.dns.search.suffix
+    # options:
+    #   - name: ndots
+    #     value: "2"
+  #   - name: edns0
+
   ## Security context to be added to node-exporter pods
   ##
   securityContext:
@@ -931,6 +955,17 @@ server:
     # containerPolicies:
     # - containerName: 'prometheus-server'
 
+  # Custom DNS configuration to be added to prometheus server pods
+  dnsConfig: {}
+    # nameservers:
+    #   - 1.2.3.4
+    # searches:
+    #   - ns1.svc.cluster-domain.example
+    #   - my.dns.search.suffix
+    # options:
+    #   - name: ndots
+    #     value: "2"
+  #   - name: edns0
   ## Security context to be added to server pods
   ##
   securityContext:
@@ -1103,6 +1138,18 @@ pushgateway:
     #   cpu: 10m
     #   memory: 32Mi
 
+  # Custom DNS configuration to be added to push-gateway pods
+  dnsConfig: {}
+    # nameservers:
+    #   - 1.2.3.4
+    # searches:
+    #   - ns1.svc.cluster-domain.example
+    #   - my.dns.search.suffix
+    # options:
+    #   - name: ndots
+    #     value: "2"
+  #   - name: edns0
+
   ## Security context to be added to push-gateway pods
   ##
   securityContext:
```

## 13.2.1 

**Release date:** 2021-01-14

![AppVersion: 2.24.0](https://img.shields.io/static/v1?label=AppVersion&message=2.24.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus] Update Helm dependency to use new kube-state-metrics chart versions (#497) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index f5b0ed3..891d899 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -413,7 +413,7 @@ kubeStateMetrics:
   enabled: true
 
 ## kube-state-metrics sub-chart configurable values
-## Please see https://github.com/helm/charts/tree/master/stable/kube-state-metrics
+## Please see https://github.com/kubernetes/kube-state-metrics/tree/master/charts/kube-state-metrics
 ##
 # kube-state-metrics:
 
```

## 13.2.0 

**Release date:** 2021-01-12

![AppVersion: 2.24.0](https://img.shields.io/static/v1?label=AppVersion&message=2.24.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus] image & dependency updates (#556) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 83c40d0..f5b0ed3 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -589,7 +589,7 @@ server:
   ##
   image:
     repository: quay.io/prometheus/prometheus
-    tag: v2.22.1
+    tag: v2.24.0
     pullPolicy: IfNotPresent
 
   ## prometheus server priorityClassName
@@ -994,7 +994,7 @@ pushgateway:
   ##
   image:
     repository: prom/pushgateway
-    tag: v1.3.0
+    tag: v1.3.1
     pullPolicy: IfNotPresent
 
   ## pushgateway priorityClassName
```

## 13.1.0 

**Release date:** 2021-01-11

![AppVersion: 2.22.1](https://img.shields.io/static/v1?label=AppVersion&message=2.22.1&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus] Support scheme annotation for pod scrapers. (#559) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index a3bc388..83c40d0 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -1509,6 +1509,8 @@ serverFiles:
       # following annotations:
       #
       # * `prometheus.io/scrape`: Only scrape pods that have a value of `true`
+      # * `prometheus.io/scheme`: If the metrics endpoint is secured then you will need
+      # to set this to `https` & most likely set the `tls_config` of the scrape config.
       # * `prometheus.io/path`: If the metrics path is not `/metrics` override this.
       # * `prometheus.io/port`: Scrape the pod on the indicated port instead of the default of `9102`.
       - job_name: 'kubernetes-pods'
@@ -1520,6 +1522,10 @@ serverFiles:
           - source_labels: [__meta_kubernetes_pod_annotation_prometheus_io_scrape]
             action: keep
             regex: true
+          - source_labels: [__meta_kubernetes_pod_annotation_prometheus_io_scheme]
+            action: replace
+            regex: (https?)
+            target_label: __scheme__
           - source_labels: [__meta_kubernetes_pod_annotation_prometheus_io_path]
             action: replace
             target_label: __metrics_path__
@@ -1548,6 +1554,8 @@ serverFiles:
       # following annotations:
       #
       # * `prometheus.io/scrape-slow`: Only scrape pods that have a value of `true`
+      # * `prometheus.io/scheme`: If the metrics endpoint is secured then you will need
+      # to set this to `https` & most likely set the `tls_config` of the scrape config.
       # * `prometheus.io/path`: If the metrics path is not `/metrics` override this.
       # * `prometheus.io/port`: Scrape the pod on the indicated port instead of the default of `9102`.
       - job_name: 'kubernetes-pods-slow'
@@ -1562,6 +1570,10 @@ serverFiles:
           - source_labels: [__meta_kubernetes_pod_annotation_prometheus_io_scrape_slow]
             action: keep
             regex: true
+          - source_labels: [__meta_kubernetes_pod_annotation_prometheus_io_scheme]
+            action: replace
+            regex: (https?)
+            target_label: __scheme__
           - source_labels: [__meta_kubernetes_pod_annotation_prometheus_io_path]
             action: replace
             target_label: __metrics_path__
```

## 13.0.3 

**Release date:** 2021-01-09

![AppVersion: 2.22.1](https://img.shields.io/static/v1?label=AppVersion&message=2.22.1&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Allow alertmanager to be deployed in IPv6 only environments  (#528) 

### Default value changes

```diff
# No changes in this release
```

## 13.0.2 

**Release date:** 2020-12-29

![AppVersion: 2.22.1](https://img.shields.io/static/v1?label=AppVersion&message=2.22.1&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* prometheus: fix chart version (#525) 
* Added absent allowedHostPath to psp for node-exporter (#470) 

### Default value changes

```diff
# No changes in this release
```

## 13.0.1 

**Release date:** 2020-12-21

![AppVersion: 2.22.1](https://img.shields.io/static/v1?label=AppVersion&message=2.22.1&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus] add .Capabilities.APIVersions for IngressClass fixes #507 (#508) 

### Default value changes

```diff
# No changes in this release
```

## 13.0.0 

**Release date:** 2020-12-08

![AppVersion: 2.22.1](https://img.shields.io/static/v1?label=AppVersion&message=2.22.1&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Prometheus - GH-417 Allow merging server sidecarContainers (#440) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 6637608..a3bc388 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -576,7 +576,14 @@ server:
   #   - yournamespace
 
   name: server
-  sidecarContainers:
+
+  # sidecarContainers - add more containers to prometheus server
+  # Key/Value where Key is the sidecar `- name: <Key>`
+  # Example:
+  #   sidecarContainers:
+  #      webserver:
+  #        image: nginx
+  sidecarContainers: {}
 
   ## Prometheus server container image
   ##
```

## 12.0.3 

**Release date:** 2020-12-08

![AppVersion: 2.22.1](https://img.shields.io/static/v1?label=AppVersion&message=2.22.1&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* change appversion to 2.22.1 matching prometheus server version (#459) 

### Default value changes

```diff
# No changes in this release
```

## 12.0.2 

**Release date:** 2020-12-05

![AppVersion: 2.21.0](https://img.shields.io/static/v1?label=AppVersion&message=2.21.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus] - Fix/node exporter host bug (#451) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 1d32377..6637608 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -530,8 +530,11 @@ nodeExporter:
 
   ## Security context to be added to node-exporter pods
   ##
-  securityContext: {}
-    # runAsUser: 0
+  securityContext:
+    fsGroup: 65534
+    runAsGroup: 65534
+    runAsNonRoot: true
+    runAsUser: 65534
 
   service:
     annotations:
```

## 12.0.1 

**Release date:** 2020-11-26

![AppVersion: 2.21.0](https://img.shields.io/static/v1?label=AppVersion&message=2.21.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus] Potentially change default liveness and readiness to abide by GKE limitations (#411) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 8bc0845..1d32377 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -892,12 +892,12 @@ server:
   ##
   readinessProbeInitialDelay: 30
   readinessProbePeriodSeconds: 5
-  readinessProbeTimeout: 30
+  readinessProbeTimeout: 4
   readinessProbeFailureThreshold: 3
   readinessProbeSuccessThreshold: 1
   livenessProbeInitialDelay: 30
   livenessProbePeriodSeconds: 15
-  livenessProbeTimeout: 30
+  livenessProbeTimeout: 10
   livenessProbeFailureThreshold: 3
   livenessProbeSuccessThreshold: 1
 
```

## 12.0.0 

**Release date:** 2020-11-22

![AppVersion: 2.21.0](https://img.shields.io/static/v1?label=AppVersion&message=2.21.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus] Migrate to chart v2 (#387) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 6244d18..8bc0845 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -47,7 +47,7 @@ alertmanager:
   ## alertmanager container image
   ##
   image:
-    repository: prom/alertmanager
+    repository: quay.io/prometheus/alertmanager
     tag: v0.21.0
     pullPolicy: IfNotPresent
 
@@ -437,7 +437,7 @@ nodeExporter:
   ## node-exporter container image
   ##
   image:
-    repository: prom/node-exporter
+    repository: quay.io/prometheus/node-exporter
     tag: v1.0.1
     pullPolicy: IfNotPresent
 
@@ -578,8 +578,8 @@ server:
   ## Prometheus server container image
   ##
   image:
-    repository: prom/prometheus
-    tag: v2.21.0
+    repository: quay.io/prometheus/prometheus
+    tag: v2.22.1
     pullPolicy: IfNotPresent
 
   ## prometheus server priorityClassName
@@ -984,7 +984,7 @@ pushgateway:
   ##
   image:
     repository: prom/pushgateway
-    tag: v1.2.0
+    tag: v1.3.0
     pullPolicy: IfNotPresent
 
   ## pushgateway priorityClassName
```

## 11.16.9 

**Release date:** 2020-11-14

![AppVersion: 2.21.0](https://img.shields.io/static/v1?label=AppVersion&message=2.21.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus] Use Alertmanager API port only for kubernetes_sd. (#353) 

### Default value changes

```diff
# No changes in this release
```

## 11.16.8 

**Release date:** 2020-11-06

![AppVersion: 2.21.0](https://img.shields.io/static/v1?label=AppVersion&message=2.21.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* charts/prometheus: fix typo (#319) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index fa2f40a..6244d18 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -1532,7 +1532,7 @@ serverFiles:
             action: drop
 
       # Example Scrape config for pods which should be scraped slower. An useful example
-      # would be stackriver-exporter which querys an API on every scrape of the pod
+      # would be stackriver-exporter which queries an API on every scrape of the pod
       #
       # The relabeling allows the actual pod scrape endpoint to be configured via the
       # following annotations:
```

## 11.16.7 

**Release date:** 2020-11-01

![AppVersion: 2.21.0](https://img.shields.io/static/v1?label=AppVersion&message=2.21.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Change helm stable repo (#290) 

### Default value changes

```diff
# No changes in this release
```

## 11.16.6 

**Release date:** 2020-10-31

![AppVersion: 2.21.0](https://img.shields.io/static/v1?label=AppVersion&message=2.21.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus] Fix rolebinding apiVersion (#287) 

### Default value changes

```diff
# No changes in this release
```

## 11.16.5 

**Release date:** 2020-10-24

![AppVersion: 2.21.0](https://img.shields.io/static/v1?label=AppVersion&message=2.21.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus] comment for prometheus base url fixing (#256) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 25bf6c7..fa2f40a 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -597,7 +597,7 @@ server:
   ## (Optional)
   prefixURL: ""
 
-  ## External URL which can access alertmanager
+  ## External URL which can access prometheus
   ## Maybe same with Ingress host name
   baseURL: ""
 
```

## 11.16.4 

**Release date:** 2020-10-21

![AppVersion: 2.21.0](https://img.shields.io/static/v1?label=AppVersion&message=2.21.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus] Make emptyDir settings consistent across components (#243) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 193542b..25bf6c7 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -224,6 +224,11 @@ alertmanager:
     ##
     subPath: ""
 
+  emptyDir:
+    ## alertmanager emptyDir volume size limit
+    ##
+    sizeLimit: ""
+
   ## Annotations to be added to alertmanager pods
   ##
   podAnnotations: {}
@@ -821,6 +826,8 @@ server:
     subPath: ""
 
   emptyDir:
+    ## Prometheus server emptyDir volume size limit
+    ##
     sizeLimit: ""
 
   ## Annotations to be added to Prometheus server pods
@@ -1114,7 +1121,6 @@ pushgateway:
 
   persistentVolume:
     ## If true, pushgateway will create/use a Persistent Volume Claim
-    ## If false, use emptyDir
     ##
     enabled: false
 
```

## 11.16.3 

**Release date:** 2020-10-19

![AppVersion: 2.21.0](https://img.shields.io/static/v1?label=AppVersion&message=2.21.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus] Create PSP only when corresponding components enabled (#233) 

### Default value changes

```diff
# No changes in this release
```

## 11.16.2 

**Release date:** 2020-10-05

![AppVersion: 2.21.0](https://img.shields.io/static/v1?label=AppVersion&message=2.21.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus] fix alertmanager meshpeer (#179) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 3a26f9c..193542b 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -277,7 +277,7 @@ alertmanager:
 
       ## Enabling peer mesh service end points for enabling the HA alert manager
       ## Ref: https://github.com/prometheus/alertmanager/blob/master/README.md
-      # enableMeshPeer : true
+      enableMeshPeer: false
 
       servicePort: 80
 
```

## 11.16.1 

**Release date:** 2020-10-05

![AppVersion: 2.21.0](https://img.shields.io/static/v1?label=AppVersion&message=2.21.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* propose to not try scrapping not running pods (#166) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 439c529..3a26f9c 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -1521,6 +1521,9 @@ serverFiles:
           - source_labels: [__meta_kubernetes_pod_name]
             action: replace
             target_label: kubernetes_pod_name
+          - source_labels: [__meta_kubernetes_pod_phase]
+            regex: Pending|Succeeded|Failed
+            action: drop
 
       # Example Scrape config for pods which should be scraped slower. An useful example
       # would be stackriver-exporter which querys an API on every scrape of the pod
@@ -1560,6 +1563,9 @@ serverFiles:
           - source_labels: [__meta_kubernetes_pod_name]
             action: replace
             target_label: kubernetes_pod_name
+          - source_labels: [__meta_kubernetes_pod_phase]
+            regex: Pending|Succeeded|Failed
+            action: drop
 
 # adds additional scrape configs to prometheus.yml
 # must be a string so you have to add a | after extraScrapeConfigs:
```

## 11.16.0 

**Release date:** 2020-10-02

![AppVersion: 2.21.0](https://img.shields.io/static/v1?label=AppVersion&message=2.21.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* update prometheus image to 2.21.0 (#173) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index cd941cf..439c529 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -574,7 +574,7 @@ server:
   ##
   image:
     repository: prom/prometheus
-    tag: v2.20.1
+    tag: v2.21.0
     pullPolicy: IfNotPresent
 
   ## prometheus server priorityClassName
```

## 11.15.0 

**Release date:** 2020-09-09

![AppVersion: 2.20.1](https://img.shields.io/static/v1?label=AppVersion&message=2.20.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus] group templated by component rather than resource type (#67) 

### Default value changes

```diff
# No changes in this release
```

## 11.14.1 

**Release date:** 2020-09-08

![AppVersion: 2.20.1](https://img.shields.io/static/v1?label=AppVersion&message=2.20.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Add naseemkullah as maintainer (#63) 

### Default value changes

```diff
# No changes in this release
```

## 11.14.0 

**Release date:** 2020-09-08

![AppVersion: 2.20.1](https://img.shields.io/static/v1?label=AppVersion&message=2.20.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus] set rbac api version with template rbac.apiVersion (#66) 

### Default value changes

```diff
# No changes in this release
```

## 11.13.1 

**Release date:** 2020-09-07

![AppVersion: 2.20.1](https://img.shields.io/static/v1?label=AppVersion&message=2.20.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus] - adds monotek to prometheus maintainers (#55) 

### Default value changes

```diff
# No changes in this release
```

## 11.13.0 

**Release date:** 2020-09-05

![AppVersion: 2.20.1](https://img.shields.io/static/v1?label=AppVersion&message=2.20.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus] unify labels and annotations across all deploymens and statefulsets (#45) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 1045f27..cd941cf 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -255,12 +255,18 @@ alertmanager:
   ##
   replicaCount: 1
 
+  ## Annotations to be added to deployment
+  ##
+  deploymentAnnotations: {}
+
   statefulSet:
     ## If true, use a statefulset instead of a deployment for pod management.
     ## This allows to scale replicas to more than 1 pod
     ##
     enabled: false
 
+    annotations: {}
+    labels: {}
     podManagementPolicy: OrderedReady
 
     ## Alertmanager headless service to use for the statefulset
@@ -848,6 +854,10 @@ server:
   ##
   replicaCount: 1
 
+  ## Annotations to be added to deployment
+  ##
+  deploymentAnnotations: {}
+
   statefulSet:
     ## If true, use a statefulset instead of a deployment for pod management.
     ## This allows to scale replicas to more than 1 pod
@@ -1034,6 +1044,10 @@ pushgateway:
   ##
   podAnnotations: {}
 
+  ## Labels to be added to pushgateway pods
+  ##
+  podLabels: {}
+
   ## Specify if a Pod Security Policy for node-exporter must be created
   ## Ref: https://kubernetes.io/docs/concepts/policy/pod-security-policy/
   ##
@@ -1050,6 +1064,10 @@ pushgateway:
 
   replicaCount: 1
 
+  ## Annotations to be added to deployment
+  ##
+  deploymentAnnotations: {}
+
   ## PodDisruptionBudget settings
   ## ref: https://kubernetes.io/docs/concepts/workloads/pods/disruptions/
   ##
```

## 11.12.1 

**Release date:** 2020-08-20

![AppVersion: 2.20.1](https://img.shields.io/static/v1?label=AppVersion&message=2.20.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Prep initial charts indexing (#14) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index c326cd6..1045f27 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -332,7 +332,7 @@ configmapReload:
     ##
     image:
       repository: jimmidyson/configmap-reload
-      tag: v0.3.0
+      tag: v0.4.0
       pullPolicy: IfNotPresent
 
     ## Additional configmap-reload container arguments
@@ -370,7 +370,7 @@ configmapReload:
     ##
     image:
       repository: jimmidyson/configmap-reload
-      tag: v0.3.0
+      tag: v0.4.0
       pullPolicy: IfNotPresent
 
     ## Additional configmap-reload container arguments
@@ -568,7 +568,7 @@ server:
   ##
   image:
     repository: prom/prometheus
-    tag: v2.19.2
+    tag: v2.20.1
     pullPolicy: IfNotPresent
 
   ## prometheus server priorityClassName
```

## 11.11.1 

**Release date:** 2020-07-30

![AppVersion: 2.19.0](https://img.shields.io/static/v1?label=AppVersion&message=2.19.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] Added init containers for prometheus server stateful set. (#23389) 

### Default value changes

```diff
# No changes in this release
```

## 11.11.0 

**Release date:** 2020-07-27

![AppVersion: 2.19.0](https://img.shields.io/static/v1?label=AppVersion&message=2.19.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] add alertmanger useExistingRole support (#23349) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 0bf2f0a..c326cd6 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -33,10 +33,13 @@ alertmanager:
   enabled: true
 
   ## Use a ClusterRole (and ClusterRoleBinding)
-  ## - If set to false, we define a Role and RoleBinding in the defined namespaces ONLY
+  ## - If set to false - we define a Role and RoleBinding in the defined namespaces ONLY
   ## This makes alertmanager work - for users who do not have ClusterAdmin privs, but wants alertmanager to operate on their own namespaces, instead of clusterwide.
   useClusterRole: true
 
+  ## Set to a rolename to use existing role - skipping role creating - but still doing serviceaccount and rolebinding to the rolename set here.
+  useExistingRole: false
+
   ## alertmanager container name
   ##
   name: alertmanager
```

## 11.10.1 

**Release date:** 2020-07-27

![AppVersion: 2.19.0](https://img.shields.io/static/v1?label=AppVersion&message=2.19.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] Document in README and Values "enableServiceLinks" (#23346) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index f653f46..0bf2f0a 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -572,7 +572,10 @@ server:
   ##
   priorityClassName: ""
 
-  # EnableServiceLinks indicates whether information about services should be injected into pod's environment variables, matching the syntax of Docker links.
+  ## EnableServiceLinks indicates whether information about services should be injected
+  ## into pod's environment variables, matching the syntax of Docker links.
+  ## WARNING: the field is unsupported and will be skipped in K8s prior to v1.13.0.
+  ##
   enableServiceLinks: true
 
   ## The URL prefix at which the container can be accessed. Useful in the case the '-web.external-url' includes a slug
```

## 11.10.0 

**Release date:** 2020-07-23

![AppVersion: 2.19.0](https://img.shields.io/static/v1?label=AppVersion&message=2.19.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] Disable "enableServiceLinks" on Kubernetes prior 1.13 (#23299) 

### Default value changes

```diff
# No changes in this release
```

## 11.9.1 

**Release date:** 2020-07-21

![AppVersion: 2.19.0](https://img.shields.io/static/v1?label=AppVersion&message=2.19.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* add a new owner and approver (#23276) 

### Default value changes

```diff
# No changes in this release
```

## 11.9.0 

**Release date:** 2020-07-21

![AppVersion: 2.19.0](https://img.shields.io/static/v1?label=AppVersion&message=2.19.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] Organized Templates Directory. Add ".gitignore". (#23186) 

### Default value changes

```diff
# No changes in this release
```

## 11.8.0 

**Release date:** 2020-07-21

![AppVersion: 2.19.0](https://img.shields.io/static/v1?label=AppVersion&message=2.19.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] bump alertmanager v0.21.0 and prometheus v2.19.2 (#23253) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 35d40a6..f653f46 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -45,7 +45,7 @@ alertmanager:
   ##
   image:
     repository: prom/alertmanager
-    tag: v0.20.0
+    tag: v0.21.0
     pullPolicy: IfNotPresent
 
   ## alertmanager priorityClassName
@@ -565,7 +565,7 @@ server:
   ##
   image:
     repository: prom/prometheus
-    tag: v2.19.0
+    tag: v2.19.2
     pullPolicy: IfNotPresent
 
   ## prometheus server priorityClassName
```

## 11.7.0 

**Release date:** 2020-07-09

![AppVersion: 2.19.0](https://img.shields.io/static/v1?label=AppVersion&message=2.19.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] Improve Prometheus Server Probes. Fix MD Syntax Issues in README. (#23116) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index ef05b89..35d40a6 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -868,10 +868,12 @@ server:
   ## Ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/
   ##
   readinessProbeInitialDelay: 30
+  readinessProbePeriodSeconds: 5
   readinessProbeTimeout: 30
   readinessProbeFailureThreshold: 3
   readinessProbeSuccessThreshold: 1
   livenessProbeInitialDelay: 30
+  livenessProbePeriodSeconds: 15
   livenessProbeTimeout: 30
   livenessProbeFailureThreshold: 3
   livenessProbeSuccessThreshold: 1
```

## 11.6.2 

**Release date:** 2020-07-08

![AppVersion: 2.19.0](https://img.shields.io/static/v1?label=AppVersion&message=2.19.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] add support for running without cluster-admin privileges (#23049) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 0e1cdad..ef05b89 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -32,6 +32,11 @@ alertmanager:
   ##
   enabled: true
 
+  ## Use a ClusterRole (and ClusterRoleBinding)
+  ## - If set to false, we define a Role and RoleBinding in the defined namespaces ONLY
+  ## This makes alertmanager work - for users who do not have ClusterAdmin privs, but wants alertmanager to operate on their own namespaces, instead of clusterwide.
+  useClusterRole: true
+
   ## alertmanager container name
   ##
   name: alertmanager
@@ -538,6 +543,21 @@ server:
   ## Prometheus server container name
   ##
   enabled: true
+
+  ## Use a ClusterRole (and ClusterRoleBinding)
+  ## - If set to false - we define a RoleBinding in the defined namespaces ONLY
+  ##
+  ## NB: because we need a Role with nonResourceURL's ("/metrics") - you must get someone with Cluster-admin privileges to define this role for you, before running with this setting enabled.
+  ##     This makes prometheus work - for users who do not have ClusterAdmin privs, but wants prometheus to operate on their own namespaces, instead of clusterwide.
+  ##
+  ## You MUST also set namespaces to the ones you have access to and want monitored by Prometheus.
+  ##
+  # useExistingClusterRoleName: nameofclusterrole
+
+  ## namespaces to monitor (instead of monitoring all - clusterwide). Needed if you want to run without Cluster-admin privileges.
+  # namespaces:
+  #   - yournamespace
+
   name: server
   sidecarContainers:
 
```

## 11.6.1 

**Release date:** 2020-07-02

![AppVersion: 2.19.0](https://img.shields.io/static/v1?label=AppVersion&message=2.19.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] Fix rbac for get ingress info (#23015) 

### Default value changes

```diff
# No changes in this release
```

## 11.6.0 

**Release date:** 2020-06-19

![AppVersion: 2.19.0](https://img.shields.io/static/v1?label=AppVersion&message=2.19.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] updated all prometheus docker images to latest versions (#22849) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 1ad0799..0e1cdad 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -419,7 +419,7 @@ nodeExporter:
   ##
   image:
     repository: prom/node-exporter
-    tag: v0.18.1
+    tag: v1.0.1
     pullPolicy: IfNotPresent
 
   ## Specify if a Pod Security Policy for node-exporter must be created
@@ -545,7 +545,7 @@ server:
   ##
   image:
     repository: prom/prometheus
-    tag: v2.18.1
+    tag: v2.19.0
     pullPolicy: IfNotPresent
 
   ## prometheus server priorityClassName
@@ -939,7 +939,7 @@ pushgateway:
   ##
   image:
     repository: prom/pushgateway
-    tag: v1.0.1
+    tag: v1.2.0
     pullPolicy: IfNotPresent
 
   ## pushgateway priorityClassName
```

## 11.5.0 

**Release date:** 2020-06-16

![AppVersion: 2.18.1](https://img.shields.io/static/v1?label=AppVersion&message=2.18.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] gRPC port on headless svc (#22784) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index b3af346..1ad0799 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -838,6 +838,11 @@ server:
       annotations: {}
       labels: {}
       servicePort: 80
+      ## Enable gRPC port on service to allow auto discovery with thanos-querier
+      gRPC:
+        enabled: false
+        servicePort: 10901
+        # nodePort: 10901
 
   ## Prometheus server readiness and liveness probe initial delay and timeout
   ## Ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/
```

## 11.4.0 

**Release date:** 2020-06-04

![AppVersion: 2.18.1](https://img.shields.io/static/v1?label=AppVersion&message=2.18.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] Add configFromSecret support to StatefulSet alertmanagers (#22647) 

### Default value changes

```diff
# No changes in this release
```

## 11.3.1 

**Release date:** 2020-06-03

![AppVersion: 2.18.1](https://img.shields.io/static/v1?label=AppVersion&message=2.18.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Fix stable/prometheus node-exporter containerPort to hostPort (#22141) 

### Default value changes

```diff
# No changes in this release
```

## 11.3.0 

**Release date:** 2020-05-18

![AppVersion: 2.18.1](https://img.shields.io/static/v1?label=AppVersion&message=2.18.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] Add Namespace to applicable Resources (#22378) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index ecd5165..b3af346 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -1541,3 +1541,6 @@ networkPolicy:
   ## Enable creation of NetworkPolicy resources.
   ##
   enabled: false
+
+# Force namespace of namespaced resources
+forceNamespace: null
```

## 11.2.3 

**Release date:** 2020-05-14

![AppVersion: 2.18.1](https://img.shields.io/static/v1?label=AppVersion&message=2.18.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] Fix incorrect rendering of enableServiceLinks (#22391) 

### Default value changes

```diff
# No changes in this release
```

## 11.2.2 

**Release date:** 2020-05-14

![AppVersion: 2.18.1](https://img.shields.io/static/v1?label=AppVersion&message=2.18.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] Fix default non-HA state for alertmanager statefulset (#22386) 

### Default value changes

```diff
# No changes in this release
```

## 11.2.1 

**Release date:** 2020-05-13

![AppVersion: 2.18.1](https://img.shields.io/static/v1?label=AppVersion&message=2.18.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* fix default non-ha state for alert-manager statefulset (#22350) 

### Default value changes

```diff
# No changes in this release
```

## 11.2.0 

**Release date:** 2020-05-12

![AppVersion: 2.18.1](https://img.shields.io/static/v1?label=AppVersion&message=2.18.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* upgrade to 2.18.1 (#22349) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 3ad64e5..ecd5165 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -545,7 +545,7 @@ server:
   ##
   image:
     repository: prom/prometheus
-    tag: v2.16.0
+    tag: v2.18.1
     pullPolicy: IfNotPresent
 
   ## prometheus server priorityClassName
```

## 11.1.6 

**Release date:** 2020-05-09

![AppVersion: 2.16.0](https://img.shields.io/static/v1?label=AppVersion&message=2.16.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] Configure enableServiceLinks for server (#22300) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 6c49edf..3ad64e5 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -552,6 +552,9 @@ server:
   ##
   priorityClassName: ""
 
+  # EnableServiceLinks indicates whether information about services should be injected into pod's environment variables, matching the syntax of Docker links.
+  enableServiceLinks: true
+
   ## The URL prefix at which the container can be accessed. Useful in the case the '-web.external-url' includes a slug
   ## so that the various internal URLs are still able to access as they are in the default case.
   ## (Optional)
```

## 11.1.5 

**Release date:** 2020-04-29

![AppVersion: 2.16.0](https://img.shields.io/static/v1?label=AppVersion&message=2.16.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] fix remoteRead/Write defaults (#22088) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index c16a2f4..6c49edf 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -605,10 +605,10 @@ server:
     evaluation_interval: 1m
   ## https://prometheus.io/docs/prometheus/latest/configuration/configuration/#remote_write
   ##
-  remoteWrite: {}
+  remoteWrite: []
   ## https://prometheus.io/docs/prometheus/latest/configuration/configuration/#remote_read
   ##
-  remoteRead: {}
+  remoteRead: []
 
   ## Additional Prometheus server container arguments
   ##
```

## 11.1.4 

**Release date:** 2020-04-27

![AppVersion: 2.16.0](https://img.shields.io/static/v1?label=AppVersion&message=2.16.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] De-duplicate "affinity" for Prometheus StatefulSet Object (#22129) 

### Default value changes

```diff
# No changes in this release
```

## 11.1.3 

**Release date:** 2020-04-25

![AppVersion: 2.16.0](https://img.shields.io/static/v1?label=AppVersion&message=2.16.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] Fix prometheus service selector with statefulset (#22098) 

### Default value changes

```diff
# No changes in this release
```

## 11.1.2 

**Release date:** 2020-04-17

![AppVersion: 2.16.0](https://img.shields.io/static/v1?label=AppVersion&message=2.16.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Fix rollingUpdate: null indent for helm3 (#21989) 

### Default value changes

```diff
# No changes in this release
```

## 11.1.1 

**Release date:** 2020-04-16

![AppVersion: 2.16.0](https://img.shields.io/static/v1?label=AppVersion&message=2.16.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Fix rollingUpdate: null indentation Recreate (#21976) 

### Default value changes

```diff
# No changes in this release
```

## 11.1.0 

**Release date:** 2020-04-11

![AppVersion: 2.16.0](https://img.shields.io/static/v1?label=AppVersion&message=2.16.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Cleanup ksm. (#21877) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 3d10f0b..c16a2f4 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -390,10 +390,14 @@ configmapReload:
 
 kubeStateMetrics:
   ## If false, kube-state-metrics sub-chart will not be installed
-  ## Please see https://github.com/helm/charts/tree/master/stable/kube-state-metrics for configurable values
   ##
   enabled: true
 
+## kube-state-metrics sub-chart configurable values
+## Please see https://github.com/helm/charts/tree/master/stable/kube-state-metrics
+##
+# kube-state-metrics:
+
 nodeExporter:
   ## If false, node-exporter will not be installed
   ##
```

## 11.0.6 

**Release date:** 2020-04-09

![AppVersion: 2.16.0](https://img.shields.io/static/v1?label=AppVersion&message=2.16.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] add hostAliases (#21805) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 9288ee8..3d10f0b 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -694,6 +694,12 @@ server:
   # strategy:
   #   type: Recreate
 
+  ## hostAliases allows adding entries to /etc/hosts inside the containers
+  hostAliases: []
+  #   - ip: "127.0.0.1"
+  #     hostnames:
+  #       - "example.com"
+
   ## Node tolerations for server scheduling to nodes with taints
   ## Ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
   ##
```

## 11.0.5 

**Release date:** 2020-04-09

![AppVersion: 2.16.0](https://img.shields.io/static/v1?label=AppVersion&message=2.16.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Added extraInitContainers to match server (#21774) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 5b3998c..9288ee8 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -51,6 +51,10 @@ alertmanager:
   ##
   extraArgs: {}
 
+  ## Additional InitContainers to initialize the pod
+  ##
+  extraInitContainers: []
+
   ## The URL prefix at which the container can be accessed. Useful in the case the '-web.external-url' includes a slug
   ## so that the various internal URLs are still able to access as they are in the default case.
   ## (Optional)
@@ -441,6 +445,10 @@ nodeExporter:
   ##
   extraArgs: {}
 
+  ## Additional InitContainers to initialize the pod
+  ##
+  extraInitContainers: []
+
   ## Additional node-exporter hostPath mounts
   ##
   extraHostPathMounts: []
@@ -925,6 +933,10 @@ pushgateway:
   ## for example: persistence.file: /data/pushgateway.data
   extraArgs: {}
 
+  ## Additional InitContainers to initialize the pod
+  ##
+  extraInitContainers: []
+
   ingress:
     ## If true, pushgateway Ingress will be created
     ##
```

## 11.0.4 

**Release date:** 2020-03-26

![AppVersion: 2.16.0](https://img.shields.io/static/v1?label=AppVersion&message=2.16.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Moving `cluster.advertise-address` param to if block for peer mesh (#21617) 

### Default value changes

```diff
# No changes in this release
```

## 11.0.3 

**Release date:** 2020-03-21

![AppVersion: 2.16.0](https://img.shields.io/static/v1?label=AppVersion&message=2.16.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] Make service account annotations configurable (#21537) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index a95df04..5b3998c 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -13,15 +13,19 @@ serviceAccounts:
   alertmanager:
     create: true
     name:
+    annotations: {}
   nodeExporter:
     create: true
     name:
+    annotations: {}
   pushgateway:
     create: true
     name:
+    annotations: {}
   server:
     create: true
     name:
+    annotations: {}
 
 alertmanager:
   ## If false, alertmanager will not be installed
```

## 11.0.2 

**Release date:** 2020-03-05

![AppVersion: 2.16.0](https://img.shields.io/static/v1?label=AppVersion&message=2.16.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] Set rollingUpdate to null when strategy is Recreate (#21289) 

### Default value changes

```diff
# No changes in this release
```

## 11.0.1 

**Release date:** 2020-03-03

![AppVersion: 2.16.0](https://img.shields.io/static/v1?label=AppVersion&message=2.16.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] Support tmpl files which are not YAML (#21208) 

### Default value changes

```diff
# No changes in this release
```

## 11.0.0 

**Release date:** 2020-03-01

![AppVersion: 2.16.0](https://img.shields.io/static/v1?label=AppVersion&message=2.16.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Use ksm chart for ksm (#21148) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index c6d8d9c..a95df04 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -13,9 +13,6 @@ serviceAccounts:
   alertmanager:
     create: true
     name:
-  kubeStateMetrics:
-    create: true
-    name:
   nodeExporter:
     create: true
     name:
@@ -383,113 +380,12 @@ configmapReload:
     ##
     resources: {}
 
-
 kubeStateMetrics:
-  ## If false, kube-state-metrics will not be installed
+  ## If false, kube-state-metrics sub-chart will not be installed
+  ## Please see https://github.com/helm/charts/tree/master/stable/kube-state-metrics for configurable values
   ##
   enabled: true
 
-  ## kube-state-metrics container name
-  ##
-  name: kube-state-metrics
-
-  ## kube-state-metrics container image
-  ##
-  image:
-    repository: quay.io/coreos/kube-state-metrics
-    tag: v1.9.5
-    pullPolicy: IfNotPresent
-
-  ## kube-state-metrics priorityClassName
-  ##
-  priorityClassName: ""
-
-  ## kube-state-metrics container arguments
-  ##
-  args: {}
-
-  ## Node tolerations for kube-state-metrics scheduling to nodes with taints
-  ## Ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
-  ##
-  tolerations: []
-    # - key: "key"
-    #   operator: "Equal|Exists"
-    #   value: "value"
-    #   effect: "NoSchedule|PreferNoSchedule|NoExecute(1.6 only)"
-
-  ## Node labels for kube-state-metrics pod assignment
-  ## Ref: https://kubernetes.io/docs/user-guide/node-selection/
-  ##
-  nodeSelector: {}
-
-  ## Annotations to be added to kube-state-metrics pods
-  ##
-  podAnnotations: {}
-
-  ## Specify if a Pod Security Policy for node-exporter must be created
-  ## Ref: https://kubernetes.io/docs/concepts/policy/pod-security-policy/
-  ##
-  podSecurityPolicy:
-    annotations: {}
-      ## Specify pod annotations
-      ## Ref: https://kubernetes.io/docs/concepts/policy/pod-security-policy/#apparmor
-      ## Ref: https://kubernetes.io/docs/concepts/policy/pod-security-policy/#seccomp
-      ## Ref: https://kubernetes.io/docs/concepts/policy/pod-security-policy/#sysctl
-      ##
-      # seccomp.security.alpha.kubernetes.io/allowedProfileNames: '*'
-      # seccomp.security.alpha.kubernetes.io/defaultProfileName: 'docker/default'
-      # apparmor.security.beta.kubernetes.io/defaultProfileName: 'runtime/default'
-
-  pod:
-    labels: {}
-
-  replicaCount: 1
-
-  ## PodDisruptionBudget settings
-  ## ref: https://kubernetes.io/docs/concepts/workloads/pods/disruptions/
-  ##
-  podDisruptionBudget:
-    enabled: false
-    maxUnavailable: 1
-
-  ## kube-state-metrics resource requests and limits
-  ## Ref: http://kubernetes.io/docs/user-guide/compute-resources/
-  ##
-  resources: {}
-    # limits:
-    #   cpu: 10m
-    #   memory: 16Mi
-    # requests:
-    #   cpu: 10m
-    #   memory: 16Mi
-
-  ## Security context to be added to kube-state-metrics pods
-  ##
-  securityContext:
-    runAsUser: 65534
-    runAsNonRoot: true
-
-  service:
-    annotations:
-      prometheus.io/scrape: "true"
-    labels: {}
-
-    # Exposed as a headless service:
-    # https://kubernetes.io/docs/concepts/services-networking/service/#headless-services
-    clusterIP: None
-
-    ## List of IP addresses at which the kube-state-metrics service is available
-    ## Ref: https://kubernetes.io/docs/user-guide/services/#external-ips
-    ##
-    externalIPs: []
-
-    loadBalancerIP: ""
-    loadBalancerSourceRanges: []
-    servicePort: 80
-    # Port for Kubestatemetric self telemetry
-    serviceTelemetryPort: 81
-    type: ClusterIP
-
 nodeExporter:
   ## If false, node-exporter will not be installed
   ##
```

## 10.6.0 

**Release date:** 2020-02-29

![AppVersion: 2.16.0](https://img.shields.io/static/v1?label=AppVersion&message=2.16.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] Added configmapReload.enabled parameter (#20578) (#20912) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 0f5ba46..c6d8d9c 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -306,39 +306,83 @@ alertmanager:
 ## Ref: https://github.com/jimmidyson/configmap-reload
 ##
 configmapReload:
-  ## configmap-reload container name
-  ##
-  name: configmap-reload
+  prometheus:
+    ## If false, the configmap-reload container will not be deployed
+    ##
+    enabled: true
 
-  ## configmap-reload container image
-  ##
-  image:
-    repository: jimmidyson/configmap-reload
-    tag: v0.3.0
-    pullPolicy: IfNotPresent
+    ## configmap-reload container name
+    ##
+    name: configmap-reload
 
-  ## Additional configmap-reload container arguments
-  ##
-  extraArgs: {}
-  ## Additional configmap-reload volume directories
-  ##
-  extraVolumeDirs: []
+    ## configmap-reload container image
+    ##
+    image:
+      repository: jimmidyson/configmap-reload
+      tag: v0.3.0
+      pullPolicy: IfNotPresent
 
+    ## Additional configmap-reload container arguments
+    ##
+    extraArgs: {}
+    ## Additional configmap-reload volume directories
+    ##
+    extraVolumeDirs: []
 
-  ## Additional configmap-reload mounts
-  ##
-  extraConfigmapMounts: []
-    # - name: prometheus-alerts
-    #   mountPath: /etc/alerts.d
-    #   subPath: ""
-    #   configMap: prometheus-alerts
-    #   readOnly: true
 
+    ## Additional configmap-reload mounts
+    ##
+    extraConfigmapMounts: []
+      # - name: prometheus-alerts
+      #   mountPath: /etc/alerts.d
+      #   subPath: ""
+      #   configMap: prometheus-alerts
+      #   readOnly: true
+
+
+    ## configmap-reload resource requests and limits
+    ## Ref: http://kubernetes.io/docs/user-guide/compute-resources/
+    ##
+    resources: {}
+  alertmanager:
+    ## If false, the configmap-reload container will not be deployed
+    ##
+    enabled: true
+
+    ## configmap-reload container name
+    ##
+    name: configmap-reload
+
+    ## configmap-reload container image
+    ##
+    image:
+      repository: jimmidyson/configmap-reload
+      tag: v0.3.0
+      pullPolicy: IfNotPresent
+
+    ## Additional configmap-reload container arguments
+    ##
+    extraArgs: {}
+    ## Additional configmap-reload volume directories
+    ##
+    extraVolumeDirs: []
+
+
+    ## Additional configmap-reload mounts
+    ##
+    extraConfigmapMounts: []
+      # - name: prometheus-alerts
+      #   mountPath: /etc/alerts.d
+      #   subPath: ""
+      #   configMap: prometheus-alerts
+      #   readOnly: true
+
+
+    ## configmap-reload resource requests and limits
+    ## Ref: http://kubernetes.io/docs/user-guide/compute-resources/
+    ##
+    resources: {}
 
-  ## configmap-reload resource requests and limits
-  ## Ref: http://kubernetes.io/docs/user-guide/compute-resources/
-  ##
-  resources: {}
 
 kubeStateMetrics:
   ## If false, kube-state-metrics will not be installed
```

## 10.5.3 

**Release date:** 2020-02-28

![AppVersion: 2.16.0](https://img.shields.io/static/v1?label=AppVersion&message=2.16.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] Use the new api for ingress (#21049) 

### Default value changes

```diff
# No changes in this release
```

## 10.5.2 

**Release date:** 2020-02-26

![AppVersion: 2.16.0](https://img.shields.io/static/v1?label=AppVersion&message=2.16.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] gRPC port on prometheus service (#21002) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 116cb6b..0f5ba46 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -930,6 +930,12 @@ server:
     sessionAffinity: None
     type: ClusterIP
 
+    ## Enable gRPC port on service to allow auto discovery with thanos-querier
+    gRPC:
+      enabled: false
+      servicePort: 10901
+      # nodePort: 10901
+
     ## If using a statefulSet (statefulSet.enabled=true), configure the
     ## service to connect to a specific replica to have a consistent view
     ## of the data.
```

## 10.5.1 

**Release date:** 2020-02-25

![AppVersion: 2.16.0](https://img.shields.io/static/v1?label=AppVersion&message=2.16.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Update kube-state-metrics (#21017) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 436fce6..116cb6b 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -353,7 +353,7 @@ kubeStateMetrics:
   ##
   image:
     repository: quay.io/coreos/kube-state-metrics
-    tag: v1.9.1
+    tag: v1.9.5
     pullPolicy: IfNotPresent
 
   ## kube-state-metrics priorityClassName
```

## 10.5.0 

**Release date:** 2020-02-25

![AppVersion: 2.16.0](https://img.shields.io/static/v1?label=AppVersion&message=2.16.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Update prometheus to 2.16.0 (#20821) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index a00a00f..436fce6 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -589,7 +589,7 @@ server:
   ##
   image:
     repository: prom/prometheus
-    tag: v2.15.2
+    tag: v2.16.0
     pullPolicy: IfNotPresent
 
   ## prometheus server priorityClassName
```

## 10.4.1 

**Release date:** 2020-02-25

![AppVersion: 2.15.2](https://img.shields.io/static/v1?label=AppVersion&message=2.15.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] Fix server.alertmanagers (#20779) 

### Default value changes

```diff
# No changes in this release
```

## 10.4.0 

**Release date:** 2020-01-30

![AppVersion: 2.15.2](https://img.shields.io/static/v1?label=AppVersion&message=2.15.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] added slow pod scrape config (#20443) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index d92f231..a00a00f 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -1496,6 +1496,45 @@ serverFiles:
             action: replace
             target_label: kubernetes_pod_name
 
+      # Example Scrape config for pods which should be scraped slower. An useful example
+      # would be stackriver-exporter which querys an API on every scrape of the pod
+      #
+      # The relabeling allows the actual pod scrape endpoint to be configured via the
+      # following annotations:
+      #
+      # * `prometheus.io/scrape-slow`: Only scrape pods that have a value of `true`
+      # * `prometheus.io/path`: If the metrics path is not `/metrics` override this.
+      # * `prometheus.io/port`: Scrape the pod on the indicated port instead of the default of `9102`.
+      - job_name: 'kubernetes-pods-slow'
+
+        scrape_interval: 5m
+        scrape_timeout: 30s
+
+        kubernetes_sd_configs:
+          - role: pod
+
+        relabel_configs:
+          - source_labels: [__meta_kubernetes_pod_annotation_prometheus_io_scrape_slow]
+            action: keep
+            regex: true
+          - source_labels: [__meta_kubernetes_pod_annotation_prometheus_io_path]
+            action: replace
+            target_label: __metrics_path__
+            regex: (.+)
+          - source_labels: [__address__, __meta_kubernetes_pod_annotation_prometheus_io_port]
+            action: replace
+            regex: ([^:]+)(?::\d+)?;(\d+)
+            replacement: $1:$2
+            target_label: __address__
+          - action: labelmap
+            regex: __meta_kubernetes_pod_label_(.+)
+          - source_labels: [__meta_kubernetes_namespace]
+            action: replace
+            target_label: kubernetes_namespace
+          - source_labels: [__meta_kubernetes_pod_name]
+            action: replace
+            target_label: kubernetes_pod_name
+
 # adds additional scrape configs to prometheus.yml
 # must be a string so you have to add a | after extraScrapeConfigs:
 # example adds prometheus-blackbox-exporter scrape config
```

## 10.3.1 

**Release date:** 2020-01-16

![AppVersion: 2.15.2](https://img.shields.io/static/v1?label=AppVersion&message=2.15.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus] fix: alertmanager to use /-/ready as readiness probe (#20190) 

### Default value changes

```diff
# No changes in this release
```

## 10.3.0 

**Release date:** 2020-01-16

![AppVersion: 2.15.2](https://img.shields.io/static/v1?label=AppVersion&message=2.15.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] prometheus.io/scrape-slow for scraping slow target (#19930) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 3ed0fd5..d92f231 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -1369,6 +1369,55 @@ serverFiles:
             action: replace
             target_label: kubernetes_node
 
+      # Scrape config for slow service endpoints; same as above, but with a larger
+      # timeout and a larger interval
+      #
+      # The relabeling allows the actual service scrape endpoint to be configured
+      # via the following annotations:
+      #
+      # * `prometheus.io/scrape-slow`: Only scrape services that have a value of `true`
+      # * `prometheus.io/scheme`: If the metrics endpoint is secured then you will need
+      # to set this to `https` & most likely set the `tls_config` of the scrape config.
+      # * `prometheus.io/path`: If the metrics path is not `/metrics` override this.
+      # * `prometheus.io/port`: If the metrics are exposed on a different port to the
+      # service then set this appropriately.
+      - job_name: 'kubernetes-service-endpoints-slow'
+
+        scrape_interval: 5m
+        scrape_timeout: 30s
+
+        kubernetes_sd_configs:
+          - role: endpoints
+
+        relabel_configs:
+          - source_labels: [__meta_kubernetes_service_annotation_prometheus_io_scrape_slow]
+            action: keep
+            regex: true
+          - source_labels: [__meta_kubernetes_service_annotation_prometheus_io_scheme]
+            action: replace
+            target_label: __scheme__
+            regex: (https?)
+          - source_labels: [__meta_kubernetes_service_annotation_prometheus_io_path]
+            action: replace
+            target_label: __metrics_path__
+            regex: (.+)
+          - source_labels: [__address__, __meta_kubernetes_service_annotation_prometheus_io_port]
+            action: replace
+            target_label: __address__
+            regex: ([^:]+)(?::\d+)?;(\d+)
+            replacement: $1:$2
+          - action: labelmap
+            regex: __meta_kubernetes_service_label_(.+)
+          - source_labels: [__meta_kubernetes_namespace]
+            action: replace
+            target_label: kubernetes_namespace
+          - source_labels: [__meta_kubernetes_service_name]
+            action: replace
+            target_label: kubernetes_name
+          - source_labels: [__meta_kubernetes_pod_node_name]
+            action: replace
+            target_label: kubernetes_node
+
       - job_name: 'prometheus-pushgateway'
         honor_labels: true
 
```

## 10.2.0 

**Release date:** 2020-01-13

![AppVersion: 2.15.2](https://img.shields.io/static/v1?label=AppVersion&message=2.15.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] Add Telemetry port to Kube-state-metrics (#20072) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 89ba207..3ed0fd5 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -353,7 +353,7 @@ kubeStateMetrics:
   ##
   image:
     repository: quay.io/coreos/kube-state-metrics
-    tag: v1.9.0
+    tag: v1.9.1
     pullPolicy: IfNotPresent
 
   ## kube-state-metrics priorityClassName
@@ -442,6 +442,8 @@ kubeStateMetrics:
     loadBalancerIP: ""
     loadBalancerSourceRanges: []
     servicePort: 80
+    # Port for Kubestatemetric self telemetry
+    serviceTelemetryPort: 81
     type: ClusterIP
 
 nodeExporter:
```

## 10.1.0 

**Release date:** 2020-01-13

![AppVersion: 2.15.2](https://img.shields.io/static/v1?label=AppVersion&message=2.15.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] Added remote_write and remote_read for prometheus (#20056) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 3cc1227..89ba207 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -645,6 +645,12 @@ server:
     ## How frequently to evaluate rules
     ##
     evaluation_interval: 1m
+  ## https://prometheus.io/docs/prometheus/latest/configuration/configuration/#remote_write
+  ##
+  remoteWrite: {}
+  ## https://prometheus.io/docs/prometheus/latest/configuration/configuration/#remote_read
+  ##
+  remoteRead: {}
 
   ## Additional Prometheus server container arguments
   ##
```

## 10.0.2 

**Release date:** 2020-01-12

![AppVersion: 2.15.2](https://img.shields.io/static/v1?label=AppVersion&message=2.15.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] Add option to set labels for Prometheus AlertManager pods (#20026) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index dfe4a13..3cc1227 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -220,6 +220,10 @@ alertmanager:
     ##
     ## prometheus.io/probe: alertmanager-teamA
 
+  ## Labels to be added to Prometheus AlertManager pods
+  ##
+  podLabels: {}
+
   ## Specify if a Pod Security Policy for node-exporter must be created
   ## Ref: https://kubernetes.io/docs/concepts/policy/pod-security-policy/
   ##
```

## 10.0.1 

**Release date:** 2020-01-10

![AppVersion: 2.15.2](https://img.shields.io/static/v1?label=AppVersion&message=2.15.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] Add Missing ClusterRole Permissions for Kube-state-metrics (#20004) 

### Default value changes

```diff
# No changes in this release
```

## 10.0.0 

**Release date:** 2020-01-09

![AppVersion: 2.15.2](https://img.shields.io/static/v1?label=AppVersion&message=2.15.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] update prometheus to 2.15.2 and also all other used images to newest … (#19927) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 356b697..dfe4a13 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -39,7 +39,7 @@ alertmanager:
   ##
   image:
     repository: prom/alertmanager
-    tag: v0.18.0
+    tag: v0.20.0
     pullPolicy: IfNotPresent
 
   ## alertmanager priorityClassName
@@ -56,8 +56,7 @@ alertmanager:
   prefixURL: ""
 
   ## External URL which can access alertmanager
-  ## Maybe same with Ingress host name
-  baseURL: "/"
+  baseURL: "http://localhost:9093"
 
   ## Additional alertmanager container environment variable
   ## For instance to add a http_proxy
@@ -311,7 +310,7 @@ configmapReload:
   ##
   image:
     repository: jimmidyson/configmap-reload
-    tag: v0.2.2
+    tag: v0.3.0
     pullPolicy: IfNotPresent
 
   ## Additional configmap-reload container arguments
@@ -350,7 +349,7 @@ kubeStateMetrics:
   ##
   image:
     repository: quay.io/coreos/kube-state-metrics
-    tag: v1.6.0
+    tag: v1.9.0
     pullPolicy: IfNotPresent
 
   ## kube-state-metrics priorityClassName
@@ -462,7 +461,7 @@ nodeExporter:
   ##
   image:
     repository: prom/node-exporter
-    tag: v0.18.0
+    tag: v0.18.1
     pullPolicy: IfNotPresent
 
   ## Specify if a Pod Security Policy for node-exporter must be created
@@ -584,7 +583,7 @@ server:
   ##
   image:
     repository: prom/prometheus
-    tag: v2.13.1
+    tag: v2.15.2
     pullPolicy: IfNotPresent
 
   ## prometheus server priorityClassName
@@ -952,7 +951,7 @@ pushgateway:
   ##
   image:
     repository: prom/pushgateway
-    tag: v0.8.0
+    tag: v1.0.1
     pullPolicy: IfNotPresent
 
   ## pushgateway priorityClassName
```

## 9.7.5 

**Release date:** 2020-01-08

![AppVersion: 2.13.1](https://img.shields.io/static/v1?label=AppVersion&message=2.13.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* stable/prometheus Updated README.md with instructions for configuring pod scraping (#19379) 

### Default value changes

```diff
# No changes in this release
```

## 9.7.4 

**Release date:** 2020-01-06

![AppVersion: 2.13.1](https://img.shields.io/static/v1?label=AppVersion&message=2.13.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] server.env default value should be an array (#19141) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 2616e6a..356b697 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -615,7 +615,7 @@ server:
   ##     secretKeyRef:
   ##       name: mysecret
   ##       key: username
-  env: {}
+  env: []
 
   extraFlags:
     - web.enable-lifecycle
```

## 9.7.3 

**Release date:** 2020-01-04

![AppVersion: 2.13.1](https://img.shields.io/static/v1?label=AppVersion&message=2.13.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Merge branch 'master' of https://github.com/helm/charts (#17071) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 871a1fa..2616e6a 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -617,12 +617,17 @@ server:
   ##       key: username
   env: {}
 
-  ## This flag controls access to the administrative HTTP API which includes functionality such as deleting time
-  ## series. This is disabled by default.
-  enableAdminApi: false
-
-  ## This flag controls BD locking
-  skipTSDBLock: false
+  extraFlags:
+    - web.enable-lifecycle
+    ## web.enable-admin-api flag controls access to the administrative HTTP API which includes functionality such as
+    ## deleting time series. This is disabled by default.
+    # - web.enable-admin-api
+    ##
+    ## storage.tsdb.no-lockfile flag controls BD locking
+    # - storage.tsdb.no-lockfile
+    ##
+    ## storage.tsdb.wal-compression flag enables compression of the write-ahead log (WAL)
+    # - storage.tsdb.wal-compression
 
   ## Path to a configuration file on prometheus server container FS
   configPath: /etc/config/prometheus.yml
```

## 9.7.2 

**Release date:** 2019-12-22

![AppVersion: 2.13.1](https://img.shields.io/static/v1?label=AppVersion&message=2.13.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] Adding Vertical Pod Autoscaler support (#17742) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 36fdfea..871a1fa 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -881,6 +881,15 @@ server:
     #   cpu: 500m
     #   memory: 512Mi
 
+  ## Vertical Pod Autoscaler config
+  ## Ref: https://github.com/kubernetes/autoscaler/tree/master/vertical-pod-autoscaler
+  verticalAutoscaler:
+    ## If true a VPA object will be created for the controller (either StatefulSet or Deployemnt, based on above configs)
+    enabled: false
+    # updateMode: "Auto"
+    # containerPolicies:
+    # - containerName: 'prometheus-server'
+
   ## Security context to be added to server pods
   ##
   securityContext:
```

## 9.7.1 

**Release date:** 2019-12-20

![AppVersion: 2.13.1](https://img.shields.io/static/v1?label=AppVersion&message=2.13.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Ability to scope alertmanager to a specific instance using probe annotation (#18039) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index c17eabe..36fdfea 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -215,6 +215,11 @@ alertmanager:
   ## Annotations to be added to alertmanager pods
   ##
   podAnnotations: {}
+    ## Tell prometheus to use a specific set of alertmanager pods
+    ## instead of all alertmanager pods found in the same namespace
+    ## Useful if you deploy multiple releases within the same namespace
+    ##
+    ## prometheus.io/probe: alertmanager-teamA
 
   ## Specify if a Pod Security Policy for node-exporter must be created
   ## Ref: https://kubernetes.io/docs/concepts/policy/pod-security-policy/
```

## 9.7.0 

**Release date:** 2019-12-18

![AppVersion: 2.13.1](https://img.shields.io/static/v1?label=AppVersion&message=2.13.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] Add support for volumeBindingMode (#18242) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 47ade67..c17eabe 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -200,6 +200,13 @@ alertmanager:
     ##
     # storageClass: "-"
 
+    ## alertmanager data Persistent Volume Binding Mode
+    ## If defined, volumeBindingMode: <volumeBindingMode>
+    ## If undefined (the default) or set to null, no volumeBindingMode spec is
+    ##   set, choosing the default mode.
+    ##
+    # volumeBindingMode: ""
+
     ## Subdirectory of alertmanager data Persistent Volume to mount
     ## Useful if the volume's root directory is not empty
     ##
@@ -783,6 +790,13 @@ server:
     ##
     # storageClass: "-"
 
+    ## Prometheus server data Persistent Volume Binding Mode
+    ## If defined, volumeBindingMode: <volumeBindingMode>
+    ## If undefined (the default) or set to null, no volumeBindingMode spec is
+    ##   set, choosing the default mode.
+    ##
+    # volumeBindingMode: ""
+
     ## Subdirectory of Prometheus server data Persistent Volume to mount
     ## Useful if the volume's root directory is not empty
     ##
@@ -1072,7 +1086,7 @@ pushgateway:
     ##
     size: 2Gi
 
-    ## alertmanager data Persistent Volume Storage Class
+    ## pushgateway data Persistent Volume Storage Class
     ## If defined, storageClassName: <storageClass>
     ## If set to "-", storageClassName: "", which disables dynamic provisioning
     ## If undefined (the default) or set to null, no storageClassName spec is
@@ -1081,7 +1095,14 @@ pushgateway:
     ##
     # storageClass: "-"
 
-    ## Subdirectory of alertmanager data Persistent Volume to mount
+    ## pushgateway data Persistent Volume Binding Mode
+    ## If defined, volumeBindingMode: <volumeBindingMode>
+    ## If undefined (the default) or set to null, no volumeBindingMode spec is
+    ##   set, choosing the default mode.
+    ##
+    # volumeBindingMode: ""
+
+    ## Subdirectory of pushgateway data Persistent Volume to mount
     ## Useful if the volume's root directory is not empty
     ##
     subPath: ""
```

## 9.6.1 

**Release date:** 2019-12-19

![AppVersion: 2.13.1](https://img.shields.io/static/v1?label=AppVersion&message=2.13.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] Update readiness and liveness probe path for prometheus pushgateway (#19554) 

### Default value changes

```diff
# No changes in this release
```

## 9.6.0 

**Release date:** 2019-12-19

![AppVersion: 2.13.1](https://img.shields.io/static/v1?label=AppVersion&message=2.13.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] Add PodDisruptionBudget support (#19670) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 51bcb7b..47ade67 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -149,6 +149,13 @@ alertmanager:
   ##
   affinity: {}
 
+  ## PodDisruptionBudget settings
+  ## ref: https://kubernetes.io/docs/concepts/workloads/pods/disruptions/
+  ##
+  podDisruptionBudget:
+    enabled: false
+    maxUnavailable: 1
+
   ## Use an alternate scheduler, e.g. "stork".
   ## ref: https://kubernetes.io/docs/tasks/administer-cluster/configure-multiple-schedulers/
   ##
@@ -379,6 +386,13 @@ kubeStateMetrics:
 
   replicaCount: 1
 
+  ## PodDisruptionBudget settings
+  ## ref: https://kubernetes.io/docs/concepts/workloads/pods/disruptions/
+  ##
+  podDisruptionBudget:
+    enabled: false
+    maxUnavailable: 1
+
   ## kube-state-metrics resource requests and limits
   ## Ref: http://kubernetes.io/docs/user-guide/compute-resources/
   ##
@@ -504,6 +518,13 @@ nodeExporter:
   pod:
     labels: {}
 
+  ## PodDisruptionBudget settings
+  ## ref: https://kubernetes.io/docs/concepts/workloads/pods/disruptions/
+  ##
+  podDisruptionBudget:
+    enabled: false
+    maxUnavailable: 1
+
   ## node-exporter resource limits & requests
   ## Ref: https://kubernetes.io/docs/user-guide/compute-resources/
   ##
@@ -711,6 +732,13 @@ server:
   ##
   affinity: {}
 
+  ## PodDisruptionBudget settings
+  ## ref: https://kubernetes.io/docs/concepts/workloads/pods/disruptions/
+  ##
+  podDisruptionBudget:
+    enabled: false
+    maxUnavailable: 1
+
   ## Use an alternate scheduler, e.g. "stork".
   ## ref: https://kubernetes.io/docs/tasks/administer-cluster/configure-multiple-schedulers/
   ##
@@ -970,6 +998,13 @@ pushgateway:
 
   replicaCount: 1
 
+  ## PodDisruptionBudget settings
+  ## ref: https://kubernetes.io/docs/concepts/workloads/pods/disruptions/
+  ##
+  podDisruptionBudget:
+    enabled: false
+    maxUnavailable: 1
+
   ## pushgateway resource requests and limits
   ## Ref: http://kubernetes.io/docs/user-guide/compute-resources/
   ##
```

## 9.5.6 

**Release date:** 2019-12-18

![AppVersion: 2.13.1](https://img.shields.io/static/v1?label=AppVersion&message=2.13.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Add optional service replica selector (#19398) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 7fba509..51bcb7b 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -858,6 +858,13 @@ server:
     sessionAffinity: None
     type: ClusterIP
 
+    ## If using a statefulSet (statefulSet.enabled=true), configure the
+    ## service to connect to a specific replica to have a consistent view
+    ## of the data.
+    statefulsetReplica:
+      enabled: false
+      replica: 0
+
   ## Prometheus server pod termination grace period
   ##
   terminationGracePeriodSeconds: 300
```

## 9.5.5 

**Release date:** 2019-12-18

![AppVersion: 2.13.1](https://img.shields.io/static/v1?label=AppVersion&message=2.13.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* additional control added (#19496) 

### Default value changes

```diff
# No changes in this release
```

## 9.5.4 

**Release date:** 2019-12-18

![AppVersion: 2.13.1](https://img.shields.io/static/v1?label=AppVersion&message=2.13.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] add failure and success threshold to probes (#19058) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index e1388af..7fba509 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -816,8 +816,12 @@ server:
   ##
   readinessProbeInitialDelay: 30
   readinessProbeTimeout: 30
+  readinessProbeFailureThreshold: 3
+  readinessProbeSuccessThreshold: 1
   livenessProbeInitialDelay: 30
   livenessProbeTimeout: 30
+  livenessProbeFailureThreshold: 3
+  livenessProbeSuccessThreshold: 1
 
   ## Prometheus server resource requests and limits
   ## Ref: http://kubernetes.io/docs/user-guide/compute-resources/
```

## 9.5.3 

**Release date:** 2019-12-17

![AppVersion: 2.13.1](https://img.shields.io/static/v1?label=AppVersion&message=2.13.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] Add option to configure AlertManagers for Prometheus server (#19557) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index a310d15..e1388af 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -772,6 +772,10 @@ server:
   ##
   podLabels: {}
 
+  ## Prometheus AlertManager configuration
+  ##
+  alertmanagers: []
+
   ## Specify if a Pod Security Policy for node-exporter must be created
   ## Ref: https://kubernetes.io/docs/concepts/policy/pod-security-policy/
   ##
```

## 9.5.2 

**Release date:** 2019-12-09

![AppVersion: 2.13.1](https://img.shields.io/static/v1?label=AppVersion&message=2.13.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Improve server file names to reflect official docs (#19472) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 31b9946..a310d15 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -1062,7 +1062,7 @@ serverFiles:
 
   ## Alerts configuration
   ## Ref: https://prometheus.io/docs/prometheus/latest/configuration/alerting_rules/
-  alerts: {}
+  alerting_rules.yml: {}
   # groups:
   #   - name: Instances
   #     rules:
@@ -1074,11 +1074,20 @@ serverFiles:
   #         annotations:
   #           description: '{{ $labels.instance }} of job {{ $labels.job }} has been down for more than 5 minutes.'
   #           summary: 'Instance {{ $labels.instance }} down'
+  ## DEPRECATED DEFAULT VALUE, unless explicitly naming your files, please use alerting_rules.yml
+  alerts: {}
 
+  ## Records configuration
+  ## Ref: https://prometheus.io/docs/prometheus/latest/configuration/recording_rules/
+  recording_rules.yml: {}
+  ## DEPRECATED DEFAULT VALUE, unless explicitly naming your files, please use recording_rules.yml
   rules: {}
 
   prometheus.yml:
     rule_files:
+      - /etc/config/recording_rules.yml
+      - /etc/config/alerting_rules.yml
+    ## Below two files are DEPRECATED will be removed from this default values file
       - /etc/config/rules
       - /etc/config/alerts
 
```

## 9.5.1 

**Release date:** 2019-12-07

![AppVersion: 2.13.1](https://img.shields.io/static/v1?label=AppVersion&message=2.13.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* stable/prometheus: Add service session affinity (#18637) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 6be5b5e..31b9946 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -277,6 +277,7 @@ alertmanager:
     loadBalancerSourceRanges: []
     servicePort: 80
     # nodePort: 30000
+    sessionAffinity: None
     type: ClusterIP
 
 ## Monitors ConfigMap changes and POSTs to a URL
@@ -846,6 +847,7 @@ server:
     loadBalancerIP: ""
     loadBalancerSourceRanges: []
     servicePort: 80
+    sessionAffinity: None
     type: ClusterIP
 
   ## Prometheus server pod termination grace period
```

## 9.5.0 

**Release date:** 2019-12-08

![AppVersion: 2.13.1](https://img.shields.io/static/v1?label=AppVersion&message=2.13.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* added deployment strategy for pushgateway (#17107) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index d74187f..6be5b5e 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -986,6 +986,10 @@ pushgateway:
     servicePort: 9091
     type: ClusterIP
 
+  ## pushgateway Deployment Strategy type
+  # strategy:
+  #   type: Recreate
+
   persistentVolume:
     ## If true, pushgateway will create/use a Persistent Volume Claim
     ## If false, use emptyDir
```

## 9.4.1 

**Release date:** 2019-12-07

![AppVersion: 2.13.1](https://img.shields.io/static/v1?label=AppVersion&message=2.13.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* fix #16690 - stable/prometheus - fails with empty clusterrole ruleset (#16692) 

### Default value changes

```diff
# No changes in this release
```

## 9.4.0 

**Release date:** 2019-12-06

![AppVersion: 2.13.1](https://img.shields.io/static/v1?label=AppVersion&message=2.13.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometehus] Support ingress.extraPaths (#17979) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index a380b79..d74187f 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -112,6 +112,13 @@ alertmanager:
     #   - alertmanager.domain.com
     #   - domain.com/alertmanager
 
+    ## Extra paths to prepend to every host configuration. This is useful when working with annotation based services.
+    extraPaths: []
+    # - path: /*
+    #   backend:
+    #     serviceName: ssl-redirect
+    #     servicePort: use-annotation
+
     ## alertmanager Ingress TLS configuration
     ## Secrets must be manually created in the namespace
     ##
@@ -666,6 +673,13 @@ server:
     #   - prometheus.domain.com
     #   - domain.com/prometheus
 
+    ## Extra paths to prepend to every host configuration. This is useful when working with annotation based services.
+    extraPaths: []
+    # - path: /*
+    #   backend:
+    #     serviceName: ssl-redirect
+    #     servicePort: use-annotation
+
     ## Prometheus server Ingress TLS configuration
     ## Secrets must be manually created in the namespace
     ##
@@ -890,6 +904,13 @@ pushgateway:
     #   - pushgateway.domain.com
     #   - domain.com/pushgateway
 
+    ## Extra paths to prepend to every host configuration. This is useful when working with annotation based services.
+    extraPaths: []
+    # - path: /*
+    #   backend:
+    #     serviceName: ssl-redirect
+    #     servicePort: use-annotation
+
     ## pushgateway Ingress TLS configuration
     ## Secrets must be manually created in the namespace
     ##
```

## 9.3.3 

**Release date:** 2019-12-06

![AppVersion: 2.13.1](https://img.shields.io/static/v1?label=AppVersion&message=2.13.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] semver compatible chart versioning breaks chart label. (#18411) 
* [stable/prometheus] Add missing `env` for Server Statefulset (#19019) 
* [stable/prometheus] allow non-yaml objects for configmap values (#18814) 

### Default value changes

```diff
# No changes in this release
```

## 9.3.2 

**Release date:** 2019-12-02

![AppVersion: 2.13.1](https://img.shields.io/static/v1?label=AppVersion&message=2.13.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] Add myself to OWNERS (#18737) 

### Default value changes

```diff
# No changes in this release
```

## 9.3.1 

**Release date:** 2019-11-05

![AppVersion: 2.13.1](https://img.shields.io/static/v1?label=AppVersion&message=2.13.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] allow multiple alert files (#18112) 

### Default value changes

```diff
# No changes in this release
```

## 9.3.0 

**Release date:** 2019-11-05

![AppVersion: 2.13.1](https://img.shields.io/static/v1?label=AppVersion&message=2.13.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] migrate API versions from deprecated, removed versions (#17268) 

### Default value changes

```diff
# No changes in this release
```

## 9.2.0 

**Release date:** 2019-10-21

![AppVersion: 2.13.1](https://img.shields.io/static/v1?label=AppVersion&message=2.13.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Add initContainers support for prometheus server (#17218) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index bbaaed2..a380b79 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -543,7 +543,7 @@ server:
   ##
   image:
     repository: prom/prometheus
-    tag: v2.11.1
+    tag: v2.13.1
     pullPolicy: IfNotPresent
 
   ## prometheus server priorityClassName
@@ -601,6 +601,10 @@ server:
   ##
   extraArgs: {}
 
+  ## Additional InitContainers to initialize the pod
+  ##
+  extraInitContainers: []
+
   ## Additional Prometheus server Volume mounts
   ##
   extraVolumeMounts: []
```

## 9.1.3 

**Release date:** 2019-10-21

![AppVersion: 2.11.1](https://img.shields.io/static/v1?label=AppVersion&message=2.11.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] Add customized readiness/liveness probe config (#17969) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index b9af5a9..bbaaed2 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -788,6 +788,14 @@ server:
       labels: {}
       servicePort: 80
 
+  ## Prometheus server readiness and liveness probe initial delay and timeout
+  ## Ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/
+  ##
+  readinessProbeInitialDelay: 30
+  readinessProbeTimeout: 30
+  livenessProbeInitialDelay: 30
+  livenessProbeTimeout: 30
+
   ## Prometheus server resource requests and limits
   ## Ref: http://kubernetes.io/docs/user-guide/compute-resources/
   ##
```

## 9.1.2 

**Release date:** 2019-10-06

![AppVersion: 2.11.1](https://img.shields.io/static/v1?label=AppVersion&message=2.11.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* add the option to add alert_relabel_configs in prometheus.yml (#15383) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 98caf84..b9af5a9 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -1317,6 +1317,14 @@ extraScrapeConfigs:
   #     - target_label: __address__
   #       replacement: prometheus-blackbox-exporter:9115
 
+# Adds option to add alert_relabel_configs to avoid duplicate alerts in alertmanager
+# useful in H/A prometheus with different external labels but the same alerts
+alertRelabelConfigs:
+  # alert_relabel_configs:
+  # - source_labels: [dc]
+  #   regex: (.+)\d+
+  #   target_label: dc
+
 networkPolicy:
   ## Enable creation of NetworkPolicy resources.
   ##
```

## 9.1.1 

**Release date:** 2019-09-14

![AppVersion: 2.11.1](https://img.shields.io/static/v1?label=AppVersion&message=2.11.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] Added /nodes/metrics permission to prometheus server (#16394) 

### Default value changes

```diff
# No changes in this release
```

## 9.1.0 

**Release date:** 2019-08-17

![AppVersion: 2.11.1](https://img.shields.io/static/v1?label=AppVersion&message=2.11.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Update psps to support all pods (#15055) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 71725d5..98caf84 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -1,6 +1,9 @@
 rbac:
   create: true
 
+podSecurityPolicy:
+  enabled: false
+
 imagePullSecrets:
 # - name: "image-pull-secret"
 
@@ -192,6 +195,20 @@ alertmanager:
   ##
   podAnnotations: {}
 
+  ## Specify if a Pod Security Policy for node-exporter must be created
+  ## Ref: https://kubernetes.io/docs/concepts/policy/pod-security-policy/
+  ##
+  podSecurityPolicy:
+    annotations: {}
+      ## Specify pod annotations
+      ## Ref: https://kubernetes.io/docs/concepts/policy/pod-security-policy/#apparmor
+      ## Ref: https://kubernetes.io/docs/concepts/policy/pod-security-policy/#seccomp
+      ## Ref: https://kubernetes.io/docs/concepts/policy/pod-security-policy/#sysctl
+      ##
+      # seccomp.security.alpha.kubernetes.io/allowedProfileNames: '*'
+      # seccomp.security.alpha.kubernetes.io/defaultProfileName: 'docker/default'
+      # apparmor.security.beta.kubernetes.io/defaultProfileName: 'runtime/default'
+
   ## Use a StatefulSet if replicaCount needs to be greater than 1 (see below)
   ##
   replicaCount: 1
@@ -335,6 +352,20 @@ kubeStateMetrics:
   ##
   podAnnotations: {}
 
+  ## Specify if a Pod Security Policy for node-exporter must be created
+  ## Ref: https://kubernetes.io/docs/concepts/policy/pod-security-policy/
+  ##
+  podSecurityPolicy:
+    annotations: {}
+      ## Specify pod annotations
+      ## Ref: https://kubernetes.io/docs/concepts/policy/pod-security-policy/#apparmor
+      ## Ref: https://kubernetes.io/docs/concepts/policy/pod-security-policy/#seccomp
+      ## Ref: https://kubernetes.io/docs/concepts/policy/pod-security-policy/#sysctl
+      ##
+      # seccomp.security.alpha.kubernetes.io/allowedProfileNames: '*'
+      # seccomp.security.alpha.kubernetes.io/defaultProfileName: 'docker/default'
+      # apparmor.security.beta.kubernetes.io/defaultProfileName: 'runtime/default'
+
   pod:
     labels: {}
 
@@ -404,7 +435,6 @@ nodeExporter:
   ## Ref: https://kubernetes.io/docs/concepts/policy/pod-security-policy/
   ##
   podSecurityPolicy:
-    enabled: False
     annotations: {}
       ## Specify pod annotations
       ## Ref: https://kubernetes.io/docs/concepts/policy/pod-security-policy/#apparmor
@@ -723,6 +753,20 @@ server:
   ##
   podLabels: {}
 
+  ## Specify if a Pod Security Policy for node-exporter must be created
+  ## Ref: https://kubernetes.io/docs/concepts/policy/pod-security-policy/
+  ##
+  podSecurityPolicy:
+    annotations: {}
+      ## Specify pod annotations
+      ## Ref: https://kubernetes.io/docs/concepts/policy/pod-security-policy/#apparmor
+      ## Ref: https://kubernetes.io/docs/concepts/policy/pod-security-policy/#seccomp
+      ## Ref: https://kubernetes.io/docs/concepts/policy/pod-security-policy/#sysctl
+      ##
+      # seccomp.security.alpha.kubernetes.io/allowedProfileNames: '*'
+      # seccomp.security.alpha.kubernetes.io/defaultProfileName: 'docker/default'
+      # apparmor.security.beta.kubernetes.io/defaultProfileName: 'runtime/default'
+
   ## Use a StatefulSet if replicaCount needs to be greater than 1 (see below)
   ##
   replicaCount: 1
@@ -860,6 +904,20 @@ pushgateway:
   ##
   podAnnotations: {}
 
+  ## Specify if a Pod Security Policy for node-exporter must be created
+  ## Ref: https://kubernetes.io/docs/concepts/policy/pod-security-policy/
+  ##
+  podSecurityPolicy:
+    annotations: {}
+      ## Specify pod annotations
+      ## Ref: https://kubernetes.io/docs/concepts/policy/pod-security-policy/#apparmor
+      ## Ref: https://kubernetes.io/docs/concepts/policy/pod-security-policy/#seccomp
+      ## Ref: https://kubernetes.io/docs/concepts/policy/pod-security-policy/#sysctl
+      ##
+      # seccomp.security.alpha.kubernetes.io/allowedProfileNames: '*'
+      # seccomp.security.alpha.kubernetes.io/defaultProfileName: 'docker/default'
+      # apparmor.security.beta.kubernetes.io/defaultProfileName: 'runtime/default'
+
   replicaCount: 1
 
   ## pushgateway resource requests and limits
```

## 9.0.0 

**Release date:** 2019-08-16

![AppVersion: 2.11.1](https://img.shields.io/static/v1?label=AppVersion&message=2.11.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus]: Support disabling the Prometheus server (#13423) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index f3ed3ae..71725d5 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -505,6 +505,7 @@ nodeExporter:
 server:
   ## Prometheus server container name
   ##
+  enabled: true
   name: server
   sidecarContainers:
 
```

## 8.15.1 

**Release date:** 2019-08-13

![AppVersion: 2.11.1](https://img.shields.io/static/v1?label=AppVersion&message=2.11.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] Add options to change volume mount propagation (same as #11194 for prometheus-node-exporter) (#13922) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index e21a25e..f3ed3ae 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -435,6 +435,7 @@ nodeExporter:
     #   mountPath: /srv/txt_collector
     #   hostPath: /var/lib/node-exporter
     #   readOnly: true
+    #   mountPropagation: HostToContainer
 
   extraConfigmapMounts: []
     # - name: certs-configmap
```

## 8.15.0 

**Release date:** 2019-07-29

![AppVersion: 2.11.1](https://img.shields.io/static/v1?label=AppVersion&message=2.11.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* stable/prometheus: run as nobody (#12175) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 4aed3b3..e21a25e 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -229,7 +229,11 @@ alertmanager:
 
   ## Security context to be added to alertmanager pods
   ##
-  securityContext: {}
+  securityContext:
+    runAsUser: 65534
+    runAsNonRoot: true
+    runAsGroup: 65534
+    fsGroup: 65534
 
   service:
     annotations: {}
@@ -289,28 +293,6 @@ configmapReload:
   ##
   resources: {}
 
-initChownData:
-  ## If false, data ownership will not be reset at startup
-  ## This allows the prometheus-server to be run with an arbitrary user
-  ##
-  enabled: true
-
-  ## initChownData container name
-  ##
-  name: init-chown-data
-
-  ## initChownData container image
-  ##
-  image:
-    repository: busybox
-    tag: latest
-    pullPolicy: IfNotPresent
-
-  ## initChownData resource requests and limits
-  ## Ref: http://kubernetes.io/docs/user-guide/compute-resources/
-  ##
-  resources: {}
-
 kubeStateMetrics:
   ## If false, kube-state-metrics will not be installed
   ##
@@ -371,7 +353,9 @@ kubeStateMetrics:
 
   ## Security context to be added to kube-state-metrics pods
   ##
-  securityContext: {}
+  securityContext:
+    runAsUser: 65534
+    runAsNonRoot: true
 
   service:
     annotations:
@@ -771,7 +755,11 @@ server:
 
   ## Security context to be added to server pods
   ##
-  securityContext: {}
+  securityContext:
+    runAsUser: 65534
+    runAsNonRoot: true
+    runAsGroup: 65534
+    fsGroup: 65534
 
   service:
     annotations: {}
@@ -885,7 +873,9 @@ pushgateway:
 
   ## Security context to be added to push-gateway pods
   ##
-  securityContext: {}
+  securityContext:
+    runAsUser: 65534
+    runAsNonRoot: true
 
   service:
     annotations:
```

## 8.14.3 

**Release date:** 2019-07-11

![AppVersion: 2.11.1](https://img.shields.io/static/v1?label=AppVersion&message=2.11.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] Upgrade alertmanager (#15408) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 14ad6b4..4aed3b3 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -36,7 +36,7 @@ alertmanager:
   ##
   image:
     repository: prom/alertmanager
-    tag: v0.17.0
+    tag: v0.18.0
     pullPolicy: IfNotPresent
 
   ## alertmanager priorityClassName
```

## 8.14.2 

**Release date:** 2019-07-10

![AppVersion: 2.11.1](https://img.shields.io/static/v1?label=AppVersion&message=2.11.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] Upgrade to 2.11.1 (#15402) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index b11cb1d..14ad6b4 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -527,7 +527,7 @@ server:
   ##
   image:
     repository: prom/prometheus
-    tag: v2.10.0
+    tag: v2.11.1
     pullPolicy: IfNotPresent
 
   ## prometheus server priorityClassName
```

## 8.14.1 

**Release date:** 2019-07-08

![AppVersion: 2.11.1](https://img.shields.io/static/v1?label=AppVersion&message=2.11.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] Add custom pod labels for prom server (#14819) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index f4d7e11..b11cb1d 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -733,6 +733,10 @@ server:
   podAnnotations: {}
     # iam.amazonaws.com/role: prometheus
 
+  ## Labels to be added to Prometheus server pods
+  ##
+  podLabels: {}
+
   ## Use a StatefulSet if replicaCount needs to be greater than 1 (see below)
   ##
   replicaCount: 1
@@ -744,6 +748,7 @@ server:
     enabled: false
 
     annotations: {}
+    labels: {}
     podManagementPolicy: OrderedReady
 
     ## Alertmanager headless service to use for the statefulset
```

## 8.14.0 

**Release date:** 2019-06-27

![AppVersion: 2.11.1](https://img.shields.io/static/v1?label=AppVersion&message=2.11.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] Updates to support an alternate scheduler (#14364) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 14d7562..f4d7e11 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -796,6 +796,11 @@ pushgateway:
   ##
   enabled: true
 
+  ## Use an alternate scheduler, e.g. "stork".
+  ## ref: https://kubernetes.io/docs/tasks/administer-cluster/configure-multiple-schedulers/
+  ##
+  # schedulerName:
+
   ## pushgateway container name
   ##
   name: pushgateway
```

## 8.13.2 

**Release date:** 2019-06-28

![AppVersion: 2.11.1](https://img.shields.io/static/v1?label=AppVersion&message=2.11.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] add alertmanager.extraSecretMounts (#12888) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 3fd241c..14d7562 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -61,6 +61,15 @@ alertmanager:
   ##
   extraEnv: {}
 
+  ## Additional alertmanager Secret mounts
+  # Defines additional mounts with secrets. Secrets must be manually created in the namespace.
+  extraSecretMounts: []
+    # - name: secret-files
+    #   mountPath: /etc/secrets
+    #   subPath: ""
+    #   secretName: alertmanager-secret-files
+    #   readOnly: true
+
   ## ConfigMap override where fullname is {{.Release.Name}}-{{.Values.alertmanager.configMapOverrideName}}
   ## Defining configMapOverrideName will cause templates/alertmanager-configmap.yaml
   ## to NOT generate a ConfigMap resource
```

## 8.13.1 

**Release date:** 2019-06-27

![AppVersion: 2.11.1](https://img.shields.io/static/v1?label=AppVersion&message=2.11.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] Fixed typo in node-exporter-podsecuritypolicy.yaml (#14154) 

### Default value changes

```diff
# No changes in this release
```

## 8.13.0 

**Release date:** 2019-06-27

![AppVersion: 2.11.0](https://img.shields.io/static/v1?label=AppVersion&message=2.11.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] Upgraded pushgateway to v0.8.0. (#14304) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 6f068f8..3fd241c 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -795,7 +795,7 @@ pushgateway:
   ##
   image:
     repository: prom/pushgateway
-    tag: v0.6.0
+    tag: v0.8.0
     pullPolicy: IfNotPresent
 
   ## pushgateway priorityClassName
```

## 8.12.0 

**Release date:** 2019-06-27

![AppVersion: 2.10.0](https://img.shields.io/static/v1?label=AppVersion&message=2.10.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] Update to Prometheus v2.10.0 (#14294) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 41002dd..6f068f8 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -518,7 +518,7 @@ server:
   ##
   image:
     repository: prom/prometheus
-    tag: v2.9.2
+    tag: v2.10.0
     pullPolicy: IfNotPresent
 
   ## prometheus server priorityClassName
```

## 8.11.6 

**Release date:** 2019-06-27

![AppVersion: 2.9.2](https://img.shields.io/static/v1?label=AppVersion&message=2.9.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] Bump alertmanager to 0.17.0 (#14034) 
* [stable/prometheus] default retention time is 15 days if not specified (#13965) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 094980c..41002dd 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -36,7 +36,7 @@ alertmanager:
   ##
   image:
     repository: prom/alertmanager
-    tag: v0.15.3
+    tag: v0.17.0
     pullPolicy: IfNotPresent
 
   ## alertmanager priorityClassName
@@ -778,9 +778,9 @@ server:
   ##
   terminationGracePeriodSeconds: 300
 
-  ## Prometheus data retention period (i.e 360h)
+  ## Prometheus data retention period (default if not specified is 15 days)
   ##
-  retention: ""
+  retention: "15d"
 
 pushgateway:
   ## If false, pushgateway will not be installed
```

## 8.11.5 

**Release date:** 2019-06-26

![AppVersion: 2.9.2](https://img.shields.io/static/v1?label=AppVersion&message=2.9.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] Ignore OWNERS file in helm package. (#13655) 

### Default value changes

```diff
# No changes in this release
```

## 8.11.4 

**Release date:** 2019-05-22

![AppVersion: 2.9.2](https://img.shields.io/static/v1?label=AppVersion&message=2.9.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] Bump node-exporter to 0.18.0 (#14036) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index a36d94d..094980c 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -404,7 +404,7 @@ nodeExporter:
   ##
   image:
     repository: prom/node-exporter
-    tag: v0.17.0
+    tag: v0.18.0
     pullPolicy: IfNotPresent
 
   ## Specify if a Pod Security Policy for node-exporter must be created
```

## 8.11.3 

**Release date:** 2019-05-20

![AppVersion: 2.9.2](https://img.shields.io/static/v1?label=AppVersion&message=2.9.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus] fix bad indentation for resources on init-container init-chown-data (#13762) 

### Default value changes

```diff
# No changes in this release
```

## 8.11.2 

**Release date:** 2019-05-15

![AppVersion: 2.9.2](https://img.shields.io/static/v1?label=AppVersion&message=2.9.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* add apiVersion (#13817) 

### Default value changes

```diff
# No changes in this release
```

## 8.11.1 

**Release date:** 2019-05-10

![AppVersion: 2.9.2](https://img.shields.io/static/v1?label=AppVersion&message=2.9.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Allow setting an emptyDir sizeLimit for Prometheus (#12281) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index e62503c..a36d94d 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -716,6 +716,9 @@ server:
     ##
     subPath: ""
 
+  emptyDir:
+    sizeLimit: ""
+
   ## Annotations to be added to Prometheus server pods
   ##
   podAnnotations: {}
```

## 8.11.0 

**Release date:** 2019-05-08

![AppVersion: 2.9.2](https://img.shields.io/static/v1?label=AppVersion&message=2.9.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] use kube-state-metrics v1.6.0 (#13623) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 1bac3b8..e62503c 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -315,7 +315,7 @@ kubeStateMetrics:
   ##
   image:
     repository: quay.io/coreos/kube-state-metrics
-    tag: v1.5.0
+    tag: v1.6.0
     pullPolicy: IfNotPresent
 
   ## kube-state-metrics priorityClassName
```

## 8.10.3 

**Release date:** 2019-04-24

![AppVersion: 2.9.2](https://img.shields.io/static/v1?label=AppVersion&message=2.9.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* upgrade to latest release (#13260) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index a733ac7..1bac3b8 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -518,7 +518,7 @@ server:
   ##
   image:
     repository: prom/prometheus
-    tag: v2.9.1
+    tag: v2.9.2
     pullPolicy: IfNotPresent
 
   ## prometheus server priorityClassName
```

## 8.10.2 

**Release date:** 2019-04-22

![AppVersion: 2.9.1](https://img.shields.io/static/v1?label=AppVersion&message=2.9.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] Upgrade to 2.9.1 (#13098) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index ddd39aa..a733ac7 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -518,7 +518,7 @@ server:
   ##
   image:
     repository: prom/prometheus
-    tag: v2.9.0
+    tag: v2.9.1
     pullPolicy: IfNotPresent
 
   ## prometheus server priorityClassName
```

## 8.10.1 

**Release date:** 2019-04-22

![AppVersion: 2.9.0](https://img.shields.io/static/v1?label=AppVersion&message=2.9.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] Remove duplicate affinity field (#12975) 

### Default value changes

```diff
# No changes in this release
```

## 8.10.0 

**Release date:** 2019-04-15

![AppVersion: 2.9.0](https://img.shields.io/static/v1?label=AppVersion&message=2.9.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* upgrade prometheus (#13067) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 1194c6b..ddd39aa 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -518,7 +518,7 @@ server:
   ##
   image:
     repository: prom/prometheus
-    tag: v2.8.0
+    tag: v2.9.0
     pullPolicy: IfNotPresent
 
   ## prometheus server priorityClassName
```

## 8.9.2 

**Release date:** 2019-04-08

![AppVersion: 2.8.0](https://img.shields.io/static/v1?label=AppVersion&message=2.8.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] Update retention time property (#11518) 

### Default value changes

```diff
# No changes in this release
```

## 8.9.1 

**Release date:** 2019-04-01

![AppVersion: 2.8.0](https://img.shields.io/static/v1?label=AppVersion&message=2.8.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Duplicate entries for affinity rules (#12468) 

### Default value changes

```diff
# No changes in this release
```

## 8.9.0 

**Release date:** 2019-03-14

![AppVersion: 2.8.0](https://img.shields.io/static/v1?label=AppVersion&message=2.8.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* upgrade prometheus (#12113) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index d94d1c2..1194c6b 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -518,7 +518,7 @@ server:
   ##
   image:
     repository: prom/prometheus
-    tag: v2.7.2
+    tag: v2.8.0
     pullPolicy: IfNotPresent
 
   ## prometheus server priorityClassName
```

## 8.8.1 

**Release date:** 2019-03-04

![AppVersion: 2.7.2](https://img.shields.io/static/v1?label=AppVersion&message=2.7.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* upgrade prometheus (#11897) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 2d57ddc..d94d1c2 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -518,7 +518,7 @@ server:
   ##
   image:
     repository: prom/prometheus
-    tag: v2.7.1
+    tag: v2.7.2
     pullPolicy: IfNotPresent
 
   ## prometheus server priorityClassName
```

## 8.8.0 

**Release date:** 2019-02-20

![AppVersion: 2.7.1](https://img.shields.io/static/v1?label=AppVersion&message=2.7.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Add ability to control prometheus TSDB locking (#11270) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 3d845e7..2d57ddc 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -555,6 +555,9 @@ server:
   ## series. This is disabled by default.
   enableAdminApi: false
 
+  ## This flag controls BD locking
+  skipTSDBLock: false
+
   ## Path to a configuration file on prometheus server container FS
   configPath: /etc/config/prometheus.yml
 
```

## 8.7.1 

**Release date:** 2019-02-12

![AppVersion: 2.7.1](https://img.shields.io/static/v1?label=AppVersion&message=2.7.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Change AlertManager reload URL from localhost to 127.0.0.1 (#11359) 

### Default value changes

```diff
# No changes in this release
```

## 8.7.0 

**Release date:** 2019-02-08

![AppVersion: 2.7.1](https://img.shields.io/static/v1?label=AppVersion&message=2.7.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] Added PVC to Pushgateway (#11234) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index c57c202..3d845e7 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -798,6 +798,7 @@ pushgateway:
 
   ## Additional pushgateway container arguments
   ##
+  ## for example: persistence.file: /data/pushgateway.data
   extraArgs: {}
 
   ingress:
@@ -877,6 +878,51 @@ pushgateway:
     servicePort: 9091
     type: ClusterIP
 
+  persistentVolume:
+    ## If true, pushgateway will create/use a Persistent Volume Claim
+    ## If false, use emptyDir
+    ##
+    enabled: false
+
+    ## pushgateway data Persistent Volume access modes
+    ## Must match those of existing PV or dynamic provisioner
+    ## Ref: http://kubernetes.io/docs/user-guide/persistent-volumes/
+    ##
+    accessModes:
+      - ReadWriteOnce
+
+    ## pushgateway data Persistent Volume Claim annotations
+    ##
+    annotations: {}
+
+    ## pushgateway data Persistent Volume existing claim name
+    ## Requires pushgateway.persistentVolume.enabled: true
+    ## If defined, PVC must be created manually before volume will be bound
+    existingClaim: ""
+
+    ## pushgateway data Persistent Volume mount root path
+    ##
+    mountPath: /data
+
+    ## pushgateway data Persistent Volume size
+    ##
+    size: 2Gi
+
+    ## alertmanager data Persistent Volume Storage Class
+    ## If defined, storageClassName: <storageClass>
+    ## If set to "-", storageClassName: "", which disables dynamic provisioning
+    ## If undefined (the default) or set to null, no storageClassName spec is
+    ##   set, choosing the default provisioner.  (gp2 on AWS, standard on
+    ##   GKE, AWS & OpenStack)
+    ##
+    # storageClass: "-"
+
+    ## Subdirectory of alertmanager data Persistent Volume to mount
+    ## Useful if the volume's root directory is not empty
+    ##
+    subPath: ""
+
+
 ## alertmanager ConfigMap entries
 ##
 alertmanagerFiles:
```

## 8.6.1 

**Release date:** 2019-01-31

![AppVersion: 2.7.1](https://img.shields.io/static/v1?label=AppVersion&message=2.7.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] upgrade to latest release (#11030) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 15389f5..c57c202 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -518,7 +518,7 @@ server:
   ##
   image:
     repository: prom/prometheus
-    tag: v2.7.0
+    tag: v2.7.1
     pullPolicy: IfNotPresent
 
   ## prometheus server priorityClassName
```

## 8.6.0 

**Release date:** 2019-01-31

![AppVersion: 2.7.0](https://img.shields.io/static/v1?label=AppVersion&message=2.7.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] Upgrade to latest release (#11020) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 32f35f4..15389f5 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -518,7 +518,7 @@ server:
   ##
   image:
     repository: prom/prometheus
-    tag: v2.6.1
+    tag: v2.7.0
     pullPolicy: IfNotPresent
 
   ## prometheus server priorityClassName
```

## 8.5.0 

**Release date:** 2019-01-30

![AppVersion: 2.6.1](https://img.shields.io/static/v1?label=AppVersion&message=2.6.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* enable watching configmaps for prometheus server (#10973) 

### Default value changes

```diff
# No changes in this release
```

## 8.4.9 

**Release date:** 2019-01-30

![AppVersion: 2.6.1](https://img.shields.io/static/v1?label=AppVersion&message=2.6.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] Added configPath parameter (#10976) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 5a53153..32f35f4 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -555,6 +555,9 @@ server:
   ## series. This is disabled by default.
   enableAdminApi: false
 
+  ## Path to a configuration file on prometheus server container FS
+  configPath: /etc/config/prometheus.yml
+
   global:
     ## How frequently to scrape targets by default
     ##
```

## 8.4.8 

**Release date:** 2019-01-29

![AppVersion: 2.6.1](https://img.shields.io/static/v1?label=AppVersion&message=2.6.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Add ability to provide environment variables to the prometheus server (#10968) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 8ec8163..5a53153 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -534,6 +534,23 @@ server:
   ## Maybe same with Ingress host name
   baseURL: ""
 
+  ## Additional server container environment variables
+  ##
+  ## You specify this manually like you would a raw deployment manifest.
+  ## This means you can bind in environment variables from secrets.
+  ##
+  ## e.g. static environment variable:
+  ##  - name: DEMO_GREETING
+  ##    value: "Hello from the environment"
+  ##
+  ## e.g. secret environment variable:
+  ## - name: USERNAME
+  ##   valueFrom:
+  ##     secretKeyRef:
+  ##       name: mysecret
+  ##       key: username
+  env: {}
+
   ## This flag controls access to the administrative HTTP API which includes functionality such as deleting time
   ## series. This is disabled by default.
   enableAdminApi: false
```

## 8.4.7 

**Release date:** 2019-01-25

![AppVersion: 2.6.1](https://img.shields.io/static/v1?label=AppVersion&message=2.6.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus]: allow adding custom containers, volumemounts and volumes to server (#10156) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 8726bc0..8ec8163 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -553,6 +553,14 @@ server:
   ##
   extraArgs: {}
 
+  ## Additional Prometheus server Volume mounts
+  ##
+  extraVolumeMounts: []
+
+  ## Additional Prometheus server Volumes
+  ##
+  extraVolumes: []
+
   ## Additional Prometheus server hostPath mounts
   ##
   extraHostPathMounts: []
```

## 8.4.6 

**Release date:** 2019-01-24

![AppVersion: 2.6.1](https://img.shields.io/static/v1?label=AppVersion&message=2.6.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* sidecar containers functionality to Prometheus helm chart.  (#10871) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 58d6ac6..8726bc0 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -512,6 +512,7 @@ server:
   ## Prometheus server container name
   ##
   name: server
+  sidecarContainers:
 
   ## Prometheus server container image
   ##
```

## 8.4.5 

**Release date:** 2019-01-16

![AppVersion: 2.6.1](https://img.shields.io/static/v1?label=AppVersion&message=2.6.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* upgrade prometheus to 2.6.1 (#10677) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 21f458e..58d6ac6 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -517,7 +517,7 @@ server:
   ##
   image:
     repository: prom/prometheus
-    tag: v2.6.0
+    tag: v2.6.1
     pullPolicy: IfNotPresent
 
   ## prometheus server priorityClassName
```

## 8.4.4 

**Release date:** 2019-01-15

![AppVersion: 2.6.0](https://img.shields.io/static/v1?label=AppVersion&message=2.6.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] updated kube-state-metrics image to v1.5.0 (#10601) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 388fe41..21f458e 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -315,7 +315,7 @@ kubeStateMetrics:
   ##
   image:
     repository: quay.io/coreos/kube-state-metrics
-    tag: v1.4.0
+    tag: v1.5.0
     pullPolicy: IfNotPresent
 
   ## kube-state-metrics priorityClassName
```

## 8.4.3 

**Release date:** 2019-01-10

![AppVersion: 2.6.0](https://img.shields.io/static/v1?label=AppVersion&message=2.6.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] Alertmanager config from Secret (#9317) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 9642c6b..388fe41 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -67,6 +67,17 @@ alertmanager:
   ##
   configMapOverrideName: ""
 
+  ## The name of a secret in the same kubernetes namespace which contains the Alertmanager config
+  ## Defining configFromSecret will cause templates/alertmanager-configmap.yaml
+  ## to NOT generate a ConfigMap resource
+  ##
+  configFromSecret: ""
+
+  ## The configuration file name to be loaded to alertmanager
+  ## Must match the key within configuration loaded from ConfigMap/Secret
+  ##
+  configFileName: alertmanager.yml
+
   ingress:
     ## If true, alertmanager Ingress will be created
     ##
```

## 8.4.2 

**Release date:** 2019-01-08

![AppVersion: 2.6.0](https://img.shields.io/static/v1?label=AppVersion&message=2.6.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Docs: Added the missing `nodeExporter.service.hostPort` configuration parameter with its default port 9100 (#9635) 

### Default value changes

```diff
# No changes in this release
```

## 8.4.1 

**Release date:** 2019-01-04

![AppVersion: 2.6.0](https://img.shields.io/static/v1?label=AppVersion&message=2.6.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* added poddisruptionbudgets view to kubeStateMetrics clusterrole (#10210) 

### Default value changes

```diff
# No changes in this release
```

## 8.4.0 

**Release date:** 2019-01-02

![AppVersion: 2.6.0](https://img.shields.io/static/v1?label=AppVersion&message=2.6.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Makes sense to have rollingUpgrade as default (#10347) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 5d92024..9642c6b 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -418,7 +418,7 @@ nodeExporter:
   ## Custom Update Strategy
   ##
   updateStrategy:
-    type: OnDelete
+    type: RollingUpdate
 
   ## Additional node-exporter container arguments
   ##
```

## 8.3.1 

**Release date:** 2019-01-02

![AppVersion: 2.6.0](https://img.shields.io/static/v1?label=AppVersion&message=2.6.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Trim global section to avoid empty line to be indented (#10204) 

### Default value changes

```diff
# No changes in this release
```

## 8.3.0 

**Release date:** 2018-12-21

![AppVersion: 2.6.0](https://img.shields.io/static/v1?label=AppVersion&message=2.6.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* upgrade prometheus (#10163) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 6714c68..5d92024 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -506,7 +506,7 @@ server:
   ##
   image:
     repository: prom/prometheus
-    tag: v2.5.0
+    tag: v2.6.0
     pullPolicy: IfNotPresent
 
   ## prometheus server priorityClassName
```

## 8.2.0 

**Release date:** 2018-12-19

![AppVersion: 2.5.0](https://img.shields.io/static/v1?label=AppVersion&message=2.5.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Mesh setup for AlertManager StatefulSet (#9874) 

### Default value changes

```diff
# No changes in this release
```

## 8.1.2 

**Release date:** 2018-12-17

![AppVersion: 2.5.0](https://img.shields.io/static/v1?label=AppVersion&message=2.5.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] harmonization of regex reference without braces (#10055) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 1963dcf..6714c68 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -970,7 +970,7 @@ serverFiles:
           - source_labels: [__meta_kubernetes_node_name]
             regex: (.+)
             target_label: __metrics_path__
-            replacement: /api/v1/nodes/${1}/proxy/metrics
+            replacement: /api/v1/nodes/$1/proxy/metrics
 
 
       - job_name: 'kubernetes-nodes-cadvisor'
@@ -1002,7 +1002,7 @@ serverFiles:
         # This configuration will work only on kubelet 1.7.3+
         # As the scrape endpoints for cAdvisor have changed
         # if you are using older version you need to change the replacement to
-        # replacement: /api/v1/nodes/${1}:4194/proxy/metrics
+        # replacement: /api/v1/nodes/$1:4194/proxy/metrics
         # more info here https://github.com/coreos/prometheus-operator/issues/633
         relabel_configs:
           - action: labelmap
@@ -1012,7 +1012,7 @@ serverFiles:
           - source_labels: [__meta_kubernetes_node_name]
             regex: (.+)
             target_label: __metrics_path__
-            replacement: /api/v1/nodes/${1}/proxy/metrics/cadvisor
+            replacement: /api/v1/nodes/$1/proxy/metrics/cadvisor
 
       # Scrape config for service endpoints.
       #
```

## 8.1.1 

**Release date:** 2018-12-15

![AppVersion: 2.5.0](https://img.shields.io/static/v1?label=AppVersion&message=2.5.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Unify labels v2 (#9803) 

### Default value changes

```diff
# No changes in this release
```

## 8.1.0 

**Release date:** 2018-12-03

![AppVersion: 2.5.0](https://img.shields.io/static/v1?label=AppVersion&message=2.5.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* upgrade pushgateway and node-exporter (#9681) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 8f29e5f..1963dcf 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -393,7 +393,7 @@ nodeExporter:
   ##
   image:
     repository: prom/node-exporter
-    tag: v0.16.0
+    tag: v0.17.0
     pullPolicy: IfNotPresent
 
   ## Specify if a Pod Security Policy for node-exporter must be created
@@ -749,7 +749,7 @@ pushgateway:
   ##
   image:
     repository: prom/pushgateway
-    tag: v0.5.2
+    tag: v0.6.0
     pullPolicy: IfNotPresent
 
   ## pushgateway priorityClassName
```

## 8.0.0 

**Release date:** 2018-12-02

![AppVersion: 2.5.0](https://img.shields.io/static/v1?label=AppVersion&message=2.5.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus]: add optional statefulset resource (#7116) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index e0a9e4c..8f29e5f 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -172,8 +172,30 @@ alertmanager:
   ##
   podAnnotations: {}
 
+  ## Use a StatefulSet if replicaCount needs to be greater than 1 (see below)
+  ##
   replicaCount: 1
 
+  statefulSet:
+    ## If true, use a statefulset instead of a deployment for pod management.
+    ## This allows to scale replicas to more than 1 pod
+    ##
+    enabled: false
+
+    podManagementPolicy: OrderedReady
+
+    ## Alertmanager headless service to use for the statefulset
+    ##
+    headless:
+      annotations: {}
+      labels: {}
+
+      ## Enabling peer mesh service end points for enabling the HA alert manager
+      ## Ref: https://github.com/prometheus/alertmanager/blob/master/README.md
+      # enableMeshPeer : true
+
+      servicePort: 80
+
   ## alertmanager resource requests and limits
   ## Ref: http://kubernetes.io/docs/user-guide/compute-resources/
   ##
@@ -656,8 +678,26 @@ server:
   podAnnotations: {}
     # iam.amazonaws.com/role: prometheus
 
+  ## Use a StatefulSet if replicaCount needs to be greater than 1 (see below)
+  ##
   replicaCount: 1
 
+  statefulSet:
+    ## If true, use a statefulset instead of a deployment for pod management.
+    ## This allows to scale replicas to more than 1 pod
+    ##
+    enabled: false
+
+    annotations: {}
+    podManagementPolicy: OrderedReady
+
+    ## Alertmanager headless service to use for the statefulset
+    ##
+    headless:
+      annotations: {}
+      labels: {}
+      servicePort: 80
+
   ## Prometheus server resource requests and limits
   ## Ref: http://kubernetes.io/docs/user-guide/compute-resources/
   ##
```

## 7.4.6 

**Release date:** 2018-12-02

![AppVersion: 2.5.0](https://img.shields.io/static/v1?label=AppVersion&message=2.5.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* add/use unified labels for prometheus components (#9555) 

### Default value changes

```diff
# No changes in this release
```

## 7.4.5 

**Release date:** 2018-11-23

![AppVersion: 2.5.0](https://img.shields.io/static/v1?label=AppVersion&message=2.5.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] Remove unnecessary reclaim policy (#9471) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 07d11de..e0a9e4c 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -651,12 +651,6 @@ server:
     ##
     subPath: ""
 
-    ## reclaimPolicy for the persistent volume.
-    ## See https://kubernetes.io/docs/tasks/administer-cluster/change-pv-reclaim-policy/
-    ## Can be Retain, Recycle, or Delete
-    ##
-    reclaimPolicy: Retain
-
   ## Annotations to be added to Prometheus server pods
   ##
   podAnnotations: {}
```

## 7.4.4 

**Release date:** 2018-11-19

![AppVersion: 2.5.0](https://img.shields.io/static/v1?label=AppVersion&message=2.5.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] Defaulting the PVC's reclaim policy to Retain (#9284) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index e0a9e4c..07d11de 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -651,6 +651,12 @@ server:
     ##
     subPath: ""
 
+    ## reclaimPolicy for the persistent volume.
+    ## See https://kubernetes.io/docs/tasks/administer-cluster/change-pv-reclaim-policy/
+    ## Can be Retain, Recycle, or Delete
+    ##
+    reclaimPolicy: Retain
+
   ## Annotations to be added to Prometheus server pods
   ##
   podAnnotations: {}
```

## 7.4.3 

**Release date:** 2018-11-19

![AppVersion: 2.5.0](https://img.shields.io/static/v1?label=AppVersion&message=2.5.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] add kubernetes_node label (#9321) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index a51a3d0..e0a9e4c 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -1015,6 +1015,9 @@ serverFiles:
           - source_labels: [__meta_kubernetes_service_name]
             action: replace
             target_label: kubernetes_name
+          - source_labels: [__meta_kubernetes_pod_node_name]
+            action: replace
+            target_label: kubernetes_node
 
       - job_name: 'prometheus-pushgateway'
         honor_labels: true
```

## 7.4.2 

**Release date:** 2018-11-19

![AppVersion: 2.5.0](https://img.shields.io/static/v1?label=AppVersion&message=2.5.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] Fixes #9033, some resources are created when a component is disabled (#9274) 

### Default value changes

```diff
# No changes in this release
```

## 7.4.1 

**Release date:** 2018-11-09

![AppVersion: 2.5.0](https://img.shields.io/static/v1?label=AppVersion&message=2.5.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* upgrade alertmanager (#9147) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 6546f26..a51a3d0 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -36,7 +36,7 @@ alertmanager:
   ##
   image:
     repository: prom/alertmanager
-    tag: v0.15.2
+    tag: v0.15.3
     pullPolicy: IfNotPresent
 
   ## alertmanager priorityClassName
```

## 7.4.0 

**Release date:** 2018-11-06

![AppVersion: 2.5.0](https://img.shields.io/static/v1?label=AppVersion&message=2.5.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* upgrade prometheus (#9037) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index e643245..6546f26 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -484,7 +484,7 @@ server:
   ##
   image:
     repository: prom/prometheus
-    tag: v2.4.3
+    tag: v2.5.0
     pullPolicy: IfNotPresent
 
   ## prometheus server priorityClassName
```

## 7.3.6 

**Release date:** 2018-11-05

![AppVersion: 2.4.3](https://img.shields.io/static/v1?label=AppVersion&message=2.4.3&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Enable deployment from secured image registry (#9018) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index dd7d123..e643245 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -1,6 +1,9 @@
 rbac:
   create: true
 
+imagePullSecrets:
+# - name: "image-pull-secret"
+
 ## Define serviceAccount names for components. Defaults to component's fully qualified name.
 ##
 serviceAccounts:
```

## 7.3.5 

**Release date:** 2018-11-02

![AppVersion: 2.4.3](https://img.shields.io/static/v1?label=AppVersion&message=2.4.3&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] Add example alert config (#966) (#8191) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index c4f60b1..dd7d123 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -816,7 +816,22 @@ alertmanagerFiles:
 ## Prometheus server ConfigMap entries
 ##
 serverFiles:
+
+  ## Alerts configuration
+  ## Ref: https://prometheus.io/docs/prometheus/latest/configuration/alerting_rules/
   alerts: {}
+  # groups:
+  #   - name: Instances
+  #     rules:
+  #       - alert: InstanceDown
+  #         expr: up == 0
+  #         for: 5m
+  #         labels:
+  #           severity: page
+  #         annotations:
+  #           description: '{{ $labels.instance }} of job {{ $labels.job }} has been down for more than 5 minutes.'
+  #           summary: 'Instance {{ $labels.instance }} down'
+
   rules: {}
 
   prometheus.yml:
```

## 7.3.4 

**Release date:** 2018-10-25

![AppVersion: 2.4.3](https://img.shields.io/static/v1?label=AppVersion&message=2.4.3&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] Make pod annotations configurable for PodSecurityPolicy (#8746) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 8c85a9c..c4f60b1 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -376,6 +376,15 @@ nodeExporter:
   ##
   podSecurityPolicy:
     enabled: False
+    annotations: {}
+      ## Specify pod annotations
+      ## Ref: https://kubernetes.io/docs/concepts/policy/pod-security-policy/#apparmor
+      ## Ref: https://kubernetes.io/docs/concepts/policy/pod-security-policy/#seccomp
+      ## Ref: https://kubernetes.io/docs/concepts/policy/pod-security-policy/#sysctl
+      ##
+      # seccomp.security.alpha.kubernetes.io/allowedProfileNames: '*'
+      # seccomp.security.alpha.kubernetes.io/defaultProfileName: 'docker/default'
+      # apparmor.security.beta.kubernetes.io/defaultProfileName: 'runtime/default'
 
   ## node-exporter priorityClassName
   ##
```

## 7.3.3 

**Release date:** 2018-10-24

![AppVersion: 2.4.3](https://img.shields.io/static/v1?label=AppVersion&message=2.4.3&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] Add cluster.advertise-address to alertmanager cmd args (#8562) 

### Default value changes

```diff
# No changes in this release
```

## 7.3.2 

**Release date:** 2018-10-23

![AppVersion: 2.4.3](https://img.shields.io/static/v1?label=AppVersion&message=2.4.3&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] Add extra volume directories to configmapreload (#8601) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 8224c98..8c85a9c 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -224,6 +224,10 @@ configmapReload:
   ## Additional configmap-reload container arguments
   ##
   extraArgs: {}
+  ## Additional configmap-reload volume directories
+  ##
+  extraVolumeDirs: []
+
 
   ## Additional configmap-reload mounts
   ##
```

## 7.3.1 

**Release date:** 2018-10-16

![AppVersion: 2.4.3](https://img.shields.io/static/v1?label=AppVersion&message=2.4.3&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] Add option to configure subpath on server volume mounts (#8113) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 3581ce0..8224c98 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -230,6 +230,7 @@ configmapReload:
   extraConfigmapMounts: []
     # - name: prometheus-alerts
     #   mountPath: /etc/alerts.d
+    #   subPath: ""
     #   configMap: prometheus-alerts
     #   readOnly: true
 
@@ -507,12 +508,14 @@ server:
   extraHostPathMounts: []
     # - name: certs-dir
     #   mountPath: /etc/kubernetes/certs
+    #   subPath: ""
     #   hostPath: /etc/kubernetes/certs
     #   readOnly: true
 
   extraConfigmapMounts: []
     # - name: certs-configmap
     #   mountPath: /prometheus
+    #   subPath: ""
     #   configMap: certs-configmap
     #   readOnly: true
 
@@ -521,6 +524,7 @@ server:
   extraSecretMounts: []
     # - name: secret-files
     #   mountPath: /etc/secrets
+    #   subPath: ""
     #   secretName: prom-secret-files
     #   readOnly: true
 
```

## 7.3.0 

**Release date:** 2018-10-09

![AppVersion: 2.4.3](https://img.shields.io/static/v1?label=AppVersion&message=2.4.3&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] add node-exporter podsecuritypolicy support (#8264) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index d0f4c24..3581ce0 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -366,6 +366,12 @@ nodeExporter:
     tag: v0.16.0
     pullPolicy: IfNotPresent
 
+  ## Specify if a Pod Security Policy for node-exporter must be created
+  ## Ref: https://kubernetes.io/docs/concepts/policy/pod-security-policy/
+  ##
+  podSecurityPolicy:
+    enabled: False
+
   ## node-exporter priorityClassName
   ##
   priorityClassName: ""
```

## 7.2.0 

**Release date:** 2018-10-07

![AppVersion: 2.4.3](https://img.shields.io/static/v1?label=AppVersion&message=2.4.3&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* added extra scrape config option (#7984) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 72d5bd6..d0f4c24 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -1053,6 +1053,25 @@ serverFiles:
             action: replace
             target_label: kubernetes_pod_name
 
+# adds additional scrape configs to prometheus.yml
+# must be a string so you have to add a | after extraScrapeConfigs:
+# example adds prometheus-blackbox-exporter scrape config
+extraScrapeConfigs:
+  # - job_name: 'prometheus-blackbox-exporter'
+  #   metrics_path: /probe
+  #   params:
+  #     module: [http_2xx]
+  #   static_configs:
+  #     - targets:
+  #       - https://example.com
+  #   relabel_configs:
+  #     - source_labels: [__address__]
+  #       target_label: __param_target
+  #     - source_labels: [__param_target]
+  #       target_label: instance
+  #     - target_label: __address__
+  #       replacement: prometheus-blackbox-exporter:9115
+
 networkPolicy:
   ## Enable creation of NetworkPolicy resources.
   ##
```

## 7.1.4 

**Release date:** 2018-10-04

![AppVersion: 2.4.3](https://img.shields.io/static/v1?label=AppVersion&message=2.4.3&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* upgrade to prometheus 2.4.3 (#8158) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index e412675..72d5bd6 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -461,7 +461,7 @@ server:
   ##
   image:
     repository: prom/prometheus
-    tag: v2.4.2
+    tag: v2.4.3
     pullPolicy: IfNotPresent
 
   ## prometheus server priorityClassName
```

## 7.1.3 

**Release date:** 2018-09-24

![AppVersion: 2.4.2](https://img.shields.io/static/v1?label=AppVersion&message=2.4.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* add option to configure hostNetwork/hostPID for node-exporter pods (#7643) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 98850c8..e412675 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -347,6 +347,14 @@ nodeExporter:
   ##
   enabled: true
 
+  ## If true, node-exporter pods share the host network namespace
+  ##
+  hostNetwork: true
+
+  ## If true, node-exporter pods share the host PID namespace
+  ##
+  hostPID: true
+
   ## node-exporter container name
   ##
   name: node-exporter
```

## 7.1.2 

**Release date:** 2018-09-21

![AppVersion: 2.4.2](https://img.shields.io/static/v1?label=AppVersion&message=2.4.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* upgrade to prometheus 2.4.2 (#7882) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 473fde8..98850c8 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -453,7 +453,7 @@ server:
   ##
   image:
     repository: prom/prometheus
-    tag: v2.4.0
+    tag: v2.4.2
     pullPolicy: IfNotPresent
 
   ## prometheus server priorityClassName
```

## 7.1.1 

**Release date:** 2018-09-18

![AppVersion: 2.4.0](https://img.shields.io/static/v1?label=AppVersion&message=2.4.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Prometheus: Fix AlertManager discovery if chart name is not the default (#7779) 

### Default value changes

```diff
# No changes in this release
```

## 7.1.0 

**Release date:** 2018-09-11

![AppVersion: 2.4.0](https://img.shields.io/static/v1?label=AppVersion&message=2.4.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* upgrade prometheus to 2.4.0 (#7663) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 62827dc..473fde8 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -453,7 +453,7 @@ server:
   ##
   image:
     repository: prom/prometheus
-    tag: v2.3.2
+    tag: v2.4.0
     pullPolicy: IfNotPresent
 
   ## prometheus server priorityClassName
```

## 7.0.3 

**Release date:** 2018-09-04

![AppVersion: 2.3.2](https://img.shields.io/static/v1?label=AppVersion&message=2.3.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* upgrade kube-state-metrics (#7528) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 9b59d12..62827dc 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -274,7 +274,7 @@ kubeStateMetrics:
   ##
   image:
     repository: quay.io/coreos/kube-state-metrics
-    tag: v1.3.1
+    tag: v1.4.0
     pullPolicy: IfNotPresent
 
   ## kube-state-metrics priorityClassName
```

## 7.0.2 

**Release date:** 2018-08-17

![AppVersion: 2.3.2](https://img.shields.io/static/v1?label=AppVersion&message=2.3.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] upgrade alertmanager (#7193) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 86ff3b7..9b59d12 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -33,7 +33,7 @@ alertmanager:
   ##
   image:
     repository: prom/alertmanager
-    tag: v0.15.1
+    tag: v0.15.2
     pullPolicy: IfNotPresent
 
   ## alertmanager priorityClassName
```

## 7.0.1 

**Release date:** 2018-08-16

![AppVersion: 2.3.2](https://img.shields.io/static/v1?label=AppVersion&message=2.3.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] Add affinity to server & alertmanager (#7215) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 4d3ec71..86ff3b7 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -112,6 +112,10 @@ alertmanager:
   ##
   nodeSelector: {}
 
+  ## Pod affinity
+  ##
+  affinity: {}
+
   ## Use an alternate scheduler, e.g. "stork".
   ## ref: https://kubernetes.io/docs/tasks/administer-cluster/configure-multiple-schedulers/
   ##
@@ -560,6 +564,10 @@ server:
   ##
   nodeSelector: {}
 
+  ## Pod affinity
+  ##
+  affinity: {}
+
   ## Use an alternate scheduler, e.g. "stork".
   ## ref: https://kubernetes.io/docs/tasks/administer-cluster/configure-multiple-schedulers/
   ##
```

## 7.0.0 

**Release date:** 2018-07-27

![AppVersion: 2.3.2](https://img.shields.io/static/v1?label=AppVersion&message=2.3.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] Switch from localhost to 127.0.0.1 (#6788) 

### Default value changes

```diff
# No changes in this release
```

## 6.10.0 

**Release date:** 2018-07-13

![AppVersion: 2.3.2](https://img.shields.io/static/v1?label=AppVersion&message=2.3.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* upgrade to latest prometheus release 2.3.2 and alertmanager 0.15.1 (#6623) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index bf377b9..4d3ec71 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -33,7 +33,7 @@ alertmanager:
   ##
   image:
     repository: prom/alertmanager
-    tag: v0.15.0
+    tag: v0.15.1
     pullPolicy: IfNotPresent
 
   ## alertmanager priorityClassName
@@ -449,7 +449,7 @@ server:
   ##
   image:
     repository: prom/prometheus
-    tag: v2.3.1
+    tag: v2.3.2
     pullPolicy: IfNotPresent
 
   ## prometheus server priorityClassName
```

## 6.9.0 

**Release date:** 2018-07-12

![AppVersion: 2.3.1](https://img.shields.io/static/v1?label=AppVersion&message=2.3.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] support priorityClasses (#6578) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index ca186a4..bf377b9 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -36,6 +36,10 @@ alertmanager:
     tag: v0.15.0
     pullPolicy: IfNotPresent
 
+  ## alertmanager priorityClassName
+  ##
+  priorityClassName: ""
+
   ## Additional alertmanager container arguments
   ##
   extraArgs: {}
@@ -269,6 +273,10 @@ kubeStateMetrics:
     tag: v1.3.1
     pullPolicy: IfNotPresent
 
+  ## kube-state-metrics priorityClassName
+  ##
+  priorityClassName: ""
+
   ## kube-state-metrics container arguments
   ##
   args: {}
@@ -346,6 +354,10 @@ nodeExporter:
     tag: v0.16.0
     pullPolicy: IfNotPresent
 
+  ## node-exporter priorityClassName
+  ##
+  priorityClassName: ""
+
   ## Custom Update Strategy
   ##
   updateStrategy:
@@ -440,6 +452,10 @@ server:
     tag: v2.3.1
     pullPolicy: IfNotPresent
 
+  ## prometheus server priorityClassName
+  ##
+  priorityClassName: ""
+
   ## The URL prefix at which the container can be accessed. Useful in the case the '-web.external-url' includes a slug
   ## so that the various internal URLs are still able to access as they are in the default case.
   ## (Optional)
@@ -654,6 +670,10 @@ pushgateway:
     tag: v0.5.2
     pullPolicy: IfNotPresent
 
+  ## pushgateway priorityClassName
+  ##
+  priorityClassName: ""
+
   ## Additional pushgateway container arguments
   ##
   extraArgs: {}
```

## 6.8.1 

**Release date:** 2018-07-10

![AppVersion: 2.3.1](https://img.shields.io/static/v1?label=AppVersion&message=2.3.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Deployment annotations (#6496) 

### Default value changes

```diff
# No changes in this release
```

## 6.8.0 

**Release date:** 2018-07-01

![AppVersion: 2.3.1](https://img.shields.io/static/v1?label=AppVersion&message=2.3.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* added annotations for node-exporter-daemonset (#6387) 

### Default value changes

```diff
# No changes in this release
```

## 6.7.5 

**Release date:** 2018-06-26

![AppVersion: 2.3.1](https://img.shields.io/static/v1?label=AppVersion&message=2.3.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* upgrade alertmanager, prometheus, configmap-reload and node-exporter (#6276) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 952a222..ca186a4 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -33,7 +33,7 @@ alertmanager:
   ##
   image:
     repository: prom/alertmanager
-    tag: v0.14.0
+    tag: v0.15.0
     pullPolicy: IfNotPresent
 
   ## Additional alertmanager container arguments
@@ -210,7 +210,7 @@ configmapReload:
   ##
   image:
     repository: jimmidyson/configmap-reload
-    tag: v0.1
+    tag: v0.2.2
     pullPolicy: IfNotPresent
 
   ## Additional configmap-reload container arguments
@@ -343,7 +343,7 @@ nodeExporter:
   ##
   image:
     repository: prom/node-exporter
-    tag: v0.15.2
+    tag: v0.16.0
     pullPolicy: IfNotPresent
 
   ## Custom Update Strategy
@@ -437,7 +437,7 @@ server:
   ##
   image:
     repository: prom/prometheus
-    tag: v2.2.1
+    tag: v2.3.1
     pullPolicy: IfNotPresent
 
   ## The URL prefix at which the container can be accessed. Useful in the case the '-web.external-url' includes a slug
@@ -651,7 +651,7 @@ pushgateway:
   ##
   image:
     repository: prom/pushgateway
-    tag: v0.5.1
+    tag: v0.5.2
     pullPolicy: IfNotPresent
 
   ## Additional pushgateway container arguments
```

## 6.7.4 

**Release date:** 2018-06-18

![AppVersion: 2.2.1](https://img.shields.io/static/v1?label=AppVersion&message=2.2.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] add required tillerVersion to Chart.yaml (#3854) 

### Default value changes

```diff
# No changes in this release
```

## 6.7.3 

**Release date:** 2018-06-17

![AppVersion: 2.2.1](https://img.shields.io/static/v1?label=AppVersion&message=2.2.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Re-implement on top of latest master changes (#6036) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 787868a..952a222 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -453,6 +453,17 @@ server:
   ## series. This is disabled by default.
   enableAdminApi: false
 
+  global:
+    ## How frequently to scrape targets by default
+    ##
+    scrape_interval: 1m
+    ## How long until a scrape request times out
+    ##
+    scrape_timeout: 10s
+    ## How frequently to evaluate rules
+    ##
+    evaluation_interval: 1m
+
   ## Additional Prometheus server container arguments
   ##
   extraArgs: {}
@@ -750,17 +761,6 @@ serverFiles:
   rules: {}
 
   prometheus.yml:
-    global:
-      ## How frequently to scrape targets by default
-      ##
-      scrape_interval: 1m
-      ## How long until a scrape request times out
-      ##
-      scrape_timeout: 10s
-      ## How frequently to evaluate rules
-      ##
-      evaluation_interval: 1m
-
     rule_files:
       - /etc/config/rules
       - /etc/config/alerts
```

## 6.7.2 

**Release date:** 2018-06-07

![AppVersion: 2.2.1](https://img.shields.io/static/v1?label=AppVersion&message=2.2.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] Update Pushgateway image tag (#5969) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index ef843bd..787868a 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -640,7 +640,7 @@ pushgateway:
   ##
   image:
     repository: prom/pushgateway
-    tag: v0.4.0
+    tag: v0.5.1
     pullPolicy: IfNotPresent
 
   ## Additional pushgateway container arguments
```

## 6.7.1 

**Release date:** 2018-06-05

![AppVersion: 2.2.1](https://img.shields.io/static/v1?label=AppVersion&message=2.2.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] #5856 Set 'global' to a dictionary (#5857) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 0ee4a00..ef843bd 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -728,7 +728,7 @@ pushgateway:
 ##
 alertmanagerFiles:
   alertmanager.yml:
-    global:
+    global: {}
       # slack_api_url: ''
 
     receivers:
```

## 6.7.0 

**Release date:** 2018-05-17

![AppVersion: 2.2.1](https://img.shields.io/static/v1?label=AppVersion&message=2.2.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] Add optional extra node-exporter pod labels (#5268) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 64a8a09..0ee4a00 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -387,6 +387,11 @@ nodeExporter:
   ##
   podAnnotations: {}
 
+  ## Labels to be added to node-exporter pods
+  ##
+  pod:
+    labels: {}
+
   ## node-exporter resource limits & requests
   ## Ref: https://kubernetes.io/docs/user-guide/compute-resources/
   ##
```

## 6.6.2 

**Release date:** 2018-05-16

![AppVersion: 2.2.1](https://img.shields.io/static/v1?label=AppVersion&message=2.2.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] Fix #5425 (#5591) 

### Default value changes

```diff
# No changes in this release
```

## 6.6.1 

**Release date:** 2018-05-15

![AppVersion: 2.2.1](https://img.shields.io/static/v1?label=AppVersion&message=2.2.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] Ability to enable admin API (#5570) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 94b96fe..64a8a09 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -444,6 +444,10 @@ server:
   ## Maybe same with Ingress host name
   baseURL: ""
 
+  ## This flag controls access to the administrative HTTP API which includes functionality such as deleting time
+  ## series. This is disabled by default.
+  enableAdminApi: false
+
   ## Additional Prometheus server container arguments
   ##
   extraArgs: {}
```

## 6.6.0 

**Release date:** 2018-05-14

![AppVersion: 2.2.1](https://img.shields.io/static/v1?label=AppVersion&message=2.2.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [prometheus] kube-state-metrics v1.3.1 (#4829) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 067b88d..94b96fe 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -265,8 +265,8 @@ kubeStateMetrics:
   ## kube-state-metrics container image
   ##
   image:
-    repository: k8s.gcr.io/kube-state-metrics
-    tag: v1.2.0
+    repository: quay.io/coreos/kube-state-metrics
+    tag: v1.3.1
     pullPolicy: IfNotPresent
 
   ## kube-state-metrics container arguments
```

## 6.5.2 

**Release date:** 2018-05-14

![AppVersion: 2.2.1](https://img.shields.io/static/v1?label=AppVersion&message=2.2.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* stable/prometheus: restore headless svcs, reverting breaking change to the default values (#5529) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 3a3ac4d..067b88d 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -316,7 +316,9 @@ kubeStateMetrics:
       prometheus.io/scrape: "true"
     labels: {}
 
-    clusterIP: ""
+    # Exposed as a headless service:
+    # https://kubernetes.io/docs/concepts/services-networking/service/#headless-services
+    clusterIP: None
 
     ## List of IP addresses at which the kube-state-metrics service is available
     ## Ref: https://kubernetes.io/docs/user-guide/services/#external-ips
@@ -406,7 +408,9 @@ nodeExporter:
       prometheus.io/scrape: "true"
     labels: {}
 
-    clusterIP: ""
+    # Exposed as a headless service:
+    # https://kubernetes.io/docs/concepts/services-networking/service/#headless-services
+    clusterIP: None
 
     ## List of IP addresses at which the node-exporter service is available
     ## Ref: https://kubernetes.io/docs/user-guide/services/#external-ips
```

## 6.5.1 

**Release date:** 2018-05-10

![AppVersion: 2.2.1](https://img.shields.io/static/v1?label=AppVersion&message=2.2.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* issues 5425 (#5430) 

### Default value changes

```diff
# No changes in this release
```

## 6.5.0 

**Release date:** 2018-05-10

![AppVersion: 2.2.1](https://img.shields.io/static/v1?label=AppVersion&message=2.2.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] optional extra Ingress labels (#5186) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 156dc52..3a3ac4d 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -71,6 +71,10 @@ alertmanager:
     #   kubernetes.io/ingress.class: nginx
     #   kubernetes.io/tls-acme: 'true'
 
+    ## alertmanager Ingress additional labels
+    ##
+    extraLabels: {}
+
     ## alertmanager Ingress hostnames with optional path
     ## Must be provided if Ingress is enabled
     ##
@@ -479,6 +483,10 @@ server:
     #   kubernetes.io/ingress.class: nginx
     #   kubernetes.io/tls-acme: 'true'
 
+    ## Prometheus server Ingress additional labels
+    ##
+    extraLabels: {}
+
     ## Prometheus server Ingress hostnames with optional path
     ## Must be provided if Ingress is enabled
     ##
```

## 6.4.1 

**Release date:** 2018-05-09

![AppVersion: 2.2.1](https://img.shields.io/static/v1?label=AppVersion&message=2.2.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Forcing a new line to be created in the configmap yaml generation (#5318) 

### Default value changes

```diff
# No changes in this release
```

## 6.4.0 

**Release date:** 2018-05-09

![AppVersion: 2.2.1](https://img.shields.io/static/v1?label=AppVersion&message=2.2.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Allows alternate scheduler for Prometheus server + alertmanager (#5212) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index fcb8627..156dc52 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -104,6 +104,11 @@ alertmanager:
   ##
   nodeSelector: {}
 
+  ## Use an alternate scheduler, e.g. "stork".
+  ## ref: https://kubernetes.io/docs/tasks/administer-cluster/configure-multiple-schedulers/
+  ##
+  # schedulerName:
+
   persistentVolume:
     ## If true, alertmanager will create/use a Persistent Volume Claim
     ## If false, use emptyDir
@@ -507,6 +512,11 @@ server:
   ##
   nodeSelector: {}
 
+  ## Use an alternate scheduler, e.g. "stork".
+  ## ref: https://kubernetes.io/docs/tasks/administer-cluster/configure-multiple-schedulers/
+  ##
+  # schedulerName:
+
   persistentVolume:
     ## If true, Prometheus server will create/use a Persistent Volume Claim
     ## If false, use emptyDir
```

## 6.3.3 

**Release date:** 2018-05-08

![AppVersion: 2.2.1](https://img.shields.io/static/v1?label=AppVersion&message=2.2.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* fix #3308 (#5217) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index c80387a..fcb8627 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -719,6 +719,17 @@ serverFiles:
   rules: {}
 
   prometheus.yml:
+    global:
+      ## How frequently to scrape targets by default
+      ##
+      scrape_interval: 1m
+      ## How long until a scrape request times out
+      ##
+      scrape_timeout: 10s
+      ## How frequently to evaluate rules
+      ##
+      evaluation_interval: 1m
+
     rule_files:
       - /etc/config/rules
       - /etc/config/alerts
```

## 6.3.2 

**Release date:** 2018-05-08

![AppVersion: 2.2.1](https://img.shields.io/static/v1?label=AppVersion&message=2.2.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] typo fix: optinal->optional (#5422) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 799a056..c80387a 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -474,7 +474,7 @@ server:
     #   kubernetes.io/ingress.class: nginx
     #   kubernetes.io/tls-acme: 'true'
 
-    ## Prometheus server Ingress hostnames with optinal path
+    ## Prometheus server Ingress hostnames with optional path
     ## Must be provided if Ingress is enabled
     ##
     hosts: []
@@ -627,7 +627,7 @@ pushgateway:
     #   kubernetes.io/ingress.class: nginx
     #   kubernetes.io/tls-acme: 'true'
 
-    ## pushgateway Ingress hostnames with optinal path
+    ## pushgateway Ingress hostnames with optional path
     ## Must be provided if Ingress is enabled
     ##
     hosts: []
```

## 6.3.1 

**Release date:** 2018-05-04

![AppVersion: 2.2.1](https://img.shields.io/static/v1?label=AppVersion&message=2.2.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* add myself as owner of prometheus chart (#5195) 

### Default value changes

```diff
# No changes in this release
```

## 6.3.0 

**Release date:** 2018-05-03

![AppVersion: 2.2.1](https://img.shields.io/static/v1?label=AppVersion&message=2.2.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus]: Make the security context customizable (#4422) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 2842f45..799a056 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -165,6 +165,10 @@ alertmanager:
     #   cpu: 10m
     #   memory: 32Mi
 
+  ## Security context to be added to alertmanager pods
+  ##
+  securityContext: {}
+
   service:
     annotations: {}
     labels: {}
@@ -294,6 +298,10 @@ kubeStateMetrics:
     #   cpu: 10m
     #   memory: 16Mi
 
+  ## Security context to be added to kube-state-metrics pods
+  ##
+  securityContext: {}
+
   service:
     annotations:
       prometheus.io/scrape: "true"
@@ -561,6 +569,10 @@ server:
     #   cpu: 500m
     #   memory: 512Mi
 
+  ## Security context to be added to server pods
+  ##
+  securityContext: {}
+
   service:
     annotations: {}
     labels: {}
@@ -661,6 +673,10 @@ pushgateway:
     #   cpu: 10m
     #   memory: 32Mi
 
+  ## Security context to be added to push-gateway pods
+  ##
+  securityContext: {}
+
   service:
     annotations:
       prometheus.io/probe: pushgateway
```

## 6.2.1 

**Release date:** 2018-04-18

![AppVersion: 2.2.1](https://img.shields.io/static/v1?label=AppVersion&message=2.2.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* fix prometheus cluster ip and support pod labels (#4915) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index e1803c0..2842f45 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -278,6 +278,9 @@ kubeStateMetrics:
   ##
   podAnnotations: {}
 
+  pod:
+    labels: {}
+
   replicaCount: 1
 
   ## kube-state-metrics resource requests and limits
@@ -296,7 +299,7 @@ kubeStateMetrics:
       prometheus.io/scrape: "true"
     labels: {}
 
-    clusterIP: None
+    clusterIP: ""
 
     ## List of IP addresses at which the kube-state-metrics service is available
     ## Ref: https://kubernetes.io/docs/user-guide/services/#external-ips
@@ -386,7 +389,7 @@ nodeExporter:
       prometheus.io/scrape: "true"
     labels: {}
 
-    clusterIP: None
+    clusterIP: ""
 
     ## List of IP addresses at which the node-exporter service is available
     ## Ref: https://kubernetes.io/docs/user-guide/services/#external-ips
```

## 6.2.0 

**Release date:** 2018-04-12

![AppVersion: 2.2.1](https://img.shields.io/static/v1?label=AppVersion&message=2.2.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus]: remove unnecessary clusterrolebindings for alert… (#4516) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index c7f0b04..e1803c0 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -1,14 +1,30 @@
 rbac:
   create: true
 
+## Define serviceAccount names for components. Defaults to component's fully qualified name.
+##
+serviceAccounts:
+  alertmanager:
+    create: true
+    name:
+  kubeStateMetrics:
+    create: true
+    name:
+  nodeExporter:
+    create: true
+    name:
+  pushgateway:
+    create: true
+    name:
+  server:
+    create: true
+    name:
+
 alertmanager:
   ## If false, alertmanager will not be installed
   ##
   enabled: true
 
-  # Defines the serviceAccountName to use when `rbac.create=false`
-  serviceAccountName: default
-
   ## alertmanager container name
   ##
   name: alertmanager
@@ -229,9 +245,6 @@ kubeStateMetrics:
   ##
   enabled: true
 
-  # Defines the serviceAccountName to use when `rbac.create=false`
-  serviceAccountName: default
-
   ## kube-state-metrics container name
   ##
   name: kube-state-metrics
@@ -300,9 +313,6 @@ nodeExporter:
   ##
   enabled: true
 
-  # Defines the serviceAccountName to use when `rbac.create=false`
-  serviceAccountName: default
-
   ## node-exporter container name
   ##
   name: node-exporter
@@ -394,9 +404,6 @@ server:
   ##
   name: server
 
-  # Defines the serviceAccountName to use when `rbac.create=false`
-  serviceAccountName: default
-
   ## Prometheus server container image
   ##
   image:
```

## 6.1.2 

**Release date:** 2018-04-09

![AppVersion: 2.2.1](https://img.shields.io/static/v1?label=AppVersion&message=2.2.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* fix typo in prometheus/values.yaml (#4796) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index c377447..c7f0b04 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -55,7 +55,7 @@ alertmanager:
     #   kubernetes.io/ingress.class: nginx
     #   kubernetes.io/tls-acme: 'true'
 
-    ## alertmanager Ingress hostnames with optinal path
+    ## alertmanager Ingress hostnames with optional path
     ## Must be provided if Ingress is enabled
     ##
     hosts: []
```

## 6.1.1 

**Release date:** 2018-04-03

![AppVersion: 2.2.1](https://img.shields.io/static/v1?label=AppVersion&message=2.2.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Fix for version comparison from strings to semver (#4605) 

### Default value changes

```diff
# No changes in this release
```

## 6.1.0 

**Release date:** 2018-04-01

![AppVersion: 2.2.1](https://img.shields.io/static/v1?label=AppVersion&message=2.2.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] Add extraArgs and extraConfigmapMounts for config-reload (#4511) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 19484a2..c377447 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -184,6 +184,19 @@ configmapReload:
     tag: v0.1
     pullPolicy: IfNotPresent
 
+  ## Additional configmap-reload container arguments
+  ##
+  extraArgs: {}
+
+  ## Additional configmap-reload mounts
+  ##
+  extraConfigmapMounts: []
+    # - name: prometheus-alerts
+    #   mountPath: /etc/alerts.d
+    #   configMap: prometheus-alerts
+    #   readOnly: true
+
+
   ## configmap-reload resource requests and limits
   ## Ref: http://kubernetes.io/docs/user-guide/compute-resources/
   ##
```

## 6.0.0 

**Release date:** 2018-03-30

![AppVersion: 2.2](https://img.shields.io/static/v1?label=AppVersion&message=2.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] Configurable name overrides / appVersion (#4397) 

### Default value changes

```diff
# No changes in this release
```

## 5.5.3 

**Release date:** 2018-03-26

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* stable/prometheus: Update prom/alertmanager to 0.14.0 (#4294) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index e307df5..19484a2 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -17,7 +17,7 @@ alertmanager:
   ##
   image:
     repository: prom/alertmanager
-    tag: v0.13.0
+    tag: v0.14.0
     pullPolicy: IfNotPresent
 
   ## Additional alertmanager container arguments
```

## 5.5.2 

**Release date:** 2018-03-27

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* change 'tables'->table" (#4382) 

### Default value changes

```diff
# No changes in this release
```

## 5.5.1 

**Release date:** 2018-03-21

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] Fix some typos (#4368) 

### Default value changes

```diff
# No changes in this release
```

## 5.5.0 

**Release date:** 2018-03-19

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* stable/prometheus: Update default version to 2.2.1 (#4288) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index b431972..e307df5 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -388,7 +388,7 @@ server:
   ##
   image:
     repository: prom/prometheus
-    tag: v2.1.0
+    tag: v2.2.1
     pullPolicy: IfNotPresent
 
   ## The URL prefix at which the container can be accessed. Useful in the case the '-web.external-url' includes a slug
```

## 5.4.4 

**Release date:** 2018-03-19

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] extraSecretMounts support (#4059) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 9abc7d7..b431972 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -418,6 +418,14 @@ server:
     #   configMap: certs-configmap
     #   readOnly: true
 
+  ## Additional Prometheus server Secret mounts
+  # Defines additional mounts with secrets. Secrets must be manually created in the namespace.
+  extraSecretMounts: []
+    # - name: secret-files
+    #   mountPath: /etc/secrets
+    #   secretName: prom-secret-files
+    #   readOnly: true
+
   ## ConfigMap override where fullname is {{.Release.Name}}-{{.Values.server.configMapOverrideName}}
   ## Defining configMapOverrideName will cause templates/server-configmap.yaml
   ## to NOT generate a ConfigMap resource
```

## 5.4.3 

**Release date:** 2018-03-12

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Add owners (#3910) 

### Default value changes

```diff
# No changes in this release
```

## 5.4.2 

**Release date:** 2018-03-09

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Allow to disable data ownership reset (#4035) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 3b9036c..9abc7d7 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -190,6 +190,11 @@ configmapReload:
   resources: {}
 
 initChownData:
+  ## If false, data ownership will not be reset at startup
+  ## This allows the prometheus-server to be run with an arbitrary user
+  ##
+  enabled: true
+
   ## initChownData container name
   ##
   name: init-chown-data
```

## 5.4.1 

**Release date:** 2018-03-06

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] update prometheus.io/port regex (#3961) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index f7e15b5..3b9036c 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -834,7 +834,7 @@ serverFiles:
           - source_labels: [__address__, __meta_kubernetes_service_annotation_prometheus_io_port]
             action: replace
             target_label: __address__
-            regex: (.+)(?::\d+);(\d+)
+            regex: ([^:]+)(?::\d+)?;(\d+)
             replacement: $1:$2
           - action: labelmap
             regex: __meta_kubernetes_service_label_(.+)
@@ -911,8 +911,8 @@ serverFiles:
             regex: (.+)
           - source_labels: [__address__, __meta_kubernetes_pod_annotation_prometheus_io_port]
             action: replace
-            regex: (.+):(?:\d+);(\d+)
-            replacement: ${1}:${2}
+            regex: ([^:]+)(?::\d+)?;(\d+)
+            replacement: $1:$2
             target_label: __address__
           - action: labelmap
             regex: __meta_kubernetes_pod_label_(.+)
```

## 5.4.0 

**Release date:** 2018-02-28

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] Use yaml not literal for alertmanagerFiles (#3777) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 47a404b..f7e15b5 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -644,7 +644,7 @@ pushgateway:
 ## alertmanager ConfigMap entries
 ##
 alertmanagerFiles:
-  alertmanager.yml: |-
+  alertmanager.yml:
     global:
       # slack_api_url: ''
 
```

## 5.3.3 

**Release date:** 2018-02-24

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Change rbac.create default to true and update README (#3813) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index af67851..47a404b 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -1,5 +1,5 @@
 rbac:
-  create: false
+  create: true
 
 alertmanager:
   ## If false, alertmanager will not be installed
```

## 5.3.2 

**Release date:** 2018-02-22

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* if alertmanager.prefixURL is set, feed that to the server alertmanagers config (#3773) 

### Default value changes

```diff
# No changes in this release
```

## 5.3.1 

**Release date:** 2018-02-20

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] Allow set securityContext for node exporter pods (#3734) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 80948d5..af67851 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -348,6 +348,11 @@ nodeExporter:
     #   cpu: 100m
     #   memory: 30Mi
 
+  ## Security context to be added to node-exporter pods
+  ##
+  securityContext: {}
+    # runAsUser: 0
+
   service:
     annotations:
       prometheus.io/scrape: "true"
@@ -570,7 +575,7 @@ pushgateway:
 
     ## pushgateway Ingress annotations
     ##
-    annotations:
+    annotations: {}
     #   kubernetes.io/ingress.class: nginx
     #   kubernetes.io/tls-acme: 'true'
 
```

## 5.3.0 

**Release date:** 2018-02-20

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* feature(prometheus): allow path element in ingress values (#3546) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index ee81b10..80948d5 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -55,11 +55,12 @@ alertmanager:
     #   kubernetes.io/ingress.class: nginx
     #   kubernetes.io/tls-acme: 'true'
 
-    ## alertmanager Ingress hostnames
+    ## alertmanager Ingress hostnames with optinal path
     ## Must be provided if Ingress is enabled
     ##
     hosts: []
     #   - alertmanager.domain.com
+    #   - domain.com/alertmanager
 
     ## alertmanager Ingress TLS configuration
     ## Secrets must be manually created in the namespace
@@ -424,11 +425,12 @@ server:
     #   kubernetes.io/ingress.class: nginx
     #   kubernetes.io/tls-acme: 'true'
 
-    ## Prometheus server Ingress hostnames
+    ## Prometheus server Ingress hostnames with optinal path
     ## Must be provided if Ingress is enabled
     ##
     hosts: []
     #   - prometheus.domain.com
+    #   - domain.com/prometheus
 
     ## Prometheus server Ingress TLS configuration
     ## Secrets must be manually created in the namespace
@@ -572,11 +574,12 @@ pushgateway:
     #   kubernetes.io/ingress.class: nginx
     #   kubernetes.io/tls-acme: 'true'
 
-    ## pushgateway Ingress hostnames
+    ## pushgateway Ingress hostnames with optinal path
     ## Must be provided if Ingress is enabled
     ##
     hosts: []
     #   - pushgateway.domain.com
+    #   - domain.com/pushgateway
 
     ## pushgateway Ingress TLS configuration
     ## Secrets must be manually created in the namespace
```

## 5.2.1 

**Release date:** 2018-02-20

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [Prometheus] Updated cAdvisor endpoint (#3684) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 8589944..ee81b10 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -780,6 +780,11 @@ serverFiles:
         kubernetes_sd_configs:
           - role: node
 
+        # This configuration will work only on kubelet 1.7.3+
+        # As the scrape endpoints for cAdvisor have changed
+        # if you are using older version you need to change the replacement to
+        # replacement: /api/v1/nodes/${1}:4194/proxy/metrics
+        # more info here https://github.com/coreos/prometheus-operator/issues/633
         relabel_configs:
           - action: labelmap
             regex: __meta_kubernetes_node_label_(.+)
@@ -788,7 +793,7 @@ serverFiles:
           - source_labels: [__meta_kubernetes_node_name]
             regex: (.+)
             target_label: __metrics_path__
-            replacement: /api/v1/nodes/${1}:4194/proxy/metrics
+            replacement: /api/v1/nodes/${1}/proxy/metrics/cadvisor
 
       # Scrape config for service endpoints.
       #
```

## 5.1.7 

**Release date:** 2018-02-19

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Mount arbitrary configmaps as files. (#3774) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 2f6d43b..8589944 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -312,6 +312,12 @@ nodeExporter:
     #   hostPath: /var/lib/node-exporter
     #   readOnly: true
 
+  extraConfigmapMounts: []
+    # - name: certs-configmap
+    #   mountPath: /prometheus
+    #   configMap: certs-configmap
+    #   readOnly: true
+
   ## Node tolerations for node-exporter scheduling to nodes with taints
   ## Ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
   ##
@@ -395,6 +401,12 @@ server:
     #   hostPath: /etc/kubernetes/certs
     #   readOnly: true
 
+  extraConfigmapMounts: []
+    # - name: certs-configmap
+    #   mountPath: /prometheus
+    #   configMap: certs-configmap
+    #   readOnly: true
+
   ## ConfigMap override where fullname is {{.Release.Name}}-{{.Values.server.configMapOverrideName}}
   ## Defining configMapOverrideName will cause templates/server-configmap.yaml
   ## to NOT generate a ConfigMap resource
```

## 5.1.6 

**Release date:** 2018-02-18

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Changes for prometheus ingress RBAC. (#3722) 

### Default value changes

```diff
# No changes in this release
```

## 5.1.5 

**Release date:** 2018-02-13

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* exposing the init container image in the config to allow the user from overriding it (#3698) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 0fddee8..2f6d43b 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -188,6 +188,23 @@ configmapReload:
   ##
   resources: {}
 
+initChownData:
+  ## initChownData container name
+  ##
+  name: init-chown-data
+
+  ## initChownData container image
+  ##
+  image:
+    repository: busybox
+    tag: latest
+    pullPolicy: IfNotPresent
+
+  ## initChownData resource requests and limits
+  ## Ref: http://kubernetes.io/docs/user-guide/compute-resources/
+  ##
+  resources: {}
+
 kubeStateMetrics:
   ## If false, kube-state-metrics will not be installed
   ##
```

## 5.1.4 

**Release date:** 2018-02-12

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* adding toleration support for missing components (#3522) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 7cfd7a5..0fddee8 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -73,6 +73,15 @@ alertmanager:
   # strategy:
   #   type: Recreate
 
+  ## Node tolerations for alertmanager scheduling to nodes with taints
+  ## Ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
+  ##
+  tolerations: []
+    # - key: "key"
+    #   operator: "Equal|Exists"
+    #   value: "value"
+    #   effect: "NoSchedule|PreferNoSchedule|NoExecute(1.6 only)"
+
   ## Node labels for alertmanager pod assignment
   ## Ref: https://kubernetes.io/docs/user-guide/node-selection/
   ##
@@ -202,6 +211,15 @@ kubeStateMetrics:
   ##
   args: {}
 
+  ## Node tolerations for kube-state-metrics scheduling to nodes with taints
+  ## Ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
+  ##
+  tolerations: []
+    # - key: "key"
+    #   operator: "Equal|Exists"
+    #   value: "value"
+    #   effect: "NoSchedule|PreferNoSchedule|NoExecute(1.6 only)"
+
   ## Node labels for kube-state-metrics pod assignment
   ## Ref: https://kubernetes.io/docs/user-guide/node-selection/
   ##
@@ -406,6 +424,7 @@ server:
 
   ## Node labels for Prometheus server pod assignment
   ## Ref: https://kubernetes.io/docs/user-guide/node-selection/
+  ##
   nodeSelector: {}
 
   persistentVolume:
@@ -538,6 +557,15 @@ pushgateway:
     #     hosts:
     #       - pushgateway.domain.com
 
+  ## Node tolerations for pushgateway scheduling to nodes with taints
+  ## Ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
+  ##
+  tolerations: []
+    # - key: "key"
+    #   operator: "Equal|Exists"
+    #   value: "value"
+    #   effect: "NoSchedule|PreferNoSchedule|NoExecute(1.6 only)"
+
   ## Node labels for pushgateway pod assignment
   ## Ref: https://kubernetes.io/docs/user-guide/node-selection/
   ##
```

## 5.1.3 

**Release date:** 2018-02-12

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Add mesh peer to prometheus alert service (#3681) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 2c935db..7cfd7a5 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -144,6 +144,10 @@ alertmanager:
     labels: {}
     clusterIP: ""
 
+    ## Enabling peer mesh service end points for enabling the HA alert manager
+    ## Ref: https://github.com/prometheus/alertmanager/blob/master/README.md
+    # enableMeshPeer : true
+
     ## List of IP addresses at which the alertmanager service is available
     ## Ref: https://kubernetes.io/docs/user-guide/services/#external-ips
     ##
```

## 5.1.2 

**Release date:** 2018-02-09

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] Fix node exporter args (#3585) 

### Default value changes

```diff
# No changes in this release
```

## 5.1.1 

**Release date:** 2018-02-09

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] fix liveness and readiness probe when the prefix has changed (#3500) 

### Default value changes

```diff
# No changes in this release
```

## 5.1.0 

**Release date:** 2018-02-07

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] update kube-state-metrics to v1.2.0 and fix rbac (#3572) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index b0c6583..2c935db 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -191,7 +191,7 @@ kubeStateMetrics:
   ##
   image:
     repository: k8s.gcr.io/kube-state-metrics
-    tag: v1.1.0
+    tag: v1.2.0
     pullPolicy: IfNotPresent
 
   ## kube-state-metrics container arguments
```

## 5.0.2 

**Release date:** 2018-01-30

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Word fixes (#3473) 

### Default value changes

```diff
# No changes in this release
```

## 5.0.1 

**Release date:** 2018-01-24

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] add livenessProbe and readinessProbe (#3042) 

### Default value changes

```diff
# No changes in this release
```

## 5.0.0 

**Release date:** 2018-01-23

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Add suport for prometheus 2.1 (#2767) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index ec023fa..b0c6583 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -17,7 +17,7 @@ alertmanager:
   ##
   image:
     repository: prom/alertmanager
-    tag: v0.9.1
+    tag: v0.13.0
     pullPolicy: IfNotPresent
 
   ## Additional alertmanager container arguments
@@ -31,7 +31,7 @@ alertmanager:
 
   ## External URL which can access alertmanager
   ## Maybe same with Ingress host name
-  baseURL: ""
+  baseURL: "/"
 
   ## Additional alertmanager container environment variable
   ## For instance to add a http_proxy
@@ -191,7 +191,7 @@ kubeStateMetrics:
   ##
   image:
     repository: k8s.gcr.io/kube-state-metrics
-    tag: v1.1.0-rc.0
+    tag: v1.1.0
     pullPolicy: IfNotPresent
 
   ## kube-state-metrics container arguments
@@ -253,7 +253,7 @@ nodeExporter:
   ##
   image:
     repository: prom/node-exporter
-    tag: v0.15.0
+    tag: v0.15.2
     pullPolicy: IfNotPresent
 
   ## Custom Update Strategy
@@ -332,13 +332,9 @@ server:
   ##
   image:
     repository: prom/prometheus
-    tag: v1.8.1
+    tag: v2.1.0
     pullPolicy: IfNotPresent
 
-  ## (optional) alertmanager URL
-  ## only used if alertmanager.enabled = false
-  alertmanagerURL: ""
-
   ## The URL prefix at which the container can be accessed. Useful in the case the '-web.external-url' includes a slug
   ## so that the various internal URLs are still able to access as they are in the default case.
   ## (Optional)
@@ -598,10 +594,10 @@ alertmanagerFiles:
 ## Prometheus server ConfigMap entries
 ##
 serverFiles:
-  alerts: ""
-  rules: ""
+  alerts: {}
+  rules: {}
 
-  prometheus.yml: |-
+  prometheus.yml:
     rule_files:
       - /etc/config/rules
       - /etc/config/alerts
```

## 4.6.18 

**Release date:** 2018-01-20

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] Fix pushgateway port number in NOTES.txt (#3329) 

### Default value changes

```diff
# No changes in this release
```

## 4.6.17 

**Release date:** 2018-01-10

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] Add support for kube-state-metrics container arguments (#2368) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 54734ba..ec023fa 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -194,6 +194,10 @@ kubeStateMetrics:
     tag: v1.1.0-rc.0
     pullPolicy: IfNotPresent
 
+  ## kube-state-metrics container arguments
+  ##
+  args: {}
+
   ## Node labels for kube-state-metrics pod assignment
   ## Ref: https://kubernetes.io/docs/user-guide/node-selection/
   ##
```

## 4.6.16 

**Release date:** 2018-01-04

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Add scrap target for kubernetes cAdvisor (#2700) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 7487b15..54734ba 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -692,6 +692,43 @@ serverFiles:
             target_label: __metrics_path__
             replacement: /api/v1/nodes/${1}/proxy/metrics
 
+
+      - job_name: 'kubernetes-nodes-cadvisor'
+
+        # Default to scraping over https. If required, just disable this or change to
+        # `http`.
+        scheme: https
+
+        # This TLS & bearer token file config is used to connect to the actual scrape
+        # endpoints for cluster components. This is separate to discovery auth
+        # configuration because discovery & scraping are two separate concerns in
+        # Prometheus. The discovery auth config is automatic if Prometheus runs inside
+        # the cluster. Otherwise, more config options have to be provided within the
+        # <kubernetes_sd_config>.
+        tls_config:
+          ca_file: /var/run/secrets/kubernetes.io/serviceaccount/ca.crt
+          # If your node certificates are self-signed or use a different CA to the
+          # master CA, then disable certificate verification below. Note that
+          # certificate verification is an integral part of a secure infrastructure
+          # so this should only be disabled in a controlled environment. You can
+          # disable certificate verification by uncommenting the line below.
+          #
+          insecure_skip_verify: true
+        bearer_token_file: /var/run/secrets/kubernetes.io/serviceaccount/token
+
+        kubernetes_sd_configs:
+          - role: node
+
+        relabel_configs:
+          - action: labelmap
+            regex: __meta_kubernetes_node_label_(.+)
+          - target_label: __address__
+            replacement: kubernetes.default.svc:443
+          - source_labels: [__meta_kubernetes_node_name]
+            regex: (.+)
+            target_label: __metrics_path__
+            replacement: /api/v1/nodes/${1}:4194/proxy/metrics
+
       # Scrape config for service endpoints.
       #
       # The relabeling allows the actual service scrape endpoint to be configured
```

## 4.6.15 

**Release date:** 2017-12-28

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Convert registry to k8s.gcr.io (#3169) 
* [stable/prometheus] Seperate between prefix and external url (#3036) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 15dc873..7487b15 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -27,6 +27,10 @@ alertmanager:
   ## The URL prefix at which the container can be accessed. Useful in the case the '-web.external-url' includes a slug
   ## so that the various internal URLs are still able to access as they are in the default case.
   ## (Optional)
+  prefixURL: ""
+
+  ## External URL which can access alertmanager
+  ## Maybe same with Ingress host name
   baseURL: ""
 
   ## Additional alertmanager container environment variable
@@ -186,7 +190,7 @@ kubeStateMetrics:
   ## kube-state-metrics container image
   ##
   image:
-    repository: gcr.io/google_containers/kube-state-metrics
+    repository: k8s.gcr.io/kube-state-metrics
     tag: v1.1.0-rc.0
     pullPolicy: IfNotPresent
 
@@ -334,6 +338,10 @@ server:
   ## The URL prefix at which the container can be accessed. Useful in the case the '-web.external-url' includes a slug
   ## so that the various internal URLs are still able to access as they are in the default case.
   ## (Optional)
+  prefixURL: ""
+
+  ## External URL which can access alertmanager
+  ## Maybe same with Ingress host name
   baseURL: ""
 
   ## Additional Prometheus server container arguments
```

## 4.6.14 

**Release date:** 2017-12-16

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Fix bugs for prometheus chart (#2915) 

### Default value changes

```diff
# No changes in this release
```

## 4.6.13 

**Release date:** 2017-11-29

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Update server/alertmanager for baseURL (#2728) 

### Default value changes

```diff
# No changes in this release
```

## 4.6.12 

**Release date:** 2017-11-09

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus]: remove outdated networkpolicy note (#2686) 

### Default value changes

```diff
# No changes in this release
```

## 4.6.11 

**Release date:** 2017-11-06

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] Add RollingUpdate for Prometheus node-exporter ds (#1861) 
* * Adding affinity to alertmanager, kube-state-metrics, pushgateway and server deployments. (#2599) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index fc46bfe..15dc873 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -248,6 +248,11 @@ nodeExporter:
     tag: v0.15.0
     pullPolicy: IfNotPresent
 
+  ## Custom Update Strategy
+  ##
+  updateStrategy:
+    type: OnDelete
+
   ## Additional node-exporter container arguments
   ##
   extraArgs: {}
```

## 4.6.10 

**Release date:** 2017-10-31

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] Update to v1.8.1 (#2604) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index c427a72..fc46bfe 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -319,7 +319,7 @@ server:
   ##
   image:
     repository: prom/prometheus
-    tag: v1.8.0
+    tag: v1.8.1
     pullPolicy: IfNotPresent
 
   ## (optional) alertmanager URL
```

## 4.6.9 

**Release date:** 2017-10-19

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] Update kube state metrics (#2477) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 3ecd82d..c427a72 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -187,7 +187,7 @@ kubeStateMetrics:
   ##
   image:
     repository: gcr.io/google_containers/kube-state-metrics
-    tag: v1.0.1
+    tag: v1.1.0-rc.0
     pullPolicy: IfNotPresent
 
   ## Node labels for kube-state-metrics pod assignment
```

## 4.6.8 

**Release date:** 2017-10-18

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Update prometheus to 1.8.0 (#2509) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 9250568..3ecd82d 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -17,7 +17,7 @@ alertmanager:
   ##
   image:
     repository: prom/alertmanager
-    tag: v0.5.1
+    tag: v0.9.1
     pullPolicy: IfNotPresent
 
   ## Additional alertmanager container arguments
@@ -187,7 +187,7 @@ kubeStateMetrics:
   ##
   image:
     repository: gcr.io/google_containers/kube-state-metrics
-    tag: v0.4.1
+    tag: v1.0.1
     pullPolicy: IfNotPresent
 
   ## Node labels for kube-state-metrics pod assignment
@@ -245,7 +245,7 @@ nodeExporter:
   ##
   image:
     repository: prom/node-exporter
-    tag: v0.14.0
+    tag: v0.15.0
     pullPolicy: IfNotPresent
 
   ## Additional node-exporter container arguments
@@ -319,7 +319,7 @@ server:
   ##
   image:
     repository: prom/prometheus
-    tag: v1.5.2
+    tag: v1.8.0
     pullPolicy: IfNotPresent
 
   ## (optional) alertmanager URL
```

## 4.6.7 

**Release date:** 2017-10-18

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] fix invalid argument for storage retention (#2515) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 00caaef..9250568 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -264,10 +264,10 @@ nodeExporter:
   ## Ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
   ##
   tolerations: []
-  #- key: "key"
-  #  operator: "Equal|Exists"
-  #  value: "value"
-  #  effect: "NoSchedule|PreferNoSchedule|NoExecute(1.6 only)"
+    # - key: "key"
+    #   operator: "Equal|Exists"
+    #   value: "value"
+    #   effect: "NoSchedule|PreferNoSchedule|NoExecute(1.6 only)"
 
   ## Node labels for node-exporter pod assignment
   ## Ref: https://kubernetes.io/docs/user-guide/node-selection/
@@ -382,10 +382,10 @@ server:
   ## Ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
   ##
   tolerations: []
-  #- key: "key"
-  #  operator: "Equal|Exists"
-  #  value: "value"
-  #  effect: "NoSchedule|PreferNoSchedule|NoExecute(1.6 only)"
+    # - key: "key"
+    #   operator: "Equal|Exists"
+    #   value: "value"
+    #   effect: "NoSchedule|PreferNoSchedule|NoExecute(1.6 only)"
 
   ## Node labels for Prometheus server pod assignment
   ## Ref: https://kubernetes.io/docs/user-guide/node-selection/
```

## 4.6.6 

**Release date:** 2017-10-03

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] Fix strategy indentation (#2390) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index aa7d469..00caaef 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -65,7 +65,7 @@ alertmanager:
     #     hosts:
     #       - alertmanager.domain.com
 
-  ## Server Deployment Strategy type
+  ## Alertmanager Deployment Strategy type
   # strategy:
   #   type: Recreate
 
```

## 4.6.5 

**Release date:** 2017-10-03

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] Add option for defining strategy (#2218) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index a18a478..aa7d469 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -374,6 +374,10 @@ server:
     #     hosts:
     #       - prometheus.domain.com
 
+  ## Server Deployment Strategy type
+  # strategy:
+  #   type: Recreate
+
   ## Node tolerations for server scheduling to nodes with taints
   ## Ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
   ##
```

## 4.6.4 

**Release date:** 2017-10-02

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Add option for defining strategy for AlertManager (#2309) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 757ec71..a18a478 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -65,6 +65,10 @@ alertmanager:
     #     hosts:
     #       - alertmanager.domain.com
 
+  ## Server Deployment Strategy type
+  # strategy:
+  #   type: Recreate
+
   ## Node labels for alertmanager pod assignment
   ## Ref: https://kubernetes.io/docs/user-guide/node-selection/
   ##
```

## 4.6.3 

**Release date:** 2017-10-01

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* add additional permissions for kube-state-metrics (#2086) 

### Default value changes

```diff
# No changes in this release
```

## 4.6.2 

**Release date:** 2017-10-01

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* add prometheus data retention config (#2108) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index c0b3a8e..757ec71 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -464,6 +464,10 @@ server:
   ##
   terminationGracePeriodSeconds: 300
 
+  ## Prometheus data retention period (i.e 360h)
+  ##
+  retention: ""
+
 pushgateway:
   ## If false, pushgateway will not be installed
   ##
```

## 4.6.1 

**Release date:** 2017-09-30

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] fix prometheus webhook URL for config reload (#2324) 

### Default value changes

```diff
# No changes in this release
```

## 4.6.0 

**Release date:** 2017-09-29

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] Add nodePort to AlertManager (#2299) 
* Updated proemetheus charts readme. (#1924) 
* [Prometheus] Make the base URL configurable (#1769) 
* name kubeStateMetrics port as 'metrics' (#1885) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index e59d53a..c0b3a8e 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -24,6 +24,11 @@ alertmanager:
   ##
   extraArgs: {}
 
+  ## The URL prefix at which the container can be accessed. Useful in the case the '-web.external-url' includes a slug
+  ## so that the various internal URLs are still able to access as they are in the default case.
+  ## (Optional)
+  baseURL: ""
+
   ## Additional alertmanager container environment variable
   ## For instance to add a http_proxy
   ##
@@ -139,6 +144,7 @@ alertmanager:
     loadBalancerIP: ""
     loadBalancerSourceRanges: []
     servicePort: 80
+    # nodePort: 30000
     type: ClusterIP
 
 ## Monitors ConfigMap changes and POSTs to a URL
@@ -316,6 +322,11 @@ server:
   ## only used if alertmanager.enabled = false
   alertmanagerURL: ""
 
+  ## The URL prefix at which the container can be accessed. Useful in the case the '-web.external-url' includes a slug
+  ## so that the various internal URLs are still able to access as they are in the default case.
+  ## (Optional)
+  baseURL: ""
+
   ## Additional Prometheus server container arguments
   ##
   extraArgs: {}
```

## 4.5.0 

**Release date:** 2017-08-18

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] RBAC, namespaces, and node-exporter default update (#1443) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 5687bf2..e59d53a 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -1,8 +1,14 @@
+rbac:
+  create: false
+
 alertmanager:
   ## If false, alertmanager will not be installed
   ##
   enabled: true
 
+  # Defines the serviceAccountName to use when `rbac.create=false`
+  serviceAccountName: default
+
   ## alertmanager container name
   ##
   name: alertmanager
@@ -160,6 +166,9 @@ kubeStateMetrics:
   ##
   enabled: true
 
+  # Defines the serviceAccountName to use when `rbac.create=false`
+  serviceAccountName: default
+
   ## kube-state-metrics container name
   ##
   name: kube-state-metrics
@@ -215,6 +224,9 @@ nodeExporter:
   ##
   enabled: true
 
+  # Defines the serviceAccountName to use when `rbac.create=false`
+  serviceAccountName: default
+
   ## node-exporter container name
   ##
   name: node-exporter
@@ -223,7 +235,7 @@ nodeExporter:
   ##
   image:
     repository: prom/node-exporter
-    tag: v0.13.0
+    tag: v0.14.0
     pullPolicy: IfNotPresent
 
   ## Additional node-exporter container arguments
@@ -290,6 +302,9 @@ server:
   ##
   name: server
 
+  # Defines the serviceAccountName to use when `rbac.create=false`
+  serviceAccountName: default
+
   ## Prometheus server container image
   ##
   image:
@@ -352,7 +367,7 @@ server:
   #  operator: "Equal|Exists"
   #  value: "value"
   #  effect: "NoSchedule|PreferNoSchedule|NoExecute(1.6 only)"
-  
+
   ## Node labels for Prometheus server pod assignment
   ## Ref: https://kubernetes.io/docs/user-guide/node-selection/
   nodeSelector: {}
```

## 4.4.0 

**Release date:** 2017-08-14

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] Allow explicit NodePort value for prometheus (#1427) 

### Default value changes

```diff
# No changes in this release
```

## 4.3.0 

**Release date:** 2017-08-14

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Prometheus Helm Chart ConfigMap override for alertmanager and server #1310 (#1311) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index b569459..5687bf2 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -23,6 +23,12 @@ alertmanager:
   ##
   extraEnv: {}
 
+  ## ConfigMap override where fullname is {{.Release.Name}}-{{.Values.alertmanager.configMapOverrideName}}
+  ## Defining configMapOverrideName will cause templates/alertmanager-configmap.yaml
+  ## to NOT generate a ConfigMap resource
+  ##
+  configMapOverrideName: ""
+
   ingress:
     ## If true, alertmanager Ingress will be created
     ##
@@ -307,6 +313,12 @@ server:
     #   hostPath: /etc/kubernetes/certs
     #   readOnly: true
 
+  ## ConfigMap override where fullname is {{.Release.Name}}-{{.Values.server.configMapOverrideName}}
+  ## Defining configMapOverrideName will cause templates/server-configmap.yaml
+  ## to NOT generate a ConfigMap resource
+  ##
+  configMapOverrideName: ""
+
   ingress:
     ## If true, Prometheus server Ingress will be created
     ##
```

## 4.1.1 

**Release date:** 2017-08-13

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* feat(server-deployment) add taint toleration to server (#1635) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index bca91e3..b569459 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -236,6 +236,10 @@ nodeExporter:
   ## Ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
   ##
   tolerations: []
+  #- key: "key"
+  #  operator: "Equal|Exists"
+  #  value: "value"
+  #  effect: "NoSchedule|PreferNoSchedule|NoExecute(1.6 only)"
 
   ## Node labels for node-exporter pod assignment
   ## Ref: https://kubernetes.io/docs/user-guide/node-selection/
@@ -328,6 +332,15 @@ server:
     #     hosts:
     #       - prometheus.domain.com
 
+  ## Node tolerations for server scheduling to nodes with taints
+  ## Ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
+  ##
+  tolerations: []
+  #- key: "key"
+  #  operator: "Equal|Exists"
+  #  value: "value"
+  #  effect: "NoSchedule|PreferNoSchedule|NoExecute(1.6 only)"
+  
   ## Node labels for Prometheus server pod assignment
   ## Ref: https://kubernetes.io/docs/user-guide/node-selection/
   nodeSelector: {}
```

## 4.1.0 

**Release date:** 2017-08-02

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] Push Gateway support (#1275) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index f62922e..bca91e3 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -413,6 +413,89 @@ server:
   ##
   terminationGracePeriodSeconds: 300
 
+pushgateway:
+  ## If false, pushgateway will not be installed
+  ##
+  enabled: true
+
+  ## pushgateway container name
+  ##
+  name: pushgateway
+
+  ## pushgateway container image
+  ##
+  image:
+    repository: prom/pushgateway
+    tag: v0.4.0
+    pullPolicy: IfNotPresent
+
+  ## Additional pushgateway container arguments
+  ##
+  extraArgs: {}
+
+  ingress:
+    ## If true, pushgateway Ingress will be created
+    ##
+    enabled: false
+
+    ## pushgateway Ingress annotations
+    ##
+    annotations:
+    #   kubernetes.io/ingress.class: nginx
+    #   kubernetes.io/tls-acme: 'true'
+
+    ## pushgateway Ingress hostnames
+    ## Must be provided if Ingress is enabled
+    ##
+    hosts: []
+    #   - pushgateway.domain.com
+
+    ## pushgateway Ingress TLS configuration
+    ## Secrets must be manually created in the namespace
+    ##
+    tls: []
+    #   - secretName: prometheus-alerts-tls
+    #     hosts:
+    #       - pushgateway.domain.com
+
+  ## Node labels for pushgateway pod assignment
+  ## Ref: https://kubernetes.io/docs/user-guide/node-selection/
+  ##
+  nodeSelector: {}
+
+  ## Annotations to be added to pushgateway pods
+  ##
+  podAnnotations: {}
+
+  replicaCount: 1
+
+  ## pushgateway resource requests and limits
+  ## Ref: http://kubernetes.io/docs/user-guide/compute-resources/
+  ##
+  resources: {}
+    # limits:
+    #   cpu: 10m
+    #   memory: 32Mi
+    # requests:
+    #   cpu: 10m
+    #   memory: 32Mi
+
+  service:
+    annotations:
+      prometheus.io/probe: pushgateway
+    labels: {}
+    clusterIP: ""
+
+    ## List of IP addresses at which the pushgateway service is available
+    ## Ref: https://kubernetes.io/docs/user-guide/services/#external-ips
+    ##
+    externalIPs: []
+
+    loadBalancerIP: ""
+    loadBalancerSourceRanges: []
+    servicePort: 9091
+    type: ClusterIP
+
 ## alertmanager ConfigMap entries
 ##
 alertmanagerFiles:
@@ -575,6 +658,17 @@ serverFiles:
             action: replace
             target_label: kubernetes_name
 
+      - job_name: 'prometheus-pushgateway'
+        honor_labels: true
+
+        kubernetes_sd_configs:
+          - role: service
+
+        relabel_configs:
+          - source_labels: [__meta_kubernetes_service_annotation_prometheus_io_probe]
+            action: keep
+            regex: pushgateway
+
       # Example scrape config for probing services via the Blackbox Exporter.
       #
       # The relabeling allows the actual service scrape endpoint to be configured
```

## 4.0.0 

**Release date:** 2017-08-01

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* prometheus: switch to Kubernetes 1.6 storage class specification (#1332) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index c443a5e..f62922e 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -84,10 +84,13 @@ alertmanager:
     size: 2Gi
 
     ## alertmanager data Persistent Volume Storage Class
-    ## If defined, volume.beta.kubernetes.io/storage-class: <storageClass>
-    ## Default: volume.alpha.kubernetes.io/storage-class: default
+    ## If defined, storageClassName: <storageClass>
+    ## If set to "-", storageClassName: "", which disables dynamic provisioning
+    ## If undefined (the default) or set to null, no storageClassName spec is
+    ##   set, choosing the default provisioner.  (gp2 on AWS, standard on
+    ##   GKE, AWS & OpenStack)
     ##
-    storageClass: ""
+    # storageClass: "-"
 
     ## Subdirectory of alertmanager data Persistent Volume to mount
     ## Useful if the volume's root directory is not empty
@@ -360,10 +363,13 @@ server:
     size: 8Gi
 
     ## Prometheus server data Persistent Volume Storage Class
-    ## If defined, volume.beta.kubernetes.io/storage-class: <storageClass>
-    ## Default: volume.alpha.kubernetes.io/storage-class: default
+    ## If defined, storageClassName: <storageClass>
+    ## If set to "-", storageClassName: "", which disables dynamic provisioning
+    ## If undefined (the default) or set to null, no storageClassName spec is
+    ##   set, choosing the default provisioner.  (gp2 on AWS, standard on
+    ##   GKE, AWS & OpenStack)
     ##
-    storageClass: ""
+    # storageClass: "-"
 
     ## Subdirectory of Prometheus server data Persistent Volume to mount
     ## Useful if the volume's root directory is not empty
```

## 3.2.0 

**Release date:** 2017-08-01

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Add NetworkPolicy for Prometheus (#1327) 
* [stable/prometheus] Support for 1.6 tolerations (#1297) 
* Use consistent whitespace in template placeholders (#1437) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 0dc510c..c443a5e 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -229,6 +229,11 @@ nodeExporter:
     #   hostPath: /var/lib/node-exporter
     #   readOnly: true
 
+  ## Node tolerations for node-exporter scheduling to nodes with taints
+  ## Ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
+  ##
+  tolerations: []
+
   ## Node labels for node-exporter pod assignment
   ## Ref: https://kubernetes.io/docs/user-guide/node-selection/
   ##
@@ -630,3 +635,8 @@ serverFiles:
           - source_labels: [__meta_kubernetes_pod_name]
             action: replace
             target_label: kubernetes_pod_name
+
+networkPolicy:
+  ## Enable creation of NetworkPolicy resources.
+  ##
+  enabled: false
```

## 3.1.0 

**Release date:** 2017-07-05

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Prometheus: modify config to support k8s 1.6 by default (#1080) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 18592d2..0dc510c 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -515,6 +515,12 @@ serverFiles:
         relabel_configs:
           - action: labelmap
             regex: __meta_kubernetes_node_label_(.+)
+          - target_label: __address__
+            replacement: kubernetes.default.svc:443
+          - source_labels: [__meta_kubernetes_node_name]
+            regex: (.+)
+            target_label: __metrics_path__
+            replacement: /api/v1/nodes/${1}/proxy/metrics
 
       # Scrape config for service endpoints.
       #
```

## 3.0.3 

**Release date:** 2017-07-06

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] extra environment variable for alert manager (#1237) 
* [stable/prometheus] Add extraHostPathMounts config (#862) 
* Fix annotations reference (#883) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index c47b510..18592d2 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -18,6 +18,11 @@ alertmanager:
   ##
   extraArgs: {}
 
+  ## Additional alertmanager container environment variable
+  ## For instance to add a http_proxy
+  ##
+  extraEnv: {}
+
   ingress:
     ## If true, alertmanager Ingress will be created
     ##
@@ -216,6 +221,14 @@ nodeExporter:
   ##
   extraArgs: {}
 
+  ## Additional node-exporter hostPath mounts
+  ##
+  extraHostPathMounts: []
+    # - name: textfile-dir
+    #   mountPath: /srv/txt_collector
+    #   hostPath: /var/lib/node-exporter
+    #   readOnly: true
+
   ## Node labels for node-exporter pod assignment
   ## Ref: https://kubernetes.io/docs/user-guide/node-selection/
   ##
@@ -274,6 +287,14 @@ server:
   ##
   extraArgs: {}
 
+  ## Additional Prometheus server hostPath mounts
+  ##
+  extraHostPathMounts: []
+    # - name: certs-dir
+    #   mountPath: /etc/kubernetes/certs
+    #   hostPath: /etc/kubernetes/certs
+    #   readOnly: true
+
   ingress:
     ## If true, Prometheus server Ingress will be created
     ##
```

## 3.0.2 

**Release date:** 2017-04-17

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Prefix defined templates (#543) 
* Add extra service labels (#886) (#887) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 4a7bf4d..c47b510 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -108,6 +108,7 @@ alertmanager:
 
   service:
     annotations: {}
+    labels: {}
     clusterIP: ""
 
     ## List of IP addresses at which the alertmanager service is available
@@ -181,6 +182,7 @@ kubeStateMetrics:
   service:
     annotations:
       prometheus.io/scrape: "true"
+    labels: {}
 
     clusterIP: None
 
@@ -237,6 +239,7 @@ nodeExporter:
   service:
     annotations:
       prometheus.io/scrape: "true"
+    labels: {}
 
     clusterIP: None
 
@@ -361,6 +364,7 @@ server:
 
   service:
     annotations: {}
+    labels: {}
     clusterIP: ""
 
     ## List of IP addresses at which the Prometheus server service is available
```

## 3.0.1 

**Release date:** 2017-03-23

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] Fixed server ingress annotations (#853) 

### Default value changes

```diff
# No changes in this release
```

## 3.0.0 

**Release date:** 2017-03-22

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] v3.0.0 (#627) 
* [stable/prometheus] Add alertmanager extra args (#533) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 4e09add..4a7bf4d 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -1,272 +1,386 @@
 alertmanager:
-  ## Alertmanager service port
+  ## If false, alertmanager will not be installed
   ##
-  httpPort: 80
+  enabled: true
 
-  ## Alertmanager service port name
-  ## Default: 'http'
+  ## alertmanager container name
   ##
-  # httpPortName: http
+  name: alertmanager
+
+  ## alertmanager container image
+  ##
+  image:
+    repository: prom/alertmanager
+    tag: v0.5.1
+    pullPolicy: IfNotPresent
 
-  ## Alertmanager Docker image
+  ## Additional alertmanager container arguments
   ##
-  image: prom/alertmanager:v0.5.1
+  extraArgs: {}
 
   ingress:
-    ## If true, Alertmanager Ingress will be created
+    ## If true, alertmanager Ingress will be created
     ##
     enabled: false
 
-    ## Alertmanager Ingress annotations
+    ## alertmanager Ingress annotations
     ##
-    # annotations:
+    annotations: {}
     #   kubernetes.io/ingress.class: nginx
     #   kubernetes.io/tls-acme: 'true'
 
-    ## Alertmanager Ingress hostnames
+    ## alertmanager Ingress hostnames
     ## Must be provided if Ingress is enabled
     ##
-    # hosts:
+    hosts: []
     #   - alertmanager.domain.com
 
-    ## Alertmanager Ingress TLS configuration
+    ## alertmanager Ingress TLS configuration
     ## Secrets must be manually created in the namespace
     ##
-    # tls:
+    tls: []
     #   - secretName: prometheus-alerts-tls
     #     hosts:
     #       - alertmanager.domain.com
 
-  ## Alertmanager container name
+  ## Node labels for alertmanager pod assignment
+  ## Ref: https://kubernetes.io/docs/user-guide/node-selection/
   ##
-  name: alertmanager
+  nodeSelector: {}
 
   persistentVolume:
-    ## If true, AlertManager will create/use a Persistent Volume Claim
+    ## If true, alertmanager will create/use a Persistent Volume Claim
     ## If false, use emptyDir
     ##
     enabled: true
 
-    ## AlertManager data Persistent Volume access modes
+    ## alertmanager data Persistent Volume access modes
     ## Must match those of existing PV or dynamic provisioner
     ## Ref: http://kubernetes.io/docs/user-guide/persistent-volumes/
     ##
     accessModes:
       - ReadWriteOnce
 
-    ## AlertManager data Persistent Volume annotations
+    ## alertmanager data Persistent Volume Claim annotations
     ##
-    # annotations:
+    annotations: {}
 
-    ## AlertManager data Persistent Volume existing claim name
+    ## alertmanager data Persistent Volume existing claim name
     ## Requires alertmanager.persistentVolume.enabled: true
     ## If defined, PVC must be created manually before volume will be bound
-    # existingClaim:
+    existingClaim: ""
+
+    ## alertmanager data Persistent Volume mount root path
+    ##
+    mountPath: /data
 
-    ## AlertManager data Persistent Volume size
+    ## alertmanager data Persistent Volume size
     ##
     size: 2Gi
 
-    ## AlertManager data Persistent Volume Storage Class
+    ## alertmanager data Persistent Volume Storage Class
     ## If defined, volume.beta.kubernetes.io/storage-class: <storageClass>
     ## Default: volume.alpha.kubernetes.io/storage-class: default
     ##
-    # storageClass:
+    storageClass: ""
 
-    ## Subdirectory of the volume to mount at
+    ## Subdirectory of alertmanager data Persistent Volume to mount
     ## Useful if the volume's root directory is not empty
     ##
     subPath: ""
 
-  ## Alertmanager resource requests and limits
+  ## Annotations to be added to alertmanager pods
+  ##
+  podAnnotations: {}
+
+  replicaCount: 1
+
+  ## alertmanager resource requests and limits
   ## Ref: http://kubernetes.io/docs/user-guide/compute-resources/
   ##
-  resources:
+  resources: {}
     # limits:
     #   cpu: 10m
     #   memory: 32Mi
-    requests:
-      cpu: 10m
-      memory: 32Mi
+    # requests:
+    #   cpu: 10m
+    #   memory: 32Mi
 
-  ## Alertmanager service type
-  ##
-  serviceType: ClusterIP
+  service:
+    annotations: {}
+    clusterIP: ""
 
-  ## Alertmanager data storage path
-  ## Default: '/data'
-  ##
-  # storagePath: /data
+    ## List of IP addresses at which the alertmanager service is available
+    ## Ref: https://kubernetes.io/docs/user-guide/services/#external-ips
+    ##
+    externalIPs: []
+
+    loadBalancerIP: ""
+    loadBalancerSourceRanges: []
+    servicePort: 80
+    type: ClusterIP
 
 ## Monitors ConfigMap changes and POSTs to a URL
 ## Ref: https://github.com/jimmidyson/configmap-reload
 ##
 configmapReload:
-  ## Configmap-reload Docker image
+  ## configmap-reload container name
   ##
-  image: jimmidyson/configmap-reload:v0.1
+  name: configmap-reload
 
-  ## Configmap-reload container name
+  ## configmap-reload container image
   ##
-  name: configmap-reload
+  image:
+    repository: jimmidyson/configmap-reload
+    tag: v0.1
+    pullPolicy: IfNotPresent
 
-## Global imagePullPolicy
-## Default: 'Always' if image tag is 'latest', else 'IfNotPresent'
-## Ref: http://kubernetes.io/docs/user-guide/images/#pre-pulling-images
-##
-# imagePullPolicy:
+  ## configmap-reload resource requests and limits
+  ## Ref: http://kubernetes.io/docs/user-guide/compute-resources/
+  ##
+  resources: {}
 
 kubeStateMetrics:
-  ## Kube-state-metrics service port
+  ## If false, kube-state-metrics will not be installed
+  ##
+  enabled: true
+
+  ## kube-state-metrics container name
   ##
-  httpPort: 80
+  name: kube-state-metrics
 
-  ## Kube-state-metrics service port name
-  ## Default: 'http'
+  ## kube-state-metrics container image
   ##
-  # httpPortName: http
+  image:
+    repository: gcr.io/google_containers/kube-state-metrics
+    tag: v0.4.1
+    pullPolicy: IfNotPresent
 
-  ## Kube-state-metrics Docker image
+  ## Node labels for kube-state-metrics pod assignment
+  ## Ref: https://kubernetes.io/docs/user-guide/node-selection/
   ##
-  image: gcr.io/google_containers/kube-state-metrics:v0.3.0
+  nodeSelector: {}
 
-  ## Kube-state-metrics container name
+  ## Annotations to be added to kube-state-metrics pods
   ##
-  name: kube-state-metrics
+  podAnnotations: {}
 
-  ## Kube-state-metrics resource requests and limits
+  replicaCount: 1
+
+  ## kube-state-metrics resource requests and limits
   ## Ref: http://kubernetes.io/docs/user-guide/compute-resources/
   ##
-  resources:
+  resources: {}
     # limits:
     #   cpu: 10m
     #   memory: 16Mi
-    requests:
-      cpu: 10m
-      memory: 16Mi
+    # requests:
+    #   cpu: 10m
+    #   memory: 16Mi
+
+  service:
+    annotations:
+      prometheus.io/scrape: "true"
+
+    clusterIP: None
+
+    ## List of IP addresses at which the kube-state-metrics service is available
+    ## Ref: https://kubernetes.io/docs/user-guide/services/#external-ips
+    ##
+    externalIPs: []
+
+    loadBalancerIP: ""
+    loadBalancerSourceRanges: []
+    servicePort: 80
+    type: ClusterIP
 
-  ## Kube-state-metrics service type
+nodeExporter:
+  ## If false, node-exporter will not be installed
   ##
-  serviceType: ClusterIP
+  enabled: true
 
-server:
-  ## Server Pod annotations:
+  ## node-exporter container name
+  ##
+  name: node-exporter
+
+  ## node-exporter container image
+  ##
+  image:
+    repository: prom/node-exporter
+    tag: v0.13.0
+    pullPolicy: IfNotPresent
+
+  ## Additional node-exporter container arguments
   ##
-  # annotations:
-  #   iam.amazonaws.com/role: prometheus
+  extraArgs: {}
 
-  ## Additional Server container arguments
+  ## Node labels for node-exporter pod assignment
+  ## Ref: https://kubernetes.io/docs/user-guide/node-selection/
   ##
-  # extraArgs:
+  nodeSelector: {}
+
+  ## Annotations to be added to node-exporter pods
+  ##
+  podAnnotations: {}
+
+  ## node-exporter resource limits & requests
+  ## Ref: https://kubernetes.io/docs/user-guide/compute-resources/
+  ##
+  resources: {}
+    # limits:
+    #   cpu: 200m
+    #   memory: 50Mi
+    # requests:
+    #   cpu: 100m
+    #   memory: 30Mi
 
-  ## Server service port
+  service:
+    annotations:
+      prometheus.io/scrape: "true"
+
+    clusterIP: None
+
+    ## List of IP addresses at which the node-exporter service is available
+    ## Ref: https://kubernetes.io/docs/user-guide/services/#external-ips
+    ##
+    externalIPs: []
+
+    hostPort: 9100
+    loadBalancerIP: ""
+    loadBalancerSourceRanges: []
+    servicePort: 9100
+    type: ClusterIP
+
+server:
+  ## Prometheus server container name
   ##
-  httpPort: 80
+  name: server
 
-  ## Server service port name
-  ## Default: 'http'
+  ## Prometheus server container image
   ##
-  # httpPortName: http
+  image:
+    repository: prom/prometheus
+    tag: v1.5.2
+    pullPolicy: IfNotPresent
+
+  ## (optional) alertmanager URL
+  ## only used if alertmanager.enabled = false
+  alertmanagerURL: ""
 
-  ## Server Docker image
+  ## Additional Prometheus server container arguments
   ##
-  image: prom/prometheus:v1.5.2
+  extraArgs: {}
 
   ingress:
-    ## If true, Server Ingress will be created
+    ## If true, Prometheus server Ingress will be created
     ##
     enabled: false
 
-    ## Server Ingress annotations
+    ## Prometheus server Ingress annotations
     ##
-    # annotations:
+    annotations: {}
     #   kubernetes.io/ingress.class: nginx
     #   kubernetes.io/tls-acme: 'true'
 
-    ## Server Ingress hostnames
+    ## Prometheus server Ingress hostnames
     ## Must be provided if Ingress is enabled
     ##
-    # hosts:
+    hosts: []
     #   - prometheus.domain.com
 
-    ## Server Ingress TLS configuration
+    ## Prometheus server Ingress TLS configuration
     ## Secrets must be manually created in the namespace
     ##
-    # tls:
+    tls: []
     #   - secretName: prometheus-server-tls
     #     hosts:
     #       - prometheus.domain.com
 
-  ## Server container name
-  ##
-  name: server
+  ## Node labels for Prometheus server pod assignment
+  ## Ref: https://kubernetes.io/docs/user-guide/node-selection/
+  nodeSelector: {}
 
   persistentVolume:
-    ## If true, Server will create/use a Persistent Volume Claim
+    ## If true, Prometheus server will create/use a Persistent Volume Claim
     ## If false, use emptyDir
     ##
     enabled: true
 
-    ## Server data Persistent Volume access modes
+    ## Prometheus server data Persistent Volume access modes
     ## Must match those of existing PV or dynamic provisioner
     ## Ref: http://kubernetes.io/docs/user-guide/persistent-volumes/
     ##
     accessModes:
       - ReadWriteOnce
 
-    ## Server data Persistent Volume annotations
+    ## Prometheus server data Persistent Volume annotations
     ##
-    # annotations:
+    annotations: {}
 
-    ## Server data Persistent Volume existing claim name
+    ## Prometheus server data Persistent Volume existing claim name
     ## Requires server.persistentVolume.enabled: true
     ## If defined, PVC must be created manually before volume will be bound
-    # existingClaim:
+    existingClaim: ""
 
-    ## Server data Persistent Volume size
+    ## Prometheus server data Persistent Volume mount root path
+    ##
+    mountPath: /data
+
+    ## Prometheus server data Persistent Volume size
     ##
     size: 8Gi
 
-    ## AlertManager data Persistent Volume Storage Class
+    ## Prometheus server data Persistent Volume Storage Class
     ## If defined, volume.beta.kubernetes.io/storage-class: <storageClass>
     ## Default: volume.alpha.kubernetes.io/storage-class: default
     ##
-    # storageClass:
+    storageClass: ""
 
-    ## Subdirectory of the volume to mount at
+    ## Subdirectory of Prometheus server data Persistent Volume to mount
     ## Useful if the volume's root directory is not empty
     ##
     subPath: ""
 
-  ## Server resource requests and limits
+  ## Annotations to be added to Prometheus server pods
+  ##
+  podAnnotations: {}
+    # iam.amazonaws.com/role: prometheus
+
+  replicaCount: 1
+
+  ## Prometheus server resource requests and limits
   ## Ref: http://kubernetes.io/docs/user-guide/compute-resources/
   ##
-  resources:
+  resources: {}
     # limits:
     #   cpu: 500m
     #   memory: 512Mi
-    requests:
-      cpu: 500m
-      memory: 512Mi
+    # requests:
+    #   cpu: 500m
+    #   memory: 512Mi
 
-  ## Server service type
-  ##
-  serviceType: ClusterIP
+  service:
+    annotations: {}
+    clusterIP: ""
 
-  ## Server local data storage path
-  ## Default: '/data'
-  ##
-  # storageLocalPath: /data
+    ## List of IP addresses at which the Prometheus server service is available
+    ## Ref: https://kubernetes.io/docs/user-guide/services/#external-ips
+    ##
+    externalIPs: []
+
+    loadBalancerIP: ""
+    loadBalancerSourceRanges: []
+    servicePort: 80
+    type: ClusterIP
 
-  ## Server Pod termination grace period
-  ## Default: 300s (5m)
+  ## Prometheus server pod termination grace period
   ##
-  # terminationGracePeriodSeconds: 300
+  terminationGracePeriodSeconds: 300
 
-## Alertmanager ConfigMap entries
+## alertmanager ConfigMap entries
 ##
 alertmanagerFiles:
-  alertmanager.yml: |
+  alertmanager.yml: |-
     global:
       # slack_api_url: ''
 
@@ -282,17 +396,13 @@ alertmanagerFiles:
       receiver: default-receiver
       repeat_interval: 3h
 
-## Server ConfigMap entries
+## Prometheus server ConfigMap entries
 ##
 serverFiles:
   alerts: ""
   rules: ""
 
-  prometheus.yml: |
-    global:
-      scrape_interval:     15s
-      evaluation_interval: 15s
-
+  prometheus.yml: |-
     rule_files:
       - /etc/config/rules
       - /etc/config/alerts
@@ -303,7 +413,7 @@ serverFiles:
           - targets:
             - localhost:9090
 
-      # Scrape configurations for running Prometheus on a Kubernetes cluster.
+      # A scrape configuration for running Prometheus on a Kubernetes cluster.
       # This uses separate scrape configs for cluster components (i.e. API server, node)
       # and services to allow each to use different authentication configs.
       #
@@ -311,7 +421,16 @@ serverFiles:
       # `labelmap` relabeling action.
 
       # Scrape config for API servers.
-      - job_name: kubernetes-apiservers
+      #
+      # Kubernetes exposes API servers as endpoints to the default/kubernetes
+      # service so this uses `endpoints` role and uses relabelling to only keep
+      # the endpoints associated with the default/kubernetes service using the
+      # default named port `https`. This works for single API server deployments as
+      # well as HA API server deployments.
+      - job_name: 'kubernetes-apiservers'
+
+        kubernetes_sd_configs:
+          - role: endpoints
 
         # Default to scraping over https. If required, just disable this or change to
         # `http`.
@@ -319,8 +438,10 @@ serverFiles:
 
         # This TLS & bearer token file config is used to connect to the actual scrape
         # endpoints for cluster components. This is separate to discovery auth
-        # configuration (`in_cluster` below) because discovery & scraping are two
-        # separate concerns in Prometheus.
+        # configuration because discovery & scraping are two separate concerns in
+        # Prometheus. The discovery auth config is automatic if Prometheus runs inside
+        # the cluster. Otherwise, more config options have to be provided within the
+        # <kubernetes_sd_config>.
         tls_config:
           ca_file: /var/run/secrets/kubernetes.io/serviceaccount/ca.crt
           # If your node certificates are self-signed or use a different CA to the
@@ -329,7 +450,7 @@ serverFiles:
           # so this should only be disabled in a controlled environment. You can
           # disable certificate verification by uncommenting the line below.
           #
-          # insecure_skip_verify: true
+          insecure_skip_verify: true
         bearer_token_file: /var/run/secrets/kubernetes.io/serviceaccount/token
 
         # Keep only the default/kubernetes service endpoints for the https port. This
@@ -340,7 +461,7 @@ serverFiles:
             action: keep
             regex: default;kubernetes;https
 
-      - job_name: kubernetes-nodes
+      - job_name: 'kubernetes-nodes'
 
         # Default to scraping over https. If required, just disable this or change to
         # `http`.
@@ -348,8 +469,10 @@ serverFiles:
 
         # This TLS & bearer token file config is used to connect to the actual scrape
         # endpoints for cluster components. This is separate to discovery auth
-        # configuration (`in_cluster` below) because discovery & scraping are two
-        # separate concerns in Prometheus.
+        # configuration because discovery & scraping are two separate concerns in
+        # Prometheus. The discovery auth config is automatic if Prometheus runs inside
+        # the cluster. Otherwise, more config options have to be provided within the
+        # <kubernetes_sd_config>.
         tls_config:
           ca_file: /var/run/secrets/kubernetes.io/serviceaccount/ca.crt
           # If your node certificates are self-signed or use a different CA to the
@@ -358,7 +481,7 @@ serverFiles:
           # so this should only be disabled in a controlled environment. You can
           # disable certificate verification by uncommenting the line below.
           #
-          # insecure_skip_verify: true
+          insecure_skip_verify: true
         bearer_token_file: /var/run/secrets/kubernetes.io/serviceaccount/token
 
         kubernetes_sd_configs:
@@ -403,7 +526,7 @@ serverFiles:
             replacement: $1:$2
           - action: labelmap
             regex: __meta_kubernetes_service_label_(.+)
-          - source_labels: [__meta_kubernetes_service_namespace]
+          - source_labels: [__meta_kubernetes_namespace]
             action: replace
             target_label: kubernetes_namespace
           - source_labels: [__meta_kubernetes_service_name]
@@ -416,31 +539,31 @@ serverFiles:
       # via the following annotations:
       #
       # * `prometheus.io/probe`: Only probe services that have a value of `true`
-      # - job_name: 'kubernetes-services'
-      #
-      #   metrics_path: /probe
-      #   params:
-      #     module: [http_2xx]
-      #
-      #   kubernetes_sd_configs:
-      #     - role: service
-      #
-      #   relabel_configs:
-      #     - source_labels: [__meta_kubernetes_service_annotation_prometheus_io_probe]
-      #       action: keep
-      #       regex: true
-      #     - source_labels: [__address__]
-      #       target_label: __param_target
-      #     - target_label: __address__
-      #       replacement: blackbox
-      #     - source_labels: [__param_target]
-      #       target_label: instance
-      #     - action: labelmap
-      #       regex: __meta_kubernetes_service_label_(.+)
-      #     - source_labels: [__meta_kubernetes_service_namespace]
-      #       target_label: kubernetes_namespace
-      #     - source_labels: [__meta_kubernetes_service_name]
-      #       target_label: kubernetes_name
+      - job_name: 'kubernetes-services'
+
+        metrics_path: /probe
+        params:
+          module: [http_2xx]
+
+        kubernetes_sd_configs:
+          - role: service
+
+        relabel_configs:
+          - source_labels: [__meta_kubernetes_service_annotation_prometheus_io_probe]
+            action: keep
+            regex: true
+          - source_labels: [__address__]
+            target_label: __param_target
+          - target_label: __address__
+            replacement: blackbox
+          - source_labels: [__param_target]
+            target_label: instance
+          - action: labelmap
+            regex: __meta_kubernetes_service_label_(.+)
+          - source_labels: [__meta_kubernetes_namespace]
+            target_label: kubernetes_namespace
+          - source_labels: [__meta_kubernetes_service_name]
+            target_label: kubernetes_name
 
       # Example scrape config for pods
       #
@@ -470,7 +593,7 @@ serverFiles:
             target_label: __address__
           - action: labelmap
             regex: __meta_kubernetes_pod_label_(.+)
-          - source_labels: [__meta_kubernetes_pod_namespace]
+          - source_labels: [__meta_kubernetes_namespace]
             action: replace
             target_label: kubernetes_namespace
           - source_labels: [__meta_kubernetes_pod_name]
```

## 2.0.4 

**Release date:** 2017-02-13

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] Bump server image version (#535) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index ed906bf..4e09add 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -72,10 +72,10 @@ alertmanager:
     ## Default: volume.alpha.kubernetes.io/storage-class: default
     ##
     # storageClass:
-    
+
     ## Subdirectory of the volume to mount at
     ## Useful if the volume's root directory is not empty
-    ## 
+    ##
     subPath: ""
 
   ## Alertmanager resource requests and limits
@@ -170,7 +170,7 @@ server:
 
   ## Server Docker image
   ##
-  image: prom/prometheus:v1.4.1
+  image: prom/prometheus:v1.5.2
 
   ingress:
     ## If true, Server Ingress will be created
@@ -232,10 +232,10 @@ server:
     ## Default: volume.alpha.kubernetes.io/storage-class: default
     ##
     # storageClass:
-    
+
     ## Subdirectory of the volume to mount at
     ## Useful if the volume's root directory is not empty
-    ## 
+    ##
     subPath: ""
 
   ## Server resource requests and limits
```

## 2.0.3 

**Release date:** 2017-02-03

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] Update to the recommended pvc patterns. (#463) 

### Default value changes

```diff
# No changes in this release
```

## 2.0.2 

**Release date:** 2017-01-31

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [stable/prometheus] Use volumeMount subPath to fix prometheus-server from crashing (#424) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 6482220..ed906bf 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -72,6 +72,11 @@ alertmanager:
     ## Default: volume.alpha.kubernetes.io/storage-class: default
     ##
     # storageClass:
+    
+    ## Subdirectory of the volume to mount at
+    ## Useful if the volume's root directory is not empty
+    ## 
+    subPath: ""
 
   ## Alertmanager resource requests and limits
   ## Ref: http://kubernetes.io/docs/user-guide/compute-resources/
@@ -227,6 +232,11 @@ server:
     ## Default: volume.alpha.kubernetes.io/storage-class: default
     ##
     # storageClass:
+    
+    ## Subdirectory of the volume to mount at
+    ## Useful if the volume's root directory is not empty
+    ## 
+    subPath: ""
 
   ## Server resource requests and limits
   ## Ref: http://kubernetes.io/docs/user-guide/compute-resources/
```

## 2.0.1 

**Release date:** 2017-01-19

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Prometheus: provide options to use existing PVCs (#252) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 865da65..6482220 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -42,7 +42,7 @@ alertmanager:
   name: alertmanager
 
   persistentVolume:
-    ## If true, AlertManager will create a Persistent Volume Claim
+    ## If true, AlertManager will create/use a Persistent Volume Claim
     ## If false, use emptyDir
     ##
     enabled: true
@@ -58,6 +58,11 @@ alertmanager:
     ##
     # annotations:
 
+    ## AlertManager data Persistent Volume existing claim name
+    ## Requires alertmanager.persistentVolume.enabled: true
+    ## If defined, PVC must be created manually before volume will be bound
+    # existingClaim:
+
     ## AlertManager data Persistent Volume size
     ##
     size: 2Gi
@@ -192,7 +197,7 @@ server:
   name: server
 
   persistentVolume:
-    ## If true, Server will create a Persistent Volume Claim
+    ## If true, Server will create/use a Persistent Volume Claim
     ## If false, use emptyDir
     ##
     enabled: true
@@ -208,6 +213,11 @@ server:
     ##
     # annotations:
 
+    ## Server data Persistent Volume existing claim name
+    ## Requires server.persistentVolume.enabled: true
+    ## If defined, PVC must be created manually before volume will be bound
+    # existingClaim:
+
     ## Server data Persistent Volume size
     ##
     size: 8Gi
```

## 2.0.0 

**Release date:** 2017-01-13

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* fixed Prometheus Ingress path rules (#375) 

### Default value changes

```diff
# No changes in this release
```

## 1.4.2 

**Release date:** 2017-01-09

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Unquote Prometheus storage paths (#350) 
* removed unneeded trim (#373) 
* Merge pull request #209 from sameersbn/misc-fixes 

### Default value changes

```diff
# No changes in this release
```

## 1.4.1 

**Release date:** 2016-12-11

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Merge pull request #251 from mgoodness/prometheus-1.4.0 
* specify namespace in `kubectl get svc` commands in NOTES.txt 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 4861c3e..865da65 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -10,7 +10,7 @@ alertmanager:
 
   ## Alertmanager Docker image
   ##
-  image: prom/alertmanager:v0.5.0
+  image: prom/alertmanager:v0.5.1
 
   ingress:
     ## If true, Alertmanager Ingress will be created
@@ -160,7 +160,7 @@ server:
 
   ## Server Docker image
   ##
-  image: prom/prometheus:v1.3.1
+  image: prom/prometheus:v1.4.1
 
   ingress:
     ## If true, Server Ingress will be created
```

## 1.3.1 

**Release date:** 2016-11-30

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Replace 'Spartakus' with 'Prometheus' (#264) 

### Default value changes

```diff
# No changes in this release
```

## 1.4.1 

**Release date:** 2016-11-29

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Prometheus 1.4.1 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 110b30d..865da65 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -160,7 +160,7 @@ server:
 
   ## Server Docker image
   ##
-  image: prom/prometheus:v1.4.0
+  image: prom/prometheus:v1.4.1
 
   ingress:
     ## If true, Server Ingress will be created
```

## 1.4.0 

**Release date:** 2016-11-25

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Bump Prometheus chart version 
* Prometheus: add kube-state-metrics (#236) 
* Default PVC SC is alpha:default 
* Default to alpha dynamic PV annotation 
* Update docs for charts moved to stable (#218) 

### Default value changes

```diff
diff --git a/charts/prometheus/values.yaml b/charts/prometheus/values.yaml
index 85d9bd2..110b30d 100644
--- a/charts/prometheus/values.yaml
+++ b/charts/prometheus/values.yaml
@@ -10,7 +10,7 @@ alertmanager:
 
   ## Alertmanager Docker image
   ##
-  image: prom/alertmanager:v0.5.0
+  image: prom/alertmanager:v0.5.1
 
   ingress:
     ## If true, Alertmanager Ingress will be created
@@ -56,13 +56,18 @@ alertmanager:
 
     ## AlertManager data Persistent Volume annotations
     ##
-    annotations:
-      volume.beta.kubernetes.io/storage-class: standard
+    # annotations:
 
     ## AlertManager data Persistent Volume size
     ##
     size: 2Gi
 
+    ## AlertManager data Persistent Volume Storage Class
+    ## If defined, volume.beta.kubernetes.io/storage-class: <storageClass>
+    ## Default: volume.alpha.kubernetes.io/storage-class: default
+    ##
+    # storageClass:
+
   ## Alertmanager resource requests and limits
   ## Ref: http://kubernetes.io/docs/user-guide/compute-resources/
   ##
@@ -101,6 +106,39 @@ configmapReload:
 ##
 # imagePullPolicy:
 
+kubeStateMetrics:
+  ## Kube-state-metrics service port
+  ##
+  httpPort: 80
+
+  ## Kube-state-metrics service port name
+  ## Default: 'http'
+  ##
+  # httpPortName: http
+
+  ## Kube-state-metrics Docker image
+  ##
+  image: gcr.io/google_containers/kube-state-metrics:v0.3.0
+
+  ## Kube-state-metrics container name
+  ##
+  name: kube-state-metrics
+
+  ## Kube-state-metrics resource requests and limits
+  ## Ref: http://kubernetes.io/docs/user-guide/compute-resources/
+  ##
+  resources:
+    # limits:
+    #   cpu: 10m
+    #   memory: 16Mi
+    requests:
+      cpu: 10m
+      memory: 16Mi
+
+  ## Kube-state-metrics service type
+  ##
+  serviceType: ClusterIP
+
 server:
   ## Server Pod annotations:
   ##
@@ -122,7 +160,7 @@ server:
 
   ## Server Docker image
   ##
-  image: prom/prometheus:v1.3.1
+  image: prom/prometheus:v1.4.0
 
   ingress:
     ## If true, Server Ingress will be created
@@ -168,13 +206,18 @@ server:
 
     ## Server data Persistent Volume annotations
     ##
-    annotations:
-      volume.beta.kubernetes.io/storage-class: standard
+    # annotations:
 
     ## Server data Persistent Volume size
     ##
     size: 8Gi
 
+    ## AlertManager data Persistent Volume Storage Class
+    ## If defined, volume.beta.kubernetes.io/storage-class: <storageClass>
+    ## Default: volume.alpha.kubernetes.io/storage-class: default
+    ##
+    # storageClass:
+
   ## Server resource requests and limits
   ## Ref: http://kubernetes.io/docs/user-guide/compute-resources/
   ##
```

## 1.3.1 

**Release date:** 2016-11-14

![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Move Prometheus to stable 

### Default value changes

```diff
alertmanager:
  ## Alertmanager service port
  ##
  httpPort: 80

  ## Alertmanager service port name
  ## Default: 'http'
  ##
  # httpPortName: http

  ## Alertmanager Docker image
  ##
  image: prom/alertmanager:v0.5.0

  ingress:
    ## If true, Alertmanager Ingress will be created
    ##
    enabled: false

    ## Alertmanager Ingress annotations
    ##
    # annotations:
    #   kubernetes.io/ingress.class: nginx
    #   kubernetes.io/tls-acme: 'true'

    ## Alertmanager Ingress hostnames
    ## Must be provided if Ingress is enabled
    ##
    # hosts:
    #   - alertmanager.domain.com

    ## Alertmanager Ingress TLS configuration
    ## Secrets must be manually created in the namespace
    ##
    # tls:
    #   - secretName: prometheus-alerts-tls
    #     hosts:
    #       - alertmanager.domain.com

  ## Alertmanager container name
  ##
  name: alertmanager

  persistentVolume:
    ## If true, AlertManager will create a Persistent Volume Claim
    ## If false, use emptyDir
    ##
    enabled: true

    ## AlertManager data Persistent Volume access modes
    ## Must match those of existing PV or dynamic provisioner
    ## Ref: http://kubernetes.io/docs/user-guide/persistent-volumes/
    ##
    accessModes:
      - ReadWriteOnce

    ## AlertManager data Persistent Volume annotations
    ##
    annotations:
      volume.beta.kubernetes.io/storage-class: standard

    ## AlertManager data Persistent Volume size
    ##
    size: 2Gi

  ## Alertmanager resource requests and limits
  ## Ref: http://kubernetes.io/docs/user-guide/compute-resources/
  ##
  resources:
    # limits:
    #   cpu: 10m
    #   memory: 32Mi
    requests:
      cpu: 10m
      memory: 32Mi

  ## Alertmanager service type
  ##
  serviceType: ClusterIP

  ## Alertmanager data storage path
  ## Default: '/data'
  ##
  # storagePath: /data

## Monitors ConfigMap changes and POSTs to a URL
## Ref: https://github.com/jimmidyson/configmap-reload
##
configmapReload:
  ## Configmap-reload Docker image
  ##
  image: jimmidyson/configmap-reload:v0.1

  ## Configmap-reload container name
  ##
  name: configmap-reload

## Global imagePullPolicy
## Default: 'Always' if image tag is 'latest', else 'IfNotPresent'
## Ref: http://kubernetes.io/docs/user-guide/images/#pre-pulling-images
##
# imagePullPolicy:

server:
  ## Server Pod annotations:
  ##
  # annotations:
  #   iam.amazonaws.com/role: prometheus

  ## Additional Server container arguments
  ##
  # extraArgs:

  ## Server service port
  ##
  httpPort: 80

  ## Server service port name
  ## Default: 'http'
  ##
  # httpPortName: http

  ## Server Docker image
  ##
  image: prom/prometheus:v1.3.1

  ingress:
    ## If true, Server Ingress will be created
    ##
    enabled: false

    ## Server Ingress annotations
    ##
    # annotations:
    #   kubernetes.io/ingress.class: nginx
    #   kubernetes.io/tls-acme: 'true'

    ## Server Ingress hostnames
    ## Must be provided if Ingress is enabled
    ##
    # hosts:
    #   - prometheus.domain.com

    ## Server Ingress TLS configuration
    ## Secrets must be manually created in the namespace
    ##
    # tls:
    #   - secretName: prometheus-server-tls
    #     hosts:
    #       - prometheus.domain.com

  ## Server container name
  ##
  name: server

  persistentVolume:
    ## If true, Server will create a Persistent Volume Claim
    ## If false, use emptyDir
    ##
    enabled: true

    ## Server data Persistent Volume access modes
    ## Must match those of existing PV or dynamic provisioner
    ## Ref: http://kubernetes.io/docs/user-guide/persistent-volumes/
    ##
    accessModes:
      - ReadWriteOnce

    ## Server data Persistent Volume annotations
    ##
    annotations:
      volume.beta.kubernetes.io/storage-class: standard

    ## Server data Persistent Volume size
    ##
    size: 8Gi

  ## Server resource requests and limits
  ## Ref: http://kubernetes.io/docs/user-guide/compute-resources/
  ##
  resources:
    # limits:
    #   cpu: 500m
    #   memory: 512Mi
    requests:
      cpu: 500m
      memory: 512Mi

  ## Server service type
  ##
  serviceType: ClusterIP

  ## Server local data storage path
  ## Default: '/data'
  ##
  # storageLocalPath: /data

  ## Server Pod termination grace period
  ## Default: 300s (5m)
  ##
  # terminationGracePeriodSeconds: 300

## Alertmanager ConfigMap entries
##
alertmanagerFiles:
  alertmanager.yml: |
    global:
      # slack_api_url: ''

    receivers:
      - name: default-receiver
        # slack_configs:
        #  - channel: '@you'
        #    send_resolved: true

    route:
      group_wait: 10s
      group_interval: 5m
      receiver: default-receiver
      repeat_interval: 3h

## Server ConfigMap entries
##
serverFiles:
  alerts: ""
  rules: ""

  prometheus.yml: |
    global:
      scrape_interval:     15s
      evaluation_interval: 15s

    rule_files:
      - /etc/config/rules
      - /etc/config/alerts

    scrape_configs:
      - job_name: prometheus
        static_configs:
          - targets:
            - localhost:9090

      # Scrape configurations for running Prometheus on a Kubernetes cluster.
      # This uses separate scrape configs for cluster components (i.e. API server, node)
      # and services to allow each to use different authentication configs.
      #
      # Kubernetes labels will be added as Prometheus labels on metrics via the
      # `labelmap` relabeling action.

      # Scrape config for API servers.
      - job_name: kubernetes-apiservers

        # Default to scraping over https. If required, just disable this or change to
        # `http`.
        scheme: https

        # This TLS & bearer token file config is used to connect to the actual scrape
        # endpoints for cluster components. This is separate to discovery auth
        # configuration (`in_cluster` below) because discovery & scraping are two
        # separate concerns in Prometheus.
        tls_config:
          ca_file: /var/run/secrets/kubernetes.io/serviceaccount/ca.crt
          # If your node certificates are self-signed or use a different CA to the
          # master CA, then disable certificate verification below. Note that
          # certificate verification is an integral part of a secure infrastructure
          # so this should only be disabled in a controlled environment. You can
          # disable certificate verification by uncommenting the line below.
          #
          # insecure_skip_verify: true
        bearer_token_file: /var/run/secrets/kubernetes.io/serviceaccount/token

        # Keep only the default/kubernetes service endpoints for the https port. This
        # will add targets for each API server which Kubernetes adds an endpoint to
        # the default/kubernetes service.
        relabel_configs:
          - source_labels: [__meta_kubernetes_namespace, __meta_kubernetes_service_name, __meta_kubernetes_endpoint_port_name]
            action: keep
            regex: default;kubernetes;https

      - job_name: kubernetes-nodes

        # Default to scraping over https. If required, just disable this or change to
        # `http`.
        scheme: https

        # This TLS & bearer token file config is used to connect to the actual scrape
        # endpoints for cluster components. This is separate to discovery auth
        # configuration (`in_cluster` below) because discovery & scraping are two
        # separate concerns in Prometheus.
        tls_config:
          ca_file: /var/run/secrets/kubernetes.io/serviceaccount/ca.crt
          # If your node certificates are self-signed or use a different CA to the
          # master CA, then disable certificate verification below. Note that
          # certificate verification is an integral part of a secure infrastructure
          # so this should only be disabled in a controlled environment. You can
          # disable certificate verification by uncommenting the line below.
          #
          # insecure_skip_verify: true
        bearer_token_file: /var/run/secrets/kubernetes.io/serviceaccount/token

        kubernetes_sd_configs:
          - role: node

        relabel_configs:
          - action: labelmap
            regex: __meta_kubernetes_node_label_(.+)

      # Scrape config for service endpoints.
      #
      # The relabeling allows the actual service scrape endpoint to be configured
      # via the following annotations:
      #
      # * `prometheus.io/scrape`: Only scrape services that have a value of `true`
      # * `prometheus.io/scheme`: If the metrics endpoint is secured then you will need
      # to set this to `https` & most likely set the `tls_config` of the scrape config.
      # * `prometheus.io/path`: If the metrics path is not `/metrics` override this.
      # * `prometheus.io/port`: If the metrics are exposed on a different port to the
      # service then set this appropriately.
      - job_name: 'kubernetes-service-endpoints'

        kubernetes_sd_configs:
          - role: endpoints

        relabel_configs:
          - source_labels: [__meta_kubernetes_service_annotation_prometheus_io_scrape]
            action: keep
            regex: true
          - source_labels: [__meta_kubernetes_service_annotation_prometheus_io_scheme]
            action: replace
            target_label: __scheme__
            regex: (https?)
          - source_labels: [__meta_kubernetes_service_annotation_prometheus_io_path]
            action: replace
            target_label: __metrics_path__
            regex: (.+)
          - source_labels: [__address__, __meta_kubernetes_service_annotation_prometheus_io_port]
            action: replace
            target_label: __address__
            regex: (.+)(?::\d+);(\d+)
            replacement: $1:$2
          - action: labelmap
            regex: __meta_kubernetes_service_label_(.+)
          - source_labels: [__meta_kubernetes_service_namespace]
            action: replace
            target_label: kubernetes_namespace
          - source_labels: [__meta_kubernetes_service_name]
            action: replace
            target_label: kubernetes_name

      # Example scrape config for probing services via the Blackbox Exporter.
      #
      # The relabeling allows the actual service scrape endpoint to be configured
      # via the following annotations:
      #
      # * `prometheus.io/probe`: Only probe services that have a value of `true`
      # - job_name: 'kubernetes-services'
      #
      #   metrics_path: /probe
      #   params:
      #     module: [http_2xx]
      #
      #   kubernetes_sd_configs:
      #     - role: service
      #
      #   relabel_configs:
      #     - source_labels: [__meta_kubernetes_service_annotation_prometheus_io_probe]
      #       action: keep
      #       regex: true
      #     - source_labels: [__address__]
      #       target_label: __param_target
      #     - target_label: __address__
      #       replacement: blackbox
      #     - source_labels: [__param_target]
      #       target_label: instance
      #     - action: labelmap
      #       regex: __meta_kubernetes_service_label_(.+)
      #     - source_labels: [__meta_kubernetes_service_namespace]
      #       target_label: kubernetes_namespace
      #     - source_labels: [__meta_kubernetes_service_name]
      #       target_label: kubernetes_name

      # Example scrape config for pods
      #
      # The relabeling allows the actual pod scrape endpoint to be configured via the
      # following annotations:
      #
      # * `prometheus.io/scrape`: Only scrape pods that have a value of `true`
      # * `prometheus.io/path`: If the metrics path is not `/metrics` override this.
      # * `prometheus.io/port`: Scrape the pod on the indicated port instead of the default of `9102`.
      - job_name: 'kubernetes-pods'

        kubernetes_sd_configs:
          - role: pod

        relabel_configs:
          - source_labels: [__meta_kubernetes_pod_annotation_prometheus_io_scrape]
            action: keep
            regex: true
          - source_labels: [__meta_kubernetes_pod_annotation_prometheus_io_path]
            action: replace
            target_label: __metrics_path__
            regex: (.+)
          - source_labels: [__address__, __meta_kubernetes_pod_annotation_prometheus_io_port]
            action: replace
            regex: (.+):(?:\d+);(\d+)
            replacement: ${1}:${2}
            target_label: __address__
          - action: labelmap
            regex: __meta_kubernetes_pod_label_(.+)
          - source_labels: [__meta_kubernetes_pod_namespace]
            action: replace
            target_label: kubernetes_namespace
          - source_labels: [__meta_kubernetes_pod_name]
            action: replace
            target_label: kubernetes_pod_name
```

---
Autogenerated from Helm Chart and git history using [helm-changelog](https://github.com/mogensen/helm-changelog)
