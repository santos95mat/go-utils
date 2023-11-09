# go-utils <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" height="48px" />

### CPF pkg

```
newCPF := CPF.Generate() // retorna um CPF válido

isCPF, cpf := CPF.IsValid("xxx.xxx.xxx-xx") // retorna um boleano [true - CPF é valido | false - CPF é invalido] e uma string com o CPF formatado para obter um padrão formatação.
```

### CNPJ pkg

```
newCNPJ := CNPJ.Generate() // retorna um CNPJ válido

isCNPJ, cnpj := CPF.IsValid("xx.xxx.xxx/xxxx-xx") // retorna um boleano [true - CNPJ é valido | false - CNPJ é invalido] e uma string com o CNPJ formatado para obter um padrão formatação.
```

### Password Generator pkg

```
password := password.Generate(X int) // retorna uma senha com X caracteres
```
