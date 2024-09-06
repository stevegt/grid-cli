# Axioms 

## Definition of "Axiom"

An axiom is a statement or proposition that is regarded as being established, accepted, or self-evidently true. In the context of PromiseGrid, axioms define the fundamental principles and assumptions upon which the system is built. These foundational truths guide the design, implementation, and operation of the PromiseGrid framework.

### Axioms on which PromiseGrid is based

#### Side effects

- An observer external to any known system will experience an alteration to its own state while observing the results of a function call.
- Every function call observed externally has a side effect.
- A system cannot be fully aware of all side effects, because the system cannot know of the existence or behavior of external observers.
- Side effects are never fully known.

#### Computability

- There are approximately 10^80 atoms in the observable universe.  (We can round this up to 2^270.)
- Any data in the observable universe that is stored at atomic scale is addressable by a byte sequence on the order of 2^270 bits.
- Any computable data can be addressed using a byte sequence that contains or references the algorithm and its input data (otherwise known as referential transparency).
- A reference to an algorithm or its input data can consist of a cryptographic hash of the algorithm or data.
- Cryptographic hashes are reasonably collision-resistant at any given moment in time and tend to be replaced by more secure algorithms as they are broken.
- A system that expects cryptographic hashes to be collision-resistant but imperfect, and that supports replaceable cryptographic hashes, can reasonably be expected to be able to reference any data or computable function, limited by the bit-length of the hashes it supports.
- A system that supports replaceable cryptographic hashes of 2^270 bits or more can reasonably be expected to be able to reference any atomic-scale data or computable function in the observable universe, now and in the foreseeable future, limited by the halting problem.

#### Byte sequence completion as a universal computation model

- Byte sequences can be used to represent any computable algorithm or data.
- Completion of a byte sequence is equivalent to executing a function given its input data (referential transparency).
- Bytes in a sequence can be organized such that the leading bytes signal the type, class, protocol, algorithm, machine, or other routing or runtime prerequisites of the computation needed to complete the sequence.
- A byte sequence can be completed by a system that can interpret the leading bytes of the sequence and route the sequence to a subsystem that can complete the sequence.
- If a subsystem is not currently available to complete a byte sequence, the sequence itself can be used to represent a promise to complete the sequence in the future.
- If a subsystem is currently available to complete a byte sequence, the promise of completion can be fulfilled in the near term, limited by the halting problem.
- A system that is able to complete byte sequences can be used to represent any computable algorithm or data, limited by the halting problem and available subsystems.

#### Considerations

The byte-sequence completion model can introduce complexity in implementation, especially on conventional computing hardware. Sequence matching, dynamic routing, and caching mechanisms are resource-intensive. Performance overheads can include latency and scalability issues, particularly with highly variable input data, high throughput requirements, and long sequences. Efficient resource management, algorithm selection, concurrent processing, and caching strategies are crucial for optimizing performance.

#### Open Questions

- How can the PromiseGrid framework be optimized to handle performance overheads introduced by the complexity of the byte-sequence completion model?
- What are the most effective strategies for managing memory and computational costs in the context of sequence matching and dynamic routing?
- How can adaptive mechanisms be implemented to adjust sequence matching strategies based on observed data patterns?
- What concurrency models and caching strategies are best suited to mitigate the latency and scalability issues inherent in the byte-sequence completion model?
