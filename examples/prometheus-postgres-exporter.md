# Change Log

## 2.2.0 

**Release date:** 2021-03-19

![AppVersion: 0.9.0](https://img.shields.io/static/v1?label=AppVersion&message=0.9.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-postgres-exporter] Update to the latest default queries.yaml (#775) 

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index af28005..47cc956 100644
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

![AppVersion: 0.9.0](https://img.shields.io/static/v1?label=AppVersion&message=0.9.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-postgres-exporter] Update to postgres_exporter 0.9.0 (#773) 

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index 802afdf..af28005 100644
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

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success&logo=)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-postgres-exporter] Migrate to chart v2 (#540) 

### Default value changes

```diff
# No changes in this release
```

## 1.10.0 

**Release date:** 2021-02-20

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Add annotations to postgres-exporter ServiceAccount (#679) 

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index 2346320..802afdf 100644
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

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-postgres-exporter] add opportunity to specify log.level in args (#672) 

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index 40ae343..2346320 100644
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

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-postgres-exporter] add initialDelaySeconds configuration for probes (#612) 

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index 17a95ef..40ae343 100644
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

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-postgres-exporter]add optional constantLabels (#600) 

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index 5752d76..17a95ef 100644
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

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [postgresql] (feat): add imagePullSecrets (#587) 

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index 51ba92e..5752d76 100644
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

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-postgres-exporter] Add ServiceMonitor metricRelabeling and targetLabels (#512) 

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index 58577a9..51ba92e 100644
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

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-postgres-exporter] add initContainers (#425) 

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index 5d0f1bc..58577a9 100644
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

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* feat(prometheus-postgres-exporter): add configurable apiVersion (#377) 

### Default value changes

```diff
# No changes in this release
```

## 1.3.4 

**Release date:** 2020-11-10

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-postgres-exporter] Add prometheusrule support (#337) 

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index fe15a62..5d0f1bc 100644
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

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-postgres-exporter] add extraVolumeMounts and examples (#107) 

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index bcd9d53..fe15a62 100644
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

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-postgres-exporter] Add zanhsieh as maintainer (#71) 

### Default value changes

```diff
# No changes in this release
```

## 1.3.1 

**Release date:** 2020-08-20

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Prep initial charts indexing (#14) 

### Default value changes

```diff
# No changes in this release
```

## 1.3.0 

**Release date:** 2020-04-20

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-postgres-exporter] Add support to configure dataso… (#21804) 

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index 6c62ce7..bcd9d53 100644
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

![AppVersion: 0.8.0](https://img.shields.io/static/v1?label=AppVersion&message=0.8.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-postgres-exporter] update app to v0.8.0 (#19749) 

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index f770b4a..6c62ce7 100644
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

![AppVersion: 0.5.1](https://img.shields.io/static/v1?label=AppVersion&message=0.5.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Update README of prometheus-postgres-exporter (#18475) 

### Default value changes

```diff
# No changes in this release
```

## 1.1.0 

**Release date:** 2019-10-11

![AppVersion: 0.5.1](https://img.shields.io/static/v1?label=AppVersion&message=0.5.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-postgres-exporter] Add feature 'auto-discover-databases' (#17826) 

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index f24fcfa..f770b4a 100644
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

![AppVersion: 0.5.1](https://img.shields.io/static/v1?label=AppVersion&message=0.5.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-postgres-exporter] Fix compatibility with k8s 1.16 (#17694) 

### Default value changes

```diff
# No changes in this release
```

## 0.7.2 

**Release date:** 2019-08-13

![AppVersion: 0.5.1](https://img.shields.io/static/v1?label=AppVersion&message=0.5.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-postgres-exporter]: add the ability to create a ServiceMonitor (#14918) 

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index 816d5f3..f24fcfa 100644
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

**Release date:** 2019-07-22

![AppVersion: 0.5.1](https://img.shields.io/static/v1?label=AppVersion&message=0.5.1&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-postgres-exporter] Bump to prometheus-postgres-exporter v0.5.1 + use prometheus-postgres-exporter.chart (#15502) 

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index 2a01c15..816d5f3 100644
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

![AppVersion: 0.5.0](https://img.shields.io/static/v1?label=AppVersion&message=0.5.0&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-postgres-exporter] upgrade to new version of postgres exporter (#15239) 

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index e40cae1..2a01c15 100644
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

![AppVersion: 0.4.7](https://img.shields.io/static/v1?label=AppVersion&message=0.4.7&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-postgres-exporter] Allow specifying a securityContext (#15074) 

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index d71cf85..e40cae1 100644
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

![AppVersion: 0.4.7](https://img.shields.io/static/v1?label=AppVersion&message=0.4.7&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-postgres-exporter] Allow to specify secret for database password (#11392) 

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index 3373cf5..d71cf85 100644
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

**Release date:** 2018-12-17

![AppVersion: 0.4.7](https://img.shields.io/static/v1?label=AppVersion&message=0.4.7&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* Update postgres-exporter to app v0.4.7 (#10079) 

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index 14f11a5..3373cf5 100644
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

![AppVersion: 0.4.6](https://img.shields.io/static/v1?label=AppVersion&message=0.4.6&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* postgres-exporter labels and port name (#10027) 

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index c3de235..14f11a5 100644
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

![AppVersion: 0.4.6](https://img.shields.io/static/v1?label=AppVersion&message=0.4.6&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-postgres-exporter] Add .Values.service.annotations (#10038) 

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index b54dd06..c3de235 100644
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

![AppVersion: 0.4.6](https://img.shields.io/static/v1?label=AppVersion&message=0.4.6&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-postgres-exporter] Added flag for disabling default metrics (#8134) 

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index b740095..b54dd06 100644
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

![AppVersion: 0.4.6](https://img.shields.io/static/v1?label=AppVersion&message=0.4.6&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-postgres-exporter] add support for Google Cloud SQL (#6009) 

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index 885153d..b740095 100644
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

![AppVersion: 0.4.6](https://img.shields.io/static/v1?label=AppVersion&message=0.4.6&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-postgres-exporter] add RBAC resources (#6008) 

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index c84faa4..885153d 100644
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

![AppVersion: 0.4.6](https://img.shields.io/static/v1?label=AppVersion&message=0.4.6&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* adding annotations (#5607) 

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index 9127267..c84faa4 100644
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

![AppVersion: 0.4.6](https://img.shields.io/static/v1?label=AppVersion&message=0.4.6&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [prometheus-postgres-exporter]Typo fix: tables lists->table lists (#5720) 

### Default value changes

```diff
# No changes in this release
```

## 0.1.2 

**Release date:** 2018-05-25

![AppVersion: 0.4.6](https://img.shields.io/static/v1?label=AppVersion&message=0.4.6&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* fix url which not exist anymore (#4589) 

### Default value changes

```diff
# No changes in this release
```

## 0.1.1 

**Release date:** 2018-05-24

![AppVersion: 0.4.6](https://img.shields.io/static/v1?label=AppVersion&message=0.4.6&color=success&logo=)
![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm)
![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)


* [stable/prometheus-postgres-exporter] Fixes #5059, #5472, version update (#5473) 

### Default value changes

```diff
diff --git a/charts/prometheus-postgres-exporter/values.yaml b/charts/prometheus-postgres-exporter/values.yaml
index 63e9682..9127267 100644
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

![AppVersion: 0.4.4](https://img.shields.io/static/v1?label=AppVersion&message=0.4.4&color=success&logo=)
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
