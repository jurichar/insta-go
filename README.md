# Insta-go

Insta-go is a simple social network application that allows users to post pictures, follow other users, like and comment on posts.

## Features

- User authentication
- User profile
- Post pictures
- Follow other users
- Like and comment on posts

## Technologies

- Go
- Fiber
- GORM
- PostgreSQL
- JWT
- template/html

## Installation

1. Clone the repository
2. Install the dependencies
3. Create a `.env` file and add the following environment variables:

```
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=your_db_name
JWT_SECRET=your_jwt_secret
```

4. Run the application

```bash
go run main.go
```

## License

This project is open-sourced software licensed under the [MIT license](https://opensource.org/licenses/MIT).
