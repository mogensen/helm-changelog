# Change Log

## 14.5.0 

**Release date:** 2021-04-03

![AppVersion: 0.46.0](https://img.shields.io/static/v1?label=AppVersion&message=0.46.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] Invalid value "None" for clusterIP (#784) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 42f6ac7..290ea1d 100644
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

![AppVersion: 0.46.0](https://img.shields.io/static/v1?label=AppVersion&message=0.46.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] add more controls to kubescheduler,kubeproxy,controllermanager,etcd exporters (#797) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 5848603..42f6ac7 100644
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

**Release date:** 2021-03-21

![AppVersion: 0.46.0](https://img.shields.io/static/v1?label=AppVersion&message=0.46.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] Exclude default rules from namespace enforcement (#770) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index efd99ac..5848603 100644
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

![AppVersion: 0.46.0](https://img.shields.io/static/v1?label=AppVersion&message=0.46.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] Update dependency version (#776) 

### Default value changes

```diff
# No changes in this release
```

## 14.1.2 

**Release date:** 2021-03-18

![AppVersion: 0.46.0](https://img.shields.io/static/v1?label=AppVersion&message=0.46.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Fix LoadBalancer typo in service type (#761) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index e59a17b..efd99ac 100644
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

![AppVersion: 0.46.0](https://img.shields.io/static/v1?label=AppVersion&message=0.46.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] Fix typo and formatting in comments (#772) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 961c619..e59a17b 100644
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

**Release date:** 2021-03-19

![AppVersion: 0.46.0](https://img.shields.io/static/v1?label=AppVersion&message=0.46.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] Added nodePort to thanosService (#769) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 2ec7360..961c619 100644
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

![AppVersion: 0.46.0](https://img.shields.io/static/v1?label=AppVersion&message=0.46.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] Update README with 13.x to 14.x upgrade instructions (#746) 

### Default value changes

```diff
# No changes in this release
```

## 14.0.0 

**Release date:** 2021-03-08

![AppVersion: 0.46.0](https://img.shields.io/static/v1?label=AppVersion&message=0.46.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] 0.46.0 operator and crds (#740) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 66114e2..2ec7360 100644
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

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Added missing upgrade command for alertmanager to do before 13.13.x (#715) 

### Default value changes

```diff
# No changes in this release
```

## 13.13.0 

**Release date:** 2021-02-23

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] Upgrade to latest kube-state-metrics chart with v1.9.8 (#703) 

### Default value changes

```diff
# No changes in this release
```

## 13.12.0 

**Release date:** 2021-02-23

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* add Prometheus and Alertmanager topologySpreadConstraints (#702) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 8b36373..66114e2 100644
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

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Added pathType (#688) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 59467c0..8b36373 100644
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

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] Upgrade grafana dependency to 6.4 (#685) 

### Default value changes

```diff
# No changes in this release
```

## 13.9.1 

**Release date:** 2021-02-17

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] Ability for custom dnsConfig (#639) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 0166596..59467c0 100644
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

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] hack: fix syncing dashboards and alerting rules (#665) 

### Default value changes

```diff
# No changes in this release
```

## 13.8.0 

**Release date:** 2021-02-17

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] update node exporter chart to 1.14.* (#677) 

### Default value changes

```diff
# No changes in this release
```

## 13.7.2 

**Release date:** 2021-02-10

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Ability for cluster-domain in kube-prometheus-stack (#647) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 0bba5e9..0166596 100644
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

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* fix yaml indent (#663) 

### Default value changes

```diff
# No changes in this release
```

## 13.7.0 

**Release date:** 2021-02-10

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] introduce optional param grafana.sidecar.datasources.defaultDatasourceScrapeInterval. (#659) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index caa38b0..0bba5e9 100644
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

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* fix: allow overriding of selector for kube-state-metrics serviceMonitor (#303) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index f367e96..caa38b0 100644
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

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] Sync kube-state-metrics (#650) 

### Default value changes

```diff
# No changes in this release
```

## 13.5.0 

**Release date:** 2021-02-04

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Add podTargetLabels to additionalServiceMonitors (#629) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 617b39c..f367e96 100644
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

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* fix: match containerPort and listen-address in operator (#624) 

### Default value changes

```diff
# No changes in this release
```

## 13.4.0 

**Release date:** 2021-01-29

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] Add support for specifying webhook port for GKE workaround (#400) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 2a0f98c..617b39c 100644
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

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] Allow override of prometheus and alertmanager default base images (#588) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 2be42c8..2a0f98c 100644
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

**Release date:** 2021-01-21

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [Kube-prometheus-stack]adding upstream for external label and url (#536) 

### Default value changes

```diff
# No changes in this release
```

## 13.2.0 

**Release date:** 2021-01-21

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Make it possible to add additional volume and allowerHostPath to prometehus psp (#583) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 8fbccb3..2be42c8 100644
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

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] Update list of excluded fs types and mountpoints in node exporter collector (#594) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index dc497e9..8fbccb3 100644
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

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] Add config.templates (#563) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 32730b5..dc497e9 100644
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

**Release date:** 2021-01-21

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Use new API version for ingress (#572) 

### Default value changes

```diff
# No changes in this release
```

## 13.0.4 

**Release date:** 2021-01-20

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success&logo=)
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

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Fix thanos grpc port when using service type NodePort (#593) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index dbf5ef6..32730b5 100644
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

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack]: enable multi cluster dashboards (#575) 

### Default value changes

```diff
# No changes in this release
```

## 13.0.1 

**Release date:** 2021-01-19

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] Add grpc thanos port to prometheus service (#504) 

### Default value changes

```diff
# No changes in this release
```

## 13.0.0 

**Release date:** 2021-01-18

![AppVersion: 0.45.0](https://img.shields.io/static/v1?label=AppVersion&message=0.45.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] Version bump prometheus-operator to 0.45.0 (#582) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index b83d389..dbf5ef6 100644
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

![AppVersion: 0.44.0](https://img.shields.io/static/v1?label=AppVersion&message=0.44.0&color=success&logo=)
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

![AppVersion: 0.44.0](https://img.shields.io/static/v1?label=AppVersion&message=0.44.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] Migrate kube-state-metrics off helm/stable (#581) 

### Default value changes

```diff
# No changes in this release
```

## 12.12.0 

**Release date:** 2021-01-13

![AppVersion: 0.44.0](https://img.shields.io/static/v1?label=AppVersion&message=0.44.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Add option to add headless svc for thanos sidecar (#564) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 4c46c13..b83d389 100644
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

![AppVersion: 0.44.0](https://img.shields.io/static/v1?label=AppVersion&message=0.44.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack]values:fix missleading comment (#553) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index b9482cf..4c46c13 100644
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

![AppVersion: 0.44.0](https://img.shields.io/static/v1?label=AppVersion&message=0.44.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] prometheus: add sharding mechanism (#552) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index bb93e35..b9482cf 100644
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

![AppVersion: 0.44.0](https://img.shields.io/static/v1?label=AppVersion&message=0.44.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] Allow setting caBundle value for admission webhook configuration (#543) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index a864b59..bb93e35 100644
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

![AppVersion: 0.44.0](https://img.shields.io/static/v1?label=AppVersion&message=0.44.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] 333 - need to template the alertmanager tls settings as well (#539) 

### Default value changes

```diff
# No changes in this release
```

## 12.10.4 

**Release date:** 2021-01-04

![AppVersion: 0.44.0](https://img.shields.io/static/v1?label=AppVersion&message=0.44.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] allow to select alertmanagerConfig resources from all namespaces (#503) 

### Default value changes

```diff
# No changes in this release
```

## 12.10.3 

**Release date:** 2021-01-04

![AppVersion: 0.44.0](https://img.shields.io/static/v1?label=AppVersion&message=0.44.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Fix accessing root in range block for datasources.timeInterval (#518) 

### Default value changes

```diff
# No changes in this release
```

## 12.10.2 

**Release date:** 2021-01-05

![AppVersion: 0.44.0](https://img.shields.io/static/v1?label=AppVersion&message=0.44.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Avoid relying on image/container working directory for TLS cert mounting in the `prometheus-operator` `Deployment` (#530) 

### Default value changes

```diff
# No changes in this release
```

## 12.10.1 

**Release date:** 2021-01-04

![AppVersion: 0.44.0](https://img.shields.io/static/v1?label=AppVersion&message=0.44.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] Fix an errant dash in prometheus-operator's --kubelet-service arg (#535) 

### Default value changes

```diff
# No changes in this release
```

## 12.10.0 

**Release date:** 2021-01-04

![AppVersion: 0.44.0](https://img.shields.io/static/v1?label=AppVersion&message=0.44.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] Upgrade grafan dependency to 6.1, bump major version to match grafana breaking change, update prometheusOperator.configReloaderMemory from 25Mi to 50Mi to follow 0.44.0 change. (#460) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 690b53e..a864b59 100644
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

**Release date:** 2021-01-02

![AppVersion: 0.44.0](https://img.shields.io/static/v1?label=AppVersion&message=0.44.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Correct path to `ServiceMonitor` ca cert when running with tls enabled (#531) 

### Default value changes

```diff
# No changes in this release
```

## 12.9.1 

**Release date:** 2021-01-01

![AppVersion: 0.44.0](https://img.shields.io/static/v1?label=AppVersion&message=0.44.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] Fix Capabilities check for IngressClass (#489) 

### Default value changes

```diff
# No changes in this release
```

## 12.9.0 

**Release date:** 2020-12-30

![AppVersion: 0.44.0](https://img.shields.io/static/v1?label=AppVersion&message=0.44.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] cert-manager support for operator webhooks (#481) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index e2c481f..690b53e 100644
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

![AppVersion: 0.44.0](https://img.shields.io/static/v1?label=AppVersion&message=0.44.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* fixed typos (#496) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 3aa810f..e2c481f 100644
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

![AppVersion: 0.44.0](https://img.shields.io/static/v1?label=AppVersion&message=0.44.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] #356 add support for multicluster in grafana dashboards (#357) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 774359f..3aa810f 100644
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

![AppVersion: 0.44.0](https://img.shields.io/static/v1?label=AppVersion&message=0.44.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] Configure possibility of providing additional rules to prometheus cluster role (#413) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index f243f8c..774359f 100644
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

![AppVersion: 0.44.0](https://img.shields.io/static/v1?label=AppVersion&message=0.44.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Align Grafana and Prometheus scrape intervals (#457) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index c2c9949..f243f8c 100644
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

![AppVersion: 0.44.0](https://img.shields.io/static/v1?label=AppVersion&message=0.44.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] Add missing fields to prometheus and alertmanager resources (#208) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index f768a47..c2c9949 100644
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

![AppVersion: 0.44.0](https://img.shields.io/static/v1?label=AppVersion&message=0.44.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] Enable AggregatedAPIDown for k8s>=1.18 only (#455) 

### Default value changes

```diff
# No changes in this release
```

## 12.4.0 

**Release date:** 2020-12-02

![AppVersion: 0.44.0](https://img.shields.io/static/v1?label=AppVersion&message=0.44.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] Update prometheus-operator to v0.44.0 (#441) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 5d0f47e..f768a47 100644
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

![AppVersion: 0.43.2](https://img.shields.io/static/v1?label=AppVersion&message=0.43.2&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] Sync grafana dashboards and prometheus rules from kube-prometheus (#432) 

### Default value changes

```diff
# No changes in this release
```

## 12.2.4 

**Release date:** 2020-11-28

![AppVersion: 0.43.2](https://img.shields.io/static/v1?label=AppVersion&message=0.43.2&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [codespell] fix readme (#427) 

### Default value changes

```diff
# No changes in this release
```

## 12.2.3 

**Release date:** 2020-11-24

![AppVersion: 0.43.2](https://img.shields.io/static/v1?label=AppVersion&message=0.43.2&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] Fix broken volumes and volumeMounts fields on alertmanager (#409) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 76b72ad..5d0f47e 100644
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

![AppVersion: 0.43.2](https://img.shields.io/static/v1?label=AppVersion&message=0.43.2&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] Allow custom TLS min version for operator (#407) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 3371203..76b72ad 100644
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

![AppVersion: 0.43.2](https://img.shields.io/static/v1?label=AppVersion&message=0.43.2&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* kube-prometheus-stack: fix ports alignment in service (#404) 

### Default value changes

```diff
# No changes in this release
```

## 12.2.0 

**Release date:** 2020-11-21

![AppVersion: 0.43.2](https://img.shields.io/static/v1?label=AppVersion&message=0.43.2&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* kube-prometheus-stack: Add optional extra ports to alertmanager (#396) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 4a7e27d..3371203 100644
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

![AppVersion: 0.43.2](https://img.shields.io/static/v1?label=AppVersion&message=0.43.2&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Allow rules be enabled independent of exporters (#330) 

### Default value changes

```diff
# No changes in this release
```

## 12.0.4 

**Release date:** 2020-11-18

![AppVersion: 0.43.2](https://img.shields.io/static/v1?label=AppVersion&message=0.43.2&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Clarify log format options in kube-prometheus-stack (#382) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 79803a7..4a7e27d 100644
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

![AppVersion: 0.43.2](https://img.shields.io/static/v1?label=AppVersion&message=0.43.2&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [fix] Correct provisioning for the --thanos-ruler-instance-namespaces prometheus-operator (#386) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 17e5d7d..79803a7 100644
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

![AppVersion: 0.43.2](https://img.shields.io/static/v1?label=AppVersion&message=0.43.2&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Fix some GitHub broken links on kube-prometheus-stack/hack/README.md (#344) 

### Default value changes

```diff
# No changes in this release
```

## 12.0.1 

**Release date:** 2020-11-17

![AppVersion: 0.43.2](https://img.shields.io/static/v1?label=AppVersion&message=0.43.2&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] Add missing parameter scrapetimeout (#370) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 03f759b..17e5d7d 100644
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

![AppVersion: 0.43.2](https://img.shields.io/static/v1?label=AppVersion&message=0.43.2&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] Migrate to chart v2 (#307) 

### Default value changes

```diff
# No changes in this release
```

## 11.1.7 

**Release date:** 2020-11-17

![AppVersion: 0.43.2](https://img.shields.io/static/v1?label=AppVersion&message=0.43.2&color=success&logo=)
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

![AppVersion: 0.43.2](https://img.shields.io/static/v1?label=AppVersion&message=0.43.2&color=success&logo=)
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

![AppVersion: 0.43.2](https://img.shields.io/static/v1?label=AppVersion&message=0.43.2&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] Show tmpfs option for Prometheus storageSpec (#363) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index efa76c7..03f759b 100644
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

![AppVersion: 0.43.2](https://img.shields.io/static/v1?label=AppVersion&message=0.43.2&color=success&logo=)
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

![AppVersion: 0.43.2](https://img.shields.io/static/v1?label=AppVersion&message=0.43.2&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] Allow modifying ServiceMonitor path (#364) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 0c83259..efa76c7 100644
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

![AppVersion: 0.43.2](https://img.shields.io/static/v1?label=AppVersion&message=0.43.2&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] Fix indentation to allow IDE collapse of sections (#350) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 8a67816..0c83259 100644
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

![AppVersion: 0.43.2](https://img.shields.io/static/v1?label=AppVersion&message=0.43.2&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* charts/kube-prometheus-stack: fix typos (#318) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index d3012b7..8a67816 100644
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

![AppVersion: 0.43.2](https://img.shields.io/static/v1?label=AppVersion&message=0.43.2&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] Add alertmanagerConfigSelector option (#311) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index a253f59..d3012b7 100644
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

![AppVersion: 0.43.2](https://img.shields.io/static/v1?label=AppVersion&message=0.43.2&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] Bump prometheus-operator to 0.43.2 (#328) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 5db649d..a253f59 100644
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

![AppVersion: 0.43.1](https://img.shields.io/static/v1?label=AppVersion&message=0.43.1&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] add ingressclass support (#321) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 1bd16bc..5db649d 100644
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

![AppVersion: 0.43.1](https://img.shields.io/static/v1?label=AppVersion&message=0.43.1&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] Drop kubectl image field (#312) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 5a0810a..1bd16bc 100644
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

![AppVersion: 0.43.1](https://img.shields.io/static/v1?label=AppVersion&message=0.43.1&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] Disable scraping kubelet resource metrics endpoint (#301) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 1aded8f..5a0810a 100644
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

![AppVersion: 0.43.1](https://img.shields.io/static/v1?label=AppVersion&message=0.43.1&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] Bump prometheus-operator to 0.43.1 (#305) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 7dd1e86..1aded8f 100644
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

![AppVersion: 0.43.0](https://img.shields.io/static/v1?label=AppVersion&message=0.43.0&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] Update prometheus-operator to v0.43.0 (#280) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index fe90381..7dd1e86 100644
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

![AppVersion: 0.42.1](https://img.shields.io/static/v1?label=AppVersion&message=0.42.1&color=success&logo=)
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

![AppVersion: 0.42.1](https://img.shields.io/static/v1?label=AppVersion&message=0.42.1&color=success&logo=)
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

![AppVersion: 0.42.1](https://img.shields.io/static/v1?label=AppVersion&message=0.42.1&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack]: add additional alert labels on prometheusRules (#247) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 9a4397d..fe90381 100644
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

![AppVersion: 0.42.1](https://img.shields.io/static/v1?label=AppVersion&message=0.42.1&color=success&logo=)
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

![AppVersion: 0.42.1](https://img.shields.io/static/v1?label=AppVersion&message=0.42.1&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] Support native TLS in prometheus-operator (#206) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 0530e48..9a4397d 100644
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

![AppVersion: 0.42.1](https://img.shields.io/static/v1?label=AppVersion&message=0.42.1&color=success&logo=)
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

![AppVersion: 0.42.1](https://img.shields.io/static/v1?label=AppVersion&message=0.42.1&color=success&logo=)
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

![AppVersion: 0.42.1](https://img.shields.io/static/v1?label=AppVersion&message=0.42.1&color=success&logo=)
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

![AppVersion: 0.42.1](https://img.shields.io/static/v1?label=AppVersion&message=0.42.1&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] Doc fixes (#246) 
* Add support for ProbeSelector and ProbeNamespaceSelector (#237) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index c672971..0530e48 100644
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

![AppVersion: 0.42.1](https://img.shields.io/static/v1?label=AppVersion&message=0.42.1&color=success&logo=)
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

![AppVersion: 0.42.1](https://img.shields.io/static/v1?label=AppVersion&message=0.42.1&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] Add instance namespaces arguments for prometheus-operator deployment (#92) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index e722503..c672971 100644
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

![AppVersion: 0.42.1](https://img.shields.io/static/v1?label=AppVersion&message=0.42.1&color=success&logo=)
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

![AppVersion: 0.42.1](https://img.shields.io/static/v1?label=AppVersion&message=0.42.1&color=success&logo=)
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

![AppVersion: 0.42.1](https://img.shields.io/static/v1?label=AppVersion&message=0.42.1&color=success&logo=)
![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=>=1.16.0-0&color=informational&logo=kubernetes)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] Update prometheus-operator to 0.42.1 (#128) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 932bf66..e722503 100644
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

![AppVersion: 0.38.1](https://img.shields.io/static/v1?label=AppVersion&message=0.38.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] fix broken link in comment (#141) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index bcfb003..932bf66 100644
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

![AppVersion: 0.38.1](https://img.shields.io/static/v1?label=AppVersion&message=0.38.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] Resync rules and dashboards (#168) 

### Default value changes

```diff
# No changes in this release
```

## 9.4.8 

**Release date:** 2020-10-05

![AppVersion: 0.38.1](https://img.shields.io/static/v1?label=AppVersion&message=0.38.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] style(grafana): fix typo (#150) 

### Default value changes

```diff
# No changes in this release
```

## 9.4.7 

**Release date:** 2020-10-04

![AppVersion: 0.38.1](https://img.shields.io/static/v1?label=AppVersion&message=0.38.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] Replacing hyperkube image with kubectl image from bitnami (#160) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index e0ef831..bcfb003 100644
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

![AppVersion: 0.38.1](https://img.shields.io/static/v1?label=AppVersion&message=0.38.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Added more options for AlertManager ServiceMonitoring (#169) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 5202bac..e0ef831 100644
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

![AppVersion: 0.38.1](https://img.shields.io/static/v1?label=AppVersion&message=0.38.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] add podmonitor docs (#164) 

### Default value changes

```diff
# No changes in this release
```

## 9.4.4 

**Release date:** 2020-09-21

![AppVersion: 0.38.1](https://img.shields.io/static/v1?label=AppVersion&message=0.38.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Add documentation on how to configure additionalPrometheusRulesMap (#109) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index 12d782d..5202bac 100644
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

![AppVersion: 0.38.1](https://img.shields.io/static/v1?label=AppVersion&message=0.38.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] Fix URL to node-exporter (#104) 

### Default value changes

```diff
# No changes in this release
```

## 9.4.2 

**Release date:** 2020-09-14

![AppVersion: 0.38.1](https://img.shields.io/static/v1?label=AppVersion&message=0.38.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] Add Artifact Hub annotations for named links and operator search facet (#95) 

### Default value changes

```diff
# No changes in this release
```

## 9.4.1 

**Release date:** 2020-09-11

![AppVersion: 0.38.1](https://img.shields.io/static/v1?label=AppVersion&message=0.38.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Update README.md in kube-prometheus-stack hack (#89) 

### Default value changes

```diff
# No changes in this release
```

## 9.4.0 

**Release date:** 2020-09-10

![AppVersion: 0.38.1](https://img.shields.io/static/v1?label=AppVersion&message=0.38.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kube-prometheus-stack] add secret-field-selector (#79) 

### Default value changes

```diff
diff --git a/charts/kube-prometheus-stack/values.yaml b/charts/kube-prometheus-stack/values.yaml
index b67195d..12d782d 100644
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

![AppVersion: 0.38.1](https://img.shields.io/static/v1?label=AppVersion&message=0.38.1&color=success&logo=)
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
