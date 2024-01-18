/*
Ce paquet sert à créer des donjons avec un algorithme d'arbre binaire.

L'arbre binaire est un arbre BSP (Binary Space Partition) :
Explication général de l'algorithme :
	1. A chaque node de l'arbre, l'on partitionne l'espace disponible en le coupant en deux,
	soit sur la largeur, soit sur la longueur, le point de coupe est choisie aléatoirement. 
	La création de branche est répété jusqu'à ce qu'il n'y ai plus d’espace disponible.
	2. Chaque feuille de l'arbre voient leur dimension modifié pour représenté les salles.
	3. A partir des node parentes les couloirs sont créer pour relier les node enfants.

Voici un site explicatif : https://www.roguebasin.com/index.php/Basic_BSP_Dungeon_generation

*/
package dungeon
