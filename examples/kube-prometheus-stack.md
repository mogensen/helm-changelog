# Change Log

## 46.8.0

**Release date:** 2023-06-08

![AppVersion: v0.65.2](https://img.shields.io/static/v1?label=AppVersion&message=v0.65.2&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add Alertmanager as default data source (#3474)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 90278ff8..86380b83 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -962,6 +962,11 @@ grafana:
       exemplarTraceIdDestinations: {}
         # datasourceUid: Jaeger
         # traceIdLabelName: trace_id
+      alertmanager:
+        enabled: true
+        uid: alertmanager
+        handleGrafanaManagedAlerts: false
+        implementation: prometheus
 
   extraConfigmapMounts: []
   # - name: certs-configmap

```

## 46.7.0

**Release date:** 2023-06-07

![AppVersion: v0.65.2](https://img.shields.io/static/v1?label=AppVersion&message=v0.65.2&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] bump chart dependencies (#3473)

### Default value changes

```diff
# No changes in this release
```

## 46.6.0

**Release date:** 2023-06-02

![AppVersion: v0.65.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.65.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] add AlertmanagerSpec clusterGossipInterval, clusterPushpullInterval and clusterPeerTimeout (#3445)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 7018316e..90278ff8 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -808,6 +808,18 @@ alertmanager:
     ##
     clusterAdvertiseAddress: false
 
+    ## clusterGossipInterval determines interval between gossip attempts.
+    ## Needs to be specified as GoDuration, a time duration that can be parsed by Go’s time.ParseDuration() (e.g. 45ms, 30s, 1m, 1h20m15s)
+    clusterGossipInterval: ""
+
+    ## clusterPeerTimeout determines timeout for cluster peering.
+    ## Needs to be specified as GoDuration, a time duration that can be parsed by Go’s time.ParseDuration() (e.g. 45ms, 30s, 1m, 1h20m15s)
+    clusterPeerTimeout: ""
+
+    ## clusterPushpullInterval determines interval between pushpull attempts.
+    ## Needs to be specified as GoDuration, a time duration that can be parsed by Go’s time.ParseDuration() (e.g. 45ms, 30s, 1m, 1h20m15s)
+    clusterPushpullInterval: ""
+
     ## ForceEnableClusterMode ensures Alertmanager does not deactivate the cluster mode when running with a single replica.
     ## Use case is e.g. spanning an Alertmanager cluster across Kubernetes clusters with a single replica in each.
     forceEnableClusterMode: false

```

## 46.5.0

**Release date:** 2023-05-31

![AppVersion: v0.65.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.65.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] - add scrapeConfig CR and tpl support for all CR's selectors (#3446)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 23180913..7018316e 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -3028,6 +3028,28 @@ prometheus:
     #   matchLabels:
     #     prometheus: somelabel
 
+    ## If true, a nil or {} value for prometheus.prometheusSpec.scrapeConfigSelector will cause the
+    ## prometheus resource to be created with selectors based on values in the helm deployment,
+    ## which will also match the scrapeConfigs created
+    ##
+    scrapeConfigSelectorNilUsesHelmValues: true
+
+    ## scrapeConfigs to be selected for target discovery.
+    ## If {}, select all scrapeConfigs
+    ##
+    scrapeConfigSelector: {}
+    ## Example which selects scrapeConfigs with label "prometheus" set to "somelabel"
+    # scrapeConfig:
+    #   matchLabels:
+    #     prometheus: somelabel
+
+    ## If nil, select own namespace. Namespaces to be selected for scrapeConfig discovery.
+    scrapeConfigNamespaceSelector: {}
+    ## Example which selects scrapeConfig in namespaces with label "prometheus" set to "somelabel"
+    # scrapeConfigsNamespaceSelector:
+    #   matchLabels:
+    #     prometheus: somelabel
+
     ## How long to retain metrics
     ##
     retention: 10d

```

## 46.4.2

**Release date:** 2023-05-30

![AppVersion: v0.65.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.65.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Make failurePolicy a string (#3442)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index daa2bf74..23180913 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1949,7 +1949,7 @@ prometheusOperator:
   admissionWebhooks:
     ## Valid values: Fail, Ignore, IgnoreOnInstallOnly
     ## IgnoreOnInstallOnly - If Release.IsInstall returns "true", set "Ignore" otherwise "Fail"
-    failurePolicy:
+    failurePolicy: ""
     ## The default timeoutSeconds is 10 and the maximum value is 30.
     timeoutSeconds: 10
     enabled: true

```

## 46.4.1

**Release date:** 2023-05-25

![AppVersion: v0.65.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.65.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Fix typo in kubeProxy additional rule labels (#3432)

### Default value changes

```diff
# No changes in this release
```

## 46.4.0

**Release date:** 2023-05-24

![AppVersion: v0.65.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.65.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add the possibility to define fine-grained labels and annotation based on the prometheus group (#3288)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 18fc0b11..daa2bf74 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -75,6 +75,66 @@ defaultRules:
   ## Additional annotations for PrometheusRule alerts
   additionalRuleAnnotations: {}
 
+  ## Additional labels for specific PrometheusRule alert groups
+  additionalRuleGroupLabels:
+    alertmanager: {}
+    etcd: {}
+    configReloaders: {}
+    general: {}
+    k8s: {}
+    kubeApiserverAvailability: {}
+    kubeApiserverBurnrate: {}
+    kubeApiserverHistogram: {}
+    kubeApiserverSlos: {}
+    kubeControllerManager: {}
+    kubelet: {}
+    kubeProxy: {}
+    kubePrometheusGeneral: {}
+    kubePrometheusNodeRecording: {}
+    kubernetesApps: {}
+    kubernetesResources: {}
+    kubernetesStorage: {}
+    kubernetesSystem: {}
+    kubeSchedulerAlerting: {}
+    kubeSchedulerRecording: {}
+    kubeStateMetrics: {}
+    network: {}
+    node: {}
+    nodeExporterAlerting: {}
+    nodeExporterRecording: {}
+    prometheus: {}
+    prometheusOperator: {}
+
+  ## Additional annotations for specific PrometheusRule alerts groups
+  additionalRuleGroupAnnotations:
+    alertmanager: {}
+    etcd: {}
+    configReloaders: {}
+    general: {}
+    k8s: {}
+    kubeApiserverAvailability: {}
+    kubeApiserverBurnrate: {}
+    kubeApiserverHistogram: {}
+    kubeApiserverSlos: {}
+    kubeControllerManager: {}
+    kubelet: {}
+    kubeProxy: {}
+    kubePrometheusGeneral: {}
+    kubePrometheusNodeRecording: {}
+    kubernetesApps: {}
+    kubernetesResources: {}
+    kubernetesStorage: {}
+    kubernetesSystem: {}
+    kubeSchedulerAlerting: {}
+    kubeSchedulerRecording: {}
+    kubeStateMetrics: {}
+    network: {}
+    node: {}
+    nodeExporterAlerting: {}
+    nodeExporterRecording: {}
+    prometheus: {}
+    prometheusOperator: {}
+
   ## Prefix for runbook URLs. Use this to override the first part of the runbookURLs that is common to all rules.
   runbookUrl: "https://runbooks.prometheus-operator.dev/runbooks"
 

```

## 46.3.0

**Release date:** 2023-05-23

![AppVersion: v0.65.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.65.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] set parameters for podsecurity restricted (#3201)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index f8989e8c..18fc0b11 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -698,6 +698,8 @@ alertmanager:
       runAsNonRoot: true
       runAsUser: 1000
       fsGroup: 2000
+      seccompProfile:
+        type: RuntimeDefault
 
     ## ListenLocal makes the Alertmanager server listen on loopback, so that it does not bind against the Pod IP.
     ## Note this is only for the Alertmanager UI, not the gossip communication.
@@ -1929,14 +1931,26 @@ prometheusOperator:
         runAsGroup: 2000
         runAsNonRoot: true
         runAsUser: 2000
+        seccompProfile:
+          type: RuntimeDefault
 
     # Security context for create job container
     createSecretJob:
-      securityContext: {}
+      securityContext:
+        allowPrivilegeEscalation: false
+        readOnlyRootFilesystem: true
+        capabilities:
+          drop:
+          - ALL
 
       # Security context for patch job container
     patchWebhookJob:
-      securityContext: {}
+      securityContext:
+        allowPrivilegeEscalation: false
+        readOnlyRootFilesystem: true
+        capabilities:
+          drop:
+          - ALL
 
     # Use certmanager to generate webhook certs
     certManager:
@@ -2178,6 +2192,8 @@ prometheusOperator:
     runAsGroup: 65534
     runAsNonRoot: true
     runAsUser: 65534
+    seccompProfile:
+      type: RuntimeDefault
 
   ## Container-specific security context configuration
   ## ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
@@ -2185,6 +2201,9 @@ prometheusOperator:
   containerSecurityContext:
     allowPrivilegeEscalation: false
     readOnlyRootFilesystem: true
+    capabilities:
+      drop:
+      - ALL
 
   # Enable vertical pod autoscaler support for prometheus-operator
   verticalPodAutoscaler:
@@ -3201,6 +3220,8 @@ prometheus:
       runAsNonRoot: true
       runAsUser: 1000
       fsGroup: 2000
+      seccompProfile:
+        type: RuntimeDefault
 
     ## Priority class assigned to the Pods
     ##
@@ -3832,6 +3853,8 @@ thanosRuler:
       runAsNonRoot: true
       runAsUser: 1000
       fsGroup: 2000
+      seccompProfile:
+        type: RuntimeDefault
 
     ## ListenLocal makes the ThanosRuler server listen on loopback, so that it does not bind against the Pod IP.
     ## Note this is only for the ThanosRuler UI, not the gossip communication.

```

## 46.2.0

**Release date:** 2023-05-23

![AppVersion: v0.65.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.65.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add support for extra manifests (#2798)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 0712ee7a..f8989e8c 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -3874,3 +3874,13 @@ thanosRuler:
 ## Setting to true produces cleaner resource names, but requires a data migration because the name of the persistent volume changes. Therefore this should only be set once on initial installation.
 ##
 cleanPrometheusOperatorObjectNames: false
+
+## Extra manifests to deploy as an array
+extraManifests: []
+  # - apiVersion: v1
+  #   kind: ConfigMap
+  #   metadata:
+  #   labels:
+  #     name: prometheus-extra
+  #   data:
+  #     extra-data: "value"

```

## 46.1.0

**Release date:** 2023-05-23

![AppVersion: v0.65.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.65.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add scheme and tlsConfig to alertmanagerSpec (#3037)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index bbdfa139..0712ee7a 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -620,6 +620,13 @@ alertmanager:
     ##
     routePrefix: /
 
+    ## scheme: HTTP scheme to use. Can be used with `tlsConfig` for example if using istio mTLS.
+    scheme: ""
+
+    ## tlsConfig: TLS configuration to use when connect to the endpoint. For example if using istio mTLS.
+    ## Of type: https://github.com/coreos/prometheus-operator/blob/main/Documentation/api.md#tlsconfig
+    tlsConfig: {}
+
     ## If set to true all actions on the underlying managed objects are not going to be performed, except for delete actions.
     ##
     paused: false

```

## 46.0.0

**Release date:** 2023-05-23

![AppVersion: v0.65.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.65.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Minify Grafana Dashboards and Update CRDs to 0.65.1 (#3416)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 6bb6cab1..bbdfa139 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -2252,7 +2252,7 @@ prometheusOperator:
   thanosImage:
     registry: quay.io
     repository: thanos/thanos
-    tag: v0.30.2
+    tag: v0.31.0
     sha: ""
 
   ## Set a Label Selector to filter watched prometheus and prometheusAgent
@@ -2755,7 +2755,7 @@ prometheus:
     image:
       registry: quay.io
       repository: prometheus/prometheus
-      tag: v2.42.0
+      tag: v2.44.0
       sha: ""
 
     ## Tolerations for use with node taints
@@ -3639,7 +3639,7 @@ thanosRuler:
     image:
       registry: quay.io
       repository: thanos/thanos
-      tag: v0.30.2
+      tag: v0.31.0
       sha: ""
 
     ## Namespaces to be selected for PrometheusRules discovery.

```

## 45.31.1

**Release date:** 2023-05-23

![AppVersion: v0.65.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.65.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-operator] Set secretFieldSelector to exclude not necessary secret types  (#3415)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index bcc433ed..6bb6cab1 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -2268,7 +2268,7 @@ prometheusOperator:
 
   ## Set a Field Selector to filter watched secrets
   ##
-  secretFieldSelector: ""
+  secretFieldSelector: "type!=kubernetes.io/dockercfg,type!=kubernetes.io/service-account-token,type!=helm.sh/release.v1"
 
 ## Deploy a Prometheus instance
 ##

```

## 45.31.0

**Release date:** 2023-05-22

![AppVersion: v0.65.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.65.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] feat: allow override of service name for alertmanager default ingress (#3361)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 07730f57..bcc433ed 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -269,8 +269,11 @@ alertmanager:
 
     labels: {}
 
-    ## Redirect ingress to an additional defined port on the service
+    ## Override ingress to a different defined port on the service
     # servicePort: 8081
+    ## Override ingress to a different service then the default, this is useful if you need to
+    ## point to a specific instance of the alertmanager (eg kube-prometheus-stack-alertmanager-0)
+    # serviceName: kube-prometheus-stack-alertmanager-0
 
     ## Hosts must be provided if Ingress is enabled.
     ##

```

## 45.30.1

**Release date:** 2023-05-22

![AppVersion: v0.65.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.65.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] - resolve issue 3396, no null labels allowed for PrometheusRules (#3400)

### Default value changes

```diff
# No changes in this release
```

## 45.30.0

**Release date:** 2023-05-22

![AppVersion: v0.65.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.65.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Fix issue with depreceted HorizontalPodAutoscaler/v1beta (#3413)

### Default value changes

```diff
# No changes in this release
```

## 45.29.0

**Release date:** 2023-05-19

![AppVersion: v0.65.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.65.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Introduce failurePolicy "IgnoreOnInstallOnly" (#3066)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 95cc696c..07730f57 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1875,6 +1875,8 @@ prometheusOperator:
   ## Admission webhook support for PrometheusRules resources added in Prometheus Operator 0.30 can be enabled to prevent incorrectly formatted
   ## rules from making their way into prometheus and potentially preventing the container from starting
   admissionWebhooks:
+    ## Valid values: Fail, Ignore, IgnoreOnInstallOnly
+    ## IgnoreOnInstallOnly - If Release.IsInstall returns "true", set "Ignore" otherwise "Fail"
     failurePolicy:
     ## The default timeoutSeconds is 10 and the maximum value is 30.
     timeoutSeconds: 10

```

## 45.28.1

**Release date:** 2023-05-17

![AppVersion: v0.65.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.65.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Resolve Issue 3340 (#3351)

### Default value changes

```diff
# No changes in this release
```

## 45.28.0

**Release date:** 2023-05-15

![AppVersion: v0.65.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.65.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add sessionAffinity support to alertmanager (#3376)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 0893c19a..95cc696c 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -379,6 +379,11 @@ alertmanager:
     ##
     externalTrafficPolicy: Cluster
 
+    ## If you want to make sure that connections from a particular client are passed to the same Pod each time
+    ## Accepts 'ClientIP' or ''
+    ##
+    sessionAffinity: ""
+
     ## Service type
     ##
     type: ClusterIP

```

## 45.27.2

**Release date:** 2023-05-10

![AppVersion: v0.65.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.65.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Fix body must be of type string bug for cilium network policy (#3360)

### Default value changes

```diff
# No changes in this release
```

## 45.27.1

**Release date:** 2023-05-09

![AppVersion: v0.65.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.65.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Avoid diff with server-side apply when enforcedNamespaceLabel is set (#3247)

### Default value changes

```diff
# No changes in this release
```

## 45.27.0

**Release date:** 2023-05-09

![AppVersion: v0.65.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.65.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Remove deprecated sha in prometheusSpec (#3191)

### Default value changes

```diff
# No changes in this release
```

## 45.26.0

**Release date:** 2023-05-08

![AppVersion: v0.65.1](https://img.shields.io/static/v1?label=AppVersion&message=v0.65.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] add prometheus config reloader liveness and readiness prob… (#3337)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index ac3b63f9..0893c19a 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -2225,6 +2225,9 @@ prometheusOperator:
       tag: ""
       sha: ""
 
+    # add prometheus config reloader liveness and readiness probe. Default: false
+    enableProbe: false
+
     # resource config for prometheusConfigReloader
     resources:
       requests:

```

## 45.25.0

**Release date:** 2023-05-04

![AppVersion: v0.63.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.63.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Added namespace selector for webhook configuration (#2969)

### Default value changes

```diff
# No changes in this release
```

## 45.24.0

**Release date:** 2023-05-02

![AppVersion: v0.63.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.63.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Upgrade Grafana to 6.56.* (#3302)

### Default value changes

```diff
# No changes in this release
```

## 45.23.0

**Release date:** 2023-04-28

![AppVersion: v0.63.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.63.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Namespace templating for Thanos sidecar ingress (#3301)

### Default value changes

```diff
# No changes in this release
```

## 45.22.0

**Release date:** 2023-04-28

![AppVersion: v0.63.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.63.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add cilium networkpolicy support (#3282)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 7db32312..ac3b63f9 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1963,6 +1963,15 @@ prometheusOperator:
     ##
     enabled: false
 
+    ## Flavor of the network policy to use.
+    #  Can be:
+    #  * kubernetes for networking.k8s.io/v1/NetworkPolicy
+    #  * cilium     for cilium.io/v2/CiliumNetworkPolicy
+    flavor: kubernetes
+
+    # cilium:
+    #   egress:
+
   ## Service account for Alertmanager to use.
   ## ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
   ##
@@ -2260,6 +2269,18 @@ prometheus:
   ## Configure network policy for the prometheus
   networkPolicy:
     enabled: false
+
+    ## Flavor of the network policy to use.
+    #  Can be:
+    #  * kubernetes for networking.k8s.io/v1/NetworkPolicy
+    #  * cilium     for cilium.io/v2/CiliumNetworkPolicy
+    flavor: kubernetes
+
+    # cilium:
+    #   endpointSelector:
+    #   egress:
+    #   ingress:
+
     # egress:
     # - {}
     # ingress:

```

## 45.21.0

**Release date:** 2023-04-25

![AppVersion: v0.63.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.63.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Sync Prometheus rules (#3223)

### Default value changes

```diff
# No changes in this release
```

## 45.20.0

**Release date:** 2023-04-24

![AppVersion: v0.63.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.63.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] allow httpMethod selection on prometheus datasource (#3053)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index e6f22d5a..7db32312 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -857,6 +857,9 @@ grafana:
       ##
       annotations: {}
 
+      ## Set method for HTTP to send query to datasource
+      httpMethod: POST
+
       ## Create datasource for each Pod of Prometheus StatefulSet;
       ## this uses headless service `prometheus-operated` which is
       ## created by Prometheus Operator

```

## 45.19.0

**Release date:** 2023-04-22

![AppVersion: v0.63.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.63.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add flag to disable all exporters (#3256)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index b949b80b..e6f22d5a 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -935,6 +935,11 @@ grafana:
     #   replacement: $1
     #   action: replace
 
+## Flag to disable all the kubernetes component scrapers
+##
+kubernetesServiceMonitors:
+  enabled: true
+
 ## Component scraping the kube api server
 ##
 kubeApiServer:

```

## 45.18.0

**Release date:** 2023-04-21

![AppVersion: v0.63.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.63.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] add a variable to manage grafana default datasource timeout (#3150)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 145e1d34..b949b80b 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -847,6 +847,9 @@ grafana:
       ##
       # url: http://prometheus-stack-prometheus:9090/
 
+      ## Prometheus request timeout in seconds
+      # timeout: 30
+
       # If not defined, will use prometheus.prometheusSpec.scrapeInterval or its default
       # defaultDatasourceScrapeInterval: 15s
 

```

## 45.17.0

**Release date:** 2023-04-20

![AppVersion: v0.63.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.63.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] adding tpl upstream for secretFieldSelector (#3151)

### Default value changes

```diff
# No changes in this release
```

## 45.16.0

**Release date:** 2023-04-20

![AppVersion: v0.63.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.63.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] add alertmanager-instance-selector and thanos… (#3113)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index d973fa24..145e1d34 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -2222,6 +2222,17 @@ prometheusOperator:
     tag: v0.30.2
     sha: ""
 
+  ## Set a Label Selector to filter watched prometheus and prometheusAgent
+  ##
+  prometheusInstanceSelector: ""
+
+  ## Set a Label Selector to filter watched alertmanager
+  ##
+  alertmanagerInstanceSelector: ""
+
+  ## Set a Label Selector to filter watched thanosRuler
+  thanosRulerInstanceSelector: ""
+
   ## Set a Field Selector to filter watched secrets
   ##
   secretFieldSelector: ""

```

## 45.15.0

**Release date:** 2023-04-19

![AppVersion: v0.63.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.63.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Support setting Prometheus version (#3112)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 89215587..d973fa24 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -2673,6 +2673,10 @@ prometheus:
     ##
     enableAdminAPI: false
 
+    ## Sets version of Prometheus overriding the Prometheus version as derived
+    ## from the image tag. Useful in cases where the tag does not follow semver v2.
+    version: ""
+
     ## WebTLSConfig defines the TLS parameters for HTTPS
     ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#webtlsconfig
     web: {}

```

## 45.14.0

**Release date:** 2023-04-19

![AppVersion: v0.63.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.63.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add automountServiceAccountToken option to kube-prometheus-stack aler… (#3180)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 0115d59d..89215587 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -158,6 +158,7 @@ alertmanager:
     create: true
     name: ""
     annotations: {}
+    automountServiceAccountToken: true
 
   ## Configure pod disruption budgets for Alertmanager
   ## ref: https://kubernetes.io/docs/tasks/run-application/configure-pdb/#specifying-a-poddisruptionbudget

```

## 45.13.0

**Release date:** 2023-04-19

![AppVersion: v0.63.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.63.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] update grafana dep (#3264)

### Default value changes

```diff
# No changes in this release
```

## 45.12.0

**Release date:** 2023-04-19

![AppVersion: v0.63.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.63.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Allow dashboard in all namespaces (#2927)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 5ca8fa41..0115d59d 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -822,6 +822,8 @@ grafana:
       enabled: true
       label: grafana_dashboard
       labelValue: "1"
+      # Allow discovery in all namespaces for dashboards
+      searchNamespace: ALL
 
       ## Annotations for Grafana dashboard configmaps
       ##

```

## 45.11.1

**Release date:** 2023-04-19

![AppVersion: v0.63.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.63.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] fix: webhook ignores namespaceOverride (#3259)

### Default value changes

```diff
# No changes in this release
```

## 45.11.0

**Release date:** 2023-04-19

![AppVersion: v0.63.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.63.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] bump chart dependencies (#3255)

### Default value changes

```diff
# No changes in this release
```

## 45.10.1

**Release date:** 2023-04-14

![AppVersion: v0.63.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.63.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Fix incorrect values.yaml descriptions (#3040)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 8a47b477..5ca8fa41 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -2780,11 +2780,12 @@ prometheus:
     ##
     query: {}
 
-    ## Namespaces to be selected for PrometheusRules discovery.
-    ## If nil, select own namespace. Namespaces to be selected for ServiceMonitor discovery.
-    ## See https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#namespaceselector for usage
-    ##
+    ## If nil, select own namespace. Namespaces to be selected for PrometheusRules discovery.
     ruleNamespaceSelector: {}
+    ## Example which selects PrometheusRules in namespaces with label "prometheus" set to "somelabel"
+    # ruleNamespaceSelector:
+    #   matchLabels:
+    #     prometheus: somelabel
 
     ## If true, a nil or {} value for prometheus.prometheusSpec.ruleSelector will cause the
     ## prometheus resource to be created with selectors based on values in the helm deployment,
@@ -2849,10 +2850,12 @@ prometheus:
     #   matchLabels:
     #     prometheus: somelabel
 
-    ## Namespaces to be selected for PodMonitor discovery.
-    ## See https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#namespaceselector for usage
-    ##
+    ## If nil, select own namespace. Namespaces to be selected for PodMonitor discovery.
     podMonitorNamespaceSelector: {}
+    ## Example which selects PodMonitor in namespaces with label "prometheus" set to "somelabel"
+    # podMonitorNamespaceSelector:
+    #   matchLabels:
+    #     prometheus: somelabel
 
     ## If true, a nil or {} value for prometheus.prometheusSpec.probeSelector will cause the
     ## prometheus resource to be created with selectors based on values in the helm deployment,
@@ -2869,10 +2872,12 @@ prometheus:
     #   matchLabels:
     #     prometheus: somelabel
 
-    ## Namespaces to be selected for Probe discovery.
-    ## See https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#namespaceselector for usage
-    ##
+    ## If nil, select own namespace. Namespaces to be selected for Probe discovery.
     probeNamespaceSelector: {}
+    ## Example which selects Probe in namespaces with label "prometheus" set to "somelabel"
+    # probeNamespaceSelector:
+    #   matchLabels:
+    #     prometheus: somelabel
 
     ## How long to retain metrics
     ##

```

## 45.10.0

**Release date:** 2023-04-13

![AppVersion: v0.63.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.63.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] bump chart dependencies (#3234)

### Default value changes

```diff
# No changes in this release
```

## 45.9.1

**Release date:** 2023-04-05

![AppVersion: v0.63.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.63.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] fix kps deployment when thanos field is set to null (#3163)

### Default value changes

```diff
# No changes in this release
```

## 45.9.0

**Release date:** 2023-04-04

![AppVersion: v0.63.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.63.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Additional labels for ServiceMonitors (#3132)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index c284b73e..8a47b477 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -420,6 +420,10 @@ alertmanager:
     interval: ""
     selfMonitor: true
 
+    ## Additional labels
+    ##
+    additionalLabels: {}
+
     ## SampleLimit defines per-scrape limit on number of scraped samples that will be accepted.
     ##
     sampleLimit: 0
@@ -2290,6 +2294,10 @@ prometheus:
     enabled: false
     interval: ""
 
+    ## Additional labels
+    ##
+    additionalLabels: {}
+
     ## scheme: HTTP scheme to use for scraping. Can be used with `tlsConfig` for example if using istio mTLS.
     scheme: ""
 
@@ -2572,6 +2580,10 @@ prometheus:
     interval: ""
     selfMonitor: true
 
+    ## Additional labels
+    ##
+    additionalLabels: {}
+
     ## SampleLimit defines per-scrape limit on number of scraped samples that will be accepted.
     ##
     sampleLimit: 0
@@ -3489,6 +3501,10 @@ thanosRuler:
     interval: ""
     selfMonitor: true
 
+    ## Additional labels
+    ##
+    additionalLabels: {}
+
     ## SampleLimit defines per-scrape limit on number of scraped samples that will be accepted.
     ##
     sampleLimit: 0

```

## 45.8.1

**Release date:** 2023-03-29

![AppVersion: v0.63.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.63.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Import Node Exporter / MacOS grafana dashboard conditionally (#3054)

### Default value changes

```diff
# No changes in this release
```

## 45.8.0

**Release date:** 2023-03-27

![AppVersion: v0.63.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.63.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] - Allow out_of_order_time_window in Prometheus spec (#2960)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 4e5c13e8..c284b73e 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -2870,6 +2870,11 @@ prometheus:
     ##
     retentionSize: ""
 
+    ## Allow out-of-order/out-of-bounds samples ingested into Prometheus for a specified duration
+    ## See https://prometheus.io/docs/prometheus/latest/configuration/configuration/#tsdb
+    tsdb:
+      outOfOrderTimeWindow: 0s
+
     ## Enable compression of the write-ahead log using Snappy.
     ##
     walCompression: true

```

## 45.7.1

**Release date:** 2023-03-08

![AppVersion: v0.63.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.63.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* prometheus-operator: add EndpointSlices for operator to support K8s 1.22+ (#2948)

### Default value changes

```diff
# No changes in this release
```

## 45.7.0

**Release date:** 2023-03-08

![AppVersion: v0.63.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.63.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* bump kube-state-metrics chart for fix global.imageRegistry usage (#3107)

### Default value changes

```diff
# No changes in this release
```

## 45.6.0

**Release date:** 2023-03-06

![AppVersion: v0.63.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.63.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Enable network policy for prometheus (#2997)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 0dbe5902..4e5c13e8 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -2222,13 +2222,23 @@ prometheusOperator:
 ## Deploy a Prometheus instance
 ##
 prometheus:
-
   enabled: true
 
   ## Annotations for Prometheus
   ##
   annotations: {}
 
+  ## Configure network policy for the prometheus
+  networkPolicy:
+    enabled: false
+    # egress:
+    # - {}
+    # ingress:
+    # - {}
+    # podSelector:
+    #   matchLabels:
+    #     app: prometheus
+
   ## Service account for Prometheuses to use.
   ## ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
   ##

```

## 45.5.0

**Release date:** 2023-03-02

![AppVersion: v0.63.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.63.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Update grafana chart (#3086)

### Default value changes

```diff
# No changes in this release
```

## 45.4.0

**Release date:** 2023-02-26

![AppVersion: v0.63.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.63.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Update dependencies to versions which support local container registry (#3059)

### Default value changes

```diff
# No changes in this release
```

## 45.3.0

**Release date:** 2023-02-23

![AppVersion: v0.63.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.63.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Bump version of kube-webhook-certgen (#3061)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 45b08c4b..0dbe5902 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1871,7 +1871,7 @@ prometheusOperator:
       image:
         registry: registry.k8s.io
         repository: ingress-nginx/kube-webhook-certgen
-        tag: v1.3.0
+        tag: v20221220-controller-v1.5.1-58-g787ea74b6
         sha: ""
         pullPolicy: IfNotPresent
       resources: {}

```

## 45.2.0

**Release date:** 2023-02-20

![AppVersion: v0.63.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.63.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Bump kube-state-metrics version (#3046)

### Default value changes

```diff
# No changes in this release
```

## 45.1.1

**Release date:** 2023-02-16

![AppVersion: v0.63.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.63.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add permission to the alertmanagers/status resource (#3016)

### Default value changes

```diff
# No changes in this release
```

## 45.1.0

**Release date:** 2023-02-14

![AppVersion: v0.63.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.63.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add stringConfig (#2992)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 5fd3dacf..45b08c4b 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -212,6 +212,13 @@ alertmanager:
     templates:
     - '/etc/alertmanager/config/*.tmpl'
 
+  ## Alertmanager configuration directives (as string type, preferred over the config hash map)
+  ## stringConfig will be used only, if tplConfig is true
+  ## ref: https://prometheus.io/docs/alerting/configuration/#configuration-file
+  ##      https://prometheus.io/webtools/alerting/routing-tree-editor/
+  ##
+  stringConfig: ""
+
   ## Pass the Alertmanager configuration directives through Helm's templating
   ## engine. If the Alertmanager configuration contains Alertmanager templates,
   ## they'll need to be properly escaped so that they are not interpreted by

```

## 45.0.1

**Release date:** 2023-02-14

![AppVersion: v0.63.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.63.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-kube-stack] Fixes deployment of Thanos Ruler - Addresses Default Thanos Ruler name too long (#3021)

### Default value changes

```diff
# No changes in this release
```

## 45.0.0

**Release date:** 2023-02-09

![AppVersion: v0.63.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.63.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* update kps to v0.63.0 (#3012)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index eef1b654..5fd3dacf 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -2205,7 +2205,7 @@ prometheusOperator:
   thanosImage:
     registry: quay.io
     repository: thanos/thanos
-    tag: v0.30.1
+    tag: v0.30.2
     sha: ""
 
   ## Set a Field Selector to filter watched secrets
@@ -2663,7 +2663,7 @@ prometheus:
     image:
       registry: quay.io
       repository: prometheus/prometheus
-      tag: v2.41.0
+      tag: v2.42.0
       sha: ""
 
     ## Tolerations for use with node taints
@@ -3533,7 +3533,7 @@ thanosRuler:
     image:
       registry: quay.io
       repository: thanos/thanos
-      tag: v0.30.1
+      tag: v0.30.2
       sha: ""
 
     ## Namespaces to be selected for PrometheusRules discovery.

```

## 44.4.1

**Release date:** 2023-02-08

![AppVersion: v0.62.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.62.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Update helm Chart.lock to get latest grafana version.  (#2990)

### Default value changes

```diff
# No changes in this release
```

## 44.4.0

**Release date:** 2023-02-07

![AppVersion: v0.62.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.62.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Support multi-cluster alerting rules for prometheus-operator (#2993)

### Default value changes

```diff
# No changes in this release
```

## 44.3.1

**Release date:** 2023-02-02

![AppVersion: v0.62.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.62.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Fix registry for Config Relaoder and Thanos (#2975)

### Default value changes

```diff
# No changes in this release
```

## 44.3.0

**Release date:** 2023-01-24

![AppVersion: v0.62.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.62.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Update grafana chart (#2946)

### Default value changes

```diff
# No changes in this release
```

## 44.2.1

**Release date:** 2023-01-16

![AppVersion: v0.62.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.62.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Registry change for upstream kube-webhook-certgen (#2924)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 790fd145..eef1b654 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1862,7 +1862,7 @@ prometheusOperator:
     patch:
       enabled: true
       image:
-        registry: k8s.gcr.io
+        registry: registry.k8s.io
         repository: ingress-nginx/kube-webhook-certgen
         tag: v1.3.0
         sha: ""

```

## 44.2.0

**Release date:** 2023-01-16

![AppVersion: v0.62.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.62.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] add alertmanagerConfigMatcherStrategy for alertmanager (#2882)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 7243e992..790fd145 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -553,6 +553,13 @@ alertmanager:
     # alertmanagerConfiguration:
     #   name: global-alertmanager-Configuration
 
+    ## Defines the strategy used by AlertmanagerConfig objects to match alerts. eg:
+    ##
+    alertmanagerConfigMatcherStrategy: {}
+    ## Example with use OnNamespace strategy
+    # alertmanagerConfigMatcherStrategy:
+    #   type: OnNamespace
+
     ## Define Log Format
     # Use logfmt (default) or json logging
     logFormat: logfmt

```

## 44.1.0

**Release date:** 2023-01-16

![AppVersion: v0.62.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.62.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add support for hostAliases in prometheus (#2780)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 9435f708..7243e992 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -3206,6 +3206,14 @@ prometheus:
     # When hostNetwork is enabled, this will set dnsPolicy to ClusterFirstWithHostNet automatically.
     hostNetwork: false
 
+    # HostAlias holds the mapping between IP and hostnames that will be injected
+    # as an entry in the pod’s hosts file.
+    hostAliases: []
+    #  - ip: 10.10.0.100
+    #    hostnames:
+    #      - a1.app.local
+    #      - b1.app.local
+
   additionalRulesForClusterRole: []
   #  - apiGroups: [ "" ]
   #    resources:

```

## 44.0.0

**Release date:** 2023-01-14

![AppVersion: v0.62.0](https://img.shields.io/static/v1?label=AppVersion&message=v0.62.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] More secure admissionWebhooks failurePolicy (#2911)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 2135df47..9435f708 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1838,7 +1838,7 @@ prometheusOperator:
   ## Admission webhook support for PrometheusRules resources added in Prometheus Operator 0.30 can be enabled to prevent incorrectly formatted
   ## rules from making their way into prometheus and potentially preventing the container from starting
   admissionWebhooks:
-    failurePolicy: Fail
+    failurePolicy:
     ## The default timeoutSeconds is 10 and the maximum value is 30.
     timeoutSeconds: 10
     enabled: true
@@ -2153,7 +2153,8 @@ prometheusOperator:
   image:
     registry: quay.io
     repository: prometheus-operator/prometheus-operator
-    tag: v0.61.1
+    # if not set appVersion field from Chart.yaml is used
+    tag: ""
     sha: ""
     pullPolicy: IfNotPresent
 
@@ -2179,7 +2180,8 @@ prometheusOperator:
     image:
       registry: quay.io
       repository: prometheus-operator/prometheus-config-reloader
-      tag: v0.61.1
+      # if not set appVersion field from Chart.yaml is used
+      tag: ""
       sha: ""
 
     # resource config for prometheusConfigReloader
@@ -2196,7 +2198,7 @@ prometheusOperator:
   thanosImage:
     registry: quay.io
     repository: thanos/thanos
-    tag: v0.29.0
+    tag: v0.30.1
     sha: ""
 
   ## Set a Field Selector to filter watched secrets
@@ -2654,7 +2656,7 @@ prometheus:
     image:
       registry: quay.io
       repository: prometheus/prometheus
-      tag: v2.40.5
+      tag: v2.41.0
       sha: ""
 
     ## Tolerations for use with node taints
@@ -3516,7 +3518,7 @@ thanosRuler:
     image:
       registry: quay.io
       repository: thanos/thanos
-      tag: v0.29.0
+      tag: v0.30.1
       sha: ""
 
     ## Namespaces to be selected for PrometheusRules discovery.

```

## 43.3.1

**Release date:** 2023-01-12

![AppVersion: 0.61.1](https://img.shields.io/static/v1?label=AppVersion&message=0.61.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] modify prometheusOperator.serviceMonitor.sampleLimit default value 0 (mean no limit) (#2913)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 55e3e032..2135df47 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -2024,7 +2024,7 @@ prometheusOperator:
 
     ## SampleLimit defines per-scrape limit on number of scraped samples that will be accepted.
     ##
-    sampleLimit: 100
+    sampleLimit: 0
 
     ## TargetLimit defines a limit on the number of scraped targets that will be accepted.
     ##

```

## 43.3.0

**Release date:** 2023-01-11

![AppVersion: 0.61.1](https://img.shields.io/static/v1?label=AppVersion&message=0.61.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] add servicemonitor/podmonitor scrape limits (#2615)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 4ad72b40..55e3e032 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -413,6 +413,26 @@ alertmanager:
     interval: ""
     selfMonitor: true
 
+    ## SampleLimit defines per-scrape limit on number of scraped samples that will be accepted.
+    ##
+    sampleLimit: 0
+
+    ## TargetLimit defines a limit on the number of scraped targets that will be accepted.
+    ##
+    targetLimit: 0
+
+    ## Per-scrape limit on number of labels that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.
+    ##
+    labelLimit: 0
+
+    ## Per-scrape limit on length of labels name that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.
+    ##
+    labelNameLengthLimit: 0
+
+    ## Per-scrape limit on length of labels value that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.
+    ##
+    labelValueLengthLimit: 0
+
     ## proxyUrl: URL of a proxy that should be used for scraping.
     ##
     proxyUrl: ""
@@ -902,6 +922,27 @@ kubeApiServer:
     ## Scrape interval. If not set, the Prometheus default scrape interval is used.
     ##
     interval: ""
+
+    ## SampleLimit defines per-scrape limit on number of scraped samples that will be accepted.
+    ##
+    sampleLimit: 0
+
+    ## TargetLimit defines a limit on the number of scraped targets that will be accepted.
+    ##
+    targetLimit: 0
+
+    ## Per-scrape limit on number of labels that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.
+    ##
+    labelLimit: 0
+
+    ## Per-scrape limit on length of labels name that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.
+    ##
+    labelNameLengthLimit: 0
+
+    ## Per-scrape limit on length of labels value that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.
+    ##
+    labelValueLengthLimit: 0
+
     ## proxyUrl: URL of a proxy that should be used for scraping.
     ##
     proxyUrl: ""
@@ -955,6 +996,26 @@ kubelet:
     ##
     interval: ""
 
+    ## SampleLimit defines per-scrape limit on number of scraped samples that will be accepted.
+    ##
+    sampleLimit: 0
+
+    ## TargetLimit defines a limit on the number of scraped targets that will be accepted.
+    ##
+    targetLimit: 0
+
+    ## Per-scrape limit on number of labels that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.
+    ##
+    labelLimit: 0
+
+    ## Per-scrape limit on length of labels name that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.
+    ##
+    labelNameLengthLimit: 0
+
+    ## Per-scrape limit on length of labels value that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.
+    ##
+    labelValueLengthLimit: 0
+
     ## proxyUrl: URL of a proxy that should be used for scraping.
     ##
     proxyUrl: ""
@@ -1141,6 +1202,26 @@ kubeControllerManager:
     ##
     interval: ""
 
+    ## SampleLimit defines per-scrape limit on number of scraped samples that will be accepted.
+    ##
+    sampleLimit: 0
+
+    ## TargetLimit defines a limit on the number of scraped targets that will be accepted.
+    ##
+    targetLimit: 0
+
+    ## Per-scrape limit on number of labels that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.
+    ##
+    labelLimit: 0
+
+    ## Per-scrape limit on length of labels name that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.
+    ##
+    labelNameLengthLimit: 0
+
+    ## Per-scrape limit on length of labels value that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.
+    ##
+    labelValueLengthLimit: 0
+
     ## proxyUrl: URL of a proxy that should be used for scraping.
     ##
     proxyUrl: ""
@@ -1195,6 +1276,26 @@ coreDns:
     ##
     interval: ""
 
+    ## SampleLimit defines per-scrape limit on number of scraped samples that will be accepted.
+    ##
+    sampleLimit: 0
+
+    ## TargetLimit defines a limit on the number of scraped targets that will be accepted.
+    ##
+    targetLimit: 0
+
+    ## Per-scrape limit on number of labels that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.
+    ##
+    labelLimit: 0
+
+    ## Per-scrape limit on length of labels name that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.
+    ##
+    labelNameLengthLimit: 0
+
+    ## Per-scrape limit on length of labels value that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.
+    ##
+    labelValueLengthLimit: 0
+
     ## proxyUrl: URL of a proxy that should be used for scraping.
     ##
     proxyUrl: ""
@@ -1241,6 +1342,26 @@ kubeDns:
     ##
     interval: ""
 
+    ## SampleLimit defines per-scrape limit on number of scraped samples that will be accepted.
+    ##
+    sampleLimit: 0
+
+    ## TargetLimit defines a limit on the number of scraped targets that will be accepted.
+    ##
+    targetLimit: 0
+
+    ## Per-scrape limit on number of labels that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.
+    ##
+    labelLimit: 0
+
+    ## Per-scrape limit on length of labels name that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.
+    ##
+    labelNameLengthLimit: 0
+
+    ## Per-scrape limit on length of labels value that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.
+    ##
+    labelValueLengthLimit: 0
+
     ## proxyUrl: URL of a proxy that should be used for scraping.
     ##
     proxyUrl: ""
@@ -1325,6 +1446,27 @@ kubeEtcd:
     ## Scrape interval. If not set, the Prometheus default scrape interval is used.
     ##
     interval: ""
+
+    ## SampleLimit defines per-scrape limit on number of scraped samples that will be accepted.
+    ##
+    sampleLimit: 0
+
+    ## TargetLimit defines a limit on the number of scraped targets that will be accepted.
+    ##
+    targetLimit: 0
+
+    ## Per-scrape limit on number of labels that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.
+    ##
+    labelLimit: 0
+
+    ## Per-scrape limit on length of labels name that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.
+    ##
+    labelNameLengthLimit: 0
+
+    ## Per-scrape limit on length of labels value that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.
+    ##
+    labelValueLengthLimit: 0
+
     ## proxyUrl: URL of a proxy that should be used for scraping.
     ##
     proxyUrl: ""
@@ -1388,6 +1530,27 @@ kubeScheduler:
     ## Scrape interval. If not set, the Prometheus default scrape interval is used.
     ##
     interval: ""
+
+    ## SampleLimit defines per-scrape limit on number of scraped samples that will be accepted.
+    ##
+    sampleLimit: 0
+
+    ## TargetLimit defines a limit on the number of scraped targets that will be accepted.
+    ##
+    targetLimit: 0
+
+    ## Per-scrape limit on number of labels that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.
+    ##
+    labelLimit: 0
+
+    ## Per-scrape limit on length of labels name that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.
+    ##
+    labelNameLengthLimit: 0
+
+    ## Per-scrape limit on length of labels value that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.
+    ##
+    labelValueLengthLimit: 0
+
     ## proxyUrl: URL of a proxy that should be used for scraping.
     ##
     proxyUrl: ""
@@ -1452,6 +1615,26 @@ kubeProxy:
     ##
     interval: ""
 
+    ## SampleLimit defines per-scrape limit on number of scraped samples that will be accepted.
+    ##
+    sampleLimit: 0
+
+    ## TargetLimit defines a limit on the number of scraped targets that will be accepted.
+    ##
+    targetLimit: 0
+
+    ## Per-scrape limit on number of labels that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.
+    ##
+    labelLimit: 0
+
+    ## Per-scrape limit on length of labels name that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.
+    ##
+    labelNameLengthLimit: 0
+
+    ## Per-scrape limit on length of labels value that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.
+    ##
+    labelValueLengthLimit: 0
+
     ## proxyUrl: URL of a proxy that should be used for scraping.
     ##
     proxyUrl: ""
@@ -1502,6 +1685,26 @@ kube-state-metrics:
       ##
       interval: ""
 
+      ## SampleLimit defines per-scrape limit on number of scraped samples that will be accepted.
+      ##
+      sampleLimit: 0
+
+      ## TargetLimit defines a limit on the number of scraped targets that will be accepted.
+      ##
+      targetLimit: 0
+
+      ## Per-scrape limit on number of labels that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.
+      ##
+      labelLimit: 0
+
+      ## Per-scrape limit on length of labels name that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.
+      ##
+      labelNameLengthLimit: 0
+
+      ## Per-scrape limit on length of labels value that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.
+      ##
+      labelValueLengthLimit: 0
+
       ## Scrape Timeout. If not set, the Prometheus default scrape timeout is used.
       ##
       scrapeTimeout: ""
@@ -1565,6 +1768,26 @@ prometheus-node-exporter:
       ##
       interval: ""
 
+      ## SampleLimit defines per-scrape limit on number of scraped samples that will be accepted.
+      ##
+      sampleLimit: 0
+
+      ## TargetLimit defines a limit on the number of scraped targets that will be accepted.
+      ##
+      targetLimit: 0
+
+      ## Per-scrape limit on number of labels that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.
+      ##
+      labelLimit: 0
+
+      ## Per-scrape limit on length of labels name that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.
+      ##
+      labelNameLengthLimit: 0
+
+      ## Per-scrape limit on length of labels value that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.
+      ##
+      labelValueLengthLimit: 0
+
       ## How long until a scrape request times out. If not set, the Prometheus default scape timeout is used.
       ##
       scrapeTimeout: ""
@@ -1798,6 +2021,27 @@ prometheusOperator:
     ## Scrape interval. If not set, the Prometheus default scrape interval is used.
     ##
     interval: ""
+
+    ## SampleLimit defines per-scrape limit on number of scraped samples that will be accepted.
+    ##
+    sampleLimit: 100
+
+    ## TargetLimit defines a limit on the number of scraped targets that will be accepted.
+    ##
+    targetLimit: 0
+
+    ## Per-scrape limit on number of labels that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.
+    ##
+    labelLimit: 0
+
+    ## Per-scrape limit on length of labels name that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.
+    ##
+    labelNameLengthLimit: 0
+
+    ## Per-scrape limit on length of labels value that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.
+    ##
+    labelValueLengthLimit: 0
+
     ## Scrape timeout. If not set, the Prometheus default scrape timeout is used.
     scrapeTimeout: ""
     selfMonitor: true
@@ -2302,6 +2546,26 @@ prometheus:
     interval: ""
     selfMonitor: true
 
+    ## SampleLimit defines per-scrape limit on number of scraped samples that will be accepted.
+    ##
+    sampleLimit: 0
+
+    ## TargetLimit defines a limit on the number of scraped targets that will be accepted.
+    ##
+    targetLimit: 0
+
+    ## Per-scrape limit on number of labels that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.
+    ##
+    labelLimit: 0
+
+    ## Per-scrape limit on length of labels name that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.
+    ##
+    labelNameLengthLimit: 0
+
+    ## Per-scrape limit on length of labels value that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.
+    ##
+    labelValueLengthLimit: 0
+
     ## scheme: HTTP scheme to use for scraping. Can be used with `tlsConfig` for example if using istio mTLS.
     scheme: ""
 
@@ -3186,6 +3450,26 @@ thanosRuler:
     interval: ""
     selfMonitor: true
 
+    ## SampleLimit defines per-scrape limit on number of scraped samples that will be accepted.
+    ##
+    sampleLimit: 0
+
+    ## TargetLimit defines a limit on the number of scraped targets that will be accepted.
+    ##
+    targetLimit: 0
+
+    ## Per-scrape limit on number of labels that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.
+    ##
+    labelLimit: 0
+
+    ## Per-scrape limit on length of labels name that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.
+    ##
+    labelNameLengthLimit: 0
+
+    ## Per-scrape limit on length of labels value that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.
+    ##
+    labelValueLengthLimit: 0
+
     ## proxyUrl: URL of a proxy that should be used for scraping.
     ##
     proxyUrl: ""

```

## 43.2.1

**Release date:** 2022-12-29

![AppVersion: 0.61.1](https://img.shields.io/static/v1?label=AppVersion&message=0.61.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Bump version to force use of new versions of subcharts (#2867)

### Default value changes

```diff
# No changes in this release
```

## 43.2.0

**Release date:** 2022-12-25

![AppVersion: 0.61.1](https://img.shields.io/static/v1?label=AppVersion&message=0.61.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] updates alertmanager image to v0.25.0 (#2860)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index e0062fd9..4ad72b40 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -463,7 +463,7 @@ alertmanager:
     image:
       registry: quay.io
       repository: prometheus/alertmanager
-      tag: v0.24.0
+      tag: v0.25.0
       sha: ""
 
     ## If true then the user will be responsible to provide a secret with alertmanager configuration

```

## 43.1.4

**Release date:** 2022-12-23

![AppVersion: 0.61.1](https://img.shields.io/static/v1?label=AppVersion&message=0.61.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add IsDefaultDatasource for editing isDefault value for another DS (#2686)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 8da2879c..e0062fd9 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -798,6 +798,7 @@ grafana:
     datasources:
       enabled: true
       defaultDatasourceEnabled: true
+      isDefaultDatasource: true
 
       uid: prometheus
 

```

## 43.1.3

**Release date:** 2022-12-22

![AppVersion: 0.61.1](https://img.shields.io/static/v1?label=AppVersion&message=0.61.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Remove ruleNamespaceSelector and ruleSelector from CRD, if enableFeatures contains agent (#2648)

### Default value changes

```diff
# No changes in this release
```

## 43.1.2

**Release date:** 2022-12-21

![AppVersion: 0.61.1](https://img.shields.io/static/v1?label=AppVersion&message=0.61.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Fix prometheus-operator VPA targetRef (#2839)

### Default value changes

```diff
# No changes in this release
```

## 43.1.1

**Release date:** 2022-12-16

![AppVersion: 0.61.1](https://img.shields.io/static/v1?label=AppVersion&message=0.61.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] fix Chart.yaml - remove `engine: gotpl` (#2814)

### Default value changes

```diff
# No changes in this release
```

## 43.1.0

**Release date:** 2022-12-15

![AppVersion: 0.61.1](https://img.shields.io/static/v1?label=AppVersion&message=0.61.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Update grafana chart version (#2805)

### Default value changes

```diff
# No changes in this release
```

## 43.0.0

**Release date:** 2022-12-13

![AppVersion: 0.61.1](https://img.shields.io/static/v1?label=AppVersion&message=0.61.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack ]Upgrade to 43.0.0 (#2789)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index e5d805cd..8da2879c 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1908,7 +1908,7 @@ prometheusOperator:
   image:
     registry: quay.io
     repository: prometheus-operator/prometheus-operator
-    tag: v0.60.1
+    tag: v0.61.1
     sha: ""
     pullPolicy: IfNotPresent
 
@@ -1934,7 +1934,7 @@ prometheusOperator:
     image:
       registry: quay.io
       repository: prometheus-operator/prometheus-config-reloader
-      tag: v0.60.1
+      tag: v0.61.1
       sha: ""
 
     # resource config for prometheusConfigReloader
@@ -1951,7 +1951,7 @@ prometheusOperator:
   thanosImage:
     registry: quay.io
     repository: thanos/thanos
-    tag: v0.28.1
+    tag: v0.29.0
     sha: ""
 
   ## Set a Field Selector to filter watched secrets
@@ -2389,7 +2389,7 @@ prometheus:
     image:
       registry: quay.io
       repository: prometheus/prometheus
-      tag: v2.39.1
+      tag: v2.40.5
       sha: ""
 
     ## Tolerations for use with node taints
@@ -3231,7 +3231,7 @@ thanosRuler:
     image:
       registry: quay.io
       repository: thanos/thanos
-      tag: v0.28.1
+      tag: v0.29.0
       sha: ""
 
     ## Namespaces to be selected for PrometheusRules discovery.

```

## 42.3.0

**Release date:** 2022-12-10

![AppVersion: 0.60.1](https://img.shields.io/static/v1?label=AppVersion&message=0.60.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Bump Grafana Chart to most recent commit (#2801)

### Default value changes

```diff
# No changes in this release
```

## 42.2.1

**Release date:** 2022-12-06

![AppVersion: 0.60.1](https://img.shields.io/static/v1?label=AppVersion&message=0.60.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Fix netpol to allow access to the api-server (#2790)

### Default value changes

```diff
# No changes in this release
```

## 42.2.0

**Release date:** 2022-12-04

![AppVersion: 0.60.1](https://img.shields.io/static/v1?label=AppVersion&message=0.60.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack]: bump grafana chart dependency to 6.45.0 and node-exporter to 4.8.0 (#2773)

### Default value changes

```diff
# No changes in this release
```

## 42.1.1

**Release date:** 2022-11-30

![AppVersion: 0.60.1](https://img.shields.io/static/v1?label=AppVersion&message=0.60.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Allow setting timeoutSeconds in ValidatingWebhook (#2761)

### Default value changes

```diff
# No changes in this release
```

## 42.1.0

**Release date:** 2022-11-26

![AppVersion: 0.60.1](https://img.shields.io/static/v1?label=AppVersion&message=0.60.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Bump kube-state-metrics version (#2740)

### Default value changes

```diff
# No changes in this release
```

## 42.0.3

**Release date:** 2022-11-26

![AppVersion: 0.60.1](https://img.shields.io/static/v1?label=AppVersion&message=0.60.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Prevent helm diff on ServiceMonitor - Kubelet (#2732)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 74290a75..e5d805cd 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1037,7 +1037,8 @@ kubelet:
     ##
     ## metrics_path is required to match upstream rules and charts
     cAdvisorRelabelings:
-      - sourceLabels: [__metrics_path__]
+      - action: replace
+        sourceLabels: [__metrics_path__]
         targetLabel: metrics_path
     # - sourceLabels: [__meta_kubernetes_pod_node_name]
     #   separator: ;
@@ -1050,7 +1051,8 @@ kubelet:
     ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#relabelconfig
     ##
     probesRelabelings:
-      - sourceLabels: [__metrics_path__]
+      - action: replace
+        sourceLabels: [__metrics_path__]
         targetLabel: metrics_path
     # - sourceLabels: [__meta_kubernetes_pod_node_name]
     #   separator: ;
@@ -1063,7 +1065,8 @@ kubelet:
     ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#relabelconfig
     ##
     resourceRelabelings:
-      - sourceLabels: [__metrics_path__]
+      - action: replace
+        sourceLabels: [__metrics_path__]
         targetLabel: metrics_path
     # - sourceLabels: [__meta_kubernetes_pod_node_name]
     #   separator: ;
@@ -1092,7 +1095,8 @@ kubelet:
     ##
     ## metrics_path is required to match upstream rules and charts
     relabelings:
-      - sourceLabels: [__metrics_path__]
+      - action: replace
+        sourceLabels: [__metrics_path__]
         targetLabel: metrics_path
     # - sourceLabels: [__meta_kubernetes_pod_node_name]
     #   separator: ;

```

## 42.0.2

**Release date:** 2022-11-24

![AppVersion: 0.60.1](https://img.shields.io/static/v1?label=AppVersion&message=0.60.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Add myself as kube-prometheus-stack maintainer (#2728)

### Default value changes

```diff
# No changes in this release
```

## 42.0.1

**Release date:** 2022-11-24

![AppVersion: 0.60.1](https://img.shields.io/static/v1?label=AppVersion&message=0.60.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add missing upgrade to 42.x info (#2727)

### Default value changes

```diff
# No changes in this release
```

## 42.0.0

**Release date:** 2022-11-22

![AppVersion: 0.60.1](https://img.shields.io/static/v1?label=AppVersion&message=0.60.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add local and global image registry support (#2692)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index c01f9c93..74290a75 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -122,6 +122,10 @@ global:
       # seccomp.security.alpha.kubernetes.io/defaultProfileName: 'docker/default'
       # apparmor.security.beta.kubernetes.io/defaultProfileName: 'runtime/default'
 
+  ## Global image registry to use if it needs to be overriden for some specific use cases (e.g local registries, custom images, ...)
+  ##
+  imageRegistry: ""
+
   ## Reference to one or more secrets to be used when pulling images
   ## ref: https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/
   ##
@@ -457,7 +461,8 @@ alertmanager:
     ## Image of Alertmanager
     ##
     image:
-      repository: quay.io/prometheus/alertmanager
+      registry: quay.io
+      repository: prometheus/alertmanager
       tag: v0.24.0
       sha: ""
 
@@ -1622,7 +1627,8 @@ prometheusOperator:
     patch:
       enabled: true
       image:
-        repository: k8s.gcr.io/ingress-nginx/kube-webhook-certgen
+        registry: k8s.gcr.io
+        repository: ingress-nginx/kube-webhook-certgen
         tag: v1.3.0
         sha: ""
         pullPolicy: IfNotPresent
@@ -1896,24 +1902,34 @@ prometheusOperator:
   ## Prometheus-operator image
   ##
   image:
-    repository: quay.io/prometheus-operator/prometheus-operator
+    registry: quay.io
+    repository: prometheus-operator/prometheus-operator
     tag: v0.60.1
     sha: ""
     pullPolicy: IfNotPresent
 
   ## Prometheus image to use for prometheuses managed by the operator
   ##
-  # prometheusDefaultBaseImage: quay.io/prometheus/prometheus
+  # prometheusDefaultBaseImage: prometheus/prometheus
+
+  ## Prometheus image registry to use for prometheuses managed by the operator
+  ##
+  # prometheusDefaultBaseImageRegistry: quay.io
 
   ## Alertmanager image to use for alertmanagers managed by the operator
   ##
-  # alertmanagerDefaultBaseImage: quay.io/prometheus/alertmanager
+  # alertmanagerDefaultBaseImage: prometheus/alertmanager
+
+  ## Alertmanager image registry to use for alertmanagers managed by the operator
+  ##
+  # alertmanagerDefaultBaseImageRegistry: quay.io
 
   ## Prometheus-config-reloader
   ##
   prometheusConfigReloader:
     image:
-      repository: quay.io/prometheus-operator/prometheus-config-reloader
+      registry: quay.io
+      repository: prometheus-operator/prometheus-config-reloader
       tag: v0.60.1
       sha: ""
 
@@ -1929,7 +1945,8 @@ prometheusOperator:
   ## Thanos side-car image when configured
   ##
   thanosImage:
-    repository: quay.io/thanos/thanos
+    registry: quay.io
+    repository: thanos/thanos
     tag: v0.28.1
     sha: ""
 
@@ -2366,7 +2383,8 @@ prometheus:
     ## Image of Prometheus.
     ##
     image:
-      repository: quay.io/prometheus/prometheus
+      registry: quay.io
+      repository: prometheus/prometheus
       tag: v2.39.1
       sha: ""
 
@@ -3207,7 +3225,8 @@ thanosRuler:
     ## Image of ThanosRuler
     ##
     image:
-      repository: quay.io/thanos/thanos
+      registry: quay.io
+      repository: thanos/thanos
       tag: v0.28.1
       sha: ""
 

```

## 41.9.1

**Release date:** 2022-11-19

![AppVersion: 0.60.1](https://img.shields.io/static/v1?label=AppVersion&message=0.60.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Fix PSP deprecation after k8s 1.25+ (#2691)

### Default value changes

```diff
# No changes in this release
```

## 41.9.0

**Release date:** 2022-11-17

![AppVersion: 0.60.1](https://img.shields.io/static/v1?label=AppVersion&message=0.60.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] add support for additional labels on operator's ServiceMonitor (#2687)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 17c57142..c01f9c93 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1781,6 +1781,9 @@ prometheusOperator:
   ## Create a servicemonitor for the operator
   ##
   serviceMonitor:
+    ## Labels for ServiceMonitor
+    additionalLabels: {}
+
     ## Scrape interval. If not set, the Prometheus default scrape interval is used.
     ##
     interval: ""

```

## 41.8.0

**Release date:** 2022-11-16

![AppVersion: 0.60.1](https://img.shields.io/static/v1?label=AppVersion&message=0.60.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] : enable "hostNetwork" support for "Prometheus" CR Object (#2693)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 47ceaf58..17c57142 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -2908,6 +2908,12 @@ prometheus:
     ## be considered available. Defaults to 0 (pod will be considered available as soon as it is ready).
     minReadySeconds: 0
 
+    # Required for use in managed kubernetes clusters (such as AWS EKS) with custom CNI (such as calico),
+    # because control-plane managed by AWS cannot communicate with pods' IP CIDR and admission webhooks are not working
+    # Use the host's network namespace if true. Make sure to understand the security implications if you want to enable it.
+    # When hostNetwork is enabled, this will set dnsPolicy to ClusterFirstWithHostNet automatically.
+    hostNetwork: false
+
   additionalRulesForClusterRole: []
   #  - apiGroups: [ "" ]
   #    resources:

```

## 41.7.4

**Release date:** 2022-11-11

![AppVersion: 0.60.1](https://img.shields.io/static/v1?label=AppVersion&message=0.60.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack]: bump grafana chart dependency to 6.43.5 (#2674)

### Default value changes

```diff
# No changes in this release
```

## 41.7.3

**Release date:** 2022-11-02

![AppVersion: 0.60.1](https://img.shields.io/static/v1?label=AppVersion&message=0.60.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Bump kube-state-metrics version (#2636)

### Default value changes

```diff
# No changes in this release
```

## 41.7.2

**Release date:** 2022-11-01

![AppVersion: 0.60.1](https://img.shields.io/static/v1?label=AppVersion&message=0.60.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] avoid alerting field in prometheus CR, if alertmanager config is empty (#2600)

### Default value changes

```diff
# No changes in this release
```

## 41.7.1

**Release date:** 2022-10-31

![AppVersion: 0.60.1](https://img.shields.io/static/v1?label=AppVersion&message=0.60.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack]: Fix for alert mgr service monitor: enableHttp2 (#2570)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 06f085db..47ceaf58 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -418,7 +418,7 @@ alertmanager:
 
     ## enableHttp2: Whether to enable HTTP2.
     ## See https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#endpoint
-    enableHttp2: false
+    enableHttp2: true
 
     ## tlsConfig: TLS configuration to use when scraping the endpoint. For example if using istio mTLS.
     ## Of type: https://github.com/coreos/prometheus-operator/blob/main/Documentation/api.md#tlsconfig

```

## 41.7.0

**Release date:** 2022-10-28

![AppVersion: 0.60.1](https://img.shields.io/static/v1?label=AppVersion&message=0.60.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Bump prometheus-node-exporter subchart version (#2619)

### Default value changes

```diff
# No changes in this release
```

## 41.6.1

**Release date:** 2022-10-25

![AppVersion: 0.60.1](https://img.shields.io/static/v1?label=AppVersion&message=0.60.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Remove desaintmartin from maintainers. (#2604)

### Default value changes

```diff
# No changes in this release
```

## 41.6.0

**Release date:** 2022-10-25

![AppVersion: 0.60.1](https://img.shields.io/static/v1?label=AppVersion&message=0.60.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] bump grafana 6.43.0 (9.2.1) (#2602)

### Default value changes

```diff
# No changes in this release
```

## 41.5.1

**Release date:** 2022-10-19

![AppVersion: 0.60.1](https://img.shields.io/static/v1?label=AppVersion&message=0.60.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* fix network policy api template (#2586)

### Default value changes

```diff
# No changes in this release
```

## 41.5.0

**Release date:** 2022-10-17

![AppVersion: 0.60.1](https://img.shields.io/static/v1?label=AppVersion&message=0.60.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] add network policies (#2580)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 00983e25..06f085db 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1692,6 +1692,11 @@ prometheusOperator:
   ##
   # clusterDomain: "cluster.local"
 
+  networkPolicy:
+    ## Enable creation of NetworkPolicy resources.
+    ##
+    enabled: false
+
   ## Service account for Alertmanager to use.
   ## ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
   ##

```

## 41.4.1

**Release date:** 2022-10-15

![AppVersion: 0.60.1](https://img.shields.io/static/v1?label=AppVersion&message=0.60.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] PodSecurityPolicy being removed from Kubernetes (#2561)

### Default value changes

```diff
# No changes in this release
```

## 41.4.0

**Release date:** 2022-10-13

![AppVersion: 0.60.1](https://img.shields.io/static/v1?label=AppVersion&message=0.60.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add Vertical Pod Autoscaler to Prometheus Operator (#2552)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 109a4f93..00983e25 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1865,6 +1865,26 @@ prometheusOperator:
     allowPrivilegeEscalation: false
     readOnlyRootFilesystem: true
 
+  # Enable vertical pod autoscaler support for prometheus-operator
+  verticalPodAutoscaler:
+    enabled: false
+    # List of resources that the vertical pod autoscaler can control. Defaults to cpu and memory
+    controlledResources: []
+
+    # Define the max allowed resources for the pod
+    maxAllowed: {}
+    # cpu: 200m
+    # memory: 100Mi
+    # Define the min allowed resources for the pod
+    minAllowed: {}
+    # cpu: 200m
+    # memory: 100Mi
+
+    updatePolicy:
+      # Specifies whether recommended updates are applied when a Pod is started and whether recommended updates
+      # are applied during the life of a Pod. Possible values are "Off", "Initial", "Recreate", and "Auto".
+      updateMode: Auto
+
   ## Prometheus-operator image
   ##
   image:

```

## 41.3.2

**Release date:** 2022-10-12

![AppVersion: 0.60.1](https://img.shields.io/static/v1?label=AppVersion&message=0.60.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Fix generated label for Prometheus Pod anti-affinity (#2554)

### Default value changes

```diff
# No changes in this release
```

## 41.3.1

**Release date:** 2022-10-12

![AppVersion: 0.60.1](https://img.shields.io/static/v1?label=AppVersion&message=0.60.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* fix servicemonitor (#2549)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 5266d0fc..109a4f93 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -418,7 +418,7 @@ alertmanager:
 
     ## enableHttp2: Whether to enable HTTP2.
     ## See https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#endpoint
-    enableHttp2: true
+    enableHttp2: false
 
     ## tlsConfig: TLS configuration to use when scraping the endpoint. For example if using istio mTLS.
     ## Of type: https://github.com/coreos/prometheus-operator/blob/main/Documentation/api.md#tlsconfig

```

## 41.3.0

**Release date:** 2022-10-12

![AppVersion: 0.60.1](https://img.shields.io/static/v1?label=AppVersion&message=0.60.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Allow annotations on the admissionWebhook Jobs (#2528)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index d9042e8e..5266d0fc 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1616,6 +1616,9 @@ prometheusOperator:
     ## On chart upgrades (or if the secret exists) the cert will not be re-generated. You can use this to provide your own
     ## certs ahead of time if you wish.
     ##
+    annotations: {}
+    #   argocd.argoproj.io/hook: PreSync
+    #   argocd.argoproj.io/hook-delete-policy: HookSucceeded
     patch:
       enabled: true
       image:
@@ -1627,6 +1630,9 @@ prometheusOperator:
       ## Provide a priority class name to the webhook patching job
       ##
       priorityClassName: ""
+      annotations: {}
+      #   argocd.argoproj.io/hook: PreSync
+      #   argocd.argoproj.io/hook-delete-policy: HookSucceeded
       podAnnotations: {}
       nodeSelector: {}
       affinity: {}

```

## 41.2.0

**Release date:** 2022-10-12

![AppVersion: 0.60.1](https://img.shields.io/static/v1?label=AppVersion&message=0.60.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack]: Alert mgr service monitor: enableHttp2 (#2540)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 844aa3eb..d9042e8e 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -416,6 +416,10 @@ alertmanager:
     ## scheme: HTTP scheme to use for scraping. Can be used with `tlsConfig` for example if using istio mTLS.
     scheme: ""
 
+    ## enableHttp2: Whether to enable HTTP2.
+    ## See https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#endpoint
+    enableHttp2: true
+
     ## tlsConfig: TLS configuration to use when scraping the endpoint. For example if using istio mTLS.
     ## Of type: https://github.com/coreos/prometheus-operator/blob/main/Documentation/api.md#tlsconfig
     tlsConfig: {}

```

## 41.1.0

**Release date:** 2022-10-12

![AppVersion: 0.60.1](https://img.shields.io/static/v1?label=AppVersion&message=0.60.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add additionalArgs support to PrometheusSpec (#2482)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 58367134..844aa3eb 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -2280,6 +2280,10 @@ prometheus:
     ##
     apiserverConfig: {}
 
+    ## Allows setting additional arguments for the Prometheus container
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#monitoring.coreos.com/v1.Prometheus
+    additionalArgs: []
+
     ## Interval between consecutive scrapes.
     ## Defaults to 30s.
     ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/release-0.44/pkg/prometheus/promcfg.go#L180-L183

```

## 41.0.1

**Release date:** 2022-10-12

![AppVersion: 0.60.1](https://img.shields.io/static/v1?label=AppVersion&message=0.60.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* fix versions (#2546)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index b24c24a6..58367134 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1859,7 +1859,7 @@ prometheusOperator:
   ##
   image:
     repository: quay.io/prometheus-operator/prometheus-operator
-    tag: v0.59.2
+    tag: v0.60.1
     sha: ""
     pullPolicy: IfNotPresent
 
@@ -1876,7 +1876,7 @@ prometheusOperator:
   prometheusConfigReloader:
     image:
       repository: quay.io/prometheus-operator/prometheus-config-reloader
-      tag: v0.59.2
+      tag: v0.60.1
       sha: ""
 
     # resource config for prometheusConfigReloader
@@ -1892,7 +1892,7 @@ prometheusOperator:
   ##
   thanosImage:
     repository: quay.io/thanos/thanos
-    tag: v0.28.0
+    tag: v0.28.1
     sha: ""
 
   ## Set a Field Selector to filter watched secrets
@@ -2325,7 +2325,7 @@ prometheus:
     ##
     image:
       repository: quay.io/prometheus/prometheus
-      tag: v2.39.0
+      tag: v2.39.1
       sha: ""
 
     ## Tolerations for use with node taints
@@ -3160,7 +3160,7 @@ thanosRuler:
     ##
     image:
       repository: quay.io/thanos/thanos
-      tag: v0.28.0
+      tag: v0.28.1
       sha: ""
 
     ## Namespaces to be selected for PrometheusRules discovery.

```

## 41.0.0

**Release date:** 2022-10-11

![AppVersion: 0.60.1](https://img.shields.io/static/v1?label=AppVersion&message=0.60.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] prom operator update / add recording and alerting switch to kubeScheduler (#2410)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 2e31c5dd..b24c24a6 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -51,7 +51,8 @@ defaultRules:
     kubernetesResources: true
     kubernetesStorage: true
     kubernetesSystem: true
-    kubeScheduler: true
+    kubeSchedulerAlerting: true
+    kubeSchedulerRecording: true
     kubeStateMetrics: true
     network: true
     node: true

```

## 40.5.0

**Release date:** 2022-10-07

![AppVersion: 0.59.2](https://img.shields.io/static/v1?label=AppVersion&message=0.59.2&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Update Prometheus to v2.39.0 (#2524)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 62397399..2e31c5dd 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -2324,7 +2324,7 @@ prometheus:
     ##
     image:
       repository: quay.io/prometheus/prometheus
-      tag: v2.38.0
+      tag: v2.39.0
       sha: ""
 
     ## Tolerations for use with node taints

```

## 40.4.0

**Release date:** 2022-10-05

![AppVersion: 0.59.2](https://img.shields.io/static/v1?label=AppVersion&message=0.59.2&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add queryConfig to thanosRulerSpec (#2522)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index ce3e59d6..62397399 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -3261,6 +3261,14 @@ thanosRuler:
     ## When used alongside with ObjectStorageConfig, ObjectStorageConfigFile takes precedence.
     objectStorageConfigFile: ""
 
+    ## QueryEndpoints defines Thanos querier endpoints from which to query metrics.
+    ## Maps to the --query flag of thanos ruler.
+    queryEndpoints: []
+
+    ## Define configuration for connecting to thanos query instances. If this is defined, the queryEndpoints field will be ignored.
+    ## Maps to the query.config CLI argument. Only available with thanos v0.11.0 and higher.
+    queryConfig: {}
+
     ## Labels configure the external label pairs to ThanosRuler. A default replica
     ## label `thanos_ruler_replica` will be always added as a label with the value
     ## of the pod's name and it will be dropped in the alerts.

```

## 40.3.1

**Release date:** 2022-10-01

![AppVersion: 0.59.2](https://img.shields.io/static/v1?label=AppVersion&message=0.59.2&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] BUMP ksm helm dep (#2509)

### Default value changes

```diff
# No changes in this release
```

## 40.3.0

**Release date:** 2022-09-30

![AppVersion: 0.59.2](https://img.shields.io/static/v1?label=AppVersion&message=0.59.2&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] - prometheus-operator0.59.2/dependencies (#2505)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 195ff31d..ce3e59d6 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1858,7 +1858,7 @@ prometheusOperator:
   ##
   image:
     repository: quay.io/prometheus-operator/prometheus-operator
-    tag: v0.59.1
+    tag: v0.59.2
     sha: ""
     pullPolicy: IfNotPresent
 
@@ -1875,7 +1875,7 @@ prometheusOperator:
   prometheusConfigReloader:
     image:
       repository: quay.io/prometheus-operator/prometheus-config-reloader
-      tag: v0.59.1
+      tag: v0.59.2
       sha: ""
 
     # resource config for prometheusConfigReloader

```

## 40.2.0

**Release date:** 2022-09-29

![AppVersion: 0.59.1](https://img.shields.io/static/v1?label=AppVersion&message=0.59.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add ability to add custom labels in the operator deployment (#2492)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index ee9cca44..195ff31d 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1727,6 +1727,10 @@ prometheusOperator:
     ##
     externalIPs: []
 
+  # ## Labels to add to the operator deployment
+  # ##
+  labels: {}
+
   ## Annotations to add to the operator deployment
   ##
   annotations: {}

```

## 40.1.2

**Release date:** 2022-09-23

![AppVersion: 0.59.1](https://img.shields.io/static/v1?label=AppVersion&message=0.59.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* helm dependency update (#2485)

### Default value changes

```diff
# No changes in this release
```

## 40.1.1

**Release date:** 2022-09-23

![AppVersion: 0.59.1](https://img.shields.io/static/v1?label=AppVersion&message=0.59.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] - Add template option to namespaces flag (#2466)

### Default value changes

```diff
# No changes in this release
```

## 40.1.0

**Release date:** 2022-09-19

![AppVersion: 0.59.1](https://img.shields.io/static/v1?label=AppVersion&message=0.59.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add support for StatefulSet minReadySeconds (#2463)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index e2d19a44..ee9cca44 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -686,6 +686,10 @@ alertmanager:
     ## Use case is e.g. spanning an Alertmanager cluster across Kubernetes clusters with a single replica in each.
     forceEnableClusterMode: false
 
+    ## Minimum number of seconds for which a newly created pod should be ready without any of its container crashing for it to
+    ## be considered available. Defaults to 0 (pod will be considered available as soon as it is ready).
+    minReadySeconds: 0
+
   ## ExtraSecret can be used to store various data in an extra secret
   ## (use it for example to store hashed basic auth credentials)
   extraSecret:
@@ -2856,6 +2860,10 @@ prometheus:
     ## in Prometheus so it may change in any upcoming release.
     allowOverlappingBlocks: false
 
+    ## Minimum number of seconds for which a newly created pod should be ready without any of its container crashing for it to
+    ## be considered available. Defaults to 0 (pod will be considered available as soon as it is ready).
+    minReadySeconds: 0
+
   additionalRulesForClusterRole: []
   #  - apiGroups: [ "" ]
   #    resources:

```

## 40.0.2

**Release date:** 2022-09-16

![AppVersion: 0.59.1](https://img.shields.io/static/v1?label=AppVersion&message=0.59.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add template to storage configurations (#2415)

### Default value changes

```diff
# No changes in this release
```

## 40.0.1

**Release date:** 2022-09-16

![AppVersion: 0.59.1](https://img.shields.io/static/v1?label=AppVersion&message=0.59.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Update port for metric of etcd (#2264)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 74e7c6df..e2d19a44 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1285,8 +1285,8 @@ kubeEtcd:
   ##
   service:
     enabled: true
-    port: 2379
-    targetPort: 2379
+    port: 2381
+    targetPort: 2381
     # selector:
     #   component: etcd
 

```

## 40.0.0

**Release date:** 2022-09-14

![AppVersion: 0.59.1](https://img.shields.io/static/v1?label=AppVersion&message=0.59.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] prometheus-operator0.59.1/Prometheus2.37.1/all dependencies (#2440)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 8de44218..74e7c6df 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -703,14 +703,6 @@ grafana:
   enabled: true
   namespaceOverride: ""
 
-  ## Image and version of grafana. If not provided (left commented out) default values from grafana charts will be used.
-  ##
-  # image:
-  #   repository: docker.io/grafana/grafana
-  #   tag: 9.0.6
-  #   sha: ""
-  #   pullPolicy: IfNotPresent
-
   ## ForceDeployDatasources Create datasource configmap even if grafana deployment has been disabled
   ##
   forceDeployDatasources: false
@@ -774,13 +766,6 @@ grafana:
     #   - grafana.example.com
 
   sidecar:
-    ## Image and version of sidecar. If not provided (left commented out) default values from grafana charts will be used.
-    ##
-    # image:
-    #   repository: quay.io/kiwigrid/k8s-sidecar
-    #   tag: 1.19.2
-    #   sha: ""
-
     dashboards:
       enabled: true
       label: grafana_dashboard
@@ -1486,14 +1471,6 @@ kubeStateMetrics:
 ## Configuration for kube-state-metrics subchart
 ##
 kube-state-metrics:
-  ##
-  ##
-  # image:
-  #  repository: registry.k8s.io/kube-state-metrics/kube-state-metrics
-  #  tag: v2.5.0
-  #  sha: ""
-  #  pullPolicy: IfNotPresent
-
   namespaceOverride: ""
   rbac:
     create: true
@@ -1548,18 +1525,12 @@ nodeExporter:
 ## Configuration for prometheus-node-exporter subchart
 ##
 prometheus-node-exporter:
-  ## Image and version of prometheus-node-exporter. If not provided (left commented out) default values from grafana charts will be used.
-  ##
-  # image:
-  #   repository: quay.io/prometheus/node-exporter
-  #   tag: v1.3.1
-  #   pullPolicy: IfNotPresent
-
   namespaceOverride: ""
   podLabels:
     ## Add the 'node-exporter' label to be used by serviceMonitor to match standard common usage in rules and grafana dashboards
     ##
     jobLabel: node-exporter
+  releaseLabel: true
   extraArgs:
     - --collector.filesystem.mount-points-exclude=^/(dev|proc|sys|var/lib/docker/.+|var/lib/kubelet/.+)($|/)
     - --collector.filesystem.fs-types-exclude=^(autofs|binfmt_misc|bpf|cgroup2?|configfs|debugfs|devpts|devtmpfs|fusectl|hugetlbfs|iso9660|mqueue|nsfs|overlay|proc|procfs|pstore|rpc_pipefs|securityfs|selinuxfs|squashfs|sysfs|tracefs)$
@@ -1640,7 +1611,7 @@ prometheusOperator:
       enabled: true
       image:
         repository: k8s.gcr.io/ingress-nginx/kube-webhook-certgen
-        tag: v1.2.0
+        tag: v1.3.0
         sha: ""
         pullPolicy: IfNotPresent
       resources: {}
@@ -1879,7 +1850,7 @@ prometheusOperator:
   ##
   image:
     repository: quay.io/prometheus-operator/prometheus-operator
-    tag: v0.58.0
+    tag: v0.59.1
     sha: ""
     pullPolicy: IfNotPresent
 
@@ -1894,10 +1865,9 @@ prometheusOperator:
   ## Prometheus-config-reloader
   ##
   prometheusConfigReloader:
-    # image to use for config and rule reloading
     image:
       repository: quay.io/prometheus-operator/prometheus-config-reloader
-      tag: v0.58.0
+      tag: v0.59.1
       sha: ""
 
     # resource config for prometheusConfigReloader
@@ -1913,7 +1883,7 @@ prometheusOperator:
   ##
   thanosImage:
     repository: quay.io/thanos/thanos
-    tag: v0.27.0
+    tag: v0.28.0
     sha: ""
 
   ## Set a Field Selector to filter watched secrets
@@ -2346,7 +2316,7 @@ prometheus:
     ##
     image:
       repository: quay.io/prometheus/prometheus
-      tag: v2.37.0
+      tag: v2.38.0
       sha: ""
 
     ## Tolerations for use with node taints
@@ -3177,7 +3147,7 @@ thanosRuler:
     ##
     image:
       repository: quay.io/thanos/thanos
-      tag: v0.27.0
+      tag: v0.28.0
       sha: ""
 
     ## Namespaces to be selected for PrometheusRules discovery.

```

## 39.13.3

**Release date:** 2022-09-13

![AppVersion: 0.58.0](https://img.shields.io/static/v1?label=AppVersion&message=0.58.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack]  Add support for templating in excludedFromEnforcement parameter (#2429)

### Default value changes

```diff
# No changes in this release
```

## 39.13.2

**Release date:** 2022-09-13

![AppVersion: 0.58.0](https://img.shields.io/static/v1?label=AppVersion&message=0.58.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] - Add support for templating in additionalAlertManagerConfigs (#2430)

### Default value changes

```diff
# No changes in this release
```

## 39.13.1

**Release date:** 2022-09-13

![AppVersion: 0.58.0](https://img.shields.io/static/v1?label=AppVersion&message=0.58.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Fix variable WalCompression (#2412)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index d1bcf166..8de44218 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -2538,7 +2538,7 @@ prometheus:
 
     ## Enable compression of the write-ahead log using Snappy.
     ##
-    walCompression: false
+    walCompression: true
 
     ## If true, the Operator won't process any Prometheus configuration changes
     ##

```

## 39.13.0

**Release date:** 2022-09-13

![AppVersion: 0.58.0](https://img.shields.io/static/v1?label=AppVersion&message=0.58.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] bump grafana 6.38.0 (9.1.4) (#2449)

### Default value changes

```diff
# No changes in this release
```

## 39.12.1

**Release date:** 2022-09-11

![AppVersion: 0.58.0](https://img.shields.io/static/v1?label=AppVersion&message=0.58.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Bumb grafana to version 6.37.3 (9.1.4) (#2443)

### Default value changes

```diff
# No changes in this release
```

## 39.12.0

**Release date:** 2022-09-10

![AppVersion: 0.58.0](https://img.shields.io/static/v1?label=AppVersion&message=0.58.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] bump grafana 6.37.1 (9.1.2) (#2431)

### Default value changes

```diff
# No changes in this release
```

## 39.11.0

**Release date:** 2022-08-31

![AppVersion: 0.58.0](https://img.shields.io/static/v1?label=AppVersion&message=0.58.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] SecurityContext for webhook job create/patch container (#2406)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 7d66ae90..d1bcf166 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1661,6 +1661,14 @@ prometheusOperator:
         runAsNonRoot: true
         runAsUser: 2000
 
+    # Security context for create job container
+    createSecretJob:
+      securityContext: {}
+
+      # Security context for patch job container
+    patchWebhookJob:
+      securityContext: {}
+
     # Use certmanager to generate webhook certs
     certManager:
       enabled: false

```

## 39.10.0

**Release date:** 2022-08-31

![AppVersion: 0.58.0](https://img.shields.io/static/v1?label=AppVersion&message=0.58.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Hint offline sources (#2411)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 0196a54c..7d66ae90 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -703,6 +703,14 @@ grafana:
   enabled: true
   namespaceOverride: ""
 
+  ## Image and version of grafana. If not provided (left commented out) default values from grafana charts will be used.
+  ##
+  # image:
+  #   repository: docker.io/grafana/grafana
+  #   tag: 9.0.6
+  #   sha: ""
+  #   pullPolicy: IfNotPresent
+
   ## ForceDeployDatasources Create datasource configmap even if grafana deployment has been disabled
   ##
   forceDeployDatasources: false
@@ -766,6 +774,13 @@ grafana:
     #   - grafana.example.com
 
   sidecar:
+    ## Image and version of sidecar. If not provided (left commented out) default values from grafana charts will be used.
+    ##
+    # image:
+    #   repository: quay.io/kiwigrid/k8s-sidecar
+    #   tag: 1.19.2
+    #   sha: ""
+
     dashboards:
       enabled: true
       label: grafana_dashboard
@@ -1471,6 +1486,14 @@ kubeStateMetrics:
 ## Configuration for kube-state-metrics subchart
 ##
 kube-state-metrics:
+  ##
+  ##
+  # image:
+  #  repository: registry.k8s.io/kube-state-metrics/kube-state-metrics
+  #  tag: v2.5.0
+  #  sha: ""
+  #  pullPolicy: IfNotPresent
+
   namespaceOverride: ""
   rbac:
     create: true
@@ -1525,6 +1548,13 @@ nodeExporter:
 ## Configuration for prometheus-node-exporter subchart
 ##
 prometheus-node-exporter:
+  ## Image and version of prometheus-node-exporter. If not provided (left commented out) default values from grafana charts will be used.
+  ##
+  # image:
+  #   repository: quay.io/prometheus/node-exporter
+  #   tag: v1.3.1
+  #   pullPolicy: IfNotPresent
+
   namespaceOverride: ""
   podLabels:
     ## Add the 'node-exporter' label to be used by serviceMonitor to match standard common usage in rules and grafana dashboards

```

## 39.9.0

**Release date:** 2022-08-21

![AppVersion: 0.58.0](https://img.shields.io/static/v1?label=AppVersion&message=0.58.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Setting a custom timeout in admission mutatingWebhookConfiguration - Prometheus Operator chart Template (#2233)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index fde3bcf5..0196a54c 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1596,6 +1596,8 @@ prometheusOperator:
   ## rules from making their way into prometheus and potentially preventing the container from starting
   admissionWebhooks:
     failurePolicy: Fail
+    ## The default timeoutSeconds is 10 and the maximum value is 30.
+    timeoutSeconds: 10
     enabled: true
     ## A PEM encoded CA bundle which will be used to validate the webhook's server certificate.
     ## If unspecified, system trust roots on the apiserver are used.

```

## 39.8.0

**Release date:** 2022-08-17

![AppVersion: 0.58.0](https://img.shields.io/static/v1?label=AppVersion&message=0.58.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add evaluationInterval to Thanos Ruler (#2374)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index b20b2e44..fde3bcf5 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -3188,6 +3188,10 @@ thanosRuler:
     ##
     retention: 24h
 
+    ## Interval between consecutive evaluations.
+    ##
+    evaluationInterval: ""
+
     ## Storage is the definition of how storage will be used by the ThanosRuler instances.
     ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/user-guides/storage.md
     ##

```

## 39.7.0

**Release date:** 2022-08-16

![AppVersion: 0.58.0](https://img.shields.io/static/v1?label=AppVersion&message=0.58.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Update Chart.yaml (#2371)
* [kube-prometheus-stack] add thanosRulerSpec.alertmanagersConfig config (#2358)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index e5f86f0d..b20b2e44 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -3201,6 +3201,22 @@ thanosRuler:
     #         storage: 50Gi
     #   selector: {}
 
+    ## AlertmanagerConfig define configuration for connecting to alertmanager.
+    ## Only available with Thanos v0.10.0 and higher. Maps to the alertmanagers.config Thanos Ruler arg.
+    alertmanagersConfig: {}
+    #   - api_version: v2
+    #     http_config:
+    #       basic_auth:
+    #         username: some_user
+    #         password: some_pass
+    #     static_configs:
+    #       - alertmanager.thanos.io
+    #     scheme: http
+    #     timeout: 10s
+
+    ## DEPRECATED. Define URLs to send alerts to Alertmanager. For Thanos v0.10.0 and higher, alertmanagersConfig should be used instead.
+    ## Note: this field will be ignored if alertmanagersConfig is specified. Maps to the alertmanagers.url Thanos Ruler arg.
+    # alertmanagersUrl:
 
     ## The external URL the Thanos Ruler instances will be available under. This is necessary to generate correct URLs. This is necessary if Thanos Ruler is not served from root of a DNS name. string false
     ##

```

## 39.6.0

**Release date:** 2022-08-10

![AppVersion: 0.58.0](https://img.shields.io/static/v1?label=AppVersion&message=0.58.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack]: add alertmanager-config-namespaces for prometheus-operator support (#2276)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 4aa29e7b..e5f86f0d 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1656,6 +1656,7 @@ prometheusOperator:
   ## Filter namespaces to look for prometheus-operator custom resources
   ##
   alertmanagerInstanceNamespaces: []
+  alertmanagerConfigNamespaces: []
   prometheusInstanceNamespaces: []
   thanosRulerInstanceNamespaces: []
 

```

## 39.5.0

**Release date:** 2022-08-07

![AppVersion: 0.58.0](https://img.shields.io/static/v1?label=AppVersion&message=0.58.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add an option to enable/disable `kubernetes-system-controller-manager` PrometheusRule (#2343)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index d5635ed8..4aa29e7b 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -42,6 +42,7 @@ defaultRules:
     kubeApiserverBurnrate: true
     kubeApiserverHistogram: true
     kubeApiserverSlos: true
+    kubeControllerManager: true
     kubelet: true
     kubeProxy: true
     kubePrometheusGeneral: true

```

## 39.4.1

**Release date:** 2022-08-05

![AppVersion: 0.58.0](https://img.shields.io/static/v1?label=AppVersion&message=0.58.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Support for import from jsonnet mixin for etcd (#2340)

### Default value changes

```diff
# No changes in this release
```

## 39.4.0

**Release date:** 2022-08-02

![AppVersion: 0.58.0](https://img.shields.io/static/v1?label=AppVersion&message=0.58.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add exemplars related settings (#2332)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 31a90f81..d5635ed8 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -2287,6 +2287,14 @@ prometheus:
     ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#webtlsconfig
     web: {}
 
+    ## Exemplars related settings that are runtime reloadable.
+    ## It requires to enable the exemplar storage feature to be effective.
+    exemplars: ""
+      ## Maximum number of exemplars stored in memory for all series.
+      ## If not set, Prometheus uses its default value.
+      ## A value of zero or less than zero disables the storage.
+      # maxSize: 100000
+
     # EnableFeatures API enables access to Prometheus disabled features.
     # ref: https://prometheus.io/docs/prometheus/latest/disabled_features/
     enableFeatures: []

```

## 39.3.0

**Release date:** 2022-08-02

![AppVersion: 0.58.0](https://img.shields.io/static/v1?label=AppVersion&message=0.58.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack]: Bump KSM version (#2333)

### Default value changes

```diff
# No changes in this release
```

## 39.2.1

**Release date:** 2022-08-01

![AppVersion: 0.58.0](https://img.shields.io/static/v1?label=AppVersion&message=0.58.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] fix new lines in _helpers (#2331)

### Default value changes

```diff
# No changes in this release
```

## 39.2.0

**Release date:** 2022-07-31

![AppVersion: 0.58.0](https://img.shields.io/static/v1?label=AppVersion&message=0.58.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Change AlertManager and Prometheus CR naming. (#2188)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index c1b548c4..31a90f81 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -3323,3 +3323,7 @@ thanosRuler:
   #   auth: |
   #     foo:$apr1$OFG3Xybp$ckL0FHDAkoXYIlH9.cysT0
   #     someoneelse:$apr1$DMZX2Z4q$6SbQIfyuLQd.xmo/P0m2c.
+
+## Setting to true produces cleaner resource names, but requires a data migration because the name of the persistent volume changes. Therefore this should only be set once on initial installation.
+##
+cleanPrometheusOperatorObjectNames: false

```

## 39.1.0

**Release date:** 2022-07-29

![AppVersion: 0.58.0](https://img.shields.io/static/v1?label=AppVersion&message=0.58.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] make additionalAlertManagerConfigs optional (#2322)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 53a2542c..c1b548c4 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -2693,6 +2693,7 @@ prometheus:
     additionalAlertManagerConfigsSecret: {}
       # name:
       # key:
+      # optional: false
 
     ## AdditionalAlertRelabelConfigs allows specifying Prometheus alert relabel configurations. Alert relabel configurations specified are appended
     ## to the configurations generated by the Prometheus Operator. Alert relabel configurations specified must have the form as specified in the

```

## 39.0.0

**Release date:** 2022-07-29

![AppVersion: 0.58.0](https://img.shields.io/static/v1?label=AppVersion&message=0.58.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* update prometheus operator to v0.58.0 (#2320)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index d17f00eb..53a2542c 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -644,7 +644,7 @@ alertmanager:
     containers: []
     # containers:
     # - name: oauth-proxy
-    #   image: quay.io/oauth2-proxy/oauth2-proxy:v7.1.2
+    #   image: quay.io/oauth2-proxy/oauth2-proxy:v7.3.0
     #   args:
     #   - --upstream=http://127.0.0.1:9093
     #   - --http-address=0.0.0.0:8081
@@ -1607,7 +1607,7 @@ prometheusOperator:
       enabled: true
       image:
         repository: k8s.gcr.io/ingress-nginx/kube-webhook-certgen
-        tag: v1.1.1
+        tag: v1.2.0
         sha: ""
         pullPolicy: IfNotPresent
       resources: {}
@@ -1837,7 +1837,7 @@ prometheusOperator:
   ##
   image:
     repository: quay.io/prometheus-operator/prometheus-operator
-    tag: v0.57.0
+    tag: v0.58.0
     sha: ""
     pullPolicy: IfNotPresent
 
@@ -1855,7 +1855,7 @@ prometheusOperator:
     # image to use for config and rule reloading
     image:
       repository: quay.io/prometheus-operator/prometheus-config-reloader
-      tag: v0.57.0
+      tag: v0.58.0
       sha: ""
 
     # resource config for prometheusConfigReloader
@@ -1871,7 +1871,7 @@ prometheusOperator:
   ##
   thanosImage:
     repository: quay.io/thanos/thanos
-    tag: v0.25.2
+    tag: v0.27.0
     sha: ""
 
   ## Set a Field Selector to filter watched secrets
@@ -2296,7 +2296,7 @@ prometheus:
     ##
     image:
       repository: quay.io/prometheus/prometheus
-      tag: v2.36.1
+      tag: v2.37.0
       sha: ""
 
     ## Tolerations for use with node taints
@@ -2748,7 +2748,7 @@ prometheus:
     containers: []
     # containers:
     # - name: oauth-proxy
-    #   image: quay.io/oauth2-proxy/oauth2-proxy:v7.1.2
+    #   image: quay.io/oauth2-proxy/oauth2-proxy:v7.3.0
     #   args:
     #   - --upstream=http://127.0.0.1:9093
     #   - --http-address=0.0.0.0:8081
@@ -3126,7 +3126,7 @@ thanosRuler:
     ##
     image:
       repository: quay.io/thanos/thanos
-      tag: v0.24.0
+      tag: v0.27.0
       sha: ""
 
     ## Namespaces to be selected for PrometheusRules discovery.

```

## 38.0.3

**Release date:** 2022-07-28

![AppVersion: 0.57.0](https://img.shields.io/static/v1?label=AppVersion&message=0.57.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack]: bump grafana version (#2283)

### Default value changes

```diff
# No changes in this release
```

## 38.0.2

**Release date:** 2022-07-25

![AppVersion: 0.57.0](https://img.shields.io/static/v1?label=AppVersion&message=0.57.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Bump grafana to 6.32.7 (#2302)

### Default value changes

```diff
# No changes in this release
```

## 38.0.1

**Release date:** 2022-07-24

![AppVersion: 0.57.0](https://img.shields.io/static/v1?label=AppVersion&message=0.57.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Removing me as maintainer (#2300)

### Default value changes

```diff
# No changes in this release
```

## 38.0.0

**Release date:** 2022-07-22

![AppVersion: 0.57.0](https://img.shields.io/static/v1?label=AppVersion&message=0.57.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Remove a cAdvisor metric relabelling (#2297)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 64b32c36..d17f00eb 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -991,10 +991,6 @@ kubelet:
       - sourceLabels: [id, pod]
         action: drop
         regex: '.+;'
-      # Drop cgroup metrics with no container.
-      - sourceLabels: [id, container]
-        action: drop
-        regex: '.+;'
     # - sourceLabels: [__name__, image]
     #   separator: ;
     #   regex: container_([a-z_]+);

```

## 37.3.0

**Release date:** 2022-07-20

![AppVersion: 0.57.0](https://img.shields.io/static/v1?label=AppVersion&message=0.57.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* chore(kube-prometheus-stack/exporters): support additional labels on servicemonitors (#2219)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 1a3b6cb1..64b32c36 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -923,6 +923,11 @@ kubeApiServer:
     # - targetLabel: __address__
     #   replacement: kubernetes.default.svc:443
 
+    ## Additional labels
+    ##
+    additionalLabels: {}
+    #  foo: bar
+
 ## Component scraping the kubelet and kubelet-hosted cAdvisor
 ##
 kubelet:
@@ -1085,6 +1090,11 @@ kubelet:
     #   replacement: $1
     #   action: replace
 
+    ## Additional labels
+    ##
+    additionalLabels: {}
+    #  foo: bar
+
 ## Component scraping the kube controller manager
 ##
 kubeControllerManager:
@@ -1150,6 +1160,11 @@ kubeControllerManager:
     #   replacement: $1
     #   action: replace
 
+    ## Additional labels
+    ##
+    additionalLabels: {}
+    #  foo: bar
+
 ## Component scraping coreDns. Use either this or kubeDns
 ##
 coreDns:
@@ -1187,6 +1202,11 @@ coreDns:
     #   replacement: $1
     #   action: replace
 
+    ## Additional labels
+    ##
+    additionalLabels: {}
+    #  foo: bar
+
 ## Component scraping kubeDns. Use either this or coreDns
 ##
 kubeDns:
@@ -1247,6 +1267,11 @@ kubeDns:
     #   replacement: $1
     #   action: replace
 
+    ## Additional labels
+    ##
+    additionalLabels: {}
+    #  foo: bar
+
 ## Component scraping etcd
 ##
 kubeEtcd:
@@ -1313,6 +1338,10 @@ kubeEtcd:
     #   replacement: $1
     #   action: replace
 
+    ## Additional labels
+    ##
+    additionalLabels: {}
+    #  foo: bar
 
 ## Component scraping kube scheduler
 ##
@@ -1377,6 +1406,10 @@ kubeScheduler:
     #   replacement: $1
     #   action: replace
 
+    ## Additional labels
+    ##
+    additionalLabels: {}
+    #  foo: bar
 
 ## Component scraping kube proxy
 ##
@@ -1428,6 +1461,10 @@ kubeProxy:
     #   regex: 'kube_(daemonset|deployment|pod|namespace|node|statefulset).+'
     #   sourceLabels: [__name__]
 
+    ## Additional labels
+    ##
+    additionalLabels: {}
+    #  foo: bar
 
 ## Component scraping kube state metrics
 ##

```

## 37.2.0

**Release date:** 2022-07-13

![AppVersion: 0.57.0](https://img.shields.io/static/v1?label=AppVersion&message=0.57.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* feat: add labels spec for thanos-ruler (#2270)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 6a0ab1c4..1a3b6cb1 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -3175,6 +3175,11 @@ thanosRuler:
     ## When used alongside with ObjectStorageConfig, ObjectStorageConfigFile takes precedence.
     objectStorageConfigFile: ""
 
+    ## Labels configure the external label pairs to ThanosRuler. A default replica
+    ## label `thanos_ruler_replica` will be always added as a label with the value
+    ## of the pod's name and it will be dropped in the alerts.
+    labels: {}
+
     ## If set to true all actions on the underlying managed objects are not going to be performed, except for delete actions.
     ##
     paused: false

```

## 37.1.0

**Release date:** 2022-07-13

![AppVersion: 0.57.0](https://img.shields.io/static/v1?label=AppVersion&message=0.57.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Bump kps chart version (#2268)

### Default value changes

```diff
# No changes in this release
```

## 37.0.0

**Release date:** 2022-07-12

![AppVersion: 0.57.0](https://img.shields.io/static/v1?label=AppVersion&message=0.57.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Update metric relabel values (#2197)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index e129574f..6a0ab1c4 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -899,7 +899,13 @@ kubeApiServer:
     ## MetricRelabelConfigs to apply to samples after scraping, but before ingestion.
     ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#relabelconfig
     ##
-    metricRelabelings: []
+    metricRelabelings:
+      # Drop excessively noisy apiserver buckets.
+      - action: drop
+        regex: apiserver_request_duration_seconds_bucket;(0.15|0.2|0.3|0.35|0.4|0.45|0.6|0.7|0.8|0.9|1.25|1.5|1.75|2|3|3.5|4|4.5|6|7|8|9|15|25|40|50)
+        sourceLabels:
+          - __name__
+          - le
     # - action: keep
     #   regex: 'kube_(daemonset|deployment|pod|namespace|node|statefulset).+'
     #   sourceLabels: [__name__]
@@ -955,7 +961,35 @@ kubelet:
     ## MetricRelabelConfigs to apply to samples after scraping, but before ingestion.
     ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#relabelconfig
     ##
-    cAdvisorMetricRelabelings: []
+    cAdvisorMetricRelabelings:
+      # Drop less useful container CPU metrics.
+      - sourceLabels: [__name__]
+        action: drop
+        regex: 'container_cpu_(cfs_throttled_seconds_total|load_average_10s|system_seconds_total|user_seconds_total)'
+      # Drop less useful container / always zero filesystem metrics.
+      - sourceLabels: [__name__]
+        action: drop
+        regex: 'container_fs_(io_current|io_time_seconds_total|io_time_weighted_seconds_total|reads_merged_total|sector_reads_total|sector_writes_total|writes_merged_total)'
+      # Drop less useful / always zero container memory metrics.
+      - sourceLabels: [__name__]
+        action: drop
+        regex: 'container_memory_(mapped_file|swap)'
+      # Drop less useful container process metrics.
+      - sourceLabels: [__name__]
+        action: drop
+        regex: 'container_(file_descriptors|tasks_state|threads_max)'
+      # Drop container spec metrics that overlap with kube-state-metrics.
+      - sourceLabels: [__name__]
+        action: drop
+        regex: 'container_spec.*'
+      # Drop cgroup metrics with no pod.
+      - sourceLabels: [id, pod]
+        action: drop
+        regex: '.+;'
+      # Drop cgroup metrics with no container.
+      - sourceLabels: [id, container]
+        action: drop
+        regex: '.+;'
     # - sourceLabels: [__name__, image]
     #   separator: ;
     #   regex: container_([a-z_]+);

```

## 36.6.2

**Release date:** 2022-07-11

![AppVersion: 0.57.0](https://img.shields.io/static/v1?label=AppVersion&message=0.57.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] introduced servicePort field to select ingress target for alertmanager/prometheus (#2163)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 74688edc..e129574f 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -255,6 +255,9 @@ alertmanager:
 
     labels: {}
 
+    ## Redirect ingress to an additional defined port on the service
+    # servicePort: 8081
+
     ## Hosts must be provided if Ingress is enabled.
     ##
     hosts: []
@@ -349,6 +352,10 @@ alertmanager:
 
     ## Additional ports to open for Alertmanager service
     additionalPorts: []
+    # additionalPorts:
+    # - name: authenticated
+    #   port: 8081
+    #   targetPort: 8081
 
     externalIPs: []
     loadBalancerIP: ""
@@ -635,6 +642,18 @@ alertmanager:
     ## Containers allows injecting additional containers. This is meant to allow adding an authentication proxy to an Alertmanager pod.
     ##
     containers: []
+    # containers:
+    # - name: oauth-proxy
+    #   image: quay.io/oauth2-proxy/oauth2-proxy:v7.1.2
+    #   args:
+    #   - --upstream=http://127.0.0.1:9093
+    #   - --http-address=0.0.0.0:8081
+    #   - ...
+    #   ports:
+    #   - containerPort: 8081
+    #     name: oauth-proxy
+    #     protocol: TCP
+    #   resources: {}
 
     # Additional volumes on the output StatefulSet definition.
     volumes: []
@@ -1939,6 +1958,10 @@ prometheus:
 
     ## Additional port to define in the Service
     additionalPorts: []
+    # additionalPorts:
+    # - name: authenticated
+    #   port: 8081
+    #   targetPort: 8081
 
     ## Consider that all endpoints are considered "ready" even if the Pods themselves are not
     ## Ref: https://kubernetes.io/docs/reference/kubernetes-api/service-resources/service-v1/#ServiceSpec
@@ -2047,6 +2070,9 @@ prometheus:
     annotations: {}
     labels: {}
 
+    ## Redirect ingress to an additional defined port on the service
+    # servicePort: 8081
+
     ## Hostnames.
     ## Must be provided if Ingress is enabled.
     ##
@@ -2653,6 +2679,18 @@ prometheus:
     ## Containers allows injecting additional containers. This is meant to allow adding an authentication proxy to a Prometheus pod.
     ## if using proxy extraContainer update targetPort with proxy container port
     containers: []
+    # containers:
+    # - name: oauth-proxy
+    #   image: quay.io/oauth2-proxy/oauth2-proxy:v7.1.2
+    #   args:
+    #   - --upstream=http://127.0.0.1:9093
+    #   - --http-address=0.0.0.0:8081
+    #   - ...
+    #   ports:
+    #   - containerPort: 8081
+    #     name: oauth-proxy
+    #     protocol: TCP
+    #   resources: {}
 
     ## InitContainers allows injecting additional initContainers. This is meant to allow doing some changes
     ## (permissions, dir tree) on mounted volumes before starting prometheus

```

## 36.6.1

**Release date:** 2022-07-08

![AppVersion: 0.57.0](https://img.shields.io/static/v1?label=AppVersion&message=0.57.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Fix field evolution error (#2256)

### Default value changes

```diff
# No changes in this release
```

## 36.6.0

**Release date:** 2022-07-08

![AppVersion: 0.57.0](https://img.shields.io/static/v1?label=AppVersion&message=0.57.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack-operator] adds deployment annotations for prometheus operator (#2246)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 029a2fdd..74688edc 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1624,6 +1624,10 @@ prometheusOperator:
     ##
     externalIPs: []
 
+  ## Annotations to add to the operator deployment
+  ##
+  annotations: {}
+
   ## Labels to add to the operator pod
   ##
   podLabels: {}

```

## 36.5.1

**Release date:** 2022-07-08

![AppVersion: 0.57.0](https://img.shields.io/static/v1?label=AppVersion&message=0.57.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] fixes #2254  (#2255)

### Default value changes

```diff
# No changes in this release
```

## 36.5.0

**Release date:** 2022-07-08

![AppVersion: 0.57.0](https://img.shields.io/static/v1?label=AppVersion&message=0.57.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add exemplar support for Grafana datasource (#2237)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 96ac5035..029a2fdd 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -786,6 +786,12 @@ grafana:
       label: grafana_datasource
       labelValue: "1"
 
+      ## Field with internal link pointing to existing data source in Grafana.
+      ## Can be provisioned via additionalDataSources
+      exemplarTraceIdDestinations: {}
+        # datasourceUid: Jaeger
+        # traceIdLabelName: trace_id
+
   extraConfigmapMounts: []
   # - name: certs-configmap
   #   mountPath: /etc/grafana/ssl/

```

## 36.4.0

**Release date:** 2022-07-08

![AppVersion: 0.57.0](https://img.shields.io/static/v1?label=AppVersion&message=0.57.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Add objectStorageConfig for thanos-ruler (#2232)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index d9b67770..96ac5035 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -3085,6 +3085,14 @@ thanosRuler:
     ##
     routePrefix: /
 
+    ## ObjectStorageConfig configures object storage in Thanos. Alternative to
+    ## ObjectStorageConfigFile, and lower order priority.
+    objectStorageConfig: {}
+
+    ## ObjectStorageConfigFile specifies the path of the object storage configuration file.
+    ## When used alongside with ObjectStorageConfig, ObjectStorageConfigFile takes precedence.
+    objectStorageConfigFile: ""
+
     ## If set to true all actions on the underlying managed objects are not going to be performed, except for delete actions.
     ##
     paused: false

```

## 36.3.1

**Release date:** 2022-07-08

![AppVersion: 0.57.0](https://img.shields.io/static/v1?label=AppVersion&message=0.57.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] fix the thanos discovery servicemonitor selector (#2244)

### Default value changes

```diff
# No changes in this release
```

## 36.3.0

**Release date:** 2022-07-08

![AppVersion: 0.57.0](https://img.shields.io/static/v1?label=AppVersion&message=0.57.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add web property to AlertmanagerSpec (#2242)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 874ea167..d9b67770 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -468,6 +468,10 @@ alertmanager:
     ##
     # configSecret:
 
+    ## WebTLSConfig defines the TLS parameters for HTTPS
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#alertmanagerwebspec
+    web: {}
+
     ## AlertmanagerConfigs to be selected to merge and configure Alertmanager with.
     ##
     alertmanagerConfigSelector: {}

```

## 36.2.1

**Release date:** 2022-07-01

![AppVersion: 0.57.0](https://img.shields.io/static/v1?label=AppVersion&message=0.57.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] fix: rename thanos sidecar servicemonitor (#2216)

### Default value changes

```diff
# No changes in this release
```

## 36.2.0

**Release date:** 2022-06-23

![AppVersion: 0.57.0](https://img.shields.io/static/v1?label=AppVersion&message=0.57.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Bump Grafana chart to 6.31.* (#2184)

### Default value changes

```diff
# No changes in this release
```

## 36.1.0

**Release date:** 2022-06-23

![AppVersion: 0.57.0](https://img.shields.io/static/v1?label=AppVersion&message=0.57.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add thanosRuler configuration (#1993)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index c2ad8e94..874ea167 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -2853,3 +2853,340 @@ prometheus:
     ## https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#podmetricsendpoint
     ##
     # podMetricsEndpoints: []
+
+## Configuration for thanosRuler
+## ref: https://thanos.io/tip/components/rule.md/
+##
+thanosRuler:
+
+  ## Deploy thanosRuler
+  ##
+  enabled: false
+
+  ## Annotations for ThanosRuler
+  ##
+  annotations: {}
+
+  ## Service account for ThanosRuler to use.
+  ## ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
+  ##
+  serviceAccount:
+    create: true
+    name: ""
+    annotations: {}
+
+  ## Configure pod disruption budgets for ThanosRuler
+  ## ref: https://kubernetes.io/docs/tasks/run-application/configure-pdb/#specifying-a-poddisruptionbudget
+  ## This configuration is immutable once created and will require the PDB to be deleted to be changed
+  ## https://github.com/kubernetes/kubernetes/issues/45398
+  ##
+  podDisruptionBudget:
+    enabled: false
+    minAvailable: 1
+    maxUnavailable: ""
+
+  ingress:
+    enabled: false
+
+    # For Kubernetes >= 1.18 you should specify the ingress-controller via the field ingressClassName
+    # See https://kubernetes.io/blog/2020/04/02/improvements-to-the-ingress-api-in-kubernetes-1.18/#specifying-the-class-of-an-ingress
+    # ingressClassName: nginx
+
+    annotations: {}
+
+    labels: {}
+
+    ## Hosts must be provided if Ingress is enabled.
+    ##
+    hosts: []
+      # - thanosruler.domain.com
+
+    ## Paths to use for ingress rules - one path should match the thanosruler.routePrefix
+    ##
+    paths: []
+    # - /
+
+    ## For Kubernetes >= 1.18 you should specify the pathType (determines how Ingress paths should be matched)
+    ## See https://kubernetes.io/blog/2020/04/02/improvements-to-the-ingress-api-in-kubernetes-1.18/#better-path-matching-with-path-types
+    # pathType: ImplementationSpecific
+
+    ## TLS configuration for ThanosRuler Ingress
+    ## Secret must be manually created in the namespace
+    ##
+    tls: []
+    # - secretName: thanosruler-general-tls
+    #   hosts:
+    #   - thanosruler.example.com
+
+  ## Configuration for ThanosRuler service
+  ##
+  service:
+    annotations: {}
+    labels: {}
+    clusterIP: ""
+
+    ## Port for ThanosRuler Service to listen on
+    ##
+    port: 10902
+    ## To be used with a proxy extraContainer port
+    ##
+    targetPort: 10902
+    ## Port to expose on each node
+    ## Only used if service.type is 'NodePort'
+    ##
+    nodePort: 30905
+    ## List of IP addresses at which the Prometheus server service is available
+    ## Ref: https://kubernetes.io/docs/user-guide/services/#external-ips
+    ##
+
+    ## Additional ports to open for ThanosRuler service
+    additionalPorts: []
+
+    externalIPs: []
+    loadBalancerIP: ""
+    loadBalancerSourceRanges: []
+
+    ## Denotes if this Service desires to route external traffic to node-local or cluster-wide endpoints
+    ##
+    externalTrafficPolicy: Cluster
+
+    ## Service type
+    ##
+    type: ClusterIP
+
+  ## If true, create a serviceMonitor for thanosRuler
+  ##
+  serviceMonitor:
+    ## Scrape interval. If not set, the Prometheus default scrape interval is used.
+    ##
+    interval: ""
+    selfMonitor: true
+
+    ## proxyUrl: URL of a proxy that should be used for scraping.
+    ##
+    proxyUrl: ""
+
+    ## scheme: HTTP scheme to use for scraping. Can be used with `tlsConfig` for example if using istio mTLS.
+    scheme: ""
+
+    ## tlsConfig: TLS configuration to use when scraping the endpoint. For example if using istio mTLS.
+    ## Of type: https://github.com/coreos/prometheus-operator/blob/main/Documentation/api.md#tlsconfig
+    tlsConfig: {}
+
+    bearerTokenFile:
+
+    ## MetricRelabelConfigs to apply to samples after scraping, but before ingestion.
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#relabelconfig
+    ##
+    metricRelabelings: []
+    # - action: keep
+    #   regex: 'kube_(daemonset|deployment|pod|namespace|node|statefulset).+'
+    #   sourceLabels: [__name__]
+
+    ## RelabelConfigs to apply to samples before scraping
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#relabelconfig
+    ##
+    relabelings: []
+    # - sourceLabels: [__meta_kubernetes_pod_node_name]
+    #   separator: ;
+    #   regex: ^(.*)$
+    #   targetLabel: nodename
+    #   replacement: $1
+    #   action: replace
+
+  ## Settings affecting thanosRulerpec
+  ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#thanosrulerspec
+  ##
+  thanosRulerSpec:
+    ## Standard object's metadata. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata
+    ## Metadata Labels and Annotations gets propagated to the ThanosRuler pods.
+    ##
+    podMetadata: {}
+
+    ## Image of ThanosRuler
+    ##
+    image:
+      repository: quay.io/thanos/thanos
+      tag: v0.24.0
+      sha: ""
+
+    ## Namespaces to be selected for PrometheusRules discovery.
+    ## If nil, select own namespace. Namespaces to be selected for ServiceMonitor discovery.
+    ## See https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#namespaceselector for usage
+    ##
+    ruleNamespaceSelector: {}
+
+    ## If true, a nil or {} value for thanosRuler.thanosRulerSpec.ruleSelector will cause the
+    ## prometheus resource to be created with selectors based on values in the helm deployment,
+    ## which will also match the PrometheusRule resources created
+    ##
+    ruleSelectorNilUsesHelmValues: true
+
+    ## PrometheusRules to be selected for target discovery.
+    ## If {}, select all PrometheusRules
+    ##
+    ruleSelector: {}
+    ## Example which select all PrometheusRules resources
+    ## with label "prometheus" with values any of "example-rules" or "example-rules-2"
+    # ruleSelector:
+    #   matchExpressions:
+    #     - key: prometheus
+    #       operator: In
+    #       values:
+    #         - example-rules
+    #         - example-rules-2
+    #
+    ## Example which select all PrometheusRules resources with label "role" set to "example-rules"
+    # ruleSelector:
+    #   matchLabels:
+    #     role: example-rules
+
+    ## Define Log Format
+    # Use logfmt (default) or json logging
+    logFormat: logfmt
+
+    ## Log level for ThanosRuler to be configured with.
+    ##
+    logLevel: info
+
+    ## Size is the expected size of the thanosRuler cluster. The controller will eventually make the size of the
+    ## running cluster equal to the expected size.
+    replicas: 1
+
+    ## Time duration ThanosRuler shall retain data for. Default is '24h', and must match the regular expression
+    ## [0-9]+(ms|s|m|h) (milliseconds seconds minutes hours).
+    ##
+    retention: 24h
+
+    ## Storage is the definition of how storage will be used by the ThanosRuler instances.
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/user-guides/storage.md
+    ##
+    storage: {}
+    # volumeClaimTemplate:
+    #   spec:
+    #     storageClassName: gluster
+    #     accessModes: ["ReadWriteOnce"]
+    #     resources:
+    #       requests:
+    #         storage: 50Gi
+    #   selector: {}
+
+
+    ## The external URL the Thanos Ruler instances will be available under. This is necessary to generate correct URLs. This is necessary if Thanos Ruler is not served from root of a DNS name. string false
+    ##
+    externalPrefix:
+
+    ## The route prefix ThanosRuler registers HTTP handlers for. This is useful, if using ExternalURL and a proxy is rewriting HTTP routes of a request, and the actual ExternalURL is still true,
+    ## but the server serves requests under a different route prefix. For example for use with kubectl proxy.
+    ##
+    routePrefix: /
+
+    ## If set to true all actions on the underlying managed objects are not going to be performed, except for delete actions.
+    ##
+    paused: false
+
+    ## Define which Nodes the Pods are scheduled on.
+    ## ref: https://kubernetes.io/docs/user-guide/node-selection/
+    ##
+    nodeSelector: {}
+
+    ## Define resources requests and limits for single Pods.
+    ## ref: https://kubernetes.io/docs/user-guide/compute-resources/
+    ##
+    resources: {}
+    # requests:
+    #   memory: 400Mi
+
+    ## Pod anti-affinity can prevent the scheduler from placing Prometheus replicas on the same node.
+    ## The default value "soft" means that the scheduler should *prefer* to not schedule two replica pods onto the same node but no guarantee is provided.
+    ## The value "hard" means that the scheduler is *required* to not schedule two replica pods onto the same node.
+    ## The value "" will disable pod anti-affinity so that no anti-affinity rules will be configured.
+    ##
+    podAntiAffinity: ""
+
+    ## If anti-affinity is enabled sets the topologyKey to use for anti-affinity.
+    ## This can be changed to, for example, failure-domain.beta.kubernetes.io/zone
+    ##
+    podAntiAffinityTopologyKey: kubernetes.io/hostname
+
+    ## Assign custom affinity rules to the thanosRuler instance
+    ## ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
+    ##
+    affinity: {}
+    # nodeAffinity:
+    #   requiredDuringSchedulingIgnoredDuringExecution:
+    #     nodeSelectorTerms:
+    #     - matchExpressions:
+    #       - key: kubernetes.io/e2e-az-name
+    #         operator: In
+    #         values:
+    #         - e2e-az1
+    #         - e2e-az2
+
+    ## If specified, the pod's tolerations.
+    ## ref: https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/
+    ##
+    tolerations: []
+    # - key: "key"
+    #   operator: "Equal"
+    #   value: "value"
+    #   effect: "NoSchedule"
+
+    ## If specified, the pod's topology spread constraints.
+    ## ref: https://kubernetes.io/docs/concepts/workloads/pods/pod-topology-spread-constraints/
+    ##
+    topologySpreadConstraints: []
+    # - maxSkew: 1
+    #   topologyKey: topology.kubernetes.io/zone
+    #   whenUnsatisfiable: DoNotSchedule
+    #   labelSelector:
+    #     matchLabels:
+    #       app: thanos-ruler
+
+    ## SecurityContext holds pod-level security attributes and common container settings.
+    ## This defaults to non root user with uid 1000 and gid 2000. *v1.PodSecurityContext  false
+    ## ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
+    ##
+    securityContext:
+      runAsGroup: 2000
+      runAsNonRoot: true
+      runAsUser: 1000
+      fsGroup: 2000
+
+    ## ListenLocal makes the ThanosRuler server listen on loopback, so that it does not bind against the Pod IP.
+    ## Note this is only for the ThanosRuler UI, not the gossip communication.
+    ##
+    listenLocal: false
+
+    ## Containers allows injecting additional containers. This is meant to allow adding an authentication proxy to an ThanosRuler pod.
+    ##
+    containers: []
+
+    # Additional volumes on the output StatefulSet definition.
+    volumes: []
+
+    # Additional VolumeMounts on the output StatefulSet definition.
+    volumeMounts: []
+
+    ## InitContainers allows injecting additional initContainers. This is meant to allow doing some changes
+    ## (permissions, dir tree) on mounted volumes before starting prometheus
+    initContainers: []
+
+    ## Priority class assigned to the Pods
+    ##
+    priorityClassName: ""
+
+    ## PortName to use for ThanosRuler.
+    ##
+    portName: "web"
+
+  ## ExtraSecret can be used to store various data in an extra secret
+  ## (use it for example to store hashed basic auth credentials)
+  extraSecret:
+    ## if not set, name will be auto generated
+    # name: ""
+    annotations: {}
+    data: {}
+  #   auth: |
+  #     foo:$apr1$OFG3Xybp$ckL0FHDAkoXYIlH9.cysT0
+  #     someoneelse:$apr1$DMZX2Z4q$6SbQIfyuLQd.xmo/P0m2c.

```

## 36.0.3

**Release date:** 2022-06-20

![AppVersion: 0.57.0](https://img.shields.io/static/v1?label=AppVersion&message=0.57.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* chore(kube-prometheus-stack): update kube-state-metrics and grafana to latest patch version (#2175)

### Default value changes

```diff
# No changes in this release
```

## 36.0.2

**Release date:** 2022-06-15

![AppVersion: 0.57.0](https://img.shields.io/static/v1?label=AppVersion&message=0.57.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Fix dynamic port assignment for exporters' Endpoints (#2162)

### Default value changes

```diff
# No changes in this release
```

## 36.0.1

**Release date:** 2022-06-13

![AppVersion: 0.57.0](https://img.shields.io/static/v1?label=AppVersion&message=0.57.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* fix: fixed remote link to prometheus-operator doc (#2150)

### Default value changes

```diff
# No changes in this release
```

## 36.0.0

**Release date:** 2022-06-11

![AppVersion: 0.57.0](https://img.shields.io/static/v1?label=AppVersion&message=0.57.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Version bump to 0.57.0 (#2120)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 259a3eed..c2ad8e94 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1737,7 +1737,7 @@ prometheusOperator:
   ##
   image:
     repository: quay.io/prometheus-operator/prometheus-operator
-    tag: v0.56.3
+    tag: v0.57.0
     sha: ""
     pullPolicy: IfNotPresent
 
@@ -1755,7 +1755,7 @@ prometheusOperator:
     # image to use for config and rule reloading
     image:
       repository: quay.io/prometheus-operator/prometheus-config-reloader
-      tag: v0.56.3
+      tag: v0.57.0
       sha: ""
 
     # resource config for prometheusConfigReloader
@@ -2189,7 +2189,7 @@ prometheus:
     ##
     image:
       repository: quay.io/prometheus/prometheus
-      tag: v2.35.0
+      tag: v2.36.1
       sha: ""
 
     ## Tolerations for use with node taints

```

## 35.6.2

**Release date:** 2022-06-11

![AppVersion: 0.56.3](https://img.shields.io/static/v1?label=AppVersion&message=0.56.3&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Add inhibit rules to default Alertmanager config (#2115)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 3c76822c..259a3eed 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -170,6 +170,27 @@ alertmanager:
   config:
     global:
       resolve_timeout: 5m
+    inhibit_rules:
+      - source_matchers:
+          - 'severity = critical'
+        target_matchers:
+          - 'severity =~ warning|info'
+        equal:
+          - 'namespace'
+          - 'alertname'
+      - source_matchers:
+          - 'severity = warning'
+        target_matchers:
+          - 'severity = info'
+        equal:
+          - 'namespace'
+          - 'alertname'
+      - source_matchers:
+          - 'alertname = InfoInhibitor'
+        target_matchers:
+          - 'severity = info'
+        equal:
+          - 'namespace'
     route:
       group_by: ['namespace']
       group_wait: 30s

```

## 35.6.1

**Release date:** 2022-06-11

![AppVersion: 0.56.3](https://img.shields.io/static/v1?label=AppVersion&message=0.56.3&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Change default grouping for Alertmanager to 'namespace' (#2114)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 690f6dd1..3c76822c 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -171,7 +171,7 @@ alertmanager:
     global:
       resolve_timeout: 5m
     route:
-      group_by: ['job']
+      group_by: ['namespace']
       group_wait: 30s
       group_interval: 5m
       repeat_interval: 12h

```

## 35.6.0

**Release date:** 2022-06-11

![AppVersion: 0.56.3](https://img.shields.io/static/v1?label=AppVersion&message=0.56.3&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack ] Add Custom annotations to alerts (#2026)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 464e46eb..690f6dd1 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -70,6 +70,9 @@ defaultRules:
   ## Additional labels for PrometheusRule alerts
   additionalRuleLabels: {}
 
+  ## Additional annotations for PrometheusRule alerts
+  additionalRuleAnnotations: {}
+
   ## Prefix for runbook URLs. Use this to override the first part of the runbookURLs that is common to all rules.
   runbookUrl: "https://runbooks.prometheus-operator.dev/runbooks"
 

```

## 35.5.4

**Release date:** 2022-06-11

![AppVersion: 0.56.3](https://img.shields.io/static/v1?label=AppVersion&message=0.56.3&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] fix ServiceMonitor for kubelet over http (#2072)

### Default value changes

```diff
# No changes in this release
```

## 35.5.3

**Release date:** 2022-06-10

![AppVersion: 0.56.3](https://img.shields.io/static/v1?label=AppVersion&message=0.56.3&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] fix typo (#2117)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 6279ef51..464e46eb 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -515,7 +515,7 @@ alertmanager:
     #     resources:
     #       requests:
     #         storage: 50Gi
-    #   selector: {}
+    #     selector: {}
 
 
     ## The external URL the Alertmanager instances will be available under. This is necessary to generate correct URLs. This is necessary if Alertmanager is not served from root of a DNS name. string  false

```

## 35.5.2

**Release date:** 2022-06-10

![AppVersion: 0.56.3](https://img.shields.io/static/v1?label=AppVersion&message=0.56.3&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add support for setting excludedFromEnforce in PrometheusSpec ADDENDUM (#2110)

### Default value changes

```diff
# No changes in this release
```

## 35.5.1

**Release date:** 2022-06-01

![AppVersion: 0.56.3](https://img.shields.io/static/v1?label=AppVersion&message=0.56.3&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Fix alertmanagerConfiguration comment (#2102)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 0ce0b53a..6279ef51 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -483,7 +483,9 @@ alertmanager:
     ## AlermanagerConfig to be used as top level configuration
     ##
     alertmanagerConfiguration: {}
-    # - name: global-alertmanager-Configuration
+    ## Example with select a global alertmanagerconfig
+    # alertmanagerConfiguration:
+    #   name: global-alertmanager-Configuration
 
     ## Define Log Format
     # Use logfmt (default) or json logging

```

## 35.5.0

**Release date:** 2022-05-31

![AppVersion: 0.56.3](https://img.shields.io/static/v1?label=AppVersion&message=0.56.3&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add support for setting excludedFromEnforce in PrometheusSpec (#2089)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 37336fa8..0ce0b53a 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1711,7 +1711,7 @@ prometheusOperator:
   ##
   image:
     repository: quay.io/prometheus-operator/prometheus-operator
-    tag: v0.56.2
+    tag: v0.56.3
     sha: ""
     pullPolicy: IfNotPresent
 
@@ -1729,7 +1729,7 @@ prometheusOperator:
     # image to use for config and rule reloading
     image:
       repository: quay.io/prometheus-operator/prometheus-config-reloader
-      tag: v0.56.2
+      tag: v0.56.3
       sha: ""
 
     # resource config for prometheusConfigReloader
@@ -2644,8 +2644,15 @@ prometheus:
 
     ## PrometheusRulesExcludedFromEnforce - list of prometheus rules to be excluded from enforcing of adding namespace labels.
     ## Works only if enforcedNamespaceLabel set to true. Make sure both ruleNamespace and ruleName are set for each pair
+    ## Deprecated, use `excludedFromEnforcement` instead
     prometheusRulesExcludedFromEnforce: []
 
+    ## ExcludedFromEnforcement - list of object references to PodMonitor, ServiceMonitor, Probe and PrometheusRule objects
+    ## to be excluded from enforcing a namespace label of origin.
+    ## Works only if enforcedNamespaceLabel set to true.
+    ## See https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#objectreference
+    excludedFromEnforcement: []
+
     ## QueryLogFile specifies the file to which PromQL queries are logged. Note that this location must be writable,
     ## and can be persisted using an attached volume. Alternatively, the location can be set to a stdout location such
     ## as /dev/stdout to log querie information to the default Prometheus log stream. This is only available in versions

```

## 35.4.2

**Release date:** 2022-05-27

![AppVersion: 0.56.2](https://img.shields.io/static/v1?label=AppVersion&message=0.56.2&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] remove deprecated kube apiserver rules (#2076)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 7414180e..37336fa8 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -38,10 +38,10 @@ defaultRules:
     configReloaders: true
     general: true
     k8s: true
-    kubeApiserver: true
     kubeApiserverAvailability: true
-    kubeApiserverSlos: true
     kubeApiserverBurnrate: true
+    kubeApiserverHistogram: true
+    kubeApiserverSlos: true
     kubelet: true
     kubeProxy: true
     kubePrometheusGeneral: true

```

## 35.4.1

**Release date:** 2022-05-27

![AppVersion: 0.56.2](https://img.shields.io/static/v1?label=AppVersion&message=0.56.2&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] add missing defaultRules.rules.kubeApiserverBurnrate (#2086)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index e613f9fd..7414180e 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -41,6 +41,7 @@ defaultRules:
     kubeApiserver: true
     kubeApiserverAvailability: true
     kubeApiserverSlos: true
+    kubeApiserverBurnrate: true
     kubelet: true
     kubeProxy: true
     kubePrometheusGeneral: true

```

## 35.4.0

**Release date:** 2022-05-27

![AppVersion: 0.56.2](https://img.shields.io/static/v1?label=AppVersion&message=0.56.2&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Allow either type of imagePullSecrets (#2056)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index a86dfba6..e613f9fd 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -121,6 +121,8 @@ global:
   ##
   imagePullSecrets: []
   # - name: "image-pull-secret"
+  # or
+  # - "image-pull-secret"
 
 ## Configuration for alertmanager
 ## ref: https://prometheus.io/docs/alerting/alertmanager/

```

## 35.3.1

**Release date:** 2022-05-22

![AppVersion: 0.56.2](https://img.shields.io/static/v1?label=AppVersion&message=0.56.2&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add null receiver to InfoInhibitor rule in default AlertManager configuration (#2000)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 108ccc7c..a86dfba6 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -171,9 +171,9 @@ alertmanager:
       repeat_interval: 12h
       receiver: 'null'
       routes:
-      - match:
-          alertname: Watchdog
-        receiver: 'null'
+      - receiver: 'null'
+        matchers:
+          - alertname =~ "InfoInhibitor|Watchdog"
     receivers:
     - name: 'null'
     templates:

```

## 35.3.0

**Release date:** 2022-05-21

![AppVersion: 0.56.2](https://img.shields.io/static/v1?label=AppVersion&message=0.56.2&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] add support for enableRemoteWriteReceiver in the prometheuses crd and… (#1963)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 3a9b311a..108ccc7c 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -2202,6 +2202,10 @@ prometheus:
     ##
     externalLabels: {}
 
+    ## enable --web.enable-remote-write-receiver flag on prometheus-server
+    ##
+    enableRemoteWriteReceiver: false
+
     ## Name of the external label used to denote replica name
     ##
     replicaExternalLabelName: ""

```

## 35.2.0

**Release date:** 2022-05-10

![AppVersion: 0.56.2](https://img.shields.io/static/v1?label=AppVersion&message=0.56.2&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Upgrade Grafana chart to 6.29 (#2029)

### Default value changes

```diff
# No changes in this release
```

## 35.1.0

**Release date:** 2022-05-10

![AppVersion: 0.56.2](https://img.shields.io/static/v1?label=AppVersion&message=0.56.2&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Version bump to 0.56.2 (#2042)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 594abb81..3a9b311a 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1708,7 +1708,7 @@ prometheusOperator:
   ##
   image:
     repository: quay.io/prometheus-operator/prometheus-operator
-    tag: v0.56.0
+    tag: v0.56.2
     sha: ""
     pullPolicy: IfNotPresent
 
@@ -1726,7 +1726,7 @@ prometheusOperator:
     # image to use for config and rule reloading
     image:
       repository: quay.io/prometheus-operator/prometheus-config-reloader
-      tag: v0.56.0
+      tag: v0.56.2
       sha: ""
 
     # resource config for prometheusConfigReloader

```

## 35.0.3

**Release date:** 2022-04-27

![AppVersion: 0.56.0](https://img.shields.io/static/v1?label=AppVersion&message=0.56.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Delete statefulset Grafana dashboard (#1892)

### Default value changes

```diff
# No changes in this release
```

## 35.0.2

**Release date:** 2022-04-27

![AppVersion: 0.56.0](https://img.shields.io/static/v1?label=AppVersion&message=0.56.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* kube-prometheus-stack adding cluster role for  prometheus-operator  prometheuses/status (#2012)

### Default value changes

```diff
# No changes in this release
```

## 35.0.1

**Release date:** 2022-04-27

![AppVersion: 0.56.0](https://img.shields.io/static/v1?label=AppVersion&message=0.56.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack]: Update prometheus-operator ref URL (#2008)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 3fa97fd0..594abb81 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -748,7 +748,7 @@ grafana:
       ## Create datasource for each Pod of Prometheus StatefulSet;
       ## this uses headless service `prometheus-operated` which is
       ## created by Prometheus Operator
-      ## ref: https://git.io/fjaBS
+      ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/0fee93e12dc7c2ea1218f19ae25ec6b893460590/pkg/prometheus/statefulset.go#L255-L286
       createPrometheusReplicasDatasources: false
       label: grafana_datasource
       labelValue: "1"

```

## 35.0.0

**Release date:** 2022-04-25

![AppVersion: 0.56.0](https://img.shields.io/static/v1?label=AppVersion&message=0.56.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Bump to prometheus-operator 0.56.0 (#2005)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index d930d939..3fa97fd0 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1708,7 +1708,7 @@ prometheusOperator:
   ##
   image:
     repository: quay.io/prometheus-operator/prometheus-operator
-    tag: v0.55.0
+    tag: v0.56.0
     sha: ""
     pullPolicy: IfNotPresent
 
@@ -1726,7 +1726,7 @@ prometheusOperator:
     # image to use for config and rule reloading
     image:
       repository: quay.io/prometheus-operator/prometheus-config-reloader
-      tag: v0.55.0
+      tag: v0.56.0
       sha: ""
 
     # resource config for prometheusConfigReloader
@@ -2160,7 +2160,7 @@ prometheus:
     ##
     image:
       repository: quay.io/prometheus/prometheus
-      tag: v2.34.0
+      tag: v2.35.0
       sha: ""
 
     ## Tolerations for use with node taints

```

## 34.10.0

**Release date:** 2022-04-13

![AppVersion: 0.55.0](https://img.shields.io/static/v1?label=AppVersion&message=0.55.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] allow for separately managed secret for additionalAlertadditionalAlertRelabelConfigs (#1977)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index a621d0da..d930d939 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -2567,6 +2567,14 @@ prometheus:
     #   replacement: $1
     #   action: labeldrop
 
+    ## If additional alert relabel configurations are already deployed in a single secret, or you want to manage
+    ## them separately from the helm deployment, you can use this section.
+    ## Expected values are the secret name and key
+    ## Cannot be used with additionalAlertRelabelConfigs
+    additionalAlertRelabelConfigsSecret: {}
+      # name:
+      # key:
+
     ## SecurityContext holds pod-level security attributes and common container settings.
     ## This defaults to non root user with uid 1000 and gid 2000.
     ## https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md

```

## 34.9.1

**Release date:** 2022-04-13

![AppVersion: 0.55.0](https://img.shields.io/static/v1?label=AppVersion&message=0.55.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* fix markdown & adjust linters (#1978)

### Default value changes

```diff
# No changes in this release
```

## 34.9.0

**Release date:** 2022-04-06

![AppVersion: 0.55.0](https://img.shields.io/static/v1?label=AppVersion&message=0.55.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Upgrade Grafana to 6.26 (#1953)

### Default value changes

```diff
# No changes in this release
```

## 34.8.0

**Release date:** 2022-04-04

![AppVersion: 0.55.0](https://img.shields.io/static/v1?label=AppVersion&message=0.55.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Allow using template for additionalScrapeConfigs (#1707)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 5c086f27..a621d0da 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -2470,6 +2470,7 @@ prometheus:
     ## appended, the user is responsible to make sure it is valid. Note that using this feature may expose the possibility
     ## to break upgrades of Prometheus. It is advised to review Prometheus release notes to ensure that no incompatible
     ## scrape configs are going to break Prometheus after the upgrade.
+    ## AdditionalScrapeConfigs can be defined as a list or as a templated string.
     ##
     ## The scrape configuration example below will find master nodes, provided they have the name .*mst.*, relabel the
     ## port to 2379 and allow etcd scraping provided it is running on all Kubernetes master nodes
@@ -2502,6 +2503,20 @@ prometheus:
     #   metric_relabel_configs:
     #   - regex: (kubernetes_io_hostname|failure_domain_beta_kubernetes_io_region|beta_kubernetes_io_os|beta_kubernetes_io_arch|beta_kubernetes_io_instance_type|failure_domain_beta_kubernetes_io_zone)
     #     action: labeldrop
+    #
+    ## If scrape config contains a repetitive section, you may want to use a template.
+    ## In the following example, you can see how to define `gce_sd_configs` for multiple zones
+    # additionalScrapeConfigs: |
+    #  - job_name: "node-exporter"
+    #    gce_sd_configs:
+    #    {{range $zone := .Values.gcp_zones}}
+    #    - project: "project1"
+    #      zone: "{{$zone}}"
+    #      port: 9100
+    #    {{end}}
+    #    relabel_configs:
+    #    ...
+
 
     ## If additional scrape configurations are already deployed in a single secret file you can use this section.
     ## Expected values are the secret name and key

```

## 34.7.1

**Release date:** 2022-04-03

![AppVersion: 0.55.0](https://img.shields.io/static/v1?label=AppVersion&message=0.55.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add values grafana.ingress.ingressClassName (#1950)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index f4806a0b..5c086f27 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -679,6 +679,11 @@ grafana:
     ##
     enabled: false
 
+    ## IngressClassName for Grafana Ingress.
+    ## Should be provided if Ingress is enable.
+    ##
+    # ingressClassName: nginx
+
     ## Annotations for Grafana Ingress
     ##
     annotations: {}

```

## 34.7.0

**Release date:** 2022-04-02

![AppVersion: 0.55.0](https://img.shields.io/static/v1?label=AppVersion&message=0.55.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Increase default CPU for config-reloader (#1700)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 90858c19..f4806a0b 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1727,10 +1727,10 @@ prometheusOperator:
     # resource config for prometheusConfigReloader
     resources:
       requests:
-        cpu: 100m
+        cpu: 200m
         memory: 50Mi
       limits:
-        cpu: 100m
+        cpu: 200m
         memory: 50Mi
 
   ## Thanos side-car image when configured

```

## 34.6.0

**Release date:** 2022-03-29

![AppVersion: 0.55.0](https://img.shields.io/static/v1?label=AppVersion&message=0.55.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] - upgrade node-exporter version (#1928)

### Default value changes

```diff
# No changes in this release
```

## 34.5.1

**Release date:** 2022-03-28

![AppVersion: 0.55.0](https://img.shields.io/static/v1?label=AppVersion&message=0.55.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Regenerate lock file (#1924)

### Default value changes

```diff
# No changes in this release
```

## 34.5.0

**Release date:** 2022-03-28

![AppVersion: 0.55.0](https://img.shields.io/static/v1?label=AppVersion&message=0.55.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add option to create aggregate ClusterRoles (#1888)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index cbf568d2..90858c19 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -101,6 +101,10 @@ additionalPrometheusRulesMap: {}
 global:
   rbac:
     create: true
+
+    ## Create ClusterRoles that extend the existing view, edit and admin ClusterRoles to interact with prometheus-operator CRDs
+    ## Ref: https://kubernetes.io/docs/reference/access-authn-authz/rbac/#aggregated-clusterroles
+    createAggregateClusterRoles: false
     pspEnabled: false
     pspAnnotations: {}
       ## Specify pod annotations

```

## 34.3.0

**Release date:** 2022-03-28

![AppVersion: 0.55.0](https://img.shields.io/static/v1?label=AppVersion&message=0.55.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Allow to override Grafana datasources UID (#1908)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 072e2c57..cbf568d2 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -723,6 +723,8 @@ grafana:
       enabled: true
       defaultDatasourceEnabled: true
 
+      uid: prometheus
+
       ## URL of prometheus datasource
       ##
       # url: http://prometheus-stack-prometheus:9090/

```

## 34.2.0

**Release date:** 2022-03-26

![AppVersion: 0.55.0](https://img.shields.io/static/v1?label=AppVersion&message=0.55.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* update prometheus images (#1918)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index ca7b2e50..072e2c57 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -414,7 +414,7 @@ alertmanager:
     ##
     image:
       repository: quay.io/prometheus/alertmanager
-      tag: v0.23.0
+      tag: v0.24.0
       sha: ""
 
     ## If true then the user will be responsible to provide a secret with alertmanager configuration
@@ -1731,7 +1731,7 @@ prometheusOperator:
   ##
   thanosImage:
     repository: quay.io/thanos/thanos
-    tag: v0.25.1
+    tag: v0.25.2
     sha: ""
 
   ## Set a Field Selector to filter watched secrets
@@ -2149,7 +2149,7 @@ prometheus:
     ##
     image:
       repository: quay.io/prometheus/prometheus
-      tag: v2.33.5
+      tag: v2.34.0
       sha: ""
 
     ## Tolerations for use with node taints

```

## 34.1.1

**Release date:** 2022-03-18

![AppVersion: 0.55.0](https://img.shields.io/static/v1?label=AppVersion&message=0.55.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* fix typo (#1887)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index b232ffee..ca7b2e50 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -475,7 +475,7 @@ alertmanager:
 
     ## AlermanagerConfig to be used as top level configuration
     ##
-    altermanagerConfiguration: {}
+    alertmanagerConfiguration: {}
     # - name: global-alertmanager-Configuration
 
     ## Define Log Format

```

## 34.1.0

**Release date:** 2022-03-17

![AppVersion: 0.55.0](https://img.shields.io/static/v1?label=AppVersion&message=0.55.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add alertmanager alertmanager configuration (#1880)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index a3de0794..b232ffee 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -473,6 +473,11 @@ alertmanager:
     #   matchLabels:
     #     alertmanagerconfig: enabled
 
+    ## AlermanagerConfig to be used as top level configuration
+    ##
+    altermanagerConfiguration: {}
+    # - name: global-alertmanager-Configuration
+
     ## Define Log Format
     # Use logfmt (default) or json logging
     logFormat: logfmt

```

## 34.0.0

**Release date:** 2022-03-16

![AppVersion: 0.55.0](https://img.shields.io/static/v1?label=AppVersion&message=0.55.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Prometheus-Operator v0.55.0 (#1873)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index bbd3f1ca..a3de0794 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1692,7 +1692,7 @@ prometheusOperator:
   ##
   image:
     repository: quay.io/prometheus-operator/prometheus-operator
-    tag: v0.54.1
+    tag: v0.55.0
     sha: ""
     pullPolicy: IfNotPresent
 
@@ -1710,7 +1710,7 @@ prometheusOperator:
     # image to use for config and rule reloading
     image:
       repository: quay.io/prometheus-operator/prometheus-config-reloader
-      tag: v0.54.1
+      tag: v0.55.0
       sha: ""
 
     # resource config for prometheusConfigReloader
@@ -1726,7 +1726,7 @@ prometheusOperator:
   ##
   thanosImage:
     repository: quay.io/thanos/thanos
-    tag: v0.24.0
+    tag: v0.25.1
     sha: ""
 
   ## Set a Field Selector to filter watched secrets
@@ -2144,7 +2144,7 @@ prometheus:
     ##
     image:
       repository: quay.io/prometheus/prometheus
-      tag: v2.33.4
+      tag: v2.33.5
       sha: ""
 
     ## Tolerations for use with node taints

```

## 33.2.1

**Release date:** 2022-03-13

![AppVersion: 0.54.1](https://img.shields.io/static/v1?label=AppVersion&message=0.54.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Replace invalid characters in the version (#1833)

### Default value changes

```diff
# No changes in this release
```

## 33.2.0

**Release date:** 2022-03-05

![AppVersion: 0.54.1](https://img.shields.io/static/v1?label=AppVersion&message=0.54.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add uid to grafana default datasources (#1816)

### Default value changes

```diff
# No changes in this release
```

## 33.1.0

**Release date:** 2022-02-28

![AppVersion: 0.54.1](https://img.shields.io/static/v1?label=AppVersion&message=0.54.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Expose the container-specific security context for Prometheus Operator (#1811)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 3ce93c58..bbd3f1ca 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1681,6 +1681,13 @@ prometheusOperator:
     runAsNonRoot: true
     runAsUser: 65534
 
+  ## Container-specific security context configuration
+  ## ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
+  ##
+  containerSecurityContext:
+    allowPrivilegeEscalation: false
+    readOnlyRootFilesystem: true
+
   ## Prometheus-operator image
   ##
   image:

```

## 33.0.0

**Release date:** 2022-02-25

![AppVersion: 0.54.1](https://img.shields.io/static/v1?label=AppVersion&message=0.54.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Version bump to 0.54.1 (#1830)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 16d4d12b..3ce93c58 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1466,8 +1466,8 @@ prometheusOperator:
       enabled: true
       image:
         repository: k8s.gcr.io/ingress-nginx/kube-webhook-certgen
-        tag: v1.0
-        sha: "f3b6b39a6062328c095337b4cadcefd1612348fdd5190b1dcbcb9b9e90bd8068"
+        tag: v1.1.1
+        sha: ""
         pullPolicy: IfNotPresent
       resources: {}
       ## Provide a priority class name to the webhook patching job
@@ -1685,7 +1685,7 @@ prometheusOperator:
   ##
   image:
     repository: quay.io/prometheus-operator/prometheus-operator
-    tag: v0.54.0
+    tag: v0.54.1
     sha: ""
     pullPolicy: IfNotPresent
 
@@ -1703,7 +1703,7 @@ prometheusOperator:
     # image to use for config and rule reloading
     image:
       repository: quay.io/prometheus-operator/prometheus-config-reloader
-      tag: v0.54.0
+      tag: v0.54.1
       sha: ""
 
     # resource config for prometheusConfigReloader
@@ -2137,7 +2137,7 @@ prometheus:
     ##
     image:
       repository: quay.io/prometheus/prometheus
-      tag: v2.33.1
+      tag: v2.33.4
       sha: ""
 
     ## Tolerations for use with node taints

```

## 32.4.0

**Release date:** 2022-02-25

![AppVersion: 0.54.0](https://img.shields.io/static/v1?label=AppVersion&message=0.54.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Bump grafana and kube-state-metrics subcharts (#1828)

### Default value changes

```diff
# No changes in this release
```

## 32.3.0

**Release date:** 2022-02-24

![AppVersion: 0.54.0](https://img.shields.io/static/v1?label=AppVersion&message=0.54.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Make burnrate and histogram apiserver rules optional (#1792)

### Default value changes

```diff
# No changes in this release
```

## 32.2.1

**Release date:** 2022-02-14

![AppVersion: 0.54.0](https://img.shields.io/static/v1?label=AppVersion&message=0.54.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Update Grafana Chart version to 6.21.5 (#1804)

### Default value changes

```diff
# No changes in this release
```

## 32.2.0

**Release date:** 2022-02-09

![AppVersion: 0.54.0](https://img.shields.io/static/v1?label=AppVersion&message=0.54.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add labelValue to overwrite grafana dashboard mark label (#1782)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index e28d904a..16d4d12b 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -702,6 +702,7 @@ grafana:
     dashboards:
       enabled: true
       label: grafana_dashboard
+      labelValue: "1"
 
       ## Annotations for Grafana dashboard configmaps
       ##
@@ -734,6 +735,7 @@ grafana:
       ## ref: https://git.io/fjaBS
       createPrometheusReplicasDatasources: false
       label: grafana_datasource
+      labelValue: "1"
 
   extraConfigmapMounts: []
   # - name: certs-configmap

```

## 32.1.0

**Release date:** 2022-02-09

![AppVersion: 0.54.0](https://img.shields.io/static/v1?label=AppVersion&message=0.54.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] add support for configure services externalTr… (#1757)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 7633d2ad..e28d904a 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -322,6 +322,11 @@ alertmanager:
     externalIPs: []
     loadBalancerIP: ""
     loadBalancerSourceRanges: []
+
+    ## Denotes if this Service desires to route external traffic to node-local or cluster-wide endpoints
+    ##
+    externalTrafficPolicy: Cluster
+
     ## Service type
     ##
     type: ClusterIP
@@ -347,6 +352,11 @@ alertmanager:
     ## Loadbalancer source IP ranges
     ## Only used if servicePerReplica.type is "LoadBalancer"
     loadBalancerSourceRanges: []
+
+    ## Denotes if this Service desires to route external traffic to node-local or cluster-wide endpoints
+    ##
+    externalTrafficPolicy: Cluster
+
     ## Service type
     ##
     type: ClusterIP
@@ -1543,6 +1553,10 @@ prometheusOperator:
     loadBalancerIP: ""
     loadBalancerSourceRanges: []
 
+    ## Denotes if this Service desires to route external traffic to node-local or cluster-wide endpoints
+    ##
+    externalTrafficPolicy: Cluster
+
   ## Service type
   ## NodePort, ClusterIP, LoadBalancer
   ##
@@ -1738,6 +1752,10 @@ prometheus:
     annotations: {}
     labels: {}
 
+    ## Denotes if this Service desires to route external traffic to node-local or cluster-wide endpoints
+    ##
+    externalTrafficPolicy: Cluster
+
     ## Service type
     ##
     type: ClusterIP
@@ -1801,6 +1819,10 @@ prometheus:
     httpPort: 10902
     targetHttpPort: "http"
 
+    ## Denotes if this Service desires to route external traffic to node-local or cluster-wide endpoints
+    ##
+    externalTrafficPolicy: Cluster
+
     ## Service type
     ##
     type: LoadBalancer
@@ -1838,6 +1860,11 @@ prometheus:
     ## Only use if service.type is "LoadBalancer"
     loadBalancerIP: ""
     loadBalancerSourceRanges: []
+
+    ## Denotes if this Service desires to route external traffic to node-local or cluster-wide endpoints
+    ##
+    externalTrafficPolicy: Cluster
+
     ## Service type
     ##
     type: ClusterIP
@@ -1872,6 +1899,11 @@ prometheus:
     ## Loadbalancer source IP ranges
     ## Only used if servicePerReplica.type is "LoadBalancer"
     loadBalancerSourceRanges: []
+
+    ## Denotes if this Service desires to route external traffic to node-local or cluster-wide endpoints
+    ##
+    externalTrafficPolicy: Cluster
+
     ## Service type
     ##
     type: ClusterIP

```

## 32.0.3

**Release date:** 2022-02-09

![AppVersion: 0.54.0](https://img.shields.io/static/v1?label=AppVersion&message=0.54.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Import Node Exporter / Nodes grafana dashboard conditionally (#1687)

### Default value changes

```diff
# No changes in this release
```

## 32.0.2

**Release date:** 2022-02-08

![AppVersion: 0.54.0](https://img.shields.io/static/v1?label=AppVersion&message=0.54.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Adding GMartinez-Sisti as new maintainer (#1778)

### Default value changes

```diff
# No changes in this release
```

## 32.0.1

**Release date:** 2022-02-07

![AppVersion: 0.54.0](https://img.shields.io/static/v1?label=AppVersion&message=0.54.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] remove depreciation warnings in prometheus-node-exporter (#1732)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 6f3feb5b..7633d2ad 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1377,8 +1377,8 @@ prometheus-node-exporter:
     ##
     jobLabel: node-exporter
   extraArgs:
-    - --collector.filesystem.ignored-mount-points=^/(dev|proc|sys|var/lib/docker/.+|var/lib/kubelet/.+)($|/)
-    - --collector.filesystem.ignored-fs-types=^(autofs|binfmt_misc|bpf|cgroup2?|configfs|debugfs|devpts|devtmpfs|fusectl|hugetlbfs|iso9660|mqueue|nsfs|overlay|proc|procfs|pstore|rpc_pipefs|securityfs|selinuxfs|squashfs|sysfs|tracefs)$
+    - --collector.filesystem.mount-points-exclude=^/(dev|proc|sys|var/lib/docker/.+|var/lib/kubelet/.+)($|/)
+    - --collector.filesystem.fs-types-exclude=^(autofs|binfmt_misc|bpf|cgroup2?|configfs|debugfs|devpts|devtmpfs|fusectl|hugetlbfs|iso9660|mqueue|nsfs|overlay|proc|procfs|pstore|rpc_pipefs|securityfs|selinuxfs|squashfs|sysfs|tracefs)$
   service:
     portName: http-metrics
   prometheus:

```

## 32.0.0

**Release date:** 2022-02-07

![AppVersion: 0.54.0](https://img.shields.io/static/v1?label=AppVersion&message=0.54.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* kube-prometheus-stack: Prometheus-Operator v0.54.0 (#1770)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 887f577a..6f3feb5b 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -760,7 +760,7 @@ grafana:
     # If true, a ServiceMonitor CRD is created for a prometheus operator
     # https://github.com/coreos/prometheus-operator
     #
-    enabled: false
+    enabled: true
 
     # Path to use for scraping metrics. Might be different if server.root_url is set
     # in grafana.ini
@@ -1669,7 +1669,7 @@ prometheusOperator:
   ##
   image:
     repository: quay.io/prometheus-operator/prometheus-operator
-    tag: v0.53.1
+    tag: v0.54.0
     sha: ""
     pullPolicy: IfNotPresent
 
@@ -1687,7 +1687,7 @@ prometheusOperator:
     # image to use for config and rule reloading
     image:
       repository: quay.io/prometheus-operator/prometheus-config-reloader
-      tag: v0.53.1
+      tag: v0.54.0
       sha: ""
 
     # resource config for prometheusConfigReloader
@@ -2103,7 +2103,7 @@ prometheus:
     ##
     image:
       repository: quay.io/prometheus/prometheus
-      tag: v2.32.1
+      tag: v2.33.1
       sha: ""
 
     ## Tolerations for use with node taints

```

## 31.0.2

**Release date:** 2022-02-07

![AppVersion: 0.53.1](https://img.shields.io/static/v1?label=AppVersion&message=0.53.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack]  Scope CoreDNS dashboard to CoreDNS job (#1730)

### Default value changes

```diff
# No changes in this release
```

## 31.0.1

**Release date:** 2022-02-07

![AppVersion: 0.53.1](https://img.shields.io/static/v1?label=AppVersion&message=0.53.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] patch updates for kube-state-metrics and grafana (#1766)

### Default value changes

```diff
# No changes in this release
```

## 31.0.0

**Release date:** 2022-02-02

![AppVersion: 0.53.1](https://img.shields.io/static/v1?label=AppVersion&message=0.53.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] rm grafana serviceMonitor (#1652)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 4782a753..887f577a 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -756,26 +756,27 @@ grafana:
   service:
     portName: http-web
 
-  ## If true, create a serviceMonitor for grafana
-  ##
   serviceMonitor:
-    ## Scrape interval. If not set, the Prometheus default scrape interval is used.
-    ##
-    interval: ""
-    selfMonitor: true
+    # If true, a ServiceMonitor CRD is created for a prometheus operator
+    # https://github.com/coreos/prometheus-operator
+    #
+    enabled: false
 
     # Path to use for scraping metrics. Might be different if server.root_url is set
     # in grafana.ini
     path: "/metrics"
 
+    #  namespace: monitoring  (defaults to use the namespace this chart is deployed to)
 
-    ## MetricRelabelConfigs to apply to samples after scraping, but before ingestion.
-    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#relabelconfig
-    ##
-    metricRelabelings: []
-    # - action: keep
-    #   regex: 'kube_(daemonset|deployment|pod|namespace|node|statefulset).+'
-    #   sourceLabels: [__name__]
+    # labels for the ServiceMonitor
+    labels: {}
+
+    # Scrape interval. If not set, the Prometheus default scrape interval is used.
+    #
+    interval: ""
+    scheme: http
+    tlsConfig: {}
+    scrapeTimeout: 30s
 
     ## RelabelConfigs to apply to samples before scraping
     ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#relabelconfig

```

## 30.2.0

**Release date:** 2022-01-25

![AppVersion: 0.53.1](https://img.shields.io/static/v1?label=AppVersion&message=0.53.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Grafana: upgrade to 6.21.0, kube-state-metrics: upgrade to 4.4.1. (#1742)

### Default value changes

```diff
# No changes in this release
```

## 30.1.0

**Release date:** 2022-01-20

![AppVersion: 0.53.1](https://img.shields.io/static/v1?label=AppVersion&message=0.53.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] final PR to finalise selectorOverrider (#1711)

### Default value changes

```diff
# No changes in this release
```

## 30.0.3

**Release date:** 2022-01-19

![AppVersion: 0.53.1](https://img.shields.io/static/v1?label=AppVersion&message=0.53.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Fix alertmanager podAntiAffinity (#1717)

### Default value changes

```diff
# No changes in this release
```

## 30.0.2

**Release date:** 2022-01-19

![AppVersion: 0.53.1](https://img.shields.io/static/v1?label=AppVersion&message=0.53.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Rename label selector for pod disruption budget of Alertmanager (#1599)

### Default value changes

```diff
# No changes in this release
```

## 30.0.1

**Release date:** 2022-01-10

![AppVersion: 0.53.1](https://img.shields.io/static/v1?label=AppVersion&message=0.53.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] set an empty list as default for namespaces.additional in prometheus-operator (#1521)

### Default value changes

```diff
# No changes in this release
```

## 30.0.0

**Release date:** 2022-01-08

![AppVersion: 0.53.1](https://img.shields.io/static/v1?label=AppVersion&message=0.53.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* use kube-state-metrics release label option (#1692)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index ab24d1b7..4782a753 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1319,6 +1319,7 @@ kube-state-metrics:
   namespaceOverride: ""
   rbac:
     create: true
+  releaseLabel: true
   prometheus:
     monitor:
       enabled: true

```

## 29.0.1

**Release date:** 2022-01-08

![AppVersion: 0.53.1](https://img.shields.io/static/v1?label=AppVersion&message=0.53.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] typo fix (#1693)

### Default value changes

```diff
# No changes in this release
```

## 29.0.0

**Release date:** 2022-01-07

![AppVersion: 0.53.1](https://img.shields.io/static/v1?label=AppVersion&message=0.53.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Support new default ports for controlplane (#1393)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 86eab1c7..ab24d1b7 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -980,8 +980,11 @@ kubeControllerManager:
   ##
   service:
     enabled: true
-    port: 10252
-    targetPort: 10252
+    ## If null or unset, the value is determined dynamically based on target Kubernetes version due to change
+    ## of default port in Kubernetes 1.22.
+    ##
+    port: null
+    targetPort: null
     # selector:
     #   component: kube-controller-manager
 
@@ -996,9 +999,10 @@ kubeControllerManager:
     proxyUrl: ""
 
     ## Enable scraping kube-controller-manager over https.
-    ## Requires proper certs (not self-signed) and delegated authentication/authorization checks
+    ## Requires proper certs (not self-signed) and delegated authentication/authorization checks.
+    ## If null or unset, the value is determined dynamically based on target Kubernetes version.
     ##
-    https: false
+    https: null
 
     # Skip TLS certificate validation when scraping
     insecureSkipVerify: null
@@ -1205,8 +1209,11 @@ kubeScheduler:
   ##
   service:
     enabled: true
-    port: 10251
-    targetPort: 10251
+    ## If null or unset, the value is determined dynamically based on target Kubernetes version due to change
+    ## of default port in Kubernetes 1.23.
+    ##
+    port: null
+    targetPort: null
     # selector:
     #   component: kube-scheduler
 
@@ -1219,9 +1226,10 @@ kubeScheduler:
     ##
     proxyUrl: ""
     ## Enable scraping kube-scheduler over https.
-    ## Requires proper certs (not self-signed) and delegated authentication/authorization checks
+    ## Requires proper certs (not self-signed) and delegated authentication/authorization checks.
+    ## If null or unset, the value is determined dynamically based on target Kubernetes version.
     ##
-    https: false
+    https: null
 
     ## Skip TLS certificate validation when scraping
     insecureSkipVerify: null

```

## 28.0.2

**Release date:** 2022-01-07

![AppVersion: 0.53.1](https://img.shields.io/static/v1?label=AppVersion&message=0.53.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] image tag and sha logic for alertmanager instance (#1684)

### Default value changes

```diff
# No changes in this release
```

## 28.0.1

**Release date:** 2022-01-07

![AppVersion: 0.53.1](https://img.shields.io/static/v1?label=AppVersion&message=0.53.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] disable all matched alerts (#1678)

### Default value changes

```diff
# No changes in this release
```

## 28.0.0

**Release date:** 2022-01-07

![AppVersion: 0.53.1](https://img.shields.io/static/v1?label=AppVersion&message=0.53.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Disable PodSecurityPolicies Creation by Default (#1657)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 7f9b1841..86eab1c7 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -101,7 +101,7 @@ additionalPrometheusRulesMap: {}
 global:
   rbac:
     create: true
-    pspEnabled: true
+    pspEnabled: false
     pspAnnotations: {}
       ## Specify pod annotations
       ## Ref: https://kubernetes.io/docs/concepts/policy/pod-security-policy/#apparmor
@@ -650,6 +650,11 @@ grafana:
 
   adminPassword: prom-operator
 
+  rbac:
+    ## If true, Grafana PSPs will be created
+    ##
+    pspEnabled: false
+
   ingress:
     ## If true, Grafana Ingress will be created
     ##
@@ -1306,8 +1311,6 @@ kube-state-metrics:
   namespaceOverride: ""
   rbac:
     create: true
-  podSecurityPolicy:
-    enabled: true
   prometheus:
     monitor:
       enabled: true
@@ -1406,6 +1409,10 @@ prometheus-node-exporter:
       #   targetLabel: nodename
       #   replacement: $1
       #   action: replace
+  rbac:
+    ## If true, create PSPs for node-exporter
+    ##
+    pspEnabled: false
 
 ## Manages Prometheus and Alertmanager components
 ##

```

## 27.2.2

**Release date:** 2022-01-06

![AppVersion: 0.53.1](https://img.shields.io/static/v1?label=AppVersion&message=0.53.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add desaintmartin as maintainer. (#1679)

### Default value changes

```diff
# No changes in this release
```

## 27.2.1

**Release date:** 2022-01-04

![AppVersion: 0.53.1](https://img.shields.io/static/v1?label=AppVersion&message=0.53.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Adding andrewgkew as new maintainer (#1672)

### Default value changes

```diff
# No changes in this release
```

## 27.2.0

**Release date:** 2022-01-04

![AppVersion: 0.53.1](https://img.shields.io/static/v1?label=AppVersion&message=0.53.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* upgrade grafana version (#1662)

### Default value changes

```diff
# No changes in this release
```

## 27.1.0

**Release date:** 2022-01-02

![AppVersion: 0.53.1](https://img.shields.io/static/v1?label=AppVersion&message=0.53.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* update thanos to 0.24.0 (#1653)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index c9cf5c6b..7f9b1841 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1686,7 +1686,7 @@ prometheusOperator:
   ##
   thanosImage:
     repository: quay.io/thanos/thanos
-    tag: v0.23.1
+    tag: v0.24.0
     sha: ""
 
   ## Set a Field Selector to filter watched secrets

```

## 27.0.1

**Release date:** 2021-12-31

![AppVersion: 0.53.1](https://img.shields.io/static/v1?label=AppVersion&message=0.53.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] fix kubeApiserver rule (#1651)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index de9a0113..c9cf5c6b 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -38,6 +38,7 @@ defaultRules:
     configReloaders: true
     general: true
     k8s: true
+    kubeApiserver: true
     kubeApiserverAvailability: true
     kubeApiserverSlos: true
     kubelet: true

```

## 27.0.0

**Release date:** 2021-12-30

![AppVersion: 0.53.1](https://img.shields.io/static/v1?label=AppVersion&message=0.53.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] make nodeexporter & config-reloaders rules configurable (#1646)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index c8f32cd5..de9a0113 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -35,18 +35,15 @@ defaultRules:
   rules:
     alertmanager: true
     etcd: true
+    configReloaders: true
     general: true
     k8s: true
-    kubeApiserver: true
     kubeApiserverAvailability: true
-    kubeApiserverError: true
     kubeApiserverSlos: true
     kubelet: true
     kubeProxy: true
     kubePrometheusGeneral: true
-    kubePrometheusNodeAlerting: true
     kubePrometheusNodeRecording: true
-    kubernetesAbsent: true
     kubernetesApps: true
     kubernetesResources: true
     kubernetesStorage: true
@@ -55,9 +52,10 @@ defaultRules:
     kubeStateMetrics: true
     network: true
     node: true
+    nodeExporterAlerting: true
+    nodeExporterRecording: true
     prometheus: true
     prometheusOperator: true
-    time: true
 
   ## Reduce app namespace alert scope
   appNamespacesTarget: ".*"
@@ -70,6 +68,9 @@ defaultRules:
   ## Additional labels for PrometheusRule alerts
   additionalRuleLabels: {}
 
+  ## Prefix for runbook URLs. Use this to override the first part of the runbookURLs that is common to all rules.
+  runbookUrl: "https://runbooks.prometheus-operator.dev/runbooks"
+
   ## Disabled PrometheusRule alerts
   disabled: {}
   # KubeAPIDown: true

```

## 26.2.0

**Release date:** 2021-12-30

![AppVersion: 0.53.1](https://img.shields.io/static/v1?label=AppVersion&message=0.53.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Disable specific alert in values.yaml (#1173)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 7e828c25..c8f32cd5 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -70,6 +70,11 @@ defaultRules:
   ## Additional labels for PrometheusRule alerts
   additionalRuleLabels: {}
 
+  ## Disabled PrometheusRule alerts
+  disabled: {}
+  # KubeAPIDown: true
+  # NodeRAIDDegraded: true
+
 ## Deprecated way to provide custom recording or alerting rules to be deployed into the cluster.
 ##
 # additionalPrometheusRules: []

```

## 26.1.0

**Release date:** 2021-12-29

![AppVersion: 0.53.1](https://img.shields.io/static/v1?label=AppVersion&message=0.53.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack]update kubernetes-system-kube-proxy.yaml (#1629)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index b0d1fa30..7e828c25 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -42,6 +42,7 @@ defaultRules:
     kubeApiserverError: true
     kubeApiserverSlos: true
     kubelet: true
+    kubeProxy: true
     kubePrometheusGeneral: true
     kubePrometheusNodeAlerting: true
     kubePrometheusNodeRecording: true

```

## 26.0.0

**Release date:** 2021-12-28

![AppVersion: 0.53.1](https://img.shields.io/static/v1?label=AppVersion&message=0.53.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* enable nodeexporter servicemonitor by default again (#1640)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index bc9ca9bf..b0d1fa30 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1362,7 +1362,7 @@ prometheus-node-exporter:
     portName: http-metrics
   prometheus:
     monitor:
-      enabled: false
+      enabled: true
 
       jobLabel: jobLabel
 

```

## 25.2.0

**Release date:** 2021-12-26

![AppVersion: 0.53.1](https://img.shields.io/static/v1?label=AppVersion&message=0.53.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] added support for cert duration customization; (#1624)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 7d335e6a..bc9ca9bf 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1453,6 +1453,11 @@ prometheusOperator:
     # Use certmanager to generate webhook certs
     certManager:
       enabled: false
+      # self-signed root certificate
+      rootCert:
+        duration: ""  # default to be 5y
+      admissionCert:
+        duration: ""  # default to be 1y
       # issuerRef:
       #   name: "issuer"
       #   kind: "ClusterIssuer"

```

## 25.1.0

**Release date:** 2021-12-22

![AppVersion: 0.53.1](https://img.shields.io/static/v1?label=AppVersion&message=0.53.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add publishNotReadyAddresses prometheus service value. (#1619)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 3b30ba23..7d335e6a 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1812,6 +1812,13 @@ prometheus:
     ##
     type: ClusterIP
 
+    ## Additional port to define in the Service
+    additionalPorts: []
+
+    ## Consider that all endpoints are considered "ready" even if the Pods themselves are not
+    ## Ref: https://kubernetes.io/docs/reference/kubernetes-api/service-resources/service-v1/#ServiceSpec
+    publishNotReadyAddresses: false
+
     sessionAffinity: ""
 
   ## Configuration for creating a separate Service for each statefulset Prometheus replica

```

## 25.0.0

**Release date:** 2021-12-20

![AppVersion: 0.53.1](https://img.shields.io/static/v1?label=AppVersion&message=0.53.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Prometheus-Operator 0.53.1 (#1614)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 10f8e9c9..3b30ba23 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -58,8 +58,6 @@ defaultRules:
     prometheusOperator: true
     time: true
 
-  ## Runbook url prefix for default rules
-  runbookUrl: https://github.com/kubernetes-monitoring/kubernetes-mixin/tree/master/runbook.md#
   ## Reduce app namespace alert scope
   appNamespacesTarget: ".*"
 
@@ -361,13 +359,13 @@ alertmanager:
     scheme: ""
 
     ## tlsConfig: TLS configuration to use when scraping the endpoint. For example if using istio mTLS.
-    ## Of type: https://github.com/coreos/prometheus-operator/blob/master/Documentation/api.md#tlsconfig
+    ## Of type: https://github.com/coreos/prometheus-operator/blob/main/Documentation/api.md#tlsconfig
     tlsConfig: {}
 
     bearerTokenFile:
 
     ## MetricRelabelConfigs to apply to samples after scraping, but before ingestion.
-    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#relabelconfig
     ##
     metricRelabelings: []
     # - action: keep
@@ -375,7 +373,7 @@ alertmanager:
     #   sourceLabels: [__name__]
 
     ## RelabelConfigs to apply to samples before scraping
-    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#relabelconfig
     ##
     relabelings: []
     # - sourceLabels: [__meta_kubernetes_pod_node_name]
@@ -386,7 +384,7 @@ alertmanager:
     #   action: replace
 
   ## Settings affecting alertmanagerSpec
-  ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#alertmanagerspec
+  ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#alertmanagerspec
   ##
   alertmanagerSpec:
     ## Standard object's metadata. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata
@@ -475,7 +473,7 @@ alertmanager:
     retention: 120h
 
     ## Storage is the definition of how storage will be used by the Alertmanager instances.
-    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/user-guides/storage.md
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/user-guides/storage.md
     ##
     storage: {}
     # volumeClaimTemplate:
@@ -759,7 +757,7 @@ grafana:
 
 
     ## MetricRelabelConfigs to apply to samples after scraping, but before ingestion.
-    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#relabelconfig
     ##
     metricRelabelings: []
     # - action: keep
@@ -767,7 +765,7 @@ grafana:
     #   sourceLabels: [__name__]
 
     ## RelabelConfigs to apply to samples before scraping
-    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#relabelconfig
     ##
     relabelings: []
     # - sourceLabels: [__meta_kubernetes_pod_node_name]
@@ -799,7 +797,7 @@ kubeApiServer:
         provider: kubernetes
 
     ## MetricRelabelConfigs to apply to samples after scraping, but before ingestion.
-    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#relabelconfig
     ##
     metricRelabelings: []
     # - action: keep
@@ -807,7 +805,7 @@ kubeApiServer:
     #   sourceLabels: [__name__]
 
     ## RelabelConfigs to apply to samples before scraping
-    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#relabelconfig
     ##
     relabelings: []
     # - sourceLabels:
@@ -855,7 +853,7 @@ kubelet:
     resourcePath: "/metrics/resource/v1alpha1"
 
     ## MetricRelabelConfigs to apply to samples after scraping, but before ingestion.
-    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#relabelconfig
     ##
     cAdvisorMetricRelabelings: []
     # - sourceLabels: [__name__, image]
@@ -870,7 +868,7 @@ kubelet:
     #   action: drop
 
     ## MetricRelabelConfigs to apply to samples after scraping, but before ingestion.
-    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#relabelconfig
     ##
     probesMetricRelabelings: []
     # - sourceLabels: [__name__, image]
@@ -885,7 +883,7 @@ kubelet:
     #   action: drop
 
     ## RelabelConfigs to apply to samples before scraping
-    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#relabelconfig
     ##
     ## metrics_path is required to match upstream rules and charts
     cAdvisorRelabelings:
@@ -899,7 +897,7 @@ kubelet:
     #   action: replace
 
     ## RelabelConfigs to apply to samples before scraping
-    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#relabelconfig
     ##
     probesRelabelings:
       - sourceLabels: [__metrics_path__]
@@ -912,7 +910,7 @@ kubelet:
     #   action: replace
 
     ## RelabelConfigs to apply to samples before scraping
-    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#relabelconfig
     ##
     resourceRelabelings:
       - sourceLabels: [__metrics_path__]
@@ -925,7 +923,7 @@ kubelet:
     #   action: replace
 
     ## MetricRelabelConfigs to apply to samples after scraping, but before ingestion.
-    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#relabelconfig
     ##
     metricRelabelings: []
     # - sourceLabels: [__name__, image]
@@ -940,7 +938,7 @@ kubelet:
     #   action: drop
 
     ## RelabelConfigs to apply to samples before scraping
-    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#relabelconfig
     ##
     ## metrics_path is required to match upstream rules and charts
     relabelings:
@@ -996,7 +994,7 @@ kubeControllerManager:
     serverName: null
 
     ## MetricRelabelConfigs to apply to samples after scraping, but before ingestion.
-    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#relabelconfig
     ##
     metricRelabelings: []
     # - action: keep
@@ -1004,7 +1002,7 @@ kubeControllerManager:
     #   sourceLabels: [__name__]
 
     ## RelabelConfigs to apply to samples before scraping
-    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#relabelconfig
     ##
     relabelings: []
     # - sourceLabels: [__meta_kubernetes_pod_node_name]
@@ -1033,7 +1031,7 @@ coreDns:
     proxyUrl: ""
 
     ## MetricRelabelConfigs to apply to samples after scraping, but before ingestion.
-    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#relabelconfig
     ##
     metricRelabelings: []
     # - action: keep
@@ -1041,7 +1039,7 @@ coreDns:
     #   sourceLabels: [__name__]
 
     ## RelabelConfigs to apply to samples before scraping
-    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#relabelconfig
     ##
     relabelings: []
     # - sourceLabels: [__meta_kubernetes_pod_node_name]
@@ -1074,7 +1072,7 @@ kubeDns:
     proxyUrl: ""
 
     ## MetricRelabelConfigs to apply to samples after scraping, but before ingestion.
-    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#relabelconfig
     ##
     metricRelabelings: []
     # - action: keep
@@ -1082,7 +1080,7 @@ kubeDns:
     #   sourceLabels: [__name__]
 
     ## RelabelConfigs to apply to samples before scraping
-    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#relabelconfig
     ##
     relabelings: []
     # - sourceLabels: [__meta_kubernetes_pod_node_name]
@@ -1093,7 +1091,7 @@ kubeDns:
     #   action: replace
 
     ## MetricRelabelConfigs to apply to samples after scraping, but before ingestion.
-    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#relabelconfig
     ##
     dnsmasqMetricRelabelings: []
     # - action: keep
@@ -1101,7 +1099,7 @@ kubeDns:
     #   sourceLabels: [__name__]
 
     ## RelabelConfigs to apply to samples before scraping
-    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#relabelconfig
     ##
     dnsmasqRelabelings: []
     # - sourceLabels: [__meta_kubernetes_pod_node_name]
@@ -1159,7 +1157,7 @@ kubeEtcd:
     keyFile: ""
 
     ## MetricRelabelConfigs to apply to samples after scraping, but before ingestion.
-    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#relabelconfig
     ##
     metricRelabelings: []
     # - action: keep
@@ -1167,7 +1165,7 @@ kubeEtcd:
     #   sourceLabels: [__name__]
 
     ## RelabelConfigs to apply to samples before scraping
-    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#relabelconfig
     ##
     relabelings: []
     # - sourceLabels: [__meta_kubernetes_pod_node_name]
@@ -1219,7 +1217,7 @@ kubeScheduler:
     serverName: null
 
     ## MetricRelabelConfigs to apply to samples after scraping, but before ingestion.
-    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#relabelconfig
     ##
     metricRelabelings: []
     # - action: keep
@@ -1227,7 +1225,7 @@ kubeScheduler:
     #   sourceLabels: [__name__]
 
     ## RelabelConfigs to apply to samples before scraping
-    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#relabelconfig
     ##
     relabelings: []
     # - sourceLabels: [__meta_kubernetes_pod_node_name]
@@ -1273,7 +1271,7 @@ kubeProxy:
     https: false
 
     ## MetricRelabelConfigs to apply to samples after scraping, but before ingestion.
-    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#relabelconfig
     ##
     metricRelabelings: []
     # - action: keep
@@ -1281,7 +1279,7 @@ kubeProxy:
     #   sourceLabels: [__name__]
 
     ## RelabelConfigs to apply to samples before scraping
-    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#relabelconfig
     ##
     relabelings: []
     # - action: keep
@@ -1323,7 +1321,7 @@ kube-state-metrics:
       honorLabels: true
 
       ## MetricRelabelConfigs to apply to samples after scraping, but before ingestion.
-      ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
+      ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#relabelconfig
       ##
       metricRelabelings: []
       # - action: keep
@@ -1331,7 +1329,7 @@ kube-state-metrics:
       #   sourceLabels: [__name__]
 
       ## RelabelConfigs to apply to samples before scraping
-      ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
+      ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#relabelconfig
       ##
       relabelings: []
       # - sourceLabels: [__meta_kubernetes_pod_node_name]
@@ -1381,7 +1379,7 @@ prometheus-node-exporter:
       proxyUrl: ""
 
       ## MetricRelabelConfigs to apply to samples after scraping, but before ingestion.
-      ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
+      ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#relabelconfig
       ##
       metricRelabelings: []
       # - sourceLabels: [__name__]
@@ -1391,7 +1389,7 @@ prometheus-node-exporter:
       #   action: drop
 
       ## RelabelConfigs to apply to samples before scraping
-      ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
+      ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#relabelconfig
       ##
       relabelings: []
       # - sourceLabels: [__meta_kubernetes_pod_node_name]
@@ -1544,7 +1542,7 @@ prometheusOperator:
   # logLevel: error
 
   ## If true, the operator will create and maintain a service for scraping kubelets
-  ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/helm/prometheus-operator/README.md
+  ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/helm/prometheus-operator/README.md
   ##
   kubeletService:
     enabled: true
@@ -1641,7 +1639,7 @@ prometheusOperator:
   ##
   image:
     repository: quay.io/prometheus-operator/prometheus-operator
-    tag: v0.52.1
+    tag: v0.53.1
     sha: ""
     pullPolicy: IfNotPresent
 
@@ -1659,7 +1657,7 @@ prometheusOperator:
     # image to use for config and rule reloading
     image:
       repository: quay.io/prometheus-operator/prometheus-config-reloader
-      tag: v0.52.1
+      tag: v0.53.1
       sha: ""
 
     # resource config for prometheusConfigReloader
@@ -1743,7 +1741,7 @@ prometheus:
     scheme: ""
 
     ## tlsConfig: TLS configuration to use when scraping the endpoint. For example if using istio mTLS.
-    ## Of type: https://github.com/coreos/prometheus-operator/blob/master/Documentation/api.md#tlsconfig
+    ## Of type: https://github.com/coreos/prometheus-operator/blob/main/Documentation/api.md#tlsconfig
     tlsConfig: {}
 
     bearerTokenFile:
@@ -1997,7 +1995,7 @@ prometheus:
     scheme: ""
 
     ## tlsConfig: TLS configuration to use when scraping the endpoint. For example if using istio mTLS.
-    ## Of type: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#tlsconfig
+    ## Of type: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#tlsconfig
     tlsConfig: {}
 
     bearerTokenFile:
@@ -2020,14 +2018,14 @@ prometheus:
     #   action: replace
 
   ## Settings affecting prometheusSpec
-  ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#prometheusspec
+  ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#prometheusspec
   ##
   prometheusSpec:
     ## If true, pass --storage.tsdb.max-block-duration=2h to prometheus. This is already done if using Thanos
     ##
     disableCompaction: false
     ## APIServerConfig
-    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#apiserverconfig
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#apiserverconfig
     ##
     apiserverConfig: {}
 
@@ -2056,7 +2054,7 @@ prometheus:
     enableAdminAPI: false
 
     ## WebTLSConfig defines the TLS parameters for HTTPS
-    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#webtlsconfig
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#webtlsconfig
     web: {}
 
     # EnableFeatures API enables access to Prometheus disabled features.
@@ -2068,7 +2066,7 @@ prometheus:
     ##
     image:
       repository: quay.io/prometheus/prometheus
-      tag: v2.31.1
+      tag: v2.32.1
       sha: ""
 
     ## Tolerations for use with node taints
@@ -2092,7 +2090,7 @@ prometheus:
     #       app: prometheus
 
     ## Alertmanagers to which alerts will be sent
-    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#alertmanagerendpoints
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#alertmanagerendpoints
     ##
     ## Default configuration will connect to the alertmanager deployed as part of this release
     ##
@@ -2148,13 +2146,13 @@ prometheus:
     configMaps: []
 
     ## QuerySpec defines the query command line flags when starting Prometheus.
-    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#queryspec
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#queryspec
     ##
     query: {}
 
     ## Namespaces to be selected for PrometheusRules discovery.
     ## If nil, select own namespace. Namespaces to be selected for ServiceMonitor discovery.
-    ## See https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#namespaceselector for usage
+    ## See https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#namespaceselector for usage
     ##
     ruleNamespaceSelector: {}
 
@@ -2222,7 +2220,7 @@ prometheus:
     #     prometheus: somelabel
 
     ## Namespaces to be selected for PodMonitor discovery.
-    ## See https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#namespaceselector for usage
+    ## See https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#namespaceselector for usage
     ##
     podMonitorNamespaceSelector: {}
 
@@ -2242,7 +2240,7 @@ prometheus:
     #     prometheus: somelabel
 
     ## Namespaces to be selected for Probe discovery.
-    ## See https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#namespaceselector for usage
+    ## See https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#namespaceselector for usage
     ##
     probeNamespaceSelector: {}
 
@@ -2323,14 +2321,14 @@ prometheus:
     #         - e2e-az2
 
     ## The remote_read spec configuration for Prometheus.
-    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#remotereadspec
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#remotereadspec
     remoteRead: []
     # - url: http://remote1/read
     ## additionalRemoteRead is appended to remoteRead
     additionalRemoteRead: []
 
     ## The remote_write spec configuration for Prometheus.
-    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#remotewritespec
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#remotewritespec
     remoteWrite: []
     # - url: http://remote1/push
     ## additionalRemoteWrite is appended to remoteWrite
@@ -2346,7 +2344,7 @@ prometheus:
     #   memory: 400Mi
 
     ## Prometheus StorageSpec for persistent data
-    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/user-guides/storage.md
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/user-guides/storage.md
     ##
     storageSpec: {}
     ## Using PersistentVolumeClaim
@@ -2462,7 +2460,7 @@ prometheus:
 
     ## SecurityContext holds pod-level security attributes and common container settings.
     ## This defaults to non root user with uid 1000 and gid 2000.
-    ## https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md
+    ## https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md
     ##
     securityContext:
       runAsGroup: 2000
@@ -2477,7 +2475,7 @@ prometheus:
     ## Thanos configuration allows configuring various aspects of a Prometheus server in a Thanos environment.
     ## This section is experimental, it may change significantly without deprecation notice in any release.
     ## This is experimental and may change significantly without backward compatibility in any release.
-    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#thanosspec
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#thanosspec
     ##
     thanos: {}
       # secretProviderClass:
@@ -2695,6 +2693,6 @@ prometheus:
       # matchNames: []
 
     ## Endpoints of the selected pods to be monitored
-    ## https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#podmetricsendpoint
+    ## https://github.com/prometheus-operator/prometheus-operator/blob/main/Documentation/api.md#podmetricsendpoint
     ##
     # podMetricsEndpoints: []

```

## 24.0.1

**Release date:** 2021-12-18

![AppVersion: 0.52.1](https://img.shields.io/static/v1?label=AppVersion&message=0.52.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Fix yaml comments kube prometheus stack (#1611)

### Default value changes

```diff
# No changes in this release
```

## 24.0.0

**Release date:** 2021-12-18

![AppVersion: 0.52.1](https://img.shields.io/static/v1?label=AppVersion&message=0.52.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Use sub-chart ServiceMonitors (#1593)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index b8d31d04..10f8e9c9 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1293,45 +1293,6 @@ kubeProxy:
 ##
 kubeStateMetrics:
   enabled: true
-  serviceMonitor:
-    ## Scrape interval. If not set, the Prometheus default scrape interval is used.
-    ##
-    interval: ""
-    ## Scrape Timeout. If not set, the Prometheus default scrape timeout is used.
-    ##
-    scrapeTimeout: ""
-    ## proxyUrl: URL of a proxy that should be used for scraping.
-    ##
-    proxyUrl: ""
-    ## Override serviceMonitor selector
-    ##
-    selectorOverride: {}
-
-    ## MetricRelabelConfigs to apply to samples after scraping, but before ingestion.
-    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
-    ##
-    metricRelabelings: []
-    # - action: keep
-    #   regex: 'kube_(daemonset|deployment|pod|namespace|node|statefulset).+'
-    #   sourceLabels: [__name__]
-
-    ## RelabelConfigs to apply to samples before scraping
-    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
-    ##
-    relabelings: []
-    # - sourceLabels: [__meta_kubernetes_pod_node_name]
-    #   separator: ;
-    #   regex: ^(.*)$
-    #   targetLabel: nodename
-    #   replacement: $1
-    #   action: replace
-
-    # Keep labels from scraped data, overriding server-side labels
-    honorLabels: true
-
-    # Enable self metrics configuration for Service Monitor
-    selfMonitor:
-      enabled: false
 
 ## Configuration for kube-state-metrics subchart
 ##
@@ -1341,50 +1302,53 @@ kube-state-metrics:
     create: true
   podSecurityPolicy:
     enabled: true
+  prometheus:
+    monitor:
+      enabled: true
+
+      ## Scrape interval. If not set, the Prometheus default scrape interval is used.
+      ##
+      interval: ""
+
+      ## Scrape Timeout. If not set, the Prometheus default scrape timeout is used.
+      ##
+      scrapeTimeout: ""
+
+      ## proxyUrl: URL of a proxy that should be used for scraping.
+      ##
+      proxyUrl: ""
+
+      # Keep labels from scraped data, overriding server-side labels
+      ##
+      honorLabels: true
+
+      ## MetricRelabelConfigs to apply to samples after scraping, but before ingestion.
+      ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
+      ##
+      metricRelabelings: []
+      # - action: keep
+      #   regex: 'kube_(daemonset|deployment|pod|namespace|node|statefulset).+'
+      #   sourceLabels: [__name__]
+
+      ## RelabelConfigs to apply to samples before scraping
+      ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
+      ##
+      relabelings: []
+      # - sourceLabels: [__meta_kubernetes_pod_node_name]
+      #   separator: ;
+      #   regex: ^(.*)$
+      #   targetLabel: nodename
+      #   replacement: $1
+      #   action: replace
+
+  selfMonitor:
+    enabled: false
 
 ## Deploy node exporter as a daemonset to all nodes
 ##
 nodeExporter:
   enabled: true
 
-  ## Use the value configured in prometheus-node-exporter.podLabels
-  ##
-  jobLabel: jobLabel
-
-  serviceMonitor:
-    ## Scrape interval. If not set, the Prometheus default scrape interval is used.
-    ##
-    interval: ""
-
-    ## proxyUrl: URL of a proxy that should be used for scraping.
-    ##
-    proxyUrl: ""
-
-    ## How long until a scrape request times out. If not set, the Prometheus default scape timeout is used.
-    ##
-    scrapeTimeout: ""
-
-    ## MetricRelabelConfigs to apply to samples after scraping, but before ingestion.
-    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
-    ##
-    metricRelabelings: []
-    # - sourceLabels: [__name__]
-    #   separator: ;
-    #   regex: ^node_mountstats_nfs_(event|operations|transport)_.+
-    #   replacement: $1
-    #   action: drop
-
-    ## RelabelConfigs to apply to samples before scraping
-    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
-    ##
-    relabelings: []
-    # - sourceLabels: [__meta_kubernetes_pod_node_name]
-    #   separator: ;
-    #   regex: ^(.*)$
-    #   targetLabel: nodename
-    #   replacement: $1
-    #   action: replace
-
 ## Configuration for prometheus-node-exporter subchart
 ##
 prometheus-node-exporter:
@@ -1398,6 +1362,44 @@ prometheus-node-exporter:
     - --collector.filesystem.ignored-fs-types=^(autofs|binfmt_misc|bpf|cgroup2?|configfs|debugfs|devpts|devtmpfs|fusectl|hugetlbfs|iso9660|mqueue|nsfs|overlay|proc|procfs|pstore|rpc_pipefs|securityfs|selinuxfs|squashfs|sysfs|tracefs)$
   service:
     portName: http-metrics
+  prometheus:
+    monitor:
+      enabled: false
+
+      jobLabel: jobLabel
+
+      ## Scrape interval. If not set, the Prometheus default scrape interval is used.
+      ##
+      interval: ""
+
+      ## How long until a scrape request times out. If not set, the Prometheus default scape timeout is used.
+      ##
+      scrapeTimeout: ""
+
+      ## proxyUrl: URL of a proxy that should be used for scraping.
+      ##
+      proxyUrl: ""
+
+      ## MetricRelabelConfigs to apply to samples after scraping, but before ingestion.
+      ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
+      ##
+      metricRelabelings: []
+      # - sourceLabels: [__name__]
+      #   separator: ;
+      #   regex: ^node_mountstats_nfs_(event|operations|transport)_.+
+      #   replacement: $1
+      #   action: drop
+
+      ## RelabelConfigs to apply to samples before scraping
+      ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
+      ##
+      relabelings: []
+      # - sourceLabels: [__meta_kubernetes_pod_node_name]
+      #   separator: ;
+      #   regex: ^(.*)$
+      #   targetLabel: nodename
+      #   replacement: $1
+      #   action: replace
 
 ## Manages Prometheus and Alertmanager components
 ##

```

## 23.3.2

**Release date:** 2021-12-15

![AppVersion: 0.52.1](https://img.shields.io/static/v1?label=AppVersion&message=0.52.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] bump grafana dependency to 6.19.2 (#1602)

### Default value changes

```diff
# No changes in this release
```

## 23.3.1

**Release date:** 2021-12-14

![AppVersion: 0.52.1](https://img.shields.io/static/v1?label=AppVersion&message=0.52.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Bump to prometheus-operator 0.52.1 (#1598)

### Default value changes

```diff
# No changes in this release
```

## 23.3.0

**Release date:** 2021-12-14

![AppVersion: 0.52.0](https://img.shields.io/static/v1?label=AppVersion&message=0.52.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Remove app label on ruleSelector (#1113)

### Default value changes

```diff
# No changes in this release
```

## 23.2.0

**Release date:** 2021-12-10

![AppVersion: 0.52.1](https://img.shields.io/static/v1?label=AppVersion&message=0.52.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] bump grafana dependency to 6.19.0 (#1586)

### Default value changes

```diff
# No changes in this release
```

## 23.1.6

**Release date:** 2021-12-09

![AppVersion: 0.52.1](https://img.shields.io/static/v1?label=AppVersion&message=0.52.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Upgrade prometheus-operator to v0.52.1 (#1584)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 8be9f182..b8d31d04 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1639,7 +1639,7 @@ prometheusOperator:
   ##
   image:
     repository: quay.io/prometheus-operator/prometheus-operator
-    tag: v0.52.0
+    tag: v0.52.1
     sha: ""
     pullPolicy: IfNotPresent
 
@@ -1657,7 +1657,7 @@ prometheusOperator:
     # image to use for config and rule reloading
     image:
       repository: quay.io/prometheus-operator/prometheus-config-reloader
-      tag: v0.52.0
+      tag: v0.52.1
       sha: ""
 
     # resource config for prometheusConfigReloader

```

## 23.1.5

**Release date:** 2021-12-09

![AppVersion: 0.52.0](https://img.shields.io/static/v1?label=AppVersion&message=0.52.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] add alertmanager unittest (#1575)

### Default value changes

```diff
# No changes in this release
```

## 23.1.4

**Release date:** 2021-12-09

![AppVersion: 0.52.0](https://img.shields.io/static/v1?label=AppVersion&message=0.52.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Fix alertmanager (#1581)

### Default value changes

```diff
# No changes in this release
```

## 23.1.3

**Release date:** 2021-12-08

![AppVersion: 0.52.0](https://img.shields.io/static/v1?label=AppVersion&message=0.52.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] fix double annotation (#1579)

### Default value changes

```diff
# No changes in this release
```

## 23.1.2

**Release date:** 2021-12-08

![AppVersion: 0.52.0](https://img.shields.io/static/v1?label=AppVersion&message=0.52.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] update kube-state-metrics dependency to 4.1.1 (#1577)

### Default value changes

```diff
# No changes in this release
```

## 23.1.1

**Release date:** 2021-12-08

![AppVersion: 0.52.0](https://img.shields.io/static/v1?label=AppVersion&message=0.52.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] update grafana dependency to 6.18.2 (#1570)

### Default value changes

```diff
# No changes in this release
```

## 23.1.0

**Release date:** 2021-12-07

![AppVersion: 0.52.0](https://img.shields.io/static/v1?label=AppVersion&message=0.52.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add support for removing old datasources in values file (#1528)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 46910d44..8be9f182 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -720,6 +720,10 @@ grafana:
   #   configMap: certs-configmap
   #   readOnly: true
 
+  deleteDatasources: []
+  # - name: example-datasource
+  #   orgId: 1
+
   ## Configure additional grafana datasources (passed through tpl)
   ## ref: http://docs.grafana.org/administration/provisioning/#datasources
   additionalDataSources: []

```

## 23.0.0

**Release date:** 2021-12-07

![AppVersion: 0.52.0](https://img.shields.io/static/v1?label=AppVersion&message=0.52.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Use Istio service port naming convention (#1407)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 21185425..46910d44 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -598,7 +598,7 @@ alertmanager:
 
     ## PortName to use for Alert Manager.
     ##
-    portName: "web"
+    portName: "http-web"
 
     ## ClusterAdvertiseAddress is the explicit address to advertise in cluster. Needs to be provided for non RFC1918 [1] (public) addresses. [1] RFC1918: https://tools.ietf.org/html/rfc1918
     ##
@@ -739,7 +739,7 @@ grafana:
   ## Passed to grafana subchart and used by servicemonitor below
   ##
   service:
-    portName: service
+    portName: http-web
 
   ## If true, create a serviceMonitor for grafana
   ##
@@ -1392,6 +1392,8 @@ prometheus-node-exporter:
   extraArgs:
     - --collector.filesystem.ignored-mount-points=^/(dev|proc|sys|var/lib/docker/.+|var/lib/kubelet/.+)($|/)
     - --collector.filesystem.ignored-fs-types=^(autofs|binfmt_misc|bpf|cgroup2?|configfs|debugfs|devpts|devtmpfs|fusectl|hugetlbfs|iso9660|mqueue|nsfs|overlay|proc|procfs|pstore|rpc_pipefs|securityfs|selinuxfs|squashfs|sysfs|tracefs)$
+  service:
+    portName: http-metrics
 
 ## Manages Prometheus and Alertmanager components
 ##
@@ -2490,7 +2492,7 @@ prometheus:
 
     ## PortName to use for Prometheus.
     ##
-    portName: "web"
+    portName: "http-web"
 
     ## ArbitraryFSAccessThroughSMs configures whether configuration based on a service monitor can access arbitrary files
     ## on the file system of the Prometheus container e.g. bearer token files.

```

## 22.1.0

**Release date:** 2021-12-07

![AppVersion: 0.52.0](https://img.shields.io/static/v1?label=AppVersion&message=0.52.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Update Grafana chart version to 6.18.* and bump prometheus-node-exporter (#1562)

### Default value changes

```diff
# No changes in this release
```

## 22.0.0

**Release date:** 2021-12-06

![AppVersion: 0.52.0](https://img.shields.io/static/v1?label=AppVersion&message=0.52.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Bump dependencies (#1518)

### Default value changes

```diff
# No changes in this release
```

## 21.0.5

**Release date:** 2021-12-06

![AppVersion: 0.52.0](https://img.shields.io/static/v1?label=AppVersion&message=0.52.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Removing me as maintainer (#1560)

### Default value changes

```diff
# No changes in this release
```

## 21.0.4

**Release date:** 2021-12-06

![AppVersion: 0.52.0](https://img.shields.io/static/v1?label=AppVersion&message=0.52.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Rename alertmanager's serviceperreplica selector key (#1524)

### Default value changes

```diff
# No changes in this release
```

## 21.0.3

**Release date:** 2021-12-03

![AppVersion: 0.52.0](https://img.shields.io/static/v1?label=AppVersion&message=0.52.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Bump grafana to 6.17.10 (#1544)

### Default value changes

```diff
# No changes in this release
```

## 21.0.2

**Release date:** 2021-12-02

![AppVersion: 0.52.0](https://img.shields.io/static/v1?label=AppVersion&message=0.52.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] missplaced end tag (#1417)

### Default value changes

```diff
# No changes in this release
```

## 21.0.1

**Release date:** 2021-12-02

![AppVersion: 0.52.0](https://img.shields.io/static/v1?label=AppVersion&message=0.52.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Apply scrapeTimeout config to all kubelet service monitors (#1408)

### Default value changes

```diff
# No changes in this release
```

## 21.0.0

**Release date:** 2021-11-27

![AppVersion: 0.52.0](https://img.shields.io/static/v1?label=AppVersion&message=0.52.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Separate config reloader limit and request (#1375)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 2a4fdf73..21185425 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1645,20 +1645,23 @@ prometheusOperator:
   ##
   # alertmanagerDefaultBaseImage: quay.io/prometheus/alertmanager
 
-  ## Prometheus-config-reloader image to use for config and rule reloading
+  ## Prometheus-config-reloader
   ##
-  prometheusConfigReloaderImage:
-    repository: quay.io/prometheus-operator/prometheus-config-reloader
-    tag: v0.52.0
-    sha: ""
+  prometheusConfigReloader:
+    # image to use for config and rule reloading
+    image:
+      repository: quay.io/prometheus-operator/prometheus-config-reloader
+      tag: v0.52.0
+      sha: ""
 
-  ## Set the prometheus config reloader side-car CPU limit
-  ##
-  configReloaderCpu: 100m
-
-  ## Set the prometheus config reloader side-car memory limit
-  ##
-  configReloaderMemory: 50Mi
+    # resource config for prometheusConfigReloader
+    resources:
+      requests:
+        cpu: 100m
+        memory: 50Mi
+      limits:
+        cpu: 100m
+        memory: 50Mi
 
   ## Thanos side-car image when configured
   ##

```

## 20.1.0

**Release date:** 2021-11-27

![AppVersion: 0.52.0](https://img.shields.io/static/v1?label=AppVersion&message=0.52.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Bump Kube-prometheus-stack versions (#1397)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 04dddd78..2a4fdf73 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -398,7 +398,7 @@ alertmanager:
     ##
     image:
       repository: quay.io/prometheus/alertmanager
-      tag: v0.22.2
+      tag: v0.23.0
       sha: ""
 
     ## If true then the user will be responsible to provide a secret with alertmanager configuration
@@ -1664,7 +1664,7 @@ prometheusOperator:
   ##
   thanosImage:
     repository: quay.io/thanos/thanos
-    tag: v0.17.2
+    tag: v0.23.1
     sha: ""
 
   ## Set a Field Selector to filter watched secrets
@@ -2057,7 +2057,7 @@ prometheus:
     ##
     image:
       repository: quay.io/prometheus/prometheus
-      tag: v2.28.1
+      tag: v2.31.1
       sha: ""
 
     ## Tolerations for use with node taints

```

## 20.0.1

**Release date:** 2021-11-10

![AppVersion: 0.52.0](https://img.shields.io/static/v1?label=AppVersion&message=0.52.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] removed app:alertmanager label (#1494)

### Default value changes

```diff
# No changes in this release
```

## 20.0.0

**Release date:** 2021-11-10

![AppVersion: 0.52.0](https://img.shields.io/static/v1?label=AppVersion&message=0.52.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Prometheus operator version upgrade v0.52.0 (#1485)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index cc0f8bbd..04dddd78 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1633,7 +1633,7 @@ prometheusOperator:
   ##
   image:
     repository: quay.io/prometheus-operator/prometheus-operator
-    tag: v0.50.0
+    tag: v0.52.0
     sha: ""
     pullPolicy: IfNotPresent
 
@@ -1649,7 +1649,7 @@ prometheusOperator:
   ##
   prometheusConfigReloaderImage:
     repository: quay.io/prometheus-operator/prometheus-config-reloader
-    tag: v0.50.0
+    tag: v0.52.0
     sha: ""
 
   ## Set the prometheus config reloader side-car CPU limit

```

## 19.3.0

**Release date:** 2021-11-09

![AppVersion: 0.50.0](https://img.shields.io/static/v1?label=AppVersion&message=0.50.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Ability to specify kubelet-service again (#1434)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index f5953614..cc0f8bbd 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1541,6 +1541,8 @@ prometheusOperator:
   kubeletService:
     enabled: true
     namespace: kube-system
+    ## Use '{{ template "kube-prometheus-stack.fullname" . }}-kubelet' by default
+    name: ""
 
   ## Create a servicemonitor for the operator
   ##

```

## 19.2.3

**Release date:** 2021-11-04

![AppVersion: 0.50.0](https://img.shields.io/static/v1?label=AppVersion&message=0.50.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Update Grafana chart version to 6.17.5 (#1487)

### Default value changes

```diff
# No changes in this release
```

## 19.2.2

**Release date:** 2021-10-19

![AppVersion: 0.50.0](https://img.shields.io/static/v1?label=AppVersion&message=0.50.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Update Grafana chart version to 6.17.2 (#1447)

### Default value changes

```diff
# No changes in this release
```

## 19.2.1

**Release date:** 2021-10-19

![AppVersion: 0.50.0](https://img.shields.io/static/v1?label=AppVersion&message=0.50.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] release new version (#1446)

### Default value changes

```diff
# No changes in this release
```

## 19.2.0

**Release date:** 2021-10-19

![AppVersion: 0.50.0](https://img.shields.io/static/v1?label=AppVersion&message=0.50.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Bump Grafana chart version to 6.17.* (#1439)

### Default value changes

```diff
# No changes in this release
```

## 19.1.0

**Release date:** 2021-10-14

![AppVersion: 0.50.0](https://img.shields.io/static/v1?label=AppVersion&message=0.50.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Move ServiceMonitor back to main namespace (#1383)

### Default value changes

```diff
# No changes in this release
```

## 19.0.3

**Release date:** 2021-10-12

![AppVersion: 0.50.0](https://img.shields.io/static/v1?label=AppVersion&message=0.50.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] 1269: Template the relabellings which are passed into the service monitors (#1271)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index a8ba31f8..f5953614 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -366,14 +366,16 @@ alertmanager:
 
     bearerTokenFile:
 
-    ## Metric relabel configs to apply to samples before ingestion.
+    ## MetricRelabelConfigs to apply to samples after scraping, but before ingestion.
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
     ##
     metricRelabelings: []
     # - action: keep
     #   regex: 'kube_(daemonset|deployment|pod|namespace|node|statefulset).+'
     #   sourceLabels: [__name__]
 
-    #   relabel configs to apply to samples before ingestion.
+    ## RelabelConfigs to apply to samples before scraping
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
     ##
     relabelings: []
     # - sourceLabels: [__meta_kubernetes_pod_node_name]
@@ -751,14 +753,17 @@ grafana:
     # in grafana.ini
     path: "/metrics"
 
-    ## Metric relabel configs to apply to samples before ingestion.
+
+    ## MetricRelabelConfigs to apply to samples after scraping, but before ingestion.
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
     ##
     metricRelabelings: []
     # - action: keep
     #   regex: 'kube_(daemonset|deployment|pod|namespace|node|statefulset).+'
     #   sourceLabels: [__name__]
 
-    #   relabel configs to apply to samples before ingestion.
+    ## RelabelConfigs to apply to samples before scraping
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
     ##
     relabelings: []
     # - sourceLabels: [__meta_kubernetes_pod_node_name]
@@ -789,12 +794,17 @@ kubeApiServer:
         component: apiserver
         provider: kubernetes
 
-    ## Metric relabel configs to apply to samples before ingestion.
+    ## MetricRelabelConfigs to apply to samples after scraping, but before ingestion.
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
     ##
     metricRelabelings: []
     # - action: keep
     #   regex: 'kube_(daemonset|deployment|pod|namespace|node|statefulset).+'
     #   sourceLabels: [__name__]
+
+    ## RelabelConfigs to apply to samples before scraping
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
+    ##
     relabelings: []
     # - sourceLabels:
     #     - __meta_kubernetes_namespace
@@ -839,7 +849,9 @@ kubelet:
     resource: false
     # From kubernetes 1.18, /metrics/resource/v1alpha1 renamed to /metrics/resource
     resourcePath: "/metrics/resource/v1alpha1"
-    ## Metric relabellings to apply to samples before ingestion
+
+    ## MetricRelabelConfigs to apply to samples after scraping, but before ingestion.
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
     ##
     cAdvisorMetricRelabelings: []
     # - sourceLabels: [__name__, image]
@@ -853,7 +865,8 @@ kubelet:
     #   replacement: $1
     #   action: drop
 
-    ## Metric relabellings to apply to samples before ingestion
+    ## MetricRelabelConfigs to apply to samples after scraping, but before ingestion.
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
     ##
     probesMetricRelabelings: []
     # - sourceLabels: [__name__, image]
@@ -867,9 +880,10 @@ kubelet:
     #   replacement: $1
     #   action: drop
 
-    #   relabel configs to apply to samples before ingestion.
-    #   metrics_path is required to match upstream rules and charts
+    ## RelabelConfigs to apply to samples before scraping
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
     ##
+    ## metrics_path is required to match upstream rules and charts
     cAdvisorRelabelings:
       - sourceLabels: [__metrics_path__]
         targetLabel: metrics_path
@@ -880,6 +894,9 @@ kubelet:
     #   replacement: $1
     #   action: replace
 
+    ## RelabelConfigs to apply to samples before scraping
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
+    ##
     probesRelabelings:
       - sourceLabels: [__metrics_path__]
         targetLabel: metrics_path
@@ -890,6 +907,9 @@ kubelet:
     #   replacement: $1
     #   action: replace
 
+    ## RelabelConfigs to apply to samples before scraping
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
+    ##
     resourceRelabelings:
       - sourceLabels: [__metrics_path__]
         targetLabel: metrics_path
@@ -900,6 +920,9 @@ kubelet:
     #   replacement: $1
     #   action: replace
 
+    ## MetricRelabelConfigs to apply to samples after scraping, but before ingestion.
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
+    ##
     metricRelabelings: []
     # - sourceLabels: [__name__, image]
     #   separator: ;
@@ -912,9 +935,10 @@ kubelet:
     #   replacement: $1
     #   action: drop
 
-    #   relabel configs to apply to samples before ingestion.
-    #   metrics_path is required to match upstream rules and charts
+    ## RelabelConfigs to apply to samples before scraping
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
     ##
+    ## metrics_path is required to match upstream rules and charts
     relabelings:
       - sourceLabels: [__metrics_path__]
         targetLabel: metrics_path
@@ -967,14 +991,16 @@ kubeControllerManager:
     # Name of the server to use when validating TLS certificate
     serverName: null
 
-    ## Metric relabel configs to apply to samples before ingestion.
+    ## MetricRelabelConfigs to apply to samples after scraping, but before ingestion.
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
     ##
     metricRelabelings: []
     # - action: keep
     #   regex: 'kube_(daemonset|deployment|pod|namespace|node|statefulset).+'
     #   sourceLabels: [__name__]
 
-    #   relabel configs to apply to samples before ingestion.
+    ## RelabelConfigs to apply to samples before scraping
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
     ##
     relabelings: []
     # - sourceLabels: [__meta_kubernetes_pod_node_name]
@@ -1002,14 +1028,16 @@ coreDns:
     ##
     proxyUrl: ""
 
-    ## Metric relabel configs to apply to samples before ingestion.
+    ## MetricRelabelConfigs to apply to samples after scraping, but before ingestion.
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
     ##
     metricRelabelings: []
     # - action: keep
     #   regex: 'kube_(daemonset|deployment|pod|namespace|node|statefulset).+'
     #   sourceLabels: [__name__]
 
-    #   relabel configs to apply to samples before ingestion.
+    ## RelabelConfigs to apply to samples before scraping
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
     ##
     relabelings: []
     # - sourceLabels: [__meta_kubernetes_pod_node_name]
@@ -1041,14 +1069,16 @@ kubeDns:
     ##
     proxyUrl: ""
 
-    ## Metric relabel configs to apply to samples before ingestion.
+    ## MetricRelabelConfigs to apply to samples after scraping, but before ingestion.
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
     ##
     metricRelabelings: []
     # - action: keep
     #   regex: 'kube_(daemonset|deployment|pod|namespace|node|statefulset).+'
     #   sourceLabels: [__name__]
 
-    #   relabel configs to apply to samples before ingestion.
+    ## RelabelConfigs to apply to samples before scraping
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
     ##
     relabelings: []
     # - sourceLabels: [__meta_kubernetes_pod_node_name]
@@ -1057,12 +1087,17 @@ kubeDns:
     #   targetLabel: nodename
     #   replacement: $1
     #   action: replace
+
+    ## MetricRelabelConfigs to apply to samples after scraping, but before ingestion.
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
+    ##
     dnsmasqMetricRelabelings: []
     # - action: keep
     #   regex: 'kube_(daemonset|deployment|pod|namespace|node|statefulset).+'
     #   sourceLabels: [__name__]
 
-    #   relabel configs to apply to samples before ingestion.
+    ## RelabelConfigs to apply to samples before scraping
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
     ##
     dnsmasqRelabelings: []
     # - sourceLabels: [__meta_kubernetes_pod_node_name]
@@ -1119,14 +1154,16 @@ kubeEtcd:
     certFile: ""
     keyFile: ""
 
-    ## Metric relabel configs to apply to samples before ingestion.
+    ## MetricRelabelConfigs to apply to samples after scraping, but before ingestion.
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
     ##
     metricRelabelings: []
     # - action: keep
     #   regex: 'kube_(daemonset|deployment|pod|namespace|node|statefulset).+'
     #   sourceLabels: [__name__]
 
-    #   relabel configs to apply to samples before ingestion.
+    ## RelabelConfigs to apply to samples before scraping
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
     ##
     relabelings: []
     # - sourceLabels: [__meta_kubernetes_pod_node_name]
@@ -1177,14 +1214,16 @@ kubeScheduler:
     ## Name of the server to use when validating TLS certificate
     serverName: null
 
-    ## Metric relabel configs to apply to samples before ingestion.
+    ## MetricRelabelConfigs to apply to samples after scraping, but before ingestion.
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
     ##
     metricRelabelings: []
     # - action: keep
     #   regex: 'kube_(daemonset|deployment|pod|namespace|node|statefulset).+'
     #   sourceLabels: [__name__]
 
-    #   relabel configs to apply to samples before ingestion.
+    ## RelabelConfigs to apply to samples before scraping
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
     ##
     relabelings: []
     # - sourceLabels: [__meta_kubernetes_pod_node_name]
@@ -1229,14 +1268,16 @@ kubeProxy:
     ##
     https: false
 
-    ## Metric relabel configs to apply to samples before ingestion.
+    ## MetricRelabelConfigs to apply to samples after scraping, but before ingestion.
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
     ##
     metricRelabelings: []
     # - action: keep
     #   regex: 'kube_(daemonset|deployment|pod|namespace|node|statefulset).+'
     #   sourceLabels: [__name__]
 
-    #   relabel configs to apply to samples before ingestion.
+    ## RelabelConfigs to apply to samples before scraping
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
     ##
     relabelings: []
     # - action: keep
@@ -1262,14 +1303,16 @@ kubeStateMetrics:
     ##
     selectorOverride: {}
 
-    ## Metric relabel configs to apply to samples before ingestion.
+    ## MetricRelabelConfigs to apply to samples after scraping, but before ingestion.
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
     ##
     metricRelabelings: []
     # - action: keep
     #   regex: 'kube_(daemonset|deployment|pod|namespace|node|statefulset).+'
     #   sourceLabels: [__name__]
 
-    #   relabel configs to apply to samples before ingestion.
+    ## RelabelConfigs to apply to samples before scraping
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
     ##
     relabelings: []
     # - sourceLabels: [__meta_kubernetes_pod_node_name]
@@ -1317,7 +1360,8 @@ nodeExporter:
     ##
     scrapeTimeout: ""
 
-    ## Metric relabel configs to apply to samples before ingestion.
+    ## MetricRelabelConfigs to apply to samples after scraping, but before ingestion.
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
     ##
     metricRelabelings: []
     # - sourceLabels: [__name__]
@@ -1326,7 +1370,8 @@ nodeExporter:
     #   replacement: $1
     #   action: drop
 
-    ## relabel configs to apply to samples before ingestion.
+    ## RelabelConfigs to apply to samples before scraping
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#relabelconfig
     ##
     relabelings: []
     # - sourceLabels: [__meta_kubernetes_pod_node_name]

```

## 19.0.2

**Release date:** 2021-10-04

![AppVersion: 0.50.0](https://img.shields.io/static/v1?label=AppVersion&message=0.50.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Support for secrets store CSI driver  (#1158)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 50ef9724..a8ba31f8 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -2422,6 +2422,13 @@ prometheus:
     ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#thanosspec
     ##
     thanos: {}
+      # secretProviderClass:
+      #   provider: gcp
+      #   parameters:
+      #     secrets: |
+      #       - resourceName: "projects/$PROJECT_ID/secrets/testsecret/versions/latest"
+      #         fileName: "objstore.yaml"
+      # objectStorageConfigFile: /var/secrets/object-store.yaml
 
     ## Containers allows injecting additional containers. This is meant to allow adding an authentication proxy to a Prometheus pod.
     ## if using proxy extraContainer update targetPort with proxy container port

```

## 19.0.1

**Release date:** 2021-09-30

![AppVersion: 0.50.0](https://img.shields.io/static/v1?label=AppVersion&message=0.50.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Bump kube-state-metrics dependency version (#1378)

### Default value changes

```diff
# No changes in this release
```

## 19.0.0

**Release date:** 2021-09-28

![AppVersion: 0.50.0](https://img.shields.io/static/v1?label=AppVersion&message=0.50.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Use namespaceOverrides when set (#1343)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index cb12c344..50ef9724 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1261,9 +1261,6 @@ kubeStateMetrics:
     ## Override serviceMonitor selector
     ##
     selectorOverride: {}
-    ## Override namespace selector
-    ##
-    namespaceOverride: ""
 
     ## Metric relabel configs to apply to samples before ingestion.
     ##

```

## 18.1.1

**Release date:** 2021-09-28

![AppVersion: 0.50.0](https://img.shields.io/static/v1?label=AppVersion&message=0.50.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Fix drilldown urls for Grafana dashboards (#1377)

### Default value changes

```diff
# No changes in this release
```

## 18.1.0

**Release date:** 2021-09-24

![AppVersion: 0.50.0](https://img.shields.io/static/v1?label=AppVersion&message=0.50.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Support all of the enforce*Limit settings (#1373)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 2fbe154d..cb12c344 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -2473,6 +2473,28 @@ prometheus:
     ## number of samples/series under the desired limit. Note that if SampleLimit is lower that value will be taken instead.
     enforcedSampleLimit: false
 
+    ## EnforcedTargetLimit defines a global limit on the number of scraped targets. This overrides any TargetLimit set
+    ## per ServiceMonitor or/and PodMonitor. It is meant to be used by admins to enforce the TargetLimit to keep the overall
+    ## number of targets under the desired limit. Note that if TargetLimit is lower, that value will be taken instead, except
+    ## if either value is zero, in which case the non-zero value will be used. If both values are zero, no limit is enforced.
+    enforcedTargetLimit: false
+
+
+    ## Per-scrape limit on number of labels that will be accepted for a sample. If more than this number of labels are present
+    ## post metric-relabeling, the entire scrape will be treated as failed. 0 means no limit. Only valid in Prometheus versions
+    ## 2.27.0 and newer.
+    enforcedLabelLimit: false
+
+    ## Per-scrape limit on length of labels name that will be accepted for a sample. If a label name is longer than this number
+    ## post metric-relabeling, the entire scrape will be treated as failed. 0 means no limit. Only valid in Prometheus versions
+    ## 2.27.0 and newer.
+    enforcedLabelNameLengthLimit: false
+
+    ## Per-scrape limit on length of labels value that will be accepted for a sample. If a label value is longer than this
+    ## number post metric-relabeling, the entire scrape will be treated as failed. 0 means no limit. Only valid in Prometheus
+    ## versions 2.27.0 and newer.
+    enforcedLabelValueLengthLimit: false
+
     ## AllowOverlappingBlocks enables vertical compaction and vertical query merge in Prometheus. This is still experimental
     ## in Prometheus so it may change in any upcoming release.
     allowOverlappingBlocks: false

```

## 18.0.13

**Release date:** 2021-09-24

![AppVersion: 0.50.0](https://img.shields.io/static/v1?label=AppVersion&message=0.50.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] bump grafana to 6.16.9 (#1367)

### Default value changes

```diff
# No changes in this release
```

## 18.0.12

**Release date:** 2021-09-20

![AppVersion: 0.50.0](https://img.shields.io/static/v1?label=AppVersion&message=0.50.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Allow to set a timezone for the default grafana dashboards (#1326)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 34b78859..2fbe154d 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -631,10 +631,15 @@ grafana:
   ##
   forceDeployDashboards: false
 
-  ## Deploy default dashboards.
+  ## Deploy default dashboards
   ##
   defaultDashboardsEnabled: true
 
+  ## Timezone for the default dashboards
+  ## Other options are: browser or a specific timezone, i.e. Europe/Luxembourg
+  ##
+  defaultDashboardsTimezone: utc
+
   adminPassword: prom-operator
 
   ingress:

```

## 18.0.11

**Release date:** 2021-09-20

![AppVersion: 0.50.0](https://img.shields.io/static/v1?label=AppVersion&message=0.50.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Fix alert manager soft anti-affinity configuration (#1282)

### Default value changes

```diff
# No changes in this release
```

## 18.0.10

**Release date:** 2021-09-17

![AppVersion: 0.50.0](https://img.shields.io/static/v1?label=AppVersion&message=0.50.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Configure Thanos Sidecar monitoring (#1339)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 6dfc1bf1..34b78859 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1649,18 +1649,50 @@ prometheus:
     enabled: false
     annotations: {}
     labels: {}
+
+    ## Service type
+    ##
+    type: ClusterIP
+
+    ## gRPC port config
     portName: grpc
     port: 10901
     targetPort: "grpc"
+
+    ## HTTP port config (for metrics)
+    httpPortName: http
+    httpPort: 10902
+    targetHttpPort: "http"
+
+    ## ClusterIP to assign
+    # Default is to make this a headless service ("None")
     clusterIP: "None"
 
-    ## Service type
-    ##
-    type: ClusterIP
-
-    ## Port to expose on each node
+    ## Port to expose on each node, if service type is NodePort
     ##
     nodePort: 30901
+    httpNodePort: 30902
+
+  # ServiceMonitor to scrape Sidecar metrics
+  # Needs thanosService to be enabled as well
+  thanosServiceMonitor:
+    enabled: false
+    interval: ""
+
+    ## scheme: HTTP scheme to use for scraping. Can be used with `tlsConfig` for example if using istio mTLS.
+    scheme: ""
+
+    ## tlsConfig: TLS configuration to use when scraping the endpoint. For example if using istio mTLS.
+    ## Of type: https://github.com/coreos/prometheus-operator/blob/master/Documentation/api.md#tlsconfig
+    tlsConfig: {}
+
+    bearerTokenFile:
+
+    ## Metric relabel configs to apply to samples before ingestion.
+    metricRelabelings: []
+
+    ## relabel configs to apply to samples before ingestion.
+    relabelings: []
 
   # Service for external access to sidecar
   # Enabling this creates a service to expose thanos-sidecar outside the cluster.
@@ -1668,11 +1700,18 @@ prometheus:
     enabled: false
     annotations: {}
     labels: {}
+    loadBalancerIP: ""
+    loadBalancerSourceRanges: []
+
+    ## gRPC port config
     portName: grpc
     port: 10901
     targetPort: "grpc"
-    loadBalancerIP: ""
-    loadBalancerSourceRanges: []
+
+    ## HTTP port config (for metrics)
+    httpPortName: http
+    httpPort: 10902
+    targetHttpPort: "http"
 
     ## Service type
     ##
@@ -1681,6 +1720,7 @@ prometheus:
     ## Port to expose on each node
     ##
     nodePort: 30901
+    httpNodePort: 30902
 
   ## Configuration for Prometheus service
   ##

```

## 18.0.9

**Release date:** 2021-09-17

![AppVersion: 0.50.0](https://img.shields.io/static/v1?label=AppVersion&message=0.50.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack]: adds templating for .Values.prometheus.thanosIngress.tls (#1341)

### Default value changes

```diff
# No changes in this release
```

## 18.0.8

**Release date:** 2021-09-15

![AppVersion: 0.50.0](https://img.shields.io/static/v1?label=AppVersion&message=0.50.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Add scrapeTimeout in kubeStateMetrics [kube-prometheus-stack] (#1346)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 8631157c..6dfc1bf1 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1247,6 +1247,9 @@ kubeStateMetrics:
     ## Scrape interval. If not set, the Prometheus default scrape interval is used.
     ##
     interval: ""
+    ## Scrape Timeout. If not set, the Prometheus default scrape timeout is used.
+    ##
+    scrapeTimeout: ""
     ## proxyUrl: URL of a proxy that should be used for scraping.
     ##
     proxyUrl: ""

```

## 18.0.7

**Release date:** 2021-09-14

![AppVersion: 0.50.0](https://img.shields.io/static/v1?label=AppVersion&message=0.50.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Optional honorLabels for kube-state-metrics ServiceMonitor (#1327)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index ee14ecb1..8631157c 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1274,6 +1274,9 @@ kubeStateMetrics:
     #   replacement: $1
     #   action: replace
 
+    # Keep labels from scraped data, overriding server-side labels
+    honorLabels: true
+
     # Enable self metrics configuration for Service Monitor
     selfMonitor:
       enabled: false

```

## 18.0.6

**Release date:** 2021-09-10

![AppVersion: 0.50.0](https://img.shields.io/static/v1?label=AppVersion&message=0.50.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] bump grafana to 6.16.* (#1318)

### Default value changes

```diff
# No changes in this release
```

## 18.0.5

**Release date:** 2021-09-05

![AppVersion: 0.50.0](https://img.shields.io/static/v1?label=AppVersion&message=0.50.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Replace admission webhook image (#1280)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 884edc9e..ee14ecb1 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1369,9 +1369,9 @@ prometheusOperator:
     patch:
       enabled: true
       image:
-        repository: jettech/kube-webhook-certgen
-        tag: v1.5.2
-        sha: ""
+        repository: k8s.gcr.io/ingress-nginx/kube-webhook-certgen
+        tag: v1.0
+        sha: "f3b6b39a6062328c095337b4cadcefd1612348fdd5190b1dcbcb9b9e90bd8068"
         pullPolicy: IfNotPresent
       resources: {}
       ## Provide a priority class name to the webhook patching job

```

## 18.0.4

**Release date:** 2021-09-03

![AppVersion: 0.50.0](https://img.shields.io/static/v1?label=AppVersion&message=0.50.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] add option to configure Prometheus image using tag and/or sha256 (#1243)

### Default value changes

```diff
# No changes in this release
```

## 18.0.3

**Release date:** 2021-09-02

![AppVersion: 0.50.0](https://img.shields.io/static/v1?label=AppVersion&message=0.50.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* `[kube-prometheus-stack]` add option to override the allowUiUpdates for grafana dashboards (#1292)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index fdceaf6d..884edc9e 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -683,6 +683,8 @@ grafana:
           enabled: false
         etcd:
           enabled: false
+      provider:
+        allowUiUpdates: false
     datasources:
       enabled: true
       defaultDatasourceEnabled: true

```

## 18.0.2

**Release date:** 2021-08-29

![AppVersion: 0.50.0](https://img.shields.io/static/v1?label=AppVersion&message=0.50.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Fix alert manager complex config templating (#1289)

### Default value changes

```diff
# No changes in this release
```

## 18.0.1

**Release date:** 2021-08-24

![AppVersion: 0.50.0](https://img.shields.io/static/v1?label=AppVersion&message=0.50.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Bump Grafana chart version (#1252) (#1263)

### Default value changes

```diff
# No changes in this release
```

## 18.0.0

**Release date:** 2021-08-17

![AppVersion: 0.50.0](https://img.shields.io/static/v1?label=AppVersion&message=0.50.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Upgrade prometheus-operator to v0.50.0 (#1257)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 08893818..fdceaf6d 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1576,7 +1576,7 @@ prometheusOperator:
   ##
   image:
     repository: quay.io/prometheus-operator/prometheus-operator
-    tag: v0.49.0
+    tag: v0.50.0
     sha: ""
     pullPolicy: IfNotPresent
 
@@ -1592,7 +1592,7 @@ prometheusOperator:
   ##
   prometheusConfigReloaderImage:
     repository: quay.io/prometheus-operator/prometheus-config-reloader
-    tag: v0.49.0
+    tag: v0.50.0
     sha: ""
 
   ## Set the prometheus config reloader side-car CPU limit

```

## 17.2.2

**Release date:** 2021-08-15

![AppVersion: 0.49.0](https://img.shields.io/static/v1?label=AppVersion&message=0.49.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add default ingress pathType for prometheus (#1249)

### Default value changes

```diff
# No changes in this release
```

## 17.2.1

**Release date:** 2021-08-11

![AppVersion: 0.49.0](https://img.shields.io/static/v1?label=AppVersion&message=0.49.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* kube-prometheus-stack: Use stable API for PodDisruptionBudget if 1.21+ clusters (#1240)

### Default value changes

```diff
# No changes in this release
```

## 17.2.0

**Release date:** 2021-08-10

![AppVersion: 0.49.0](https://img.shields.io/static/v1?label=AppVersion&message=0.49.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Support multicluster dashboards only for etcd (#1132)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 8ed4b4fe..08893818 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -197,7 +197,7 @@ alertmanager:
   #       {{- $root := . -}}
   #       {{ range .Alerts }}
   #         *Alert:* {{ .Annotations.summary }} - `{{ .Labels.severity }}`
-  #         *Cluster:*  {{ template "cluster" $root }}
+  #         *Cluster:* {{ template "cluster" $root }}
   #         *Description:* {{ .Annotations.description }}
   #         *Graph:* <{{ .GeneratorURL }}|:chart_with_upwards_trend:>
   #         *Runbook:* <{{ .Annotations.runbook }}|:spiral_note_pad:>
@@ -366,14 +366,14 @@ alertmanager:
 
     bearerTokenFile:
 
-    ## 	metric relabel configs to apply to samples before ingestion.
+    ## Metric relabel configs to apply to samples before ingestion.
     ##
     metricRelabelings: []
     # - action: keep
     #   regex: 'kube_(daemonset|deployment|pod|namespace|node|statefulset).+'
     #   sourceLabels: [__name__]
 
-    # 	relabel configs to apply to samples before ingestion.
+    #   relabel configs to apply to samples before ingestion.
     ##
     relabelings: []
     # - sourceLabels: [__meta_kubernetes_pod_node_name]
@@ -486,11 +486,11 @@ alertmanager:
     #   selector: {}
 
 
-    ## 	The external URL the Alertmanager instances will be available under. This is necessary to generate correct URLs. This is necessary if Alertmanager is not served from root of a DNS name.	string	false
+    ## The external URL the Alertmanager instances will be available under. This is necessary to generate correct URLs. This is necessary if Alertmanager is not served from root of a DNS name. string  false
     ##
     externalUrl:
 
-    ## 	The route prefix Alertmanager registers HTTP handlers for. This is useful, if using ExternalURL and a proxy is rewriting HTTP routes of a request, and the actual ExternalURL is still true,
+    ## The route prefix Alertmanager registers HTTP handlers for. This is useful, if using ExternalURL and a proxy is rewriting HTTP routes of a request, and the actual ExternalURL is still true,
     ## but the server serves requests under a different route prefix. For example for use with kubectl proxy.
     ##
     routePrefix: /
@@ -558,7 +558,7 @@ alertmanager:
     #       app: alertmanager
 
     ## SecurityContext holds pod-level security attributes and common container settings.
-    ## This defaults to non root user with uid 1000 and gid 2000.	*v1.PodSecurityContext	false
+    ## This defaults to non root user with uid 1000 and gid 2000. *v1.PodSecurityContext  false
     ## ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
     ##
     securityContext:
@@ -678,7 +678,11 @@ grafana:
       ## Annotations for Grafana dashboard configmaps
       ##
       annotations: {}
-      multicluster: false
+      multicluster:
+        global:
+          enabled: false
+        etcd:
+          enabled: false
     datasources:
       enabled: true
       defaultDatasourceEnabled: true
@@ -740,14 +744,14 @@ grafana:
     # in grafana.ini
     path: "/metrics"
 
-    ## 	metric relabel configs to apply to samples before ingestion.
+    ## Metric relabel configs to apply to samples before ingestion.
     ##
     metricRelabelings: []
     # - action: keep
     #   regex: 'kube_(daemonset|deployment|pod|namespace|node|statefulset).+'
     #   sourceLabels: [__name__]
 
-    # 	relabel configs to apply to samples before ingestion.
+    #   relabel configs to apply to samples before ingestion.
     ##
     relabelings: []
     # - sourceLabels: [__meta_kubernetes_pod_node_name]
@@ -778,7 +782,7 @@ kubeApiServer:
         component: apiserver
         provider: kubernetes
 
-    ## 	metric relabel configs to apply to samples before ingestion.
+    ## Metric relabel configs to apply to samples before ingestion.
     ##
     metricRelabelings: []
     # - action: keep
@@ -856,7 +860,7 @@ kubelet:
     #   replacement: $1
     #   action: drop
 
-    # 	relabel configs to apply to samples before ingestion.
+    #   relabel configs to apply to samples before ingestion.
     #   metrics_path is required to match upstream rules and charts
     ##
     cAdvisorRelabelings:
@@ -901,7 +905,7 @@ kubelet:
     #   replacement: $1
     #   action: drop
 
-    # 	relabel configs to apply to samples before ingestion.
+    #   relabel configs to apply to samples before ingestion.
     #   metrics_path is required to match upstream rules and charts
     ##
     relabelings:
@@ -956,14 +960,14 @@ kubeControllerManager:
     # Name of the server to use when validating TLS certificate
     serverName: null
 
-    ## 	metric relabel configs to apply to samples before ingestion.
+    ## Metric relabel configs to apply to samples before ingestion.
     ##
     metricRelabelings: []
     # - action: keep
     #   regex: 'kube_(daemonset|deployment|pod|namespace|node|statefulset).+'
     #   sourceLabels: [__name__]
 
-    # 	relabel configs to apply to samples before ingestion.
+    #   relabel configs to apply to samples before ingestion.
     ##
     relabelings: []
     # - sourceLabels: [__meta_kubernetes_pod_node_name]
@@ -991,14 +995,14 @@ coreDns:
     ##
     proxyUrl: ""
 
-    ## 	metric relabel configs to apply to samples before ingestion.
+    ## Metric relabel configs to apply to samples before ingestion.
     ##
     metricRelabelings: []
     # - action: keep
     #   regex: 'kube_(daemonset|deployment|pod|namespace|node|statefulset).+'
     #   sourceLabels: [__name__]
 
-    # 	relabel configs to apply to samples before ingestion.
+    #   relabel configs to apply to samples before ingestion.
     ##
     relabelings: []
     # - sourceLabels: [__meta_kubernetes_pod_node_name]
@@ -1030,14 +1034,14 @@ kubeDns:
     ##
     proxyUrl: ""
 
-    ## 	metric relabel configs to apply to samples before ingestion.
+    ## Metric relabel configs to apply to samples before ingestion.
     ##
     metricRelabelings: []
     # - action: keep
     #   regex: 'kube_(daemonset|deployment|pod|namespace|node|statefulset).+'
     #   sourceLabels: [__name__]
 
-    # 	relabel configs to apply to samples before ingestion.
+    #   relabel configs to apply to samples before ingestion.
     ##
     relabelings: []
     # - sourceLabels: [__meta_kubernetes_pod_node_name]
@@ -1051,7 +1055,7 @@ kubeDns:
     #   regex: 'kube_(daemonset|deployment|pod|namespace|node|statefulset).+'
     #   sourceLabels: [__name__]
 
-    # 	relabel configs to apply to samples before ingestion.
+    #   relabel configs to apply to samples before ingestion.
     ##
     dnsmasqRelabelings: []
     # - sourceLabels: [__meta_kubernetes_pod_node_name]
@@ -1108,14 +1112,14 @@ kubeEtcd:
     certFile: ""
     keyFile: ""
 
-    ## 	metric relabel configs to apply to samples before ingestion.
+    ## Metric relabel configs to apply to samples before ingestion.
     ##
     metricRelabelings: []
     # - action: keep
     #   regex: 'kube_(daemonset|deployment|pod|namespace|node|statefulset).+'
     #   sourceLabels: [__name__]
 
-    # 	relabel configs to apply to samples before ingestion.
+    #   relabel configs to apply to samples before ingestion.
     ##
     relabelings: []
     # - sourceLabels: [__meta_kubernetes_pod_node_name]
@@ -1166,14 +1170,14 @@ kubeScheduler:
     ## Name of the server to use when validating TLS certificate
     serverName: null
 
-    ## 	metric relabel configs to apply to samples before ingestion.
+    ## Metric relabel configs to apply to samples before ingestion.
     ##
     metricRelabelings: []
     # - action: keep
     #   regex: 'kube_(daemonset|deployment|pod|namespace|node|statefulset).+'
     #   sourceLabels: [__name__]
 
-    # 	relabel configs to apply to samples before ingestion.
+    #   relabel configs to apply to samples before ingestion.
     ##
     relabelings: []
     # - sourceLabels: [__meta_kubernetes_pod_node_name]
@@ -1218,14 +1222,14 @@ kubeProxy:
     ##
     https: false
 
-    ## 	metric relabel configs to apply to samples before ingestion.
+    ## Metric relabel configs to apply to samples before ingestion.
     ##
     metricRelabelings: []
     # - action: keep
     #   regex: 'kube_(daemonset|deployment|pod|namespace|node|statefulset).+'
     #   sourceLabels: [__name__]
 
-    # 	relabel configs to apply to samples before ingestion.
+    #   relabel configs to apply to samples before ingestion.
     ##
     relabelings: []
     # - action: keep
@@ -1251,14 +1255,14 @@ kubeStateMetrics:
     ##
     namespaceOverride: ""
 
-    ## 	metric relabel configs to apply to samples before ingestion.
+    ## Metric relabel configs to apply to samples before ingestion.
     ##
     metricRelabelings: []
     # - action: keep
     #   regex: 'kube_(daemonset|deployment|pod|namespace|node|statefulset).+'
     #   sourceLabels: [__name__]
 
-    # 	relabel configs to apply to samples before ingestion.
+    #   relabel configs to apply to samples before ingestion.
     ##
     relabelings: []
     # - sourceLabels: [__meta_kubernetes_pod_node_name]
@@ -1303,7 +1307,7 @@ nodeExporter:
     ##
     scrapeTimeout: ""
 
-    ## 	metric relabel configs to apply to samples before ingestion.
+    ## Metric relabel configs to apply to samples before ingestion.
     ##
     metricRelabelings: []
     # - sourceLabels: [__name__]
@@ -1312,7 +1316,7 @@ nodeExporter:
     #   replacement: $1
     #   action: drop
 
-    ## 	relabel configs to apply to samples before ingestion.
+    ## relabel configs to apply to samples before ingestion.
     ##
     relabelings: []
     # - sourceLabels: [__meta_kubernetes_pod_node_name]
@@ -1377,7 +1381,7 @@ prometheusOperator:
       tolerations: []
 
       ## SecurityContext holds pod-level security attributes and common container settings.
-      ## This defaults to non root user with uid 2000 and gid 2000.	*v1.PodSecurityContext	false
+      ## This defaults to non root user with uid 2000 and gid 2000. *v1.PodSecurityContext  false
       ## ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
       ##
       securityContext:
@@ -1493,14 +1497,14 @@ prometheusOperator:
     scrapeTimeout: ""
     selfMonitor: true
 
-    ## 	metric relabel configs to apply to samples before ingestion.
+    ## Metric relabel configs to apply to samples before ingestion.
     ##
     metricRelabelings: []
     # - action: keep
     #   regex: 'kube_(daemonset|deployment|pod|namespace|node|statefulset).+'
     #   sourceLabels: [__name__]
 
-    # 	relabel configs to apply to samples before ingestion.
+    #   relabel configs to apply to samples before ingestion.
     ##
     relabelings: []
     # - sourceLabels: [__meta_kubernetes_pod_node_name]
@@ -1890,14 +1894,14 @@ prometheus:
 
     bearerTokenFile:
 
-    ## 	metric relabel configs to apply to samples before ingestion.
+    ## Metric relabel configs to apply to samples before ingestion.
     ##
     metricRelabelings: []
     # - action: keep
     #   regex: 'kube_(daemonset|deployment|pod|namespace|node|statefulset).+'
     #   sourceLabels: [__name__]
 
-    # 	relabel configs to apply to samples before ingestion.
+    #   relabel configs to apply to samples before ingestion.
     ##
     relabelings: []
     # - sourceLabels: [__meta_kubernetes_pod_node_name]
@@ -2358,7 +2362,7 @@ prometheus:
       runAsUser: 1000
       fsGroup: 2000
 
-    ## 	Priority class assigned to the Pods
+    ## Priority class assigned to the Pods
     ##
     priorityClassName: ""
 
@@ -2370,7 +2374,7 @@ prometheus:
     thanos: {}
 
     ## Containers allows injecting additional containers. This is meant to allow adding an authentication proxy to a Prometheus pod.
-    ##  if using proxy extraContainer  update targetPort with proxy container port
+    ## if using proxy extraContainer update targetPort with proxy container port
     containers: []
 
     ## InitContainers allows injecting additional initContainers. This is meant to allow doing some changes

```

## 17.1.3

**Release date:** 2021-08-07

![AppVersion: 0.49.0](https://img.shields.io/static/v1?label=AppVersion&message=0.49.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] The default for pathType mustn't be empty if "networking.k8s.io/v1" is available (#1150)

### Default value changes

```diff
# No changes in this release
```

## 17.1.2

**Release date:** 2021-08-07

![AppVersion: 0.49.0](https://img.shields.io/static/v1?label=AppVersion&message=0.49.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* add support for web TLS configuration (#1233)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index e7b5a65e..8ed4b4fe 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1943,6 +1943,10 @@ prometheus:
     ##
     enableAdminAPI: false
 
+    ## WebTLSConfig defines the TLS parameters for HTTPS
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#webtlsconfig
+    web: {}
+
     # EnableFeatures API enables access to Prometheus disabled features.
     # ref: https://prometheus.io/docs/prometheus/latest/disabled_features/
     enableFeatures: []

```

## 17.1.1

**Release date:** 2021-08-04

![AppVersion: 0.49.0](https://img.shields.io/static/v1?label=AppVersion&message=0.49.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Fix ingress logic (#1218)

### Default value changes

```diff
# No changes in this release
```

## 17.1.0

**Release date:** 2021-08-02

![AppVersion: 0.49.0](https://img.shields.io/static/v1?label=AppVersion&message=0.49.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] allow to create secrets for basic auth usage (#1193)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 0b2a607b..e7b5a65e 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -606,6 +606,16 @@ alertmanager:
     ## Use case is e.g. spanning an Alertmanager cluster across Kubernetes clusters with a single replica in each.
     forceEnableClusterMode: false
 
+  ## ExtraSecret can be used to store various data in an extra secret
+  ## (use it for example to store hashed basic auth credentials)
+  extraSecret:
+    ## if not set, name will be auto generated
+    # name: ""
+    annotations: {}
+    data: {}
+  #   auth: |
+  #     foo:$apr1$OFG3Xybp$ckL0FHDAkoXYIlH9.cysT0
+  #     someoneelse:$apr1$DMZX2Z4q$6SbQIfyuLQd.xmo/P0m2c.
 
 ## Using default values from https://github.com/grafana/helm-charts/blob/main/charts/grafana/values.yaml
 ##
@@ -1767,6 +1777,18 @@ prometheus:
     # - secretName: thanos-gateway-tls
     #   hosts:
     #   - thanos-gateway.domain.com
+    #
+
+  ## ExtraSecret can be used to store various data in an extra secret
+  ## (use it for example to store hashed basic auth credentials)
+  extraSecret:
+    ## if not set, name will be auto generated
+    # name: ""
+    annotations: {}
+    data: {}
+  #   auth: |
+  #     foo:$apr1$OFG3Xybp$ckL0FHDAkoXYIlH9.cysT0
+  #     someoneelse:$apr1$DMZX2Z4q$6SbQIfyuLQd.xmo/P0m2c.
 
   ingress:
     enabled: false

```

## 17.0.3

**Release date:** 2021-07-29

![AppVersion: 0.49.0](https://img.shields.io/static/v1?label=AppVersion&message=0.49.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Restore the semver check to AggregatedAPIDown and add it to the sync script (#1215)

### Default value changes

```diff
# No changes in this release
```

## 17.0.2

**Release date:** 2021-07-23

![AppVersion: 0.49.0](https://img.shields.io/static/v1?label=AppVersion&message=0.49.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add loadBalancerIP to external Thanos Service (#1190)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index f045ba41..0b2a607b 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1649,6 +1649,8 @@ prometheus:
     portName: grpc
     port: 10901
     targetPort: "grpc"
+    loadBalancerIP: ""
+    loadBalancerSourceRanges: []
 
     ## Service type
     ##

```

## 17.0.1

**Release date:** 2021-07-23

![AppVersion: 0.49.0](https://img.shields.io/static/v1?label=AppVersion&message=0.49.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Update README.md - Fix version mismatch (#1201)

### Default value changes

```diff
# No changes in this release
```

## 17.0.0

**Release date:** 2021-07-22

![AppVersion: 0.49.0](https://img.shields.io/static/v1?label=AppVersion&message=0.49.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Version bump prometheus-operator to 0.49.0 (#1199)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 276af37f..f045ba41 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1562,7 +1562,7 @@ prometheusOperator:
   ##
   image:
     repository: quay.io/prometheus-operator/prometheus-operator
-    tag: v0.48.1
+    tag: v0.49.0
     sha: ""
     pullPolicy: IfNotPresent
 
@@ -1578,7 +1578,7 @@ prometheusOperator:
   ##
   prometheusConfigReloaderImage:
     repository: quay.io/prometheus-operator/prometheus-config-reloader
-    tag: v0.48.1
+    tag: v0.49.0
     sha: ""
 
   ## Set the prometheus config reloader side-car CPU limit

```

## 16.15.0

**Release date:** 2021-07-20

![AppVersion: 0.48.1](https://img.shields.io/static/v1?label=AppVersion&message=0.48.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Update dependencies, rules and dashboard (#1179)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 3dddd432..276af37f 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1928,7 +1928,7 @@ prometheus:
     ##
     image:
       repository: quay.io/prometheus/prometheus
-      tag: v2.27.1
+      tag: v2.28.1
       sha: ""
 
     ## Tolerations for use with node taints

```

## 16.14.1

**Release date:** 2021-07-15

![AppVersion: 0.48.1](https://img.shields.io/static/v1?label=AppVersion&message=0.48.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* dashboard templates - only remove lines with defaults (#1160)

### Default value changes

```diff
# No changes in this release
```

## 16.14.0

**Release date:** 2021-07-15

![AppVersion: 0.48.1](https://img.shields.io/static/v1?label=AppVersion&message=0.48.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Sync rules (#1165)

### Default value changes

```diff
# No changes in this release
```

## 16.13.0

**Release date:** 2021-07-13

![AppVersion: 0.48.1](https://img.shields.io/static/v1?label=AppVersion&message=0.48.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add ability to specify existing secret with additional alertmanager configs (#1108)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index c3a9ffd4..3dddd432 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -2299,6 +2299,14 @@ prometheus:
     #     services:
     #       - metrics-prometheus-alertmanager
 
+    ## If additional alertmanager configurations are already deployed in a single secret, or you want to manage
+    ## them separately from the helm deployment, you can use this section.
+    ## Expected values are the secret name and key
+    ## Cannot be used with additionalAlertManagerConfigs
+    additionalAlertManagerConfigsSecret: {}
+      # name:
+      # key:
+
     ## AdditionalAlertRelabelConfigs allows specifying Prometheus alert relabel configurations. Alert relabel configurations specified are appended
     ## to the configurations generated by the Prometheus Operator. Alert relabel configurations specified must have the form as specified in the
     ## official Prometheus documentation: https://prometheus.io/docs/prometheus/latest/configuration/configuration/#alert_relabel_configs.

```

## 16.12.2

**Release date:** 2021-07-08

![AppVersion: 0.48.1](https://img.shields.io/static/v1?label=AppVersion&message=0.48.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* sync rules and update chart (#1137)

### Default value changes

```diff
# No changes in this release
```

## 16.12.1

**Release date:** 2021-07-01

![AppVersion: 0.48.1](https://img.shields.io/static/v1?label=AppVersion&message=0.48.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Dont create webhookconfigurations when prometheus operator is disabled. (#1100)

### Default value changes

```diff
# No changes in this release
```

## 16.12.0

**Release date:** 2021-06-26

![AppVersion: 0.48.1](https://img.shields.io/static/v1?label=AppVersion&message=0.48.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Adding the possibility to override thanos side-car container registry URL and version. (#1110)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 6a0ac267..c3a9ffd4 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1589,6 +1589,13 @@ prometheusOperator:
   ##
   configReloaderMemory: 50Mi
 
+  ## Thanos side-car image when configured
+  ##
+  thanosImage:
+    repository: quay.io/thanos/thanos
+    tag: v0.17.2
+    sha: ""
+
   ## Set a Field Selector to filter watched secrets
   ##
   secretFieldSelector: ""

```

## 16.11.0

**Release date:** 2021-06-24

![AppVersion: 0.48.1](https://img.shields.io/static/v1?label=AppVersion&message=0.48.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Sync rules, dbs and subcharts (#1103)

### Default value changes

```diff
# No changes in this release
```

## 16.10.0

**Release date:** 2021-06-18

![AppVersion: 0.48.1](https://img.shields.io/static/v1?label=AppVersion&message=0.48.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] add endpoint port mertics for kube-state-metrics self monitor (#1090)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index c4b34a1a..6a0ac267 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1258,6 +1258,10 @@ kubeStateMetrics:
     #   replacement: $1
     #   action: replace
 
+    # Enable self metrics configuration for Service Monitor
+    selfMonitor:
+      enabled: false
+
 ## Configuration for kube-state-metrics subchart
 ##
 kube-state-metrics:

```

## 16.9.3

**Release date:** 2021-06-18

![AppVersion: 0.48.1](https://img.shields.io/static/v1?label=AppVersion&message=0.48.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Update Grafana helm chart (#1092)

### Default value changes

```diff
# No changes in this release
```

## 16.9.2

**Release date:** 2021-06-18

![AppVersion: 0.48.1](https://img.shields.io/static/v1?label=AppVersion&message=0.48.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Bump Grafana-chart to 6.13.1 (#1089)

### Default value changes

```diff
# No changes in this release
```

## 16.9.1

**Release date:** 2021-06-17

![AppVersion: 0.48.1](https://img.shields.io/static/v1?label=AppVersion&message=0.48.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* bump grafana dep in kube-prometheus-stack to 6.13.* (#1088)

### Default value changes

```diff
# No changes in this release
```

## 16.9.0

**Release date:** 2021-06-16

![AppVersion: 0.48.1](https://img.shields.io/static/v1?label=AppVersion&message=0.48.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] add extra support for Amazon EKS and Google GKE KubeVersion on ingress (#1083)

### Default value changes

```diff
# No changes in this release
```

## 16.8.0

**Release date:** 2021-06-15

![AppVersion: 0.48.1](https://img.shields.io/static/v1?label=AppVersion&message=0.48.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Sync rules and dashboards (#1075)

### Default value changes

```diff
# No changes in this release
```

## 16.7.1

**Release date:** 2021-06-14

![AppVersion: 0.48.1](https://img.shields.io/static/v1?label=AppVersion&message=0.48.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Revert c4a7d10fdc6a0f694d9b97e9446207ba67d997dd (#1072)

### Default value changes

```diff
# No changes in this release
```

## 16.7.0

**Release date:** 2021-06-11

![AppVersion: 0.48.1](https://img.shields.io/static/v1?label=AppVersion&message=0.48.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Remove old rulesets and sync new (#1066)

### Default value changes

```diff
# No changes in this release
```

## 16.6.4

**Release date:** 2021-06-11

![AppVersion: 0.48.1](https://img.shields.io/static/v1?label=AppVersion&message=0.48.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Update kube-state-metrics ServiceMonitor port (#1065)

### Default value changes

```diff
# No changes in this release
```

## 16.6.3

**Release date:** 2021-06-11

![AppVersion: 0.48.1](https://img.shields.io/static/v1?label=AppVersion&message=0.48.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Update grafana helm chart version (#1064)

### Default value changes

```diff
# No changes in this release
```

## 16.6.2

**Release date:** 2021-06-10

![AppVersion: 0.48.1](https://img.shields.io/static/v1?label=AppVersion&message=0.48.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Following prometheus-operator change #860 (#1048)

### Default value changes

```diff
# No changes in this release
```

## 16.6.1

**Release date:** 2021-06-10

![AppVersion: 0.48.1](https://img.shields.io/static/v1?label=AppVersion&message=0.48.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Update Prometheus-Operator to v0.48.1 (#1053)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 94f9576a..c4b34a1a 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1558,7 +1558,7 @@ prometheusOperator:
   ##
   image:
     repository: quay.io/prometheus-operator/prometheus-operator
-    tag: v0.48.0
+    tag: v0.48.1
     sha: ""
     pullPolicy: IfNotPresent
 
@@ -1574,7 +1574,7 @@ prometheusOperator:
   ##
   prometheusConfigReloaderImage:
     repository: quay.io/prometheus-operator/prometheus-config-reloader
-    tag: v0.48.0
+    tag: v0.48.1
     sha: ""
 
   ## Set the prometheus config reloader side-car CPU limit

```

## 16.6.0

**Release date:** 2021-06-09

![AppVersion: 0.48.0](https://img.shields.io/static/v1?label=AppVersion&message=0.48.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] update grafana to 8.0 via chart version 6.12.0 (#1056)

### Default value changes

```diff
# No changes in this release
```

## 16.5.0

**Release date:** 2021-06-08

![AppVersion: 0.48.0](https://img.shields.io/static/v1?label=AppVersion&message=0.48.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Sync grafana dashboards (#1050)

### Default value changes

```diff
# No changes in this release
```

## 16.4.1

**Release date:** 2021-06-07

![AppVersion: 0.48.0](https://img.shields.io/static/v1?label=AppVersion&message=0.48.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] move relabelings filed to servicemonitor in kubeApiServer (#1014)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index b125a831..94f9576a 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -754,19 +754,6 @@ kubeApiServer:
   tlsConfig:
     serverName: kubernetes
     insecureSkipVerify: false
-
-  ## If your API endpoint address is not reachable (as in AKS) you can replace it with the kubernetes service
-  ##
-  relabelings: []
-  # - sourceLabels:
-  #     - __meta_kubernetes_namespace
-  #     - __meta_kubernetes_service_name
-  #     - __meta_kubernetes_endpoint_port_name
-  #   action: keep
-  #   regex: default;kubernetes;https
-  # - targetLabel: __address__
-  #   replacement: kubernetes.default.svc:443
-
   serviceMonitor:
     ## Scrape interval. If not set, the Prometheus default scrape interval is used.
     ##
@@ -787,6 +774,15 @@ kubeApiServer:
     # - action: keep
     #   regex: 'kube_(daemonset|deployment|pod|namespace|node|statefulset).+'
     #   sourceLabels: [__name__]
+    relabelings: []
+    # - sourceLabels:
+    #     - __meta_kubernetes_namespace
+    #     - __meta_kubernetes_service_name
+    #     - __meta_kubernetes_endpoint_port_name
+    #   action: keep
+    #   regex: default;kubernetes;https
+    # - targetLabel: __address__
+    #   replacement: kubernetes.default.svc:443
 
 ## Component scraping the kubelet and kubelet-hosted cAdvisor
 ##

```

## 16.4.0

**Release date:** 2021-06-07

![AppVersion: 0.48.0](https://img.shields.io/static/v1?label=AppVersion&message=0.48.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Version bump for grafana, node-exporter and kube-state-metrics (#1049)

### Default value changes

```diff
# No changes in this release
```

## 16.3.2

**Release date:** 2021-06-07

![AppVersion: 0.48.0](https://img.shields.io/static/v1?label=AppVersion&message=0.48.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] fix: CA injection for admission webhook (#897)

### Default value changes

```diff
# No changes in this release
```

## 16.3.1

**Release date:** 2021-06-06

![AppVersion: 0.48.0](https://img.shields.io/static/v1?label=AppVersion&message=0.48.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] AlertManager v0.22.2 (#1046)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index caa43d7f..b125a831 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -396,7 +396,7 @@ alertmanager:
     ##
     image:
       repository: quay.io/prometheus/alertmanager
-      tag: v0.22.0
+      tag: v0.22.2
       sha: ""
 
     ## If true then the user will be responsible to provide a secret with alertmanager configuration

```

## 16.3.0

**Release date:** 2021-06-06

![AppVersion: 0.48.0](https://img.shields.io/static/v1?label=AppVersion&message=0.48.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Fix issue #1038 (#1045)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 7a327961..caa43d7f 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1365,6 +1365,16 @@ prometheusOperator:
       nodeSelector: {}
       affinity: {}
       tolerations: []
+
+      ## SecurityContext holds pod-level security attributes and common container settings.
+      ## This defaults to non root user with uid 2000 and gid 2000.	*v1.PodSecurityContext	false
+      ## ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
+      ##
+      securityContext:
+        runAsGroup: 2000
+        runAsNonRoot: true
+        runAsUser: 2000
+
     # Use certmanager to generate webhook certs
     certManager:
       enabled: false

```

## 16.2.0

**Release date:** 2021-06-05

![AppVersion: 0.48.0](https://img.shields.io/static/v1?label=AppVersion&message=0.48.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Limit grafana deployment to dashboards (#1044)
* fix nodeExporter serviceMonitor's namespaceSelector (#985)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 109820a3..7a327961 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -613,6 +613,14 @@ grafana:
   enabled: true
   namespaceOverride: ""
 
+  ## ForceDeployDatasources Create datasource configmap even if grafana deployment has been disabled
+  ##
+  forceDeployDatasources: false
+
+  ## ForceDeployDashboard Create dashboard configmap even if grafana deployment has been disabled
+  ##
+  forceDeployDashboards: false
+
   ## Deploy default dashboards.
   ##
   defaultDashboardsEnabled: true
@@ -665,6 +673,10 @@ grafana:
       enabled: true
       defaultDatasourceEnabled: true
 
+      ## URL of prometheus datasource
+      ##
+      # url: http://prometheus-stack-prometheus:9090/
+
       # If not defined, will use prometheus.prometheusSpec.scrapeInterval or its default
       # defaultDatasourceScrapeInterval: 15s
 

```

## 16.1.2

**Release date:** 2021-05-30

![AppVersion: 0.48.0](https://img.shields.io/static/v1?label=AppVersion&message=0.48.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Update prometheus service selector for 0.48.0 (#1010)

### Default value changes

```diff
# No changes in this release
```

## 16.1.1

**Release date:** 2021-05-28

![AppVersion: 0.48.0](https://img.shields.io/static/v1?label=AppVersion&message=0.48.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Set pathType on Ingress only when supported (#960)

### Default value changes

```diff
# No changes in this release
```

## 16.1.0

**Release date:** 2021-05-27

![AppVersion: 0.48.0](https://img.shields.io/static/v1?label=AppVersion&message=0.48.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Prometheus-Operator v0.48.0 and update used images (#1006)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 9098f9fa..109820a3 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -396,7 +396,7 @@ alertmanager:
     ##
     image:
       repository: quay.io/prometheus/alertmanager
-      tag: v0.21.0
+      tag: v0.22.0
       sha: ""
 
     ## If true then the user will be responsible to provide a secret with alertmanager configuration
@@ -1342,7 +1342,7 @@ prometheusOperator:
       enabled: true
       image:
         repository: jettech/kube-webhook-certgen
-        tag: v1.5.0
+        tag: v1.5.2
         sha: ""
         pullPolicy: IfNotPresent
       resources: {}
@@ -1540,7 +1540,7 @@ prometheusOperator:
   ##
   image:
     repository: quay.io/prometheus-operator/prometheus-operator
-    tag: v0.47.1
+    tag: v0.48.0
     sha: ""
     pullPolicy: IfNotPresent
 
@@ -1556,7 +1556,7 @@ prometheusOperator:
   ##
   prometheusConfigReloaderImage:
     repository: quay.io/prometheus-operator/prometheus-config-reloader
-    tag: v0.47.1
+    tag: v0.48.0
     sha: ""
 
   ## Set the prometheus config reloader side-car CPU limit
@@ -1899,7 +1899,7 @@ prometheus:
     ##
     image:
       repository: quay.io/prometheus/prometheus
-      tag: v2.26.0
+      tag: v2.27.1
       sha: ""
 
     ## Tolerations for use with node taints

```

## 16.0.1

**Release date:** 2021-05-21

![AppVersion: 0.47.1](https://img.shields.io/static/v1?label=AppVersion&message=0.47.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* fix kube-state-metrics url (#993)

### Default value changes

```diff
# No changes in this release
```

## 16.0.0

**Release date:** 2021-05-20

![AppVersion: 0.47.1](https://img.shields.io/static/v1?label=AppVersion&message=0.47.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Update dependencies / prometheus-operator 0.47.1 (#951)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 51788366..9098f9fa 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1540,7 +1540,7 @@ prometheusOperator:
   ##
   image:
     repository: quay.io/prometheus-operator/prometheus-operator
-    tag: v0.47.0
+    tag: v0.47.1
     sha: ""
     pullPolicy: IfNotPresent
 
@@ -1556,7 +1556,7 @@ prometheusOperator:
   ##
   prometheusConfigReloaderImage:
     repository: quay.io/prometheus-operator/prometheus-config-reloader
-    tag: v0.47.0
+    tag: v0.47.1
     sha: ""
 
   ## Set the prometheus config reloader side-car CPU limit

```

## 15.4.6

**Release date:** 2021-05-12

![AppVersion: 0.47.0](https://img.shields.io/static/v1?label=AppVersion&message=0.47.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Fix proxy url error (#943)

### Default value changes

```diff
# No changes in this release
```

## 15.4.5

**Release date:** 2021-05-11

![AppVersion: 0.47.0](https://img.shields.io/static/v1?label=AppVersion&message=0.47.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Fix ingressperreplica not working (#946)

### Default value changes

```diff
# No changes in this release
```

## 15.4.4

**Release date:** 2021-05-05

![AppVersion: 0.47.0](https://img.shields.io/static/v1?label=AppVersion&message=0.47.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Omit imagePullSecrets element if no secrets provided (#936)

### Default value changes

```diff
# No changes in this release
```

## 15.4.3

**Release date:** 2021-05-04

![AppVersion: 0.47.0](https://img.shields.io/static/v1?label=AppVersion&message=0.47.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] fix lint in additional-prometheus-rules (#931)

### Default value changes

```diff
# No changes in this release
```

## 15.4.2

**Release date:** 2021-05-04

![AppVersion: 0.47.0](https://img.shields.io/static/v1?label=AppVersion&message=0.47.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack]: Allow kubeVersion to be overridden due to helm V3 rendering issue (#920)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 1321efa9..51788366 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -14,6 +14,10 @@ namespaceOverride: ""
 ##
 kubeTargetVersionOverride: ""
 
+## Allow kubeVersion to be overridden while creating the ingress
+##
+kubeVersionOverride: ""
+
 ## Provide a name to substitute for the full names of resources
 ##
 fullnameOverride: ""

```

## 15.4.1

**Release date:** 2021-05-04

![AppVersion: 0.47.0](https://img.shields.io/static/v1?label=AppVersion&message=0.47.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] docs: upgrade CRD to v0.44 for chart@v12 (#926)

### Default value changes

```diff
# No changes in this release
```

## 15.4.0

**Release date:** 2021-05-04

![AppVersion: 0.47.0](https://img.shields.io/static/v1?label=AppVersion&message=0.47.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add external service for thanos-sidecar (#916)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 6b210fa6..1321efa9 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1607,6 +1607,24 @@ prometheus:
     ##
     nodePort: 30901
 
+  # Service for external access to sidecar
+  # Enabling this creates a service to expose thanos-sidecar outside the cluster.
+  thanosServiceExternal:
+    enabled: false
+    annotations: {}
+    labels: {}
+    portName: grpc
+    port: 10901
+    targetPort: "grpc"
+
+    ## Service type
+    ##
+    type: LoadBalancer
+
+    ## Port to expose on each node
+    ##
+    nodePort: 30901
+
   ## Configuration for Prometheus service
   ##
   service:

```

## 15.3.1

**Release date:** 2021-04-29

![AppVersion: 0.47.0](https://img.shields.io/static/v1?label=AppVersion&message=0.47.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* chore: bump chart version (#911)

### Default value changes

```diff
# No changes in this release
```

## 15.3.0

**Release date:** 2021-04-29

![AppVersion: 0.47.0](https://img.shields.io/static/v1?label=AppVersion&message=0.47.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Changed Url for etcd rules and synced prometh… (#908)

### Default value changes

```diff
# No changes in this release
```

## 15.2.4

**Release date:** 2021-04-29

![AppVersion: 0.47.0](https://img.shields.io/static/v1?label=AppVersion&message=0.47.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* add templating helm function to kube-prometheus remote config. On using kube-prometheus as a dependent helm chart passing {{ .Release.Name}} & {{ .Release.Namespace }} the respective values aren't parsed as tpl function isn't used. (#875)

### Default value changes

```diff
# No changes in this release
```

## 15.2.3

**Release date:** 2021-04-29

![AppVersion: 0.47.0](https://img.shields.io/static/v1?label=AppVersion&message=0.47.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] New location for kube-state-metrics (#892)

### Default value changes

```diff
# No changes in this release
```

## 15.2.2

**Release date:** 2021-04-28

![AppVersion: 0.47.0](https://img.shields.io/static/v1?label=AppVersion&message=0.47.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack]: fix Capabilities.APIVersions.Has check for ingress API version (#885)

### Default value changes

```diff
# No changes in this release
```

## 15.2.1

**Release date:** 2021-04-26

![AppVersion: 0.47.0](https://img.shields.io/static/v1?label=AppVersion&message=0.47.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add matchLabel to prevent duplicate kubelet targets (#860)

### Default value changes

```diff
# No changes in this release
```

## 15.2.0

**Release date:** 2021-04-22

![AppVersion: 0.47.0](https://img.shields.io/static/v1?label=AppVersion&message=0.47.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* #876 Update kube-prometheus-stack grafana dependency to 6.8.0 (#877)

### Default value changes

```diff
# No changes in this release
```

## 15.1.3

**Release date:** 2021-04-22

![AppVersion: 0.47.0](https://img.shields.io/static/v1?label=AppVersion&message=0.47.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* fixes oversight in #850 (#872)
* [kube-prometheus-stack] Recommanded Kubernetes labels (#692)

### Default value changes

```diff
# No changes in this release
```

## 15.1.2

**Release date:** 2021-04-21

![AppVersion: 0.47.0](https://img.shields.io/static/v1?label=AppVersion&message=0.47.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* update ingress for v1 (#850)

### Default value changes

```diff
# No changes in this release
```

## 15.1.1

**Release date:** 2021-04-21

![AppVersion: 0.47.0](https://img.shields.io/static/v1?label=AppVersion&message=0.47.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] fix(#857): Adjusted the anti-affinity rule to match the statefulset p… (#867)

### Default value changes

```diff
# No changes in this release
```

## 15.1.0

**Release date:** 2021-04-20

![AppVersion: 0.47.0](https://img.shields.io/static/v1?label=AppVersion&message=0.47.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] support enableFeatures API (#828)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 0e2d6d6e..6b210fa6 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1868,6 +1868,11 @@ prometheus:
     ##
     enableAdminAPI: false
 
+    # EnableFeatures API enables access to Prometheus disabled features.
+    # ref: https://prometheus.io/docs/prometheus/latest/disabled_features/
+    enableFeatures: []
+    # - exemplar-storage
+
     ## Image of Prometheus.
     ##
     image:

```

## 15.0.0

**Release date:** 2021-04-19

![AppVersion: 0.47.0](https://img.shields.io/static/v1?label=AppVersion&message=0.47.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Update prometheus-operator to 0.47.0 (#858)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index c7abeefb..0e2d6d6e 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1536,7 +1536,7 @@ prometheusOperator:
   ##
   image:
     repository: quay.io/prometheus-operator/prometheus-operator
-    tag: v0.46.0
+    tag: v0.47.0
     sha: ""
     pullPolicy: IfNotPresent
 
@@ -1552,7 +1552,7 @@ prometheusOperator:
   ##
   prometheusConfigReloaderImage:
     repository: quay.io/prometheus-operator/prometheus-config-reloader
-    tag: v0.46.0
+    tag: v0.47.0
     sha: ""
 
   ## Set the prometheus config reloader side-car CPU limit

```

## 14.9.0

**Release date:** 2021-04-14

![AppVersion: 0.46.0](https://img.shields.io/static/v1?label=AppVersion&message=0.46.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add proxyUrl variable to exporters (#789)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index bc546254..c7abeefb 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -349,6 +349,10 @@ alertmanager:
     interval: ""
     selfMonitor: true
 
+    ## proxyUrl: URL of a proxy that should be used for scraping.
+    ##
+    proxyUrl: ""
+
     ## scheme: HTTP scheme to use for scraping. Can be used with `tlsConfig` for example if using istio mTLS.
     scheme: ""
 
@@ -751,6 +755,10 @@ kubeApiServer:
     ## Scrape interval. If not set, the Prometheus default scrape interval is used.
     ##
     interval: ""
+    ## proxyUrl: URL of a proxy that should be used for scraping.
+    ##
+    proxyUrl: ""
+
     jobLabel: component
     selector:
       matchLabels:
@@ -775,6 +783,10 @@ kubelet:
     ##
     interval: ""
 
+    ## proxyUrl: URL of a proxy that should be used for scraping.
+    ##
+    proxyUrl: ""
+
     ## Enable scraping the kubelet over https. For requirements to enable this see
     ## https://github.com/prometheus-operator/prometheus-operator/issues/926
     ##
@@ -907,6 +919,10 @@ kubeControllerManager:
     ##
     interval: ""
 
+    ## proxyUrl: URL of a proxy that should be used for scraping.
+    ##
+    proxyUrl: ""
+
     ## Enable scraping kube-controller-manager over https.
     ## Requires proper certs (not self-signed) and delegated authentication/authorization checks
     ##
@@ -949,6 +965,10 @@ coreDns:
     ##
     interval: ""
 
+    ## proxyUrl: URL of a proxy that should be used for scraping.
+    ##
+    proxyUrl: ""
+
     ## 	metric relabel configs to apply to samples before ingestion.
     ##
     metricRelabelings: []
@@ -984,6 +1004,10 @@ kubeDns:
     ##
     interval: ""
 
+    ## proxyUrl: URL of a proxy that should be used for scraping.
+    ##
+    proxyUrl: ""
+
     ## 	metric relabel configs to apply to samples before ingestion.
     ##
     metricRelabelings: []
@@ -1052,6 +1076,9 @@ kubeEtcd:
     ## Scrape interval. If not set, the Prometheus default scrape interval is used.
     ##
     interval: ""
+    ## proxyUrl: URL of a proxy that should be used for scraping.
+    ##
+    proxyUrl: ""
     scheme: http
     insecureSkipVerify: false
     serverName: ""
@@ -1103,6 +1130,9 @@ kubeScheduler:
     ## Scrape interval. If not set, the Prometheus default scrape interval is used.
     ##
     interval: ""
+    ## proxyUrl: URL of a proxy that should be used for scraping.
+    ##
+    proxyUrl: ""
     ## Enable scraping kube-scheduler over https.
     ## Requires proper certs (not self-signed) and delegated authentication/authorization checks
     ##
@@ -1157,6 +1187,10 @@ kubeProxy:
     ##
     interval: ""
 
+    ## proxyUrl: URL of a proxy that should be used for scraping.
+    ##
+    proxyUrl: ""
+
     ## Enable scraping kube-proxy over https.
     ## Requires proper certs (not self-signed) and delegated authentication/authorization checks
     ##
@@ -1185,6 +1219,9 @@ kubeStateMetrics:
     ## Scrape interval. If not set, the Prometheus default scrape interval is used.
     ##
     interval: ""
+    ## proxyUrl: URL of a proxy that should be used for scraping.
+    ##
+    proxyUrl: ""
     ## Override serviceMonitor selector
     ##
     selectorOverride: {}
@@ -1232,6 +1269,10 @@ nodeExporter:
     ##
     interval: ""
 
+    ## proxyUrl: URL of a proxy that should be used for scraping.
+    ##
+    proxyUrl: ""
+
     ## How long until a scrape request times out. If not set, the Prometheus default scape timeout is used.
     ##
     scrapeTimeout: ""

```

## 14.8.2

**Release date:** 2021-04-14

![AppVersion: 0.46.0](https://img.shields.io/static/v1?label=AppVersion&message=0.46.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] add optional kube-state-metrics namespace selector (#827)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 4d2840f6..bc546254 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1188,6 +1188,9 @@ kubeStateMetrics:
     ## Override serviceMonitor selector
     ##
     selectorOverride: {}
+    ## Override namespace selector
+    ##
+    namespaceOverride: ""
 
     ## 	metric relabel configs to apply to samples before ingestion.
     ##

```

## 14.8.1

**Release date:** 2021-04-14

![AppVersion: 0.46.0](https://img.shields.io/static/v1?label=AppVersion&message=0.46.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add support for Alertmanager CR annotations (#752)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index b85ee806..4d2840f6 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -117,6 +117,10 @@ alertmanager:
   ##
   enabled: true
 
+  ## Annotations for Alertmanager
+  ##
+  annotations: {}
+
   ## Api that prometheus will use to communicate with alertmanager. Possible values are v1, v2
   ##
   apiVersion: v2

```

## 14.8.0

**Release date:** 2021-04-14

![AppVersion: 0.46.0](https://img.shields.io/static/v1?label=AppVersion&message=0.46.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Update node-exporter dep to 1.17.0 + kube-state-metrics to 2.13.2 + grafana to 6.7.4 (#847)

### Default value changes

```diff
# No changes in this release
```

## 14.7.2

**Release date:** 2021-04-14

![AppVersion: 0.46.0](https://img.shields.io/static/v1?label=AppVersion&message=0.46.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add service account annotations (#759)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index b2c5c4ef..b85ee806 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1535,6 +1535,7 @@ prometheus:
   serviceAccount:
     create: true
     name: ""
+    annotations: {}
 
   # Service for thanos service discovery on sidecar
   # Enable this can make Thanos Query can use

```

## 14.7.1

**Release date:** 2021-04-14

![AppVersion: 0.46.0](https://img.shields.io/static/v1?label=AppVersion&message=0.46.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Avoid using non-ascii char in values.yaml (#843)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 575a1f1b..b2c5c4ef 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -194,7 +194,7 @@ alertmanager:
   #         *Graph:* <{{ .GeneratorURL }}|:chart_with_upwards_trend:>
   #         *Runbook:* <{{ .Annotations.runbook }}|:spiral_note_pad:>
   #         *Details:*
-  #           {{ range .Labels.SortedPairs }} • *{{ .Name }}:* `{{ .Value }}`
+  #           {{ range .Labels.SortedPairs }} - *{{ .Name }}:* `{{ .Value }}`
   #           {{ end }}
   #       {{ end }}
   #       {{ end }}
@@ -375,7 +375,7 @@ alertmanager:
   ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#alertmanagerspec
   ##
   alertmanagerSpec:
-    ## Standard object’s metadata. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata
+    ## Standard object's metadata. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata
     ## Metadata Labels and Annotations gets propagated to the Alertmanager pods.
     ##
     podMetadata: {}
@@ -2044,7 +2044,7 @@ prometheus:
     ##
     routePrefix: /
 
-    ## Standard object’s metadata. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata
+    ## Standard object's metadata. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata
     ## Metadata Labels and Annotations gets propagated to the prometheus pods.
     ##
     podMetadata: {}

```

## 14.7.0

**Release date:** 2021-04-14

![AppVersion: 0.46.0](https://img.shields.io/static/v1?label=AppVersion&message=0.46.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Bump prometheus image version to fix the target's issue (#842)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 691b2cf4..575a1f1b 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1823,7 +1823,7 @@ prometheus:
     ##
     image:
       repository: quay.io/prometheus/prometheus
-      tag: v2.24.0
+      tag: v2.26.0
       sha: ""
 
     ## Tolerations for use with node taints

```

## 14.6.3

**Release date:** 2021-04-14

![AppVersion: 0.46.0](https://img.shields.io/static/v1?label=AppVersion&message=0.46.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Bump PyYAML to 5.4 (#829)

### Default value changes

```diff
# No changes in this release
```

## 14.6.2

**Release date:** 2021-04-09

![AppVersion: 0.46.0](https://img.shields.io/static/v1?label=AppVersion&message=0.46.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] add additionalRemoteRead and additionalRemoteWrite (#747)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 290ea1d1..691b2cf4 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -2081,11 +2081,15 @@ prometheus:
     ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#remotereadspec
     remoteRead: []
     # - url: http://remote1/read
+    ## additionalRemoteRead is appended to remoteRead
+    additionalRemoteRead: []
 
     ## The remote_write spec configuration for Prometheus.
     ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#remotewritespec
     remoteWrite: []
     # - url: http://remote1/push
+    ## additionalRemoteWrite is appended to remoteWrite
+    additionalRemoteWrite: []
 
     ## Enable/Disable Grafana dashboards provisioning for prometheus remote write feature
     remoteWriteDashboards: false

```

## 14.6.1

**Release date:** 2021-04-09

![AppVersion: 0.46.0](https://img.shields.io/static/v1?label=AppVersion&message=0.46.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* refactor(kube-prometheus-stack): set cert-manager generated certificate duration to include minutes and seconds complying to cert-manager representation on updates (#750)

### Default value changes

```diff
# No changes in this release
```

## 14.6.0

**Release date:** 2021-04-09

![AppVersion: 0.46.0](https://img.shields.io/static/v1?label=AppVersion&message=0.46.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Update grafana dep to 6.7 (#814)

### Default value changes

```diff
# No changes in this release
```

## 14.5.0

**Release date:** 2021-04-02

![AppVersion: 0.46.0](https://img.shields.io/static/v1?label=AppVersion&message=0.46.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Invalid value "None" for clusterIP (#784)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 42f6ac7c..290ea1d1 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1548,6 +1548,7 @@ prometheus:
     portName: grpc
     port: 10901
     targetPort: "grpc"
+    clusterIP: "None"
 
     ## Service type
     ##

```

## 14.4.0

**Release date:** 2021-03-27

![AppVersion: 0.46.0](https://img.shields.io/static/v1?label=AppVersion&message=0.46.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] add more controls to kubescheduler,kubeproxy,controllermanager,etcd exporters (#797)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 5848603d..42f6ac7c 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -891,12 +891,14 @@ kubeControllerManager:
   ## If using kubeControllerManager.endpoints only the port and targetPort are used
   ##
   service:
+    enabled: true
     port: 10252
     targetPort: 10252
     # selector:
     #   component: kube-controller-manager
 
   serviceMonitor:
+    enabled: true
     ## Scrape interval. If not set, the Prometheus default scrape interval is used.
     ##
     interval: ""
@@ -1024,6 +1026,7 @@ kubeEtcd:
   ## Etcd service. If using kubeEtcd.endpoints only the port and targetPort are used
   ##
   service:
+    enabled: true
     port: 2379
     targetPort: 2379
     # selector:
@@ -1041,6 +1044,7 @@ kubeEtcd:
   ##   keyFile: /etc/prometheus/secrets/etcd-client-cert/etcd-client-key
   ##
   serviceMonitor:
+    enabled: true
     ## Scrape interval. If not set, the Prometheus default scrape interval is used.
     ##
     interval: ""
@@ -1084,12 +1088,14 @@ kubeScheduler:
   ## If using kubeScheduler.endpoints only the port and targetPort are used
   ##
   service:
+    enabled: true
     port: 10251
     targetPort: 10251
     # selector:
     #   component: kube-scheduler
 
   serviceMonitor:
+    enabled: true
     ## Scrape interval. If not set, the Prometheus default scrape interval is used.
     ##
     interval: ""
@@ -1135,12 +1141,14 @@ kubeProxy:
   # - 10.141.4.24
 
   service:
+    enabled: true
     port: 10249
     targetPort: 10249
     # selector:
     #   k8s-app: kube-proxy
 
   serviceMonitor:
+    enabled: true
     ## Scrape interval. If not set, the Prometheus default scrape interval is used.
     ##
     interval: ""

```

## 14.3.0

**Release date:** 2021-03-20

![AppVersion: 0.46.0](https://img.shields.io/static/v1?label=AppVersion&message=0.46.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Exclude default rules from namespace enforcement (#770)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index efd99acc..5848603d 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -2242,9 +2242,14 @@ prometheus:
     ## configs, and they will only discover endpoints within their current namespace. Defaults to false.
     ignoreNamespaceSelectors: false
 
+    ## EnforcedNamespaceLabel enforces adding a namespace label of origin for each alert and metric that is user created.
+    ## The label value will always be the namespace of the object that is being created.
+    ## Disabled by default
+    enforcedNamespaceLabel: ""
+
     ## PrometheusRulesExcludedFromEnforce - list of prometheus rules to be excluded from enforcing of adding namespace labels.
     ## Works only if enforcedNamespaceLabel set to true. Make sure both ruleNamespace and ruleName are set for each pair
-    prometheusRulesExcludedFromEnforce: false
+    prometheusRulesExcludedFromEnforce: []
 
     ## QueryLogFile specifies the file to which PromQL queries are logged. Note that this location must be writable,
     ## and can be persisted using an attached volume. Alternatively, the location can be set to a stdout location such

```

## 14.2.0

**Release date:** 2021-03-19

![AppVersion: 0.46.0](https://img.shields.io/static/v1?label=AppVersion&message=0.46.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Update dependency version (#776)

### Default value changes

```diff
# No changes in this release
```

## 14.1.2

**Release date:** 2021-03-18

![AppVersion: 0.46.0](https://img.shields.io/static/v1?label=AppVersion&message=0.46.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Fix LoadBalancer typo in service type (#761)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index e59a17bd..efd99acc 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -331,7 +331,7 @@ alertmanager:
     nodePort: 30904
 
     ## Loadbalancer source IP ranges
-    ## Only used if servicePerReplica.type is "loadbalancer"
+    ## Only used if servicePerReplica.type is "LoadBalancer"
     loadBalancerSourceRanges: []
     ## Service type
     ##
@@ -1351,13 +1351,13 @@ prometheusOperator:
     additionalPorts: []
 
   ## Loadbalancer IP
-  ## Only use if service.type is "loadbalancer"
+  ## Only use if service.type is "LoadBalancer"
   ##
     loadBalancerIP: ""
     loadBalancerSourceRanges: []
 
   ## Service type
-  ## NodePort, ClusterIP, loadbalancer
+  ## NodePort, ClusterIP, LoadBalancer
   ##
     type: ClusterIP
 
@@ -1574,7 +1574,7 @@ prometheus:
     nodePort: 30090
 
     ## Loadbalancer IP
-    ## Only use if service.type is "loadbalancer"
+    ## Only use if service.type is "LoadBalancer"
     loadBalancerIP: ""
     loadBalancerSourceRanges: []
     ## Service type
@@ -1602,7 +1602,7 @@ prometheus:
     nodePort: 30091
 
     ## Loadbalancer source IP ranges
-    ## Only used if servicePerReplica.type is "loadbalancer"
+    ## Only used if servicePerReplica.type is "LoadBalancer"
     loadBalancerSourceRanges: []
     ## Service type
     ##

```

## 14.1.1

**Release date:** 2021-03-18

![AppVersion: 0.46.0](https://img.shields.io/static/v1?label=AppVersion&message=0.46.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Fix typo and formatting in comments (#772)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 961c6198..e59a17bd 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1911,10 +1911,10 @@ prometheus:
     ruleSelectorNilUsesHelmValues: true
 
     ## PrometheusRules to be selected for target discovery.
-    ## If {}, select all ServiceMonitors
+    ## If {}, select all PrometheusRules
     ##
     ruleSelector: {}
-    ## Example which select all prometheusrules resources
+    ## Example which select all PrometheusRules resources
     ## with label "prometheus" with values any of "example-rules" or "example-rules-2"
     # ruleSelector:
     #   matchExpressions:
@@ -1924,7 +1924,7 @@ prometheus:
     #         - example-rules
     #         - example-rules-2
     #
-    ## Example which select all prometheusrules resources with label "role" set to "example-rules"
+    ## Example which select all PrometheusRules resources with label "role" set to "example-rules"
     # ruleSelector:
     #   matchLabels:
     #     role: example-rules

```

## 14.1.0

**Release date:** 2021-03-18

![AppVersion: 0.46.0](https://img.shields.io/static/v1?label=AppVersion&message=0.46.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Added nodePort to thanosService (#769)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 2ec73605..961c6198 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1541,6 +1541,14 @@ prometheus:
     port: 10901
     targetPort: "grpc"
 
+    ## Service type
+    ##
+    type: ClusterIP
+
+    ## Port to expose on each node
+    ##
+    nodePort: 30901
+
   ## Configuration for Prometheus service
   ##
   service:

```

## 14.0.1

**Release date:** 2021-03-09

![AppVersion: 0.46.0](https://img.shields.io/static/v1?label=AppVersion&message=0.46.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Update README with 13.x to 14.x upgrade instructions (#746)

### Default value changes

```diff
# No changes in this release
```

## 14.0.0

**Release date:** 2021-03-08

![AppVersion: 0.46.0](https://img.shields.io/static/v1?label=AppVersion&message=0.46.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] 0.46.0 operator and crds (#740)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 66114e2e..2ec73605 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1480,7 +1480,7 @@ prometheusOperator:
   ##
   image:
     repository: quay.io/prometheus-operator/prometheus-operator
-    tag: v0.45.0
+    tag: v0.46.0
     sha: ""
     pullPolicy: IfNotPresent
 
@@ -1496,7 +1496,7 @@ prometheusOperator:
   ##
   prometheusConfigReloaderImage:
     repository: quay.io/prometheus-operator/prometheus-config-reloader
-    tag: v0.45.0
+    tag: v0.46.0
     sha: ""
 
   ## Set the prometheus config reloader side-car CPU limit

```

## 13.13.1

**Release date:** 2021-03-01

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Added missing upgrade command for alertmanager to do before 13.13.x (#715)

### Default value changes

```diff
# No changes in this release
```

## 13.13.0

**Release date:** 2021-02-23

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Upgrade to latest kube-state-metrics chart with v1.9.8 (#703)

### Default value changes

```diff
# No changes in this release
```

## 13.12.0

**Release date:** 2021-02-23

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* add Prometheus and Alertmanager topologySpreadConstraints (#702)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 8b363731..66114e2e 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -534,6 +534,17 @@ alertmanager:
     #   value: "value"
     #   effect: "NoSchedule"
 
+    ## If specified, the pod's topology spread constraints.
+    ## ref: https://kubernetes.io/docs/concepts/workloads/pods/pod-topology-spread-constraints/
+    ##
+    topologySpreadConstraints: []
+    # - maxSkew: 1
+    #   topologyKey: topology.kubernetes.io/zone
+    #   whenUnsatisfiable: DoNotSchedule
+    #   labelSelector:
+    #     matchLabels:
+    #       app: alertmanager
+
     ## SecurityContext holds pod-level security attributes and common container settings.
     ## This defaults to non root user with uid 1000 and gid 2000.	*v1.PodSecurityContext	false
     ## ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
@@ -1807,6 +1818,17 @@ prometheus:
     #    value: "value"
     #    effect: "NoSchedule"
 
+    ## If specified, the pod's topology spread constraints.
+    ## ref: https://kubernetes.io/docs/concepts/workloads/pods/pod-topology-spread-constraints/
+    ##
+    topologySpreadConstraints: []
+    # - maxSkew: 1
+    #   topologyKey: topology.kubernetes.io/zone
+    #   whenUnsatisfiable: DoNotSchedule
+    #   labelSelector:
+    #     matchLabels:
+    #       app: prometheus
+
     ## Alertmanagers to which alerts will be sent
     ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#alertmanagerendpoints
     ##

```

## 13.11.0

**Release date:** 2021-02-23

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Added pathType (#688)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 59467c07..8b363731 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -220,6 +220,10 @@ alertmanager:
     paths: []
     # - /
 
+    ## For Kubernetes >= 1.18 you should specify the pathType (determines how Ingress paths should be matched)
+    ## See https://kubernetes.io/blog/2020/04/02/improvements-to-the-ingress-api-in-kubernetes-1.18/#better-path-matching-with-path-types
+    # pathType: ImplementationSpecific
+
     ## TLS configuration for Alertmanager Ingress
     ## Secret must be manually created in the namespace
     ##
@@ -260,6 +264,10 @@ alertmanager:
     paths: []
     # - /
 
+    ## For Kubernetes >= 1.18 you should specify the pathType (determines how Ingress paths should be matched)
+    ## See https://kubernetes.io/blog/2020/04/02/improvements-to-the-ingress-api-in-kubernetes-1.18/#better-path-matching-with-path-types
+    # pathType: ImplementationSpecific
+
     ## Secret name containing the TLS certificate for alertmanager per replica ingress
     ## Secret must be manually created in the namespace
     tlsSecretName: ""
@@ -1618,6 +1626,10 @@ prometheus:
     paths: []
     # - /
 
+    ## For Kubernetes >= 1.18 you should specify the pathType (determines how Ingress paths should be matched)
+    ## See https://kubernetes.io/blog/2020/04/02/improvements-to-the-ingress-api-in-kubernetes-1.18/#better-path-matching-with-path-types
+    # pathType: ImplementationSpecific
+
     ## TLS configuration for Thanos Ingress
     ## Secret must be manually created in the namespace
     ##
@@ -1648,6 +1660,10 @@ prometheus:
     paths: []
     # - /
 
+    ## For Kubernetes >= 1.18 you should specify the pathType (determines how Ingress paths should be matched)
+    ## See https://kubernetes.io/blog/2020/04/02/improvements-to-the-ingress-api-in-kubernetes-1.18/#better-path-matching-with-path-types
+    # pathType: ImplementationSpecific
+
     ## TLS configuration for Prometheus Ingress
     ## Secret must be manually created in the namespace
     ##
@@ -1683,6 +1699,10 @@ prometheus:
     paths: []
     # - /
 
+    ## For Kubernetes >= 1.18 you should specify the pathType (determines how Ingress paths should be matched)
+    ## See https://kubernetes.io/blog/2020/04/02/improvements-to-the-ingress-api-in-kubernetes-1.18/#better-path-matching-with-path-types
+    # pathType: ImplementationSpecific
+
     ## Secret name containing the TLS certificate for Prometheus per replica ingress
     ## Secret must be manually created in the namespace
     tlsSecretName: ""

```

## 13.10.0

**Release date:** 2021-02-18

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Upgrade grafana dependency to 6.4 (#685)

### Default value changes

```diff
# No changes in this release
```

## 13.9.1

**Release date:** 2021-02-17

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Ability for custom dnsConfig (#639)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 01665960..59467c07 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1441,7 +1441,16 @@ prometheusOperator:
     #         values:
     #         - e2e-az1
     #         - e2e-az2
-
+  dnsConfig: {}
+    # nameservers:
+    #   - 1.2.3.4
+    # searches:
+    #   - ns1.svc.cluster-domain.example
+    #   - my.dns.search.suffix
+    # options:
+    #   - name: ndots
+    #     value: "2"
+  #   - name: edns0
   securityContext:
     fsGroup: 65534
     runAsGroup: 65534

```

## 13.9.0

**Release date:** 2021-02-17

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] hack: fix syncing dashboards and alerting rules (#665)

### Default value changes

```diff
# No changes in this release
```

## 13.8.0

**Release date:** 2021-02-17

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] update node exporter chart to 1.14.* (#677)

### Default value changes

```diff
# No changes in this release
```

## 13.7.2

**Release date:** 2021-02-10

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Ability for cluster-domain in kube-prometheus-stack (#647)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 0bba5e98..01665960 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1299,6 +1299,12 @@ prometheusOperator:
   prometheusInstanceNamespaces: []
   thanosRulerInstanceNamespaces: []
 
+  ## The clusterDomain value will be added to the cluster.peer option of the alertmanager.
+  ## Without this specified option cluster.peer will have value alertmanager-monitoring-alertmanager-0.alertmanager-operated:9094 (default value)
+  ## With this specified option cluster.peer will have value alertmanager-monitoring-alertmanager-0.alertmanager-operated.namespace.svc.cluster-domain:9094
+  ##
+  # clusterDomain: "cluster.local"
+
   ## Service account for Alertmanager to use.
   ## ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
   ##

```

## 13.7.1

**Release date:** 2021-02-10

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* fix yaml indent (#663)

### Default value changes

```diff
# No changes in this release
```

## 13.7.0

**Release date:** 2021-02-10

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] introduce optional param grafana.sidecar.datasources.defaultDatasourceScrapeInterval. (#659)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index caa38b02..0bba5e98 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -634,6 +634,9 @@ grafana:
       enabled: true
       defaultDatasourceEnabled: true
 
+      # If not defined, will use prometheus.prometheusSpec.scrapeInterval or its default
+      # defaultDatasourceScrapeInterval: 15s
+
       ## Annotations for Grafana datasource configmaps
       ##
       annotations: {}

```

## 13.6.0

**Release date:** 2021-02-09

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* fix: allow overriding of selector for kube-state-metrics serviceMonitor (#303)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index f367e960..caa38b02 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1151,6 +1151,9 @@ kubeStateMetrics:
     ## Scrape interval. If not set, the Prometheus default scrape interval is used.
     ##
     interval: ""
+    ## Override serviceMonitor selector
+    ##
+    selectorOverride: {}
 
     ## 	metric relabel configs to apply to samples before ingestion.
     ##

```

## 13.5.1

**Release date:** 2021-02-08

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Sync kube-state-metrics (#650)

### Default value changes

```diff
# No changes in this release
```

## 13.5.0

**Release date:** 2021-02-04

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Add podTargetLabels to additionalServiceMonitors (#629)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 617b39cc..f367e960 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -2213,7 +2213,11 @@ prometheus:
 
     ## labels to transfer from the kubernetes service to the target
     ##
-    # targetLabels: ""
+    # targetLabels: []
+
+    ## labels to transfer from the kubernetes pods to the target
+    ##
+    # podTargetLabels: []
 
     ## Label selector for services to which this ServiceMonitor applies
     ##

```

## 13.4.1

**Release date:** 2021-01-29

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* fix: match containerPort and listen-address in operator (#624)

### Default value changes

```diff
# No changes in this release
```

## 13.4.0

**Release date:** 2021-01-29

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add support for specifying webhook port for GKE workaround (#400)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 2a0f98c0..617b39cc 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1232,11 +1232,14 @@ prometheus-node-exporter:
 prometheusOperator:
   enabled: true
 
-  # Prometheus-Operator v0.39.0 and later support TLS natively.
+  ## Prometheus-Operator v0.39.0 and later support TLS natively.
+  ##
   tls:
     enabled: true
     # Value must match version names from https://golang.org/pkg/crypto/tls/#pkg-constants
     tlsMinVersion: VersionTLS13
+    # The default webhook port is 10250 in order to work out-of-the-box in GKE private clusters and avoid adding firewall rules.
+    internalPort: 10250
 
   ## Admission webhook support for PrometheusRules resources added in Prometheus Operator 0.30 can be enabled to prevent incorrectly formatted
   ## rules from making their way into prometheus and potentially preventing the container from starting

```

## 13.3.0

**Release date:** 2021-01-27

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Allow override of prometheus and alertmanager default base images (#588)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 2be42c80..2a0f98c0 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1441,6 +1441,14 @@ prometheusOperator:
     sha: ""
     pullPolicy: IfNotPresent
 
+  ## Prometheus image to use for prometheuses managed by the operator
+  ##
+  # prometheusDefaultBaseImage: quay.io/prometheus/prometheus
+
+  ## Alertmanager image to use for alertmanagers managed by the operator
+  ##
+  # alertmanagerDefaultBaseImage: quay.io/prometheus/alertmanager
+
   ## Prometheus-config-reloader image to use for config and rule reloading
   ##
   prometheusConfigReloaderImage:

```

## 13.2.1

**Release date:** 2021-01-22

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [Kube-prometheus-stack]adding upstream for external label and url (#536)

### Default value changes

```diff
# No changes in this release
```

## 13.2.0

**Release date:** 2021-01-21

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Make it possible to add additional volume and allowerHostPath to prometehus psp (#583)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 8fbccb3d..2be42c80 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1668,6 +1668,8 @@ prometheus:
   ## ref: https://kubernetes.io/docs/concepts/policy/pod-security-policy/
   podSecurityPolicy:
     allowedCapabilities: []
+    allowedHostPaths: []
+    volumes: []
 
   serviceMonitor:
     ## Scrape interval. If not set, the Prometheus default scrape interval is used.

```

## 13.1.2

**Release date:** 2021-01-21

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Update list of excluded fs types and mountpoints in node exporter collector (#594)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index dc497e96..8fbccb3d 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1224,8 +1224,8 @@ prometheus-node-exporter:
     ##
     jobLabel: node-exporter
   extraArgs:
-    - --collector.filesystem.ignored-mount-points=^/(dev|proc|sys|var/lib/docker/.+)($|/)
-    - --collector.filesystem.ignored-fs-types=^(autofs|binfmt_misc|cgroup|configfs|debugfs|devpts|devtmpfs|fusectl|hugetlbfs|mqueue|overlay|proc|procfs|pstore|rpc_pipefs|securityfs|sysfs|tracefs)$
+    - --collector.filesystem.ignored-mount-points=^/(dev|proc|sys|var/lib/docker/.+|var/lib/kubelet/.+)($|/)
+    - --collector.filesystem.ignored-fs-types=^(autofs|binfmt_misc|bpf|cgroup2?|configfs|debugfs|devpts|devtmpfs|fusectl|hugetlbfs|iso9660|mqueue|nsfs|overlay|proc|procfs|pstore|rpc_pipefs|securityfs|selinuxfs|squashfs|sysfs|tracefs)$
 
 ## Manages Prometheus and Alertmanager components
 ##

```

## 13.1.1

**Release date:** 2021-01-20

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add config.templates (#563)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 32730b57..dc497e96 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -158,6 +158,8 @@ alertmanager:
         receiver: 'null'
     receivers:
     - name: 'null'
+    templates:
+    - '/etc/alertmanager/config/*.tmpl'
 
   ## Pass the Alertmanager configuration directives through Helm's templating
   ## engine. If the Alertmanager configuration contains Alertmanager templates,
@@ -170,6 +172,10 @@ alertmanager:
   tplConfig: false
 
   ## Alertmanager template files to format alerts
+  ## By default, templateFiles are placed in /etc/alertmanager/config/ and if
+  ## they have a .tmpl file suffix will be loaded. See config.templates above
+  ## to change, add other suffixes. If adding other suffixes, be sure to update
+  ## config.templates above to include those suffixes.
   ## ref: https://prometheus.io/docs/alerting/notifications/
   ##      https://prometheus.io/docs/alerting/notification_examples/
   ##

```

## 13.0.5

**Release date:** 2021-01-20

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Use new API version for ingress (#572)

### Default value changes

```diff
# No changes in this release
```

## 13.0.4

**Release date:** 2021-01-20

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Document cert-manager usage (#574)
* [kube-prometheus-stack]: Add upgrade notes for v13 (#596)

### Default value changes

```diff
# No changes in this release
```

## 13.0.3

**Release date:** 2021-01-20

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Fix thanos grpc port when using service type NodePort (#593)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index dbf5ef62..32730b57 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1564,6 +1564,12 @@ prometheus:
     annotations: {}
     labels: {}
     servicePort: 10901
+
+    ## Port to expose on each node
+    ## Only used if service.type is 'NodePort'
+    ##
+    nodePort: 30901
+
     ## Hosts must be provided if Ingress is enabled.
     ##
     hosts: []

```

## 13.0.2

**Release date:** 2021-01-19

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack]: enable multi cluster dashboards (#575)

### Default value changes

```diff
# No changes in this release
```

## 13.0.1

**Release date:** 2021-01-19

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add grpc thanos port to prometheus service (#504)

### Default value changes

```diff
# No changes in this release
```

## 13.0.0

**Release date:** 2021-01-18

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Version bump prometheus-operator to 0.45.0 (#582)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index b83d3894..dbf5ef62 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1431,22 +1431,15 @@ prometheusOperator:
   ##
   image:
     repository: quay.io/prometheus-operator/prometheus-operator
-    tag: v0.44.0
+    tag: v0.45.0
     sha: ""
     pullPolicy: IfNotPresent
 
-  ## Configmap-reload image to use for reloading configmaps
-  ##
-  configmapReloadImage:
-    repository: docker.io/jimmidyson/configmap-reload
-    tag: v0.4.0
-    sha: ""
-
   ## Prometheus-config-reloader image to use for config and rule reloading
   ##
   prometheusConfigReloaderImage:
     repository: quay.io/prometheus-operator/prometheus-config-reloader
-    tag: v0.44.0
+    tag: v0.45.0
     sha: ""
 
   ## Set the prometheus config reloader side-car CPU limit
@@ -1736,7 +1729,7 @@ prometheus:
     ##
     image:
       repository: quay.io/prometheus/prometheus
-      tag: v2.22.1
+      tag: v2.24.0
       sha: ""
 
     ## Tolerations for use with node taints

```

## 12.12.2

**Release date:** 2021-01-18

![AppVersion: 0.44.0](https://img.shields.io/static/v1?label=AppVersion&message=0.44.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* fix chart version (#589)
* Use matchExpressions instead of matchLabels in podAntiAffinity rules (#438)

### Default value changes

```diff
# No changes in this release
```

## 12.12.1

**Release date:** 2021-01-14

![AppVersion: 0.44.0](https://img.shields.io/static/v1?label=AppVersion&message=0.44.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Migrate kube-state-metrics off helm/stable (#581)

### Default value changes

```diff
# No changes in this release
```

## 12.12.0

**Release date:** 2021-01-13

![AppVersion: 0.44.0](https://img.shields.io/static/v1?label=AppVersion&message=0.44.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Add option to add headless svc for thanos sidecar (#564)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 4c46c131..b83d3894 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1478,6 +1478,19 @@ prometheus:
     create: true
     name: ""
 
+  # Service for thanos service discovery on sidecar
+  # Enable this can make Thanos Query can use
+  # `--store=dnssrv+_grpc._tcp.${kube-prometheus-stack.fullname}-thanos-discovery.${namespace}.svc.cluster.local` to discovery
+  # Thanos sidecar on prometheus nodes
+  # (Please remember to change ${kube-prometheus-stack.fullname} and ${namespace}. Not just copy and paste!)
+  thanosService:
+    enabled: false
+    annotations: {}
+    labels: {}
+    portName: grpc
+    port: 10901
+    targetPort: "grpc"
+
   ## Configuration for Prometheus service
   ##
   service:

```

## 12.11.3

**Release date:** 2021-01-10

![AppVersion: 0.44.0](https://img.shields.io/static/v1?label=AppVersion&message=0.44.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack]values:fix missleading comment (#553)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index b9482cf7..4c46c131 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1843,9 +1843,12 @@ prometheus:
     #     prometheus: somelabel
 
     ## Namespaces to be selected for ServiceMonitor discovery.
-    ## See https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#namespaceselector for usage
     ##
     serviceMonitorNamespaceSelector: {}
+    ## Example which selects ServiceMonitors in namespaces with label "prometheus" set to "somelabel"
+    # serviceMonitorNamespaceSelector:
+    #   matchLabels:
+    #     prometheus: somelabel
 
     ## If true, a nil or {} value for prometheus.prometheusSpec.podMonitorSelector will cause the
     ## prometheus resource to be created with selectors based on values in the helm deployment,

```

## 12.11.2

**Release date:** 2021-01-10

![AppVersion: 0.44.0](https://img.shields.io/static/v1?label=AppVersion&message=0.44.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] prometheus: add sharding mechanism (#552)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index bb93e359..b9482cf7 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1903,10 +1903,20 @@ prometheus:
     ##
     paused: false
 
-    ## Number of Prometheus replicas desired
+    ## Number of replicas of each shard to deploy for a Prometheus deployment.
+    ## Number of replicas multiplied by shards is the total number of Pods created.
     ##
     replicas: 1
 
+    ## EXPERIMENTAL: Number of shards to distribute targets onto.
+    ## Number of replicas multiplied by shards is the total number of Pods created.
+    ## Note that scaling down shards will not reshard data onto remaining instances, it must be manually moved.
+    ## Increasing shards will not reshard data either but it will continue to be available from the same instances.
+    ## To query globally use Thanos sidecar and Thanos querier or remote write data to a central location.
+    ## Sharding is done on the content of the `__address__` target meta-label.
+    ##
+    shards: 1
+
     ## Log level for Prometheus be configured in
     ##
     logLevel: info

```

## 12.10.6

**Release date:** 2021-01-05

![AppVersion: 0.44.0](https://img.shields.io/static/v1?label=AppVersion&message=0.44.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Allow setting caBundle value for admission webhook configuration (#543)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index a864b592..bb93e359 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1237,6 +1237,9 @@ prometheusOperator:
   admissionWebhooks:
     failurePolicy: Fail
     enabled: true
+    ## A PEM encoded CA bundle which will be used to validate the webhook's server certificate.
+    ## If unspecified, system trust roots on the apiserver are used.
+    caBundle: ""
     ## If enabled, generate a self-signed certificate, then patch the webhook configurations with the generated data.
     ## On chart upgrades (or if the secret exists) the cert will not be re-generated. You can use this to provide your own
     ## certs ahead of time if you wish.

```

## 12.10.5

**Release date:** 2021-01-05

![AppVersion: 0.44.0](https://img.shields.io/static/v1?label=AppVersion&message=0.44.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] 333 - need to template the alertmanager tls settings as well (#539)

### Default value changes

```diff
# No changes in this release
```

## 12.10.4

**Release date:** 2021-01-04

![AppVersion: 0.44.0](https://img.shields.io/static/v1?label=AppVersion&message=0.44.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] allow to select alertmanagerConfig resources from all namespaces (#503)

### Default value changes

```diff
# No changes in this release
```

## 12.10.3

**Release date:** 2021-01-04

![AppVersion: 0.44.0](https://img.shields.io/static/v1?label=AppVersion&message=0.44.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Fix accessing root in range block for datasources.timeInterval (#518)

### Default value changes

```diff
# No changes in this release
```

## 12.10.2

**Release date:** 2021-01-04

![AppVersion: 0.44.0](https://img.shields.io/static/v1?label=AppVersion&message=0.44.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Avoid relying on image/container working directory for TLS cert mounting in the `prometheus-operator` `Deployment` (#530)

### Default value changes

```diff
# No changes in this release
```

## 12.10.1

**Release date:** 2021-01-04

![AppVersion: 0.44.0](https://img.shields.io/static/v1?label=AppVersion&message=0.44.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Fix an errant dash in prometheus-operator's --kubelet-service arg (#535)

### Default value changes

```diff
# No changes in this release
```

## 12.10.0

**Release date:** 2021-01-04

![AppVersion: 0.44.0](https://img.shields.io/static/v1?label=AppVersion&message=0.44.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Upgrade grafan dependency to 6.1, bump major version to match grafana breaking change, update prometheusOperator.configReloaderMemory from 25Mi to 50Mi to follow 0.44.0 change. (#460)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 690b53ef..a864b592 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1452,7 +1452,7 @@ prometheusOperator:
 
   ## Set the prometheus config reloader side-car memory limit
   ##
-  configReloaderMemory: 25Mi
+  configReloaderMemory: 50Mi
 
   ## Set a Field Selector to filter watched secrets
   ##

```

## 12.9.2

**Release date:** 2021-01-01

![AppVersion: 0.44.0](https://img.shields.io/static/v1?label=AppVersion&message=0.44.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Correct path to `ServiceMonitor` ca cert when running with tls enabled (#531)

### Default value changes

```diff
# No changes in this release
```

## 12.9.1

**Release date:** 2021-01-01

![AppVersion: 0.44.0](https://img.shields.io/static/v1?label=AppVersion&message=0.44.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Fix Capabilities check for IngressClass (#489)

### Default value changes

```diff
# No changes in this release
```

## 12.9.0

**Release date:** 2020-12-30

![AppVersion: 0.44.0](https://img.shields.io/static/v1?label=AppVersion&message=0.44.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] cert-manager support for operator webhooks (#481)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index e2c481fe..690b53ef 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1256,6 +1256,12 @@ prometheusOperator:
       nodeSelector: {}
       affinity: {}
       tolerations: []
+    # Use certmanager to generate webhook certs
+    certManager:
+      enabled: false
+      # issuerRef:
+      #   name: "issuer"
+      #   kind: "ClusterIssuer"
 
   ## Namespaces to scope the interaction of the Prometheus Operator and the apiserver (allow list).
   ## This is mutually exclusive with denyNamespaces. Setting this to an empty object will disable the configuration

```

## 12.8.1

**Release date:** 2020-12-22

![AppVersion: 0.44.0](https://img.shields.io/static/v1?label=AppVersion&message=0.44.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* fixed typos (#496)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 3aa810fa..e2c481fe 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1538,7 +1538,7 @@ prometheus:
     minAvailable: 1
     maxUnavailable: ""
 
-  # Ingress exposes thanos sidecar outside the clsuter
+  # Ingress exposes thanos sidecar outside the cluster
   thanosIngress:
     enabled: false
 
@@ -1559,7 +1559,7 @@ prometheus:
     paths: []
     # - /
 
-    ## TLS configuration for Alertmanager Ingress
+    ## TLS configuration for Thanos Ingress
     ## Secret must be manually created in the namespace
     ##
     tls: []

```

## 12.8.0

**Release date:** 2020-12-11

![AppVersion: 0.44.0](https://img.shields.io/static/v1?label=AppVersion&message=0.44.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] #356 add support for multicluster in grafana dashboards (#357)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 774359fe..3aa810fa 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -623,6 +623,7 @@ grafana:
       ## Annotations for Grafana dashboard configmaps
       ##
       annotations: {}
+      multicluster: false
     datasources:
       enabled: true
       defaultDatasourceEnabled: true

```

## 12.7.0

**Release date:** 2020-12-07

![AppVersion: 0.44.0](https://img.shields.io/static/v1?label=AppVersion&message=0.44.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Configure possibility of providing additional rules to prometheus cluster role (#413)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index f243f8c4..774359fe 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -2136,6 +2136,12 @@ prometheus:
     ## in Prometheus so it may change in any upcoming release.
     allowOverlappingBlocks: false
 
+  additionalRulesForClusterRole: []
+  #  - apiGroups: [ "" ]
+  #    resources:
+  #      - nodes/proxy
+  #    verbs: [ "get", "list", "watch" ]
+
   additionalServiceMonitors: []
   ## Name of the ServiceMonitor to create
   ##

```

## 12.6.0

**Release date:** 2020-12-07

![AppVersion: 0.44.0](https://img.shields.io/static/v1?label=AppVersion&message=0.44.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Align Grafana and Prometheus scrape intervals (#457)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index c2c99494..f243f8c4 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1686,6 +1686,8 @@ prometheus:
     apiserverConfig: {}
 
     ## Interval between consecutive scrapes.
+    ## Defaults to 30s.
+    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/release-0.44/pkg/prometheus/promcfg.go#L180-L183
     ##
     scrapeInterval: ""
 

```

## 12.5.0

**Release date:** 2020-12-07

![AppVersion: 0.44.0](https://img.shields.io/static/v1?label=AppVersion&message=0.44.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add missing fields to prometheus and alertmanager resources (#208)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index f768a47b..c2c99494 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -545,6 +545,10 @@ alertmanager:
     # Additional VolumeMounts on the output StatefulSet definition.
     volumeMounts: []
 
+    ## InitContainers allows injecting additional initContainers. This is meant to allow doing some changes
+    ## (permissions, dir tree) on mounted volumes before starting prometheus
+    initContainers: []
+
     ## Priority class assigned to the Pods
     ##
     priorityClassName: ""
@@ -561,6 +565,10 @@ alertmanager:
     ##
     clusterAdvertiseAddress: false
 
+    ## ForceEnableClusterMode ensures Alertmanager does not deactivate the cluster mode when running with a single replica.
+    ## Use case is e.g. spanning an Alertmanager cluster across Kubernetes clusters with a single replica in each.
+    forceEnableClusterMode: false
+
 
 ## Using default values from https://github.com/grafana/helm-charts/blob/main/charts/grafana/values.yaml
 ##
@@ -2092,6 +2100,40 @@ prometheus:
     ##
     portName: "web"
 
+    ## ArbitraryFSAccessThroughSMs configures whether configuration based on a service monitor can access arbitrary files
+    ## on the file system of the Prometheus container e.g. bearer token files.
+    arbitraryFSAccessThroughSMs: false
+
+    ## OverrideHonorLabels if set to true overrides all user configured honor_labels. If HonorLabels is set in ServiceMonitor
+    ## or PodMonitor to true, this overrides honor_labels to false.
+    overrideHonorLabels: false
+
+    ## OverrideHonorTimestamps allows to globally enforce honoring timestamps in all scrape configs.
+    overrideHonorTimestamps: false
+
+    ## IgnoreNamespaceSelectors if set to true will ignore NamespaceSelector settings from the podmonitor and servicemonitor
+    ## configs, and they will only discover endpoints within their current namespace. Defaults to false.
+    ignoreNamespaceSelectors: false
+
+    ## PrometheusRulesExcludedFromEnforce - list of prometheus rules to be excluded from enforcing of adding namespace labels.
+    ## Works only if enforcedNamespaceLabel set to true. Make sure both ruleNamespace and ruleName are set for each pair
+    prometheusRulesExcludedFromEnforce: false
+
+    ## QueryLogFile specifies the file to which PromQL queries are logged. Note that this location must be writable,
+    ## and can be persisted using an attached volume. Alternatively, the location can be set to a stdout location such
+    ## as /dev/stdout to log querie information to the default Prometheus log stream. This is only available in versions
+    ## of Prometheus >= 2.16.0. For more details, see the Prometheus docs (https://prometheus.io/docs/guides/query-log/)
+    queryLogFile: false
+
+    ## EnforcedSampleLimit defines global limit on number of scraped samples that will be accepted. This overrides any SampleLimit
+    ## set per ServiceMonitor or/and PodMonitor. It is meant to be used by admins to enforce the SampleLimit to keep overall
+    ## number of samples/series under the desired limit. Note that if SampleLimit is lower that value will be taken instead.
+    enforcedSampleLimit: false
+
+    ## AllowOverlappingBlocks enables vertical compaction and vertical query merge in Prometheus. This is still experimental
+    ## in Prometheus so it may change in any upcoming release.
+    allowOverlappingBlocks: false
+
   additionalServiceMonitors: []
   ## Name of the ServiceMonitor to create
   ##

```

## 12.4.1

**Release date:** 2020-12-07

![AppVersion: 0.44.0](https://img.shields.io/static/v1?label=AppVersion&message=0.44.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Enable AggregatedAPIDown for k8s>=1.18 only (#455)

### Default value changes

```diff
# No changes in this release
```

## 12.4.0

**Release date:** 2020-12-02

![AppVersion: 0.44.0](https://img.shields.io/static/v1?label=AppVersion&message=0.44.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Update prometheus-operator to v0.44.0 (#441)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 5d0f47e3..f768a47b 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1413,7 +1413,7 @@ prometheusOperator:
   ##
   image:
     repository: quay.io/prometheus-operator/prometheus-operator
-    tag: v0.43.2
+    tag: v0.44.0
     sha: ""
     pullPolicy: IfNotPresent
 
@@ -1428,7 +1428,7 @@ prometheusOperator:
   ##
   prometheusConfigReloaderImage:
     repository: quay.io/prometheus-operator/prometheus-config-reloader
-    tag: v0.43.2
+    tag: v0.44.0
     sha: ""
 
   ## Set the prometheus config reloader side-car CPU limit

```

## 12.3.0

**Release date:** 2020-11-30

![AppVersion: 0.43.2](https://img.shields.io/static/v1?label=AppVersion&message=0.43.2&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Sync grafana dashboards and prometheus rules from kube-prometheus (#432)

### Default value changes

```diff
# No changes in this release
```

## 12.2.4

**Release date:** 2020-11-28

![AppVersion: 0.43.2](https://img.shields.io/static/v1?label=AppVersion&message=0.43.2&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [codespell] fix readme (#427)

### Default value changes

```diff
# No changes in this release
```

## 12.2.3

**Release date:** 2020-11-24

![AppVersion: 0.43.2](https://img.shields.io/static/v1?label=AppVersion&message=0.43.2&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Fix broken volumes and volumeMounts fields on alertmanager (#409)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 76b72adb..5d0f47e3 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -539,6 +539,12 @@ alertmanager:
     ##
     containers: []
 
+    # Additional volumes on the output StatefulSet definition.
+    volumes: []
+
+    # Additional VolumeMounts on the output StatefulSet definition.
+    volumeMounts: []
+
     ## Priority class assigned to the Pods
     ##
     priorityClassName: ""
@@ -1968,6 +1974,7 @@ prometheus:
 
     # Additional volumes on the output StatefulSet definition.
     volumes: []
+
     # Additional VolumeMounts on the output StatefulSet definition.
     volumeMounts: []
 

```

## 12.2.2

**Release date:** 2020-11-23

![AppVersion: 0.43.2](https://img.shields.io/static/v1?label=AppVersion&message=0.43.2&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Allow custom TLS min version for operator (#407)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 3371203e..76b72adb 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1214,6 +1214,8 @@ prometheusOperator:
   # Prometheus-Operator v0.39.0 and later support TLS natively.
   tls:
     enabled: true
+    # Value must match version names from https://golang.org/pkg/crypto/tls/#pkg-constants
+    tlsMinVersion: VersionTLS13
 
   ## Admission webhook support for PrometheusRules resources added in Prometheus Operator 0.30 can be enabled to prevent incorrectly formatted
   ## rules from making their way into prometheus and potentially preventing the container from starting

```

## 12.2.1

**Release date:** 2020-11-23

![AppVersion: 0.43.2](https://img.shields.io/static/v1?label=AppVersion&message=0.43.2&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* kube-prometheus-stack: fix ports alignment in service (#404)

### Default value changes

```diff
# No changes in this release
```

## 12.2.0

**Release date:** 2020-11-21

![AppVersion: 0.43.2](https://img.shields.io/static/v1?label=AppVersion&message=0.43.2&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* kube-prometheus-stack: Add optional extra ports to alertmanager (#396)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 4a7e27df..3371203e 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -287,6 +287,10 @@ alertmanager:
     ## List of IP addresses at which the Prometheus server service is available
     ## Ref: https://kubernetes.io/docs/user-guide/services/#external-ips
     ##
+
+    ## Additional ports to open for Alertmanager service
+    additionalPorts: []
+
     externalIPs: []
     loadBalancerIP: ""
     loadBalancerSourceRanges: []

```

## 12.1.0

**Release date:** 2020-11-19

![AppVersion: 0.43.2](https://img.shields.io/static/v1?label=AppVersion&message=0.43.2&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Allow rules be enabled independent of exporters (#330)

### Default value changes

```diff
# No changes in this release
```

## 12.0.4

**Release date:** 2020-11-18

![AppVersion: 0.43.2](https://img.shields.io/static/v1?label=AppVersion&message=0.43.2&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Clarify log format options in kube-prometheus-stack (#382)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 79803a72..4a7e27df 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -426,7 +426,7 @@ alertmanager:
     #     alertmanagerconfig: enabled
 
     ## Define Log Format
-    # Use logfmt (default) or json-formatted logging
+    # Use logfmt (default) or json logging
     logFormat: logfmt
 
     ## Log level for Alertmanager to be configured with.
@@ -1308,7 +1308,7 @@ prometheusOperator:
   # priorityClassName: ""
 
   ## Define Log Format
-  # Use logfmt (default) or json-formatted logging
+  # Use logfmt (default) or json logging
   # logFormat: logfmt
 
   ## Decrease log verbosity to errors only

```

## 12.0.3

**Release date:** 2020-11-18

![AppVersion: 0.43.2](https://img.shields.io/static/v1?label=AppVersion&message=0.43.2&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [fix] Correct provisioning for the --thanos-ruler-instance-namespaces prometheus-operator (#386)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 17e5d7d7..79803a72 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1252,7 +1252,7 @@ prometheusOperator:
   ##
   alertmanagerInstanceNamespaces: []
   prometheusInstanceNamespaces: []
-  thanosInstanceNamespaces: []
+  thanosRulerInstanceNamespaces: []
 
   ## Service account for Alertmanager to use.
   ## ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/

```

## 12.0.2

**Release date:** 2020-11-18

![AppVersion: 0.43.2](https://img.shields.io/static/v1?label=AppVersion&message=0.43.2&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Fix some GitHub broken links on kube-prometheus-stack/hack/README.md (#344)

### Default value changes

```diff
# No changes in this release
```

## 12.0.1

**Release date:** 2020-11-17

![AppVersion: 0.43.2](https://img.shields.io/static/v1?label=AppVersion&message=0.43.2&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add missing parameter scrapetimeout (#370)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 03f759bb..17e5d7d7 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1669,6 +1669,10 @@ prometheus:
     ##
     scrapeInterval: ""
 
+    ## Number of seconds to wait for target to respond before erroring
+    ##
+    scrapeTimeout: ""
+
     ## Interval between consecutive evaluations.
     ##
     evaluationInterval: ""

```

## 12.0.0

**Release date:** 2020-11-17

![AppVersion: 0.43.2](https://img.shields.io/static/v1?label=AppVersion&message=0.43.2&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Migrate to chart v2 (#307)

### Default value changes

```diff
# No changes in this release
```

## 11.1.7

**Release date:** 2020-11-17

![AppVersion: 0.43.2](https://img.shields.io/static/v1?label=AppVersion&message=0.43.2&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Fix indentation for applying labels to ingressPerReplica (#347)

### Default value changes

```diff
# No changes in this release
```

## 11.1.6

**Release date:** 2020-11-16

![AppVersion: 0.43.2](https://img.shields.io/static/v1?label=AppVersion&message=0.43.2&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Issue #333 (#334)

### Default value changes

```diff
# No changes in this release
```

## 11.1.5

**Release date:** 2020-11-16

![AppVersion: 0.43.2](https://img.shields.io/static/v1?label=AppVersion&message=0.43.2&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Show tmpfs option for Prometheus storageSpec (#363)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index efa76c75..03f759bb 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1940,6 +1940,8 @@ prometheus:
     ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/user-guides/storage.md
     ##
     storageSpec: {}
+    ## Using PersistentVolumeClaim
+    ##
     #  volumeClaimTemplate:
     #    spec:
     #      storageClassName: gluster
@@ -1949,6 +1951,11 @@ prometheus:
     #          storage: 50Gi
     #    selector: {}
 
+    ## Using tmpfs volume
+    ##
+    #  emptyDir:
+    #    medium: Memory
+
     # Additional volumes on the output StatefulSet definition.
     volumes: []
     # Additional VolumeMounts on the output StatefulSet definition.

```

## 11.1.4

**Release date:** 2020-11-16

![AppVersion: 0.43.2](https://img.shields.io/static/v1?label=AppVersion&message=0.43.2&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Fix CoreDNS dashboard for v1.7.0+ (#264)

### Default value changes

```diff
# No changes in this release
```

## 11.1.3

**Release date:** 2020-11-15

![AppVersion: 0.43.2](https://img.shields.io/static/v1?label=AppVersion&message=0.43.2&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Allow modifying ServiceMonitor path (#364)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 0c83259e..efa76c75 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -655,6 +655,10 @@ grafana:
     interval: ""
     selfMonitor: true
 
+    # Path to use for scraping metrics. Might be different if server.root_url is set
+    # in grafana.ini
+    path: "/metrics"
+
     ## 	metric relabel configs to apply to samples before ingestion.
     ##
     metricRelabelings: []

```

## 11.1.2

**Release date:** 2020-11-15

![AppVersion: 0.43.2](https://img.shields.io/static/v1?label=AppVersion&message=0.43.2&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Fix indentation to allow IDE collapse of sections (#350)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 8a678163..0c83259e 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1513,7 +1513,7 @@ prometheus:
     minAvailable: 1
     maxUnavailable: ""
 
-# Ingress exposes thanos sidecar outside the clsuter
+  # Ingress exposes thanos sidecar outside the clsuter
   thanosIngress:
     enabled: false
 

```

## 11.1.1

**Release date:** 2020-11-09

![AppVersion: 0.43.2](https://img.shields.io/static/v1?label=AppVersion&message=0.43.2&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* charts/kube-prometheus-stack: fix typos (#318)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index d3012b76..8a678163 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1958,7 +1958,7 @@ prometheus:
     ## to break upgrades of Prometheus. It is advised to review Prometheus release notes to ensure that no incompatible
     ## scrape configs are going to break Prometheus after the upgrade.
     ##
-    ## The scrape configuraiton example below will find master nodes, provided they have the name .*mst.*, relabel the
+    ## The scrape configuration example below will find master nodes, provided they have the name .*mst.*, relabel the
     ## port to 2379 and allow etcd scraping provided it is running on all Kubernetes master nodes
     ##
     additionalScrapeConfigs: []

```

## 11.1.0

**Release date:** 2020-11-09

![AppVersion: 0.43.2](https://img.shields.io/static/v1?label=AppVersion&message=0.43.2&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add alertmanagerConfigSelector option (#311)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index a253f599..d3012b76 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -389,6 +389,42 @@ alertmanager:
     ##
     # configSecret:
 
+    ## AlertmanagerConfigs to be selected to merge and configure Alertmanager with.
+    ##
+    alertmanagerConfigSelector: {}
+    ## Example which selects all alertmanagerConfig resources
+    ## with label "alertconfig" with values any of "example-config" or "example-config-2"
+    # alertmanagerConfigSelector:
+    #   matchExpressions:
+    #     - key: alertconfig
+    #       operator: In
+    #       values:
+    #         - example-config
+    #         - example-config-2
+    #
+    ## Example which selects all alertmanagerConfig resources with label "role" set to "example-config"
+    # alertmanagerConfigSelector:
+    #   matchLabels:
+    #     role: example-config
+
+    ## Namespaces to be selected for AlertmanagerConfig discovery. If nil, only check own namespace.
+    ##
+    alertmanagerConfigNamespaceSelector: {}
+    ## Example which selects all namespaces
+    ## with label "alertmanagerconfig" with values any of "example-namespace" or "example-namespace-2"
+    # alertmanagerConfigNamespaceSelector:
+    #   matchExpressions:
+    #     - key: alertmanagerconfig
+    #       operator: In
+    #       values:
+    #         - example-namespace
+    #         - example-namespace-2
+
+    ## Example which selects all namespaces with label "alertmanagerconfig" set to "enabled"
+    # alertmanagerConfigNamespaceSelector:
+    #   matchLabels:
+    #     alertmanagerconfig: enabled
+
     ## Define Log Format
     # Use logfmt (default) or json-formatted logging
     logFormat: logfmt

```

## 11.0.5

**Release date:** 2020-11-09

![AppVersion: 0.43.2](https://img.shields.io/static/v1?label=AppVersion&message=0.43.2&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Bump prometheus-operator to 0.43.2 (#328)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 5db649d5..a253f599 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1361,7 +1361,7 @@ prometheusOperator:
   ##
   image:
     repository: quay.io/prometheus-operator/prometheus-operator
-    tag: v0.43.1
+    tag: v0.43.2
     sha: ""
     pullPolicy: IfNotPresent
 
@@ -1376,7 +1376,7 @@ prometheusOperator:
   ##
   prometheusConfigReloaderImage:
     repository: quay.io/prometheus-operator/prometheus-config-reloader
-    tag: v0.43.1
+    tag: v0.43.2
     sha: ""
 
   ## Set the prometheus config reloader side-car CPU limit

```

## 11.0.4

**Release date:** 2020-11-06

![AppVersion: 0.43.1](https://img.shields.io/static/v1?label=AppVersion&message=0.43.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] add ingressclass support (#321)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 1bd16bc1..5db649d5 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -196,6 +196,10 @@ alertmanager:
   ingress:
     enabled: false
 
+    # For Kubernetes >= 1.18 you should specify the ingress-controller via the field ingressClassName
+    # See https://kubernetes.io/blog/2020/04/02/improvements-to-the-ingress-api-in-kubernetes-1.18/#specifying-the-class-of-an-ingress
+    # ingressClassName: nginx
+
     annotations: {}
 
     labels: {}
@@ -228,6 +232,11 @@ alertmanager:
   ##
   ingressPerReplica:
     enabled: false
+
+    # For Kubernetes >= 1.18 you should specify the ingress-controller via the field ingressClassName
+    # See https://kubernetes.io/blog/2020/04/02/improvements-to-the-ingress-api-in-kubernetes-1.18/#specifying-the-class-of-an-ingress
+    # ingressClassName: nginx
+
     annotations: {}
     labels: {}
 
@@ -1471,6 +1480,11 @@ prometheus:
 # Ingress exposes thanos sidecar outside the clsuter
   thanosIngress:
     enabled: false
+
+    # For Kubernetes >= 1.18 you should specify the ingress-controller via the field ingressClassName
+    # See https://kubernetes.io/blog/2020/04/02/improvements-to-the-ingress-api-in-kubernetes-1.18/#specifying-the-class-of-an-ingress
+    # ingressClassName: nginx
+
     annotations: {}
     labels: {}
     servicePort: 10901
@@ -1494,6 +1508,11 @@ prometheus:
 
   ingress:
     enabled: false
+
+    # For Kubernetes >= 1.18 you should specify the ingress-controller via the field ingressClassName
+    # See https://kubernetes.io/blog/2020/04/02/improvements-to-the-ingress-api-in-kubernetes-1.18/#specifying-the-class-of-an-ingress
+    # ingressClassName: nginx
+
     annotations: {}
     labels: {}
 
@@ -1522,6 +1541,11 @@ prometheus:
   ##
   ingressPerReplica:
     enabled: false
+
+    # For Kubernetes >= 1.18 you should specify the ingress-controller via the field ingressClassName
+    # See https://kubernetes.io/blog/2020/04/02/improvements-to-the-ingress-api-in-kubernetes-1.18/#specifying-the-class-of-an-ingress
+    # ingressClassName: nginx
+
     annotations: {}
     labels: {}
 

```

## 11.0.3

**Release date:** 2020-11-06

![AppVersion: 0.43.1](https://img.shields.io/static/v1?label=AppVersion&message=0.43.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Drop kubectl image field (#312)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 5a0810ac..1bd16bc1 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1382,14 +1382,6 @@ prometheusOperator:
   ##
   secretFieldSelector: ""
 
-  ## kubectl image to use when cleaning up
-  ##
-  kubectlImage:
-    repository: docker.io/bitnami/kubectl
-    tag: 1.16.15
-    sha: ""
-    pullPolicy: IfNotPresent
-
 ## Deploy a Prometheus instance
 ##
 prometheus:

```

## 11.0.2

**Release date:** 2020-11-05

![AppVersion: 0.43.1](https://img.shields.io/static/v1?label=AppVersion&message=0.43.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Disable scraping kubelet resource metrics endpoint (#301)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 1aded8ff..5a0810ac 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -689,8 +689,9 @@ kubelet:
     probes: true
 
     ## Enable scraping /metrics/resource from kubelet's service
+    ## This is disabled by default because container metrics are already exposed by cAdvisor
     ##
-    resource: true
+    resource: false
     # From kubernetes 1.18, /metrics/resource/v1alpha1 renamed to /metrics/resource
     resourcePath: "/metrics/resource/v1alpha1"
     ## Metric relabellings to apply to samples before ingestion
@@ -1630,7 +1631,7 @@ prometheus:
     ##
     image:
       repository: quay.io/prometheus/prometheus
-      tag: v2.22.0
+      tag: v2.22.1
       sha: ""
 
     ## Tolerations for use with node taints

```

## 11.0.1

**Release date:** 2020-11-05

![AppVersion: 0.43.1](https://img.shields.io/static/v1?label=AppVersion&message=0.43.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Bump prometheus-operator to 0.43.1 (#305)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 7dd1e861..1aded8ff 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1351,7 +1351,7 @@ prometheusOperator:
   ##
   image:
     repository: quay.io/prometheus-operator/prometheus-operator
-    tag: v0.43.0
+    tag: v0.43.1
     sha: ""
     pullPolicy: IfNotPresent
 
@@ -1366,7 +1366,7 @@ prometheusOperator:
   ##
   prometheusConfigReloaderImage:
     repository: quay.io/prometheus-operator/prometheus-config-reloader
-    tag: v0.43.0
+    tag: v0.43.1
     sha: ""
 
   ## Set the prometheus config reloader side-car CPU limit

```

## 11.0.0

**Release date:** 2020-11-02

![AppVersion: 0.43.0](https://img.shields.io/static/v1?label=AppVersion&message=0.43.0&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Update prometheus-operator to v0.43.0 (#280)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index fe903812..7dd1e861 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1157,23 +1157,9 @@ prometheus-node-exporter:
 prometheusOperator:
   enabled: true
 
-  # If true prometheus operator will create and update its CRDs on startup
-  # Only for prometheusOperator.image.tag < v0.39.0
-  manageCrds: true
-
-  # This is deprecated and will be removed in v11 of the chart. Please use native TLS support below.
-  tlsProxy:
-    enabled: true
-    image:
-      repository: squareup/ghostunnel
-      tag: v1.5.2
-      sha: ""
-      pullPolicy: IfNotPresent
-    resources: {}
-
-  # Prometheus-Operator v0.39.0 and later support TLS natively. To preserve compatibility with tlsProxy, the secret tls-proxy-secret is reused
+  # Prometheus-Operator v0.39.0 and later support TLS natively.
   tls:
-    enabled: false
+    enabled: true
 
   ## Admission webhook support for PrometheusRules resources added in Prometheus Operator 0.30 can be enabled to prevent incorrectly formatted
   ## rules from making their way into prometheus and potentially preventing the container from starting
@@ -1188,7 +1174,7 @@ prometheusOperator:
       enabled: true
       image:
         repository: jettech/kube-webhook-certgen
-        tag: v1.2.1
+        tag: v1.5.0
         sha: ""
         pullPolicy: IfNotPresent
       resources: {}
@@ -1260,14 +1246,6 @@ prometheusOperator:
     ##
     externalIPs: []
 
-  ## Deploy CRDs used by Prometheus Operator.
-  ##
-  createCustomResource: true
-
-  ## Attempt to clean up CRDs created by Prometheus Operator.
-  ##
-  cleanupCustomResource: false
-
   ## Labels to add to the operator pod
   ##
   podLabels: {}
@@ -1373,7 +1351,7 @@ prometheusOperator:
   ##
   image:
     repository: quay.io/prometheus-operator/prometheus-operator
-    tag: v0.42.1
+    tag: v0.43.0
     sha: ""
     pullPolicy: IfNotPresent
 
@@ -1388,7 +1366,7 @@ prometheusOperator:
   ##
   prometheusConfigReloaderImage:
     repository: quay.io/prometheus-operator/prometheus-config-reloader
-    tag: v0.42.1
+    tag: v0.43.0
     sha: ""
 
   ## Set the prometheus config reloader side-car CPU limit
@@ -1652,7 +1630,7 @@ prometheus:
     ##
     image:
       repository: quay.io/prometheus/prometheus
-      tag: v2.21.0
+      tag: v2.22.0
       sha: ""
 
     ## Tolerations for use with node taints

```

## 10.3.5

**Release date:** 2020-11-02

![AppVersion: 0.42.1](https://img.shields.io/static/v1?label=AppVersion&message=0.42.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Update chart dependencies (#293)

### Default value changes

```diff
# No changes in this release
```

## 10.3.4

**Release date:** 2020-11-01

![AppVersion: 0.42.1](https://img.shields.io/static/v1?label=AppVersion&message=0.42.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Change helm stable repo in docs (#291)

### Default value changes

```diff
# No changes in this release
```

## 10.3.3

**Release date:** 2020-10-30

![AppVersion: 0.42.1](https://img.shields.io/static/v1?label=AppVersion&message=0.42.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack]: add additional alert labels on prometheusRules (#247)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 9a4397d8..fe903812 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -64,6 +64,9 @@ defaultRules:
   ## Annotations for default rules
   annotations: {}
 
+  ## Additional labels for PrometheusRule alerts
+  additionalRuleLabels: {}
+
 ## Deprecated way to provide custom recording or alerting rules to be deployed into the cluster.
 ##
 # additionalPrometheusRules: []

```

## 10.3.2

**Release date:** 2020-10-29

![AppVersion: 0.42.1](https://img.shields.io/static/v1?label=AppVersion&message=0.42.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Upgrade admission-webhook api (#267)

### Default value changes

```diff
# No changes in this release
```

## 10.3.1

**Release date:** 2020-10-28

![AppVersion: 0.42.1](https://img.shields.io/static/v1?label=AppVersion&message=0.42.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Support native TLS in prometheus-operator (#206)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 0530e486..9a4397d8 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1158,6 +1158,7 @@ prometheusOperator:
   # Only for prometheusOperator.image.tag < v0.39.0
   manageCrds: true
 
+  # This is deprecated and will be removed in v11 of the chart. Please use native TLS support below.
   tlsProxy:
     enabled: true
     image:
@@ -1167,6 +1168,10 @@ prometheusOperator:
       pullPolicy: IfNotPresent
     resources: {}
 
+  # Prometheus-Operator v0.39.0 and later support TLS natively. To preserve compatibility with tlsProxy, the secret tls-proxy-secret is reused
+  tls:
+    enabled: false
+
   ## Admission webhook support for PrometheusRules resources added in Prometheus Operator 0.30 can be enabled to prevent incorrectly formatted
   ## rules from making their way into prometheus and potentially preventing the container from starting
   admissionWebhooks:

```

## 10.2.0

**Release date:** 2020-10-27

![AppVersion: 0.42.1](https://img.shields.io/static/v1?label=AppVersion&message=0.42.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Allow use of Prometheus Operator master image tag (#224)

### Default value changes

```diff
# No changes in this release
```

## 10.1.4

**Release date:** 2020-10-27

![AppVersion: 0.42.1](https://img.shields.io/static/v1?label=AppVersion&message=0.42.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Fix requirements.yaml (#269)

### Default value changes

```diff
# No changes in this release
```

## 10.1.3

**Release date:** 2020-10-26

![AppVersion: 0.42.1](https://img.shields.io/static/v1?label=AppVersion&message=0.42.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Bump grafana to include PSP fix (#258)

### Default value changes

```diff
# No changes in this release
```

## 10.1.2

**Release date:** 2020-10-22

![AppVersion: 0.42.1](https://img.shields.io/static/v1?label=AppVersion&message=0.42.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Doc fixes (#246)
* Add support for ProbeSelector and ProbeNamespaceSelector (#237)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index c6729719..0530e486 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1788,6 +1788,26 @@ prometheus:
     ##
     podMonitorNamespaceSelector: {}
 
+    ## If true, a nil or {} value for prometheus.prometheusSpec.probeSelector will cause the
+    ## prometheus resource to be created with selectors based on values in the helm deployment,
+    ## which will also match the probes created
+    ##
+    probeSelectorNilUsesHelmValues: true
+
+    ## Probes to be selected for target discovery.
+    ## If {}, select all Probes
+    ##
+    probeSelector: {}
+    ## Example which selects Probes with label "prometheus" set to "somelabel"
+    # probeSelector:
+    #   matchLabels:
+    #     prometheus: somelabel
+
+    ## Namespaces to be selected for Probe discovery.
+    ## See https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#namespaceselector for usage
+    ##
+    probeNamespaceSelector: {}
+
     ## How long to retain metrics
     ##
     retention: 10d

```

## 10.1.1

**Release date:** 2020-10-21

![AppVersion: 0.42.1](https://img.shields.io/static/v1?label=AppVersion&message=0.42.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add hint for zero downtime migration from stable/prometheus-operator (#239)

### Default value changes

```diff
# No changes in this release
```

## 10.1.0

**Release date:** 2020-10-13

![AppVersion: 0.42.1](https://img.shields.io/static/v1?label=AppVersion&message=0.42.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add instance namespaces arguments for prometheus-operator deployment (#92)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index e7225038..c6729719 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1204,6 +1204,12 @@ prometheusOperator:
   ##
   denyNamespaces: []
 
+  ## Filter namespaces to look for prometheus-operator custom resources
+  ##
+  alertmanagerInstanceNamespaces: []
+  prometheusInstanceNamespaces: []
+  thanosInstanceNamespaces: []
+
   ## Service account for Alertmanager to use.
   ## ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
   ##

```

## 10.0.2

**Release date:** 2020-10-11

![AppVersion: 0.42.1](https://img.shields.io/static/v1?label=AppVersion&message=0.42.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Update grafana requirement (#197)

### Default value changes

```diff
# No changes in this release
```

## 10.0.1

**Release date:** 2020-10-09

![AppVersion: 0.42.1](https://img.shields.io/static/v1?label=AppVersion&message=0.42.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add upgrade path from stable/prometheus-operator (#119)

### Default value changes

```diff
# No changes in this release
```

## 10.0.0

**Release date:** 2020-10-09

![AppVersion: 0.42.1](https://img.shields.io/static/v1?label=AppVersion&message=0.42.1&color=success)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Update prometheus-operator to 0.42.1 (#128)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 932bf667..e7225038 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -499,6 +499,10 @@ alertmanager:
     ##
     portName: "web"
 
+    ## ClusterAdvertiseAddress is the explicit address to advertise in cluster. Needs to be provided for non RFC1918 [1] (public) addresses. [1] RFC1918: https://tools.ietf.org/html/rfc1918
+    ##
+    clusterAdvertiseAddress: false
+
 
 ## Using default values from https://github.com/grafana/helm-charts/blob/main/charts/grafana/values.yaml
 ##
@@ -1354,8 +1358,8 @@ prometheusOperator:
   ## Prometheus-operator image
   ##
   image:
-    repository: quay.io/coreos/prometheus-operator
-    tag: v0.38.1
+    repository: quay.io/prometheus-operator/prometheus-operator
+    tag: v0.42.1
     sha: ""
     pullPolicy: IfNotPresent
 
@@ -1363,14 +1367,14 @@ prometheusOperator:
   ##
   configmapReloadImage:
     repository: docker.io/jimmidyson/configmap-reload
-    tag: v0.3.0
+    tag: v0.4.0
     sha: ""
 
   ## Prometheus-config-reloader image to use for config and rule reloading
   ##
   prometheusConfigReloaderImage:
-    repository: quay.io/coreos/prometheus-config-reloader
-    tag: v0.38.1
+    repository: quay.io/prometheus-operator/prometheus-config-reloader
+    tag: v0.42.1
     sha: ""
 
   ## Set the prometheus config reloader side-car CPU limit
@@ -1634,7 +1638,7 @@ prometheus:
     ##
     image:
       repository: quay.io/prometheus/prometheus
-      tag: v2.18.2
+      tag: v2.21.0
       sha: ""
 
     ## Tolerations for use with node taints

```

## 9.4.10

**Release date:** 2020-10-07

![AppVersion: 0.38.1](https://img.shields.io/static/v1?label=AppVersion&message=0.38.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] fix broken link in comment (#141)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index bcfb0035..932bf667 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -161,7 +161,7 @@ alertmanager:
   ## they'll need to be properly escaped so that they are not interpreted by
   ## Helm
   ## ref: https://helm.sh/docs/developing_charts/#using-the-tpl-function
-  ##      https://prometheus.io/docs/alerting/configuration/#%3Ctmpl_string%3E
+  ##      https://prometheus.io/docs/alerting/configuration/#tmpl_string
   ##      https://prometheus.io/docs/alerting/notifications/
   ##      https://prometheus.io/docs/alerting/notification_examples/
   tplConfig: false

```

## 9.4.9

**Release date:** 2020-10-06

![AppVersion: 0.38.1](https://img.shields.io/static/v1?label=AppVersion&message=0.38.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Resync rules and dashboards (#168)

### Default value changes

```diff
# No changes in this release
```

## 9.4.8

**Release date:** 2020-10-05

![AppVersion: 0.38.1](https://img.shields.io/static/v1?label=AppVersion&message=0.38.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] style(grafana): fix typo (#150)

### Default value changes

```diff
# No changes in this release
```

## 9.4.7

**Release date:** 2020-10-05

![AppVersion: 0.38.1](https://img.shields.io/static/v1?label=AppVersion&message=0.38.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Replacing hyperkube image with kubectl image from bitnami (#160)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index e0ef831c..bcfb0035 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1385,11 +1385,11 @@ prometheusOperator:
   ##
   secretFieldSelector: ""
 
-  ## Hyperkube image to use when cleaning up
+  ## kubectl image to use when cleaning up
   ##
-  hyperkubeImage:
-    repository: k8s.gcr.io/hyperkube
-    tag: v1.16.12
+  kubectlImage:
+    repository: docker.io/bitnami/kubectl
+    tag: 1.16.15
     sha: ""
     pullPolicy: IfNotPresent
 

```

## 9.4.6

**Release date:** 2020-10-03

![AppVersion: 0.38.1](https://img.shields.io/static/v1?label=AppVersion&message=0.38.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Added more options for AlertManager ServiceMonitoring (#169)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 5202bac5..e0ef831c 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -315,6 +315,15 @@ alertmanager:
     interval: ""
     selfMonitor: true
 
+    ## scheme: HTTP scheme to use for scraping. Can be used with `tlsConfig` for example if using istio mTLS.
+    scheme: ""
+
+    ## tlsConfig: TLS configuration to use when scraping the endpoint. For example if using istio mTLS.
+    ## Of type: https://github.com/coreos/prometheus-operator/blob/master/Documentation/api.md#tlsconfig
+    tlsConfig: {}
+
+    bearerTokenFile:
+
     ## 	metric relabel configs to apply to samples before ingestion.
     ##
     metricRelabelings: []

```

## 9.4.5

**Release date:** 2020-09-30

![AppVersion: 0.38.1](https://img.shields.io/static/v1?label=AppVersion&message=0.38.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] add podmonitor docs (#164)

### Default value changes

```diff
# No changes in this release
```

## 9.4.4

**Release date:** 2020-09-21

![AppVersion: 0.38.1](https://img.shields.io/static/v1?label=AppVersion&message=0.38.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Add documentation on how to configure additionalPrometheusRulesMap (#109)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 12d782d4..5202bac5 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -64,9 +64,9 @@ defaultRules:
   ## Annotations for default rules
   annotations: {}
 
-## Provide custom recording or alerting rules to be deployed into the cluster.
+## Deprecated way to provide custom recording or alerting rules to be deployed into the cluster.
 ##
-additionalPrometheusRules: []
+# additionalPrometheusRules: []
 #  - name: my-rule-file
 #    groups:
 #      - name: my_group
@@ -74,6 +74,16 @@ additionalPrometheusRules: []
 #        - record: my_record
 #          expr: 100 * my_record
 
+## Provide custom recording or alerting rules to be deployed into the cluster.
+##
+additionalPrometheusRulesMap: {}
+#  rule-name:
+#    groups:
+#    - name: my_group
+#      rules:
+#      - record: my_record
+#        expr: 100 * my_record
+
 ##
 global:
   rbac:

```

## 9.4.3

**Release date:** 2020-09-16

![AppVersion: 0.38.1](https://img.shields.io/static/v1?label=AppVersion&message=0.38.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Fix URL to node-exporter (#104)

### Default value changes

```diff
# No changes in this release
```

## 9.4.2

**Release date:** 2020-09-14

![AppVersion: 0.38.1](https://img.shields.io/static/v1?label=AppVersion&message=0.38.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] Add Artifact Hub annotations for named links and operator search facet (#95)

### Default value changes

```diff
# No changes in this release
```

## 9.4.1

**Release date:** 2020-09-12

![AppVersion: 0.38.1](https://img.shields.io/static/v1?label=AppVersion&message=0.38.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Update README.md in kube-prometheus-stack hack (#89)

### Default value changes

```diff
# No changes in this release
```

## 9.4.0

**Release date:** 2020-09-10

![AppVersion: 0.38.1](https://img.shields.io/static/v1?label=AppVersion&message=0.38.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [kube-prometheus-stack] add secret-field-selector (#79)

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index b67195d0..12d782d4 100644
--- a/charts/kube-prometheus-stack/values.yaml
+++ b/charts/kube-prometheus-stack/values.yaml
@@ -1362,6 +1362,10 @@ prometheusOperator:
   ##
   configReloaderMemory: 25Mi
 
+  ## Set a Field Selector to filter watched secrets
+  ##
+  secretFieldSelector: ""
+
   ## Hyperkube image to use when cleaning up
   ##
   hyperkubeImage:

```

## 9.3.4

**Release date:** 2020-09-09

![AppVersion: 0.38.1](https://img.shields.io/static/v1?label=AppVersion&message=0.38.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-operator] Rename to kube-prometheus-stack (#1)

### Default value changes

```diff
# Default values for kube-prometheus-stack.
# This is a YAML-formatted file.
# Declare variables to be passed into your templates.

## Provide a name in place of kube-prometheus-stack for `app:` labels
##
nameOverride: ""

## Override the deployment namespace
##
namespaceOverride: ""

## Provide a k8s version to auto dashboard import script example: kubeTargetVersionOverride: 1.16.6
##
kubeTargetVersionOverride: ""

## Provide a name to substitute for the full names of resources
##
fullnameOverride: ""

## Labels to apply to all resources
##
commonLabels: {}
# scmhash: abc123
# myLabel: aakkmd

## Create default rules for monitoring the cluster
##
defaultRules:
  create: true
  rules:
    alertmanager: true
    etcd: true
    general: true
    k8s: true
    kubeApiserver: true
    kubeApiserverAvailability: true
    kubeApiserverError: true
    kubeApiserverSlos: true
    kubelet: true
    kubePrometheusGeneral: true
    kubePrometheusNodeAlerting: true
    kubePrometheusNodeRecording: true
    kubernetesAbsent: true
    kubernetesApps: true
    kubernetesResources: true
    kubernetesStorage: true
    kubernetesSystem: true
    kubeScheduler: true
    kubeStateMetrics: true
    network: true
    node: true
    prometheus: true
    prometheusOperator: true
    time: true

  ## Runbook url prefix for default rules
  runbookUrl: https://github.com/kubernetes-monitoring/kubernetes-mixin/tree/master/runbook.md#
  ## Reduce app namespace alert scope
  appNamespacesTarget: ".*"

  ## Labels for default rules
  labels: {}
  ## Annotations for default rules
  annotations: {}

## Provide custom recording or alerting rules to be deployed into the cluster.
##
additionalPrometheusRules: []
#  - name: my-rule-file
#    groups:
#      - name: my_group
#        rules:
#        - record: my_record
#          expr: 100 * my_record

##
global:
  rbac:
    create: true
    pspEnabled: true
    pspAnnotations: {}
      ## Specify pod annotations
      ## Ref: https://kubernetes.io/docs/concepts/policy/pod-security-policy/#apparmor
      ## Ref: https://kubernetes.io/docs/concepts/policy/pod-security-policy/#seccomp
      ## Ref: https://kubernetes.io/docs/concepts/policy/pod-security-policy/#sysctl
      ##
      # seccomp.security.alpha.kubernetes.io/allowedProfileNames: '*'
      # seccomp.security.alpha.kubernetes.io/defaultProfileName: 'docker/default'
      # apparmor.security.beta.kubernetes.io/defaultProfileName: 'runtime/default'

  ## Reference to one or more secrets to be used when pulling images
  ## ref: https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/
  ##
  imagePullSecrets: []
  # - name: "image-pull-secret"

## Configuration for alertmanager
## ref: https://prometheus.io/docs/alerting/alertmanager/
##
alertmanager:

  ## Deploy alertmanager
  ##
  enabled: true

  ## Api that prometheus will use to communicate with alertmanager. Possible values are v1, v2
  ##
  apiVersion: v2

  ## Service account for Alertmanager to use.
  ## ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
  ##
  serviceAccount:
    create: true
    name: ""
    annotations: {}

  ## Configure pod disruption budgets for Alertmanager
  ## ref: https://kubernetes.io/docs/tasks/run-application/configure-pdb/#specifying-a-poddisruptionbudget
  ## This configuration is immutable once created and will require the PDB to be deleted to be changed
  ## https://github.com/kubernetes/kubernetes/issues/45398
  ##
  podDisruptionBudget:
    enabled: false
    minAvailable: 1
    maxUnavailable: ""

  ## Alertmanager configuration directives
  ## ref: https://prometheus.io/docs/alerting/configuration/#configuration-file
  ##      https://prometheus.io/webtools/alerting/routing-tree-editor/
  ##
  config:
    global:
      resolve_timeout: 5m
    route:
      group_by: ['job']
      group_wait: 30s
      group_interval: 5m
      repeat_interval: 12h
      receiver: 'null'
      routes:
      - match:
          alertname: Watchdog
        receiver: 'null'
    receivers:
    - name: 'null'

  ## Pass the Alertmanager configuration directives through Helm's templating
  ## engine. If the Alertmanager configuration contains Alertmanager templates,
  ## they'll need to be properly escaped so that they are not interpreted by
  ## Helm
  ## ref: https://helm.sh/docs/developing_charts/#using-the-tpl-function
  ##      https://prometheus.io/docs/alerting/configuration/#%3Ctmpl_string%3E
  ##      https://prometheus.io/docs/alerting/notifications/
  ##      https://prometheus.io/docs/alerting/notification_examples/
  tplConfig: false

  ## Alertmanager template files to format alerts
  ## ref: https://prometheus.io/docs/alerting/notifications/
  ##      https://prometheus.io/docs/alerting/notification_examples/
  ##
  templateFiles: {}
  #
  ## An example template:
  #   template_1.tmpl: |-
  #       {{ define "cluster" }}{{ .ExternalURL | reReplaceAll ".*alertmanager\\.(.*)" "$1" }}{{ end }}
  #
  #       {{ define "slack.myorg.text" }}
  #       {{- $root := . -}}
  #       {{ range .Alerts }}
  #         *Alert:* {{ .Annotations.summary }} - `{{ .Labels.severity }}`
  #         *Cluster:*  {{ template "cluster" $root }}
  #         *Description:* {{ .Annotations.description }}
  #         *Graph:* <{{ .GeneratorURL }}|:chart_with_upwards_trend:>
  #         *Runbook:* <{{ .Annotations.runbook }}|:spiral_note_pad:>
  #         *Details:*
  #           {{ range .Labels.SortedPairs }} • *{{ .Name }}:* `{{ .Value }}`
  #           {{ end }}
  #       {{ end }}
  #       {{ end }}

  ingress:
    enabled: false

    annotations: {}

    labels: {}

    ## Hosts must be provided if Ingress is enabled.
    ##
    hosts: []
      # - alertmanager.domain.com

    ## Paths to use for ingress rules - one path should match the alertmanagerSpec.routePrefix
    ##
    paths: []
    # - /

    ## TLS configuration for Alertmanager Ingress
    ## Secret must be manually created in the namespace
    ##
    tls: []
    # - secretName: alertmanager-general-tls
    #   hosts:
    #   - alertmanager.example.com

  ## Configuration for Alertmanager secret
  ##
  secret:
    annotations: {}

  ## Configuration for creating an Ingress that will map to each Alertmanager replica service
  ## alertmanager.servicePerReplica must be enabled
  ##
  ingressPerReplica:
    enabled: false
    annotations: {}
    labels: {}

    ## Final form of the hostname for each per replica ingress is
    ## {{ ingressPerReplica.hostPrefix }}-{{ $replicaNumber }}.{{ ingressPerReplica.hostDomain }}
    ##
    ## Prefix for the per replica ingress that will have `-$replicaNumber`
    ## appended to the end
    hostPrefix: ""
    ## Domain that will be used for the per replica ingress
    hostDomain: ""

    ## Paths to use for ingress rules
    ##
    paths: []
    # - /

    ## Secret name containing the TLS certificate for alertmanager per replica ingress
    ## Secret must be manually created in the namespace
    tlsSecretName: ""

    ## Separated secret for each per replica Ingress. Can be used together with cert-manager
    ##
    tlsSecretPerReplica:
      enabled: false
      ## Final form of the secret for each per replica ingress is
      ## {{ tlsSecretPerReplica.prefix }}-{{ $replicaNumber }}
      ##
      prefix: "alertmanager"

  ## Configuration for Alertmanager service
  ##
  service:
    annotations: {}
    labels: {}
    clusterIP: ""

    ## Port for Alertmanager Service to listen on
    ##
    port: 9093
    ## To be used with a proxy extraContainer port
    ##
    targetPort: 9093
    ## Port to expose on each node
    ## Only used if service.type is 'NodePort'
    ##
    nodePort: 30903
    ## List of IP addresses at which the Prometheus server service is available
    ## Ref: https://kubernetes.io/docs/user-guide/services/#external-ips
    ##
    externalIPs: []
    loadBalancerIP: ""
    loadBalancerSourceRanges: []
    ## Service type
    ##
    type: ClusterIP

  ## Configuration for creating a separate Service for each statefulset Alertmanager replica
  ##
  servicePerReplica:
    enabled: false
    annotations: {}

    ## Port for Alertmanager Service per replica to listen on
    ##
    port: 9093

    ## To be used with a proxy extraContainer port
    targetPort: 9093

    ## Port to expose on each node
    ## Only used if servicePerReplica.type is 'NodePort'
    ##
    nodePort: 30904

    ## Loadbalancer source IP ranges
    ## Only used if servicePerReplica.type is "loadbalancer"
    loadBalancerSourceRanges: []
    ## Service type
    ##
    type: ClusterIP

  ## If true, create a serviceMonitor for alertmanager
  ##
  serviceMonitor:
    ## Scrape interval. If not set, the Prometheus default scrape interval is used.
    ##
    interval: ""
    selfMonitor: true

    ## 	metric relabel configs to apply to samples before ingestion.
    ##
    metricRelabelings: []
    # - action: keep
    #   regex: 'kube_(daemonset|deployment|pod|namespace|node|statefulset).+'
    #   sourceLabels: [__name__]

    # 	relabel configs to apply to samples before ingestion.
    ##
    relabelings: []
    # - sourceLabels: [__meta_kubernetes_pod_node_name]
    #   separator: ;
    #   regex: ^(.*)$
    #   targetLabel: nodename
    #   replacement: $1
    #   action: replace

  ## Settings affecting alertmanagerSpec
  ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#alertmanagerspec
  ##
  alertmanagerSpec:
    ## Standard object’s metadata. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata
    ## Metadata Labels and Annotations gets propagated to the Alertmanager pods.
    ##
    podMetadata: {}

    ## Image of Alertmanager
    ##
    image:
      repository: quay.io/prometheus/alertmanager
      tag: v0.21.0
      sha: ""

    ## If true then the user will be responsible to provide a secret with alertmanager configuration
    ## So when true the config part will be ignored (including templateFiles) and the one in the secret will be used
    ##
    useExistingSecret: false

    ## Secrets is a list of Secrets in the same namespace as the Alertmanager object, which shall be mounted into the
    ## Alertmanager Pods. The Secrets are mounted into /etc/alertmanager/secrets/.
    ##
    secrets: []

    ## ConfigMaps is a list of ConfigMaps in the same namespace as the Alertmanager object, which shall be mounted into the Alertmanager Pods.
    ## The ConfigMaps are mounted into /etc/alertmanager/configmaps/.
    ##
    configMaps: []

    ## ConfigSecret is the name of a Kubernetes Secret in the same namespace as the Alertmanager object, which contains configuration for
    ## this Alertmanager instance. Defaults to 'alertmanager-' The secret is mounted into /etc/alertmanager/config.
    ##
    # configSecret:

    ## Define Log Format
    # Use logfmt (default) or json-formatted logging
    logFormat: logfmt

    ## Log level for Alertmanager to be configured with.
    ##
    logLevel: info

    ## Size is the expected size of the alertmanager cluster. The controller will eventually make the size of the
    ## running cluster equal to the expected size.
    replicas: 1

    ## Time duration Alertmanager shall retain data for. Default is '120h', and must match the regular expression
    ## [0-9]+(ms|s|m|h) (milliseconds seconds minutes hours).
    ##
    retention: 120h

    ## Storage is the definition of how storage will be used by the Alertmanager instances.
    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/user-guides/storage.md
    ##
    storage: {}
    # volumeClaimTemplate:
    #   spec:
    #     storageClassName: gluster
    #     accessModes: ["ReadWriteOnce"]
    #     resources:
    #       requests:
    #         storage: 50Gi
    #   selector: {}


    ## 	The external URL the Alertmanager instances will be available under. This is necessary to generate correct URLs. This is necessary if Alertmanager is not served from root of a DNS name.	string	false
    ##
    externalUrl:

    ## 	The route prefix Alertmanager registers HTTP handlers for. This is useful, if using ExternalURL and a proxy is rewriting HTTP routes of a request, and the actual ExternalURL is still true,
    ## but the server serves requests under a different route prefix. For example for use with kubectl proxy.
    ##
    routePrefix: /

    ## If set to true all actions on the underlying managed objects are not going to be performed, except for delete actions.
    ##
    paused: false

    ## Define which Nodes the Pods are scheduled on.
    ## ref: https://kubernetes.io/docs/user-guide/node-selection/
    ##
    nodeSelector: {}

    ## Define resources requests and limits for single Pods.
    ## ref: https://kubernetes.io/docs/user-guide/compute-resources/
    ##
    resources: {}
    # requests:
    #   memory: 400Mi

    ## Pod anti-affinity can prevent the scheduler from placing Prometheus replicas on the same node.
    ## The default value "soft" means that the scheduler should *prefer* to not schedule two replica pods onto the same node but no guarantee is provided.
    ## The value "hard" means that the scheduler is *required* to not schedule two replica pods onto the same node.
    ## The value "" will disable pod anti-affinity so that no anti-affinity rules will be configured.
    ##
    podAntiAffinity: ""

    ## If anti-affinity is enabled sets the topologyKey to use for anti-affinity.
    ## This can be changed to, for example, failure-domain.beta.kubernetes.io/zone
    ##
    podAntiAffinityTopologyKey: kubernetes.io/hostname

    ## Assign custom affinity rules to the alertmanager instance
    ## ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
    ##
    affinity: {}
    # nodeAffinity:
    #   requiredDuringSchedulingIgnoredDuringExecution:
    #     nodeSelectorTerms:
    #     - matchExpressions:
    #       - key: kubernetes.io/e2e-az-name
    #         operator: In
    #         values:
    #         - e2e-az1
    #         - e2e-az2

    ## If specified, the pod's tolerations.
    ## ref: https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/
    ##
    tolerations: []
    # - key: "key"
    #   operator: "Equal"
    #   value: "value"
    #   effect: "NoSchedule"

    ## SecurityContext holds pod-level security attributes and common container settings.
    ## This defaults to non root user with uid 1000 and gid 2000.	*v1.PodSecurityContext	false
    ## ref: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
    ##
    securityContext:
      runAsGroup: 2000
      runAsNonRoot: true
      runAsUser: 1000
      fsGroup: 2000

    ## ListenLocal makes the Alertmanager server listen on loopback, so that it does not bind against the Pod IP.
    ## Note this is only for the Alertmanager UI, not the gossip communication.
    ##
    listenLocal: false

    ## Containers allows injecting additional containers. This is meant to allow adding an authentication proxy to an Alertmanager pod.
    ##
    containers: []

    ## Priority class assigned to the Pods
    ##
    priorityClassName: ""

    ## AdditionalPeers allows injecting a set of additional Alertmanagers to peer with to form a highly available cluster.
    ##
    additionalPeers: []

    ## PortName to use for Alert Manager.
    ##
    portName: "web"


## Using default values from https://github.com/grafana/helm-charts/blob/main/charts/grafana/values.yaml
##
grafana:
  enabled: true
  namespaceOverride: ""

  ## Deploy default dashboards.
  ##
  defaultDashboardsEnabled: true

  adminPassword: prom-operator

  ingress:
    ## If true, Grafana Ingress will be created
    ##
    enabled: false

    ## Annotations for Grafana Ingress
    ##
    annotations: {}
      # kubernetes.io/ingress.class: nginx
      # kubernetes.io/tls-acme: "true"

    ## Labels to be added to the Ingress
    ##
    labels: {}

    ## Hostnames.
    ## Must be provided if Ingress is enable.
    ##
    # hosts:
    #   - grafana.domain.com
    hosts: []

    ## Path for grafana ingress
    path: /

    ## TLS configuration for grafana Ingress
    ## Secret must be manually created in the namespace
    ##
    tls: []
    # - secretName: grafana-general-tls
    #   hosts:
    #   - grafana.example.com

  sidecar:
    dashboards:
      enabled: true
      label: grafana_dashboard

      ## Annotations for Grafana dashboard configmaps
      ##
      annotations: {}
    datasources:
      enabled: true
      defaultDatasourceEnabled: true

      ## Annotations for Grafana datasource configmaps
      ##
      annotations: {}

      ## Create datasource for each Pod of Prometheus StatefulSet;
      ## this uses headless service `prometheus-operated` which is
      ## created by Prometheus Operator
      ## ref: https://git.io/fjaBS
      createPrometheusReplicasDatasources: false
      label: grafana_datasource

  extraConfigmapMounts: []
  # - name: certs-configmap
  #   mountPath: /etc/grafana/ssl/
  #   configMap: certs-configmap
  #   readOnly: true

  ## Configure additional grafana datasources (passed through tpl)
  ## ref: http://docs.grafana.org/administration/provisioning/#datasources
  additionalDataSources: []
  # - name: prometheus-sample
  #   access: proxy
  #   basicAuth: true
  #   basicAuthPassword: pass
  #   basicAuthUser: daco
  #   editable: false
  #   jsonData:
  #       tlsSkipVerify: true
  #   orgId: 1
  #   type: prometheus
  #   url: https://{{ printf "%s-prometheus.svc" .Release.Name }}:9090
  #   version: 1

  ## Passed to grafana subchart and used by servicemonitor below
  ##
  service:
    portName: service

  ## If true, create a serviceMonitor for grafana
  ##
  serviceMonitor:
    ## Scrape interval. If not set, the Prometheus default scrape interval is used.
    ##
    interval: ""
    selfMonitor: true

    ## 	metric relabel configs to apply to samples before ingestion.
    ##
    metricRelabelings: []
    # - action: keep
    #   regex: 'kube_(daemonset|deployment|pod|namespace|node|statefulset).+'
    #   sourceLabels: [__name__]

    # 	relabel configs to apply to samples before ingestion.
    ##
    relabelings: []
    # - sourceLabels: [__meta_kubernetes_pod_node_name]
    #   separator: ;
    #   regex: ^(.*)$
    #   targetLabel: nodename
    #   replacement: $1
    #   action: replace

## Component scraping the kube api server
##
kubeApiServer:
  enabled: true
  tlsConfig:
    serverName: kubernetes
    insecureSkipVerify: false

  ## If your API endpoint address is not reachable (as in AKS) you can replace it with the kubernetes service
  ##
  relabelings: []
  # - sourceLabels:
  #     - __meta_kubernetes_namespace
  #     - __meta_kubernetes_service_name
  #     - __meta_kubernetes_endpoint_port_name
  #   action: keep
  #   regex: default;kubernetes;https
  # - targetLabel: __address__
  #   replacement: kubernetes.default.svc:443

  serviceMonitor:
    ## Scrape interval. If not set, the Prometheus default scrape interval is used.
    ##
    interval: ""
    jobLabel: component
    selector:
      matchLabels:
        component: apiserver
        provider: kubernetes

    ## 	metric relabel configs to apply to samples before ingestion.
    ##
    metricRelabelings: []
    # - action: keep
    #   regex: 'kube_(daemonset|deployment|pod|namespace|node|statefulset).+'
    #   sourceLabels: [__name__]

## Component scraping the kubelet and kubelet-hosted cAdvisor
##
kubelet:
  enabled: true
  namespace: kube-system

  serviceMonitor:
    ## Scrape interval. If not set, the Prometheus default scrape interval is used.
    ##
    interval: ""

    ## Enable scraping the kubelet over https. For requirements to enable this see
    ## https://github.com/prometheus-operator/prometheus-operator/issues/926
    ##
    https: true

    ## Enable scraping /metrics/cadvisor from kubelet's service
    ##
    cAdvisor: true

    ## Enable scraping /metrics/probes from kubelet's service
    ##
    probes: true

    ## Enable scraping /metrics/resource from kubelet's service
    ##
    resource: true
    # From kubernetes 1.18, /metrics/resource/v1alpha1 renamed to /metrics/resource
    resourcePath: "/metrics/resource/v1alpha1"
    ## Metric relabellings to apply to samples before ingestion
    ##
    cAdvisorMetricRelabelings: []
    # - sourceLabels: [__name__, image]
    #   separator: ;
    #   regex: container_([a-z_]+);
    #   replacement: $1
    #   action: drop
    # - sourceLabels: [__name__]
    #   separator: ;
    #   regex: container_(network_tcp_usage_total|network_udp_usage_total|tasks_state|cpu_load_average_10s)
    #   replacement: $1
    #   action: drop

    ## Metric relabellings to apply to samples before ingestion
    ##
    probesMetricRelabelings: []
    # - sourceLabels: [__name__, image]
    #   separator: ;
    #   regex: container_([a-z_]+);
    #   replacement: $1
    #   action: drop
    # - sourceLabels: [__name__]
    #   separator: ;
    #   regex: container_(network_tcp_usage_total|network_udp_usage_total|tasks_state|cpu_load_average_10s)
    #   replacement: $1
    #   action: drop

    # 	relabel configs to apply to samples before ingestion.
    #   metrics_path is required to match upstream rules and charts
    ##
    cAdvisorRelabelings:
      - sourceLabels: [__metrics_path__]
        targetLabel: metrics_path
    # - sourceLabels: [__meta_kubernetes_pod_node_name]
    #   separator: ;
    #   regex: ^(.*)$
    #   targetLabel: nodename
    #   replacement: $1
    #   action: replace

    probesRelabelings:
      - sourceLabels: [__metrics_path__]
        targetLabel: metrics_path
    # - sourceLabels: [__meta_kubernetes_pod_node_name]
    #   separator: ;
    #   regex: ^(.*)$
    #   targetLabel: nodename
    #   replacement: $1
    #   action: replace

    resourceRelabelings:
      - sourceLabels: [__metrics_path__]
        targetLabel: metrics_path
    # - sourceLabels: [__meta_kubernetes_pod_node_name]
    #   separator: ;
    #   regex: ^(.*)$
    #   targetLabel: nodename
    #   replacement: $1
    #   action: replace

    metricRelabelings: []
    # - sourceLabels: [__name__, image]
    #   separator: ;
    #   regex: container_([a-z_]+);
    #   replacement: $1
    #   action: drop
    # - sourceLabels: [__name__]
    #   separator: ;
    #   regex: container_(network_tcp_usage_total|network_udp_usage_total|tasks_state|cpu_load_average_10s)
    #   replacement: $1
    #   action: drop

    # 	relabel configs to apply to samples before ingestion.
    #   metrics_path is required to match upstream rules and charts
    ##
    relabelings:
      - sourceLabels: [__metrics_path__]
        targetLabel: metrics_path
    # - sourceLabels: [__meta_kubernetes_pod_node_name]
    #   separator: ;
    #   regex: ^(.*)$
    #   targetLabel: nodename
    #   replacement: $1
    #   action: replace

## Component scraping the kube controller manager
##
kubeControllerManager:
  enabled: true

  ## If your kube controller manager is not deployed as a pod, specify IPs it can be found on
  ##
  endpoints: []
  # - 10.141.4.22
  # - 10.141.4.23
  # - 10.141.4.24

  ## If using kubeControllerManager.endpoints only the port and targetPort are used
  ##
  service:
    port: 10252
    targetPort: 10252
    # selector:
    #   component: kube-controller-manager

  serviceMonitor:
    ## Scrape interval. If not set, the Prometheus default scrape interval is used.
    ##
    interval: ""

    ## Enable scraping kube-controller-manager over https.
    ## Requires proper certs (not self-signed) and delegated authentication/authorization checks
    ##
    https: false

    # Skip TLS certificate validation when scraping
    insecureSkipVerify: null

    # Name of the server to use when validating TLS certificate
    serverName: null

    ## 	metric relabel configs to apply to samples before ingestion.
    ##
    metricRelabelings: []
    # - action: keep
    #   regex: 'kube_(daemonset|deployment|pod|namespace|node|statefulset).+'
    #   sourceLabels: [__name__]

    # 	relabel configs to apply to samples before ingestion.
    ##
    relabelings: []
    # - sourceLabels: [__meta_kubernetes_pod_node_name]
    #   separator: ;
    #   regex: ^(.*)$
    #   targetLabel: nodename
    #   replacement: $1
    #   action: replace

## Component scraping coreDns. Use either this or kubeDns
##
coreDns:
  enabled: true
  service:
    port: 9153
    targetPort: 9153
    # selector:
    #   k8s-app: kube-dns
  serviceMonitor:
    ## Scrape interval. If not set, the Prometheus default scrape interval is used.
    ##
    interval: ""

    ## 	metric relabel configs to apply to samples before ingestion.
    ##
    metricRelabelings: []
    # - action: keep
    #   regex: 'kube_(daemonset|deployment|pod|namespace|node|statefulset).+'
    #   sourceLabels: [__name__]

    # 	relabel configs to apply to samples before ingestion.
    ##
    relabelings: []
    # - sourceLabels: [__meta_kubernetes_pod_node_name]
    #   separator: ;
    #   regex: ^(.*)$
    #   targetLabel: nodename
    #   replacement: $1
    #   action: replace

## Component scraping kubeDns. Use either this or coreDns
##
kubeDns:
  enabled: false
  service:
    dnsmasq:
      port: 10054
      targetPort: 10054
    skydns:
      port: 10055
      targetPort: 10055
    # selector:
    #   k8s-app: kube-dns
  serviceMonitor:
    ## Scrape interval. If not set, the Prometheus default scrape interval is used.
    ##
    interval: ""

    ## 	metric relabel configs to apply to samples before ingestion.
    ##
    metricRelabelings: []
    # - action: keep
    #   regex: 'kube_(daemonset|deployment|pod|namespace|node|statefulset).+'
    #   sourceLabels: [__name__]

    # 	relabel configs to apply to samples before ingestion.
    ##
    relabelings: []
    # - sourceLabels: [__meta_kubernetes_pod_node_name]
    #   separator: ;
    #   regex: ^(.*)$
    #   targetLabel: nodename
    #   replacement: $1
    #   action: replace
    dnsmasqMetricRelabelings: []
    # - action: keep
    #   regex: 'kube_(daemonset|deployment|pod|namespace|node|statefulset).+'
    #   sourceLabels: [__name__]

    # 	relabel configs to apply to samples before ingestion.
    ##
    dnsmasqRelabelings: []
    # - sourceLabels: [__meta_kubernetes_pod_node_name]
    #   separator: ;
    #   regex: ^(.*)$
    #   targetLabel: nodename
    #   replacement: $1
    #   action: replace

## Component scraping etcd
##
kubeEtcd:
  enabled: true

  ## If your etcd is not deployed as a pod, specify IPs it can be found on
  ##
  endpoints: []
  # - 10.141.4.22
  # - 10.141.4.23
  # - 10.141.4.24

  ## Etcd service. If using kubeEtcd.endpoints only the port and targetPort are used
  ##
  service:
    port: 2379
    targetPort: 2379
    # selector:
    #   component: etcd

  ## Configure secure access to the etcd cluster by loading a secret into prometheus and
  ## specifying security configuration below. For example, with a secret named etcd-client-cert
  ##
  ## serviceMonitor:
  ##   scheme: https
  ##   insecureSkipVerify: false
  ##   serverName: localhost
  ##   caFile: /etc/prometheus/secrets/etcd-client-cert/etcd-ca
  ##   certFile: /etc/prometheus/secrets/etcd-client-cert/etcd-client
  ##   keyFile: /etc/prometheus/secrets/etcd-client-cert/etcd-client-key
  ##
  serviceMonitor:
    ## Scrape interval. If not set, the Prometheus default scrape interval is used.
    ##
    interval: ""
    scheme: http
    insecureSkipVerify: false
    serverName: ""
    caFile: ""
    certFile: ""
    keyFile: ""

    ## 	metric relabel configs to apply to samples before ingestion.
    ##
    metricRelabelings: []
    # - action: keep
    #   regex: 'kube_(daemonset|deployment|pod|namespace|node|statefulset).+'
    #   sourceLabels: [__name__]

    # 	relabel configs to apply to samples before ingestion.
    ##
    relabelings: []
    # - sourceLabels: [__meta_kubernetes_pod_node_name]
    #   separator: ;
    #   regex: ^(.*)$
    #   targetLabel: nodename
    #   replacement: $1
    #   action: replace


## Component scraping kube scheduler
##
kubeScheduler:
  enabled: true

  ## If your kube scheduler is not deployed as a pod, specify IPs it can be found on
  ##
  endpoints: []
  # - 10.141.4.22
  # - 10.141.4.23
  # - 10.141.4.24

  ## If using kubeScheduler.endpoints only the port and targetPort are used
  ##
  service:
    port: 10251
    targetPort: 10251
    # selector:
    #   component: kube-scheduler

  serviceMonitor:
    ## Scrape interval. If not set, the Prometheus default scrape interval is used.
    ##
    interval: ""
    ## Enable scraping kube-scheduler over https.
    ## Requires proper certs (not self-signed) and delegated authentication/authorization checks
    ##
    https: false

    ## Skip TLS certificate validation when scraping
    insecureSkipVerify: null

    ## Name of the server to use when validating TLS certificate
    serverName: null

    ## 	metric relabel configs to apply to samples before ingestion.
    ##
    metricRelabelings: []
    # - action: keep
    #   regex: 'kube_(daemonset|deployment|pod|namespace|node|statefulset).+'
    #   sourceLabels: [__name__]

    # 	relabel configs to apply to samples before ingestion.
    ##
    relabelings: []
    # - sourceLabels: [__meta_kubernetes_pod_node_name]
    #   separator: ;
    #   regex: ^(.*)$
    #   targetLabel: nodename
    #   replacement: $1
    #   action: replace


## Component scraping kube proxy
##
kubeProxy:
  enabled: true

  ## If your kube proxy is not deployed as a pod, specify IPs it can be found on
  ##
  endpoints: []
  # - 10.141.4.22
  # - 10.141.4.23
  # - 10.141.4.24

  service:
    port: 10249
    targetPort: 10249
    # selector:
    #   k8s-app: kube-proxy

  serviceMonitor:
    ## Scrape interval. If not set, the Prometheus default scrape interval is used.
    ##
    interval: ""

    ## Enable scraping kube-proxy over https.
    ## Requires proper certs (not self-signed) and delegated authentication/authorization checks
    ##
    https: false

    ## 	metric relabel configs to apply to samples before ingestion.
    ##
    metricRelabelings: []
    # - action: keep
    #   regex: 'kube_(daemonset|deployment|pod|namespace|node|statefulset).+'
    #   sourceLabels: [__name__]

    # 	relabel configs to apply to samples before ingestion.
    ##
    relabelings: []
    # - action: keep
    #   regex: 'kube_(daemonset|deployment|pod|namespace|node|statefulset).+'
    #   sourceLabels: [__name__]


## Component scraping kube state metrics
##
kubeStateMetrics:
  enabled: true
  serviceMonitor:
    ## Scrape interval. If not set, the Prometheus default scrape interval is used.
    ##
    interval: ""

    ## 	metric relabel configs to apply to samples before ingestion.
    ##
    metricRelabelings: []
    # - action: keep
    #   regex: 'kube_(daemonset|deployment|pod|namespace|node|statefulset).+'
    #   sourceLabels: [__name__]

    # 	relabel configs to apply to samples before ingestion.
    ##
    relabelings: []
    # - sourceLabels: [__meta_kubernetes_pod_node_name]
    #   separator: ;
    #   regex: ^(.*)$
    #   targetLabel: nodename
    #   replacement: $1
    #   action: replace

## Configuration for kube-state-metrics subchart
##
kube-state-metrics:
  namespaceOverride: ""
  rbac:
    create: true
  podSecurityPolicy:
    enabled: true

## Deploy node exporter as a daemonset to all nodes
##
nodeExporter:
  enabled: true

  ## Use the value configured in prometheus-node-exporter.podLabels
  ##
  jobLabel: jobLabel

  serviceMonitor:
    ## Scrape interval. If not set, the Prometheus default scrape interval is used.
    ##
    interval: ""

    ## How long until a scrape request times out. If not set, the Prometheus default scape timeout is used.
    ##
    scrapeTimeout: ""

    ## 	metric relabel configs to apply to samples before ingestion.
    ##
    metricRelabelings: []
    # - sourceLabels: [__name__]
    #   separator: ;
    #   regex: ^node_mountstats_nfs_(event|operations|transport)_.+
    #   replacement: $1
    #   action: drop

    ## 	relabel configs to apply to samples before ingestion.
    ##
    relabelings: []
    # - sourceLabels: [__meta_kubernetes_pod_node_name]
    #   separator: ;
    #   regex: ^(.*)$
    #   targetLabel: nodename
    #   replacement: $1
    #   action: replace

## Configuration for prometheus-node-exporter subchart
##
prometheus-node-exporter:
  namespaceOverride: ""
  podLabels:
    ## Add the 'node-exporter' label to be used by serviceMonitor to match standard common usage in rules and grafana dashboards
    ##
    jobLabel: node-exporter
  extraArgs:
    - --collector.filesystem.ignored-mount-points=^/(dev|proc|sys|var/lib/docker/.+)($|/)
    - --collector.filesystem.ignored-fs-types=^(autofs|binfmt_misc|cgroup|configfs|debugfs|devpts|devtmpfs|fusectl|hugetlbfs|mqueue|overlay|proc|procfs|pstore|rpc_pipefs|securityfs|sysfs|tracefs)$

## Manages Prometheus and Alertmanager components
##
prometheusOperator:
  enabled: true

  # If true prometheus operator will create and update its CRDs on startup
  # Only for prometheusOperator.image.tag < v0.39.0
  manageCrds: true

  tlsProxy:
    enabled: true
    image:
      repository: squareup/ghostunnel
      tag: v1.5.2
      sha: ""
      pullPolicy: IfNotPresent
    resources: {}

  ## Admission webhook support for PrometheusRules resources added in Prometheus Operator 0.30 can be enabled to prevent incorrectly formatted
  ## rules from making their way into prometheus and potentially preventing the container from starting
  admissionWebhooks:
    failurePolicy: Fail
    enabled: true
    ## If enabled, generate a self-signed certificate, then patch the webhook configurations with the generated data.
    ## On chart upgrades (or if the secret exists) the cert will not be re-generated. You can use this to provide your own
    ## certs ahead of time if you wish.
    ##
    patch:
      enabled: true
      image:
        repository: jettech/kube-webhook-certgen
        tag: v1.2.1
        sha: ""
        pullPolicy: IfNotPresent
      resources: {}
      ## Provide a priority class name to the webhook patching job
      ##
      priorityClassName: ""
      podAnnotations: {}
      nodeSelector: {}
      affinity: {}
      tolerations: []

  ## Namespaces to scope the interaction of the Prometheus Operator and the apiserver (allow list).
  ## This is mutually exclusive with denyNamespaces. Setting this to an empty object will disable the configuration
  ##
  namespaces: {}
    # releaseNamespace: true
    # additional:
    # - kube-system

  ## Namespaces not to scope the interaction of the Prometheus Operator (deny list).
  ##
  denyNamespaces: []

  ## Service account for Alertmanager to use.
  ## ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
  ##
  serviceAccount:
    create: true
    name: ""

  ## Configuration for Prometheus operator service
  ##
  service:
    annotations: {}
    labels: {}
    clusterIP: ""

  ## Port to expose on each node
  ## Only used if service.type is 'NodePort'
  ##
    nodePort: 30080

    nodePortTls: 30443

  ## Additional ports to open for Prometheus service
  ## ref: https://kubernetes.io/docs/concepts/services-networking/service/#multi-port-services
  ##
    additionalPorts: []

  ## Loadbalancer IP
  ## Only use if service.type is "loadbalancer"
  ##
    loadBalancerIP: ""
    loadBalancerSourceRanges: []

  ## Service type
  ## NodePort, ClusterIP, loadbalancer
  ##
    type: ClusterIP

    ## List of IP addresses at which the Prometheus server service is available
    ## Ref: https://kubernetes.io/docs/user-guide/services/#external-ips
    ##
    externalIPs: []

  ## Deploy CRDs used by Prometheus Operator.
  ##
  createCustomResource: true

  ## Attempt to clean up CRDs created by Prometheus Operator.
  ##
  cleanupCustomResource: false

  ## Labels to add to the operator pod
  ##
  podLabels: {}

  ## Annotations to add to the operator pod
  ##
  podAnnotations: {}

  ## Assign a PriorityClassName to pods if set
  # priorityClassName: ""

  ## Define Log Format
  # Use logfmt (default) or json-formatted logging
  # logFormat: logfmt

  ## Decrease log verbosity to errors only
  # logLevel: error

  ## If true, the operator will create and maintain a service for scraping kubelets
  ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/helm/prometheus-operator/README.md
  ##
  kubeletService:
    enabled: true
    namespace: kube-system

  ## Create a servicemonitor for the operator
  ##
  serviceMonitor:
    ## Scrape interval. If not set, the Prometheus default scrape interval is used.
    ##
    interval: ""
    ## Scrape timeout. If not set, the Prometheus default scrape timeout is used.
    scrapeTimeout: ""
    selfMonitor: true

    ## 	metric relabel configs to apply to samples before ingestion.
    ##
    metricRelabelings: []
    # - action: keep
    #   regex: 'kube_(daemonset|deployment|pod|namespace|node|statefulset).+'
    #   sourceLabels: [__name__]

    # 	relabel configs to apply to samples before ingestion.
    ##
    relabelings: []
    # - sourceLabels: [__meta_kubernetes_pod_node_name]
    #   separator: ;
    #   regex: ^(.*)$
    #   targetLabel: nodename
    #   replacement: $1
    #   action: replace

  ## Resource limits & requests
  ##
  resources: {}
  # limits:
  #   cpu: 200m
  #   memory: 200Mi
  # requests:
  #   cpu: 100m
  #   memory: 100Mi

  # Required for use in managed kubernetes clusters (such as AWS EKS) with custom CNI (such as calico),
  # because control-plane managed by AWS cannot communicate with pods' IP CIDR and admission webhooks are not working
  ##
  hostNetwork: false

  ## Define which Nodes the Pods are scheduled on.
  ## ref: https://kubernetes.io/docs/user-guide/node-selection/
  ##
  nodeSelector: {}

  ## Tolerations for use with node taints
  ## ref: https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/
  ##
  tolerations: []
  # - key: "key"
  #   operator: "Equal"
  #   value: "value"
  #   effect: "NoSchedule"

  ## Assign custom affinity rules to the prometheus operator
  ## ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
  ##
  affinity: {}
    # nodeAffinity:
    #   requiredDuringSchedulingIgnoredDuringExecution:
    #     nodeSelectorTerms:
    #     - matchExpressions:
    #       - key: kubernetes.io/e2e-az-name
    #         operator: In
    #         values:
    #         - e2e-az1
    #         - e2e-az2

  securityContext:
    fsGroup: 65534
    runAsGroup: 65534
    runAsNonRoot: true
    runAsUser: 65534

  ## Prometheus-operator image
  ##
  image:
    repository: quay.io/coreos/prometheus-operator
    tag: v0.38.1
    sha: ""
    pullPolicy: IfNotPresent

  ## Configmap-reload image to use for reloading configmaps
  ##
  configmapReloadImage:
    repository: docker.io/jimmidyson/configmap-reload
    tag: v0.3.0
    sha: ""

  ## Prometheus-config-reloader image to use for config and rule reloading
  ##
  prometheusConfigReloaderImage:
    repository: quay.io/coreos/prometheus-config-reloader
    tag: v0.38.1
    sha: ""

  ## Set the prometheus config reloader side-car CPU limit
  ##
  configReloaderCpu: 100m

  ## Set the prometheus config reloader side-car memory limit
  ##
  configReloaderMemory: 25Mi

  ## Hyperkube image to use when cleaning up
  ##
  hyperkubeImage:
    repository: k8s.gcr.io/hyperkube
    tag: v1.16.12
    sha: ""
    pullPolicy: IfNotPresent

## Deploy a Prometheus instance
##
prometheus:

  enabled: true

  ## Annotations for Prometheus
  ##
  annotations: {}

  ## Service account for Prometheuses to use.
  ## ref: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
  ##
  serviceAccount:
    create: true
    name: ""

  ## Configuration for Prometheus service
  ##
  service:
    annotations: {}
    labels: {}
    clusterIP: ""

    ## Port for Prometheus Service to listen on
    ##
    port: 9090

    ## To be used with a proxy extraContainer port
    targetPort: 9090

    ## List of IP addresses at which the Prometheus server service is available
    ## Ref: https://kubernetes.io/docs/user-guide/services/#external-ips
    ##
    externalIPs: []

    ## Port to expose on each node
    ## Only used if service.type is 'NodePort'
    ##
    nodePort: 30090

    ## Loadbalancer IP
    ## Only use if service.type is "loadbalancer"
    loadBalancerIP: ""
    loadBalancerSourceRanges: []
    ## Service type
    ##
    type: ClusterIP

    sessionAffinity: ""

  ## Configuration for creating a separate Service for each statefulset Prometheus replica
  ##
  servicePerReplica:
    enabled: false
    annotations: {}

    ## Port for Prometheus Service per replica to listen on
    ##
    port: 9090

    ## To be used with a proxy extraContainer port
    targetPort: 9090

    ## Port to expose on each node
    ## Only used if servicePerReplica.type is 'NodePort'
    ##
    nodePort: 30091

    ## Loadbalancer source IP ranges
    ## Only used if servicePerReplica.type is "loadbalancer"
    loadBalancerSourceRanges: []
    ## Service type
    ##
    type: ClusterIP

  ## Configure pod disruption budgets for Prometheus
  ## ref: https://kubernetes.io/docs/tasks/run-application/configure-pdb/#specifying-a-poddisruptionbudget
  ## This configuration is immutable once created and will require the PDB to be deleted to be changed
  ## https://github.com/kubernetes/kubernetes/issues/45398
  ##
  podDisruptionBudget:
    enabled: false
    minAvailable: 1
    maxUnavailable: ""

# Ingress exposes thanos sidecar outside the clsuter
  thanosIngress:
    enabled: false
    annotations: {}
    labels: {}
    servicePort: 10901
    ## Hosts must be provided if Ingress is enabled.
    ##
    hosts: []
      # - thanos-gateway.domain.com

    ## Paths to use for ingress rules
    ##
    paths: []
    # - /

    ## TLS configuration for Alertmanager Ingress
    ## Secret must be manually created in the namespace
    ##
    tls: []
    # - secretName: thanos-gateway-tls
    #   hosts:
    #   - thanos-gateway.domain.com

  ingress:
    enabled: false
    annotations: {}
    labels: {}

    ## Hostnames.
    ## Must be provided if Ingress is enabled.
    ##
    # hosts:
    #   - prometheus.domain.com
    hosts: []

    ## Paths to use for ingress rules - one path should match the prometheusSpec.routePrefix
    ##
    paths: []
    # - /

    ## TLS configuration for Prometheus Ingress
    ## Secret must be manually created in the namespace
    ##
    tls: []
      # - secretName: prometheus-general-tls
      #   hosts:
      #     - prometheus.example.com

  ## Configuration for creating an Ingress that will map to each Prometheus replica service
  ## prometheus.servicePerReplica must be enabled
  ##
  ingressPerReplica:
    enabled: false
    annotations: {}
    labels: {}

    ## Final form of the hostname for each per replica ingress is
    ## {{ ingressPerReplica.hostPrefix }}-{{ $replicaNumber }}.{{ ingressPerReplica.hostDomain }}
    ##
    ## Prefix for the per replica ingress that will have `-$replicaNumber`
    ## appended to the end
    hostPrefix: ""
    ## Domain that will be used for the per replica ingress
    hostDomain: ""

    ## Paths to use for ingress rules
    ##
    paths: []
    # - /

    ## Secret name containing the TLS certificate for Prometheus per replica ingress
    ## Secret must be manually created in the namespace
    tlsSecretName: ""

    ## Separated secret for each per replica Ingress. Can be used together with cert-manager
    ##
    tlsSecretPerReplica:
      enabled: false
      ## Final form of the secret for each per replica ingress is
      ## {{ tlsSecretPerReplica.prefix }}-{{ $replicaNumber }}
      ##
      prefix: "prometheus"

  ## Configure additional options for default pod security policy for Prometheus
  ## ref: https://kubernetes.io/docs/concepts/policy/pod-security-policy/
  podSecurityPolicy:
    allowedCapabilities: []

  serviceMonitor:
    ## Scrape interval. If not set, the Prometheus default scrape interval is used.
    ##
    interval: ""
    selfMonitor: true

    ## scheme: HTTP scheme to use for scraping. Can be used with `tlsConfig` for example if using istio mTLS.
    scheme: ""

    ## tlsConfig: TLS configuration to use when scraping the endpoint. For example if using istio mTLS.
    ## Of type: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#tlsconfig
    tlsConfig: {}

    bearerTokenFile:

    ## 	metric relabel configs to apply to samples before ingestion.
    ##
    metricRelabelings: []
    # - action: keep
    #   regex: 'kube_(daemonset|deployment|pod|namespace|node|statefulset).+'
    #   sourceLabels: [__name__]

    # 	relabel configs to apply to samples before ingestion.
    ##
    relabelings: []
    # - sourceLabels: [__meta_kubernetes_pod_node_name]
    #   separator: ;
    #   regex: ^(.*)$
    #   targetLabel: nodename
    #   replacement: $1
    #   action: replace

  ## Settings affecting prometheusSpec
  ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#prometheusspec
  ##
  prometheusSpec:
    ## If true, pass --storage.tsdb.max-block-duration=2h to prometheus. This is already done if using Thanos
    ##
    disableCompaction: false
    ## APIServerConfig
    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#apiserverconfig
    ##
    apiserverConfig: {}

    ## Interval between consecutive scrapes.
    ##
    scrapeInterval: ""

    ## Interval between consecutive evaluations.
    ##
    evaluationInterval: ""

    ## ListenLocal makes the Prometheus server listen on loopback, so that it does not bind against the Pod IP.
    ##
    listenLocal: false

    ## EnableAdminAPI enables Prometheus the administrative HTTP API which includes functionality such as deleting time series.
    ## This is disabled by default.
    ## ref: https://prometheus.io/docs/prometheus/latest/querying/api/#tsdb-admin-apis
    ##
    enableAdminAPI: false

    ## Image of Prometheus.
    ##
    image:
      repository: quay.io/prometheus/prometheus
      tag: v2.18.2
      sha: ""

    ## Tolerations for use with node taints
    ## ref: https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/
    ##
    tolerations: []
    #  - key: "key"
    #    operator: "Equal"
    #    value: "value"
    #    effect: "NoSchedule"

    ## Alertmanagers to which alerts will be sent
    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#alertmanagerendpoints
    ##
    ## Default configuration will connect to the alertmanager deployed as part of this release
    ##
    alertingEndpoints: []
    # - name: ""
    #   namespace: ""
    #   port: http
    #   scheme: http
    #   pathPrefix: ""
    #   tlsConfig: {}
    #   bearerTokenFile: ""
    #   apiVersion: v2

    ## External labels to add to any time series or alerts when communicating with external systems
    ##
    externalLabels: {}

    ## Name of the external label used to denote replica name
    ##
    replicaExternalLabelName: ""

    ## If true, the Operator won't add the external label used to denote replica name
    ##
    replicaExternalLabelNameClear: false

    ## Name of the external label used to denote Prometheus instance name
    ##
    prometheusExternalLabelName: ""

    ## If true, the Operator won't add the external label used to denote Prometheus instance name
    ##
    prometheusExternalLabelNameClear: false

    ## External URL at which Prometheus will be reachable.
    ##
    externalUrl: ""

    ## Define which Nodes the Pods are scheduled on.
    ## ref: https://kubernetes.io/docs/user-guide/node-selection/
    ##
    nodeSelector: {}

    ## Secrets is a list of Secrets in the same namespace as the Prometheus object, which shall be mounted into the Prometheus Pods.
    ## The Secrets are mounted into /etc/prometheus/secrets/. Secrets changes after initial creation of a Prometheus object are not
    ## reflected in the running Pods. To change the secrets mounted into the Prometheus Pods, the object must be deleted and recreated
    ## with the new list of secrets.
    ##
    secrets: []

    ## ConfigMaps is a list of ConfigMaps in the same namespace as the Prometheus object, which shall be mounted into the Prometheus Pods.
    ## The ConfigMaps are mounted into /etc/prometheus/configmaps/.
    ##
    configMaps: []

    ## QuerySpec defines the query command line flags when starting Prometheus.
    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#queryspec
    ##
    query: {}

    ## Namespaces to be selected for PrometheusRules discovery.
    ## If nil, select own namespace. Namespaces to be selected for ServiceMonitor discovery.
    ## See https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#namespaceselector for usage
    ##
    ruleNamespaceSelector: {}

    ## If true, a nil or {} value for prometheus.prometheusSpec.ruleSelector will cause the
    ## prometheus resource to be created with selectors based on values in the helm deployment,
    ## which will also match the PrometheusRule resources created
    ##
    ruleSelectorNilUsesHelmValues: true

    ## PrometheusRules to be selected for target discovery.
    ## If {}, select all ServiceMonitors
    ##
    ruleSelector: {}
    ## Example which select all prometheusrules resources
    ## with label "prometheus" with values any of "example-rules" or "example-rules-2"
    # ruleSelector:
    #   matchExpressions:
    #     - key: prometheus
    #       operator: In
    #       values:
    #         - example-rules
    #         - example-rules-2
    #
    ## Example which select all prometheusrules resources with label "role" set to "example-rules"
    # ruleSelector:
    #   matchLabels:
    #     role: example-rules

    ## If true, a nil or {} value for prometheus.prometheusSpec.serviceMonitorSelector will cause the
    ## prometheus resource to be created with selectors based on values in the helm deployment,
    ## which will also match the servicemonitors created
    ##
    serviceMonitorSelectorNilUsesHelmValues: true

    ## ServiceMonitors to be selected for target discovery.
    ## If {}, select all ServiceMonitors
    ##
    serviceMonitorSelector: {}
    ## Example which selects ServiceMonitors with label "prometheus" set to "somelabel"
    # serviceMonitorSelector:
    #   matchLabels:
    #     prometheus: somelabel

    ## Namespaces to be selected for ServiceMonitor discovery.
    ## See https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#namespaceselector for usage
    ##
    serviceMonitorNamespaceSelector: {}

    ## If true, a nil or {} value for prometheus.prometheusSpec.podMonitorSelector will cause the
    ## prometheus resource to be created with selectors based on values in the helm deployment,
    ## which will also match the podmonitors created
    ##
    podMonitorSelectorNilUsesHelmValues: true

    ## PodMonitors to be selected for target discovery.
    ## If {}, select all PodMonitors
    ##
    podMonitorSelector: {}
    ## Example which selects PodMonitors with label "prometheus" set to "somelabel"
    # podMonitorSelector:
    #   matchLabels:
    #     prometheus: somelabel

    ## Namespaces to be selected for PodMonitor discovery.
    ## See https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#namespaceselector for usage
    ##
    podMonitorNamespaceSelector: {}

    ## How long to retain metrics
    ##
    retention: 10d

    ## Maximum size of metrics
    ##
    retentionSize: ""

    ## Enable compression of the write-ahead log using Snappy.
    ##
    walCompression: false

    ## If true, the Operator won't process any Prometheus configuration changes
    ##
    paused: false

    ## Number of Prometheus replicas desired
    ##
    replicas: 1

    ## Log level for Prometheus be configured in
    ##
    logLevel: info

    ## Log format for Prometheus be configured in
    ##
    logFormat: logfmt

    ## Prefix used to register routes, overriding externalUrl route.
    ## Useful for proxies that rewrite URLs.
    ##
    routePrefix: /

    ## Standard object’s metadata. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata
    ## Metadata Labels and Annotations gets propagated to the prometheus pods.
    ##
    podMetadata: {}
    # labels:
    #   app: prometheus
    #   k8s-app: prometheus

    ## Pod anti-affinity can prevent the scheduler from placing Prometheus replicas on the same node.
    ## The default value "soft" means that the scheduler should *prefer* to not schedule two replica pods onto the same node but no guarantee is provided.
    ## The value "hard" means that the scheduler is *required* to not schedule two replica pods onto the same node.
    ## The value "" will disable pod anti-affinity so that no anti-affinity rules will be configured.
    podAntiAffinity: ""

    ## If anti-affinity is enabled sets the topologyKey to use for anti-affinity.
    ## This can be changed to, for example, failure-domain.beta.kubernetes.io/zone
    ##
    podAntiAffinityTopologyKey: kubernetes.io/hostname

    ## Assign custom affinity rules to the prometheus instance
    ## ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
    ##
    affinity: {}
    # nodeAffinity:
    #   requiredDuringSchedulingIgnoredDuringExecution:
    #     nodeSelectorTerms:
    #     - matchExpressions:
    #       - key: kubernetes.io/e2e-az-name
    #         operator: In
    #         values:
    #         - e2e-az1
    #         - e2e-az2

    ## The remote_read spec configuration for Prometheus.
    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#remotereadspec
    remoteRead: []
    # - url: http://remote1/read

    ## The remote_write spec configuration for Prometheus.
    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#remotewritespec
    remoteWrite: []
    # - url: http://remote1/push

    ## Enable/Disable Grafana dashboards provisioning for prometheus remote write feature
    remoteWriteDashboards: false

    ## Resource limits & requests
    ##
    resources: {}
    # requests:
    #   memory: 400Mi

    ## Prometheus StorageSpec for persistent data
    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/user-guides/storage.md
    ##
    storageSpec: {}
    #  volumeClaimTemplate:
    #    spec:
    #      storageClassName: gluster
    #      accessModes: ["ReadWriteOnce"]
    #      resources:
    #        requests:
    #          storage: 50Gi
    #    selector: {}

    # Additional volumes on the output StatefulSet definition.
    volumes: []
    # Additional VolumeMounts on the output StatefulSet definition.
    volumeMounts: []

    ## AdditionalScrapeConfigs allows specifying additional Prometheus scrape configurations. Scrape configurations
    ## are appended to the configurations generated by the Prometheus Operator. Job configurations must have the form
    ## as specified in the official Prometheus documentation:
    ## https://prometheus.io/docs/prometheus/latest/configuration/configuration/#scrape_config. As scrape configs are
    ## appended, the user is responsible to make sure it is valid. Note that using this feature may expose the possibility
    ## to break upgrades of Prometheus. It is advised to review Prometheus release notes to ensure that no incompatible
    ## scrape configs are going to break Prometheus after the upgrade.
    ##
    ## The scrape configuraiton example below will find master nodes, provided they have the name .*mst.*, relabel the
    ## port to 2379 and allow etcd scraping provided it is running on all Kubernetes master nodes
    ##
    additionalScrapeConfigs: []
    # - job_name: kube-etcd
    #   kubernetes_sd_configs:
    #     - role: node
    #   scheme: https
    #   tls_config:
    #     ca_file:   /etc/prometheus/secrets/etcd-client-cert/etcd-ca
    #     cert_file: /etc/prometheus/secrets/etcd-client-cert/etcd-client
    #     key_file:  /etc/prometheus/secrets/etcd-client-cert/etcd-client-key
    #   relabel_configs:
    #   - action: labelmap
    #     regex: __meta_kubernetes_node_label_(.+)
    #   - source_labels: [__address__]
    #     action: replace
    #     targetLabel: __address__
    #     regex: ([^:;]+):(\d+)
    #     replacement: ${1}:2379
    #   - source_labels: [__meta_kubernetes_node_name]
    #     action: keep
    #     regex: .*mst.*
    #   - source_labels: [__meta_kubernetes_node_name]
    #     action: replace
    #     targetLabel: node
    #     regex: (.*)
    #     replacement: ${1}
    #   metric_relabel_configs:
    #   - regex: (kubernetes_io_hostname|failure_domain_beta_kubernetes_io_region|beta_kubernetes_io_os|beta_kubernetes_io_arch|beta_kubernetes_io_instance_type|failure_domain_beta_kubernetes_io_zone)
    #     action: labeldrop

    ## If additional scrape configurations are already deployed in a single secret file you can use this section.
    ## Expected values are the secret name and key
    ## Cannot be used with additionalScrapeConfigs
    additionalScrapeConfigsSecret: {}
      # enabled: false
      # name:
      # key:

    ## additionalPrometheusSecretsAnnotations allows to add annotations to the kubernetes secret. This can be useful
    ## when deploying via spinnaker to disable versioning on the secret, strategy.spinnaker.io/versioned: 'false'
    additionalPrometheusSecretsAnnotations: {}

    ## AdditionalAlertManagerConfigs allows for manual configuration of alertmanager jobs in the form as specified
    ## in the official Prometheus documentation https://prometheus.io/docs/prometheus/latest/configuration/configuration/#<alertmanager_config>.
    ## AlertManager configurations specified are appended to the configurations generated by the Prometheus Operator.
    ## As AlertManager configs are appended, the user is responsible to make sure it is valid. Note that using this
    ## feature may expose the possibility to break upgrades of Prometheus. It is advised to review Prometheus release
    ## notes to ensure that no incompatible AlertManager configs are going to break Prometheus after the upgrade.
    ##
    additionalAlertManagerConfigs: []
    # - consul_sd_configs:
    #   - server: consul.dev.test:8500
    #     scheme: http
    #     datacenter: dev
    #     tag_separator: ','
    #     services:
    #       - metrics-prometheus-alertmanager

    ## AdditionalAlertRelabelConfigs allows specifying Prometheus alert relabel configurations. Alert relabel configurations specified are appended
    ## to the configurations generated by the Prometheus Operator. Alert relabel configurations specified must have the form as specified in the
    ## official Prometheus documentation: https://prometheus.io/docs/prometheus/latest/configuration/configuration/#alert_relabel_configs.
    ## As alert relabel configs are appended, the user is responsible to make sure it is valid. Note that using this feature may expose the
    ## possibility to break upgrades of Prometheus. It is advised to review Prometheus release notes to ensure that no incompatible alert relabel
    ## configs are going to break Prometheus after the upgrade.
    ##
    additionalAlertRelabelConfigs: []
    # - separator: ;
    #   regex: prometheus_replica
    #   replacement: $1
    #   action: labeldrop

    ## SecurityContext holds pod-level security attributes and common container settings.
    ## This defaults to non root user with uid 1000 and gid 2000.
    ## https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md
    ##
    securityContext:
      runAsGroup: 2000
      runAsNonRoot: true
      runAsUser: 1000
      fsGroup: 2000

    ## 	Priority class assigned to the Pods
    ##
    priorityClassName: ""

    ## Thanos configuration allows configuring various aspects of a Prometheus server in a Thanos environment.
    ## This section is experimental, it may change significantly without deprecation notice in any release.
    ## This is experimental and may change significantly without backward compatibility in any release.
    ## ref: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#thanosspec
    ##
    thanos: {}

    ## Containers allows injecting additional containers. This is meant to allow adding an authentication proxy to a Prometheus pod.
    ##  if using proxy extraContainer  update targetPort with proxy container port
    containers: []

    ## InitContainers allows injecting additional initContainers. This is meant to allow doing some changes
    ## (permissions, dir tree) on mounted volumes before starting prometheus
    initContainers: []

    ## PortName to use for Prometheus.
    ##
    portName: "web"

  additionalServiceMonitors: []
  ## Name of the ServiceMonitor to create
  ##
  # - name: ""

    ## Additional labels to set used for the ServiceMonitorSelector. Together with standard labels from
    ## the chart
    ##
    # additionalLabels: {}

    ## Service label for use in assembling a job name of the form <label value>-<port>
    ## If no label is specified, the service name is used.
    ##
    # jobLabel: ""

    ## labels to transfer from the kubernetes service to the target
    ##
    # targetLabels: ""

    ## Label selector for services to which this ServiceMonitor applies
    ##
    # selector: {}

    ## Namespaces from which services are selected
    ##
    # namespaceSelector:
      ## Match any namespace
      ##
      # any: false

      ## Explicit list of namespace names to select
      ##
      # matchNames: []

    ## Endpoints of the selected service to be monitored
    ##
    # endpoints: []
      ## Name of the endpoint's service port
      ## Mutually exclusive with targetPort
      # - port: ""

      ## Name or number of the endpoint's target port
      ## Mutually exclusive with port
      # - targetPort: ""

      ## File containing bearer token to be used when scraping targets
      ##
      #   bearerTokenFile: ""

      ## Interval at which metrics should be scraped
      ##
      #   interval: 30s

      ## HTTP path to scrape for metrics
      ##
      #   path: /metrics

      ## HTTP scheme to use for scraping
      ##
      #   scheme: http

      ## TLS configuration to use when scraping the endpoint
      ##
      #   tlsConfig:

          ## Path to the CA file
          ##
          # caFile: ""

          ## Path to client certificate file
          ##
          # certFile: ""

          ## Skip certificate verification
          ##
          # insecureSkipVerify: false

          ## Path to client key file
          ##
          # keyFile: ""

          ## Server name used to verify host name
          ##
          # serverName: ""

  additionalPodMonitors: []
  ## Name of the PodMonitor to create
  ##
  # - name: ""

    ## Additional labels to set used for the PodMonitorSelector. Together with standard labels from
    ## the chart
    ##
    # additionalLabels: {}

    ## Pod label for use in assembling a job name of the form <label value>-<port>
    ## If no label is specified, the pod endpoint name is used.
    ##
    # jobLabel: ""

    ## Label selector for pods to which this PodMonitor applies
    ##
    # selector: {}

    ## PodTargetLabels transfers labels on the Kubernetes Pod onto the target.
    ##
    # podTargetLabels: {}

    ## SampleLimit defines per-scrape limit on number of scraped samples that will be accepted.
    ##
    # sampleLimit: 0

    ## Namespaces from which pods are selected
    ##
    # namespaceSelector:
      ## Match any namespace
      ##
      # any: false

      ## Explicit list of namespace names to select
      ##
      # matchNames: []

    ## Endpoints of the selected pods to be monitored
    ## https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/api.md#podmetricsendpoint
    ##
    # podMetricsEndpoints: []

```

---
Autogenerated from Helm Chart and git history using [helm-changelog](https://github.com/mogensen/helm-changelog)
