# Contribuindo

Olá!
Se você está aqui provavelmente sentiu necessidade de incrementar essa biblioteca.

Temos um Makefile que contém nossas tasks utilitárias. Por isso, certifique-se de ter o software [make] instalado.

Antes de tudo, execute:

```shell
make setup
```

## Análise de código (linting)

Como ferramenta de análise de código utilizamos o gometalinter, que agrupa inúmeros linters para contextos específicos.

### Execução manual

Para utiliza-lo execute o comando:

```shell
make check
```

ou para uma análise mais profunda:

```shell
make deepcheck
```

### Utilizando no editor de texto

É possível automatizar a execução do gometalinter em alguns [editores de texto](https://github.com/alecthomas/gometalinter#editor-integration), se o fizer, configure o editor para passar o argumento `--fast` ao gometalinter para não sobrecarregar sua máquina.

## Testando

A base para estruturação dos testes na biblioteca é o BDD, e o framework para implementação e execução dos testes é o [ginkgo]. Para entender sua utilização, por favor olhe sua documentação ou alguns dos testes já implementados.

### Executando testes

- Unitários

```shell
make test
```

- Coverage

```shell
make coverage
```

## Mocks

A biblioteca faz uso extensivo de interfaces permitindo baixo acoplamento entre componentes. Com isso dependências podem ser facilmente substituidas por estruturas que permitem definir o comportamento necessário para testes de um fluxo. Para isso, os mocks são compostos das estruturas de mock do repositorio [testify]. Como a construção do mock pode ser um processo trabalhoso, utilizamos um code generator chamado [mockery].

Para regerar os mocks execute o comando:

```shell
make mocks
```

### Registrando mocks para novas interfaces

Há um script em scripts/mocks que contém descrito os mocks a serem gerados das interfaces da biblioteca. Caso haja alguma interface faltante, por favor adicione na aréa descrita no script.

_Caso tenha uma solução melhor para o gerenciamento dos mocks, por favor, abra um pull request com melhorias._

[ginkgo]: https://onsi.github.io/ginkgo/
[testify]: https://github.com/stretchr/testify#mock-package
[mockery]: https://github.com/vektra/mockery
[gometalinter]: https://github.com/alecthomas/gometalinter
[make]: https://en.wikipedia.org/wiki/Make_(software)