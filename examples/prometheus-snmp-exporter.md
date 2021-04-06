# Change Log

## 0.1.2 

**Release date:** 2021-02-09

![AppVersion: 0.19.0](https://img.shields.io/static/v1?label=AppVersion&message=0.19.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* chore: bump configmap reloader (#646) 

### Default value changes

```diff
diff --git a/charts/prometheus-snmp-exporter/values.yaml b/charts/prometheus-snmp-exporter/values.yaml
index 32a7bc7..b067cd8 100644
--- a/charts/prometheus-snmp-exporter/values.yaml
+++ b/charts/prometheus-snmp-exporter/values.yaml
@@ -87,7 +87,7 @@ configmapReload:
   ##
   image:
     repository: jimmidyson/configmap-reload
-    tag: v0.2.2
+    tag: v0.5.0
     pullPolicy: IfNotPresent
 
   ## configmap-reload resource requests and limits
```

## 0.1.1 

**Release date:** 2020-11-26

![AppVersion: 0.19.0](https://img.shields.io/static/v1?label=AppVersion&message=0.19.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-snmp-exporter] Fixes #398 and adds some features (#399) 

### Default value changes

```diff
# No changes in this release
```

## 0.1.0 

**Release date:** 2020-11-22

![AppVersion: 0.19.0](https://img.shields.io/static/v1?label=AppVersion&message=0.19.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-snmp-exporter] Add support mount extra secret and configmap  (#218) 

### Default value changes

```diff
diff --git a/charts/prometheus-snmp-exporter/values.yaml b/charts/prometheus-snmp-exporter/values.yaml
index 0695068..32a7bc7 100644
--- a/charts/prometheus-snmp-exporter/values.yaml
+++ b/charts/prometheus-snmp-exporter/values.yaml
@@ -2,7 +2,7 @@ restartPolicy: Always
 
 image:
   repository: prom/snmp-exporter
-  tag: v0.14.0
+  tag: v0.19.0
   pullPolicy: IfNotPresent
 
 nodeSelector: {}
@@ -11,6 +11,23 @@ affinity: {}
 
 # config:
 
+extraConfigmapMounts: []
+  # - name: snmp-exporter-configmap
+  #   mountPath: /run/secrets/snmp-exporter
+  #   subPath: snmp.yaml # (optional)
+  #   configMap: snmp-exporter-configmap-configmap
+  #   readOnly: true
+  #   defaultMode: 420
+
+## Additional secret mounts
+# Defines additional mounts with secrets. Secrets must be manually created in the namespace.
+extraSecretMounts: []
+  # - name: secret-files
+  #   mountPath: /run/secrets/snmp-exporter
+  #   secretName: snmp-exporter-secret-files
+  #   readOnly: true
+  #   defaultMode: 420
+
 ## For RBAC support:
 rbac:
   # Specifies whether RBAC resources should be created
```

## 0.0.6 

**Release date:** 2020-08-20

![AppVersion: 0.14.0](https://img.shields.io/static/v1?label=AppVersion&message=0.14.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Prep initial charts indexing (#14) 

### Default value changes

```diff
# No changes in this release
```

## 0.0.5 

**Release date:** 2020-03-14

![AppVersion: 0.14.0](https://img.shields.io/static/v1?label=AppVersion&message=0.14.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-snmp-exporter] Fix serviceMonitor params (#19985) 

### Default value changes

```diff
# No changes in this release
```

## 0.0.4 

**Release date:** 2019-06-27

![AppVersion: 0.14.0](https://img.shields.io/static/v1?label=AppVersion&message=0.14.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-snmp-exporter] Don't run config through toYaml (#13884) 

### Default value changes

```diff
# No changes in this release
```

## 0.0.3 

**Release date:** 2019-06-11

![AppVersion: 0.14.0](https://img.shields.io/static/v1?label=AppVersion&message=0.14.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-snmp-exporter] add namespace metadata (#14328) 

### Default value changes

```diff
# No changes in this release
```

## 0.0.2 

**Release date:** 2019-03-19

![AppVersion: 0.14.0](https://img.shields.io/static/v1?label=AppVersion&message=0.14.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Add servicemonitor to prometheus-snmp-exporter chart (#12320) 

### Default value changes

```diff
diff --git a/charts/prometheus-snmp-exporter/values.yaml b/charts/prometheus-snmp-exporter/values.yaml
index 38dee25..0695068 100644
--- a/charts/prometheus-snmp-exporter/values.yaml
+++ b/charts/prometheus-snmp-exporter/values.yaml
@@ -77,3 +77,31 @@ configmapReload:
   ## Ref: http://kubernetes.io/docs/user-guide/compute-resources/
   ##
   resources: {}
+
+# Enable this if you're using https://github.com/coreos/prometheus-operator
+serviceMonitor:
+  enabled: false
+  namespace: monitoring
+
+  # fallback to the prometheus default unless specified
+  # interval: 10s
+
+  ## Defaults to what's used if you follow CoreOS [Prometheus Install Instructions](https://github.com/helm/charts/tree/master/stable/prometheus-operator#tldr)
+  ## [Prometheus Selector Label](https://github.com/helm/charts/tree/master/stable/prometheus-operator#prometheus-operator-1)
+  ## [Kube Prometheus Selector Label](https://github.com/helm/charts/tree/master/stable/prometheus-operator#exporters)
+  selector:
+    prometheus: kube-prometheus
+  # Retain the job and instance labels of the metrics pushed to the snmp-exporter
+  # [Scraping SNMP-exporter](https://github.com/prometheus/snmp_exporter#configure-the-snmp_exporter-as-a-target-to-scrape)
+  honorLabels: true
+
+  params:
+    enabled: false
+    conf:
+      module:
+      - if_mib
+      target:
+      - 127.0.0.1
+
+  path: /snmp
+  scrapeTimeout: 10s
```

## 0.0.1 

**Release date:** 2019-01-21

![AppVersion: 0.14.0](https://img.shields.io/static/v1?label=AppVersion&message=0.14.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Add prometheus-snmp-exporter chart (#10734) 

### Default value changes

```diff
restartPolicy: Always

image:
  repository: prom/snmp-exporter
  tag: v0.14.0
  pullPolicy: IfNotPresent

nodeSelector: {}
tolerations: []
affinity: {}

# config:

## For RBAC support:
rbac:
  # Specifies whether RBAC resources should be created
  create: true

serviceAccount:
  # Specifies whether a ServiceAccount should be created
  create: true

  # The name of the ServiceAccount to use.
  # If not set and create is true, a name is generated using the fullname template
  name:

resources: {}
  # limits:
  #   memory: 300Mi
  # requests:
  #   memory: 50Mi

service:
  annotations: {}
  type: ClusterIP
  port: 9116

## An Ingress resource can provide name-based virtual hosting and TLS
## termination among other things for CouchDB deployments which are accessed
## from outside the Kubernetes cluster.
## ref: https://kubernetes.io/docs/concepts/services-networking/ingress/
ingress:
  enabled: false
  hosts: []
     # - chart-example.local
  annotations: {}
    # kubernetes.io/ingress.class: nginx
    # kubernetes.io/tls-acme: "true"
  tls: []
    # Secrets must be manually created in the namespace.
    # - secretName: chart-example-tls
    #   hosts:
    #     - chart-example.local

podAnnotations: {}

extraArgs: []
#  --history.limit=1000

replicas: 1
## Monitors ConfigMap changes and POSTs to a URL
## Ref: https://github.com/jimmidyson/configmap-reload
##
configmapReload:
  ## configmap-reload container name
  ##
  name: configmap-reload

  ## configmap-reload container image
  ##
  image:
    repository: jimmidyson/configmap-reload
    tag: v0.2.2
    pullPolicy: IfNotPresent

  ## configmap-reload resource requests and limits
  ## Ref: http://kubernetes.io/docs/user-guide/compute-resources/
  ##
  resources: {}
```

---
Autogenerated from Helm Chart and git history using [helm-changelog](https://github.com/mogensen/helm-changelog)
