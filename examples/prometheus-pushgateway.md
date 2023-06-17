# Change Log

## 2.2.0

**Release date:** 2023-05-25

![AppVersion: v1.5.1](https://img.shields.io/static/v1?label=AppVersion&message=v1.5.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-pushgateway] Make path configurable in service monitor (#3373)

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index 0df94441..011dbf29 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -209,6 +209,11 @@ serviceMonitor:
   enabled: false
   namespace: monitoring
 
+  # telemetryPath: HTTP resource path from which to fetch metrics.
+  # Telemetry path, default /metrics, has to be prefixed accordingly if pushgateway sets a route prefix at start-up.
+  #
+  telemetryPath: "/metrics"
+
   # Fallback to the prometheus default unless specified
   # interval: 10s
 

```

## 2.1.6

**Release date:** 2023-05-08

![AppVersion: v1.5.1](https://img.shields.io/static/v1?label=AppVersion&message=v1.5.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-pushgateway] fix apiVersion for PDB (#3300)

### Default value changes

```diff
# No changes in this release
```

## 2.1.5

**Release date:** 2023-05-08

![AppVersion: v1.5.1](https://img.shields.io/static/v1?label=AppVersion&message=v1.5.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-pushgateway] Remove label duplication (#3266)

### Default value changes

```diff
# No changes in this release
```

## 2.1.4

**Release date:** 2023-05-08

![AppVersion: v1.5.1](https://img.shields.io/static/v1?label=AppVersion&message=v1.5.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-pushgateway] Add a maintainer (#3350)

### Default value changes

```diff
# No changes in this release
```

## 2.1.3

**Release date:** 2023-02-14

![AppVersion: v1.5.1](https://img.shields.io/static/v1?label=AppVersion&message=v1.5.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-pushgateway] Correct spelling mistake (#3003)

### Default value changes

```diff
# No changes in this release
```

## 2.1.2

**Release date:** 2023-02-10

![AppVersion: v1.5.1](https://img.shields.io/static/v1?label=AppVersion&message=v1.5.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-pushgateway] Adjust Liveness probe default to /healthy (#3017)

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index 23d0a0ea..0df94441 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -103,7 +103,7 @@ liveness:
   enabled: true
   probe:
     httpGet:
-      path: /-/ready
+      path: /-/healthy
       port: 9091
     initialDelaySeconds: 10
     timeoutSeconds: 10

```

## 2.1.1

**Release date:** 2023-02-03

![AppVersion: v1.5.1](https://img.shields.io/static/v1?label=AppVersion&message=v1.5.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-pushgateway] helm template error (#2985)

### Default value changes

```diff
# No changes in this release
```

## 2.1.0

**Release date:** 2023-02-02

![AppVersion: v1.5.1](https://img.shields.io/static/v1?label=AppVersion&message=v1.5.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-pushgateway] Add namespaceOverride to all resources (#2967)

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index f2fd5cd1..23d0a0ea 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -8,6 +8,9 @@ nameOverride: ""
 # Provide a name to substitute for the full names of resources
 fullnameOverride: ""
 
+# Provide a namespace to substitude for the namespace on resources
+namespaceOverride: ""
+
 image:
   repository: prom/pushgateway
   # if not set appVersion field from Chart.yaml is used

```

## 2.0.4

**Release date:** 2023-01-30

![AppVersion: v1.5.1](https://img.shields.io/static/v1?label=AppVersion&message=v1.5.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-pushgateway] Fix NodePort service type (#2958)

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index 90cee5c8..f2fd5cd1 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -21,6 +21,7 @@ service:
   type: ClusterIP
   port: 9091
   targetPort: 9091
+  # nodePort: 32100
 
   # Optional - Can be used for headless if value is "None"
   clusterIP: ""

```

## 2.0.3

**Release date:** 2022-12-21

![AppVersion: v1.5.1](https://img.shields.io/static/v1?label=AppVersion&message=v1.5.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-pushgateway] Fix additionalLabels for service monitor (#2820)

### Default value changes

```diff
# No changes in this release
```

## 2.0.2

**Release date:** 2022-12-05

![AppVersion: v1.5.1](https://img.shields.io/static/v1?label=AppVersion&message=v1.5.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-pushgateway] Fix service monitor selector (#2782) (#2783)

### Default value changes

```diff
# No changes in this release
```

## 2.0.1

**Release date:** 2022-12-03

![AppVersion: v1.5.1](https://img.shields.io/static/v1?label=AppVersion&message=v1.5.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-pushgateway] Hotfix new line problem (#2779)

### Default value changes

```diff
# No changes in this release
```

## 2.0.0

**Release date:** 2022-12-02

![AppVersion: v1.5.1](https://img.shields.io/static/v1?label=AppVersion&message=v1.5.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-pushgateway] Refactor (#2746)

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index e2e9ebed..90cee5c8 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -10,7 +10,8 @@ fullnameOverride: ""
 
 image:
   repository: prom/pushgateway
-  tag: v1.4.2
+  # if not set appVersion field from Chart.yaml is used
+  tag: ""
   pullPolicy: IfNotPresent
 
 # Optional pod imagePullSecrets

```

## 1.21.1

**Release date:** 2022-11-30

![AppVersion: 1.4.2](https://img.shields.io/static/v1?label=AppVersion&message=1.4.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-pushgateway] Ensure PDB labels are same as the matchLabels of the owner app (#2666)

### Default value changes

```diff
# No changes in this release
```

## 1.21.0

**Release date:** 2022-11-20

![AppVersion: 1.4.2](https://img.shields.io/static/v1?label=AppVersion&message=1.4.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-pushgateway] Support ingress.extraPaths (#2709)

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index aed41a8f..e2e9ebed 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -132,25 +132,32 @@ ingress:
   path: /
   pathType: ImplementationSpecific
 
-    ## Annotations.
-    ##
-    # annotations:
-    #   kubernetes.io/ingress.class: nginx
-    #   kubernetes.io/tls-acme: 'true'
+  ## Extra paths to prepend to every host configuration. This is useful when working with annotation based services.
+  extraPaths: []
+  # - path: /*
+  #   backend:
+  #     serviceName: ssl-redirect
+  #     servicePort: use-annotation
 
-    ## Hostnames.
-    ## Must be provided if Ingress is enabled.
-    ##
-    # hosts:
-    #   - pushgateway.domain.com
+  ## Annotations.
+  ##
+  # annotations:
+  #   kubernetes.io/ingress.class: nginx
+  #   kubernetes.io/tls-acme: 'true'
 
-    ## TLS configuration.
-    ## Secrets must be manually created in the namespace.
-    ##
-    # tls:
-    #   - secretName: pushgateway-tls
-    #     hosts:
-    #       - pushgateway.domain.com
+  ## Hostnames.
+  ## Must be provided if Ingress is enabled.
+  ##
+  # hosts:
+  #   - pushgateway.domain.com
+
+  ## TLS configuration.
+  ## Secrets must be manually created in the namespace.
+  ##
+  # tls:
+  #   - secretName: pushgateway-tls
+  #     hosts:
+  #       - pushgateway.domain.com
 
 tolerations: {}
   # - effect: NoSchedule

```

## 1.20.1

**Release date:** 2022-10-19

![AppVersion: 1.4.2](https://img.shields.io/static/v1?label=AppVersion&message=1.4.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Omit clusterIP if service is of type LoadBalancer (#2579)

### Default value changes

```diff
# No changes in this release
```

## 1.20.0

**Release date:** 2022-09-30

![AppVersion: 1.4.2](https://img.shields.io/static/v1?label=AppVersion&message=1.4.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-pushgateway] - Add extra init containers for pushgateway helm chart (#2480)

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index cb81baa2..aed41a8f 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -56,6 +56,10 @@ extraVars: []
 ##  - --persistence.interval=5m
 extraArgs: []
 
+## Additional InitContainers to initialize the pod
+##
+extraInitContainers: []
+
 # Optional additional containers (sidecar)
 extraContainers: []
   # - name: oAuth2-proxy

```

## 1.19.2

**Release date:** 2022-09-30

![AppVersion: 1.4.2](https://img.shields.io/static/v1?label=AppVersion&message=1.4.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-pushgateway] extraVolumes and extraVolumeMounts should be empty arrays (#2495)

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index aefa2bbc..cb81baa2 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -286,10 +286,10 @@ persistentVolume:
   ##
   subPath: ""
 
-extraVolumes: {}
+extraVolumes: []
   # - name: extra
   #   emptyDir: {}
-extraVolumeMounts: {}
+extraVolumeMounts: []
   # - name: extra
   #   mountPath: /usr/share/extras
   #   readOnly: true

```

## 1.19.1

**Release date:** 2022-09-28

![AppVersion: 1.4.2](https://img.shields.io/static/v1?label=AppVersion&message=1.4.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Fix topologySpreadConstraints (#2501)

### Default value changes

```diff
# No changes in this release
```

## 1.19.0

**Release date:** 2022-09-27

![AppVersion: 1.4.2](https://img.shields.io/static/v1?label=AppVersion&message=1.4.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-pushgateway] Add topologySpreadConstraints (#2497)

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index 5bb8eb2b..aefa2bbc 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -184,6 +184,10 @@ containerSecurityContext: {}
 ## Ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#affinity-and-anti-affinity
 affinity: {}
 
+## Topology spread constraints for pods
+## Ref: https://kubernetes.io/docs/concepts/workloads/pods/pod-topology-spread-constraints/
+topologySpreadConstraints: []
+
 # Enable this if you're using https://github.com/coreos/prometheus-operator
 serviceMonitor:
   enabled: false

```

## 1.18.3

**Release date:** 2022-09-19

![AppVersion: 1.4.2](https://img.shields.io/static/v1?label=AppVersion&message=1.4.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-pushgateway] update notes.txt notify port (#2462)

### Default value changes

```diff
# No changes in this release
```

## 1.18.2

**Release date:** 2022-06-10

![AppVersion: 1.4.2](https://img.shields.io/static/v1?label=AppVersion&message=1.4.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Fix: Metadata labels consistency (#2126)

### Default value changes

```diff
# No changes in this release
```

## 1.18.1

**Release date:** 2022-05-30

![AppVersion: 1.4.2](https://img.shields.io/static/v1?label=AppVersion&message=1.4.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-pushgateway] - Upgrade pdb version to apiVersion: policy/v1 (#2096)

### Default value changes

```diff
# No changes in this release
```

## 1.18.0

**Release date:** 2022-05-11

![AppVersion: 1.4.2](https://img.shields.io/static/v1?label=AppVersion&message=1.4.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-pushgateway] Allow prometheus gateway to run as a StatefulSet (#2031)

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index 8c5992e5..5bb8eb2b 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -159,6 +159,12 @@ nodeSelector: {}
 
 replicaCount: 1
 
+## When running more than one replica alongside with persistence, different volumes are needed
+## per replica, since sharing a `persistence.file` across replicas does not keep metrics synced.
+## For this purpose, you can enable the `runAsStatefulSet` to deploy the pushgateway as a
+## StatefulSet instead of as a Deployment.
+runAsStatefulSet: false
+
 ## Security context to be added to push-gateway pods
 ##
 securityContext:

```

## 1.17.0

**Release date:** 2022-05-03

![AppVersion: 1.4.2](https://img.shields.io/static/v1?label=AppVersion&message=1.4.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-pushgateway] Adding extra containers (#2023)

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index ecc536b6..8c5992e5 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -56,6 +56,28 @@ extraVars: []
 ##  - --persistence.interval=5m
 extraArgs: []
 
+# Optional additional containers (sidecar)
+extraContainers: []
+  # - name: oAuth2-proxy
+  #   args:
+  #     - -https-address=:9092
+  #     - -upstream=http://localhost:9091
+  #     - -skip-auth-regex=^/metrics
+  #     - -openshift-delegate-urls={"/":{"group":"monitoring.coreos.com","resource":"prometheuses","verb":"get"}}
+  #   image: openshift/oauth-proxy:v1.1.0
+  #   ports:
+  #       - containerPort: 9092
+  #         name: proxy
+  #   resources:
+  #       limits:
+  #         memory: 16Mi
+  #       requests:
+  #         memory: 4Mi
+  #         cpu: 20m
+  #   volumeMounts:
+  #     - mountPath: /etc/prometheus/secrets/pushgateway-tls
+  #       name: secret-pushgateway-tls
+
 resources: {}
   # We usually recommend not to specify default resources and to leave this as a conscious
   # choice for the user. This also increases chances charts run on environments with little
@@ -164,6 +186,14 @@ serviceMonitor:
   # Fallback to the prometheus default unless specified
   # interval: 10s
 
+  ## scheme: HTTP scheme to use for scraping. Can be used with `tlsConfig` for example if using istio mTLS.
+  # scheme: ""
+
+  ## tlsConfig: TLS configuration to use when scraping the endpoint. For example if using istio mTLS.
+  ## Of type: https://github.com/coreos/prometheus-operator/blob/master/Documentation/api.md#tlsconfig
+  # tlsConfig: {}
+
+  # bearerTokenFile:
   # Fallback to the prometheus default unless specified
   # scrapeTimeout: 30s
 

```

## 1.16.1

**Release date:** 2022-03-01

![AppVersion: 1.4.2](https://img.shields.io/static/v1?label=AppVersion&message=1.4.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-pushgateway] Fix indentation in resources example comment (#1837)

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index 33635c92..ecc536b6 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -63,7 +63,7 @@ resources: {}
   # lines, adjust them as necessary, and remove the curly braces after 'resources:'.
   # limits:
   #   cpu: 200m
-  #    memory: 50Mi
+  #   memory: 50Mi
   # requests:
   #   cpu: 100m
   #   memory: 30Mi

```

## 1.16.0

**Release date:** 2022-02-24

![AppVersion: 1.4.2](https://img.shields.io/static/v1?label=AppVersion&message=1.4.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-pushgateway] give the pushgateway live/ready probes the ability to add httpHeaders (#1797)

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index 5d691766..33635c92 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -68,6 +68,24 @@ resources: {}
   #   cpu: 100m
   #   memory: 30Mi
 
+liveness:
+  enabled: true
+  probe:
+    httpGet:
+      path: /-/ready
+      port: 9091
+    initialDelaySeconds: 10
+    timeoutSeconds: 10
+
+readiness:
+  enabled: true
+  probe:
+    httpGet:
+      path: /-/ready
+      port: 9091
+    initialDelaySeconds: 10
+    timeoutSeconds: 10
+
 serviceAccount:
   # Specifies whether a ServiceAccount should be created
   create: true

```

## 1.15.0

**Release date:** 2022-01-28

![AppVersion: 1.4.2](https://img.shields.io/static/v1?label=AppVersion&message=1.4.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-pushgateway] Add support for containerSecurityContext to fix #1741 (#1744)

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index 9b730d95..5d691766 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -126,6 +126,14 @@ securityContext:
   runAsUser: 65534
   runAsNonRoot: true
 
+## Security context to be added to push-gateway containers
+## Having a separate variable as securityContext differs for pods and containers.
+containerSecurityContext: {}
+#  allowPrivilegeEscalation: false
+#  readOnlyRootFilesystem: true
+#  runAsUser: 65534
+#  runAsNonRoot: true
+
 ## Affinity for pod assignment
 ## Ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#affinity-and-anti-affinity
 affinity: {}

```

## 1.14.0

**Release date:** 2021-12-22

![AppVersion: 1.4.2](https://img.shields.io/static/v1?label=AppVersion&message=1.4.2&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-pushgateway] bump prom pushgateway to latest version - 1.4.2 (#1589)

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index fbb49040..9b730d95 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -10,7 +10,7 @@ fullnameOverride: ""
 
 image:
   repository: prom/pushgateway
-  tag: v1.4.1
+  tag: v1.4.2
   pullPolicy: IfNotPresent
 
 # Optional pod imagePullSecrets

```

## 1.13.0

**Release date:** 2021-11-02

![AppVersion: 1.4.1](https://img.shields.io/static/v1?label=AppVersion&message=1.4.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-pushgateway] Support networking.k8s.io/v1 (#1443)

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index 9340c15c..fbb49040 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -84,7 +84,9 @@ ingress:
   ##
   enabled: false
   # AWS ALB requires path of /*
+  className: ""
   path: /
+  pathType: ImplementationSpecific
 
     ## Annotations.
     ##

```

## 1.12.0

**Release date:** 2021-10-19

![AppVersion: 1.4.1](https://img.shields.io/static/v1?label=AppVersion&message=1.4.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-pushgateway] increase version number (#1445)
* Update pushgateway latest (#1306)

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index d9ccdf5b..9340c15c 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -10,7 +10,7 @@ fullnameOverride: ""
 
 image:
   repository: prom/pushgateway
-  tag: v1.3.0
+  tag: v1.4.1
   pullPolicy: IfNotPresent
 
 # Optional pod imagePullSecrets

```

## 1.11.0

**Release date:** 2021-09-14

![AppVersion: 1.3.0](https://img.shields.io/static/v1?label=AppVersion&message=1.3.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-pushgateway] add extraVolumes and extraVolumeMounts to chart (#1059)

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index c3f78ea1..d9ccdf5b 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -218,6 +218,14 @@ persistentVolume:
   ##
   subPath: ""
 
+extraVolumes: {}
+  # - name: extra
+  #   emptyDir: {}
+extraVolumeMounts: {}
+  # - name: extra
+  #   mountPath: /usr/share/extras
+  #   readOnly: true
+
 # Configuration for clusters with restrictive network policies in place:
 # - allowAll allows access to the PushGateway from any namespace
 # - customSelector is a list of pod/namespaceSelectors to allow access from

```

## 1.10.1

**Release date:** 2021-06-29

![AppVersion: 1.3.0](https://img.shields.io/static/v1?label=AppVersion&message=1.3.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Only apply clusterIP if we have a value (#954)

### Default value changes

```diff
# No changes in this release
```

## 1.10.0

**Release date:** 2021-06-17

![AppVersion: 1.3.0](https://img.shields.io/static/v1?label=AppVersion&message=1.3.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-pushgateway] PriorityClassName for prom-pushgateway (#948)

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index ae22dba7..c3f78ea1 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -168,6 +168,8 @@ serviceMonitor:
 # If not set then a PodDisruptionBudget will not be created
 podDisruptionBudget: {}
 
+priorityClassName:
+
 # Deployment Strategy type
 strategy:
   type: Recreate

```

## 1.9.0

**Release date:** 2021-04-26

![AppVersion: 1.3.0](https://img.shields.io/static/v1?label=AppVersion&message=1.3.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-pushgateway] added ClusterIP flag in values.yaml to support headless service (#324)

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index cd586269..ae22dba7 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -21,8 +21,12 @@ service:
   port: 9091
   targetPort: 9091
 
+  # Optional - Can be used for headless if value is "None"
+  clusterIP: ""
+
   loadBalancerIP: ""
   loadBalancerSourceRanges: []
+
 # Optional pod annotations
 podAnnotations: {}
 

```

## 1.8.0

**Release date:** 2021-04-12

![AppVersion: 1.3.0](https://img.shields.io/static/v1?label=AppVersion&message=1.3.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-pushgateway] add loadBalancerIP and sourceranges to pushgateway service (#671)

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index e05ee201..cd586269 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -21,6 +21,8 @@ service:
   port: 9091
   targetPort: 9091
 
+  loadBalancerIP: ""
+  loadBalancerSourceRanges: []
 # Optional pod annotations
 podAnnotations: {}
 

```

## 1.7.1

**Release date:** 2021-03-05

![AppVersion: 1.3.0](https://img.shields.io/static/v1?label=AppVersion&message=1.3.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-pushgateway] Align resource label handling (#456)

### Default value changes

```diff
# No changes in this release
```

## 1.7.0

**Release date:** 2021-02-05

![AppVersion: 1.3.0](https://img.shields.io/static/v1?label=AppVersion&message=1.3.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-pushgateway] fix failed release (#637)
* [prometheus-pushgateway] Adds ability to specify metricRelabelings and relabelings (#453)

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index 12e9e3d0..e05ee201 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -141,6 +141,23 @@ serviceMonitor:
   # [Scraping Pushgateway](https://github.com/prometheus/pushgateway#configure-the-pushgateway-as-a-target-to-scrape)
   honorLabels: true
 
+  ## Metric relabel configs to apply to samples before ingestion.
+  ## [Metric Relabeling](https://prometheus.io/docs/prometheus/latest/configuration/configuration/#metric_relabel_configs)
+  metricRelabelings: []
+  # - action: keep
+  #   regex: 'kube_(daemonset|deployment|pod|namespace|node|statefulset).+'
+  #   sourceLabels: [__name__]
+
+  ## Relabel configs to apply to samples before ingestion.
+  ## [Relabeling](https://prometheus.io/docs/prometheus/latest/configuration/configuration/#relabel_config)
+  relabelings: []
+  # - sourceLabels: [__meta_kubernetes_pod_node_name]
+  #   separator: ;
+  #   regex: ^(.*)$
+  #   targetLabel: nodename
+  #   replacement: $1
+  #   action: replace
+
 # The values to set in the PodDisruptionBudget spec (minAvailable/maxUnavailable)
 # If not set then a PodDisruptionBudget will not be created
 podDisruptionBudget: {}

```

## 1.6.0

**Release date:** 2021-02-04

![AppVersion: 1.3.0](https://img.shields.io/static/v1?label=AppVersion&message=1.3.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-pushgateway] Adds imagePullSecrets value for deployment (#608)

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index 4a56f4f4..12e9e3d0 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -13,6 +13,9 @@ image:
   tag: v1.3.0
   pullPolicy: IfNotPresent
 
+# Optional pod imagePullSecrets
+imagePullSecrets: []
+
 service:
   type: ClusterIP
   port: 9091

```

## 1.5.1

**Release date:** 2020-12-18

![AppVersion: 1.3.0](https://img.shields.io/static/v1?label=AppVersion&message=1.3.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-pushgateway] Add missing fields within values.yaml (#219)

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index 169a5cd9..4a56f4f4 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -1,6 +1,13 @@
 # Default values for prometheus-pushgateway.
 # This is a YAML-formatted file.
 # Declare variables to be passed into your templates.
+
+# Provide a name in place of prometheus-pushgateway for `app:` labels
+nameOverride: ""
+
+# Provide a name to substitute for the full names of resources
+fullnameOverride: ""
+
 image:
   repository: prom/pushgateway
   tag: v1.3.0

```

## 1.5.0

**Release date:** 2020-11-04

![AppVersion: 1.3.0](https://img.shields.io/static/v1?label=AppVersion&message=1.3.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* pushgateway 1.3.0 - update chart version, pushgateway version and values (#195)

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index 63b2a43e..169a5cd9 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -3,7 +3,7 @@
 # Declare variables to be passed into your templates.
 image:
   repository: prom/pushgateway
-  tag: v1.2.0
+  tag: v1.3.0
   pullPolicy: IfNotPresent
 
 service:

```

## 1.4.2

**Release date:** 2020-08-20

![AppVersion: 1.2.0](https://img.shields.io/static/v1?label=AppVersion&message=1.2.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Prep initial charts indexing (#14)

### Default value changes

```diff
# No changes in this release
```

## 1.4.1

**Release date:** 2020-07-07

![AppVersion: 1.2.0](https://img.shields.io/static/v1?label=AppVersion&message=1.2.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-pushgateway] Fix networkPolicy podSelector and ingress rule (#23053)

### Default value changes

```diff
# No changes in this release
```

## 1.4.0

**Release date:** 2020-04-06

![AppVersion: 1.2.0](https://img.shields.io/static/v1?label=AppVersion&message=1.2.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-pushgateway] configureable deployment strategy (#21685)

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index 1ae09068..63b2a43e 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -135,6 +135,10 @@ serviceMonitor:
 # If not set then a PodDisruptionBudget will not be created
 podDisruptionBudget: {}
 
+# Deployment Strategy type
+strategy:
+  type: Recreate
+
 persistentVolume:
   ## If true, pushgateway will create/use a Persistent Volume Claim
   ## If false, use emptyDir

```

## 1.3.0

**Release date:** 2020-03-12

![AppVersion: 1.2.0](https://img.shields.io/static/v1?label=AppVersion&message=1.2.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-pushgateway] upgrade to latest release (#21424)

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index 7709c9d6..1ae09068 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -3,7 +3,7 @@
 # Declare variables to be passed into your templates.
 image:
   repository: prom/pushgateway
-  tag: v1.0.1
+  tag: v1.2.0
   pullPolicy: IfNotPresent
 
 service:

```

## 1.2.15

**Release date:** 2020-03-04

![AppVersion: 1.0.1](https://img.shields.io/static/v1?label=AppVersion&message=1.0.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* fix hard-coded path and default to "/" (#21253)

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index 55d0bf8b..7709c9d6 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -67,6 +67,8 @@ ingress:
   ## Enable Ingress.
   ##
   enabled: false
+  # AWS ALB requires path of /*
+  path: /
 
     ## Annotations.
     ##

```

## 1.2.14

**Release date:** 2020-01-31

![AppVersion: 1.0.1](https://img.shields.io/static/v1?label=AppVersion&message=1.0.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Add annotation support to pushgateway service (#20351)

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index d9ff624b..55d0bf8b 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -17,6 +17,9 @@ podAnnotations: {}
 # Optional pod labels
 podLabels: {}
 
+# Optional service annotations
+serviceAnnotations: {}
+
 # Optional service labels
 serviceLabels: {}
 

```

## 1.2.13

**Release date:** 2020-01-13

![AppVersion: 1.0.1](https://img.shields.io/static/v1?label=AppVersion&message=1.0.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-pushgateway] Fix persistence support (#20048)

### Default value changes

```diff
# No changes in this release
```

## 1.2.12

**Release date:** 2020-01-10

![AppVersion: 1.0.1](https://img.shields.io/static/v1?label=AppVersion&message=1.0.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-pushgateway] Add persistence support (#16040)

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index a714a399..d9ff624b 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -23,12 +23,20 @@ serviceLabels: {}
 # Optional serviceAccount labels
 serviceAccountLabels: {}
 
-# Optional additional arguments
-extraArgs: []
+# Optional persistentVolume labels
+persistentVolumeLabels: {}
 
 # Optional additional environment variables
 extraVars: []
 
+## Additional pushgateway container arguments
+##
+## example:
+## extraArgs:
+##  - --persistence.file=/data/pushgateway.data
+##  - --persistence.interval=5m
+extraArgs: []
+
 resources: {}
   # We usually recommend not to specify default resources and to leave this as a conscious
   # choice for the user. This also increases chances charts run on environments with little
@@ -88,6 +96,13 @@ nodeSelector: {}
 
 replicaCount: 1
 
+## Security context to be added to push-gateway pods
+##
+securityContext:
+  fsGroup: 65534
+  runAsUser: 65534
+  runAsNonRoot: true
+
 ## Affinity for pod assignment
 ## Ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#affinity-and-anti-affinity
 affinity: {}
@@ -115,6 +130,50 @@ serviceMonitor:
 # If not set then a PodDisruptionBudget will not be created
 podDisruptionBudget: {}
 
+persistentVolume:
+  ## If true, pushgateway will create/use a Persistent Volume Claim
+  ## If false, use emptyDir
+  ##
+  enabled: false
+
+  ## pushgateway data Persistent Volume access modes
+  ## Must match those of existing PV or dynamic provisioner
+  ## Ref: http://kubernetes.io/docs/user-guide/persistent-volumes/
+  ##
+  accessModes:
+    - ReadWriteOnce
+
+  ## pushgateway data Persistent Volume Claim annotations
+  ##
+  annotations: {}
+
+  ## pushgateway data Persistent Volume existing claim name
+  ## Requires pushgateway.persistentVolume.enabled: true
+  ## If defined, PVC must be created manually before volume will be bound
+  existingClaim: ""
+
+  ## pushgateway data Persistent Volume mount root path
+  ##
+  mountPath: /data
+
+  ## pushgateway data Persistent Volume size
+  ##
+  size: 2Gi
+
+  ## pushgateway data Persistent Volume Storage Class
+  ## If defined, storageClassName: <storageClass>
+  ## If set to "-", storageClassName: "", which disables dynamic provisioning
+  ## If undefined (the default) or set to null, no storageClassName spec is
+  ##   set, choosing the default provisioner.  (gp2 on AWS, standard on
+  ##   GKE, AWS & OpenStack)
+  ##
+  # storageClass: "-"
+
+  ## Subdirectory of pushgateway data Persistent Volume to mount
+  ## Useful if the volume's root directory is not empty
+  ##
+  subPath: ""
+
 # Configuration for clusters with restrictive network policies in place:
 # - allowAll allows access to the PushGateway from any namespace
 # - customSelector is a list of pod/namespaceSelectors to allow access from

```

## 1.2.11

**Release date:** 2020-01-09

![AppVersion: 1.0.1](https://img.shields.io/static/v1?label=AppVersion&message=1.0.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-pushgateway] Allow to specify the nodePort when service type is NodePort (#19936)

### Default value changes

```diff
# No changes in this release
```

## 1.2.10

**Release date:** 2019-12-21

![AppVersion: 1.0.1](https://img.shields.io/static/v1?label=AppVersion&message=1.0.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-pushgateway] Upgrade push-gateway (#19732)

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index 5f204c2c..a714a399 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -3,7 +3,7 @@
 # Declare variables to be passed into your templates.
 image:
   repository: prom/pushgateway
-  tag: v1.0.0
+  tag: v1.0.1
   pullPolicy: IfNotPresent
 
 service:

```

## 1.2.9

**Release date:** 2019-12-20

![AppVersion: 1.0.0](https://img.shields.io/static/v1?label=AppVersion&message=1.0.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-pushgateway] Fix missleading documentation about `metrics.enabled` in values documentation (#19730)

### Default value changes

```diff
# No changes in this release
```

## 1.2.8

**Release date:** 2019-12-19

![AppVersion: 1.0.0](https://img.shields.io/static/v1?label=AppVersion&message=1.0.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-pushgateway] fix default values for podDisruptionBudget and networkPolicy (#19680)

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index c9752632..5f204c2c 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -113,13 +113,13 @@ serviceMonitor:
 
 # The values to set in the PodDisruptionBudget spec (minAvailable/maxUnavailable)
 # If not set then a PodDisruptionBudget will not be created
-podDisruptionBudget:
+podDisruptionBudget: {}
 
 # Configuration for clusters with restrictive network policies in place:
 # - allowAll allows access to the PushGateway from any namespace
 # - customSelector is a list of pod/namespaceSelectors to allow access from
 # These options are mutually exclusive and the latter will take precedence.
-networkPolicy:
+networkPolicy: {}
   # allowAll: true
   # customSelectors:
   #   - namespaceSelector:

```

## 1.2.7

**Release date:** 2019-12-16

![AppVersion: 1.0.0](https://img.shields.io/static/v1?label=AppVersion&message=1.0.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Update readiness liveness probe path for prometheus pushgateway (#19580)

### Default value changes

```diff
# No changes in this release
```

## 1.2.6

**Release date:** 2019-11-22

![AppVersion: 1.0.0](https://img.shields.io/static/v1?label=AppVersion&message=1.0.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-pushgateway] support for networkpolicies (#19057)

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index 8a60aa02..c9752632 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -114,3 +114,17 @@ serviceMonitor:
 # The values to set in the PodDisruptionBudget spec (minAvailable/maxUnavailable)
 # If not set then a PodDisruptionBudget will not be created
 podDisruptionBudget:
+
+# Configuration for clusters with restrictive network policies in place:
+# - allowAll allows access to the PushGateway from any namespace
+# - customSelector is a list of pod/namespaceSelectors to allow access from
+# These options are mutually exclusive and the latter will take precedence.
+networkPolicy:
+  # allowAll: true
+  # customSelectors:
+  #   - namespaceSelector:
+  #       matchLabels:
+  #         type: admin
+  #   - podSelector:
+  #       matchLabels:
+  #         app: myapp

```

## 1.2.5

**Release date:** 2019-11-15

![AppVersion: 1.0.0](https://img.shields.io/static/v1?label=AppVersion&message=1.0.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-pushgateway] Add servicemonitor configuration options (#18907)

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index 26d3ceb7..8a60aa02 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -96,9 +96,13 @@ affinity: {}
 serviceMonitor:
   enabled: false
   namespace: monitoring
-  # fallback to the prometheus default unless specified
+
+  # Fallback to the prometheus default unless specified
   # interval: 10s
 
+  # Fallback to the prometheus default unless specified
+  # scrapeTimeout: 30s
+
   ## Used to pass Labels that are used by the Prometheus installed in your cluster to select Service Monitors to work with
   ## ref: https://github.com/coreos/prometheus-operator/blob/master/Documentation/api.md#prometheusspec
   additionalLabels: {}

```

## 1.2.4

**Release date:** 2019-11-15

![AppVersion: 1.0.0](https://img.shields.io/static/v1?label=AppVersion&message=1.0.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-pushgateway] Adjust servicemonitor labels (#18877)

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index bc19c116..26d3ceb7 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -98,11 +98,11 @@ serviceMonitor:
   namespace: monitoring
   # fallback to the prometheus default unless specified
   # interval: 10s
-  ## Defaults to what's used if you follow CoreOS [Prometheus Install Instructions](https://github.com/helm/charts/tree/master/stable/prometheus-operator#tldr)
-  ## [Prometheus Selector Label](https://github.com/helm/charts/tree/master/stable/prometheus-operator#prometheus-operator-1)
-  ## [Kube Prometheus Selector Label](https://github.com/helm/charts/tree/master/stable/prometheus-operator#exporters)
-  selector:
-    prometheus: kube-prometheus
+
+  ## Used to pass Labels that are used by the Prometheus installed in your cluster to select Service Monitors to work with
+  ## ref: https://github.com/coreos/prometheus-operator/blob/master/Documentation/api.md#prometheusspec
+  additionalLabels: {}
+
   # Retain the job and instance labels of the metrics pushed to the Pushgateway
   # [Scraping Pushgateway](https://github.com/prometheus/pushgateway#configure-the-pushgateway-as-a-target-to-scrape)
   honorLabels: true

```

## 1.2.3

**Release date:** 2019-11-14

![AppVersion: 1.0.0](https://img.shields.io/static/v1?label=AppVersion&message=1.0.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-pushgateway] Use fullname for the servicemonitor name (#18872)

### Default value changes

```diff
# No changes in this release
```

## 1.2.2

**Release date:** 2019-11-08

![AppVersion: 1.0.0](https://img.shields.io/static/v1?label=AppVersion&message=1.0.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-pushgateway] Compat k8s 1.16 (#18598)

### Default value changes

```diff
# No changes in this release
```

## 1.2.1

**Release date:** 2019-11-07

![AppVersion: 1.0.0](https://img.shields.io/static/v1?label=AppVersion&message=1.0.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-pushgateway] Follow the Call for maintainers (#18662)

### Default value changes

```diff
# No changes in this release
```

## 1.2.0

**Release date:** 2019-11-05

![AppVersion: 1.0.0](https://img.shields.io/static/v1?label=AppVersion&message=1.0.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Fix deprecated apiVersion for Deployment (#18586)

### Default value changes

```diff
# No changes in this release
```

## 1.1.1

**Release date:** 2019-10-30

![AppVersion: 1.0.0](https://img.shields.io/static/v1?label=AppVersion&message=1.0.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-pushgateway] Fix podDisruptionBudget template (#18187)

### Default value changes

```diff
# No changes in this release
```

## 1.1.0

**Release date:** 2019-10-21

![AppVersion: 1.0.0](https://img.shields.io/static/v1?label=AppVersion&message=1.0.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-pushgateway] Update default version to v1.0.0 (#18091)

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index b08f2a74..bc19c116 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -3,7 +3,7 @@
 # Declare variables to be passed into your templates.
 image:
   repository: prom/pushgateway
-  tag: v0.9.1
+  tag: v1.0.0
   pullPolicy: IfNotPresent
 
 service:

```

## 1.0.1

**Release date:** 2019-08-03

![AppVersion: 0.9.1](https://img.shields.io/static/v1?label=AppVersion&message=0.9.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* upgrade pushgateway (#16061)

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index adb90803..b08f2a74 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -3,7 +3,7 @@
 # Declare variables to be passed into your templates.
 image:
   repository: prom/pushgateway
-  tag: v0.9.0
+  tag: v0.9.1
   pullPolicy: IfNotPresent
 
 service:

```

## 1.0.0

**Release date:** 2019-07-29

![AppVersion: 0.9.0](https://img.shields.io/static/v1?label=AppVersion&message=0.9.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-pushgateway] upgrade to latest release, set chart to v1 (#15912)

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index 2eaa38ee..adb90803 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -3,7 +3,7 @@
 # Declare variables to be passed into your templates.
 image:
   repository: prom/pushgateway
-  tag: v0.8.0
+  tag: v0.9.0
   pullPolicy: IfNotPresent
 
 service:

```

## 0.4.1

**Release date:** 2019-07-03

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-pushgateway] add optional PodDisruptionBudget (#13293)

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index 77a1921a..2eaa38ee 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -106,3 +106,7 @@ serviceMonitor:
   # Retain the job and instance labels of the metrics pushed to the Pushgateway
   # [Scraping Pushgateway](https://github.com/prometheus/pushgateway#configure-the-pushgateway-as-a-target-to-scrape)
   honorLabels: true
+
+# The values to set in the PodDisruptionBudget spec (minAvailable/maxUnavailable)
+# If not set then a PodDisruptionBudget will not be created
+podDisruptionBudget:

```

## 0.4.0

**Release date:** 2019-04-14

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* upgrade promteheus pushgateway (#13038)

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index ad5a4ce5..77a1921a 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -3,7 +3,7 @@
 # Declare variables to be passed into your templates.
 image:
   repository: prom/pushgateway
-  tag: v0.6.0
+  tag: v0.8.0
   pullPolicy: IfNotPresent
 
 service:

```

## 0.3.1

**Release date:** 2019-03-25

![AppVersion: 0.6.0](https://img.shields.io/static/v1?label=AppVersion&message=0.6.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Allow for customization of env vars on prometheus-pushgateway (#12196)

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index cc06ed09..ad5a4ce5 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -26,6 +26,9 @@ serviceAccountLabels: {}
 # Optional additional arguments
 extraArgs: []
 
+# Optional additional environment variables
+extraVars: []
+
 resources: {}
   # We usually recommend not to specify default resources and to leave this as a conscious
   # choice for the user. This also increases chances charts run on environments with little

```

## 0.3.0

**Release date:** 2019-01-17

![AppVersion: 0.6.0](https://img.shields.io/static/v1?label=AppVersion&message=0.6.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Fix pushgateway labels honoring (#10728)

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index 9900bda6..cc06ed09 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -100,3 +100,6 @@ serviceMonitor:
   ## [Kube Prometheus Selector Label](https://github.com/helm/charts/tree/master/stable/prometheus-operator#exporters)
   selector:
     prometheus: kube-prometheus
+  # Retain the job and instance labels of the metrics pushed to the Pushgateway
+  # [Scraping Pushgateway](https://github.com/prometheus/pushgateway#configure-the-pushgateway-as-a-target-to-scrape)
+  honorLabels: true

```

## 0.2.0

**Release date:** 2018-11-19

![AppVersion: 0.6.0](https://img.shields.io/static/v1?label=AppVersion&message=0.6.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* add prometheus servicemonitor support to pushgateway (#9385)

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index 5dd7a076..9900bda6 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -88,3 +88,15 @@ replicaCount: 1
 ## Affinity for pod assignment
 ## Ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#affinity-and-anti-affinity
 affinity: {}
+
+# Enable this if you're using https://github.com/coreos/prometheus-operator
+serviceMonitor:
+  enabled: false
+  namespace: monitoring
+  # fallback to the prometheus default unless specified
+  # interval: 10s
+  ## Defaults to what's used if you follow CoreOS [Prometheus Install Instructions](https://github.com/helm/charts/tree/master/stable/prometheus-operator#tldr)
+  ## [Prometheus Selector Label](https://github.com/helm/charts/tree/master/stable/prometheus-operator#prometheus-operator-1)
+  ## [Kube Prometheus Selector Label](https://github.com/helm/charts/tree/master/stable/prometheus-operator#exporters)
+  selector:
+    prometheus: kube-prometheus

```

## 0.1.6

**Release date:** 2018-11-06

![AppVersion: 0.6.0](https://img.shields.io/static/v1?label=AppVersion&message=0.6.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Add the serviceAccountLabels variable i forgot to document in #8567 (#8913)

### Default value changes

```diff
# No changes in this release
```

## 0.1.5

**Release date:** 2018-10-30

![AppVersion: 0.6.0](https://img.shields.io/static/v1?label=AppVersion&message=0.6.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Add pod and service labels support (#8567)

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index c6abed5d..5dd7a076 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -14,6 +14,15 @@ service:
 # Optional pod annotations
 podAnnotations: {}
 
+# Optional pod labels
+podLabels: {}
+
+# Optional service labels
+serviceLabels: {}
+
+# Optional serviceAccount labels
+serviceAccountLabels: {}
+
 # Optional additional arguments
 extraArgs: []
 

```

## 0.1.4

**Release date:** 2018-10-30

![AppVersion: 0.6.0](https://img.shields.io/static/v1?label=AppVersion&message=0.6.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Update appVersion and image tag (#8864)

### Default value changes

```diff
# No changes in this release
```

## 0.1.3

**Release date:** 2018-10-26

![AppVersion: 0.6.0](https://img.shields.io/static/v1?label=AppVersion&message=0.6.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Update appVersion and image tag (#8794)

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index 90e08214..c6abed5d 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -3,7 +3,7 @@
 # Declare variables to be passed into your templates.
 image:
   repository: prom/pushgateway
-  tag: v0.4.0
+  tag: v0.6.0
   pullPolicy: IfNotPresent
 
 service:

```

## 0.1.2

**Release date:** 2018-05-09

![AppVersion: 0.4.0](https://img.shields.io/static/v1?label=AppVersion&message=0.4.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* pod annotations on prometheus-pushgateway pod (#5469)

### Default value changes

```diff
diff --git a/charts/prometheus-pushgateway/values.yaml b/charts/prometheus-pushgateway/values.yaml
index 70a2bc42..90e08214 100644
--- a/charts/prometheus-pushgateway/values.yaml
+++ b/charts/prometheus-pushgateway/values.yaml
@@ -11,6 +11,9 @@ service:
   port: 9091
   targetPort: 9091
 
+# Optional pod annotations
+podAnnotations: {}
+
 # Optional additional arguments
 extraArgs: []
 

```

## 0.1.1

**Release date:** 2018-04-04

![AppVersion: 0.4.0](https://img.shields.io/static/v1?label=AppVersion&message=0.4.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* review wrong service port in pushgateway readme (#4679)

### Default value changes

```diff
# No changes in this release
```

## 0.1.0

**Release date:** 2018-04-03

![AppVersion: 0.4.0](https://img.shields.io/static/v1?label=AppVersion&message=0.4.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Add prometheus pushgateway (#4620)

### Default value changes

```diff
# Default values for prometheus-pushgateway.
# This is a YAML-formatted file.
# Declare variables to be passed into your templates.
image:
  repository: prom/pushgateway
  tag: v0.4.0
  pullPolicy: IfNotPresent

service:
  type: ClusterIP
  port: 9091
  targetPort: 9091

# Optional additional arguments
extraArgs: []

resources: {}
  # We usually recommend not to specify default resources and to leave this as a conscious
  # choice for the user. This also increases chances charts run on environments with little
  # resources, such as Minikube. If you do want to specify resources, uncomment the following
  # lines, adjust them as necessary, and remove the curly braces after 'resources:'.
  # limits:
  #   cpu: 200m
  #    memory: 50Mi
  # requests:
  #   cpu: 100m
  #   memory: 30Mi

serviceAccount:
  # Specifies whether a ServiceAccount should be created
  create: true
  # The name of the ServiceAccount to use.
  # If not set and create is true, a name is generated using the fullname template
  name:

## Configure ingress resource that allow you to access the
## pushgateway installation. Set up the URL
## ref: http://kubernetes.io/docs/user-guide/ingress/
##
ingress:
  ## Enable Ingress.
  ##
  enabled: false

    ## Annotations.
    ##
    # annotations:
    #   kubernetes.io/ingress.class: nginx
    #   kubernetes.io/tls-acme: 'true'

    ## Hostnames.
    ## Must be provided if Ingress is enabled.
    ##
    # hosts:
    #   - pushgateway.domain.com

    ## TLS configuration.
    ## Secrets must be manually created in the namespace.
    ##
    # tls:
    #   - secretName: pushgateway-tls
    #     hosts:
    #       - pushgateway.domain.com

tolerations: {}
  # - effect: NoSchedule
  #   operator: Exists

## Node labels for pushgateway pod assignment
## Ref: https://kubernetes.io/docs/user-guide/node-selection/
##
nodeSelector: {}

replicaCount: 1

## Affinity for pod assignment
## Ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#affinity-and-anti-affinity
affinity: {}

```

---
Autogenerated from Helm Chart and git history using [helm-changelog](https://github.com/mogensen/helm-changelog)
