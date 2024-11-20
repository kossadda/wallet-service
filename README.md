# Wallet Management System

## Overview

This project is a wallet management system that allows users to deposit and withdraw funds from their wallets. The system is designed to handle high concurrency (up to 1000 `RPS`) and ensures that all requests are processed without `50X` errors.

**Application Requirements**:
- must be written in `Golang`
- must use `PostgreSQL` as the database
- must be containerized using `Docker`
- system deployment must be done using `Docker Compose`
- environment variables must be read from a `config.env` file

## Features

- **Deposit and Withdraw Funds**: Users can deposit or withdraw funds from their wallets.
- **Check Balance**: Users can check the current balance of their wallets.
- **High Concurrency Handling**: The system is designed to handle high concurrency without 50X errors.
- **Containerization**: The application and database are containerized using Docker and can be deployed using Docker Compose.

## API Endpoints

### Deposit or Withdraw Funds

- **Endpoint**: `POST /api/v1/wallet`
- **Request Body**:
  ```json
  {
    "walletId": "UUID",
    "operationType": "DEPOSIT", // or "WITHDRAW"
    "amount": 1000
  }

- **Description**: This `endpoint` allows users to deposit or withdraw funds from their wallets.

### Check Wallet Balance

- **Endpoint**: `GET /api/v1/wallets/{WALLET_UUID}`
- **Description**: This `endpoint` allows users to check the current balance of their wallets.

## Technology Stack

- **Backend**: `Golang`
- **Database**: `PostgreSQL`
- **Containerization**: `Docker`
- **Orchestration**: `Docker Compose`
