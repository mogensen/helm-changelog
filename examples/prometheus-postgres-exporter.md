# Change Log

## 4.5.0

**Release date:** 2023-06-15

![AppVersion: 0.11.1](https://img.shields.io/static/v1?label=AppVersion&message=0.11.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-postgres-exporter] user from secret (#3492)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index 9177e9f4..af33ea33 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -133,6 +133,11 @@ config:
     # Specify one of both datasource or datasourceSecret
     host:
     user: postgres
+    userSecret: {}
+    # Secret name
+    #  name:
+    # User key inside secret
+    #  key:
     # Only one of password, passwordFile, passwordSecret and pgpassfile can be specified
     password:
     # Specify passwordFile if DB password is stored in a file.

```

## 4.4.4

**Release date:** 2023-05-20

![AppVersion: 0.11.1](https://img.shields.io/static/v1?label=AppVersion&message=0.11.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-postgres-exporter] Correct role (#3395)

### Default value changes

```diff
# No changes in this release
```

## 4.4.3

**Release date:** 2023-05-04

![AppVersion: 0.11.1](https://img.shields.io/static/v1?label=AppVersion&message=0.11.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-postgres-exporter] template apiVersion based on capabilities for the pod disruption budget (#3322)

### Default value changes

```diff
# No changes in this release
```

## 4.4.2

**Release date:** 2023-05-04

![AppVersion: 0.11.1](https://img.shields.io/static/v1?label=AppVersion&message=0.11.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-postgres-exporter] Add a maintainer (#3312)

### Default value changes

```diff
# No changes in this release
```

## 4.4.1

**Release date:** 2023-04-06

![AppVersion: 0.11.1](https://img.shields.io/static/v1?label=AppVersion&message=0.11.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-postgres-exporter] Add comments regarding excludeDatabases and includeDatabases (#3173)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index 47d2f6ac..9177e9f4 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -164,7 +164,9 @@ config:
   disableSettingsMetrics: false
   autoDiscoverDatabases: false
   excludeDatabases: []
+  # autoDiscoverDatabases must be true for excludeDatabases to be considered
   includeDatabases: []
+  # autoDiscoverDatabases must be true for includeDatabases to be considered
   constantLabels: {}
   # possible values debug, info, warn, error, fatal
   logLevel: ""

```

## 4.4.0

**Release date:** 2023-03-09

![AppVersion: 0.11.1](https://img.shields.io/static/v1?label=AppVersion&message=0.11.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-postgres-exporter] Add automountServiceAccountToken (#3104)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index 15fb5286..47d2f6ac 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -26,6 +26,8 @@ service:
   labels: {}
   annotations: {}
 
+automountServiceAccountToken: false
+
 serviceMonitor:
   # When set true then use a ServiceMonitor to configure scraping
   enabled: false

```

## 4.3.0

**Release date:** 2023-03-01

![AppVersion: 0.11.1](https://img.shields.io/static/v1?label=AppVersion&message=0.11.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-postgres-exporter] Add option to change deployment command (#3070)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index d847bc53..15fb5286 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -16,6 +16,8 @@ image:
   # pullSecrets:
   #   - myRegistrKeySecretName
 
+command: []
+
 service:
   type: ClusterIP
   port: 80

```

## 4.2.1

**Release date:** 2023-01-26

![AppVersion: 0.11.1](https://img.shields.io/static/v1?label=AppVersion&message=0.11.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* fix(postgres-exporter): fix flag format (#2952)

### Default value changes

```diff
# No changes in this release
```

## 4.2.0

**Release date:** 2023-01-24

![AppVersion: 0.11.1](https://img.shields.io/static/v1?label=AppVersion&message=0.11.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-postgres-exporter] add parameters to helm (#2943)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index b0f31b5b..d847bc53 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -154,6 +154,8 @@ config:
     #  name:
     # Connection string key inside secret
     #  key:
+  disableCollectorDatabase: false
+  disableCollectorBgwriter: false
   disableDefaultMetrics: false
   disableSettingsMetrics: false
   autoDiscoverDatabases: false
@@ -164,6 +166,7 @@ config:
   logLevel: ""
   # possible values logfmt, json
   logFormat: ""
+  extraArgs: []
   # Enable queries from an external configmap, enable it will disable inline queries below
   externalQueries:
     enabled: false

```

## 4.1.4

**Release date:** 2023-01-24

![AppVersion: 0.11.1](https://img.shields.io/static/v1?label=AppVersion&message=0.11.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-postgres-exporter] Fix bug in targetPort (#2942)

### Default value changes

```diff
# No changes in this release
```

## 4.1.3

**Release date:** 2022-12-24

![AppVersion: 0.11.1](https://img.shields.io/static/v1?label=AppVersion&message=0.11.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [postgres-exporter] Fix placement of securityContext (#2853)

### Default value changes

```diff
# No changes in this release
```

## 4.1.2

**Release date:** 2022-12-23

![AppVersion: 0.11.1](https://img.shields.io/static/v1?label=AppVersion&message=0.11.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [postgres-exporter] Fix indentation at securityContext (#2840) (#2850)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index bf91f88f..b0f31b5b 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -24,13 +24,6 @@ service:
   labels: {}
   annotations: {}
 
-podSecurityContext: {}
-  # runAsUser: 1001
-  # runAsGroup: 1001
-  # runAsNonRoot: true
-  # seccompProfile:
-    # type: RuntimeDefault
-
 serviceMonitor:
   # When set true then use a ServiceMonitor to configure scraping
   enabled: false
@@ -84,6 +77,7 @@ rbac:
   create: true
   # Specifies whether a PodSecurityPolicy should be created
   pspEnabled: true
+
 serviceAccount:
   # Specifies whether a ServiceAccount should be created
   create: true
@@ -101,18 +95,27 @@ networkPolicy:
   # Set labels for the NetworkPolicy
   labels: {}
 
+# The securityContext of the pod.
+# See https://kubernetes.io/docs/concepts/policy/security-context/ for more.
+podSecurityContext: {}
+  # runAsUser: 1001
+  # runAsGroup: 1001
+  # runAsNonRoot: true
+  # seccompProfile:
+  #   type: RuntimeDefault
+
+# The securityContext of the container.
+# See https://kubernetes.io/docs/concepts/policy/security-context/ for more.
 securityContext: {}
-  # The securityContext this Pod should use. See https://kubernetes.io/docs/concepts/policy/security-context/ for more.
-  # runAsUser: 65534
   # runAsUser: 1001
   # runAsGroup: 1001
   # readOnlyRootFilesystem: true
   # runAsNonRoot: true
   # allowPrivilegeEscalation: false
   # capabilities:
-   # drop: ["ALL"]
+  #   drop: ["ALL"]
   # seccompProfile:
-    # type: RuntimeDefault
+  #   type: RuntimeDefault
 
 hostAliases: []
   # Set Host Aliases as per https://kubernetes.io/docs/tasks/network/customize-hosts-file-for-pods/
@@ -475,7 +478,6 @@ extraContainers: []
 
 # Additional volumes, e. g. for secrets used in an extraContainer
 extraVolumes: []
-
 # Uncomment for mounting custom ca-certificates
 #  - name: ssl-certs
 #    secret:
@@ -487,7 +489,6 @@ extraVolumes: []
 
 # Additional volume mounts
 extraVolumeMounts: []
-
 # Uncomment for mounting custom ca-certificates file into container
 #  - name: ssl-certs
 #    mountPath: /etc/ssl/certs/ca-certificates.crt

```

## 4.1.1

**Release date:** 2022-12-23

![AppVersion: 0.11.1](https://img.shields.io/static/v1?label=AppVersion&message=0.11.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [postgres-exporter] feat: propagate pg_stats_statement WARNING (#2838)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index 97b54f86..bf91f88f 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -315,6 +315,7 @@ config:
             usage: "COUNTER"
             description: "Number of buffer hits in this table's TOAST table indexes (if any)"
 
+    # WARNING: This set of metrics can be very expensive on a busy server as every unique query executed will create an additional time series
     pg_stat_statements:
       query: "SELECT t2.rolname, t3.datname, queryid, calls, ( total_plan_time + total_exec_time ) / 1000 as total_time_seconds, ( min_plan_time + min_exec_time ) / 1000 as min_time_seconds, ( max_plan_time + max_exec_time ) / 1000 as max_time_seconds, ( mean_plan_time + mean_exec_time ) / 1000 as mean_time_seconds, ( stddev_plan_time + stddev_exec_time )  / 1000 as stddev_time_seconds, rows, shared_blks_hit, shared_blks_read, shared_blks_dirtied, shared_blks_written, local_blks_hit, local_blks_read, local_blks_dirtied, local_blks_written, temp_blks_read, temp_blks_written, blk_read_time / 1000 as blk_read_time_seconds, blk_write_time / 1000 as blk_write_time_seconds FROM pg_stat_statements t1 JOIN pg_roles t2 ON (t1.userid=t2.oid) JOIN pg_database t3 ON (t1.dbid=t3.oid) WHERE t2.rolname != 'rdsadmin' AND queryid IS NOT NULL"
       master: true

```

## 4.1.0

**Release date:** 2022-12-22

![AppVersion: 0.11.1](https://img.shields.io/static/v1?label=AppVersion&message=0.11.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Refactor extra envs (#2844)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index e571d53d..97b54f86 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -453,8 +453,13 @@ readinessProbe:
   timeoutSeconds: 1
 
 # ExtraEnvs
-extraEnvs: {}
-  # EXTRA_ENV: value
+extraEnvs: []
+  # - name: EXTRA_ENV
+  #   value: value
+  # - name: POD_NAMESPACE
+  #   valueFrom:
+  #     fieldRef:
+  #       fieldPath: metadata.namespace
 
 # Init containers, e. g. for secrets creation before the exporter
 initContainers: []

```

## 4.0.0

**Release date:** 2022-12-13

![AppVersion: 0.11.1](https://img.shields.io/static/v1?label=AppVersion&message=0.11.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-postgres-exporter] Remove pg_database query (#2752)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index 7e119454..e571d53d 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -315,18 +315,6 @@ config:
             usage: "COUNTER"
             description: "Number of buffer hits in this table's TOAST table indexes (if any)"
 
-    pg_database:
-      query: "SELECT pg_database.datname, pg_database_size(pg_database.datname) as size_bytes FROM pg_database"
-      master: true
-      cache_seconds: 30
-      metrics:
-        - datname:
-            usage: "LABEL"
-            description: "Name of the database"
-        - size_bytes:
-            usage: "GAUGE"
-            description: "Disk space used by the database"
-
     pg_stat_statements:
       query: "SELECT t2.rolname, t3.datname, queryid, calls, ( total_plan_time + total_exec_time ) / 1000 as total_time_seconds, ( min_plan_time + min_exec_time ) / 1000 as min_time_seconds, ( max_plan_time + max_exec_time ) / 1000 as max_time_seconds, ( mean_plan_time + mean_exec_time ) / 1000 as mean_time_seconds, ( stddev_plan_time + stddev_exec_time )  / 1000 as stddev_time_seconds, rows, shared_blks_hit, shared_blks_read, shared_blks_dirtied, shared_blks_written, local_blks_hit, local_blks_read, local_blks_dirtied, local_blks_written, temp_blks_read, temp_blks_written, blk_read_time / 1000 as blk_read_time_seconds, blk_write_time / 1000 as blk_write_time_seconds FROM pg_stat_statements t1 JOIN pg_roles t2 ON (t1.userid=t2.oid) JOIN pg_database t3 ON (t1.dbid=t3.oid) WHERE t2.rolname != 'rdsadmin' AND queryid IS NOT NULL"
       master: true

```

## 3.3.0

**Release date:** 2022-11-22

![AppVersion: 0.11.1](https://img.shields.io/static/v1?label=AppVersion&message=0.11.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-postgre-exporter] new securityContext for POD (#2698)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index 7da996dd..7e119454 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -24,6 +24,13 @@ service:
   labels: {}
   annotations: {}
 
+podSecurityContext: {}
+  # runAsUser: 1001
+  # runAsGroup: 1001
+  # runAsNonRoot: true
+  # seccompProfile:
+    # type: RuntimeDefault
+
 serviceMonitor:
   # When set true then use a ServiceMonitor to configure scraping
   enabled: false
@@ -97,6 +104,15 @@ networkPolicy:
 securityContext: {}
   # The securityContext this Pod should use. See https://kubernetes.io/docs/concepts/policy/security-context/ for more.
   # runAsUser: 65534
+  # runAsUser: 1001
+  # runAsGroup: 1001
+  # readOnlyRootFilesystem: true
+  # runAsNonRoot: true
+  # allowPrivilegeEscalation: false
+  # capabilities:
+   # drop: ["ALL"]
+  # seccompProfile:
+    # type: RuntimeDefault
 
 hostAliases: []
   # Set Host Aliases as per https://kubernetes.io/docs/tasks/network/customize-hosts-file-for-pods/

```

## 3.2.0

**Release date:** 2022-11-22

![AppVersion: 0.11.1](https://img.shields.io/static/v1?label=AppVersion&message=0.11.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-postgre-exporter] Add extraEnvs to deployment (#2715)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index e0b937e5..7da996dd 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -6,7 +6,7 @@ replicaCount: 1
 
 image:
   repository: quay.io/prometheuscommunity/postgres-exporter
-  tag: v0.10.1
+  tag: v0.11.1
   pullPolicy: IfNotPresent
 
   ## Optionally specify an array of imagePullSecrets.
@@ -448,6 +448,9 @@ readinessProbe:
   initialDelaySeconds: 0
   timeoutSeconds: 1
 
+# ExtraEnvs
+extraEnvs: {}
+  # EXTRA_ENV: value
 
 # Init containers, e. g. for secrets creation before the exporter
 initContainers: []

```

## 3.1.5

**Release date:** 2022-10-21

![AppVersion: 0.10.1](https://img.shields.io/static/v1?label=AppVersion&message=0.10.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Allow user queries to merge with default queries (#2592)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index e06ea955..e0b937e5 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -426,6 +426,9 @@ config:
             usage: "HISTOGRAM"
             description: "Idle time of server processes"
 
+  # These are user-specified queries that are deep merged with the queries above
+  userQueries: ""
+
 nodeSelector: {}
 
 tolerations: []

```

## 3.1.4

**Release date:** 2022-10-16

![AppVersion: 0.10.1](https://img.shields.io/static/v1?label=AppVersion&message=0.10.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-postgres-exporter] Align PSP deprecation with other charts (#2573)

### Default value changes

```diff
# No changes in this release
```

## 3.1.3

**Release date:** 2022-09-09

![AppVersion: 0.10.1](https://img.shields.io/static/v1?label=AppVersion&message=0.10.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-postgres-exporter] DATA_SOURCE_PASS_FILE support (#2437)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index bd8c2f38..e06ea955 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -110,8 +110,11 @@ config:
     # Specify one of both datasource or datasourceSecret
     host:
     user: postgres
-    # Only one of password, passwordSecret and pgpassfile can be specified
+    # Only one of password, passwordFile, passwordSecret and pgpassfile can be specified
     password:
+    # Specify passwordFile if DB password is stored in a file.
+    # For example, to use with vault-injector from Hashicorp
+    passwordFile: ''
     # Specify passwordSecret if DB password is stored in secret.
     passwordSecret: {}
     # Secret name

```

## 3.1.2

**Release date:** 2022-08-23

![AppVersion: 0.10.1](https://img.shields.io/static/v1?label=AppVersion&message=0.10.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-postgres-exporter] Fix indentation (#2388)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index 27e14c02..bd8c2f38 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -166,31 +166,31 @@ config:
 
     pg_stat_user_tables:
       query: |
-       SELECT
-         current_database() datname,
-         schemaname,
-         relname,
-         seq_scan,
-         seq_tup_read,
-         idx_scan,
-         idx_tup_fetch,
-         n_tup_ins,
-         n_tup_upd,
-         n_tup_del,
-         n_tup_hot_upd,
-         n_live_tup,
-         n_dead_tup,
-         n_mod_since_analyze,
-         COALESCE(last_vacuum, '1970-01-01Z') as last_vacuum,
-         COALESCE(last_autovacuum, '1970-01-01Z') as last_autovacuum,
-         COALESCE(last_analyze, '1970-01-01Z') as last_analyze,
-         COALESCE(last_autoanalyze, '1970-01-01Z') as last_autoanalyze,
-         vacuum_count,
-         autovacuum_count,
-         analyze_count,
-         autoanalyze_count
-       FROM
-         pg_stat_user_tables
+        SELECT
+          current_database() datname,
+          schemaname,
+          relname,
+          seq_scan,
+          seq_tup_read,
+          idx_scan,
+          idx_tup_fetch,
+          n_tup_ins,
+          n_tup_upd,
+          n_tup_del,
+          n_tup_hot_upd,
+          n_live_tup,
+          n_dead_tup,
+          n_mod_since_analyze,
+          COALESCE(last_vacuum, '1970-01-01Z') as last_vacuum,
+          COALESCE(last_autovacuum, '1970-01-01Z') as last_autovacuum,
+          COALESCE(last_analyze, '1970-01-01Z') as last_analyze,
+          COALESCE(last_autoanalyze, '1970-01-01Z') as last_autoanalyze,
+          vacuum_count,
+          autovacuum_count,
+          analyze_count,
+          autoanalyze_count
+        FROM
+          pg_stat_user_tables
       metrics:
         - datname:
             usage: "LABEL"

```

## 3.1.1

**Release date:** 2022-08-18

![AppVersion: 0.10.1](https://img.shields.io/static/v1?label=AppVersion&message=0.10.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* fix(prometheus-postgres-exporter): convert port variable to int (#2372)

### Default value changes

```diff
# No changes in this release
```

## 3.1.0

**Release date:** 2022-08-06

![AppVersion: 0.10.1](https://img.shields.io/static/v1?label=AppVersion&message=0.10.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-postgres-exporter] Add possibility to have hostAliases defined in deployment spec (#2351)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index a2ab8ff3..27e14c02 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -98,6 +98,13 @@ securityContext: {}
   # The securityContext this Pod should use. See https://kubernetes.io/docs/concepts/policy/security-context/ for more.
   # runAsUser: 65534
 
+hostAliases: []
+  # Set Host Aliases as per https://kubernetes.io/docs/tasks/network/customize-hosts-file-for-pods/
+  # - ip: "127.0.0.1"
+  #   hostnames:
+  #   - "foo.local"
+  #   - "bar.local"
+
 config:
   datasource:
     # Specify one of both datasource or datasourceSecret

```

## 3.0.3

**Release date:** 2022-06-29

![AppVersion: 0.10.1](https://img.shields.io/static/v1?label=AppVersion&message=0.10.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* uses policy/v1 when available (#2208)

### Default value changes

```diff
# No changes in this release
```

## 3.0.2

**Release date:** 2022-06-24

![AppVersion: 0.10.1](https://img.shields.io/static/v1?label=AppVersion&message=0.10.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-postgres-exporter] Use correct port in port-forward command (#2189)

### Default value changes

```diff
# No changes in this release
```

## 3.0.1

**Release date:** 2022-06-12

![AppVersion: 0.10.1](https://img.shields.io/static/v1?label=AppVersion&message=0.10.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-postgres-exporter] fix cve-2021-25742 (#2149)

### Default value changes

```diff
# No changes in this release
```

## 3.0.0

**Release date:** 2022-05-05

![AppVersion: 0.10.1](https://img.shields.io/static/v1?label=AppVersion&message=0.10.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-postgres-exporter]:  use new time column scheme of PG13 (#2021)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index ea33595c..a2ab8ff3 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -302,7 +302,7 @@ config:
             description: "Disk space used by the database"
 
     pg_stat_statements:
-      query: "SELECT t2.rolname, t3.datname, queryid, calls, total_time / 1000 as total_time_seconds, min_time / 1000 as min_time_seconds, max_time / 1000 as max_time_seconds, mean_time / 1000 as mean_time_seconds, stddev_time / 1000 as stddev_time_seconds, rows, shared_blks_hit, shared_blks_read, shared_blks_dirtied, shared_blks_written, local_blks_hit, local_blks_read, local_blks_dirtied, local_blks_written, temp_blks_read, temp_blks_written, blk_read_time / 1000 as blk_read_time_seconds, blk_write_time / 1000 as blk_write_time_seconds FROM pg_stat_statements t1 JOIN pg_roles t2 ON (t1.userid=t2.oid) JOIN pg_database t3 ON (t1.dbid=t3.oid) WHERE t2.rolname != 'rdsadmin' AND queryid IS NOT NULL"
+      query: "SELECT t2.rolname, t3.datname, queryid, calls, ( total_plan_time + total_exec_time ) / 1000 as total_time_seconds, ( min_plan_time + min_exec_time ) / 1000 as min_time_seconds, ( max_plan_time + max_exec_time ) / 1000 as max_time_seconds, ( mean_plan_time + mean_exec_time ) / 1000 as mean_time_seconds, ( stddev_plan_time + stddev_exec_time )  / 1000 as stddev_time_seconds, rows, shared_blks_hit, shared_blks_read, shared_blks_dirtied, shared_blks_written, local_blks_hit, local_blks_read, local_blks_dirtied, local_blks_written, temp_blks_read, temp_blks_written, blk_read_time / 1000 as blk_read_time_seconds, blk_write_time / 1000 as blk_write_time_seconds FROM pg_stat_statements t1 JOIN pg_roles t2 ON (t1.userid=t2.oid) JOIN pg_database t3 ON (t1.dbid=t3.oid) WHERE t2.rolname != 'rdsadmin' AND queryid IS NOT NULL"
       master: true
       metrics:
         - rolname:

```

## 2.10.1

**Release date:** 2022-05-02

![AppVersion: 0.10.1](https://img.shields.io/static/v1?label=AppVersion&message=0.10.1&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* chore(postgres-exporter): update image to 0.10.1 (#2019)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index 51969102..ea33595c 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -6,7 +6,7 @@ replicaCount: 1
 
 image:
   repository: quay.io/prometheuscommunity/postgres-exporter
-  tag: v0.10.0
+  tag: v0.10.1
   pullPolicy: IfNotPresent
 
   ## Optionally specify an array of imagePullSecrets.

```

## 2.10.0

**Release date:** 2022-04-12

![AppVersion: 0.10.0](https://img.shields.io/static/v1?label=AppVersion&message=0.10.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-postgres-exporter] Add simple Ingress to make service.targetPort available in namespace (#1973)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index 10739694..51969102 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -86,6 +86,14 @@ serviceAccount:
   # Add annotations to the ServiceAccount, useful for EKS IAM Roles for Service Accounts or Google Workload Identity.
   annotations: {}
 
+# Add a default ingress to allow namespace access to service.targetPort
+# Helpful if other NetworkPolicies are configured in the namespace
+networkPolicy:
+  # Specifies whether a NetworkPolicy should be created
+  enabled: false
+  # Set labels for the NetworkPolicy
+  labels: {}
+
 securityContext: {}
   # The securityContext this Pod should use. See https://kubernetes.io/docs/concepts/policy/security-context/ for more.
   # runAsUser: 65534

```

## 2.9.0

**Release date:** 2022-04-08

![AppVersion: 0.10.0](https://img.shields.io/static/v1?label=AppVersion&message=0.10.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Add poddisruptionbudget for prometheus-postgres-exporter chart (#1964)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index 659a1b6b..10739694 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -458,3 +458,7 @@ extraVolumeMounts: []
 #  - name: ssl-certs
 #    mountPath: /etc/ssl/certs/ca-certificates.crt
 #    subPath: ca-certificates.crt
+
+podDisruptionBudget:
+  enabled: false
+  maxUnavailable: 1

```

## 2.8.0

**Release date:** 2022-03-31

![AppVersion: 0.10.0](https://img.shields.io/static/v1?label=AppVersion&message=0.10.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* prometheus-postgres-exporter added pgpassfile configuration value (#1937)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index 1330bc6f..659a1b6b 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -95,7 +95,7 @@ config:
     # Specify one of both datasource or datasourceSecret
     host:
     user: postgres
-    # Only one of password and passwordSecret can be specified
+    # Only one of password, passwordSecret and pgpassfile can be specified
     password:
     # Specify passwordSecret if DB password is stored in secret.
     passwordSecret: {}
@@ -103,6 +103,9 @@ config:
     #  name:
     # Password key inside secret
     #  key:
+    pgpassfile: ''
+    # If pgpassfile is set, it is used to initialize the PGPASSFILE environment variable.
+    # See https://www.postgresql.org/docs/14/libpq-pgpass.html for more info.
     port: "5432"
     database: ''
     sslmode: disable

```

## 2.7.0

**Release date:** 2022-03-22

![AppVersion: 0.10.0](https://img.shields.io/static/v1?label=AppVersion&message=0.10.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-postgres-exporter] Add relabelings to servicemonitor (#1895)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index 433d3540..1330bc6f 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -41,6 +41,8 @@ serviceMonitor:
   # targetLabels: []
   # MetricRelabelConfigs to apply to samples before ingestion
   # metricRelabelings: []
+  # Set relabel_configs as per https://prometheus.io/docs/prometheus/latest/configuration/configuration/#relabel_config
+  # relabelings: []
 
 prometheusRule:
   enabled: false

```

## 2.6.1

**Release date:** 2022-03-11

![AppVersion: 0.10.0](https://img.shields.io/static/v1?label=AppVersion&message=0.10.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* fix(postgres-exporter): trimmed data_source_uri if no extra options are provided (#1865)

### Default value changes

```diff
# No changes in this release
```

## 2.6.0

**Release date:** 2022-03-07

![AppVersion: 0.10.0](https://img.shields.io/static/v1?label=AppVersion&message=0.10.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-postgres-exporter] Support custom DB connection parameters (#1850)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index 55b7829f..433d3540 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -104,6 +104,7 @@ config:
     port: "5432"
     database: ''
     sslmode: disable
+    extraParams: ''
   datasourceSecret: {}
     # Specifies if datasource should be sourced from secret value in format: postgresql://login:password@hostname:port/dbname?sslmode=disable
     # Multiple Postgres databases can be configured by comma separated postgres connection strings

```

## 2.5.0

**Release date:** 2022-01-18

![AppVersion: 0.10.0](https://img.shields.io/static/v1?label=AppVersion&message=0.10.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-postgres-exporter] Add log.format args  (#1719)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index f9307859..55b7829f 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -119,6 +119,8 @@ config:
   constantLabels: {}
   # possible values debug, info, warn, error, fatal
   logLevel: ""
+  # possible values logfmt, json
+  logFormat: ""
   # Enable queries from an external configmap, enable it will disable inline queries below
   externalQueries:
     enabled: false

```

## 2.4.0

**Release date:** 2021-12-02

![AppVersion: 0.10.0](https://img.shields.io/static/v1?label=AppVersion&message=0.10.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-postgres-exporter] Add support for queries in external configmap (#1545)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index ef023792..f9307859 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -119,6 +119,10 @@ config:
   constantLabels: {}
   # possible values debug, info, warn, error, fatal
   logLevel: ""
+  # Enable queries from an external configmap, enable it will disable inline queries below
+  externalQueries:
+    enabled: false
+    configmap: postgresql-common-exporter-queries
   # These are the default queries that the exporter will run, extracted from: https://github.com/prometheus-community/postgres_exporter/blob/master/queries.yaml
   queries: |-
     pg_replication:

```

## 2.3.7

**Release date:** 2021-09-23

![AppVersion: 0.10.0](https://img.shields.io/static/v1?label=AppVersion&message=0.10.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* fix(postgres-exporter): ignore pg_stat_statememts with empty queryId (#1364)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index efe7e135..ef023792 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -282,7 +282,7 @@ config:
             description: "Disk space used by the database"
 
     pg_stat_statements:
-      query: "SELECT t2.rolname, t3.datname, queryid, calls, total_time / 1000 as total_time_seconds, min_time / 1000 as min_time_seconds, max_time / 1000 as max_time_seconds, mean_time / 1000 as mean_time_seconds, stddev_time / 1000 as stddev_time_seconds, rows, shared_blks_hit, shared_blks_read, shared_blks_dirtied, shared_blks_written, local_blks_hit, local_blks_read, local_blks_dirtied, local_blks_written, temp_blks_read, temp_blks_written, blk_read_time / 1000 as blk_read_time_seconds, blk_write_time / 1000 as blk_write_time_seconds FROM pg_stat_statements t1 JOIN pg_roles t2 ON (t1.userid=t2.oid) JOIN pg_database t3 ON (t1.dbid=t3.oid) WHERE t2.rolname != 'rdsadmin'"
+      query: "SELECT t2.rolname, t3.datname, queryid, calls, total_time / 1000 as total_time_seconds, min_time / 1000 as min_time_seconds, max_time / 1000 as max_time_seconds, mean_time / 1000 as mean_time_seconds, stddev_time / 1000 as stddev_time_seconds, rows, shared_blks_hit, shared_blks_read, shared_blks_dirtied, shared_blks_written, local_blks_hit, local_blks_read, local_blks_dirtied, local_blks_written, temp_blks_read, temp_blks_written, blk_read_time / 1000 as blk_read_time_seconds, blk_write_time / 1000 as blk_write_time_seconds FROM pg_stat_statements t1 JOIN pg_roles t2 ON (t1.userid=t2.oid) JOIN pg_database t3 ON (t1.dbid=t3.oid) WHERE t2.rolname != 'rdsadmin' AND queryid IS NOT NULL"
       master: true
       metrics:
         - rolname:

```

## 2.3.6

**Release date:** 2021-08-24

![AppVersion: 0.10.0](https://img.shields.io/static/v1?label=AppVersion&message=0.10.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-postgres-exporter] Update to v0.10.0 (#1281)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index 2baec0a2..efe7e135 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -6,7 +6,7 @@ replicaCount: 1
 
 image:
   repository: quay.io/prometheuscommunity/postgres-exporter
-  tag: v0.9.0
+  tag: v0.10.0
   pullPolicy: IfNotPresent
 
   ## Optionally specify an array of imagePullSecrets.

```

## 2.3.5

**Release date:** 2021-06-23

![AppVersion: 0.9.0](https://img.shields.io/static/v1?label=AppVersion&message=0.9.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Add --include-databases arguments to Postgres exporter values (#1099)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index b6d3435a..2baec0a2 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -115,6 +115,7 @@ config:
   disableSettingsMetrics: false
   autoDiscoverDatabases: false
   excludeDatabases: []
+  includeDatabases: []
   constantLabels: {}
   # possible values debug, info, warn, error, fatal
   logLevel: ""

```

## 2.3.4

**Release date:** 2021-05-30

![AppVersion: 0.9.0](https://img.shields.io/static/v1?label=AppVersion&message=0.9.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Update pg_stat_activity query (#1012)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index 4fd3883e..b6d3435a 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -351,14 +351,14 @@ config:
             usage: "COUNTER"
             description: "Total time the statement spent writing blocks, in milliseconds (if track_io_timing is enabled, otherwise zero)"
 
-    pg_stat_activity:
+    pg_stat_activity_idle:
       query: |
         WITH
           metrics AS (
             SELECT
               application_name,
-              SUM(EXTRACT(EPOCH FROM (CURRENT_TIMESTAMP - state_change))::bigint)::float AS process_idle_seconds_sum,
-              COUNT(*) AS process_idle_seconds_count
+              SUM(EXTRACT(EPOCH FROM (CURRENT_TIMESTAMP - state_change))::bigint)::float AS process_seconds_sum,
+              COUNT(*) AS process_seconds_count
             FROM pg_stat_activity
             WHERE state = 'idle'
             GROUP BY application_name
@@ -381,17 +381,17 @@ config:
           )
         SELECT
           application_name,
-          process_idle_seconds_sum,
-          process_idle_seconds_count,
-          ARRAY_AGG(le) AS process_idle_seconds,
-          ARRAY_AGG(bucket) AS process_idle_seconds_bucket
+          process_seconds_sum,
+          process_seconds_count,
+          ARRAY_AGG(le) AS process_seconds,
+          ARRAY_AGG(bucket) AS process_seconds_bucket
         FROM metrics JOIN buckets USING (application_name)
         GROUP BY 1, 2, 3
       metrics:
         - application_name:
             usage: "LABEL"
             description: "Application Name"
-        - process_idle_seconds:
+        - process_seconds:
             usage: "HISTOGRAM"
             description: "Idle time of server processes"
 

```

## 2.3.3

**Release date:** 2021-05-21

![AppVersion: 0.9.0](https://img.shields.io/static/v1?label=AppVersion&message=0.9.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* don't apply secret when using existing (#992)

### Default value changes

```diff
# No changes in this release
```

## 2.3.2

**Release date:** 2021-05-21

![AppVersion: 0.9.0](https://img.shields.io/static/v1?label=AppVersion&message=0.9.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-postgres-exporter] drop capability on service monitor (#991)

### Default value changes

```diff
# No changes in this release
```

## 2.3.1

**Release date:** 2021-05-11

![AppVersion: 0.9.0](https://img.shields.io/static/v1?label=AppVersion&message=0.9.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-postgres-exporter]Update default password or passwordSecret in values.yaml (#958)
* prometheus-postgres-exporter - Allow readiness timeout to be configurable via config (#815)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index 32a0fdfa..4fd3883e 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -94,7 +94,7 @@ config:
     host:
     user: postgres
     # Only one of password and passwordSecret can be specified
-    password: somepassword
+    password:
     # Specify passwordSecret if DB password is stored in secret.
     passwordSecret: {}
     # Secret name
@@ -408,9 +408,12 @@ podLabels: {}
 # Configurable health checks
 livenessProbe:
   initialDelaySeconds: 0
+  timeoutSeconds: 1
 
 readinessProbe:
   initialDelaySeconds: 0
+  timeoutSeconds: 1
+
 
 # Init containers, e. g. for secrets creation before the exporter
 initContainers: []

```

## 2.3.0

**Release date:** 2021-04-17

![AppVersion: 0.9.0](https://img.shields.io/static/v1?label=AppVersion&message=0.9.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-postgres-exporter] Use arrays instead of strings for extraContainers, extraVolumes and extraVolumeMounts (#856)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index 47cc9561..32a0fdfa 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -421,10 +421,10 @@ initContainers: []
   #       mountPath: /creds
 
 # Additional sidecar containers, e. g. for a database proxy, such as Google's cloudsql-proxy
-extraContainers: |
+extraContainers: []
 
 # Additional volumes, e. g. for secrets used in an extraContainer
-extraVolumes: |
+extraVolumes: []
 
 # Uncomment for mounting custom ca-certificates
 #  - name: ssl-certs
@@ -436,7 +436,7 @@ extraVolumes: |
 #      secretName: ssl-certs
 
 # Additional volume mounts
-extraVolumeMounts: |
+extraVolumeMounts: []
 
 # Uncomment for mounting custom ca-certificates file into container
 #  - name: ssl-certs

```

## 2.2.0

**Release date:** 2021-03-19

![AppVersion: 0.9.0](https://img.shields.io/static/v1?label=AppVersion&message=0.9.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-postgres-exporter] Update to the latest default queries.yaml (#775)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index af280051..47cc9561 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -118,10 +118,10 @@ config:
   constantLabels: {}
   # possible values debug, info, warn, error, fatal
   logLevel: ""
-  # this are the defaults queries that the exporter will run, extracted from: https://github.com/wrouesnel/postgres_exporter/blob/master/queries.yaml
+  # These are the default queries that the exporter will run, extracted from: https://github.com/prometheus-community/postgres_exporter/blob/master/queries.yaml
   queries: |-
     pg_replication:
-      query: "SELECT EXTRACT(EPOCH FROM (now() - pg_last_xact_replay_timestamp())) as lag"
+      query: "SELECT CASE WHEN NOT pg_is_in_recovery() THEN 0 ELSE GREATEST (0, EXTRACT(EPOCH FROM (now() - pg_last_xact_replay_timestamp()))) END AS lag"
       master: true
       metrics:
         - lag:
@@ -137,7 +137,32 @@ config:
             description: "Time at which postmaster started"
 
     pg_stat_user_tables:
-      query: "SELECT current_database() datname, schemaname, relname, seq_scan, seq_tup_read, idx_scan, idx_tup_fetch, n_tup_ins, n_tup_upd, n_tup_del, n_tup_hot_upd, n_live_tup, n_dead_tup, n_mod_since_analyze, COALESCE(last_vacuum, '1970-01-01Z'), COALESCE(last_vacuum, '1970-01-01Z') as last_vacuum, COALESCE(last_autovacuum, '1970-01-01Z') as last_autovacuum, COALESCE(last_analyze, '1970-01-01Z') as last_analyze, COALESCE(last_autoanalyze, '1970-01-01Z') as last_autoanalyze, vacuum_count, autovacuum_count, analyze_count, autoanalyze_count FROM pg_stat_user_tables"
+      query: |
+       SELECT
+         current_database() datname,
+         schemaname,
+         relname,
+         seq_scan,
+         seq_tup_read,
+         idx_scan,
+         idx_tup_fetch,
+         n_tup_ins,
+         n_tup_upd,
+         n_tup_del,
+         n_tup_hot_upd,
+         n_live_tup,
+         n_dead_tup,
+         n_mod_since_analyze,
+         COALESCE(last_vacuum, '1970-01-01Z') as last_vacuum,
+         COALESCE(last_autovacuum, '1970-01-01Z') as last_autovacuum,
+         COALESCE(last_analyze, '1970-01-01Z') as last_analyze,
+         COALESCE(last_autoanalyze, '1970-01-01Z') as last_autoanalyze,
+         vacuum_count,
+         autovacuum_count,
+         analyze_count,
+         autoanalyze_count
+       FROM
+         pg_stat_user_tables
       metrics:
         - datname:
             usage: "LABEL"
@@ -244,7 +269,7 @@ config:
             description: "Number of buffer hits in this table's TOAST table indexes (if any)"
 
     pg_database:
-      query: "SELECT pg_database.datname, pg_database_size(pg_database.datname) as size FROM pg_database"
+      query: "SELECT pg_database.datname, pg_database_size(pg_database.datname) as size_bytes FROM pg_database"
       master: true
       cache_seconds: 30
       metrics:
@@ -256,7 +281,7 @@ config:
             description: "Disk space used by the database"
 
     pg_stat_statements:
-      query: "SELECT t2.rolname, t3.datname, queryid, calls, total_time / 1000 as total_time_seconds, min_time / 1000 as min_time_seconds, max_time / 1000 as max_time_seconds, mean_time / 1000 as mean_time_seconds, stddev_time / 1000 as stddev_time_seconds, rows, shared_blks_hit, shared_blks_read, shared_blks_dirtied, shared_blks_written, local_blks_hit, local_blks_read, local_blks_dirtied, local_blks_written, temp_blks_read, temp_blks_written, blk_read_time / 1000 as blk_read_time_seconds, blk_write_time / 1000 as blk_write_time_seconds FROM pg_stat_statements t1 join pg_roles t2 on (t1.userid=t2.oid) join pg_database t3 on (t1.dbid=t3.oid)"
+      query: "SELECT t2.rolname, t3.datname, queryid, calls, total_time / 1000 as total_time_seconds, min_time / 1000 as min_time_seconds, max_time / 1000 as max_time_seconds, mean_time / 1000 as mean_time_seconds, stddev_time / 1000 as stddev_time_seconds, rows, shared_blks_hit, shared_blks_read, shared_blks_dirtied, shared_blks_written, local_blks_hit, local_blks_read, local_blks_dirtied, local_blks_written, temp_blks_read, temp_blks_written, blk_read_time / 1000 as blk_read_time_seconds, blk_write_time / 1000 as blk_write_time_seconds FROM pg_stat_statements t1 JOIN pg_roles t2 ON (t1.userid=t2.oid) JOIN pg_database t3 ON (t1.dbid=t3.oid) WHERE t2.rolname != 'rdsadmin'"
       master: true
       metrics:
         - rolname:
@@ -326,6 +351,50 @@ config:
             usage: "COUNTER"
             description: "Total time the statement spent writing blocks, in milliseconds (if track_io_timing is enabled, otherwise zero)"
 
+    pg_stat_activity:
+      query: |
+        WITH
+          metrics AS (
+            SELECT
+              application_name,
+              SUM(EXTRACT(EPOCH FROM (CURRENT_TIMESTAMP - state_change))::bigint)::float AS process_idle_seconds_sum,
+              COUNT(*) AS process_idle_seconds_count
+            FROM pg_stat_activity
+            WHERE state = 'idle'
+            GROUP BY application_name
+          ),
+          buckets AS (
+            SELECT
+              application_name,
+              le,
+              SUM(
+                CASE WHEN EXTRACT(EPOCH FROM (CURRENT_TIMESTAMP - state_change)) <= le
+                  THEN 1
+                  ELSE 0
+                END
+              )::bigint AS bucket
+            FROM
+              pg_stat_activity,
+              UNNEST(ARRAY[1, 2, 5, 15, 30, 60, 90, 120, 300]) AS le
+            GROUP BY application_name, le
+            ORDER BY application_name, le
+          )
+        SELECT
+          application_name,
+          process_idle_seconds_sum,
+          process_idle_seconds_count,
+          ARRAY_AGG(le) AS process_idle_seconds,
+          ARRAY_AGG(bucket) AS process_idle_seconds_bucket
+        FROM metrics JOIN buckets USING (application_name)
+        GROUP BY 1, 2, 3
+      metrics:
+        - application_name:
+            usage: "LABEL"
+            description: "Application Name"
+        - process_idle_seconds:
+            usage: "HISTOGRAM"
+            description: "Idle time of server processes"
+
 nodeSelector: {}
 
 tolerations: []

```

## 2.1.0

**Release date:** 2021-03-18

![AppVersion: 0.9.0](https://img.shields.io/static/v1?label=AppVersion&message=0.9.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-postgres-exporter] Update to postgres_exporter 0.9.0 (#773)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index 802afdf2..af280051 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -5,8 +5,8 @@
 replicaCount: 1
 
 image:
-  repository: wrouesnel/postgres_exporter
-  tag: v0.8.0
+  repository: quay.io/prometheuscommunity/postgres-exporter
+  tag: v0.9.0
   pullPolicy: IfNotPresent
 
   ## Optionally specify an array of imagePullSecrets.

```

## 2.0.0

**Release date:** 2021-03-07

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-postgres-exporter] Migrate to chart v2 (#540)

### Default value changes

```diff
# No changes in this release
```

## 1.10.0

**Release date:** 2021-02-20

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Add annotations to postgres-exporter ServiceAccount (#679)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index 2346320d..802afdf2 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -81,6 +81,8 @@ serviceAccount:
   # The name of the ServiceAccount to use.
   # If not set and create is true, a name is generated using the fullname template
   name:
+  # Add annotations to the ServiceAccount, useful for EKS IAM Roles for Service Accounts or Google Workload Identity.
+  annotations: {}
 
 securityContext: {}
   # The securityContext this Pod should use. See https://kubernetes.io/docs/concepts/policy/security-context/ for more.

```

## 1.9.1

**Release date:** 2021-02-12

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-postgres-exporter] add opportunity to specify log.level in args (#672)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index 40ae3434..2346320d 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -114,6 +114,8 @@ config:
   autoDiscoverDatabases: false
   excludeDatabases: []
   constantLabels: {}
+  # possible values debug, info, warn, error, fatal
+  logLevel: ""
   # this are the defaults queries that the exporter will run, extracted from: https://github.com/wrouesnel/postgres_exporter/blob/master/queries.yaml
   queries: |-
     pg_replication:

```

## 1.9.0

**Release date:** 2021-01-26

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-postgres-exporter] add initialDelaySeconds configuration for probes (#612)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index 17a95efc..40ae3434 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -332,6 +332,13 @@ annotations: {}
 
 podLabels: {}
 
+# Configurable health checks
+livenessProbe:
+  initialDelaySeconds: 0
+
+readinessProbe:
+  initialDelaySeconds: 0
+
 # Init containers, e. g. for secrets creation before the exporter
 initContainers: []
   # - name:

```

## 1.8.0

**Release date:** 2021-01-20

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-postgres-exporter]add optional constantLabels (#600)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index 5752d76c..17a95efc 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -113,6 +113,7 @@ config:
   disableSettingsMetrics: false
   autoDiscoverDatabases: false
   excludeDatabases: []
+  constantLabels: {}
   # this are the defaults queries that the exporter will run, extracted from: https://github.com/wrouesnel/postgres_exporter/blob/master/queries.yaml
   queries: |-
     pg_replication:

```

## 1.7.0

**Release date:** 2021-01-16

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [postgresql] (feat): add imagePullSecrets (#587)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index 51ba92ea..5752d76c 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -9,6 +9,13 @@ image:
   tag: v0.8.0
   pullPolicy: IfNotPresent
 
+  ## Optionally specify an array of imagePullSecrets.
+  ## Secrets must be manually created in the namespace.
+  ## ref: https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/
+  ##
+  # pullSecrets:
+  #   - myRegistrKeySecretName
+
 service:
   type: ClusterIP
   port: 80

```

## 1.6.0

**Release date:** 2020-12-22

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-postgres-exporter] Add ServiceMonitor metricRelabeling and targetLabels (#512)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index 58577a92..51ba92ea 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -30,6 +30,10 @@ serviceMonitor:
   # labels:
   # Set timeout for scrape
   # timeout: 10s
+  # Set of labels to transfer from the Kubernetes Service onto the target
+  # targetLabels: []
+  # MetricRelabelConfigs to apply to samples before ingestion
+  # metricRelabelings: []
 
 prometheusRule:
   enabled: false

```

## 1.5.0

**Release date:** 2020-11-30

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-postgres-exporter] add initContainers (#425)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index 5d0f1bc9..58577a92 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -320,6 +320,14 @@ annotations: {}
 
 podLabels: {}
 
+# Init containers, e. g. for secrets creation before the exporter
+initContainers: []
+  # - name:
+  #   image:
+  #   volumeMounts:
+  #     - name: creds
+  #       mountPath: /creds
+
 # Additional sidecar containers, e. g. for a database proxy, such as Google's cloudsql-proxy
 extraContainers: |
 

```

## 1.4.0

**Release date:** 2020-11-18

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* feat(prometheus-postgres-exporter): add configurable apiVersion (#377)

### Default value changes

```diff
# No changes in this release
```

## 1.3.4

**Release date:** 2020-11-10

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-postgres-exporter] Add prometheusrule support (#337)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index fe15a627..5d0f1bc9 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -31,6 +31,22 @@ serviceMonitor:
   # Set timeout for scrape
   # timeout: 10s
 
+prometheusRule:
+  enabled: false
+  additionalLabels: {}
+  namespace: ""
+  rules: []
+    ## These are just examples rules, please adapt them to your needs.
+    ## Make sure to constraint the rules to the current prometheus-postgres-exporter service.
+    # - alert: HugeReplicationLag
+    #   expr: pg_replication_lag{service="{{ template "prometheus-postgres-exporter.fullname" . }}"} / 3600 > 1
+    #   for: 1m
+    #   labels:
+    #     severity: critical
+    #   annotations:
+    #     description: replication for {{ template "prometheus-postgres-exporter.fullname" . }} PostgreSQL is lagging by {{ "{{ $value }}" }} hour(s).
+    #     summary: PostgreSQL replication is lagging by {{ "{{ $value }}" }} hour(s).
+
 resources: {}
   # We usually recommend not to specify default resources and to leave this as a conscious
   # choice for the user. This also increases chances charts run on environments with little

```

## 1.3.3

**Release date:** 2020-09-17

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-postgres-exporter] add extraVolumeMounts and examples (#107)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index bcd9d53c..fe15a627 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -309,3 +309,20 @@ extraContainers: |
 
 # Additional volumes, e. g. for secrets used in an extraContainer
 extraVolumes: |
+
+# Uncomment for mounting custom ca-certificates
+#  - name: ssl-certs
+#    secret:
+#      defaultMode: 420
+#      items:
+#      - key: ca-certificates.crt
+#        path: ca-certificates.crt
+#      secretName: ssl-certs
+
+# Additional volume mounts
+extraVolumeMounts: |
+
+# Uncomment for mounting custom ca-certificates file into container
+#  - name: ssl-certs
+#    mountPath: /etc/ssl/certs/ca-certificates.crt
+#    subPath: ca-certificates.crt

```

## 1.3.2

**Release date:** 2020-09-09

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-postgres-exporter] Add zanhsieh as maintainer (#71)

### Default value changes

```diff
# No changes in this release
```

## 1.3.1

**Release date:** 2020-08-20

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Prep initial charts indexing (#14)

### Default value changes

```diff
# No changes in this release
```

## 1.3.0

**Release date:** 2020-04-20

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-postgres-exporter] Add support to configure dataso… (#21804)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index 6c62ce75..bcd9d53c 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -61,6 +61,7 @@ securityContext: {}
 
 config:
   datasource:
+    # Specify one of both datasource or datasourceSecret
     host:
     user: postgres
     # Only one of password and passwordSecret can be specified
@@ -74,6 +75,13 @@ config:
     port: "5432"
     database: ''
     sslmode: disable
+  datasourceSecret: {}
+    # Specifies if datasource should be sourced from secret value in format: postgresql://login:password@hostname:port/dbname?sslmode=disable
+    # Multiple Postgres databases can be configured by comma separated postgres connection strings
+    # Secret name
+    #  name:
+    # Connection string key inside secret
+    #  key:
   disableDefaultMetrics: false
   disableSettingsMetrics: false
   autoDiscoverDatabases: false

```

## 1.2.0

**Release date:** 2019-12-23

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-postgres-exporter] update app to v0.8.0 (#19749)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index f770b4a4..6c62ce75 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -6,7 +6,7 @@ replicaCount: 1
 
 image:
   repository: wrouesnel/postgres_exporter
-  tag: v0.5.1
+  tag: v0.8.0
   pullPolicy: IfNotPresent
 
 service:
@@ -75,12 +75,14 @@ config:
     database: ''
     sslmode: disable
   disableDefaultMetrics: false
+  disableSettingsMetrics: false
   autoDiscoverDatabases: false
   excludeDatabases: []
   # this are the defaults queries that the exporter will run, extracted from: https://github.com/wrouesnel/postgres_exporter/blob/master/queries.yaml
   queries: |-
     pg_replication:
-      query: "SELECT EXTRACT(EPOCH FROM (now() - pg_last_xact_replay_timestamp()))::INT as lag"
+      query: "SELECT EXTRACT(EPOCH FROM (now() - pg_last_xact_replay_timestamp())) as lag"
+      master: true
       metrics:
         - lag:
             usage: "GAUGE"
@@ -88,14 +90,18 @@ config:
 
     pg_postmaster:
       query: "SELECT pg_postmaster_start_time as start_time_seconds from pg_postmaster_start_time()"
+      master: true
       metrics:
         - start_time_seconds:
             usage: "GAUGE"
             description: "Time at which postmaster started"
 
     pg_stat_user_tables:
-      query: "SELECT schemaname, relname, seq_scan, seq_tup_read, idx_scan, idx_tup_fetch, n_tup_ins, n_tup_upd, n_tup_del, n_tup_hot_upd, n_live_tup, n_dead_tup, n_mod_since_analyze, last_vacuum, last_autovacuum, last_analyze, last_autoanalyze, vacuum_count, autovacuum_count, analyze_count, autoanalyze_count FROM pg_stat_user_tables"
+      query: "SELECT current_database() datname, schemaname, relname, seq_scan, seq_tup_read, idx_scan, idx_tup_fetch, n_tup_ins, n_tup_upd, n_tup_del, n_tup_hot_upd, n_live_tup, n_dead_tup, n_mod_since_analyze, COALESCE(last_vacuum, '1970-01-01Z'), COALESCE(last_vacuum, '1970-01-01Z') as last_vacuum, COALESCE(last_autovacuum, '1970-01-01Z') as last_autovacuum, COALESCE(last_analyze, '1970-01-01Z') as last_analyze, COALESCE(last_autoanalyze, '1970-01-01Z') as last_autoanalyze, vacuum_count, autovacuum_count, analyze_count, autoanalyze_count FROM pg_stat_user_tables"
       metrics:
+        - datname:
+            usage: "LABEL"
+            description: "Name of current database"
         - schemaname:
             usage: "LABEL"
             description: "Name of the schema that this table is in"
@@ -161,8 +167,11 @@ config:
             description: "Number of times this table has been analyzed by the autovacuum daemon"
 
     pg_statio_user_tables:
-      query: "SELECT schemaname, relname, heap_blks_read, heap_blks_hit, idx_blks_read, idx_blks_hit, toast_blks_read, toast_blks_hit, tidx_blks_read, tidx_blks_hit FROM pg_statio_user_tables"
+      query: "SELECT current_database() datname, schemaname, relname, heap_blks_read, heap_blks_hit, idx_blks_read, idx_blks_hit, toast_blks_read, toast_blks_hit, tidx_blks_read, tidx_blks_hit FROM pg_statio_user_tables"
       metrics:
+        - datname:
+            usage: "LABEL"
+            description: "Name of current database"
         - schemaname:
             usage: "LABEL"
             description: "Name of the schema that this table is in"
@@ -195,15 +204,88 @@ config:
             description: "Number of buffer hits in this table's TOAST table indexes (if any)"
 
     pg_database:
-      query: " SELECT pg_database.datname, pg_database_size(pg_database.datname) as size FROM pg_database"
+      query: "SELECT pg_database.datname, pg_database_size(pg_database.datname) as size FROM pg_database"
+      master: true
+      cache_seconds: 30
       metrics:
         - datname:
             usage: "LABEL"
             description: "Name of the database"
-        - size:
+        - size_bytes:
             usage: "GAUGE"
             description: "Disk space used by the database"
 
+    pg_stat_statements:
+      query: "SELECT t2.rolname, t3.datname, queryid, calls, total_time / 1000 as total_time_seconds, min_time / 1000 as min_time_seconds, max_time / 1000 as max_time_seconds, mean_time / 1000 as mean_time_seconds, stddev_time / 1000 as stddev_time_seconds, rows, shared_blks_hit, shared_blks_read, shared_blks_dirtied, shared_blks_written, local_blks_hit, local_blks_read, local_blks_dirtied, local_blks_written, temp_blks_read, temp_blks_written, blk_read_time / 1000 as blk_read_time_seconds, blk_write_time / 1000 as blk_write_time_seconds FROM pg_stat_statements t1 join pg_roles t2 on (t1.userid=t2.oid) join pg_database t3 on (t1.dbid=t3.oid)"
+      master: true
+      metrics:
+        - rolname:
+            usage: "LABEL"
+            description: "Name of user"
+        - datname:
+            usage: "LABEL"
+            description: "Name of database"
+        - queryid:
+            usage: "LABEL"
+            description: "Query ID"
+        - calls:
+            usage: "COUNTER"
+            description: "Number of times executed"
+        - total_time_seconds:
+            usage: "COUNTER"
+            description: "Total time spent in the statement, in milliseconds"
+        - min_time_seconds:
+            usage: "GAUGE"
+            description: "Minimum time spent in the statement, in milliseconds"
+        - max_time_seconds:
+            usage: "GAUGE"
+            description: "Maximum time spent in the statement, in milliseconds"
+        - mean_time_seconds:
+            usage: "GAUGE"
+            description: "Mean time spent in the statement, in milliseconds"
+        - stddev_time_seconds:
+            usage: "GAUGE"
+            description: "Population standard deviation of time spent in the statement, in milliseconds"
+        - rows:
+            usage: "COUNTER"
+            description: "Total number of rows retrieved or affected by the statement"
+        - shared_blks_hit:
+            usage: "COUNTER"
+            description: "Total number of shared block cache hits by the statement"
+        - shared_blks_read:
+            usage: "COUNTER"
+            description: "Total number of shared blocks read by the statement"
+        - shared_blks_dirtied:
+            usage: "COUNTER"
+            description: "Total number of shared blocks dirtied by the statement"
+        - shared_blks_written:
+            usage: "COUNTER"
+            description: "Total number of shared blocks written by the statement"
+        - local_blks_hit:
+            usage: "COUNTER"
+            description: "Total number of local block cache hits by the statement"
+        - local_blks_read:
+            usage: "COUNTER"
+            description: "Total number of local blocks read by the statement"
+        - local_blks_dirtied:
+            usage: "COUNTER"
+            description: "Total number of local blocks dirtied by the statement"
+        - local_blks_written:
+            usage: "COUNTER"
+            description: "Total number of local blocks written by the statement"
+        - temp_blks_read:
+            usage: "COUNTER"
+            description: "Total number of temp blocks read by the statement"
+        - temp_blks_written:
+            usage: "COUNTER"
+            description: "Total number of temp blocks written by the statement"
+        - blk_read_time_seconds:
+            usage: "COUNTER"
+            description: "Total time the statement spent reading blocks, in milliseconds (if track_io_timing is enabled, otherwise zero)"
+        - blk_write_time_seconds:
+            usage: "COUNTER"
+            description: "Total time the statement spent writing blocks, in milliseconds (if track_io_timing is enabled, otherwise zero)"
+
 nodeSelector: {}
 
 tolerations: []

```

## 1.1.1

**Release date:** 2019-11-12

![AppVersion: 0.5.1](https://img.shields.io/static/v1?label=AppVersion&message=0.5.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Update README of prometheus-postgres-exporter (#18475)

### Default value changes

```diff
# No changes in this release
```

## 1.1.0

**Release date:** 2019-10-11

![AppVersion: 0.5.1](https://img.shields.io/static/v1?label=AppVersion&message=0.5.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-postgres-exporter] Add feature 'auto-discover-databases' (#17826)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index f24fcfa2..f770b4a4 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -75,6 +75,8 @@ config:
     database: ''
     sslmode: disable
   disableDefaultMetrics: false
+  autoDiscoverDatabases: false
+  excludeDatabases: []
   # this are the defaults queries that the exporter will run, extracted from: https://github.com/wrouesnel/postgres_exporter/blob/master/queries.yaml
   queries: |-
     pg_replication:

```

## 1.0.0

**Release date:** 2019-10-06

![AppVersion: 0.5.1](https://img.shields.io/static/v1?label=AppVersion&message=0.5.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-postgres-exporter] Fix compatibility with k8s 1.16 (#17694)

### Default value changes

```diff
# No changes in this release
```

## 0.7.2

**Release date:** 2019-08-13

![AppVersion: 0.5.1](https://img.shields.io/static/v1?label=AppVersion&message=0.5.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-postgres-exporter]: add the ability to create a ServiceMonitor (#14918)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index 816d5f3f..f24fcfa2 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -17,6 +17,20 @@ service:
   labels: {}
   annotations: {}
 
+serviceMonitor:
+  # When set true then use a ServiceMonitor to configure scraping
+  enabled: false
+  # Set the namespace the ServiceMonitor should be deployed
+  # namespace: monitoring
+  # Set how frequently Prometheus should scrape
+  # interval: 30s
+  # Set path to cloudwatch-exporter telemtery-path
+  # telemetryPath: /metrics
+  # Set labels for the ServiceMonitor, use this to define your scrape label for Prometheus Operator
+  # labels:
+  # Set timeout for scrape
+  # timeout: 10s
+
 resources: {}
   # We usually recommend not to specify default resources and to leave this as a conscious
   # choice for the user. This also increases chances charts run on environments with little

```

## 0.7.1

**Release date:** 2019-07-23

![AppVersion: 0.5.1](https://img.shields.io/static/v1?label=AppVersion&message=0.5.1&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-postgres-exporter] Bump to prometheus-postgres-exporter v0.5.1 + use prometheus-postgres-exporter.chart (#15502)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index 2a01c159..816d5f3f 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -6,7 +6,7 @@ replicaCount: 1
 
 image:
   repository: wrouesnel/postgres_exporter
-  tag: v0.5.0
+  tag: v0.5.1
   pullPolicy: IfNotPresent
 
 service:

```

## 0.7.0

**Release date:** 2019-07-05

![AppVersion: 0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=0.5.0&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-postgres-exporter] upgrade to new version of postgres exporter (#15239)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index e40cae10..2a01c159 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -6,7 +6,7 @@ replicaCount: 1
 
 image:
   repository: wrouesnel/postgres_exporter
-  tag: v0.4.7
+  tag: v0.5.0
   pullPolicy: IfNotPresent
 
 service:

```

## 0.6.3

**Release date:** 2019-06-27

![AppVersion: 0.4.7](https://img.shields.io/static/v1?label=AppVersion&message=0.4.7&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-postgres-exporter] Allow specifying a securityContext (#15074)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index d71cf852..e40cae10 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -41,6 +41,10 @@ serviceAccount:
   # If not set and create is true, a name is generated using the fullname template
   name:
 
+securityContext: {}
+  # The securityContext this Pod should use. See https://kubernetes.io/docs/concepts/policy/security-context/ for more.
+  # runAsUser: 65534
+
 config:
   datasource:
     host:

```

## 0.6.2

**Release date:** 2019-03-16

![AppVersion: 0.4.7](https://img.shields.io/static/v1?label=AppVersion&message=0.4.7&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-postgres-exporter] Allow to specify secret for database password (#11392)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index 3373cf55..d71cf852 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -44,8 +44,15 @@ serviceAccount:
 config:
   datasource:
     host:
-    user:
-    password:
+    user: postgres
+    # Only one of password and passwordSecret can be specified
+    password: somepassword
+    # Specify passwordSecret if DB password is stored in secret.
+    passwordSecret: {}
+    # Secret name
+    #  name:
+    # Password key inside secret
+    #  key:
     port: "5432"
     database: ''
     sslmode: disable

```

## 0.6.1

**Release date:** 2018-12-18

![AppVersion: 0.4.7](https://img.shields.io/static/v1?label=AppVersion&message=0.4.7&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Update postgres-exporter to app v0.4.7 (#10079)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index 14f11a55..3373cf55 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -6,7 +6,7 @@ replicaCount: 1
 
 image:
   repository: wrouesnel/postgres_exporter
-  tag: v0.4.6
+  tag: v0.4.7
   pullPolicy: IfNotPresent
 
 service:
@@ -133,6 +133,40 @@ config:
             usage: "COUNTER"
             description: "Number of times this table has been analyzed by the autovacuum daemon"
 
+    pg_statio_user_tables:
+      query: "SELECT schemaname, relname, heap_blks_read, heap_blks_hit, idx_blks_read, idx_blks_hit, toast_blks_read, toast_blks_hit, tidx_blks_read, tidx_blks_hit FROM pg_statio_user_tables"
+      metrics:
+        - schemaname:
+            usage: "LABEL"
+            description: "Name of the schema that this table is in"
+        - relname:
+            usage: "LABEL"
+            description: "Name of this table"
+        - heap_blks_read:
+            usage: "COUNTER"
+            description: "Number of disk blocks read from this table"
+        - heap_blks_hit:
+            usage: "COUNTER"
+            description: "Number of buffer hits in this table"
+        - idx_blks_read:
+            usage: "COUNTER"
+            description: "Number of disk blocks read from all indexes on this table"
+        - idx_blks_hit:
+            usage: "COUNTER"
+            description: "Number of buffer hits in all indexes on this table"
+        - toast_blks_read:
+            usage: "COUNTER"
+            description: "Number of disk blocks read from this table's TOAST table (if any)"
+        - toast_blks_hit:
+            usage: "COUNTER"
+            description: "Number of buffer hits in this table's TOAST table (if any)"
+        - tidx_blks_read:
+            usage: "COUNTER"
+            description: "Number of disk blocks read from this table's TOAST table indexes (if any)"
+        - tidx_blks_hit:
+            usage: "COUNTER"
+            description: "Number of buffer hits in this table's TOAST table indexes (if any)"
+
     pg_database:
       query: " SELECT pg_database.datname, pg_database_size(pg_database.datname) as size FROM pg_database"
       metrics:

```

## 0.6.0

**Release date:** 2018-12-17

![AppVersion: 0.4.6](https://img.shields.io/static/v1?label=AppVersion&message=0.4.6&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* postgres-exporter labels and port name (#10027)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index c3de2355..14f11a55 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -13,6 +13,8 @@ service:
   type: ClusterIP
   port: 80
   targetPort: 9187
+  name: http
+  labels: {}
   annotations: {}
 
 resources: {}
@@ -149,6 +151,8 @@ affinity: {}
 
 annotations: {}
 
+podLabels: {}
+
 # Additional sidecar containers, e. g. for a database proxy, such as Google's cloudsql-proxy
 extraContainers: |
 

```

## 0.5.1

**Release date:** 2018-12-16

![AppVersion: 0.4.6](https://img.shields.io/static/v1?label=AppVersion&message=0.4.6&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-postgres-exporter] Add .Values.service.annotations (#10038)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index b54dd067..c3de2355 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -13,6 +13,7 @@ service:
   type: ClusterIP
   port: 80
   targetPort: 9187
+  annotations: {}
 
 resources: {}
   # We usually recommend not to specify default resources and to leave this as a conscious

```

## 0.5.0

**Release date:** 2018-10-13

![AppVersion: 0.4.6](https://img.shields.io/static/v1?label=AppVersion&message=0.4.6&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-postgres-exporter] Added flag for disabling default metrics (#8134)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index b740095d..b54dd067 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -46,6 +46,7 @@ config:
     port: "5432"
     database: ''
     sslmode: disable
+  disableDefaultMetrics: false
   # this are the defaults queries that the exporter will run, extracted from: https://github.com/wrouesnel/postgres_exporter/blob/master/queries.yaml
   queries: |-
     pg_replication:

```

## 0.4.0

**Release date:** 2018-06-24

![AppVersion: 0.4.6](https://img.shields.io/static/v1?label=AppVersion&message=0.4.6&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-postgres-exporter] add support for Google Cloud SQL (#6009)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index 885153d7..b740095d 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -44,6 +44,7 @@ config:
     user:
     password:
     port: "5432"
+    database: ''
     sslmode: disable
   # this are the defaults queries that the exporter will run, extracted from: https://github.com/wrouesnel/postgres_exporter/blob/master/queries.yaml
   queries: |-
@@ -145,3 +146,9 @@ tolerations: []
 affinity: {}
 
 annotations: {}
+
+# Additional sidecar containers, e. g. for a database proxy, such as Google's cloudsql-proxy
+extraContainers: |
+
+# Additional volumes, e. g. for secrets used in an extraContainer
+extraVolumes: |

```

## 0.3.0

**Release date:** 2018-06-17

![AppVersion: 0.4.6](https://img.shields.io/static/v1?label=AppVersion&message=0.4.6&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-postgres-exporter] add RBAC resources (#6008)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index c84faa4b..885153d7 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -26,6 +26,11 @@ resources: {}
   #   cpu: 100m
   #   memory: 128Mi
 
+rbac:
+  # Specifies whether RBAC resources should be created
+  create: true
+  # Specifies whether a PodSecurityPolicy should be created
+  pspEnabled: true
 serviceAccount:
   # Specifies whether a ServiceAccount should be created
   create: true

```

## 0.2.0

**Release date:** 2018-06-11

![AppVersion: 0.4.6](https://img.shields.io/static/v1?label=AppVersion&message=0.4.6&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* adding annotations (#5607)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index 91272671..c84faa4b 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -138,3 +138,5 @@ nodeSelector: {}
 tolerations: []
 
 affinity: {}
+
+annotations: {}

```

## 0.1.3

**Release date:** 2018-05-27

![AppVersion: 0.4.6](https://img.shields.io/static/v1?label=AppVersion&message=0.4.6&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [prometheus-postgres-exporter]Typo fix: tables lists->table lists (#5720)

### Default value changes

```diff
# No changes in this release
```

## 0.1.2

**Release date:** 2018-05-25

![AppVersion: 0.4.6](https://img.shields.io/static/v1?label=AppVersion&message=0.4.6&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* fix url which not exist anymore (#4589)

### Default value changes

```diff
# No changes in this release
```

## 0.1.1

**Release date:** 2018-05-24

![AppVersion: 0.4.6](https://img.shields.io/static/v1?label=AppVersion&message=0.4.6&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* [stable/prometheus-postgres-exporter] Fixes #5059, #5472, version update (#5473)

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index 63e96829..91272671 100644
--- a/charts/prometheus-postgres-exporter/values.yaml
+++ b/charts/prometheus-postgres-exporter/values.yaml
@@ -6,7 +6,7 @@ replicaCount: 1
 
 image:
   repository: wrouesnel/postgres_exporter
-  tag: v0.4.4
+  tag: v0.4.6
   pullPolicy: IfNotPresent
 
 service:

```

## 0.1.0

**Release date:** 2018-03-16

![AppVersion: 0.4.4](https://img.shields.io/static/v1?label=AppVersion&message=0.4.4&color=success)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* Create prometheus postgres exporter chart (#4076)

### Default value changes

```diff
# Default values for prometheus-postgres-exporter.
# This is a YAML-formatted file.
# Declare variables to be passed into your templates.

replicaCount: 1

image:
  repository: wrouesnel/postgres_exporter
  tag: v0.4.4
  pullPolicy: IfNotPresent

service:
  type: ClusterIP
  port: 80
  targetPort: 9187

resources: {}
  # We usually recommend not to specify default resources and to leave this as a conscious
  # choice for the user. This also increases chances charts run on environments with little
  # resources, such as Minikube. If you do want to specify resources, uncomment the following
  # lines, adjust them as necessary, and remove the curly braces after 'resources:'.
  # limits:
  #   cpu: 100m
  #    memory: 128Mi
  # requests:
  #   cpu: 100m
  #   memory: 128Mi

serviceAccount:
  # Specifies whether a ServiceAccount should be created
  create: true
  # The name of the ServiceAccount to use.
  # If not set and create is true, a name is generated using the fullname template
  name:

config:
  datasource:
    host:
    user:
    password:
    port: "5432"
    sslmode: disable
  # this are the defaults queries that the exporter will run, extracted from: https://github.com/wrouesnel/postgres_exporter/blob/master/queries.yaml
  queries: |-
    pg_replication:
      query: "SELECT EXTRACT(EPOCH FROM (now() - pg_last_xact_replay_timestamp()))::INT as lag"
      metrics:
        - lag:
            usage: "GAUGE"
            description: "Replication lag behind master in seconds"

    pg_postmaster:
      query: "SELECT pg_postmaster_start_time as start_time_seconds from pg_postmaster_start_time()"
      metrics:
        - start_time_seconds:
            usage: "GAUGE"
            description: "Time at which postmaster started"

    pg_stat_user_tables:
      query: "SELECT schemaname, relname, seq_scan, seq_tup_read, idx_scan, idx_tup_fetch, n_tup_ins, n_tup_upd, n_tup_del, n_tup_hot_upd, n_live_tup, n_dead_tup, n_mod_since_analyze, last_vacuum, last_autovacuum, last_analyze, last_autoanalyze, vacuum_count, autovacuum_count, analyze_count, autoanalyze_count FROM pg_stat_user_tables"
      metrics:
        - schemaname:
            usage: "LABEL"
            description: "Name of the schema that this table is in"
        - relname:
            usage: "LABEL"
            description: "Name of this table"
        - seq_scan:
            usage: "COUNTER"
            description: "Number of sequential scans initiated on this table"
        - seq_tup_read:
            usage: "COUNTER"
            description: "Number of live rows fetched by sequential scans"
        - idx_scan:
            usage: "COUNTER"
            description: "Number of index scans initiated on this table"
        - idx_tup_fetch:
            usage: "COUNTER"
            description: "Number of live rows fetched by index scans"
        - n_tup_ins:
            usage: "COUNTER"
            description: "Number of rows inserted"
        - n_tup_upd:
            usage: "COUNTER"
            description: "Number of rows updated"
        - n_tup_del:
            usage: "COUNTER"
            description: "Number of rows deleted"
        - n_tup_hot_upd:
            usage: "COUNTER"
            description: "Number of rows HOT updated (i.e., with no separate index update required)"
        - n_live_tup:
            usage: "GAUGE"
            description: "Estimated number of live rows"
        - n_dead_tup:
            usage: "GAUGE"
            description: "Estimated number of dead rows"
        - n_mod_since_analyze:
            usage: "GAUGE"
            description: "Estimated number of rows changed since last analyze"
        - last_vacuum:
            usage: "GAUGE"
            description: "Last time at which this table was manually vacuumed (not counting VACUUM FULL)"
        - last_autovacuum:
            usage: "GAUGE"
            description: "Last time at which this table was vacuumed by the autovacuum daemon"
        - last_analyze:
            usage: "GAUGE"
            description: "Last time at which this table was manually analyzed"
        - last_autoanalyze:
            usage: "GAUGE"
            description: "Last time at which this table was analyzed by the autovacuum daemon"
        - vacuum_count:
            usage: "COUNTER"
            description: "Number of times this table has been manually vacuumed (not counting VACUUM FULL)"
        - autovacuum_count:
            usage: "COUNTER"
            description: "Number of times this table has been vacuumed by the autovacuum daemon"
        - analyze_count:
            usage: "COUNTER"
            description: "Number of times this table has been manually analyzed"
        - autoanalyze_count:
            usage: "COUNTER"
            description: "Number of times this table has been analyzed by the autovacuum daemon"

    pg_database:
      query: " SELECT pg_database.datname, pg_database_size(pg_database.datname) as size FROM pg_database"
      metrics:
        - datname:
            usage: "LABEL"
            description: "Name of the database"
        - size:
            usage: "GAUGE"
            description: "Disk space used by the database"

nodeSelector: {}

tolerations: []

affinity: {}

```

---
Autogenerated from Helm Chart and git history using [helm-changelog](https://github.com/mogensen/helm-changelog)
