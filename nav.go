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
