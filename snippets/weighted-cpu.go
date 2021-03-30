// init data container
processed := make(map[string]map[string]float64)
for _, broker := range endpoints {
  processed[broker.ID] = make(map[string]float64)
}

// process average values over all stored metrics
for _, broker := range endpoints {
  for metric, data := range broker.Metrics {
    mean := 0.0
    for _, d := range data {
      mean += d.Value
    }

    processed[broker.ID][metric] = mean / float64(len(data))
  }
}

for _, broker := range endpoints {
  // calculate weight based on free cpu capacity
  val := 100.0 - processed[broker.ID][Metrics.CpuLoad]

  // check for NaN
  if val != val || val < 1.0 {
    val = 1.0
  }

  // set the actual weight to the node
  broker.Weight = uint32(math.Round(val))
  log.Printf(
    "Broker: %v / free CPU: %v / Weight: %v / Healthy: %v",
    broker.ID, val, broker.Weight, broker.Healthy,
  )
}

// update snapshot
