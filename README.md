# Midtrans Payment Gateway Implementation using Golang

## Description
This implementation demonstrates how to use Midtrans as a Payment Gateway in your Golang project. Midtrans provides an API that allows you to handle payments easily and securely.

##  Getting Midtrans Client Key
Before you begin, make sure you have a Midtrans account. If you don't have one, you can sign up at [Midtrans Dashboard](https://dashboard.midtrans.com/).

1. **Login to Midtrans Dashboard**

    Log in to your Midtrans Dashboard using your credentials. This is where you will manage your Midtrans account and access the necessary configurations.

2. **Retrieve Sandbox Server Key**

    In the Dashboard you can select the sanbox or production environment, then navigate to the **Settings** section. In this case I choose sanbox because it was for testing purposes. You may find it in the sidebar or top navigation, depending on the Dashboard version. Look for options like "API Keys" or "Access Keys". Here, you will find your **Sandbox Server Key**. It usually starts with "SB-Mid-server-..." or similar.

    Or you can also click this link [Midtrans Sanbox](https://dashboard.sandbox.midtrans.com/settings/config_info)

    **Note:** Understand the key differences between using the sandbox and production environments. Sandbox is meant for development and testing without involving real money. In contrast, the production environment is used for real transactions with actual financial implications.

3. **Copy Your Sandbox Server Key**

    Copy your Sandbox Servey Key to your env.

## Installation
1. Install the Midtrans Go library in your project:

    ```go
    go get -u github.com/midtrans/midtrans-go
    ```

2. Import the Midtrans library in your Golang code:

    ```go
    import (
        "github.com/midtrans/midtrans-go"
        "github.com/midtrans/midtrans-go/snap"
    )
    ```

3. Set up the Midtrans configuration using the server key obtained from the Midtrans Dashboard:

    ```go
    snapClient := &snap.Client{}
    snapClient.New("your_server_key", midtrans.Sandbox)
    ```

## Creating a Payment
```go
// Example payment creation
chargeReq := &midtrans.ChargeReq{
    PaymentType: midtrans.SourceCreditCard,
    TransactionDetails: midtrans.TransactionDetails{
        OrderID:  "ORDER-ID-12345",
        GrossAmt: 200000,
    },
    CreditCard: &midtrans.CreditCardDetails{
        TokenID: "credit_card_token",
        SaveCard: true,
    },
}

chargeResp, err := snapClient.CreateTransaction(chargeReq)
if err != nil {
    // Handle error
    fmt.Println("Error:", err)
    return
}
```
