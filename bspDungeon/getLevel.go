package bspDungeon

func GetLevel(seed string, width, height int) (level BSP_tree, floorContent [][]int) {
	level = CreateLevel(seed, NBR_DIV, width, height)
	return level, level.ToMap()
}

func GetLevelNoTree(seed string, width, height int) (floorContent [][]int) {
	level := CreateLevel(seed, NBR_DIV, width, height)
	return level.ToMap()
}
