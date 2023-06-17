# Change Log

## 0.25.1

**Release date:** 2023-06-16

![AppVersion: 0.15.4](https://img.shields.io/static/v1?label=AppVersion&message=0.15.4&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* update image version (#3500)

### Default value changes

```diff
# No changes in this release
```

## 0.25.0

**Release date:** 2023-05-01

![AppVersion: 0.15.3](https://img.shields.io/static/v1?label=AppVersion&message=0.15.3&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-cloudwatch-exporter] Appversion bump to 0.15.3 (#3241)

### Default value changes

```diff
# No changes in this release
```

## 0.24.0

**Release date:** 2023-02-10

![AppVersion: 0.15.1](https://img.shields.io/static/v1?label=AppVersion&message=0.15.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-cloudwatch-exporter] Add labels to deployment resource (#3001)

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index c5923ac1..42a63fe3 100644
--- a/charts/prometheus-cloudwatch-exporter/values.yaml
+++ b/charts/prometheus-cloudwatch-exporter/values.yaml
@@ -39,6 +39,11 @@ pod:
   labels: {}
   annotations: {}
 
+# Labels and annotations to attach to the deployment resource
+deployment:
+  labels: {}
+  annotations: {}
+
 # Extra environment variables
 extraEnv:
   # - name: foo

```

## 0.23.0

**Release date:** 2023-02-08

![AppVersion: 0.15.1](https://img.shields.io/static/v1?label=AppVersion&message=0.15.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-cloudwatch-exporter]: upgrade to v0.15.1 (#3008)

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index 1cb5c20a..c5923ac1 100644
--- a/charts/prometheus-cloudwatch-exporter/values.yaml
+++ b/charts/prometheus-cloudwatch-exporter/values.yaml
@@ -6,7 +6,8 @@ replicaCount: 1
 
 image:
   repository: prom/cloudwatch-exporter
-  tag: v0.15.0
+  # if not set appVersion field from Chart.yaml is used
+  tag:
   pullPolicy: IfNotPresent
   pullSecrets:
   # - name: "image-pull-secret"

```

## 0.22.0

**Release date:** 2022-10-14

![AppVersion: 0.15.0](https://img.shields.io/static/v1?label=AppVersion&message=0.15.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Support networking.k8s.io/v1 in helpers file for k8s clusters > 1.21 (#2560)

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index 37715907..1cb5c20a 100644
--- a/charts/prometheus-cloudwatch-exporter/values.yaml
+++ b/charts/prometheus-cloudwatch-exporter/values.yaml
@@ -218,6 +218,13 @@ ingress:
   #    hosts:
   #      - chart-example.local
 
+  # For Kubernetes >= 1.18 you should specify the ingress-controller via the field ingressClassName
+  # See https://kubernetes.io/blog/2020/04/02/improvements-to-the-ingress-api-in-kubernetes-1.18/#specifying-the-class-of-an-ingress
+  # ingressClassName: nginx
+
+  # pathType is only for k8s >= 1.18
+  pathType: Prefix
+
 securityContext:
   runAsUser: 65534  # run as nobody user instead of root
   fsGroup: 65534  # necessary to be able to read the EKS IAM token

```

## 0.21.1

**Release date:** 2022-10-13

![AppVersion: 0.15.0](https://img.shields.io/static/v1?label=AppVersion&message=0.15.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Support networking.k8s.io/v1 in helpers file for k8s clusters > 1.21 (#2558)

### Default value changes

```diff
# No changes in this release
```

## 0.21.0

**Release date:** 2022-09-29

![AppVersion: 0.15.0](https://img.shields.io/static/v1?label=AppVersion&message=0.15.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-cloudwatch-exporter] Add extraEnv to values file to specify extra env vars being added to the deployment (#2500)

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index fd470670..37715907 100644
--- a/charts/prometheus-cloudwatch-exporter/values.yaml
+++ b/charts/prometheus-cloudwatch-exporter/values.yaml
@@ -38,6 +38,11 @@ pod:
   labels: {}
   annotations: {}
 
+# Extra environment variables
+extraEnv:
+  # - name: foo
+  #   value: baa
+
 resources: {}
   # We usually recommend not to specify default resources and to leave this as a conscious
   # choice for the user. This also increases chances charts run on environments with little

```

## 0.20.1

**Release date:** 2022-09-28

![AppVersion: 0.15.0](https://img.shields.io/static/v1?label=AppVersion&message=0.15.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-cloudwatch-exporter] Fix typo in README (#2392)

### Default value changes

```diff
# No changes in this release
```

## 0.20.0

**Release date:** 2022-09-15

![AppVersion: 0.15.0](https://img.shields.io/static/v1?label=AppVersion&message=0.15.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-cloudwatch-exporter] Update image tag v0.14.3 to v0.15.0 (#2458)

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index 45f73344..fd470670 100644
--- a/charts/prometheus-cloudwatch-exporter/values.yaml
+++ b/charts/prometheus-cloudwatch-exporter/values.yaml
@@ -6,7 +6,7 @@ replicaCount: 1
 
 image:
   repository: prom/cloudwatch-exporter
-  tag: v0.14.3
+  tag: v0.15.0
   pullPolicy: IfNotPresent
   pullSecrets:
   # - name: "image-pull-secret"

```

## 0.19.2

**Release date:** 2022-07-08

![AppVersion: 0.14.3](https://img.shields.io/static/v1?label=AppVersion&message=0.14.3&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-cloudwatch-exporter] correct container service port  Signed-… (#2235)

### Default value changes

```diff
# No changes in this release
```

## 0.19.1

**Release date:** 2022-06-27

![AppVersion: 0.14.3](https://img.shields.io/static/v1?label=AppVersion&message=0.14.3&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-cloudwatch-exporter] update the cloudwatch-exporter image from 0.14.0 to 0.14.3 (#2171)

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index 7789e917..45f73344 100644
--- a/charts/prometheus-cloudwatch-exporter/values.yaml
+++ b/charts/prometheus-cloudwatch-exporter/values.yaml
@@ -6,7 +6,7 @@ replicaCount: 1
 
 image:
   repository: prom/cloudwatch-exporter
-  tag: v0.14.0
+  tag: v0.14.3
   pullPolicy: IfNotPresent
   pullSecrets:
   # - name: "image-pull-secret"

```

## 0.19.0

**Release date:** 2022-05-30

![AppVersion: 0.14.0](https://img.shields.io/static/v1?label=AppVersion&message=0.14.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-cloudwatch-exporter] Update default image tag for Cloudwatch exporter chart (#2097)

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index 9d7f1eda..7789e917 100644
--- a/charts/prometheus-cloudwatch-exporter/values.yaml
+++ b/charts/prometheus-cloudwatch-exporter/values.yaml
@@ -6,7 +6,7 @@ replicaCount: 1
 
 image:
   repository: prom/cloudwatch-exporter
-  tag: v0.12.2
+  tag: v0.14.0
   pullPolicy: IfNotPresent
   pullSecrets:
   # - name: "image-pull-secret"

```

## 0.18.0

**Release date:** 2022-01-14

![AppVersion: 0.12.2](https://img.shields.io/static/v1?label=AppVersion&message=0.12.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* fix(cloudwatch-exporter): update the image from 0.10.0 to 0.12.2 (#1597)

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index 33ff85da..9d7f1eda 100644
--- a/charts/prometheus-cloudwatch-exporter/values.yaml
+++ b/charts/prometheus-cloudwatch-exporter/values.yaml
@@ -6,7 +6,7 @@ replicaCount: 1
 
 image:
   repository: prom/cloudwatch-exporter
-  tag: cloudwatch_exporter-0.10.0
+  tag: v0.12.2
   pullPolicy: IfNotPresent
   pullSecrets:
   # - name: "image-pull-secret"

```

## 0.17.2

**Release date:** 2022-01-04

![AppVersion: 0.10.0](https://img.shields.io/static/v1?label=AppVersion&message=0.10.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* fix: Indentation for multiple annotations of ServiceAccounts (#1667)

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index 83e13e9b..33ff85da 100644
--- a/charts/prometheus-cloudwatch-exporter/values.yaml
+++ b/charts/prometheus-cloudwatch-exporter/values.yaml
@@ -80,6 +80,7 @@ serviceAccount:
   # e.g.
   # annotations:
   #   eks.amazonaws.com/role-arn: arn:aws:iam::1234567890:role/prom-cloudwatch-exporter-oidc
+  #   eks.amazonaws.com/sts-regional-endpoints: "true"
   # Specifies whether to automount API credentials for the ServiceAccount.
   automountServiceAccountToken: true
 

```

## 0.17.1

**Release date:** 2021-11-24

![AppVersion: 0.10.0](https://img.shields.io/static/v1?label=AppVersion&message=0.10.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Add missing whitespace to cloudwatch-exporter chart (#1533)

### Default value changes

```diff
# No changes in this release
```

## 0.17.0

**Release date:** 2021-11-07

![AppVersion: 0.10.0](https://img.shields.io/static/v1?label=AppVersion&message=0.10.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-cloudwatch-exporter] configurable pod annotations (#1294)

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index 956b1b07..83e13e9b 100644
--- a/charts/prometheus-cloudwatch-exporter/values.yaml
+++ b/charts/prometheus-cloudwatch-exporter/values.yaml
@@ -36,6 +36,7 @@ service:
 
 pod:
   labels: {}
+  annotations: {}
 
 resources: {}
   # We usually recommend not to specify default resources and to leave this as a conscious

```

## 0.16.0

**Release date:** 2021-05-03

![AppVersion: 0.10.0](https://img.shields.io/static/v1?label=AppVersion&message=0.10.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-cloudwatch-exporter] Allow for specifying automountserviceaccounttoken for prometheus-cloudwatch-exporter (#861)

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index 6be9d2f4..956b1b07 100644
--- a/charts/prometheus-cloudwatch-exporter/values.yaml
+++ b/charts/prometheus-cloudwatch-exporter/values.yaml
@@ -79,6 +79,8 @@ serviceAccount:
   # e.g.
   # annotations:
   #   eks.amazonaws.com/role-arn: arn:aws:iam::1234567890:role/prom-cloudwatch-exporter-oidc
+  # Specifies whether to automount API credentials for the ServiceAccount.
+  automountServiceAccountToken: true
 
 rbac:
   # Specifies whether RBAC resources should be created

```

## 0.15.0

**Release date:** 2021-05-03

![AppVersion: 0.10.0](https://img.shields.io/static/v1?label=AppVersion&message=0.10.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Add container security context (#849)

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index 25223659..6be9d2f4 100644
--- a/charts/prometheus-cloudwatch-exporter/values.yaml
+++ b/charts/prometheus-cloudwatch-exporter/values.yaml
@@ -213,6 +213,10 @@ securityContext:
   runAsUser: 65534  # run as nobody user instead of root
   fsGroup: 65534  # necessary to be able to read the EKS IAM token
 
+containerSecurityContext: {}
+  # allowPrivilegeEscalation: false
+  # readOnlyRootFilesystem: true
+
 # Leverage a PriorityClass to ensure your pods survive resource shortages
 # ref: https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption/
 # priorityClassName: system-cluster-critical

```

## 0.14.1

**Release date:** 2021-02-16

![AppVersion: 0.10.0](https://img.shields.io/static/v1?label=AppVersion&message=0.10.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-cloudwatch-exporter] Set namespace (#660)

### Default value changes

```diff
# No changes in this release
```

## 0.14.0

**Release date:** 2021-02-12

![AppVersion: 0.10.0](https://img.shields.io/static/v1?label=AppVersion&message=0.10.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-cloudwatch-exporter] Update to cloudwatch exporter v0.10.0 (#668)

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index 6602d830..25223659 100644
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

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-cloudwatch-exporter] - Added option to enable sts regional endpoints (#565)

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index 229026f6..6602d830 100644
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

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-cloudwatch-exporter] Add the fsGroup:65534 to the security context (#502)

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index ed13fa56..229026f6 100644
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

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-cloudwatch-exporter] Allow user-defined matchLabels for deployment and servicemonitor in prometheus-cloudwatch-exporter (#406)

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index 3e77bb81..ed13fa56 100644
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

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Add imagePullSecrets to prometheus-cloudwatch-exporter (#391)

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index fa9fade9..3e77bb81 100644
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

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* charts/prometheus-cloudwatch-exporter: fix typo (#315)

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index 43452954..fa9fade9 100644
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

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-cloudwatch-exporter] Allow Helm variables in config (#220)

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index 2b77b682..43452954 100644
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

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-cloudwatch-exporter] Add priorityClassName to cloudwatch-exporter (#31)

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index 2b10317e..2b77b682 100644
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

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Prep initial charts indexing (#14)

### Default value changes

```diff
# No changes in this release
```

## 0.8.3

**Release date:** 2020-07-27

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* prometheus-cloudwatch-exporter: Add containerPort, livenessProbe.path and readinessProbe.path variables (#22923)

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index c4da6cb3..2b10317e 100644
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

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-cloudwatch-exporter] additional docu on service account usage (#23028)

### Default value changes

```diff
# No changes in this release
```

## 0.8.1

**Release date:** 2020-06-23

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* stable/prometheus-cloudwatch-exporter Add shasum of secrets yaml render to annotations (#22906)

### Default value changes

```diff
# No changes in this release
```

## 0.8.0

**Release date:** 2020-05-21

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-cloudwatch-exporter] Enable to set an existing service account name (#22272)

### Default value changes

```diff
# No changes in this release
```

## 0.7.0

**Release date:** 2020-04-09

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* stable/prometheus-cloudwatch-exporter Add prometheusRule creation. (#21838)

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index 2ddb9e23..c4da6cb3 100644
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

![AppVersion: 0.6.0](https://img.shields.io/static/v1?label=AppVersion&message=0.6.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-cloudwatch-exporter] - Add annotations to serviceAccounts for EKS IAM (#20162)

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index f7c9f83a..2ddb9e23 100644
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

![AppVersion: 0.6.0](https://img.shields.io/static/v1?label=AppVersion&message=0.6.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-cloudwatch-exporter] Upgrade api versions (#17764)

### Default value changes

```diff
# No changes in this release
```

## 0.4.13

**Release date:** 2019-10-28

![AppVersion: 0.6.0](https://img.shields.io/static/v1?label=AppVersion&message=0.6.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-cloudwatch-exporter] Update cloudwatch-exporter version. (#18003)

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index 279be115..f7c9f83a 100644
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

![AppVersion: 0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=0.5.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-cloudwatch-exporter] Add myself to owners/maintainers. (#18093)

### Default value changes

```diff
# No changes in this release
```

## 0.4.11

**Release date:** 2019-10-03

![AppVersion: 0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=0.5.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-cloudwatch-exporter] added changelog (#16579)

### Default value changes

```diff
# No changes in this release
```

## 0.4.10

**Release date:** 2019-08-21

![AppVersion: 0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=0.5.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* include ServiceMonitor relabelings and metricRelabelings (#16443)

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index ecf6e51d..279be115 100644
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

![AppVersion: 0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=0.5.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-cloudwatch-exporter] Fix documentation typos (README.md) (#15745)

### Default value changes

```diff
# No changes in this release
```

## 0.4.8

**Release date:** 2019-07-19

![AppVersion: 0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=0.5.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-cloudwatch-exporter] Added securityContext configuration capabilities (#12405)

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index 2d3094c8..ecf6e51d 100644
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

![AppVersion: 0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=0.5.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-cloudwatch-exporter] configurable container command (#14803)

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index 70faf357..2d3094c8 100644
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

![AppVersion: 0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=0.5.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* add torstenwalter as maintainer (#14788)

### Default value changes

```diff
# No changes in this release
```

## 0.4.5

**Release date:** 2019-05-13

![AppVersion: 0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=0.5.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-cloudwatch-exporter] Add scrape timeout option (#13686)

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index 405bfcc3..70faf357 100644
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

![AppVersion: 0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=0.5.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-cloudwatch-exporter] Added ingress configuration (#12407)

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index eeccd3f2..405bfcc3 100644
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

![AppVersion: 0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=0.5.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Add optional ServiceMonitor and values stub with it disabled by default (#12503)

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index 0f1c8be2..eeccd3f2 100644
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

![AppVersion: 0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=0.5.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Fix default service port and target port.  (#11264)

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index 6117ea98..0f1c8be2 100644
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

![AppVersion: 0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=0.5.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-cloudwatch-exporter] Adjust service targetPort to reference container port name. (#11462)

### Default value changes

```diff
# No changes in this release
```

## 0.4.0

**Release date:** 2019-01-29

![AppVersion: 0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=0.5.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Add support for pre-created secrets and AWS STS session tokens (#10878)

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index e26ef687..6117ea98 100644
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

![AppVersion: 0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=0.5.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Use entrypoint from Docker image and fix probe paths (#10873)

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index c089a7d1..e26ef687 100644
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

![AppVersion: 0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=0.5.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* cloudwatch-exporter - Support custom service labels and portName (#8488)

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index bb986cfd..c089a7d1 100644
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

![AppVersion: 0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=0.5.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-cloudwatch-exporter] health/liveness endpoints added (#7681)

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index 68dda8c2..bb986cfd 100644
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

![AppVersion: 0.1.0](https://img.shields.io/static/v1?label=AppVersion&message=0.1.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-cloudwatch-exporter] Fixed secret variable parsing for access/secret keys. (#6511)

### Default value changes

```diff
# No changes in this release
```

## 0.1.3

**Release date:** 2018-06-29

![AppVersion: 0.1.0](https://img.shields.io/static/v1?label=AppVersion&message=0.1.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Fixed image configuration in README for the prometheus-cloudwatch-exporter (#6400)

### Default value changes

```diff
# No changes in this release
```

## 0.1.2

**Release date:** 2018-04-16

![AppVersion: 0.1.0](https://img.shields.io/static/v1?label=AppVersion&message=0.1.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Annotations and checksum (#5067)

### Default value changes

```diff
diff --git a/charts/prometheus-cloudwatch-exporter/values.yaml b/charts/prometheus-cloudwatch-exporter/values.yaml
index 2c4a90ce..68dda8c2 100644
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

**Release date:** 2018-04-02

![AppVersion: 0.1.0](https://img.shields.io/static/v1?label=AppVersion&message=0.1.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Update table/prometheus-cloudwatch-exporte (#4551)

### Default value changes

```diff
# No changes in this release
```

## 0.1.0

**Release date:** 2018-03-30

![AppVersion: 0.1.0](https://img.shields.io/static/v1?label=AppVersion&message=0.1.0&color=success)
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
