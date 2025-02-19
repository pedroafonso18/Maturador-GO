# Maturador-GO

Um sistema automatizado para maturação de contas WhatsApp através de trocas de mensagens controladas entre instâncias.

## Sobre o Projeto

O Maturador-GO é uma ferramenta desenvolvida em Go que gerencia a troca automática de mensagens entre diferentes instâncias do WhatsApp para manter as contas ativas e "maduras". O sistema suporta tanto a API Evolution quanto WuzAPI para envio de mensagens.

## Funcionalidades

- Troca automática de mensagens entre instâncias
- Suporte para múltiplos tipos de conteúdo:
  - Mensagens de texto
  - Áudios
  - Stickers
- Delays aleatórios entre mensagens
- Integração com banco de dados PostgreSQL
- Suporte para múltiplas APIs de WhatsApp (Evolution e WuzAPI)

## Configuração

1. Clone o repositório
2. Crie um arquivo `.env` com as seguintes variáveis:
```env
DB_URL=sua_url_do_postgres
EVO_URL=url_da_api_evolution
EVO_TOKEN=seu_token_evolution
WUZ_URL=url_da_api_wuz
```

3. Certifique-se de ter a estrutura de pastas correta para os recursos:
```
resources/
├── audio/
│   └── (arquivos .ogg)
└── sticker/
    └── (arquivos .webp)
```

## Estrutura do Banco de Dados

A tabela `instances` deve conter os seguintes campos:
- name
- instance_id
- limite
- is_evo
- numero
- maturador (boolean)

## Como Funciona

O sistema:
1. Busca todas as instâncias marcadas para maturação no banco de dados
2. Cria um ciclo de troca de mensagens entre as instâncias
3. Seleciona aleatoriamente o tipo de conteúdo (texto, áudio ou sticker)
4. Envia as mensagens com delays aleatórios entre 8 e 30 segundos

## Tecnologias Utilizadas

- Go
- PostgreSQL (pgx driver)
- godotenv
- APIs de WhatsApp (Evolution e WuzAPI)
