package grafo

type Grafo[K comparable] interface {
	AgregarVertice(K)
	SacarVertice(K)
	Adyacente(K) []K
	EstanUnidos(K, K) bool
	//
	AgregarArista(K, K, int)
	SacarArista(K, K)
	PesoArista(K, K) int
	//
	ObtenerVertices() []K
	VerticeAleatorio() K
	//
	Dirigido() bool
	Imprimir()
}
