/*
quadtree est le paquet qui fournit la structure de données
pour les arbres quaternaires ainsi que les fonctions et
méthodes pour manipuler ces arbres.
*/
package quadtree

/*
Ce ficher servira à documenter l'ensemble des fonctions réalisées
dans le dossier Quadtree.
Pour des raisons d'efficacité lorsque le champs
FloorKind du fichier de config vaut 2 alors la
bibliothèque Quadtree est utilisée:
Nous avons dans un premier temps implémenter les
2 fonctions demandées MakeFromArray et GetContent.
Par la suite, nous les avons modifié pour les améliorer en rajoutant
d'autres fonctions liées à celles-ci.
*/

/*
Fonction MakeFromArray (fichier make.go)

Entrées: floorContent -> un tableau en 2 dimensions d'entiers
Sorties:  q Quadtree -> une structure de l'arbre finale contenant le terrain
Fonction permettant d'initialiser un Quadtree de manière récursive,
en stockant les différents types de terrains et leur emplacement dans
les feuilles de l'arbre. Si il n'y a pas de terrain, la valeur
est -1 sinon en fonction du type de terrain, un entier associé
à celui-ci.
*/

/*
Fonction addContent (fichier addContent.go)

Entrées: currentNode *node -> le pointeur de la node à insérer,
content -> le contenu, targetX / targetY -> les coordonnées
de la feuille
Sorties:
Fonction permettant d'insérer la node mise en entrée dans
le Quadtree de fin. Fonction récursive qui vient optimiser
l'arbre après insertion.
*/

/*
Fonction GetContent (fichier get.go)
Entrées: topLeftX / topLeftY int -> coordonnées de la feuille (point de départ de l'itération),
contentHolder [][]int -> vide au départ
Sorties: contentHolder [][]int -> tableau en 2 dimensions d'entiers avec le contenu de chaque tuile
Fonction permettant de remonter une branche de l'arbre jusqu'a ce que node.content != nil
	- si coordonnées dans contentHolder alors append
	- sinon échec
Si toutes les branches ont été explorées et contentHolder pas remplit,
alors le remplir de tuiles vide (tuile par default)
*/

/*
Fonction seek (fichier seek.go)
Entrées: targetX / targetY -> coordonnées de la feuille (point à atteindre)
Sorties: nodeFound *node -> pointeur de la node trouvée ou emplacement juste avant celle-ci dans le cas
où elle n'existe pas
Fonction appelant l'autre fonction du fichier: seeker, l'ensemble permet alors de trouver le pointeur
de la node où se trouvent ces coordonnées sinon elle renvoie le pointeur de la dernière node non vide
*/
