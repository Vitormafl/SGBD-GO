// cada documento ocupa no minimo 1 e no máximo  bytes
// documento de tamanho variavel é uma sequencia de caractéres(classe)
// a estrutura do documento é <int page_id, int seq, int tam> e é chamada DID 

// classe para representar uma página de disco que contenha uma coleção de documentos
// cada página possui 5 bytes
// o sistema de armazenamento devera ter um total de 20 páginas
// criar uma interface para página a fim de criar métodos para manipular os documentos 

//operações do sistema: 

// Scan()

// Seek(char[])

// Delete (char[])

// Insert (char[])

package main

import "fmt"

type DID struct {
	page_id	int
	seq int
	tam int
}

type Documento struct {
	did DID
	valor []byte
}

type Page struct {
	page_id int
	tam_disp int
	next_page_id *Page
}


func NewDID(tam int) DID {
	return DID{
		page_id: -1,
		seq:     -1,
		tam:     tam,
	}
}

func NewDocumento(tam int, valor []byte) Documento {
	did := NewDID(tam)
	return Documento{
		did:   did,
		valor: valor,
	}
}


func createDocumento(dado []byte) (*Documento, error) {
	if len(dado) < 1 || len(dado) > 5 {
		return nil, fmt.Errorf("Tamanho de dado inválido! Os dados devem ter entre 1 e 5 bytes")
	}

	data := NewDID(len(dado))

	documento := &Documento{
		did:   data,
		valor: dado,
	}

	return documento, nil
}

func createDataBase() *Page{
	var head *Page
	var tail *Page

	for i := 0; i < 20; i++ {
		page := &Page{
			page_id:     i + 1,
			tam_disp:    5,
			next_page_id: nil,
		}

		if head == nil {
			head = page
			tail = page
		} else {
			tail.next_page_id = page
			tail = page
		}
	}

	return head
}
 

// func Seek(docs []*DID, content byte) (*DID, error) {
// 	for _, doc := range docs {
// 		if doc.valor == content {
// 			return doc, nil
// 		}
// 	}

// 	return nil, fmt.Errorf("Document with content '%s' not found", content)
// }

func main() {
	// Exemplo de uso
	pagina_inicial := createDataBase()

	valor := []byte("Hello")

	documento, _ := createDocumento(valor)
	
	fmt.Printf("%d \n",documento.did.page_id)
	fmt.Printf("%d \n",pagina_inicial.page_id)

	// Você pode acessar os campos das instâncias criadas
	// fmt.Println(did.page_id)         // Saída: 1
	// fmt.Println(documento.did.seq)   // Saída: 2
	// fmt.Println(string(documento.valor)) // Saída: Hello, world!

}