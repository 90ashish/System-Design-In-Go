## 📋 Problem Description

Implements the **Abstract Factory Design Pattern** to create **sports products** (shoes and shirts) for different brands — specifically **Adidas** and **Nike**. The **Abstract Factory Pattern** is ideal for creating families of related objects without specifying their concrete classes.

---

## ❓ What Does This Code Do?

- Provides a **flexible** and **scalable** solution for producing **Adidas** and **Nike** branded sports products like shoes and shirts.  
- Utilizes the **Abstract Factory Pattern** to dynamically create appropriate brand-specific products.  
- Demonstrates **polymorphism** by providing a **common interface** to interact with different product variants.  

---

## 🧩 Design Pattern Explanation — Abstract Factory

The **Abstract Factory Pattern** is used to create families of related products without specifying their concrete classes.

---

## 🔧 Key Components in the Code

1. **Abstract Factory Interface (`ISportsFactory`)**  
   - Declares methods for creating abstract product types like `makeShoe()` and `makeShirt()`.  

2. **Concrete Factories (`Adidas`, `Nike`)**  
   - Implement the `ISportsFactory` interface to produce brand-specific products.  

3. **Abstract Product Interfaces (`IShoe`, `IShirt`)**  
   - Define methods for setting and retrieving product properties like **logo** and **size**.  

4. **Concrete Products (`AdidasShoe`, `NikeShoe`, `AdidasShirt`, `NikeShirt`)**  
   - Embed a common base structure (Shoe or Shirt) while applying brand-specific configurations.  

5. **Client Code (`main()`)**  
   - Uses the `GetSportsFactory()` method to obtain the appropriate factory instance.  
   - Uses each factory to create respective products (**Adidas** or **Nike**).  

---

## 🏗️ Flow of Execution

1. The `main()` function calls:  
   - `GetSportsFactory("adidas")` → Returns Adidas Factory  
   - `GetSportsFactory("nike")` → Returns Nike Factory  

2. Each factory creates both shoes and shirts using their respective methods:  
   - `adidasFactory.makeShoe()` → **Adidas Shoe**  
   - `adidasFactory.makeShirt()` → **Adidas Shirt**  
   - `nikeFactory.makeShoe()` → **Nike Shoe**  
   - `nikeFactory.makeShirt()` → **Nike Shirt**  

3. The products are printed using helper functions — `printShoeDetails()` and `printShirtDetails()`.  

---

## 🧪 Expected Output

Shoe Logo: nike
Shoe Size: 14
Shirt Logo: nike
Shirt Size: 14
Shoe Logo: adidas
Shoe Size: 14
Shirt Logo: adidas
Shirt Size: 14


---

## 📚 Key Takeaways

✅ Implements the **Abstract Factory Pattern** effectively to produce related product families.  
✅ Ensures **loose coupling** — client code interacts only with the factory interface, not the concrete implementations.  
✅ Ensures **scalability** — adding new brands like **Puma**, **Reebok**, etc., can be done without modifying the existing logic.  
✅ Promotes **code maintainability** by separating the logic for creating related product families.  

---

## 💬 Bonus Challenge

🔹 Add a **Puma** or **Reebok** factory to demonstrate scalability.  
🔹 Implement additional product categories like **Caps** or **Jackets** following the same pattern.  
🔹 Introduce **product discounts** or **customization features** to enhance flexibility.  
