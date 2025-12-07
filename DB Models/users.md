User Entity : 

{
    "userId": "123",
    "name": "Ayush",
    "email": "ayush@rapido.bike",
    "password": "123456",
    "role": "admin",
    "active": true,
    "addressId": "add123",
    "createdAt": "2025-11-28",
    "updatedAt": "2025-11-28"
}

Read Patterns & APIs : 

1. Get User by Mobile Number/Email
- API: GET /api/users?email={email} or GET /api/users?phone={phone}
- Lookup Key: email or phone
- Pattern: Non read-heavy
- Use Case: User lookup for authentication, profile access
- Query Pattern: 
  db.users.findOne({ email: "ayush@rapido.bike" })
  // or
  db.users.findOne({ phone: "6386709252" })

2. Get User for Delivery Confirmation
- API: GET /api/users/{userId}
- Lookup Key: userId
- Pattern: Read heavy (during delivery operations)
- Use Case: Fetching user details for delivery confirmation
- Query Pattern:
  db.users.findOne({ userId: "123" })

3. Get User Profile
- API: GET /api/users/{userId}/profile
- Lookup Key: userId
- Pattern: Non read-heavy
- Use Case: Displaying user profile information
- Query Pattern:
  db.users.findOne({ userId: "123" })

4. User Logout
- API: POST /api/users/{userId}/logout
- Lookup Key: userId
- Pattern: Non read-heavy
- Use Case: Logging out user session
- Query Pattern:
  db.users.findOne({ userId: "123" })

5. User Login
- API: POST /auth/login
- Lookup Key: phone or email
- Pattern: Read heavy
- Use Case: User authentication/login
- Query Pattern:
  db.users.findOne({ phone: "6386709252" })
  // or
  db.users.findOne({ email: "ayush@rapido.bike" })


Write Patterns & APIs :

1. User Registration
- API: POST /api/users/register
- Lookup Key: email (unique check)
- Pattern: Non write-heavy (unless campaign/advertisement active)
- Use Case: Creating new user account
- Query Pattern:
  // Check uniqueness
  db.users.findOne({ email: "newuser@example.com" })
  // Insert if not exists
  db.users.insertOne({ userId, name, email, password, ... })

2. Update Profile Information
- API: PUT /api/users/{userId}/profile
- Lookup Key: userId
- Pattern: Non write-heavy
- Use Case: Updating user profile details
- Query Pattern:
  db.users.updateOne(
    { userId: "123" },
    { $set: { name: "Updated Name", updatedAt: new Date() } }
  )

3. Password Reset
- API: POST /api/users/reset-password
- Lookup Key: userId or email
- Pattern: Non write-heavy
- Use Case: Resetting user password
- Query Pattern:
  db.users.updateOne(
    { email: "user@example.com" },
    { $set: { password: "newHashedPassword", updatedAt: new Date() } }
  )

4. Deactivate User
- API: PUT /api/users/{userId}/deactivate
- Lookup Key: userId
- Pattern: Non write-heavy
- Use Case: Deactivating user account
- Query Pattern:
  db.users.updateOne(
    { userId: "123" },
    { $set: { active: false, updatedAt: new Date() } }
  )

5. Authenticate User
- API: POST /api/users/authenticate
- Lookup Key: email, password
- Pattern: Non write-heavy (read operation for authentication)
- Use Case: User login/authentication
- Query Pattern:
  db.users.findOne({ 
    email: "user@example.com",
    password: "hashedPassword" 
  })


6. Update Status (Active / Blocked)
- API: PUT /users/{userId}/status
- Lookup Key: phone
- Pattern: Non write-heavy
- Use Case: Admin changing user status by phone number
- Query Pattern:
  db.users.updateOne(
    { phone: "6386709252" },
    { $set: { status: "active" } }
  )


Indexes : 

Primary Index
- userId
  - Primary key for all user operations

Secondary Indexes
1. email
   - Used for: Login and user updates
   - Command: db.users.createIndex({ email: 1 })

2. phone
   - Used for: Login and user updates
   - Command: db.users.createIndex({ phone: 1 })

3. role
   - Used for: Admin dashboard filtering
   - Command: db.users.createIndex({ role: 1 })

4. status
   - Used for: Admin dashboard filtering
   - Command: db.users.createIndex({ status: 1 })
