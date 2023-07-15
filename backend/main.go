package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

type Graph struct {
	adjacencyMap map[string][]string
	indexMap     map[string]int
	lowLinkMap   map[string]int
	onStackMap   map[string]bool
	stack        []string
	index        int
	SCC          [][]string
}

func NewGraph() *Graph {
	return &Graph{
		adjacencyMap: make(map[string][]string),
		indexMap:     make(map[string]int),
		lowLinkMap:   make(map[string]int),
		onStackMap:   make(map[string]bool),
		stack:        []string{},
		index:        0,
		SCC:          [][]string{},
	}
}

func NewGraphBridge() *GraphBridge {
	return &GraphBridge{
		Nodes:   make(map[string]*NodeBridge),
		Bridges: make([]*Bridge, 0),
		time:    0,
	}
}

type GraphBridge struct {
	Nodes   map[string]*NodeBridge
	Bridges []*Bridge
	time    int
}

type NodeBridge struct {
	Name     string
	Index    int
	LowLink  int
	Visited  bool
	Adjacent []*NodeBridge
}

type Bridge struct {
	Start *NodeBridge
	End   *NodeBridge
}

func (g *GraphBridge) AddNodeBridge(name string) {
	if _, ok := g.Nodes[name]; !ok {
		node := &NodeBridge{
			Name:     name,
			Index:    -1,
			LowLink:  -1,
			Visited:  false,
			Adjacent: make([]*NodeBridge, 0),
		}
		g.Nodes[name] = node
	}
}

func (g *GraphBridge) AddEdgeBridge(start, end string) {
	g.AddNodeBridge(start)
	g.AddNodeBridge(end)
	startNode := g.Nodes[start]
	endNode := g.Nodes[end]
	startNode.Adjacent = append(startNode.Adjacent, endNode)
}

func (g *Graph) AddEdge(u, v string) {
	g.adjacencyMap[u] = append(g.adjacencyMap[u], v)
}

func (g *Graph) Tarjan() {
	for node := range g.adjacencyMap {
		if g.indexMap[node] == 0 {
			g.dfs(node)
		}
	}
}

func (g *Graph) dfs(node string) {
	g.index++
	g.indexMap[node] = g.index
	g.lowLinkMap[node] = g.index
	g.stack = append(g.stack, node)
	g.onStackMap[node] = true

	for _, neighbor := range g.adjacencyMap[node] {
		if g.indexMap[neighbor] == 0 {
			g.dfs(neighbor)
			g.lowLinkMap[node] = min(g.lowLinkMap[node], g.lowLinkMap[neighbor])
		} else if g.onStackMap[neighbor] {
			g.lowLinkMap[node] = min(g.lowLinkMap[node], g.indexMap[neighbor])
		}
	}

	if g.lowLinkMap[node] == g.indexMap[node] {
		scc := []string{}
		for {
			lastIndex := len(g.stack) - 1
			lastNode := g.stack[lastIndex]
			g.stack = g.stack[:lastIndex]
			g.onStackMap[lastNode] = false
			scc = append(scc, lastNode)
			if lastNode == node {
				break
			}
		}
		g.SCC = append(g.SCC, scc)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (g *GraphBridge) FindBridges() {
	for _, node := range g.Nodes {
		if !node.Visited {
			g.FindBridgesDFS(node, nil)
		}
	}
}

func (g *GraphBridge) FindBridgesDFS(current, parent *NodeBridge) {
	current.Visited = true
	current.Index = g.time
	current.LowLink = g.time
	g.time++

	for _, neighbor := range current.Adjacent {
		if neighbor == parent {
			continue
		}

		if !neighbor.Visited {
			g.FindBridgesDFS(neighbor, current)
			current.LowLink = min(current.LowLink, neighbor.LowLink)
			if neighbor.LowLink > current.Index {
				g.Bridges = append(g.Bridges, &Bridge{current, neighbor})
			}
		} else {
			current.LowLink = min(current.LowLink, neighbor.Index)
		}
	}
}

func visualizeGraph(graph *Graph, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	fmt.Fprintln(writer, "digraph G {")
	for u, neighbors := range graph.adjacencyMap {
		for _, v := range neighbors {
			fmt.Fprintf(writer, "  %s -> %s;\n", u, v)
		}
	}
	fmt.Fprintln(writer, "}")
	writer.Flush()

	outputPath := filepath.Join("uploads", "graf.png")
	cmd := exec.Command("dot", "-Tpng", "-o", outputPath, filename)
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func visualizeBridges(graph *Graph, graphBridge *GraphBridge, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	fmt.Fprintln(writer, "digraph G {")
	for _, bridge := range graphBridge.Bridges {
		fmt.Fprintf(writer, "  %s -> %s;\n", bridge.Start.Name, bridge.End.Name)
	}
	fmt.Fprintln(writer, "}")
	writer.Flush()

	outputPath := filepath.Join("uploads", "bridge.png")
	cmd := exec.Command("dot", "-Tpng", "-o", outputPath, filename)
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func visualizeSCC(graph *Graph, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	fmt.Fprintln(writer, "digraph G {")
	for _, scc := range graph.SCC {
		if len(scc) > 1 {
			for i := 0; i < len(scc)-1; i++ {
				fmt.Fprintf(writer, "  %s -> %s;\n", scc[i+1], scc[i])
			}
			fmt.Fprintf(writer, "  %s -> %s;\n", scc[0], scc[len(scc)-1])
		} else {
			fmt.Fprintf(writer, "  %s;\n", scc[0])
		}
	}
	fmt.Fprintln(writer, "}")
	writer.Flush()

	outputPath := filepath.Join("uploads", "scc.png")
	cmd := exec.Command("dot", "-Tpng", "-o", outputPath, filename)
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	// Hapus semua file dalam folder "uploads"
	err := os.RemoveAll("uploads")
	if err != nil {
		http.Error(w, "Gagal menghapus file-file di server", http.StatusInternalServerError)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Gagal mengunggah file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileName := handler.Filename
	err = os.MkdirAll("uploads", os.ModePerm)
	if err != nil {
		http.Error(w, "Gagal membuat direktori unggahan", http.StatusInternalServerError)
		return
	}

	f, err := os.OpenFile(fmt.Sprintf("uploads/%s", fileName), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Gagal menyimpan file", http.StatusInternalServerError)
		return
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Gagal menyimpan file", http.StatusInternalServerError)
		return
	}

	// Proses ulang file yang diunggah dan menghasilkan data bridge dan SCC
	graph, graphBridge, err := processFile(fmt.Sprintf("uploads/%s", fileName))
	if err != nil {
		http.Error(w, "Gagal memproses file", http.StatusInternalServerError)
		return
	}

	// Membuat visualisasi
	visualizeGraph(graph, "uploads/graph.txt")
	visualizeBridges(graph, graphBridge, "uploads/bridge.txt")
	visualizeSCC(graph, "uploads/scc.txt")

	bridgeData := make([]string, 0, len(graphBridge.Bridges))
	for _, bridge := range graphBridge.Bridges {
		bridgeData = append(bridgeData, fmt.Sprintf("%s %s", bridge.Start.Name, bridge.End.Name))
	}

	sccData := make([]string, 0, len(graph.SCC))
	for _, scc := range graph.SCC {
		sccData = append(sccData, strings.Join(scc, " "))
	}

	graphURL := fmt.Sprintf("http://localhost:8080/images/graf.png?%d", time.Now().Unix())
	bridgesURL := fmt.Sprintf("http://localhost:8080/images/bridge.png?%d", time.Now().Unix())
	sccURL := fmt.Sprintf("http://localhost:8080/images/scc.png?%d", time.Now().Unix())

	response := map[string]interface{}{
		"bridges":    bridgeData,
		"scc":        sccData,
		"graphUrl":   graphURL,
		"bridgesUrl": bridgesURL,
		"sccUrl":     sccURL,
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func processFile(filename string) (*Graph, *GraphBridge, error) {
	graph := NewGraph()
	graphBridge := NewGraphBridge()

	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		vertices := strings.Fields(line)
		if len(vertices) != 2 {
			return nil, nil, fmt.Errorf("Format input salah: %s", line)
		}
		u := vertices[0]
		v := vertices[1]
		edges := strings.Split(line, " ")
		start := edges[0]
		end := edges[1]
		graphBridge.AddEdgeBridge(start, end)
		graph.AddEdge(u, v)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	graph.Tarjan()
	graphBridge.FindBridges()

	bridgeData := make([]string, 0, len(graphBridge.Bridges))
	for _, bridge := range graphBridge.Bridges {
		bridgeData = append(bridgeData, fmt.Sprintf("%s %s", bridge.Start.Name, bridge.End.Name))
	}

	sccData := make([]string, 0, len(graph.SCC))
	for _, scc := range graph.SCC {
		sccData = append(sccData, strings.Join(scc, " "))
	}

	return graph, graphBridge, nil
}

func main() {
	http.HandleFunc("/upload", handleUpload)
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("uploads"))))
	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
