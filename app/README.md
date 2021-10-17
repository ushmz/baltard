# ratri application

## Requirements

- Go 1.15+
- firebase SDK

You need to set up firebase SDK for user authentication.
The credential file must be placed in this directory.

## Structure

```
.
├── cmd
│  └── main file
└── internal
   ├── domain
   │  ├── model
   │  │  └── # Entities (interfaces)
   │  └── repository
   │     └── # Data access interfaces
   ├── handler
   │  └── # handlers(Controllers)
   ├── infra
   │  └── mysql
   │     └── # DB conection and operation codes
   ├── middleware
   │  └── # Middleware for framework(auth, logging ...)
   └── usecase
      └── # Application business logics
```
