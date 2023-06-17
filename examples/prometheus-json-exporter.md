# Change Log

## 0.7.1

**Release date:** 2023-05-25

![AppVersion: v0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.5.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-json-exporter] Add xiu as a maintainer (#3406)

### Default value changes

```diff
# No changes in this release
```

## 0.7.0

**Release date:** 2023-05-21

![AppVersion: v0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.5.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-json-exporter] Add extra init containers (#3403)

### Default value changes

```diff
diff --git a/charts/prometheus-json-exporter/values.yaml b/charts/prometheus-json-exporter/values.yaml
index 5d3c2b6b..5a62e7a8 100644
--- a/charts/prometheus-json-exporter/values.yaml
+++ b/charts/prometheus-json-exporter/values.yaml
@@ -185,3 +185,10 @@ additionalVolumeMounts: []
   # - name: password-file
   #   mountPath: "/tmp/mysecret.txt"
   #   subPath: mysecret.txt
+
+## Additional init containers
+# These will be added to the prometheus-json-exporter pod.
+extraInitContainers: []
+  # - name: init-myservice
+  #   image: busybox:1.28
+  #   command: [ 'sh', '-c', "sleep 10; done" ]

```

## 0.6.1

**Release date:** 2022-12-22

![AppVersion: v0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.5.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Automatically rollover deployment when configmap changed. (#2831)

### Default value changes

```diff
# No changes in this release
```

## 0.6.0

**Release date:** 2022-12-21

![AppVersion: v0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.5.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-json-exporter] Add support for env variables (#2807)

### Default value changes

```diff
diff --git a/charts/prometheus-json-exporter/values.yaml b/charts/prometheus-json-exporter/values.yaml
index fb1c5c09..5d3c2b6b 100644
--- a/charts/prometheus-json-exporter/values.yaml
+++ b/charts/prometheus-json-exporter/values.yaml
@@ -98,6 +98,15 @@ resources: {}
 #   cpu: 100m
 #   memory: 128Mi
 
+environmentVariables: []
+  # - name: some-secret-variable
+  #   valueFrom:
+  #     secretKeyRef:
+  #       key: some-key
+  #       name: some-secret-name
+  # - name: some-other-variable
+  #   value: some-value
+
 autoscaling:
   enabled: false
   minReplicas: 1

```

## 0.5.0

**Release date:** 2022-10-02

![AppVersion: v0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.5.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-json-exporter] - Allow to pass name of the module for scraping targets for json exporter (#2512)

### Default value changes

```diff
diff --git a/charts/prometheus-json-exporter/values.yaml b/charts/prometheus-json-exporter/values.yaml
index c822dafe..fb1c5c09 100644
--- a/charts/prometheus-json-exporter/values.yaml
+++ b/charts/prometheus-json-exporter/values.yaml
@@ -68,6 +68,7 @@ serviceMonitor:
 #      interval: 60s                    # Scraping interval. Overrides value set in `defaults`
 #      scrapeTimeout: 60s               # Scrape timeout. Overrides value set in `defaults`
 #      additionalMetricsRelabels: {}    # Map of metric labels and values to add
+#      module: example_module           # Name of the module to pick up from `config.yaml` for scraping this target. Optional. Default is `default` provided by the exporter itself.
 
 ingress:
   enabled: false

```

## 0.4.0

**Release date:** 2022-09-03

![AppVersion: v0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.5.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-json-exporter] open service configuration (#2414)

### Default value changes

```diff
diff --git a/charts/prometheus-json-exporter/values.yaml b/charts/prometheus-json-exporter/values.yaml
index dd3d3516..c822dafe 100644
--- a/charts/prometheus-json-exporter/values.yaml
+++ b/charts/prometheus-json-exporter/values.yaml
@@ -42,6 +42,8 @@ securityContext: {}
 service:
   type: ClusterIP
   port: 7979
+  targetPort: http
+  name: http
 
 serviceMonitor:
   ## If true, a ServiceMonitor CRD is created for a prometheus operator

```

## 0.3.0

**Release date:** 2022-08-08

![AppVersion: v0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.5.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-json-exporter] feat: Implement mounting of additional volumes. (#2352)

### Default value changes

```diff
diff --git a/charts/prometheus-json-exporter/values.yaml b/charts/prometheus-json-exporter/values.yaml
index 57d8ce95..dd3d3516 100644
--- a/charts/prometheus-json-exporter/values.yaml
+++ b/charts/prometheus-json-exporter/values.yaml
@@ -110,49 +110,51 @@ affinity: []
 configuration:
   config: |
     ---
-    metrics:
-      - name: example_global_value
-        path: "{ .counter }"
-        help: Example of a top-level global value scrape in the json
-        labels:
-          environment: beta # static label
-          location: "planet-{.location}"          # dynamic label
+    modules:
+      default:
+        metrics:
+          - name: example_global_value
+            path: "{ .counter }"
+            help: Example of a top-level global value scrape in the json
+            labels:
+              environment: beta # static label
+              location: "planet-{.location}"          # dynamic label
 
-      - name: example_value
-        type: object
-        help: Example of sub-level value scrapes from a json
-        path: '{.values[?(@.state == "ACTIVE")]}'
-        labels:
-          environment: beta # static label
-          id: '{.id}'          # dynamic label
-        values:
-          active: 1      # static value
-          count: '{.count}' # dynamic value
-          boolean: '{.some_boolean}'
+          - name: example_value
+            type: object
+            help: Example of sub-level value scrapes from a json
+            path: '{.values[?(@.state == "ACTIVE")]}'
+            labels:
+              environment: beta # static label
+              id: '{.id}'          # dynamic label
+            values:
+              active: 1      # static value
+              count: '{.count}' # dynamic value
+              boolean: '{.some_boolean}'
 
-    headers:
-      X-Dummy: my-test-header
+        headers:
+          X-Dummy: my-test-header
 
-    # If 'body' is set, it will be sent by the exporter as the body content in the scrape request. The HTTP method will also be set as 'POST' in this case.
-    # body:
-    #   content: |
-    #     {"time_diff": "1m25s", "anotherVar": "some value"}
+        # If 'body' is set, it will be sent by the exporter as the body content in the scrape request. The HTTP method will also be set as 'POST' in this case.
+        # body:
+        #   content: |
+        #     {"time_diff": "1m25s", "anotherVar": "some value"}
 
-    # The body content can also be a Go Template (https://golang.org/pkg/text/template), with all the functions from the Sprig library (https://masterminds.github.io/sprig/) available. All the query parameters sent by prometheus in the scrape query to the exporter, are available in the template.
-    # body:
-    #   content: |
-    #     {"time_diff": "{{ duration `95` }}","anotherVar": "{{ .myVal | first }}"}
-    #   templatize: true
+        # The body content can also be a Go Template (https://golang.org/pkg/text/template), with all the functions from the Sprig library (https://masterminds.github.io/sprig/) available. All the query parameters sent by prometheus in the scrape query to the exporter, are available in the template.
+        # body:
+        #   content: |
+        #     {"time_diff": "{{ duration `95` }}","anotherVar": "{{ .myVal | first }}"}
+        #   templatize: true
 
-    # For full http client config parameters, ref: https://pkg.go.dev/github.com/prometheus/common/config?tab=doc#HTTPClientConfig
-    #
-    # http_client_config:
-    #   tls_config:
-    #     insecure_skip_verify: true
-    #   basic_auth:
-    #     username: myuser
-    #     #password: veryverysecret
-    #     password_file: /tmp/mysecret.txt
+        # For full http client config parameters, ref: https://pkg.go.dev/github.com/prometheus/common/config?tab=doc#HTTPClientConfig
+        #
+        # http_client_config:
+        #   tls_config:
+        #     insecure_skip_verify: true
+        #   basic_auth:
+        #     username: myuser
+        #     #password: veryverysecret
+        #     password_file: /tmp/mysecret.txt
 
 ## Custom PrometheusRules to be defined
 ## ref: https://github.com/coreos/prometheus-operator#customresourcedefinitions
@@ -161,3 +163,13 @@ prometheusRule:
   additionalLabels: {}
   namespace: ""
   rules: []
+
+additionalVolumes: []
+  # - name: password-file
+  #   secret:
+  #     secretName: secret-name
+
+additionalVolumeMounts: []
+  # - name: password-file
+  #   mountPath: "/tmp/mysecret.txt"
+  #   subPath: mysecret.txt

```

## 0.2.3

**Release date:** 2022-07-31

![AppVersion: v0.3.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.3.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* fix: Multiple targets are now created as differently named ServiceMonitors. (#2328)

### Default value changes

```diff
# No changes in this release
```

## 0.2.2

**Release date:** 2022-06-12

![AppVersion: v0.3.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.3.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-json-exporter] fix cve-2021-25742 (#2148)

### Default value changes

```diff
# No changes in this release
```

## 0.2.1

**Release date:** 2022-06-07

![AppVersion: v0.3.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.3.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-json-exporter] Allow specifying additional pod labels (#2113)

### Default value changes

```diff
diff --git a/charts/prometheus-json-exporter/values.yaml b/charts/prometheus-json-exporter/values.yaml
index bc8dbced..57d8ce95 100644
--- a/charts/prometheus-json-exporter/values.yaml
+++ b/charts/prometheus-json-exporter/values.yaml
@@ -28,6 +28,9 @@ podAnnotations: []
 podSecurityContext: {}
 # fsGroup: 2000
 
+# podLabels:
+  # Custom labels for the pod
+
 securityContext: {}
 # capabilities:
 #   drop:

```

## 0.2.0

**Release date:** 2022-03-09

![AppVersion: v0.3.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.3.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-json-exporter] Add ability to create a serviceMonitor and prometheus rules (#1858)

### Default value changes

```diff
diff --git a/charts/prometheus-json-exporter/values.yaml b/charts/prometheus-json-exporter/values.yaml
index f7cd487e..bc8dbced 100644
--- a/charts/prometheus-json-exporter/values.yaml
+++ b/charts/prometheus-json-exporter/values.yaml
@@ -40,6 +40,30 @@ service:
   type: ClusterIP
   port: 7979
 
+serviceMonitor:
+  ## If true, a ServiceMonitor CRD is created for a prometheus operator
+  ## https://github.com/coreos/prometheus-operator
+  ##
+  enabled: false
+
+  namespace: ""
+  scheme: http
+
+  # Default values that will be used for all ServiceMonitors created by `targets`
+  defaults:
+    additionalMetricsRelabels: {}
+    interval: 10s
+    labels: {}
+    scrapeTimeout: 30s
+
+  targets:
+#    - name: example                    # Human readable URL that will appear in Prometheus / AlertManager
+#      url: http://example.com/healthz  # The URL that json-exporter will scrape
+#      labels: {}                       # Map of labels for ServiceMonitor. Overrides value set in `defaults`
+#      interval: 60s                    # Scraping interval. Overrides value set in `defaults`
+#      scrapeTimeout: 60s               # Scrape timeout. Overrides value set in `defaults`
+#      additionalMetricsRelabels: {}    # Map of metric labels and values to add
+
 ingress:
   enabled: false
   className: ""
@@ -126,3 +150,11 @@ configuration:
     #     username: myuser
     #     #password: veryverysecret
     #     password_file: /tmp/mysecret.txt
+
+## Custom PrometheusRules to be defined
+## ref: https://github.com/coreos/prometheus-operator#customresourcedefinitions
+prometheusRule:
+  enabled: false
+  additionalLabels: {}
+  namespace: ""
+  rules: []

```

## 0.1.2

**Release date:** 2022-02-23

![AppVersion: v0.3.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.3.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-json-exporter] Add zanhsieh as one of maintainers (#1821)

### Default value changes

```diff
# No changes in this release
```

## 0.1.1

**Release date:** 2022-01-20

![AppVersion: v0.3.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.3.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-json-exporter] Fix typo in deployment nodeselector  (#1716)

### Default value changes

```diff
diff --git a/charts/prometheus-json-exporter/values.yaml b/charts/prometheus-json-exporter/values.yaml
index 93b6df0f..f7cd487e 100644
--- a/charts/prometheus-json-exporter/values.yaml
+++ b/charts/prometheus-json-exporter/values.yaml
@@ -8,7 +8,7 @@ image:
   repository: quay.io/prometheuscommunity/json-exporter
   pullPolicy: IfNotPresent
   # Overrides the image tag whose default is the chart appVersion.
-  tag: "v0.3.0"
+  tag: ""
 
 imagePullSecrets: []
 nameOverride: ""

```

## 0.1.0

**Release date:** 2021-10-29

![AppVersion: 1.0.2](https://img.shields.io/static/v1?label=AppVersion&message=1.0.2&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-json-exporter] New Exporter (#1460)

### Default value changes

```diff
# Default values for prometheus-json-exporter.
# This is a YAML-formatted file.
# Declare variables to be passed into your templates.

replicaCount: 1

image:
  repository: quay.io/prometheuscommunity/json-exporter
  pullPolicy: IfNotPresent
  # Overrides the image tag whose default is the chart appVersion.
  tag: "v0.3.0"

imagePullSecrets: []
nameOverride: ""
fullnameOverride: ""

serviceAccount:
  # Specifies whether a service account should be created
  create: true
  # Annotations to add to the service account
  annotations: []
  # The name of the service account to use.
  # If not set and create is true, a name is generated using the fullname template
  name: ""

podAnnotations: []

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
  port: 7979

ingress:
  enabled: false
  className: ""
  annotations: []
  # kubernetes.io/ingress.class: nginx
  # kubernetes.io/tls-acme: "true"
  hosts:
    - host: chart-example.local
      paths:
        - path: /
          pathType: ImplementationSpecific
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

nodeSelector: []

tolerations: []

affinity: []
configuration:
  config: |
    ---
    metrics:
      - name: example_global_value
        path: "{ .counter }"
        help: Example of a top-level global value scrape in the json
        labels:
          environment: beta # static label
          location: "planet-{.location}"          # dynamic label

      - name: example_value
        type: object
        help: Example of sub-level value scrapes from a json
        path: '{.values[?(@.state == "ACTIVE")]}'
        labels:
          environment: beta # static label
          id: '{.id}'          # dynamic label
        values:
          active: 1      # static value
          count: '{.count}' # dynamic value
          boolean: '{.some_boolean}'

    headers:
      X-Dummy: my-test-header

    # If 'body' is set, it will be sent by the exporter as the body content in the scrape request. The HTTP method will also be set as 'POST' in this case.
    # body:
    #   content: |
    #     {"time_diff": "1m25s", "anotherVar": "some value"}

    # The body content can also be a Go Template (https://golang.org/pkg/text/template), with all the functions from the Sprig library (https://masterminds.github.io/sprig/) available. All the query parameters sent by prometheus in the scrape query to the exporter, are available in the template.
    # body:
    #   content: |
    #     {"time_diff": "{{ duration `95` }}","anotherVar": "{{ .myVal | first }}"}
    #   templatize: true

    # For full http client config parameters, ref: https://pkg.go.dev/github.com/prometheus/common/config?tab=doc#HTTPClientConfig
    #
    # http_client_config:
    #   tls_config:
    #     insecure_skip_verify: true
    #   basic_auth:
    #     username: myuser
    #     #password: veryverysecret
    #     password_file: /tmp/mysecret.txt

```

---
Autogenerated from Helm Chart and git history using [helm-changelog](https://github.com/mogensen/helm-changelog)
