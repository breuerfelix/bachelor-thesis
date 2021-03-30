func discoverCluster(cluster *Cluster) {
  for {
    ips, _ := net.LookupIP(cluster.Domain)
    newCluster := make([]*HiveMQ, 0)

    for _, ip := range ips {
      found := false
      for _, hivemq := range cluster.Nodes {
        if ip.String() == hivemq.HostName {
          newCluster = append(newCluster, hivemq)
          found = true
          break
        }
      }

      if found {
        continue
      }

      fmt.Println("Discovered New Node:", ip.String())
      hivemq := NewHiveMQ(ip.String(), 1883, 9399)
      newCluster = append(newCluster, hivemq)
    }

    cluster.Nodes = newCluster
    time.Sleep(2 * time.Second)
  }
}
