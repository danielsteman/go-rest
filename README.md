# go-rest

Inspired on [learning-cloud-native-go](https://github.com/learning-cloud-native-go/myapp).

## Architecture

```
HTTP Request
    ↓
[REST Layer]       ← Handles HTTP (parsing, status codes, JSON)
    ↓
[Service Layer]    ← Business logic (validation, rules, orchestration)
    ↓
[Repository Layer] ← Data access (database queries)
    ↓
Database
```
