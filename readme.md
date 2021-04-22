- ssl erwähnen

TODO bis zur ersten abgabe:
- envoy kapitel ausführlicher schreiben

TODO chrstian:
- hardware lb mehr erklären -> performance ist nicht direkt schlechter
- unter osi layer = hardware
- layer 4 und höher -> softwarelösung
- consistent hashing -> warum?
- betrieb der LB und des clusters sollte transparent sein -> dies ist besonders schwer bei langlebigen tcp verbindungen
- client hat nen takeover -> client hat keinen takeover -> client merkt sobald die verbindung wegkommt
- problembeschreibung circuit breaking
- 4.5 verbindungsorientiertes
- discovery selber mechanismus -> selbe datenbasis
- stricted weighted round-robin
  - es geht nicht darum jede einzelne verbindung perfekt zu optimieren
  - es geht darum auf viele verschiedene verbindungen zu optimieren
  - future work -> es geht um die verteilung eine verbindung spezifische verbindung zu verteilen
- ausblick: verteilung findet auf warscheinlichkeiten statt
- ist envoy eine gute basis ? ist das vorgehen gut ? basis ist ausbaufähig -> man kann mehrere pakete parsen außer connect paket

- welche probleme löse ich ?
- smartness -> anpassungsfähigkeit des load balancers und seiner entscheidungen
- hivemq success stories
- die probleme besser beschreiben, beispiele?
- quellen für die probleme, warum hab ich die probleme?
- mqtt kennt keine cluster

TODO letzter check:
- zahlen 0 - 12 ausgeschrieben, runde zahlen auch ausschreiben
- fragezeichen quellen?

feedback neu:
- grafiken überarbeiten
- panel -> repeat option
- slow start -> ramp up
- hivemq kann das implementieren
- ist ein operatives ding -> kann metriken berücksichtigen

load balancing:
- wichtig! ist kern der arbeit

cluster discovery:
- distributed system hat das problem split brain
- splitbrain is nicht mein problem, das wird durch mqtt health check gelöst, muss der hivemq lösen

hivemq:
- industrial iot iiot
- mqtt hat viele konzepte die man rekombinieren kann

## wichtige formale fragen

- endet ein absatz wenn eine liste anfängt ? also muss ich bevor die liste anfängt den quellenverweis schreiben?
- Überschriften auf english oder deutsch?
- titel der arbeit, englisch oder deutsch?
- muss die caption am anfang groß geschrieben und mit einem punkt beendet werden?

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

ganz wichtig -> wir brauchen eine extension die clients abschießen kann
-> member einer shared subscription können immer sorgenfrei abgeschossen werden

rolling updates strategy, settings: search
neue latest version: alles geht auf

## fragen

open source mqtt stresser ?

# fragen an hivemq

- wie kann ich den cluster überlasten ? er schmeißt mir regelmäßig die publisher weg allerdings hat er trotzdem immer ca 30k available credits
- warum sind bytes read immer gleich ? auch bei 80 / 20 ?
- clients with backpressure immens auf großem node, trotzdem schlägt die overload protection auf dem kleinen an

# guidelines

guidelines: https://ilias.th-koeln.de/goto.php?target=file_1364683_download&client_id=ILIAS_FH_Koeln

# links

hivemq:

hivemq extensions: https://www.hivemq.com/docs/hivemq/4.4/extensions/registries.html#modifiable-client-settings
client takeover: https://www.hivemq.com/blog/mqtt-essentials-part-10-alive-client-take-over/
resilience: https://www.hivemq.com/blog/are-your-mqtt-applications-resilient-enough/
mqtt stresser: https://github.com/inovex/mqtt-stresser
custom tlv fields in proxy protocol: https://www.hivemq.com/docs/hivemq/4.4/user-guide/proxy-protocol.html#custom-tlv

lwt and client takeover issue: https://github.com/eclipse/mosquitto/issues/904

envoy:

draining: https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/operations/draining
life of a request: https://www.envoyproxy.io/docs/envoy/latest/intro/life_of_a_request
edge proxy: https://www.envoyproxy.io/docs/envoy/latest/configuration/best_practices/edge
tcp proxy and link to reference: https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/listeners/tcp_proxy
timeouts: https://www.envoyproxy.io/docs/envoy/latest/faq/configuration/timeouts
hello world example: https://salmaan-rashid.medium.com/envoy-control-plane-hello-world-2f49b2865f29
exmaple with python: https://salmaan-rashid.medium.com/envoy-discovery-hello-world-d3e44d19603d
envoy discovery repo: https://github.com/salrashid123/envoy_discovery
go control plane: https://github.com/envoyproxy/go-control-plane
microservices resilience: https://www.getambassador.io/resources/getting-started-envoyproxy-microservices-resilience/
matt klein talk: https://www.youtube.com/watch?v=gQF23Vw0keg
per listener overhead too high: https://github.com/envoyproxy/envoy/issues/4196
istio performance: https://istio.io/latest/docs/ops/deployment/performance-and-scalability/

podcasts:
podcast service mesh or a platform: https://www.weave.works/blog/is-envoy-a-service-mesh-or-a-platform
envoy: https://kubernetespodcast.com/episode/033-envoy/
matt klein 1: https://developertea.simplecast.com/episodes/50464d4b-4071c0b8
matt klein 2: https://developertea.simplecast.com/episodes/10bb00cd-a148f095

papers:

LB survey: https://dl.acm.org/doi/10.1145/3281010

inovex:

confluence: https://confluence.inovex.de/display/ITKB/%28WIP%29+Smart+load+balancing+von+MQTT+traffic
jira: https://jira.inovex.de/browse/ST-2828
