## structure

einleitung:
hivemq
mqtt

load balancer
envoy
- envoy hat einen weighted round robin aber nur basierend auf connections -> muss modular sein
- maybe this ? https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/endpoint/v3/endpoint_components.proto#envoy-v3-api-field-config-endpoint-v3-lbendpoint-load-balancing-weight
istio
- maybe not istio?

load shedding
circuit breaking

optional:
multiplexing

## idea

## what is smart?

- cluster discovery that works with hivemq cluster discovery
- sticky session via clientid
  - weil clients auch oft ihre ip ändern -> client id meist persistent
- weighted round robin (based on client credits for example)
  - ist dies besser als cpu round robin?
  - auf jeden fall besser als least connection -> wegen fetten downstream clients
- least connection on new broker (on new broker... maybe handle this with weight?)
  - how to handle rolling updates with a smart LB ?
  - this should solve the problem!
- shared subscriptions
- circuit breaking

- ist envoy grundlage oder teil der arbeit das wir uns für envoy entscheiden ?
- wasm network filter für client id




- hivemq testcontainer junit

- weighted round robin + sticky session based on client id + shared round robin based on sni
  - what if the broker do not want the client id
- LB as a huge state machine
  - different LB techniques in different states
  - state balanced: sticky sessions with clientid parsed (with lua filter)
  - state new broker: weighted round robin
- dynamic dns (with hivemq dynamic dns)

use wasm network filter:
https://www.envoyproxy.io/docs/envoy/latest/configuration/listeners/network_filters/wasm_filter

!! lua only works as http filter
lua filter https://github.com/envoyproxy/envoy/blob/main/examples/lua/envoy.yaml
https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/lua_filter#config-http-filters-lua

## schlachtplan

- k8s cluster mit hivemq aufsetzen
- istio rumspielen / circuit breaking etc
- hivemq overload -> rate limiting -> scaling cluster -> more connections

TODO:
- dominik fragen nach test suit

# what is smart

## istio

- multi tenant
  - unterscheidung mit zertifikaten based on sni?
  - unterscheidung mit ?
- limiting clusters
  - client limit
  - payload limit / funktioniert nicht out of the box
- hot reload hivemq nodes / auto discovery
- changing client tls certificates via api
  - bedeutet tcp verbindungen reißen ab?
- circuit breaking
  - apply client limit based on what ?
  - problem: wenn die load hoch ist, und man ein client limit setzt, dann sind die bisherigen clients immernoch verbunden! also bringt circuit breaking gar nichts?
- load shedding
  - tcp backpressure ??
  - error to client -> client has to resend
  - wie kann das load shedding mit der internen overload protection harmonieren ?
  - metriken scrapen -> siehe bmw projekt
- autoscaling
  - overload -> neuer hivemq node oder cpu / ram anpassen -> clients loslassen
  - wie kann ich die load denn stoppen? die tcp connections sind ja alle aktiv. ich müsste also künstlich tcp backpressure erzeugen

typische traffic muster:
- 'star' clients downstream
  - shared subscriptions
- billige iot publisher
- teure publisher
  - publish to ALL clients

problematische mqtt clients / hotspots:
- shared clients auf dem selben node
- teure publisher die an alle clients publishen

how to solve shared subscriptions:
- invers maglev
  - do NOT put clients with same 'identifier' on the same node
  - 'identifier' ist zb clients die zu einer shared group gehören
- gracefully teardown clients from hivemq
  - route them to other node on reconnect aka rebalancing

was ist das skalierungsproblem ??
  - shared subscription rebalancing
  - gracefully teardown

no more clients

maglev invers
read slaves datenbank, maximale verteilung

mqtt 5 redirect

### was ist nicht smart

- smartes verteilen auf einen hivemq node
  - je nach topic hat ein client intern sowieso viele verbindungen zu anderen nodes
- hivemq hat bereits kluge overload protection
  - diese muss nicht neu erfunden werden
  - basiert auf client credits
  - löst es mit hilfe von tcp backpressure
  - basiert auf ausstehende tasks

### probleme

- istio ist 'statisch' konfiguriert
  - muss man einen k8s controller schreiben der dann immer die werte im etcd ändert?
- was ist der kern der arbeit?
- istio kann man nicht einfach so erweitern mit quellcode zb custom health checks
  - eigener controller?
- wie soll ich die load aufhalten vor dem skalieren? ich hab ja bereits alle tcp connections

### roadmap

- k8s mit istio aufsetzen
- circuit breaking einrichten
- auto discovery
- custom k8s controller programmieren ?

## multiplexing

- ich bin in der bisherigen projektskizze nur auf multiplexing eingegangen
- viele hivemq instanzen auf einem worker -> viele tcp sockets
- mehrere millionen clients pro node
- man braucht mehrere listener -> max port limit -> da LB zwischen client / broker
- tcp memory allocation / buffer bloat am LB und broker
- es ist ja immernoch load balancing. bzw man kann den titel noch anpacken
- QUIC ist nicht mein hauptaugenmerk sondern
  - memory usage
  - warum ist multiplexing gerade bei mqtt gut ? (idle clients)

- https://github.com/envoyproxy/envoy/issues/4196
  - 90 GB per 1mio tcp connections
  -> 90 GB ram sparen durch double buffer bloat

## problem

- MVP zeigt: weniger tcp connections am broker = absolut keine RAM reduktion

## other

- connection limit per client id
- tls handshakes
- mqtt tracing

# alte multiplexing projektplanung

Konkrete Ziele:
- TCP - gRPC Bridge in einer beliebigen Programmiersprache programmieren
- gRPC Server im HiveMQ implementieren
- Dummy MQTT Client über die Bridge mit dem Broker verbinden
- TCP - gRPC Bridge in Envoy mit WASM implementieren
- Envoy Control Plane programmieren welche alle HiveMQ Nodes discovered

Risiken:
- Es ist zu komplex den gRPC Server in den HiveMQ Broker zu implementieren.
In diesem Fall kann zu der TCP - gRPC Bridge eine TCP - gRPC Reverse Bridge programmiert werden. Diese läuft dann als Sidecar neben dem HiveMQ um aus der gRPC Verbindung einzelne TCP Verbindungen zum HiveMQ zu machen. Dies ist natürlich keineswegs für die Produktion geeignet, da man mit dieser Lösung nichts gewinnt, aber ich könnte im Rahmen der Bachelorarbeit die TCP - gRPC Bridge Ende zu Ende testen. Der gRPC Server kann dann im Anschluss an meine Bachelorarbeit in den HiveMQ Broker implementiert werden.
- HTTP/2 Streams sind nicht performanter als TCP Sockets oder haben ein hard Limit
Falls die HTTP/2 Streams an ihre Grenzen stoßen, gibt es die Möglichkeit diese zu Poolen. Dabei werden mehrere Clients über einen Stream abgebildet und man müsste pro Paket einen eindeutigen Identifier mitschicken, zu welchem Client dieses gehört.
- Es ist nicht möglich einen solchen Filter in Envoy zu implementieren.
Envoy bietet erst seit kurzem (November 2020) die Möglichkeit einen Filter als WASM Modul einzubinden. Dementsprechend ist die Schnittstelle noch nicht 100% ausgereift und dokumentiert. Da Envoy Open Source ist, kann ein solcher Filter nach wie vor nativ implementiert werden. Falls dies zu aufwendig ist, kann für den Use - Case auch die TCP - gRPC Bridge aus Schritt 1 erweitert und optimiert werden.

# TODO Multiplexing

https://github.com/tidwall/evio
implement with evio

# meeting marcel

## wie wird das cluster betrieben

6 pods im k8s (interne kommunikation)
3 ecs plain docker (für fahrzeuge)

alles gleich groß

## monitoring von node auslastung

gemittelte 5 min
current connections

backend 4k clients, shared subscriptions

- cpu
- network traffic
- average credits accross all clients
- overload protection

cluster join ist teuer
credits 
overload protection level ab lvl 5

cluster rollover / cluser join ist problematisch
szenario:
- zuerst
- readiness 

node dazu / node weg cluster join sehr teuer, dauert lange

marcel: wir können mal gerne nen node abschießen und gucken was passiert

# tweekly 5.2

clustergröße
screenshots mqtt stresser
indikatoren, wie lange average bilden ?

meeting mit marcel besprechen
vpn ??

einfach mehr shared runner ?? 4k auf 3 clients -> keine probleme mehr
-> es gibt keine teuren clients mehr

warum? hohe packungsdichte erzielen -> keine ressourcen verschwenden indem man broker überdimensioniert

cluster join -> wie kann man da helfen ?

ganz wichtig -> wir brauchen eine extension die clients abschießen kann
-> member einer shared subscription können immer sorgenfrei abgeschossen werden

rolling updates strategy, settings: search
neue latest version: alles geht auf

## fragen

open source mqtt stresser ?
persistent connection
qos level

was sind problemclients

terraform domain + ip , dns auflösung in go control plane
https://www.hivemq.com/extension/dns-discovery-extension/

vpn
https://gitlab.inovex.de/inovex-cloud/terraform-module-openstack-openvpn/-/tree/master/

# fragen an hivemq

- wie kann ich den cluster überlasten ? er schmeißt mir regelmäßig die publisher weg allerdings hat er trotzdem immer ca 30k available credits
- warum sind bytes read immer gleich ? auch bei 80 / 20 ?
- clients with backpressure immens auf großem node, trotzdem schlägt die overload protection auf dem kleinen an
