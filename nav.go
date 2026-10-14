package nav

// cloneItems devolve uma copia do slice de itens.
//
// Os construtores recebem o slice variadico do chamador e os metodos Items()
// devolvem a lista interna. Sem a copia, quem guarda o slice original (ou o
// resultado de Items()) consegue alterar o estado do componente por fora.
func cloneItems(items []string) []string {
	if items == nil {
		return nil
	}
	copied := make([]string, len(items))
	copy(copied, items)
	return copied
}

// clampIndex limita um indice ao intervalo valido da lista.
// Retorna -1 quando a lista esta vazia ou o indice e invalido.
func clampIndex(index, length int) int {
	if length <= 0 || index < 0 || index >= length {
		return -1
	}
	return index
}
