# Change Log

## 5.3.2

**Release date:** 2023-04-13

![AppVersion: v1.44.0](https://img.shields.io/static/v1?label=AppVersion&message=v1.44.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-redis-exporter] explicit default for metricRelabelings[].action (#3224)

### Default value changes

```diff
# No changes in this release
```

## 5.3.1

**Release date:** 2023-03-22

![AppVersion: v1.44.0](https://img.shields.io/static/v1?label=AppVersion&message=v1.44.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-redis-exporter] Added additionalLabels for service monitor (#3121)

### Default value changes

```diff
# No changes in this release
```

## 5.3.0

**Release date:** 2022-11-08

![AppVersion: v1.44.0](https://img.shields.io/static/v1?label=AppVersion&message=v1.44.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* feat(deployment): support tls config using secrets (#2660)

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index 7ad918a1..d0ff352f 100644
--- a/charts/prometheus-redis-exporter/values.yaml
+++ b/charts/prometheus-redis-exporter/values.yaml
@@ -66,6 +66,40 @@ redisAddressConfig:
     name: ""
     key: ""
 
+redisTlsConfig:
+  # Use TLS configuration
+  enabled: false
+  # Whether to to skip TLS verification
+  skipTlsVerification: false
+  # All secrets key about TLS config will be mounted into this path
+  mountPath: /tls
+
+  # REDIS_EXPORTER_TLS_CA_CERT_FILE will be set to /tls/tls-ca-cert.crt
+  caCertFile:
+    secret:
+      name: ""
+      key: ""
+  # REDIS_EXPORTER_TLS_CLIENT_KEY_FILE  will be set to /tls/tls-client-key.key
+  clientKeyFile:
+    secret:
+      name: ""
+      key: ""
+  # REDIS_EXPORTER_TLS_CLIENT_CERT_FILE will be set to /tls/tls-client-cert.crt
+  clientCertFile:
+    secret:
+      name: ""
+      key: ""
+  # REDIS_EXPORTER_TLS_SERVER_KEY_FILE will be set to /tls/tls-server-key.key
+  serverKeyFile:
+    secret:
+      name: ""
+      key: ""
+  # REDIS_EXPORTER_TLS_SERVER_CERT_FILE will be set to /tls/tls-server-cert.crt
+  serverCertFile:
+    secret:
+      name: ""
+      key: ""
+
 serviceMonitor:
   # When set true then use a ServiceMonitor to configure scraping
   enabled: false

```

## 5.2.1

**Release date:** 2022-10-16

![AppVersion: v1.44.0](https://img.shields.io/static/v1?label=AppVersion&message=v1.44.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-redis-exporter] Fix PSP deprecation after k8s 1.25+ (#2572)

### Default value changes

```diff
# No changes in this release
```

## 5.2.0

**Release date:** 2022-09-26

![AppVersion: v1.44.0](https://img.shields.io/static/v1?label=AppVersion&message=v1.44.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometehus-redis-exporter] update exporter to 1.44 & use appVersion as image tag (#2490)

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index 97d0320e..7ad918a1 100644
--- a/charts/prometheus-redis-exporter/values.yaml
+++ b/charts/prometheus-redis-exporter/values.yaml
@@ -13,8 +13,10 @@ serviceAccount:
 replicaCount: 1
 image:
   repository: oliver006/redis_exporter
-  tag: v1.27.0
   pullPolicy: IfNotPresent
+  # Overrides the image tag whose default is the chart appVersion.
+  tag: ""
+
 extraArgs: {}
 
 # global custom labels, applied to all resrouces

```

## 5.1.0

**Release date:** 2022-08-28

![AppVersion: 1.43.0](https://img.shields.io/static/v1?label=AppVersion&message=1.43.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-redis-exporter] support redis password file secret volume mount (#2401)

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index e7df3a2a..97d0320e 100644
--- a/charts/prometheus-redis-exporter/values.yaml
+++ b/charts/prometheus-redis-exporter/values.yaml
@@ -162,3 +162,14 @@ auth:
   redisPassword: ""
   # Redis user (version 6.X and above)
   redisUser: ""
+  # Redis password file (e.g., https://github.com/oliver006/redis_exporter/blob/v1.27.0/contrib/sample-pwd-file.json)
+  # secret (useful for multiple redis instances with different passwords). If secret name and key are set
+  # this will ignore the single password auth.secret.*
+  redisPasswordFile:
+    # The secret key will be mounted into this path as a file
+    # e.g., if secret key is pass.json, the env variable
+    # REDIS_PASSWORD_FILE will be set to /auth/pass.json
+    mountPath: /auth
+    secret:
+      name: ""
+      key: ""

```

## 5.0.0

**Release date:** 2022-06-30

![AppVersion: 1.43.0](https://img.shields.io/static/v1?label=AppVersion&message=1.43.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-redis-exporter] support multiple targets (#2199)

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index e94e3fda..e7df3a2a 100644
--- a/charts/prometheus-redis-exporter/values.yaml
+++ b/charts/prometheus-redis-exporter/values.yaml
@@ -17,6 +17,7 @@ image:
   pullPolicy: IfNotPresent
 extraArgs: {}
 
+# global custom labels, applied to all resrouces
 customLabels: {}
 
 securityContext: {}
@@ -45,7 +46,10 @@ tolerations: []
 
 affinity: {}
 
+# If serviceMonitor.multipleTarget is enabled, this configuration is actually not used
 redisAddress: redis://myredis:6379
+
+# deployment additional annotations and labels
 annotations: {}
 labels: {}
 #  prometheus.io/path: /metrics
@@ -63,11 +67,29 @@ redisAddressConfig:
 serviceMonitor:
   # When set true then use a ServiceMonitor to configure scraping
   enabled: false
+  multipleTarget: false
+  targets: []
+  # for every targets, url and name must be set,
+  # an individual additionalRelabeling can be set for every target
+  # - url: "redis://myredis:6379"
+  #   name: "my-redis"
+  # - url: "redis://my-redis-cluster:6379"
+  #   name: "bar"
+  #   additionalRelabeling:
+  #   - sourceLabels: [type]
+  #     targetLabel: type
+  #     replacement: cluster
+  #   additionalMetricsRelabels:
+  #     type: cluster
+  additionalMetricsRelabels: {}
+  additionalRelabeling: []
+
   # Set the namespace the ServiceMonitor should be deployed
   # namespace: monitoring
   # Set how frequently Prometheus should scrape
   # interval: 30s
   # Set path to redis-exporter telemtery-path
+  # Please set telemetryPath to /scrape if you are using multiple targets
   # telemetryPath: /metrics
   # Set labels for the ServiceMonitor, use this to define your scrape label for Prometheus Operator
   # labels:

```

## 4.8.0

**Release date:** 2022-06-15

![AppVersion: 1.27.0](https://img.shields.io/static/v1?label=AppVersion&message=1.27.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-redis-exporter] Add tls config options for servicemonitor (#2158)

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index 060ec474..e94e3fda 100644
--- a/charts/prometheus-redis-exporter/values.yaml
+++ b/charts/prometheus-redis-exporter/values.yaml
@@ -31,6 +31,7 @@ env: {}
 service:
   type: ClusterIP
   port: 9121
+  portName: redis-exporter
   annotations: {}
   labels: {}
     # prometheus.io/path: /metrics
@@ -77,6 +78,9 @@ serviceMonitor:
   # Set of labels to transfer on the Kubernetes Service onto the target.
   # targetLabels: []
   # metricRelabelings: []
+  # Set tls options
+  # scheme: ""
+  # tlsConfig: {}
 
 ## Custom PrometheusRules to be defined
 ## The value is evaluated as a template, so, for example, the value can depend on .Release or .Chart

```

## 4.7.0

**Release date:** 2022-05-17

![AppVersion: 1.27.0](https://img.shields.io/static/v1?label=AppVersion&message=1.27.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-redis-exporter] Add support for REDIS_USER (#2063)

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index 658029e9..060ec474 100644
--- a/charts/prometheus-redis-exporter/values.yaml
+++ b/charts/prometheus-redis-exporter/values.yaml
@@ -134,3 +134,5 @@ auth:
     key: ""
   # Redis password (when not stored in a secret)
   redisPassword: ""
+  # Redis user (version 6.X and above)
+  redisUser: ""

```

## 4.6.0

**Release date:** 2021-09-10

![AppVersion: 1.27.0](https://img.shields.io/static/v1?label=AppVersion&message=1.27.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-redis-exporter] allow using securityContext in containers (#1329)

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index d39d8a7b..658029e9 100644
--- a/charts/prometheus-redis-exporter/values.yaml
+++ b/charts/prometheus-redis-exporter/values.yaml
@@ -19,6 +19,8 @@ extraArgs: {}
 
 customLabels: {}
 
+securityContext: {}
+
 # Additional Environment variables
 env: {}
 # - name: REDIS_PASSWORD

```

## 4.5.0

**Release date:** 2021-09-09

![AppVersion: 1.27.0](https://img.shields.io/static/v1?label=AppVersion&message=1.27.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* allow using annotations in deployment (#1325)

### Default value changes

```diff
# No changes in this release
```

## 4.4.0

**Release date:** 2021-09-07

![AppVersion: 1.27.0](https://img.shields.io/static/v1?label=AppVersion&message=1.27.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* allow using custom labels (#1322)

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index 2dc3548d..d39d8a7b 100644
--- a/charts/prometheus-redis-exporter/values.yaml
+++ b/charts/prometheus-redis-exporter/values.yaml
@@ -16,6 +16,9 @@ image:
   tag: v1.27.0
   pullPolicy: IfNotPresent
 extraArgs: {}
+
+customLabels: {}
+
 # Additional Environment variables
 env: {}
 # - name: REDIS_PASSWORD

```

## 4.3.0

**Release date:** 2021-08-31

![AppVersion: 1.27.0](https://img.shields.io/static/v1?label=AppVersion&message=1.27.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-redis-exporter] Version bump to 1.27.0 (#1297)

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index 5b28ff50..2dc3548d 100644
--- a/charts/prometheus-redis-exporter/values.yaml
+++ b/charts/prometheus-redis-exporter/values.yaml
@@ -13,7 +13,7 @@ serviceAccount:
 replicaCount: 1
 image:
   repository: oliver006/redis_exporter
-  tag: v1.11.1
+  tag: v1.27.0
   pullPolicy: IfNotPresent
 extraArgs: {}
 # Additional Environment variables

```

## 4.2.0

**Release date:** 2021-07-23

![AppVersion: 1.11.1](https://img.shields.io/static/v1?label=AppVersion&message=1.11.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-redis-exporter] address via configmap (#1203)

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index ccbdd43f..5b28ff50 100644
--- a/charts/prometheus-redis-exporter/values.yaml
+++ b/charts/prometheus-redis-exporter/values.yaml
@@ -46,6 +46,14 @@ labels: {}
 #  prometheus.io/port: "9121"
 #  prometheus.io/scrape: "true"
 
+redisAddressConfig:
+  # Use config from configmap
+  enabled: false
+  # Use existing configmap (ignores redisAddress)
+  configmap:
+    name: ""
+    key: ""
+
 serviceMonitor:
   # When set true then use a ServiceMonitor to configure scraping
   enabled: false

```

## 4.1.0

**Release date:** 2021-06-23

![AppVersion: 1.11.1](https://img.shields.io/static/v1?label=AppVersion&message=1.11.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-redis-exporter] Add custom labels support (#1102)

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index a65310a8..ccbdd43f 100644
--- a/charts/prometheus-redis-exporter/values.yaml
+++ b/charts/prometheus-redis-exporter/values.yaml
@@ -41,6 +41,7 @@ affinity: {}
 
 redisAddress: redis://myredis:6379
 annotations: {}
+labels: {}
 #  prometheus.io/path: /metrics
 #  prometheus.io/port: "9121"
 #  prometheus.io/scrape: "true"

```

## 4.0.2

**Release date:** 2021-04-20

![AppVersion: 1.11.1](https://img.shields.io/static/v1?label=AppVersion&message=1.11.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* prometheus-redis-exporter : add relabeling configs for redis exporter (#866)

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index 967184e3..a65310a8 100644
--- a/charts/prometheus-redis-exporter/values.yaml
+++ b/charts/prometheus-redis-exporter/values.yaml
@@ -58,6 +58,8 @@ serviceMonitor:
   # labels:
   # Set timeout for scrape
   # timeout: 10s
+  # Set relabel_configs as per https://prometheus.io/docs/prometheus/latest/configuration/configuration/#relabel_config
+  # relabelings: []
   # Set of labels to transfer on the Kubernetes Service onto the target.
   # targetLabels: []
   # metricRelabelings: []

```

## 4.0.1

**Release date:** 2021-04-12

![AppVersion: 1.11.1](https://img.shields.io/static/v1?label=AppVersion&message=1.11.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* fix readme (#831)

### Default value changes

```diff
# No changes in this release
```

## 4.0.0

**Release date:** 2020-11-23

![AppVersion: 1.11.1](https://img.shields.io/static/v1?label=AppVersion&message=1.11.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-redis-exporter] Migrate to chart v2 (#389)

### Default value changes

```diff
# No changes in this release
```

## 3.7.0

**Release date:** 2020-11-18

![AppVersion: 1.11.1](https://img.shields.io/static/v1?label=AppVersion&message=1.11.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* feat(prometheus-redis-exporter): add configurable apiVersion (#369)

### Default value changes

```diff
# No changes in this release
```

## 3.6.0

**Release date:** 2020-09-12

![AppVersion: 1.11.1](https://img.shields.io/static/v1?label=AppVersion&message=1.11.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-redis-exporter] Upgrade exporter image to version 1.11.1 (#87)

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index 68151d3a..967184e3 100644
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

![AppVersion: 1.3.4](https://img.shields.io/static/v1?label=AppVersion&message=1.3.4&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-redis-exporter] Add zanhsieh as maintainer (#46)

### Default value changes

```diff
# No changes in this release
```

## 3.5.1

**Release date:** 2020-08-20

![AppVersion: 1.3.4](https://img.shields.io/static/v1?label=AppVersion&message=1.3.4&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Prep initial charts indexing (#14)

### Default value changes

```diff
# No changes in this release
```

## 3.5.0

**Release date:** 2020-07-22

![AppVersion: 1.3.4](https://img.shields.io/static/v1?label=AppVersion&message=1.3.4&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-redis-exporter] Add ServiceMonitor metricRelabelings (#23289)

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index c19f7a4a..68151d3a 100644
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

![AppVersion: 1.3.4](https://img.shields.io/static/v1?label=AppVersion&message=1.3.4&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-redis-exporter] fix targetLabels (#22541)

### Default value changes

```diff
# No changes in this release
```

## 3.4.0

**Release date:** 2020-04-09

![AppVersion: 1.3.4](https://img.shields.io/static/v1?label=AppVersion&message=1.3.4&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-redis-exporter] Add support for tolerations and affinity (#21848)

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index ba021953..c19f7a4a 100644
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

![AppVersion: 1.3.4](https://img.shields.io/static/v1?label=AppVersion&message=1.3.4&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-redis-exporter] Add support for Prometheus Operator Rules (#21387)

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index 2fce50cd..ba021953 100644
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

![AppVersion: 1.3.4](https://img.shields.io/static/v1?label=AppVersion&message=1.3.4&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-redis-exporter] Fix secret key path in deployment (#21259)

### Default value changes

```diff
# No changes in this release
```

## 3.2.2

**Release date:** 2020-02-04

![AppVersion: 1.3.4](https://img.shields.io/static/v1?label=AppVersion&message=1.3.4&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-redis-exporter] Update image for exporter (#19317)

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index aacdb14e..2fce50cd 100644
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

![AppVersion: 1.0.4](https://img.shields.io/static/v1?label=AppVersion&message=1.0.4&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-redis-exporter]Add redis authoration configuration in chart (#19447)

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index 439d8c5c..aacdb14e 100644
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

![AppVersion: 1.0.4](https://img.shields.io/static/v1?label=AppVersion&message=1.0.4&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/promethues-redis-exporter] Add target labels (#18818)

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index 8bab250b..439d8c5c 100644
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

![AppVersion: 1.0.4](https://img.shields.io/static/v1?label=AppVersion&message=1.0.4&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Update api versions to support k8s 1.16 (#17640)

### Default value changes

```diff
# No changes in this release
```

## 3.0.1

**Release date:** 2019-08-08

![AppVersion: 1.0.4](https://img.shields.io/static/v1?label=AppVersion&message=1.0.4&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-redis-exporter] Updated version of redis_exporter by oliver006 to version 1.0.4 (#16179)

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index 0886ae5d..8bab250b 100644
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

![AppVersion: 1.0.3](https://img.shields.io/static/v1?label=AppVersion&message=1.0.3&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Update image (#15217)

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index 50bc214c..0886ae5d 100644
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

![AppVersion: 0.28.0](https://img.shields.io/static/v1?label=AppVersion&message=0.28.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-redis-exporter] Adding acondrat to prometheus-redis-exporter owners (#14826)

### Default value changes

```diff
# No changes in this release
```

## 2.0.1

**Release date:** 2019-06-22

![AppVersion: 0.28.0](https://img.shields.io/static/v1?label=AppVersion&message=0.28.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-redis-exporter] Adding abbility to use custom script (#14552)

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index 0790d82d..50bc214c 100644
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

![AppVersion: 0.28.0](https://img.shields.io/static/v1?label=AppVersion&message=0.28.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* - Make redisAddress value a list instead of comma serparated string (#14696)

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index 917078e4..0790d82d 100644
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

![AppVersion: 0.28.0](https://img.shields.io/static/v1?label=AppVersion&message=0.28.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-redis-exporter] Add servicemonitor support (#14534)

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index ea0869f1..917078e4 100644
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

![AppVersion: 0.28.0](https://img.shields.io/static/v1?label=AppVersion&message=0.28.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-redis-exporter] Support image pullSecrets (#14517)

### Default value changes

```diff
# No changes in this release
```

## 1.0.2

**Release date:** 2019-02-12

![AppVersion: 0.28.0](https://img.shields.io/static/v1?label=AppVersion&message=0.28.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Update image (#11314)

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index f42021a8..ea0869f1 100644
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

![AppVersion: 0.25.0](https://img.shields.io/static/v1?label=AppVersion&message=0.25.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Update image to 0.25.0 (#10272)

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index fdc93ae5..f42021a8 100644
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

![AppVersion: 0.22.1](https://img.shields.io/static/v1?label=AppVersion&message=0.22.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Update image to 0.22.1, and chart to 1.0.0 because stable (#9683)

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index deec2444..fdc93ae5 100644
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

![AppVersion: 0.21.2](https://img.shields.io/static/v1?label=AppVersion&message=0.21.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Updating image of prometheus-redis-exporter to latest version (#8994)

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index 62c6f824..deec2444 100644
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

![AppVersion: 0.21.1](https://img.shields.io/static/v1?label=AppVersion&message=0.21.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Set redis-exporter name for port in service (#6547)

### Default value changes

```diff
# No changes in this release
```

## 0.3.2

**Release date:** 2018-09-11

![AppVersion: 0.21.1](https://img.shields.io/static/v1?label=AppVersion&message=0.21.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Update image (#7630)

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index 9f34cc15..62c6f824 100644
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

![AppVersion: 0.16.0](https://img.shields.io/static/v1?label=AppVersion&message=0.16.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* adding additional env variable for redis (#7482)

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index a8881202..9f34cc15 100644
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

![AppVersion: 0.16.0](https://img.shields.io/static/v1?label=AppVersion&message=0.16.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* add option to supply extra arguments (#6516)

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index adecdb98..a8881202 100644
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

![AppVersion: 0.16.0](https://img.shields.io/static/v1?label=AppVersion&message=0.16.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-redis-exporter] add RBAC resources (#6083)

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index 76840cf8..adecdb98 100644
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

![AppVersion: 0.16.0](https://img.shields.io/static/v1?label=AppVersion&message=0.16.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-redis-exporter] typo fix: tables lists->table lists (#5721)

### Default value changes

```diff
# No changes in this release
```

## 0.1.1

**Release date:** 2018-05-02

![AppVersion: 0.16.0](https://img.shields.io/static/v1?label=AppVersion&message=0.16.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Added prometheus scrape annotations (#5256)

### Default value changes

```diff
diff --git a/charts/prometheus-redis-exporter/values.yaml b/charts/prometheus-redis-exporter/values.yaml
index e00b48d3..76840cf8 100644
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

![AppVersion: 0.16.0](https://img.shields.io/static/v1?label=AppVersion&message=0.16.0&color=success)
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
