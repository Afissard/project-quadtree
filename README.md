# Project-Quadtree
Projet d'introduction au développement (R1.01) et de SAÉ implémentation d'un besoin client (SAE1.01) réalisé par Léa Cortet et Sacha Chauvel, année 2023-2024.

# Sommaire
1. [Contenu du projet](#contenu-du-projet)
    - [SAE (partie 5)](#sae-partie-5)
    - [Extensions implémenté (partie 6)](#extensions-implémenté-partie-6)
    - [Extensions en cours d'implémentation](#extensions-en-cours-dimplémentation)
2. [Contrôles](#contrôles)
3. [Build & exécution du projet](#build--exécution-du-projet)
    - [Linux](#linux)
    - [Window](#window)
4. [Bugs connus](#bugs-connus)
5. [Ressources](#ressources)

# Contenu du projet
## SAE (partie 5)
- floor : readFromFile
- floor : updateFromFile
- quadtree : MakeFromArray
- quadtree : GetContent
## Extensions implémenté (partie 6)
- Optimisation quadtree 
- Sauvegarde avancé : sauvé la position du joueur et autres informations utiles dans un json
- UI : interface utilisateur (menu principal, sauvegarde, chargement de la sauvegarde)
- Zoom camera (sert aussi de minimap)
- Nouvelle tuiles etc (collision et autre interaction), (FloorKind = 3)
- Monde procédural (génération aléatoire, taille fixe) : donjons -> algo de génération BSP, (FloorKind = 3)
- Portail de téléportation (générer dans le donjon (FloorKind = 3))
## Extensions en cours d'implémentation
- Monde infini à la volé (sans algo de génération), (FloorKind = 4)

# Contrôles
| Action | Touches associées |
| ---- | ---- |
| Déplacement haut | flèche haut |
| Déplacement bas | flèche bas |
| Déplacement gauche | flèches gauche |
| Déplacement droite | flèches droite |
| Retour au menu | Escape |
| Quitter le jeu en cours | Backspace |
| Ouverture de l'interface de debug | F1 |
| Dé-zoom de la camera | F2 |
| Zoom de la camera | F3 |


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

# Bugs connus
- Le dé-zoom peu causé des chute de FPS, de plus pour une grande map affiché avec un dé-zoom au maximum (supérieur à 6) la rendering pipeline d'Ebitengine peut causé un crash du jeu
- En floorKing 4, lorsque la taille du quadtree est augmenté, son contenu disparais (la génération infinie à la volé est toujours WIP)
- Seed non déterministe lors de la génération de donjons.
- Jouer dans un niveau charger depuis une sauvegarde rend les téléporteur inutilisable pour cet étage.
- Dans de très rare cas il est possible d'allez sur une tuile normalement incessible (exemple mur) en sortant d'un téléporteur et en continuant de marcher.

# Ressources
## Ebitengine Documentation :
- Ebiten website : https://ebitengine.org
- Documentation : https://ebitengine.org/en/documents
- Exemples : https://ebitengine.org/en/examples/
## Exemple de projets fait avec Ebitengine :
- Showcase du site d'Ebiten : https://ebitengine.org/en/showcase.html
- Jeux fait avec Ebitengine sur Itch.io : https://itch.io/games/tag-ebitengine
- Github de Loig Jezequel : https://github.com/loig
- Page Itch de Loig Jezequel : https://loig.itch.io
## Exemple et explication des Quadtree :
- Site interactif : https://jimkang.com/quadtreevis/
- Exemple d'implémentation d'un quadtree en go : https://pkg.go.dev/github.com/paulmach/orb@v0.10.0/quadtree
## Binary Space Partition Dungeon
- Site explicatif pour l'algorithme: https://www.roguebasin.com/index.php/Basic_BSP_Dungeon_generation
## Wave Fonction Collapse
- Wave Fonction Collapse explication : https://robertheaton.com/2018/12/17/wavefunction-collapse-algorithm/
- WFC repository : https://github.com/mxgmn/WaveFunctionCollapse
- Exemple d'implémentation en Golang : https://github.com/shawnridgeway/wfc