# Jolokia2 Input Plugin

The [Jolokia](http://jolokia.org) input plugin collects JVM metrics exposed as JMX MBean attributes
through the Jolokia REST endpoint and its [JSON-over-HTTP protocol](https://jolokia.org/reference/html/protocol.html).

### Configuration:

```toml
# Read JMX metrics from a Jolokia REST endpoint

[[inputs.jolokia2]]
  # default_field_separator = "."
  # default_field_prefix    = ""
  # default_tag_separator   = "_"
  # default_tag_prefix      = ""

  # Add agents to query
  [inputs.jolokia2.agents]
    urls = ["http://kafka:8080/jolokia"]
    # tag_url  = true
    # tag_host = false
    # tag_addr = false

  # Supply a 'path' to collect a simple scalar value called 'Uptime'.
  [[inputs.jolokia2.metric]]
    name  = "jvm_runtime"
    mbean = "java.lang:type=Runtime"
    paths = ["Uptime"]

  # More complex values may be collected as well.
  [[inputs.jolokia2.metric]]
    name  = "jvm_memory"
    mbean = "java.lang:type=Memory"
    paths = ["HeapMemoryUsage", "NonHeapMemoryUsage", "ObjectPendingFinalizationCount"]

  # Use 'mbean' object patterns to create distinct series, and
  # use 'tag_keys' to specify exactly the keys to add as tags.
  [[inputs.jolokia2.metric]]
    name     = "jvm_garbage_collector"
    mbean    = "java.lang:name=*,type=GarbageCollector"
    paths    = ["CollectionTime", "CollectionCount"]
    tag_keys = ["name"]

  # Use 'tag_prefix' to add detail to tag names.
  [[inputs.jolokia2.metric]]
    name       = "jvm_memory_pool"
    mbean      = "java.lang:name=*,type=MemoryPool"
    paths      = ["Usage", "PeakUsage", "CollectionUsage"]
    tag_keys   = ["name"]
    tag_prefix = "pool"

  # Use simple substitutions to alter field prefixes with mbean properties.
  [[inputs.jolokia2.metric]]
    name         = "kafka_topic"
    mbean        = "kafka.server:name=*,topic=*,type=BrokerTopicMetrics"
    field_prefix = "$1"
    tag_keys     = ["topic"]

  # In some cases, it makes sense to substitute entire field names for
  # mbean property keys.
  [[inputs.jolokia2.metric]]
    name       = "kafka_log"
    mbean      = "kafka.log:name=*,partition=*,topic=*,type=Log"
    field_name = "$1"
    tag_keys = ["topic", "partition"]
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
