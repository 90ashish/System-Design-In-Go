## 🧪 Expected Output :

```
2025/03/23 01:22:42 [DataAnalytics] Bulkhead full for request 'Req6'. Executing fallback.
2025/03/23 01:22:42 Fallback: Analytics deferred for: Req6
2025/03/23 01:22:42 [DataAnalytics] Bulkhead full for request 'Req2'. Executing fallback.
2025/03/23 01:22:42 Fallback: Analytics deferred for: Req2
2025/03/23 01:22:42 [EmailNotification] Bulkhead full for request 'Req4'. Executing fallback.
2025/03/23 01:22:42 [EmailNotification] Bulkhead full for request 'Req6'. Executing fallback.
2025/03/23 01:22:42 Fallback: Email will be retried later: Req4
2025/03/23 01:22:42 [EmailNotification] Bulkhead full for request 'Req8'. Executing fallback.
2025/03/23 01:22:42 Fallback: Email will be retried later: Req8
2025/03/23 01:22:42 [EmailNotification] Bulkhead full for request 'Req7'. Executing fallback.
2025/03/23 01:22:42 [DocumentProcessing] Bulkhead full for request 'Req4'. Executing fallback.
2025/03/23 01:22:42 [DocumentProcessing] Bulkhead full for request 'Req5'. Executing fallback.
2025/03/23 01:22:42 Fallback: Document queued for later processing: Req5
2025/03/23 01:22:42 [DocumentProcessing] Bulkhead full for request 'Req8'. Executing fallback.
2025/03/23 01:22:42 Fallback: Document queued for later processing: Req4
2025/03/23 01:22:42 [DataAnalytics] Bulkhead full for request 'Req3'. Executing fallback.
2025/03/23 01:22:42 Fallback: Analytics deferred for: Req3
2025/03/23 01:22:42 Fallback: Document queued for later processing: Req8
2025/03/23 01:22:42 Fallback: Email will be retried later: Req6
2025/03/23 01:22:42 Fallback: Email will be retried later: Req7
2025/03/23 01:22:42 [DataAnalytics] Bulkhead full for request 'Req1'. Executing fallback.
2025/03/23 01:22:42 Fallback: Analytics deferred for: Req1
2025/03/23 01:22:42 [EmailNotification] Bulkhead full for request 'Req2'. Executing fallback.
2025/03/23 01:22:42 Fallback: Email will be retried later: Req2
2025/03/23 01:22:42 [EmailNotification] Bulkhead full for request 'Req5'. Executing fallback.
2025/03/23 01:22:42 Fallback: Email will be retried later: Req5
2025/03/23 01:22:42 [DataAnalytics] Bulkhead full for request 'Req8'. Executing fallback.
2025/03/23 01:22:42 Fallback: Analytics deferred for: Req8
2025/03/23 01:22:42 [DocumentProcessing] Bulkhead full for request 'Req6'. Executing fallback.
2025/03/23 01:22:42 Fallback: Document queued for later processing: Req6
2025/03/23 01:22:42 [DocumentProcessing] Bulkhead full for request 'Req7'. Executing fallback.
2025/03/23 01:22:42 Fallback: Document queued for later processing: Req7
2025/03/23 01:22:42 [DataAnalytics] Bulkhead full for request 'Req7'. Executing fallback.
2025/03/23 01:22:42 Fallback: Analytics deferred for: Req7
2025/03/23 01:22:42 Document processed: Req1
2025/03/23 01:22:42 Email sent: Req3
2025/03/23 01:22:42 Document processed: Req2
2025/03/23 01:22:42 Document processed: Req3
2025/03/23 01:22:42 Analytics complete for: Req5
2025/03/23 01:22:42 Analytics complete for: Req4
2025/03/23 01:22:42 Email sent: Req1
2025/03/23 01:22:44 All requests processed.
```