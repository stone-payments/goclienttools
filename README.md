# GoClientTools

Este repositório contém pacotes utilitarios para se construir um Client em Go.


## Utilizando no seu projeto

### Setup

Basta baixar a biblioteca para o seu workspace utilizando o comando tradicional do go, ou qualquer gerenciador de dependencia de sua escolha.

```bash
go get -u github.com/stone-payments/goclienttools
```

### Show me the code

A interface principal da biblioteca é o Manager, ele é responsável por gerenciar as requisições.

```go
import (
    "github.com/stone-payments/goclienttools/http"
    "github.com/stone-payments/goclienttools/errors"    
)

type Client struct {
    manager http.Manager
}

func (c *Client) Get(id int) (Something, errors.Error){
    url := r.manager.BuildURL(endpoint, id)
    resp, err := c.manager.Request(http.Requester.Get, url, http.Options())

    var something Something
	if err == nil {
		err = resp.JSON(&something)
	}

	return &something, err
}

client := Client{baseUrl}
```

É possível registrar interceptors que processem requests e responses. Basta implementar a interface Interceptor do pacote http.

```go
import (
    "github.com/stone-payments/goclienttools/errors"
    "github.com/stone-payments/goclienttools/http"
)

func init() {
    http.Interceptors = append(http.Interceptors,
        http.NewInterceptor(func(req http.Request) http.Request {
            println("on request event")
            return req
        }, func(resp http.Response) (http.Response, errors.Error) {
            println("on response event")
            return resp, nil
        }))
}

```

Também é possível substituir o http requester padrão
```go
import (
    "github.com/stone-payments/goclienttools/http"
)

func init() {
    http.Requester = requester
}

```

## Feito com

[Go](https://golang.org/) 1.9

## Como contribuir?

Todo o processo necessário para você contribuir encontra-se em nosso [CONTRIBUTING](CONTRIBUTING.md).

## Versionamento

Por favor, consulte nosso [CHANGELOG](CHANGELOG.md).

## Autores

Build with :heart: by Dev RC Team!

## License

Copyright (C) 2018. Stone Co. All rights reserved.
