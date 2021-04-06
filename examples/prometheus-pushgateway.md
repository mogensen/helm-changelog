# Change Log

## 1.7.1 

**Release date:** 2021-03-05

![AppVersion: 1.3.0](https://img.shields.io/static/v1?label=AppVersion&message=1.3.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-pushgateway] Align resource label handling (#456) 

### Default value changes

```diff
# No changes in this release
```

## 1.7.0 

**Release date:** 2021-02-05

![AppVersion: 1.3.0](https://img.shields.io/static/v1?label=AppVersion&message=1.3.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-pushgateway] fix failed release (#637) 
* [prometheus-pushgateway] Adds ability to specify metricRelabelings and relabelings (#453) 

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index 12e9e3d..e05ee20 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -141,6 +141,23 @@ serviceMonitor:
   # [Scraping Pushgateway](https://github.com/prometheus/pushgateway#configure-the-pushgateway-as-a-target-to-scrape)
   honorLabels: true
 
+  ## Metric relabel configs to apply to samples before ingestion.
+  ## [Metric Relabeling](https://prometheus.io/docs/prometheus/latest/configuration/configuration/#metric_relabel_configs)
+  metricRelabelings: []
+  # - action: keep
+  #   regex: 'kube_(daemonset|deployment|pod|namespace|node|statefulset).+'
+  #   sourceLabels: [__name__]
+
+  ## Relabel configs to apply to samples before ingestion.
+  ## [Relabeling](https://prometheus.io/docs/prometheus/latest/configuration/configuration/#relabel_config)
+  relabelings: []
+  # - sourceLabels: [__meta_kubernetes_pod_node_name]
+  #   separator: ;
+  #   regex: ^(.*)$
+  #   targetLabel: nodename
+  #   replacement: $1
+  #   action: replace
+
 # The values to set in the PodDisruptionBudget spec (minAvailable/maxUnavailable)
 # If not set then a PodDisruptionBudget will not be created
 podDisruptionBudget: {}
```

## 1.6.0 

**Release date:** 2021-02-04

![AppVersion: 1.3.0](https://img.shields.io/static/v1?label=AppVersion&message=1.3.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-pushgateway] Adds imagePullSecrets value for deployment (#608) 

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index 4a56f4f..12e9e3d 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -13,6 +13,9 @@ image:
   tag: v1.3.0
   pullPolicy: IfNotPresent
 
+# Optional pod imagePullSecrets
+imagePullSecrets: []
+
 service:
   type: ClusterIP
   port: 9091
```

## 1.5.1 

**Release date:** 2020-12-18

![AppVersion: 1.3.0](https://img.shields.io/static/v1?label=AppVersion&message=1.3.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-pushgateway] Add missing fields within values.yaml (#219) 

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index 169a5cd..4a56f4f 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -1,6 +1,13 @@
 # Default values for prometheus-pushgateway.
 # This is a YAML-formatted file.
 # Declare variables to be passed into your templates.
+
+# Provide a name in place of prometheus-pushgateway for `app:` labels
+nameOverride: ""
+
+# Provide a name to substitute for the full names of resources
+fullnameOverride: ""
+
 image:
   repository: prom/pushgateway
   tag: v1.3.0
```

## 1.5.0 

**Release date:** 2020-11-04

![AppVersion: 1.3.0](https://img.shields.io/static/v1?label=AppVersion&message=1.3.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* pushgateway 1.3.0 - update chart version, pushgateway version and values (#195) 

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index 63b2a43..169a5cd 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -3,7 +3,7 @@
 # Declare variables to be passed into your templates.
 image:
   repository: prom/pushgateway
-  tag: v1.2.0
+  tag: v1.3.0
   pullPolicy: IfNotPresent
 
 service:
```

## 1.4.2 

**Release date:** 2020-08-20

![AppVersion: 1.2.0](https://img.shields.io/static/v1?label=AppVersion&message=1.2.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Prep initial charts indexing (#14) 

### Default value changes

```diff
# No changes in this release
```

## 1.4.1 

**Release date:** 2020-07-07

![AppVersion: 1.2.0](https://img.shields.io/static/v1?label=AppVersion&message=1.2.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-pushgateway] Fix networkPolicy podSelector and ingress rule (#23053) 

### Default value changes

```diff
# No changes in this release
```

## 1.4.0 

**Release date:** 2020-04-06

![AppVersion: 1.2.0](https://img.shields.io/static/v1?label=AppVersion&message=1.2.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-pushgateway] configureable deployment strategy (#21685) 

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index 1ae0906..63b2a43 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -135,6 +135,10 @@ serviceMonitor:
 # If not set then a PodDisruptionBudget will not be created
 podDisruptionBudget: {}
 
+# Deployment Strategy type
+strategy:
+  type: Recreate
+
 persistentVolume:
   ## If true, pushgateway will create/use a Persistent Volume Claim
   ## If false, use emptyDir
```

## 1.3.0 

**Release date:** 2020-03-12

![AppVersion: 1.2.0](https://img.shields.io/static/v1?label=AppVersion&message=1.2.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-pushgateway] upgrade to latest release (#21424) 

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index 7709c9d..1ae0906 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -3,7 +3,7 @@
 # Declare variables to be passed into your templates.
 image:
   repository: prom/pushgateway
-  tag: v1.0.1
+  tag: v1.2.0
   pullPolicy: IfNotPresent
 
 service:
```

## 1.2.15 

**Release date:** 2020-03-04

![AppVersion: 1.0.1](https://img.shields.io/static/v1?label=AppVersion&message=1.0.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* fix hard-coded path and default to "/" (#21253) 

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index 55d0bf8..7709c9d 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -67,6 +67,8 @@ ingress:
   ## Enable Ingress.
   ##
   enabled: false
+  # AWS ALB requires path of /*
+  path: /
 
     ## Annotations.
     ##
```

## 1.2.14 

**Release date:** 2020-01-31

![AppVersion: 1.0.1](https://img.shields.io/static/v1?label=AppVersion&message=1.0.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Add annotation support to pushgateway service (#20351) 

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index d9ff624..55d0bf8 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -17,6 +17,9 @@ podAnnotations: {}
 # Optional pod labels
 podLabels: {}
 
+# Optional service annotations
+serviceAnnotations: {}
+
 # Optional service labels
 serviceLabels: {}
 
```

## 1.2.13 

**Release date:** 2020-01-13

![AppVersion: 1.0.1](https://img.shields.io/static/v1?label=AppVersion&message=1.0.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-pushgateway] Fix persistence support (#20048) 

### Default value changes

```diff
# No changes in this release
```

## 1.2.12 

**Release date:** 2020-01-10

![AppVersion: 1.0.1](https://img.shields.io/static/v1?label=AppVersion&message=1.0.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-pushgateway] Add persistence support (#16040) 

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index a714a39..d9ff624 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -23,12 +23,20 @@ serviceLabels: {}
 # Optional serviceAccount labels
 serviceAccountLabels: {}
 
-# Optional additional arguments
-extraArgs: []
+# Optional persistentVolume labels
+persistentVolumeLabels: {}
 
 # Optional additional environment variables
 extraVars: []
 
+## Additional pushgateway container arguments
+##
+## example:
+## extraArgs:
+##  - --persistence.file=/data/pushgateway.data
+##  - --persistence.interval=5m
+extraArgs: []
+
 resources: {}
   # We usually recommend not to specify default resources and to leave this as a conscious
   # choice for the user. This also increases chances charts run on environments with little
@@ -88,6 +96,13 @@ nodeSelector: {}
 
 replicaCount: 1
 
+## Security context to be added to push-gateway pods
+##
+securityContext:
+  fsGroup: 65534
+  runAsUser: 65534
+  runAsNonRoot: true
+
 ## Affinity for pod assignment
 ## Ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#affinity-and-anti-affinity
 affinity: {}
@@ -115,6 +130,50 @@ serviceMonitor:
 # If not set then a PodDisruptionBudget will not be created
 podDisruptionBudget: {}
 
+persistentVolume:
+  ## If true, pushgateway will create/use a Persistent Volume Claim
+  ## If false, use emptyDir
+  ##
+  enabled: false
+
+  ## pushgateway data Persistent Volume access modes
+  ## Must match those of existing PV or dynamic provisioner
+  ## Ref: http://kubernetes.io/docs/user-guide/persistent-volumes/
+  ##
+  accessModes:
+    - ReadWriteOnce
+
+  ## pushgateway data Persistent Volume Claim annotations
+  ##
+  annotations: {}
+
+  ## pushgateway data Persistent Volume existing claim name
+  ## Requires pushgateway.persistentVolume.enabled: true
+  ## If defined, PVC must be created manually before volume will be bound
+  existingClaim: ""
+
+  ## pushgateway data Persistent Volume mount root path
+  ##
+  mountPath: /data
+
+  ## pushgateway data Persistent Volume size
+  ##
+  size: 2Gi
+
+  ## pushgateway data Persistent Volume Storage Class
+  ## If defined, storageClassName: <storageClass>
+  ## If set to "-", storageClassName: "", which disables dynamic provisioning
+  ## If undefined (the default) or set to null, no storageClassName spec is
+  ##   set, choosing the default provisioner.  (gp2 on AWS, standard on
+  ##   GKE, AWS & OpenStack)
+  ##
+  # storageClass: "-"
+
+  ## Subdirectory of pushgateway data Persistent Volume to mount
+  ## Useful if the volume's root directory is not empty
+  ##
+  subPath: ""
+
 # Configuration for clusters with restrictive network policies in place:
 # - allowAll allows access to the PushGateway from any namespace
 # - customSelector is a list of pod/namespaceSelectors to allow access from
```

## 1.2.11 

**Release date:** 2020-01-09

![AppVersion: 1.0.1](https://img.shields.io/static/v1?label=AppVersion&message=1.0.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-pushgateway] Allow to specify the nodePort when service type is NodePort (#19936) 

### Default value changes

```diff
# No changes in this release
```

## 1.2.10 

**Release date:** 2019-12-21

![AppVersion: 1.0.1](https://img.shields.io/static/v1?label=AppVersion&message=1.0.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-pushgateway] Upgrade push-gateway (#19732) 

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index 5f204c2..a714a39 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -3,7 +3,7 @@
 # Declare variables to be passed into your templates.
 image:
   repository: prom/pushgateway
-  tag: v1.0.0
+  tag: v1.0.1
   pullPolicy: IfNotPresent
 
 service:
```

## 1.2.9 

**Release date:** 2019-12-20

![AppVersion: 1.0.0](https://img.shields.io/static/v1?label=AppVersion&message=1.0.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-pushgateway] Fix missleading documentation about `metrics.enabled` in values documentation (#19730) 

### Default value changes

```diff
# No changes in this release
```

## 1.2.8 

**Release date:** 2019-12-18

![AppVersion: 1.0.0](https://img.shields.io/static/v1?label=AppVersion&message=1.0.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-pushgateway] fix default values for podDisruptionBudget and networkPolicy (#19680) 

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index c975263..5f204c2 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -113,13 +113,13 @@ serviceMonitor:
 
 # The values to set in the PodDisruptionBudget spec (minAvailable/maxUnavailable)
 # If not set then a PodDisruptionBudget will not be created
-podDisruptionBudget:
+podDisruptionBudget: {}
 
 # Configuration for clusters with restrictive network policies in place:
 # - allowAll allows access to the PushGateway from any namespace
 # - customSelector is a list of pod/namespaceSelectors to allow access from
 # These options are mutually exclusive and the latter will take precedence.
-networkPolicy:
+networkPolicy: {}
   # allowAll: true
   # customSelectors:
   #   - namespaceSelector:
```

## 1.2.7 

**Release date:** 2019-12-16

![AppVersion: 1.0.0](https://img.shields.io/static/v1?label=AppVersion&message=1.0.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Update readiness liveness probe path for prometheus pushgateway (#19580) 

### Default value changes

```diff
# No changes in this release
```

## 1.2.6 

**Release date:** 2019-11-22

![AppVersion: 1.0.0](https://img.shields.io/static/v1?label=AppVersion&message=1.0.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-pushgateway] support for networkpolicies (#19057) 

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index 8a60aa0..c975263 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -114,3 +114,17 @@ serviceMonitor:
 # The values to set in the PodDisruptionBudget spec (minAvailable/maxUnavailable)
 # If not set then a PodDisruptionBudget will not be created
 podDisruptionBudget:
+
+# Configuration for clusters with restrictive network policies in place:
+# - allowAll allows access to the PushGateway from any namespace
+# - customSelector is a list of pod/namespaceSelectors to allow access from
+# These options are mutually exclusive and the latter will take precedence.
+networkPolicy:
+  # allowAll: true
+  # customSelectors:
+  #   - namespaceSelector:
+  #       matchLabels:
+  #         type: admin
+  #   - podSelector:
+  #       matchLabels:
+  #         app: myapp
```

## 1.2.5 

**Release date:** 2019-11-15

![AppVersion: 1.0.0](https://img.shields.io/static/v1?label=AppVersion&message=1.0.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-pushgateway] Add servicemonitor configuration options (#18907) 

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index 26d3ceb..8a60aa0 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -96,9 +96,13 @@ affinity: {}
 serviceMonitor:
   enabled: false
   namespace: monitoring
-  # fallback to the prometheus default unless specified
+
+  # Fallback to the prometheus default unless specified
   # interval: 10s
 
+  # Fallback to the prometheus default unless specified
+  # scrapeTimeout: 30s
+
   ## Used to pass Labels that are used by the Prometheus installed in your cluster to select Service Monitors to work with
   ## ref: https://github.com/coreos/prometheus-operator/blob/master/Documentation/api.md#prometheusspec
   additionalLabels: {}
```

## 1.2.4 

**Release date:** 2019-11-15

![AppVersion: 1.0.0](https://img.shields.io/static/v1?label=AppVersion&message=1.0.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-pushgateway] Adjust servicemonitor labels (#18877) 

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index bc19c11..26d3ceb 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -98,11 +98,11 @@ serviceMonitor:
   namespace: monitoring
   # fallback to the prometheus default unless specified
   # interval: 10s
-  ## Defaults to what's used if you follow CoreOS [Prometheus Install Instructions](https://github.com/helm/charts/tree/master/stable/prometheus-operator#tldr)
-  ## [Prometheus Selector Label](https://github.com/helm/charts/tree/master/stable/prometheus-operator#prometheus-operator-1)
-  ## [Kube Prometheus Selector Label](https://github.com/helm/charts/tree/master/stable/prometheus-operator#exporters)
-  selector:
-    prometheus: kube-prometheus
+
+  ## Used to pass Labels that are used by the Prometheus installed in your cluster to select Service Monitors to work with
+  ## ref: https://github.com/coreos/prometheus-operator/blob/master/Documentation/api.md#prometheusspec
+  additionalLabels: {}
+
   # Retain the job and instance labels of the metrics pushed to the Pushgateway
   # [Scraping Pushgateway](https://github.com/prometheus/pushgateway#configure-the-pushgateway-as-a-target-to-scrape)
   honorLabels: true
```

## 1.2.3 

**Release date:** 2019-11-14

![AppVersion: 1.0.0](https://img.shields.io/static/v1?label=AppVersion&message=1.0.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-pushgateway] Use fullname for the servicemonitor name (#18872) 

### Default value changes

```diff
# No changes in this release
```

## 1.2.2 

**Release date:** 2019-11-08

![AppVersion: 1.0.0](https://img.shields.io/static/v1?label=AppVersion&message=1.0.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-pushgateway] Compat k8s 1.16 (#18598) 

### Default value changes

```diff
# No changes in this release
```

## 1.2.1 

**Release date:** 2019-11-07

![AppVersion: 1.0.0](https://img.shields.io/static/v1?label=AppVersion&message=1.0.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-pushgateway] Follow the Call for maintainers (#18662) 

### Default value changes

```diff
# No changes in this release
```

## 1.2.0 

**Release date:** 2019-11-05

![AppVersion: 1.0.0](https://img.shields.io/static/v1?label=AppVersion&message=1.0.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Fix deprecated apiVersion for Deployment (#18586) 

### Default value changes

```diff
# No changes in this release
```

## 1.1.1 

**Release date:** 2019-10-30

![AppVersion: 1.0.0](https://img.shields.io/static/v1?label=AppVersion&message=1.0.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-pushgateway] Fix podDisruptionBudget template (#18187) 

### Default value changes

```diff
# No changes in this release
```

## 1.1.0 

**Release date:** 2019-10-21

![AppVersion: 1.0.0](https://img.shields.io/static/v1?label=AppVersion&message=1.0.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-pushgateway] Update default version to v1.0.0 (#18091) 

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index b08f2a7..bc19c11 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -3,7 +3,7 @@
 # Declare variables to be passed into your templates.
 image:
   repository: prom/pushgateway
-  tag: v0.9.1
+  tag: v1.0.0
   pullPolicy: IfNotPresent
 
 service:
```

## 1.0.1 

**Release date:** 2019-08-03

![AppVersion: 0.9.1](https://img.shields.io/static/v1?label=AppVersion&message=0.9.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* upgrade pushgateway (#16061) 

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index adb9080..b08f2a7 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -3,7 +3,7 @@
 # Declare variables to be passed into your templates.
 image:
   repository: prom/pushgateway
-  tag: v0.9.0
+  tag: v0.9.1
   pullPolicy: IfNotPresent
 
 service:
```

## 1.0.0 

**Release date:** 2019-07-29

![AppVersion: 0.9.0](https://img.shields.io/static/v1?label=AppVersion&message=0.9.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-pushgateway] upgrade to latest release, set chart to v1 (#15912) 

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index 2eaa38e..adb9080 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -3,7 +3,7 @@
 # Declare variables to be passed into your templates.
 image:
   repository: prom/pushgateway
-  tag: v0.8.0
+  tag: v0.9.0
   pullPolicy: IfNotPresent
 
 service:
```

## 0.4.1 

**Release date:** 2019-07-03

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-pushgateway] add optional PodDisruptionBudget (#13293) 

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index 77a1921..2eaa38e 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -106,3 +106,7 @@ serviceMonitor:
   # Retain the job and instance labels of the metrics pushed to the Pushgateway
   # [Scraping Pushgateway](https://github.com/prometheus/pushgateway#configure-the-pushgateway-as-a-target-to-scrape)
   honorLabels: true
+
+# The values to set in the PodDisruptionBudget spec (minAvailable/maxUnavailable)
+# If not set then a PodDisruptionBudget will not be created
+podDisruptionBudget:
```

## 0.4.0 

**Release date:** 2019-04-14

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* upgrade promteheus pushgateway (#13038) 

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index ad5a4ce..77a1921 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -3,7 +3,7 @@
 # Declare variables to be passed into your templates.
 image:
   repository: prom/pushgateway
-  tag: v0.6.0
+  tag: v0.8.0
   pullPolicy: IfNotPresent
 
 service:
```

## 0.3.1 

**Release date:** 2019-03-25

![AppVersion: 0.6.0](https://img.shields.io/static/v1?label=AppVersion&message=0.6.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Allow for customization of env vars on prometheus-pushgateway (#12196) 

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index cc06ed0..ad5a4ce 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -26,6 +26,9 @@ serviceAccountLabels: {}
 # Optional additional arguments
 extraArgs: []
 
+# Optional additional environment variables
+extraVars: []
+
 resources: {}
   # We usually recommend not to specify default resources and to leave this as a conscious
   # choice for the user. This also increases chances charts run on environments with little
```

## 0.3.0 

**Release date:** 2019-01-17

![AppVersion: 0.6.0](https://img.shields.io/static/v1?label=AppVersion&message=0.6.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Fix pushgateway labels honoring (#10728) 

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index 9900bda..cc06ed0 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -100,3 +100,6 @@ serviceMonitor:
   ## [Kube Prometheus Selector Label](https://github.com/helm/charts/tree/master/stable/prometheus-operator#exporters)
   selector:
     prometheus: kube-prometheus
+  # Retain the job and instance labels of the metrics pushed to the Pushgateway
+  # [Scraping Pushgateway](https://github.com/prometheus/pushgateway#configure-the-pushgateway-as-a-target-to-scrape)
+  honorLabels: true
```

## 0.2.0 

**Release date:** 2018-11-19

![AppVersion: 0.6.0](https://img.shields.io/static/v1?label=AppVersion&message=0.6.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* add prometheus servicemonitor support to pushgateway (#9385) 

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index 5dd7a07..9900bda 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -88,3 +88,15 @@ replicaCount: 1
 ## Affinity for pod assignment
 ## Ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#affinity-and-anti-affinity
 affinity: {}
+
+# Enable this if you're using https://github.com/coreos/prometheus-operator
+serviceMonitor:
+  enabled: false
+  namespace: monitoring
+  # fallback to the prometheus default unless specified
+  # interval: 10s
+  ## Defaults to what's used if you follow CoreOS [Prometheus Install Instructions](https://github.com/helm/charts/tree/master/stable/prometheus-operator#tldr)
+  ## [Prometheus Selector Label](https://github.com/helm/charts/tree/master/stable/prometheus-operator#prometheus-operator-1)
+  ## [Kube Prometheus Selector Label](https://github.com/helm/charts/tree/master/stable/prometheus-operator#exporters)
+  selector:
+    prometheus: kube-prometheus
```

## 0.1.6 

**Release date:** 2018-11-06

![AppVersion: 0.6.0](https://img.shields.io/static/v1?label=AppVersion&message=0.6.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Add the serviceAccountLabels variable i forgot to document in #8567 (#8913) 

### Default value changes

```diff
# No changes in this release
```

## 0.1.5 

**Release date:** 2018-10-30

![AppVersion: 0.6.0](https://img.shields.io/static/v1?label=AppVersion&message=0.6.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Add pod and service labels support (#8567) 

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index c6abed5..5dd7a07 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -14,6 +14,15 @@ service:
 # Optional pod annotations
 podAnnotations: {}
 
+# Optional pod labels
+podLabels: {}
+
+# Optional service labels
+serviceLabels: {}
+
+# Optional serviceAccount labels
+serviceAccountLabels: {}
+
 # Optional additional arguments
 extraArgs: []
 
```

## 0.1.4 

**Release date:** 2018-10-30

![AppVersion: 0.6.0](https://img.shields.io/static/v1?label=AppVersion&message=0.6.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Update appVersion and image tag (#8864) 

### Default value changes

```diff
# No changes in this release
```

## 0.1.3 

**Release date:** 2018-10-26

![AppVersion: 0.6.0](https://img.shields.io/static/v1?label=AppVersion&message=0.6.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Update appVersion and image tag (#8794) 

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index 90e0821..c6abed5 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -3,7 +3,7 @@
 # Declare variables to be passed into your templates.
 image:
   repository: prom/pushgateway
-  tag: v0.4.0
+  tag: v0.6.0
   pullPolicy: IfNotPresent
 
 service:
```

## 0.1.2 

**Release date:** 2018-05-09

![AppVersion: 0.4.0](https://img.shields.io/static/v1?label=AppVersion&message=0.4.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* pod annotations on prometheus-pushgateway pod (#5469) 

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index 70a2bc4..90e0821 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -11,6 +11,9 @@ service:
   port: 9091
   targetPort: 9091
 
+# Optional pod annotations
+podAnnotations: {}
+
 # Optional additional arguments
 extraArgs: []
 
```

## 0.1.1 

**Release date:** 2018-04-04

![AppVersion: 0.4.0](https://img.shields.io/static/v1?label=AppVersion&message=0.4.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* review wrong service port in pushgateway readme (#4679) 

### Default value changes

```diff
# No changes in this release
```

## 0.1.0 

**Release date:** 2018-04-03

![AppVersion: 0.4.0](https://img.shields.io/static/v1?label=AppVersion&message=0.4.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Add prometheus pushgateway (#4620) 

### Default value changes

```diff
# Default values for prometheus-pushgateway.
# This is a YAML-formatted file.
# Declare variables to be passed into your templates.
image:
  repository: prom/pushgateway
  tag: v0.4.0
  pullPolicy: IfNotPresent

service:
  type: ClusterIP
  port: 9091
  targetPort: 9091

# Optional additional arguments
extraArgs: []

resources: {}
  # We usually recommend not to specify default resources and to leave this as a conscious
  # choice for the user. This also increases chances charts run on environments with little
  # resources, such as Minikube. If you do want to specify resources, uncomment the following
  # lines, adjust them as necessary, and remove the curly braces after 'resources:'.
  # limits:
  #   cpu: 200m
  #    memory: 50Mi
  # requests:
  #   cpu: 100m
  #   memory: 30Mi

serviceAccount:
  # Specifies whether a ServiceAccount should be created
  create: true
  # The name of the ServiceAccount to use.
  # If not set and create is true, a name is generated using the fullname template
  name:

## Configure ingress resource that allow you to access the
## pushgateway installation. Set up the URL
## ref: http://kubernetes.io/docs/user-guide/ingress/
##
ingress:
  ## Enable Ingress.
  ##
  enabled: false

    ## Annotations.
    ##
    # annotations:
    #   kubernetes.io/ingress.class: nginx
    #   kubernetes.io/tls-acme: 'true'

    ## Hostnames.
    ## Must be provided if Ingress is enabled.
    ##
    # hosts:
    #   - pushgateway.domain.com

    ## TLS configuration.
    ## Secrets must be manually created in the namespace.
    ##
    # tls:
    #   - secretName: pushgateway-tls
    #     hosts:
    #       - pushgateway.domain.com

tolerations: {}
  # - effect: NoSchedule
  #   operator: Exists

## Node labels for pushgateway pod assignment
## Ref: https://kubernetes.io/docs/user-guide/node-selection/
##
nodeSelector: {}

replicaCount: 1

## Affinity for pod assignment
## Ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#affinity-and-anti-affinity
affinity: {}
```

---
Autogenerated from Helm Chart and git history using [helm-changelog](https://github.com/mogensen/helm-changelog)
