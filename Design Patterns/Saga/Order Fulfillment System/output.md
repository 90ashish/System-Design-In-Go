## 🧪 Expected Output :

```
>>> Processing Successful Order

Starting Order Fulfillment Process...
[Payment Service] Payment processed successfully for: Order#123
[Inventory Service] Inventory reserved for: Order#123
[Shipping Service] Shipment arranged for: Order#123
[Notification Service] Customer notified for: Order#123
✅ Order Fulfillment Process Completed Successfully!

>>> Processing Failed Order

Starting Order Fulfillment Process...
[Payment Service] Payment processed successfully for: Order#456
[Inventory Service] Inventory reservation FAILED for: Order#456
[ERROR] Transaction failed. Initiating rollback...
[Payment Service] Payment rollback initiated for: Order#456
[Notification Service] Notification rollback initiated for: Order#456
[Shipping Service] Shipment cancellation initiated for: Order#456
[Inventory Service] Inventory release initiated for: Order#456
[Payment Service] Payment rollback initiated for: Order#456
❌ Order Fulfillment Rolled Back Successfully!
```