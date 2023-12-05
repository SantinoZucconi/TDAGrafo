package grafo

import (
	"fmt"
	"math/rand"
	TDADICC "tdas/diccionario"
)

type grafo[K comparable] struct {
	vertices    TDADICC.Diccionario[K, K]
	adyacencias TDADICC.Diccionario[K, *TDADICC.Diccionario[K, int]]
	dirigido    bool
}

func CrearGrafoPesado[K comparable](dirigido bool) GrafoPesado[K] {
	return &grafo[K]{vertices: TDADICC.CrearHash[K, K](), adyacencias: TDADICC.CrearHash[K, *TDADICC.Diccionario[K, int]](), dirigido: dirigido}
}

func CrearGrafoNoPesado[K comparable](dirigido bool) GrafoNoPesado[K] {
	return &grafo[K]{vertices: TDADICC.CrearHash[K, K](), adyacencias: TDADICC.CrearHash[K, *TDADICC.Diccionario[K, int]](), dirigido: dirigido}

}

func (g *grafo[K]) panicSiNoPerteneceVertice(v K) {
	if !g.vertices.Pertenece(v) {
		panic("El vertice no pertenece al grafo")
	}
}

func (g *grafo[K]) panicSiNoPerteneceArista(v, w K) {
	g.panicSiNoPerteneceVertice(v)
	if !(*g.adyacencias.Obtener(v)).Pertenece(w) {
		panic("La arista no pertenece al grafo")
	}
}

func (g *grafo[K]) EsVertice(v K) bool {
	return g.vertices.Pertenece(v)
}

func (g *grafo[K]) AgregarVertice(vertice K) {
	if g.vertices.Pertenece(vertice) {
		panic("El vertice ya pertenece al grafo")
	}
	g.vertices.Guardar(vertice, vertice)
	d := TDADICC.CrearHash[K, int]()
	g.adyacencias.Guardar(vertice, &d)
}

func (g *grafo[K]) AgregarArista(v, w K, peso int) {
	g.panicSiNoPerteneceVertice(v)
	g.panicSiNoPerteneceVertice(w)
	(*g.adyacencias.Obtener(v)).Guardar(w, peso)
	if !g.dirigido {
		(*g.adyacencias.Obtener(w)).Guardar(v, peso)
	}
}

func (g *grafo[K]) AgregarAristaNP(v, w K) {
	g.AgregarArista(v, w, 1)
}

func (g *grafo[K]) HayArista(v, w K) bool {
	g.panicSiNoPerteneceVertice(v)
	g.panicSiNoPerteneceVertice(w)
	return (*g.adyacencias.Obtener(v)).Pertenece(w)
}

func (g *grafo[K]) Adyacente(v K) []K {
	g.panicSiNoPerteneceVertice(v)
	res := []K{}
	for iter := (*g.adyacencias.Obtener(v)).Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		adyacente, _ := iter.VerActual()
		res = append(res, adyacente)
	}
	return res
}

func (g *grafo[K]) SacarArista(v, w K) {
	g.panicSiNoPerteneceArista(v, w)
	(*g.adyacencias.Obtener(v)).Borrar(w)
	if !g.dirigido {
		(*g.adyacencias.Obtener(w)).Borrar(v)
	}
}

func (g *grafo[K]) SacarVertice(v K) {
	g.panicSiNoPerteneceVertice(v)
	g.vertices.Borrar(v)
	g.adyacencias.Borrar(v)

	for iterV := g.adyacencias.Iterador(); iterV.HaySiguiente(); iterV.Siguiente() {
		_, dicW := iterV.VerActual()
		if (*dicW).Pertenece(v) {
			(*dicW).Borrar(v)
		}
	}
}

func (g *grafo[K]) PesoArista(v, w K) int {
	g.panicSiNoPerteneceArista(v, w)
	return (*g.adyacencias.Obtener(v)).Obtener(w)
}

func (g *grafo[K]) ObtenerVertices() []K {
	res := []K{}
	for iter := g.vertices.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		v, _ := iter.VerActual()
		res = append(res, v)
	}
	return res
}

func (g *grafo[K]) VerticeAleatorio() K {
	if g.Cantidad() == 0 {
		panic("El grafo no contiene vertices")
	}
	vertices := g.ObtenerVertices()
	num := rand.Intn(len(vertices))
	return vertices[num]
}

func (g *grafo[K]) Dirigido() bool {
	return g.dirigido
}

func (g *grafo[K]) Cantidad() int {
	return g.vertices.Cantidad()
}

func (g *grafo[K]) Imprimir() {
	fmt.Println("Vertices:")
	fmt.Println(g.ObtenerVertices())
	fmt.Println("Aristas:")
	for i := g.adyacencias.Iterador(); i.HaySiguiente(); i.Siguiente() {
		vertice, ady := i.VerActual()
		for k := (*ady).Iterador(); k.HaySiguiente(); k.Siguiente() {
			adyacente, peso := k.VerActual()
			fmt.Println(vertice, "<--->", adyacente, "peso:", peso)
		}
	}
}
