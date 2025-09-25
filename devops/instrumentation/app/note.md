# **API Endpoints**

### **1. Order**

#### **POST** `/api/orders`

- **Description**: Create a new order.
- **Request Body**:

  ```json
  {
    "customer_id": "12345",
    "items": [
      {
        "product_id": "P001",
        "quantity": 2,
        "unit_price": 100.0
      },
      {
        "product_id": "P002",
        "quantity": 1,
        "unit_price": 50.0
      }
    ],
    "shipping_address": "Jl. Merdeka No. 10, Jakarta"
  }
  ```

- **Response**:

  ```json
  {
    "order_id": "ORD12345",
    "total_amount": 250.0,
    "status": "Pending",
    "shipping_address": "Jl. Merdeka No. 10, Jakarta"
  }
  ```

- **Description**: This endpoint creates a new order by providing customer details, items they wish to purchase, and a shipping address.

---

#### **GET** `/api/orders/{order_id}`

- **Description**: Get the details of an order by `order_id`.
- **Response**:

  ```json
  {
    "order_id": "ORD12345",
    "customer_id": "12345",
    "order_date": "2025-09-25T10:00:00Z",
    "status": "Pending",
    "total_amount": 250.0,
    "shipping_address": "Jl. Merdeka No. 10, Jakarta",
    "items": [
      {
        "product_id": "P001",
        "quantity": 2,
        "unit_price": 100.0,
        "total_price": 200.0
      },
      {
        "product_id": "P002",
        "quantity": 1,
        "unit_price": 50.0,
        "total_price": 50.0
      }
    ]
  }
  ```

- **Description**: This endpoint retrieves the order details such as the list of items, status, and shipping information.

---

### **2. Payment**

#### **POST** `/api/payments`

- **Description**: Make a payment for an order.
- **Request Body**:

  ```json
  {
    "order_id": "ORD12345",
    "payment_method": "Credit Card",
    "amount_paid": 250.0
  }
  ```

- **Response**:

  ```json
  {
    "payment_id": "PAY12345",
    "order_id": "ORD12345",
    "payment_date": "2025-09-25T10:15:00Z",
    "payment_method": "Credit Card",
    "amount_paid": 250.0,
    "payment_status": "Completed",
    "transaction_id": "TXN98765"
  }
  ```

- **Description**: This endpoint creates a payment for the specified order. The user provides the payment method and amount paid.

---

#### **GET** `/api/payments/{payment_id}`

- **Description**: Get the details of a payment by `payment_id`.
- **Response**:

  ```json
  {
    "payment_id": "PAY12345",
    "order_id": "ORD12345",
    "payment_date": "2025-09-25T10:15:00Z",
    "payment_method": "Credit Card",
    "amount_paid": 250.0,
    "payment_status": "Completed",
    "transaction_id": "TXN98765"
  }
  ```

- **Description**: This endpoint retrieves the details of a payment, including the method, status, and transaction ID.

---

### **3. Shipping**

#### **POST** `/api/shippings`

- **Description**: Create a shipping entry for an order.
- **Request Body**:

  ```json
  {
    "order_id": "ORD12345",
    "shipping_method": "Express",
    "tracking_number": "TRACK12345"
  }
  ```

- **Response**:

  ```json
  {
    "shipping_id": "SHIP12345",
    "order_id": "ORD12345",
    "shipping_date": "2025-09-25T10:20:00Z",
    "shipping_method": "Express",
    "tracking_number": "TRACK12345",
    "shipping_status": "Shipped",
    "estimated_arrival": "2025-09-27T10:00:00Z"
  }
  ```

- **Description**: This endpoint creates a shipping record for an order with a specified shipping method and tracking number.

---

#### **GET** `/api/shippings/{shipping_id}`

- **Description**: Get the details of a shipping by `shipping_id`.
- **Response**:

  ```json
  {
    "shipping_id": "SHIP12345",
    "order_id": "ORD12345",
    "shipping_date": "2025-09-25T10:20:00Z",
    "shipping_method": "Express",
    "tracking_number": "TRACK12345",
    "shipping_status": "Shipped",
    "estimated_arrival": "2025-09-27T10:00:00Z"
  }
  ```

- **Description**: This endpoint retrieves shipping details, such as the shipping method, tracking number, and estimated arrival.

---

### **4. Inventory**

#### **POST** `/api/inventory`

- **Description**: Add or update product inventory.
- **Request Body**:

  ```json
  {
    "product_id": "P001",
    "product_name": "Product A",
    "quantity": 100,
    "unit_price": 150.0
  }
  ```

- **Response**:

  ```json
  {
    "product_id": "P001",
    "product_name": "Product A",
    "quantity": 100,
    "unit_price": 150.0,
    "total_stock_value": 15000.0
  }
  ```

- **Description**: This endpoint adds or updates the inventory for a product, including its quantity and unit price.

---

#### **GET** `/api/inventory/{product_id}`

- **Description**: Get the details of a product in inventory by `product_id`.
- **Response**:

  ```json
  {
    "product_id": "P001",
    "product_name": "Product A",
    "quantity": 100,
    "unit_price": 150.0,
    "total_stock_value": 15000.0
  }
  ```

- **Description**: This endpoint retrieves the current stock and price information for a product in inventory.

---

### **Summary**

| **Entity**    | **Endpoint**                   | **Method** | **Description**               |
| ------------- | ------------------------------ | ---------- | ----------------------------- |
| **Order**     | `/api/orders`                  | `POST`     | Create a new order            |
|               | `/api/orders/{order_id}`       | `GET`      | Get order details             |
| **Payment**   | `/api/payments`                | `POST`     | Create a payment              |
|               | `/api/payments/{payment_id}`   | `GET`      | Get payment details           |
| **Shipping**  | `/api/shippings`               | `POST`     | Create a shipping             |
|               | `/api/shippings/{shipping_id}` | `GET`      | Get shipping details          |
| **Inventory** | `/api/inventory`               | `POST`     | Add or update inventory stock |
|               | `/api/inventory/{product_id}`  | `GET`      | Get inventory details         |

---

### **Conclusion**

These API endpoints will allow you to manage **orders**, **payments**, **shipping**, and **inventory** in a simple and efficient way. Each endpoint is designed to handle the basic CRUD operations (Create, Read) for the respective entities.