## 🧪 Expected Output

```
[✅ Transaction 1] Payment Processed Successfully!
✅ Circuit Breaker RESET to CLOSED after successful retry.
[✅ Transaction 2] Payment Processed Successfully!
✅ Circuit Breaker RESET to CLOSED after successful retry.
[❌ Transaction 3] Failed: Payment Service Failed
[✅ Transaction 4] Payment Processed Successfully!
✅ Circuit Breaker RESET to CLOSED after successful retry.
[❌ Transaction 5] Failed: Payment Service Failed
[❌ Transaction 6] Failed: Payment Service Failed
[❌ Transaction 7] Failed: Payment Service Failed
🚨 Circuit Breaker OPEN: Blocking requests to prevent overload.
[❌ Transaction 8] Request blocked due to Circuit Breaker OPEN state.
[❌ Transaction 9] Request blocked due to Circuit Breaker OPEN state.
[❌ Transaction 10] Request blocked due to Circuit Breaker OPEN state.
[❌ Transaction 11] Request blocked due to Circuit Breaker OPEN state.
[❌ Transaction 12] Request blocked due to Circuit Breaker OPEN state.
[❌ Transaction 13] Request blocked due to Circuit Breaker OPEN state.
[❌ Transaction 14] Request blocked due to Circuit Breaker OPEN state.
[❌ Transaction 15] Request blocked due to Circuit Breaker OPEN state.
✅ All transactions processed successfully.
```