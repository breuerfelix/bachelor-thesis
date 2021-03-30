// calculations based on metrics
for _, broker := range endpoints {
  // calculate health status based on overload protection
  metrics := broker.Metrics[Metrics.OverloadProtection]
  if len(metrics) < 1 {
    continue
  }

  // get the latest value
  overload := metrics[len(metrics)-1]
  // if value is <= 5, set to healthy
  broker.OverloadProtectionHealthy = overload.Value <= 5.0
}

// determine health status of nodes
for _, broker := range endpoints {
  // only set to healthy if all healh conditions are met
  if broker.MqttHealthy &&
    broker.OverloadProtectionHealthy {
    broker.Healthy = true
    continue
  }

  broker.Healthy = false
}

// update snapshot
