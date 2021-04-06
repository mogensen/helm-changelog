# Change Log

## 1.0.0 

**Release date:** 2021-01-24

![AppVersion: v1.2.0](https://img.shields.io/static/v1?label=AppVersion&message=v1.2.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-kafka-exporter] Use helm test hooks (#498) 

### Default value changes

```diff
diff --git a/charts/prometheus-kafka-exporter/values.yaml b/charts/prometheus-kafka-exporter/values.yaml
index f238a90..cb81089 100644
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

![AppVersion: v1.2.0](https://img.shields.io/static/v1?label=AppVersion&message=v1.2.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Add environment variables for prometheus-kafka-exporter (#601) 

### Default value changes

```diff
diff --git a/charts/prometheus-kafka-exporter/values.yaml b/charts/prometheus-kafka-exporter/values.yaml
index d15a735..f238a90 100644
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

![AppVersion: v1.2.0](https://img.shields.io/static/v1?label=AppVersion&message=v1.2.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* feat(prometheus-kafka-exporter): add configurable apiVersion (#385) 

### Default value changes

```diff
# No changes in this release
```

## 0.1.5 

**Release date:** 2020-11-18

![AppVersion: v1.2.0](https://img.shields.io/static/v1?label=AppVersion&message=v1.2.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-kafka-exporter] Servicemonitor is not able to find endpoints for kafka-exporter target (shows 0/0 up) #343 (#374) 

### Default value changes

```diff
diff --git a/charts/prometheus-kafka-exporter/values.yaml b/charts/prometheus-kafka-exporter/values.yaml
index b4a72e0..d15a735 100644
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

![AppVersion: v1.2.0](https://img.shields.io/static/v1?label=AppVersion&message=v1.2.0&color=success&logo=)
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
