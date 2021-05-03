Grafiken:
14:20 - 14:40 = szenario 1, least conn
15:40 - 16:10 = szenario 1, round robin
16:20 - 16:40 = szenario 1, weighted cpu
19:22 - 19:30 = szenario 2, least conn
19:32 - 19:40 = szenario 2, round robin
11:47 - 11:59 = szenario 2, weighted cpu
19:45 - 20:20 = szenario 3, round robin
12:55 - = szenario 3, least conn
12:00 - 12:40 = szenario 3, weighted cpu

- komprimierbare oder erschöpfbare ressourcen sind egal beim regeln -> ich kann ja genauso die erschöpfbaren ressourcen regeln
-> ist wichtig beim circuit breaking
- hivemq macht in ihren testszenarien oft: alle clients verbinden -> testen ... weighted = sinnlos

TODO:
- abstract schreiben
- korrekturen einarbeiten
- bilder an die richtigen stellen
- quellen überprüfen / seitenangabe bei quellen
- überlastschutz kapitel
- envoy architecture
- split brain erwähnen bei cluster discovery

überlastschutz bearbeiten:
- ist circuit breaking das richtige wort ?
- circuit breaking -> ich muss vorher schon regeln sonst schädige ich die verbundenen clients
- begriff circuit breaking nennen
- ist health status wirklich das was ich bei circuit breaking machen will ?
- welche ressourcen kann man sich angucken?

- prototypische
- ausblick dann fazit
- fazit nurnoch persönliche worte
- klug im sinne von integriert
- eher was die arbeit festgestellt hat

ausblick:
- andeuten was es noch so gut
- server redirect funktion -> mqtt protokoll entwicklung
- user properties -> kann man für load balancing benutzen

fazit:
- nicht zu technisch -> mehr beobachtung der branche
- mqtt ist ein wenig verschlossen -> gegenüber http
- es gibt keine erweiterungn für load balancer für mqtt
- es ist am anfang das man für dieses protokoll die features versteht
- iot welt oft groß kommerziell -> nicht einzusehen

- kluger lb -> gute leistung -> kluge metriken aus der applikation

- lb hat alle nodes im überblick

fazit: halbe seite

abstract:
- iot, mqtt
- skalierbar -> hivemq cluster
- neue herausforderungen
- envoy bietet große chancen, programmierbare lb wie envoy

debriefing mit hivemq

## letzter check
- zahlen 0 - 12 ausgeschrieben, runde zahlen auch ausschreiben
- quellen überarbeiten -> firmen, online magazine
- fragezeichen quellen?

## wichtige formale fragen

- Überschriften auf english oder deutsch?
- titel der arbeit, englisch oder deutsch?

## fragen an inovex

- open source mqtt stresser ?
- open source control plane ?

# guidelines

guidelines: https://ilias.th-koeln.de/goto.php?target=file_1364683_download&client_id=ILIAS_FH_Koeln

# links

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
