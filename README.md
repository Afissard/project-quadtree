# Project-Quadtree
Projet d'introduction au développement (R1.01) et de SAÉ implémentation d'un besoin client (SAE1.01), année 2023-2024.

# Build & exécution du projet
## Linux
Si la librairie X11-dev n'est pas installé :
```bash
sudo apt-get install libx11-dev xserver-xorg-dev xorg-dev
```
Build :
```bash
cd cmd && go build . && ./cmd
```
Run :
```bash
cd cmd && ./cmd
```
## Window
Build :
```powershell
cd cmd && go build . && .\cmd
```
Run :
Build :
```powershell
cd cmd && .\cmd
```

# Divers informations
## Libraries :
- Ebiten engine : https://ebitengine.org
- Dear ImGUI for Ebiten-engine : https://github.com/gabstv/ebiten-imgui
## Autres informations
- repo source projet : https://gitlab.univ-nantes.fr/jezequel-l/project-quadtree/-/releases
- Google doc des idées : https://docs.google.com/document/d/1e6MvH3YVjEUX3oUGH9tsGn1JIZyKfRrJr4DWm7M6GVA/edit

# TODO list:
Suivre ordre de création des fonction proposer par le pdf (le temps de comprendre la structure du projet).
Créer des fonctions de tests pour chaque fonction créé/modifié.
## Projet en 2 grande partie :
### Partie 1.1 : Affichage map, depuis un fichier
- Récupérer la map depuis un fichier : floor/readFloorFromFile
- Déterminé la région de la map à afficher : floor/updateFromFileFloor
### Partie 1.2 : Algo quadtree
Générer la map avec l'algo quadtree, détail dans le pdf
- makeFromArray : stock la map dans un quadtree (à optimiser)
- getContent : donne une region de la map depuis un quadtree (à modifié quand MakeFromArray sera optimisé)
### Partie 2 : Extensions
- Monde procédural en direct (taille fixe) (donjons -> algo de génération (Wave Fonction Collapse en live))
- Déplacement du joueur indépendant des tuiles (=> fix le scrolling de la caméra)
- GUI (paramètres, info joueur, (dialogue ?))
- Collisions
- Ajouter de la vie (ia state machine algo + path-finding algo)
- dynamic light (voir technique de dithering)
- Monde infini & procédural
- etc