package floor

import (
	"bufio"
	"os"
	"strconv"

	// "github.com/Afissard/project-quadtree/bspDungeon"
	// "github.com/Afissard/project-quadtree/configuration"
	// "github.com/Afissard/project-quadtree/quadtree"

	"gitlab.univ-nantes.fr/pub/but/but1/r1.01/sae1.01/groupe4/eq-4-05_chauvel-sacha_cortet-lea/configuration"
	"gitlab.univ-nantes.fr/pub/but/but1/r1.01/sae1.01/groupe4/eq-4-05_chauvel-sacha_cortet-lea/dungeon"
	"gitlab.univ-nantes.fr/pub/but/but1/r1.01/sae1.01/groupe4/eq-4-05_chauvel-sacha_cortet-lea/quadtree"
)

// Initialise les structures de données internes de f.
func (f *Floor) Init() (spawnX, spawnY int, listPortal []dungeon.Portal) {
	f.content = make([][]int, configuration.Global.NumTileY)
	for y := 0; y < len(f.content); y++ {
		f.content[y] = make([]int, configuration.Global.NumTileX)
	}

	spawnX, spawnY = 0, 0
	switch configuration.Global.FloorKind { // INFO: dépend du config.json
	case fromFileFloor: // INFO: utilise d'abord un fichier pour la map avant algo quadtree
		f.fullContent = readFloorFromFile(configuration.Global.FilePath + configuration.Global.FloorFile)
	case quadTreeFloor:
		f.QuadtreeContent = quadtree.MakeFromArray(readFloorFromFile(configuration.Global.FilePath + configuration.Global.FloorFile))
	case bspDungeonFloor:
		if configuration.Global.FloorFile == configuration.Global.SaveFile {
			f.QuadtreeContent = quadtree.MakeFromArray(readFloorFromFile(configuration.Global.FilePath + configuration.Global.FloorFile))
		} else {
			f.QuadtreeContent, spawnX, spawnY, listPortal = dungeon.NewDungeon(configuration.Global.FixedGenWidth, configuration.Global.FixedGenHeight, configuration.Global.Seed)
		}
	case infiniteWorld:
		f.QuadtreeContent.Init(2, 2) // taille minimal du quadtree (1, 1 est possible aussi)
	}
	return spawnX, spawnY, listPortal
}

/*
lecture du contenu d'un fichier représentant un terrain
pour le stocker dans un tableau
*/
func readFloorFromFile(fileName string) (floorContent [][]int) {
	/*
		Récupère les données du fichier map : floor-files/filename
		Retourne les données dans une liste de int à deux
		dimensions : floorContent
	*/
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	file_scan := bufio.NewScanner(file)
	file_scan.Split(bufio.ScanRunes) // scan char par char
	lineContent := []int{}

	for file_scan.Scan() {
		str_content := file_scan.Text() // récupère les données au format texte
		if err != nil {
			panic(err)
		}
		// si nouvelle ligne, n'ajoute pas le caractère d’échappement
		if str_content == "\r" { // est parfois détecté par Atoi()
			continue
		} else if str_content == "\n" { // retour à la ligne
			floorContent = append(floorContent, lineContent) // ajoute la ligne à floorContent
			lineContent = nil
		} else {
			content, err := strconv.Atoi(str_content) // str -> int
			if err != nil {
				// panic(err)
				content = -1
			}
			lineContent = append(lineContent, content)
		}
	}
	// else implicite pour la dernière ligne (il y a sûrement une meilleur solution)
	floorContent = append(floorContent, lineContent)

	file.Close()
	return floorContent
}
