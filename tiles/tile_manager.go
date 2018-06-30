package tiles

import (
	"fmt"
	"image"
	"os"

	"github.com/gandrin/ASharedJourney/shared"

	_ "image/png"

	"log"
	"path"
	"runtime"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/lafriks/go-tiled"
)

func extractAndPlaceSprites(
	layerTiles []*tiled.LayerTile,
	spritesheet pixel.Picture,
	tilesFrames []pixel.Rect,
	originPosition pixel.Vec,
) (positionedSprites []SpriteWithPosition) {
	for index, layerTile := range layerTiles {
		if !layerTile.IsNil() {
			sprite := pixel.NewSprite(spritesheet, tilesFrames[layerTile.ID])
			spritePosition := getSpritePosition(index, originPosition)
			positionedSprites = append(positionedSprites, SpriteWithPosition{
				Sprite:   sprite,
				Position: spritePosition,
			})
		}
	}
	return positionedSprites
}

const mapPath = "tiles/tilemap.tmx"   // path to your map
const tilesPath = "tiles/tileset.png" // path to your tileset
const tileSize = 16
const mapWidth = 30
const mapHeight = 30

type World struct {
	BackgroundTiles []SpriteWithPosition
	Players         [1]SpriteWithPosition
	Movables 		[1]SpriteWithPosition
}

//SpriteWithPosition holds the sprite and its position into the window
type SpriteWithPosition struct {
	Sprite   *pixel.Sprite
	Position pixel.Vec
}

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

// GenerateMap generates the map from a .tmx file
func GenerateMap() (pixel.Picture, []pixel.Rect, World) {
	//get path to file from current programme root
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		log.Fatal("error loading called")
	}
	filemap := path.Join(path.Dir(filename), mapPath)
	filetile := path.Join(path.Dir(filename), tilesPath)

	gameMap, err := tiled.LoadFromFile(filemap)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Error parsing map")
		os.Exit(2)
	}

	spritesheet, err := loadPicture(filetile)
	if err != nil {
		panic(err)
	}

	tilesFrames := getTilesFrames(spritesheet)

	originPosition := getOrigin(shared.Win)

	positionedSprites := extractAndPlaceSprites(gameMap.Layers[0].Tiles, spritesheet, tilesFrames, originPosition)

	// TODO iterate over objects to look for "player" object
	// TODO make sure the given input is a multiple of tileSize
	// playerLayer := gameMap.Layers[2].Tiles
	playerTiledObject := gameMap.ObjectGroups[0].Objects[0]
	player1X := playerTiledObject.X + int(originPosition.X)
	player1Y := -playerTiledObject.Y + int(originPosition.Y)

	fmt.Println(player1X)
	fmt.Println(player1Y)

	player1 := SpriteWithPosition{Sprite: pixel.NewSprite(spritesheet, tilesFrames[203]), Position: pixel.V(float64(player1X), float64(player1Y))}
	var players [1]SpriteWithPosition
	players[0] = player1
	world := World{BackgroundTiles: positionedSprites, Players: players}

	return spritesheet, tilesFrames, world
}

//DrawMap draws into window the given sprites
func DrawMap(positionedSprites []SpriteWithPosition) {
	for _, positionedSprite := range positionedSprites {
		positionedSprite.Sprite.Draw(shared.Win, pixel.IM.Moved(positionedSprite.Position))
	}
}
