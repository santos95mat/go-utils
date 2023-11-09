# go-utils <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" height="48px" />

### CPF pkg
```
// Gera um CPF aleatório
newCPF := CPF.Generate()

// Valida um CPF
isCPF, cpf := CPF.IsValid("xxx.xxx.xxx-xx")
```

### CNPJ pkg

```
// Gera um CNPJ aleatório
newCNPJ := CNPJ.Generate()

// Valida um CNPJ
isCNPJ, cnpj := CPF.IsValid("xx.xxx.xxx/xxxx-xx")
```

### Password Generator pkg

```
// Cria um objeto com as informções para gerar um senha
passwordDTO := &password.CreatePasswordDTO{
		LowCaseQuantity:     3,
		UpCaseQuantity:      3,
		NumbersQuantity:     3,
		SpecialCharQuantity: 1,
	}

// Gera uma senha aleatória
pw := password.Generate(passwordDTO)
```
