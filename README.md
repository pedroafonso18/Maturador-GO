# Maturador-GO

Um sistema automatizado para maturação de contas WhatsApp através de trocas de mensagens controladas entre instâncias.

## Sobre o Projeto

O Maturador-GO é uma ferramenta desenvolvida em Go que gerencia a troca automática de mensagens entre diferentes instâncias do WhatsApp para manter as contas ativas e "maduras". O sistema suporta tanto a API Evolution quanto WuzAPI para envio de mensagens.

## Funcionalidades

- Troca automática de mensagens entre instâncias
- Funcionamento restrito ao horário comercial (8h às 17h)
- Suporte para múltiplos tipos de conteúdo:
  - Mensagens de texto pré-definidas
  - Áudios
  - Stickers
- Delays aleatórios entre mensagens (8-30 segundos)
- Criação automática do banco de dados
- Integração com banco de dados PostgreSQL
- Suporte para múltiplas APIs de WhatsApp:
  - Evolution API
  - WuzAPI

## Configuração

1. Clone o repositório
2. Crie um arquivo `.env` com as seguintes variáveis:
```env
DB_URL=sua_url_do_postgres
EVO_URL=url_da_api_evolution
EVO_TOKEN=seu_token_evolution
WUZ_URL=url_da_api_wuz
```

3. Prepare a estrutura de recursos:
```
resources/
├── audio/
│   └── (arquivos .ogg)
└── sticker/
    └── (arquivos .webp)
```

## Estrutura do Banco de Dados

O sistema criará automaticamente o banco de dados caso não exista.

Tabela `instances`:
- name (string): Nome da instância
- instance_id (string): ID único da instância
- limite (int): Limite de mensagens
- is_evo (boolean): Indica se usa Evolution API
- numero (string): Número do WhatsApp
- maturador (boolean): Flag para participar da maturação

## Como Funciona

1. Verifica se está dentro do horário permitido (8h-17h)
2. Busca instâncias marcadas para maturação no banco
3. Cria um ciclo de troca de mensagens:
   - Seleciona aleatoriamente o tipo de conteúdo
   - Aplica delays aleatórios entre mensagens
   - Alterna entre instâncias para simular conversas reais

## Tecnologias Utilizadas

- Go 1.22+
- PostgreSQL com pgx driver
- godotenv para configuração
- APIs:
  - Evolution API
  - WuzAPI
