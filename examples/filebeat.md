# Change Log

## Next Release 

![AppVersion: 8.0.0-SNAPSHOT](https://img.shields.io/static/v1?label=AppVersion&message=8.0.0-SNAPSHOT&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [elasticsearch][kibana] remove oss examples (#1046) 
* [filebeat] variable maxUnavailable updateStrategy (#809) 
* [meta] add build status and artifact hub badges (#1033) 
* NIT Update backport config and small doc fixes (#1017) 
* [filebeat][metricbeat] Update documentation on port collisions for multiple beats agents with hostNetworking enabled. (#997) 
* [filebeat] Configurable ClusterRole (#978) 
* Filebeat deployment support feature (#964) 
* [all] add hostaliases (#970) 
* [meta] stabilize CI tests (#935) 
* [meta] remove support for k8s <1.14 & helm <2.17.0 (#916) 
* [meta] upgrade test (#907) 
* Helm 3 (#516) 
* [meta] increase helm timeout (#891) 
* [meta] update rbac.authorization.k8s.io api (#890) 
* Add warning comment placeholder (main branch) (#886) 
* Update warning on charts readme (#863) 
* [filebeat] introduce dnsConfig values for the containers (#659) 
* [meta] update changelog for 7.9.2 release  (#824) 
* Fix typo in FAQ (#744) 
* [meta] Update changelog for 6.8.12 / 7.9.0 releases (#789) 
* [meta] add helm 3 beta support (#759) 
* [doc] update doc links (#758) 
* [meta] update changelog for 7.8.1 and 6.8.11 releases (#755) 
* [doc] fix some links (#737) 
* [filebeat] document probe workaround for kafka output (#699) 
* [filebeat] add permission to list nodes (#704) 

### Default value changes

```diff
diff --git a/filebeat/values.yaml b/filebeat/values.yaml
index e8c4ce22..8e9daf8d 100755
--- a/filebeat/values.yaml
+++ b/filebeat/values.yaml
@@ -1,38 +1,129 @@
 ---
-# Allows you to add any config files in /usr/share/filebeat
-# such as filebeat.yml
-filebeatConfig:
-  filebeat.yml: |
-    filebeat.inputs:
-    - type: container
-      paths:
-        - /var/log/containers/*.log
-      processors:
-      - add_kubernetes_metadata:
-          host: ${NODE_NAME}
-          matchers:
-          - logs_path:
-              logs_path: "/var/log/containers/"
-
-    output.elasticsearch:
-      host: '${NODE_NAME}'
-      hosts: '${ELASTICSEARCH_HOSTS:elasticsearch-master:9200}'
-
-# Extra environment variables to append to the DaemonSet pod spec.
-# This will be appended to the current 'env:' key. You can use any of the kubernetes env
-# syntax here
-extraEnvs: []
-#  - name: MY_ENVIRONMENT_VAR
-#    value: the_value_goes_here
-
-extraVolumeMounts: []
+daemonset:
+  # Annotations to apply to the daemonset
+  annotations: {}
+  # additionals labels
+  labels: {}
+  affinity: {}
+  # Include the daemonset
+  enabled: true
+  # Extra environment variables for Filebeat container.
+  envFrom: []
+  # - configMapRef:
+  #     name: config-secret
+  extraEnvs: []
+  #  - name: MY_ENVIRONMENT_VAR
+  #    value: the_value_goes_here
+  extraVolumes: []
+    # - name: extras
+    #   emptyDir: {}
+  extraVolumeMounts: []
+    # - name: extras
+    #   mountPath: /usr/share/extras
+    #   readOnly: true
+  hostNetworking: false
+  # Allows you to add any config files in /usr/share/filebeat
+  # such as filebeat.yml for daemonset
+  filebeatConfig:
+    filebeat.yml: |
+      filebeat.inputs:
+      - type: container
+        paths:
+          - /var/log/containers/*.log
+        processors:
+        - add_kubernetes_metadata:
+            host: ${NODE_NAME}
+            matchers:
+            - logs_path:
+                logs_path: "/var/log/containers/"
+
+      output.elasticsearch:
+        host: '${NODE_NAME}'
+        hosts: '${ELASTICSEARCH_HOSTS:elasticsearch-master:9200}'
+  # Only used when updateStrategy is set to "RollingUpdate"
+  maxUnavailable: 1
+  nodeSelector: {}
+  # A list of secrets and their paths to mount inside the pod
+  # This is useful for mounting certificates for security other sensitive values
+  secretMounts: []
+  #  - name: filebeat-certificates
+  #    secretName: filebeat-certificates
+  #    path: /usr/share/filebeat/certs
+  # Various pod security context settings. Bear in mind that many of these have an impact on Filebeat functioning properly.
+  #
+  # - User that the container will execute as. Typically necessary to run as root (0) in order to properly collect host container logs.
+  # - Whether to execute the Filebeat containers as privileged containers. Typically not necessarily unless running within environments such as OpenShift.
+  securityContext:
+    runAsUser: 0
+    privileged: false
+  resources:
+    requests:
+      cpu: "100m"
+      memory: "100Mi"
+    limits:
+      cpu: "1000m"
+      memory: "200Mi"
+  tolerations: []
+
+deployment:
+  # Annotations to apply to the deployment
+  annotations: {}
+  # additionals labels
+  labels: {}
+  affinity: {}
+  # Include the deployment
+  enabled: false
+  # Extra environment variables for Filebeat container.
+  envFrom: []
+  # - configMapRef:
+  #     name: config-secret
+  extraEnvs: []
+  #  - name: MY_ENVIRONMENT_VAR
+  #    value: the_value_goes_here
+  # Allows you to add any config files in /usr/share/filebeat
+  extraVolumes: []
+  # - name: extras
+  #   emptyDir: {}
+  extraVolumeMounts: []
   # - name: extras
   #   mountPath: /usr/share/extras
   #   readOnly: true
-
-extraVolumes: []
-  # - name: extras
-  #   emptyDir: {}
+  # such as filebeat.yml for deployment
+  filebeatConfig:
+    filebeat.yml: |
+      filebeat.inputs:
+      - type: tcp
+        max_message_size: 10MiB
+        host: "localhost:9000"
+
+      output.elasticsearch:
+        host: '${NODE_NAME}'
+        hosts: '${ELASTICSEARCH_HOSTS:elasticsearch-master:9200}'
+  nodeSelector: {}
+  # A list of secrets and their paths to mount inside the pod
+  # This is useful for mounting certificates for security other sensitive values
+  secretMounts: []
+  #  - name: filebeat-certificates
+  #    secretName: filebeat-certificates
+  #    path: /usr/share/filebeat/certs
+  #
+  # - User that the container will execute as.
+  # Not necessary to run as root (0) as the Filebeat Deployment use cases do not need access to Kubernetes Node internals
+  # - Typically not necessarily unless running within environments such as OpenShift.
+  securityContext:
+    runAsUser: 0
+    privileged: false
+  resources:
+    requests:
+      cpu: "100m"
+      memory: "100Mi"
+    limits:
+      cpu: "1000m"
+      memory: "200Mi"
+  tolerations: []
+
+# Replicas being used for the filebeat deployment
+replicas: 1
 
 extraContainers: ""
 # - name: dummy-init
@@ -41,16 +132,19 @@ extraContainers: ""
 
 extraInitContainers: []
 # - name: dummy-init
-#   image: busybox
-#   command: ['echo', 'hey']
-
-envFrom: []
-# - configMapRef:
-#     name: configmap-name
 
 # Root directory where Filebeat will write data to in order to persist registry data across pod restarts (file position and other metadata).
 hostPathRoot: /var/lib
-hostNetworking: false
+
+dnsConfig: {}
+# options:
+#   - name: ndots
+#     value: "2"
+hostAliases: []
+#- ip: "127.0.0.1"
+#  hostnames:
+#  - "foo.local"
+#  - "bar.local"
 image: "docker.elastic.co/beats/filebeat"
 imageTag: "8.0.0-SNAPSHOT"
 imagePullPolicy: "IfNotPresent"
@@ -85,51 +179,31 @@ readinessProbe:
 # Whether this chart should self-manage its service account, role, and associated role binding.
 managedServiceAccount: true
 
-# additionals labels
-labels: {}
+clusterRoleRules:
+- apiGroups:
+  - ""
+  resources:
+  - namespaces
+  - nodes
+  - pods
+  verbs:
+  - get
+  - list
+  - watch
 
 podAnnotations: {}
   # iam.amazonaws.com/role: es-cluster
 
-# Various pod security context settings. Bear in mind that many of these have an impact on Filebeat functioning properly.
-#
-# - User that the container will execute as. Typically necessary to run as root (0) in order to properly collect host container logs.
-# - Whether to execute the Filebeat containers as privileged containers. Typically not necessarily unless running within environments such as OpenShift.
-podSecurityContext:
-  runAsUser: 0
-  privileged: false
-
-resources:
-  requests:
-    cpu: "100m"
-    memory: "100Mi"
-  limits:
-    cpu: "1000m"
-    memory: "200Mi"
-
 # Custom service account override that the pod will use
 serviceAccount: ""
 
 # Annotations to add to the ServiceAccount that is created if the serviceAccount value isn't set.
 serviceAccountAnnotations: {}
-  # eks.amazonaws.com/role-arn: arn:aws:iam::111111111111:role/k8s.clustername.namespace.serviceaccount
 
-# A list of secrets and their paths to mount inside the pod
-# This is useful for mounting certificates for security other sensitive values
-secretMounts: []
-#  - name: filebeat-certificates
-#    secretName: filebeat-certificates
-#    path: /usr/share/filebeat/certs
+  # eks.amazonaws.com/role-arn: arn:aws:iam::111111111111:role/k8s.clustername.namespace.serviceaccount
 
 # How long to wait for Filebeat pods to stop gracefully
 terminationGracePeriod: 30
-
-tolerations: []
-
-nodeSelector: {}
-
-affinity: {}
-
 # This is the PriorityClass settings as defined in
 # https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption/#priorityclass
 priorityClassName: ""
@@ -140,3 +214,19 @@ updateStrategy: RollingUpdate
 # Only edit these if you know what you're doing
 nameOverride: ""
 fullnameOverride: ""
+
+# DEPRECATED
+affinity: {}
+envFrom: []
+extraEnvs: []
+extraVolumes: []
+extraVolumeMounts: []
+# Allows you to add any config files in /usr/share/filebeat
+# such as filebeat.yml for both daemonset and deployment
+filebeatConfig: {}
+nodeSelector: {}
+podSecurityContext: {}
+resources: {}
+secretMounts: []
+tolerations: []
+labels: {}
```

## 8.0.0-SNAPSHOT 

**Release date:** 2020-06-29

![AppVersion: 8.0.0-SNAPSHOT](https://img.shields.io/static/v1?label=AppVersion&message=8.0.0-SNAPSHOT&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Bump master branch to 8.0.0-SNAPSHOT (#682) 
* Add ServiceAccount annotations (#686) 

### Default value changes

```diff
diff --git a/filebeat/values.yaml b/filebeat/values.yaml
index 5118993b..e8c4ce22 100755
--- a/filebeat/values.yaml
+++ b/filebeat/values.yaml
@@ -52,7 +52,7 @@ envFrom: []
 hostPathRoot: /var/lib
 hostNetworking: false
 image: "docker.elastic.co/beats/filebeat"
-imageTag: "7.7.1"
+imageTag: "8.0.0-SNAPSHOT"
 imagePullPolicy: "IfNotPresent"
 imagePullSecrets: []
 
@@ -110,6 +110,10 @@ resources:
 # Custom service account override that the pod will use
 serviceAccount: ""
 
+# Annotations to add to the ServiceAccount that is created if the serviceAccount value isn't set.
+serviceAccountAnnotations: {}
+  # eks.amazonaws.com/role-arn: arn:aws:iam::111111111111:role/k8s.clustername.namespace.serviceaccount
+
 # A list of secrets and their paths to mount inside the pod
 # This is useful for mounting certificates for security other sensitive values
 secretMounts: []
```

## 7.7.1 

**Release date:** 2020-06-03

![AppVersion: 7.7.1](https://img.shields.io/static/v1?label=AppVersion&message=7.7.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Update changelog for 7.7.1 and 6.8.10 releases (#652) 

### Default value changes

```diff
diff --git a/filebeat/values.yaml b/filebeat/values.yaml
index 1b396af7..5118993b 100755
--- a/filebeat/values.yaml
+++ b/filebeat/values.yaml
@@ -52,7 +52,7 @@ envFrom: []
 hostPathRoot: /var/lib
 hostNetworking: false
 image: "docker.elastic.co/beats/filebeat"
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
* [meta] move breaking changes into dedicated file and update readme 
* [meta] centralize development and testing doc in CONTRIBUTING.md 
* [filebeat] add missing value in readme 
* [filebeat] update default values 
* [filebeat] format README 
* [metricbeat] move back to docker input and in_cluster for 6.x 
* fixup! [filebeat] use container input instead of docker input 
* [filebeat] use recommended values for add_kubernetes_metadata 
* [filebeat] use container input instead of docker input 
* [filebeat] remove in_cluster config from add_kubernetes_metadata 

### Default value changes

```diff
diff --git a/filebeat/values.yaml b/filebeat/values.yaml
index 6405b2c9..1b396af7 100755
--- a/filebeat/values.yaml
+++ b/filebeat/values.yaml
@@ -4,11 +4,15 @@
 filebeatConfig:
   filebeat.yml: |
     filebeat.inputs:
-    - type: docker
-      containers.ids:
-      - '*'
+    - type: container
+      paths:
+        - /var/log/containers/*.log
       processors:
-      - add_kubernetes_metadata: ~
+      - add_kubernetes_metadata:
+          host: ${NODE_NAME}
+          matchers:
+          - logs_path:
+              logs_path: "/var/log/containers/"
 
     output.elasticsearch:
       host: '${NODE_NAME}'
@@ -48,7 +52,7 @@ envFrom: []
 hostPathRoot: /var/lib
 hostNetworking: false
 image: "docker.elastic.co/beats/filebeat"
-imageTag: "7.6.2"
+imageTag: "7.7.0"
 imagePullPolicy: "IfNotPresent"
 imagePullSecrets: []
 
```

## 7.6.2 

**Release date:** 2020-03-31

![AppVersion: 7.6.2](https://img.shields.io/static/v1?label=AppVersion&message=7.6.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [7.6.2] bump version 
* Accept a string as extraInitContainers value for filebeat 

### Default value changes

```diff
diff --git a/filebeat/values.yaml b/filebeat/values.yaml
index 1b99157f..6405b2c9 100755
--- a/filebeat/values.yaml
+++ b/filebeat/values.yaml
@@ -48,7 +48,7 @@ envFrom: []
 hostPathRoot: /var/lib
 hostNetworking: false
 image: "docker.elastic.co/beats/filebeat"
-imageTag: "7.6.1"
+imageTag: "7.6.2"
 imagePullPolicy: "IfNotPresent"
 imagePullSecrets: []
 
```

## 7.6.1 

**Release date:** 2020-03-04

![AppVersion: 7.6.1](https://img.shields.io/static/v1?label=AppVersion&message=7.6.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [7.6.1] bump version 

### Default value changes

```diff
diff --git a/filebeat/values.yaml b/filebeat/values.yaml
index d3c575c9..1b99157f 100755
--- a/filebeat/values.yaml
+++ b/filebeat/values.yaml
@@ -48,7 +48,7 @@ envFrom: []
 hostPathRoot: /var/lib
 hostNetworking: false
 image: "docker.elastic.co/beats/filebeat"
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
* [filebeat] add extracontainers 
* Add tests for extraInitContainers to filebeat chart 
* Add extraInitContainers to filebeat chart 

### Default value changes

```diff
diff --git a/filebeat/values.yaml b/filebeat/values.yaml
index 6cce7feb..d3c575c9 100755
--- a/filebeat/values.yaml
+++ b/filebeat/values.yaml
@@ -30,6 +30,16 @@ extraVolumes: []
   # - name: extras
   #   emptyDir: {}
 
+extraContainers: ""
+# - name: dummy-init
+#   image: busybox
+#   command: ['echo', 'hey']
+
+extraInitContainers: []
+# - name: dummy-init
+#   image: busybox
+#   command: ['echo', 'hey']
+
 envFrom: []
 # - configMapRef:
 #     name: configmap-name
@@ -38,7 +48,7 @@ envFrom: []
 hostPathRoot: /var/lib
 hostNetworking: false
 image: "docker.elastic.co/beats/filebeat"
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
* Correct default value of extraVolumeMounts and extraVolumes in filebeat/README.md 
* Merge pull request #420 from jmlrt/override-probe 
* [filebeat] allow override readiness and liveness probes commands 
* Merge branch 'master' into add-support-for-envfrom 

### Default value changes

```diff
diff --git a/filebeat/values.yaml b/filebeat/values.yaml
index b1f2d941..6cce7feb 100755
--- a/filebeat/values.yaml
+++ b/filebeat/values.yaml
@@ -30,21 +30,39 @@ extraVolumes: []
   # - name: extras
   #   emptyDir: {}
 
+envFrom: []
+# - configMapRef:
+#     name: configmap-name
+
 # Root directory where Filebeat will write data to in order to persist registry data across pod restarts (file position and other metadata).
 hostPathRoot: /var/lib
 hostNetworking: false
 image: "docker.elastic.co/beats/filebeat"
-imageTag: "7.5.1"
+imageTag: "7.5.2"
 imagePullPolicy: "IfNotPresent"
 imagePullSecrets: []
 
 livenessProbe:
+  exec:
+    command:
+      - sh
+      - -c
+      - |
+        #!/usr/bin/env bash -e
+        curl --fail 127.0.0.1:5066
   failureThreshold: 3
   initialDelaySeconds: 10
   periodSeconds: 10
   timeoutSeconds: 5
 
 readinessProbe:
+  exec:
+    command:
+      - sh
+      - -c
+      - |
+        #!/usr/bin/env bash -e
+        filebeat test output
   failureThreshold: 3
   initialDelaySeconds: 10
   periodSeconds: 10
```

## 7.5.1 

**Release date:** 2019-12-19

![AppVersion: 7.5.1](https://img.shields.io/static/v1?label=AppVersion&message=7.5.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Merge pull request #415 from jmlrt/add-custom-labels-to-pods 

### Default value changes

```diff
diff --git a/filebeat/values.yaml b/filebeat/values.yaml
index a9dc5fd3..b1f2d941 100755
--- a/filebeat/values.yaml
+++ b/filebeat/values.yaml
@@ -34,7 +34,7 @@ extraVolumes: []
 hostPathRoot: /var/lib
 hostNetworking: false
 image: "docker.elastic.co/beats/filebeat"
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


* [filebeat] apply labels to all pods 

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
* Merge pull request #403 from ChrsMark/patch-1 
* Remove in_cluster config from add_kubernetes_metadata 
* [helm] make more explicit that helm 3 is not supported 
* [filebeat] update link for hostNetworking setting 
* [filebeat] disable host networking by default 

### Default value changes

```diff
diff --git a/filebeat/values.yaml b/filebeat/values.yaml
index 5d28831a..b1f2d941 100755
--- a/filebeat/values.yaml
+++ b/filebeat/values.yaml
@@ -8,8 +8,7 @@ filebeatConfig:
       containers.ids:
       - '*'
       processors:
-      - add_kubernetes_metadata:
-          in_cluster: true
+      - add_kubernetes_metadata: ~
 
     output.elasticsearch:
       host: '${NODE_NAME}'
@@ -33,9 +32,9 @@ extraVolumes: []
 
 # Root directory where Filebeat will write data to in order to persist registry data across pod restarts (file position and other metadata).
 hostPathRoot: /var/lib
-hostNetworking: true
+hostNetworking: false
 image: "docker.elastic.co/beats/filebeat"
-imageTag: "7.5.0"
+imageTag: "7.5.1"
 imagePullPolicy: "IfNotPresent"
 imagePullSecrets: []
 
```

## 7.5.0 

**Release date:** 2019-12-02

![AppVersion: 7.5.0](https://img.shields.io/static/v1?label=AppVersion&message=7.5.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [7.5.0] release 
* bump elastic stack versions to 6.8.5 and 7.4.2 
* Merge branch 'master' into add-support-for-envfrom 
* Merge pull request #322 from pbecotte/FilebeatVolumes 
* Merge pull request #321 from pbecotte/HostnetworkForFilebeat 
* Merge branch 'master' into issue_267 
* change ordering 
* readme tweak 
* tweaks to name and test 
* tests 
* update README 
* added support for envFrom 
* update install doc 
* [filebeat] fix beats icon 
* [helm] add some notice about tested helm version 

### Default value changes

```diff
diff --git a/filebeat/values.yaml b/filebeat/values.yaml
index 0adcae1f..5d28831a 100755
--- a/filebeat/values.yaml
+++ b/filebeat/values.yaml
@@ -12,6 +12,7 @@ filebeatConfig:
           in_cluster: true
 
     output.elasticsearch:
+      host: '${NODE_NAME}'
       hosts: '${ELASTICSEARCH_HOSTS:elasticsearch-master:9200}'
 
 # Extra environment variables to append to the DaemonSet pod spec.
@@ -21,20 +22,20 @@ extraEnvs: []
 #  - name: MY_ENVIRONMENT_VAR
 #    value: the_value_goes_here
 
-extraVolumeMounts: ""
+extraVolumeMounts: []
   # - name: extras
   #   mountPath: /usr/share/extras
   #   readOnly: true
 
-extraVolumes: ""
+extraVolumes: []
   # - name: extras
   #   emptyDir: {}
 
 # Root directory where Filebeat will write data to in order to persist registry data across pod restarts (file position and other metadata).
 hostPathRoot: /var/lib
-
+hostNetworking: true
 image: "docker.elastic.co/beats/filebeat"
-imageTag: "7.4.1"
+imageTag: "7.5.0"
 imagePullPolicy: "IfNotPresent"
 imagePullSecrets: []
 
```

## 7.4.1 

**Release date:** 2019-10-22

![AppVersion: 7.4.1](https://img.shields.io/static/v1?label=AppVersion&message=7.4.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* 7.4.1 release 
* [helm] add chart api version This is required witrh Helm 2.15.0 following https://github.com/helm/helm/pull/6180 
* [Issue #267] Support fullnameOverride in kibana, filebeat, metricbeat charts 
* use a list for extra volume mounts to match the comments and other values 
* use host networking for filebeat so that the stats have the correct node information 

### Default value changes

```diff
diff --git a/filebeat/values.yaml b/filebeat/values.yaml
index 1e6af758..0adcae1f 100755
--- a/filebeat/values.yaml
+++ b/filebeat/values.yaml
@@ -34,7 +34,7 @@ extraVolumes: ""
 hostPathRoot: /var/lib
 
 image: "docker.elastic.co/beats/filebeat"
-imageTag: "7.4.0"
+imageTag: "7.4.1"
 imagePullPolicy: "IfNotPresent"
 imagePullSecrets: []
 
```

## 7.4.0 

**Release date:** 2019-10-01

![AppVersion: 7.4.0](https://img.shields.io/static/v1?label=AppVersion&message=7.4.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* 7.4.0 Release 
* [filebeat] change min k8s version due to daemonset apiVersion 

### Default value changes

```diff
diff --git a/filebeat/values.yaml b/filebeat/values.yaml
index 284185ac..1e6af758 100755
--- a/filebeat/values.yaml
+++ b/filebeat/values.yaml
@@ -34,7 +34,7 @@ extraVolumes: ""
 hostPathRoot: /var/lib
 
 image: "docker.elastic.co/beats/filebeat"
-imageTag: "7.3.2"
+imageTag: "7.4.0"
 imagePullPolicy: "IfNotPresent"
 imagePullSecrets: []
 
```

## 7.3.2 

**Release date:** 2019-09-19

![AppVersion: 7.3.2](https://img.shields.io/static/v1?label=AppVersion&message=7.3.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* 7.3.2 Release 
* Clarify priorityClassName default for filebeat chart 
* Add priorityClassName to filebeat chart 
* [filebeat][metricbeat] Add configurable nodeSelector and affinit… (#243) 
* [filebeat][metricbeat] Add configurable nodeSelector and affinity spec 
* Update documentation and defaults for tmpl values 

### Default value changes

```diff
diff --git a/filebeat/values.yaml b/filebeat/values.yaml
index 3da889cd..284185ac 100755
--- a/filebeat/values.yaml
+++ b/filebeat/values.yaml
@@ -21,12 +21,12 @@ extraEnvs: []
 #  - name: MY_ENVIRONMENT_VAR
 #    value: the_value_goes_here
 
-extraVolumeMounts: []
+extraVolumeMounts: ""
   # - name: extras
   #   mountPath: /usr/share/extras
   #   readOnly: true
 
-extraVolumes: []
+extraVolumes: ""
   # - name: extras
   #   emptyDir: {}
 
@@ -34,7 +34,7 @@ extraVolumes: []
 hostPathRoot: /var/lib
 
 image: "docker.elastic.co/beats/filebeat"
-imageTag: "7.3.0"
+imageTag: "7.3.2"
 imagePullPolicy: "IfNotPresent"
 imagePullSecrets: []
 
@@ -90,6 +90,14 @@ terminationGracePeriod: 30
 
 tolerations: []
 
+nodeSelector: {}
+
+affinity: {}
+
+# This is the PriorityClass settings as defined in
+# https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption/#priorityclass
+priorityClassName: ""
+
 updateStrategy: RollingUpdate
 
 # Override various naming aspects of this chart
```

## 7.3.0 

**Release date:** 2019-07-31

![AppVersion: 7.3.0](https://img.shields.io/static/v1?label=AppVersion&message=7.3.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* 7.3.0 Release 

### Default value changes

```diff
diff --git a/filebeat/values.yaml b/filebeat/values.yaml
index 9cf90069..3da889cd 100755
--- a/filebeat/values.yaml
+++ b/filebeat/values.yaml
@@ -34,7 +34,7 @@ extraVolumes: []
 hostPathRoot: /var/lib
 
 image: "docker.elastic.co/beats/filebeat"
-imageTag: "7.2.0"
+imageTag: "7.3.0"
 imagePullPolicy: "IfNotPresent"
 imagePullSecrets: []
 
```

## 7.2.1-0 

**Release date:** 2019-07-17

![AppVersion: 7.2.0](https://img.shields.io/static/v1?label=AppVersion&message=7.2.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* 7.2.1-0 Release 
* [filebeat] additionals labels 

### Default value changes

```diff
diff --git a/filebeat/values.yaml b/filebeat/values.yaml
index 9bb996c5..9cf90069 100755
--- a/filebeat/values.yaml
+++ b/filebeat/values.yaml
@@ -53,6 +53,9 @@ readinessProbe:
 # Whether this chart should self-manage its service account, role, and associated role binding.
 managedServiceAccount: true
 
+# additionals labels
+labels: {}
+
 podAnnotations: {}
   # iam.amazonaws.com/role: es-cluster
 
```

## 7.2.0 

**Release date:** 2019-07-01

![AppVersion: 7.2.0](https://img.shields.io/static/v1?label=AppVersion&message=7.2.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Merge pull request #194 from elastic/seven_too_oh 

### Default value changes

```diff
diff --git a/filebeat/values.yaml b/filebeat/values.yaml
index 1bb78011..9bb996c5 100755
--- a/filebeat/values.yaml
+++ b/filebeat/values.yaml
@@ -34,7 +34,7 @@ extraVolumes: []
 hostPathRoot: /var/lib
 
 image: "docker.elastic.co/beats/filebeat"
-imageTag: "7.1.1"
+imageTag: "7.2.0"
 imagePullPolicy: "IfNotPresent"
 imagePullSecrets: []
 
```

## 7.1.1 

**Release date:** 2019-07-01

![AppVersion: 7.1.1](https://img.shields.io/static/v1?label=AppVersion&message=7.1.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Merge pull request #188 from cclauss/patch-1 

### Default value changes

```diff
# No changes in this release
```

## 7.2.0 

**Release date:** 2019-07-01

![AppVersion: 7.2.0](https://img.shields.io/static/v1?label=AppVersion&message=7.2.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Release 7.2.0 
* Fix octal literal to work in both Python 2 and Python 3 
* Fix filebeat_test with secretName 
* Change value for secretName 
* Fix unintended new line 
* change secretName field 
* Drop support for 5.x 
* Update beta notice and add chart descriptions 

### Default value changes

```diff
diff --git a/filebeat/values.yaml b/filebeat/values.yaml
index 1bb78011..9bb996c5 100755
--- a/filebeat/values.yaml
+++ b/filebeat/values.yaml
@@ -34,7 +34,7 @@ extraVolumes: []
 hostPathRoot: /var/lib
 
 image: "docker.elastic.co/beats/filebeat"
-imageTag: "7.1.1"
+imageTag: "7.2.0"
 imagePullPolicy: "IfNotPresent"
 imagePullSecrets: []
 
```

## 7.1.1 

**Release date:** 2019-06-06

![AppVersion: 7.1.1](https://img.shields.io/static/v1?label=AppVersion&message=7.1.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Release 7.1.1 
* Remove fsGroup from container level security context 

### Default value changes

```diff
diff --git a/filebeat/values.yaml b/filebeat/values.yaml
index a49d1328..1bb78011 100755
--- a/filebeat/values.yaml
+++ b/filebeat/values.yaml
@@ -34,7 +34,7 @@ extraVolumes: []
 hostPathRoot: /var/lib
 
 image: "docker.elastic.co/beats/filebeat"
-imageTag: "7.1.0"
+imageTag: "7.1.1"
 imagePullPolicy: "IfNotPresent"
 imagePullSecrets: []
 
@@ -58,11 +58,9 @@ podAnnotations: {}
 
 # Various pod security context settings. Bear in mind that many of these have an impact on Filebeat functioning properly.
 #
-# - Filesystem group for the Filebeat user. The official elastic docker images always have an id of 1000.
 # - User that the container will execute as. Typically necessary to run as root (0) in order to properly collect host container logs.
 # - Whether to execute the Filebeat containers as privileged containers. Typically not necessarily unless running within environments such as OpenShift.
 podSecurityContext:
-  fsGroup: 1000
   runAsUser: 0
   privileged: false
 
```

## 7.1.0 

**Release date:** 2019-05-21

![AppVersion: 7.1.0](https://img.shields.io/static/v1?label=AppVersion&message=7.1.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* 7.1.0 release and promotion to beta status 
* Add usage notes for Filebeat 
* Update default value for filebeatConfig 
* Add namespace to hostPath to make sure it is unique per cluster 
* Fix 6.x example and statefulset links in readme 
* Avoid name collisions for the default serviceAccount 

### Default value changes

```diff
diff --git a/filebeat/values.yaml b/filebeat/values.yaml
index e8c1c6da..a49d1328 100755
--- a/filebeat/values.yaml
+++ b/filebeat/values.yaml
@@ -34,7 +34,7 @@ extraVolumes: []
 hostPathRoot: /var/lib
 
 image: "docker.elastic.co/beats/filebeat"
-imageTag: "7.0.1"
+imageTag: "7.1.0"
 imagePullPolicy: "IfNotPresent"
 imagePullSecrets: []
 
@@ -74,8 +74,8 @@ resources:
     cpu: "1000m"
     memory: "200Mi"
 
-# Service account that the pod will use
-serviceAccount: filebeat
+# Custom service account override that the pod will use
+serviceAccount: ""
 
 # A list of secrets and their paths to mount inside the pod
 # This is useful for mounting certificates for security other sensitive values
```

## 7.0.1-alpha1 

**Release date:** 2019-05-07

![AppVersion: 7.0.1](https://img.shields.io/static/v1?label=AppVersion&message=7.0.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Bump versions 
* Add readinessProbe and more tests 
* Only use CA for filebeat security example 
* Add 6.x and security examples (certificates not working yet) 
* Use fullname for the configmap map. 
* WIP: Add integration tests and other tweaks for filebeat 
* [filebeat] python templating tests 
* Remove unused filebeat PDB template 

### Default value changes

```diff
diff --git a/filebeat/values.yaml b/filebeat/values.yaml
index 6fbe4493..e8c1c6da 100755
--- a/filebeat/values.yaml
+++ b/filebeat/values.yaml
@@ -1,10 +1,18 @@
 ---
 # Allows you to add any config files in /usr/share/filebeat
 # such as filebeat.yml
-filebeatConfig: {}
-#  filebeat.yml: |
-#    key:
-#      nestedkey: value
+filebeatConfig:
+  filebeat.yml: |
+    filebeat.inputs:
+    - type: docker
+      containers.ids:
+      - '*'
+      processors:
+      - add_kubernetes_metadata:
+          in_cluster: true
+
+    output.elasticsearch:
+      hosts: '${ELASTICSEARCH_HOSTS:elasticsearch-master:9200}'
 
 # Extra environment variables to append to the DaemonSet pod spec.
 # This will be appended to the current 'env:' key. You can use any of the kubernetes env
@@ -26,7 +34,7 @@ extraVolumes: []
 hostPathRoot: /var/lib
 
 image: "docker.elastic.co/beats/filebeat"
-imageTag: "7.0.0"
+imageTag: "7.0.1"
 imagePullPolicy: "IfNotPresent"
 imagePullSecrets: []
 
@@ -36,6 +44,12 @@ livenessProbe:
   periodSeconds: 10
   timeoutSeconds: 5
 
+readinessProbe:
+  failureThreshold: 3
+  initialDelaySeconds: 10
+  periodSeconds: 10
+  timeoutSeconds: 5
+
 # Whether this chart should self-manage its service account, role, and associated role binding.
 managedServiceAccount: true
 
```

## 7.0.0-alpha1 

**Release date:** 2019-04-26

![AppVersion: 7.0.0](https://img.shields.io/static/v1?label=AppVersion&message=7.0.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Initial commit of Filebeat chart 

### Default value changes

```diff
---
# Allows you to add any config files in /usr/share/filebeat
# such as filebeat.yml
filebeatConfig: {}
#  filebeat.yml: |
#    key:
#      nestedkey: value

# Extra environment variables to append to the DaemonSet pod spec.
# This will be appended to the current 'env:' key. You can use any of the kubernetes env
# syntax here
extraEnvs: []
#  - name: MY_ENVIRONMENT_VAR
#    value: the_value_goes_here

extraVolumeMounts: []
  # - name: extras
  #   mountPath: /usr/share/extras
  #   readOnly: true

extraVolumes: []
  # - name: extras
  #   emptyDir: {}

# Root directory where Filebeat will write data to in order to persist registry data across pod restarts (file position and other metadata).
hostPathRoot: /var/lib

image: "docker.elastic.co/beats/filebeat"
imageTag: "7.0.0"
imagePullPolicy: "IfNotPresent"
imagePullSecrets: []

livenessProbe:
  failureThreshold: 3
  initialDelaySeconds: 10
  periodSeconds: 10
  timeoutSeconds: 5

# Whether this chart should self-manage its service account, role, and associated role binding.
managedServiceAccount: true

podAnnotations: {}
  # iam.amazonaws.com/role: es-cluster

# Various pod security context settings. Bear in mind that many of these have an impact on Filebeat functioning properly.
#
# - Filesystem group for the Filebeat user. The official elastic docker images always have an id of 1000.
# - User that the container will execute as. Typically necessary to run as root (0) in order to properly collect host container logs.
# - Whether to execute the Filebeat containers as privileged containers. Typically not necessarily unless running within environments such as OpenShift.
podSecurityContext:
  fsGroup: 1000
  runAsUser: 0
  privileged: false

resources:
  requests:
    cpu: "100m"
    memory: "100Mi"
  limits:
    cpu: "1000m"
    memory: "200Mi"

# Service account that the pod will use
serviceAccount: filebeat

# A list of secrets and their paths to mount inside the pod
# This is useful for mounting certificates for security other sensitive values
secretMounts: []
#  - name: filebeat-certificates
#    secretName: filebeat-certificates
#    path: /usr/share/filebeat/certs

# How long to wait for Filebeat pods to stop gracefully
terminationGracePeriod: 30

tolerations: []

updateStrategy: RollingUpdate

# Override various naming aspects of this chart
# Only edit these if you know what you're doing
nameOverride: ""
fullnameOverride: ""
```

---
Autogenerated from Helm Chart and git history using [helm-changelog](https://github.com/mogensen/helm-changelog)
