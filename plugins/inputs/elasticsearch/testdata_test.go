package elasticsearch

const clusterHealthResponse = `
{
   "cluster_name": "elasticsearch_telegraf",
   "status": "green",
   "timed_out": false,
   "number_of_nodes": 3,
   "number_of_data_nodes": 3,
   "active_primary_shards": 5,
   "active_shards": 15,
   "relocating_shards": 0,
   "initializing_shards": 0,
   "unassigned_shards": 0,
   "delayed_unassigned_shards": 0,
   "number_of_pending_tasks": 0,
   "number_of_in_flight_fetch": 0,
   "task_max_waiting_in_queue_millis": 0,
   "active_shards_percent_as_number": 100.0
}
`

const clusterHealthResponseWithIndices = `
{
   "cluster_name": "elasticsearch_telegraf",
   "status": "green",
   "timed_out": false,
   "number_of_nodes": 3,
   "number_of_data_nodes": 3,
   "active_primary_shards": 5,
   "active_shards": 15,
   "relocating_shards": 0,
   "initializing_shards": 0,
   "unassigned_shards": 0,
   "delayed_unassigned_shards": 0,
   "number_of_pending_tasks": 0,
   "number_of_in_flight_fetch": 0,
   "task_max_waiting_in_queue_millis": 0,
   "active_shards_percent_as_number": 100.0,
   "indices": {
      "v1": {
         "status": "green",
         "number_of_shards": 10,
         "number_of_replicas": 1,
         "active_primary_shards": 10,
         "active_shards": 20,
         "relocating_shards": 0,
         "initializing_shards": 0,
         "unassigned_shards": 0
      },
      "v2": {
         "status": "red",
         "number_of_shards": 10,
         "number_of_replicas": 1,
         "active_primary_shards": 0,
         "active_shards": 0,
         "relocating_shards": 0,
         "initializing_shards": 0,
         "unassigned_shards": 20
      }
   }
}
`

var clusterHealthExpected = map[string]interface{}{
"active_primary_shards":      5,
"active_shards":      15,
"active_shards_percent_as_number":      100.0,
"initializing_shards":      0,
"number_of_data_nodes":      3,
"number_of_nodes":      3,
"number_of_pending_tasks":      0,
"relocating_shards":      0,
"status":      "green",
"status_code":      1,
"task_max_waiting_in_queue_millis":      0,
"timed_out":      false,
"unassigned_shards":      0,
}

var v1IndexExpected = map[string]interface{}{
	"status":                "green",
	"status_code":           1,
	"number_of_shards":      10,
	"number_of_replicas":    1,
	"active_primary_shards": 10,
	"active_shards":         20,
	"relocating_shards":     0,
	"initializing_shards":   0,
	"unassigned_shards":     0,
}

var v2IndexExpected = map[string]interface{}{
	"status":                "red",
	"status_code":           3,
	"number_of_shards":      10,
	"number_of_replicas":    1,
	"active_primary_shards": 0,
	"active_shards":         0,
	"relocating_shards":     0,
	"initializing_shards":   0,
	"unassigned_shards":     20,
}

const nodeStatsResponse = `
{
  "_nodes" : {
    "total" : 3,
    "successful" : 3,
    "failed" : 0
  },
  "cluster_name": "es-testcluster",
  "nodes": {
    "SDFsfSDFsdfFSDSDfSFDSDF": {
      "timestamp": 1544727513050,
      "name": "test.host.com",
      "transport_address": "10.43.168.58:9300",
      "host": "test",
      "ip": "10.43.168.58:9300",
      "roles": [
        "master",
        "ingest"
      ],
      "attributes": {
        "ml.machine_memory": "6442450944",
        "ml.max_open_jobs": "20",
        "xpack.installed": "true",
        "ml.enabled": "true"
      },
      "indices": {
        "docs": {
          "count": 0,
          "deleted": 0
        },
        "store": {
          "size_in_bytes": 0
        },
        "indexing": {
          "index_total": 0,
          "index_time_in_millis": 0,
          "index_current": 0,
          "index_failed": 0,
          "delete_total": 0,
          "delete_time_in_millis": 0,
          "delete_current": 0,
          "noop_update_total": 0,
          "is_throttled": false,
          "throttle_time_in_millis": 0
        },
        "get": {
          "total": 0,
          "time_in_millis": 0,
          "exists_total": 0,
          "exists_time_in_millis": 0,
          "missing_total": 0,
          "missing_time_in_millis": 0,
          "current": 0
        },
        "search": {
          "open_contexts": 0,
          "query_total": 0,
          "query_time_in_millis": 0,
          "query_current": 0,
          "fetch_total": 0,
          "fetch_time_in_millis": 0,
          "fetch_current": 0,
          "scroll_total": 0,
          "scroll_time_in_millis": 0,
          "scroll_current": 0,
          "suggest_total": 0,
          "suggest_time_in_millis": 0,
          "suggest_current": 0
        },
        "merges": {
          "current": 0,
          "current_docs": 0,
          "current_size_in_bytes": 0,
          "total": 0,
          "total_time_in_millis": 0,
          "total_docs": 0,
          "total_size_in_bytes": 0,
          "total_stopped_time_in_millis": 0,
          "total_throttled_time_in_millis": 0,
          "total_auto_throttle_in_bytes": 0
        },
        "refresh": {
          "total": 0,
          "total_time_in_millis": 0,
          "listeners": 0
        },
        "flush": {
          "total": 0,
          "periodic": 0,
          "total_time_in_millis": 0
        },
        "warmer": {
          "current": 0,
          "total": 0,
          "total_time_in_millis": 0
        },
        "query_cache": {
          "memory_size_in_bytes": 0,
          "total_count": 0,
          "hit_count": 0,
          "miss_count": 0,
          "cache_size": 0,
          "cache_count": 0,
          "evictions": 0
        },
        "fielddata": {
          "memory_size_in_bytes": 0,
          "evictions": 0
        },
        "completion": {
          "size_in_bytes": 0
        },
        "segments": {
          "count": 0,
          "memory_in_bytes": 0,
          "terms_memory_in_bytes": 0,
          "stored_fields_memory_in_bytes": 0,
          "term_vectors_memory_in_bytes": 0,
          "norms_memory_in_bytes": 0,
          "points_memory_in_bytes": 0,
          "doc_values_memory_in_bytes": 0,
          "index_writer_memory_in_bytes": 0,
          "version_map_memory_in_bytes": 0,
          "fixed_bit_set_memory_in_bytes": 0,
          "max_unsafe_auto_id_timestamp": -9223372036854776000,
          "file_sizes": {}
        },
        "translog": {
          "operations": 0,
          "size_in_bytes": 0,
          "uncommitted_operations": 0,
          "uncommitted_size_in_bytes": 0,
          "earliest_last_modified_age": 0
        },
        "request_cache": {
          "memory_size_in_bytes": 0,
          "evictions": 0,
          "hit_count": 0,
          "miss_count": 0
        },
        "recovery": {
          "current_as_source": 0,
          "current_as_target": 0,
          "throttle_time_in_millis": 0
        }
      },
      "os": {
        "timestamp": 1544727513051,
        "cpu": {
          "percent": 0,
          "load_average": {
            "1m": 0,
            "5m": 0.01,
            "15m": 0
          }
        },
        "mem": {
          "total_in_bytes": 66726227968,
          "free_in_bytes": 56699928576,
          "used_in_bytes": 10026299392,
          "free_percent": 85,
          "used_percent": 15
        },
        "swap": {
          "total_in_bytes": 0,
          "free_in_bytes": 0,
          "used_in_bytes": 0
        },
        "cgroup": {
          "cpuacct": {
            "control_group": "/",
            "usage_nanos": 6648413872901
          },
          "cpu": {
            "control_group": "/",
            "cfs_period_micros": 100000,
            "cfs_quota_micros": 100000,
            "stat": {
              "number_of_elapsed_periods": 2332636,
              "number_of_times_throttled": 21112,
              "time_throttled_nanos": 2087989489827
            }
          },
          "memory": {
            "control_group": "/",
            "limit_in_bytes": "6442450944",
            "usage_in_bytes": "5251289088"
          }
        }
      },
      "process": {
        "timestamp": 1544727513051,
        "open_file_descriptors": 280,
        "max_file_descriptors": 1048576,
        "cpu": {
          "percent": 0,
          "total_in_millis": 6639430
        },
        "mem": {
          "total_virtual_in_bytes": 10407563264
        }
      },
      "jvm": {
        "timestamp": 1544727513051,
        "uptime_in_millis": 581892320,
        "mem": {
          "heap_used_in_bytes": 227332320,
          "heap_used_percent": 5,
          "heap_committed_in_bytes": 4286251008,
          "heap_max_in_bytes": 4286251008,
          "non_heap_used_in_bytes": 139787192,
          "non_heap_committed_in_bytes": 168996864,
          "pools": {
            "young": {
              "used_in_bytes": 26252696,
              "max_in_bytes": 69795840,
              "peak_used_in_bytes": 69795840,
              "peak_max_in_bytes": 69795840
            },
            "survivor": {
              "used_in_bytes": 3779488,
              "max_in_bytes": 8716288,
              "peak_used_in_bytes": 8716288,
              "peak_max_in_bytes": 8716288
            },
            "old": {
              "used_in_bytes": 197315272,
              "max_in_bytes": 4207738880,
              "peak_used_in_bytes": 197315272,
              "peak_max_in_bytes": 4207738880
            }
          }
        },
        "threads": {
          "count": 76,
          "peak_count": 76
        },
        "gc": {
          "collectors": {
            "young": {
              "collection_count": 42793,
              "collection_time_in_millis": 1340364
            },
            "old": {
              "collection_count": 3,
              "collection_time_in_millis": 140
            }
          }
        },
        "buffer_pools": {
          "mapped": {
            "count": 0,
            "used_in_bytes": 0,
            "total_capacity_in_bytes": 0
          },
          "direct": {
            "count": 17,
            "used_in_bytes": 34005104,
            "total_capacity_in_bytes": 34005103
          }
        },
        "classes": {
          "current_loaded_count": 16731,
          "total_loaded_count": 16733,
          "total_unloaded_count": 2
        }
      },
      "thread_pool": {
        "analyze": {
          "threads": 0,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 0,
          "completed": 0
        },
        "fetch_shard_started": {
          "threads": 0,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 0,
          "completed": 0
        },
        "fetch_shard_store": {
          "threads": 0,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 0,
          "completed": 0
        },
        "flush": {
          "threads": 0,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 0,
          "completed": 0
        },
        "force_merge": {
          "threads": 0,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 0,
          "completed": 0
        },
        "generic": {
          "threads": 8,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 8,
          "completed": 1335639
        },
        "get": {
          "threads": 0,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 0,
          "completed": 0
        },
        "index": {
          "threads": 0,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 0,
          "completed": 0
        },
        "listener": {
          "threads": 0,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 0,
          "completed": 0
        },
        "management": {
          "threads": 4,
          "queue": 0,
          "active": 1,
          "rejected": 0,
          "largest": 4,
          "completed": 622652
        },
        "ml_autodetect": {
          "threads": 0,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 0,
          "completed": 0
        },
        "ml_datafeed": {
          "threads": 0,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 0,
          "completed": 0
        },
        "ml_utility": {
          "threads": 36,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 36,
          "completed": 36
        },
        "refresh": {
          "threads": 0,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 0,
          "completed": 0
        },
        "rollup_indexing": {
          "threads": 0,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 0,
          "completed": 0
        },
        "search": {
          "threads": 2,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 2,
          "completed": 281218
        },
        "security-token-key": {
          "threads": 0,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 0,
          "completed": 0
        },
        "snapshot": {
          "threads": 0,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 0,
          "completed": 0
        },
        "warmer": {
          "threads": 0,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 0,
          "completed": 0
        },
        "watcher": {
          "threads": 0,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 0,
          "completed": 0
        },
        "write": {
          "threads": 1,
          "queue": 0,
          "active": 0,
          "rejected": 0,
          "largest": 1,
          "completed": 65008
        }
      },
      "fs": {
        "timestamp": 1544727513051,
        "total": {
          "total_in_bytes": 21003583488,
          "free_in_bytes": 20682870784,
          "available_in_bytes": 19592351744
        },
        "data": [
          {
            "path": "/usr/share/elasticsearch/data/nodes/0",
            "mount": "/usr/share/elasticsearch/data (/dev/nvme1n1)",
            "type": "ext4",
            "total_in_bytes": 21003583488,
            "free_in_bytes": 20682870784,
            "available_in_bytes": 19592351744
          }
        ],
        "io_stats": {
          "devices": [
            {
              "device_name": "nvme1n1",
              "operations": 222457,
              "read_operations": 25,
              "write_operations": 222432,
              "read_kilobytes": 156,
              "write_kilobytes": 1618188
            }
          ],
          "total": {
            "operations": 222457,
            "read_operations": 25,
            "write_operations": 222432,
            "read_kilobytes": 156,
            "write_kilobytes": 1618188
          }
        }
      },
      "transport": {
        "server_open": 46,
        "rx_count": 6950414,
        "rx_size_in_bytes": 21891347203,
        "tx_count": 6950411,
        "tx_size_in_bytes": 17902587800
      },
      "http": {
        "current_open": 0,
        "total_opened": 38
      },
      "breakers": {
        "request": {
          "limit_size_in_bytes": 2571750604,
          "limit_size": "2.3gb",
          "estimated_size_in_bytes": 0,
          "estimated_size": "0b",
          "overhead": 1,
          "tripped": 0
        },
        "fielddata": {
          "limit_size_in_bytes": 2571750604,
          "limit_size": "2.3gb",
          "estimated_size_in_bytes": 0,
          "estimated_size": "0b",
          "overhead": 1.03,
          "tripped": 0
        },
        "in_flight_requests": {
          "limit_size_in_bytes": 4286251008,
          "limit_size": "3.9gb",
          "estimated_size_in_bytes": 1384,
          "estimated_size": "1.3kb",
          "overhead": 1,
          "tripped": 0
        },
        "accounting": {
          "limit_size_in_bytes": 4286251008,
          "limit_size": "3.9gb",
          "estimated_size_in_bytes": 0,
          "estimated_size": "0b",
          "overhead": 1,
          "tripped": 0
        },
        "parent": {
          "limit_size_in_bytes": 3000375705,
          "limit_size": "2.7gb",
          "estimated_size_in_bytes": 1384,
          "estimated_size": "1.3kb",
          "overhead": 1,
          "tripped": 0
        }
      },
      "script": {
        "compilations": 18,
        "cache_evictions": 0
      },
      "discovery": {
        "cluster_state_queue": {
          "total": 0,
          "pending": 0,
          "committed": 0
        },
        "published_cluster_states": {
          "full_states": 1,
          "incompatible_diffs": 0,
          "compatible_diffs": 24
        }
      },
      "ingest": {
        "total": {
          "count": 0,
          "time_in_millis": 0,
          "current": 0,
          "failed": 0
        },
        "pipelines": {
          "xpack_monitoring_2": {
            "count": 0,
            "time_in_millis": 0,
            "current": 0,
            "failed": 0
          },
          "xpack_monitoring_6": {
            "count": 0,
            "time_in_millis": 0,
            "current": 0,
            "failed": 0
          }
        }
      },
      "adaptive_selection": {
        "NGqSLsYuQ7uhfaQD5zdXpA": {
          "outgoing_searches": 0,
          "avg_queue_size": 0,
          "avg_service_time_ns": 227630,
          "avg_response_time_ns": 528265,
          "rank": "0.5"
        },
        "yeEZjv04SXqxNeIOKrj4gw": {
          "outgoing_searches": 0,
          "avg_queue_size": 0,
          "avg_service_time_ns": 398806,
          "avg_response_time_ns": 544477,
          "rank": "0.5"
        },
        "OGOD423-RZaMi9odaozE-A": {
          "outgoing_searches": 0,
          "avg_queue_size": 0,
          "avg_service_time_ns": 286288,
          "avg_response_time_ns": 526151,
          "rank": "0.5"
        }
      }
    }
  }
}
`

const nodeStatsResponseJVMProcess = `
{
  "_nodes": {
    "total": 1,
    "successful": 1,
    "failed": 0
  },
  "cluster_name": "es-testcluster",
  "nodes": {
    "SDFsfSDFsdfFSDSDfSFDSDF": {
      "timestamp": 1544736554065,
      "name": "test.host.com",
      "transport_address": "10.43.168.250:9300",
      "host": "test",
      "ip": "10.43.168.250:9300",
      "roles": [
        "ingest",
        "master"
      ],
      "attributes": {
        "ml.machine_memory": "21474836480",
        "ml.max_open_jobs": "20",
        "xpack.installed": "true",
        "ml.enabled": "true"
      },
      "process": {
        "timestamp": 1544727513051,
        "open_file_descriptors": 280,
        "max_file_descriptors": 1048576,
        "cpu": {
          "percent": 0,
          "total_in_millis": 6639430
        },
        "mem": {
          "total_virtual_in_bytes": 10407563264
        }
      },
      "jvm": {
        "timestamp": 1544727513051,
        "uptime_in_millis": 581892320,
        "mem": {
          "heap_used_in_bytes": 227332320,
          "heap_used_percent": 5,
          "heap_committed_in_bytes": 4286251008,
          "heap_max_in_bytes": 4286251008,
          "non_heap_used_in_bytes": 139787192,
          "non_heap_committed_in_bytes": 168996864,
          "pools": {
            "young": {
              "used_in_bytes": 26252696,
              "max_in_bytes": 69795840,
              "peak_used_in_bytes": 69795840,
              "peak_max_in_bytes": 69795840
            },
            "survivor": {
              "used_in_bytes": 3779488,
              "max_in_bytes": 8716288,
              "peak_used_in_bytes": 8716288,
              "peak_max_in_bytes": 8716288
            },
            "old": {
              "used_in_bytes": 197315272,
              "max_in_bytes": 4207738880,
              "peak_used_in_bytes": 197315272,
              "peak_max_in_bytes": 4207738880
            }
          }
        },
        "threads": {
          "count": 76,
          "peak_count": 76
        },
        "gc": {
          "collectors": {
            "young": {
              "collection_count": 42793,
              "collection_time_in_millis": 1340364
            },
            "old": {
              "collection_count": 3,
              "collection_time_in_millis": 140
            }
          }
        },
        "buffer_pools": {
          "mapped": {
            "count": 0,
            "used_in_bytes": 0,
            "total_capacity_in_bytes": 0
          },
          "direct": {
            "count": 17,
            "used_in_bytes": 34005104,
            "total_capacity_in_bytes": 34005103
          }
        },
        "classes": {
          "current_loaded_count": 16731,
          "total_loaded_count": 16733,
          "total_unloaded_count": 2
        }
      }
    }
  }
}
`

var nodestatsIndicesExpected = map[string]interface{}{
  "completion_size_in_bytes":      float64(0),
  "docs_count":      float64(0),
  "docs_deleted":      float64(0),
  "fielddata_evictions":      float64(0),
  "fielddata_memory_size_in_bytes":      float64(0),
  "flush_periodic":      float64(0),
  "flush_total":      float64(0),
  "flush_total_time_in_millis":      float64(0),
  "get_current":      float64(0),
  "get_exists_time_in_millis":      float64(0),
  "get_exists_total":      float64(0),
  "get_missing_time_in_millis":      float64(0),
  "get_missing_total":      float64(0),
  "get_time_in_millis":      float64(0),
  "get_total":      float64(0),
  "indexing_delete_current":      float64(0),
  "indexing_delete_time_in_millis":      float64(0),
  "indexing_delete_total":      float64(0),
  "indexing_index_current":      float64(0),
  "indexing_index_failed":      float64(0),
  "indexing_index_time_in_millis":      float64(0),
  "indexing_index_total":      float64(0),
  "indexing_noop_update_total":      float64(0),
  "indexing_throttle_time_in_millis":      float64(0),
  "merges_current":      float64(0),
  "merges_current_docs":      float64(0),
  "merges_current_size_in_bytes":      float64(0),
  "merges_total":      float64(0),
  "merges_total_auto_throttle_in_bytes":      float64(0),
  "merges_total_docs":      float64(0),
  "merges_total_size_in_bytes":      float64(0),
  "merges_total_stopped_time_in_millis":      float64(0),
  "merges_total_throttled_time_in_millis":      float64(0),
  "merges_total_time_in_millis":      float64(0),
  "query_cache_cache_count":      float64(0),
  "query_cache_cache_size":      float64(0),
  "query_cache_evictions":      float64(0),
  "query_cache_hit_count":      float64(0),
  "query_cache_memory_size_in_bytes":      float64(0),
  "query_cache_miss_count":      float64(0),
  "query_cache_total_count":      float64(0),
  "recovery_current_as_source":      float64(0),
  "recovery_current_as_target":      float64(0),
  "recovery_throttle_time_in_millis":      float64(0),
  "refresh_listeners":      float64(0),
  "refresh_total":      float64(0),
  "refresh_total_time_in_millis":      float64(0),
  "request_cache_evictions":      float64(0),
  "request_cache_hit_count":      float64(0),
  "request_cache_memory_size_in_bytes":      float64(0),
  "request_cache_miss_count":      float64(0),
  "search_fetch_current":      float64(0),
  "search_fetch_time_in_millis":      float64(0),
  "search_fetch_total":      float64(0),
  "search_open_contexts":      float64(0),
  "search_query_current":      float64(0),
  "search_query_time_in_millis":      float64(0),
  "search_query_total":      float64(0),
  "search_scroll_current":      float64(0),
  "search_scroll_time_in_millis":      float64(0),
  "search_scroll_total":      float64(0),
  "search_suggest_current":      float64(0),
  "search_suggest_time_in_millis":      float64(0),
  "search_suggest_total":      float64(0),
  "segments_count":      float64(0),
  "segments_doc_values_memory_in_bytes":      float64(0),
  "segments_fixed_bit_set_memory_in_bytes":      float64(0),
  "segments_index_writer_memory_in_bytes":      float64(0),
  "segments_max_unsafe_auto_id_timestamp":      float64(-9.223372036854776e+18),
  "segments_memory_in_bytes":      float64(0),
  "segments_norms_memory_in_bytes":      float64(0),
  "segments_points_memory_in_bytes":      float64(0),
  "segments_stored_fields_memory_in_bytes":      float64(0),
  "segments_term_vectors_memory_in_bytes":      float64(0),
  "segments_terms_memory_in_bytes":      float64(0),
  "segments_version_map_memory_in_bytes":      float64(0),
  "store_size_in_bytes":      float64(0),
  "translog_earliest_last_modified_age":      float64(0),
  "translog_operations":      float64(0),
  "translog_size_in_bytes":      float64(0),
  "translog_uncommitted_operations":      float64(0),
  "translog_uncommitted_size_in_bytes":      float64(0),
  "warmer_current":      float64(0),
  "warmer_total":      float64(0),
  "warmer_total_time_in_millis":      float64(0),
}

var nodestatsOsExpected = map[string]interface{}{
  "cgroup_cpu_cfs_period_micros":      float64(100000),
  "cgroup_cpu_cfs_quota_micros":      float64(100000),
  "cgroup_cpu_stat_number_of_elapsed_periods":      float64(2332636.0),
  "cgroup_cpu_stat_number_of_times_throttled":      float64(21112),
  "cgroup_cpu_stat_time_throttled_nanos":      float64(2087989489827.0),
  "cgroup_cpuacct_usage_nanos":      float64(6648413872901.0),
  "cpu_load_average_15m":      float64(0),
  "cpu_load_average_1m":      float64(0),
  "cpu_load_average_5m":      float64(0.01),
  "cpu_percent":      float64(0),
  "mem_free_in_bytes":      float64(56699928576.0),
  "mem_free_percent":      float64(85),
  "mem_total_in_bytes":      float64(66726227968.0),
  "mem_used_in_bytes":      float64(10026299392.0),
  "mem_used_percent":      float64(15),
  "swap_free_in_bytes":      float64(0),
  "swap_total_in_bytes":      float64(0),
  "swap_used_in_bytes":      float64(0),
  "timestamp":      float64(1544727513051.0),
}

var nodestatsProcessExpected = map[string]interface{}{
  "cpu_percent":      float64(0),
  "cpu_total_in_millis":      float64(6639430.0),
  "max_file_descriptors":      float64(1048576.0),
  "mem_total_virtual_in_bytes":      float64(10407563264.0),
  "open_file_descriptors":      float64(280),
  "timestamp":      float64(1544727513051.0),
}

var nodestatsJvmExpected = map[string]interface{}{
  "buffer_pools_direct_count":      float64(17),
  "buffer_pools_direct_total_capacity_in_bytes":      float64(34005103.0),
  "buffer_pools_direct_used_in_bytes":      float64(34005104.0),
  "buffer_pools_mapped_count":      float64(0),
  "buffer_pools_mapped_total_capacity_in_bytes":      float64(0),
  "buffer_pools_mapped_used_in_bytes":      float64(0),
  "classes_current_loaded_count":      float64(16731),
  "classes_total_loaded_count":      float64(16733),
  "classes_total_unloaded_count":      float64(2),
  "gc_collectors_old_collection_count":      float64(3),
  "gc_collectors_old_collection_time_in_millis":      float64(140),
  "gc_collectors_young_collection_count":      float64(42793),
  "gc_collectors_young_collection_time_in_millis":      float64(1340364.0),
  "mem_heap_committed_in_bytes":      float64(4286251008.0),
  "mem_heap_max_in_bytes":      float64(4286251008.0),
  "mem_heap_used_in_bytes":      float64(227332320.0),
  "mem_heap_used_percent":      float64(5),
  "mem_non_heap_committed_in_bytes":      float64(168996864.0),
  "mem_non_heap_used_in_bytes":      float64(139787192.0),
  "mem_pools_old_max_in_bytes":      float64(4207738880.0),
  "mem_pools_old_peak_max_in_bytes":      float64(4207738880.0),
  "mem_pools_old_peak_used_in_bytes":      float64(197315272.0),
  "mem_pools_old_used_in_bytes":      float64(197315272.0),
  "mem_pools_survivor_max_in_bytes":      float64(8716288.0),
  "mem_pools_survivor_peak_max_in_bytes":      float64(8716288.0),
  "mem_pools_survivor_peak_used_in_bytes":      float64(8716288.0),
  "mem_pools_survivor_used_in_bytes":      float64(3779488.0),
  "mem_pools_young_max_in_bytes":      float64(69795840.0),
  "mem_pools_young_peak_max_in_bytes":      float64(69795840.0),
  "mem_pools_young_peak_used_in_bytes":      float64(69795840.0),
  "mem_pools_young_used_in_bytes":      float64(26252696.0),
  "threads_count":      float64(76),
  "threads_peak_count":      float64(76),
  "timestamp":      float64(1544727513051.0),
  "uptime_in_millis":      float64(581892320.0),
}

var nodestatsThreadPoolExpected = map[string]interface{}{
  "analyze_active":      float64(0),
  "analyze_completed":      float64(0),
  "analyze_largest":      float64(0),
  "analyze_queue":      float64(0),
  "analyze_rejected":      float64(0),
  "analyze_threads":      float64(0),
  "fetch_shard_started_active":      float64(0),
  "fetch_shard_started_completed":      float64(0),
  "fetch_shard_started_largest":      float64(0),
  "fetch_shard_started_queue":      float64(0),
  "fetch_shard_started_rejected":      float64(0),
  "fetch_shard_started_threads":      float64(0),
  "fetch_shard_store_active":      float64(0),
  "fetch_shard_store_completed":      float64(0),
  "fetch_shard_store_largest":      float64(0),
  "fetch_shard_store_queue":      float64(0),
  "fetch_shard_store_rejected":      float64(0),
  "fetch_shard_store_threads":      float64(0),
  "flush_active":      float64(0),
  "flush_completed":      float64(0),
  "flush_largest":      float64(0),
  "flush_queue":      float64(0),
  "flush_rejected":      float64(0),
  "flush_threads":      float64(0),
  "force_merge_active":      float64(0),
  "force_merge_completed":      float64(0),
  "force_merge_largest":      float64(0),
  "force_merge_queue":      float64(0),
  "force_merge_rejected":      float64(0),
  "force_merge_threads":      float64(0),
  "generic_active":      float64(0),
  "generic_completed":      float64(1335639.0),
  "generic_largest":      float64(8),
  "generic_queue":      float64(0),
  "generic_rejected":      float64(0),
  "generic_threads":      float64(8),
  "get_active":      float64(0),
  "get_completed":      float64(0),
  "get_largest":      float64(0),
  "get_queue":      float64(0),
  "get_rejected":      float64(0),
  "get_threads":      float64(0),
  "index_active":      float64(0),
  "index_completed":      float64(0),
  "index_largest":      float64(0),
  "index_queue":      float64(0),
  "index_rejected":      float64(0),
  "index_threads":      float64(0),
  "listener_active":      float64(0),
  "listener_completed":      float64(0),
  "listener_largest":      float64(0),
  "listener_queue":      float64(0),
  "listener_rejected":      float64(0),
  "listener_threads":      float64(0),
  "management_active":      float64(1),
  "management_completed":      float64(622652),
  "management_largest":      float64(4),
  "management_queue":      float64(0),
  "management_rejected":      float64(0),
  "management_threads":      float64(4),
  "ml_autodetect_active":      float64(0),
  "ml_autodetect_completed":      float64(0),
  "ml_autodetect_largest":      float64(0),
  "ml_autodetect_queue":      float64(0),
  "ml_autodetect_rejected":      float64(0),
  "ml_autodetect_threads":      float64(0),
  "ml_datafeed_active":      float64(0),
  "ml_datafeed_completed":      float64(0),
  "ml_datafeed_largest":      float64(0),
  "ml_datafeed_queue":      float64(0),
  "ml_datafeed_rejected":      float64(0),
  "ml_datafeed_threads":      float64(0),
  "ml_utility_active":      float64(0),
  "ml_utility_completed":      float64(36),
  "ml_utility_largest":      float64(36),
  "ml_utility_queue":      float64(0),
  "ml_utility_rejected":      float64(0),
  "ml_utility_threads":      float64(36),
  "refresh_active":      float64(0),
  "refresh_completed":      float64(0),
  "refresh_largest":      float64(0),
  "refresh_queue":      float64(0),
  "refresh_rejected":      float64(0),
  "refresh_threads":      float64(0),
  "rollup_indexing_active":      float64(0),
  "rollup_indexing_completed":      float64(0),
  "rollup_indexing_largest":      float64(0),
  "rollup_indexing_queue":      float64(0),
  "rollup_indexing_rejected":      float64(0),
  "rollup_indexing_threads":      float64(0),
  "search_active":      float64(0),
  "search_completed":      float64(281218),
  "search_largest":      float64(2),
  "search_queue":      float64(0),
  "search_rejected":      float64(0),
  "search_threads":      float64(2),
  "security-token-key_active":      float64(0),
  "security-token-key_completed":      float64(0),
  "security-token-key_largest":      float64(0),
  "security-token-key_queue":      float64(0),
  "security-token-key_rejected":      float64(0),
  "security-token-key_threads":      float64(0),
  "snapshot_active":      float64(0),
  "snapshot_completed":      float64(0),
  "snapshot_largest":      float64(0),
  "snapshot_queue":      float64(0),
  "snapshot_rejected":      float64(0),
  "snapshot_threads":      float64(0),
  "warmer_active":      float64(0),
  "warmer_completed":      float64(0),
  "warmer_largest":      float64(0),
  "warmer_queue":      float64(0),
  "warmer_rejected":      float64(0),
  "warmer_threads":      float64(0),
  "watcher_active":      float64(0),
  "watcher_completed":      float64(0),
  "watcher_largest":      float64(0),
  "watcher_queue":      float64(0),
  "watcher_rejected":      float64(0),
  "watcher_threads":      float64(0),
  "write_active":      float64(0),
  "write_completed":      float64(65008),
  "write_largest":      float64(1),
  "write_queue":      float64(0),
  "write_rejected":      float64(0),
  "write_threads":      float64(1),
}

var nodestatsFsExpected = map[string]interface{}{
  "data_0_available_in_bytes":      float64(19592351744.0),
  "data_0_free_in_bytes":      float64(20682870784.0),
  "data_0_total_in_bytes":      float64(21003583488.0),
  "io_stats_devices_0_operations":      float64(222457),
  "io_stats_devices_0_read_kilobytes":      float64(156),
  "io_stats_devices_0_read_operations":      float64(25),
  "io_stats_devices_0_write_kilobytes":      float64(1618188.0),
  "io_stats_devices_0_write_operations":      float64(222432),
  "io_stats_total_operations":      float64(222457),
  "io_stats_total_read_kilobytes":      float64(156),
  "io_stats_total_read_operations":      float64(25),
  "io_stats_total_write_kilobytes":      float64(1618188.0),
  "io_stats_total_write_operations":      float64(222432),
  "timestamp":      float64(1544727513051.0),
  "total_available_in_bytes":      float64(19592351744.0),
  "total_free_in_bytes":      float64(20682870784.0),
  "total_total_in_bytes":      float64(21003583488.0),
}

var nodestatsTransportExpected = map[string]interface{}{
  "rx_count":      float64(6950414.0),
  "rx_size_in_bytes":      float64(21891347203.0),
  "server_open":      float64(46),
  "tx_count":      float64(6950411.0),
  "tx_size_in_bytes":      float64(17902587800.0),
}

var nodestatsHttpExpected = map[string]interface{}{
	"current_open": float64(0),
	"total_opened": float64(38),
}

var nodestatsBreakersExpected = map[string]interface{}{
  "accounting_estimated_size_in_bytes":      float64(0),
  "accounting_limit_size_in_bytes":      float64(4286251008.0),
  "accounting_overhead":      float64(1),
  "accounting_tripped":      float64(0),
  "fielddata_estimated_size_in_bytes":      float64(0),
  "fielddata_limit_size_in_bytes":      float64(2571750604.0),
  "fielddata_overhead":      float64(1.03),
  "fielddata_tripped":      float64(0),
  "in_flight_requests_estimated_size_in_bytes":      float64(1384),
  "in_flight_requests_limit_size_in_bytes":      float64(4286251008.0),
  "in_flight_requests_overhead":      float64(1),
  "in_flight_requests_tripped":      float64(0),
  "parent_estimated_size_in_bytes":      float64(1384),
  "parent_limit_size_in_bytes":      float64(3000375705.0),
  "parent_overhead":      float64(1),
  "parent_tripped":      float64(0),
  "request_estimated_size_in_bytes":      float64(0),
  "request_limit_size_in_bytes":      float64(2571750604.0),
  "request_overhead":      float64(1),
  "request_tripped":      float64(0),
}

const clusterStatsResponse = `
{
   "host":"ip-10-0-1-214",
   "log_type":"metrics",
   "timestamp":1475767451229,
   "log_level":"INFO",
   "node_name":"test.host.com",
   "cluster_name":"es-testcluster",
   "status":"red",
   "indices":{
      "count":1,
      "shards":{
         "total":4,
         "primaries":4,
         "replication":0.0,
         "index":{
            "shards":{
               "min":4,
               "max":4,
               "avg":4.0
            },
            "primaries":{
               "min":4,
               "max":4,
               "avg":4.0
            },
            "replication":{
               "min":0.0,
               "max":0.0,
               "avg":0.0
            }
         }
      },
      "docs":{
         "count":4,
         "deleted":0
      },
      "store":{
         "size_in_bytes":17084,
         "throttle_time_in_millis":0
      },
      "fielddata":{
         "memory_size_in_bytes":0,
         "evictions":0
      },
      "query_cache":{
         "memory_size_in_bytes":0,
         "total_count":0,
         "hit_count":0,
         "miss_count":0,
         "cache_size":0,
         "cache_count":0,
         "evictions":0
      },
      "completion":{
         "size_in_bytes":0
      },
      "segments":{
         "count":4,
         "memory_in_bytes":11828,
         "terms_memory_in_bytes":8932,
         "stored_fields_memory_in_bytes":1248,
         "term_vectors_memory_in_bytes":0,
         "norms_memory_in_bytes":1280,
         "doc_values_memory_in_bytes":368,
         "index_writer_memory_in_bytes":0,
         "index_writer_max_memory_in_bytes":2048000,
         "version_map_memory_in_bytes":0,
         "fixed_bit_set_memory_in_bytes":0
      },
      "percolate":{
         "total":0,
         "time_in_millis":0,
         "current":0,
         "memory_size_in_bytes":-1,
         "memory_size":"-1b",
         "queries":0
      }
   },
   "nodes":{
      "count":{
         "total":1,
         "master_only":0,
         "data_only":0,
         "master_data":1,
         "client":0
      },
      "versions":[
         {
         "version": "2.3.3"
         }
      ],
      "os":{
         "available_processors":1,
         "allocated_processors":1,
         "mem":{
            "total_in_bytes":593301504
         },
         "names":[
            {
               "name":"Linux",
               "count":1
            }
         ]
      },
      "process":{
         "cpu":{
            "percent":0
         },
         "open_file_descriptors":{
            "min":145,
            "max":145,
            "avg":145
         }
      },
      "jvm":{
         "max_uptime_in_millis":11580527,
         "versions":[
            {
               "version":"1.8.0_101",
               "vm_name":"OpenJDK 64-Bit Server VM",
               "vm_version":"25.101-b13",
               "vm_vendor":"Oracle Corporation",
               "count":1
            }
         ],
         "mem":{
            "heap_used_in_bytes":70550288,
            "heap_max_in_bytes":1065025536
         },
         "threads":30
      },
      "fs":{
         "total_in_bytes":8318783488,
         "free_in_bytes":6447439872,
         "available_in_bytes":6344785920
      },
      "plugins":[
         {
            "name":"cloud-aws",
            "version":"2.3.3",
            "description":"The Amazon Web Service (AWS) Cloud plugin allows to use AWS API for the unicast discovery mechanism and add S3 repositories.",
            "jvm":true,
            "classname":"org.elasticsearch.plugin.cloud.aws.CloudAwsPlugin",
            "isolated":true,
            "site":false
         },
         {
            "name":"kopf",
            "version":"2.0.1",
            "description":"kopf - simple web administration tool for Elasticsearch",
            "url":"/_plugin/kopf/",
            "jvm":false,
            "site":true
         },
         {
            "name":"tr-metrics",
            "version":"7bd5b4b",
            "description":"Logs cluster and node stats for performance monitoring.",
            "jvm":true,
            "classname":"com.trgr.elasticsearch.plugin.metrics.MetricsPlugin",
            "isolated":true,
            "site":false
         }
      ]
   }
}
`

var clusterstatsIndicesExpected = map[string]interface{}{
  "completion_size_in_bytes":      float64(0),
  "count":      float64(1),
  "docs_count":      float64(4),
  "docs_deleted":      float64(0),
  "fielddata_evictions":      float64(0),
  "fielddata_memory_size_in_bytes":      float64(0),
  "percolate_current":      float64(0),
  "percolate_memory_size":      "-1b",
  "percolate_memory_size_in_bytes":      float64(-1),
  "percolate_queries":      float64(0),
  "percolate_time_in_millis":      float64(0),
  "percolate_total":      float64(0),
  "query_cache_cache_count":      float64(0),
  "query_cache_cache_size":      float64(0),
  "query_cache_evictions":      float64(0),
  "query_cache_hit_count":      float64(0),
  "query_cache_memory_size_in_bytes":      float64(0),
  "query_cache_miss_count":      float64(0),
  "query_cache_total_count":      float64(0),
  "segments_count":      float64(4),
  "segments_doc_values_memory_in_bytes":      float64(368),
  "segments_fixed_bit_set_memory_in_bytes":      float64(0),
  "segments_index_writer_max_memory_in_bytes":      float64(2048000.0),
  "segments_index_writer_memory_in_bytes":      float64(0),
  "segments_memory_in_bytes":      float64(11828),
  "segments_norms_memory_in_bytes":      float64(1280),
  "segments_stored_fields_memory_in_bytes":      float64(1248),
  "segments_term_vectors_memory_in_bytes":      float64(0),
  "segments_terms_memory_in_bytes":      float64(8932),
  "segments_version_map_memory_in_bytes":      float64(0),
  "shards_index_primaries_avg":      float64(4),
  "shards_index_primaries_max":      float64(4),
  "shards_index_primaries_min":      float64(4),
  "shards_index_replication_avg":      float64(0),
  "shards_index_replication_max":      float64(0),
  "shards_index_replication_min":      float64(0),
  "shards_index_shards_avg":      float64(4),
  "shards_index_shards_max":      float64(4),
  "shards_index_shards_min":      float64(4),
  "shards_primaries":      float64(4),
  "shards_replication":      float64(0),
  "shards_total":      float64(4),
  "store_size_in_bytes":      float64(17084),
  "store_throttle_time_in_millis":      float64(0),
}

var clusterstatsNodesExpected = map[string]interface{}{
	"count_client":                      float64(0),
	"count_data_only":                   float64(0),
	"count_master_data":                 float64(1),
	"count_master_only":                 float64(0),
	"count_total":                       float64(1),
	"fs_available_in_bytes":             float64(6.34478592e+09),
	"fs_free_in_bytes":                  float64(6.447439872e+09),
	"fs_total_in_bytes":                 float64(8.318783488e+09),
	"jvm_max_uptime_in_millis":          float64(1.1580527e+07),
	"jvm_mem_heap_max_in_bytes":         float64(1.065025536e+09),
	"jvm_mem_heap_used_in_bytes":        float64(7.0550288e+07),
	"jvm_threads":                       float64(30),
	"jvm_versions_0_count":              float64(1),
	"jvm_versions_0_version":            "1.8.0_101",
	"jvm_versions_0_vm_name":            "OpenJDK 64-Bit Server VM",
	"jvm_versions_0_vm_vendor":          "Oracle Corporation",
	"jvm_versions_0_vm_version":         "25.101-b13",
	"os_allocated_processors":           float64(1),
	"os_available_processors":           float64(1),
	"os_mem_total_in_bytes":             float64(5.93301504e+08),
	"os_names_0_count":                  float64(1),
	"os_names_0_name":                   "Linux",
	"process_cpu_percent":               float64(0),
	"process_open_file_descriptors_avg": float64(145),
	"process_open_file_descriptors_max": float64(145),
	"process_open_file_descriptors_min": float64(145),
	"versions_0_version":                "2.3.3",
	"plugins_0_classname":               "org.elasticsearch.plugin.cloud.aws.CloudAwsPlugin",
	"plugins_0_description":             "The Amazon Web Service (AWS) Cloud plugin allows to use AWS API for the unicast discovery mechanism and add S3 repositories.",
	"plugins_0_isolated":                true,
	"plugins_0_jvm":                     true,
	"plugins_0_name":                    "cloud-aws",
	"plugins_0_site":                    false,
	"plugins_0_version":                 "2.3.3",
	"plugins_1_description":             "kopf - simple web administration tool for Elasticsearch",
	"plugins_1_jvm":                     false,
	"plugins_1_name":                    "kopf",
	"plugins_1_site":                    true,
	"plugins_1_url":                     "/_plugin/kopf/",
	"plugins_1_version":                 "2.0.1",
	"plugins_2_classname":               "com.trgr.elasticsearch.plugin.metrics.MetricsPlugin",
	"plugins_2_description":             "Logs cluster and node stats for performance monitoring.",
	"plugins_2_isolated":                true,
	"plugins_2_jvm":                     true,
	"plugins_2_name":                    "tr-metrics",
	"plugins_2_site":                    false,
	"plugins_2_version":                 "7bd5b4b",
}

const IsMasterResult = "SDFsfSDFsdfFSDSDfSFDSDF 10.206.124.66 10.206.124.66 test.host.com "

const IsNotMasterResult = "junk 10.206.124.66 10.206.124.66 test.junk.com "
