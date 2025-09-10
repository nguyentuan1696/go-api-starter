# Go API Starter

A modern, production-ready Go API backend built with clean architecture principles and modular design. This project serves as a comprehensive starter template for building scalable web APIs with authentication, file storage, and business logic modules.

## 🚀 Features

### Core Functionality
- **Authentication & Authorization**: JWT-based auth with role-based access control (RBAC)
- **Product Management**: Comprehensive e-commerce product catalog system
- **File Storage**: Cloud storage integration with AWS S3/R2 support
- **User Management**: Complete user lifecycle with OAuth provider support
- **Order Processing**: Full order management with payment and shipping methods

### Technical Stack
- **Go 1.24.4**: Modern Go with latest features
- **Echo v4**: High-performance HTTP web framework
- **PostgreSQL**: Primary database with SQLx ORM
- **Redis**: Caching and session storage
- **JWT**: Secure token-based authentication
- **Docker**: Containerization with multi-stage builds
- **AWS S3/R2**: Cloud storage for file uploads

## 🏗️ Architecture

The project follows **Clean Architecture** principles with clear separation of concerns:

```
├── core/                 # Core infrastructure layer
│   ├── cache/           # Redis cache implementation
│   ├── config/          # Configuration management
│   ├── database/        # Database connections (PostgreSQL, Redis)
│   ├── middleware/      # HTTP middleware (auth, CORS, logging)
│   ├── server/          # HTTP server setup
│   └── utils/           # Shared utilities
├── modules/             # Business logic modules
│   ├── auth/           # Authentication & authorization
│   ├── product/        # Product management (cosmetics)
│   └── storage/        # File upload & management
└── templates/          # Email templates
```

### Module Structure
Each module follows a consistent layered architecture:
- **Controller**: HTTP handlers and request/response logic
- **Service**: Business logic and orchestration
- **Repository**: Data access layer
- **Entity**: Domain models
- **DTO**: Data transfer objects
- **Router**: Route definitions
- **Validator**: Input validation

## 🛠️ Quick Start

### Prerequisites
- Go 1.24.4+
- PostgreSQL 15+
- Redis 7+
- Docker (optional)

### Local Development

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd go-api-starter
   ```

2. **Setup environment**
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

3. **Start dependencies**
   ```bash
   docker-compose up -d postgres redis
   ```

4. **Run the application**
   ```bash
   make dev
   # or
   go run main.go -env dev
   ```

### Docker Deployment

```bash
# Build and run with Docker
make docker-build
make docker-run

# View logs
make docker-logs
```

## 📝 Configuration

The application uses environment variables for configuration. See `.env.example` for all available options:

- **Server**: Host, port, timeouts
- **Database**: PostgreSQL connection settings
- **Redis**: Cache configuration
- **JWT**: Secret key for token signing
- **SMTP**: Email service configuration
- **R2/S3**: Cloud storage credentials
- **Frontend**: CORS and integration URLs

## 🔐 Authentication

The auth module provides:
- User registration and login
- JWT token management with refresh tokens
- Role-based permissions system
- OAuth provider integration (ready)
- Password reset functionality
- Email verification

## 📦 Product Management

The product module includes:
- **Products**: Full CRUD with categories and brands
- **Brands**: Brand management system
- **Categories**: Hierarchical product categorization
- **Ingredients**: Cosmetic ingredient tracking
- **Tags**: Product tagging system
- **Orders**: Complete order processing
- **Payment Methods**: Payment integration ready
- **Shipping Methods**: Shipping options management

## 📁 File Storage

The storage module provides:
- Secure file upload to R2/S3
- Image processing and optimization
- File metadata management
- CDN-ready file serving
- Access control for uploaded files

## 🚀 API Endpoints

### Authentication
- `POST /auth/register` - User registration
- `POST /auth/login` - User login
- `POST /auth/refresh` - Refresh JWT token
- `POST /auth/forgot-password` - Password reset request
- `GET /auth/profile` - Get user profile

### Products (Protected)
- `GET /products/items` - List products
- `POST /products/items` - Create product
- `PUT /products/items/:id` - Update product
- `DELETE /products/items/:id` - Delete product
- `GET /products/brands` - List brands
- `GET /products/categories` - List categories

### Storage (Protected)
- `POST /storage/upload` - Upload file
- `GET /storage/files` - List files
- `DELETE /storage/files/:id` - Delete file

## 🧪 Development

### Available Commands

```bash
# Development
make dev          # Run in development mode
make build        # Build the application
make test         # Run tests

# Docker
make docker-build # Build Docker image
make docker-run   # Run Docker container
make docker-stop  # Stop container
make docker-logs  # View logs

# Utilities
make clean        # Clean build artifacts
make help         # Show all commands
```

### Code Structure Guidelines

1. **Follow Clean Architecture**: Keep business logic separate from infrastructure
2. **Use Dependency Injection**: All dependencies are injected through constructors
3. **Error Handling**: Consistent error handling with custom error types
4. **Validation**: Input validation at controller level
5. **Logging**: Structured logging throughout the application
6. **Testing**: Unit tests for business logic, integration tests for APIs

## 🔒 Security Features

- JWT-based authentication with secure token handling
- Role-based access control (RBAC)
- Input validation and sanitization
- SQL injection prevention with parameterized queries
- CORS configuration
- Rate limiting ready
- Secure file upload with type validation
- Password hashing with bcrypt

## 📊 Monitoring & Logging

- Structured logging with contextual information
- Request/response logging middleware
- Health check endpoints
- Performance metrics ready
- Error tracking and reporting

## 🚀 Production Deployment

### Docker Production Build
The Dockerfile uses multi-stage builds for optimized production images:
- Minimal Alpine Linux base image
- Non-root user for security
- Optimized binary size
- Health checks included

### Environment Setup
1. Configure production environment variables
2. Set up PostgreSQL and Redis instances
3. Configure R2/S3 storage
4. Set up SMTP for email services
5. Deploy with Docker or Kubernetes

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🆘 Support

For questions and support:
- Create an issue in the repository
- Check the documentation
- Review the code examples

---

**Go API Starter** - A solid foundation for your next Go API project! 🚀