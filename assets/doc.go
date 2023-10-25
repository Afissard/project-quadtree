/*
assets est le paquet en charge d'importer les images du jeu
pour qu'elles soient utilisables par Ebitengine. Ceci se fait
en deux étapes :

 1. À la compilation, les directive go:embed permettent de
    recopier les contenus binaires des fichiers indiqués dans les
    variables xxxBytes qui viennent juste après ces directives.

 2. À l'exécution, les séquences d'octets représentant les
    fichiers sont converties en images pour Ebitengine (type *ebiten.Image).

Cela signifie qu'il faut toujours avoir appelé la fonction Load
de ce paquet avant d'utiliser ces images. Pour cela, cette
fonction est appelée au début du main dans le paquet quadtree/cmd.
*/
package assets
