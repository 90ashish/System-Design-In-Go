## 🧪 Expected Output :

```
=== Successful Ride Booking ===
[RideBookingService] Ride RIDE001 booked successfully.
[DriverAssignmentService] Driver assigned for Ride RIDE001.
[PaymentService] Payment of $20.00 for Ride RIDE001 processed successfully.
[NotificationService] Notification Sent: Your ride has been confirmed! (Ride RIDE001)
✅ Ride booking completed successfully.

=== Failed Scenario: Payment Service Failure ===
[RideBookingService] Ride RIDE002 booked successfully.
[DriverAssignmentService] Driver assigned for Ride RIDE002.
[ERROR] Payment processing failed. Rolling back...
[DriverAssignmentService] Driver unassigned for Ride RIDE002.
[RideBookingService] Booking for Ride RIDE002 canceled.
```