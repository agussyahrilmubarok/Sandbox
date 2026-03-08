### 1. **Sign Up (Register)**

**Endpoint:** `POST /api/v1/auth/signup`

**cURL Request:**
```bash
curl -X POST http://localhost:8000/api/v1/auth/signup \
-H "Content-Type: application/json" \
-d '{
    "username": "johndoe123",
    "email": "johndoe123@mail.com",
    "password": "securePassword123"
}'
```

**Expected Response (on Success):**
```json
{
  "username": "johndoe123",
  "email": "johndoe123@mail.com",
  "avatar": "default-user.png",
  "createdAt": "2024-12-06T10:38:29.262Z",
  "updatedAt": "2024-12-06T10:38:29.262Z"
}
```

**Explanation:**
- This request will create a new user with the provided `username`, `email`, and `password`.
- If the `email` or `username` already exists, the server will return a `422` status with an error message.

---

### 2. **Sign In (Login)**

**Endpoint:** `POST /api/v1/auth/signin`

**cURL Request:**
```bash
curl -X POST http://localhost:8000/api/v1/auth/signin \
-H "Content-Type: application/json" \
-d '{
    "email": "johndoe123@mail.com",
    "password": "securePassword123"
}'
```

**Expected Response (on Success):**
```json
{
  "username": "johndoe123",
  "email": "johndoe123@mail.com",
  "avatar": "default-user.png",
  "createdAt": "2024-12-06T10:38:29.262Z",
  "updatedAt": "2024-12-06T10:38:29.262Z"
}
```

**Explanation:**
- This request will authenticate the user using the provided `email` and `password`.
- If the user is found and the password is correct, the server will return the user's data along with a JWT token in a `taskly_token` cookie.
- If the credentials are incorrect or the user is not found, you will get a `404` or `401` response.

---

### 3. **Sign Out (Logout)**

**Endpoint:** `GET /api/v1/auth/signout`

**cURL Request:**
```bash
curl -X GET http://localhost:8000/api/v1/auth/signout \
--cookie "taskly_token=your_jwt_token_here"
```

**Expected Response (on Success):**
```json
{
  "message": "Sign out successful"
}
```

**Explanation:**
- This request will clear the `taskly_token` cookie, effectively logging out the user.

---

### 4. **Sign Up Failure (Email or Username Already Taken)**

**Endpoint:** `POST /api/v1/auth/signup`

**cURL Request:**
```bash
curl -X POST http://localhost:8000/api/v1/auth/signup \
-H "Content-Type: application/json" \
-d '{
    "username": "johndoe123",   // Username already taken
    "email": "johndoe123@mail.com",  // Email already taken
    "password": "anotherPassword123"
}'
```

**Expected Response (on Failure):**
```json
{
  "status": 422,
  "message": "Email or Username is already registered."
}
```

**Explanation:**
- If the `email` or `username` already exists in the database, the server will return a `422` status code with an appropriate error message.

---

### 5. **Sign In Failure (Wrong Password)**

**Endpoint:** `POST /api/v1/auth/signin`

**cURL Request:**
```bash
curl -X POST http://localhost:8000/api/v1/auth/signin \
-H "Content-Type: application/json" \
-d '{
    "email": "johndoe123@mail.com",
    "password": "wrongPassword123"  // Incorrect password
}'
```

**Expected Response (on Failure):**
```json
{
  "status": 401,
  "message": "Wrong password!"
}
```

**Explanation:**
- If the provided password does not match the stored hashed password, the server will return a `401` status code with a "Wrong password!" message.

---

### Notes:
- Make sure your **server** is running on the correct port (e.g., `localhost:8000`) and that MongoDB is connected and operational.
- If you are using **HTTPS** in production, ensure your requests are sent to `https://localhost:8000` (or your production URL) instead of `http://`.
- The `taskly_token` will be stored as an **HTTP-only cookie** in the browser, and you can use it to authenticate subsequent requests that require user authorization.

### Optional: Checking the Cookies
You can check if the cookie is properly set using the following command:

```bash
curl -X POST http://localhost:8000/api/v1/auth/signin \
-H "Content-Type: application/json" \
-d '{
    "email": "johndoe123@mail.com",
    "password": "securePassword123"
}' \
-c cookies.txt  # This saves the cookies in a file
```

To use the cookies in a subsequent request:

```bash
curl -X GET http://localhost:8000/api/v1/tasks \
-b cookies.txt  # This sends the saved cookies with the request
```