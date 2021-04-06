# Change Log

## 0.9.0 

**Release date:** 2021-03-29

![AppVersion: v0.21.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.21.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [alertmanager] add NodePort support (#790) 

### Default value changes

```diff
diff --git a/charts/alertmanager/values.yaml b/charts/alertmanager/values.yaml
index 12f277f..bcc193b 100644
--- a/charts/alertmanager/values.yaml
+++ b/charts/alertmanager/values.yaml
@@ -52,6 +52,8 @@ service:
   annotations: {}
   type: ClusterIP
   port: 9093
+  # if you want to force a specific nodePort. Must be use with service.type=NodePort
+  # nodePort:
 
 ingress:
   enabled: false
```

## 0.8.0 

**Release date:** 2021-03-19

![AppVersion: v0.21.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.21.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [Alertmanager] command podannotations (#781) 

### Default value changes

```diff
diff --git a/charts/alertmanager/values.yaml b/charts/alertmanager/values.yaml
index aa695a4..12f277f 100644
--- a/charts/alertmanager/values.yaml
+++ b/charts/alertmanager/values.yaml
@@ -85,6 +85,10 @@ affinity: {}
 statefulSet:
   annotations: {}
 
+podAnnotations: {}
+
+command: []
+
 persistence:
   enabled: true
   ## Persistent Volume Storage Class
```

## 0.7.1 

**Release date:** 2021-03-19

![AppVersion: v0.21.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.21.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [Alertmanager]: Fix indent for service tpl annotation (#777) 

### Default value changes

```diff
# No changes in this release
```

## 0.7.0 

**Release date:** 2021-03-18

![AppVersion: v0.21.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.21.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [Alertmanager]: Enable annotations for service tpl (#767) 

### Default value changes

```diff
diff --git a/charts/alertmanager/values.yaml b/charts/alertmanager/values.yaml
index 0e07e8f..aa695a4 100644
--- a/charts/alertmanager/values.yaml
+++ b/charts/alertmanager/values.yaml
@@ -49,6 +49,7 @@ securityContext:
 additionalPeers: []
 
 service:
+  annotations: {}
   type: ClusterIP
   port: 9093
 
```

## 0.6.0 

**Release date:** 2021-02-09

![AppVersion: v0.21.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.21.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Ability for custom dnsConfig in alertmanager (#642) 

### Default value changes

```diff
diff --git a/charts/alertmanager/values.yaml b/charts/alertmanager/values.yaml
index b8acca9..0e07e8f 100644
--- a/charts/alertmanager/values.yaml
+++ b/charts/alertmanager/values.yaml
@@ -27,7 +27,16 @@ serviceAccount:
 
 podSecurityContext:
   fsGroup: 65534
-
+dnsConfig: {}
+  # nameservers:
+  #   - 1.2.3.4
+  # searches:
+  #   - ns1.svc.cluster-domain.example
+  #   - my.dns.search.suffix
+  # options:
+  #   - name: ndots
+  #     value: "2"
+  #   - name: edns0
 securityContext:
   # capabilities:
   #   drop:
```

## 0.5.1 

**Release date:** 2021-02-07

![AppVersion: v0.21.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.21.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* chore: bump configmap reloader (#645) 

### Default value changes

```diff
diff --git a/charts/alertmanager/values.yaml b/charts/alertmanager/values.yaml
index 782a531..b8acca9 100644
--- a/charts/alertmanager/values.yaml
+++ b/charts/alertmanager/values.yaml
@@ -123,7 +123,7 @@ configmapReload:
   ##
   image:
     repository: jimmidyson/configmap-reload
-    tag: v0.4.0
+    tag: v0.5.0
     pullPolicy: IfNotPresent
 
   ## configmap-reload resource requests and limits
```

## 0.5.0 

**Release date:** 2021-01-23

![AppVersion: v0.21.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.21.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [alertmanager] allow possibility to use alertmanager in HA (#609) 

### Default value changes

```diff
diff --git a/charts/alertmanager/values.yaml b/charts/alertmanager/values.yaml
index eb6188d..782a531 100644
--- a/charts/alertmanager/values.yaml
+++ b/charts/alertmanager/values.yaml
@@ -37,6 +37,8 @@ securityContext:
   runAsNonRoot: true
   runAsGroup: 65534
 
+additionalPeers: []
+
 service:
   type: ClusterIP
   port: 9093
```

## 0.4.0 

**Release date:** 2021-01-23

![AppVersion: v0.21.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.21.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [alertmanager] Add optional config reloader (#348) 

### Default value changes

```diff
diff --git a/charts/alertmanager/values.yaml b/charts/alertmanager/values.yaml
index de6954a..eb6188d 100644
--- a/charts/alertmanager/values.yaml
+++ b/charts/alertmanager/values.yaml
@@ -105,5 +105,29 @@ config:
     receiver: default-receiver
     repeat_interval: 3h
 
+## Monitors ConfigMap changes and POSTs to a URL
+## Ref: https://github.com/jimmidyson/configmap-reload
+##
+configmapReload:
+  ## If false, the configmap-reload container will not be deployed
+  ##
+  enabled: false
+
+  ## configmap-reload container name
+  ##
+  name: configmap-reload
+
+  ## configmap-reload container image
+  ##
+  image:
+    repository: jimmidyson/configmap-reload
+    tag: v0.4.0
+    pullPolicy: IfNotPresent
+
+  ## configmap-reload resource requests and limits
+  ## Ref: http://kubernetes.io/docs/user-guide/compute-resources/
+  ##
+  resources: {}
+
 templates: {}
 #   alertmanager.tmpl: |-
```

## 0.3.0 

**Release date:** 2020-12-12

![AppVersion: v0.21.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.21.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [alertmanager] Allow passing notification templates (#484) 

### Default value changes

```diff
diff --git a/charts/alertmanager/values.yaml b/charts/alertmanager/values.yaml
index acbd738..de6954a 100644
--- a/charts/alertmanager/values.yaml
+++ b/charts/alertmanager/values.yaml
@@ -90,6 +90,9 @@ config:
   global: {}
     # slack_api_url: ''
 
+  templates:
+    - '/etc/alertmanager/*.tmpl'
+
   receivers:
     - name: default-receiver
       # slack_configs:
@@ -101,3 +104,6 @@ config:
     group_interval: 5m
     receiver: default-receiver
     repeat_interval: 3h
+
+templates: {}
+#   alertmanager.tmpl: |-
```

## 0.2.0 

**Release date:** 2020-11-25

![AppVersion: v0.21.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.21.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [alertmanager] add statefulSet annotations (#414) 

### Default value changes

```diff
diff --git a/charts/alertmanager/values.yaml b/charts/alertmanager/values.yaml
index 70e9a7d..acbd738 100644
--- a/charts/alertmanager/values.yaml
+++ b/charts/alertmanager/values.yaml
@@ -70,6 +70,9 @@ tolerations: []
 
 affinity: {}
 
+statefulSet:
+  annotations: {}
+
 persistence:
   enabled: true
   ## Persistent Volume Storage Class
```

## 0.1.4 

**Release date:** 2020-10-28

![AppVersion: v0.21.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.21.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* alertmanager: add maintainer (#276) 

### Default value changes

```diff
# No changes in this release
```

## 0.1.3 

**Release date:** 2020-10-21

![AppVersion: v0.21.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.21.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [alertmanager] Add support change default config (#240) 

### Default value changes

```diff
diff --git a/charts/alertmanager/values.yaml b/charts/alertmanager/values.yaml
index 6f4a487..70e9a7d 100644
--- a/charts/alertmanager/values.yaml
+++ b/charts/alertmanager/values.yaml
@@ -10,6 +10,8 @@ image:
   # Overrides the image tag whose default is the chart appVersion.
   tag: ""
 
+extraArgs: {}
+
 imagePullSecrets: []
 nameOverride: ""
 fullnameOverride: ""
```

## 0.1.2 

**Release date:** 2020-10-20

![AppVersion: v0.21.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.21.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Fix alertmanager statefulset template (#236) 

### Default value changes

```diff
# No changes in this release
```

## 0.1.1 

**Release date:** 2020-10-02

![AppVersion: v0.21.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.21.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [alertmanager ]fix persistence (#174) 

### Default value changes

```diff
diff --git a/charts/alertmanager/values.yaml b/charts/alertmanager/values.yaml
index b8dac66..6f4a487 100644
--- a/charts/alertmanager/values.yaml
+++ b/charts/alertmanager/values.yaml
@@ -69,7 +69,7 @@ tolerations: []
 affinity: {}
 
 persistence:
-  enabled: false
+  enabled: true
   ## Persistent Volume Storage Class
   ## If defined, storageClassName: <storageClass>
   ## If set to "-", storageClassName: "", which disables dynamic provisioning
@@ -79,7 +79,7 @@ persistence:
   # storageClass: "-"
   accessModes:
     - ReadWriteOnce
-  size: 1Gi
+  size: 50Mi
 
 config:
   global: {}
```

## 0.1.0 

**Release date:** 2020-09-27

![AppVersion: v0.21.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.21.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [alertmanager] add new chart (#64) 

### Default value changes

```diff
# Default values for alertmanager.
# This is a YAML-formatted file.
# Declare variables to be passed into your templates.

replicaCount: 1

image:
  repository: prom/alertmanager
  pullPolicy: IfNotPresent
  # Overrides the image tag whose default is the chart appVersion.
  tag: ""

imagePullSecrets: []
nameOverride: ""
fullnameOverride: ""

serviceAccount:
  # Specifies whether a service account should be created
  create: true
  # Annotations to add to the service account
  annotations: {}
  # The name of the service account to use.
  # If not set and create is true, a name is generated using the fullname template
  name:

podSecurityContext:
  fsGroup: 65534

securityContext:
  # capabilities:
  #   drop:
  #   - ALL
  # readOnlyRootFilesystem: true
  runAsUser: 65534
  runAsNonRoot: true
  runAsGroup: 65534

service:
  type: ClusterIP
  port: 9093

ingress:
  enabled: false
  annotations: {}
  hosts:
    - host: alertmanager.domain.com
      paths: []
  tls: []
  #  - secretName: chart-example-tls
  #    hosts:
  #      - alertmanager.domain.com

resources: {}
  # We usually recommend not to specify default resources and to leave this as a conscious
  # choice for the user. This also increases chances charts run on environments with little
  # resources, such as Minikube. If you do want to specify resources, uncomment the following
  # lines, adjust them as necessary, and remove the curly braces after 'resources:'.
  # limits:
  #   cpu: 100m
  #   memory: 128Mi
  # requests:
  #   cpu: 10m
  #   memory: 32Mi

nodeSelector: {}

tolerations: []

affinity: {}

persistence:
  enabled: false
  ## Persistent Volume Storage Class
  ## If defined, storageClassName: <storageClass>
  ## If set to "-", storageClassName: "", which disables dynamic provisioning
  ## If undefined (the default) or set to null, no storageClassName spec is
  ## set, choosing the default provisioner.
  ##
  # storageClass: "-"
  accessModes:
    - ReadWriteOnce
  size: 1Gi

config:
  global: {}
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
```

---
Autogenerated from Helm Chart and git history using [helm-changelog](https://github.com/mogensen/helm-changelog)
