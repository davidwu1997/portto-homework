create table blocks (
  block_num bigint not null unique,
  block_hash char(66) not null unique,
  block_time bigint not null,
  parent_hash char(66) not null unique,
  transactions text,
  primary key(block_num)
);

create index block_num_index on blocks(block_num);
create index block_hash_index on blocks(block_hash);

create table transactions (
  tx_hash char(66) not null unique,
  from_acc varchar(66),
  to_acc varchar(66),
  nonce bigint,
  data text,
  value bigint,
  logs text,
  primary key(tx_hash)
);

create index transaction_hash_index on transactions(tx_hash);
