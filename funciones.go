package grafo

import (
	TDACola "tdas/cola"
	TDADicc "tdas/diccionario"
	TDAHeap "tdas/heap"
	TDAPila "tdas/pila"
)

type distancia_vertice[K comparable] struct {
	vertice   K
	distancia int
}

type arista[K comparable] struct {
	origen  K
	destino K
	peso    int
}

func Bfs[K comparable](g Grafo[K], origen, destino K, con_destino bool) (TDADicc.Diccionario[K, K], TDADicc.Diccionario[K, int]) {
	var none K
	padres := TDADicc.CrearHash[K, K]()
	visitados := TDADicc.CrearHash[K, bool]()
	orden := TDADicc.CrearHash[K, int]()
	q := TDACola.CrearColaEnlazada[K]()
	q.Encolar(origen)
	padres.Guardar(origen, none)
	orden.Guardar(origen, 0)
	visitados.Guardar(origen, true)
	for !q.EstaVacia() {
		v := q.Desencolar()
		for _, w := range g.Adyacente(v) {
			if !visitados.Pertenece(w) {
				padres.Guardar(w, v)
				orden.Guardar(w, orden.Obtener(v)+1)
				visitados.Guardar(w, true)
				if w == destino && con_destino {
					return padres, orden
				}
				q.Encolar(w)
			}
		}
	}
	return padres, orden
}

func Dijkstra[K comparable](g Grafo[K], origen K) (TDADicc.Diccionario[K, K], TDADicc.Diccionario[K, int]) {
	var none K
	h := TDAHeap.CrearHeap[distancia_vertice[K]](func(k1, k2 distancia_vertice[K]) int { return k2.distancia - k1.distancia })
	padres := TDADicc.CrearHash[K, K]()
	distancias := TDADicc.CrearHash[K, int]()
	padres.Guardar(origen, none)
	distancias.Guardar(origen, 0)
	h.Encolar(distancia_vertice[K]{origen, 0})
	for !h.EstaVacia() {
		v := h.Desencolar()
		for _, w := range g.Adyacente(v.vertice) {
			nueva_dist := distancias.Obtener(v.vertice) + g.PesoArista(v.vertice, w)
			if !distancias.Pertenece(w) {
				distancias.Guardar(w, nueva_dist)
				padres.Guardar(w, v.vertice)
				h.Encolar(distancia_vertice[K]{w, nueva_dist})
			} else if nueva_dist < distancias.Obtener(w) {
				distancias.Guardar(w, nueva_dist)
				padres.Guardar(w, v.vertice)
				h.Encolar(distancia_vertice[K]{w, nueva_dist})
			}
		}
	}
	return padres, distancias
}

func Prim[K comparable](g Grafo[K]) Grafo[K] {
	origen := g.VerticeAleatorio()
	a := CrearGrafo[K](g.Dirigido())
	for _, v := range g.ObtenerVertices() {
		a.AgregarVertice(v)
	}
	h := TDAHeap.CrearHeap[arista[K]](func(a1, a2 arista[K]) int { return a2.peso - a1.peso })
	for _, v := range g.Adyacente(origen) {
		h.Encolar(arista[K]{origen, v, g.PesoArista(origen, v)})
	}
	visitados := TDADicc.CrearHash[K, bool]()
	visitados.Guardar(origen, true)
	for !h.EstaVacia() {
		e := h.Desencolar()
		if !visitados.Pertenece(e.destino) {
			a.AgregarArista(e.origen, e.destino, g.PesoArista(e.origen, e.destino))
			visitados.Guardar(e.destino, true)
			for _, u := range g.Adyacente(e.destino) {
				if !visitados.Pertenece(u) {
					h.Encolar(arista[K]{e.destino, u, g.PesoArista(e.destino, u)})
				}
			}
		}
	}
	return a
}

func OrdenTopologico[K comparable](g Grafo[K]) []K {
	if !g.Dirigido() {
		panic("El grafo no es dirigido.")
	}
	res := []K{}
	grados := gradoDeEntrada[K](g)
	q := TDACola.CrearColaEnlazada[K]()
	for i := grados.Iterador(); i.HaySiguiente(); i.Siguiente() {
		vertice, grado := i.VerActual()
		if grado == 0 {
			q.Encolar(vertice)
		}
	}
	for !q.EstaVacia() {
		v := q.Desencolar()
		for _, w := range g.Adyacente(v) {
			grado_anterior := grados.Obtener(w)
			grados.Guardar(w, grado_anterior-1)
			if grado_anterior-1 == 0 {
				res = append(res, w)
				q.Encolar(w)
			}
		}
	}
	return res
}

func gradoDeEntrada[K comparable](g Grafo[K]) TDADicc.Diccionario[K, int] {
	gr_entrada := TDADicc.CrearHash[K, int]()
	for _, v := range g.ObtenerVertices() {
		for _, w := range g.Adyacente(v) {
			if !gr_entrada.Pertenece(w) {
				gr_entrada.Guardar(w, 1)
			} else {
				grado_anterior := gr_entrada.Obtener(w)
				gr_entrada.Guardar(w, grado_anterior+1)
			}
		}
	}
	return gr_entrada
}

func ReconstruirCamino[K comparable](padres TDADicc.Diccionario[K, K], in, fin K) []K {
	p := TDAPila.CrearPilaDinamica[K]()
	v := fin
	res := []K{}
	for v != in {
		p.Apilar(v)
		v = padres.Obtener(v)
	}
	p.Apilar(in)
	for !p.EstaVacia() {
		res = append(res, p.Desapilar())
	}
	return res
}
