CREATE TABLE customer (
  id uuid PRIMARY KEY,
  name varchar(255) NOT NULL,
  email varchar(255) UNIQUE NOT NULL,
  cpf varchar(11) NOT NULL
);
