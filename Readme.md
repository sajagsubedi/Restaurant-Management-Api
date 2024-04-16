# Restaurant Management API 
This RESTful API, built with Golang, manages various aspects of a restaurant, including food items, invoices, menus, order items, orders, and user profiles. It utilizes PostgreSQL as the database system.
## Routes

### Food Routes
- **GET /api/v1/foods**: Retrieve all food items.
- **GET /api/v1/foods/:foodid**: Retrieve a specific food item by ID.
- **POST /api/v1/foods/add**: Add a new food item (admin privilege required).
- **PATCH /api/v1/foods/update/:foodid**: Update a food item (admin privilege required).
- **DELETE /api/v1/foods/delete/:foodid**: Delete a food item (admin privilege required).

### Invoice Routes
- **GET /api/v1/invoices**: Retrieve all invoices.
- **GET /api/v1/invoices/:invoiceid**: Retrieve a specific invoice by ID.
- **POST /api/v1/invoices/add**: Create a new invoice.
- **PATCH /api/v1/invoices/update/:invoiceid**: Update an invoice.
### Order Routes
- **POST /api/v1/orders/add**: Create a new order (user privilege required).
- **PATCH /api/v1/orders/update/:orderid**: Update an order (user privilege required).
- **GET /api/v1/orders**: Retrieve all orders (admin privilege required).
- **GET /api/v1/orders/:orderid**: Retrieve a specific order by ID (admin privilege required).
### Menu Routes
- **GET /api/v1/menus**: Retrieve all menus.
- **GET /api/v1/menus/:menuid**: Retrieve a specific menu by ID.
- **POST /api/v1/menus/add**: Add a new menu (admin privilege required).
- **PATCH /api/v1/menus/update/:menuid**: Update a menu (admin privilege required).

### Order Item Routes
- **POST /api/v1/orderitems/create**: Create a new order item (user privilege required).
- **PATCH /api/v1/orderitems/:orderitemid**: Update an order item (user privilege required).
- **DELETE /api/v1/orderitems/:orderitemid**: Delete an order item (user privilege required).
- **GET /api/v1/orderitems**: Retrieve all order items (admin privilege required).
- **GET /api/v1/orderitems/:orderitemid**: Retrieve a specific order item by ID (admin privilege required).
- 
### User Routes
- **POST /api/v1/users/signup**: Register a new user.
- **POST /api/v1/users/signin**: Authenticate and retrieve access token.
- **POST /api/v1/users/getaccesstoken**: Refresh access token.
- **GET /api/v1/users/profile**: Retrieve user profile (user privilege required).
- **PATCH /api/v1/users/updateprofile**: Update user profile (user privilege required).
- **GET /api/v1/users**: Retrieve all users (admin privilege required).
- **GET /api/v1/users/:userid**: Retrieve a specific user by ID (admin privilege required).
- **PATCH /api/v1/users/update/:userid**: Update user profile (admin privilege required).
## Authentication and Authorization
- Authentication is handled using JWT tokens.
- User roles (admin or regular user) determine access to certain routes.
- Middleware functions are used to check user roles before granting access to specific endpoints.
## Dependencies
- Gin (`github.com/gin-gonic/gin`) for HTTP routing.
- PostgreSQL as the database system.

## Getting Started
1. Clone the repository.
2. Install dependencies (`go mod tidy`).
3. Set up your PostgreSQL database and create the necessary tables (see `/database` folder).
4. Create a `.env` file with the following variables:
   - `PORT`: Port number for the server.
   - `SECRET_KEY`: Secret key for JWT token encryption.
   - `POSTGRES_URI`: Connection URI for PostgreSQL database.
5. Run the server (`go run main.go`).

## Folder Structure
- `/controllers`: Handlers for request processing.
- `/middlewares`: Middleware functions for authentication and authorization.
- `/routes`: Route definitions for API endpoints.
- `/models`: Database models.
- `/helpers`: Helper functions.
- `/database`: Scripts for database setup.

## Contributors
- [Sajag Subedi](https://github.com/sajagsubedi)

