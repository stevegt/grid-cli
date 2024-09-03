# PromiseGrid Routing

## Overview

Routing is the act of forwarding a message to the correct recipient(s) or to an intermediary host that can further route the message to its destination. 

In PromiseGrid, routing is a critical aspect of the network's operation, facilitating the exchange of messages, resources, and services between hosts and between the modules (apps) running on those hosts.

## Exchange Rate-Based Routing

### 1. Exchange Rate-Based Routing
- Hosts route based on the exchange rates of "personal currencies" --
  see below for more on personal currencies.
- Prefer routing to hosts with higher-valued currencies to ensure stability and reliability.
- Encourages hosts to maintain strong reputations and economic value.

### 2. Reputation and Currency Mechanisms
- Reputation points function as a currency, influencing routing decisions.
- Hosts decide to route based on the perceived value and reliability of other hosts.
- Market dynamics dictate routing paths, promoting decentralized control.

### 3. Handling Misbehaving Hosts
- Hosts that misbehave receive less network traffic and forwarding services.
- Poor behavior leads to reduced currency value, decreasing a host's network influence.
- This self-regulating mechanism ensures network stability and security.

## How Personal Currencies Work: "Everyone is Their Own Central Bank"

### Concept of Personal Currencies
- In PromiseGrid, **every host acts as its own central bank**, issuing its own currency in the form of reputation points.
- These currencies represent the host's trustworthiness and reliability, similar to how fiat currencies reflect the stability of a government.

### Incentives and Disincentives for Maintaining Currency Value
- Hosts are incentivized to maintain the value of their currency (reputation points) by keeping promises and behaving reliably.
- A **high currency value** means other hosts are more likely to route traffic through or trade with that host, increasing its influence and participation in the network.
- Conversely, **bad behavior** (breaking promises, excessive querying, etc.) leads to a **devaluation** of the currency, reducing a host’s ability to transact and interact with others effectively.

### Tools for Market Participants to Maintain Value
- **Reputation Management**: Hosts can actively manage their reputation by fulfilling promises and avoiding behaviors that could devalue their currency.
- **Dynamic Exchange Rates**: By adjusting their exchange rates with other hosts based on behavior and market conditions, hosts can strategically protect their currency's value.
- **Selective Interactions**: Hosts may choose to interact only with other hosts that maintain a stable or high-value currency, reinforcing trust and reliability in the network.

## Import and Export of Resources in PromiseGrid

### Understanding Import and Export
- In PromiseGrid, **importing resources** means paying another host for storage space or computational power. This could involve storing data on another host’s hardware or outsourcing computation tasks.
- **Exporting resources** refers to being paid by other hosts to provide storage space or computational power. Hosts that export resources effectively "sell" their excess capacity to others on the network.

### Analogies to International Trade
- The concepts of importing and exporting in PromiseGrid are similar to international trade in the global economy:
  - **Importing resources** (buying storage or computation) can be likened to a country purchasing goods or services from abroad. Hosts spend their currency to acquire necessary resources they cannot efficiently produce themselves.
  - **Exporting resources** (selling storage or computation) is akin to a country selling its goods or services to other nations. Hosts receive payments in other currencies, which strengthens their economic standing and increases their influence within the network.
- Just as countries aim to maintain a favorable balance of trade to enhance their currency’s value, hosts in PromiseGrid seek to optimize their import and export balance to maintain or increase their reputation and currency value.

### Governance and Geopolitical Factors Influencing Currency Values
- In the real world, geopolitical factors, governance, and economic policies heavily influence currency values. Similarly, in PromiseGrid:
  - **Governance Decisions**: Hosts might make decisions that affect their ability to import or export resources. For example, a host that reliably exports high-quality computational resources might see its currency increase in value, akin to how a country with a robust export economy strengthens its currency.
  - **Network Effects**: The behavior of a host's peers and its relationships within the network influence its currency's value. If a host is part of a trusted coalition or "alliance" of hosts that consistently fulfill promises, this could positively impact its currency, similar to how international alliances and trade agreements can stabilize or increase a nation’s currency value.
  - **Economic Policies**: Hosts might implement policies regarding how much they are willing to export or import based on current network conditions. This is analogous to a country setting tariffs or trade restrictions to protect its economy or encourage domestic production.

## Concept that "A Physical Host is Sovereign"

### Physical Host Sovereignty Analogy
- A physical host in PromiseGrid is considered sovereign, analogous to a geographical nation-state.
- This sovereignty implies that a host has ultimate control over its resources, policies, and interactions within its domain.

### Use of Force and Sovereignty
- Just as a nation-state can enforce its laws and defend its territory, a physical host can enforce its own rules and policies to maintain control -- ultimately, the host's system administrator can simply pull the plug.
- It's important to ensure the the host's owner/administrator is incentivized to play well with others (ensure the common good).
- We can do this by incentivizing the owner/administrator to maintain the host's reputation as expressed in its currency value.

## Evolution of Routing Strategies: From Concept to Current Model

### Early Considerations and Alternatives

We explored several other routing strategies, including:

#### Pay-for-Flow 

Hosts paying others in their currency to receive packets, effectively incentivizing traffic flow to specific routes.

#### Bond-Based Routing

Treating routing decisions as financial investments, where sending a message through a host could be seen as buying a bond in that host's currency.

Essentially this is still pay-for-flow, with deferred payment.

#### ZKOS 

Derived from Zero-Kernel Operating Systems, ZKOS routing used zero-knowledge proofs for privacy-preserving validation.  [https://web.mit.edu/ha22286/www/papers/MEng20_4.pdf](https://github.com/stevegt/grid-cli/blob/main/v2/doc/420-routing.md)

- **Zero-Knowledge Proofs (ZKPs)**: Used for privacy-preserving validation.
- **Decentralized Control**: hosts operate independently, ensuring security and resilience.

#### Antikernel

Paper by Andrew Zonenberg on the concept of moving much of the intra-kernel routing into
hardware: [https://www.iacr.org/archive/ches2016/98130227/98130227.pdf](https://www.iacr.org/archive/ches2016/98130227/98130227.pdf)

- **Modularity**: OS functionalities are divided into independent modules.
- **Hardware State Machines**: Utilizes hardware-based state machines for enhanced security and performance.

#### Capability Tokens

- **Token-Based Access Control**: Ensures that only authorized modules can handle specific messages.
- **Decentralized Capabilities**: Each module can issue and manage its capability tokens, promoting autonomy and security.

#### Hash-Based Routing

- **Hash-Based Addresses**: Each route is determined by hashing the first N bytes of the message. This can work because the grid's function is byte sequence completion.
- **Consistency and Collision Avoidance**: Uses cryptographic hashes to ensure unique and collision-resistant addresses.

#### Byte Sequence Completion

- **Trie Structure**: Uses a trie to manage and match byte sequences efficiently.
- **Dynamic Path Selection**: Adapts routing paths dynamically based on sequence completions and historical success rates.

#### Genetic Algorithms

- **Evolutionary Optimization**: Routes evolve through selection, crossover, and mutation, adapting to changing network conditions.
- **Fitness Function**: Evaluates routes based on performance metrics such as latency, throughput, and reliability.

### Machine Learning

- **Predictive Analytics**: Models analyze past routing decisions to forecast the best paths for new messages.
- **Reinforcement Learning**: The system learns optimal routes through reinforcement learning, continuously improving based on feedback.

## Shift to Exchange Rate Routing

- The concept of **exchange rate-based routing** emerged as a more straightforward alternative, leveraging existing market dynamics without requiring complex transaction management.
- In this model, hosts prefer routes that align with the highest-valued currencies, ensuring they interact with the most reliable and stable participants.
- This approach naturally encourages hosts to maintain a strong currency, as higher value directly translates to better routing opportunities and network influence.

### Considerations for the Current Model

- While the exchange rate-based routing simplifies decision-making, it introduces new considerations:
  - **Market Volatility**: Exchange rates can fluctuate, potentially leading to instability in routing decisions if not managed carefully.
  - **Equity and Access**: Ensuring that less established hosts can still participate meaningfully in the network without being overshadowed by hosts with established high-value currencies.
- These challenges are part of the ongoing discussion to refine and finalize routing strategies that align with PromiseGrid's decentralized ethos.

## Open Questions

### 1. Complexity of Implementation
   - What are the challenges in implementing these routing strategies?
   - Are there simpler alternatives that align with PromiseGrid’s goals?

### 2. Potential Pitfalls
   - What unforeseen issues might arise from using personal currencies in routing?
   - How do we handle situations where hosts frequently change their routing behavior?

### 3. Identity evasion

One potential pitfall in using personal currencies and reputation-based routing is the possibility that a participant can evade their past record of bad behavior by creating a new identity. Solutions to this issue are likely related to robust onboarding processes, which may include measures such as:

1. **Identity Verification**: Implementing strong identity verification steps could help ensure that new identities are tied to real-world anchors, making it harder to create disposable identities.
  
2. **Reputation Migration**: Allowing participants to carry some of their reputation (positive or negative) when creating a new identity can help mitigate the effects of identity hopping.

3. **Initial Restrictions**: Imposing temporary restrictions or limitations on new identities until they have built up a certain level of trust within the network.

4. **Behavioral Analysis**: Using automated systems to analyze behavior patterns that suggest identity evasion, thereby identifying and mitigating malicious actors faster.

5. **Exchange rates and Sovereignty**: Ultimately, the value of a host's currency is determined by the network's collective perception of its reliability and trustworthiness. Hosts that consistently fulfill promises and maintain good behavior will naturally experience higher currency values as reflected in the exchange rates quoted by their peers, increasing their currency value and influence within the network.

