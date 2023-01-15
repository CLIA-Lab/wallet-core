# Wallet Core Service
Base core for any type of service requiring wallet functionality

## Blockchain

- Bitcoin
- Ethereum
- Near
- Cardano

## Functionalities

- GetWallet() -> (pubK, privK): Generate pair public-private key
- SendTrx(from, to, amount) -> bool: Send a transaction and return if it was possible (true) or not (false)
- GetBalance(address) ->string: Get the balance of an address and return a big number representing the amount 
