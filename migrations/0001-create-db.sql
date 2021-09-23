CREATE TABLE transaction(
  id UUID primary key,
  title varchar,
  value integer,
  created_at timestamp default current_timestamp,
  updated_at timestamp default current_timestamp
);
