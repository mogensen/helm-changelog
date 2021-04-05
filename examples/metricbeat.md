# Change Log

## Next Release 

![AppVersion: 8.0.0-SNAPSHOT](https://img.shields.io/static/v1?label=AppVersion&message=8.0.0-SNAPSHOT&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [elasticsearch][kibana] remove oss examples (#1046) 
* [metricbeat] Fixing the respository of kube-state-metrics for metricbeats (#1039) 
* [meta] add build status and artifact hub badges (#1033) 
* [meta] fix transient errors with stable repository (#1018) 
* [filebeat][metricbeat] Update documentation on port collisions for multiple beats agents with hostNetworking enabled. (#997) 
* [all] add hostaliases (#970) 
* [meta] enable metricbeat upgrade test (#940) 
* [meta] stabilize CI tests (#935) 
* [meta] remove support for k8s <1.14 & helm <2.17.0 (#916) 
* [meta] upgrade test (#907) 
* Helm 3 (#516) 
* [meta] increase helm timeout (#891) 
* [meta] update rbac.authorization.k8s.io api (#890) 
* Add warning comment placeholder (main branch) (#886) 
* [metricbeat] use relocated stable repo for kube-state-metrics (#882) 
* Update warning on charts readme (#863) 
* [Metricbeat] Dont generate config if not enabled (#767) 
* [metricbeat] support deployment/daemonset specific metrics (#820) 
* [meta] update changelog for 7.9.2 release  (#824) 
* [metricbeat] Support secrets (#778) 
* [metricbeat] Add missing labels for deployment (#770) 
* [meta] Update changelog for 6.8.12 / 7.9.0 releases (#789) 
* [meta] add helm 3 beta support (#759) 
* [meta] update changelog for 7.8.1 and 6.8.11 releases (#755) 
* [doc] fix some links (#737) 
* Make kube-state-metrics optional for installation (#387) 
* Metricbeat daemonset deployment optional (#716) 
* feat(metricbeat): annotation support for daemonset and deployment (#713) 

### Default value changes

```diff
diff --git a/metricbeat/values.yaml b/metricbeat/values.yaml
index ad670fdb..02f03d74 100755
--- a/metricbeat/values.yaml
+++ b/metricbeat/values.yaml
@@ -1,7 +1,13 @@
 ---
 
 daemonset:
+  # Annotations to apply to the daemonset
+  annotations: {}
+  # additionals labels
+  labels: {}
   affinity: {}
+  # Include the daemonset
+  enabled: true
   # Extra environment variables for Metricbeat container.
   envFrom: []
   # - configMapRef:
@@ -16,6 +22,11 @@ daemonset:
   # - name: extras
   #   mountPath: /usr/share/extras
   #   readOnly: true
+  hostAliases: []
+  #- ip: "127.0.0.1"
+  #  hostnames:
+  #  - "foo.local"
+  #  - "bar.local"
   hostNetworking: false
   # Allows you to add any config files in /usr/share/metricbeat
   # such as metricbeat.yml for daemonset
@@ -91,7 +102,13 @@ daemonset:
   tolerations: []
 
 deployment:
+  # Annotations to apply to the deployment
+  annotations: {}
+  # additionals labels
+  labels: {}
   affinity: {}
+  # Include the deployment
+  enabled: true
   # Extra environment variables for Metricbeat container.
   envFrom: []
   # - configMapRef:
@@ -108,6 +125,11 @@ deployment:
   #   mountPath: /usr/share/extras
   #   readOnly: true
   # such as metricbeat.yml for deployment
+  hostAliases: []
+  #- ip: "127.0.0.1"
+  #  hostnames:
+  #  - "foo.local"
+  #  - "bar.local"
   metricbeatConfig:
     metricbeat.yml: |
       metricbeat.modules:
@@ -189,9 +211,6 @@ readinessProbe:
   periodSeconds: 10
   timeoutSeconds: 5
 
-# additionals labels
-labels: {}
-
 # Whether this chart should self-manage its service account, role, and associated role binding.
 managedServiceAccount: true
 
@@ -242,6 +261,27 @@ updateStrategy: RollingUpdate
 nameOverride: ""
 fullnameOverride: ""
 
+kube_state_metrics:
+  enabled: true
+  # host is used only when kube_state_metrics.enabled: false
+  host: ""
+
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
 # DEPRECATED
 affinity: {}
 envFrom: []
@@ -256,3 +296,4 @@ podSecurityContext: {}
 resources: {}
 secretMounts: []
 tolerations: []
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
diff --git a/metricbeat/values.yaml b/metricbeat/values.yaml
index 59b03239..ad670fdb 100755
--- a/metricbeat/values.yaml
+++ b/metricbeat/values.yaml
@@ -159,7 +159,7 @@ extraInitContainers: ""
 hostPathRoot: /var/lib
 
 image: "docker.elastic.co/beats/metricbeat"
-imageTag: "7.7.1"
+imageTag: "8.0.0-SNAPSHOT"
 imagePullPolicy: "IfNotPresent"
 imagePullSecrets: []
 
@@ -224,6 +224,10 @@ podAnnotations: {}
 # Custom service account override that the pod will use
 serviceAccount: ""
 
+# Annotations to add to the ServiceAccount that is created if the serviceAccount value isn't set.
+serviceAccountAnnotations: {}
+  # eks.amazonaws.com/role-arn: arn:aws:iam::111111111111:role/k8s.clustername.namespace.serviceaccount
+
 # How long to wait for metricbeat pods to stop gracefully
 terminationGracePeriod: 30
 
```

## 7.7.1 

**Release date:** 2020-06-03

![AppVersion: 7.7.1](https://img.shields.io/static/v1?label=AppVersion&message=7.7.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Update changelog for 7.7.1 and 6.8.10 releases (#652) 
* [metricbeat] use deprecated kube-state-metrics config when existing (#624) 
* Fix configchecksum not being set (#634) 
* [metricbeat] fix deployment upgrade by removing chart label from .spec.selector.matchLabels (#622) 

### Default value changes

```diff
diff --git a/metricbeat/values.yaml b/metricbeat/values.yaml
index d89a947d..59b03239 100755
--- a/metricbeat/values.yaml
+++ b/metricbeat/values.yaml
@@ -159,7 +159,7 @@ extraInitContainers: ""
 hostPathRoot: /var/lib
 
 image: "docker.elastic.co/beats/metricbeat"
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
* [metricbeat] update default values 
* [metricbeat] format README 
* [metricbeat] add host networking option 
* [metricbeat] nit - add missing trailing lines 
* [metricbeat] remove unneeded if statement for envFrom 
* [metricbeat] split secretMounts for daemonset and deployment 
* [metricbeat] split volumes and volumeMounts for daemonset and deployment 
* [metricbeat] split envFrom for daemonset and deployment 
* metricbeat] split extraEnvs for daemonset and deployment 
* [metricbeat] nit - fix default values in README 
* [metricbeat] split security context for daemonset and deployment 
* [metricbeat] nit 
* [metricbeat] split tolerations for daemonset and deployment 
* [metricbeat] split nodeSelector for daemonset and deployment 
* [metricbeat] split resources for daemonset and deployment 
* [metricbeat] split affinity for daemonset and deployment 
* [metricbeat] use markdown implicit link name for config table 
* [metricbeat] split configmap for daemonset and deployment 
* fixup! [metricbeat] remove unused /var/lib/docker/container mount 
* [metricbeat] add replicasets.apps to clusterRoleRules 
* [metricbeat] remove unused /var/lib/docker/container mount 
* fixup! [metricbeat] update clusterRoleRules with recommended values 
* [metricbeat] update clusterRoleRules with recommended values 
* [metricbeat] remove in_cluster config from add_kubernetes_metadata 
* Add nodes/stats clusterRoleRules 
* Replace HOSTNAME with NODE_NAME to access host without hostNetwork 
* Merge remote-tracking branch 'upstream/master' into patch-2 

### Default value changes

```diff
diff --git a/metricbeat/values.yaml b/metricbeat/values.yaml
index 141b02c3..d89a947d 100755
--- a/metricbeat/values.yaml
+++ b/metricbeat/values.yaml
@@ -1,67 +1,148 @@
 ---
-# Allows you to add any config files in /usr/share/metricbeat
-# such as metricbeat.yml
-metricbeatConfig:
-  metricbeat.yml: |
-    metricbeat.modules:
-    - module: kubernetes
-      metricsets:
-        - container
-        - node
-        - pod
-        - system
-        - volume
-      period: 10s
-      host: "${NODE_NAME}"
-      hosts: ["${NODE_NAME}:10255"]
-      processors:
-      - add_kubernetes_metadata:
-          in_cluster: true
-    - module: kubernetes
-      enabled: true
-      metricsets:
-        - event
-    - module: system
-      period: 10s
-      metricsets:
-        - cpu
-        - load
-        - memory
-        - network
-        - process
-        - process_summary
-      processes: ['.*']
-      process.include_top_n:
-        by_cpu: 5
-        by_memory: 5
-    - module: system
-      period: 1m
-      metricsets:
-        - filesystem
-        - fsstat
-      processors:
-      - drop_event.when.regexp:
-          system.filesystem.mount_point: '^/(sys|cgroup|proc|dev|etc|host|lib)($|/)'
-    output.elasticsearch:
-      hosts: '${ELASTICSEARCH_HOSTS:elasticsearch-master:9200}'
-
-  kube-state-metrics-metricbeat.yml: |
-    metricbeat.modules:
-    - module: kubernetes
-      enabled: true
-      metricsets:
-        - state_node
-        - state_deployment
-        - state_replicaset
-        - state_pod
-        - state_container
-      period: 10s
-      hosts: ["${KUBE_STATE_METRICS_HOSTS}"]
-    output.elasticsearch:
-      hosts: '${ELASTICSEARCH_HOSTS:elasticsearch-master:9200}'
 
-# Replicas being used for the kube-state-metrics metricbeat deployment
+daemonset:
+  affinity: {}
+  # Extra environment variables for Metricbeat container.
+  envFrom: []
+  # - configMapRef:
+  #     name: config-secret
+  extraEnvs: []
+  #  - name: MY_ENVIRONMENT_VAR
+  #    value: the_value_goes_here
+  extraVolumes: []
+  # - name: extras
+  #   emptyDir: {}
+  extraVolumeMounts: []
+  # - name: extras
+  #   mountPath: /usr/share/extras
+  #   readOnly: true
+  hostNetworking: false
+  # Allows you to add any config files in /usr/share/metricbeat
+  # such as metricbeat.yml for daemonset
+  metricbeatConfig:
+    metricbeat.yml: |
+      metricbeat.modules:
+      - module: kubernetes
+        metricsets:
+          - container
+          - node
+          - pod
+          - system
+          - volume
+        period: 10s
+        host: "${NODE_NAME}"
+        hosts: ["https://${NODE_NAME}:10250"]
+        bearer_token_file: /var/run/secrets/kubernetes.io/serviceaccount/token
+        ssl.verification_mode: "none"
+        # If using Red Hat OpenShift remove ssl.verification_mode entry and
+        # uncomment these settings:
+        #ssl.certificate_authorities:
+          #- /var/run/secrets/kubernetes.io/serviceaccount/service-ca.crt
+        processors:
+        - add_kubernetes_metadata: ~
+      - module: kubernetes
+        enabled: true
+        metricsets:
+          - event
+      - module: system
+        period: 10s
+        metricsets:
+          - cpu
+          - load
+          - memory
+          - network
+          - process
+          - process_summary
+        processes: ['.*']
+        process.include_top_n:
+          by_cpu: 5
+          by_memory: 5
+      - module: system
+        period: 1m
+        metricsets:
+          - filesystem
+          - fsstat
+        processors:
+        - drop_event.when.regexp:
+            system.filesystem.mount_point: '^/(sys|cgroup|proc|dev|etc|host|lib)($|/)'
+      output.elasticsearch:
+        hosts: '${ELASTICSEARCH_HOSTS:elasticsearch-master:9200}'
+  nodeSelector: {}
+  # A list of secrets and their paths to mount inside the pod
+  # This is useful for mounting certificates for security other sensitive values
+  secretMounts: []
+  #  - name: metricbeat-certificates
+  #    secretName: metricbeat-certificates
+  #    path: /usr/share/metricbeat/certs
+  # Various pod security context settings. Bear in mind that many of these have an impact on metricbeat functioning properly.
+  # - Filesystem group for the metricbeat user. The official elastic docker images always have an id of 1000.
+  # - User that the container will execute as. Typically necessary to run as root (0) in order to properly collect host container logs.
+  # - Whether to execute the metricbeat containers as privileged containers. Typically not necessarily unless running within environments such as OpenShift.
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
+  affinity: {}
+  # Extra environment variables for Metricbeat container.
+  envFrom: []
+  # - configMapRef:
+  #     name: config-secret
+  extraEnvs: []
+  #  - name: MY_ENVIRONMENT_VAR
+  #    value: the_value_goes_here
+  # Allows you to add any config files in /usr/share/metricbeat
+  extraVolumes: []
+  # - name: extras
+  #   emptyDir: {}
+  extraVolumeMounts: []
+  # - name: extras
+  #   mountPath: /usr/share/extras
+  #   readOnly: true
+  # such as metricbeat.yml for deployment
+  metricbeatConfig:
+    metricbeat.yml: |
+      metricbeat.modules:
+      - module: kubernetes
+        enabled: true
+        metricsets:
+          - state_node
+          - state_deployment
+          - state_replicaset
+          - state_pod
+          - state_container
+        period: 10s
+        hosts: ["${KUBE_STATE_METRICS_HOSTS}"]
+      output.elasticsearch:
+        hosts: '${ELASTICSEARCH_HOSTS:elasticsearch-master:9200}'
+  nodeSelector: {}
+  # A list of secrets and their paths to mount inside the pod
+  # This is useful for mounting certificates for security other sensitive values
+  secretMounts: []
+  #  - name: metricbeat-certificates
+  #    secretName: metricbeat-certificates
+  #    path: /usr/share/metricbeat/certs
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
 
+# Replicas being used for the kube-state-metrics metricbeat deployment
 replicas: 1
 
 extraContainers: ""
@@ -74,31 +155,11 @@ extraInitContainers: ""
 #   image: busybox
 #   command: ['echo', 'hey']
 
-# Extra environment variables to append to the DaemonSet pod spec.
-# This will be appended to the current 'env:' key. You can use any of the kubernetes env
-# syntax here
-extraEnvs: []
-#  - name: MY_ENVIRONMENT_VAR
-#    value: the_value_goes_here
-
-extraVolumeMounts: []
-  # - name: extras
-  #   mountPath: /usr/share/extras
-  #   readOnly: true
-
-extraVolumes: []
-  # - name: extras
-  #   emptyDir: {}
-
-envFrom: []
-  # - configMapRef:
-  #     name: config-secret
-
 # Root directory where metricbeat will write data to in order to persist registry data across pod restarts (file position and other metadata).
 hostPathRoot: /var/lib
 
 image: "docker.elastic.co/beats/metricbeat"
-imageTag: "7.6.2"
+imageTag: "7.7.0"
 imagePullPolicy: "IfNotPresent"
 imagePullSecrets: []
 
@@ -135,61 +196,37 @@ labels: {}
 managedServiceAccount: true
 
 clusterRoleRules:
-  - apiGroups:
-    - "extensions"
-    - "apps"
-    - ""
-    resources:
-    - namespaces
-    - pods
-    - events
-    - deployments
-    - nodes
-    - replicasets
-    verbs:
-    - get
-    - list
-    - watch
+- apiGroups: [""]
+  resources:
+  - nodes
+  - namespaces
+  - events
+  - pods
+  verbs: ["get", "list", "watch"]
+- apiGroups: ["extensions"]
+  resources:
+  - replicasets
+  verbs: ["get", "list", "watch"]
+- apiGroups: ["apps"]
+  resources:
+  - statefulsets
+  - deployments
+  - replicasets
+  verbs: ["get", "list", "watch"]
+- apiGroups: [""]
+  resources:
+  - nodes/stats
+  verbs: ["get"]
 
 podAnnotations: {}
   # iam.amazonaws.com/role: es-cluster
 
-# Various pod security context settings. Bear in mind that many of these have an impact on metricbeat functioning properly.
-#
-# - Filesystem group for the metricbeat user. The official elastic docker images always have an id of 1000.
-# - User that the container will execute as. Typically necessary to run as root (0) in order to properly collect host container logs.
-# - Whether to execute the metricbeat containers as privileged containers. Typically not necessarily unless running within environments such as OpenShift.
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
 
-# A list of secrets and their paths to mount inside the pod
-# This is useful for mounting certificates for security other sensitive values
-secretMounts: []
-#  - name: metricbeat-certificates
-#    secretName: metricbeat-certificates
-#    path: /usr/share/metricbeat/certs
-
 # How long to wait for metricbeat pods to stop gracefully
 terminationGracePeriod: 30
 
-tolerations: []
-
-nodeSelector: {}
-
-affinity: {}
-
 # This is the PriorityClass settings as defined in
 # https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption/#priorityclass
 priorityClassName: ""
@@ -200,3 +237,18 @@ updateStrategy: RollingUpdate
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
+# Allows you to add any config files in /usr/share/metricbeat
+# such as metricbeat.yml for both daemonset and deployment
+metricbeatConfig: {}
+nodeSelector: {}
+podSecurityContext: {}
+resources: {}
+secretMounts: []
+tolerations: []
```

## 7.6.2 

**Release date:** 2020-03-31

![AppVersion: 7.6.2](https://img.shields.io/static/v1?label=AppVersion&message=7.6.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [7.6.2] bump version 

### Default value changes

```diff
diff --git a/metricbeat/values.yaml b/metricbeat/values.yaml
index 53032198..141b02c3 100755
--- a/metricbeat/values.yaml
+++ b/metricbeat/values.yaml
@@ -98,7 +98,7 @@ envFrom: []
 hostPathRoot: /var/lib
 
 image: "docker.elastic.co/beats/metricbeat"
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
diff --git a/metricbeat/values.yaml b/metricbeat/values.yaml
index 3ea4af31..53032198 100755
--- a/metricbeat/values.yaml
+++ b/metricbeat/values.yaml
@@ -98,7 +98,7 @@ envFrom: []
 hostPathRoot: /var/lib
 
 image: "docker.elastic.co/beats/metricbeat"
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
* [metricbeat] add extra init containers 
* [metricbeat]  add extracontainers 
* Update values.yaml 
* Update values.yaml 
* Make use of secure port when accessing Kubelet API 

### Default value changes

```diff
diff --git a/metricbeat/values.yaml b/metricbeat/values.yaml
index 4e34369e..3ea4af31 100755
--- a/metricbeat/values.yaml
+++ b/metricbeat/values.yaml
@@ -64,6 +64,16 @@ metricbeatConfig:
 
 replicas: 1
 
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
 # Extra environment variables to append to the DaemonSet pod spec.
 # This will be appended to the current 'env:' key. You can use any of the kubernetes env
 # syntax here
@@ -88,7 +98,7 @@ envFrom: []
 hostPathRoot: /var/lib
 
 image: "docker.elastic.co/beats/metricbeat"
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
* added priorityClassName in metrics deployment 
* Merge branch 'master' into master 
* [metricbeat] typo 
* [doc] switch relative URLs to absolute URLs 
* [metricbeat] add test for priorityClassName 
* [metricbeat] add doc for priorityClassName 
* [metricbeat] Add priorityClassName config 
* Merge pull request #425 from pbecotte/424 
* update hostfs to be a CLI option instead of a config option 
* Merge pull request #420 from jmlrt/override-probe 
* forgot end block. doh 
* [metricbeat] allow override readiness and liveness probes commands 
* Merge branch 'master' into add-support-for-envfrom 

### Default value changes

```diff
diff --git a/metricbeat/values.yaml b/metricbeat/values.yaml
index 4e8b501f..4e34369e 100755
--- a/metricbeat/values.yaml
+++ b/metricbeat/values.yaml
@@ -3,8 +3,6 @@
 # such as metricbeat.yml
 metricbeatConfig:
   metricbeat.yml: |
-    system:
-      hostfs: /hostfs
     metricbeat.modules:
     - module: kubernetes
       metricsets:
@@ -82,21 +80,39 @@ extraVolumes: []
   # - name: extras
   #   emptyDir: {}
 
+envFrom: []
+  # - configMapRef:
+  #     name: config-secret
+
 # Root directory where metricbeat will write data to in order to persist registry data across pod restarts (file position and other metadata).
 hostPathRoot: /var/lib
 
 image: "docker.elastic.co/beats/metricbeat"
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
+        metricbeat test output
   failureThreshold: 3
   initialDelaySeconds: 10
   periodSeconds: 10
@@ -164,6 +180,10 @@ nodeSelector: {}
 
 affinity: {}
 
+# This is the PriorityClass settings as defined in
+# https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption/#priorityclass
+priorityClassName: ""
+
 updateStrategy: RollingUpdate
 
 # Override various naming aspects of this chart
```

## 7.5.1 

**Release date:** 2019-12-19

![AppVersion: 7.5.1](https://img.shields.io/static/v1?label=AppVersion&message=7.5.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Merge pull request #415 from jmlrt/add-custom-labels-to-pods 

### Default value changes

```diff
diff --git a/metricbeat/values.yaml b/metricbeat/values.yaml
index 47087591..4e8b501f 100755
--- a/metricbeat/values.yaml
+++ b/metricbeat/values.yaml
@@ -86,7 +86,7 @@ extraVolumes: []
 hostPathRoot: /var/lib
 
 image: "docker.elastic.co/beats/metricbeat"
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


* [metricbeat] apply labels to all pods 

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
* [metricbeat] add notice about breaking change 

### Default value changes

```diff
diff --git a/metricbeat/values.yaml b/metricbeat/values.yaml
index 47087591..4e8b501f 100755
--- a/metricbeat/values.yaml
+++ b/metricbeat/values.yaml
@@ -86,7 +86,7 @@ extraVolumes: []
 hostPathRoot: /var/lib
 
 image: "docker.elastic.co/beats/metricbeat"
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
* Merge pull request #314 from pbecotte/ExtraMounts 
* Merge branch 'master' into issue_267 
* Merge pull request #352 from masterkain/master 
* change ordering 
* readme tweak 
* tweaks to name and test 
* tests 
* update README 
* added support for envFrom 
* update install doc 
* bump metricbeat to latest chart and app version 
* [metricbeat] fix beats icon 
* [helm] add some notice about tested helm version 

### Default value changes

```diff
diff --git a/metricbeat/values.yaml b/metricbeat/values.yaml
index fea0deaf..47087591 100755
--- a/metricbeat/values.yaml
+++ b/metricbeat/values.yaml
@@ -73,12 +73,12 @@ extraEnvs: []
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
 
@@ -86,7 +86,7 @@ extraVolumes: ""
 hostPathRoot: /var/lib
 
 image: "docker.elastic.co/beats/metricbeat"
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
* Merge pull request #338 from jmlrt/helm-2-15 
* [helm] add chart api version This is required witrh Helm 2.15.0 following https://github.com/helm/helm/pull/6180 
* [metricbeat] add pod labels 

### Default value changes

```diff
diff --git a/metricbeat/values.yaml b/metricbeat/values.yaml
index 321c85d4..fea0deaf 100755
--- a/metricbeat/values.yaml
+++ b/metricbeat/values.yaml
@@ -86,7 +86,7 @@ extraVolumes: ""
 hostPathRoot: /var/lib
 
 image: "docker.elastic.co/beats/metricbeat"
-imageTag: "7.4.0"
+imageTag: "7.4.1"
 imagePullPolicy: "IfNotPresent"
 imagePullSecrets: []
 
@@ -102,6 +102,9 @@ readinessProbe:
   periodSeconds: 10
   timeoutSeconds: 5
 
+# additionals labels
+labels: {}
+
 # Whether this chart should self-manage its service account, role, and associated role binding.
 managedServiceAccount: true
 
```

## 7.4.0 

**Release date:** 2019-10-15

![AppVersion: 7.4.0](https://img.shields.io/static/v1?label=AppVersion&message=7.4.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* [Issue #267] Support fullnameOverride in kibana, filebeat, metricbeat charts 

### Default value changes

```diff
diff --git a/metricbeat/values.yaml b/metricbeat/values.yaml
index 5c852497..321c85d4 100755
--- a/metricbeat/values.yaml
+++ b/metricbeat/values.yaml
@@ -86,7 +86,7 @@ extraVolumes: ""
 hostPathRoot: /var/lib
 
 image: "docker.elastic.co/beats/metricbeat"
-imageTag: "7.3.2"
+imageTag: "7.4.0"
 imagePullPolicy: "IfNotPresent"
 imagePullSecrets: []
 
@@ -105,6 +105,23 @@ readinessProbe:
 # Whether this chart should self-manage its service account, role, and associated role binding.
 managedServiceAccount: true
 
+clusterRoleRules:
+  - apiGroups:
+    - "extensions"
+    - "apps"
+    - ""
+    resources:
+    - namespaces
+    - pods
+    - events
+    - deployments
+    - nodes
+    - replicasets
+    verbs:
+    - get
+    - list
+    - watch
+
 podAnnotations: {}
   # iam.amazonaws.com/role: es-cluster
 
```

## 7.3.2 

**Release date:** 2019-10-04

![AppVersion: 7.3.2](https://img.shields.io/static/v1?label=AppVersion&message=7.3.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Had to add a couple extra mounts to get MetricBeat to pick up all the metrics from the host nodes on Digital Ocean. In the process, noticed that the extraMount setting is configured different then the rest- changed it to be consistent 

### Default value changes

```diff
diff --git a/metricbeat/values.yaml b/metricbeat/values.yaml
index 5c852497..77fea733 100755
--- a/metricbeat/values.yaml
+++ b/metricbeat/values.yaml
@@ -73,12 +73,12 @@ extraEnvs: []
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
 
```

## 7.4.0 

**Release date:** 2019-10-01

![AppVersion: 7.4.0](https://img.shields.io/static/v1?label=AppVersion&message=7.4.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* 7.4.0 Release 
* [metricbeat] Make cluster role rules configurable 
* [metricbeat] change min k8s version due to used apiVersions 

### Default value changes

```diff
diff --git a/metricbeat/values.yaml b/metricbeat/values.yaml
index 5c852497..321c85d4 100755
--- a/metricbeat/values.yaml
+++ b/metricbeat/values.yaml
@@ -86,7 +86,7 @@ extraVolumes: ""
 hostPathRoot: /var/lib
 
 image: "docker.elastic.co/beats/metricbeat"
-imageTag: "7.3.2"
+imageTag: "7.4.0"
 imagePullPolicy: "IfNotPresent"
 imagePullSecrets: []
 
@@ -105,6 +105,23 @@ readinessProbe:
 # Whether this chart should self-manage its service account, role, and associated role binding.
 managedServiceAccount: true
 
+clusterRoleRules:
+  - apiGroups:
+    - "extensions"
+    - "apps"
+    - ""
+    resources:
+    - namespaces
+    - pods
+    - events
+    - deployments
+    - nodes
+    - replicasets
+    verbs:
+    - get
+    - list
+    - watch
+
 podAnnotations: {}
   # iam.amazonaws.com/role: es-cluster
 
```

## 7.3.2 

**Release date:** 2019-09-19

![AppVersion: 7.3.2](https://img.shields.io/static/v1?label=AppVersion&message=7.3.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* 7.3.2 Release 
* Merge pull request #254 from Azuka/feature/metrics-cluster-role-events 
* Merge branch 'master' into feature/metrics-cluster-role-events 
* [metricbeat] Remove default kube static metrics host to avoid confusion 
* Fix octal literal to work in both Python 2 and Python 3 
* [metricbeat] Enable events access to cluster role 
* [metricbeat] Fix default configuration for kubernetes module 
* [filebeat][metricbeat] Add configurable nodeSelector and affinit… (#243) 
* [filebeat][metricbeat] Add configurable nodeSelector and affinity spec 
* Update documentation and defaults for tmpl values 

### Default value changes

```diff
diff --git a/metricbeat/values.yaml b/metricbeat/values.yaml
index 29996d3f..5c852497 100755
--- a/metricbeat/values.yaml
+++ b/metricbeat/values.yaml
@@ -14,7 +14,8 @@ metricbeatConfig:
         - system
         - volume
       period: 10s
-      hosts: ["localhost:10255"]
+      host: "${NODE_NAME}"
+      hosts: ["${NODE_NAME}:10255"]
       processors:
       - add_kubernetes_metadata:
           in_cluster: true
@@ -57,7 +58,7 @@ metricbeatConfig:
         - state_pod
         - state_container
       period: 10s
-      hosts: ["${KUBE_STATE_METRICS_HOSTS:kube-state-metrics:8080}"]
+      hosts: ["${KUBE_STATE_METRICS_HOSTS}"]
     output.elasticsearch:
       hosts: '${ELASTICSEARCH_HOSTS:elasticsearch-master:9200}'
 
@@ -72,12 +73,12 @@ extraEnvs: []
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
 
@@ -85,7 +86,7 @@ extraVolumes: []
 hostPathRoot: /var/lib
 
 image: "docker.elastic.co/beats/metricbeat"
-imageTag: "7.3.0"
+imageTag: "7.3.2"
 imagePullPolicy: "IfNotPresent"
 imagePullSecrets: []
 
@@ -139,6 +140,10 @@ terminationGracePeriod: 30
 
 tolerations: []
 
+nodeSelector: {}
+
+affinity: {}
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
diff --git a/metricbeat/values.yaml b/metricbeat/values.yaml
index 7cae07ef..29996d3f 100755
--- a/metricbeat/values.yaml
+++ b/metricbeat/values.yaml
@@ -85,7 +85,7 @@ extraVolumes: []
 hostPathRoot: /var/lib
 
 image: "docker.elastic.co/beats/metricbeat"
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
* Remove data mount from the kube-state-metrics metricbeat 
* Update file pattern to include all goss tests 
* Update readme and backport fixes from filebeat chart 

### Default value changes

```diff
diff --git a/metricbeat/values.yaml b/metricbeat/values.yaml
index 2d9e85a8..7cae07ef 100755
--- a/metricbeat/values.yaml
+++ b/metricbeat/values.yaml
@@ -113,7 +113,6 @@ podAnnotations: {}
 # - User that the container will execute as. Typically necessary to run as root (0) in order to properly collect host container logs.
 # - Whether to execute the metricbeat containers as privileged containers. Typically not necessarily unless running within environments such as OpenShift.
 podSecurityContext:
-  fsGroup: 1000
   runAsUser: 0
   privileged: false
 
```

## 7.2.0 

**Release date:** 2019-07-01

![AppVersion: 7.2.0](https://img.shields.io/static/v1?label=AppVersion&message=7.2.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* Rebase on master and bump to latest version 
* Add security example 
* Add oss example 
* Add 6.x example 
* Add tests to make sure that data is flowing in 
* Update goss testing to support testing multiple containers 
* Fix up linting and tests 
* Update description and beta disclaimer 
* Copy in all of the missing fields from the daemonset 

### Default value changes

```diff
diff --git a/metricbeat/values.yaml b/metricbeat/values.yaml
index 145821bc..2d9e85a8 100755
--- a/metricbeat/values.yaml
+++ b/metricbeat/values.yaml
@@ -46,8 +46,7 @@ metricbeatConfig:
     output.elasticsearch:
       hosts: '${ELASTICSEARCH_HOSTS:elasticsearch-master:9200}'
 
-metricbeatKubeStateMetricsConfig:
-  metricbeat.yml: |
+  kube-state-metrics-metricbeat.yml: |
     metricbeat.modules:
     - module: kubernetes
       enabled: true
@@ -62,6 +61,10 @@ metricbeatKubeStateMetricsConfig:
     output.elasticsearch:
       hosts: '${ELASTICSEARCH_HOSTS:elasticsearch-master:9200}'
 
+# Replicas being used for the kube-state-metrics metricbeat deployment
+
+replicas: 1
+
 # Extra environment variables to append to the DaemonSet pod spec.
 # This will be appended to the current 'env:' key. You can use any of the kubernetes env
 # syntax here
@@ -82,7 +85,7 @@ extraVolumes: []
 hostPathRoot: /var/lib
 
 image: "docker.elastic.co/beats/metricbeat"
-imageTag: "7.0.1"
+imageTag: "7.2.0"
 imagePullPolicy: "IfNotPresent"
 imagePullSecrets: []
 
```

## 7.0.1-alpha1 

**Release date:** 2019-05-09

![AppVersion: 7.0.1](https://img.shields.io/static/v1?label=AppVersion&message=7.0.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)


* WIP: Add metricbeat chart 

### Default value changes

```diff
---
# Allows you to add any config files in /usr/share/metricbeat
# such as metricbeat.yml
metricbeatConfig:
  metricbeat.yml: |
    system:
      hostfs: /hostfs
    metricbeat.modules:
    - module: kubernetes
      metricsets:
        - container
        - node
        - pod
        - system
        - volume
      period: 10s
      hosts: ["localhost:10255"]
      processors:
      - add_kubernetes_metadata:
          in_cluster: true
    - module: kubernetes
      enabled: true
      metricsets:
        - event
    - module: system
      period: 10s
      metricsets:
        - cpu
        - load
        - memory
        - network
        - process
        - process_summary
      processes: ['.*']
      process.include_top_n:
        by_cpu: 5
        by_memory: 5
    - module: system
      period: 1m
      metricsets:
        - filesystem
        - fsstat
      processors:
      - drop_event.when.regexp:
          system.filesystem.mount_point: '^/(sys|cgroup|proc|dev|etc|host|lib)($|/)'
    output.elasticsearch:
      hosts: '${ELASTICSEARCH_HOSTS:elasticsearch-master:9200}'

metricbeatKubeStateMetricsConfig:
  metricbeat.yml: |
    metricbeat.modules:
    - module: kubernetes
      enabled: true
      metricsets:
        - state_node
        - state_deployment
        - state_replicaset
        - state_pod
        - state_container
      period: 10s
      hosts: ["${KUBE_STATE_METRICS_HOSTS:kube-state-metrics:8080}"]
    output.elasticsearch:
      hosts: '${ELASTICSEARCH_HOSTS:elasticsearch-master:9200}'

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

# Root directory where metricbeat will write data to in order to persist registry data across pod restarts (file position and other metadata).
hostPathRoot: /var/lib

image: "docker.elastic.co/beats/metricbeat"
imageTag: "7.0.1"
imagePullPolicy: "IfNotPresent"
imagePullSecrets: []

livenessProbe:
  failureThreshold: 3
  initialDelaySeconds: 10
  periodSeconds: 10
  timeoutSeconds: 5

readinessProbe:
  failureThreshold: 3
  initialDelaySeconds: 10
  periodSeconds: 10
  timeoutSeconds: 5

# Whether this chart should self-manage its service account, role, and associated role binding.
managedServiceAccount: true

podAnnotations: {}
  # iam.amazonaws.com/role: es-cluster

# Various pod security context settings. Bear in mind that many of these have an impact on metricbeat functioning properly.
#
# - Filesystem group for the metricbeat user. The official elastic docker images always have an id of 1000.
# - User that the container will execute as. Typically necessary to run as root (0) in order to properly collect host container logs.
# - Whether to execute the metricbeat containers as privileged containers. Typically not necessarily unless running within environments such as OpenShift.
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

# Custom service account override that the pod will use
serviceAccount: ""

# A list of secrets and their paths to mount inside the pod
# This is useful for mounting certificates for security other sensitive values
secretMounts: []
#  - name: metricbeat-certificates
#    secretName: metricbeat-certificates
#    path: /usr/share/metricbeat/certs

# How long to wait for metricbeat pods to stop gracefully
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
