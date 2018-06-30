package tiles

import (
	"fmt"
	"image"
	"os"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/lafriks/go-tiled"
)

const mapPath = "tiles/tilemap.tmx" // path to your map
const tileSize = 16
const mapWidth = 30
const mapHeight = 30

// loadPicture load the picture
func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func getTilesFrames(spritesheet pixel.Picture) []pixel.Rect {
	var tilesFrames []pixel.Rect
	for y := spritesheet.Bounds().Max.Y - tileSize; y > spritesheet.Bounds().Min.Y; y -= tileSize {
		for x := spritesheet.Bounds().Min.X; x < spritesheet.Bounds().Max.X; x += tileSize {
			tilesFrames = append(tilesFrames, pixel.R(x, y, x+tileSize, y+tileSize))
		}
	}

	return tilesFrames
}

func getOrigin(win *pixelgl.Window) pixel.Vec {
	centerPosition := win.Bounds().Center()
	originXPosition := centerPosition.X - mapWidth/2*tileSize
	originYPosition := centerPosition.Y + mapHeight/2*tileSize - tileSize

	return pixel.V(originXPosition, originYPosition)
}

func getSpritePosition(spriteIndex int, origin pixel.Vec) pixel.Vec {
	spriteXPosition := origin.X + float64((spriteIndex%mapWidth)*tileSize) + tileSize/2
	spriteYPosition := origin.Y + tileSize/2 - float64((spriteIndex/mapWidth)*tileSize)

	return pixel.V(spriteXPosition, spriteYPosition)
}

// GenerateMap generates the map
func GenerateMap(win *pixelgl.Window) (pixel.Picture, []pixel.Rect) {
	// parse tmx file
	gameMap, err := tiled.LoadFromFile(mapPath)
	if err != nil {
		fmt.Println("Error parsing map")
		os.Exit(2)
	}

	spritesheet, err := loadPicture("tiles/tileset.png")
	if err != nil {
		panic(err)
	}

	tilesFrames := getTilesFrames(spritesheet)

	originPosition := getOrigin(win)

	for index, layerTile := range gameMap.Layers[0].Tiles {
		sprite := pixel.NewSprite(spritesheet, tilesFrames[layerTile.ID])
		spritePosition := getSpritePosition(index, originPosition)
		sprite.Draw(win, pixel.IM.Moved(spritePosition))
	}

	return spritesheet, tilesFrames
}
