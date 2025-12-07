Product Entity : 

{
    "productId": "123456",
    "productName": "Tennis Bat",
    "description": "Cricket bat",
    "price": 1000,
    "discount": 10,
    "productQuantity": 100,
    "productImage": "product Image",
    "categoryId": "cat1",
    "subCategoryId": "subcat1",
    "brand": "Brand 1",
    "color": "Black",
    "isActive": true,
    "createdAt": "2025-11-28",
    "updatedAt": "2025-11-28"
}

Read Patterns & APIs

1. Browse Products
- API: GET /api/products?page={page}&limit={limit}
- Lookup Key: None (scan all active products)
- Pattern: Read heavy
- Use Case: Displaying all products in catalog
- Query Pattern:
  db.products.find({ isActive: true })
    .skip((page - 1) * limit)
    .limit(limit)
    .sort({ createdAt: -1 })
  

2. Search Products by Name/Category
- API: GET /api/products/search?q={query}&categoryId={categoryId}
- Lookup Key: productName, categoryId
- Pattern: Read heavy
- Use Case: Product search functionality
- Query Pattern:
  db.products.find({
    $or: [
      { productName: { $regex: query, $options: "i" } },
      { categoryId: categoryId }
    ],
    isActive: true
  })
  

3. View Product Details
- API: GET /api/products/{productId}
- Lookup Key: productId
- Pattern: Non read-heavy
- Use Case: Displaying individual product details
- Query Pattern:
  db.products.findOne({ productId: "123456" })
  

4. Filter Products by Category/Brand/Price
- API: GET /api/products/filter?categoryId={cat}&brand={brand}&minPrice={min}&maxPrice={max}
- Lookup Key: categoryId, brand, price range
- Pattern: Read heavy
- Use Case: Product filtering and browsing
- Query Pattern:
  db.products.find({
    categoryId: "cat1",
    brand: "Brand 1",
    price: { $gte: minPrice, $lte: maxPrice },
    isActive: true
  })
  

5. View Product Availability/Quantity
- API: GET /api/products/{productId}/availability
- Lookup Key: productId
- Pattern: Non read-heavy
- Use Case: Checking product stock availability
- Query Pattern:
  db.products.findOne(
    { productId: "123456" },
    { productId: 1, productQuantity: 1, isActive: 1 }
  )
  
  

Write Patterns & APIs

1. Add New Product
- API: POST /api/products
- Lookup Key: productId (unique check)
- Pattern: Non write-heavy
- Use Case: Adding new product to catalog
- Query Pattern:
  // Check uniqueness
  db.products.findOne({ productId: "newProductId" })
  // Insert if not exists
  db.products.insertOne({ productId, productName, price, ... })
  

2. Update Product Information
- API: PUT /api/products/{productId}
- Lookup Key: productId
- Pattern: Non write-heavy
- Use Case: Updating product details
- Query Pattern:
  db.products.updateOne(
    { productId: "123456" },
    { $set: { productName: "Updated Name", description: "...", updatedAt: new Date() } }
  )
  

3. Update Product Quantity/Stock
- API: PUT /api/products/{productId}/stock
- Lookup Key: productId
- Pattern: Write heavy
- Use Case: Updating inventory after orders/restocking
- Query Pattern:
  db.products.updateOne(
    { productId: "123456" },
    { $set: { productQuantity: newQuantity, updatedAt: new Date() } }
  )
  

4. Deactivate Product
- API: PUT /api/products/{productId}/deactivate
- Lookup Key: productId
- Pattern: Non write-heavy
- Use Case: Removing product from active catalog
- Query Pattern:
  db.products.updateOne(
    { productId: "123456" },
    { $set: { isActive: false, updatedAt: new Date() } }
  )
  

5. Update Product Price/Discount
- API: PUT /api/products/{productId}/pricing
- Lookup Key: productId
- Pattern: Non write-heavy
- Use Case: Updating product pricing or discounts
- Query Pattern:
  db.products.updateOne(
    { productId: "123456" },
    { $set: { price: newPrice, discount: newDiscount, updatedAt: new Date() } }
  )
  


Indexes 

Primary Index
- productId
  - Primary key for all product operations

Secondary Indexes
1. categoryId + subCategoryId
   - Used for: Filter products by category and subcategory together
   - Command: db.products.createIndex({ categoryId: 1, subCategoryId: 1 })

2. productName
   - Used for: Search products by name
   - Command: db.products.createIndex({ productName: 1 })

3. categoryId + brand
   - Used for: Filter products by category and brand together
   - Command: db.products.createIndex({ categoryId: 1, brand: 1 })

4. categoryId + price
   - Used for: Filter products by category and price range
   - Command: db.products.createIndex({ categoryId: 1, price: 1 })

5. isActive + categoryId
   - Used for: Browse active products in a category
   - Command: db.products.createIndex({ isActive: 1, categoryId: 1 })
