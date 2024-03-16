package  main

import "fmt"

type Endereco struct {
    Logradouro string
    Numero int
    Cidade string
    Estado string

}
type Empresa struct {
    Nome string
}
func (e Empresa) Desativar() {

}

type Pessoa interface{
    Desativar()
}

type Cliente struct {
    Nome string
    Idade int
    Ativo bool
    Endereco
}

func (c Cliente) Desativar() {
    c.Ativo = false
    fmt.Printf("O cliente %s foi desativado", c.Nome)
}

func Desativacao (pessoa Pessoa) {
    pessoa.Desativar()
}

func main() {

    lutcha:= Cliente{
        Nome: "luciano",
        Idade: 30,
        Ativo: true,
    }
    minhaEmpresa := Empresa{}

    Desativacao(minhaEmpresa)
}