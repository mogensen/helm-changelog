# Change Log

## Next Release 

![AppVersion: 8.0.0-SNAPSHOT](https://img.shields.io/static/v1?label=AppVersion&message=8.0.0-SNAPSHOT&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [elasticsearch] only configure ES_JAVA_OPTS when value is set (#1089) 
* [elasticsearch] Mark esMajorVersion as deprecated (#1109) 
* [elasticsearch] fix network policies http additional rules (#1111) 
* [elasticsearch] Add namespace to helm test command in NOTES.txt (#1105) 
* [elasticsearch][kibana] remove oss examples (#1046) 
* fix: security.yaml is 404 (#1040) 
* [meta] add build status and artifact hub badges (#1033) 
* [elasticsearch] fix secrets in config example (#1012) 
* Add support for NetworkPolicy. (#498) 
* [elasticsearch][kibana] Add flexible ingress (#994) 
* [all] add hostaliases (#970) 
* elasticsearch: add emptyDir to podSecurityPolicy as allowed volume-type (#975) 
* [meta] stabilize CI tests (#935) 
* ES Statefulset empty initContainers fix (#795) 
* [meta] remove support for k8s <1.14 & helm <2.17.0 (#916) 
* [meta] upgrade test (#907) 
* [elasticsearch] Fix spelling (#897) 
* [elasticsearch] update test hook annotations (#911) 
* Helm 3 (#516) 
* [meta] increase helm timeout (#891) 
* Add warning comment placeholder (main branch) (#886) 
* [elasticsearch] add coordinator node to multi test (#854) 
* Update warning on charts readme (#863) 
* [elasticsearch][kibana] disable nss dentry cache (#818) 
* [meta] update changelog for 7.9.2 release  (#824) 
* [elasticsearch] fix secrets names in examples (#811) 
* Include pre-releases in the semver range. (#729) 
* [elasticsearch] add loadBalancer externalTrafficPolicy option (#810) 
* Missing deletion of "elastic-certificate-crt" (#752) 
* [meta] Update changelog for 6.8.12 / 7.9.0 releases (#789) 
* [meta] add helm 3 beta support (#759) 
* [doc] update doc links (#758) 
* [meta] update changelog for 7.8.1 and 6.8.11 releases (#755) 
* [doc] fix some links (#737) 
* [elasticsearch] Update test image pull policy. (#727) 
* [elasticsearch] _helpers.tpl - elasticsearch.endpoints to use elasticsearch.uname (#670) 

### Default value changes

```diff
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index 42cefab1..9ef799b8 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -12,6 +12,8 @@ roles:
   master: "true"
   ingest: "true"
   data: "true"
+  remote_cluster_client: "true"
+  ml: "true"
 
 replicas: 3
 minimumMasterNodes: 2
@@ -50,6 +52,12 @@ secretMounts: []
 #    path: /usr/share/elasticsearch/config/certs
 #    defaultMode: 0755
 
+hostAliases: []
+#- ip: "127.0.0.1"
+#  hostnames:
+#  - "foo.local"
+#  - "bar.local"
+
 image: "docker.elastic.co/elasticsearch/elasticsearch"
 imageTag: "8.0.0-SNAPSHOT"
 imagePullPolicy: "IfNotPresent"
@@ -60,7 +68,7 @@ podAnnotations: {}
 # additionals labels
 labels: {}
 
-esJavaOpts: "-Xmx1g -Xms1g"
+esJavaOpts: "" # example: "-Xmx1g -Xms1g"
 
 resources:
   requests:
@@ -116,11 +124,12 @@ podSecurityPolicy:
       - secret
       - configMap
       - persistentVolumeClaim
+      - emptyDir
 
 persistence:
   enabled: true
   labels:
-    # Add default labels for the volumeClaimTemplate fo the StatefulSet
+    # Add default labels for the volumeClaimTemplate of the StatefulSet
     enabled: false
   annotations: {}
 
@@ -182,6 +191,7 @@ service:
   transportPortName: transport
   loadBalancerIP: ""
   loadBalancerSourceRanges: []
+  externalTrafficPolicy: ""
 
 updateStrategy: RollingUpdate
 
@@ -233,9 +243,10 @@ ingress:
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
@@ -272,6 +283,61 @@ sysctlInitContainer:
 
 keystore: []
 
+networkPolicy:
+  ## Enable creation of NetworkPolicy resources. Only Ingress traffic is filtered for now.
+  ## In order for a Pod to access Elasticsearch, it needs to have the following label:
+  ## {{ template "uname" . }}-client: "true"
+  ## Example for default configuration to access HTTP port:
+  ## elasticsearch-master-http-client: "true"
+  ## Example for default configuration to access transport port:
+  ## elasticsearch-master-transport-client: "true"
+
+  http:
+    enabled: false
+    ## if explicitNamespacesSelector is not set or set to {}, only client Pods being in the networkPolicy's namespace
+    ## and matching all criteria can reach the DB.
+    ## But sometimes, we want the Pods to be accessible to clients from other namespaces, in this case, we can use this
+    ## parameter to select these namespaces
+    ##
+    # explicitNamespacesSelector:
+    #   # Accept from namespaces with all those different rules (only from whitelisted Pods)
+    #   matchLabels:
+    #     role: frontend
+    #   matchExpressions:
+    #     - {key: role, operator: In, values: [frontend]}
+
+    ## Additional NetworkPolicy Ingress "from" rules to set. Note that all rules are OR-ed.
+    ##
+    # additionalRules:
+    #   - podSelector:
+    #       matchLabels:
+    #         role: frontend
+    #   - podSelector:
+    #       matchExpressions:
+    #         - key: role
+    #           operator: In
+    #           values:
+    #             - frontend
+
+  transport:
+    ## Note that all Elasticsearch Pods can talks to themselves using transport port even if enabled.
+    enabled: false
+    # explicitNamespacesSelector:
+    #   matchLabels:
+    #     role: frontend
+    #   matchExpressions:
+    #     - {key: role, operator: In, values: [frontend]}
+    # additionalRules:
+    #   - podSelector:
+    #       matchLabels:
+    #         role: frontend
+    #   - podSelector:
+    #       matchExpressions:
+    #         - key: role
+    #           operator: In
+    #           values:
+    #             - frontend
+
 # Deprecated
 # please use the above podSecurityContext.fsGroup instead
 fsGroup: ""
```

## 8.0.0-SNAPSHOT 

**Release date:** 2020-06-29

![AppVersion: 8.0.0-SNAPSHOT](https://img.shields.io/static/v1?label=AppVersion&message=8.0.0-SNAPSHOT&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Bump master branch to 8.0.0-SNAPSHOT (#682) 
* Elasticsearch: set PVC labels through setting all StatefulSet labels to its volumeClaimTemplate (#665) 
* Add ServiceAccount annotations (#686) 
* podSecurityContext.runAsUser needs to be nulled as well for Openshift (#655) 
* [logstash] add security example (#392) 
* [elasticsearch] update kind example for version >= 0.7.0 (#669) 
* [elasticsearch] Disable service links to prevent very long startup times (#542) 
* Elasticsearch: do not include heritage selector (#437) 

### Default value changes

```diff
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index f5a162b9..42cefab1 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -51,7 +51,7 @@ secretMounts: []
 #    defaultMode: 0755
 
 image: "docker.elastic.co/elasticsearch/elasticsearch"
-imageTag: "7.7.1"
+imageTag: "8.0.0-SNAPSHOT"
 imagePullPolicy: "IfNotPresent"
 
 podAnnotations: {}
@@ -96,6 +96,7 @@ volumeClaimTemplate:
 
 rbac:
   create: false
+  serviceAccountAnnotations: {}
   serviceAccountName: ""
 
 podSecurityPolicy:
@@ -118,6 +119,9 @@ podSecurityPolicy:
 
 persistence:
   enabled: true
+  labels:
+    # Add default labels for the volumeClaimTemplate fo the StatefulSet
+    enabled: false
   annotations: {}
 
 extraVolumes: []
@@ -159,6 +163,11 @@ nodeAffinity: {}
 # the same time when bootstrapping the cluster
 podManagementPolicy: "Parallel"
 
+# The environment variables injected by service links are not used, but can lead to slow Elasticsearch boot times when
+# there are many services in the current namespace.
+# If you experience slow pod startups you probably want to set this to `false`.
+enableServiceLinks: true
+
 protocol: http
 httpPort: 9200
 transportPort: 9300
```

## 7.7.1 

**Release date:** 2020-06-03

![AppVersion: 7.7.1](https://img.shields.io/static/v1?label=AppVersion&message=7.7.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Update changelog for 7.7.1 and 6.8.10 releases (#652) 
* [meta] add support for k8s 1.16 (#635) 
* Elasticsearch secret mountmode (#596) 
* [elasticsearch] Fix issue with `readinessProbe` causing outages (#638) 
* Fix values links in examples/multi/README.md (#639) 

### Default value changes

```diff
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index 0029b992..f5a162b9 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -48,9 +48,10 @@ secretMounts: []
 #  - name: elastic-certificates
 #    secretName: elastic-certificates
 #    path: /usr/share/elasticsearch/config/certs
+#    defaultMode: 0755
 
 image: "docker.elastic.co/elasticsearch/elasticsearch"
-imageTag: "7.7.0"
+imageTag: "7.7.1"
 imagePullPolicy: "IfNotPresent"
 
 podAnnotations: {}
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
* [elasticseach] move deprecated value 
* [elasticsearch] update default values 
* [elasticsearch] format examples README 
* [elasticsearch] format README 
* Merge branch 'master' into master 
* [elasticsearch] Add envFrom parameter 
* Merge pull request #586 from jmlrt/readiness-to-503 
* Fix linting errors 
* [elasticsearch] don't print BASIC_AUTH in k8s event log 
* fixup! [elasticsearch] split readiness http code check condition 
* [elasticsearch] exit with return code from curl 
* Updated tests to test both string and yaml 
* [elasticsearch] split readiness http code check condition 
* Added ability for string or yaml for extra* templates 
* Merge branch 'master' of github.com:elastic/helm-charts into helm-default-patch 
* Merge pull request #584 from michelesr/master 
* [elasticsearch] Adds imagePullSecrets for test Pod 
* [elasticsearch] update readiness probe endpoint 
* [elasticsearch] Set securityContext for test pod 
* Use busybox for ES testing password generation 
* [elasticsearch] fix helm test command 

### Default value changes

```diff
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index 4486d53c..0029b992 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -34,6 +34,13 @@ extraEnvs: []
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
@@ -43,7 +50,7 @@ secretMounts: []
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/elasticsearch/elasticsearch"
-imageTag: "7.6.2"
+imageTag: "7.7.0"
 imagePullPolicy: "IfNotPresent"
 
 podAnnotations: {}
@@ -112,21 +119,21 @@ persistence:
   enabled: true
   annotations: {}
 
-extraVolumes: ""
+extraVolumes: []
   # - name: extras
   #   emptyDir: {}
 
-extraVolumeMounts: ""
+extraVolumeMounts: []
   # - name: extras
   #   mountPath: /usr/share/extras
   #   readOnly: true
 
-extraContainers: ""
+extraContainers: []
   # - name: do-something
   #   image: busybox
   #   command: ['do', 'something']
 
-extraInitContainers: ""
+extraInitContainers: []
   # - name: do-something
   #   image: busybox
   #   command: ['do', 'something']
@@ -177,10 +184,6 @@ podSecurityContext:
   fsGroup: 1000
   runAsUser: 1000
 
-# The following value is deprecated,
-# please use the above podSecurityContext.fsGroup instead
-fsGroup: ""
-
 securityContext:
   capabilities:
     drop:
@@ -258,3 +261,7 @@ sysctlInitContainer:
   enabled: true
 
 keystore: []
+
+# Deprecated
+# please use the above podSecurityContext.fsGroup instead
+fsGroup: ""
```

## 7.6.2 

**Release date:** 2020-03-31

![AppVersion: 7.6.2](https://img.shields.io/static/v1?label=AppVersion&message=7.6.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [7.6.2] bump version 
* Support staging ES security tests 
* added loadBalancerIP option to service 
* Update defaults for extra values to support lists 
* Add namespace parameter to the test function to NOTES.txt 
* Add possibility to define custom readinessProbe (#485) 

### Default value changes

```diff
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index 053c0209..4486d53c 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -43,7 +43,7 @@ secretMounts: []
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/elasticsearch/elasticsearch"
-imageTag: "7.6.1"
+imageTag: "7.6.2"
 imagePullPolicy: "IfNotPresent"
 
 podAnnotations: {}
@@ -163,6 +163,7 @@ service:
   annotations: {}
   httpPortName: http
   transportPortName: transport
+  loadBalancerIP: ""
   loadBalancerSourceRanges: []
 
 updateStrategy: RollingUpdate
```

## 7.6.1 

**Release date:** 2020-03-04

![AppVersion: 7.6.1](https://img.shields.io/static/v1?label=AppVersion&message=7.6.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [7.6.1] bump version 
* [elasticsearch] fix format of test for loadBalancerSourceRanges 
* correct dictionary tol list 
* [elasticsearch] add support for loadBalancerSourceRanges 
* [elasticsearch] add notice about cpu request change merged in c56682285e6d47e5fed9e000b045ec98ee1d3555 

### Default value changes

```diff
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index cad0b582..053c0209 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -43,7 +43,7 @@ secretMounts: []
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/elasticsearch/elasticsearch"
-imageTag: "7.6.0"
+imageTag: "7.6.1"
 imagePullPolicy: "IfNotPresent"
 
 podAnnotations: {}
@@ -163,6 +163,7 @@ service:
   annotations: {}
   httpPortName: http
   transportPortName: transport
+  loadBalancerSourceRanges: []
 
 updateStrategy: RollingUpdate
 
```

## 7.6.0 

**Release date:** 2020-02-11

![AppVersion: 7.6.0](https://img.shields.io/static/v1?label=AppVersion&message=7.6.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [7.6.0] bump version 
* [meta] format with black python scripts 
* [elasticsearch] add extracontainers 
* Merge pull request #458 from jmlrt/request-equals-limits 
* [elasticsearch] set cpu request = cpu limit 
* Fixing typo 

### Default value changes

```diff
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index 9cd502d7..cad0b582 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -43,7 +43,7 @@ secretMounts: []
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/elasticsearch/elasticsearch"
-imageTag: "7.5.2"
+imageTag: "7.6.0"
 imagePullPolicy: "IfNotPresent"
 
 podAnnotations: {}
@@ -56,7 +56,7 @@ esJavaOpts: "-Xmx1g -Xms1g"
 
 resources:
   requests:
-    cpu: "100m"
+    cpu: "1000m"
     memory: "2Gi"
   limits:
     cpu: "1000m"
@@ -121,6 +121,11 @@ extraVolumeMounts: ""
   #   mountPath: /usr/share/extras
   #   readOnly: true
 
+extraContainers: ""
+  # - name: do-something
+  #   image: busybox
+  #   command: ['do', 'something']
+
 extraInitContainers: ""
   # - name: do-something
   #   image: busybox
```

## 7.5.2 

**Release date:** 2020-01-21

![AppVersion: 7.5.2](https://img.shields.io/static/v1?label=AppVersion&message=7.5.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [7.5.2] bump versions 
* Add commented out example of a useful post start hook 
* [doc] switch relative URLs to absolute URLs 
* [elasticsearch]  add workaround to fix kind example 
* Merge pull request #433 from jmlrt/elasticsearch-microk8s 
* [elasticsearch] add example for microk8s 
* duplicate label removed 
* Merge pull request #382 from jaumann/es-nameoverride 

### Default value changes

```diff
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index c667444f..9cd502d7 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -43,7 +43,7 @@ secretMounts: []
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/elasticsearch/elasticsearch"
-imageTag: "7.5.1"
+imageTag: "7.5.2"
 imagePullPolicy: "IfNotPresent"
 
 podAnnotations: {}
@@ -233,7 +233,19 @@ lifecycle: {}
   #     command: ["/bin/sh", "-c", "echo Hello from the postStart handler > /usr/share/message"]
   # postStart:
   #   exec:
-  #     command: ["/bin/sh", "-c", "echo Hello from the postStart handler > /usr/share/message"]
+  #     command:
+  #       - bash
+  #       - -c
+  #       - |
+  #         #!/bin/bash
+  #         # Add a template to adjust number of shards/replicas
+  #         TEMPLATE_NAME=my_template
+  #         INDEX_PATTERN="logstash-*"
+  #         SHARD_COUNT=8
+  #         REPLICA_COUNT=1
+  #         ES_URL=http://localhost:9200
+  #         while [[ "$(curl -s -o /dev/null -w '%{http_code}\n' $ES_URL)" != "200" ]]; do sleep 1; done
+  #         curl -XPUT "$ES_URL/_template/$TEMPLATE_NAME" -H 'Content-Type: application/json' -d'{"index_patterns":['\""$INDEX_PATTERN"\"'],"settings":{"number_of_shards":'$SHARD_COUNT',"number_of_replicas":'$REPLICA_COUNT'}}'
 
 sysctlInitContainer:
   enabled: true
```

## 7.5.1 

**Release date:** 2019-12-18

![AppVersion: 7.5.1](https://img.shields.io/static/v1?label=AppVersion&message=7.5.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [7.5.1] bump versions 
* Fix misnamed masterService 
* Remove merge conflict 
* Allow for name overrides of resources 
* Prefixed helper functions with chart name 
* [helm] make more explicit that helm 3 is not supported 

### Default value changes

```diff
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index 8f2d943a..c667444f 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -43,7 +43,7 @@ secretMounts: []
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/elasticsearch/elasticsearch"
-imageTag: "7.5.0"
+imageTag: "7.5.1"
 imagePullPolicy: "IfNotPresent"
 
 podAnnotations: {}
```

## 7.5.0 

**Release date:** 2019-12-02

![AppVersion: 7.5.0](https://img.shields.io/static/v1?label=AppVersion&message=7.5.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [7.5.0] release 
* bump elastic stack versions to 6.8.5 and 7.4.2 
* [elasticsearch] Apply labels to all pods 
* [elasticsearch] Tweak the 'readinessProbe' command to verify that master nodes are available. 
* update install doc 
* use same imagePullPolicy in initContainer 
* add tests 
* update readme 
* add support for labels for services 
* [helm] add some notice about tested helm version 

### Default value changes

```diff
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index 72bf2b47..8f2d943a 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -43,7 +43,7 @@ secretMounts: []
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/elasticsearch/elasticsearch"
-imageTag: "7.4.1"
+imageTag: "7.5.0"
 imagePullPolicy: "IfNotPresent"
 
 podAnnotations: {}
@@ -151,6 +151,8 @@ httpPort: 9200
 transportPort: 9300
 
 service:
+  labels: {}
+  labelsHeadless: {}
   type: ClusterIP
   nodePort: ""
   annotations: {}
```

## 7.4.1 

**Release date:** 2019-10-22

![AppVersion: 7.4.1](https://img.shields.io/static/v1?label=AppVersion&message=7.4.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* 7.4.1 release 
* [helm] workaround to adress regression for numeric value type in helm v2.15 - Apply workaround from https://github.com/helm/helm/issues/6708 - Related to https://github.com/helm/helm/pull/6010 
* [helm ] force empty string to "null" values This fix the following errors during `helm lint --strict`: ``` ==> Linting . [ERROR] templates/: render error in "elasticsearch/templates/service.yaml": template: elasticsearch/templates/service.yaml:24:14: executing "elasticsearch/templates/service.yaml" at <.Values.service.nodePort>: map has no entry for key "nodePort" ``` 
* [helm] add chart api version This is required witrh Helm 2.15.0 following https://github.com/helm/helm/pull/6180 
* Merge pull request #337 from jmlrt/es-unused-default-value 
* [elasticsearch] remove unused default value 
* Update elasticsearch/README.md 
* [elasticsearch] fix deprecated note related to https://github.com/elastic/helm-charts/pull/333/files#r335988708 
* Merge pull request #274 from salaboy/master 
* [elasticsearch] Add logging when adding password to keystore 

### Default value changes

```diff
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index eff281ba..72bf2b47 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -43,7 +43,7 @@ secretMounts: []
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/elasticsearch/elasticsearch"
-imageTag: "7.4.0"
+imageTag: "7.4.1"
 imagePullPolicy: "IfNotPresent"
 
 podAnnotations: {}
@@ -152,7 +152,7 @@ transportPort: 9300
 
 service:
   type: ClusterIP
-  nodePort:
+  nodePort: ""
   annotations: {}
   httpPortName: http
   transportPortName: transport
```

## 7.4.0 

**Release date:** 2019-10-01

![AppVersion: 7.4.0](https://img.shields.io/static/v1?label=AppVersion&message=7.4.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* 7.4.0 Release 
* Merge pull request #302 from code-chris/elastic-k8s-1.16 
* [elasticsearch] Add compatibility for k8s 1.16 

### Default value changes

```diff
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index ccab4623..eff281ba 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -43,7 +43,7 @@ secretMounts: []
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/elasticsearch/elasticsearch"
-imageTag: "7.3.2"
+imageTag: "7.4.0"
 imagePullPolicy: "IfNotPresent"
 
 podAnnotations: {}
```

## 7.3.2 

**Release date:** 2019-09-27

![AppVersion: 7.3.2](https://img.shields.io/static/v1?label=AppVersion&message=7.3.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Fix bug when checking if a variable is set. 

### Default value changes

```diff
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index 7ce13c31..ccab4623 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -43,7 +43,7 @@ secretMounts: []
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/elasticsearch/elasticsearch"
-imageTag: "7.3.0"
+imageTag: "7.3.2"
 imagePullPolicy: "IfNotPresent"
 
 podAnnotations: {}
@@ -86,6 +86,28 @@ volumeClaimTemplate:
     requests:
       storage: 30Gi
 
+rbac:
+  create: false
+  serviceAccountName: ""
+
+podSecurityPolicy:
+  create: false
+  name: ""
+  spec:
+    privileged: true
+    fsGroup:
+      rule: RunAsAny
+    runAsUser:
+      rule: RunAsAny
+    seLinux:
+      rule: RunAsAny
+    supplementalGroups:
+      rule: RunAsAny
+    volumes:
+      - secret
+      - configMap
+      - persistentVolumeClaim
+
 persistence:
   enabled: true
   annotations: {}
@@ -132,6 +154,8 @@ service:
   type: ClusterIP
   nodePort:
   annotations: {}
+  httpPortName: http
+  transportPortName: transport
 
 updateStrategy: RollingUpdate
 
@@ -142,6 +166,7 @@ maxUnavailable: 1
 
 podSecurityContext:
   fsGroup: 1000
+  runAsUser: 1000
 
 # The following value is deprecated,
 # please use the above podSecurityContext.fsGroup instead
```

## 7.3.0 

**Release date:** 2019-09-24

![AppVersion: 7.3.0](https://img.shields.io/static/v1?label=AppVersion&message=7.3.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Update values.yaml 
* Add working examples for running Elasticsearch and Kibana on ope… (#263) 

### Default value changes

```diff
# No changes in this release
```

## 7.3.2 

**Release date:** 2019-09-23

![AppVersion: 7.3.2](https://img.shields.io/static/v1?label=AppVersion&message=7.3.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Make it possible to override the endpoint template. (#298) 

### Default value changes

```diff
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index 7ce13c31..ccab4623 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -43,7 +43,7 @@ secretMounts: []
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/elasticsearch/elasticsearch"
-imageTag: "7.3.0"
+imageTag: "7.3.2"
 imagePullPolicy: "IfNotPresent"
 
 podAnnotations: {}
@@ -86,6 +86,28 @@ volumeClaimTemplate:
     requests:
       storage: 30Gi
 
+rbac:
+  create: false
+  serviceAccountName: ""
+
+podSecurityPolicy:
+  create: false
+  name: ""
+  spec:
+    privileged: true
+    fsGroup:
+      rule: RunAsAny
+    runAsUser:
+      rule: RunAsAny
+    seLinux:
+      rule: RunAsAny
+    supplementalGroups:
+      rule: RunAsAny
+    volumes:
+      - secret
+      - configMap
+      - persistentVolumeClaim
+
 persistence:
   enabled: true
   annotations: {}
@@ -132,6 +154,8 @@ service:
   type: ClusterIP
   nodePort:
   annotations: {}
+  httpPortName: http
+  transportPortName: transport
 
 updateStrategy: RollingUpdate
 
@@ -142,6 +166,7 @@ maxUnavailable: 1
 
 podSecurityContext:
   fsGroup: 1000
+  runAsUser: 1000
 
 # The following value is deprecated,
 # please use the above podSecurityContext.fsGroup instead
```

## 7.3.0 

**Release date:** 2019-09-23

![AppVersion: 7.3.0](https://img.shields.io/static/v1?label=AppVersion&message=7.3.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* updating storageClassName 
* Make it possible to override the endpoint template. 
* fix "; \" when there is no additional command in the makefiles 

### Default value changes

```diff
# No changes in this release
```

## 7.3.2 

**Release date:** 2019-09-19

![AppVersion: 7.3.2](https://img.shields.io/static/v1?label=AppVersion&message=7.3.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* 7.3.2 Release 
* Hardening of the pod permissions. (#265) 
* [elasticsearch] Set default runAsUser for pod security context (#259) 
* Merge pull request #286 from NorDroN/fix-certs 
* Merge branch 'master' into podsecuritypolicy 
* [elasticsearch] Set default runAsUser for pod security context 
* Fix test following removal of some permissions in the PSP spec. 
* Changes based on feedback. 
* ES Variable Port Name (#270) 
* Remove unnecessary semicolon and newline escape 
* Update values file and added templates to both services 
* [elasticsearch] Drop version from chart label in service 
* Fix 'Certificate has no `subjectAltName`, falling back to check for a `commonName` for now. This feature is being removed by major browsers and deprecated by RFC 2818' 
* adding Example for Kubernetes KIND 
* ES Variable Port Name 
* Hardening of the pod permissions. 
* Properly quote bootstrap password 
* [elasticsearch] Keystore integration 
* Add working examples for running Elasticsearch and Kibana on openshift 
* Clarify priorityClassName default for es chart 
* Fixed indent on elasticsearch extraVolumes tpl. Was causing parsing errors 
* Update documentation and defaults for tmpl values 

### Default value changes

```diff
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index 13ca0626..ccab4623 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -43,7 +43,7 @@ secretMounts: []
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/elasticsearch/elasticsearch"
-imageTag: "7.3.0"
+imageTag: "7.3.2"
 imagePullPolicy: "IfNotPresent"
 
 podAnnotations: {}
@@ -86,20 +86,42 @@ volumeClaimTemplate:
     requests:
       storage: 30Gi
 
+rbac:
+  create: false
+  serviceAccountName: ""
+
+podSecurityPolicy:
+  create: false
+  name: ""
+  spec:
+    privileged: true
+    fsGroup:
+      rule: RunAsAny
+    runAsUser:
+      rule: RunAsAny
+    seLinux:
+      rule: RunAsAny
+    supplementalGroups:
+      rule: RunAsAny
+    volumes:
+      - secret
+      - configMap
+      - persistentVolumeClaim
+
 persistence:
   enabled: true
   annotations: {}
 
-extraVolumes: []
+extraVolumes: ""
   # - name: extras
   #   emptyDir: {}
 
-extraVolumeMounts: []
+extraVolumeMounts: ""
   # - name: extras
   #   mountPath: /usr/share/extras
   #   readOnly: true
 
-extraInitContainers: []
+extraInitContainers: ""
   # - name: do-something
   #   image: busybox
   #   command: ['do', 'something']
@@ -132,6 +154,8 @@ service:
   type: ClusterIP
   nodePort:
   annotations: {}
+  httpPortName: http
+  transportPortName: transport
 
 updateStrategy: RollingUpdate
 
@@ -142,6 +166,7 @@ maxUnavailable: 1
 
 podSecurityContext:
   fsGroup: 1000
+  runAsUser: 1000
 
 # The following value is deprecated,
 # please use the above podSecurityContext.fsGroup instead
@@ -210,3 +235,5 @@ lifecycle: {}
 
 sysctlInitContainer:
   enabled: true
+
+keystore: []
```

## 7.3.0 

**Release date:** 2019-07-31

![AppVersion: 7.3.0](https://img.shields.io/static/v1?label=AppVersion&message=7.3.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* 7.3.0 Release 
* publishNotReadyAddresses not working 

### Default value changes

```diff
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index 9caf0bcc..13ca0626 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -43,7 +43,7 @@ secretMounts: []
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/elasticsearch/elasticsearch"
-imageTag: "7.2.0"
+imageTag: "7.3.0"
 imagePullPolicy: "IfNotPresent"
 
 podAnnotations: {}
```

## 7.2.1-0 

**Release date:** 2019-07-18

![AppVersion: 7.2.0](https://img.shields.io/static/v1?label=AppVersion&message=7.2.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Merge pull request #226 from MichaelMarieJulie/add-elasticsearch-additionals-labels 

### Default value changes

```diff
# No changes in this release
```

## 7.2.0 

**Release date:** 2019-07-17

![AppVersion: 7.2.0](https://img.shields.io/static/v1?label=AppVersion&message=7.2.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [elasticsearch] additionals labels 

### Default value changes

```diff
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index 6e53ee60..9caf0bcc 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -49,6 +49,9 @@ imagePullPolicy: "IfNotPresent"
 podAnnotations: {}
   # iam.amazonaws.com/role: es-cluster
 
+# additionals labels
+labels: {}
+
 esJavaOpts: "-Xmx1g -Xms1g"
 
 resources:
```

## 7.2.1-0 

**Release date:** 2019-07-17

![AppVersion: 7.2.0](https://img.shields.io/static/v1?label=AppVersion&message=7.2.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* 7.2.1-0 Release 
* Merge pull request #205 from elastic/alright_keep_your_secrets 
* Add tests and fix README.md 
* Fixup the instructions for how to actual interact with the deployment 
* Add option to disable sysctlInitContainer 
* [elasticsearch] Adding testing for security context 
* Use generated encryption key for Kibana 
* Don't mount in the current directory when generating certs 
* Update Kibana and Elasticsearch security examples to be generated 
* Add test coverage for esMajorVersion for oss and default value 
* [elasticsearch] Automatically detect esMajorVersion for default image 
* Merge pull request #197 from tetianakravchenko/elasticsearch-add-hooks 
* default values: update example to make it more generic 
* remove not needed comment 
* revertchanges: remove hooks variable 
* Merge pull request #171 from naseemkullah/elasticsearch-runasnonroot 
* values: fix duplication, move hooks closer to the lifecycle; tests: add sugested test 
* revert some changes 
* revert some changes 

### Default value changes

```diff
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index 75507647..6e53ee60 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -16,7 +16,7 @@ roles:
 replicas: 3
 minimumMasterNodes: 2
 
-esMajorVersion: 7
+esMajorVersion: ""
 
 # Allows you to add any config files in /usr/share/elasticsearch/config/
 # such as elasticsearch.yml and log4j2.properties
@@ -27,14 +27,6 @@ esConfig: {}
 #  log4j2.properties: |
 #    key = value
 
-# Add custom preStop, postStart lifecycle hooks
-hooks: {}
-  ## (string) Script to execute prior the pod stops.
-  # preStop: |-
-
-  ## (string) Script to execute after the pod starts.
-  # postStart: |-
-
 # Extra environment variables to append to this nodeGroup
 # This will be appended to the current 'env:' key. You can use any of the kubernetes env
 # syntax here
@@ -145,8 +137,20 @@ updateStrategy: RollingUpdate
 # of your pods to be unavailable during maintenance
 maxUnavailable: 1
 
- # GroupID for the elasticsearch user. The official elastic docker images always have the id of 1000
-fsGroup: 1000
+podSecurityContext:
+  fsGroup: 1000
+
+# The following value is deprecated,
+# please use the above podSecurityContext.fsGroup instead
+fsGroup: ""
+
+securityContext:
+  capabilities:
+    drop:
+    - ALL
+  # readOnlyRootFilesystem: true
+  runAsNonRoot: true
+  runAsUser: 1000
 
 # How long to wait for elasticsearch to stop gracefully
 terminationGracePeriod: 120
@@ -196,7 +200,10 @@ masterTerminationFix: false
 lifecycle: {}
   # preStop:
   #   exec:
-  #     command: ["/bin/bash","/preStop"]
-  # preStop:
+  #     command: ["/bin/sh", "-c", "echo Hello from the postStart handler > /usr/share/message"]
+  # postStart:
   #   exec:
-  #     command: ["/bin/bash","/preStop"]
+  #     command: ["/bin/sh", "-c", "echo Hello from the postStart handler > /usr/share/message"]
+
+sysctlInitContainer:
+  enabled: true
```

## 7.2.0 

**Release date:** 2019-07-02

![AppVersion: 7.2.0](https://img.shields.io/static/v1?label=AppVersion&message=7.2.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Add option to provide custom start/stop hooks 

### Default value changes

```diff
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index 14d28f71..75507647 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -27,6 +27,14 @@ esConfig: {}
 #  log4j2.properties: |
 #    key = value
 
+# Add custom preStop, postStart lifecycle hooks
+hooks: {}
+  ## (string) Script to execute prior the pod stops.
+  # preStop: |-
+
+  ## (string) Script to execute after the pod starts.
+  # postStart: |-
+
 # Extra environment variables to append to this nodeGroup
 # This will be appended to the current 'env:' key. You can use any of the kubernetes env
 # syntax here
@@ -43,7 +51,7 @@ secretMounts: []
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/elasticsearch/elasticsearch"
-imageTag: "7.1.1"
+imageTag: "7.2.0"
 imagePullPolicy: "IfNotPresent"
 
 podAnnotations: {}
@@ -67,6 +75,14 @@ initResources: {}
   #   cpu: "25m"
   #   memory: "128Mi"
 
+sidecarResources: {}
+  # limits:
+  #   cpu: "25m"
+  #   # memory: "128Mi"
+  # requests:
+  #   cpu: "25m"
+  #   memory: "128Mi"
+
 networkHost: "0.0.0.0"
 
 volumeClaimTemplate:
@@ -173,3 +189,14 @@ ingress:
 
 nameOverride: ""
 fullnameOverride: ""
+
+# https://github.com/elastic/helm-charts/issues/63
+masterTerminationFix: false
+
+lifecycle: {}
+  # preStop:
+  #   exec:
+  #     command: ["/bin/bash","/preStop"]
+  # preStop:
+  #   exec:
+  #     command: ["/bin/bash","/preStop"]
```

## 7.1.1 

**Release date:** 2019-06-17

![AppVersion: 7.1.1](https://img.shields.io/static/v1?label=AppVersion&message=7.1.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Run as 1000 
* [elasticsearch] Fix sidecar container resources test 
* Merge pull request #194 from elastic/seven_too_oh 
* [elasticsearch] Disable masterTerminationFix by default 

### Default value changes

```diff
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index 14d28f71..5f459823 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -129,8 +129,20 @@ updateStrategy: RollingUpdate
 # of your pods to be unavailable during maintenance
 maxUnavailable: 1
 
- # GroupID for the elasticsearch user. The official elastic docker images always have the id of 1000
-fsGroup: 1000
+podSecurityContext:
+  fsGroup: 1000
+
+# The following value is deprecated,
+# please use the above podSecurityContext.fsGroup instead
+fsGroup: ""
+
+securityContext:
+  capabilities:
+    drop:
+    - ALL
+  # readOnlyRootFilesystem: true
+  runAsNonRoot: true
+  runAsUser: 1000
 
 # How long to wait for elasticsearch to stop gracefully
 terminationGracePeriod: 120
```

## 7.2.0 

**Release date:** 2019-07-01

![AppVersion: 7.2.0](https://img.shields.io/static/v1?label=AppVersion&message=7.2.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Release 7.2.0 
* Merge pull request #189 from gnatpat/master 
* Merge pull request #182 from hxquangnhat/patch-1 
* Fix test. 
* Update README. 
* Add resources to sidecar container. 
* Wrong value of secretName 
* Wrong secretName value 
* [elasticsearch] Fix pvc annotations with multiple fields 
* Always run tests against localhost or the service endpoint 
* Drop support for 5.x 
* Remove goss port test and instead do http calls on 0.0.0.0 and service 
* [elasticsearch] Update security example docs to match reality 
* [elasticsearch] Add configurable nodePort to service spec 
* Merge pull request #123 from kimxogus/elasticsearch/fix-master-check 
* Update beta notice and add chart descriptions 
* fix test failure 
* Merge pull request #141 from satchpx/master 
* fix README.md 

### Default value changes

```diff
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index d7a7b416..2c947219 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -43,7 +43,7 @@ secretMounts: []
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/elasticsearch/elasticsearch"
-imageTag: "7.1.1"
+imageTag: "7.2.0"
 imagePullPolicy: "IfNotPresent"
 
 podAnnotations: {}
@@ -119,6 +119,7 @@ transportPort: 9300
 
 service:
   type: ClusterIP
+  nodePort:
   annotations: {}
 
 updateStrategy: RollingUpdate
@@ -146,6 +147,11 @@ readinessProbe:
 # https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-health.html#request-params wait_for_status
 clusterHealthCheckParams: "wait_for_status=green&timeout=1s"
 
+## Use an alternate scheduler.
+## ref: https://kubernetes.io/docs/tasks/administer-cluster/configure-multiple-schedulers/
+##
+schedulerName: ""
+
 imagePullSecrets: []
 nodeSelector: {}
 tolerations: []
```

## 7.1.1 

**Release date:** 2019-06-07

![AppVersion: 7.1.1](https://img.shields.io/static/v1?label=AppVersion&message=7.1.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Merge branch 'master' into elasticsearch/fix-master-check 

### Default value changes

```diff
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index af8c6e1f..d7a7b416 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -43,7 +43,7 @@ secretMounts: []
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/elasticsearch/elasticsearch"
-imageTag: "7.1.0"
+imageTag: "7.1.1"
 imagePullPolicy: "IfNotPresent"
 
 podAnnotations: {}
@@ -117,6 +117,10 @@ protocol: http
 httpPort: 9200
 transportPort: 9300
 
+service:
+  type: ClusterIP
+  annotations: {}
+
 updateStrategy: RollingUpdate
 
 # This is the max unavailable setting for the pod disruption budget
```

## 7.1.0 

**Release date:** 2019-06-06

![AppVersion: 7.1.0](https://img.shields.io/static/v1?label=AppVersion&message=7.1.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* update values for default schedulerName 

### Default value changes

```diff
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index af8c6e1f..788769ee 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -142,6 +142,11 @@ readinessProbe:
 # https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-health.html#request-params wait_for_status
 clusterHealthCheckParams: "wait_for_status=green&timeout=1s"
 
+## Use an alternate scheduler.
+## ref: https://kubernetes.io/docs/tasks/administer-cluster/configure-multiple-schedulers/
+##
+schedulerName: ""
+
 imagePullSecrets: []
 nodeSelector: {}
 tolerations: []
```

## 7.1.1 

**Release date:** 2019-06-06

![AppVersion: 7.1.1](https://img.shields.io/static/v1?label=AppVersion&message=7.1.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Release 7.1.1 
* update README/values; Add test for schedulerName 
* [elasticsearch] Add instructions for how to enable snapshots 
* add capability to specify alternalte scheduler 
* Use basic license for security example 

### Default value changes

```diff
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index af8c6e1f..c1686e5d 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -43,7 +43,7 @@ secretMounts: []
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/elasticsearch/elasticsearch"
-imageTag: "7.1.0"
+imageTag: "7.1.1"
 imagePullPolicy: "IfNotPresent"
 
 podAnnotations: {}
```

## 7.1.0 

**Release date:** 2019-05-21

![AppVersion: 7.1.0](https://img.shields.io/static/v1?label=AppVersion&message=7.1.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* 7.1.0 release and promotion to beta status 
* Merge remote-tracking branch 'upstream/master' into elasticsearch/fix-master-check 
* change required master prefix 
* Merge pull request #109 from lancespeelmon/priorityClassName 
* RE: @Crazybus add a basic template test 

### Default value changes

```diff
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index a3d1da94..af8c6e1f 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -43,7 +43,7 @@ secretMounts: []
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/elasticsearch/elasticsearch"
-imageTag: "7.0.1"
+imageTag: "7.1.0"
 imagePullPolicy: "IfNotPresent"
 
 podAnnotations: {}
```

## 7.0.1-alpha1 

**Release date:** 2019-05-08

![AppVersion: 7.0.1](https://img.shields.io/static/v1?label=AppVersion&message=7.0.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Merge branch 'master' into priorityClassName 
* RE: @Crazybus Could you also add these new value file into each charts readme 

### Default value changes

```diff
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index 637469f5..a3d1da94 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -43,7 +43,7 @@ secretMounts: []
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/elasticsearch/elasticsearch"
-imageTag: "7.0.0"
+imageTag: "7.0.1"
 imagePullPolicy: "IfNotPresent"
 
 podAnnotations: {}
@@ -76,6 +76,7 @@ volumeClaimTemplate:
       storage: 30Gi
 
 persistence:
+  enabled: true
   annotations: {}
 
 extraVolumes: []
```

## 7.0.0-alpha1 

**Release date:** 2019-05-08

![AppVersion: 7.0.0](https://img.shields.io/static/v1?label=AppVersion&message=7.0.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* RE: @Crazybus I think this should be an empty string 
* Bump versions 
* WIP: Add integration tests and other tweaks for filebeat 
* set empty string if curl fails 
* fix test failure, update README 
* rearrange values 
* update pytest spec 
* refactor services 
* check only if master node string is empty 

### Default value changes

```diff
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index 57ebe9f4..637469f5 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -92,6 +92,10 @@ extraInitContainers: []
   #   image: busybox
   #   command: ['do', 'something']
 
+# This is the PriorityClass settings as defined in
+# https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption/#priorityclass
+priorityClassName: ""
+
 # By default this will make sure two pods don't end up on the same node
 # Changing this to a region would allow you to spread pods across regions
 antiAffinityTopologyKey: "kubernetes.io/hostname"
```

## 7.0.1-alpha1 

**Release date:** 2019-05-03

![AppVersion: 7.0.1](https://img.shields.io/static/v1?label=AppVersion&message=7.0.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Merge pull request #119 from kimxogus/elasticsearch/master-sidecar 

### Default value changes

```diff
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index 83ac6799..8a9caf01 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -43,7 +43,7 @@ secretMounts: []
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/elasticsearch/elasticsearch"
-imageTag: "7.0.0"
+imageTag: "7.0.1"
 imagePullPolicy: "IfNotPresent"
 
 podAnnotations: {}
```

## 7.0.0-alpha1 

**Release date:** 2019-05-03

![AppVersion: 7.0.0](https://img.shields.io/static/v1?label=AppVersion&message=7.0.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* refactor reviewed code 
* Merge pull request #117 from elastic/filebeat-chart 
* Bump the version number for the new upgrade test 

### Default value changes

```diff
# No changes in this release
```

## 7.0.1-alpha1 

**Release date:** 2019-05-02

![AppVersion: 7.0.1](https://img.shields.io/static/v1?label=AppVersion&message=7.0.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Merge pull request #115 from elastic/up_and_at_them 
* rename sidecar container name 
* rename container name 
* rewrite without voting_config_exclusion 

### Default value changes

```diff
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index 83ac6799..8a9caf01 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -43,7 +43,7 @@ secretMounts: []
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/elasticsearch/elasticsearch"
-imageTag: "7.0.0"
+imageTag: "7.0.1"
 imagePullPolicy: "IfNotPresent"
 
 podAnnotations: {}
```

## 7.0.0-alpha1 

**Release date:** 2019-05-01

![AppVersion: 7.0.0](https://img.shields.io/static/v1?label=AppVersion&message=7.0.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Init helm so that this works properly in a fresh CI environment 

### Default value changes

```diff
# No changes in this release
```

## 7.0.1-alpha1 

**Release date:** 2019-05-01

![AppVersion: 7.0.1](https://img.shields.io/static/v1?label=AppVersion&message=7.0.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Release 7.0.1 and update the changelog 
* [elasticsearch] Add upgrade integration test 
* [elasticsearch] Make persistent volumes optional 
* add support for k8s PriorityClass 
* Fix some incorrect es/kibana default values in docs 
* Add explicit chart version to the imageTag example 
* add a test, fix some tests 
* set the change for 7+ versions 
* add discovery.seed_hosts setting 
* remove discovery.zen.ping.unicast.hosts settign on esMajorVersion > 7 

### Default value changes

```diff
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index 57ebe9f4..8a9caf01 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -43,7 +43,7 @@ secretMounts: []
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/elasticsearch/elasticsearch"
-imageTag: "7.0.0"
+imageTag: "7.0.1"
 imagePullPolicy: "IfNotPresent"
 
 podAnnotations: {}
@@ -76,6 +76,7 @@ volumeClaimTemplate:
       storage: 30Gi
 
 persistence:
+  enabled: true
   annotations: {}
 
 extraVolumes: []
```

## 7.0.0-alpha1 

**Release date:** 2019-04-17

![AppVersion: 7.0.0](https://img.shields.io/static/v1?label=AppVersion&message=7.0.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Merge branch 'master' into es/update-values 

### Default value changes

```diff
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index 6e12f739..57ebe9f4 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -16,7 +16,7 @@ roles:
 replicas: 3
 minimumMasterNodes: 2
 
-esMajorVersion: 6
+esMajorVersion: 7
 
 # Allows you to add any config files in /usr/share/elasticsearch/config/
 # such as elasticsearch.yml and log4j2.properties
@@ -43,7 +43,7 @@ secretMounts: []
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/elasticsearch/elasticsearch"
-imageTag: "6.6.2"
+imageTag: "7.0.0"
 imagePullPolicy: "IfNotPresent"
 
 podAnnotations: {}
@@ -78,16 +78,16 @@ volumeClaimTemplate:
 persistence:
   annotations: {}
 
-extraVolumes: |
+extraVolumes: []
   # - name: extras
   #   emptyDir: {}
 
-extraVolumeMounts: |
+extraVolumeMounts: []
   # - name: extras
   #   mountPath: /usr/share/extras
   #   readOnly: true
 
-extraInitContainers: |
+extraInitContainers: []
   # - name: do-something
   #   image: busybox
   #   command: ['do', 'something']
```

## 6.6.2-alpha1 

**Release date:** 2019-04-17

![AppVersion: 6.6.2](https://img.shields.io/static/v1?label=AppVersion&message=6.6.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* revert version bumps 
* Latest 5.x release is actually 5.6.16 
* Switch the new extra values into arrays 

### Default value changes

```diff
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index db65b429..6e12f739 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -71,7 +71,6 @@ networkHost: "0.0.0.0"
 
 volumeClaimTemplate:
   accessModes: [ "ReadWriteOnce" ]
-  storageClassName: "standard"
   resources:
     requests:
       storage: 30Gi
```

## 7.0.0-alpha1 

**Release date:** 2019-04-16

![AppVersion: 7.0.0](https://img.shields.io/static/v1?label=AppVersion&message=7.0.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Switch to 7.x as the default version 
* update README.md 

### Default value changes

```diff
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index db65b429..33ed2168 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -16,7 +16,7 @@ roles:
 replicas: 3
 minimumMasterNodes: 2
 
-esMajorVersion: 6
+esMajorVersion: 7
 
 # Allows you to add any config files in /usr/share/elasticsearch/config/
 # such as elasticsearch.yml and log4j2.properties
@@ -43,7 +43,7 @@ secretMounts: []
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/elasticsearch/elasticsearch"
-imageTag: "6.6.2"
+imageTag: "7.0.0"
 imagePullPolicy: "IfNotPresent"
 
 podAnnotations: {}
```

## 6.6.2-alpha1 

**Release date:** 2019-04-16

![AppVersion: 6.7.1](https://img.shields.io/static/v1?label=AppVersion&message=6.7.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* revert chart version bump 

### Default value changes

```diff
# No changes in this release
```

## 6.7.1-alpha1 

**Release date:** 2019-04-16

![AppVersion: 6.7.1](https://img.shields.io/static/v1?label=AppVersion&message=6.7.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Merge branch 'master' into es/update-values 
* Merge pull request #70 from tproenca/tproenca/customization 
* Merge pull request #95 from Conky5/doc-tweaks 
* Merge pull request #86 from elastic/migrating_west_for_the_winter 
* Tweak masterService description 
* Use 6.7 for tls link 

### Default value changes

```diff
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index e0f9f2a4..dad9bdb0 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -43,7 +43,7 @@ secretMounts: []
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/elasticsearch/elasticsearch"
-imageTag: "6.6.2"
+imageTag: "6.7.1"
 imagePullPolicy: "IfNotPresent"
 
 podAnnotations: {}
@@ -71,11 +71,27 @@ networkHost: "0.0.0.0"
 
 volumeClaimTemplate:
   accessModes: [ "ReadWriteOnce" ]
-  storageClassName: "standard"
   resources:
     requests:
       storage: 30Gi
 
+persistence:
+  annotations: {}
+
+extraVolumes: |
+  # - name: extras
+  #   emptyDir: {}
+
+extraVolumeMounts: |
+  # - name: extras
+  #   mountPath: /usr/share/extras
+  #   readOnly: true
+
+extraInitContainers: |
+  # - name: do-something
+  #   image: busybox
+  #   command: ['do', 'something']
+
 # By default this will make sure two pods don't end up on the same node
 # Changing this to a region would allow you to spread pods across regions
 antiAffinityTopologyKey: "kubernetes.io/hostname"
```

## 6.6.2-alpha1 

**Release date:** 2019-04-12

![AppVersion: 6.6.2](https://img.shields.io/static/v1?label=AppVersion&message=6.6.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Update documentation links 
* remove standard storageClassName in default values 

### Default value changes

```diff
# No changes in this release
```

## 6.7.1-alpha1 

**Release date:** 2019-04-12

![AppVersion: 6.7.1](https://img.shields.io/static/v1?label=AppVersion&message=6.7.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* bump es version 
* Update examples to 7.0.0 
* Add extra check to make sure that all masters are present. 
* Add link to migration guide in main readme. 
* Remove unnecessary "and" before leftover 
* Update README.md 
* Enable strict mode for helm lint test 
* Add migration guide for migrating from the community helm/charts repo 
* Merge pull request #59 from weirdwater/master 
* Updated the documentation 
* Added tests 
* Added extra volumes, volumeMounts, initContainers 
* Added default values for customization 

### Default value changes

```diff
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index 4d98d971..4b53e595 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -4,7 +4,7 @@ nodeGroup: "master"
 
 # The service that non master groups will try to connect to when joining the cluster
 # This should be set to clusterName + "-" + nodeGroup for your master group
-# masterService: "elasticsearch-master"
+masterService: ""
 
 # Elasticsearch roles that will be applied to this nodeGroup
 # These will be set as environment variables. E.g. node.master=true
@@ -20,7 +20,7 @@ esMajorVersion: 6
 
 # Allows you to add any config files in /usr/share/elasticsearch/config/
 # such as elasticsearch.yml and log4j2.properties
-esConfig:
+esConfig: {}
 #  elasticsearch.yml: |
 #    key:
 #      nestedkey: value
@@ -30,20 +30,20 @@ esConfig:
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
 
 image: "docker.elastic.co/elasticsearch/elasticsearch"
-imageTag: "6.6.2"
+imageTag: "6.7.1"
 imagePullPolicy: "IfNotPresent"
 
 podAnnotations: {}
@@ -139,3 +139,6 @@ ingress:
   #  - secretName: chart-example-tls
   #    hosts:
   #      - chart-example.local
+
+nameOverride: ""
+fullnameOverride: ""
```

## 6.6.2-alpha1 

**Release date:** 2019-03-18

![AppVersion: 6.6.2](https://img.shields.io/static/v1?label=AppVersion&message=6.6.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Merge pull request #79 from elastic/bumper_cars 
* Merge pull request #76 from Conky5/seven-one 

### Default value changes

```diff
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index f03f1360..4d98d971 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -43,7 +43,7 @@ secretMounts:
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/elasticsearch/elasticsearch"
-imageTag: "6.6.0"
+imageTag: "6.6.2"
 imagePullPolicy: "IfNotPresent"
 
 podAnnotations: {}
```

## 6.6.0-alpha1 

**Release date:** 2019-03-14

![AppVersion: 6.6.0](https://img.shields.io/static/v1?label=AppVersion&message=6.6.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* small typo in 'parallel' 

### Default value changes

```diff
# No changes in this release
```

## 6.6.2-alpha1 

**Release date:** 2019-03-15

![AppVersion: 6.6.2](https://img.shields.io/static/v1?label=AppVersion&message=6.6.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Bump version to 6.6.2-alpha1 
* Only require setting masterService for custom master nodeGroups 
* update README.md 

### Default value changes

```diff
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index f65d1cc9..4d98d971 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -4,7 +4,7 @@ nodeGroup: "master"
 
 # The service that non master groups will try to connect to when joining the cluster
 # This should be set to clusterName + "-" + nodeGroup for your master group
-masterService: ""
+# masterService: "elasticsearch-master"
 
 # Elasticsearch roles that will be applied to this nodeGroup
 # These will be set as environment variables. E.g. node.master=true
@@ -43,7 +43,7 @@ secretMounts:
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/elasticsearch/elasticsearch"
-imageTag: "6.6.0"
+imageTag: "6.6.2"
 imagePullPolicy: "IfNotPresent"
 
 podAnnotations: {}
```

## 6.6.0-alpha1 

**Release date:** 2019-03-13

![AppVersion: 6.6.0](https://img.shields.io/static/v1?label=AppVersion&message=6.6.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* remove version bump 
* remove unnecessary masterService values 

### Default value changes

```diff
# No changes in this release
```

## 6.6.0-alpha2 

**Release date:** 2019-03-13

![AppVersion: 6.6.0](https://img.shields.io/static/v1?label=AppVersion&message=6.6.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* add masterService template defaults uname 
* Add 7.1.0 snapshot testing 
* [elasticsearch] Add transport port to headless service 
* :TableFormat 
* Set types of empty Elasticsearch values 
* PR feedback 
* Bump 7 alpha testing to 7 beta 1 
* issue-60:  Update clusterHealthCheckParams description 
* issue-60: Widen scope of configurable health param string that the readinessProbe will wait for 
* issue-60: Add a configurable health status that the readinessProbe will wait for 
* Add FAQ section with keystore and plugin examples 

### Default value changes

```diff
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index 657faf0f..f65d1cc9 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -4,7 +4,7 @@ nodeGroup: "master"
 
 # The service that non master groups will try to connect to when joining the cluster
 # This should be set to clusterName + "-" + nodeGroup for your master group
-masterService: "elasticsearch-master"
+masterService: ""
 
 # Elasticsearch roles that will be applied to this nodeGroup
 # These will be set as environment variables. E.g. node.master=true
@@ -118,6 +118,9 @@ readinessProbe:
   successThreshold: 3
   timeoutSeconds: 5
 
+# https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-health.html#request-params wait_for_status
+clusterHealthCheckParams: "wait_for_status=green&timeout=1s"
+
 imagePullSecrets: []
 nodeSelector: {}
 tolerations: []
```

## 6.6.0-alpha1 

**Release date:** 2019-01-29

![AppVersion: 6.6.0](https://img.shields.io/static/v1?label=AppVersion&message=6.6.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Bump version to 6.6.0 
* [elasticsearch] Add configurable pod annotations 
* Merge pull request #40 from powerhome/add-resources-for-initContainers 
* Add tests for initcontainer resources 

### Default value changes

```diff
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index 203d6cec..657faf0f 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -43,9 +43,12 @@ secretMounts:
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/elasticsearch/elasticsearch"
-imageTag: "6.5.4"
+imageTag: "6.6.0"
 imagePullPolicy: "IfNotPresent"
 
+podAnnotations: {}
+  # iam.amazonaws.com/role: es-cluster
+
 esJavaOpts: "-Xmx1g -Xms1g"
 
 resources:
@@ -81,6 +84,10 @@ antiAffinityTopologyKey: "kubernetes.io/hostname"
 # and that they will never end up on the same node. Setting this to soft will do this "best effort"
 antiAffinity: "hard"
 
+# This is the node affinity settings as defined in
+# https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#node-affinity-beta-feature
+nodeAffinity: {}
+
 # The default is to deploy all pods serially. By setting this to parallel all pods are started at
 # the same time when bootstrapping the cluster
 podManagementPolicy: "Parallel"
```

## 6.5.4-alpha3 

**Release date:** 2019-01-24

![AppVersion: 6.5.4](https://img.shields.io/static/v1?label=AppVersion&message=6.5.4&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Reset version to original 

### Default value changes

```diff
# No changes in this release
```

## 6.5.4-alpha4 

**Release date:** 2019-01-24

![AppVersion: 6.5.4](https://img.shields.io/static/v1?label=AppVersion&message=6.5.4&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Add extra variable to to elasticsearch README 

### Default value changes

```diff
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index 3561f8d1..203d6cec 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -37,7 +37,7 @@ extraEnvs:
 # A list of secrets and their paths to mount inside the pod
 # This is useful for mounting certificates for security and for mounting
 # the X-Pack license
-secretMounts: 
+secretMounts:
 #  - name: elastic-certificates
 #    secretName: elastic-certificates
 #    path: /usr/share/elasticsearch/config/certs
@@ -56,6 +56,14 @@ resources:
     cpu: "1000m"
     memory: "2Gi"
 
+initResources: {}
+  # limits:
+  #   cpu: "25m"
+  #   # memory: "128Mi"
+  # requests:
+  #   cpu: "25m"
+  #   memory: "128Mi"
+
 networkHost: "0.0.0.0"
 
 volumeClaimTemplate:
@@ -67,7 +75,7 @@ volumeClaimTemplate:
 
 # By default this will make sure two pods don't end up on the same node
 # Changing this to a region would allow you to spread pods across regions
-antiAffinityTopologyKey: "kubernetes.io/hostname" 
+antiAffinityTopologyKey: "kubernetes.io/hostname"
 
 # Hard means that by default pods will only be scheduled if there are enough nodes for them
 # and that they will never end up on the same node. Setting this to soft will do this "best effort"
```

## 6.5.4-alpha3 

**Release date:** 2019-01-21

![AppVersion: 6.5.4](https://img.shields.io/static/v1?label=AppVersion&message=6.5.4&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [elasticsearch] Add nodeAffinity support. 

### Default value changes

```diff
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index 3561f8d1..ebd62c29 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -73,6 +73,10 @@ antiAffinityTopologyKey: "kubernetes.io/hostname"
 # and that they will never end up on the same node. Setting this to soft will do this "best effort"
 antiAffinity: "hard"
 
+# This is the node affinity settings as defined in
+# https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#node-affinity-beta-feature
+nodeAffinity: {}
+
 # The default is to deploy all pods serially. By setting this to parallel all pods are started at
 # the same time when bootstrapping the cluster
 podManagementPolicy: "Parallel"
```

## 6.5.4-alpha4 

**Release date:** 2019-01-18

![AppVersion: 6.5.4](https://img.shields.io/static/v1?label=AppVersion&message=6.5.4&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Add resources to InitContainers 
* Add support for setting configuration via configmaps 
* Merge pull request #33 from elastic/shufalaughagus 

### Default value changes

```diff
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index 998ad94d..203d6cec 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -18,6 +18,15 @@ minimumMasterNodes: 2
 
 esMajorVersion: 6
 
+# Allows you to add any config files in /usr/share/elasticsearch/config/
+# such as elasticsearch.yml and log4j2.properties
+esConfig:
+#  elasticsearch.yml: |
+#    key:
+#      nestedkey: value
+#  log4j2.properties: |
+#    key = value
+
 # Extra environment variables to append to this nodeGroup
 # This will be appended to the current 'env:' key. You can use any of the kubernetes env
 # syntax here
@@ -28,7 +37,7 @@ extraEnvs:
 # A list of secrets and their paths to mount inside the pod
 # This is useful for mounting certificates for security and for mounting
 # the X-Pack license
-secretMounts: 
+secretMounts:
 #  - name: elastic-certificates
 #    secretName: elastic-certificates
 #    path: /usr/share/elasticsearch/config/certs
@@ -47,6 +56,14 @@ resources:
     cpu: "1000m"
     memory: "2Gi"
 
+initResources: {}
+  # limits:
+  #   cpu: "25m"
+  #   # memory: "128Mi"
+  # requests:
+  #   cpu: "25m"
+  #   memory: "128Mi"
+
 networkHost: "0.0.0.0"
 
 volumeClaimTemplate:
@@ -58,7 +75,7 @@ volumeClaimTemplate:
 
 # By default this will make sure two pods don't end up on the same node
 # Changing this to a region would allow you to spread pods across regions
-antiAffinityTopologyKey: "kubernetes.io/hostname" 
+antiAffinityTopologyKey: "kubernetes.io/hostname"
 
 # Hard means that by default pods will only be scheduled if there are enough nodes for them
 # and that they will never end up on the same node. Setting this to soft will do this "best effort"
```

## 6.5.4-alpha3 

**Release date:** 2019-01-14

![AppVersion: 6.5.4](https://img.shields.io/static/v1?label=AppVersion&message=6.5.4&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Switch helm chart icons to png 
* Properly set release variable so we can programmatically test with goss 

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
* Merge pull request #31 from elastic/rolly_polly 
* Generate node roles instead of hard coding them 

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
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index 29377ed9..998ad94d 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -16,6 +16,8 @@ roles:
 replicas: 3
 minimumMasterNodes: 2
 
+esMajorVersion: 6
+
 # Extra environment variables to append to this nodeGroup
 # This will be appended to the current 'env:' key. You can use any of the kubernetes env
 # syntax here
@@ -32,7 +34,7 @@ secretMounts:
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/elasticsearch/elasticsearch"
-imageTag: "6.5.0"
+imageTag: "6.5.4"
 imagePullPolicy: "IfNotPresent"
 
 esJavaOpts: "-Xmx1g -Xms1g"
```

## 6.5.0 

**Release date:** 2019-01-04

![AppVersion: 6.5.0](https://img.shields.io/static/v1?label=AppVersion&message=6.5.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Added icons to ES and Kibana Chart.yaml 
* Add Kibana 7 example and update Elasticsearch 7 to also use goss 

### Default value changes

```diff
# No changes in this release
```

## 6.5.4-alpha1 

**Release date:** 2019-01-04

![AppVersion: 6.5.4](https://img.shields.io/static/v1?label=AppVersion&message=6.5.4&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Merge branch 'master' into zen_master 
* Fix up test name 
* Add `esMajorVersion` to determine discovery method 

### Default value changes

```diff
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index 74ffb112..998ad94d 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -16,7 +16,7 @@ roles:
 replicas: 3
 minimumMasterNodes: 2
 
-esDiscoveryModule: "zen"
+esMajorVersion: 6
 
 # Extra environment variables to append to this nodeGroup
 # This will be appended to the current 'env:' key. You can use any of the kubernetes env
@@ -34,7 +34,7 @@ secretMounts:
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/elasticsearch/elasticsearch"
-imageTag: "6.5.3"
+imageTag: "6.5.4"
 imagePullPolicy: "IfNotPresent"
 
 esJavaOpts: "-Xmx1g -Xms1g"
```

## 6.5.3-alpha1 

**Release date:** 2019-01-02

![AppVersion: 6.5.3](https://img.shields.io/static/v1?label=AppVersion&message=6.5.3&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Don't set minimum master nodes when using zen2 discovery 
* Get username and password from environment 
* Add goss integration tests for the multi and security examples 
* Add goss integration tests for the 5.x example 
* Fix testing and install ordering 

### Default value changes

```diff
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index 219b2647..74ffb112 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -16,6 +16,8 @@ roles:
 replicas: 3
 minimumMasterNodes: 2
 
+esDiscoveryModule: "zen"
+
 # Extra environment variables to append to this nodeGroup
 # This will be appended to the current 'env:' key. You can use any of the kubernetes env
 # syntax here
```

## 6.5.4-alpha1 

**Release date:** 2018-12-31

![AppVersion: 6.5.4](https://img.shields.io/static/v1?label=AppVersion&message=6.5.4&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Add goss integration tests for the default example 
* Update the node.name to match the resolvable name 

### Default value changes

```diff
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index 219b2647..eb8f93ae 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -32,7 +32,7 @@ secretMounts:
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/elasticsearch/elasticsearch"
-imageTag: "6.5.3"
+imageTag: "6.5.4"
 imagePullPolicy: "IfNotPresent"
 
 esJavaOpts: "-Xmx1g -Xms1g"
```

## 6.5.3-alpha1 

**Release date:** 2018-12-28

![AppVersion: 6.5.3](https://img.shields.io/static/v1?label=AppVersion&message=6.5.3&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Add testing suite for v7 using the new zen2 discovery method 

### Default value changes

```diff
# No changes in this release
```

## 6.5.4-alpha1 

**Release date:** 2018-12-28

![AppVersion: 6.5.4](https://img.shields.io/static/v1?label=AppVersion&message=6.5.4&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Update maintainer mailing list and bump to 6.5.4 
* Add documentation and example for running the chart on docker-for-mac 

### Default value changes

```diff
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index 219b2647..eb8f93ae 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -32,7 +32,7 @@ secretMounts:
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/elasticsearch/elasticsearch"
-imageTag: "6.5.3"
+imageTag: "6.5.4"
 imagePullPolicy: "IfNotPresent"
 
 esJavaOpts: "-Xmx1g -Xms1g"
```

## 6.5.3-alpha1 

**Release date:** 2018-12-13

![AppVersion: 6.5.3](https://img.shields.io/static/v1?label=AppVersion&message=6.5.3&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Bump default version to 6.5.3 

### Default value changes

```diff
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index 39e41033..219b2647 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -32,7 +32,7 @@ secretMounts:
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/elasticsearch/elasticsearch"
-imageTag: "6.5.2"
+imageTag: "6.5.3"
 imagePullPolicy: "IfNotPresent"
 
 esJavaOpts: "-Xmx1g -Xms1g"
```

## 6.5.2-alpha1 

**Release date:** 2018-12-06

![AppVersion: 6.5.2](https://img.shields.io/static/v1?label=AppVersion&message=6.5.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Bump to version 6.5.2 and append the alpha1 release tag 
* Update Elasticsearch README 
* Update wording a spelling of the usage notes 
* Add usage notes and some getting started tips 

### Default value changes

```diff
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index 29377ed9..39e41033 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -32,7 +32,7 @@ secretMounts:
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/elasticsearch/elasticsearch"
-imageTag: "6.5.0"
+imageTag: "6.5.2"
 imagePullPolicy: "IfNotPresent"
 
 esJavaOpts: "-Xmx1g -Xms1g"
```

## 6.5.0 

**Release date:** 2018-11-20

![AppVersion: 6.5.0](https://img.shields.io/static/v1?label=AppVersion&message=6.5.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Bump default version to 6.5.0 

### Default value changes

```diff
diff --git a/elasticsearch/values.yaml b/elasticsearch/values.yaml
index c488844b..29377ed9 100755
--- a/elasticsearch/values.yaml
+++ b/elasticsearch/values.yaml
@@ -32,7 +32,7 @@ secretMounts:
 #    path: /usr/share/elasticsearch/config/certs
 
 image: "docker.elastic.co/elasticsearch/elasticsearch"
-imageTag: "6.4.3"
+imageTag: "6.5.0"
 imagePullPolicy: "IfNotPresent"
 
 esJavaOpts: "-Xmx1g -Xms1g"
```

## 6.4.3 

**Release date:** 2018-11-13

![AppVersion: 6.4.3](https://img.shields.io/static/v1?label=AppVersion&message=6.4.3&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* first commit 

### Default value changes

```diff
---
clusterName: "elasticsearch"
nodeGroup: "master"

# The service that non master groups will try to connect to when joining the cluster
# This should be set to clusterName + "-" + nodeGroup for your master group
masterService: "elasticsearch-master"

# Elasticsearch roles that will be applied to this nodeGroup
# These will be set as environment variables. E.g. node.master=true
roles:
  master: "true"
  ingest: "true"
  data: "true"

replicas: 3
minimumMasterNodes: 2

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

image: "docker.elastic.co/elasticsearch/elasticsearch"
imageTag: "6.4.3"
imagePullPolicy: "IfNotPresent"

esJavaOpts: "-Xmx1g -Xms1g"

resources:
  requests:
    cpu: "100m"
    memory: "2Gi"
  limits:
    cpu: "1000m"
    memory: "2Gi"

networkHost: "0.0.0.0"

volumeClaimTemplate:
  accessModes: [ "ReadWriteOnce" ]
  storageClassName: "standard"
  resources:
    requests:
      storage: 30Gi

# By default this will make sure two pods don't end up on the same node
# Changing this to a region would allow you to spread pods across regions
antiAffinityTopologyKey: "kubernetes.io/hostname" 

# Hard means that by default pods will only be scheduled if there are enough nodes for them
# and that they will never end up on the same node. Setting this to soft will do this "best effort"
antiAffinity: "hard"

# The default is to deploy all pods serially. By setting this to parallel all pods are started at
# the same time when bootstrapping the cluster
podManagementPolicy: "Parallel"

protocol: http
httpPort: 9200
transportPort: 9300

updateStrategy: RollingUpdate

# This is the max unavailable setting for the pod disruption budget
# The default value of 1 will make sure that kubernetes won't allow more than 1
# of your pods to be unavailable during maintenance
maxUnavailable: 1

 # GroupID for the elasticsearch user. The official elastic docker images always have the id of 1000
fsGroup: 1000

# How long to wait for elasticsearch to stop gracefully
terminationGracePeriod: 120

sysctlVmMaxMapCount: 262144

readinessProbe:
  failureThreshold: 3
  initialDelaySeconds: 10
  periodSeconds: 10
  successThreshold: 3
  timeoutSeconds: 5

imagePullSecrets: []
nodeSelector: {}
tolerations: []

# Enabling this will publically expose your Elasticsearch instance.
# Only enable this if you have security enabled on your cluster
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
```

---
Autogenerated from Helm Chart and git history using [helm-changelog](https://github.com/mogensen/helm-changelog)
