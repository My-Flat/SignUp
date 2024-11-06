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
  "id_type": "string",
  "id_number": "string", // For later verification if someone registered throught more than one method, then link the accounts
  "password_hash": "string",
  "is_verified": "boolean",
  "name": "string",
  "surname": "string",
  "phone": "string",
  "address": "string",
  "city": "string",
  "zip": "string",
  "country": "string",
  "createdAt": "string",
  "updatedAt": "string"
}
```

Later versions of the data model will include more fields, in order to handle auth providers and other features.