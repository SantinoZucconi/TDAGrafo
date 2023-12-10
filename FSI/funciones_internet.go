package funciones_sobre_internet

import (
	TDACola "tdas/cola"
	TDAHeap "tdas/cola_prioridad"
	TDADicc "tdas/diccionario"
	TDAGrafo "tdas/grafo"
	TDALista "tdas/lista"
	TDAPila "tdas/pila"
	ERROR "tp3/errores"
)

const d float64 = 0.85
const MAX_ITERACIONES int = 300
const MIN_DIFPORCENTUAL float64 = 1

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func difAbs(x, y float64) float64 {
	if x > y {
		return x - y
	}

	return y - x
}

func difPorcentual(original, nuevo float64) float64 {
	return (difAbs(original, nuevo) / original) * 100
}

type verticePR[K comparable] struct {
	vertice  K
	pageRank float64
}

func BFS[K comparable](g TDAGrafo.GrafoNoPesado[K], origen K, cond func(K, TDADicc.Diccionario[K, int]) bool) (TDADicc.Diccionario[K, K], TDADicc.Diccionario[K, int]) {
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
				if cond(w, orden) {
					return padres, orden
				}
				q.Encolar(w)
			}
		}
	}
	return padres, orden
}

func _dfs[K comparable](g TDAGrafo.GrafoNoPesado[K], vertice, padre K, padres TDADicc.Diccionario[K, K], orden TDADicc.Diccionario[K, int], visitados TDADicc.Diccionario[K, bool]) {
	var none K
	if padre != none {
		padres.Guardar(vertice, padre)
		orden.Guardar(vertice, orden.Obtener(padre)+1)
	}
	visitados.Guardar(vertice, true)

	adyacentes := g.Adyacente(vertice)
	cantAdyacentes := len(adyacentes)

	for i := 0; i < cantAdyacentes; i++ {
		w := adyacentes[i]
		if !visitados.Obtener(w) {
			_dfs(g, w, vertice, padres, orden, visitados)
		}
	}
}

func DFS[K comparable](g TDAGrafo.GrafoNoPesado[K], origen K) (TDADicc.Diccionario[K, K], TDADicc.Diccionario[K, int]) {
	var none K
	padres := TDADicc.CrearHash[K, K]()
	visitados := TDADicc.CrearHash[K, bool]()
	orden := TDADicc.CrearHash[K, int]()

	padres.Guardar(origen, none)
	orden.Guardar(origen, 0)
	visitados.Guardar(origen, true)

	_dfs(g, origen, none, padres, orden, visitados)

	return padres, orden
}

func GradoDeEntrada[K comparable](g TDAGrafo.GrafoNoPesado[K]) TDADicc.Diccionario[K, int] {
	grEntrada := TDADicc.CrearHash[K, int]()
	for _, v := range g.ObtenerVertices() {
		grEntrada.Guardar(v, 0)
	}
	for _, v := range g.ObtenerVertices() {
		for _, w := range g.Adyacente(v) {
			grado_anterior := grEntrada.Obtener(w)
			grEntrada.Guardar(w, grado_anterior+1)
		}
	}
	return grEntrada
}

func GradoDeSalida[K comparable](g TDAGrafo.GrafoNoPesado[K]) TDADicc.Diccionario[K, int] {
	grSalida := TDADicc.CrearHash[K, int]()
	for _, v := range g.ObtenerVertices() {
		cantAdyacentes := len(g.Adyacente(v))
		grSalida.Guardar(v, cantAdyacentes)
	}

	return grSalida
}

func ReconstruirCamino[K comparable](padres TDADicc.Diccionario[K, K], in, fin K) (TDALista.Lista[K], error) {
	lista := TDALista.CrearListaEnlazada[K]()

	if !padres.Pertenece(fin) {
		return lista, &ERROR.ErrorNoExisteRecorrido{}
	}
	v := fin
	for v != in {
		lista.InsertarPrimero(v)
		v = padres.Obtener(v)
	}
	lista.InsertarPrimero(v)
	return lista, nil
}

////

func CaminoMinimo[K comparable](g TDAGrafo.GrafoNoPesado[K], origen, destino K) (TDALista.Lista[K], error) {
	padres, _ := BFS(g, origen, func(vertice K, orden TDADicc.Diccionario[K, int]) bool {
		return vertice == destino
	})
	camino, err := ReconstruirCamino(padres, origen, destino)
	return camino, err
}

func TodosEnRango[K comparable](g TDAGrafo.GrafoNoPesado[K], origen K, rango int) int {
	_, orden := BFS(g, origen, func(vertice K, orden TDADicc.Diccionario[K, int]) bool {
		return orden.Obtener(vertice) > rango
	})

	contador := 0
	for iter := orden.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		_, n := iter.VerActual()
		if n == rango {
			contador++
		}
	}

	return contador
}

func NavegacionPrimerLink[K comparable](g TDAGrafo.GrafoNoPesado[K], primerosLinks TDADicc.Diccionario[K, K], inicio K) TDALista.Lista[K] {
	v := inicio
	contador := 0
	hayLinks := true
	camino := TDALista.CrearListaEnlazada[K]()
	camino.InsertarUltimo(inicio)
	contador++

	for contador < 21 && hayLinks {
		if primerosLinks.Pertenece(v) {
			v = primerosLinks.Obtener(v)
			camino.InsertarUltimo(v)
			contador++
		} else {
			hayLinks = false
		}
	}

	return camino
}

func Diametro[K comparable](g TDAGrafo.GrafoNoPesado[K]) TDALista.Lista[K] {
	diametro := 0
	var verticeInicio, verticeMasLejos K
	var padresDiametro TDADicc.Diccionario[K, K]

	vertices := g.ObtenerVertices()
	for i := 0; i < g.Cantidad(); i++ {
		v := vertices[i]
		padres, orden := BFS(g, v, func(vertice K, orden TDADicc.Diccionario[K, int]) bool {
			return false
		})

		for iter := orden.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
			w, n := iter.VerActual()
			if n > diametro {
				verticeInicio = v
				verticeMasLejos = w
				diametro = n
				padresDiametro = padres
			}
		}
	}

	camino, _ := ReconstruirCamino(padresDiametro, verticeInicio, verticeMasLejos)

	return camino
}

////

func _dfs_cfc[K comparable](g TDAGrafo.GrafoNoPesado[K], v K, visitados, apilados TDADicc.Diccionario[K, K], orden, masbajo TDADicc.Diccionario[K, int], pila TDAPila.Pila[K], cfc *[]TDADicc.Diccionario[K, K], contador_global *int) {
	visitados.Guardar(v, v)
	pila.Apilar(v)
	apilados.Guardar(v, v)
	orden.Guardar(v, *contador_global)
	masbajo.Guardar(v, *contador_global)
	*contador_global++

	adyacentes := g.Adyacente(v)
	for i := 0; i < len(adyacentes); i++ {
		w := adyacentes[i]
		if !visitados.Pertenece(w) {
			_dfs_cfc(g, w, visitados, apilados, orden, masbajo, pila, cfc, contador_global)
		}
		if apilados.Pertenece(w) {
			masbajo.Guardar(v, min(masbajo.Obtener(v), masbajo.Obtener(w)))
		}
	}

	if orden.Obtener(v) == masbajo.Obtener(v) {
		finCFC := false
		nueva_cfc := TDADicc.CrearHash[K, K]()
		for !finCFC {
			max := pila.Desapilar()
			apilados.Borrar(max)
			if max == v {
				finCFC = true
			}
			nueva_cfc.Guardar(max, max)
		}
		*cfc = append(*cfc, nueva_cfc)
	}

}

func CFC[K comparable](g TDAGrafo.GrafoNoPesado[K]) ([]TDALista.Lista[K], TDADicc.Diccionario[K, int]) {
	var v K

	orden := TDADicc.CrearHash[K, int]()
	masbajo := TDADicc.CrearHash[K, int]()
	visitados := TDADicc.CrearHash[K, K]()
	pila := TDAPila.CrearPilaDinamica[K]()
	apilados := TDADicc.CrearHash[K, K]()

	cfc := make([]TDADicc.Diccionario[K, K], 0)
	pertenencia := TDADicc.CrearHash[K, int]()

	listas := make([]TDALista.Lista[K], 0)

	cont := 0
	contador_global := &cont

	vertices := g.ObtenerVertices()
	for i := 0; i < g.Cantidad(); i++ {
		v = vertices[i]
		if !visitados.Pertenece(v) {
			_dfs_cfc(g, v, visitados, apilados, orden, masbajo, pila, &cfc, contador_global)
		}
	}

	for i := 0; i < g.Cantidad(); i++ {
		v = vertices[i]
		for d := 0; d < len(cfc); d++ {
			comp := cfc[d]
			if comp.Pertenece(vertices[i]) {
				pertenencia.Guardar(v, d)
			}
		}
	}

	for d := 0; d < len(cfc); d++ {
		lista := TDALista.CrearListaEnlazada[K]()
		comp := cfc[d]
		comp.Iterar(func(clave, dato K) bool {
			lista.InsertarUltimo(clave)
			return true
		})
		listas = append(listas, lista)
	}

	return listas, pertenencia
}

func Comunidades[K comparable](g TDAGrafo.GrafoNoPesado[K]) (TDADicc.Diccionario[K, int], TDADicc.Diccionario[int, TDALista.Lista[K]]) {
	label := TDADicc.CrearHash[K, int]()
	var comunidad TDALista.Lista[K]
	_comunidades := TDADicc.CrearHash[int, TDALista.Lista[K]]()
	vEntradas := vertices_entrada[K](g)
	for i, v := range g.ObtenerVertices() {
		label.Guardar(v, i)
	}
	for i := 0; i < 4; i++ {
		for _, v := range g.ObtenerVertices() {
			labelNeighbor := []int{}
			for _, w := range g.Adyacente(v) {
				labelNeighbor = append(labelNeighbor, label.Obtener(w))
			}
			if g.Dirigido() {
				if vEntradas.Pertenece(v) {
					colaAdy := vEntradas.Obtener(v)
					for !colaAdy.EstaVacia() {
						labelNeighbor = append(labelNeighbor, label.Obtener(colaAdy.BorrarPrimero()))
					}
				}
			}
			newLabel := max_freq(labelNeighbor)
			label.Guardar(v, newLabel)
		}
	}
	label.Iterar(func(clave K, dato int) bool {
		if !_comunidades.Pertenece(dato) {
			comunidad = TDALista.CrearListaEnlazada[K]()
			comunidad.InsertarPrimero(clave)
			_comunidades.Guardar(dato, comunidad)
		} else {
			comunidad = _comunidades.Obtener(dato)
			comunidad.InsertarUltimo(clave)
			_comunidades.Guardar(dato, comunidad)
		}
		return true
	})
	return label, _comunidades
}

func max_freq(arr []int) int {
	maxFreqLabel := TDADicc.CrearHash[int, int]()
	var max int
	var maxLabel int
	for _, label := range arr {
		if !maxFreqLabel.Pertenece(label) {
			maxFreqLabel.Guardar(label, 0)
		}
		maxFreqLabel.Guardar(label, maxFreqLabel.Obtener(label)+1)
	}
	maxFreqLabel.Iterar(func(clave, dato int) bool {
		if dato > max {
			max = dato
			maxLabel = clave
		}
		return true
	})
	return maxLabel
}

func vertices_entrada[K comparable](g TDAGrafo.GrafoNoPesado[K]) TDADicc.Diccionario[K, TDALista.Lista[K]] {
	verticesEntrada := TDADicc.CrearHash[K, TDALista.Lista[K]]()
	for _, v := range g.ObtenerVertices() {
		verticesEntrada.Guardar(v, TDALista.CrearListaEnlazada[K]())
	}

	for _, v := range g.ObtenerVertices() {
		for _, w := range g.Adyacente(v) {
			verticesEntrada.Obtener(w).InsertarUltimo(v)
		}
	}
	return verticesEntrada
}

func ClusteringIndividual[K comparable](g TDAGrafo.GrafoNoPesado[K], vertice K) float64 {
	adyacentes := g.Adyacente(vertice)
	cantAdyacentes := len(adyacentes)
	unionAdyacentes := 0

	if cantAdyacentes < 2 {
		return 0
	}

	for j := 0; j < cantAdyacentes; j++ {
		w1 := adyacentes[j]
		for k := 0; k < cantAdyacentes; k++ {
			w2 := adyacentes[k]
			if w1 != vertice && w1 != w2 && g.HayArista(w1, w2) {
				unionAdyacentes++
			}
		}
	}

	return float64(unionAdyacentes) / float64(cantAdyacentes*(cantAdyacentes-1))
}

func ClusteringRed[K comparable](g TDAGrafo.GrafoNoPesado[K]) float64 {
	vertices := g.ObtenerVertices()
	cantidadVertices := g.Cantidad()
	var total float64

	for i := 0; i < cantidadVertices; i++ {
		v := vertices[i]
		clust := ClusteringIndividual(g, v)
		total += clust
	}

	return total / float64(cantidadVertices)
}

////

func PageRank[K comparable](g TDAGrafo.GrafoNoPesado[K]) TDALista.Lista[K] {
	vertices := g.ObtenerVertices()
	cantidadVertices := g.Cantidad()
	pageRanks := TDADicc.CrearHash[K, []float64]()
	entradas := vertices_entrada(g)
	gradoSalidas := GradoDeSalida(g)
	pageOrdenadas := TDALista.CrearListaEnlazada[K]()
	contador := 0

	heap := TDAHeap.CrearHeap[verticePR[K]](func(v1, v2 verticePR[K]) int {
		comp := v1.pageRank - v2.pageRank
		if comp < 0 {
			return -1
		} else if comp > 0 {
			return 1
		}
		return 0
	})

	for i := 0; i < cantidadVertices; i++ {
		v := vertices[i]
		pageRanks.Guardar(v, []float64{1})
	}

	seguirIterando := true

	for seguirIterando && contador < MAX_ITERACIONES {
		seguirIterando = false
		for k := 0; k < cantidadVertices; k++ {
			v := vertices[k]
			prNuevo := _arrastrePR(gradoSalidas, v, cantidadVertices, pageRanks, contador, entradas.Obtener(v))
			if difPorcentual(pageRanks.Obtener(v)[contador], prNuevo) > MIN_DIFPORCENTUAL && !seguirIterando {
				seguirIterando = true
			}
			prPagina := pageRanks.Obtener(v)
			prPagina = append(prPagina, prNuevo)
			pageRanks.Guardar(v, prPagina)
		}

		contador++
	}

	pageRanks.Iterar(func(clave K, dato []float64) bool {
		pr := dato[contador]
		heap.Encolar(verticePR[K]{clave, pr})
		return true
	})

	for !heap.EstaVacia() {
		pageOrdenadas.InsertarUltimo(heap.Desencolar().vertice)
	}

	return pageOrdenadas
}

func _arrastrePR[K comparable](gradoSalidas TDADicc.Diccionario[K, int], vertice K, N int, pageRanks TDADicc.Diccionario[K, []float64], pos int, entradas TDALista.Lista[K]) float64 {
	var primerTermino, segundoTermino float64
	primerTermino = (1 - d) / float64(N)
	st := &segundoTermino

	entradas.Iterar(func(k K) bool {
		cantAdyacentes := gradoSalidas.Obtener(k)
		*st += pageRanks.Obtener(k)[pos] / float64(cantAdyacentes)
		return true
	})

	segundoTermino *= d

	return primerTermino + segundoTermino
}

func OrdenLectura[K comparable](grafo TDAGrafo.GrafoNoPesado[K], paginas []K) (TDALista.Lista[K], error) {
	dicc := TDADicc.CrearHash[K, bool]()
	subgrafo := TDAGrafo.CrearGrafoNoPesado[K](true)
	for _, pagina := range paginas {
		dicc.Guardar(pagina, true)
		subgrafo.AgregarVertice(pagina)
	}
	for _, pagina := range paginas {
		lista := grafo.Adyacente(pagina)
		for _, link := range lista {
			if dicc.Pertenece(link) {
				subgrafo.AgregarAristaNP(link, pagina)
			}
		}
	}
	return OrdenTopologico[K](subgrafo)
}

func OrdenTopologico[K comparable](g TDAGrafo.GrafoNoPesado[K]) (TDALista.Lista[K], error) {
	res := TDALista.CrearListaEnlazada[K]()
	if !g.Dirigido() {
		return res, ERROR.ErrorNoExisteOrden{}
	}
	grados := GradoDeEntrada[K](g)
	q := TDACola.CrearColaEnlazada[K]()
	for iter := grados.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		vertice, grado := iter.VerActual()
		if grado == 0 {
			q.Encolar(vertice)
			res.InsertarUltimo(vertice)
		}
	}
	for !q.EstaVacia() {
		v := q.Desencolar()
		for _, w := range g.Adyacente(v) {
			grado_anterior := grados.Obtener(w)
			grados.Guardar(w, grado_anterior-1)
			if grado_anterior-1 == 0 {
				res.InsertarUltimo(w)
				q.Encolar(w)
			}
		}
	}
	if res.Largo() != len(g.ObtenerVertices()) {
		return res, ERROR.ErrorNoExisteOrden{}
	}
	return res, nil
}

func CicloN[K comparable](g TDAGrafo.GrafoNoPesado[K], pagina K, saltos int) (TDALista.Lista[K], error) {
	distancia := TDADicc.CrearHash[K, int]()
	visitados := TDADicc.CrearHash[K, bool]()
	distancia.Guardar(pagina, 0)
	subgrafo := TDAGrafo.CrearGrafoNoPesado[K](true)
	q := TDACola.CrearColaEnlazada[K]()
	q.Encolar(pagina)
	for !q.EstaVacia() {
		v := q.Desencolar()
		if !visitados.Pertenece(v) {
			visitados.Guardar(v, true)
			subgrafo.AgregarVertice(v)
		}
		for _, w := range g.Adyacente(v) {
			distancia_w := distancia.Obtener(v) + 1
			if !visitados.Pertenece(w) && distancia_w < saltos {
				visitados.Guardar(w, true)
				subgrafo.AgregarVertice(w)
			}
			if distancia_w == saltos && w == pagina || distancia_w < saltos {
				subgrafo.AgregarAristaNP(v, w)
			}
			if distancia_w <= saltos {
				distancia.Guardar(w, distancia_w)
				q.Encolar(w)
			}
		}
	}
	return dfs_cicloN[K](subgrafo, pagina, saltos)
}

func dfs_cicloN[K comparable](g TDAGrafo.Grafo[K], origen K, dist int) (TDALista.Lista[K], error) {
	var none K
	padres := TDADicc.CrearHash[K, K]()
	visitados := TDADicc.CrearHash[K, bool]()
	padres.Guardar(origen, none)
	return _dfs_cicloN_aux[K](g, 1, dist, origen, origen, padres, visitados)
}

func _dfs_cicloN_aux[K comparable](g TDAGrafo.Grafo[K], contador, n int, v, destino K, padres TDADicc.Diccionario[K, K], visitados TDADicc.Diccionario[K, bool]) (TDALista.Lista[K], error) {
	visitados.Guardar(v, true)
	if contador == n {
		for _, w := range g.Adyacente(v) {
			if w == destino {
				camino, err := ReconstruirCamino[K](padres, destino, v)
				camino.InsertarUltimo(w)
				return camino, err
			}
		}
	} else {
		for _, w := range g.Adyacente(v) {
			if !visitados.Pertenece(w) {
				padres.Guardar(w, v)
				solucion, err := _dfs_cicloN_aux[K](g, contador+1, n, w, destino, padres, visitados)
				if err == nil {
					return solucion, err
				}
			}
		}
	}
	visitados.Borrar(v)
	return TDALista.CrearListaEnlazada[K](), &ERROR.ErrorNoExisteRecorrido{}
}
