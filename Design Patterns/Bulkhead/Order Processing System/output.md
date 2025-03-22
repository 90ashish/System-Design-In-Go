## 🧪 Expected Output :

```
[❌ Inventory Service] Request 3 rejected due to overload
[❌ Payment Service] Request 5 rejected due to overload
[❌ Inventory Service] Request 5 rejected due to overload
[❌ Payment Service] Request 4 rejected due to overload
[❌ Inventory Service] Request 4 rejected due to overload
[✅ Notification Service] Sent confirmation for Request 5
[✅ Payment Service] Processed Payment Request 1
[✅ Notification Service] Sent confirmation for Request 3
[✅ Notification Service] Sent confirmation for Request 4
[✅ Payment Service] Processed Payment Request 2
[✅ Notification Service] Sent confirmation for Request 2
[✅ Notification Service] Sent confirmation for Request 1
[✅ Inventory Service] Updated stock for Request 1
[✅ Inventory Service] Updated stock for Request 2
[✅ Payment Service] Processed Payment Request 3
✅ All requests have been processed with Bulkhead Isolation.
```