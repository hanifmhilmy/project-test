# Database Schema

For this case I want to use the relational database, since the relation between each of Entity already defined, and relational DB known to be have a strong data integrity semantics through ACID (Atomicity, Consistency, Isolation, Durability), I think it suite better to use relational DB.

Even tho I know if using this relation DB we will facing challange to update the schema (since in most of the case will require downtime), doing horizontal scale and also prone to SPOF (single point of failure) that need to be considered. 

Relation:

- User 1...m Wallet
- User 1...m Team
- Team 1...m User

Schema:

