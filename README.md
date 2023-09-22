# go-utils <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" height="48px" />

### CPF package

```
// retorna uma string de um CPF
newCPF := CPF.Generate()

// retorna um boleano [true - cpf é valido | false - cpf é invalido]
// retorna um error caso o CPF digitado não satisfaça a regras ["xxx.xxx.xxx-xx" ou "xxxxxxxxxxx"]
isCPF, err := CPF.IsValid("xxx.xxx.xxx-xx")
```
