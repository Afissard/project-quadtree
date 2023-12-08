# Project-Quadtree
Projet d'introduction au développement (R1.01) et de SAÉ implémentation d'un besoin client (SAE1.01), année 2023-2024.

# Contenu du projet
## SAE
- floor : readFromFile
- floor : updateFromFile
- quadtree : MakeFromArray
- quadtree : GetContent
## Extensions implémenté
- optimisation quadtree 
- map aléatoire : donjon BSP (TODO : BSP -> WFC)
- sauvegarde de la map
## En cours d'implémentation
- Sauvegarde avancé : sauvé la position du joueur et autres informations utiles dans un json
- GUI : Graphical User Interface (paramètres, info joueur, etc)
## TODO list
- Déplacement du joueur indépendant des tuiles (=> fix le scrolling de la caméra)
- Monde procédural en direct (taille fixe) (donjons -> algo de génération (Wave Fonction Collapse en live))
- Nouvelle tuiles etc (collision et autre interaction)
- Monde infini & procédural
- dynamic light (utilisé technique de dithering ?)
- Ajouter de la vie ("ia" algo state machine + path-finding : Dijkstra (ou A* All Star si temps))

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
```powershell
cd cmd && .\cmd
```

# Divers informations
## Libraries :
- Ebiten engine : https://ebitengine.org
## Autres informations
- repo source projet : https://gitlab.univ-nantes.fr/jezequel-l/project-quadtree/-/releases
- Google doc des idées : https://docs.google.com/document/d/1e6MvH3YVjEUX3oUGH9tsGn1JIZyKfRrJr4DWm7M6GVA/edit
# Ressources
## Exemple de projet avec Ebiten :
- Github de Loig Jezequel : https://github.com/loig
- Page itch de Loig Jezequel : https://loig.itch.io
- Discord de Ebiten Engine
## Ressources Wave Fonction Collapse
- WFC repository : https://github.com/mxgmn/WaveFunctionCollapse
- WFC implemented in Golang : https://github.com/shawnridgeway/wfc
