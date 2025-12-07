Cart Entity : 

{
    "cartId": "cart1",
    "userId": "123",
    "products": [
        {
            "productId": "123456",
            "quantity": 1,
            "unitPrice": 100,
            "subtotal": 100
        },
        {
            "productId": "12345",
            "quantity": 5,
            "unitPrice": 500,
            "subtotal": 2500
        }
    ]
}

Read Patterns & APIs 

1. View Cart
- API: GET /api/users/{userId}/cart
- Lookup Key: userId
- Pattern: Read heavy
- Use Case: Displaying user's shopping cart
- Query Pattern:
  db.carts.findOne({ userId: "123" })
  

2. View Cart Items
- API: GET /api/users/{userId}/cart/items
- Lookup Key: userId, cartId
- Pattern: Read heavy
- Use Case: Getting list of items in cart
- Query Pattern:
  db.carts.findOne({ userId: "123", cartId: "cart1" }, { products: 1 })
  

Write Patterns & APIs

1. Add Item to Cart
- API: POST /api/users/{userId}/cart/items
- Lookup Key: userId, cartId, productId
- Pattern: Write heavy
- Use Case: Adding product to shopping cart
- Query Pattern:
  // Add new product to cart
  db.carts.updateOne(
    { userId: "123", cartId: "cart1" },
    { $push: { products: { productId: "123456", quantity: 1, unitPrice: 100, subtotal: 100 } } }
  )
  

2. Update Item Quantity in Cart
- API: PUT /api/users/{userId}/cart/items/{productId}
- Lookup Key: userId, cartId, productId
- Pattern: Write heavy
- Use Case: Updating quantity of item in cart
- Query Pattern:
  db.carts.updateOne(
    { userId: "123", cartId: "cart1", "products.productId": "123456" },
    { 
      $set: { 
        "products.$.quantity": newQuantity,
        "products.$.subtotal": newQuantity * unitPrice
      }
    }
  )
  

3. Remove Item from Cart
- API: DELETE /api/users/{userId}/cart/items/{productId}
- Lookup Key: userId, cartId, productId
- Pattern: Write heavy
- Use Case: Removing product from cart
- Query Pattern:
  db.carts.updateOne(
    { userId: "123", cartId: "cart1" },
    { $pull: { products: { productId: "123456" } } }
  )
  

4. Clear Cart
- API: DELETE /api/users/{userId}/cart
- Lookup Key: userId, cartId
- Pattern: Write heavy
- Use Case: Clearing all items from cart (after order creation)
- Query Pattern:
  db.carts.updateOne({ userId: "123", cartId: "cart1" }, { $set: { products: [] } })
  

Indexes

Primary Index
- cartId
  - Primary key for cart operations

Secondary Indexes
1. userId + cartId
   - Used for: All cart operations
   - Command: db.carts.createIndex({ userId: 1, cartId: 1 })

2. userId + cartId + productId
   - Used for: Add/update/remove item, update quantity
   - Command: db.carts.createIndex({ userId: 1, cartId: 1, "products.productId": 1 })
