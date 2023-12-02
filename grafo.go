package grafo

type Grafo[K comparable] interface {

	// EsVertice devuelve true si el dato pasado por parametro es un vertice del grafo, false si no lo es
	EsVertice(K) bool

	// AgregarVertice crea un nuevo vertice en el grafo con el dato pasado por parametro
	AgregarVertice(K)

	// SacarVertice elimina el vertice que tiene como dato el pasado por parametro. Entra en panico si el dato no es un vertice del grafo.
	SacarVertice(K)

	// Adyacente devuelve una lista de vertices adyacentes al vertice del dato pasado por parametro
	Adyacente(K) []K

	// HayArista devuelve true si hay una arista que va desde el vertice del primer dato pasado por parametro con el segundo, false si no la hay
	HayArista(K, K) bool

	// SacarArista elimina la conexion entre el vertice del primer dato con el segundo. Entra en panico si no existia esa arista previamente.
	SacarArista(K, K)

	// ObtenerVertices devuelve la lista de vertices del grafo
	ObtenerVertices() []K

	// VerticeAleatorio devuelve un vertice al azar
	VerticeAleatorio() K

	// Cantidad devuelve el numero de vertices que hay en el grafo
	Cantidad() int

	// Dirigido devuelve true si el grafo es dirigido, false si no lo es
	Dirigido() bool

	//
	Imprimir()
}

type GrafoNoPesado[K comparable] interface {
	Grafo[K]
	// AgregarAristaNP crea una arista que une el primer vertice pasado por parametro con el segundo
	AgregarAristaNP(K, K)
}

type GrafoPesado[K comparable] interface {
	Grafo[K]

	// AgregarArista crea una arista que une el primer vertice pasado por parametro con el segundo y se le asigna su peso correspondiente
	AgregarArista(K, K, int)

	// PesoArista devuelve el peso de la arista que conecta el primer vertice pasado por parametro con el segundo. Entra en panico si no existe la arista.
	PesoArista(K, K) int
}
