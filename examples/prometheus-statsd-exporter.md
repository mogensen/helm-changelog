# Change Log

## 0.3.1 

**Release date:** 2021-03-10

![AppVersion: 0.20.0](https://img.shields.io/static/v1?label=AppVersion&message=0.20.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* prometheus-statsd-exporter - Fix statsd service monitor config (#732) 

### Default value changes

```diff
# No changes in this release
```

## 0.3.0 

**Release date:** 2021-02-09

![AppVersion: 0.20.0](https://img.shields.io/static/v1?label=AppVersion&message=0.20.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* update statsd-exporter version (#658) 

### Default value changes

```diff
diff --git a/charts/prometheus-statsd-exporter/values.yaml b/charts/prometheus-statsd-exporter/values.yaml
index a7fc28f..26a99a0 100644
--- a/charts/prometheus-statsd-exporter/values.yaml
+++ b/charts/prometheus-statsd-exporter/values.yaml
@@ -9,7 +9,7 @@ image:
   repository: prom/statsd-exporter
   pullPolicy: IfNotPresent
   # Overrides the image tag whose default is the chart appVersion.
-  tag: v0.18.0
+  tag: v0.20.0
 
 imagePullSecrets: []
 nameOverride: ""
```

## 0.2.0 

**Release date:** 2021-01-09

![AppVersion: 0.18.0](https://img.shields.io/static/v1?label=AppVersion&message=0.18.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-statsd-exporter] Adding `honorLabels` to `statsd-exporter` `ServiceMonitor` (#534) 

### Default value changes

```diff
diff --git a/charts/prometheus-statsd-exporter/values.yaml b/charts/prometheus-statsd-exporter/values.yaml
index d10d3d1..a7fc28f 100644
--- a/charts/prometheus-statsd-exporter/values.yaml
+++ b/charts/prometheus-statsd-exporter/values.yaml
@@ -47,6 +47,7 @@ serviceMonitor:
   interval: 30s
   scrapeTimeout: 10s
   namespace: monitoring
+  honorLabels: false
   additionalLabels: {}
 
 serviceAccount:
```

## 0.1.0 

**Release date:** 2020-12-08

![AppVersion: 0.18.0](https://img.shields.io/static/v1?label=AppVersion&message=0.18.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-statsd-exporter] add first release (#449) 

### Default value changes

```diff
# Default values for prometheus-statsd-exporter.
# This is a YAML-formatted file.
# Declare variables to be passed into your templates.

replicaCount: 1
# deploymentRevisionHistoryLimit: 10

image:
  repository: prom/statsd-exporter
  pullPolicy: IfNotPresent
  # Overrides the image tag whose default is the chart appVersion.
  tag: v0.18.0

imagePullSecrets: []
nameOverride: ""
fullnameOverride: ""


statsd:
  # The UDP port on which to receive statsd metric lines.
  udpPort: 9125

  # The TCP port on which to receive statsd metric lines.
  tcpPort: 9125

  # Maximum size of your metric mapping cache.
  # Relies on least recently used replacement policy if max size is reached.
  cacheSize: 1000

  # Size of internal queue for processing events.
  eventQueueSize: 10000

  # Number of events to hold in queue before flushing.
  eventFlushThreshold: 1000

  # Time interval before flushing events in queue.
  eventFlushInterval: 200ms

  # Metric mapping ConfigMap
#  mappingConfigMapName: ""

  # Metric mapping configuration
#  mappingConfig: |-

serviceMonitor:
  enabled: false
  interval: 30s
  scrapeTimeout: 10s
  namespace: monitoring
  additionalLabels: {}

serviceAccount:
  # Specifies whether a service account should be created
  create: true
  # Annotations to add to the service account
  annotations: {}
  # The name of the service account to use.
  # If not set and create is true, a name is generated using the fullname template
  name: ""

podAnnotations: {}

podSecurityContext: {}
  # fsGroup: 2000

securityContext: {}
  # capabilities:
  #   drop:
  #   - ALL
  # readOnlyRootFilesystem: true
  # runAsNonRoot: true
  # runAsUser: 1000

service:
  type: ClusterIP
  # The address on which to expose the web interface and generated Prometheus metrics.
  port: 9102
  # Path under which to expose metrics.
  path: /metrics
  annotations: {}

ingress:
  enabled: false
  annotations: {}
    # kubernetes.io/ingress.class: nginx
    # kubernetes.io/tls-acme: "true"
  hosts:
    - host: chart-example.local
      #  pathType: ImplementationSpecific # only Kubernetes v1.19+
      paths: []
  tls: []
  #  - secretName: chart-example-tls
  #    hosts:
  #      - chart-example.local

resources: {}
  # We usually recommend not to specify default resources and to leave this as a conscious
  # choice for the user. This also increases chances charts run on environments with little
  # resources, such as Minikube. If you do want to specify resources, uncomment the following
  # lines, adjust them as necessary, and remove the curly braces after 'resources:'.
  # limits:
  #   cpu: 100m
  #   memory: 128Mi
  # requests:
  #   cpu: 100m
  #   memory: 128Mi

autoscaling:
  enabled: false
  minReplicas: 1
  maxReplicas: 100
  targetCPUUtilizationPercentage: 80
  # targetMemoryUtilizationPercentage: 80

nodeSelector: {}

tolerations: []

affinity: {}
```

---
Autogenerated from Helm Chart and git history using [helm-changelog](https://github.com/mogensen/helm-changelog)
