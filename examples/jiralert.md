# Change Log

## 1.4.1

**Release date:** 2023-06-06

![AppVersion: v1.3.0](https://img.shields.io/static/v1?label=AppVersion&message=v1.3.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [jiralert] fix invalid PodDisruptionBudget (#3458)

### Default value changes

```diff
# No changes in this release
```

## 1.4.0

**Release date:** 2023-06-02

![AppVersion: v1.3.0](https://img.shields.io/static/v1?label=AppVersion&message=v1.3.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [jiralert] add PodDisruptionBudget and topologySpreadConstraints (#3452)

### Default value changes

```diff
diff --git a/charts/jiralert/values.yaml b/charts/jiralert/values.yaml
index 7096adda..541a485e 100755
--- a/charts/jiralert/values.yaml
+++ b/charts/jiralert/values.yaml
@@ -61,6 +61,21 @@ podSecurityContext:
   seccompProfile:
     type: RuntimeDefault
 
+# Ref: https://kubernetes.io/docs/tasks/run-application/configure-pdb/
+podDisruptionBudget: {}
+  # maxUnavailable: 1
+  # minAvailable: 1
+
+## Topology spread constraints rely on node labels to identify the topology domain(s) that each Node is in.
+## Ref: https://kubernetes.io/docs/concepts/workloads/pods/pod-topology-spread-constraints/
+topologySpreadConstraints: []
+  # - maxSkew: 1
+  #   topologyKey: failure-domain.beta.kubernetes.io/zone
+  #   whenUnsatisfiable: DoNotSchedule
+  #   labelSelector:
+  #     matchLabels:
+  #       app.kubernetes.io/instance: jiralert
+
 # Container security context
 securityContext:
   runAsUser: 1001
@@ -88,6 +103,7 @@ serviceAccount:
   annotations: {}
   # -- Labels for service account. Evaluated as a template.
   labels: {}
+  automountServiceAccountToken: false
 
 existingConfigSecret: ~
 

```

## 1.3.1

**Release date:** 2023-05-29

![AppVersion: v1.3.0](https://img.shields.io/static/v1?label=AppVersion&message=v1.3.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [jiralert] fix podSecurityContext and define defaults (#3438)

### Default value changes

```diff
diff --git a/charts/jiralert/values.yaml b/charts/jiralert/values.yaml
index 809dda6c..7096adda 100755
--- a/charts/jiralert/values.yaml
+++ b/charts/jiralert/values.yaml
@@ -57,14 +57,16 @@ annotations: {}
 podAnnotations: {}
 
 # -- jiralert pods' Security Context.
-podSecurityContext: {}
-# fsGroup: 2000
+podSecurityContext:
+  seccompProfile:
+    type: RuntimeDefault
 
 # Container security context
 securityContext:
   runAsUser: 1001
   runAsGroup: 1001
   runAsNonRoot: true
+  allowPrivilegeEscalation: false
   readOnlyRootFilesystem: true
 
 livenessProbe:

```

## 1.3.0

**Release date:** 2023-05-25

![AppVersion: v1.3.0](https://img.shields.io/static/v1?label=AppVersion&message=v1.3.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [jiralert] add extra env and volumes for use with csi-secrets (#3356)

### Default value changes

```diff
diff --git a/charts/jiralert/values.yaml b/charts/jiralert/values.yaml
index 4c36bfbe..809dda6c 100755
--- a/charts/jiralert/values.yaml
+++ b/charts/jiralert/values.yaml
@@ -21,6 +21,13 @@ image:
 extraArgs:
   - -log.level=debug
 
+
+# -- Additional Volume mounts
+extraVolumeMounts: []
+
+# -- Additional Volumes
+extraVolumes: []
+
 # Number of pod replicas
 replicaCount: 1
 

```

## 1.2.1

**Release date:** 2023-05-12

![AppVersion: v1.3.0](https://img.shields.io/static/v1?label=AppVersion&message=v1.3.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* jiralert: values: fixed config: there is no need to do escaping (#3372)

### Default value changes

```diff
diff --git a/charts/jiralert/values.yaml b/charts/jiralert/values.yaml
index 9b3e78a4..4c36bfbe 100755
--- a/charts/jiralert/values.yaml
+++ b/charts/jiralert/values.yaml
@@ -99,8 +99,8 @@ config:
     api_url: "https://example.atlassian.net"
     # user: {{ .config.jiraUser }}
     # password: '{{ .config.jiraToken }}'
-    summary: '{{`{{ template "jira.summary" . }}`}}'
-    description: '{{`{{ template "jira.description" . }}`}}'
+    summary: '{{ template "jira.summary" . }}'
+    description: '{{ template "jira.description" . }}'
     issue_type: Bug
     reopen_state: "To Do"
     reopen_duration: 0h

```

## 1.2.0

**Release date:** 2023-02-23

![AppVersion: v1.3.0](https://img.shields.io/static/v1?label=AppVersion&message=v1.3.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [jiralert] Bump version to v1.3.0 (#3065)

### Default value changes

```diff
# No changes in this release
```

## 1.1.0

**Release date:** 2023-02-10

![AppVersion: 1.2](https://img.shields.io/static/v1?label=AppVersion&message=1.2&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [jiralert] Make tplConfig toggable and introduce a value for complex configurations (#3004)

### Default value changes

```diff
diff --git a/charts/jiralert/values.yaml b/charts/jiralert/values.yaml
index f8d4e4e7..9b3e78a4 100755
--- a/charts/jiralert/values.yaml
+++ b/charts/jiralert/values.yaml
@@ -82,6 +82,13 @@ serviceAccount:
 
 existingConfigSecret: ~
 
+## Pass the jiralert configuration directives through Helm's templating
+## engine. If the jiralert configuration contains Alertmanager templates,
+## they'll need to be properly escaped so that they are not interpreted by
+## Helm
+## ref: https://helm.sh/docs/developing_charts/#using-the-tpl-function
+tplConfig: false
+
 config:
   # File containing template definitions. Required.
   template: jiralert.tmpl
@@ -111,6 +118,9 @@ config:
   #      # Set the Organization field
   #      "customfield_10002": [ { { .config.receiver.organizationID } } ]
 
+# Alternative to config. Must be used with tplConfig=true. Could be used for complex configurations.
+configString: ""
+
 # -- Affinity for pod assignment
 affinity: {}
 

```

## 1.0.1

**Release date:** 2022-12-03

![AppVersion: 1.2](https://img.shields.io/static/v1?label=AppVersion&message=1.2&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [jiralert] Add annotations for deployment and remove annotation config hash if existingSecret (#2776)

### Default value changes

```diff
diff --git a/charts/jiralert/values.yaml b/charts/jiralert/values.yaml
index de1a5859..f8d4e4e7 100755
--- a/charts/jiralert/values.yaml
+++ b/charts/jiralert/values.yaml
@@ -43,6 +43,9 @@ resources:
     cpu: 100m
     memory: 64Mi
 
+# -- Annotations for jiralert deployment
+annotations: {}
+
 # -- Annotations for jiralert pods
 podAnnotations: {}
 

```

## 1.0.0

**Release date:** 2022-11-25

![AppVersion: 1.2](https://img.shields.io/static/v1?label=AppVersion&message=1.2&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [jiralert] Standardized labels (#2735)

### Default value changes

```diff
# No changes in this release
```

## 0.2.0

**Release date:** 2022-11-25

![AppVersion: 1.2](https://img.shields.io/static/v1?label=AppVersion&message=1.2&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [jiralert] Add ServiceMonitor (#2734)

### Default value changes

```diff
diff --git a/charts/jiralert/values.yaml b/charts/jiralert/values.yaml
index f49cd51e..de1a5859 100755
--- a/charts/jiralert/values.yaml
+++ b/charts/jiralert/values.yaml
@@ -145,3 +145,47 @@ issueTemplate: |
   {{ end }}
   Source: {{ .GeneratorURL }}
   {{ end }}{{ end }}`}}
+
+# Enable this if you're using https://github.com/coreos/prometheus-operator
+serviceMonitor:
+  enabled: false
+  # namespace: monitoring
+
+  # Fallback to the prometheus default unless specified
+  # interval: 10s
+
+  ## scheme: HTTP scheme to use for scraping. Can be used with `tlsConfig` for example if using istio mTLS.
+  # scheme: ""
+
+  ## tlsConfig: TLS configuration to use when scraping the endpoint. For example if using istio mTLS.
+  ## Of type: https://github.com/coreos/prometheus-operator/blob/master/Documentation/api.md#tlsconfig
+  # tlsConfig: {}
+
+  # bearerTokenFile:
+  # Fallback to the prometheus default unless specified
+  # scrapeTimeout: 30s
+
+  ## Used to pass Labels that are used by the Prometheus installed in your cluster to select Service Monitors to work with
+  ## ref: https://github.com/coreos/prometheus-operator/blob/master/Documentation/api.md#prometheusspec
+  additionalLabels: {}
+
+  # Retain the job and instance labels of the metrics pushed to the Pushgateway
+  # [Scraping Pushgateway](https://github.com/prometheus/pushgateway#configure-the-pushgateway-as-a-target-to-scrape)
+  honorLabels: true
+
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

```

## 0.1.0

**Release date:** 2022-10-30

![AppVersion: 1.2](https://img.shields.io/static/v1?label=AppVersion&message=1.2&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [jiralert] Setup chart (#2591)

### Default value changes

```diff
global:
  imagePullSecrets: []

# -- String to partially override jiralert.fullname template (will maintain the release name)
nameOverride: ""

# -- Override the namespace
namespaceOverride: ""

# -- String to fully override amazon-eks-pod-identity.fullname template
fullnameOverride: ""

image:
  pullPolicy: IfNotPresent
  # -- jiralert image registry
  repository: quay.io/jiralert/jiralert-linux-amd64
  # -- jiralert image tag (immutable tags are recommended).
  # @default -- `.Chart.AppVersion`
  tag: ""

extraArgs:
  - -log.level=debug

# Number of pod replicas
replicaCount: 1

# Container resources
resources:
  # -- The resources limits for the jiralert container
  ## Example:
  ## limits:
  ##    cpu: 100m
  ##    memory: 128Mi
  limits:
    cpu: 200m
    memory: 128Mi
  # -- The requested resources for the jiralert container
  ## Examples:
  ## requests:
  ##    cpu: 100m
  ##    memory: 128Mi
  requests:
    cpu: 100m
    memory: 64Mi

# -- Annotations for jiralert pods
podAnnotations: {}

# -- jiralert pods' Security Context.
podSecurityContext: {}
# fsGroup: 2000

# Container security context
securityContext:
  runAsUser: 1001
  runAsGroup: 1001
  runAsNonRoot: true
  readOnlyRootFilesystem: true

livenessProbe:
  httpGet:
    path: /healthz
    port: http
readinessProbe:
  httpGet:
    path: /healthz
    port: http

serviceAccount:
  # -- Enable creation of ServiceAccount for nginx pod
  create: true
  # -- The name of the ServiceAccount to use.
  # @default -- A name is generated using the `jiralert.fullname` template
  name: ''
  # -- Annotations for service account. Evaluated as a template.
  annotations: {}
  # -- Labels for service account. Evaluated as a template.
  labels: {}

existingConfigSecret: ~

config:
  # File containing template definitions. Required.
  template: jiralert.tmpl

  # Example: https://github.com/prometheus-community/jiralert/blob/master/examples/jiralert.yml
  defaults:
    # API access fields.
    api_url: "https://example.atlassian.net"
    # user: {{ .config.jiraUser }}
    # password: '{{ .config.jiraToken }}'
    summary: '{{`{{ template "jira.summary" . }}`}}'
    description: '{{`{{ template "jira.description" . }}`}}'
    issue_type: Bug
    reopen_state: "To Do"
    reopen_duration: 0h

  # Receiver definitions. At least one must be defined.
  receivers: []
  #  - name: 'default'
  #    # JIRA project to create the issue in. Required.
  #    project:
  #    # Copy all Prometheus labels into separate JIRA labels. Optional (default: false).
  #    add_group_labels: false
  #    fields:
  #      # Set the Environment field
  #      "customfield_10078": { "value": { { .config.receiver.environment | quote } } }
  #      # Set the Organization field
  #      "customfield_10002": [ { { .config.receiver.organizationID } } ]

# -- Affinity for pod assignment
affinity: {}

# -- Node labels for pod assignment. Evaluated as a template.
nodeSelector: {}

# -- Tolerations for pod assignment. Evaluated as a template.
tolerations: []

ingress:
  enabled: false
  className: ""
  labels: {}
  annotations: {}
    # kubernetes.io/ingress.class: nginx
  # kubernetes.io/tls-acme: "true"
  hosts:
    - host: chart-example.local
      paths:
        - path: /
          pathType: ImplementationSpecific
  tls: []
  #  - secretName: chart-example-tls
  #    hosts:
  #      - chart-example.local

issueTemplate: |
  {{`{{ define "jira.summary" }}[{{ .Status | toUpper }}{{ if eq .Status "firing" }}:{{ .Alerts.Firing | len }}{{ end }}] {{ .GroupLabels.SortedPairs.Values | join " " }} {{ if gt (len .CommonLabels) (len .GroupLabels) }}({{ with .CommonLabels.Remove .GroupLabels.Names }}{{ .Values | join " " }}{{ end }}){{ end }}{{ end }}

  {{ define "jira.description" }}{{ range .Alerts.Firing }}Labels:
  {{ range .Labels.SortedPairs }} - {{ .Name }} = {{ .Value }}
  {{ end }}
  Annotations:
  {{ range .Annotations.SortedPairs }} - {{ .Name }} = {{ .Value }}
  {{ end }}
  Source: {{ .GeneratorURL }}
  {{ end }}{{ end }}`}}

```

---
Autogenerated from Helm Chart and git history using [helm-changelog](https://github.com/mogensen/helm-changelog)
