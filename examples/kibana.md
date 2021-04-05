# Change Log

## Next Release 

![AppVersion: 8.0.0-SNAPSHOT](https://img.shields.io/static/v1?label=AppVersion&message=8.0.0-SNAPSHOT&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [elasticsearch][kibana] remove oss examples (#1046) 
* [meta] add build status and artifact hub badges (#1033) 
* Fix post-lifecycle hook example (#1028) 
* [elasticsearch][kibana] Add flexible ingress (#994) 
* [all] add hostaliases (#970) 
* [kibana] add service.httpPortName config in chart (#843) 
* [meta] stabilize CI tests (#935) 
* [meta] remove support for k8s <1.14 & helm <2.17.0 (#916) 
* [meta] upgrade test (#907) 
* Helm 3 (#516) 
* [meta] increase helm timeout (#891) 
* Add warning comment placeholder (main branch) (#886) 
* Update warning on charts readme (#863) 
* support tpl in logstashConfig, logstashPipeline and kibanaConfig (#717) 
* [elasticsearch][kibana] disable nss dentry cache (#818) 
* [meta] update changelog for 7.9.2 release  (#824) 
* [meta] Update changelog for 6.8.12 / 7.9.0 releases (#789) 
* [meta] add helm 3 beta support (#759) 
* [meta] update changelog for 7.8.1 and 6.8.11 releases (#755) 
* Update README.md (#749) 
* [doc] fix some links (#737) 
* [Kibana] Add loadbalancerIP to service spec in kibana (#726) 
* [doc] fix copy-paste errors (#700) 

### Default value changes

```diff
diff --git a/kibana/values.yaml b/kibana/values.yaml
index ba6746c9..9f59286a 100755
--- a/kibana/values.yaml
+++ b/kibana/values.yaml
@@ -28,6 +28,12 @@ secretMounts: []
 #    path: /usr/share/kibana/data/kibana.keystore
 #    subPath: kibana.keystore # optional
 
+hostAliases: []
+#- ip: "127.0.0.1"
+#  hostnames:
+#  - "foo.local"
+#  - "bar.local"
+
 image: "docker.elastic.co/kibana/kibana"
 imageTag: "8.0.0-SNAPSHOT"
 imagePullPolicy: "IfNotPresent"
@@ -95,6 +101,7 @@ updateStrategy:
 
 service:
   type: ClusterIP
+  loadBalancerIP: ""
   port: 5601
   nodePort: ""
   labels: {}
@@ -106,15 +113,17 @@ service:
     # service.beta.kubernetes.io/cce-load-balancer-internal-vpc: "true"
   loadBalancerSourceRanges: []
     # 0.0.0.0/0
+  httpPortName: http
 
 ingress:
   enabled: false
   annotations: {}
     # kubernetes.io/ingress.class: nginx
     # kubernetes.io/tls-acme: "true"
-  path: /
   hosts:
-    - chart-example.local
+    - host: chart-example.local
+      paths:
+        - path: /
   tls: []
   #  - secretName: chart-example-tls
   #    hosts:
```

## 8.0.0-SNAPSHOT 

**Release date:** 2020-06-29

![AppVersion: 8.0.0-SNAPSHOT](https://img.shields.io/static/v1?label=AppVersion&message=8.0.0-SNAPSHOT&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Bump master branch to 8.0.0-SNAPSHOT (#682) 

### Default value changes

```diff
diff --git a/kibana/values.yaml b/kibana/values.yaml
index 486954a4..ba6746c9 100755
--- a/kibana/values.yaml
+++ b/kibana/values.yaml
@@ -29,7 +29,7 @@ secretMounts: []
 #    subPath: kibana.keystore # optional
 
 image: "docker.elastic.co/kibana/kibana"
-imageTag: "7.7.1"
+imageTag: "8.0.0-SNAPSHOT"
 imagePullPolicy: "IfNotPresent"
 
 # additionals labels
```

## 7.7.1 

**Release date:** 2020-06-03

![AppVersion: 7.7.1](https://img.shields.io/static/v1?label=AppVersion&message=7.7.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Update changelog for 7.7.1 and 6.8.10 releases (#652) 
* [kibana] Add extensible label support (#555) 
* [kibana] String/YAML conditions for `.Values.{extraContainers,extraInitContainers}` (#637) 

### Default value changes

```diff
diff --git a/kibana/values.yaml b/kibana/values.yaml
index b843b7f7..486954a4 100755
--- a/kibana/values.yaml
+++ b/kibana/values.yaml
@@ -29,7 +29,7 @@ secretMounts: []
 #    subPath: kibana.keystore # optional
 
 image: "docker.elastic.co/kibana/kibana"
-imageTag: "7.7.0"
+imageTag: "7.7.1"
 imagePullPolicy: "IfNotPresent"
 
 # additionals labels
```

## 7.7.0 

**Release date:** 2020-05-13

![AppVersion: 7.7.0](https://img.shields.io/static/v1?label=AppVersion&message=7.7.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [meta] 7.70 pre-release (#612) 
* FAQ and examples improvements (#598) 
* [meta] add toc to markdown files 
* Merge branch 'master' into doc-improvements 
* nit, add missing trailing lines 
* [meta] move breaking changes into dedicated file and update readme 
* [meta] centralize development and testing doc in CONTRIBUTING.md 
* [kibana] add missing values in readme 
* [kibana] move deprecated value 
* [kibana] update default values 
* [kibana] format README 
* Merge branch 'master' into master 
* [kibana] Add envFrom parameter 
* Use busybox for ES testing password generation 
* Use busybox for key generation in testing 
* Merge pull request #549 from kuisathaverat/fix-readinessProbe 

### Default value changes

```diff
diff --git a/kibana/values.yaml b/kibana/values.yaml
index 468a9853..b843b7f7 100755
--- a/kibana/values.yaml
+++ b/kibana/values.yaml
@@ -1,6 +1,4 @@
 ---
-
-elasticsearchURL: "" # "http://elasticsearch-master:9200"
 elasticsearchHosts: "http://elasticsearch-master:9200"
 
 replicas: 1
@@ -14,6 +12,13 @@ extraEnvs:
 #  - name: MY_ENVIRONMENT_VAR
 #    value: the_value_goes_here
 
+# Allows you to load environment variables from kubernetes secret or config map
+envFrom: []
+# - secretRef:
+#     name: env-secret
+# - configMapRef:
+#     name: config-map
+
 # A list of secrets and their paths to mount inside the pod
 # This is useful for mounting certificates for security and for mounting
 # the X-Pack license
@@ -24,7 +29,7 @@ secretMounts: []
 #    subPath: kibana.keystore # optional
 
 image: "docker.elastic.co/kibana/kibana"
-imageTag: "7.6.2"
+imageTag: "7.7.0"
 imagePullPolicy: "IfNotPresent"
 
 # additionals labels
@@ -137,3 +142,6 @@ lifecycle: {}
   # postStart:
   #   exec:
   #     command: ["/bin/sh", "-c", "echo Hello from the postStart handler > /usr/share/message"]
+
+# Deprecated - use only with versions < 6.6
+elasticsearchURL: "" # "http://elasticsearch-master:9200"
```

## 7.6.2 

**Release date:** 2020-03-31

![AppVersion: 7.6.2](https://img.shields.io/static/v1?label=AppVersion&message=7.6.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [7.6.2] bump version 
* fix: allow redirection on the readinessProbe 
* Merge pull request #540 from jmlrt/kibana-memory 
* [kibana] update default values 
* [kibana] set cpu request = limits 
* [kibana] increase nodejs memory limit 
* [kibana] increase memory requests/limits 

### Default value changes

```diff
diff --git a/kibana/values.yaml b/kibana/values.yaml
index 9b172945..468a9853 100755
--- a/kibana/values.yaml
+++ b/kibana/values.yaml
@@ -8,7 +8,9 @@ replicas: 1
 # Extra environment variables to append to this nodeGroup
 # This will be appended to the current 'env:' key. You can use any of the kubernetes env
 # syntax here
-extraEnvs: []
+extraEnvs:
+  - name: "NODE_OPTIONS"
+    value: "--max-old-space-size=1800"
 #  - name: MY_ENVIRONMENT_VAR
 #    value: the_value_goes_here
 
@@ -22,7 +24,7 @@ secretMounts: []
 #    subPath: kibana.keystore # optional
 
 image: "docker.elastic.co/kibana/kibana"
-imageTag: "7.6.1"
+imageTag: "7.6.2"
 imagePullPolicy: "IfNotPresent"
 
 # additionals labels
@@ -33,11 +35,11 @@ podAnnotations: {}
 
 resources:
   requests:
-    cpu: "100m"
-    memory: "500Mi"
+    cpu: "1000m"
+    memory: "2Gi"
   limits:
     cpu: "1000m"
-    memory: "1Gi"
+    memory: "2Gi"
 
 protocol: http
 
```

## 7.6.1 

**Release date:** 2020-03-20

![AppVersion: 7.6.1](https://img.shields.io/static/v1?label=AppVersion&message=7.6.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Merge branch 'master' of https://github.com/elastic/helm-charts 

### Default value changes

```diff
diff --git a/kibana/values.yaml b/kibana/values.yaml
index b9e71c39..9b172945 100755
--- a/kibana/values.yaml
+++ b/kibana/values.yaml
@@ -22,7 +22,7 @@ secretMounts: []
 #    subPath: kibana.keystore # optional
 
 image: "docker.elastic.co/kibana/kibana"
-imageTag: "7.6.0"
+imageTag: "7.6.1"
 imagePullPolicy: "IfNotPresent"
 
 # additionals labels
```

## 7.6.0 

**Release date:** 2020-03-19

![AppVersion: 7.6.0](https://img.shields.io/static/v1?label=AppVersion&message=7.6.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Fix linting issues 

### Default value changes

```diff
# No changes in this release
```

## 7.6.1 

**Release date:** 2020-03-04

![AppVersion: 7.6.1](https://img.shields.io/static/v1?label=AppVersion&message=7.6.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [7.6.1] bump version 
* Add tests to ensure the service selectors bind correctly to pods created by the deployment 
* Fix app label to match service if chart is deployed with an alias 

### Default value changes

```diff
diff --git a/kibana/values.yaml b/kibana/values.yaml
index b9e71c39..9b172945 100755
--- a/kibana/values.yaml
+++ b/kibana/values.yaml
@@ -22,7 +22,7 @@ secretMounts: []
 #    subPath: kibana.keystore # optional
 
 image: "docker.elastic.co/kibana/kibana"
-imageTag: "7.6.0"
+imageTag: "7.6.1"
 imagePullPolicy: "IfNotPresent"
 
 # additionals labels
```

## 7.6.0 

**Release date:** 2020-02-11

![AppVersion: 7.6.0](https://img.shields.io/static/v1?label=AppVersion&message=7.6.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [7.6.0] bump version 
* [meta] format with black python scripts 
* [kibana] fix typo 
* [kibana] add extra init containers 
* [kibana] add extracontainers 

### Default value changes

```diff
diff --git a/kibana/values.yaml b/kibana/values.yaml
index f708013e..b9e71c39 100755
--- a/kibana/values.yaml
+++ b/kibana/values.yaml
@@ -22,7 +22,7 @@ secretMounts: []
 #    subPath: kibana.keystore # optional
 
 image: "docker.elastic.co/kibana/kibana"
-imageTag: "7.5.2"
+imageTag: "7.6.0"
 imagePullPolicy: "IfNotPresent"
 
 # additionals labels
@@ -73,6 +73,16 @@ priorityClassName: ""
 
 httpPort: 5601
 
+extraContainers: ""
+# - name: dummy-init
+#   image: busybox
+#   command: ['echo', 'hey']
+
+extraInitContainers: ""
+# - name: dummy-init
+#   image: busybox
+#   command: ['echo', 'hey']
+
 updateStrategy:
   type: "Recreate"
 
```

## 7.5.2 

**Release date:** 2020-01-21

![AppVersion: 7.5.2](https://img.shields.io/static/v1?label=AppVersion&message=7.5.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [7.5.2] bump versions 
* [doc] switch relative URLs to absolute URLs 
* Merge pull request #419 from jmlrt/kibana-plugins-doc 
* Merge pull request #408 from maedadev/master 
* [Kibana] remove useless maxUnavailable in Kibana chart (#422) 
* [kibana] add doc for plugin install 

### Default value changes

```diff
diff --git a/kibana/values.yaml b/kibana/values.yaml
index 12cfc09f..f708013e 100755
--- a/kibana/values.yaml
+++ b/kibana/values.yaml
@@ -22,7 +22,7 @@ secretMounts: []
 #    subPath: kibana.keystore # optional
 
 image: "docker.elastic.co/kibana/kibana"
-imageTag: "7.5.1"
+imageTag: "7.5.2"
 imagePullPolicy: "IfNotPresent"
 
 # additionals labels
@@ -73,11 +73,6 @@ priorityClassName: ""
 
 httpPort: 5601
 
-# This is the max unavailable setting for the pod disruption budget
-# The default value of 1 will make sure that kubernetes won't allow more than 1
-# of your pods to be unavailable during maintenance
-maxUnavailable: 1
-
 updateStrategy:
   type: "Recreate"
 
```

## 7.5.1 

**Release date:** 2019-12-21

![AppVersion: 7.5.1](https://img.shields.io/static/v1?label=AppVersion&message=7.5.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Merge branch 'master' into master 

### Default value changes

```diff
diff --git a/kibana/values.yaml b/kibana/values.yaml
index cb3ee52b..12cfc09f 100755
--- a/kibana/values.yaml
+++ b/kibana/values.yaml
@@ -22,7 +22,7 @@ secretMounts: []
 #    subPath: kibana.keystore # optional
 
 image: "docker.elastic.co/kibana/kibana"
-imageTag: "7.5.0"
+imageTag: "7.5.1"
 imagePullPolicy: "IfNotPresent"
 
 # additionals labels
@@ -85,6 +85,7 @@ service:
   type: ClusterIP
   port: 5601
   nodePort: ""
+  labels: {}
   annotations: {}
     # cloud.google.com/load-balancer-type: "Internal"
     # service.beta.kubernetes.io/aws-load-balancer-internal: 0.0.0.0/0
```

## 7.5.0 

**Release date:** 2019-12-21

![AppVersion: 7.5.0](https://img.shields.io/static/v1?label=AppVersion&message=7.5.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* update README about default service values 

### Default value changes

```diff
diff --git a/kibana/values.yaml b/kibana/values.yaml
index c67ba204..cb3ee52b 100755
--- a/kibana/values.yaml
+++ b/kibana/values.yaml
@@ -91,6 +91,8 @@ service:
     # service.beta.kubernetes.io/azure-load-balancer-internal: "true"
     # service.beta.kubernetes.io/openstack-internal-load-balancer: "true"
     # service.beta.kubernetes.io/cce-load-balancer-internal-vpc: "true"
+  loadBalancerSourceRanges: []
+    # 0.0.0.0/0
 
 ingress:
   enabled: false
```

## 7.5.1 

**Release date:** 2019-12-19

![AppVersion: 7.5.1](https://img.shields.io/static/v1?label=AppVersion&message=7.5.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Merge pull request #415 from jmlrt/add-custom-labels-to-pods 

### Default value changes

```diff
diff --git a/kibana/values.yaml b/kibana/values.yaml
index 947e47dd..7ed80bc4 100755
--- a/kibana/values.yaml
+++ b/kibana/values.yaml
@@ -22,7 +22,7 @@ secretMounts: []
 #    subPath: kibana.keystore # optional
 
 image: "docker.elastic.co/kibana/kibana"
-imageTag: "7.5.0"
+imageTag: "7.5.1"
 imagePullPolicy: "IfNotPresent"
 
 # additionals labels
```

## 7.5.0 

**Release date:** 2019-12-18

![AppVersion: 7.5.0](https://img.shields.io/static/v1?label=AppVersion&message=7.5.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [kibana] apply labels to all pods 

### Default value changes

```diff
# No changes in this release
```

## 7.5.1 

**Release date:** 2019-12-18

![AppVersion: 7.5.1](https://img.shields.io/static/v1?label=AppVersion&message=7.5.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [7.5.1] bump versions 
* add test for service labels 
* default values for labels 
* add support for service labels 
* update docs for service labels 
* correct dictionary to list 
* add support for loadBalancerSourceRanges 
* Prefixed helper functions with chart name 
* [helm] make more explicit that helm 3 is not supported 

### Default value changes

```diff
diff --git a/kibana/values.yaml b/kibana/values.yaml
index c67ba204..7ed80bc4 100755
--- a/kibana/values.yaml
+++ b/kibana/values.yaml
@@ -22,7 +22,7 @@ secretMounts: []
 #    subPath: kibana.keystore # optional
 
 image: "docker.elastic.co/kibana/kibana"
-imageTag: "7.5.0"
+imageTag: "7.5.1"
 imagePullPolicy: "IfNotPresent"
 
 # additionals labels
@@ -85,6 +85,7 @@ service:
   type: ClusterIP
   port: 5601
   nodePort: ""
+  labels: {}
   annotations: {}
     # cloud.google.com/load-balancer-type: "Internal"
     # service.beta.kubernetes.io/aws-load-balancer-internal: 0.0.0.0/0
```

## 7.5.0 

**Release date:** 2019-12-02

![AppVersion: 7.5.0](https://img.shields.io/static/v1?label=AppVersion&message=7.5.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [7.5.0] release 
* bump elastic stack versions to 6.8.5 and 7.4.2 
* Merge branch 'master' into issue_267 
* update install doc 
* [helm] add some notice about tested helm version 

### Default value changes

```diff
diff --git a/kibana/values.yaml b/kibana/values.yaml
index 77ed94f4..c67ba204 100755
--- a/kibana/values.yaml
+++ b/kibana/values.yaml
@@ -22,7 +22,7 @@ secretMounts: []
 #    subPath: kibana.keystore # optional
 
 image: "docker.elastic.co/kibana/kibana"
-imageTag: "7.4.1"
+imageTag: "7.5.0"
 imagePullPolicy: "IfNotPresent"
 
 # additionals labels
```

## 7.4.1 

**Release date:** 2019-10-22

![AppVersion: 7.4.1](https://img.shields.io/static/v1?label=AppVersion&message=7.4.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* 7.4.1 release 
* [helm ] force empty string to "null" values This fix the following errors during `helm lint --strict`: ``` ==> Linting . [ERROR] templates/: render error in "elasticsearch/templates/service.yaml": template: elasticsearch/templates/service.yaml:24:14: executing "elasticsearch/templates/service.yaml" at <.Values.service.nodePort>: map has no entry for key "nodePort" ``` 
* [helm] add chart api version This is required witrh Helm 2.15.0 following https://github.com/helm/helm/pull/6180 
* [Issue #267] Support fullnameOverride in kibana, filebeat, metricbeat charts 
* [kibana] remove unused antiAffinity keys This keys are present in README and values but aren't used in the templates 

### Default value changes

```diff
diff --git a/kibana/values.yaml b/kibana/values.yaml
index 4e3668c8..77ed94f4 100755
--- a/kibana/values.yaml
+++ b/kibana/values.yaml
@@ -22,7 +22,7 @@ secretMounts: []
 #    subPath: kibana.keystore # optional
 
 image: "docker.elastic.co/kibana/kibana"
-imageTag: "7.4.0"
+imageTag: "7.4.1"
 imagePullPolicy: "IfNotPresent"
 
 # additionals labels
@@ -71,14 +71,6 @@ serviceAccount: ""
 # https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption/#priorityclass
 priorityClassName: ""
 
-# By default this will make sure two pods don't end up on the same node
-# Changing this to a region would allow you to spread pods across regions
-antiAffinityTopologyKey: "kubernetes.io/hostname"
-
-# Hard means that by default pods will only be scheduled if there are enough nodes for them
-# and that they will never end up on the same node. Setting this to soft will do this "best effort"
-antiAffinity: "hard"
-
 httpPort: 5601
 
 # This is the max unavailable setting for the pod disruption budget
@@ -92,7 +84,7 @@ updateStrategy:
 service:
   type: ClusterIP
   port: 5601
-  nodePort:
+  nodePort: ""
   annotations: {}
     # cloud.google.com/load-balancer-type: "Internal"
     # service.beta.kubernetes.io/aws-load-balancer-internal: 0.0.0.0/0
```

## 7.4.0 

**Release date:** 2019-10-01

![AppVersion: 7.4.0](https://img.shields.io/static/v1?label=AppVersion&message=7.4.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* 7.4.0 Release 
* [kibana] change min k8s version due to deployment apiVersion 
* [kibana] Add compatibility for k8s 1.16 
* Add working examples for running Elasticsearch and Kibana on ope… (#263) 

### Default value changes

```diff
diff --git a/kibana/values.yaml b/kibana/values.yaml
index 1af52b81..4e3668c8 100755
--- a/kibana/values.yaml
+++ b/kibana/values.yaml
@@ -22,7 +22,7 @@ secretMounts: []
 #    subPath: kibana.keystore # optional
 
 image: "docker.elastic.co/kibana/kibana"
-imageTag: "7.3.2"
+imageTag: "7.4.0"
 imagePullPolicy: "IfNotPresent"
 
 # additionals labels
```

## 7.3.2 

**Release date:** 2019-09-23

![AppVersion: 7.3.2](https://img.shields.io/static/v1?label=AppVersion&message=7.3.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [kibana] Allow configuring lifecycle events (#295) 
* Update kibana/examples/openshift/Makefile 

### Default value changes

```diff
diff --git a/kibana/values.yaml b/kibana/values.yaml
index cef949e6..1af52b81 100755
--- a/kibana/values.yaml
+++ b/kibana/values.yaml
@@ -22,7 +22,7 @@ secretMounts: []
 #    subPath: kibana.keystore # optional
 
 image: "docker.elastic.co/kibana/kibana"
-imageTag: "7.3.0"
+imageTag: "7.3.2"
 imagePullPolicy: "IfNotPresent"
 
 # additionals labels
@@ -34,7 +34,7 @@ podAnnotations: {}
 resources:
   requests:
     cpu: "100m"
-    memory: "500m"
+    memory: "500Mi"
   limits:
     cpu: "1000m"
     memory: "1Gi"
@@ -127,3 +127,11 @@ affinity: {}
 
 nameOverride: ""
 fullnameOverride: ""
+
+lifecycle: {}
+  # preStop:
+  #   exec:
+  #     command: ["/bin/sh", "-c", "echo Hello from the postStart handler > /usr/share/message"]
+  # postStart:
+  #   exec:
+  #     command: ["/bin/sh", "-c", "echo Hello from the postStart handler > /usr/share/message"]
```

## 7.3.0 

**Release date:** 2019-09-20

![AppVersion: 7.3.0](https://img.shields.io/static/v1?label=AppVersion&message=7.3.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Fix fsGroup and runAsUser for Kibana 
* fix "; \" when there is no additional command in the makefiles 
* [kibana] Allow configuring lifecycle events 
* 7.3.2 Release (#293) 

### Default value changes

```diff
# No changes in this release
```

## 7.3.2 

**Release date:** 2019-09-19

![AppVersion: 7.3.2](https://img.shields.io/static/v1?label=AppVersion&message=7.3.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* 7.3.2 Release 
* [kibana] Explicitly test for a 200 for readinessProbe 
* use same env variable as application (#272) 
* use same env variable as application 
* Update README.md 
* kibana: fixed bogus request of 500 millibytes mem 
* Add working examples for running Elasticsearch and Kibana on openshift 
* Clarify priorityClassName default for kibana chart 

### Default value changes

```diff
diff --git a/kibana/values.yaml b/kibana/values.yaml
index cef949e6..51a76470 100755
--- a/kibana/values.yaml
+++ b/kibana/values.yaml
@@ -22,7 +22,7 @@ secretMounts: []
 #    subPath: kibana.keystore # optional
 
 image: "docker.elastic.co/kibana/kibana"
-imageTag: "7.3.0"
+imageTag: "7.3.2"
 imagePullPolicy: "IfNotPresent"
 
 # additionals labels
@@ -34,7 +34,7 @@ podAnnotations: {}
 resources:
   requests:
     cpu: "100m"
-    memory: "500m"
+    memory: "500Mi"
   limits:
     cpu: "1000m"
     memory: "1Gi"
```

## 7.3.0 

**Release date:** 2019-07-31

![AppVersion: 7.3.0](https://img.shields.io/static/v1?label=AppVersion&message=7.3.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* 7.3.0 Release 
* [kibana] Add subPath support to secretMounts 

### Default value changes

```diff
diff --git a/kibana/values.yaml b/kibana/values.yaml
index cf08ef6f..cef949e6 100755
--- a/kibana/values.yaml
+++ b/kibana/values.yaml
@@ -16,12 +16,13 @@ extraEnvs: []
 # This is useful for mounting certificates for security and for mounting
 # the X-Pack license
 secretMounts: []
-#  - name: elastic-certificates
-#    secretName: elastic-certificates
-#    path: /usr/share/elasticsearch/config/certs
+#  - name: kibana-keystore
+#    secretName: kibana-keystore
+#    path: /usr/share/kibana/data/kibana.keystore
+#    subPath: kibana.keystore # optional
 
 image: "docker.elastic.co/kibana/kibana"
-imageTag: "7.2.0"
+imageTag: "7.3.0"
 imagePullPolicy: "IfNotPresent"
 
 # additionals labels
```

## 7.2.1-0 

**Release date:** 2019-07-19

![AppVersion: 7.2.0](https://img.shields.io/static/v1?label=AppVersion&message=7.2.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Merge pull request #225 from plumcraft/add-kibana-additionals-labels 
* Update kibana/tests/kibana_test.py 

### Default value changes

```diff
# No changes in this release
```

## 7.2.0 

**Release date:** 2019-07-17

![AppVersion: 7.2.0](https://img.shields.io/static/v1?label=AppVersion&message=7.2.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [kibana] - Add additionals labels 

### Default value changes

```diff
diff --git a/kibana/values.yaml b/kibana/values.yaml
index e34407e2..cf08ef6f 100755
--- a/kibana/values.yaml
+++ b/kibana/values.yaml
@@ -24,6 +24,9 @@ image: "docker.elastic.co/kibana/kibana"
 imageTag: "7.2.0"
 imagePullPolicy: "IfNotPresent"
 
+# additionals labels
+labels: {}
+
 podAnnotations: {}
   # iam.amazonaws.com/role: es-cluster
 
```

## 7.2.1-0 

**Release date:** 2019-07-17

![AppVersion: 7.2.0](https://img.shields.io/static/v1?label=AppVersion&message=7.2.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* 7.2.1-0 Release 
* Merge pull request #205 from elastic/alright_keep_your_secrets 
* Fixup the instructions for how to actual interact with the deployment 
* [kibana] Update healthCheckPath to mention basePath usage 
* Use generated encryption key for Kibana 
* Merge branch 'master' into feat/kibana-add-pod-annotations 
* Added kibana pod annotations 
* Update Kibana and Elasticsearch security examples to be generated 
* [kibana] Fixup security install docs 
* [kibana] Make imagePullPolicy actually do something 
* Merge pull request #194 from elastic/seven_too_oh 

### Default value changes

```diff
diff --git a/kibana/values.yaml b/kibana/values.yaml
index e1581bef..e34407e2 100755
--- a/kibana/values.yaml
+++ b/kibana/values.yaml
@@ -24,6 +24,9 @@ image: "docker.elastic.co/kibana/kibana"
 imageTag: "7.2.0"
 imagePullPolicy: "IfNotPresent"
 
+podAnnotations: {}
+  # iam.amazonaws.com/role: es-cluster
+
 resources:
   requests:
     cpu: "100m"
```

## 7.2.0 

**Release date:** 2019-07-01

![AppVersion: 7.2.0](https://img.shields.io/static/v1?label=AppVersion&message=7.2.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Release 7.2.0 
* Merge pull request #182 from hxquangnhat/patch-1 
* wrong secretName value 
* Merge pull request #172 from naseemkullah/kibana-non-root 
* Merge pull request #184 from diegofernandes/master 
* fix support wildcard tls host on ingress 
* Run as 1000 
* Always run tests against localhost or the service endpoint 
* Test kibana directly on 0.0.0.0 instead of checking the port 
* Drop support for 5.x 
* [kibana] Always set server.host to the docker default 
* [kibana] Add configurable nodePort to service spec 
* Update beta notice and add chart descriptions 

### Default value changes

```diff
diff --git a/kibana/values.yaml b/kibana/values.yaml
index d3a212f5..e1581bef 100755
--- a/kibana/values.yaml
+++ b/kibana/values.yaml
@@ -21,7 +21,7 @@ secretMounts: []
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/kibana/kibana"
-imageTag: "7.1.1"
+imageTag: "7.2.0"
 imagePullPolicy: "IfNotPresent"
 
 resources:
@@ -34,6 +34,8 @@ resources:
 
 protocol: http
 
+serverHost: "0.0.0.0"
+
 healthCheckPath: "/app/kibana"
 
 # Allows you to add any config files in /usr/share/kibana/config/
@@ -44,8 +46,17 @@ kibanaConfig: {}
 #       nestedkey: value
 
 # If Pod Security Policy in use it may be required to specify security context as well as service account
-podSecurityContext: {}
-  #runAsUser: "place the user id here"
+
+podSecurityContext:
+  fsGroup: 1000
+
+securityContext:
+  capabilities:
+    drop:
+    - ALL
+  # readOnlyRootFilesystem: true
+  runAsNonRoot: true
+  runAsUser: 1000
 
 serviceAccount: ""
 
@@ -74,6 +85,7 @@ updateStrategy:
 service:
   type: ClusterIP
   port: 5601
+  nodePort:
   annotations: {}
     # cloud.google.com/load-balancer-type: "Internal"
     # service.beta.kubernetes.io/aws-load-balancer-internal: 0.0.0.0/0
```

## 7.1.1 

**Release date:** 2019-06-07

![AppVersion: 7.1.1](https://img.shields.io/static/v1?label=AppVersion&message=7.1.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Merge pull request #151 from natebwangsut/master 
* Fix annotation unit test error 
* Fix wrong README.md service defaults 

### Default value changes

```diff
diff --git a/kibana/values.yaml b/kibana/values.yaml
index e7bfac25..d3a212f5 100755
--- a/kibana/values.yaml
+++ b/kibana/values.yaml
@@ -21,7 +21,7 @@ secretMounts: []
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/kibana/kibana"
-imageTag: "7.1.0"
+imageTag: "7.1.1"
 imagePullPolicy: "IfNotPresent"
 
 resources:
```

## 7.1.0 

**Release date:** 2019-06-06

![AppVersion: 7.1.0](https://img.shields.io/static/v1?label=AppVersion&message=7.1.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Added kibana_test for service annotation 

### Default value changes

```diff
diff --git a/kibana/values.yaml b/kibana/values.yaml
index 95e576e3..e7bfac25 100755
--- a/kibana/values.yaml
+++ b/kibana/values.yaml
@@ -74,6 +74,12 @@ updateStrategy:
 service:
   type: ClusterIP
   port: 5601
+  annotations: {}
+    # cloud.google.com/load-balancer-type: "Internal"
+    # service.beta.kubernetes.io/aws-load-balancer-internal: 0.0.0.0/0
+    # service.beta.kubernetes.io/azure-load-balancer-internal: "true"
+    # service.beta.kubernetes.io/openstack-internal-load-balancer: "true"
+    # service.beta.kubernetes.io/cce-load-balancer-internal-vpc: "true"
 
 ingress:
   enabled: false
```

## 7.1.1 

**Release date:** 2019-06-06

![AppVersion: 7.1.1](https://img.shields.io/static/v1?label=AppVersion&message=7.1.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Release 7.1.1 
* Added service annotations example 
* Added README section about k8s service 
* Update Kibana service.yaml 
* Remove fsGroup from container level security context 
* Merge pull request #138 from elastic/the_sooner_the_beta 

### Default value changes

```diff
diff --git a/kibana/values.yaml b/kibana/values.yaml
index c6f3e940..297e9df3 100755
--- a/kibana/values.yaml
+++ b/kibana/values.yaml
@@ -21,7 +21,7 @@ secretMounts: []
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/kibana/kibana"
-imageTag: "7.1.0"
+imageTag: "7.1.1"
 imagePullPolicy: "IfNotPresent"
 
 resources:
@@ -46,7 +46,6 @@ kibanaConfig: {}
 # If Pod Security Policy in use it may be required to specify security context as well as service account
 podSecurityContext: {}
   #runAsUser: "place the user id here"
-  #fsGroup: "place the group id here"
 
 serviceAccount: ""
 
```

## 7.1.0 

**Release date:** 2019-05-21

![AppVersion: 7.1.0](https://img.shields.io/static/v1?label=AppVersion&message=7.1.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* 7.1.0 release and promotion to beta status 
* Merge pull request #134 from elastic/glass_o_port 
* Merge pull request #109 from lancespeelmon/priorityClassName 
* [kibana] Explictly set the targetPort to the defined http port 
* RE: @Crazybus add a basic template test 

### Default value changes

```diff
diff --git a/kibana/values.yaml b/kibana/values.yaml
index 5f44f1cd..c6f3e940 100755
--- a/kibana/values.yaml
+++ b/kibana/values.yaml
@@ -21,7 +21,7 @@ secretMounts: []
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/kibana/kibana"
-imageTag: "7.0.1"
+imageTag: "7.1.0"
 imagePullPolicy: "IfNotPresent"
 
 resources:
```

## 7.0.1-alpha1 

**Release date:** 2019-05-08

![AppVersion: 7.0.1](https://img.shields.io/static/v1?label=AppVersion&message=7.0.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Merge branch 'master' into priorityClassName 
* RE: @Crazybus Could you also add these new value file into each charts readme 

### Default value changes

```diff
diff --git a/kibana/values.yaml b/kibana/values.yaml
index 116b626e..5f44f1cd 100755
--- a/kibana/values.yaml
+++ b/kibana/values.yaml
@@ -21,7 +21,7 @@ secretMounts: []
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/kibana/kibana"
-imageTag: "7.0.0"
+imageTag: "7.0.1"
 imagePullPolicy: "IfNotPresent"
 
 resources:
```

## 7.0.0-alpha1 

**Release date:** 2019-05-08

![AppVersion: 7.0.0](https://img.shields.io/static/v1?label=AppVersion&message=7.0.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* RE: @Crazybus I think this should be an empty string 
* Bump versions 
* WIP: Add integration tests and other tweaks for filebeat 
* Merge pull request #117 from elastic/filebeat-chart 
* Merge pull request #115 from elastic/up_and_at_them 

### Default value changes

```diff
diff --git a/kibana/values.yaml b/kibana/values.yaml
index 3261206f..116b626e 100755
--- a/kibana/values.yaml
+++ b/kibana/values.yaml
@@ -50,6 +50,10 @@ podSecurityContext: {}
 
 serviceAccount: ""
 
+# This is the PriorityClass settings as defined in
+# https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption/#priorityclass
+priorityClassName: ""
+
 # By default this will make sure two pods don't end up on the same node
 # Changing this to a region would allow you to spread pods across regions
 antiAffinityTopologyKey: "kubernetes.io/hostname"
```

## 7.0.1-alpha1 

**Release date:** 2019-05-01

![AppVersion: 7.0.1](https://img.shields.io/static/v1?label=AppVersion&message=7.0.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Release 7.0.1 and update the changelog 
* [elasticsearch] Add upgrade integration test 
* add support for k8s PriorityClass 
* Fix some incorrect es/kibana default values in docs 
* [kibana] Make health check path configurable 
* Add explicit chart version to the imageTag example 
* Match yml and yaml when bumping versions 
* Latest 5.x release is actually 5.6.16 

### Default value changes

```diff
diff --git a/kibana/values.yaml b/kibana/values.yaml
index 299f5e7c..7b79683b 100755
--- a/kibana/values.yaml
+++ b/kibana/values.yaml
@@ -21,7 +21,7 @@ secretMounts: []
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/kibana/kibana"
-imageTag: "7.0.0"
+imageTag: "7.0.1"
 imagePullPolicy: "IfNotPresent"
 
 resources:
@@ -34,6 +34,8 @@ resources:
 
 protocol: http
 
+healthCheckPath: "/app/kibana"
+
 # Allows you to add any config files in /usr/share/kibana/config/
 # such as kibana.yml
 kibanaConfig: {}
```

## 7.0.0-alpha1 

**Release date:** 2019-04-16

![AppVersion: 7.0.0](https://img.shields.io/static/v1?label=AppVersion&message=7.0.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Switch to 7.x as the default version 
* Update documentation links 
* Update examples to 7.0.0 
* Enable strict mode for helm lint test 
* Merge pull request #79 from elastic/bumper_cars 

### Default value changes

```diff
diff --git a/kibana/values.yaml b/kibana/values.yaml
index 3ee12e3c..299f5e7c 100755
--- a/kibana/values.yaml
+++ b/kibana/values.yaml
@@ -1,27 +1,27 @@
 ---
 
-elasticsearchURL: "http://elasticsearch-master:9200"
-#elasticsearchHosts: "http://elasticsearch-master:9200"
+elasticsearchURL: "" # "http://elasticsearch-master:9200"
+elasticsearchHosts: "http://elasticsearch-master:9200"
 
 replicas: 1
 
 # Extra environment variables to append to this nodeGroup
 # This will be appended to the current 'env:' key. You can use any of the kubernetes env
 # syntax here
-extraEnvs:
+extraEnvs: []
 #  - name: MY_ENVIRONMENT_VAR
 #    value: the_value_goes_here
 
 # A list of secrets and their paths to mount inside the pod
 # This is useful for mounting certificates for security and for mounting
 # the X-Pack license
-secretMounts: 
+secretMounts: []
 #  - name: elastic-certificates
 #    secretName: elastic-certificates
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/kibana/kibana"
-imageTag: "6.6.2"
+imageTag: "7.0.0"
 imagePullPolicy: "IfNotPresent"
 
 resources:
@@ -36,21 +36,21 @@ protocol: http
 
 # Allows you to add any config files in /usr/share/kibana/config/
 # such as kibana.yml
-# kibanaConfig:
+kibanaConfig: {}
 #   kibana.yml: |
 #     key:
 #       nestedkey: value
 
 # If Pod Security Policy in use it may be required to specify security context as well as service account
-#podSecurityContext:
+podSecurityContext: {}
   #runAsUser: "place the user id here"
   #fsGroup: "place the group id here"
- 
-#serviceAccount:
+
+serviceAccount: ""
 
 # By default this will make sure two pods don't end up on the same node
 # Changing this to a region would allow you to spread pods across regions
-antiAffinityTopologyKey: "kubernetes.io/hostname" 
+antiAffinityTopologyKey: "kubernetes.io/hostname"
 
 # Hard means that by default pods will only be scheduled if there are enough nodes for them
 # and that they will never end up on the same node. Setting this to soft will do this "best effort"
@@ -94,3 +94,6 @@ imagePullSecrets: []
 nodeSelector: {}
 tolerations: []
 affinity: {}
+
+nameOverride: ""
+fullnameOverride: ""
```

## 6.6.2-alpha1 

**Release date:** 2019-03-15

![AppVersion: 6.6.2](https://img.shields.io/static/v1?label=AppVersion&message=6.6.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Bump version to 6.6.2-alpha1 
* Drop snapshot qualifier from kibana test 
* Add 7.1.0 snapshot testing 
* :TableFormat 
* Clarify elasticsearchURL deprecation vs support 
* fix url index 
* Re-arrange elasticsearch url and hosts 
* PR feedback 
* Kibana 7 beta1 variable name change 
* Bump 7 alpha testing to 7 beta 1 

### Default value changes

```diff
diff --git a/kibana/values.yaml b/kibana/values.yaml
index 3def480a..3ee12e3c 100755
--- a/kibana/values.yaml
+++ b/kibana/values.yaml
@@ -1,6 +1,7 @@
 ---
 
 elasticsearchURL: "http://elasticsearch-master:9200"
+#elasticsearchHosts: "http://elasticsearch-master:9200"
 
 replicas: 1
 
@@ -20,7 +21,7 @@ secretMounts:
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/kibana/kibana"
-imageTag: "6.6.0"
+imageTag: "6.6.2"
 imagePullPolicy: "IfNotPresent"
 
 resources:
```

## 6.6.0-alpha1 

**Release date:** 2019-01-29

![AppVersion: 6.6.0](https://img.shields.io/static/v1?label=AppVersion&message=6.6.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Bump version to 6.6.0 
* Specify http/https protocol instead of trying to determine it 
* WIP: Add configmap support for Kibana 
* Add test coverage for serviceAccount and podSecurityContext 
* requested corrections 
* adding securityContext/PSP support for Kibana 
* Merge pull request #33 from elastic/shufalaughagus 

### Default value changes

```diff
diff --git a/kibana/values.yaml b/kibana/values.yaml
index 6f2d0a6a..3def480a 100755
--- a/kibana/values.yaml
+++ b/kibana/values.yaml
@@ -20,7 +20,7 @@ secretMounts:
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/kibana/kibana"
-imageTag: "6.5.4"
+imageTag: "6.6.0"
 imagePullPolicy: "IfNotPresent"
 
 resources:
@@ -31,6 +31,22 @@ resources:
     cpu: "1000m"
     memory: "1Gi"
 
+protocol: http
+
+# Allows you to add any config files in /usr/share/kibana/config/
+# such as kibana.yml
+# kibanaConfig:
+#   kibana.yml: |
+#     key:
+#       nestedkey: value
+
+# If Pod Security Policy in use it may be required to specify security context as well as service account
+#podSecurityContext:
+  #runAsUser: "place the user id here"
+  #fsGroup: "place the group id here"
+ 
+#serviceAccount:
+
 # By default this will make sure two pods don't end up on the same node
 # Changing this to a region would allow you to spread pods across regions
 antiAffinityTopologyKey: "kubernetes.io/hostname" 
@@ -53,11 +69,6 @@ service:
   type: ClusterIP
   port: 5601
 
-# Set this if you are setting server.ssl.enabled in Kibana.
-# This value is a hostname accepted by the SSL certificate provided to kibana.
-# Configuring this allows the readinessProbe to successfully check Kibana over HTTPS.
-kibanaSSLHostname: localhost
-
 ingress:
   enabled: false
   annotations: {}
```

## 6.5.4-alpha3 

**Release date:** 2019-01-14

![AppVersion: 6.5.4](https://img.shields.io/static/v1?label=AppVersion&message=6.5.4&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Switch helm chart icons to png 
* Update the regex pattern for Kibana 5.x version 
* Set proper credentials for connecting to Kibana 

### Default value changes

```diff
# No changes in this release
```

## 6.5.4-alpha1 

**Release date:** 2019-01-07

![AppVersion: 6.5.4](https://img.shields.io/static/v1?label=AppVersion&message=6.5.4&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Switch Kibana over to goss testing setup 

### Default value changes

```diff
# No changes in this release
```

## 6.5.4-alpha2 

**Release date:** 2019-01-14

![AppVersion: 6.5.4](https://img.shields.io/static/v1?label=AppVersion&message=6.5.4&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Bump version to 6.5.4-alpha2 

### Default value changes

```diff
# No changes in this release
```

## 6.5.4-alpha1 

**Release date:** 2019-01-08

![AppVersion: 6.5.4](https://img.shields.io/static/v1?label=AppVersion&message=6.5.4&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Merge pull request #30 from jethr0null/add_logos 
* Updated icon links 

### Default value changes

```diff
diff --git a/kibana/values.yaml b/kibana/values.yaml
index 62f105ae..6f2d0a6a 100755
--- a/kibana/values.yaml
+++ b/kibana/values.yaml
@@ -20,7 +20,7 @@ secretMounts:
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/kibana/kibana"
-imageTag: "6.5.0"
+imageTag: "6.5.4"
 imagePullPolicy: "IfNotPresent"
 
 resources:
@@ -53,6 +53,11 @@ service:
   type: ClusterIP
   port: 5601
 
+# Set this if you are setting server.ssl.enabled in Kibana.
+# This value is a hostname accepted by the SSL certificate provided to kibana.
+# Configuring this allows the readinessProbe to successfully check Kibana over HTTPS.
+kibanaSSLHostname: localhost
+
 ingress:
   enabled: false
   annotations: {}
```

## 6.5.0 

**Release date:** 2019-01-04

![AppVersion: 6.5.0](https://img.shields.io/static/v1?label=AppVersion&message=6.5.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Added icons to ES and Kibana Chart.yaml 
* Add Kibana 7 example and update Elasticsearch 7 to also use goss 
* Merge pull request #26 from elastic/respect_mah_securitay 
* Add https to the security example for Kibana 
* Set fullname to start with the overridden name 

### Default value changes

```diff
# No changes in this release
```

## 6.5.4-alpha1 

**Release date:** 2018-12-31

![AppVersion: 6.5.4](https://img.shields.io/static/v1?label=AppVersion&message=6.5.4&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Merge branch 'master' into issue/16 
* Add note about the setting being specifically for readinessProbe. 

### Default value changes

```diff
diff --git a/kibana/values.yaml b/kibana/values.yaml
index 135df800..6f2d0a6a 100755
--- a/kibana/values.yaml
+++ b/kibana/values.yaml
@@ -20,7 +20,7 @@ secretMounts:
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/kibana/kibana"
-imageTag: "6.5.3"
+imageTag: "6.5.4"
 imagePullPolicy: "IfNotPresent"
 
 resources:
@@ -55,6 +55,7 @@ service:
 
 # Set this if you are setting server.ssl.enabled in Kibana.
 # This value is a hostname accepted by the SSL certificate provided to kibana.
+# Configuring this allows the readinessProbe to successfully check Kibana over HTTPS.
 kibanaSSLHostname: localhost
 
 ingress:
```

## 6.5.3-alpha1 

**Release date:** 2018-12-28

![AppVersion: 6.5.3](https://img.shields.io/static/v1?label=AppVersion&message=6.5.3&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Add kibanaSSLHostname to README 

### Default value changes

```diff
diff --git a/kibana/values.yaml b/kibana/values.yaml
index f9975fb8..135df800 100755
--- a/kibana/values.yaml
+++ b/kibana/values.yaml
@@ -53,6 +53,10 @@ service:
   type: ClusterIP
   port: 5601
 
+# Set this if you are setting server.ssl.enabled in Kibana.
+# This value is a hostname accepted by the SSL certificate provided to kibana.
+kibanaSSLHostname: localhost
+
 ingress:
   enabled: false
   annotations: {}
```

## 6.5.4-alpha1 

**Release date:** 2018-12-28

![AppVersion: 6.5.4](https://img.shields.io/static/v1?label=AppVersion&message=6.5.4&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Update maintainer mailing list and bump to 6.5.4 
* Add https support to readinessProbe 

### Default value changes

```diff
diff --git a/kibana/values.yaml b/kibana/values.yaml
index f9975fb8..03f50203 100755
--- a/kibana/values.yaml
+++ b/kibana/values.yaml
@@ -20,7 +20,7 @@ secretMounts:
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/kibana/kibana"
-imageTag: "6.5.3"
+imageTag: "6.5.4"
 imagePullPolicy: "IfNotPresent"
 
 resources:
```

## 6.5.3-alpha1 

**Release date:** 2018-12-13

![AppVersion: 6.5.3](https://img.shields.io/static/v1?label=AppVersion&message=6.5.3&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Bump default version to 6.5.3 

### Default value changes

```diff
diff --git a/kibana/values.yaml b/kibana/values.yaml
index 50e4e3c8..f9975fb8 100755
--- a/kibana/values.yaml
+++ b/kibana/values.yaml
@@ -20,7 +20,7 @@ secretMounts:
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/kibana/kibana"
-imageTag: "6.5.2"
+imageTag: "6.5.3"
 imagePullPolicy: "IfNotPresent"
 
 resources:
```

## 6.5.2-alpha1 

**Release date:** 2018-12-06

![AppVersion: 6.5.2](https://img.shields.io/static/v1?label=AppVersion&message=6.5.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Bump to version 6.5.2 and append the alpha1 release tag 
* Update Kibana README 
* Switch Kibana to use the recreate strategy by default 

### Default value changes

```diff
diff --git a/kibana/values.yaml b/kibana/values.yaml
index 32a2794c..50e4e3c8 100755
--- a/kibana/values.yaml
+++ b/kibana/values.yaml
@@ -20,7 +20,7 @@ secretMounts:
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/kibana/kibana"
-imageTag: "6.5.0"
+imageTag: "6.5.2"
 imagePullPolicy: "IfNotPresent"
 
 resources:
@@ -46,6 +46,9 @@ httpPort: 5601
 # of your pods to be unavailable during maintenance
 maxUnavailable: 1
 
+updateStrategy:
+  type: "Recreate"
+
 service:
   type: ClusterIP
   port: 5601
```

## 6.5.0 

**Release date:** 2018-11-20

![AppVersion: 6.5.0](https://img.shields.io/static/v1?label=AppVersion&message=6.5.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Bump default version to 6.5.0 

### Default value changes

```diff
diff --git a/kibana/values.yaml b/kibana/values.yaml
index 99dbcd23..32a2794c 100755
--- a/kibana/values.yaml
+++ b/kibana/values.yaml
@@ -20,7 +20,7 @@ secretMounts:
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/kibana/kibana"
-imageTag: "6.4.3"
+imageTag: "6.5.0"
 imagePullPolicy: "IfNotPresent"
 
 resources:
```

## 6.4.3 

**Release date:** 2018-11-13

![AppVersion: 6.4.3](https://img.shields.io/static/v1?label=AppVersion&message=6.4.3&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* first commit 

### Default value changes

```diff
---

elasticsearchURL: "http://elasticsearch-master:9200"

replicas: 1

# Extra environment variables to append to this nodeGroup
# This will be appended to the current 'env:' key. You can use any of the kubernetes env
# syntax here
extraEnvs:
#  - name: MY_ENVIRONMENT_VAR
#    value: the_value_goes_here

# A list of secrets and their paths to mount inside the pod
# This is useful for mounting certificates for security and for mounting
# the X-Pack license
secretMounts: 
#  - name: elastic-certificates
#    secretName: elastic-certificates
#    path: /usr/share/elasticsearch/config/certs

image: "docker.elastic.co/kibana/kibana"
imageTag: "6.4.3"
imagePullPolicy: "IfNotPresent"

resources:
  requests:
    cpu: "100m"
    memory: "500m"
  limits:
    cpu: "1000m"
    memory: "1Gi"

# By default this will make sure two pods don't end up on the same node
# Changing this to a region would allow you to spread pods across regions
antiAffinityTopologyKey: "kubernetes.io/hostname" 

# Hard means that by default pods will only be scheduled if there are enough nodes for them
# and that they will never end up on the same node. Setting this to soft will do this "best effort"
antiAffinity: "hard"

httpPort: 5601

# This is the max unavailable setting for the pod disruption budget
# The default value of 1 will make sure that kubernetes won't allow more than 1
# of your pods to be unavailable during maintenance
maxUnavailable: 1

service:
  type: ClusterIP
  port: 5601

ingress:
  enabled: false
  annotations: {}
    # kubernetes.io/ingress.class: nginx
    # kubernetes.io/tls-acme: "true"
  path: /
  hosts:
    - chart-example.local
  tls: []
  #  - secretName: chart-example-tls
  #    hosts:
  #      - chart-example.local

readinessProbe:
  failureThreshold: 3
  initialDelaySeconds: 10
  periodSeconds: 10
  successThreshold: 3
  timeoutSeconds: 5

imagePullSecrets: []
nodeSelector: {}
tolerations: []
affinity: {}
```

---
Autogenerated from Helm Chart and git history using [helm-changelog](https://github.com/mogensen/helm-changelog)
