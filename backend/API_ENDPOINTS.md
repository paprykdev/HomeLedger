# HomeLedger API Endpoints (Frontend Reference)

Base URL: `http://localhost:<PORT>`

## Auth

Protected endpoints require:

`Authorization: Bearer <jwt_token>`

Token is returned by register/login endpoints.

---

## Health

### `GET /health`
Public health check.

**Response 200**
```json
{
  "status": "ok"
}
```

---

## Authentication

### `POST /auth/register`
Create a user and return JWT token + user data.

**Request body**
```json
{
  "username": "john",
  "email": "john@example.com",
  "password": "secret123"
}
```

**Response 201**
```json
{
  "token": "<jwt>",
  "user": {
    "id": "uuid",
    "username": "john",
    "email": "john@example.com",
    "created_at": "2026-01-01T10:00:00Z",
    "updated_at": "2026-01-01T10:00:00Z"
  }
}
```

**Common errors**
- `400` invalid request body / missing fields / password too short
- `409` username already exists / email already exists

### `POST /auth/login`
Authenticate existing user and return JWT token + user data.

**Request body**
```json
{
  "username": "john",
  "password": "secret123"
}
```

**Response 200**
```json
{
  "token": "<jwt>",
  "user": {
    "id": "uuid",
    "username": "john",
    "email": "john@example.com",
    "created_at": "2026-01-01T10:00:00Z",
    "updated_at": "2026-01-01T10:00:00Z"
  }
}
```

**Common errors**
- `400` username and password required
- `401` invalid credentials

---

## Accounts (JWT required)

### `POST /accounts/`
Create account for authenticated user.

**Request body**
```json
{
  "name": "Main Wallet",
  "currency": "PLN"
}
```

**currency allowed values:** `PLN`, `USD`, `EUR`

**Response 201**
```json
{
  "id": "uuid",
  "name": "Main Wallet",
  "currency": "PLN",
  "user_id": "uuid",
  "created_at": "2026-01-01T10:00:00Z",
  "updated_at": "2026-01-01T10:00:00Z"
}
```

### `GET /accounts/`
List accounts for authenticated user.

**Response 200**
```json
[
  {
    "id": "uuid",
    "name": "Main Wallet",
    "currency": "PLN",
    "user_id": "uuid",
    "created_at": "2026-01-01T10:00:00Z",
    "updated_at": "2026-01-01T10:00:00Z"
  }
]
```

### `GET /accounts/{id}`
Get single account (must belong to authenticated user).

**Response 200**
```json
{
  "id": "uuid",
  "name": "Main Wallet",
  "currency": "PLN",
  "user_id": "uuid",
  "created_at": "2026-01-01T10:00:00Z",
  "updated_at": "2026-01-01T10:00:00Z"
}
```

**Errors**
- `404` account not found

### `DELETE /accounts/{id}`
Soft-delete account.

**Response 204** (no body)

---

## Transactions (JWT required)

### `POST /transactions/`
Create transaction for authenticated user.

**Request body**
```json
{
  "type": "expense",
  "amount": 50.25,
  "description": "groceries",
  "category": "food",
  "account_id": "account-uuid"
}
```

**type allowed values:** `income`, `expense`

**Response 201**
```json
{
  "id": "uuid",
  "user_id": "uuid",
  "type": "expense",
  "amount": 50.25,
  "description": "groceries",
  "category": "food",
  "account_id": "account-uuid",
  "created_at": "2026-01-01T10:00:00Z",
  "updated_at": "2026-01-01T10:00:00Z"
}
```

Notes:
- `category` is normalized to lowercase.
- Empty `category` becomes `"other"`.
- `account_id` must belong to authenticated user.

### `GET /transactions/`
List transactions for authenticated user.

**Response 200**
```json
[
  {
    "id": "uuid",
    "user_id": "uuid",
    "type": "income",
    "amount": 5000,
    "description": "salary",
    "category": "job",
    "account_id": "account-uuid",
    "created_at": "2026-01-01T10:00:00Z",
    "updated_at": "2026-01-01T10:00:00Z"
  }
]
```

### `GET /transactions/{id}`
Get single transaction (must belong to authenticated user).

**Response 200**
```json
{
  "id": "uuid",
  "user_id": "uuid",
  "type": "expense",
  "amount": 50.25,
  "description": "groceries",
  "category": "food",
  "account_id": "account-uuid",
  "created_at": "2026-01-01T10:00:00Z",
  "updated_at": "2026-01-01T10:00:00Z"
}
```

**Errors**
- `404` transaction not found

### `PATCH /transactions/{id}`
Partially update transaction.

Supported fields in request body:
- `type`
- `amount`
- `description`
- `category`

**Request body example**
```json
{
  "amount": 65.1,
  "category": "shopping"
}
```

**Response 200**
Returns full updated transaction object.

Important behavior:
- `amount` updates only when non-zero.
- `description` updates only when non-empty.
- `category` and `type` are normalized (trim + lowercase).

### `DELETE /transactions/{id}`
Soft-delete transaction.

**Response 204** (no body)

---

## Error response format

Current API mostly returns plain text errors via `http.Error`, for example:

- `unauthorized`
- `invalid request body`
- `failed to create transaction`

For frontend integration, treat non-2xx as text responses (not guaranteed JSON).
