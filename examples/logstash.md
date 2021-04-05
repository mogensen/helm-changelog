# Change Log

## Next Release 

![AppVersion: 8.0.0-SNAPSHOT](https://img.shields.io/static/v1?label=AppVersion&message=8.0.0-SNAPSHOT&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [logstash] Add option loadBalancerIP to service (#1099) 
* [elasticsearch][kibana] remove oss examples (#1046) 
* [logstash] Add support to use pattern files (#883) 
* [meta] add build status and artifact hub badges (#1033) 
* [logstash] disable privileged container in psp (#1000) 
* [all] add hostaliases (#970) 
* [meta] stabilize CI tests (#935) 
* [logstash] add rbac custom annotations (#764) 
* [meta] remove support for k8s <1.14 & helm <2.17.0 (#916) 
* [meta] upgrade test (#907) 
* Helm 3 (#516) 
* [meta] increase helm timeout (#891) 
* Add warning comment placeholder (main branch) (#886) 
* Update warning on charts readme (#863) 
* Added ingress support to the logstash chart (#793) 
* [logstash] use only httpPort in headless service (#839) 
* support tpl in logstashConfig, logstashPipeline and kibanaConfig (#717) 
* [meta] update changelog for 7.9.2 release  (#824) 
* [logstash] Fix headless service ports spec (#776) 
* [meta] Update changelog for 6.8.12 / 7.9.0 releases (#789) 
* [meta] add helm 3 beta support (#759) 
* [meta] update changelog for 7.8.1 and 6.8.11 releases (#755) 
* [doc] fix some links (#737) 
* [logstash] Restart the logstash pod when the secrets have changed (#723) 
* Support creating k8s helm secrets for logstash helm chart #705 (#712) 
* [doc] fix copy-paste errors (#700) 
* [logstash] add headless service for statefulset (#695) 

### Default value changes

```diff
diff --git a/logstash/values.yaml b/logstash/values.yaml
index d3677222..cfccbd3d 100755
--- a/logstash/values.yaml
+++ b/logstash/values.yaml
@@ -25,6 +25,12 @@ logstashPipeline: {}
 #    }
 #    output { stdout { } }
 
+# Allows you to add any pattern files in your custom pattern dir
+logstashPatternDir: "/usr/share/logstash/patterns/"
+logstashPattern: {}
+#    pattern.conf: |
+#      DPKG_VERSION [-+~<>\.0-9a-zA-Z]+
+
 # Extra environment variables to append to this nodeGroup
 # This will be appended to the current 'env:' key. You can use any of the kubernetes env
 # syntax here
@@ -39,9 +45,32 @@ envFrom: []
 # - configMapRef:
 #     name: config-map
 
+# Add sensitive data to k8s secrets
+secrets: []
+#  - name: "env"
+#    value:
+#      ELASTICSEARCH_PASSWORD: "LS1CRUdJTiBgUFJJVkFURSB"
+#      api_key: ui2CsdUadTiBasRJRkl9tvNnw
+#  - name: "tls"
+#    value:
+#      ca.crt: |
+#        LS0tLS1CRUdJT0K
+#        LS0tLS1CRUdJT0K
+#        LS0tLS1CRUdJT0K
+#        LS0tLS1CRUdJT0K
+#      cert.crt: "LS0tLS1CRUdJTiBlRJRklDQVRFLS0tLS0K"
+#      cert.key.filepath: "secrets.crt" # The path to file should be relative to the `values.yaml` file.
+
+
 # A list of secrets and their paths to mount inside the pod
 secretMounts: []
 
+hostAliases: []
+#- ip: "127.0.0.1"
+#  hostnames:
+#  - "foo.local"
+#  - "bar.local"
+
 image: "docker.elastic.co/logstash/logstash"
 imageTag: "8.0.0-SNAPSHOT"
 imagePullPolicy: "IfNotPresent"
@@ -72,12 +101,16 @@ rbac:
   create: false
   serviceAccountAnnotations: {}
   serviceAccountName: ""
+  annotations: {}
+    #annotation1: "value1"
+    #annotation2: "value2"
+    #annotation3: "value3"
 
 podSecurityPolicy:
   create: false
   name: ""
   spec:
-    privileged: true
+    privileged: false
     fsGroup:
       rule: RunAsAny
     runAsUser:
@@ -220,6 +253,7 @@ lifecycle: {}
 service: {}
 #  annotations: {}
 #  type: ClusterIP
+#  loadBalancerIP: ""
 #  ports:
 #    - name: beats
 #      port: 5044
@@ -229,3 +263,13 @@ service: {}
 #      port: 8080
 #      protocol: TCP
 #      targetPort: 8080
+
+ingress:
+  enabled: false
+#  annotations: {}
+#  hosts:
+#    - host: logstash.local
+#      paths:
+#        - path: /logs
+#          servicePort: 8080
+#  tls: []
```

## 8.0.0-SNAPSHOT 

**Release date:** 2020-06-29

![AppVersion: 8.0.0-SNAPSHOT](https://img.shields.io/static/v1?label=AppVersion&message=8.0.0-SNAPSHOT&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Bump master branch to 8.0.0-SNAPSHOT (#682) 
* Add ServiceAccount annotations (#686) 
* [logstash] add security example (#392) 

### Default value changes

```diff
diff --git a/logstash/values.yaml b/logstash/values.yaml
index 9911fada..d3677222 100755
--- a/logstash/values.yaml
+++ b/logstash/values.yaml
@@ -43,7 +43,7 @@ envFrom: []
 secretMounts: []
 
 image: "docker.elastic.co/logstash/logstash"
-imageTag: "7.7.1"
+imageTag: "8.0.0-SNAPSHOT"
 imagePullPolicy: "IfNotPresent"
 imagePullSecrets: []
 
@@ -70,6 +70,7 @@ volumeClaimTemplate:
 
 rbac:
   create: false
+  serviceAccountAnnotations: {}
   serviceAccountName: ""
 
 podSecurityPolicy:
```

## 7.7.1 

**Release date:** 2020-06-03

![AppVersion: 7.7.1](https://img.shields.io/static/v1?label=AppVersion&message=7.7.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Update changelog for 7.7.1 and 6.8.10 releases (#652) 

### Default value changes

```diff
diff --git a/logstash/values.yaml b/logstash/values.yaml
index 4c8c475d..9911fada 100755
--- a/logstash/values.yaml
+++ b/logstash/values.yaml
@@ -43,7 +43,7 @@ envFrom: []
 secretMounts: []
 
 image: "docker.elastic.co/logstash/logstash"
-imageTag: "7.7.0"
+imageTag: "7.7.1"
 imagePullPolicy: "IfNotPresent"
 imagePullSecrets: []
 
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
* Merge pull request #591 from jmlrt/logstash-http-host-var 
* Merge branch 'master' into doc-improvements 
* [meta] move breaking changes into dedicated file and update readme 
* [meta] centralize development and testing doc in CONTRIBUTING.md 
* [logstash] add missing values in readme 
* [logstash] update default values 
* [logstash] format README 
* [logstash] Add envFrom parameter 
* [logstash] update doc and values.yaml for http.host issues 

### Default value changes

```diff
diff --git a/logstash/values.yaml b/logstash/values.yaml
index 41225251..4c8c475d 100755
--- a/logstash/values.yaml
+++ b/logstash/values.yaml
@@ -3,6 +3,9 @@ replicas: 1
 
 # Allows you to add any config files in /usr/share/logstash/config/
 # such as logstash.yml and log4j2.properties
+#
+# Note that when overriding logstash.yml, `http.host: 0.0.0.0` should always be included
+# to make default probes work.
 logstashConfig: {}
 #  logstash.yml: |
 #    key:
@@ -29,11 +32,18 @@ extraEnvs: []
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
 secretMounts: []
 
 image: "docker.elastic.co/logstash/logstash"
-imageTag: "7.6.2"
+imageTag: "7.7.0"
 imagePullPolicy: "IfNotPresent"
 imagePullSecrets: []
 
@@ -152,6 +162,21 @@ securityContext:
 # How long to wait for logstash to stop gracefully
 terminationGracePeriod: 120
 
+# Probes
+# Default probes are using `httpGet` which requires that `http.host: 0.0.0.0` is part of
+# `logstash.yml`. If needed probes can be disabled or overrided using the following syntaxes:
+#
+# disable livenessProbe
+# livenessProbe: null
+#
+# replace httpGet default readinessProbe by some exec probe
+# readinessProbe:
+#   httpGet: null
+#   exec:
+#     command:
+#       - curl
+#      - localhost:9600
+
 livenessProbe:
   httpGet:
     path: /
```

## 7.6.2 

**Release date:** 2020-03-31

![AppVersion: 7.6.2](https://img.shields.io/static/v1?label=AppVersion&message=7.6.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [7.6.2] bump version 

### Default value changes

```diff
diff --git a/logstash/values.yaml b/logstash/values.yaml
index fae3b75e..41225251 100755
--- a/logstash/values.yaml
+++ b/logstash/values.yaml
@@ -33,7 +33,7 @@ extraEnvs: []
 secretMounts: []
 
 image: "docker.elastic.co/logstash/logstash"
-imageTag: "7.6.1"
+imageTag: "7.6.2"
 imagePullPolicy: "IfNotPresent"
 imagePullSecrets: []
 
```

## 7.6.1 

**Release date:** 2020-03-09

![AppVersion: 7.6.1](https://img.shields.io/static/v1?label=AppVersion&message=7.6.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Merge pull request #505 from ChiefAlexander/master 
* Update logstash/tests/logstash_test.py 

### Default value changes

```diff
diff --git a/logstash/values.yaml b/logstash/values.yaml
index 434a2a26..fae3b75e 100755
--- a/logstash/values.yaml
+++ b/logstash/values.yaml
@@ -11,8 +11,9 @@ logstashConfig: {}
 #    key = value
 
 # Allows you to add any pipeline files in /usr/share/logstash/pipeline/
+### ***warn*** there is a hardcoded logstash.conf in the image, override it first
 logstashPipeline: {}
-#  uptime.conf: |
+#  logstash.conf: |
 #    input {
 #      exec {
 #        command => "uptime"
@@ -32,7 +33,7 @@ extraEnvs: []
 secretMounts: []
 
 image: "docker.elastic.co/logstash/logstash"
-imageTag: "7.6.0"
+imageTag: "7.6.1"
 imagePullPolicy: "IfNotPresent"
 imagePullSecrets: []
 
```

## 7.6.0 

**Release date:** 2020-03-07

![AppVersion: 7.6.0](https://img.shields.io/static/v1?label=AppVersion&message=7.6.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Update based on code review 
* Merge pull request #500 from zeph/master 

### Default value changes

```diff
diff --git a/logstash/values.yaml b/logstash/values.yaml
index 498136a5..434a2a26 100755
--- a/logstash/values.yaml
+++ b/logstash/values.yaml
@@ -124,6 +124,11 @@ podManagementPolicy: "Parallel"
 
 httpPort: 9600
 
+# Custom ports to add to logstash
+extraPorts: []
+  # - name: beats
+  #   containerPort: 5001
+
 updateStrategy: RollingUpdate
 
 # This is the max unavailable setting for the pod disruption budget
```

## 7.6.1 

**Release date:** 2020-03-04

![AppVersion: 7.6.1](https://img.shields.io/static/v1?label=AppVersion&message=7.6.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [7.6.1] bump version 
* Update logstash chart to support custom ports 
* warn in reference to  #499 

### Default value changes

```diff
diff --git a/logstash/values.yaml b/logstash/values.yaml
index 498136a5..f63913fd 100755
--- a/logstash/values.yaml
+++ b/logstash/values.yaml
@@ -32,7 +32,7 @@ extraEnvs: []
 secretMounts: []
 
 image: "docker.elastic.co/logstash/logstash"
-imageTag: "7.6.0"
+imageTag: "7.6.1"
 imagePullPolicy: "IfNotPresent"
 imagePullSecrets: []
 
```

## 7.6.0 

**Release date:** 2020-02-11

![AppVersion: 7.6.0](https://img.shields.io/static/v1?label=AppVersion&message=7.6.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [7.6.0] bump version 
* [meta] format with black python scripts 
* [logstash] remove duplicate line 
* [logstash] Update tests 
* [logstash] Add fullnameOverride setting 

### Default value changes

```diff
diff --git a/logstash/values.yaml b/logstash/values.yaml
index b0286707..498136a5 100755
--- a/logstash/values.yaml
+++ b/logstash/values.yaml
@@ -32,7 +32,7 @@ extraEnvs: []
 secretMounts: []
 
 image: "docker.elastic.co/logstash/logstash"
-imageTag: "7.5.2"
+imageTag: "7.6.0"
 imagePullPolicy: "IfNotPresent"
 imagePullSecrets: []
 
```

## 7.5.2 

**Release date:** 2020-01-21

![AppVersion: 7.5.2](https://img.shields.io/static/v1?label=AppVersion&message=7.5.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [7.5.2] bump versions 
* [doc] switch relative URLs to absolute URLs 

### Default value changes

```diff
diff --git a/logstash/values.yaml b/logstash/values.yaml
index 97e8668e..b0286707 100755
--- a/logstash/values.yaml
+++ b/logstash/values.yaml
@@ -32,7 +32,7 @@ extraEnvs: []
 secretMounts: []
 
 image: "docker.elastic.co/logstash/logstash"
-imageTag: "7.5.1"
+imageTag: "7.5.2"
 imagePullPolicy: "IfNotPresent"
 imagePullSecrets: []
 
```

## 7.5.1 

**Release date:** 2019-12-19

![AppVersion: 7.5.1](https://img.shields.io/static/v1?label=AppVersion&message=7.5.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Merge pull request #415 from jmlrt/add-custom-labels-to-pods 

### Default value changes

```diff
diff --git a/logstash/values.yaml b/logstash/values.yaml
index 42f9cbf1..97e8668e 100755
--- a/logstash/values.yaml
+++ b/logstash/values.yaml
@@ -32,7 +32,7 @@ extraEnvs: []
 secretMounts: []
 
 image: "docker.elastic.co/logstash/logstash"
-imageTag: "7.5.0"
+imageTag: "7.5.1"
 imagePullPolicy: "IfNotPresent"
 imagePullSecrets: []
 
```

## 7.5.0 

**Release date:** 2019-12-18

![AppVersion: 7.5.0](https://img.shields.io/static/v1?label=AppVersion&message=7.5.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [logstash] apply labels to all pods 

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
* Prefixed helper functions with chart name 
* [helm] make more explicit that helm 3 is not supported 
* [logstash] increase liveness probe initial delay again 
* [logstash] increase liveness probe initial delay logstash can take longer than 60s to fully start and we can sometime reach the point where Logstash full start happen a few second after liveness probe sending the kill signal. 
* [logstash] add plugin install faq 
* [logstash] update install doc with helm repo instructions 

### Default value changes

```diff
diff --git a/logstash/values.yaml b/logstash/values.yaml
index 2efe2a61..97e8668e 100755
--- a/logstash/values.yaml
+++ b/logstash/values.yaml
@@ -32,7 +32,7 @@ extraEnvs: []
 secretMounts: []
 
 image: "docker.elastic.co/logstash/logstash"
-imageTag: "7.5.0"
+imageTag: "7.5.1"
 imagePullPolicy: "IfNotPresent"
 imagePullSecrets: []
 
@@ -150,7 +150,7 @@ livenessProbe:
   httpGet:
     path: /
     port: http
-  initialDelaySeconds: 60
+  initialDelaySeconds: 300
   periodSeconds: 10
   timeoutSeconds: 5
   failureThreshold: 3
```

## 7.5.0 

**Release date:** 2019-12-02

![AppVersion: 7.5.0](https://img.shields.io/static/v1?label=AppVersion&message=7.5.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [7.5.0] release 
* bump elastic stack versions to 6.8.5 and 7.4.2 
* [logstash] increase pod memory resources 
* [logstash] update logstash default values for memory requirements 
* update install doc 
* [logstash] increase helm timeout 
* [logstash] remove goss port test Goss port test for `tcp:9600` is failing on GKE 1.12 due to https://github.com/aelsabbahy/goss/issues/149. 
* [logstash] add 6.x example 

### Default value changes

```diff
diff --git a/logstash/values.yaml b/logstash/values.yaml
index 2470c794..2efe2a61 100755
--- a/logstash/values.yaml
+++ b/logstash/values.yaml
@@ -32,7 +32,7 @@ extraEnvs: []
 secretMounts: []
 
 image: "docker.elastic.co/logstash/logstash"
-imageTag: "7.4.1"
+imageTag: "7.5.0"
 imagePullPolicy: "IfNotPresent"
 imagePullSecrets: []
 
@@ -46,10 +46,10 @@ logstashJavaOpts: "-Xmx1g -Xms1g"
 resources:
   requests:
     cpu: "100m"
-    memory: "2Gi"
+    memory: "1536Mi"
   limits:
     cpu: "1000m"
-    memory: "2Gi"
+    memory: "1536Mi"
 
 volumeClaimTemplate:
   accessModes: [ "ReadWriteOnce" ]
```

## 7.4.1 

**Release date:** 2019-10-24

![AppVersion: 7.4.1](https://img.shields.io/static/v1?label=AppVersion&message=7.4.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [logstash] use helm 2.15.1 
* [logstash] manage services for inputs implementing a listener 
* [logstash] manage extra containers 
* [logstash] increase probes initial delay 
* [logstash] disable persistence by default 
* [logstash] cleanup/update values in README and values.yaml 
* [logstash] add first python tests 
* [logstash] fix goss test 
* [logstash] add missing lifecycle 
* [logstash] use a different annotation to handle logstash restart if pipeline configmap has changed 
* [logstash] remove unused default value 
* [logstash] stop enforcing node.name as the default value is already the hostname 
* [logstash] remove another deprecated note about storage class 
* [logstash] improve pvc purge `release` label need to be added StatefulSet `spec.selector.matchLabels` to be included in pvc so we can select only pvc matching release 
* [logstash] add oss and default examples 
* [logstash] cleanup remaining pv(c) after goss test 
* [logstash] fix deprecated note about default storage class 
* [logstash] use elasticsearch in logstash example 
* [logstash] add first goss test 
* [logstash] add notes 
* [logstash] allow using either configmap or env variable for logstash.yml 
* [logstash] add first default example 
* [logstash] remove service and ingress resources Logstash HTTP API shouldn't be exposed externally as Logstash doesn't manage any authentication on this API. Also as Logstash doesn't support clustering, being able to access this API via a service doesn't make sense. 

### Default value changes

```diff
diff --git a/logstash/values.yaml b/logstash/values.yaml
index 6c5a076b..2470c794 100755
--- a/logstash/values.yaml
+++ b/logstash/values.yaml
@@ -32,7 +32,7 @@ extraEnvs: []
 secretMounts: []
 
 image: "docker.elastic.co/logstash/logstash"
-imageTag: "7.4.0"
+imageTag: "7.4.1"
 imagePullPolicy: "IfNotPresent"
 imagePullSecrets: []
 
@@ -51,22 +51,6 @@ resources:
     cpu: "1000m"
     memory: "2Gi"
 
-initResources: {}
-  # limits:
-  #   cpu: "25m"
-  #   # memory: "128Mi"
-  # requests:
-  #   cpu: "25m"
-  #   memory: "128Mi"
-
-sidecarResources: {}
-  # limits:
-  #   cpu: "25m"
-  #   # memory: "128Mi"
-  # requests:
-  #   cpu: "25m"
-  #   memory: "128Mi"
-
 volumeClaimTemplate:
   accessModes: [ "ReadWriteOnce" ]
   resources:
@@ -96,7 +80,7 @@ podSecurityPolicy:
       - persistentVolumeClaim
 
 persistence:
-  enabled: true
+  enabled: false
   annotations: {}
 
 extraVolumes: ""
@@ -108,6 +92,11 @@ extraVolumeMounts: ""
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
@@ -133,15 +122,8 @@ nodeAffinity: {}
 # the same time when bootstrapping the cluster
 podManagementPolicy: "Parallel"
 
-protocol: http
 httpPort: 9600
 
-service:
-  type: ClusterIP
-  port: 9600
-  nodePort:
-  annotations: {}
-
 updateStrategy: RollingUpdate
 
 # This is the max unavailable setting for the pod disruption budget
@@ -168,7 +150,7 @@ livenessProbe:
   httpGet:
     path: /
     port: http
-  initialDelaySeconds: 30
+  initialDelaySeconds: 60
   periodSeconds: 10
   timeoutSeconds: 5
   failureThreshold: 3
@@ -178,7 +160,7 @@ readinessProbe:
   httpGet:
     path: /
     port: http
-  initialDelaySeconds: 30
+  initialDelaySeconds: 60
   periodSeconds: 10
   timeoutSeconds: 5
   failureThreshold: 3
@@ -192,21 +174,6 @@ schedulerName: ""
 nodeSelector: {}
 tolerations: []
 
-# Enabling this will publically expose your Logstash instance.
-# Only enable this if you have security enabled on your cluster
-ingress:
-  enabled: false
-  annotations: {}
-    # kubernetes.io/ingress.class: nginx
-    # kubernetes.io/tls-acme: "true"
-  path: /
-  hosts:
-    - chart-example.local
-  tls: []
-  #  - secretName: chart-example-tls
-  #    hosts:
-  #      - chart-example.local
-
 nameOverride: ""
 fullnameOverride: ""
 
@@ -217,3 +184,16 @@ lifecycle: {}
   # postStart:
   #   exec:
   #     command: ["/bin/sh", "-c", "echo Hello from the postStart handler > /usr/share/message"]
+
+service: {}
+#  annotations: {}
+#  type: ClusterIP
+#  ports:
+#    - name: beats
+#      port: 5044
+#      protocol: TCP
+#      targetPort: 5044
+#    - name: http
+#      port: 8080
+#      protocol: TCP
+#      targetPort: 8080
```

## 7.4.0 

**Release date:** 2019-10-15

![AppVersion: 7.4.0](https://img.shields.io/static/v1?label=AppVersion&message=7.4.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [logstash] first version of logstash helm chart 

### Default value changes

```diff
---
replicas: 1

# Allows you to add any config files in /usr/share/logstash/config/
# such as logstash.yml and log4j2.properties
logstashConfig: {}
#  logstash.yml: |
#    key:
#      nestedkey: value
#  log4j2.properties: |
#    key = value

# Allows you to add any pipeline files in /usr/share/logstash/pipeline/
logstashPipeline: {}
#  uptime.conf: |
#    input {
#      exec {
#        command => "uptime"
#        interval => 30
#      }
#    }
#    output { stdout { } }

# Extra environment variables to append to this nodeGroup
# This will be appended to the current 'env:' key. You can use any of the kubernetes env
# syntax here
extraEnvs: []
#  - name: MY_ENVIRONMENT_VAR
#    value: the_value_goes_here

# A list of secrets and their paths to mount inside the pod
secretMounts: []

image: "docker.elastic.co/logstash/logstash"
imageTag: "7.4.0"
imagePullPolicy: "IfNotPresent"
imagePullSecrets: []

podAnnotations: {}

# additionals labels
labels: {}

logstashJavaOpts: "-Xmx1g -Xms1g"

resources:
  requests:
    cpu: "100m"
    memory: "2Gi"
  limits:
    cpu: "1000m"
    memory: "2Gi"

initResources: {}
  # limits:
  #   cpu: "25m"
  #   # memory: "128Mi"
  # requests:
  #   cpu: "25m"
  #   memory: "128Mi"

sidecarResources: {}
  # limits:
  #   cpu: "25m"
  #   # memory: "128Mi"
  # requests:
  #   cpu: "25m"
  #   memory: "128Mi"

volumeClaimTemplate:
  accessModes: [ "ReadWriteOnce" ]
  resources:
    requests:
      storage: 1Gi

rbac:
  create: false
  serviceAccountName: ""

podSecurityPolicy:
  create: false
  name: ""
  spec:
    privileged: true
    fsGroup:
      rule: RunAsAny
    runAsUser:
      rule: RunAsAny
    seLinux:
      rule: RunAsAny
    supplementalGroups:
      rule: RunAsAny
    volumes:
      - secret
      - configMap
      - persistentVolumeClaim

persistence:
  enabled: true
  annotations: {}

extraVolumes: ""
  # - name: extras
  #   emptyDir: {}

extraVolumeMounts: ""
  # - name: extras
  #   mountPath: /usr/share/extras
  #   readOnly: true

extraInitContainers: ""
  # - name: do-something
  #   image: busybox
  #   command: ['do', 'something']

# This is the PriorityClass settings as defined in
# https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption/#priorityclass
priorityClassName: ""

# By default this will make sure two pods don't end up on the same node
# Changing this to a region would allow you to spread pods across regions
antiAffinityTopologyKey: "kubernetes.io/hostname"

# Hard means that by default pods will only be scheduled if there are enough nodes for them
# and that they will never end up on the same node. Setting this to soft will do this "best effort"
antiAffinity: "hard"

# This is the node affinity settings as defined in
# https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#node-affinity-beta-feature
nodeAffinity: {}

# The default is to deploy all pods serially. By setting this to parallel all pods are started at
# the same time when bootstrapping the cluster
podManagementPolicy: "Parallel"

protocol: http
httpPort: 9600

service:
  type: ClusterIP
  port: 9600
  nodePort:
  annotations: {}

updateStrategy: RollingUpdate

# This is the max unavailable setting for the pod disruption budget
# The default value of 1 will make sure that kubernetes won't allow more than 1
# of your pods to be unavailable during maintenance
maxUnavailable: 1

podSecurityContext:
  fsGroup: 1000
  runAsUser: 1000

securityContext:
  capabilities:
    drop:
    - ALL
  # readOnlyRootFilesystem: true
  runAsNonRoot: true
  runAsUser: 1000

# How long to wait for logstash to stop gracefully
terminationGracePeriod: 120

livenessProbe:
  httpGet:
    path: /
    port: http
  initialDelaySeconds: 30
  periodSeconds: 10
  timeoutSeconds: 5
  failureThreshold: 3
  successThreshold: 1

readinessProbe:
  httpGet:
    path: /
    port: http
  initialDelaySeconds: 30
  periodSeconds: 10
  timeoutSeconds: 5
  failureThreshold: 3
  successThreshold: 3

## Use an alternate scheduler.
## ref: https://kubernetes.io/docs/tasks/administer-cluster/configure-multiple-schedulers/
##
schedulerName: ""

nodeSelector: {}
tolerations: []

# Enabling this will publically expose your Logstash instance.
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

nameOverride: ""
fullnameOverride: ""

lifecycle: {}
  # preStop:
  #   exec:
  #     command: ["/bin/sh", "-c", "echo Hello from the postStart handler > /usr/share/message"]
  # postStart:
  #   exec:
  #     command: ["/bin/sh", "-c", "echo Hello from the postStart handler > /usr/share/message"]
```

---
Autogenerated from Helm Chart and git history using [helm-changelog](https://github.com/mogensen/helm-changelog)
