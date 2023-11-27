package grafo_test

import (
	TDAGrafo "tdas/grafo"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGrafoDirigido(t *testing.T) {
	grafo := TDAGrafo.CrearGrafo[string](true)
	grafo.AgregarVertice("a")
	grafo.AgregarVertice("b")
	grafo.AgregarArista("a", "b", 1)
	require.Equal(t, true, grafo.EstanUnidos("a", "b"))
	require.Equal(t, false, grafo.EstanUnidos("b", "a"))
}

func TestGrafoNoDirigido(t *testing.T) {
	grafo := TDAGrafo.CrearGrafo[string](false)
	grafo.AgregarVertice("a")
	grafo.AgregarVertice("b")
	grafo.AgregarArista("a", "b", 1)
	require.Equal(t, true, grafo.EstanUnidos("a", "b"))
	require.Equal(t, true, grafo.EstanUnidos("b", "a"))
}

func TestGrafoNoDirigido2(t *testing.T) {
	grafo := TDAGrafo.CrearGrafo[string](false)
	grafo.AgregarVertice("a")
	grafo.AgregarVertice("b")
	grafo.AgregarVertice("c")
	grafo.AgregarArista("a", "b", 1)
	require.Equal(t, true, grafo.EstanUnidos("a", "b"))
	require.Equal(t, true, grafo.EstanUnidos("b", "a"))
	require.Equal(t, grafo.PesoArista("a", "b"), 1)
	grafo.SacarArista("a", "b")
	require.PanicsWithValue(t, "La arista no pertenece al grafo", func() { grafo.PesoArista("a", "b") })
	require.Equal(t, false, grafo.EstanUnidos("a", "b"))
	require.Equal(t, false, grafo.EstanUnidos("b", "a"))
}

func TestRecorrido(t *testing.T) {
	g := TDAGrafo.CrearGrafo[string](false)
	g.AgregarVertice("a")
	g.AgregarVertice("b")
	g.AgregarVertice("c")
	g.AgregarVertice("d")
	g.AgregarVertice("e")
	g.AgregarArista("a", "b", 1)
	g.AgregarArista("a", "c", 1)
	g.AgregarArista("b", "d", 1)
	g.AgregarArista("d", "e", 1)
	padres, ordenes := TDAGrafo.Bfs[string](g, "a", "", false)
	require.Equal(t, "", padres.Obtener("a"))
	require.Equal(t, 0, ordenes.Obtener("a"))
	require.Equal(t, "a", padres.Obtener("b"))
	require.Equal(t, 1, ordenes.Obtener("b"))
	require.Equal(t, "a", padres.Obtener("c"))
	require.Equal(t, 1, ordenes.Obtener("c"))
	require.Equal(t, "b", padres.Obtener("d"))
	require.Equal(t, 2, ordenes.Obtener("d"))
	require.Equal(t, "d", padres.Obtener("e"))
	require.Equal(t, 3, ordenes.Obtener("e"))
}

func TestDijkstra(t *testing.T) {
	g := TDAGrafo.CrearGrafo[string](false)
	g.AgregarVertice("a")
	g.AgregarVertice("b")
	g.AgregarVertice("c")
	g.AgregarVertice("d")
	g.AgregarArista("a", "b", 1)
	g.AgregarArista("a", "c", 1)
	g.AgregarArista("b", "d", 4)
	g.AgregarArista("c", "d", 2)
	padres, distancias := TDAGrafo.Dijkstra[string](g, "a")
	require.Equal(t, "", padres.Obtener("a"))
	require.Equal(t, 0, distancias.Obtener("a"))
	require.Equal(t, "a", padres.Obtener("b"))
	require.Equal(t, 1, distancias.Obtener("b"))
	require.Equal(t, "a", padres.Obtener("c"))
	require.Equal(t, 1, distancias.Obtener("c"))
	require.Equal(t, "c", padres.Obtener("d"))
	require.Equal(t, 3, distancias.Obtener("d"))
}

func TestPrim(t *testing.T) {
	g := TDAGrafo.CrearGrafo[string](false)
	g.AgregarVertice("a")
	g.AgregarVertice("b")
	g.AgregarVertice("c")
	g.AgregarVertice("d")
	g.AgregarVertice("e")
	g.AgregarVertice("f")
	g.AgregarVertice("g")
	g.AgregarVertice("h")
	g.AgregarVertice("i")
	g.AgregarArista("b", "a", 5)
	g.AgregarArista("b", "f", 1)
	g.AgregarArista("f", "a", 8)
	g.AgregarArista("f", "d", 6)
	g.AgregarArista("d", "c", 8)
	g.AgregarArista("a", "c", 7)
	g.AgregarArista("a", "g", 4)
	g.AgregarArista("a", "e", 2)
	g.AgregarArista("g", "e", 3)
	g.AgregarArista("e", "i", 3)
	g.AgregarArista("e", "c", 3)
	g.AgregarArista("c", "h", 3)
	g.AgregarArista("h", "i", 2)
	g.AgregarArista("d", "h", 6)
	a := TDAGrafo.Prim[string](g)
	a.Imprimir()
	require.Equal(t, false, a.EstanUnidos("a", "f"))
	require.Equal(t, false, a.EstanUnidos("a", "c"))
	require.Equal(t, false, a.EstanUnidos("d", "c"))
	require.Equal(t, false, a.EstanUnidos("a", "g"))
	require.Equal(t, false, a.EstanUnidos("d", "f"))
	require.Equal(t, false, a.EstanUnidos("c", "h"))
}
