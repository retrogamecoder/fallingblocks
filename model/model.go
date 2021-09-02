package model

import (
	"math/rand"
	"time"
)

const (
	Width  = 15
	Height = 60

	ShapeWidth  = 4
	ShapeHeight = 2
)

type TileKind uint8

const (
	TileNone TileKind = iota
	TileS
	TileZ
	TileL
	TileJ
	TileI
	TileO
	TileT
	TileLast
)

type Direction uint8

const (
	DirectionNone Direction = iota
	DirectionUp
	DirectionDown
	DirectionLeft
	DirectionRight
)

type Block struct {
	X     uint8
	Y     uint8
	Angle Direction
	Kind  TileKind
}

func bottomRow(b *Block) [][2]uint8 {
	return nil
}

type Model interface {
	Tiles() [Width][Height + 1]TileKind
	CurrentBlock() *Block
	NextBlock() *Block
	Score() uint32
	Level() uint8

	Update(dt time.Duration)
	MoveCurrentBlock(direction Direction)
}

type model struct {
	tiles [Width][Height + 1]TileKind
	score uint32
	level uint8

	currentBlock *Block
	nextBlock    *Block

	lastUpdated time.Time

	rand *rand.Rand
}

func New() Model {
	r := rand.New(rand.NewSource(42))
	return &model{
		rand: r,
	}
}

func (m *model) Tiles() [Width][Height + 1]TileKind { return m.tiles }
func (m *model) CurrentBlock() *Block               { return m.currentBlock }
func (m *model) NextBlock() *Block                  { return m.nextBlock }
func (m *model) Score() uint32                      { return m.score }
func (m *model) Level() uint8                       { return m.level }

func (m *model) Update(dt time.Duration) {
	if time.Now().Before(m.nextUpdate()) {
		return
	}

	m.checkBlocks()

	if m.isCollision() {
		m.freezeCurrent()
		m.currentBlock = nil
		return
	}

	m.MoveCurrentBlock(DirectionDown)
}

func (m *model) isCollision() bool {
	return false
}

func (m *model) freezeCurrent() {
}

func (m *model) nextUpdate() time.Time {
	return m.lastUpdated.Add(m.levelGap())
}

func (m *model) checkBlocks() {
	if m.nextBlock == nil {
		m.nextBlock = m.randomBlock()
	}

	if m.currentBlock != nil {
		return
	}

	m.currentBlock = m.nextBlock
	m.nextBlock = m.randomBlock()
}

func (m *model) randomBlock() *Block {
	return &Block{
		X:    m.startingX(),
		Y:    m.startingY(),
		Kind: TileKind(m.rand.Intn(int(TileLast)-1) + 1),
	}
}

func (m *model) startingX() uint8 {
	return uint8(Width / 2)
}

func (m *model) startingY() uint8 {
	return uint8(Height)
}

func (m *model) levelGap() time.Duration {
	return 2 * time.Second
}

func (m *model) MoveCurrentBlock(direction Direction) {}
