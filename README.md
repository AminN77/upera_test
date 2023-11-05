# Golang Product Versioning/Revisions

A comprehensive system in Golang for managing the revisions of products on a shopping site. This system allows you to keep track of product attributes' changes, such as price, color, description, and more, and maintains a history of these revisions for analysis and traceability purposes.

## Installation and Run

For a streamlined setup and running of the system, it's recommended to use the provided Docker Compose file. To do this, make sure you have the `make` tool installed on your system. If it's not already installed, you can typically install it via your system's package manager.

Once you have `make` installed, follow these steps to run the project using Docker Compose:

1. Clone the repository:

   ```bash
   git clone https://github.com/AminN77/upera_test.git
   ```

2. Navigate to the project directory:

   ```bash
   cd upera_test
   ```

3. Run the following command:

   ```bash
   make run
   ```

The provided Docker Compose file will set up the necessary services, including PostgreSQL, MongoDB, and Kafka. It will also launch your Golang application.

**Heads Up**: During the startup, you might encounter some initial issues related to the databases and Kafka. This is normal, and the services will automatically retry connecting until the dependencies are ready. There's no need to worry about these connection retries.

## APIs

### Add Product

- **Description**: Add a new product to the system.

- **HTTP Method**: POST

- **Endpoint**: `{{product}}/api/v1/product`

- **Request Body**:

  ```json
  {
    "name": "ruler 3",
    "description": "30cm ruler",
    "color": "silver",
    "price": 350,
    "imageUrl": "http://someurl/someid"
  }
  ```

### Update Product

- **Description**: Update an existing product.

- **HTTP Method**: PUT

- **Endpoint**: `{{product}}/api/v1/product/:id`

- **Request Body**:

  ```json
  {
    "name": "ruler 3",
    "description": "30cm ruler",
    "color": "green",
    "price": 350,
    "imageUrl": "http://someurl/someid",
    "token": "7451375a-7b3d-11ee-a5bc-1afc20633765"
  }
  ```

### Fetch Product

- **Description**: Fetch details of a specific product (latest attributes).

- **HTTP Method**: GET

- **Endpoint**: `{{product}}/api/v1/product/:id`

### Fetch All Revisions of One Product

- **Description**: Fetch all revisions of a product.

- **HTTP Method**: GET

- **Endpoint**: `{{history}}/api/v1/history/:productId/revision?pageSize=10&pageIndex=1`

- **Query Parameters**:

  - `pageSize`: 10
  - `pageIndex`: 1

### Fetch Detail Of Specific Revision

- **Description**: Fetch details of a specific product for a specific revision number.

- **HTTP Method**: GET

- **Endpoint**: `{{history}}/api/v1/history/:revisionNumber`

Ensure that you replace `{{product}}` and `{{history}}` with the actual base URLs of the product and history services. Additionally, for endpoints that include `:id`, `:productId`, and `:revisionNumber`, you should replace them with the actual values of the respective identifiers.

These API endpoints facilitate the core functionality of the system, allowing you to add, update, and retrieve product information, as well as access historical product revisions.

## Architecture Explanation

The system is designed around two main services:

- **Product Service**: This service handles the logic for managing products. It is responsible for creating, updating, and retrieving product details. The product service communicates with the history service through Kafka, our event bus, to record revisions when product attributes change.

- **History Service**: This service is dedicated to managing product revisions. It captures and stores historical records of product attribute changes. When a product is updated, the history service records the changes in MongoDB, providing a comprehensive history of product revisions.

The two services, product_service and history_service, are distinct and operate independently. They communicate through Kafka to ensure seamless and efficient event-driven handling of product revisions. This architecture allows for scalability and separation of concerns between product management and revision tracking.

### Databases

- **Product Service Database**: PostgreSQL is used to store the latest state of each product. This relational database provides a reliable and structured way to manage product data.

- **History Service Database**: MongoDB is employed to store product revisions. MongoDB's flexibility and document-oriented storage make it well-suited for storing and querying historical data.

This architecture ensures that the system is both performant and capable of maintaining a comprehensive history of product revisions. The separation of databases allows each service to focus on its specific responsibilities, providing a scalable and robust solution for product versioning and revision management.

## Design Decisions

### Token on Product Model

In situations where I anticipated heavy loads and the possibility of race conditions, I decided to introduce a token within the Product Model. This token is intended to grant access to anyone who possesses the latest version of a product, enabling him to update the product. It's important to note that to perform an update, he must first fetch the product to obtain the latest token. After each update, the token is automatically changed.This approach enhances the overall system's stability and prevents simultaneous conflicting updates.

### MongoDB as History Database

When considering the implementation of an event sourcing solution, I chose to leverage MongoDB as the history database for product revisions. Rather than creating a complex event sourcing architecture, I store each version of an update in a revision. This approach provides a more straightforward and efficient way to manage historical data.

MongoDB's document-oriented storage is well-suited for this purpose, allowing me to embed revisions within the product document. This approach not only reduces join overheads but also simplifies queries for historical data. It offers a pragmatic solution to track product revisions and maintain a comprehensive history without the complexity of a full event sourcing system.

### Kafka as Event Bus

I selected Kafka as my event bus for its advantages over other message brokers. Kafka's architecture is optimized for high-throughput and low-latency data streaming, making it an ideal choice for my event-driven system.

Kafka's key benefits include:

- **Scalability**: Kafka's distributed nature and support for partitioning enable horizontal scalability, allowing my system to handle high loads and data volumes.

- **Durability**: Kafka ensures message durability, meaning that once a message is written to a topic, it won't be lost.

- **Low Latency**: Kafka offers low-latency message processing, crucial for real-time systems.

- **Built-in Zookeeper Integration**: Kafka's integration with Zookeeper provides a robust and semi-ready cluster for high-load scenarios, ensuring my system's stability and reliability.

By utilizing Kafka as my event bus, I benefit from its performance, scalability, and built-in features, making it a solid choice for managing event-driven communication and revision tracking within my system.
