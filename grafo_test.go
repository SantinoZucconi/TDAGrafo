package grafo_test

import (
	"math/rand"
	TDAGrafo "tdas/grafo"
	"testing"

	"github.com/stretchr/testify/require"
)

const MEDIO int = 10
const VOLUMEN int = 125000

func Grafo1(g *TDAGrafo.GrafoNoPesado[string]) {
	(*g).AgregarVertice("A")
	(*g).AgregarVertice("B")
	(*g).AgregarVertice("C")
	(*g).AgregarAristaNP("A", "B")
	(*g).AgregarAristaNP("A", "C")
}

func GrafoA(g *TDAGrafo.GrafoPesado[string]) {
	(*g).AgregarVertice("A")
	(*g).AgregarVertice("B")
	(*g).AgregarVertice("C")
	(*g).AgregarArista("A", "B", 1)
	(*g).AgregarArista("A", "C", 2)
}

func TestGrafoVacio(t *testing.T) {
	t.Log("Creamos un grafo vacio y comprobamos que se comporte como tal")
	grafo := TDAGrafo.CrearGrafoPesado[string](false)

	require.Equal(t, 0, grafo.Cantidad())
	require.Equal(t, 0, len(grafo.ObtenerVertices()))
	require.False(t, grafo.EsVertice(""))

	require.PanicsWithValue(t, "El vertice no pertenece al grafo", func() { grafo.SacarVertice("A") })
	require.PanicsWithValue(t, "El vertice no pertenece al grafo", func() { grafo.HayArista("A", "B") })
	require.PanicsWithValue(t, "El vertice no pertenece al grafo", func() { grafo.SacarArista("A", "B") })
	require.PanicsWithValue(t, "El vertice no pertenece al grafo", func() { grafo.Adyacente("A") })
	require.PanicsWithValue(t, "El grafo no contiene vertices", func() { grafo.VerticeAleatorio() })

}

func TestGrafoDirigido(t *testing.T) {
	t.Log("Creamos un grafo no pesado dirigido, comprobando que los vertices y aristas se agreguen correectamente")
	grafo := TDAGrafo.CrearGrafoNoPesado[string](true)
	Grafo1(&grafo)

	require.Equal(t, 3, grafo.Cantidad())
	require.True(t, grafo.HayArista("A", "B"))
	require.True(t, grafo.HayArista("A", "C"))
	require.False(t, grafo.HayArista("B", "C"))
	require.False(t, grafo.HayArista("C", "B"))
	require.False(t, grafo.HayArista("B", "A"))
	require.False(t, grafo.HayArista("C", "A"))

}

func TestGrafoNoDirigido(t *testing.T) {
	t.Log("Creamos un grafo no pesado no dirigido, comprobando que los vertices y aristas se agreguen correectamente")

	grafo := TDAGrafo.CrearGrafoNoPesado[string](false)
	Grafo1(&grafo)

	require.True(t, grafo.HayArista("A", "B"))
	require.True(t, grafo.HayArista("A", "C"))
	require.False(t, grafo.HayArista("B", "C"))
	require.False(t, grafo.HayArista("C", "B"))
	require.True(t, grafo.HayArista("B", "A"))
	require.True(t, grafo.HayArista("C", "A"))

}

func TestGrafoPesado(t *testing.T) {
	t.Log("Creamos un grafo pesado dirigido, comprobando que los vertices y aristas se agreguen correectamente")
	grafo := TDAGrafo.CrearGrafoPesado[string](false)
	GrafoA(&grafo)

	require.True(t, grafo.HayArista("A", "B"))
	require.True(t, grafo.HayArista("A", "C"))
	require.False(t, grafo.HayArista("B", "C"))
	require.False(t, grafo.HayArista("C", "B"))
	require.True(t, grafo.HayArista("B", "A"))
	require.True(t, grafo.HayArista("C", "A"))

	require.Equal(t, 1, grafo.PesoArista("A", "B"))
	require.Equal(t, 2, grafo.PesoArista("A", "C"))
	require.PanicsWithValue(t, "La arista no pertenece al grafo", func() { grafo.PesoArista("B", "C") })

}

func TestVerticeAleatorio(t *testing.T) {
	t.Log("Comprobamos que al obtener un vertice aleatorio del grado, este efectivamente pertenezca al grafo")

	grafo := TDAGrafo.CrearGrafoNoPesado[int](false)

	for i := 0; i < MEDIO; i++ {
		grafo.AgregarVertice(i)
	}

	verticeAleatorio := grafo.VerticeAleatorio()
	require.True(t, grafo.EsVertice(verticeAleatorio))
}

func TestVertices(t *testing.T) {
	t.Log("Agregamos y borramos vertices comprobando que siempre este con la cantidad correcta")
	grafo := TDAGrafo.CrearGrafoNoPesado[int](false)

	grafo.AgregarVertice(1)
	grafo.AgregarVertice(2)
	grafo.AgregarVertice(3)
	require.Equal(t, 3, grafo.Cantidad())
	require.True(t, grafo.EsVertice(1))
	require.True(t, grafo.EsVertice(2))
	require.True(t, grafo.EsVertice(3))

	grafo.SacarVertice(2)
	require.Equal(t, 2, grafo.Cantidad())
	require.True(t, grafo.EsVertice(1))
	require.False(t, grafo.EsVertice(2))
	require.True(t, grafo.EsVertice(3))

	grafo.SacarVertice(1)
	require.Equal(t, 1, grafo.Cantidad())
	require.False(t, grafo.EsVertice(1))
	require.False(t, grafo.EsVertice(2))
	require.True(t, grafo.EsVertice(3))

	grafo.SacarVertice(3)
	require.Equal(t, 0, grafo.Cantidad())
	require.False(t, grafo.EsVertice(1))
	require.False(t, grafo.EsVertice(2))
	require.False(t, grafo.EsVertice(3))
}

func TestAristas(t *testing.T) {
	t.Log("Agregamos y borramos aristas comprobando que siempre este con la cantidad correcta")
	grafo := TDAGrafo.CrearGrafoPesado[string](true)
	GrafoA(&grafo)

	grafo.SacarArista("A", "B")
	require.True(t, grafo.HayArista("A", "C"))
	require.False(t, grafo.HayArista("C", "A"))
	require.Equal(t, 2, grafo.PesoArista("A", "C"))
	require.False(t, grafo.HayArista("A", "B"))
	require.False(t, grafo.HayArista("B", "C"))

	grafo.SacarArista("A", "C")
	require.False(t, grafo.HayArista("A", "C"))
	require.False(t, grafo.HayArista("A", "B"))
	require.False(t, grafo.HayArista("B", "C"))

}

func TestVolumen(t *testing.T) {
	t.Log("Agregamos vertices y aristas en volumen comprobando que se agreguen correctamente")
	grafo := TDAGrafo.CrearGrafoNoPesado[int](true)
	cantAdyacentes := make([]int, 0)

	for i := 0; i < VOLUMEN; i++ {
		grafo.AgregarVertice(i)
	}

	for i := 0; i < VOLUMEN; i++ {
		adyacentes := rand.Intn(5)
		cantAdyacentes = append(cantAdyacentes, adyacentes)
		j := 0
		for j < adyacentes {
			if !grafo.HayArista(i, j) {
				grafo.AgregarAristaNP(i, j)
				j++
			}
		}
	}

	for i := 0; i < VOLUMEN; i++ {
		require.True(t, grafo.EsVertice(i))
		require.Equal(t, len(grafo.Adyacente(i)), cantAdyacentes[i])
	}

}

/*
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
*/
