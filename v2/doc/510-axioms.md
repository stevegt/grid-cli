# Axioms 

## Definition of "Axiom"

An axiom is a statement or proposition regarded as established, accepted, or self-evidently true. In PromiseGrid, axioms define the fundamental principles guiding the system's design, implementation, and operation.

### Axioms on which PromiseGrid is based

#### Side Effects

- External observers experience state changes when observing a function call's results.
- A function call always has side effects from the perspective of external observers.
- The system cannot fully know all side effects due to the existence of external observers.

#### Computability

- There are approximately 10^80 atoms in the observable universe.  (We can round this up to 2^270.)
- Any data in the observable universe that is stored at atomic scale is addressable by a byte sequence on the order of 2^270 bits.
- Any computable data can be addressed using a byte sequence that contains or references the algorithm and its input data (otherwise known as referential transparency).
- A reference to an algorithm or its input data can consist of a cryptographic hash of the algorithm or data.
- Cryptographic hashes are reasonably collision-resistant at any given moment in time and tend to be replaced by more secure algorithms as they are broken.
- A system that expects cryptographic hashes to be collision-resistant but imperfect, and that supports replaceable cryptographic hashes, can reasonably be expected to be able to reference any data or computable function, limited by the bit-length of the hashes it supports.
- A system that supports replaceable cryptographic hashes of 2^270 bits or more can reasonably be expected to be able to reference any atomic-scale data or computable function in the observable universe, now and in the foreseeable future, limited by the halting problem.

#### Byte Sequence Completion as a Universal Computation Model

- Byte sequences can be used to represent any computable algorithm or data.
- Completion of a byte sequence is equivalent to executing a function given its input data (referential transparency).
- Bytes in a sequence can be organized such that the leading bytes signal the type, class, protocol, algorithm, machine, or other routing or runtime prerequisites of the computation needed to complete the sequence.
- A byte sequence can be completed by a system that can interpret the leading bytes of the sequence and route the sequence to a subsystem that can complete the sequence.
- If a subsystem is not currently available to complete a byte sequence, the sequence itself can be used to represent a promise to complete the sequence in the future.
- If a subsystem is currently available to complete a byte sequence, the promise of completion can be fulfilled in the near term, limited by the halting problem.
- A system that is able to complete byte sequences can be used to represent any computable algorithm or data, limited by the halting problem and available subsystems.

