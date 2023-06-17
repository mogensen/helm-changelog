# Change Log

## 2.2.0

**Release date:** 2023-06-01

![AppVersion: v1.6.0](https://img.shields.io/static/v1?label=AppVersion&message=v1.6.0&color=success)
![Kubernetes: >=1.19.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.19.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-kafka-exporter] Change chart PSP default state to off (#3127)

### Default value changes

```diff
diff --git a/charts/prometheus-kafka-exporter/values.yaml b/charts/prometheus-kafka-exporter/values.yaml
index 5e52908b..f7cdf6bc 100644
--- a/charts/prometheus-kafka-exporter/values.yaml
+++ b/charts/prometheus-kafka-exporter/values.yaml
@@ -39,7 +39,7 @@ rbac:
   # Specifies whether RBAC resources should be created
   create: true
   # Specifies whether a PodSecurityPolicy should be created
-  pspEnabled: true
+  pspEnabled: false
 
 serviceAccount:
   # Specifies whether a ServiceAccount should be created

```

## 2.1.0

**Release date:** 2023-05-04

![AppVersion: v1.6.0](https://img.shields.io/static/v1?label=AppVersion&message=v1.6.0&color=success)
![Kubernetes: >=1.19.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.19.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-kafka-exporter] Add imagePullSecrets to kafka-exporter (#3291)

### Default value changes

```diff
diff --git a/charts/prometheus-kafka-exporter/values.yaml b/charts/prometheus-kafka-exporter/values.yaml
index f33dc2ed..5e52908b 100644
--- a/charts/prometheus-kafka-exporter/values.yaml
+++ b/charts/prometheus-kafka-exporter/values.yaml
@@ -4,6 +4,22 @@ image:
   tag: ""
   pullPolicy: IfNotPresent
 
+imagePullSecrets: []
+
+global:
+  # To help compatibility with other charts which use global.imagePullSecrets.
+  # Allow either an array of {name: pullSecret} maps (k8s-style), or an array of strings (more common helm-style).
+  # global:
+  #   imagePullSecrets:
+  #   - name: pullSecret1
+  #   - name: pullSecret2
+  # or
+  # global:
+  #   imagePullSecrets:
+  #   - pullSecret1
+  #   - pullSecret2
+  imagePullSecrets: []
+
 replicaCount: 1
 
 kafkaServer:

```

## 2.0.0

**Release date:** 2023-05-03

![AppVersion: v1.6.0](https://img.shields.io/static/v1?label=AppVersion&message=v1.6.0&color=success)
![Kubernetes: >=1.19.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.19.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-kafka-exporter] fix Chart.yaml - remove engine: gotpl (#3199)

### Default value changes

```diff
# No changes in this release
```

## 1.8.1

**Release date:** 2023-05-01

![AppVersion: v1.6.0](https://img.shields.io/static/v1?label=AppVersion&message=v1.6.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-kafka-exporter] Add a maintainer (#3317)

### Default value changes

```diff
# No changes in this release
```

## 1.8.0

**Release date:** 2023-01-09

![AppVersion: v1.6.0](https://img.shields.io/static/v1?label=AppVersion&message=v1.6.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-kafka-exporter] Add support for securityContext (#2903)

### Default value changes

```diff
diff --git a/charts/prometheus-kafka-exporter/values.yaml b/charts/prometheus-kafka-exporter/values.yaml
index 941825f6..f33dc2ed 100644
--- a/charts/prometheus-kafka-exporter/values.yaml
+++ b/charts/prometheus-kafka-exporter/values.yaml
@@ -143,8 +143,22 @@ server:
 kafka:
   enabled: false
 
-# Set security context for the kafka-exporter container. Useful when you need to adapt to an existing PSP
+# Set security context for the kafka-exporter pod. Useful when you need to adapt to an existing PSP
 securityContext: {}
-  # runAsUser:
-  # runAsGroup:
-  # fsGroup:
+  # fsGroup: 2000
+  # runAsGroup: 10002
+  # runAsUser: 10001
+  # seccompProfile:
+  #   type: RuntimeDefault
+
+# Set securityContext for the kafka-exporter container
+containerSecurityContext: {}
+  # allowPrivilegeEscalation: false
+  # capabilities:
+  #   drop: ["all"]
+  # readOnlyRootFilesystem: true
+  # runAsGroup: 10002
+  # runAsNonRoot: true
+  # runAsUser: 10001
+  # seccompProfile:
+  #   type: RuntimeDefault

```

## 1.7.0

**Release date:** 2022-10-26

![AppVersion: v1.6.0](https://img.shields.io/static/v1?label=AppVersion&message=v1.6.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-kafka-exporter] Bump kafka-exporter to 1.6.0 (#2613)

### Default value changes

```diff
diff --git a/charts/prometheus-kafka-exporter/values.yaml b/charts/prometheus-kafka-exporter/values.yaml
index d76f76c8..941825f6 100644
--- a/charts/prometheus-kafka-exporter/values.yaml
+++ b/charts/prometheus-kafka-exporter/values.yaml
@@ -1,6 +1,7 @@
 image:
   repository: danielqsj/kafka-exporter
-  tag: v1.4.2
+  # Overrides the image tag whose default is the chart appVersion.
+  tag: ""
   pullPolicy: IfNotPresent
 
 replicaCount: 1

```

## 1.6.1

**Release date:** 2022-10-16

![AppVersion: v1.4.2](https://img.shields.io/static/v1?label=AppVersion&message=v1.4.2&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-kafka-exporter] Fix PSP deprecation after k8s 1.25+ (#2575)

### Default value changes

```diff
# No changes in this release
```

## 1.6.0

**Release date:** 2022-06-11

![AppVersion: v1.4.2](https://img.shields.io/static/v1?label=AppVersion&message=v1.4.2&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Bump kafka-exporter to 1.4.2, bump bitnami kafka (#2116)

### Default value changes

```diff
diff --git a/charts/prometheus-kafka-exporter/values.yaml b/charts/prometheus-kafka-exporter/values.yaml
index b8962a11..d76f76c8 100644
--- a/charts/prometheus-kafka-exporter/values.yaml
+++ b/charts/prometheus-kafka-exporter/values.yaml
@@ -1,6 +1,6 @@
 image:
   repository: danielqsj/kafka-exporter
-  tag: v1.4.1
+  tag: v1.4.2
   pullPolicy: IfNotPresent
 
 replicaCount: 1
@@ -102,6 +102,8 @@ affinity: {}
 tls:
   enabled: false
   insecureSkipVerify: false
+  # The kafka server's name. Used to verify the hostname on the returned certificates unless tls.insecureSkipVerify is set to true.
+  serverName: ""
   # mountPath: /secret/certs
   # secretName: <name of an existing secret>
 
@@ -127,6 +129,14 @@ sasl:
     # mountPath: /secret/kerberos
     # secretName: <name of an existing secret>
 
+# This enables TLS for web server
+server:
+  tls:
+    enabled: false
+    mutualAuthEnabled: false
+  # mountPath: /secret/certs
+  # secretName: <name of an existing secret>
+
 # If enabled Kafka dependency chart will be used.
 # This is only needed for the CI job so it should always be disabled.
 kafka:

```

## 1.5.0

**Release date:** 2021-09-18

![AppVersion: v1.4.1](https://img.shields.io/static/v1?label=AppVersion&message=v1.4.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-kafka-exporter] - added SASL support and upgrade to latest version (#1349)

### Default value changes

```diff
diff --git a/charts/prometheus-kafka-exporter/values.yaml b/charts/prometheus-kafka-exporter/values.yaml
index c5a977d1..b8962a11 100644
--- a/charts/prometheus-kafka-exporter/values.yaml
+++ b/charts/prometheus-kafka-exporter/values.yaml
@@ -1,6 +1,6 @@
 image:
   repository: danielqsj/kafka-exporter
-  tag: v1.3.1
+  tag: v1.4.1
   pullPolicy: IfNotPresent
 
 replicaCount: 1
@@ -11,8 +11,8 @@ kafkaServer:
 # Specifies broker version to use, leave empty for default
 kafkaBrokerVersion:
 
-# Valid levels: [debug, info, warn, error, fatal]
-logLevel: info
+# Specifies log verbosity
+verbosity: 0
 
 # Sarama logging
 sarama:
@@ -32,6 +32,8 @@ serviceAccount:
   name:
 
 env: []
+# - name: <ENV_VAR_NAME>
+#   value: <value>
 
 # List of additional cli arguments to configure kafka-exporter
 # for example: --log.enable-sarama, --log.level=debug, etc.
@@ -103,6 +105,28 @@ tls:
   # mountPath: /secret/certs
   # secretName: <name of an existing secret>
 
+sasl:
+  enabled: false
+  handshake: true
+  scram:
+    enabled: false
+    # mechanism: <plain|scram-sha512|scram-sha256>
+
+    # add username and password
+    # username:
+    # password:
+
+    # or use an existing secret
+    # secretName: <name of an existing secret>
+
+  kerberos:
+    enabled: false
+    # serviceName:
+    # realm:
+    # kerberosAuthType: <keytabAuth|userAuth>
+    # mountPath: /secret/kerberos
+    # secretName: <name of an existing secret>
+
 # If enabled Kafka dependency chart will be used.
 # This is only needed for the CI job so it should always be disabled.
 kafka:

```

## 1.4.0

**Release date:** 2021-09-14

![AppVersion: v1.3.1](https://img.shields.io/static/v1?label=AppVersion&message=v1.3.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-kafka-exporter] Add: securitycontext to deployment.yaml (#1334)

### Default value changes

```diff
diff --git a/charts/prometheus-kafka-exporter/values.yaml b/charts/prometheus-kafka-exporter/values.yaml
index ed32c7ba..c5a977d1 100644
--- a/charts/prometheus-kafka-exporter/values.yaml
+++ b/charts/prometheus-kafka-exporter/values.yaml
@@ -107,3 +107,9 @@ tls:
 # This is only needed for the CI job so it should always be disabled.
 kafka:
   enabled: false
+
+# Set security context for the kafka-exporter container. Useful when you need to adapt to an existing PSP
+securityContext: {}
+  # runAsUser:
+  # runAsGroup:
+  # fsGroup:

```

## 1.3.0

**Release date:** 2021-09-02

![AppVersion: v1.3.1](https://img.shields.io/static/v1?label=AppVersion&message=v1.3.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-kafka-exporter] Kafka exporter chart not propagating service labels (#1300)

### Default value changes

```diff
diff --git a/charts/prometheus-kafka-exporter/values.yaml b/charts/prometheus-kafka-exporter/values.yaml
index 894e072d..ed32c7ba 100644
--- a/charts/prometheus-kafka-exporter/values.yaml
+++ b/charts/prometheus-kafka-exporter/values.yaml
@@ -41,6 +41,7 @@ extraArgs: []
 service:
   type: ClusterIP
   port: 9308
+  labels: {}
   annotations: {}
 
 liveness:
@@ -69,6 +70,7 @@ prometheus:
     # additionalLabels:
     #   release: kube-prometheus
     additionalLabels: {}
+    targetLabels: []
 
 resources: {}
   # We usually recommend not to specify default resources and to leave this as a conscious

```

## 1.2.0

**Release date:** 2021-08-24

![AppVersion: v1.3.1](https://img.shields.io/static/v1?label=AppVersion&message=v1.3.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-kafka-exporter] Configure log level and extra args (#1275)

### Default value changes

```diff
diff --git a/charts/prometheus-kafka-exporter/values.yaml b/charts/prometheus-kafka-exporter/values.yaml
index 98156b10..894e072d 100644
--- a/charts/prometheus-kafka-exporter/values.yaml
+++ b/charts/prometheus-kafka-exporter/values.yaml
@@ -1,6 +1,6 @@
 image:
   repository: danielqsj/kafka-exporter
-  tag: latest
+  tag: v1.3.1
   pullPolicy: IfNotPresent
 
 replicaCount: 1
@@ -11,6 +11,13 @@ kafkaServer:
 # Specifies broker version to use, leave empty for default
 kafkaBrokerVersion:
 
+# Valid levels: [debug, info, warn, error, fatal]
+logLevel: info
+
+# Sarama logging
+sarama:
+  logEnabled: false
+
 rbac:
   # Specifies whether RBAC resources should be created
   create: true
@@ -26,6 +33,11 @@ serviceAccount:
 
 env: []
 
+# List of additional cli arguments to configure kafka-exporter
+# for example: --log.enable-sarama, --log.level=debug, etc.
+# all the possible args can be found here: https://github.com/danielqsj/kafka_exporter#flags
+extraArgs: []
+
 service:
   type: ClusterIP
   port: 9308
@@ -85,6 +97,7 @@ affinity: {}
 # tls.key
 tls:
   enabled: false
+  insecureSkipVerify: false
   # mountPath: /secret/certs
   # secretName: <name of an existing secret>
 

```

## 1.1.1

**Release date:** 2021-06-16

![AppVersion: v1.2.0](https://img.shields.io/static/v1?label=AppVersion&message=v1.2.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-kafka-exporter] Add missing key `annotations` in `values.yml` example (#1078)

### Default value changes

```diff
diff --git a/charts/prometheus-kafka-exporter/values.yaml b/charts/prometheus-kafka-exporter/values.yaml
index 4785de2e..98156b10 100644
--- a/charts/prometheus-kafka-exporter/values.yaml
+++ b/charts/prometheus-kafka-exporter/values.yaml
@@ -72,6 +72,8 @@ resources: {}
 
 nodeSelector: {}
 
+annotations: {}
+
 tolerations: []
 
 affinity: {}

```

## 1.1.0

**Release date:** 2021-04-28

![AppVersion: v1.2.0](https://img.shields.io/static/v1?label=AppVersion&message=v1.2.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-kafka-exporter] Add kafka broker version (#896)

### Default value changes

```diff
diff --git a/charts/prometheus-kafka-exporter/values.yaml b/charts/prometheus-kafka-exporter/values.yaml
index cb810890..4785de2e 100644
--- a/charts/prometheus-kafka-exporter/values.yaml
+++ b/charts/prometheus-kafka-exporter/values.yaml
@@ -8,6 +8,9 @@ replicaCount: 1
 kafkaServer:
   - kafka-server:9092
 
+# Specifies broker version to use, leave empty for default
+kafkaBrokerVersion:
+
 rbac:
   # Specifies whether RBAC resources should be created
   create: true

```

## 1.0.0

**Release date:** 2021-01-24

![AppVersion: v1.2.0](https://img.shields.io/static/v1?label=AppVersion&message=v1.2.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-kafka-exporter] Use helm test hooks (#498)

### Default value changes

```diff
diff --git a/charts/prometheus-kafka-exporter/values.yaml b/charts/prometheus-kafka-exporter/values.yaml
index f238a90b..cb810890 100644
--- a/charts/prometheus-kafka-exporter/values.yaml
+++ b/charts/prometheus-kafka-exporter/values.yaml
@@ -1,8 +1,19 @@
+image:
+  repository: danielqsj/kafka-exporter
+  tag: latest
+  pullPolicy: IfNotPresent
+
+replicaCount: 1
+
+kafkaServer:
+  - kafka-server:9092
+
 rbac:
   # Specifies whether RBAC resources should be created
   create: true
   # Specifies whether a PodSecurityPolicy should be created
   pspEnabled: true
+
 serviceAccount:
   # Specifies whether a ServiceAccount should be created
   create: true
@@ -10,15 +21,6 @@ serviceAccount:
   # If not set and create is true, a name is generated using the fullname template
   name:
 
-replicaCount: 1
-image:
-  repository: danielqsj/kafka-exporter
-  tag: latest
-  pullPolicy: IfNotPresent
-
-kafkaServer:
-  - kafka-server:9092
-
 env: []
 
 service:
@@ -32,6 +34,7 @@ liveness:
     httpGet:
       path: /
       port: exporter-port
+
 readiness:
   enabled: false
   probe:
@@ -79,3 +82,8 @@ tls:
   enabled: false
   # mountPath: /secret/certs
   # secretName: <name of an existing secret>
+
+# If enabled Kafka dependency chart will be used.
+# This is only needed for the CI job so it should always be disabled.
+kafka:
+  enabled: false

```

## 0.2.1

**Release date:** 2021-01-20

![AppVersion: v1.2.0](https://img.shields.io/static/v1?label=AppVersion&message=v1.2.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Add environment variables for prometheus-kafka-exporter (#601)

### Default value changes

```diff
diff --git a/charts/prometheus-kafka-exporter/values.yaml b/charts/prometheus-kafka-exporter/values.yaml
index d15a7353..f238a90b 100644
--- a/charts/prometheus-kafka-exporter/values.yaml
+++ b/charts/prometheus-kafka-exporter/values.yaml
@@ -19,6 +19,8 @@ image:
 kafkaServer:
   - kafka-server:9092
 
+env: []
+
 service:
   type: ClusterIP
   port: 9308

```

## 0.2.0

**Release date:** 2020-11-18

![AppVersion: v1.2.0](https://img.shields.io/static/v1?label=AppVersion&message=v1.2.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* feat(prometheus-kafka-exporter): add configurable apiVersion (#385)

### Default value changes

```diff
# No changes in this release
```

## 0.1.5

**Release date:** 2020-11-18

![AppVersion: v1.2.0](https://img.shields.io/static/v1?label=AppVersion&message=v1.2.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-kafka-exporter] Servicemonitor is not able to find endpoints for kafka-exporter target (shows 0/0 up) #343 (#374)

### Default value changes

```diff
diff --git a/charts/prometheus-kafka-exporter/values.yaml b/charts/prometheus-kafka-exporter/values.yaml
index b4a72e05..d15a7353 100644
--- a/charts/prometheus-kafka-exporter/values.yaml
+++ b/charts/prometheus-kafka-exporter/values.yaml
@@ -42,6 +42,12 @@ prometheus:
     enabled: false
     namespace: monitoring
     interval: "30s"
+    # If serviceMonitor is enabled and you want prometheus to automatically register
+    # target using serviceMonitor, add additionalLabels with prometheus release name
+    # e.g. If you have deployed kube-prometheus-stack with release name kube-prometheus
+    # then additionalLabels will be
+    # additionalLabels:
+    #   release: kube-prometheus
     additionalLabels: {}
 
 resources: {}

```

## 0.1.4

**Release date:** 2020-10-18

![AppVersion: v1.2.0](https://img.shields.io/static/v1?label=AppVersion&message=v1.2.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-kafka-exporter] new chart donation (#216)

### Default value changes

```diff
rbac:
  # Specifies whether RBAC resources should be created
  create: true
  # Specifies whether a PodSecurityPolicy should be created
  pspEnabled: true
serviceAccount:
  # Specifies whether a ServiceAccount should be created
  create: true
  # The name of the ServiceAccount to use.
  # If not set and create is true, a name is generated using the fullname template
  name:

replicaCount: 1
image:
  repository: danielqsj/kafka-exporter
  tag: latest
  pullPolicy: IfNotPresent

kafkaServer:
  - kafka-server:9092

service:
  type: ClusterIP
  port: 9308
  annotations: {}

liveness:
  enabled: false
  probe:
    httpGet:
      path: /
      port: exporter-port
readiness:
  enabled: false
  probe:
    httpGet:
      path: /
      port: exporter-port

prometheus:
  serviceMonitor:
    enabled: false
    namespace: monitoring
    interval: "30s"
    additionalLabels: {}

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

# this allows the usage of tls connection to your kafka cluster based on a secret in tls format
# mandatory keys:
# ca.crt
# tls.crt
# tls.key
tls:
  enabled: false
  # mountPath: /secret/certs
  # secretName: <name of an existing secret>

```

---
Autogenerated from Helm Chart and git history using [helm-changelog](https://github.com/mogensen/helm-changelog)
