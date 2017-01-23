# Seedbot - POC

## Documentatie 
Je kan mijn logs en powerpoint vinden in de /doc folder

## Ik wil je source bekijken.

Om de source van de bot te bekijken ga naar: src/bot/src/github.com/vincentdc94

Om de source van de cnc server te bekijken ga naar het zelfde pad maar verwissel alleen 'bot' met 'cnc'

## Het programma draaien.
In de bin folder zijn statisch gecompileerde '.exe' bestanden om het botnet uit te voeren. Deze versie werkt via localhost.

Als je liever vanuit de source code vertrekt. Installeer De Go toolchain. En doe de volgende zaken voor de bot en de cnc server.

**Beiden**

Ga naar 'src/bot/src/' of 'src/bot/src/' en voer via cmd setpath.bat uit zodat je go omgeving nu de directories van mijn applicatie zijn.
Ik heb per project een aparte omgeving zodat ik later geen problemen ga hebben met overlappende dependencies onder go projecten waarbij functionaliteit breek als ik andere versies heb.

**Bot**

Ga naar de src/bot/src/github.com/vincentdc94/ folder en run het commando

```bash
go run bot.go
```

**CNC Server**

Ga naar src/cnc/src/github.com/vincentdc94/ en run het commando

```bash
go run cnc.go
```

