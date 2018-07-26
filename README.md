Hypergold

Official Web

http://hglod.com/

Innovations To date, existing decentralized cryptocurrencies adopt either PoW consensus scheme or hybrid consensus model of PoW and PoS. However, these systems still encounter the issue of very limited efficiency/throughput. Meanwhile, upcoming quantum computers threaten existing classical cryptography which is the foundation of blockchain security. In particular, the quantum algorithm by Shor for computing discrete logarithms breaks the ECDSA signature scheme used by almost all cryptocurrencies, such as Bitcoin, Ethereum, Decred and Monero. However, if post-quantum cryptographic schemes are equipped in these systems, the throughput of them will become worse and even unbearable.

Hcash project aims to build a secure, efficient, robust and reliable decentralized system. Highlighted features such as newly-proposed hybrid consensus scheme, post-quantum digital signature, linkability among various blockchain-based and DAG-based decentralized cryptocurrencies, smart contract mechanism and post-quantum privacy-preserving scheme will be proposed and implemented in Hcash eventually.

Novel Consensus Scheme To deal with the performance issue, we implement a novel hybrid consensus scheme with strong robustness, high throughput as well as sufficient flexibility in Hcash. On the one hand, with a newly-proposed two-layer framework of block chain, significant improvement of the efficiency is offered without compromising the security. On the other hand, with a hybrid consensus model, both PoW and PoS miners are incentivized to take part in the consensus process, thereby enhancing the security and flexibility of the consensus scheme, and providing a mechanism that supports basic DAO for future protocol updating and project investments.

Post-Quantum Features To address security issues stemming from quantum computers, we design and implement post-quantum solutions in Hcash. Our proposals achieve the following 4 features:

Compatibility: Compatible with existing ECDSA signature solution; Flexibility: Support multiple post-quantum signature solutions that are thoroughly analyzed, assessed and proved by international cryptography research institutions, meanwhile their security and performance must be outstanding; Security: the post-quantum solution must be proved secure in theory, and side-channel attack proof in practice; High performance: Signing and signature verification must be fast. Most importantly, the public key and signature must be short.

Starting Hgold Hgold is a Hypergold full node implementation written in Go (golang).

This acts as a chain daemon for the Hypergold cryptocurrency. Hgold maintains the entire past transactional ledger of Hypergold and allows relaying of transactions to other Hypergold nodes across the world.

Hgold Start hgold solo mining

hgoldctl setgenerate true x # where x represents the number of CPU threads Stop hgold solo mining

hgoldctl setgenerate false
