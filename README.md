# Projeto Book

## Sobre a API

Este projeto é uma API que oferece funcionalidades relacionadas a livros.

A estrutura da API é dividida nos seguintes pacotes:

### Pacotes Principais

- **Main**: Este pacote contém a configuração principal da aplicação.
- **Router**: Responsável por lidar com as rotas da API e direcionar as solicitações para os controladores apropriados.
- **Controllers**: Contém os controladores da API que processam as solicitações recebidas pelas rotas.
- **Model**: Aqui estão os modelos de dados usados pela API, definindo a estrutura e os atributos dos objetos manipulados.
- **Repository**: Responsável pela interação com o banco de dados ou fontes externas para buscar ou persistir dados.

### Pacotes Auxiliares

Além dos pacotes principais, existem os seguintes pacotes auxiliares que fornecem utilidades adicionais:

- **Config**: Contém arquivos de configuração necessários para o funcionamento da aplicação.
- **Banco**: Oferece funcionalidades relacionadas ao banco de dados, como conexão e configurações.
- **Autenticação**: Lida com processos de autenticação e autorização na API.
- **Middleware**: Contém middlewares utilizados na aplicação para processar requisições antes de chegarem aos controladores.
- **Segurança**: Pacote dedicado a funcionalidades de segurança da aplicação.
- **Respostas**: Fornece estruturas e métodos para gerar respostas padronizadas nas requisições da API.
