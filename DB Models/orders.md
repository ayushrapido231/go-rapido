Orders Entity : 

{
    "orderId": "12345",
    "userId": "123",
    "products": [
        {
            "productId": "12345",
            "productName": "Product Name",
            "quantity": 1,
            "subtotal": 1000
        }
    ],
    "orderStatus": "pending",
    "orderDate": "2025-11-28",
    "orderTotal": 1000,
    "paymentMethod": "cash",
    "paymentStatus": "pending",
    "paymentDate": "2025-11-28",
    "paymentAmount": 1000,
    "createdAt": "2025-11-28",
    "updatedAt": "2025-11-28"
}

Read Patterns & APIs

1. View Order History
- API: GET /api/users/{userId}/orders
- Lookup Key: userId
- Pattern: Read heavy
- Use Case: Displaying user's order history
- Query Pattern:
  db.orders.find({ userId: "123" })
    .sort({ orderDate: -1 })
    .limit(limit)
  

2. View Order Details
- API: GET /api/orders/{orderId}
- Lookup Key: orderId
- Pattern: Non read-heavy
- Use Case: Displaying specific order details
- Query Pattern:
  db.orders.findOne({ orderId: "12345" })
  

3. Track Order Status
- API: GET /api/orders/{orderId}/status
- Lookup Key: orderId
- Pattern: Read heavy
- Use Case: Order tracking functionality
- Query Pattern:
  db.orders.findOne(
    { orderId: "12345" },
    { orderId: 1, orderStatus: 1, orderDate: 1, updatedAt: 1 }
  )
  

4. View Orders by Date Range
- API: GET /api/users/{userId}/orders?startDate={start}&endDate={end}
- Lookup Key: orderDate, userId
- Pattern: Read heavy
- Use Case: Filtering orders by date range
- Query Pattern:
  db.orders.find({
    userId: "123",
    orderDate: { $gte: startDate, $lte: endDate }
  }).sort({ orderDate: -1 })
  

5. View Orders by Status
- API: GET /api/users/{userId}/orders?status={status}
- Lookup Key: orderStatus, userId
- Pattern: Read heavy
- Use Case: Filtering orders by status (pending, completed, cancelled)
- Query Pattern:
  db.orders.find({
    userId: "123",
    orderStatus: "pending"
  }).sort({ orderDate: -1 })
  

Write Patterns & APIs

1. Create New Order
- API: POST /api/orders
- Lookup Key: orderId (unique check), userId
- Pattern: Write heavy
- Use Case: Creating new order from cart
- Query Pattern:
  // Check uniqueness
  db.orders.findOne({ orderId: "newOrderId" })
  // Insert order
  db.orders.insertOne({
    orderId: "newOrderId",
    userId: "123",
    products: [...],
    orderStatus: "pending",
    orderDate: new Date(),
    ...
  })
  // Update product stock
  db.products.updateMany(
    { productId: { $in: productIds } },
    { $inc: { productQuantity: -quantity } }
  )
  

2. Update Order Status
- API: PUT /api/orders/{orderId}/status
- Lookup Key: orderId
- Pattern: Write heavy
- Use Case: Updating order status (pending → processing → shipped → delivered)
- Query Pattern:
  db.orders.updateOne(
    { orderId: "12345" },
    { $set: { orderStatus: "shipped", updatedAt: new Date() } }
  )
  

3. Cancel Order
- API: PUT /api/orders/{orderId}/cancel
- Lookup Key: orderId
- Pattern: Write heavy
- Use Case: Cancelling an order
- Query Pattern:
  // Update order status
  db.orders.updateOne(
    { orderId: "12345" },
    { $set: { orderStatus: "cancelled", updatedAt: new Date() } }
  )
  // Restore product stock
  db.products.updateMany(
    { productId: { $in: productIds } },
    { $inc: { productQuantity: quantity } }
  )
  

4. Update Payment Status
- API: PUT /api/orders/{orderId}/payment
- Lookup Key: orderId
- Pattern: Write heavy
- Use Case: Updating payment status after payment processing
- Query Pattern:
  db.orders.updateOne(
    { orderId: "12345" },
    { 
      $set: { 
        paymentStatus: "completed",
        paymentDate: new Date(),
        updatedAt: new Date()
      } 
    }
  )
  

5. Update Order Items
- API: PUT /api/orders/{orderId}/items
- Lookup Key: orderId, productId
- Pattern: Write heavy
- Use Case: Modifying items in an order
- Query Pattern:
  db.orders.updateOne(
    { orderId: "12345", "products.productId": "productId" },
    { 
      $set: { 
        "products.$.quantity": newQuantity,
        "products.$.subtotal": newSubtotal,
        orderTotal: newTotal,
        updatedAt: new Date()
      } 
    }
  )
  


Indexes 

Primary Index
- orderId
  - Primary key for all order operations

Secondary Indexes
1. userId + orderDate
   - Used for: View orders by user and date range
   - Command: db.orders.createIndex({ userId: 1, orderDate: -1 })

2. userId + orderStatus
   - Used for: View orders by user and status
   - Command: db.orders.createIndex({ userId: 1, orderStatus: 1 })

3. orderId + productId
   - Used for: Update order items
   - Command: db.orders.createIndex({ orderId: 1, "products.productId": 1 })
