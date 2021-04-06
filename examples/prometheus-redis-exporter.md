# Change Log

## 4.0.0 

**Release date:** 2020-11-23

![AppVersion: 1.11.1](https://img.shields.io/static/v1?label=AppVersion&message=1.11.1&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-redis-exporter] Migrate to chart v2 (#389) 

### Default value changes

```diff
# No changes in this release
```

## 3.7.0 

**Release date:** 2020-11-18

![AppVersion: 1.11.1](https://img.shields.io/static/v1?label=AppVersion&message=1.11.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* feat(prometheus-redis-exporter): add configurable apiVersion (#369) 

### Default value changes

```diff
# No changes in this release
```

## 3.6.0 

**Release date:** 2020-09-11

![AppVersion: 1.11.1](https://img.shields.io/static/v1?label=AppVersion&message=1.11.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-redis-exporter] Upgrade exporter image to version 1.11.1 (#87) 

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index 68151d3..967184e 100644
--- a/charts/prometheus-redis-exporter/values.yaml
+++ b/charts/prometheus-redis-exporter/values.yaml
@@ -13,7 +13,7 @@ serviceAccount:
 replicaCount: 1
 image:
   repository: oliver006/redis_exporter
-  tag: v1.3.4
+  tag: v1.11.1
   pullPolicy: IfNotPresent
 extraArgs: {}
 # Additional Environment variables
```

## 3.5.2 

**Release date:** 2020-09-06

![AppVersion: 1.3.4](https://img.shields.io/static/v1?label=AppVersion&message=1.3.4&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-redis-exporter] Add zanhsieh as maintainer (#46) 

### Default value changes

```diff
# No changes in this release
```

## 3.5.1 

**Release date:** 2020-08-20

![AppVersion: 1.3.4](https://img.shields.io/static/v1?label=AppVersion&message=1.3.4&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Prep initial charts indexing (#14) 

### Default value changes

```diff
# No changes in this release
```

## 3.5.0 

**Release date:** 2020-07-22

![AppVersion: 1.3.4](https://img.shields.io/static/v1?label=AppVersion&message=1.3.4&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-redis-exporter] Add ServiceMonitor metricRelabelings (#23289) 

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index c19f7a4..68151d3 100644
--- a/charts/prometheus-redis-exporter/values.yaml
+++ b/charts/prometheus-redis-exporter/values.yaml
@@ -60,6 +60,7 @@ serviceMonitor:
   # timeout: 10s
   # Set of labels to transfer on the Kubernetes Service onto the target.
   # targetLabels: []
+  # metricRelabelings: []
 
 ## Custom PrometheusRules to be defined
 ## The value is evaluated as a template, so, for example, the value can depend on .Release or .Chart
```

## 3.4.1 

**Release date:** 2020-05-28

![AppVersion: 1.3.4](https://img.shields.io/static/v1?label=AppVersion&message=1.3.4&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-redis-exporter] fix targetLabels (#22541) 

### Default value changes

```diff
# No changes in this release
```

## 3.4.0 

**Release date:** 2020-04-09

![AppVersion: 1.3.4](https://img.shields.io/static/v1?label=AppVersion&message=1.3.4&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-redis-exporter] Add support for tolerations and affinity (#21848) 

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index ba02195..c19f7a4 100644
--- a/charts/prometheus-redis-exporter/values.yaml
+++ b/charts/prometheus-redis-exporter/values.yaml
@@ -32,6 +32,13 @@ service:
     # prometheus.io/port: "9121"
     # prometheus.io/scrape: "true"
 resources: {}
+
+nodeSelector: {}
+
+tolerations: []
+
+affinity: {}
+
 redisAddress: redis://myredis:6379
 annotations: {}
 #  prometheus.io/path: /metrics
```

## 3.3.3 

**Release date:** 2020-03-11

![AppVersion: 1.3.4](https://img.shields.io/static/v1?label=AppVersion&message=1.3.4&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-redis-exporter] Add support for Prometheus Operator Rules (#21387) 

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index 2fce50c..ba02195 100644
--- a/charts/prometheus-redis-exporter/values.yaml
+++ b/charts/prometheus-redis-exporter/values.yaml
@@ -54,6 +54,48 @@ serviceMonitor:
   # Set of labels to transfer on the Kubernetes Service onto the target.
   # targetLabels: []
 
+## Custom PrometheusRules to be defined
+## The value is evaluated as a template, so, for example, the value can depend on .Release or .Chart
+## ref: https://github.com/coreos/prometheus-operator#customresourcedefinitions
+prometheusRule:
+  enabled: false
+  additionalLabels: {}
+  namespace: ""
+  rules: []
+    ## These are just examples rules, please adapt them to your needs.
+    ## Make sure to constraint the rules to the current service.
+    #  - alert: RedisDown
+    #    expr: redis_up{service="{{ template "prometheus-redis-exporter.fullname" . }}"} == 0
+    #    for: 2m
+    #    labels:
+    #      severity: error
+    #    annotations:
+    #      summary: Redis instance {{ "{{ $labels.instance }}" }} down
+    #      description: Redis instance {{ "{{ $labels.instance }}" }} is down.
+    #  - alert: RedisMemoryHigh
+    #    expr: >
+    #       redis_memory_used_bytes{service="{{ template "prometheus-redis-exporter.fullname" . }}"} * 100
+    #       /
+    #       redis_memory_max_bytes{service="{{ template "prometheus-redis-exporter.fullname" . }}"}
+    #       > 90 <= 100
+    #    for: 2m
+    #    labels:
+    #      severity: error
+    #    annotations:
+    #      summary: Redis instance {{ "{{ $labels.instance }}" }} is using too much memory
+    #      description: |
+    #         Redis instance {{ "{{ $labels.instance }}" }} is using {{ "{{ $value }}" }}% of its available memory.
+    #  - alert: RedisKeyEviction
+    #    expr: |
+    #      increase(redis_evicted_keys_total{service="{{ template "prometheus-redis-exporter.fullname" . }}"}[5m]) > 0
+    #    for: 1s
+    #    labels:
+    #      severity: error
+    #    annotations:
+    #      summary: Redis instance {{ "{{ $labels.instance }}" }} has evicted keys
+    #      description: |
+    #        Redis instance {{ "{{ $labels.instance }}" }} has evicted {{ "{{ $value }}" }} keys in the last 5 minutes.
+
 # Used to mount a LUA-Script via config map and use it for metrics-collection
 # script:
 #   configmap: prometheus-redis-exporter-script
```

## 3.2.3 

**Release date:** 2020-03-05

![AppVersion: 1.3.4](https://img.shields.io/static/v1?label=AppVersion&message=1.3.4&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-redis-exporter] Fix secret key path in deployment (#21259) 

### Default value changes

```diff
# No changes in this release
```

## 3.2.2 

**Release date:** 2020-02-03

![AppVersion: 1.3.4](https://img.shields.io/static/v1?label=AppVersion&message=1.3.4&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-redis-exporter] Update image for exporter (#19317) 

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index aacdb14..2fce50c 100644
--- a/charts/prometheus-redis-exporter/values.yaml
+++ b/charts/prometheus-redis-exporter/values.yaml
@@ -13,7 +13,7 @@ serviceAccount:
 replicaCount: 1
 image:
   repository: oliver006/redis_exporter
-  tag: v1.0.4
+  tag: v1.3.4
   pullPolicy: IfNotPresent
 extraArgs: {}
 # Additional Environment variables
```

## 3.2.1 

**Release date:** 2019-12-23

![AppVersion: 1.0.4](https://img.shields.io/static/v1?label=AppVersion&message=1.0.4&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-redis-exporter]Add redis authoration configuration in chart (#19447) 

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index 439d8c5..aacdb14 100644
--- a/charts/prometheus-redis-exporter/values.yaml
+++ b/charts/prometheus-redis-exporter/values.yaml
@@ -58,3 +58,13 @@ serviceMonitor:
 # script:
 #   configmap: prometheus-redis-exporter-script
 #   keyname: script
+
+auth:
+  # Use password authentication
+  enabled: false
+  # Use existing secret (ignores redisPassword)
+  secret:
+    name: ""
+    key: ""
+  # Redis password (when not stored in a secret)
+  redisPassword: ""
```

## 3.2.0 

**Release date:** 2019-11-13

![AppVersion: 1.0.4](https://img.shields.io/static/v1?label=AppVersion&message=1.0.4&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/promethues-redis-exporter] Add target labels (#18818) 

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index 8bab250..439d8c5 100644
--- a/charts/prometheus-redis-exporter/values.yaml
+++ b/charts/prometheus-redis-exporter/values.yaml
@@ -51,6 +51,8 @@ serviceMonitor:
   # labels:
   # Set timeout for scrape
   # timeout: 10s
+  # Set of labels to transfer on the Kubernetes Service onto the target.
+  # targetLabels: []
 
 # Used to mount a LUA-Script via config map and use it for metrics-collection
 # script:
```

## 3.1.0 

**Release date:** 2019-10-03

![AppVersion: 1.0.4](https://img.shields.io/static/v1?label=AppVersion&message=1.0.4&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Update api versions to support k8s 1.16 (#17640) 

### Default value changes

```diff
# No changes in this release
```

## 3.0.1 

**Release date:** 2019-08-08

![AppVersion: 1.0.4](https://img.shields.io/static/v1?label=AppVersion&message=1.0.4&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-redis-exporter] Updated version of redis_exporter by oliver006 to version 1.0.4 (#16179) 

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index 0886ae5..8bab250 100644
--- a/charts/prometheus-redis-exporter/values.yaml
+++ b/charts/prometheus-redis-exporter/values.yaml
@@ -13,7 +13,7 @@ serviceAccount:
 replicaCount: 1
 image:
   repository: oliver006/redis_exporter
-  tag: v1.0.3
+  tag: v1.0.4
   pullPolicy: IfNotPresent
 extraArgs: {}
 # Additional Environment variables
```

## 3.0.0 

**Release date:** 2019-07-08

![AppVersion: 1.0.3](https://img.shields.io/static/v1?label=AppVersion&message=1.0.3&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Update image (#15217) 

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index 50bc214..0886ae5 100644
--- a/charts/prometheus-redis-exporter/values.yaml
+++ b/charts/prometheus-redis-exporter/values.yaml
@@ -13,7 +13,7 @@ serviceAccount:
 replicaCount: 1
 image:
   repository: oliver006/redis_exporter
-  tag: v0.28.0
+  tag: v1.0.3
   pullPolicy: IfNotPresent
 extraArgs: {}
 # Additional Environment variables
@@ -32,8 +32,7 @@ service:
     # prometheus.io/port: "9121"
     # prometheus.io/scrape: "true"
 resources: {}
-redisAddress:
-  - redis://myredis:6379
+redisAddress: redis://myredis:6379
 annotations: {}
 #  prometheus.io/path: /metrics
 #  prometheus.io/port: "9121"
```

## 2.0.2 

**Release date:** 2019-06-30

![AppVersion: 0.28.0](https://img.shields.io/static/v1?label=AppVersion&message=0.28.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-redis-exporter] Adding acondrat to prometheus-redis-exporter owners (#14826) 

### Default value changes

```diff
# No changes in this release
```

## 2.0.1 

**Release date:** 2019-06-22

![AppVersion: 0.28.0](https://img.shields.io/static/v1?label=AppVersion&message=0.28.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-redis-exporter] Adding abbility to use custom script (#14552) 

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index 0790d82..50bc214 100644
--- a/charts/prometheus-redis-exporter/values.yaml
+++ b/charts/prometheus-redis-exporter/values.yaml
@@ -18,11 +18,11 @@ image:
 extraArgs: {}
 # Additional Environment variables
 env: {}
-  # - name: REDIS_PASSWORD
-  #   valueFrom:
-  #     secretKeyRef:
-  #       key: redis-password
-  #       name: redis-config-0.0.2
+# - name: REDIS_PASSWORD
+#   valueFrom:
+#     secretKeyRef:
+#       key: redis-password
+#       name: redis-config-0.0.2
 service:
   type: ClusterIP
   port: 9121
@@ -52,3 +52,8 @@ serviceMonitor:
   # labels:
   # Set timeout for scrape
   # timeout: 10s
+
+# Used to mount a LUA-Script via config map and use it for metrics-collection
+# script:
+#   configmap: prometheus-redis-exporter-script
+#   keyname: script
```

## 2.0.0 

**Release date:** 2019-06-11

![AppVersion: 0.28.0](https://img.shields.io/static/v1?label=AppVersion&message=0.28.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* - Make redisAddress value a list instead of comma serparated string (#14696) 

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index 917078e..0790d82 100644
--- a/charts/prometheus-redis-exporter/values.yaml
+++ b/charts/prometheus-redis-exporter/values.yaml
@@ -32,7 +32,8 @@ service:
     # prometheus.io/port: "9121"
     # prometheus.io/scrape: "true"
 resources: {}
-redisAddress: redis://myredis:6379
+redisAddress:
+  - redis://myredis:6379
 annotations: {}
 #  prometheus.io/path: /metrics
 #  prometheus.io/port: "9121"
```

## 1.1.0 

**Release date:** 2019-06-09

![AppVersion: 0.28.0](https://img.shields.io/static/v1?label=AppVersion&message=0.28.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-redis-exporter] Add servicemonitor support (#14534) 

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index ea0869f..917078e 100644
--- a/charts/prometheus-redis-exporter/values.yaml
+++ b/charts/prometheus-redis-exporter/values.yaml
@@ -27,6 +27,7 @@ service:
   type: ClusterIP
   port: 9121
   annotations: {}
+  labels: {}
     # prometheus.io/path: /metrics
     # prometheus.io/port: "9121"
     # prometheus.io/scrape: "true"
@@ -36,3 +37,17 @@ annotations: {}
 #  prometheus.io/path: /metrics
 #  prometheus.io/port: "9121"
 #  prometheus.io/scrape: "true"
+
+serviceMonitor:
+  # When set true then use a ServiceMonitor to configure scraping
+  enabled: false
+  # Set the namespace the ServiceMonitor should be deployed
+  # namespace: monitoring
+  # Set how frequently Prometheus should scrape
+  # interval: 30s
+  # Set path to redis-exporter telemtery-path
+  # telemetryPath: /metrics
+  # Set labels for the ServiceMonitor, use this to define your scrape label for Prometheus Operator
+  # labels:
+  # Set timeout for scrape
+  # timeout: 10s
```

## 1.0.3 

**Release date:** 2019-06-05

![AppVersion: 0.28.0](https://img.shields.io/static/v1?label=AppVersion&message=0.28.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-redis-exporter] Support image pullSecrets (#14517) 

### Default value changes

```diff
# No changes in this release
```

## 1.0.2 

**Release date:** 2019-02-12

![AppVersion: 0.28.0](https://img.shields.io/static/v1?label=AppVersion&message=0.28.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Update image (#11314) 

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index f42021a..ea0869f 100644
--- a/charts/prometheus-redis-exporter/values.yaml
+++ b/charts/prometheus-redis-exporter/values.yaml
@@ -13,7 +13,7 @@ serviceAccount:
 replicaCount: 1
 image:
   repository: oliver006/redis_exporter
-  tag: v0.25.0
+  tag: v0.28.0
   pullPolicy: IfNotPresent
 extraArgs: {}
 # Additional Environment variables
```

## 1.0.1 

**Release date:** 2019-01-02

![AppVersion: 0.25.0](https://img.shields.io/static/v1?label=AppVersion&message=0.25.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Update image to 0.25.0 (#10272) 

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index fdc93ae..f42021a 100644
--- a/charts/prometheus-redis-exporter/values.yaml
+++ b/charts/prometheus-redis-exporter/values.yaml
@@ -13,7 +13,7 @@ serviceAccount:
 replicaCount: 1
 image:
   repository: oliver006/redis_exporter
-  tag: v0.22.1
+  tag: v0.25.0
   pullPolicy: IfNotPresent
 extraArgs: {}
 # Additional Environment variables
```

## 1.0.0 

**Release date:** 2018-12-02

![AppVersion: 0.22.1](https://img.shields.io/static/v1?label=AppVersion&message=0.22.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Update image to 0.22.1, and chart to 1.0.0 because stable (#9683) 

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index deec244..fdc93ae 100644
--- a/charts/prometheus-redis-exporter/values.yaml
+++ b/charts/prometheus-redis-exporter/values.yaml
@@ -13,7 +13,7 @@ serviceAccount:
 replicaCount: 1
 image:
   repository: oliver006/redis_exporter
-  tag: v0.21.2
+  tag: v0.22.1
   pullPolicy: IfNotPresent
 extraArgs: {}
 # Additional Environment variables
@@ -27,6 +27,9 @@ service:
   type: ClusterIP
   port: 9121
   annotations: {}
+    # prometheus.io/path: /metrics
+    # prometheus.io/port: "9121"
+    # prometheus.io/scrape: "true"
 resources: {}
 redisAddress: redis://myredis:6379
 annotations: {}
```

## 0.3.4 

**Release date:** 2018-11-06

![AppVersion: 0.21.2](https://img.shields.io/static/v1?label=AppVersion&message=0.21.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Updating image of prometheus-redis-exporter to latest version (#8994) 

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index 62c6f82..deec244 100644
--- a/charts/prometheus-redis-exporter/values.yaml
+++ b/charts/prometheus-redis-exporter/values.yaml
@@ -13,7 +13,7 @@ serviceAccount:
 replicaCount: 1
 image:
   repository: oliver006/redis_exporter
-  tag: v0.21.1
+  tag: v0.21.2
   pullPolicy: IfNotPresent
 extraArgs: {}
 # Additional Environment variables
```

## 0.3.3 

**Release date:** 2018-11-04

![AppVersion: 0.21.1](https://img.shields.io/static/v1?label=AppVersion&message=0.21.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Set redis-exporter name for port in service (#6547) 

### Default value changes

```diff
# No changes in this release
```

## 0.3.2 

**Release date:** 2018-09-11

![AppVersion: 0.21.1](https://img.shields.io/static/v1?label=AppVersion&message=0.21.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Update image (#7630) 

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index 9f34cc1..62c6f82 100644
--- a/charts/prometheus-redis-exporter/values.yaml
+++ b/charts/prometheus-redis-exporter/values.yaml
@@ -13,7 +13,7 @@ serviceAccount:
 replicaCount: 1
 image:
   repository: oliver006/redis_exporter
-  tag: v0.16.0
+  tag: v0.21.1
   pullPolicy: IfNotPresent
 extraArgs: {}
 # Additional Environment variables
```

## 0.3.1 

**Release date:** 2018-09-02

![AppVersion: 0.16.0](https://img.shields.io/static/v1?label=AppVersion&message=0.16.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* adding additional env variable for redis (#7482) 

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index a888120..9f34cc1 100644
--- a/charts/prometheus-redis-exporter/values.yaml
+++ b/charts/prometheus-redis-exporter/values.yaml
@@ -16,6 +16,13 @@ image:
   tag: v0.16.0
   pullPolicy: IfNotPresent
 extraArgs: {}
+# Additional Environment variables
+env: {}
+  # - name: REDIS_PASSWORD
+  #   valueFrom:
+  #     secretKeyRef:
+  #       key: redis-password
+  #       name: redis-config-0.0.2
 service:
   type: ClusterIP
   port: 9121
```

## 0.3.0 

**Release date:** 2018-07-25

![AppVersion: 0.16.0](https://img.shields.io/static/v1?label=AppVersion&message=0.16.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* add option to supply extra arguments (#6516) 

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index adecdb9..a888120 100644
--- a/charts/prometheus-redis-exporter/values.yaml
+++ b/charts/prometheus-redis-exporter/values.yaml
@@ -15,6 +15,7 @@ image:
   repository: oliver006/redis_exporter
   tag: v0.16.0
   pullPolicy: IfNotPresent
+extraArgs: {}
 service:
   type: ClusterIP
   port: 9121
```

## 0.2.0 

**Release date:** 2018-06-17

![AppVersion: 0.16.0](https://img.shields.io/static/v1?label=AppVersion&message=0.16.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-redis-exporter] add RBAC resources (#6083) 

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index 76840cf..adecdb9 100644
--- a/charts/prometheus-redis-exporter/values.yaml
+++ b/charts/prometheus-redis-exporter/values.yaml
@@ -1,3 +1,15 @@
+rbac:
+  # Specifies whether RBAC resources should be created
+  create: true
+  # Specifies whether a PodSecurityPolicy should be created
+  pspEnabled: true
+serviceAccount:
+  # Specifies whether a ServiceAccount should be created
+  create: true
+  # The name of the ServiceAccount to use.
+  # If not set and create is true, a name is generated using the fullname template
+  name:
+
 replicaCount: 1
 image:
   repository: oliver006/redis_exporter
```

## 0.1.2 

**Release date:** 2018-05-25

![AppVersion: 0.16.0](https://img.shields.io/static/v1?label=AppVersion&message=0.16.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-redis-exporter] typo fix: tables lists->table lists (#5721) 

### Default value changes

```diff
# No changes in this release
```

## 0.1.1 

**Release date:** 2018-05-02

![AppVersion: 0.16.0](https://img.shields.io/static/v1?label=AppVersion&message=0.16.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Added prometheus scrape annotations (#5256) 

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index e00b48d..76840cf 100644
--- a/charts/prometheus-redis-exporter/values.yaml
+++ b/charts/prometheus-redis-exporter/values.yaml
@@ -9,3 +9,7 @@ service:
   annotations: {}
 resources: {}
 redisAddress: redis://myredis:6379
+annotations: {}
+#  prometheus.io/path: /metrics
+#  prometheus.io/port: "9121"
+#  prometheus.io/scrape: "true"
```

## 0.1.0 

**Release date:** 2018-04-08

![AppVersion: 0.16.0](https://img.shields.io/static/v1?label=AppVersion&message=0.16.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* adding redis-exporter chart (#2743) 

### Default value changes

```diff
replicaCount: 1
image:
  repository: oliver006/redis_exporter
  tag: v0.16.0
  pullPolicy: IfNotPresent
service:
  type: ClusterIP
  port: 9121
  annotations: {}
resources: {}
redisAddress: redis://myredis:6379
```

---
Autogenerated from Helm Chart and git history using [helm-changelog](https://github.com/mogensen/helm-changelog)
