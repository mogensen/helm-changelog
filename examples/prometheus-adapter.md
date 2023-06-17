# Change Log

## 4.2.0

**Release date:** 2023-04-23

![AppVersion: v0.10.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.10.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-adapter] Add support for configuring probes (#3273)

### Default value changes

```diff
diff --git a/charts/prometheus-adapter/values.yaml b/charts/prometheus-adapter/values.yaml
index 2eb1a037..5058838a 100644
--- a/charts/prometheus-adapter/values.yaml
+++ b/charts/prometheus-adapter/values.yaml
@@ -93,6 +93,25 @@ resources: {}
   #   cpu: 100m
   #   memory: 128Mi
 
+# Configure liveness probe
+# https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/pod-v1/#Probe
+livenessProbe:
+  httpGet:
+    path: /healthz
+    port: https
+    scheme: HTTPS
+  initialDelaySeconds: 30
+  timeoutSeconds: 5
+
+# Configure readiness probe
+readinessProbe:
+  httpGet:
+    path: /healthz
+    port: https
+    scheme: HTTPS
+  initialDelaySeconds: 30
+  timeoutSeconds: 5
+
 rules:
   default: true
 

```

## 4.1.1

**Release date:** 2023-01-26

![AppVersion: v0.10.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.10.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-adapter]: fix usage of with for deploymentAnnotations (#2954)

### Default value changes

```diff
# No changes in this release
```

## 4.1.0

**Release date:** 2023-01-26

![AppVersion: v0.10.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.10.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-adapter]: allow setting annotations on the deployment (#2951)

### Default value changes

```diff
diff --git a/charts/prometheus-adapter/values.yaml b/charts/prometheus-adapter/values.yaml
index e9d0714d..2eb1a037 100644
--- a/charts/prometheus-adapter/values.yaml
+++ b/charts/prometheus-adapter/values.yaml
@@ -217,6 +217,9 @@ podLabels: {}
 # Annotations added to the pod
 podAnnotations: {}
 
+# Annotations added to the deployment
+deploymentAnnotations: {}
+
 hostNetwork:
   # Specifies if prometheus-adapter should be started in hostNetwork mode.
   #

```

## 4.0.2

**Release date:** 2023-01-21

![AppVersion: v0.10.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.10.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-adapter] Update PromQL queries (#2827)

### Default value changes

```diff
diff --git a/charts/prometheus-adapter/values.yaml b/charts/prometheus-adapter/values.yaml
index e48ca0a5..e9d0714d 100644
--- a/charts/prometheus-adapter/values.yaml
+++ b/charts/prometheus-adapter/values.yaml
@@ -120,8 +120,14 @@ rules:
 
   # resource:
   #   cpu:
-  #     containerQuery: sum(rate(container_cpu_usage_seconds_total{<<.LabelMatchers>>, container!=""}[3m])) by (<<.GroupBy>>)
-  #     nodeQuery: sum(rate(container_cpu_usage_seconds_total{<<.LabelMatchers>>, id='/'}[3m])) by (<<.GroupBy>>)
+  #     containerQuery: |
+  #       sum by (<<.GroupBy>>) (
+  #         rate(container_cpu_usage_seconds_total{container!="",<<.LabelMatchers>>}[3m])
+  #       )
+  #     nodeQuery: |
+  #       sum  by (<<.GroupBy>>) (
+  #         rate(node_cpu_seconds_total{mode!="idle",mode!="iowait",mode!="steal",<<.LabelMatchers>>}[3m])
+  #       )
   #     resources:
   #       overrides:
   #         node:
@@ -132,8 +138,16 @@ rules:
   #           resource: pod
   #     containerLabel: container
   #   memory:
-  #     containerQuery: sum(container_memory_working_set_bytes{<<.LabelMatchers>>, container!=""}) by (<<.GroupBy>>)
-  #     nodeQuery: sum(container_memory_working_set_bytes{<<.LabelMatchers>>,id='/'}) by (<<.GroupBy>>)
+  #     containerQuery: |
+  #       sum by (<<.GroupBy>>) (
+  #         avg_over_time(container_memory_working_set_bytes{container!="",<<.LabelMatchers>>}[3m])
+  #       )
+  #     nodeQuery: |
+  #       sum by (<<.GroupBy>>) (
+  #         avg_over_time(node_memory_MemTotal_bytes{<<.LabelMatchers>>}[3m])
+  #         -
+  #         avg_over_time(node_memory_MemAvailable_bytes{<<.LabelMatchers>>}[3m])
+  #       )
   #     resources:
   #       overrides:
   #         node:

```

## 4.0.1

**Release date:** 2023-01-17

![AppVersion: v0.10.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.10.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* use registry.k8s.io (#2926)

### Default value changes

```diff
diff --git a/charts/prometheus-adapter/values.yaml b/charts/prometheus-adapter/values.yaml
index a8141114..e48ca0a5 100644
--- a/charts/prometheus-adapter/values.yaml
+++ b/charts/prometheus-adapter/values.yaml
@@ -3,7 +3,7 @@ affinity: {}
 topologySpreadConstraints: []
 
 image:
-  repository: k8s.gcr.io/prometheus-adapter/prometheus-adapter
+  repository: registry.k8s.io/prometheus-adapter/prometheus-adapter
   tag: v0.10.0
   pullPolicy: IfNotPresent
 

```

## 4.0.0

**Release date:** 2023-01-10

![AppVersion: v0.10.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.10.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-adapter] Make securityContext user-settable (#2897)

### Default value changes

```diff
diff --git a/charts/prometheus-adapter/values.yaml b/charts/prometheus-adapter/values.yaml
index 51a75136..a8141114 100644
--- a/charts/prometheus-adapter/values.yaml
+++ b/charts/prometheus-adapter/values.yaml
@@ -13,9 +13,6 @@ metricsRelistInterval: 1m
 
 listenPort: 6443
 
-# User to run adapter container as
-runAsUser: 10001
-
 nodeSelector: {}
 
 priorityClassName: ""
@@ -45,6 +42,18 @@ replicas: 1
 podSecurityContext:
   fsGroup: 10001
 
+# SecurityContext of the container
+# ref. https://kubernetes.io/docs/tasks/configure-pod-container/security-context
+securityContext:
+  allowPrivilegeEscalation: false
+  capabilities:
+    drop: ["all"]
+  readOnlyRootFilesystem: true
+  runAsNonRoot: true
+  runAsUser: 10001
+  seccompProfile:
+    type: RuntimeDefault
+
 rbac:
   # Specifies whether RBAC resources should be created
   create: true

```

## 3.5.0

**Release date:** 2022-12-23

![AppVersion: v0.10.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.10.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-adapter] Support environment variables (#2842) (#2845)

### Default value changes

```diff
diff --git a/charts/prometheus-adapter/values.yaml b/charts/prometheus-adapter/values.yaml
index a1c61147..51a75136 100644
--- a/charts/prometheus-adapter/values.yaml
+++ b/charts/prometheus-adapter/values.yaml
@@ -151,6 +151,17 @@ tls:
   certificate: |-
     # Public key of the APIService
 
+# Set environment variables from secrets, configmaps or by setting them as name/value
+env: []
+  # - name: TMP_DIR
+  #   value: /tmp
+  # - name: PASSWORD
+  #   valueFrom:
+  #     secretKeyRef:
+  #       name: mysecret
+  #       key: password
+  #       optional: false
+
 # Any extra arguments
 extraArguments: []
   # - --tls-private-key-file=/etc/tls/tls.key

```

## 3.4.2

**Release date:** 2022-11-08

![AppVersion: v0.10.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.10.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-adapter] Prevent helm diff on insecureSkipTLSVerify  (#2637)

### Default value changes

```diff
# No changes in this release
```

## 3.4.1

**Release date:** 2022-10-16

![AppVersion: v0.10.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.10.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-adapter] Fix PSP deprecation after k8s 1.25+ (#2565)

### Default value changes

```diff
# No changes in this release
```

## 3.4.0

**Release date:** 2022-08-13

![AppVersion: v0.10.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.10.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-adapter] Update To v0.10.0 (#2366)

### Default value changes

```diff
diff --git a/charts/prometheus-adapter/values.yaml b/charts/prometheus-adapter/values.yaml
index c857ab62..a1c61147 100644
--- a/charts/prometheus-adapter/values.yaml
+++ b/charts/prometheus-adapter/values.yaml
@@ -4,7 +4,7 @@ topologySpreadConstraints: []
 
 image:
   repository: k8s.gcr.io/prometheus-adapter/prometheus-adapter
-  tag: v0.9.1
+  tag: v0.10.0
   pullPolicy: IfNotPresent
 
 logLevel: 4

```

## 3.3.1

**Release date:** 2022-05-27

![AppVersion: v0.9.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.9.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-adapter] conditionally use the policy/v1 apiversion for PDBs (#2084)

### Default value changes

```diff
# No changes in this release
```

## 3.3.0

**Release date:** 2022-05-19

![AppVersion: v0.9.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.9.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-adapter] Add runAsUser value to Deployment (#2075)

### Default value changes

```diff
diff --git a/charts/prometheus-adapter/values.yaml b/charts/prometheus-adapter/values.yaml
index da1466ed..c857ab62 100644
--- a/charts/prometheus-adapter/values.yaml
+++ b/charts/prometheus-adapter/values.yaml
@@ -13,6 +13,9 @@ metricsRelistInterval: 1m
 
 listenPort: 6443
 
+# User to run adapter container as
+runAsUser: 10001
+
 nodeSelector: {}
 
 priorityClassName: ""

```

## 3.2.2

**Release date:** 2022-04-27

![AppVersion: v0.9.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.9.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-adapter] fix serviceaccount annotations (#2014)

### Default value changes

```diff
# No changes in this release
```

## 3.2.1

**Release date:** 2022-04-13

![AppVersion: v0.9.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.9.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* remove helm 2 commands from readme / cleanup values.yaml (#1975)

### Default value changes

```diff
diff --git a/charts/prometheus-adapter/values.yaml b/charts/prometheus-adapter/values.yaml
index 3315b0c4..da1466ed 100644
--- a/charts/prometheus-adapter/values.yaml
+++ b/charts/prometheus-adapter/values.yaml
@@ -1,4 +1,3 @@
-# Default values for k8s-prometheus-adapter..
 affinity: {}
 
 topologySpreadConstraints: []
@@ -21,7 +20,6 @@ priorityClassName: ""
 ## Override the release namespace (for multi-namespace deployments in combined charts)
 namespaceOverride: ""
 
-
 ## Additional annotations to add to all resources
 customAnnotations: {}
   # role: custom-metrics
@@ -65,15 +63,16 @@ serviceAccount:
 
 # Custom DNS configuration to be added to prometheus-adapter pods
 dnsConfig: {}
-# nameservers:
-#   - 1.2.3.4
-# searches:
-#   - ns1.svc.cluster-domain.example
-#   - my.dns.search.suffix
-# options:
-#   - name: ndots
-#     value: "2"
-#   - name: edns0
+  # nameservers:
+  #   - 1.2.3.4
+  # searches:
+  #   - ns1.svc.cluster-domain.example
+  #   - my.dns.search.suffix
+  # options:
+  #   - name: ndots
+  #     value: "2"
+  #   - name: edns0
+
 resources: {}
   # requests:
   #   cpu: 100m
@@ -84,57 +83,61 @@ resources: {}
 
 rules:
   default: true
+
   custom: []
-# - seriesQuery: '{__name__=~"^some_metric_count$"}'
-#   resources:
-#     template: <<.Resource>>
-#   name:
-#     matches: ""
-#     as: "my_custom_metric"
-#   metricsQuery: sum(<<.Series>>{<<.LabelMatchers>>}) by (<<.GroupBy>>)
+    # - seriesQuery: '{__name__=~"^some_metric_count$"}'
+    #   resources:
+    #     template: <<.Resource>>
+    #   name:
+    #     matches: ""
+    #     as: "my_custom_metric"
+    #   metricsQuery: sum(<<.Series>>{<<.LabelMatchers>>}) by (<<.GroupBy>>)
+
   # Mounts a configMap with pre-generated rules for use. Overrides the
   # default, custom, external and resource entries
   existing:
+
   external: []
-# - seriesQuery: '{__name__=~"^some_metric_count$"}'
-#   resources:
-#     template: <<.Resource>>
-#   name:
-#     matches: ""
-#     as: "my_external_metric"
-#   metricsQuery: sum(<<.Series>>{<<.LabelMatchers>>}) by (<<.GroupBy>>)
-  resource: {}
-#   cpu:
-#     containerQuery: sum(rate(container_cpu_usage_seconds_total{<<.LabelMatchers>>, container!=""}[3m])) by (<<.GroupBy>>)
-#     nodeQuery: sum(rate(container_cpu_usage_seconds_total{<<.LabelMatchers>>, id='/'}[3m])) by (<<.GroupBy>>)
-#     resources:
-#       overrides:
-#         node:
-#           resource: node
-#         namespace:
-#           resource: namespace
-#         pod:
-#           resource: pod
-#     containerLabel: container
-#   memory:
-#     containerQuery: sum(container_memory_working_set_bytes{<<.LabelMatchers>>, container!=""}) by (<<.GroupBy>>)
-#     nodeQuery: sum(container_memory_working_set_bytes{<<.LabelMatchers>>,id='/'}) by (<<.GroupBy>>)
-#     resources:
-#       overrides:
-#         node:
-#           resource: node
-#         namespace:
-#           resource: namespace
-#         pod:
-#           resource: pod
-#     containerLabel: container
-#   window: 3m
+    # - seriesQuery: '{__name__=~"^some_metric_count$"}'
+    #   resources:
+    #     template: <<.Resource>>
+    #   name:
+    #     matches: ""
+    #     as: "my_external_metric"
+    #   metricsQuery: sum(<<.Series>>{<<.LabelMatchers>>}) by (<<.GroupBy>>)
+
+  # resource:
+  #   cpu:
+  #     containerQuery: sum(rate(container_cpu_usage_seconds_total{<<.LabelMatchers>>, container!=""}[3m])) by (<<.GroupBy>>)
+  #     nodeQuery: sum(rate(container_cpu_usage_seconds_total{<<.LabelMatchers>>, id='/'}[3m])) by (<<.GroupBy>>)
+  #     resources:
+  #       overrides:
+  #         node:
+  #           resource: node
+  #         namespace:
+  #           resource: namespace
+  #         pod:
+  #           resource: pod
+  #     containerLabel: container
+  #   memory:
+  #     containerQuery: sum(container_memory_working_set_bytes{<<.LabelMatchers>>, container!=""}) by (<<.GroupBy>>)
+  #     nodeQuery: sum(container_memory_working_set_bytes{<<.LabelMatchers>>,id='/'}) by (<<.GroupBy>>)
+  #     resources:
+  #       overrides:
+  #         node:
+  #           resource: node
+  #         namespace:
+  #           resource: namespace
+  #         pod:
+  #           resource: pod
+  #     containerLabel: container
+  #   window: 3m
 
 service:
   annotations: {}
   port: 443
   type: ClusterIP
-# clusterIP: 1.2.3.4
+  # clusterIP: 1.2.3.4
 
 tls:
   enable: false

```

## 3.2.0

**Release date:** 2022-04-03

![AppVersion: v0.9.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.9.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-adapter] Allow namespace override (#1933)

### Default value changes

```diff
diff --git a/charts/prometheus-adapter/values.yaml b/charts/prometheus-adapter/values.yaml
index 9f5de9bc..3315b0c4 100644
--- a/charts/prometheus-adapter/values.yaml
+++ b/charts/prometheus-adapter/values.yaml
@@ -18,6 +18,9 @@ nodeSelector: {}
 
 priorityClassName: ""
 
+## Override the release namespace (for multi-namespace deployments in combined charts)
+namespaceOverride: ""
+
 
 ## Additional annotations to add to all resources
 customAnnotations: {}

```

## 3.1.0

**Release date:** 2022-03-29

![AppVersion: v0.9.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.9.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-adapter] support for topology spread constraints (#1925)

### Default value changes

```diff
diff --git a/charts/prometheus-adapter/values.yaml b/charts/prometheus-adapter/values.yaml
index 04c0a0ef..9f5de9bc 100644
--- a/charts/prometheus-adapter/values.yaml
+++ b/charts/prometheus-adapter/values.yaml
@@ -1,6 +1,8 @@
 # Default values for k8s-prometheus-adapter..
 affinity: {}
 
+topologySpreadConstraints: []
+
 image:
   repository: k8s.gcr.io/prometheus-adapter/prometheus-adapter
   tag: v0.9.1

```

## 3.0.3

**Release date:** 2022-03-04

![AppVersion: v0.9.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.9.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-adapter] Correct prometheus-adapter external metrics rbac (#1838)

### Default value changes

```diff
# No changes in this release
```

## 3.0.2

**Release date:** 2022-02-10

![AppVersion: v0.9.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.9.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-adapter] Fix missing namespace field (#1788)

### Default value changes

```diff
# No changes in this release
```

## 3.0.1

**Release date:** 2022-01-12

![AppVersion: v0.9.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.9.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-adapter] Fix external-metric role binding service account (#1703)

### Default value changes

```diff
# No changes in this release
```

## 3.0.0

**Release date:** 2021-10-20

![AppVersion: v0.9.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.9.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-adapter] Recommanded Kubernetes labels , custom labels and annotations (#1302)

### Default value changes

```diff
diff --git a/charts/prometheus-adapter/values.yaml b/charts/prometheus-adapter/values.yaml
index 7f6b262e..04c0a0ef 100644
--- a/charts/prometheus-adapter/values.yaml
+++ b/charts/prometheus-adapter/values.yaml
@@ -16,6 +16,15 @@ nodeSelector: {}
 
 priorityClassName: ""
 
+
+## Additional annotations to add to all resources
+customAnnotations: {}
+  # role: custom-metrics
+
+## Additional labels to add to all resources
+customLabels: {}
+  # monitoring: prometheus-adapter
+
 # Url to access prometheus
 prometheus:
   # Value is templated

```

## 2.17.1

**Release date:** 2021-10-18

![AppVersion: v0.9.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.9.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Update prometheus-adapter patch version (#1440)

### Default value changes

```diff
diff --git a/charts/prometheus-adapter/values.yaml b/charts/prometheus-adapter/values.yaml
index eff56fe1..7f6b262e 100644
--- a/charts/prometheus-adapter/values.yaml
+++ b/charts/prometheus-adapter/values.yaml
@@ -3,7 +3,7 @@ affinity: {}
 
 image:
   repository: k8s.gcr.io/prometheus-adapter/prometheus-adapter
-  tag: v0.9.0
+  tag: v0.9.1
   pullPolicy: IfNotPresent
 
 logLevel: 4

```

## 2.17.0

**Release date:** 2021-09-04

![AppVersion: v0.9.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.9.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-adapter] Add support for configurable service ip (#1315)

### Default value changes

```diff
diff --git a/charts/prometheus-adapter/values.yaml b/charts/prometheus-adapter/values.yaml
index 689fcfe2..eff56fe1 100644
--- a/charts/prometheus-adapter/values.yaml
+++ b/charts/prometheus-adapter/values.yaml
@@ -120,6 +120,7 @@ service:
   annotations: {}
   port: 443
   type: ClusterIP
+# clusterIP: 1.2.3.4
 
 tls:
   enable: false

```

## 2.16.0

**Release date:** 2021-08-26

![AppVersion: v0.9.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.9.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-adapter] Upgrade image to v0.9.0, use new prometheus adapter repo (#1288)

### Default value changes

```diff
diff --git a/charts/prometheus-adapter/values.yaml b/charts/prometheus-adapter/values.yaml
index 51676f1b..689fcfe2 100644
--- a/charts/prometheus-adapter/values.yaml
+++ b/charts/prometheus-adapter/values.yaml
@@ -2,8 +2,8 @@
 affinity: {}
 
 image:
-  repository: directxman12/k8s-prometheus-adapter-amd64
-  tag: v0.8.4
+  repository: k8s.gcr.io/prometheus-adapter/prometheus-adapter
+  tag: v0.9.0
   pullPolicy: IfNotPresent
 
 logLevel: 4

```

## 2.15.2

**Release date:** 2021-07-23

![AppVersion: v0.8.4](https://img.shields.io/static/v1?label=AppVersion&message=v0.8.4&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* update chart version to fix release (#1195)

### Default value changes

```diff
# No changes in this release
```

## 2.15.1

**Release date:** 2021-07-18

![AppVersion: v0.8.4](https://img.shields.io/static/v1?label=AppVersion&message=v0.8.4&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-adapter] add pod securityContext (#1136)

### Default value changes

```diff
diff --git a/charts/prometheus-adapter/values.yaml b/charts/prometheus-adapter/values.yaml
index 18151643..51676f1b 100644
--- a/charts/prometheus-adapter/values.yaml
+++ b/charts/prometheus-adapter/values.yaml
@@ -25,6 +25,11 @@ prometheus:
 
 replicas: 1
 
+# k8s 1.21 needs fsGroup to be set for non root deployments
+# ref: https://github.com/kubernetes/kubernetes/issues/70679
+podSecurityContext:
+  fsGroup: 10001
+
 rbac:
   # Specifies whether RBAC resources should be created
   create: true

```

## 2.15.0

**Release date:** 2021-07-10

![AppVersion: v0.8.4](https://img.shields.io/static/v1?label=AppVersion&message=v0.8.4&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-adapter] Allow to specify deployment strategy (#1147)

### Default value changes

```diff
diff --git a/charts/prometheus-adapter/values.yaml b/charts/prometheus-adapter/values.yaml
index 7d42f5c5..18151643 100644
--- a/charts/prometheus-adapter/values.yaml
+++ b/charts/prometheus-adapter/values.yaml
@@ -168,6 +168,13 @@ hostNetwork:
 # When hostNetwork is enabled, you probably want to set this to ClusterFirstWithHostNet
 # dnsPolicy: ClusterFirstWithHostNet
 
+# Deployment strategy type
+strategy:
+  type: RollingUpdate
+  rollingUpdate:
+    maxUnavailable: 25%
+    maxSurge: 25%
+
 podDisruptionBudget:
   # Specifies if PodDisruptionBudget should be enabled
   # When enabled, minAvailable or maxUnavailable should also be defined.

```

## 2.14.2

**Release date:** 2021-06-15

![AppVersion: v0.8.4](https://img.shields.io/static/v1?label=AppVersion&message=v0.8.4&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-adapter] fix(README.md): Fixed the default docs to use "node" label rather tha… (#889)

### Default value changes

```diff
diff --git a/charts/prometheus-adapter/values.yaml b/charts/prometheus-adapter/values.yaml
index a9cb21be..7d42f5c5 100644
--- a/charts/prometheus-adapter/values.yaml
+++ b/charts/prometheus-adapter/values.yaml
@@ -90,7 +90,7 @@ rules:
 #     nodeQuery: sum(rate(container_cpu_usage_seconds_total{<<.LabelMatchers>>, id='/'}[3m])) by (<<.GroupBy>>)
 #     resources:
 #       overrides:
-#         instance:
+#         node:
 #           resource: node
 #         namespace:
 #           resource: namespace
@@ -102,7 +102,7 @@ rules:
 #     nodeQuery: sum(container_memory_working_set_bytes{<<.LabelMatchers>>,id='/'}) by (<<.GroupBy>>)
 #     resources:
 #       overrides:
-#         instance:
+#         node:
 #           resource: node
 #         namespace:
 #           resource: namespace

```

## 2.14.1

**Release date:** 2021-06-08

![AppVersion: v0.8.4](https://img.shields.io/static/v1?label=AppVersion&message=v0.8.4&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Increase timeout seconds (#1043)

### Default value changes

```diff
# No changes in this release
```

## 2.14.0

**Release date:** 2021-06-08

![AppVersion: v0.8.4](https://img.shields.io/static/v1?label=AppVersion&message=v0.8.4&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-adapter] Allow specifying service account annotations (#1029)

### Default value changes

```diff
diff --git a/charts/prometheus-adapter/values.yaml b/charts/prometheus-adapter/values.yaml
index 86046b50..a9cb21be 100644
--- a/charts/prometheus-adapter/values.yaml
+++ b/charts/prometheus-adapter/values.yaml
@@ -39,6 +39,11 @@ serviceAccount:
   # The name of the service account to use.
   # If not set and create is true, a name is generated using the fullname template
   name:
+  # ServiceAccount annotations.
+  # Use case: AWS EKS IAM roles for service accounts
+  # ref: https://docs.aws.amazon.com/eks/latest/userguide/specify-service-account-role.html
+  annotations: {}
+
 # Custom DNS configuration to be added to prometheus-adapter pods
 dnsConfig: {}
 # nameservers:

```

## 2.13.0

**Release date:** 2021-05-29

![AppVersion: v0.8.4](https://img.shields.io/static/v1?label=AppVersion&message=v0.8.4&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-adapter] Add hostPorts in psp (#982)

### Default value changes

```diff
# No changes in this release
```

## 2.12.3

**Release date:** 2021-05-10

![AppVersion: v0.8.4](https://img.shields.io/static/v1?label=AppVersion&message=v0.8.4&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-adapter] filter resource metrics for pod for real row (#950)

### Default value changes

```diff
diff --git a/charts/prometheus-adapter/values.yaml b/charts/prometheus-adapter/values.yaml
index f8538387..86046b50 100644
--- a/charts/prometheus-adapter/values.yaml
+++ b/charts/prometheus-adapter/values.yaml
@@ -81,7 +81,7 @@ rules:
 #   metricsQuery: sum(<<.Series>>{<<.LabelMatchers>>}) by (<<.GroupBy>>)
   resource: {}
 #   cpu:
-#     containerQuery: sum(rate(container_cpu_usage_seconds_total{<<.LabelMatchers>>}[3m])) by (<<.GroupBy>>)
+#     containerQuery: sum(rate(container_cpu_usage_seconds_total{<<.LabelMatchers>>, container!=""}[3m])) by (<<.GroupBy>>)
 #     nodeQuery: sum(rate(container_cpu_usage_seconds_total{<<.LabelMatchers>>, id='/'}[3m])) by (<<.GroupBy>>)
 #     resources:
 #       overrides:
@@ -93,7 +93,7 @@ rules:
 #           resource: pod
 #     containerLabel: container
 #   memory:
-#     containerQuery: sum(container_memory_working_set_bytes{<<.LabelMatchers>>}) by (<<.GroupBy>>)
+#     containerQuery: sum(container_memory_working_set_bytes{<<.LabelMatchers>>, container!=""}) by (<<.GroupBy>>)
 #     nodeQuery: sum(container_memory_working_set_bytes{<<.LabelMatchers>>,id='/'}) by (<<.GroupBy>>)
 #     resources:
 #       overrides:

```

## 2.12.2

**Release date:** 2021-04-27

![AppVersion: v0.8.4](https://img.shields.io/static/v1?label=AppVersion&message=v0.8.4&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* prometheus-adapter: upgrade appversdion to v0.8.4 (#895)

### Default value changes

```diff
diff --git a/charts/prometheus-adapter/values.yaml b/charts/prometheus-adapter/values.yaml
index 1baceeb8..f8538387 100644
--- a/charts/prometheus-adapter/values.yaml
+++ b/charts/prometheus-adapter/values.yaml
@@ -3,7 +3,7 @@ affinity: {}
 
 image:
   repository: directxman12/k8s-prometheus-adapter-amd64
-  tag: v0.8.3
+  tag: v0.8.4
   pullPolicy: IfNotPresent
 
 logLevel: 4

```

## 2.12.1

**Release date:** 2021-02-20

![AppVersion: v0.8.3](https://img.shields.io/static/v1?label=AppVersion&message=v0.8.3&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Fix overzealous whitespace trimming. (#689)

### Default value changes

```diff
# No changes in this release
```

## 2.12.0

**Release date:** 2021-02-09

![AppVersion: v0.8.3](https://img.shields.io/static/v1?label=AppVersion&message=v0.8.3&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Ability for custom dnsConfig to prometheus-adapter (#640)

### Default value changes

```diff
diff --git a/charts/prometheus-adapter/values.yaml b/charts/prometheus-adapter/values.yaml
index 2f93f39d..1baceeb8 100644
--- a/charts/prometheus-adapter/values.yaml
+++ b/charts/prometheus-adapter/values.yaml
@@ -39,7 +39,17 @@ serviceAccount:
   # The name of the service account to use.
   # If not set and create is true, a name is generated using the fullname template
   name:
-
+# Custom DNS configuration to be added to prometheus-adapter pods
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
 resources: {}
   # requests:
   #   cpu: 100m

```

## 2.11.1

**Release date:** 2021-02-02

![AppVersion: v0.8.3](https://img.shields.io/static/v1?label=AppVersion&message=v0.8.3&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* chore: upgrade prometheus-adapter to v0.8.3 (#627)

### Default value changes

```diff
diff --git a/charts/prometheus-adapter/values.yaml b/charts/prometheus-adapter/values.yaml
index a73eeb36..2f93f39d 100644
--- a/charts/prometheus-adapter/values.yaml
+++ b/charts/prometheus-adapter/values.yaml
@@ -3,7 +3,7 @@ affinity: {}
 
 image:
   repository: directxman12/k8s-prometheus-adapter-amd64
-  tag: v0.8.2
+  tag: v0.8.3
   pullPolicy: IfNotPresent
 
 logLevel: 4

```

## 2.11.0

**Release date:** 2021-01-16

![AppVersion: v0.8.2](https://img.shields.io/static/v1?label=AppVersion&message=v0.8.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-adapter] cert-manager support (#480)

### Default value changes

```diff
diff --git a/charts/prometheus-adapter/values.yaml b/charts/prometheus-adapter/values.yaml
index 0637f101..a73eeb36 100644
--- a/charts/prometheus-adapter/values.yaml
+++ b/charts/prometheus-adapter/values.yaml
@@ -159,3 +159,8 @@ podDisruptionBudget:
   enabled: false
   minAvailable:
   maxUnavailable: 1
+
+certManager:
+  enabled: false
+  caCertDuration: 43800h
+  certDuration: 8760h

```

## 2.10.1

**Release date:** 2020-12-15

![AppVersion: v0.8.2](https://img.shields.io/static/v1?label=AppVersion&message=v0.8.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-adapter] Use the supported APIService API version (#463)

### Default value changes

```diff
# No changes in this release
```

## 2.10.0

**Release date:** 2020-12-12

![AppVersion: v0.8.2](https://img.shields.io/static/v1?label=AppVersion&message=v0.8.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-adapter] support extra arguments for the prometheus-adapter deployment (#487)

### Default value changes

```diff
diff --git a/charts/prometheus-adapter/values.yaml b/charts/prometheus-adapter/values.yaml
index b7b7e877..0637f101 100644
--- a/charts/prometheus-adapter/values.yaml
+++ b/charts/prometheus-adapter/values.yaml
@@ -110,6 +110,11 @@ tls:
   certificate: |-
     # Public key of the APIService
 
+# Any extra arguments
+extraArguments: []
+  # - --tls-private-key-file=/etc/tls/tls.key
+  # - --tls-cert-file=/etc/tls/tls.crt
+
 # Any extra volumes
 extraVolumes: []
   # - name: example-name

```

## 2.9.0

**Release date:** 2020-12-07

![AppVersion: v0.8.2](https://img.shields.io/static/v1?label=AppVersion&message=v0.8.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-adapter] Upgrade to v0.8.2 (#462)

### Default value changes

```diff
diff --git a/charts/prometheus-adapter/values.yaml b/charts/prometheus-adapter/values.yaml
index 230c926b..b7b7e877 100644
--- a/charts/prometheus-adapter/values.yaml
+++ b/charts/prometheus-adapter/values.yaml
@@ -3,7 +3,7 @@ affinity: {}
 
 image:
   repository: directxman12/k8s-prometheus-adapter-amd64
-  tag: v0.7.0
+  tag: v0.8.2
   pullPolicy: IfNotPresent
 
 logLevel: 4

```

## 2.8.1

**Release date:** 2020-12-05

![AppVersion: v0.7.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.7.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-adapter] Default to not dnsPolicy to keep cluster default (#446)

### Default value changes

```diff
diff --git a/charts/prometheus-adapter/values.yaml b/charts/prometheus-adapter/values.yaml
index c7cf3d9d..230c926b 100644
--- a/charts/prometheus-adapter/values.yaml
+++ b/charts/prometheus-adapter/values.yaml
@@ -146,7 +146,7 @@ hostNetwork:
   enabled: false
 
 # When hostNetwork is enabled, you probably want to set this to ClusterFirstWithHostNet
-dnsPolicy: Default
+# dnsPolicy: ClusterFirstWithHostNet
 
 podDisruptionBudget:
   # Specifies if PodDisruptionBudget should be enabled

```

## 2.8.0

**Release date:** 2020-12-02

![AppVersion: v0.7.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.7.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-adapter] Add missing dnsPolicy for hostNetwork (#375)

### Default value changes

```diff
diff --git a/charts/prometheus-adapter/values.yaml b/charts/prometheus-adapter/values.yaml
index 23548a5d..c7cf3d9d 100644
--- a/charts/prometheus-adapter/values.yaml
+++ b/charts/prometheus-adapter/values.yaml
@@ -142,9 +142,12 @@ hostNetwork:
   #
   # You would require this enabled if you use alternate overlay networking for pods and
   # API server unable to communicate with metrics-server. As an example, this is required
-  # if you use Weave network on EKS
+  # if you use Weave network on EKS. See also dnsPolicy
   enabled: false
 
+# When hostNetwork is enabled, you probably want to set this to ClusterFirstWithHostNet
+dnsPolicy: Default
+
 podDisruptionBudget:
   # Specifies if PodDisruptionBudget should be enabled
   # When enabled, minAvailable or maxUnavailable should also be defined.

```

## 2.7.2

**Release date:** 2020-12-02

![AppVersion: v0.7.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.7.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-adapter] Template prometheus.url & rename files (#205)

### Default value changes

```diff
diff --git a/charts/prometheus-adapter/values.yaml b/charts/prometheus-adapter/values.yaml
index 187124e0..23548a5d 100644
--- a/charts/prometheus-adapter/values.yaml
+++ b/charts/prometheus-adapter/values.yaml
@@ -18,6 +18,7 @@ priorityClassName: ""
 
 # Url to access prometheus
 prometheus:
+  # Value is templated
   url: http://prometheus.default.svc
   port: 9090
   path: ""

```

## 2.7.1

**Release date:** 2020-11-14

![AppVersion: v0.7.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.7.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-adapter] Add missing namespace fields to custom-metric resources to allow install dedicated ns (#358)

### Default value changes

```diff
# No changes in this release
```

## 2.7.0

**Release date:** 2020-09-25

![AppVersion: v0.7.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.7.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-adapter] Add a pod disruption budget (#143)

### Default value changes

```diff
diff --git a/charts/prometheus-adapter/values.yaml b/charts/prometheus-adapter/values.yaml
index 7857c06d..187124e0 100644
--- a/charts/prometheus-adapter/values.yaml
+++ b/charts/prometheus-adapter/values.yaml
@@ -143,3 +143,10 @@ hostNetwork:
   # API server unable to communicate with metrics-server. As an example, this is required
   # if you use Weave network on EKS
   enabled: false
+
+podDisruptionBudget:
+  # Specifies if PodDisruptionBudget should be enabled
+  # When enabled, minAvailable or maxUnavailable should also be defined.
+  enabled: false
+  minAvailable:
+  maxUnavailable: 1

```

## 2.6.2

**Release date:** 2020-09-23

![AppVersion: v0.7.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.7.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-adapter] bugfix psp port range (#135)

### Default value changes

```diff
# No changes in this release
```

## 2.6.1

**Release date:** 2020-09-23

![AppVersion: v0.7.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.7.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-adapter] Fix auth-reader rolebinding (#120)

### Default value changes

```diff
# No changes in this release
```

## 2.6.0

**Release date:** 2020-09-22

![AppVersion: v0.7.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.7.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-adapter] Add PodSecurityPolicy and fix custom metrics role binding (#115)

### Default value changes

```diff
diff --git a/charts/prometheus-adapter/values.yaml b/charts/prometheus-adapter/values.yaml
index 4cc5e406..7857c06d 100644
--- a/charts/prometheus-adapter/values.yaml
+++ b/charts/prometheus-adapter/values.yaml
@@ -28,6 +28,10 @@ rbac:
   # Specifies whether RBAC resources should be created
   create: true
 
+psp:
+  # Specifies whether PSP resources should be created
+  create: false
+
 serviceAccount:
   # Specifies whether a service account should be created
   create: true

```

## 2.5.2

**Release date:** 2020-09-14

![AppVersion: v0.7.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.7.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-adapter] Remove colons from the ClusterRoleBinding resource (#94)

### Default value changes

```diff
# No changes in this release
```

## 2.5.1

**Release date:** 2020-08-20

![AppVersion: v0.7.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.7.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Prep initial charts indexing (#14)

### Default value changes

```diff
# No changes in this release
```

## 2.5.0

**Release date:** 2020-07-17

![AppVersion: v0.7.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.7.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-adapter] Upgrading to v0.7.0 (#23245)

### Default value changes

```diff
diff --git a/charts/prometheus-adapter/values.yaml b/charts/prometheus-adapter/values.yaml
index 1d414b54..4cc5e406 100644
--- a/charts/prometheus-adapter/values.yaml
+++ b/charts/prometheus-adapter/values.yaml
@@ -3,7 +3,7 @@ affinity: {}
 
 image:
   repository: directxman12/k8s-prometheus-adapter-amd64
-  tag: v0.6.0
+  tag: v0.7.0
   pullPolicy: IfNotPresent
 
 logLevel: 4

```

## 2.4.0

**Release date:** 2020-06-12

![AppVersion: v0.6.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.6.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-adapter] Allow specifying a path for the Prometheus URL (#22760)

### Default value changes

```diff
diff --git a/charts/prometheus-adapter/values.yaml b/charts/prometheus-adapter/values.yaml
index d27bc9a2..1d414b54 100644
--- a/charts/prometheus-adapter/values.yaml
+++ b/charts/prometheus-adapter/values.yaml
@@ -20,6 +20,7 @@ priorityClassName: ""
 prometheus:
   url: http://prometheus.default.svc
   port: 9090
+  path: ""
 
 replicas: 1
 

```

## 2.3.1

**Release date:** 2020-04-19

![AppVersion: v0.6.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.6.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* add watch to resource-reader ClusterRole (#22010)

### Default value changes

```diff
# No changes in this release
```

## 2.3.0

**Release date:** 2020-04-14

![AppVersion: v0.6.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.6.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-adapter] Defining a variable for listen port (#21918)

### Default value changes

```diff
diff --git a/charts/prometheus-adapter/values.yaml b/charts/prometheus-adapter/values.yaml
index 6d6c74c2..d27bc9a2 100644
--- a/charts/prometheus-adapter/values.yaml
+++ b/charts/prometheus-adapter/values.yaml
@@ -10,6 +10,8 @@ logLevel: 4
 
 metricsRelistInterval: 1m
 
+listenPort: 6443
+
 nodeSelector: {}
 
 priorityClassName: ""

```

## 2.2.0

**Release date:** 2020-03-27

![AppVersion: v0.6.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.6.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-adapter] Add podLabels support (#21658)

### Default value changes

```diff
diff --git a/charts/prometheus-adapter/values.yaml b/charts/prometheus-adapter/values.yaml
index 8e960bbc..6d6c74c2 100644
--- a/charts/prometheus-adapter/values.yaml
+++ b/charts/prometheus-adapter/values.yaml
@@ -123,6 +123,9 @@ extraVolumeMounts: []
 
 tolerations: []
 
+# Labels added to the pod
+podLabels: {}
+
 # Annotations added to the pod
 podAnnotations: {}
 

```

## 2.1.3

**Release date:** 2020-02-28

![AppVersion: v0.6.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.6.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-adapter] Do not create custom metrics rbac (#20900)

### Default value changes

```diff
# No changes in this release
```

## 2.1.2

**Release date:** 2020-02-21

![AppVersion: v0.6.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.6.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-adapter]: do not create apiservice when not needed (#20883)

### Default value changes

```diff
# No changes in this release
```

## 2.1.1

**Release date:** 2020-02-18

![AppVersion: v0.6.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.6.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* pass empty array if rules are empty (#19616) (#20819)

### Default value changes

```diff
# No changes in this release
```

## 2.1.0

**Release date:** 2020-02-17

![AppVersion: v0.6.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.6.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-adapter] Upgrading to v0.6.0 (#20750)

### Default value changes

```diff
diff --git a/charts/prometheus-adapter/values.yaml b/charts/prometheus-adapter/values.yaml
index 9e4b9a5e..8e960bbc 100644
--- a/charts/prometheus-adapter/values.yaml
+++ b/charts/prometheus-adapter/values.yaml
@@ -3,7 +3,7 @@ affinity: {}
 
 image:
   repository: directxman12/k8s-prometheus-adapter-amd64
-  tag: v0.5.0
+  tag: v0.6.0
   pullPolicy: IfNotPresent
 
 logLevel: 4

```

## 2.0.1

**Release date:** 2020-01-16

![AppVersion: v0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.5.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-adapter] Allow prometheus-adapter to start in hostNetwork mode (#19466)

### Default value changes

```diff
diff --git a/charts/prometheus-adapter/values.yaml b/charts/prometheus-adapter/values.yaml
index ba1a5c4c..9e4b9a5e 100644
--- a/charts/prometheus-adapter/values.yaml
+++ b/charts/prometheus-adapter/values.yaml
@@ -125,3 +125,11 @@ tolerations: []
 
 # Annotations added to the pod
 podAnnotations: {}
+
+hostNetwork:
+  # Specifies if prometheus-adapter should be started in hostNetwork mode.
+  #
+  # You would require this enabled if you use alternate overlay networking for pods and
+  # API server unable to communicate with metrics-server. As an example, this is required
+  # if you use Weave network on EKS
+  enabled: false

```

## 2.0.0

**Release date:** 2020-01-10

![AppVersion: v0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.5.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [promethesus-adapter] use the new tag name (#19961)

### Default value changes

```diff
diff --git a/charts/prometheus-adapter/values.yaml b/charts/prometheus-adapter/values.yaml
index 4e13ad6c..ba1a5c4c 100644
--- a/charts/prometheus-adapter/values.yaml
+++ b/charts/prometheus-adapter/values.yaml
@@ -71,9 +71,9 @@ rules:
 #           resource: node
 #         namespace:
 #           resource: namespace
-#         pod_name:
+#         pod:
 #           resource: pod
-#     containerLabel: container_name
+#     containerLabel: container
 #   memory:
 #     containerQuery: sum(container_memory_working_set_bytes{<<.LabelMatchers>>}) by (<<.GroupBy>>)
 #     nodeQuery: sum(container_memory_working_set_bytes{<<.LabelMatchers>>,id='/'}) by (<<.GroupBy>>)
@@ -83,9 +83,9 @@ rules:
 #           resource: node
 #         namespace:
 #           resource: namespace
-#         pod_name:
+#         pod:
 #           resource: pod
-#     containerLabel: container_name
+#     containerLabel: container
 #   window: 3m
 
 service:

```

## 1.4.0

**Release date:** 2019-10-15

![AppVersion: v0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.5.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-adapter] Support extra volumes (#17982)

### Default value changes

```diff
diff --git a/charts/prometheus-adapter/values.yaml b/charts/prometheus-adapter/values.yaml
index 1370ae01..4e13ad6c 100644
--- a/charts/prometheus-adapter/values.yaml
+++ b/charts/prometheus-adapter/values.yaml
@@ -102,6 +102,25 @@ tls:
   certificate: |-
     # Public key of the APIService
 
+# Any extra volumes
+extraVolumes: []
+  # - name: example-name
+  #   hostPath:
+  #     path: /path/on/host
+  #     type: DirectoryOrCreate
+  # - name: ssl-certs
+  #   hostPath:
+  #     path: /etc/ssl/certs/ca-bundle.crt
+  #     type: File
+
+# Any extra volume mounts
+extraVolumeMounts: []
+  #   - name: example-name
+  #     mountPath: /path/in/container
+  #   - name: ssl-certs
+  #     mountPath: /etc/ssl/certs/ca-certificates.crt
+  #     readOnly: true
+
 tolerations: []
 
 # Annotations added to the pod

```

## 1.3.0

**Release date:** 2019-08-22

![AppVersion: v0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.5.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-adapter] Added pod annotations to chart.  (#16446)

### Default value changes

```diff
diff --git a/charts/prometheus-adapter/values.yaml b/charts/prometheus-adapter/values.yaml
index 79942e46..1370ae01 100644
--- a/charts/prometheus-adapter/values.yaml
+++ b/charts/prometheus-adapter/values.yaml
@@ -103,3 +103,6 @@ tls:
     # Public key of the APIService
 
 tolerations: []
+
+# Annotations added to the pod
+podAnnotations: {}

```

## 1.2.0

**Release date:** 2019-06-11

![AppVersion: v0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.5.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-adapter] Add priorityClassName (#14570)

### Default value changes

```diff
diff --git a/charts/prometheus-adapter/values.yaml b/charts/prometheus-adapter/values.yaml
index 251d3a39..79942e46 100644
--- a/charts/prometheus-adapter/values.yaml
+++ b/charts/prometheus-adapter/values.yaml
@@ -12,6 +12,8 @@ metricsRelistInterval: 1m
 
 nodeSelector: {}
 
+priorityClassName: ""
+
 # Url to access prometheus
 prometheus:
   url: http://prometheus.default.svc

```

## 1.1.0

**Release date:** 2019-06-06

![AppVersion: v0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.5.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-adapter] add new resourceRules in values (#14508)

### Default value changes

```diff
diff --git a/charts/prometheus-adapter/values.yaml b/charts/prometheus-adapter/values.yaml
index 28fc9f14..251d3a39 100644
--- a/charts/prometheus-adapter/values.yaml
+++ b/charts/prometheus-adapter/values.yaml
@@ -49,7 +49,7 @@ rules:
 #     as: "my_custom_metric"
 #   metricsQuery: sum(<<.Series>>{<<.LabelMatchers>>}) by (<<.GroupBy>>)
   # Mounts a configMap with pre-generated rules for use. Overrides the
-  # default, custom and external entries
+  # default, custom, external and resource entries
   existing:
   external: []
 # - seriesQuery: '{__name__=~"^some_metric_count$"}'
@@ -57,8 +57,34 @@ rules:
 #     template: <<.Resource>>
 #   name:
 #     matches: ""
-#     as: "my_custom_metric"
+#     as: "my_external_metric"
 #   metricsQuery: sum(<<.Series>>{<<.LabelMatchers>>}) by (<<.GroupBy>>)
+  resource: {}
+#   cpu:
+#     containerQuery: sum(rate(container_cpu_usage_seconds_total{<<.LabelMatchers>>}[3m])) by (<<.GroupBy>>)
+#     nodeQuery: sum(rate(container_cpu_usage_seconds_total{<<.LabelMatchers>>, id='/'}[3m])) by (<<.GroupBy>>)
+#     resources:
+#       overrides:
+#         instance:
+#           resource: node
+#         namespace:
+#           resource: namespace
+#         pod_name:
+#           resource: pod
+#     containerLabel: container_name
+#   memory:
+#     containerQuery: sum(container_memory_working_set_bytes{<<.LabelMatchers>>}) by (<<.GroupBy>>)
+#     nodeQuery: sum(container_memory_working_set_bytes{<<.LabelMatchers>>,id='/'}) by (<<.GroupBy>>)
+#     resources:
+#       overrides:
+#         instance:
+#           resource: node
+#         namespace:
+#           resource: namespace
+#         pod_name:
+#           resource: pod
+#     containerLabel: container_name
+#   window: 3m
 
 service:
   annotations: {}

```

## 1.0.4

**Release date:** 2019-06-05

![AppVersion: v0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.5.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* owners: add a new owner (#14530)

### Default value changes

```diff
# No changes in this release
```

## 1.0.3

**Release date:** 2019-06-02

![AppVersion: v0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.5.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-adapter] conditionally register external metrics API (#13716) (#14389)

### Default value changes

```diff
# No changes in this release
```

## 1.0.2

**Release date:** 2019-05-21

![AppVersion: v0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.5.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-adapter] Add myself to OWNERS (#13887)

### Default value changes

```diff
# No changes in this release
```

## 1.0.1

**Release date:** 2019-05-20

![AppVersion: v0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.5.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* prometheus-adapter: allow to set custom prometheus URL path (#13174)

### Default value changes

```diff
# No changes in this release
```

## 1.0.0

**Release date:** 2019-05-15

![AppVersion: v0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.5.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* add apiVersion (#13818)

### Default value changes

```diff
# No changes in this release
```

## v0.5.0

**Release date:** 2019-04-29

![AppVersion: v0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.5.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)

* [stable/prometheus-adapter] Adds support for external metrics v0.5.0 (#13133)

### Default value changes

```diff
diff --git a/charts/prometheus-adapter/values.yaml b/charts/prometheus-adapter/values.yaml
index 070bc462..28fc9f14 100644
--- a/charts/prometheus-adapter/values.yaml
+++ b/charts/prometheus-adapter/values.yaml
@@ -3,7 +3,7 @@ affinity: {}
 
 image:
   repository: directxman12/k8s-prometheus-adapter-amd64
-  tag: v0.4.1
+  tag: v0.5.0
   pullPolicy: IfNotPresent
 
 logLevel: 4
@@ -49,8 +49,16 @@ rules:
 #     as: "my_custom_metric"
 #   metricsQuery: sum(<<.Series>>{<<.LabelMatchers>>}) by (<<.GroupBy>>)
   # Mounts a configMap with pre-generated rules for use. Overrides the
-  # default and custom entries
+  # default, custom and external entries
   existing:
+  external: []
+# - seriesQuery: '{__name__=~"^some_metric_count$"}'
+#   resources:
+#     template: <<.Resource>>
+#   name:
+#     matches: ""
+#     as: "my_custom_metric"
+#   metricsQuery: sum(<<.Series>>{<<.LabelMatchers>>}) by (<<.GroupBy>>)
 
 service:
   annotations: {}

```

## v0.4.2

**Release date:** 2019-04-25

![AppVersion: v0.4.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.4.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)

* fixes incompatibility with 1.11 (#13261)

### Default value changes

```diff
# No changes in this release
```

## v0.4.1

**Release date:** 2019-01-17

![AppVersion: v0.4.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.4.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)

* Fix/prometheus adapter image pull secrets (#10725)

### Default value changes

```diff
# No changes in this release
```

## v0.4.0

**Release date:** 2019-01-15

![AppVersion: v0.4.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.4.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)

* Add support for image pull secrets for api server deployment (#10269)

### Default value changes

```diff
# No changes in this release
```

## v0.3.0

**Release date:** 2019-01-02

![AppVersion: v0.4.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.4.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)

* Adding support for externally defined adapter rules (#8959)

### Default value changes

```diff
diff --git a/charts/prometheus-adapter/values.yaml b/charts/prometheus-adapter/values.yaml
index f91f200e..070bc462 100644
--- a/charts/prometheus-adapter/values.yaml
+++ b/charts/prometheus-adapter/values.yaml
@@ -48,6 +48,9 @@ rules:
 #     matches: ""
 #     as: "my_custom_metric"
 #   metricsQuery: sum(<<.Series>>{<<.LabelMatchers>>}) by (<<.GroupBy>>)
+  # Mounts a configMap with pre-generated rules for use. Overrides the
+  # default and custom entries
+  existing:
 
 service:
   annotations: {}

```

## v0.2.3

**Release date:** 2018-12-20

![AppVersion: v0.4.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.4.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)

* [stable/prometheus-adapter] readOnlyFileSystem (#10042)

### Default value changes

```diff
# No changes in this release
```

## v0.2.2

**Release date:** 2018-12-16

![AppVersion: v0.4.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.4.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)

* [stable/prometheus-adapter] v0.4.1, readOnlyRootFilesystem, fixes (#10036)

### Default value changes

```diff
diff --git a/charts/prometheus-adapter/values.yaml b/charts/prometheus-adapter/values.yaml
index 23ade966..f91f200e 100644
--- a/charts/prometheus-adapter/values.yaml
+++ b/charts/prometheus-adapter/values.yaml
@@ -3,12 +3,12 @@ affinity: {}
 
 image:
   repository: directxman12/k8s-prometheus-adapter-amd64
-  tag: v0.4.0
+  tag: v0.4.1
   pullPolicy: IfNotPresent
 
 logLevel: 4
 
-metricsRelistInterval: 30s
+metricsRelistInterval: 1m
 
 nodeSelector: {}
 
@@ -30,7 +30,7 @@ serviceAccount:
   # If not set and create is true, a name is generated using the fullname template
   name:
 
-resources:
+resources: {}
   # requests:
   #   cpu: 100m
   #   memory: 128Mi

```

## v0.2.1

**Release date:** 2018-12-04

![AppVersion: v0.4.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.4.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)

* Upgrading to v0.4.0 (#9711)

### Default value changes

```diff
diff --git a/charts/prometheus-adapter/values.yaml b/charts/prometheus-adapter/values.yaml
index da65e027..23ade966 100644
--- a/charts/prometheus-adapter/values.yaml
+++ b/charts/prometheus-adapter/values.yaml
@@ -3,7 +3,7 @@ affinity: {}
 
 image:
   repository: directxman12/k8s-prometheus-adapter-amd64
-  tag: v0.2.1
+  tag: v0.4.0
   pullPolicy: IfNotPresent
 
 logLevel: 4

```

## v0.2.0

**Release date:** 2018-10-11

![AppVersion: v0.2.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.2.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)

* [stable/prometheus-adapter] Add checksum of configMap as annotation (#8334)

### Default value changes

```diff
# No changes in this release
```

## v0.1.2

**Release date:** 2018-09-19

![AppVersion: v0.2.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.2.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)

* Changed syntax error in custom-metrics-apiserver-service and secret (label to labels) (#7295)

### Default value changes

```diff
# No changes in this release
```

## v0.1.1

**Release date:** 2018-09-08

![AppVersion: v0.2.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.2.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)

* Fixing warning when overriding rules.custom (#7586)

### Default value changes

```diff
diff --git a/charts/prometheus-adapter/values.yaml b/charts/prometheus-adapter/values.yaml
index 460d0c19..da65e027 100644
--- a/charts/prometheus-adapter/values.yaml
+++ b/charts/prometheus-adapter/values.yaml
@@ -40,7 +40,7 @@ resources:
 
 rules:
   default: true
-  custom: {}
+  custom: []
 # - seriesQuery: '{__name__=~"^some_metric_count$"}'
 #   resources:
 #     template: <<.Resource>>

```

## v0.1.0

**Release date:** 2018-08-15

![AppVersion: v0.2.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.2.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)

* [stable/prometheus-adapter] Enhancements (#7183)

### Default value changes

```diff
diff --git a/charts/prometheus-adapter/values.yaml b/charts/prometheus-adapter/values.yaml
index 624268ff..460d0c19 100644
--- a/charts/prometheus-adapter/values.yaml
+++ b/charts/prometheus-adapter/values.yaml
@@ -1,9 +1,22 @@
 # Default values for k8s-prometheus-adapter..
+affinity: {}
+
 image:
   repository: directxman12/k8s-prometheus-adapter-amd64
   tag: v0.2.1
   pullPolicy: IfNotPresent
 
+logLevel: 4
+
+metricsRelistInterval: 30s
+
+nodeSelector: {}
+
+# Url to access prometheus
+prometheus:
+  url: http://prometheus.default.svc
+  port: 9090
+
 replicas: 1
 
 rbac:
@@ -25,6 +38,22 @@ resources:
   #   cpu: 100m
   #   memory: 128Mi
 
+rules:
+  default: true
+  custom: {}
+# - seriesQuery: '{__name__=~"^some_metric_count$"}'
+#   resources:
+#     template: <<.Resource>>
+#   name:
+#     matches: ""
+#     as: "my_custom_metric"
+#   metricsQuery: sum(<<.Series>>{<<.LabelMatchers>>}) by (<<.GroupBy>>)
+
+service:
+  annotations: {}
+  port: 443
+  type: ClusterIP
+
 tls:
   enable: false
   ca: |-
@@ -34,10 +63,4 @@ tls:
   certificate: |-
     # Public key of the APIService
 
-# Url to access prometheus
-prometheus:
-  url: http://prometheus.default.svc
-  port: 9090
-
-metricsRelistInterval: 30s
-logLevel: 10
+tolerations: []

```

## v0.0.1

**Release date:** 2018-08-13

![AppVersion: v0.2.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.2.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)

* Add prometheus adapter chart (#6082)

### Default value changes

```diff
# Default values for k8s-prometheus-adapter..
image:
  repository: directxman12/k8s-prometheus-adapter-amd64
  tag: v0.2.1
  pullPolicy: IfNotPresent

replicas: 1

rbac:
  # Specifies whether RBAC resources should be created
  create: true

serviceAccount:
  # Specifies whether a service account should be created
  create: true
  # The name of the service account to use.
  # If not set and create is true, a name is generated using the fullname template
  name:

resources:
  # requests:
  #   cpu: 100m
  #   memory: 128Mi
  # limits:
  #   cpu: 100m
  #   memory: 128Mi

tls:
  enable: false
  ca: |-
    # Public CA file that signed the APIService
  key: |-
    # Private key of the APIService
  certificate: |-
    # Public key of the APIService

# Url to access prometheus
prometheus:
  url: http://prometheus.default.svc
  port: 9090

metricsRelistInterval: 30s
logLevel: 10

```

---
Autogenerated from Helm Chart and git history using [helm-changelog](https://github.com/mogensen/helm-changelog)
