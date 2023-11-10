# go-utils <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" height="48px" />

### Repositório com packages criados por mim para serem utilizados em outros projetos

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
pw := &pwg.PWG{
		LowCaseQuantity:     4,
		UpCaseQuantity:      1,
		NumbersQuantity:     2,
		SpecialCharQuantity: 1,
	}

// Gera uma senha aleatória
password := pw.Generate()
```
