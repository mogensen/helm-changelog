# Change Log

## 0.14.1 

**Release date:** 2021-02-15

![AppVersion: 0.10.0](https://img.shields.io/static/v1?label=AppVersion&message=0.10.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-cloudwatch-exporter] Set namespace (#660) 

### Default value changes

```diff
# No changes in this release
```

## 0.14.0 

**Release date:** 2021-02-12

![AppVersion: 0.10.0](https://img.shields.io/static/v1?label=AppVersion&message=0.10.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-cloudwatch-exporter] Update to cloudwatch exporter v0.10.0 (#668) 

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index 6602d83..2522365 100644
--- a/charts/prometheus-cloudwatch-exporter/values.yaml
+++ b/charts/prometheus-cloudwatch-exporter/values.yaml
@@ -6,7 +6,7 @@ replicaCount: 1
 
 image:
   repository: prom/cloudwatch-exporter
-  tag: cloudwatch_exporter-0.8.0
+  tag: cloudwatch_exporter-0.10.0
   pullPolicy: IfNotPresent
   pullSecrets:
   # - name: "image-pull-secret"
```

## 0.13.0 

**Release date:** 2021-01-14

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-cloudwatch-exporter] - Added option to enable sts regional endpoints (#565) 

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index 229026f..6602d83 100644
--- a/charts/prometheus-cloudwatch-exporter/values.yaml
+++ b/charts/prometheus-cloudwatch-exporter/values.yaml
@@ -51,6 +51,9 @@ resources: {}
 
 aws:
   role:
+  # Enables usage of regional STS endpoints rather than global which is default
+  stsRegional:
+    enabled: false
 
   # The name of a pre-created secret in which AWS credentials are stored. When
   # set, aws_access_key_id is assumed to be in a field called access_key,
```

## 0.12.1 

**Release date:** 2020-12-20

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-cloudwatch-exporter] Add the fsGroup:65534 to the security context (#502) 

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index ed13fa5..229026f 100644
--- a/charts/prometheus-cloudwatch-exporter/values.yaml
+++ b/charts/prometheus-cloudwatch-exporter/values.yaml
@@ -208,6 +208,7 @@ ingress:
 
 securityContext:
   runAsUser: 65534  # run as nobody user instead of root
+  fsGroup: 65534  # necessary to be able to read the EKS IAM token
 
 # Leverage a PriorityClass to ensure your pods survive resource shortages
 # ref: https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption/
```

## 0.12.0 

**Release date:** 2020-11-24

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-cloudwatch-exporter] Allow user-defined matchLabels for deployment and servicemonitor in prometheus-cloudwatch-exporter (#406) 

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index 3e77bb8..ed13fa5 100644
--- a/charts/prometheus-cloudwatch-exporter/values.yaml
+++ b/charts/prometheus-cloudwatch-exporter/values.yaml
@@ -34,6 +34,9 @@ service:
   annotations: {}
   labels: {}
 
+pod:
+  labels: {}
+
 resources: {}
   # We usually recommend not to specify default resources and to leave this as a conscious
   # choice for the user. This also increases chances charts run on environments with little
```

## 0.11.0 

**Release date:** 2020-11-21

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Add imagePullSecrets to prometheus-cloudwatch-exporter (#391) 

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index fa9fade..3e77bb8 100644
--- a/charts/prometheus-cloudwatch-exporter/values.yaml
+++ b/charts/prometheus-cloudwatch-exporter/values.yaml
@@ -8,6 +8,8 @@ image:
   repository: prom/cloudwatch-exporter
   tag: cloudwatch_exporter-0.8.0
   pullPolicy: IfNotPresent
+  pullSecrets:
+  # - name: "image-pull-secret"
 
 # Example proxy configuration:
 # command:
```

## 0.10.1 

**Release date:** 2020-11-06

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* charts/prometheus-cloudwatch-exporter: fix typo (#315) 

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index 4345295..fa9fade 100644
--- a/charts/prometheus-cloudwatch-exporter/values.yaml
+++ b/charts/prometheus-cloudwatch-exporter/values.yaml
@@ -67,7 +67,7 @@ serviceAccount:
   # If not set and create is true, a name is generated using the fullname template
   name:
   # annotations:
-  # Will add the provided map to the annotations for the crated serviceAccount
+  # Will add the provided map to the annotations for the created serviceAccount
   # e.g.
   # annotations:
   #   eks.amazonaws.com/role-arn: arn:aws:iam::1234567890:role/prom-cloudwatch-exporter-oidc
```

## 0.10.0 

**Release date:** 2020-10-21

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-cloudwatch-exporter] Allow Helm variables in config (#220) 

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index 2b77b68..4345295 100644
--- a/charts/prometheus-cloudwatch-exporter/values.yaml
+++ b/charts/prometheus-cloudwatch-exporter/values.yaml
@@ -76,6 +76,7 @@ rbac:
   # Specifies whether RBAC resources should be created
   create: true
 
+# Configuration is rendered with `tpl` function, therefore you can use any Helm variables and/or templates here
 config: |-
   # This is the default configuration for prometheus-cloudwatch-exporter
   region: eu-west-1
```

## 0.9.0 

**Release date:** 2020-09-10

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-cloudwatch-exporter] Add priorityClassName to cloudwatch-exporter (#31) 

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index 2b10317..2b77b68 100644
--- a/charts/prometheus-cloudwatch-exporter/values.yaml
+++ b/charts/prometheus-cloudwatch-exporter/values.yaml
@@ -202,3 +202,8 @@ ingress:
 
 securityContext:
   runAsUser: 65534  # run as nobody user instead of root
+
+# Leverage a PriorityClass to ensure your pods survive resource shortages
+# ref: https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption/
+# priorityClassName: system-cluster-critical
+priorityClassName: ""
```

## 0.8.4 

**Release date:** 2020-08-20

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Prep initial charts indexing (#14) 

### Default value changes

```diff
# No changes in this release
```

## 0.8.3 

**Release date:** 2020-07-27

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* prometheus-cloudwatch-exporter: Add containerPort, livenessProbe.path and readinessProbe.path variables (#22923) 

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index c4da6cb..2b10317 100644
--- a/charts/prometheus-cloudwatch-exporter/values.yaml
+++ b/charts/prometheus-cloudwatch-exporter/values.yaml
@@ -23,6 +23,8 @@ image:
 
 command: []
 
+containerPort: 9106
+
 service:
   type: ClusterIP
   port: 9106
@@ -113,6 +115,7 @@ affinity: {}
 
 # Configurable health checks against the /healthy and /ready endpoints
 livenessProbe:
+  path: /-/healthy
   initialDelaySeconds: 30
   periodSeconds: 5
   timeoutSeconds: 5
@@ -120,6 +123,7 @@ livenessProbe:
   failureThreshold: 3
 
 readinessProbe:
+  path: /-/ready
   initialDelaySeconds: 30
   periodSeconds: 5
   timeoutSeconds: 5
```

## 0.8.2 

**Release date:** 2020-07-01

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-cloudwatch-exporter] additional docu on service account usage (#23028) 

### Default value changes

```diff
# No changes in this release
```

## 0.8.1 

**Release date:** 2020-06-23

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* stable/prometheus-cloudwatch-exporter Add shasum of secrets yaml render to annotations (#22906) 

### Default value changes

```diff
# No changes in this release
```

## 0.8.0 

**Release date:** 2020-05-22

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-cloudwatch-exporter] Enable to set an existing service account name (#22272) 

### Default value changes

```diff
# No changes in this release
```

## 0.7.0 

**Release date:** 2020-04-09

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* stable/prometheus-cloudwatch-exporter Add prometheusRule creation. (#21838) 

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index 2ddb9e2..c4da6cb 100644
--- a/charts/prometheus-cloudwatch-exporter/values.yaml
+++ b/charts/prometheus-cloudwatch-exporter/values.yaml
@@ -6,7 +6,7 @@ replicaCount: 1
 
 image:
   repository: prom/cloudwatch-exporter
-  tag: cloudwatch_exporter-0.6.0
+  tag: cloudwatch_exporter-0.8.0
   pullPolicy: IfNotPresent
 
 # Example proxy configuration:
@@ -151,6 +151,37 @@ serviceMonitor:
   #     replacement: mydbname
   #     targetLabel: dbname
 
+prometheusRule:
+  # Specifies whether a PrometheusRule should be created
+  enabled: false
+  # Set the namespace the PrometheusRule should be deployed
+  # namespace: monitoring
+  # Set labels for the PrometheusRule, use this to define your scrape label for Prometheus Operator
+  # labels:
+  # Example - note the Kubernetes convention of camelCase instead of Prometheus'
+  # rules:
+  #    - alert: ELB-Low-BurstBalance
+  #      annotations:
+  #        message: The ELB BurstBalance during the last 10 minutes is lower than 80%.
+  #      expr: aws_ebs_burst_balance_average < 80
+  #      for: 10m
+  #      labels:
+  #        severity: warning
+  #    - alert: ELB-Low-BurstBalance
+  #      annotations:
+  #        message: The ELB BurstBalance during the last 10 minutes is lower than 50%.
+  #      expr: aws_ebs_burst_balance_average < 50
+  #      for: 10m
+  #      labels:
+  #        severity: warning
+  #    - alert: ELB-Low-BurstBalance
+  #      annotations:
+  #        message: The ELB BurstBalance during the last 10 minutes is lower than 30%.
+  #      expr: aws_ebs_burst_balance_average < 30
+  #      for: 10m
+  #      labels:
+  #        severity: critical
+
 ingress:
   enabled: false
   annotations: {}
```

## 0.6.0 

**Release date:** 2020-01-22

![AppVersion: 0.6.0](https://img.shields.io/static/v1?label=AppVersion&message=0.6.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-cloudwatch-exporter] - Add annotations to serviceAccounts for EKS IAM (#20162) 

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index f7c9f83..2ddb9e2 100644
--- a/charts/prometheus-cloudwatch-exporter/values.yaml
+++ b/charts/prometheus-cloudwatch-exporter/values.yaml
@@ -64,6 +64,11 @@ serviceAccount:
   # The name of the ServiceAccount to use.
   # If not set and create is true, a name is generated using the fullname template
   name:
+  # annotations:
+  # Will add the provided map to the annotations for the crated serviceAccount
+  # e.g.
+  # annotations:
+  #   eks.amazonaws.com/role-arn: arn:aws:iam::1234567890:role/prom-cloudwatch-exporter-oidc
 
 rbac:
   # Specifies whether RBAC resources should be created
```

## 0.5.0 

**Release date:** 2019-11-02

![AppVersion: 0.6.0](https://img.shields.io/static/v1?label=AppVersion&message=0.6.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-cloudwatch-exporter] Upgrade api versions (#17764) 

### Default value changes

```diff
# No changes in this release
```

## 0.4.13 

**Release date:** 2019-10-28

![AppVersion: 0.6.0](https://img.shields.io/static/v1?label=AppVersion&message=0.6.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-cloudwatch-exporter] Update cloudwatch-exporter version. (#18003) 

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index 279be11..f7c9f83 100644
--- a/charts/prometheus-cloudwatch-exporter/values.yaml
+++ b/charts/prometheus-cloudwatch-exporter/values.yaml
@@ -6,7 +6,7 @@ replicaCount: 1
 
 image:
   repository: prom/cloudwatch-exporter
-  tag: cloudwatch_exporter-0.5.0
+  tag: cloudwatch_exporter-0.6.0
   pullPolicy: IfNotPresent
 
 # Example proxy configuration:
```

## 0.4.12 

**Release date:** 2019-10-28

![AppVersion: 0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=0.5.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-cloudwatch-exporter] Add myself to owners/maintainers. (#18093) 

### Default value changes

```diff
# No changes in this release
```

## 0.4.11 

**Release date:** 2019-10-03

![AppVersion: 0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=0.5.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-cloudwatch-exporter] added changelog (#16579) 

### Default value changes

```diff
# No changes in this release
```

## 0.4.10 

**Release date:** 2019-08-21

![AppVersion: 0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=0.5.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* include ServiceMonitor relabelings and metricRelabelings (#16443) 

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index ecf6e51..279be11 100644
--- a/charts/prometheus-cloudwatch-exporter/values.yaml
+++ b/charts/prometheus-cloudwatch-exporter/values.yaml
@@ -134,6 +134,17 @@ serviceMonitor:
   # labels:
   # Set timeout for scrape
   # timeout: 10s
+  # Set relabelings for the ServiceMonitor, use to apply to samples before scraping
+  # relabelings: []
+  # Set metricRelabelings for the ServiceMonitor, use to apply to samples for ingestion
+  # metricRelabelings: []
+  #
+  # Example - note the Kubernetes convention of camelCase instead of Prometheus' snake_case
+  # metricRelabelings:
+  #   - sourceLabels: [dbinstance_identifier]
+  #     action: replace
+  #     replacement: mydbname
+  #     targetLabel: dbname
 
 ingress:
   enabled: false
```

## 0.4.9 

**Release date:** 2019-08-04

![AppVersion: 0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=0.5.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-cloudwatch-exporter] Fix documentation typos (README.md) (#15745) 

### Default value changes

```diff
# No changes in this release
```

## 0.4.8 

**Release date:** 2019-07-19

![AppVersion: 0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=0.5.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-cloudwatch-exporter] Added securityContext configuration capabilities (#12405) 

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index 2d3094c..ecf6e51 100644
--- a/charts/prometheus-cloudwatch-exporter/values.yaml
+++ b/charts/prometheus-cloudwatch-exporter/values.yaml
@@ -148,3 +148,6 @@ ingress:
   #  - secretName: chart-example-tls
   #    hosts:
   #      - chart-example.local
+
+securityContext:
+  runAsUser: 65534  # run as nobody user instead of root
```

## 0.4.7 

**Release date:** 2019-06-14

![AppVersion: 0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=0.5.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-cloudwatch-exporter] configurable container command (#14803) 

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index 70faf35..2d3094c 100644
--- a/charts/prometheus-cloudwatch-exporter/values.yaml
+++ b/charts/prometheus-cloudwatch-exporter/values.yaml
@@ -9,6 +9,20 @@ image:
   tag: cloudwatch_exporter-0.5.0
   pullPolicy: IfNotPresent
 
+# Example proxy configuration:
+# command:
+#   - 'java'
+#   - '-Dhttp.proxyHost=proxy.example.com'
+#   - '-Dhttp.proxyPort=3128'
+#   - '-Dhttps.proxyHost=proxy.example.com'
+#   - '-Dhttps.proxyPort=3128'
+#   - '-jar'
+#   - '/cloudwatch_exporter.jar'
+#   - '9106'
+#   - '/config/config.yml'
+
+command: []
+
 service:
   type: ClusterIP
   port: 9106
```

## 0.4.6 

**Release date:** 2019-06-13

![AppVersion: 0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=0.5.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* add torstenwalter as maintainer (#14788) 

### Default value changes

```diff
# No changes in this release
```

## 0.4.5 

**Release date:** 2019-05-13

![AppVersion: 0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=0.5.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-cloudwatch-exporter] Add scrape timeout option (#13686) 

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index 405bfcc..70faf35 100644
--- a/charts/prometheus-cloudwatch-exporter/values.yaml
+++ b/charts/prometheus-cloudwatch-exporter/values.yaml
@@ -118,6 +118,8 @@ serviceMonitor:
   # telemetryPath: /metrics
   # Set labels for the ServiceMonitor, use this to define your scrape label for Prometheus Operator
   # labels:
+  # Set timeout for scrape
+  # timeout: 10s
 
 ingress:
   enabled: false
```

## 0.4.4 

**Release date:** 2019-04-02

![AppVersion: 0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=0.5.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-cloudwatch-exporter] Added ingress configuration (#12407) 

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index eeccd3f..405bfcc 100644
--- a/charts/prometheus-cloudwatch-exporter/values.yaml
+++ b/charts/prometheus-cloudwatch-exporter/values.yaml
@@ -118,3 +118,17 @@ serviceMonitor:
   # telemetryPath: /metrics
   # Set labels for the ServiceMonitor, use this to define your scrape label for Prometheus Operator
   # labels:
+
+ingress:
+  enabled: false
+  annotations: {}
+    # kubernetes.io/ingress.class: nginx
+    # kubernetes.io/tls-acme: "true"
+  labels: {}
+  path: /
+  hosts:
+    - chart-example.local
+  tls: []
+  #  - secretName: chart-example-tls
+  #    hosts:
+  #      - chart-example.local
```

## 0.4.3 

**Release date:** 2019-04-01

![AppVersion: 0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=0.5.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Add optional ServiceMonitor and values stub with it disabled by default (#12503) 

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index 0f1c8be..eeccd3f 100644
--- a/charts/prometheus-cloudwatch-exporter/values.yaml
+++ b/charts/prometheus-cloudwatch-exporter/values.yaml
@@ -106,3 +106,15 @@ readinessProbe:
   timeoutSeconds: 5
   successThreshold: 1
   failureThreshold: 3
+
+serviceMonitor:
+  # When set true then use a ServiceMonitor to configure scraping
+  enabled: false
+  # Set the namespace the ServiceMonitor should be deployed
+  # namespace: monitoring
+  # Set how frequently Prometheus should scrape
+  # interval: 30s
+  # Set path to cloudwatch-exporter telemtery-path
+  # telemetryPath: /metrics
+  # Set labels for the ServiceMonitor, use this to define your scrape label for Prometheus Operator
+  # labels:
```

## 0.4.2 

**Release date:** 2019-02-20

![AppVersion: 0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=0.5.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Fix default service port and target port.  (#11264) 

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index 6117ea9..0f1c8be 100644
--- a/charts/prometheus-cloudwatch-exporter/values.yaml
+++ b/charts/prometheus-cloudwatch-exporter/values.yaml
@@ -11,7 +11,7 @@ image:
 
 service:
   type: ClusterIP
-  port: 80
+  port: 9106
   portName: http
   annotations: {}
   labels: {}
```

## 0.4.1 

**Release date:** 2019-02-18

![AppVersion: 0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=0.5.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-cloudwatch-exporter] Adjust service targetPort to reference container port name. (#11462) 

### Default value changes

```diff
# No changes in this release
```

## 0.4.0 

**Release date:** 2019-01-29

![AppVersion: 0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=0.5.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Add support for pre-created secrets and AWS STS session tokens (#10878) 

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index e26ef68..6117ea9 100644
--- a/charts/prometheus-cloudwatch-exporter/values.yaml
+++ b/charts/prometheus-cloudwatch-exporter/values.yaml
@@ -29,9 +29,18 @@ resources: {}
   #   memory: 128Mi
 
 aws:
-  region: eu-west-1
   role:
-  # Note: Do not specify the aws_access_key_id abd aws_secret_access_key if you specified role before
+
+  # The name of a pre-created secret in which AWS credentials are stored. When
+  # set, aws_access_key_id is assumed to be in a field called access_key,
+  # aws_secret_access_key is assumed to be in a field called secret_key, and the
+  # session token, if it exists, is assumed to be in a field called
+  # security_token
+  secret:
+    name:
+    includesSessionToken: false
+
+  # Note: Do not specify the aws_access_key_id and aws_secret_access_key if you specified role or secret.name before
   aws_access_key_id:
   aws_secret_access_key:
 
```

## 0.3.0 

**Release date:** 2019-01-24

![AppVersion: 0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=0.5.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Use entrypoint from Docker image and fix probe paths (#10873) 

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index c089a7d..e26ef68 100644
--- a/charts/prometheus-cloudwatch-exporter/values.yaml
+++ b/charts/prometheus-cloudwatch-exporter/values.yaml
@@ -12,7 +12,6 @@ image:
 service:
   type: ClusterIP
   port: 80
-  targetPort: 9100
   portName: http
   annotations: {}
   labels: {}
```

## 0.2.1 

**Release date:** 2018-10-16

![AppVersion: 0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=0.5.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* cloudwatch-exporter - Support custom service labels and portName (#8488) 

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index bb986cf..c089a7d 100644
--- a/charts/prometheus-cloudwatch-exporter/values.yaml
+++ b/charts/prometheus-cloudwatch-exporter/values.yaml
@@ -13,7 +13,9 @@ service:
   type: ClusterIP
   port: 80
   targetPort: 9100
+  portName: http
   annotations: {}
+  labels: {}
 
 resources: {}
   # We usually recommend not to specify default resources and to leave this as a conscious
```

## 0.2.0 

**Release date:** 2018-09-13

![AppVersion: 0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=0.5.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-cloudwatch-exporter] health/liveness endpoints added (#7681) 

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index 68dda8c..bb986cf 100644
--- a/charts/prometheus-cloudwatch-exporter/values.yaml
+++ b/charts/prometheus-cloudwatch-exporter/values.yaml
@@ -6,7 +6,7 @@ replicaCount: 1
 
 image:
   repository: prom/cloudwatch-exporter
-  tag: latest
+  tag: cloudwatch_exporter-0.5.0
   pullPolicy: IfNotPresent
 
 service:
@@ -81,3 +81,18 @@ nodeSelector: {}
 tolerations: []
 
 affinity: {}
+
+# Configurable health checks against the /healthy and /ready endpoints
+livenessProbe:
+  initialDelaySeconds: 30
+  periodSeconds: 5
+  timeoutSeconds: 5
+  successThreshold: 1
+  failureThreshold: 3
+
+readinessProbe:
+  initialDelaySeconds: 30
+  periodSeconds: 5
+  timeoutSeconds: 5
+  successThreshold: 1
+  failureThreshold: 3
```

## 0.1.4 

**Release date:** 2018-07-10

![AppVersion: 0.1.0](https://img.shields.io/static/v1?label=AppVersion&message=0.1.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-cloudwatch-exporter] Fixed secret variable parsing for access/secret keys. (#6511) 

### Default value changes

```diff
# No changes in this release
```

## 0.1.3 

**Release date:** 2018-06-29

![AppVersion: 0.1.0](https://img.shields.io/static/v1?label=AppVersion&message=0.1.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Fixed image configuration in README for the prometheus-cloudwatch-exporter (#6400) 

### Default value changes

```diff
# No changes in this release
```

## 0.1.2 

**Release date:** 2018-04-16

![AppVersion: 0.1.0](https://img.shields.io/static/v1?label=AppVersion&message=0.1.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Annotations and checksum (#5067) 

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index 2c4a90c..68dda8c 100644
--- a/charts/prometheus-cloudwatch-exporter/values.yaml
+++ b/charts/prometheus-cloudwatch-exporter/values.yaml
@@ -13,6 +13,7 @@ service:
   type: ClusterIP
   port: 80
   targetPort: 9100
+  annotations: {}
 
 resources: {}
   # We usually recommend not to specify default resources and to leave this as a conscious
```

## 0.1.1 

**Release date:** 2018-04-03

![AppVersion: 0.1.0](https://img.shields.io/static/v1?label=AppVersion&message=0.1.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Update table/prometheus-cloudwatch-exporte (#4551) 

### Default value changes

```diff
# No changes in this release
```

## 0.1.0 

**Release date:** 2018-03-30

![AppVersion: 0.1.0](https://img.shields.io/static/v1?label=AppVersion&message=0.1.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Add cloudwatch exporter (#4022) 

### Default value changes

```diff
# Default values for prometheus-cloudwatch-exporter.
# This is a YAML-formatted file.
# Declare variables to be passed into your templates.

replicaCount: 1

image:
  repository: prom/cloudwatch-exporter
  tag: latest
  pullPolicy: IfNotPresent

service:
  type: ClusterIP
  port: 80
  targetPort: 9100

resources: {}
  # We usually recommend not to specify default resources and to leave this as a conscious
  # choice for the user. This also increases chances charts run on environments with little
  # resources, such as Minikube. If you do want to specify resources, uncomment the following
  # lines, adjust them as necessary, and remove the curly braces after 'resources:'.
  # limits:
  #   cpu: 100m
  #    memory: 128Mi
  # requests:
  #   cpu: 100m
  #   memory: 128Mi

aws:
  region: eu-west-1
  role:
  # Note: Do not specify the aws_access_key_id abd aws_secret_access_key if you specified role before
  aws_access_key_id:
  aws_secret_access_key:

serviceAccount:
  # Specifies whether a ServiceAccount should be created
  create: true
  # The name of the ServiceAccount to use.
  # If not set and create is true, a name is generated using the fullname template
  name:

rbac:
  # Specifies whether RBAC resources should be created
  create: true

config: |-
  # This is the default configuration for prometheus-cloudwatch-exporter
  region: eu-west-1
  period_seconds: 240
  metrics:
  - aws_namespace: AWS/ELB
    aws_metric_name: HealthyHostCount
    aws_dimensions: [AvailabilityZone, LoadBalancerName]
    aws_statistics: [Average]

  - aws_namespace: AWS/ELB
    aws_metric_name: UnHealthyHostCount
    aws_dimensions: [AvailabilityZone, LoadBalancerName]
    aws_statistics: [Average]

  - aws_namespace: AWS/ELB
    aws_metric_name: RequestCount
    aws_dimensions: [AvailabilityZone, LoadBalancerName]
    aws_statistics: [Sum]

  - aws_namespace: AWS/ELB
    aws_metric_name: Latency
    aws_dimensions: [AvailabilityZone, LoadBalancerName]
    aws_statistics: [Average]

  - aws_namespace: AWS/ELB
    aws_metric_name: SurgeQueueLength
    aws_dimensions: [AvailabilityZone, LoadBalancerName]
    aws_statistics: [Maximum, Sum]


nodeSelector: {}

tolerations: []

affinity: {}
```

---
Autogenerated from Helm Chart and git history using [helm-changelog](https://github.com/mogensen/helm-changelog)
