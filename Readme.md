# MyFlat - Signup

## Description

This is the Signup Microservice for the MyFlat project. It is responsible for handling the signup process for new users.

## Data Model V1.0

The data model is as follows:

```json
{
  "id": "string",
  "username": "string",
  "email": "string",
  "password_hash": "string",
  "createdAt": "string",
  "updatedAt": "string",
  "verified_at": "string",
  "is_verified": "boolean"
}
```

Later versions of the data model will include more fields, in order to handle auth providers and other features. Also auth_providers will have its own collection.

## Endpoints

### POST /signup

This endpoint is used to create a new user.

#### Request Body

```json
{
  "email": "string",
  "password": "string"
}
```

### Behavior

1. The service will check if the email is already in use.
2. If the email is not in use, the service will create a new user with the provided email and password, storing the password hash and the encrypted email.
3. The service will generate a random username for the user that is not already in use.
4. The service will return the username

### Response

```json
{
  "username": "string"
}
```

