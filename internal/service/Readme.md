```
internal/
└── service/
    └── auth/
        ├── service.go          # AuthService (orchestrator)
        ├── login.go            # Login logic
        ├── register.go         # CreateUser
        ├── totp.go             # VerifyTOTP
        ├── session.go          # CheckUserStatus, Logout
        ├── password.go         # ForgotPassword, ResetPassword
        ├── refresh.go          # RefreshAuth
        ├── interfaces.go       # DB / crypto / mail abstractions
        ├── models.go           # Domain structs
        └── errors.go           # Domain errors
```