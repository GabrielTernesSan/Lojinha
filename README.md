# Go para Web

​	Este projeto tem como objetivo demostrar o uso da linguagem Go voltado para aplicações web. Através deste projeto irei mostrar como desenvolver uma aplicação web simples, onde é possível criar produtos, salvá-los e editar sua informações, isso será possível porque tanto os produtos quanto suas respectivas descrições serão salvas em um banco de dados.

​	![Loja](https://github.com/GabrielTernesSan/Projetos/blob/master/Imagens/Loja.png)

​	Este código foi feito somente com as bibliotecas do Go. Apesar de ter uma página web o curso não é voltado para Front, por esta razão foi utilizado a framework chamada [bootstrap](https://www.alura.com.br/artigos/bootstrap?gclid=Cj0KCQiA-qGNBhD3ARIsAO_o7yl1yxoccENn9_9MmUgOz3cBL4ZZCtnSndrRCr4lKYhCDYiYaHk9BZoaAoumEALw_wcB). O projeto foi todo modularizado usando a padrão de arquitetura [MVC](https://www.lewagon.com/pt-BR/blog/o-que-e-padrao-mvc).

```go
- Controllers
- DB
- Models
- Routes
- Templates
```

​	O primeiro passo para comerçarmos é criar um arquivo, ele vai ser o código Go principal, por convenção chamamos de `main.go`. É neste arquivo que iremos subir nosso servidor, para que quando digitarmos `localhost`, na porta 8000, apareça nossa página principal.

​	Dentro deste arquivo iremos criar uma função, a função `main` do ``package main``, vai ser dentro desta função que iremos subir o servidor, desta forma conseguiremos ouvir uma requisição e respondê-la. Para isso vamos usar a função `http.ListenAndServe(":8000", nil)` da biblioteca "net/hhttp". Essa função vai escutar e responder as requisições da página, que no nosso caso virão da porta 8000, como não precisamos passsar nenhuma informação para a requisição colocaremos `nil`. 

Servidor online!! :)

​	Agora para conseguirmos ver uma página precisamos criá-la. Geralmente quando criamos páginas web em projetos GO chamamos elas de **Templates** e como nosso projeto terá várias páginas web, criaremos uma pasta chamada "Templates" onde todas elas se concentrarão. Primeiro criaremos nossa página inicial, que por convenção é chamada de **Index.html** ela será a porta de entrada da nossa página.
​	Como este projeto não tem como foco o desenvolvimento web vou deixar aqui o código da página principal.
​	

``` html
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css" integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T"
        crossorigin="anonymous">
    <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js" integrity="sha384-JjSmVgyd0p3pXB1rRibZUAYoIIy6OrQ6VrjIEaFf/nJGzIxFDsf4x0xIM+B07jRM"
        crossorigin="anonymous"></script>
    <title>Alura loja</title>
</head>
<nav class="navbar navbar-light bg-light mb-4">
    <a class="navbar-brand" href="/">Alura Loja</a>
</nav>

<body>
    <div class="container">
        <section class="card">
            <div>
                <table class="table table-striped table-hover mb-0">
                    <thead>
                        <tr>
                            <th>Nome</th>
                            <th>Descrição</th>
                            <th>Preço</th>
                            <th>Quantidade</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr>
                            <td>Camiseta</td>
                            <td>Bem bonita</td>
                            <td>29</td>
                            <td>10</td>
                        </tr>
                        <tr>
                            <td>Notebook</td>
                            <td>Muito rápido</td>
                            <td>1999</td>
                            <td>2</td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </section>

    </div>
</body>

</html>
```

​	Agora para conseguirmos de fato atender a requisição e encaminhar para a página "Index" eu preciso chamar mais uma função da biblioteca "net/http", a função `http.HandleFunc()`, o primeiro parâmentro que ela espera é o caminho que ela vai atender. 
​	Como a página "Index" será nossa página inicial, o caminho será o "/", desta forma a função ficará da seguinte maneira `http.HandleFunc("/")`, o segundo parâmetro que ela espera é quem irá atender a requisição, e para este objetivo iremos criar uma função responsável por isso.

A função "Index" terá dois parâmetros, o primeiro será o ```http.ResponseWriter```, com esse parâmetro será possível exibir a página do site toda vez que surgir uma requisição, o segundo parâmetro será o ```*http.Request```. 

