/*
Package include imports all Module and MetricSet packages so that they register
their factories with the global registry. This package can be imported in the
main package to automatically register all of the standard supported Metricbeat
modules.
*/
package include

import (
	// This list is automatically generated by `make imports`
	_ "github.com/elastic/beats/metricbeat/module/apache"
	_ "github.com/elastic/beats/metricbeat/module/apache/status"
	_ "github.com/elastic/beats/metricbeat/module/couchbase"
	_ "github.com/elastic/beats/metricbeat/module/couchbase/bucket"
	_ "github.com/elastic/beats/metricbeat/module/couchbase/cluster"
	_ "github.com/elastic/beats/metricbeat/module/couchbase/node"
	_ "github.com/elastic/beats/metricbeat/module/docker"
	_ "github.com/elastic/beats/metricbeat/module/docker/container"
	_ "github.com/elastic/beats/metricbeat/module/docker/cpu"
	_ "github.com/elastic/beats/metricbeat/module/docker/diskio"
	_ "github.com/elastic/beats/metricbeat/module/docker/info"
	_ "github.com/elastic/beats/metricbeat/module/docker/memory"
	_ "github.com/elastic/beats/metricbeat/module/docker/network"
	_ "github.com/elastic/beats/metricbeat/module/docker/node"
	_ "github.com/elastic/beats/metricbeat/module/haproxy"
	_ "github.com/elastic/beats/metricbeat/module/haproxy/info"
	_ "github.com/elastic/beats/metricbeat/module/haproxy/stat"
	_ "github.com/elastic/beats/metricbeat/module/kafka"
	_ "github.com/elastic/beats/metricbeat/module/kafka/consumergroup"
	_ "github.com/elastic/beats/metricbeat/module/kafka/partition"
	_ "github.com/elastic/beats/metricbeat/module/mongodb"
	_ "github.com/elastic/beats/metricbeat/module/mongodb/status"
	_ "github.com/elastic/beats/metricbeat/module/mysql"
	_ "github.com/elastic/beats/metricbeat/module/mysql/status"
	_ "github.com/elastic/beats/metricbeat/module/nginx"
	_ "github.com/elastic/beats/metricbeat/module/nginx/stubstatus"
	_ "github.com/elastic/beats/metricbeat/module/postgresql"
	_ "github.com/elastic/beats/metricbeat/module/postgresql/activity"
	_ "github.com/elastic/beats/metricbeat/module/postgresql/bgwriter"
	_ "github.com/elastic/beats/metricbeat/module/postgresql/database"
	_ "github.com/elastic/beats/metricbeat/module/prometheus"
	_ "github.com/elastic/beats/metricbeat/module/prometheus/collector"
	_ "github.com/elastic/beats/metricbeat/module/prometheus/stats"
	_ "github.com/elastic/beats/metricbeat/module/redis"
	_ "github.com/elastic/beats/metricbeat/module/redis/info"
	_ "github.com/elastic/beats/metricbeat/module/redis/keyspace"
	_ "github.com/elastic/beats/metricbeat/module/system"
	_ "github.com/elastic/beats/metricbeat/module/system/core"
	_ "github.com/elastic/beats/metricbeat/module/system/cpu"
	_ "github.com/elastic/beats/metricbeat/module/system/diskio"
	_ "github.com/elastic/beats/metricbeat/module/system/filesystem"
	_ "github.com/elastic/beats/metricbeat/module/system/fsstat"
	_ "github.com/elastic/beats/metricbeat/module/system/load"
	_ "github.com/elastic/beats/metricbeat/module/system/memory"
	_ "github.com/elastic/beats/metricbeat/module/system/network"
	_ "github.com/elastic/beats/metricbeat/module/system/process"
	_ "github.com/elastic/beats/metricbeat/module/system/socket"
	_ "github.com/elastic/beats/metricbeat/module/zookeeper"
	_ "github.com/elastic/beats/metricbeat/module/zookeeper/mntr"
)
