# Projet : Mac Monitoring en Go (Prometheus & Grafana)

## À propos de ce projet
L'objectif de ce projet est d'explorer le monde du monitoring de bout en bout. Plutôt que d'utiliser des solutions "boîtes noires", l'idée était de comprendre comment les métriques sont générées, collectées et affichées en partant de zéro. 

Ce projet utilise :
- **Go (gopsutil)** : Pour extraire directement les métriques matérielles (Charge CPU, Pression RAM) du système.
- **Prometheus** : Pour aspirer et stocker ces données temporelles.
- **Grafana** : Pour la visualisation et la création de tableaux de bord.
- **Docker Compose** : Pour orchestrer toute cette architecture proprement.

## Les Métriques Exposées
Le script Go personnalise et expose actuellement :
- `mac_cpu_charge` : Le nombre de processus en cours d'exécution par rapport au nombre de cœurs.
- `mac_ram_surcharge` : Le pourcentage de pression de la mémoire vive.

## Prérequis
- Docker et Docker Compose installés sur la machine.

## Lancement du projet

1. Placez-vous à la racine du projet (là où se trouve le fichier `docker-compose.yml`).
2. Lancez la construction et le démarrage des conteneurs en arrière-plan :
```bash
   docker compose up -d --build
```
3. Pour vérifier que tout fonctionne :
```bash
docker compose ps
```

## Accès aux services
- **Conteneur Go** (Métriques brutes) : http://localhost:1221/metrics
- **Serveur Prometheus** : http://localhost:9090
- **Interface Grafana** : http://localhost:3000

## Guide Rapide : Créer son premier graphique sur Grafana

Une fois les conteneurs lancés, voici comment visualiser les données :

1. Ouvrez http://localhost:3000.
2. Connectez-vous avec les identifiants par défaut (admin / admin).
3. Allez dans Connections > Data Sources > Add data source > Prometheus.
4. Dans l'URL, entrez le nom du service Docker : http://prometheus:9090 puis cliquez sur Save & test.
5. Allez dans Dashboards > New dashboard > Add visualization.
6. Sélectionnez votre source Prometheus.
7. Dans le champ de requête (PromQL), tapez l'une des métriques, par exemple : mac_ram_surcharge
8. En haut à droite, changez le type de graphique (passez de "Time series" à "Gauge" pour avoir un compteur style tableau de bord).
9. Cliquez sur Apply et sauvegardez votre Dashboard.

## Arrêt du système
Pour stopper proprement tous les services et libérer les ports, tapez simplement :
```bash
docker compose down
```
## Licence et Contact
Ce projet est obligatoirement distribué sous licence MIT, imposant la conservation de la notice de copyright pour toute redistribution ou modification.
[GitHub](https://github.com/amd1-7)