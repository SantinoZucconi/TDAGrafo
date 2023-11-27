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

func CrearGrafo[K comparable](dirigido bool) Grafo[K] {
	return &grafo[K]{vertices: TDADICC.CrearHash[K, K](), adyacencias: TDADICC.CrearHash[K, *TDADICC.Diccionario[K, int]](), dirigido: dirigido}
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
	panicSiNoPerteneceVertice[K](v, g)
	panicSiNoPerteneceVertice[K](w, g)
	(*g.adyacencias.Obtener(v)).Guardar(w, peso)
	if !g.dirigido {
		(*g.adyacencias.Obtener(w)).Guardar(v, peso)
	}
}

func (g *grafo[K]) EstanUnidos(v, w K) bool {
	panicSiNoPerteneceVertice[K](v, g)
	panicSiNoPerteneceVertice[K](w, g)
	if g.dirigido {
		return g.adyacencias.Pertenece(v) && (*g.adyacencias.Obtener(v)).Pertenece(w)
	} else {
		return g.adyacencias.Pertenece(v) && (*g.adyacencias.Obtener(v)).Pertenece(w) || g.adyacencias.Pertenece(w) && (*g.adyacencias.Obtener(w)).Pertenece(v)
	}
}

func (g *grafo[K]) Adyacente(v K) []K {
	panicSiNoPerteneceVertice[K](v, g)
	res := []K{}
	for iter := (*g.adyacencias.Obtener(v)).Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		adyacente, _ := iter.VerActual()
		res = append(res, adyacente)
	}
	return res
}

func (g *grafo[K]) SacarArista(v, w K) {
	panicSiNoPerteneceArista[K](v, w, g)
	(*g.adyacencias.Obtener(v)).Borrar(w)
	if !g.dirigido {
		panicSiNoPerteneceArista[K](w, v, g)
		(*g.adyacencias.Obtener(w)).Borrar(v)
	}
}

func (g *grafo[K]) SacarVertice(v K) {
	panicSiNoPerteneceVertice[K](v, g)
	if !g.adyacencias.Pertenece(v) {
		panic("El vertice no pertenece al grafo")
	}
	g.adyacencias.Borrar(v)
}

func (g *grafo[K]) PesoArista(v, w K) int {
	panicSiNoPerteneceArista[K](v, w, g)
	return (*g.adyacencias.Obtener(v)).Obtener(w)
}

func (g *grafo[K]) ObtenerVertices() []K {
	res := []K{}
	for i := g.vertices.Iterador(); i.HaySiguiente(); i.Siguiente() {
		v, _ := i.VerActual()
		res = append(res, v)
	}
	return res
}

func (g *grafo[K]) VerticeAleatorio() K {
	vertices := g.ObtenerVertices()
	num := rand.Intn(len(vertices))
	return vertices[num]
}

func panicSiNoPerteneceArista[K comparable](v, w K, g *grafo[K]) {
	if !g.vertices.Pertenece(v) {
		panic("El vertice no pertenece al grafo")
	}
	if !(*g.adyacencias.Obtener(v)).Pertenece(w) {
		panic("La arista no pertenece al grafo")
	}
}

func panicSiNoPerteneceVertice[K comparable](v K, g *grafo[K]) {
	if !g.vertices.Pertenece(v) {
		panic("El vertice no pertenece al grafo")
	}
}

func (g *grafo[K]) Dirigido() bool {
	return g.dirigido
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
