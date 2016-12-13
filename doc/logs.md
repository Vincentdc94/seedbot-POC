# Botnet research

Botnet heeft typisch twee elementen in zijn netwerk. Een CNC server om de commando's te sturen naar de bot en bots die de commando's uitvoeren.
Een andere variant van een botnet is de P2P variant die commando's onderling sturen zonder centrale CNC server. 

## We hebben torrents nodig!

*8/12/2016*

Omdat we natuurlijk torrents nodig hebben om bij te seeden begin ik nu met het maken van een cnc server die torrents haalt van de yify JSON api.
Met als plan om de torrents onder 10 seeders hun hashes door te sturen naar de bots om die torrents te downloaden en achteraf te seeden.

*12/12/2016*

Op naar de torrent downloaden om later gewoon te laten seeden. Ik ben op zoek naar een goede torrent library voor golang.
Op het moment ga ik voor *"libtorrent"* (github.com/axet/libtorrent). Momenteel ga ik proberen om adhv van de hashes die ik kan verkrijgen via de Yify API. 
Magnet links opbouwen om dan de torrents te downloaden en hopen dat libtorrent automatisch seed als ik een torrent start via de startTorrent() functie.

### libtorrent en windows

Het ziet er naar uit dat libtorrent zich niet zo graag op windows bevindt. Het lijkt me tijd om eens libtorrent-go te proberen.
Deze heeft dan ook weer buiten gcc, pkg-config nodig. 

*13/12/2016*

Ben uiteindelijk uitgekomen op *"github.com/anacrolix/torrent"* en bekijk gebruik inspiratie van https://github.com/Sioro-Neoku/go-peerflix om mijn eigen simpele torrentclient te maken voor de bots.

Ok even later alles werkt all basis ik krijg hashes doorgestuurd via de cnc server en de bots bouwen hiermee magnet links op waarvan ze gewoon de torrents downloaden al. 

**Paar problemen wel**

1) Download 20 - 30 torrents tegelijk en eet resources weg.
2) Nog niet zeker of het er na gaat seeden en hoe ik kan reseeden als de computer opnieuw start
3) Veel optimalizaties zijn nodig om gebied van juiste torrents te halen en downloads te verdelen en eventueel zelfs torrents toewijzen aan maar een selecte groep bots omdat er een limiet is aan hoeveel er nodig zijn per torrent. 
Ook is het probleem dat als alle bots al 20 torrents gaan opslaan dat de computers snel volgeraken misschien een configuratie maken waarmee dat kan aangepast kan worden.

