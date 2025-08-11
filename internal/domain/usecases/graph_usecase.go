package usecases

import (
	"fmt"
	"markmind/internal/core/algebra"
	"markmind/internal/core/iterators"
	"markmind/internal/core/moner"
	"markmind/internal/data/entities"
	"markmind/internal/data/repos"
	"math"
	"strings"
)

type GraphUseCase struct {
	explorerRepo *repos.ExplorerRepo
	fileRepo     *repos.FileRepo
}

func NewGraphUseCase(
	explorerRepo *repos.ExplorerRepo,
	fileRepo *repos.FileRepo,
) *GraphUseCase {
	return &GraphUseCase{
		explorerRepo: explorerRepo,
		fileRepo:     fileRepo,
	}
}

var center = algebra.NewVec2D(700, 400)

func (self GraphUseCase) getRotMatrix(alpha float64) algebra.Matrix2D {
	cos := math.Cos(alpha)
	sin := math.Sin(alpha)
	return algebra.NewMatrix2DFromRows(
		[2]float64{cos, -sin},
		[2]float64{sin, cos},
	)
}

func (self GraphUseCase) GetGraph() ([]*entities.GraphNode, error) {
	fmt.Println("call to GetGraph")
	nodes := make([]*entities.GraphNode, 0, 100)
	files, err := self.getAllFiles("/")
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		return nil, err
	}

	for _, file := range files {
		node := file.ToGraphNode()
		// node.X = 100
		// node.Y = 100 + 25*i
		nodes = append(nodes, node)
	}

	fmt.Println("find links")

	lenght := len(nodes)
	for i := 0; i < lenght; i++ {
		node := nodes[i]
		links, err := self.getLinks(node.Path)
		fmt.Printf("%d links have been found in %s\n", len(links), node.Path)
		if err != nil {
			return nil, err
		}
		for _, link := range links {
			for j := 0; j < lenght; j++ {
				tmp := nodes[j]
				if tmp.Path == link {
					node.LinkedNodes = append(node.LinkedNodes, tmp)
					break
				}
			}
		}
	}
	fmt.Println("map node positions")

	mostConnected := self.findMostConnectedFile(nodes)
	nodeStack := make([]*entities.GraphNode, 0, len(nodes)+1)

	mostConnected.Pos.X = center.X
	mostConnected.Pos.Y = center.Y
	nodeStack = append(nodeStack, mostConnected.LinkedNodes...)
	segmentWidth := 2 * math.Pi / float64(len(mostConnected.LinkedNodes))
	distance := algebra.NewVec2D(-20, -20)
	distMatrix := algebra.Identity2D().MultiplyMatrix(algebra.NewMatrix2D(distance, algebra.NewZeroVec2D()))
  fmt.Printf("identity2Matrix: (%v)\n", algebra.Identity2D())
  fmt.Printf("distMatrix (%v)\n", distMatrix)

	for i, neighbour := range mostConnected.LinkedNodes {
		direction := distMatrix.MultiplyMatrix(self.getRotMatrix(segmentWidth * float64(i)))
		self.mapNodePos(mostConnected, neighbour, direction)
	}

	positioned := iterators.Iter(nodes).Filter(func(i int, value *entities.GraphNode) bool {
		return value.Pos.X != 0 || value.Pos.Y != 0
	})

	var maxDist float64 = 0
	maxDiff := algebra.NewVec2D(0, 0)
	for _, node := range positioned {
		diff := node.Pos.Subtract(mostConnected.Pos)
		dist := diff.X*diff.X + diff.Y*diff.Y
		if dist > maxDist {
			maxDist = dist
			maxDiff = diff
		}
	}

	unpositioned := iterators.Iter(nodes).Filter(func(i int, value *entities.GraphNode) bool {
		return value.Pos.X == 0 && value.Pos.Y == 0
	})
	segmentWidth = 2 * math.Pi / float64(lenght)
	distVec := maxDiff.Add(algebra.NewVec2D(20, 20))
	fmt.Printf("distVector: %v\n", distVec)
	distMatrix = algebra.Identity2D().MultiplyMatrix(
		algebra.NewMatrix2D(
			distVec,
			algebra.NewZeroVec2D(),
		),
	)
	for i, node := range unpositioned {
		direction := distMatrix.MultiplyMatrix(self.getRotMatrix(segmentWidth * float64(i)))
		self.mapNodePos(mostConnected, node, direction)
	}

	return nodes, nil
}

func (self GraphUseCase) mapNodePos(
	predecessor,
	current *entities.GraphNode,
	direction algebra.Matrix2D,
) {
	fmt.Printf("pred x: %f, pred y: %f\n", predecessor.Pos.X, predecessor.Pos.Y)
	if current.Pos.X == 0 && current.Pos.Y == 0 {
		newPos := direction.Multiply(algebra.NewVec2D(
			float64(predecessor.Pos.X),
			float64(predecessor.Pos.Y),
		))
		current.Pos.X = newPos.X
		current.Pos.Y = newPos.Y

		segmentWidth := 2 * math.Pi / float64(len(current.LinkedNodes))
		distance := algebra.NewVec2D(-20, -20)
		distMatrix := algebra.Identity2D().MultiplyMatrix(algebra.NewMatrix2D(distance, algebra.NewVec2D(0, 0)))

		for i, node := range current.LinkedNodes {
			direction := distMatrix.MultiplyMatrix(self.getRotMatrix(segmentWidth * float64(i)))
			self.mapNodePos(current, node, direction)
		}
	}
}

func (self GraphUseCase) findMostConnectedFile(nodes []*entities.GraphNode) *entities.GraphNode {
	var mostConnected *entities.GraphNode
	connectionCount := 0
	for _, node := range nodes {
		length := len(node.LinkedNodes)
		if length >= connectionCount {
			mostConnected = node
		}
	}
	return mostConnected
}

func (self GraphUseCase) getLinks(path string) ([]string, error) {
	getAddresses := moner.Fmap(self.getAllAddresses, self.fileRepo.GetFileContent)
	return getAddresses(path)
}

func (self GraphUseCase) getAllAddresses(byteContent []byte) []string {
	lines := strings.Split(string(byteContent), "\n")
	links := make([]string, 0, 10)
	for _, line := range lines {
		links = append(links, self.extractLinkAddresses(line)...)
	}
	return links
}

func (self GraphUseCase) extractLinkAddresses(line string) []string {
	const (
		irrelevant = iota
		desc       = iota
		descEnd    = iota
		address    = iota
	)
	linkBuilder := strings.Builder{}
	links := make([]string, 0, 2)
	state := irrelevant
	for _, c := range line {
		switch state {
		case irrelevant:
			if c == '[' {
				state = desc
			}
		case desc:
			if c == ']' {
				state = descEnd
			}
		case descEnd:
			if c == '(' {
				state = address
			}
		case address:
			if c == ')' {
				state = irrelevant
				links = append(links, linkBuilder.String())
				linkBuilder.Reset()
			} else {
				linkBuilder.WriteRune(c)
			}
		}
	}
	return links
}

func (self GraphUseCase) getAllFiles(path string) ([]entities.MarkdownFileMeta, error) {
	fmt.Printf("get files of: %s\n", path)
	entries, err := self.explorerRepo.GetEntriesOfDirectory(path)
	if err != nil {
		return nil, err
	}

	leaves := make([]entities.MarkdownFileMeta, 0, len(entries))
	for _, entry := range entries {
		fmt.Println(entry)

		if entry.Name == ".." {
			continue
		} else if entry.IsDir {
			tmp, err := self.getAllFiles(entry.Path)
			if err != nil {
				return nil, err
			}
			leaves = append(leaves, tmp...)
		} else {
			leaves = append(leaves, entry)
		}
	}
	return leaves, nil
}
