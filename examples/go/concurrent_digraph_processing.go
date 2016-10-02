package main

import (
	"math/rand"
	"time"
	"fmt"
)

const (
	NumberOfConcurrentProcessors = 32
	NumberOfCrawlIterations = 4
)

// DiGraph
//
// This is a directional graph where edges can point in any direction between vertices.
// The graph has 1 root vertex, which is where the Crawl() starts from.
type DiGraph struct {
	RootVertex       *Vertex       // The root vertex of the graph
	Processors       [] *Processor // List of concurrent processors
	ProcessorIndex   int           // The current index of the next processor to use
	Edges            []*Edge       // All directional edges that make up the graph
	Iterations       int           // The total number of times to iterate over the graph
	TotalVertices    int           // Count of the total number of vertices that make up the graph
	ProcessedChannel chan int      // Channel to track processed vertices
	ProcessedCount   int           // Total number of processed vertices
	TotalOperations  int           // Total number of expected operations | [(TotalVertices * Iterations) - Iterations] + 1
}

// Vertex
//
// A single unit that composes the graph. Each vertex has relationships with other vertices,
// and should represent a single entity or unit of work.
type Vertex struct {
	Name   string // Unique name of this Vertex
	Edges  []*Edge
	Status int
}

// Edge
//
// Edges connect vertices together. Edges have a concept of how many times they have been processed
// And a To and From direction
type Edge struct {
	To             *Vertex
	From           *Vertex
	ProcessedCount int
}

// Processor
//
// This represents a single concurrent process that will operate on N number of vertices
type Processor struct {
	Function func(*Vertex) int
	Channel  chan *Vertex
}

// Init the graph with a literal definition
var TheGraph *DiGraph = NewGraph()

func main() {
	TheGraph.Init(NumberOfConcurrentProcessors, NumberOfCrawlIterations)
	TheGraph.Crawl()
}

func (d *DiGraph) Init(n, i int) {
	noProcs := n
	d.TotalVertices = d.RootVertex.recursiveCount()
	d.Iterations = i
	for ; n > 0; n-- {
		p := Processor{Channel: make(chan *Vertex)}
		d.Processors = append(d.Processors, &p)
		p.Function = Process
		go p.Exec()
	}
	d.TotalOperations = (d.TotalVertices * d.Iterations) - d.Iterations + 1 //Math is hard
	fmt.Printf("Total Vertices              : %d\n", d.TotalVertices)
	fmt.Printf("Total Iterations            : %d\n", d.Iterations)
	fmt.Printf("Total Concurrent Processors : %d\n", noProcs)
	fmt.Printf("Total Assumed Operations    : %d\n", d.TotalOperations)
}

func (d *DiGraph) Crawl() {
	d.ProcessedChannel = make(chan int)
	go d.RootVertex.recursiveProcess(d.getProcessor().Channel)
	fmt.Printf("---\n")
	for d.ProcessedCount < d.TotalOperations {
		d.ProcessedCount += <-d.ProcessedChannel
		printColor(fmt.Sprintf("%d ", d.ProcessedCount))
	}
	fmt.Printf("\n---\n")
	fmt.Printf("Total Comlpeted Operations  : %d\n", d.ProcessedCount)
}

func (d *DiGraph) getProcessor() *Processor {
	maxIndex := len(d.Processors) - 1
	if d.ProcessorIndex == maxIndex {
		d.ProcessorIndex = 0
	} else {
		d.ProcessorIndex += 1
	}
	return d.Processors[d.ProcessorIndex]
}

func Process(v *Vertex) int {

	// Simulate some work with a random sleep
	rand.Seed(time.Now().Unix())
	sleep := rand.Intn(100 - 0) + 100
	time.Sleep(time.Millisecond * time.Duration(sleep))

	// Return a status code
	return 1
}

func (v *Vertex) recursiveProcess(ch chan *Vertex) {
	ch <- v
	for _, e := range v.Edges {
		if e.ProcessedCount < TheGraph.Iterations {
			e.ProcessedCount += 1
			go e.To.recursiveProcess(TheGraph.getProcessor().Channel)
		}
	}
}

func (v *Vertex) recursiveCount() int {
	i := 1
	for _, e := range v.Edges {
		if e.ProcessedCount != 0 {
			e.ProcessedCount = 0
			i += e.To.recursiveCount()
		}
	}
	return i
}

func (v *Vertex) AddVertex(name string) *Vertex {
	newVertex := &Vertex{Name: name}
	newEdge := &Edge{To: newVertex, From: v, ProcessedCount: -1}
	newVertex.Edges = append(newVertex.Edges, newEdge)
	v.Edges = append(v.Edges, newEdge)
	return newVertex
}

func (p *Processor) Exec() {
	for {
		v := <-p.Channel
		v.Status = p.Function(v)
		TheGraph.ProcessedChannel <- 1
	}
}

func NewGraph() *DiGraph {
	rootVertex := &Vertex{Name: "0"}
	v1 := rootVertex.AddVertex("1")
	rootVertex.AddVertex("2")
	rootVertex.AddVertex("3")
	v1.AddVertex("1-1")
	v1.AddVertex("1-2")
	v1_3 := v1.AddVertex("1-3")
	v1_3.AddVertex("1-3-1")
	v1_3.AddVertex("1-3-2")
	v1_3_3 := v1_3.AddVertex("1-3-3")
	v1_3_3.AddVertex("1-3-3-1")
	v1_3_3.AddVertex("1-3-3-2")
	v1_3_3.AddVertex("1-3-3-3")
	v1_3_3.AddVertex("1-3-3-4")
	v1_3_3.AddVertex("1-3-3-5")
	v1_3_3.AddVertex("1-3-3-6")
	v1_3_3.AddVertex("1-3-3-7")
	graph := &DiGraph{}
	graph.RootVertex = rootVertex
	return graph
}

func printColor(str string) {
	fmt.Printf("\033[0;34m%s\033[0m", str)
}