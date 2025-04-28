# 📋 Problem Statement

Flipkart's **E-commerce Platform** is undergoing a **modernization project** to migrate its legacy **product catalog system** to a new **microservices-based architecture**.

### Current Legacy System Responsibilities:
- Manages a **massive inventory database**.  
- Serves customer-facing product details like **descriptions**, **pricing**, and **availability**.  
- Has become increasingly **difficult to maintain** due to outdated code.  

The goal is to gradually migrate the functionality from the **legacy system** to the **new microservices architecture** without disrupting the ongoing e-commerce operations.

---

## ❓ Question

Design a solution using the **Strangler Pattern** that:

✅ Enables **incremental migration** of the product catalog system to a modern architecture.  
✅ Ensures **existing functionality** continues to operate during migration.  
✅ Routes certain **API requests** to the legacy system and others to the new system.  
✅ Allows future **product data updates** to be handled by the new system.  
✅ Ensures a **seamless switch-over** from the legacy system to the new system when the migration is complete.  

---

## 🧩 Solution Overview

The solution will utilize:

✅ **Strangler Pattern** to gradually replace the old system with the new **microservices architecture**.  
✅ A **Proxy (Router) Layer** that intelligently directs requests to either the **legacy system** or the **new system** based on feature availability.  
✅ Uses **feature flags** or **endpoint routing** to control the migration flow.  
✅ Ensures **backward compatibility** to avoid disruptions during the transition.  

---

## 📚 Key Takeaways

✅ Uses the **Strangler Pattern** to support **incremental migration** without disrupting ongoing operations.  
✅ Introduces a **Router Layer** that directs requests to either the **legacy** or **new system** based on product data availability.  
✅ Ensures **seamless transition** by enabling partial functionality in the new system while still supporting the old system.  
✅ Supports **feature toggles** or **routing logic** for easier control over the migration flow.  
✅ Ensures improved **code maintainability** by reducing technical debt over time.  

---

## 💬 Bonus Challenge

🔹 Implement a **caching mechanism** to improve performance during the migration phase.  
🔹 Introduce a **monitoring system** to track performance differences between the two systems.  
🔹 Add a **data synchronization process** to ensure the legacy and new systems stay in sync during the transition.  
🔹 Implement **rollback mechanisms** to switch back to the legacy system in case of migration issues.  
