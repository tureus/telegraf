# Jolokia2 Input Plugin

The [Jolokia](http://jolokia.org) input plugin collects JVM metrics exposed as JMX MBean attributes
through the Jolokia REST endpoint and its [JSON-over-HTTP protocol](https://jolokia.org/reference/html/protocol.html).

### Configuration:

```toml

# Read JMX metrics through Jolokia

[[inputs.jolokia2]]
  #default_field_delimiter = "."
  #default_field_prefix    = ""
  #default_tag_delimiter   = "_"
  #default_tag_prefix      = "mbean"

  # Add agents to query
  [inputs.jolokia2.agents]
    urls = ["http://kafka:8080/jolokia"]

  [[inputs.jolokia2.metric]]
    name  = "jvm_runtime"
    mbean = "java.lang:type=Runtime"
    paths = ["Uptime"]

  [[inputs.jolokia2.metric]]
    name  = "jvm_memory"
    mbean = "java.lang:type=Memory"
    paths = ["HeapMemoryUsage", "NonHeapMemoryUsage", "ObjectPendingFinalizationCount"]

  # By default, all mbean keys are added as tags
  # Use 'taginclude' to specify the exact tags to add.
  [[inputs.jolokia2.metric]]
    name  = "jvm_g1_garbage_collector"
    mbean = "java.lang:name=G1*,type=GarbageCollector"
    paths = [
      "CollectionTime",
      "CollectionCount",
      "LastGcInfo/duration",
      "LastGcInfo/GcThreadCount",
    ]
    taginclude = ["name"]

  # Use 'tagexclude' to specify just the tags to remove.
  [[inputs.jolokia2.metric]]
    name       = "jvm_memory_pool"
    mbean      = "java.lang:name=*,type=MemoryPool"
    paths      = ["Usage", "PeakUsage, "CollectionUsage"]
    tagexclude = ["type"]

  [[inputs.jolokia2.metric]]
    name         = "kafka_topic"
    mbean        = "kafka.server:name=*,topic=*,type=BrokerTopicMetrics"
    field_prefix = "$1"
    taginclude   = ["topic"]

  [[inputs.jolokia2.metric]]
    name       = "kafka_log"
    mbean      = "kafka.log:name=*,partition=*,topic=*,type=Log"
    field_name = "$1"
    taginclude = ["topic", "partition"]
```

To specify timeouts for slower/over-loaded clients:

```
[[inputs.jolokia2]]
  [inputs.jolokia2.agents]
    urls = ["http://kafka:8080/jolokia"]

    # The amount of time to wait for any requests made by this client.
    # Includes connection time, any redirects, and reading the response body.
    # (default is 5s)
    response_timeout = "10s"
```

To specify SSL options, add details to the `agents` configuration:

```
[[inputs.jolokia2]]
  [inputs.jolokia2.agents]
    urls = [
      "https://kafka:8080/jolokia",
    ]
    #username = ""
    #password = ""
    ssl_ca   = "/var/private/ca.pem"
    ssl_cert = "/var/private/client.pem"
    ssl_key  = "/var/private/client-key.pem"
    #insecure_skip_verify = false
```

To interact with agents via a Jolokia proxy, use a `proxy` configuration instead:

```
[[inputs.jolokia2]]
  [inputs.jolokia2.proxy]
    url = "https://proxy:8080/jolokia"
    response_timeout = "10s"
    #default_target_username = ""
    #default_target_password = ""
    ssl_ca   = "/var/private/ca.pem"
    ssl_cert = "/var/private/client.pem"
    ssl_key  = "/var/private/client-key.pem"

    [[inputs.jolokia2.proxy.target]]
      url = "service:jmx:rmi:///jndi/rmi://targethost:9999/jmxrmi"
      #username = ""
      #password = ""
```
