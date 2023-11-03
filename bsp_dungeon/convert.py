import PIL, sys
from PIL import Image

map_path = "../floor-files/map"
width, height = int(sys.argv[1]), int(sys.argv[2])

map_file = open(map_path, 'r')

img_output = PIL.Image.new(mode="RGBA", size=(width, height))

for y in range(height) :
    for x in range(width) :
        char = map_file.read(1) 
        if char == "1":
            img_output.putpixel((x, y), (50, 50, 50))
        else :
            img_output.putpixel((x, y), (100, 100, 100))
    char = map_file.read(1) # pour passer le \n

map_file.close()
img_output.show()
img_output.save(map_path + ".png")
print("done")