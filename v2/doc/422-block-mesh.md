# Routing and Block Mesh

### **1. Decentralized Routing**:
- The routing mechanism is decentralized, meaning that each node in the network can route messages independently, without reliance on a central authority.
- **Source Routing**: There is flexibility in allowing the originator of the message to specify intermediate nodes, giving nodes the option of choosing their own routes based on their needs.
- Routing choices take into account factors like **latency**, **reputation**, and **response time**, and can be influenced by **market dynamics**.

### **2. Personal Currency and Reputation**:
- The system integrates a reputation mechanism where nodes issue their own **personal currency**, akin to promises. Reputation is **relative** and depends on the **exchange rate** of a node's personal currency relative to other currencies.
- Nodes earn reputation points or currency by fulfilling promises, and their **currency’s exchange rate** reflects how much trust or value other nodes place on their output.
  
### **3. Exchange Rate-Based Routing**:
- We explored routing strategies where nodes may choose routes based on the **exchange rates** of other nodes’ currencies. The idea is that routing could be influenced by which nodes offer the most valuable currency relative to the sender's currency.
- Nodes could send periodic queries asking for exchange rate bids/asks, and routes may favor higher-value currencies.

### **4. Hybrid Matching**:
- There may be a **hybrid system** that combines routing via address resolution protocols (e.g., ARC-style) and reputation-based market forces.
- **Usenet-style bang path routing** and **market forces** may also play a role in how routing decisions are made, with nodes advertising the sequences they can complete and selecting routes accordingly.

### **5. Atomicity and Merge Transactions**:
- We shifted from thinking of ownership transfers to viewing transactions as **bid and ask orders**, akin to an open market. Nodes post orders on unit-specific ledgers, and a third party (like a matchmaker) matches orders and coordinates atomic swaps.
- The **swap transaction** is treated as a merge operation similar to a version control system, where the swap includes the most recent transaction hashes from both unit ledgers. This ensures **atomicity** and prevents one-sided swaps.

### **6. Unit-Specific Ledgers**:
- Each currency unit (identified by a serial number) has its own **unit ledger**, which tracks ownership and transactions for that unit.
- Transactions are handled through **merge transactions** that reference the most recent transactions on both ledgers, ensuring an atomic and verifiable swap.
- This model eliminates the need for a central ledger, making each currency unit ledger independent and decentralized.
- Thinking of unit ledgers as being akin to version control branches, and atomic transactions as merges between branches, and subsequent ledger entries as branches off the merge commit, leads to a concept of a **block mesh** of unit ledgers.

### **7. Gas Charges for Routing**:
- We discussed introducing **gas charges** to the system, where kernels charge modules for the resources consumed in routing, such as CPU cycles and message traffic. This mechanism could ensure that heavily used resources are properly compensated.

### **8. Dynamic Pricing and Incentives**:
- Kernels can dynamically adjust the pricing (gas fees) for routing and other services based on their available resources and demand. This creates a self-regulating economy where well-performing nodes (kernels) with high demand can charge higher fees.

### **9. Capabilities and Currency**:
- Capabilities were also explored as a form of currency. A node's **capabilities** (such as handling particular message sequences or offering specific services) might be issued in the same way as currency, where **capabilities are a form of promises** that can be exchanged.
  
### **Conclusion**:
PromiseGrid's routing design is built around **decentralized, market-driven mechanisms** where routing decisions depend on the **value of personal currencies, promises, and node capabilities**. Nodes make routing choices based on their trust in other nodes, as measured by exchange rates, and the system ensures atomicity through merge transactions that reference the latest state of unit-specific ledgers. Overall, the system is flexible, decentralized, and driven by economic and reputation-based incentives.

