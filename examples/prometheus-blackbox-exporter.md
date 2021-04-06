# Change Log

## 4.10.3 

**Release date:** 2021-01-21

![AppVersion: 0.18.0](https://img.shields.io/static/v1?label=AppVersion&message=0.18.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-blackbox-exporter] Fix chart name in readme (#604) 

### Default value changes

```diff
# No changes in this release
```

## 4.10.2 

**Release date:** 2021-01-15

![AppVersion: 0.18.0](https://img.shields.io/static/v1?label=AppVersion&message=0.18.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Changing default container port in prometheus-blackbox-exporter (#445) 

### Default value changes

```diff
diff --git a/charts/prometheus-blackbox-exporter/values.yaml b/charts/prometheus-blackbox-exporter/values.yaml
index 6b467c2..ccf34d9 100644
--- a/charts/prometheus-blackbox-exporter/values.yaml
+++ b/charts/prometheus-blackbox-exporter/values.yaml
@@ -132,7 +132,11 @@ service:
   labels: {}
   type: ClusterIP
   port: 9115
-  targetPort: 9115
+
+# Only changes container port. Application port can be changed with extraArgs (--web.listen-address=:9115)
+# https://github.com/prometheus/blackbox_exporter/blob/998037b5b40c1de5fee348ffdea8820509d85171/main.go#L55
+containerPort: 9115
+
 serviceAccount:
   # Specifies whether a ServiceAccount should be created
   create: true
```

## 4.10.1 

**Release date:** 2020-11-19

![AppVersion: 0.18.0](https://img.shields.io/static/v1?label=AppVersion&message=0.18.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-blackbox-exporter] add detail comment to the option allowMonitoringNamespace in values.yaml (#378) 

### Default value changes

```diff
diff --git a/charts/prometheus-blackbox-exporter/values.yaml b/charts/prometheus-blackbox-exporter/values.yaml
index bb7cdfe..6b467c2 100644
--- a/charts/prometheus-blackbox-exporter/values.yaml
+++ b/charts/prometheus-blackbox-exporter/values.yaml
@@ -211,6 +211,8 @@ networkPolicy:
   # Enable network policy and allow access from anywhere
   enabled: false
   # Limit access only from monitoring namespace
+  # Before setting this value to true, you must add the name=monitoring label to the monitoring namespace
+  # Network Policy uses label filtering
   allowMonitoringNamespace: false
 
 ## dnsPolicy and dnsConfig for Deployments and Daemonsets if you want non-default settings.
```

## 4.10.0 

**Release date:** 2020-10-25

![AppVersion: 0.18.0](https://img.shields.io/static/v1?label=AppVersion&message=0.18.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-blackbox-exporter] Feat use external secret for configura… (#139) 

### Default value changes

```diff
diff --git a/charts/prometheus-blackbox-exporter/values.yaml b/charts/prometheus-blackbox-exporter/values.yaml
index 1b9778e..bb7cdfe 100644
--- a/charts/prometheus-blackbox-exporter/values.yaml
+++ b/charts/prometheus-blackbox-exporter/values.yaml
@@ -84,6 +84,11 @@ nodeSelector: {}
 tolerations: []
 affinity: {}
 
+# if the configuration is managed as secret outside the chart, using SealedSecret for example,
+# provide the name of the secret here. If secretConfig is set to true, configExistingSecretName will be ignored
+# in favor of the config value.
+configExistingSecretName: ""
+# Store the configuration as a `Secret` instead of a `ConfigMap`, useful in case it contains sensitive data
 secretConfig: false
 config:
   modules:
```

## 4.9.1 

**Release date:** 2020-10-18

![AppVersion: 0.18.0](https://img.shields.io/static/v1?label=AppVersion&message=0.18.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-blackbox-exporter] fix service monitor template (#227) 

### Default value changes

```diff
# No changes in this release
```

## 4.9.0 

**Release date:** 2020-10-17

![AppVersion: 0.18.0](https://img.shields.io/static/v1?label=AppVersion&message=0.18.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Bump blackbox exporter to v0.18.0. (#226) 

### Default value changes

```diff
diff --git a/charts/prometheus-blackbox-exporter/values.yaml b/charts/prometheus-blackbox-exporter/values.yaml
index 2542314..1b9778e 100644
--- a/charts/prometheus-blackbox-exporter/values.yaml
+++ b/charts/prometheus-blackbox-exporter/values.yaml
@@ -54,7 +54,7 @@ strategy:
 
 image:
   repository: prom/blackbox-exporter
-  tag: v0.17.0
+  tag: v0.18.0
   pullPolicy: IfNotPresent
 
   ## Optionally specify an array of imagePullSecrets.
```

## 4.8.0 

**Release date:** 2020-10-16

![AppVersion: 0.17.0](https://img.shields.io/static/v1?label=AppVersion&message=0.17.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-blackbox-exporter] Added support of sidecars and extraVolumes (#194) 

### Default value changes

```diff
diff --git a/charts/prometheus-blackbox-exporter/values.yaml b/charts/prometheus-blackbox-exporter/values.yaml
index 4d395d1..2542314 100644
--- a/charts/prometheus-blackbox-exporter/values.yaml
+++ b/charts/prometheus-blackbox-exporter/values.yaml
@@ -13,6 +13,36 @@ podDisruptionBudget: {}
 ##   NO_PROXY: "localhost,127.0.0.1"
 extraEnv: {}
 
+extraVolumes: []
+  # - name: secret-blackbox-oauth-htpasswd
+  #   secret:
+  #     defaultMode: 420
+  #     secretName: blackbox-oauth-htpasswd
+  # - name: storage-volume
+  #   persistentVolumeClaim:
+  #     claimName: example
+
+extraContainers: []
+  # - name: oAuth2-proxy
+  #   args:
+  #     - -https-address=:9116
+  #     - -upstream=http://localhost:9115
+  #     - -skip-auth-regex=^/metrics
+  #     - -openshift-delegate-urls={"/":{"group":"monitoring.coreos.com","resource":"prometheuses","verb":"get"}}
+  #   image: openshift/oauth-proxy:v1.1.0
+  #   ports:
+  #       - containerPort: 9116
+  #         name: proxy
+  #   resources:
+  #       limits:
+  #         memory: 16Mi
+  #       requests:
+  #         memory: 4Mi
+  #         cpu: 20m
+  #   volumeMounts:
+  #     - mountPath: /etc/prometheus/secrets/blackbox-tls
+  #       name: secret-blackbox-tls
+
 ## Enable pod security policy
 pspEnabled: true
 
@@ -39,6 +69,7 @@ runAsUser: 1000
 readOnlyRootFilesystem: true
 runAsNonRoot: true
 
+
 livenessProbe:
   httpGet:
     path: /health
@@ -96,7 +127,7 @@ service:
   labels: {}
   type: ClusterIP
   port: 9115
-
+  targetPort: 9115
 serviceAccount:
   # Specifies whether a ServiceAccount should be created
   create: true
@@ -146,6 +177,12 @@ serviceMonitor:
     interval: 30s
     scrapeTimeout: 30s
     module: http_2xx
+  ## scheme: HTTP scheme to use for scraping. Can be used with `tlsConfig` for example if using istio mTLS.
+  scheme: http
+  ## tlsConfig: TLS configuration to use when scraping the endpoint. For example if using istio mTLS.
+  ## Of type: https://github.com/coreos/prometheus-operator/blob/master/Documentation/api.md#tlsconfig
+  tlsConfig: {}
+  bearerTokenFile:
 
   targets:
 #    - name: example                    # Human readable URL that will appear in Prometheus / AlertManager
```

## 4.7.0 

**Release date:** 2020-10-07

![AppVersion: 0.17.0](https://img.shields.io/static/v1?label=AppVersion&message=0.17.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-blackbox-exporter] Add additional env variables to blackbox-exporter container (#182) 

### Default value changes

```diff
diff --git a/charts/prometheus-blackbox-exporter/values.yaml b/charts/prometheus-blackbox-exporter/values.yaml
index 564ef2e..4d395d1 100644
--- a/charts/prometheus-blackbox-exporter/values.yaml
+++ b/charts/prometheus-blackbox-exporter/values.yaml
@@ -5,6 +5,14 @@ kind: Deployment
 podDisruptionBudget: {}
   # maxUnavailable: 0
 
+## Additional blackbox-exporter container environment variables
+## For instance to add a http_proxy
+##
+## extraEnv:
+##   HTTP_PROXY: "http://superproxy.com:3128"
+##   NO_PROXY: "localhost,127.0.0.1"
+extraEnv: {}
+
 ## Enable pod security policy
 pspEnabled: true
 
```

## 4.6.0 

**Release date:** 2020-09-22

![AppVersion: 0.17.0](https://img.shields.io/static/v1?label=AppVersion&message=0.17.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-blackbox-exporter] Allow custom labels for pods. Add service labels to values.yaml. (#125) 

### Default value changes

```diff
diff --git a/charts/prometheus-blackbox-exporter/values.yaml b/charts/prometheus-blackbox-exporter/values.yaml
index 3fae201..564ef2e 100644
--- a/charts/prometheus-blackbox-exporter/values.yaml
+++ b/charts/prometheus-blackbox-exporter/values.yaml
@@ -85,6 +85,7 @@ priorityClassName: ""
 
 service:
   annotations: {}
+  labels: {}
   type: ClusterIP
   port: 9115
 
@@ -116,6 +117,9 @@ ingress:
 
 podAnnotations: {}
 
+pod:
+  labels: {}
+
 extraArgs: []
 #  --history.limit=1000
 
```

## 4.5.2 

**Release date:** 2020-09-18

![AppVersion: 0.17.0](https://img.shields.io/static/v1?label=AppVersion&message=0.17.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* fix http_2xx defaults (#101) 

### Default value changes

```diff
diff --git a/charts/prometheus-blackbox-exporter/values.yaml b/charts/prometheus-blackbox-exporter/values.yaml
index b7db0e0..3fae201 100644
--- a/charts/prometheus-blackbox-exporter/values.yaml
+++ b/charts/prometheus-blackbox-exporter/values.yaml
@@ -52,7 +52,7 @@ config:
       prober: http
       timeout: 5s
       http:
-        valid_http_versions: ["HTTP/1.1", "HTTP/2"]
+        valid_http_versions: ["HTTP/1.1", "HTTP/2.0"]
         no_follow_redirects: false
         preferred_ip_protocol: "ip4"
 
```

## 4.5.1 

**Release date:** 2020-09-08

![AppVersion: 0.17.0](https://img.shields.io/static/v1?label=AppVersion&message=0.17.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [blackbox-exporter] added monotek to maintainers (#58) 

### Default value changes

```diff
# No changes in this release
```

## 4.5.0 

**Release date:** 2020-09-08

![AppVersion: 0.17.0](https://img.shields.io/static/v1?label=AppVersion&message=0.17.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-blackbox-exporter] Add additional metrics relabels blackbox and update to 0.17.0 (#69) 

### Default value changes

```diff
diff --git a/charts/prometheus-blackbox-exporter/values.yaml b/charts/prometheus-blackbox-exporter/values.yaml
index e6e8f1e..b7db0e0 100644
--- a/charts/prometheus-blackbox-exporter/values.yaml
+++ b/charts/prometheus-blackbox-exporter/values.yaml
@@ -16,7 +16,7 @@ strategy:
 
 image:
   repository: prom/blackbox-exporter
-  tag: v0.16.0
+  tag: v0.17.0
   pullPolicy: IfNotPresent
 
   ## Optionally specify an array of imagePullSecrets.
@@ -129,6 +129,7 @@ serviceMonitor:
 
   # Default values that will be used for all ServiceMonitors created by `targets`
   defaults:
+    additionalMetricsRelabels: {}
     labels: {}
     interval: 30s
     scrapeTimeout: 30s
@@ -137,10 +138,11 @@ serviceMonitor:
   targets:
 #    - name: example                    # Human readable URL that will appear in Prometheus / AlertManager
 #      url: http://example.com/healthz  # The URL that blackbox will scrape
-#      labels: {}                       # List of labels for ServiceMonitor. Overrides value set in `defaults`
+#      labels: {}                       # Map of labels for ServiceMonitor. Overrides value set in `defaults`
 #      interval: 60s                    # Scraping interval. Overrides value set in `defaults`
 #      scrapeTimeout: 60s               # Scrape timeout. Overrides value set in `defaults`
 #      module: http_2xx                 # Module used for scraping. Overrides value set in `defaults`
+#      additionalMetricsRelabels: {}    # Map of metric labels and values to add
 
 ## Custom PrometheusRules to be defined
 ## ref: https://github.com/coreos/prometheus-operator#customresourcedefinitions
```

## 4.4.0 

**Release date:** 2020-09-08

![AppVersion: 0.16.0](https://img.shields.io/static/v1?label=AppVersion&message=0.16.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-blackbox-exporter] allow configurable DNS settings (#42) 

### Default value changes

```diff
diff --git a/charts/prometheus-blackbox-exporter/values.yaml b/charts/prometheus-blackbox-exporter/values.yaml
index 5797022..e6e8f1e 100644
--- a/charts/prometheus-blackbox-exporter/values.yaml
+++ b/charts/prometheus-blackbox-exporter/values.yaml
@@ -156,3 +156,8 @@ networkPolicy:
   enabled: false
   # Limit access only from monitoring namespace
   allowMonitoringNamespace: false
+
+## dnsPolicy and dnsConfig for Deployments and Daemonsets if you want non-default settings.
+## These will be passed directly to the PodSpec of same.
+dnsPolicy:
+dnsConfig:
```

## 4.3.2 

**Release date:** 2020-09-07

![AppVersion: 0.16.0](https://img.shields.io/static/v1?label=AppVersion&message=0.16.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-blackbox-exporter] fix linting failure due to deprecated api version (see issue #56) (#57) 

### Default value changes

```diff
# No changes in this release
```

## 4.3.1 

**Release date:** 2020-08-20

![AppVersion: 0.16.0](https://img.shields.io/static/v1?label=AppVersion&message=0.16.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Prep initial charts indexing (#14) 

### Default value changes

```diff
# No changes in this release
```

## 4.3.0 

**Release date:** 2020-08-04

![AppVersion: 0.16.0](https://img.shields.io/static/v1?label=AppVersion&message=0.16.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-blackbox-exporter] added network policy (#23273) 

### Default value changes

```diff
diff --git a/charts/prometheus-blackbox-exporter/values.yaml b/charts/prometheus-blackbox-exporter/values.yaml
index e8488f6..5797022 100644
--- a/charts/prometheus-blackbox-exporter/values.yaml
+++ b/charts/prometheus-blackbox-exporter/values.yaml
@@ -149,3 +149,10 @@ prometheusRule:
   additionalLabels: {}
   namespace: ""
   rules: []
+
+## Network policy for chart
+networkPolicy:
+  # Enable network policy and allow access from anywhere
+  enabled: false
+  # Limit access only from monitoring namespace
+  allowMonitoringNamespace: false
```

## 4.2.2 

**Release date:** 2020-08-03

![AppVersion: 0.16.0](https://img.shields.io/static/v1?label=AppVersion&message=0.16.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* fix(helm#23318): [stable/prometheus-blackbox-exporter] (#23319) 

### Default value changes

```diff
# No changes in this release
```

## 4.2.1 

**Release date:** 2020-08-03

![AppVersion: 0.16.0](https://img.shields.io/static/v1?label=AppVersion&message=0.16.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-blackbox-exporter] remove replicas and strategy keys from daemonset template (#23421) 

### Default value changes

```diff
# No changes in this release
```

## 4.2.0 

**Release date:** 2020-07-27

![AppVersion: 0.16.0](https://img.shields.io/static/v1?label=AppVersion&message=0.16.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Fix blackbox target's metrics labels (#22546) 

### Default value changes

```diff
# No changes in this release
```

## 4.1.1 

**Release date:** 2020-06-15

![AppVersion: 0.16.0](https://img.shields.io/static/v1?label=AppVersion&message=0.16.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-blackbox-exporter] Fix broken additionalLabels for protmetheusrule (#22800) 

### Default value changes

```diff
# No changes in this release
```

## 4.1.0 

**Release date:** 2020-05-14

![AppVersion: 0.16.0](https://img.shields.io/static/v1?label=AppVersion&message=0.16.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Add vars for liveness and readyness (#22201) 

### Default value changes

```diff
diff --git a/charts/prometheus-blackbox-exporter/values.yaml b/charts/prometheus-blackbox-exporter/values.yaml
index 59c0955..e8488f6 100644
--- a/charts/prometheus-blackbox-exporter/values.yaml
+++ b/charts/prometheus-blackbox-exporter/values.yaml
@@ -31,6 +31,16 @@ runAsUser: 1000
 readOnlyRootFilesystem: true
 runAsNonRoot: true
 
+livenessProbe:
+  httpGet:
+    path: /health
+    port: http
+
+readinessProbe:
+  httpGet:
+    path: /health
+    port: http
+
 nodeSelector: {}
 tolerations: []
 affinity: {}
```

## 4.0.0 

**Release date:** 2020-05-14

![AppVersion: 0.16.0](https://img.shields.io/static/v1?label=AppVersion&message=0.16.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Add pod security policy support and create service account by default (#21796) 

### Default value changes

```diff
diff --git a/charts/prometheus-blackbox-exporter/values.yaml b/charts/prometheus-blackbox-exporter/values.yaml
index d8f27ad..59c0955 100644
--- a/charts/prometheus-blackbox-exporter/values.yaml
+++ b/charts/prometheus-blackbox-exporter/values.yaml
@@ -5,6 +5,9 @@ kind: Deployment
 podDisruptionBudget: {}
   # maxUnavailable: 0
 
+## Enable pod security policy
+pspEnabled: true
+
 strategy:
   rollingUpdate:
     maxSurge: 1
@@ -77,7 +80,7 @@ service:
 
 serviceAccount:
   # Specifies whether a ServiceAccount should be created
-  create: false
+  create: true
   # The name of the ServiceAccount to use.
   # If not set and create is true, a name is generated using the fullname template
   name:
```

## 3.5.0 

**Release date:** 2020-04-28

![AppVersion: 0.15.1](https://img.shields.io/static/v1?label=AppVersion&message=0.15.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Add Ingress path (#22142) 

### Default value changes

```diff
diff --git a/charts/prometheus-blackbox-exporter/values.yaml b/charts/prometheus-blackbox-exporter/values.yaml
index c0409be..d8f27ad 100644
--- a/charts/prometheus-blackbox-exporter/values.yaml
+++ b/charts/prometheus-blackbox-exporter/values.yaml
@@ -91,6 +91,7 @@ ingress:
   enabled: false
   hosts: []
      # - chart-example.local
+  path: '/'
   annotations: {}
     # kubernetes.io/ingress.class: nginx
     # kubernetes.io/tls-acme: "true"
```

## 3.4.1 

**Release date:** 2020-04-27

![AppVersion: 0.15.1](https://img.shields.io/static/v1?label=AppVersion&message=0.15.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-blackbox-exporter] prometheusrule additionalLabels have too many indents/spaces (#22051) 

### Default value changes

```diff
# No changes in this release
```

## 3.4.0 

**Release date:** 2020-04-07

![AppVersion: 0.15.1](https://img.shields.io/static/v1?label=AppVersion&message=0.15.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-blackbox-exporter] Add PrometheusRule support (#21776) 

### Default value changes

```diff
diff --git a/charts/prometheus-blackbox-exporter/values.yaml b/charts/prometheus-blackbox-exporter/values.yaml
index 4c2851d..c0409be 100644
--- a/charts/prometheus-blackbox-exporter/values.yaml
+++ b/charts/prometheus-blackbox-exporter/values.yaml
@@ -127,3 +127,11 @@ serviceMonitor:
 #      interval: 60s                    # Scraping interval. Overrides value set in `defaults`
 #      scrapeTimeout: 60s               # Scrape timeout. Overrides value set in `defaults`
 #      module: http_2xx                 # Module used for scraping. Overrides value set in `defaults`
+
+## Custom PrometheusRules to be defined
+## ref: https://github.com/coreos/prometheus-operator#customresourcedefinitions
+prometheusRule:
+  enabled: false
+  additionalLabels: {}
+  namespace: ""
+  rules: []
```

## 3.3.0 

**Release date:** 2020-04-02

![AppVersion: 0.15.1](https://img.shields.io/static/v1?label=AppVersion&message=0.15.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-blackbox-exporter] Add serviceAccount support (#21726) 

### Default value changes

```diff
diff --git a/charts/prometheus-blackbox-exporter/values.yaml b/charts/prometheus-blackbox-exporter/values.yaml
index f6e7eac..4c2851d 100644
--- a/charts/prometheus-blackbox-exporter/values.yaml
+++ b/charts/prometheus-blackbox-exporter/values.yaml
@@ -75,6 +75,14 @@ service:
   type: ClusterIP
   port: 9115
 
+serviceAccount:
+  # Specifies whether a ServiceAccount should be created
+  create: false
+  # The name of the ServiceAccount to use.
+  # If not set and create is true, a name is generated using the fullname template
+  name:
+  annotations: {}
+
 ## An Ingress resource can provide name-based virtual hosting and TLS
 ## termination among other things for CouchDB deployments which are accessed
 ## from outside the Kubernetes cluster.
```

## 3.2.0 

**Release date:** 2020-02-11

![AppVersion: 0.15.1](https://img.shields.io/static/v1?label=AppVersion&message=0.15.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Add possibility to set blackbox to daemonset (#19097) 

### Default value changes

```diff
diff --git a/charts/prometheus-blackbox-exporter/values.yaml b/charts/prometheus-blackbox-exporter/values.yaml
index 53e2f33..f6e7eac 100644
--- a/charts/prometheus-blackbox-exporter/values.yaml
+++ b/charts/prometheus-blackbox-exporter/values.yaml
@@ -1,5 +1,7 @@
 restartPolicy: Always
 
+kind: Deployment
+
 podDisruptionBudget: {}
   # maxUnavailable: 0
 
```

## 3.1.0 

**Release date:** 2020-02-03

![AppVersion: 0.15.1](https://img.shields.io/static/v1?label=AppVersion&message=0.15.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-blackbox-exporter] add "allowIcmp" setting (#20409) 

### Default value changes

```diff
diff --git a/charts/prometheus-blackbox-exporter/values.yaml b/charts/prometheus-blackbox-exporter/values.yaml
index e52e351..53e2f33 100644
--- a/charts/prometheus-blackbox-exporter/values.yaml
+++ b/charts/prometheus-blackbox-exporter/values.yaml
@@ -58,6 +58,8 @@ extraSecretMounts: []
   #   readOnly: true
   #   defaultMode: 420
 
+allowIcmp: false
+
 resources: {}
   # limits:
   #   memory: 300Mi
```

## 3.0.1 

**Release date:** 2020-01-30

![AppVersion: 0.15.1](https://img.shields.io/static/v1?label=AppVersion&message=0.15.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-blackbox-exporter] Fix multiple target servicemonitor  (#20296) 

### Default value changes

```diff
# No changes in this release
```

## 3.0.0 

**Release date:** 2020-01-20

![AppVersion: 0.15.1](https://img.shields.io/static/v1?label=AppVersion&message=0.15.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-blackbox-exporter] Allow scraping multiple targets (#19620) 

### Default value changes

```diff
diff --git a/charts/prometheus-blackbox-exporter/values.yaml b/charts/prometheus-blackbox-exporter/values.yaml
index 5469f07..e52e351 100644
--- a/charts/prometheus-blackbox-exporter/values.yaml
+++ b/charts/prometheus-blackbox-exporter/values.yaml
@@ -100,11 +100,18 @@ serviceMonitor:
   ## https://github.com/coreos/prometheus-operator
   ##
   enabled: false
-  labels: {}
-  interval: 30s
-  scrapeTimeout: 30s
-  module: http_2xx
-  # The URL that blackbox will scrape
-  url: http://example.com/healthz
-  # Optional human readable URL that will appear in Prometheus / AlertManager
-  urlHumanReadable:  # www.changemeoriwillfail.com
+
+  # Default values that will be used for all ServiceMonitors created by `targets`
+  defaults:
+    labels: {}
+    interval: 30s
+    scrapeTimeout: 30s
+    module: http_2xx
+
+  targets:
+#    - name: example                    # Human readable URL that will appear in Prometheus / AlertManager
+#      url: http://example.com/healthz  # The URL that blackbox will scrape
+#      labels: {}                       # List of labels for ServiceMonitor. Overrides value set in `defaults`
+#      interval: 60s                    # Scraping interval. Overrides value set in `defaults`
+#      scrapeTimeout: 60s               # Scrape timeout. Overrides value set in `defaults`
+#      module: http_2xx                 # Module used for scraping. Overrides value set in `defaults`
```

## 2.0.0 

**Release date:** 2020-01-08

![AppVersion: 0.15.1](https://img.shields.io/static/v1?label=AppVersion&message=0.15.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-blackbox-exporter] Fix PodDisruptionBudget configuration (#19707) 

### Default value changes

```diff
diff --git a/charts/prometheus-blackbox-exporter/values.yaml b/charts/prometheus-blackbox-exporter/values.yaml
index 71d7a01..5469f07 100644
--- a/charts/prometheus-blackbox-exporter/values.yaml
+++ b/charts/prometheus-blackbox-exporter/values.yaml
@@ -1,8 +1,7 @@
 restartPolicy: Always
 
-podDisruptionBudget:
-  enabled: true
-  maxUnavailable: 0
+podDisruptionBudget: {}
+  # maxUnavailable: 0
 
 strategy:
   rollingUpdate:
```

## 1.6.0 

**Release date:** 2019-12-16

![AppVersion: 0.15.1](https://img.shields.io/static/v1?label=AppVersion&message=0.15.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-blackbox-exporter] Able to mount extra secrets int… (#19119) 

### Default value changes

```diff
diff --git a/charts/prometheus-blackbox-exporter/values.yaml b/charts/prometheus-blackbox-exporter/values.yaml
index 0c44610..71d7a01 100644
--- a/charts/prometheus-blackbox-exporter/values.yaml
+++ b/charts/prometheus-blackbox-exporter/values.yaml
@@ -12,7 +12,7 @@ strategy:
 
 image:
   repository: prom/blackbox-exporter
-  tag: v0.15.1
+  tag: v0.16.0
   pullPolicy: IfNotPresent
 
   ## Optionally specify an array of imagePullSecrets.
@@ -42,6 +42,23 @@ config:
         no_follow_redirects: false
         preferred_ip_protocol: "ip4"
 
+extraConfigmapMounts: []
+  # - name: certs-configmap
+  #   mountPath: /etc/secrets/ssl/
+  #   subPath: certificates.crt # (optional)
+  #   configMap: certs-configmap
+  #   readOnly: true
+  #   defaultMode: 420
+
+## Additional secret mounts
+# Defines additional mounts with secrets. Secrets must be manually created in the namespace.
+extraSecretMounts: []
+  # - name: secret-files
+  #   mountPath: /etc/secrets
+  #   secretName: blackbox-secret-files
+  #   readOnly: true
+  #   defaultMode: 420
+
 resources: {}
   # limits:
   #   memory: 300Mi
```

## 1.5.1 

**Release date:** 2019-11-13

![AppVersion: 0.15.1](https://img.shields.io/static/v1?label=AppVersion&message=0.15.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-blackbox-exporter] Update the Ingress apiversion to networking.k8s.io/v1beta1 (#18645) 

### Default value changes

```diff
# No changes in this release
```

## 1.5.0 

**Release date:** 2019-11-04

![AppVersion: 0.15.1](https://img.shields.io/static/v1?label=AppVersion&message=0.15.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-blackbox-exporter] removes configmap-reload in favor of a configmap checksum (#17687) 

### Default value changes

```diff
diff --git a/charts/prometheus-blackbox-exporter/values.yaml b/charts/prometheus-blackbox-exporter/values.yaml
index 618baa3..0c44610 100644
--- a/charts/prometheus-blackbox-exporter/values.yaml
+++ b/charts/prometheus-blackbox-exporter/values.yaml
@@ -78,30 +78,6 @@ extraArgs: []
 #  --history.limit=1000
 
 replicas: 1
-## Monitors ConfigMap changes and POSTs to a URL
-## Ref: https://github.com/jimmidyson/configmap-reload
-##
-configmapReload:
-  ## configmap-reload container name
-  ##
-  name: configmap-reload
-
-  ## User to run configmap-reload container as
-  ##
-  runAsUser: 65534
-  runAsNonRoot: true
-
-  ## configmap-reload container image
-  ##
-  image:
-    repository: jimmidyson/configmap-reload
-    tag: v0.2.2
-    pullPolicy: IfNotPresent
-
-  ## configmap-reload resource requests and limits
-  ## Ref: http://kubernetes.io/docs/user-guide/compute-resources/
-  ##
-  resources: {}
 
 serviceMonitor:
   ## If true, a ServiceMonitor CRD is created for a prometheus operator
```

## 1.4.0 

**Release date:** 2019-10-23

![AppVersion: 0.15.1](https://img.shields.io/static/v1?label=AppVersion&message=0.15.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-blackbox-exporter] Add option to enable/disable PodDisruptionBudget. (#18235) 

### Default value changes

```diff
diff --git a/charts/prometheus-blackbox-exporter/values.yaml b/charts/prometheus-blackbox-exporter/values.yaml
index 9777a5e..618baa3 100644
--- a/charts/prometheus-blackbox-exporter/values.yaml
+++ b/charts/prometheus-blackbox-exporter/values.yaml
@@ -1,6 +1,7 @@
 restartPolicy: Always
 
 podDisruptionBudget:
+  enabled: true
   maxUnavailable: 0
 
 strategy:
```

## 1.3.1 

**Release date:** 2019-10-02

![AppVersion: 0.15.1](https://img.shields.io/static/v1?label=AppVersion&message=0.15.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-blackbox-exporter] updated docker image to 0.15.1 (#17560) 

### Default value changes

```diff
diff --git a/charts/prometheus-blackbox-exporter/values.yaml b/charts/prometheus-blackbox-exporter/values.yaml
index 62305b7..9777a5e 100644
--- a/charts/prometheus-blackbox-exporter/values.yaml
+++ b/charts/prometheus-blackbox-exporter/values.yaml
@@ -11,7 +11,7 @@ strategy:
 
 image:
   repository: prom/blackbox-exporter
-  tag: v0.15.0
+  tag: v0.15.1
   pullPolicy: IfNotPresent
 
   ## Optionally specify an array of imagePullSecrets.
```

## 1.3.0 

**Release date:** 2019-09-12

![AppVersion: 0.15.0](https://img.shields.io/static/v1?label=AppVersion&message=0.15.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/blackbox-exporter]: upgrade (#17083) 

### Default value changes

```diff
diff --git a/charts/prometheus-blackbox-exporter/values.yaml b/charts/prometheus-blackbox-exporter/values.yaml
index 8dfb9d0..62305b7 100644
--- a/charts/prometheus-blackbox-exporter/values.yaml
+++ b/charts/prometheus-blackbox-exporter/values.yaml
@@ -11,7 +11,7 @@ strategy:
 
 image:
   repository: prom/blackbox-exporter
-  tag: v0.14.0
+  tag: v0.15.0
   pullPolicy: IfNotPresent
 
   ## Optionally specify an array of imagePullSecrets.
```

## 1.2.0 

**Release date:** 2019-09-05

![AppVersion: 0.14.0](https://img.shields.io/static/v1?label=AppVersion&message=0.14.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometehus-blackbox-exporter] Improve configurability. (#16815) 

### Default value changes

```diff
diff --git a/charts/prometheus-blackbox-exporter/values.yaml b/charts/prometheus-blackbox-exporter/values.yaml
index 2558b75..8dfb9d0 100644
--- a/charts/prometheus-blackbox-exporter/values.yaml
+++ b/charts/prometheus-blackbox-exporter/values.yaml
@@ -1,5 +1,14 @@
 restartPolicy: Always
 
+podDisruptionBudget:
+  maxUnavailable: 0
+
+strategy:
+  rollingUpdate:
+    maxSurge: 1
+    maxUnavailable: 0
+  type: RollingUpdate
+
 image:
   repository: prom/blackbox-exporter
   tag: v0.14.0
@@ -38,6 +47,8 @@ resources: {}
   # requests:
   #   memory: 50Mi
 
+priorityClassName: ""
+
 service:
   annotations: {}
   type: ClusterIP
```

## 1.1.0 

**Release date:** 2019-08-22

![AppVersion: 0.14.0](https://img.shields.io/static/v1?label=AppVersion&message=0.14.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Allow custom module for blackbox service monitor (#15987) 

### Default value changes

```diff
diff --git a/charts/prometheus-blackbox-exporter/values.yaml b/charts/prometheus-blackbox-exporter/values.yaml
index 81e4e65..2558b75 100644
--- a/charts/prometheus-blackbox-exporter/values.yaml
+++ b/charts/prometheus-blackbox-exporter/values.yaml
@@ -99,6 +99,7 @@ serviceMonitor:
   labels: {}
   interval: 30s
   scrapeTimeout: 30s
+  module: http_2xx
   # The URL that blackbox will scrape
   url: http://example.com/healthz
   # Optional human readable URL that will appear in Prometheus / AlertManager
```

## 1.0.1 

**Release date:** 2019-08-05

![AppVersion: 0.14.0](https://img.shields.io/static/v1?label=AppVersion&message=0.14.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-blackbox-exporter] Fixes issue where containers in deployment were running as root (#15154) 

### Default value changes

```diff
diff --git a/charts/prometheus-blackbox-exporter/values.yaml b/charts/prometheus-blackbox-exporter/values.yaml
index d3b96be..81e4e65 100644
--- a/charts/prometheus-blackbox-exporter/values.yaml
+++ b/charts/prometheus-blackbox-exporter/values.yaml
@@ -12,6 +12,11 @@ image:
   # pullSecrets:
   #   - myRegistrKeySecretName
 
+## User to run blackbox-exporter container as
+runAsUser: 1000
+readOnlyRootFilesystem: true
+runAsNonRoot: true
+
 nodeSelector: {}
 tolerations: []
 affinity: {}
@@ -69,6 +74,11 @@ configmapReload:
   ##
   name: configmap-reload
 
+  ## User to run configmap-reload container as
+  ##
+  runAsUser: 65534
+  runAsNonRoot: true
+
   ## configmap-reload container image
   ##
   image:
```

## 1.0.0 

**Release date:** 2019-07-24

![AppVersion: 0.14.0](https://img.shields.io/static/v1?label=AppVersion&message=0.14.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-blackbox-exporter] Breaking change: general upgrade and ServiceMonitor (#11426) 

### Default value changes

```diff
diff --git a/charts/prometheus-blackbox-exporter/values.yaml b/charts/prometheus-blackbox-exporter/values.yaml
index ad96e18..d3b96be 100644
--- a/charts/prometheus-blackbox-exporter/values.yaml
+++ b/charts/prometheus-blackbox-exporter/values.yaml
@@ -80,3 +80,16 @@ configmapReload:
   ## Ref: http://kubernetes.io/docs/user-guide/compute-resources/
   ##
   resources: {}
+
+serviceMonitor:
+  ## If true, a ServiceMonitor CRD is created for a prometheus operator
+  ## https://github.com/coreos/prometheus-operator
+  ##
+  enabled: false
+  labels: {}
+  interval: 30s
+  scrapeTimeout: 30s
+  # The URL that blackbox will scrape
+  url: http://example.com/healthz
+  # Optional human readable URL that will appear in Prometheus / AlertManager
+  urlHumanReadable:  # www.changemeoriwillfail.com
```

## 0.4.0 

**Release date:** 2019-06-27

![AppVersion: 0.14.0](https://img.shields.io/static/v1?label=AppVersion&message=0.14.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-blackbox-exporter] Add support for imagePullSecrets (#14903) 

### Default value changes

```diff
diff --git a/charts/prometheus-blackbox-exporter/values.yaml b/charts/prometheus-blackbox-exporter/values.yaml
index 46041d0..ad96e18 100644
--- a/charts/prometheus-blackbox-exporter/values.yaml
+++ b/charts/prometheus-blackbox-exporter/values.yaml
@@ -5,6 +5,13 @@ image:
   tag: v0.14.0
   pullPolicy: IfNotPresent
 
+  ## Optionally specify an array of imagePullSecrets.
+  ## Secrets must be manually created in the namespace.
+  ## ref: https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/
+  ##
+  # pullSecrets:
+  #   - myRegistrKeySecretName
+
 nodeSelector: {}
 tolerations: []
 affinity: {}
```

## 0.3.0 

**Release date:** 2019-04-29

![AppVersion: 0.14.0](https://img.shields.io/static/v1?label=AppVersion&message=0.14.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-blackbox-exporter] add option to store config as a Secret (#12731) 

### Default value changes

```diff
diff --git a/charts/prometheus-blackbox-exporter/values.yaml b/charts/prometheus-blackbox-exporter/values.yaml
index 3b942ad..46041d0 100644
--- a/charts/prometheus-blackbox-exporter/values.yaml
+++ b/charts/prometheus-blackbox-exporter/values.yaml
@@ -9,6 +9,7 @@ nodeSelector: {}
 tolerations: []
 affinity: {}
 
+secretConfig: false
 config:
   modules:
     http_2xx:
```

## 0.2.1 

**Release date:** 2019-04-19

![AppVersion: 0.14.0](https://img.shields.io/static/v1?label=AppVersion&message=0.14.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Update blackbox exporter to v0.14.0 (#11933) 

### Default value changes

```diff
diff --git a/charts/prometheus-blackbox-exporter/values.yaml b/charts/prometheus-blackbox-exporter/values.yaml
index 6b56a83..3b942ad 100644
--- a/charts/prometheus-blackbox-exporter/values.yaml
+++ b/charts/prometheus-blackbox-exporter/values.yaml
@@ -2,7 +2,7 @@ restartPolicy: Always
 
 image:
   repository: prom/blackbox-exporter
-  tag: v0.12.0
+  tag: v0.14.0
   pullPolicy: IfNotPresent
 
 nodeSelector: {}
```

## 0.2.0 

**Release date:** 2018-11-05

![AppVersion: 0.12.0](https://img.shields.io/static/v1?label=AppVersion&message=0.12.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Blackbox-exporter: Add support for affinity and tolerations (#8865) 

### Default value changes

```diff
diff --git a/charts/prometheus-blackbox-exporter/values.yaml b/charts/prometheus-blackbox-exporter/values.yaml
index 0a6d91e..6b56a83 100644
--- a/charts/prometheus-blackbox-exporter/values.yaml
+++ b/charts/prometheus-blackbox-exporter/values.yaml
@@ -6,6 +6,8 @@ image:
   pullPolicy: IfNotPresent
 
 nodeSelector: {}
+tolerations: []
+affinity: {}
 
 config:
   modules:
```

## 0.1.3 

**Release date:** 2018-10-24

![AppVersion: 0.12.0](https://img.shields.io/static/v1?label=AppVersion&message=0.12.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-blackbox-exporter] Fixing minor syntax error (#8674) 

### Default value changes

```diff
diff --git a/charts/prometheus-blackbox-exporter/values.yaml b/charts/prometheus-blackbox-exporter/values.yaml
index c8859ab..0a6d91e 100644
--- a/charts/prometheus-blackbox-exporter/values.yaml
+++ b/charts/prometheus-blackbox-exporter/values.yaml
@@ -17,7 +17,7 @@ config:
         no_follow_redirects: false
         preferred_ip_protocol: "ip4"
 
-resources:
+resources: {}
   # limits:
   #   memory: 300Mi
   # requests:
```

## 0.1.2 

**Release date:** 2018-10-24

![AppVersion: 0.12.0](https://img.shields.io/static/v1?label=AppVersion&message=0.12.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-blackbox-exporter] Allowing custom service annotations and labels (#8673) 

### Default value changes

```diff
diff --git a/charts/prometheus-blackbox-exporter/values.yaml b/charts/prometheus-blackbox-exporter/values.yaml
index 55d7de9..c8859ab 100644
--- a/charts/prometheus-blackbox-exporter/values.yaml
+++ b/charts/prometheus-blackbox-exporter/values.yaml
@@ -24,6 +24,7 @@ resources:
   #   memory: 50Mi
 
 service:
+  annotations: {}
   type: ClusterIP
   port: 9115
 
```

## 0.1.1 

**Release date:** 2018-07-15

![AppVersion: 0.12.0](https://img.shields.io/static/v1?label=AppVersion&message=0.12.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Review readme (#6399) 

### Default value changes

```diff
diff --git a/charts/prometheus-blackbox-exporter/values.yaml b/charts/prometheus-blackbox-exporter/values.yaml
index a26fdc1..55d7de9 100644
--- a/charts/prometheus-blackbox-exporter/values.yaml
+++ b/charts/prometheus-blackbox-exporter/values.yaml
@@ -7,14 +7,15 @@ image:
 
 nodeSelector: {}
 
-config: {}
-  # http_2xx:
-  #   prober: http
-  #   timeout: 5s
-  #   http:
-  #     valid_http_versions: ["HTTP/1.1", "HTTP/2"]
-  #     no_follow_redirects: false
-  #     preferred_ip_protocol: "ip4" # defaults to "ip6"
+config:
+  modules:
+    http_2xx:
+      prober: http
+      timeout: 5s
+      http:
+        valid_http_versions: ["HTTP/1.1", "HTTP/2"]
+        no_follow_redirects: false
+        preferred_ip_protocol: "ip4"
 
 resources:
   # limits:
```

## 0.1.0 

**Release date:** 2018-06-28

![AppVersion: 0.12.0](https://img.shields.io/static/v1?label=AppVersion&message=0.12.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Create helm chart for blackbox exporter (#6134) 

### Default value changes

```diff
restartPolicy: Always

image:
  repository: prom/blackbox-exporter
  tag: v0.12.0
  pullPolicy: IfNotPresent

nodeSelector: {}

config: {}
  # http_2xx:
  #   prober: http
  #   timeout: 5s
  #   http:
  #     valid_http_versions: ["HTTP/1.1", "HTTP/2"]
  #     no_follow_redirects: false
  #     preferred_ip_protocol: "ip4" # defaults to "ip6"

resources:
  # limits:
  #   memory: 300Mi
  # requests:
  #   memory: 50Mi

service:
  type: ClusterIP
  port: 9115

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
