# AGENTS.md - Auth+ Notification Service

> **Propósito:** Ponto de entrada técnico e guia de regras de negócio para Agentes de IA atuando no microsserviço `auth-plus-notification`.
>
> **Bounded Context:** Notificações Transacionais Multicanal (Email, SMS, WhatsApp, Push Notifications, Telegram), Balanceamento de Provedores e Estratégias de Entrega (IP Warming, Failover).

---

## 1. Visão Geral e Stack Tecnológica

O `auth-plus-notification` atua como a central de comunicações externas da plataforma, desacoplando os provedores de entrega do restante dos microsserviços.

- **Runtime / Linguagem:** Go (Golang 1.20+)
- **Estrutura:** Clean Architecture / Driven Adapters
- **Mensageria:** Apache Kafka (`kafka:9092`) via `segmentio/kafka-go`
- **Observabilidade:** OpenTelemetry Go SDK (`go.opentelemetry.io/otel`) integrado ao Uptrace (`uptrace:4318` via HTTP / gRPC) com `otelzap`
- **Servidores / Binários:**
  - `cmd/http/http_server.go`: API REST na porta `5001` (host e container)
  - `cmd/kafka/kafka_server.go`: Consumidor concorrente de eventos Kafka

---

## 2. Regras de Negócio e Estratégias de Disparo

### 2.1. Canais e Provedores Suportados
O serviço suporta envio unificado através de múltiplos canais com diferentes provedores terceirizados:

| Canal | Provedores Integrados |
| :--- | :--- |
| **Email** | SendGrid, Mailgun, OneSignal |
| **SMS** | Twilio, AWS SNS |
| **WhatsApp** | Twilio API |
| **Push Notification** | Firebase Cloud Messaging (FCM), OneSignal |
| **Telegram** | Telegram Bot API |

### 2.2. Padrão Manager de Roteamento e IP Warming
O envio de notificações nunca é acoplado diretamente a um provedor específico. O caso de uso consulta um `Manager`:
1. **`IPWarmimgEmailManager`**:
   - Implementa estratégia de aquecimento gradual de reputação de IP para envio de e-mails em escala.
   - Distribui percentualmente os disparos entre múltiplos provedores (ex.: cota prioritária no SendGrid, cota intermediária no OneSignal e fallback no Mailgun).
2. **`RandomManager`**:
   - Distribui requisições aleatoriamente entre provedores disponíveis para balanceamento de tráfego e rate limits.
3. **Resiliência e Fallback:** Se um provedor apresentar indisponibilidade ou esgotamento de quota, o manager pode selecionar a rota de contingência.

### 2.3. Consumo Reativo de Eventos do Kafka
O serviço escuta e processa os tópicos operacionais da organização:
- **`2FA_EMAIL_CREATED` / `2FA_PHONE_CREATED`**: Despacha códigos e links de desafio de verificação de autenticação de dois fatores.
- **`2FA_EMAIL_SENT` / `2FA_PHONE_SENT`**: Dispara códigos OTP para o destino solicitado.
- **`USER_CREATED`**: Dispara mensagens de boas-vindas e instruções iniciais de onboarding.
- **`ORGANIZATION_CREATED`**: Notifica administradores sobre provisionamento de novos tenants.

---

## 3. Contratos de API (Endpoints Principais)

| Método | Rota | Payload Principal | Descrição |
| :--- | :--- | :--- | :--- |
| `POST` | `/email` | `{ "email": "...", "subject": "...", "content": "..." }` | Disparo transacional síncrono de e-mail |
| `POST` | `/sms` | `{ "phone": "...", "content": "..." }` | Disparo transacional de SMS |
| `POST` | `/push-notification` | `{ "device_id": "...", "content": "..." }` | Envio de push notification via FCM/OneSignal |

---

## 4. Instruções de Desenvolvimento e Testes

```bash
# Baixar módulos e dependências Go
go mod download

# Executar servidor HTTP
go run ./cmd/http/http_server.go

# Executar consumidor Kafka
go run ./cmd/kafka/kafka_server.go

# Executar suíte de testes unitários e de integração
go test -v ./...

# Cobertura de testes
go test -coverprofile=coverage.out ./...
```

---

## 5. Diretrizes para Agentes de IA

1. **Desacoplamento de Credenciais:** Nunca versione chaves de API de provedores reais (SendGrid, Twilio, AWS). Em ambiente local/testes, utilize mocks configurados em `test/mocks/` ou logs emulados.
2. **Propagação de Trace Context no Kafka:** No consumidor Kafka (`api/kafka/index.go`), extraia sempre os metadados do W3C Trace Context usando o `KafkaHeaderCarrier` para que os spans de envio no Uptrace sejam filhos diretos da requisição originária do `authentication` ou `billing`.
3. **Tratamento de Erros:** Disparos para canais externos devem registrar logs enriquecidos com `otelzap` contendo identificadores anonimizados e status de retorno do provedor.
