# TO-DO file

### Global TO-DO

- [ ] **Arsenal** - Tools marketplace
- [ ] **One-liners** - One-liners marketplace/run
- [ ] **Forge** - Forge of commands/tools
- [ ] **Cyberchief** - Cyberchief like tool (Other name to find)
- [ ] **Report** - Report tool
- [ ] **CVE/CWE** - Common Vulnerabilities and Exposures/Common Weakness Enumeration
- [ ] **OSINT** - Open Source Intelligence

### Commandes de base pour le projet CLI HecateArsenal

hecatearsenal
ha-cli : Command Line Interface
ha-gui : Graphical User Interface
ha-tui : Text User Interface

ha-cli ol use <one-liner_name>
ha-cli ol list
ha-cli ol search <one-liner_name>
ha-cli ol show <one-liner_name>
ha-cli tool show <tool_name>
ha-cli tool list
ha-cli tool search <tool_name>
ha-cli tool use <tool_name>
ha-cli tool install <tool_name>

(ol = one-liner)

### Fonctionnalités de base pour le projet GUI HecateArsenal

- Interface graphique pour les one-liners
- Interface graphique pour les outils (marketplace, installation, configuration, exécution avec forge de commandes, suite de commandes/d'outils)
- Interface graphique pour les rapports (exportation, partage, stockage, visualisation)
- Interface graphique pour brute-forcing (dictionnaires, attaques par force brute, attaques par dictionnaire)
- Interface graphique pour la gestion des sessions (stockage, partage, visualisation, exportation)


### Base de données

- Choix de la base de données? : Neo4j
- Pourquoi ? : Base de données orientée graphe, permettant de modéliser les données de manière plus naturelle, beaucoup d'interconnexions entre les données, permet de faire des requêtes plus complexes, plus rapides et plus efficaces.

- Intégration de la base de données Neo4j
- Création des collections
- Création des méthodes de CRUD
- Création des méthodes de recherche
- Création des méthodes de filtres
- Création des méthodes de tri
- Traduction objet / document en base de données
- Librairie base de données : Neo4j

https://dev.to/makepad/creating-a-simple-go-application-for-crud-operations-with-neo4j-and-docker-9nh
https://neo4j.com/docs/go-manual/current/
https://pkg.go.dev/github.com/neo4j/neo4j-go-driver/v5

### GUI Project ?

- Wails : https://wails.io/
- Environnement de Developpement avec image Docker pour build le projet Wails / Go / VueJS

### TUI Projet ?
- https://github.com/charmbracelet/bubbletea

### Maybe interesting for the project ?

- gomplate : https://github.com/hairyhenderson/gomplate
- 