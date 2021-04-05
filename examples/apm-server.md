# Change Log

## Next Release 

![AppVersion: 8.0.0-SNAPSHOT](https://img.shields.io/static/v1?label=AppVersion&message=8.0.0-SNAPSHOT&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [apm-server] Add  option loadBalancerIP to service (#1075) 
* [elasticsearch][kibana] remove oss examples (#1046) 
* [meta] add build status and artifact hub badges (#1033) 
* [apm-server] run as non root user (#996) 
* [all] add hostaliases (#970) 
* [apm-server] Add missing fields to HPA (#782) 
* [meta] stabilize CI tests (#935) 
* [meta] remove support for k8s <1.14 & helm <2.17.0 (#916) 
* [meta] upgrade test (#907) 
* Helm 3 (#516) 
* [meta] increase helm timeout (#891) 
* [meta] update rbac.authorization.k8s.io api (#890) 
* Add warning comment placeholder (main branch) (#886) 
* Update warning on charts readme (#863) 
* [meta] update changelog for 7.9.2 release  (#824) 
* Fix serviceAccount for APM server (#786) 
* Remove duplicate "initialDelaySeconds" field (#763) 
* [meta] Update changelog for 6.8.12 / 7.9.0 releases (#789) 
* [meta] add helm 3 beta support (#759) 
* [meta] update changelog for 7.8.1 and 6.8.11 releases (#755) 
* [doc] fix some links (#737) 

### Default value changes

```diff
diff --git a/apm-server/values.yaml b/apm-server/values.yaml
index 55b3d06a..fa8a2e95 100755
--- a/apm-server/values.yaml
+++ b/apm-server/values.yaml
@@ -61,6 +61,12 @@ extraVolumes: []
   # - name: extras
   #   emptyDir: {}
 
+hostAliases: []
+#- ip: "127.0.0.1"
+#  hostnames:
+#  - "foo.local"
+#  - "bar.local"
+
 image: "docker.elastic.co/apm/apm-server"
 imageTag: "8.0.0-SNAPSHOT"
 imagePullPolicy: "IfNotPresent"
@@ -76,8 +82,15 @@ podAnnotations: {}
 labels: {}
 
 podSecurityContext:
-  runAsUser: 0
+  fsGroup: 1000
+  runAsUser: 1000
+  runAsGroup: 0
+
+securityContext:
   privileged: false
+  runAsNonRoot: true
+  runAsUser: 1000
+  runAsGroup: 0
 
 livenessProbe:
   httpGet:
@@ -85,7 +98,6 @@ livenessProbe:
     port: http
   initialDelaySeconds: 30
   failureThreshold: 3
-  initialDelaySeconds: 10
   periodSeconds: 10
   timeoutSeconds: 5
 
@@ -95,7 +107,6 @@ readinessProbe:
     port: http
   initialDelaySeconds: 30
   failureThreshold: 3
-  initialDelaySeconds: 10
   periodSeconds: 10
   timeoutSeconds: 5
 
@@ -142,6 +153,9 @@ fullnameOverride: ""
 
 autoscaling:
   enabled: false
+  minReplicas: 1
+  maxReplicas: 3
+  averageCpuUtilization: 50
 
 ingress:
   enabled: false
@@ -158,6 +172,7 @@ ingress:
 
 service:
   type: ClusterIP
+  loadBalancerIP: ""
   port: 8200
   nodePort: ""
   annotations: {}
```

## 8.0.0-SNAPSHOT 

**Release date:** 2020-06-29

![AppVersion: 8.0.0-SNAPSHOT](https://img.shields.io/static/v1?label=AppVersion&message=8.0.0-SNAPSHOT&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Bump master branch to 8.0.0-SNAPSHOT (#682) 
* Add ServiceAccount annotations (#686) 
* [apm-server] allow customizing probes (#671) 
* [apm-server] increase memory limit (#664) 

### Default value changes

```diff
diff --git a/apm-server/values.yaml b/apm-server/values.yaml
index 2d2b9314..55b3d06a 100755
--- a/apm-server/values.yaml
+++ b/apm-server/values.yaml
@@ -62,7 +62,7 @@ extraVolumes: []
   #   emptyDir: {}
 
 image: "docker.elastic.co/apm/apm-server"
-imageTag: "7.7.1"
+imageTag: "8.0.0-SNAPSHOT"
 imagePullPolicy: "IfNotPresent"
 imagePullSecrets: []
 
@@ -80,12 +80,20 @@ podSecurityContext:
   privileged: false
 
 livenessProbe:
+  httpGet:
+    path: /
+    port: http
+  initialDelaySeconds: 30
   failureThreshold: 3
   initialDelaySeconds: 10
   periodSeconds: 10
   timeoutSeconds: 5
 
 readinessProbe:
+  httpGet:
+    path: /
+    port: http
+  initialDelaySeconds: 30
   failureThreshold: 3
   initialDelaySeconds: 10
   periodSeconds: 10
@@ -97,11 +105,15 @@ resources:
       memory: "100Mi"
     limits:
       cpu: "1000m"
-      memory: "200Mi"
+      memory: "512Mi"
 
 # Custom service account override that the pod will use
 serviceAccount: ""
 
+# Annotations to add to the ServiceAccount that is created if the serviceAccount value isn't set.
+serviceAccountAnnotations: {}
+  # eks.amazonaws.com/role-arn: arn:aws:iam::111111111111:role/k8s.clustername.namespace.serviceaccount
+
 # A list of secrets and their paths to mount inside the pod
 secretMounts: []
 #  - name: elastic-certificate-pem
```

## 7.7.1 

**Release date:** 2020-06-03

![AppVersion: 7.7.1](https://img.shields.io/static/v1?label=AppVersion&message=7.7.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Update changelog for 7.7.1 and 6.8.10 releases (#652) 

### Default value changes

```diff
diff --git a/apm-server/values.yaml b/apm-server/values.yaml
index 22bdf745..2d2b9314 100755
--- a/apm-server/values.yaml
+++ b/apm-server/values.yaml
@@ -62,7 +62,7 @@ extraVolumes: []
   #   emptyDir: {}
 
 image: "docker.elastic.co/apm/apm-server"
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
* [meta] move breaking changes into dedicated file and update readme 
* [meta] centralize development and testing doc in CONTRIBUTING.md 
* [apm-server] update default values 
* [apm-server] format README 
* [apm-server] Add envFrom parameter 

### Default value changes

```diff
diff --git a/apm-server/values.yaml b/apm-server/values.yaml
index f1af513e..22bdf745 100755
--- a/apm-server/values.yaml
+++ b/apm-server/values.yaml
@@ -45,6 +45,13 @@ extraEnvs: []
   #        name: elastic-credentials
   #        key: password
 
+# Allows you to load environment variables from kubernetes secret or config map
+envFrom: []
+# - secretRef:
+#     name: env-secret
+# - configMapRef:
+#     name: config-map
+
 extraVolumeMounts: []
   # - name: extras
   #   mountPath: /usr/share/extras
@@ -55,7 +62,7 @@ extraVolumes: []
   #   emptyDir: {}
 
 image: "docker.elastic.co/apm/apm-server"
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
* fix apiVersion of HPA 
* [apm-server]Fix fullnameOverride setting (#508) 

### Default value changes

```diff
diff --git a/apm-server/values.yaml b/apm-server/values.yaml
index 342cfee7..f1af513e 100755
--- a/apm-server/values.yaml
+++ b/apm-server/values.yaml
@@ -55,7 +55,7 @@ extraVolumes: []
   #   emptyDir: {}
 
 image: "docker.elastic.co/apm/apm-server"
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
* Fix template name 

### Default value changes

```diff
diff --git a/apm-server/values.yaml b/apm-server/values.yaml
index b6c30c29..342cfee7 100755
--- a/apm-server/values.yaml
+++ b/apm-server/values.yaml
@@ -55,7 +55,7 @@ extraVolumes: []
   #   emptyDir: {}
 
 image: "docker.elastic.co/apm/apm-server"
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
* [apm-server] add extra init container 
* [apm-server] add empty value to node port 
* [apm-server]  add extracontainers 
* [apm-server] fix oss image name 
* [apm-server] fix goss http test 
* [apm-server] update doc 

### Default value changes

```diff
diff --git a/apm-server/values.yaml b/apm-server/values.yaml
index 187454bf..b6c30c29 100755
--- a/apm-server/values.yaml
+++ b/apm-server/values.yaml
@@ -20,6 +20,16 @@ apmConfig:
 
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
@@ -45,7 +55,7 @@ extraVolumes: []
   #   emptyDir: {}
 
 image: "docker.elastic.co/apm/apm-server"
-imageTag: "7.5.2"
+imageTag: "7.6.0"
 imagePullPolicy: "IfNotPresent"
 imagePullSecrets: []
 
@@ -130,7 +140,7 @@ ingress:
 service:
   type: ClusterIP
   port: 8200
-  nodePort:
+  nodePort: ""
   annotations: {}
     # cloud.google.com/load-balancer-type: "Internal"
     # service.beta.kubernetes.io/aws-load-balancer-internal: 0.0.0.0/0
```

## 7.5.2 

**Release date:** 2020-01-24

![AppVersion: 7.5.2](https://img.shields.io/static/v1?label=AppVersion&message=7.5.2&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [apm-server] bump to 7.5.2 
* address the last few (hopefully) comments 

### Default value changes

```diff
diff --git a/apm-server/values.yaml b/apm-server/values.yaml
index 18b41d2d..187454bf 100755
--- a/apm-server/values.yaml
+++ b/apm-server/values.yaml
@@ -45,7 +45,7 @@ extraVolumes: []
   #   emptyDir: {}
 
 image: "docker.elastic.co/apm/apm-server"
-imageTag: "7.5.1"
+imageTag: "7.5.2"
 imagePullPolicy: "IfNotPresent"
 imagePullSecrets: []
 
@@ -74,13 +74,13 @@ readinessProbe:
   periodSeconds: 10
   timeoutSeconds: 5
 
-resources: {}
-  #  requests:
-  #    cpu: "100m"
-  #    memory: "100Mi"
-  #  limits:
-  #    cpu: "1000m"
-  #    memory: "200Mi"
+resources:
+    requests:
+      cpu: "100m"
+      memory: "100Mi"
+    limits:
+      cpu: "1000m"
+      memory: "200Mi"
 
 # Custom service account override that the pod will use
 serviceAccount: ""
```

## 7.5.1 

**Release date:** 2020-01-17

![AppVersion: 7.5.1](https://img.shields.io/static/v1?label=AppVersion&message=7.5.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* add apm-server helm chart 

### Default value changes

```diff
---
# Allows you to add config files
apmConfig:
  apm-server.yml: |
    apm-server:
      host: "0.0.0.0:8200"

    queue: {}

    output.elasticsearch:
      hosts: ["http://elasticsearch-master:9200"]
      ## If you have security enabled- you'll need to add the credentials
      ## as environment variables
      # username: "${ELASTICSEARCH_USERNAME}"
      # password: "${ELASTICSEARCH_PASSWORD}"
      ## If SSL is enabled
      # protocol: https
      # ssl.certificate_authorities:
      #  - /usr/share/apm-server/config/certs/elastic-ca.pem

replicas: 1

# Extra environment variables to append to the DaemonSet pod spec.
# This will be appended to the current 'env:' key. You can use any of the kubernetes env
# syntax here
extraEnvs: []
  #  - name: 'ELASTICSEARCH_USERNAME'
  #    valueFrom:
  #      secretKeyRef:
  #        name: elastic-credentials
  #        key: username
  #  - name: 'ELASTICSEARCH_PASSWORD'
  #    valueFrom:
  #      secretKeyRef:
  #        name: elastic-credentials
  #        key: password

extraVolumeMounts: []
  # - name: extras
  #   mountPath: /usr/share/extras
  #   readOnly: true

extraVolumes: []
  # - name: extras
  #   emptyDir: {}

image: "docker.elastic.co/apm/apm-server"
imageTag: "7.5.1"
imagePullPolicy: "IfNotPresent"
imagePullSecrets: []

# Whether this chart should self-manage its service account, role, and associated role binding.
managedServiceAccount: true

podAnnotations: {}
  # iam.amazonaws.com/role: es-cluster

# additionals labels
labels: {}

podSecurityContext:
  runAsUser: 0
  privileged: false

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

resources: {}
  #  requests:
  #    cpu: "100m"
  #    memory: "100Mi"
  #  limits:
  #    cpu: "1000m"
  #    memory: "200Mi"

# Custom service account override that the pod will use
serviceAccount: ""

# A list of secrets and their paths to mount inside the pod
secretMounts: []
#  - name: elastic-certificate-pem
#    secretName: elastic-certificates
#    path: /usr/share/apm-server/config/certs

terminationGracePeriod: 30

tolerations: []

nodeSelector: {}

affinity: {}

# This is the PriorityClass settings as defined in
# https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption/#priorityclass
priorityClassName: ""

updateStrategy:
  type: "RollingUpdate"

# Override various naming aspects of this chart
# Only edit these if you know what you're doing
nameOverride: ""
fullnameOverride: ""

autoscaling:
  enabled: false

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

service:
  type: ClusterIP
  port: 8200
  nodePort:
  annotations: {}
    # cloud.google.com/load-balancer-type: "Internal"
    # service.beta.kubernetes.io/aws-load-balancer-internal: 0.0.0.0/0
    # service.beta.kubernetes.io/azure-load-balancer-internal: "true"
    # service.beta.kubernetes.io/openstack-internal-load-balancer: "true"
    # service.beta.kubernetes.io/cce-load-balancer-internal-vpc: "true"

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
