package logic

import (
	"container/list"
	"sync"
)

const LOGICAL_WIDTH = 100
const LOGICAL_HEIGHT = 10
const WORLD_WIDTH = 6400
const WORLD_HEIGHT = 600

type Level interface {
	Iterator() Iterator
}

type Iterator interface {
	HasNext() bool
	Next()
	Get() *Block
}

type ListLevel struct {
	blocks *list.List
}

type LevelIterator struct {
	element *list.Element
}

func (level ListLevel) Iterator() Iterator {
	return &LevelIterator{level.blocks.Front()}
}

func (iterator LevelIterator) HasNext() bool {
	return iterator.element != nil
}

func (iterator *LevelIterator) Next() {
	iterator.element = iterator.element.Next()
}

func (iterator LevelIterator) Get() *Block {
	return iterator.element.Value.(*Block)
}

func NewLevel() *ListLevel {
	level := new(ListLevel)
	level.blocks = list.New()
	return level
}

func initLevel(level *ListLevel) {
	addFloor(level.blocks)
	addCeiling(level.blocks)
}

func addFloor(level *list.List) {
	for i := 0; i < WORLD_WIDTH; i += BLOCK_WIDTH_PIXELS {
		object := new(Block)
		object.accelerationX = 0
		object.accelerationY = 0
		object.blockType = BLOCK_TYPE_FLOOR
		object.friction = -1
		object.id = i
		object.posX = i
		object.posY = WORLD_HEIGHT - BLOCK_HEIGHT_PIXELS
		object.speedX = 0
		object.speedY = 0
		object.width = BLOCK_WIDTH_PIXELS
		object.height = BLOCK_HEIGHT_PIXELS
		object.solid = true
		level.PushFront(object)
	}
}

func addCeiling(level *list.List) {
	for i := 0; i < WORLD_WIDTH/BLOCK_WIDTH_PIXELS; i++ {
		object := new(Block)
		object.accelerationX = 0
		object.accelerationY = 0
		if i != 0 && i%5 == 0 {
			object.blockType = BLOCK_TYPE_FALLING_ROCK
		} else {
			object.blockType = BLOCK_TYPE_ROCK
		}
		object.friction = 0
		object.id = i
		object.posX = i * BLOCK_WIDTH_PIXELS
		object.posY = 0
		object.speedX = 0
		object.speedY = 0
		object.width = BLOCK_WIDTH_PIXELS
		object.height = BLOCK_HEIGHT_PIXELS
		object.solid = true
		level.PushFront(object)
	}
}

func applyLevelLogic(game *Game) {
	waitGroup := new(sync.WaitGroup)

	for iterator := game.level.Iterator(); iterator.HasNext(); iterator.Next() {
		block := iterator.Get()
		if block != nil {
			waitGroup.Add(1)
			go func() {
				applyBlockLogic(block, game)
				waitGroup.Done()
			}()
		}
	}
	waitGroup.Wait()
}
