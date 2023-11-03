/*
BSP Dungeon :
Paquet destiné à la génération procédural de donjons
à l'aide d'un algorithme de Binary Space Partition.

Idéalement le temps de génération d'un niveau n'est pas
important, en effet, celui est générer en parallèle du
reste du jeu (ne peut être embêtant que si le joueur
complète trop vite le niveau précédant).
-> TODO: utilisé les coroutines/threads

Contient :
  - création d'un arbre binaire équilibré.
  - génération d'un fichier map à partir de la
    structure de données créer.
*/
package bsp_dungeon