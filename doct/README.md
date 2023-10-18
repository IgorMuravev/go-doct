doct is a package that modifies a document using transactions.

In packet terms, a "Document" is a set of bytes.
A transaction is an operation on this set of bytes.

4 types of transactions supported. Each transaction implements an interface ITransaction

- AppendTransaction is a transaction which append data to the end of a document. Contains field with DATA

- EraseTransaction is a transaction which blanks existing document

- InsertTransaction is a transaction which inserts some data into a specified position. Contains fields DATA + POSITION

- RemoveTransaction is a transaction which removes some data in specified range. Contains fields POSITION + COUNT