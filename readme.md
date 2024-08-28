
# AuthServiceAPI

This project provides a simple REST API designed for user authentication and secure access to protected resources. It includes endpoints for user registration, login, and accessing protected resources. Ideal for developers needing a foundational backend service to handle user authentication in their applications. The API ensures secure storage of user credentials, supports authentication through access tokens, and demonstrates best practices in API security and user management.


## Design Decisions
To ensure a structured and maintainable development process, this project follows the principle of separation of concerns. The system is divided into three main components:

1. Database Layer: Responsible for interacting with the database, including storing and retrieving user data. This layer ensures that data-related logic is managed separately from the application logic.

2. Routing Layer: Manages API routes and directs requests to the appropriate controllers. This layer is responsible for determining which endpoint should be accessed based on the URL and HTTP method used.

3. Controller Layer: Handles business logic and application flow. The controller receives requests from the routing layer, interacts with the database layer to access or modify data, and returns the appropriate response to the client.

By separating these components, the project not only ensures cleaner and more manageable code but also enhances the security and scalability of the application.
## Library Choices
For this project, the following technologies and libraries have been utilized:

1. Golang: The entire backend service is implemented using Golang. Golang was chosen for its efficiency, performance, and strong support for concurrent programming, which is ideal for building scalable and high-performance APIs.

2. MySQL: MySQL is used as the database for storing user information. MySQL was selected for its reliability, robustness, and extensive support for data management, which ensures secure and efficient handling of user data.


## Challenges Encountered
One of the primary challenges faced during this project was organizing the code to adhere to the principle of separation of concerns. Ensuring that the logic for database interactions, routing, and application control was distinctly separated required careful planning and refactoring.

The main difficulty was structuring the codebase so that each component—database layer, routing layer, and controller layer—was effectively isolated, yet functioned seamlessly together. This involved creating clear boundaries between different layers, avoiding tight coupling, and ensuring that each component was responsible for a specific aspect of the application’s functionality.

By addressing these challenges, the project achieved a well-organized and maintainable codebase, enhancing both its security and scalability.
## Routes
This API provides the following endpoints:

- '/signin' : Used for user authentication. You need to provide the email and password in JSON format to receive an access token.

Request:
```json
{
  "email": "user@example.com",
  "password": "yourpassword"
}
```
- '/signup' : Used for user registration. You need to provide username, email, and password in JSON format.

Request:
```json
{
  "username": "yourusername",
  "email": "user@example.com",
  "password": "yourpassword"
}
```
- '/user' : A protected route that requires a bearer token obtained from the /signin route. The token must be included in the Authorization header as a Bearer token.

Request Header:
```
Authorization: Bearer <your_access_token>
```

## Usage
#### Build Yourself
Follow these steps:
1. Clone the repository:
```bash
  git clone https://github.com/fossyy/AuthServiceAPI.git
```

2. Install dependencies::
```bash
  cd AuthServiceAPI
  go mod tidy
```

3. Build and run the application:
```bash
  go run main.go
```

4. Access the application in Postman or a similar tool:
http://localhost:8000

#### Download the Binary
Alternatively, you can download the pre-built binary:

1. Download the binary from the repository's releases.

2. Run the binary directly:

```bash
./AuthServiceAPI
```

3. Access the application in Postman or a similar tool: http://localhost:8000

