# Drawght

Drawght é um manipulador de dados para textos sem declarações lógicas. O
objetivo é usar um conjunto de dados (como assunto de um texto) para fazer
rascunho de um modelo de documento. Pode ser considerado um mini processador de
templates.

Os dados são acessados através de chaves `{}`, substituídos por seus
respectivos valores.

Considerando os seguintes dados:

```yaml
title: Drawght é um rascunho muito útil
author:
  name: Hallison Batista
  email: email@hallison.dev.br
  networks:
  - name: Github
    url: //github.com/hallison
  - name: Twitter
    url: //twitter.com/hallison
creation-date: 2021-06-28
publishing date: 2021-07-01
references:
- name: JSON Template
  url: //code.google.com/archive/p/json-template
- name: Mustache
  url: //mustache.github.io
- name: Handlebars
  url: //handlebarsjs.com
tags:
- Template
- Draft
```

Observe que os campos `creation-date` e `publising date` são identificados
normalmente pelo conversor.

Em um _template_ escrito em Markdown:

```markdown
# {title}

Drawght é uma boa ferramenta para escrever rascunhos de documentos usando
conjunto de dados sem declarações lógicas.

Escrito por {author.name} <{author.email}>, criado em {creation-date},
publicado em {publishing date} e marcado por {tags#1}.

- [{author.networks:name}]({author.networks:url})

Acompanhe as novidades em [{author.networks#1.name}]({author.networks#1.url}).

A sintaxe foi inspirada por:

- [{references:name}]({references:url})

Marcações:

- {tags} (marcado por {author.name}).
```

O	processamento de Drawght retorna o seguinte resultado:

```
# Drawght é um rascunho muito útil

Drawght é uma boa ferramenta para escrever rascunhos de documentos usando
conjunto de dados sem declarações lógicas.

Escrito por Hallison Batista <email@hallison.dev.br>, criado em 2021-06-28,
publicado em 2021-07-01 e marcado por Template.

- [Dev.to](//dev.to/hallison)
- [Github](//github.com/hallison)
- [Twitter](//twitter.com/hallison)

Acompanhe as novidades em [Dev.to](//dev.to/hallison).

A sintaxe foi inspirada por:

- [JSON Template](//code.google.com/archive/p/json-template)
- [Mustache](//mustache.github.io)
- [Handlebars](//handlebarsjs.com)

Marcações:

- Template (marcado por Hallison Batista).
- Draft (marcado por Hallison Batista).
```

Em um _template_ escrito em HTML:

```html
<h1>{title}</h1>

<p>
Drawght é uma boa ferramenta para escrever rascunhos de documentos usando dados
sem declarações lógicas.
</p>

<p>
Escrito por {author.name} &lt;<a href="mailto:{author.email}">{author.email}</a>&gt;
em {creation-date}, publicado em {publishing date} e marcado como {tags#1}.
</p>

<ul>
  <li><a href="{author.networks:url}">{author.networks:name}</a></li>
</ul>

<p>
Acompanhe as novidades em <a href="{author.networks#1.url}">{author.networks#1.name}</a>.
</p>

<p>
A sintaxe foi inspirada por:
</p>

<ul>
  <li><a href="{references:url}">{references:name}</a></li>
</ul>

<p>
Marcações:
</p>

<ul>
  <li>{tags} (marcado por {author.name}).</li>
</ul>
```

O processamento de Drawght retorna o seguinte resultado:

```html
<h1>Drawght é um rascunho muito útil</h1>

<p>
Drawght é uma boa ferramenta para escrever rascunhos de documentos usando dados
sem declarações lógicas.
</p>

<p>
Escrito por Hallison Batista &lt;<a href="mailto:email@hallison.dev.br">email@hallison.dev.br</a>&gt;
em 2021-06-28 e marcado por Template.
</p>

<ul>
  <li><a href="//dev.to/hallison">Dev.to</a></li>
  <li><a href="//github.com/hallison">Github</a></li>
  <li><a href="//twitter.com/hallison">Twitter</a></li>
</ul>

<p>
Acompanhe as novidades em <a href="//dev.to/hallison">Dev.to</a>.
</p>

<p>
A sintaxe foi inspirada por:
</p>

<ul>
  <li><a href="//code.google.com/archive/p/json-template">JSON Template</a></li>
  <li><a href="//mustache.github.io">Mustache</a></li>
  <li><a href="//handlebarsjs.com">Handlebars</a></li>
</ul>

<p>
Marcações:
</p>

<ul>
  <li>Template (marcado por Hallison Batista).</li>
  <li>Draft (marcado por Hallison Batista).</li>
</ul>
```

## Sintaxe

Drawght possui uma sintaxe simples:

- `{chave}`: converte `chave` para o seu respectivo valor. Se o valor for uma
  lista, então a linha será replicada e convertida com os respectivos valores.

- `{objeto.chave}`: converte `chave` para o seu respectivo valor dentro de
  `objeto`, assumindo o mesmo comportamento de `{chave}`.

- `{lista:chave}`: seleciona `lista`, replica a linha para cada item da lista
  e converte `chave` para seu respectivo valor contido em um objeto dentro de
  `lista`. A chave `lista` também pode ser acessado por `objeto.lista`, assim
  como `chave` também pode ser acessado por `objeto.chave` que será convertido
  seguindo o mesmo processo no caso de ser uma lista.

- `{lista#n.chave}`: seleciona o objeto do item `n` (a partir de 1) de `lista`
  e converte `chave` para seu respectivo valor. Todo o processo é semelhante
  ao de `objeto.chave`.

## Licença (MIT)

### Copyright (c) 2021, Hallison Batista

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
of the Software, and to permit persons to whom the Software is furnished to do
so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
