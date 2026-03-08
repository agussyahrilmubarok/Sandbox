#!/bin/sh

# Sign up a new user
curl -X POST http://localhost:3000/api/users \
     -H "Content-Type: application/json" \
     -d '{"name": "John Doe", "email": "john@example.com", "password": "password123"}'

# Sign in
curl -X POST http://localhost:3001/api/v1/auth/signin \
     -H "Content-Type: application/json" \
     -d '{"email": "john@example.com", "password": "password123"}'

# Get list of users
curl -X GET http://localhost:3000/api/users \
     -H "Content-Type: application/json"

# Get a single user by ID
curl -X GET http://localhost:3000/api/users/{userId} \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer {TOKEN}"

# Update user by ID
curl -X PUT http://localhost:3000/api/users/{userId} \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer {TOKEN}" \
     -d '{"name": "John Updated"}'

# Delete user by ID
curl -X DELETE http://localhost:3000/api/users/{userId} \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer {TOKEN}"

# Sign out
curl -X GET http://localhost:3000/auth/signout
