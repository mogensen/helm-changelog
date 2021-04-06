# Change Log

## 1.16.2 

**Release date:** 2021-03-18

![AppVersion: 1.1.2](https://img.shields.io/static/v1?label=AppVersion&message=1.1.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* prometheus-node-exporter - use custom service accounts (#755) 

### Default value changes

```diff
# No changes in this release
```

## 1.16.1 

**Release date:** 2021-03-18

![AppVersion: 1.1.2](https://img.shields.io/static/v1?label=AppVersion&message=1.1.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-node-exporter] Allow to opt-out node exporter rootfs mount (#757) 

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 63ff136..93f8ef8 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -80,6 +80,10 @@ endpoints: []
 # Expose the service to the host network
 hostNetwork: true
 
+## If true, node-exporter pods mounts host / at /host/root
+##
+hostRootFsMount: true
+
 ## Assign a group of affinity scheduling rules
 ##
 affinity: {}
```

## 1.16.0 

**Release date:** 2021-03-18

![AppVersion: 1.1.2](https://img.shields.io/static/v1?label=AppVersion&message=1.1.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-node-exporter] AppVersion bump to 1.1.2 (#774) 

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 966e417..63ff136 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -3,7 +3,7 @@
 # Declare variables to be passed into your templates.
 image:
   repository: quay.io/prometheus/node-exporter
-  tag: v1.0.1
+  tag: v1.1.2
   pullPolicy: IfNotPresent
 
 service:
```

## 1.15.0 

**Release date:** 2021-02-10

![AppVersion: 1.0.1](https://img.shields.io/static/v1?label=AppVersion&message=1.0.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Ability for custom dnsConfig on prometheus-node-exporter (#644) 

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 4be3f9c..966e417 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -100,6 +100,18 @@ podAnnotations:
 # Extra labels to be added to node exporter pods
 podLabels: {}
 
+# Custom DNS configuration to be added to prometheus-node-exporter pods
+dnsConfig: {}
+# nameservers:
+#   - 1.2.3.4
+# searches:
+#   - ns1.svc.cluster-domain.example
+#   - my.dns.search.suffix
+# options:
+#   - name: ndots
+#     value: "2"
+#   - name: edns0
+
 ## Assign a nodeSelector if operating a hybrid cluster
 ##
 nodeSelector: {}
```

## 1.14.2 

**Release date:** 2021-02-03

![AppVersion: 1.0.1](https://img.shields.io/static/v1?label=AppVersion&message=1.0.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* add containerSecurityContext to node-exporter (#544) 

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 2b2205c..4be3f9c 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -60,6 +60,11 @@ securityContext:
   runAsNonRoot: true
   runAsUser: 65534
 
+containerSecurityContext: {}
+  # capabilities:
+  #   add:
+  #   - SYS_TIME
+
 rbac:
   ## If true, create & use RBAC resources
   ##
```

## 1.14.1 

**Release date:** 2021-02-01

![AppVersion: 1.0.1](https://img.shields.io/static/v1?label=AppVersion&message=1.0.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-node-exporter] Add bismarck as maintainer (#625) 

### Default value changes

```diff
# No changes in this release
```

## 1.14.0 

**Release date:** 2021-02-01

![AppVersion: 1.0.1](https://img.shields.io/static/v1?label=AppVersion&message=1.0.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-node-exporter] Fix very slow GKE cluster upgrades (#485) 

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 8ef7748..2b2205c 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -88,7 +88,9 @@ affinity: {}
 #                 - target-host-name
 
 # Annotations to be added to node exporter pods
-podAnnotations: {}
+podAnnotations:
+  # Fix for very slow GKE cluster upgrades
+  cluster-autoscaler.kubernetes.io/safe-to-evict: "true"
 
 # Extra labels to be added to node exporter pods
 podLabels: {}
```

## 1.13.0 

**Release date:** 2021-01-29

![AppVersion: 1.0.1](https://img.shields.io/static/v1?label=AppVersion&message=1.0.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Adding secret support for oauth psidecar (#486) 

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 7edd893..8ef7748 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -20,6 +20,9 @@ prometheus:
     enabled: false
     additionalLabels: {}
     namespace: ""
+    scheme: http
+    bearerTokenFile:
+    tlsConfig: {}
 
     relabelings: []
     scrapeTimeout: 10s
@@ -48,6 +51,7 @@ serviceAccount:
   # The name of the ServiceAccount to use.
   # If not set and create is true, a name is generated using the fullname template
   name:
+  annotations: {}
   imagePullSecrets: []
 
 securityContext:
@@ -122,7 +126,9 @@ extraHostVolumeMounts: []
 configmaps: []
 # - name: <configMapName>
 #   mountPath: <mountPath>
-
+secrets: []
+# - name: <secretName>
+#   mountPath: <mountPatch>
 ## Override the deployment namespace
 ##
 namespaceOverride: ""
```

## 1.12.0 

**Release date:** 2020-10-30

![AppVersion: 1.0.1](https://img.shields.io/static/v1?label=AppVersion&message=1.0.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-node-exporter] Add rootfs mount by default (#80) 

### Default value changes

```diff
# No changes in this release
```

## 1.11.2 

**Release date:** 2020-08-20

![AppVersion: 1.0.1](https://img.shields.io/static/v1?label=AppVersion&message=1.0.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Prep initial charts indexing (#14) 

### Default value changes

```diff
# No changes in this release
```

## 1.11.1 

**Release date:** 2020-07-15

![AppVersion: 1.0.1](https://img.shields.io/static/v1?label=AppVersion&message=1.0.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-node-exporter] Version bump to 1.0.1 (#23151) 

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 3e0acd0..7edd893 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -3,7 +3,7 @@
 # Declare variables to be passed into your templates.
 image:
   repository: quay.io/prometheus/node-exporter
-  tag: v1.0.0
+  tag: v1.0.1
   pullPolicy: IfNotPresent
 
 service:
```

## 1.11.0 

**Release date:** 2020-06-16

![AppVersion: 1.0.0](https://img.shields.io/static/v1?label=AppVersion&message=1.0.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Add relabelings config for node-exporter ServiceMonitor (#22764) 

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 5e4b2f2..3e0acd0 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -21,6 +21,7 @@ prometheus:
     additionalLabels: {}
     namespace: ""
 
+    relabelings: []
     scrapeTimeout: 10s
 
 ## Customize the updateStrategy if set
```

## 1.10.0 

**Release date:** 2020-05-31

![AppVersion: 1.0.0](https://img.shields.io/static/v1?label=AppVersion&message=1.0.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-node-exporter] Upgrade to 1.0.0 & set group to non-root (#22603) 

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 868bfcd..5e4b2f2 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -3,7 +3,7 @@
 # Declare variables to be passed into your templates.
 image:
   repository: quay.io/prometheus/node-exporter
-  tag: v0.18.1
+  tag: v1.0.0
   pullPolicy: IfNotPresent
 
 service:
@@ -50,6 +50,8 @@ serviceAccount:
   imagePullSecrets: []
 
 securityContext:
+  fsGroup: 65534
+  runAsGroup: 65534
   runAsNonRoot: true
   runAsUser: 65534
 
```

## 1.9.1 

**Release date:** 2020-03-19

![AppVersion: 0.18.1](https://img.shields.io/static/v1?label=AppVersion&message=0.18.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* prometheus node exporter: Configure exposing IP (#21490) 

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 1567c87..868bfcd 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -11,6 +11,7 @@ service:
   port: 9100
   targetPort: 9100
   nodePort:
+  listenOnAllInterfaces: true
   annotations:
     prometheus.io/scrape: "true"
 
```

## 1.9.0 

**Release date:** 2020-02-22

![AppVersion: 0.18.1](https://img.shields.io/static/v1?label=AppVersion&message=0.18.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-node-exporter] node-exporter add sidecar (#20952) 

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 1074662..1567c87 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -102,6 +102,7 @@ tolerations:
 ##
 extraArgs: []
 #   - --collector.diskstats.ignored-devices=^(ram|loop|fd|(h|s|v)d[a-z]|nvme\\d+n\\d+p)\\d+$
+#   - --collector.textfile.directory=/run/prometheus
 
 ## Additional mounts from the host
 ##
@@ -121,3 +122,16 @@ configmaps: []
 ## Override the deployment namespace
 ##
 namespaceOverride: ""
+
+## Additional containers for export metrics to text file
+##
+sidecars: []
+##  - name: nvidia-dcgm-exporter
+##    image: nvidia/dcgm-exporter:1.4.3
+
+## Volume for sidecar containers
+##
+sidecarVolumeMount: []
+##  - name: collector-textfiles
+##    mountPath: /run/prometheus
+##    readOnly: false
```

## 1.8.2 

**Release date:** 2020-02-19

![AppVersion: 0.18.1](https://img.shields.io/static/v1?label=AppVersion&message=0.18.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-node-exporter] Add updateStrategy to template (#20858) 

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index bdac990..1074662 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -22,6 +22,12 @@ prometheus:
 
     scrapeTimeout: 10s
 
+## Customize the updateStrategy if set
+updateStrategy:
+  type: RollingUpdate
+  rollingUpdate:
+    maxUnavailable: 1
+
 resources: {}
   # We usually recommend not to specify default resources and to leave this as a conscious
   # choice for the user. This also increases chances charts run on environments with little
```

## 1.8.1 

**Release date:** 2019-12-17

![AppVersion: 0.18.1](https://img.shields.io/static/v1?label=AppVersion&message=0.18.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-node-exporter] - add podAnnotations value (#17740) 

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index cf79367..bdac990 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -73,6 +73,12 @@ affinity: {}
 #               values:
 #                 - target-host-name
 
+# Annotations to be added to node exporter pods
+podAnnotations: {}
+
+# Extra labels to be added to node exporter pods
+podLabels: {}
+
 ## Assign a nodeSelector if operating a hybrid cluster
 ##
 nodeSelector: {}
```

## 1.8.0 

**Release date:** 2019-11-19

![AppVersion: 0.18.1](https://img.shields.io/static/v1?label=AppVersion&message=0.18.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Prometheus node exporter namespace override (#18982) 

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 900111c..cf79367 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -105,3 +105,7 @@ extraHostVolumeMounts: []
 configmaps: []
 # - name: <configMapName>
 #   mountPath: <mountPath>
+
+## Override the deployment namespace
+##
+namespaceOverride: ""
```

## 1.7.3 

**Release date:** 2019-11-04

![AppVersion: 0.18.1](https://img.shields.io/static/v1?label=AppVersion&message=0.18.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-node-exporter] Fix maps -> slices, bump version (#18527) 

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 9fa0686..900111c 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -3,7 +3,7 @@
 # Declare variables to be passed into your templates.
 image:
   repository: quay.io/prometheus/node-exporter
-  tag: v0.18.0
+  tag: v0.18.1
   pullPolicy: IfNotPresent
 
 service:
@@ -88,12 +88,12 @@ tolerations:
 
 ## Additional container arguments
 ##
-extraArgs: {}
+extraArgs: []
 #   - --collector.diskstats.ignored-devices=^(ram|loop|fd|(h|s|v)d[a-z]|nvme\\d+n\\d+p)\\d+$
 
 ## Additional mounts from the host
 ##
-extraHostVolumeMounts: {}
+extraHostVolumeMounts: []
 #  - name: <mountName>
 #    hostPath: <hostPath>
 #    mountPath: <mountPath>
@@ -102,6 +102,6 @@ extraHostVolumeMounts: {}
 
 ## Additional configmaps to be mounted.
 ##
-configmaps: {}
+configmaps: []
 # - name: <configMapName>
 #   mountPath: <mountPath>
```

## 1.7.2 

**Release date:** 2019-10-31

![AppVersion: 0.18.0](https://img.shields.io/static/v1?label=AppVersion&message=0.18.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-node-exporter] Add vsliouniaev as owner (#18461) 

### Default value changes

```diff
# No changes in this release
```

## 1.7.1 

**Release date:** 2019-10-28

![AppVersion: 0.18.0](https://img.shields.io/static/v1?label=AppVersion&message=0.18.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-node-exporter] update NOTES.txt to show correct port (#17931) 

### Default value changes

```diff
# No changes in this release
```

## 1.7.0 

**Release date:** 2019-10-10

![AppVersion: 0.18.0](https://img.shields.io/static/v1?label=AppVersion&message=0.18.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-node-exporter] Use latest API versions (#17785) 

### Default value changes

```diff
# No changes in this release
```

## 1.6.0 

**Release date:** 2019-07-28

![AppVersion: 0.18.0](https://img.shields.io/static/v1?label=AppVersion&message=0.18.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* allow mounting configmaps (#15356) 

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index b73b75c..9fa0686 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -99,3 +99,9 @@ extraHostVolumeMounts: {}
 #    mountPath: <mountPath>
 #    readOnly: true|false
 #    mountPropagation: None|HostToContainer|Bidirectional
+
+## Additional configmaps to be mounted.
+##
+configmaps: {}
+# - name: <configMapName>
+#   mountPath: <mountPath>
```

## 1.5.2 

**Release date:** 2019-07-26

![AppVersion: 0.18.0](https://img.shields.io/static/v1?label=AppVersion&message=0.18.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* removed errant space in commented line (#15661) 

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 2da6571..b73b75c 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -29,7 +29,7 @@ resources: {}
   # lines, adjust them as necessary, and remove the curly braces after 'resources:'.
   # limits:
   #   cpu: 200m
-  #    memory: 50Mi
+  #   memory: 50Mi
   # requests:
   #   cpu: 100m
   #   memory: 30Mi
```

## 1.5.1 

**Release date:** 2019-06-27

![AppVersion: 0.18.0](https://img.shields.io/static/v1?label=AppVersion&message=0.18.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-node-exporter] add scrapeTimeout for the service monitor (#14293) 

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 7ca723f..2da6571 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -20,6 +20,8 @@ prometheus:
     additionalLabels: {}
     namespace: ""
 
+    scrapeTimeout: 10s
+
 resources: {}
   # We usually recommend not to specify default resources and to leave this as a conscious
   # choice for the user. This also increases chances charts run on environments with little
```

## 1.5.0 

**Release date:** 2019-05-20

![AppVersion: 0.18.0](https://img.shields.io/static/v1?label=AppVersion&message=0.18.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-node-exporter] Upgrade to 0.18.0 (#13739) 

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 09261c4..7ca723f 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -3,7 +3,7 @@
 # Declare variables to be passed into your templates.
 image:
   repository: quay.io/prometheus/node-exporter
-  tag: v0.17.0
+  tag: v0.18.0
   pullPolicy: IfNotPresent
 
 service:
```

## 1.4.2 

**Release date:** 2019-04-03

![AppVersion: 0.17.0](https://img.shields.io/static/v1?label=AppVersion&message=0.17.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-node-exporter] configurable host network (#12747) 

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 14bcfc5..09261c4 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -56,6 +56,9 @@ rbac:
 # their addresses here
 endpoints: []
 
+# Expose the service to the host network
+hostNetwork: true
+
 ## Assign a group of affinity scheduling rules
 ##
 affinity: {}
```

## 1.3.2 

**Release date:** 2019-03-07

![AppVersion: 0.17.0](https://img.shields.io/static/v1?label=AppVersion&message=0.17.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* add prometheus service monitor to prometheus node exporter (#12004) 

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index fcd7acb..14bcfc5 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -14,6 +14,12 @@ service:
   annotations:
     prometheus.io/scrape: "true"
 
+prometheus:
+  monitor:
+    enabled: false
+    additionalLabels: {}
+    namespace: ""
+
 resources: {}
   # We usually recommend not to specify default resources and to leave this as a conscious
   # choice for the user. This also increases chances charts run on environments with little
```

## 1.3.1 

**Release date:** 2019-03-08

![AppVersion: 0.17.0](https://img.shields.io/static/v1?label=AppVersion&message=0.17.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-node-exporter] Fix imagePullSecrets value path (#11222) 

### Default value changes

```diff
# No changes in this release
```

## 1.3.0 

**Release date:** 2019-02-18

![AppVersion: 0.17.0](https://img.shields.io/static/v1?label=AppVersion&message=0.17.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-node-exporter] Add options to change volume mount propagation  (#11194) 

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 5f4298c..fcd7acb 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -87,3 +87,4 @@ extraHostVolumeMounts: {}
 #    hostPath: <hostPath>
 #    mountPath: <mountPath>
 #    readOnly: true|false
+#    mountPropagation: None|HostToContainer|Bidirectional
```

## 1.2.0 

**Release date:** 2019-01-31

![AppVersion: 0.17.0](https://img.shields.io/static/v1?label=AppVersion&message=0.17.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-node-exporter] Add nodeSelector to DaemonSet (#10952) 

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 6df2436..5f4298c 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -50,6 +50,24 @@ rbac:
 # their addresses here
 endpoints: []
 
+## Assign a group of affinity scheduling rules
+##
+affinity: {}
+#   nodeAffinity:
+#     requiredDuringSchedulingIgnoredDuringExecution:
+#       nodeSelectorTerms:
+#         - matchFields:
+#             - key: metadata.name
+#               operator: In
+#               values:
+#                 - target-host-name
+
+## Assign a nodeSelector if operating a hybrid cluster
+##
+nodeSelector: {}
+#   beta.kubernetes.io/arch: amd64
+#   beta.kubernetes.io/os: linux
+
 tolerations:
   - effect: NoSchedule
     operator: Exists
```

## 1.1.0 

**Release date:** 2018-12-31

![AppVersion: 0.17.0](https://img.shields.io/static/v1?label=AppVersion&message=0.17.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-node-exporter] Node exporter endpoints (#10243) 

### Default value changes

```diff
# No changes in this release
```

## 1.0.1 

**Release date:** 2018-12-20

![AppVersion: 0.17.0](https://img.shields.io/static/v1?label=AppVersion&message=0.17.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-node-exporter] - Bump app version to 0.17.0 ADDENDUM (#10125) 

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 3169424..6df2436 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -3,7 +3,7 @@
 # Declare variables to be passed into your templates.
 image:
   repository: quay.io/prometheus/node-exporter
-  tag: v0.16.0
+  tag: v0.17.0
   pullPolicy: IfNotPresent
 
 service:
```

## 1.0.0 

**Release date:** 2018-12-17

![AppVersion: 0.17.0](https://img.shields.io/static/v1?label=AppVersion&message=0.17.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-node-exporter]Add support for nodePort (#9754) 

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 7d3fc76..3169424 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -10,6 +10,7 @@ service:
   type: ClusterIP
   port: 9100
   targetPort: 9100
+  nodePort:
   annotations:
     prometheus.io/scrape: "true"
 
```

## 0.6.2 

**Release date:** 2018-12-17

![AppVersion: 0.17.0](https://img.shields.io/static/v1?label=AppVersion&message=0.17.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-node-exporter] - Bump app version to 0.17.0 (#10058) 

### Default value changes

```diff
# No changes in this release
```

## 0.6.1 

**Release date:** 2018-12-03

![AppVersion: 0.16.0](https://img.shields.io/static/v1?label=AppVersion&message=0.16.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-node-exporter] add securityContext value (#9618) 

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 38a7059..7d3fc76 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -33,6 +33,10 @@ serviceAccount:
   name:
   imagePullSecrets: []
 
+securityContext:
+  runAsNonRoot: true
+  runAsUser: 65534
+
 rbac:
   ## If true, create & use RBAC resources
   ##
```

## 0.6.0 

**Release date:** 2018-11-22

![AppVersion: 0.16.0](https://img.shields.io/static/v1?label=AppVersion&message=0.16.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Add priorityClassName option to values (#9483) 

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 6fac1b1..38a7059 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -49,6 +49,9 @@ tolerations:
   - effect: NoSchedule
     operator: Exists
 
+## Assign a PriorityClassName to pods if set
+# priorityClassName: ""
+
 ## Additional container arguments
 ##
 extraArgs: {}
```

## 0.5.1 

**Release date:** 2018-11-05

![AppVersion: 0.16.0](https://img.shields.io/static/v1?label=AppVersion&message=0.16.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Fix typo causing RBAC resources not being created (#8991) 

### Default value changes

```diff
# No changes in this release
```

## 0.5.0 

**Release date:** 2018-10-04

![AppVersion: 0.16.0](https://img.shields.io/static/v1?label=AppVersion&message=0.16.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-node-exporter] allow setting additional host mounts (#8117) 

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index effe0b4..6fac1b1 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -53,3 +53,11 @@ tolerations:
 ##
 extraArgs: {}
 #   - --collector.diskstats.ignored-devices=^(ram|loop|fd|(h|s|v)d[a-z]|nvme\\d+n\\d+p)\\d+$
+
+## Additional mounts from the host
+##
+extraHostVolumeMounts: {}
+#  - name: <mountName>
+#    hostPath: <hostPath>
+#    mountPath: <mountPath>
+#    readOnly: true|false
```

## 0.4.0 

**Release date:** 2018-10-02

![AppVersion: 0.16.0](https://img.shields.io/static/v1?label=AppVersion&message=0.16.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* added service annotations with the default value (#8086) 

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 18a6693..effe0b4 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -3,13 +3,15 @@
 # Declare variables to be passed into your templates.
 image:
   repository: quay.io/prometheus/node-exporter
-  tag: v0.15.2
+  tag: v0.16.0
   pullPolicy: IfNotPresent
 
 service:
   type: ClusterIP
   port: 9100
   targetPort: 9100
+  annotations:
+    prometheus.io/scrape: "true"
 
 resources: {}
   # We usually recommend not to specify default resources and to leave this as a conscious
```

## 0.3.0 

**Release date:** 2018-07-26

![AppVersion: 0.15.2](https://img.shields.io/static/v1?label=AppVersion&message=0.15.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-node-exporter] Add ability to customize labels and endpoints, Pod security policy (#6742) 

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 30db7c5..18a6693 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -29,6 +29,19 @@ serviceAccount:
   # The name of the ServiceAccount to use.
   # If not set and create is true, a name is generated using the fullname template
   name:
+  imagePullSecrets: []
+
+rbac:
+  ## If true, create & use RBAC resources
+  ##
+  create: true
+  ## If true, create & use Pod Security Policy resources
+  ## https://kubernetes.io/docs/concepts/policy/pod-security-policy/
+  pspEnabled: true
+
+# for deployments that have node_exporter deployed outside of the cluster, list
+# their addresses here
+endpoints: []
 
 tolerations:
   - effect: NoSchedule
@@ -36,5 +49,5 @@ tolerations:
 
 ## Additional container arguments
 ##
-# extraArgs:
+extraArgs: {}
 #   - --collector.diskstats.ignored-devices=^(ram|loop|fd|(h|s|v)d[a-z]|nvme\\d+n\\d+p)\\d+$
```

## 0.2.0 

**Release date:** 2018-05-23

![AppVersion: 0.15.2](https://img.shields.io/static/v1?label=AppVersion&message=0.15.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [Node-exporter] Add extra flags arguments and rollingUpdate strategy (#5699) 

### Default value changes

```diff
diff --git a/charts/prometheus-node-exporter/values.yaml b/charts/prometheus-node-exporter/values.yaml
index 0a63536..30db7c5 100644
--- a/charts/prometheus-node-exporter/values.yaml
+++ b/charts/prometheus-node-exporter/values.yaml
@@ -33,3 +33,8 @@ serviceAccount:
 tolerations:
   - effect: NoSchedule
     operator: Exists
+
+## Additional container arguments
+##
+# extraArgs:
+#   - --collector.diskstats.ignored-devices=^(ram|loop|fd|(h|s|v)d[a-z]|nvme\\d+n\\d+p)\\d+$
```

## 0.1.2 

**Release date:** 2018-04-27

![AppVersion: 0.15.2](https://img.shields.io/static/v1?label=AppVersion&message=0.15.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/promehteus-node-exporter] Fix health checks (#5295) 

### Default value changes

```diff
# No changes in this release
```

## 0.1.1 

**Release date:** 2018-04-03

![AppVersion: 0.15.2](https://img.shields.io/static/v1?label=AppVersion&message=0.15.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* stable/prometheus-node-exporter/README.md (#4550) 

### Default value changes

```diff
# No changes in this release
```

## 0.1.0 

**Release date:** 2018-03-30

![AppVersion: 0.15.2](https://img.shields.io/static/v1?label=AppVersion&message=0.15.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* add prometheus node exporter chart (#4260) 

### Default value changes

```diff
# Default values for prometheus-node-exporter.
# This is a YAML-formatted file.
# Declare variables to be passed into your templates.
image:
  repository: quay.io/prometheus/node-exporter
  tag: v0.15.2
  pullPolicy: IfNotPresent

service:
  type: ClusterIP
  port: 9100
  targetPort: 9100

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

tolerations:
  - effect: NoSchedule
    operator: Exists
```

---
Autogenerated from Helm Chart and git history using [helm-changelog](https://github.com/mogensen/helm-changelog)
