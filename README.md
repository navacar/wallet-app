
# Wallet App

This is a simple Wallet App that allows you to store and manage the balances of multiple wallets. You can use the app to track and modify wallet balances via a set of APIs.

## API Endpoints

### 1. Get Wallet Balance

**GET** `/api/v1/wallets/{id}`

- **Description**: Retrieve the balance of a wallet with the specified ID `{id}`.
- **Response**: Returns the balance of the wallet.

### 2. Update Wallet Balance

**POST** `/api/v1/wallet`

- **Description**: Update the balance of a specific wallet.
- **Request Body** (JSON format):
  
  ```json
  {
      "walletId": {id},
      "operationType": "{operation_type}",
      "amount": {amount}
  }
  ```

  - `walletId` (integer): The ID of the wallet to update.
  - `operationType` (string): The type of operation (e.g., deposit, withdrawal).
  - `amount` (number): The amount to change the wallet balance by.

## How It Works

- This app allows you to perform CRUD operations on wallet balances.
- You can retrieve the balance of any wallet and update it by adding or subtracting a specified amount.
