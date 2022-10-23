drop table if exists tasks;
drop table if exists sessions;
drop table if exists users;

create table users (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  username   varchar(255) not null unique,
  name       varchar(255),
  email      varchar(255) not null unique,
  password   varchar(255) not null,
  created_at timestamp not null   
);

create table tasks (
  id      serial primary key,
  title   text,
  user_id integer references users(id),
  created_at timestamp
);



create table sessions (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  email      varchar(255),
  user_id    integer references users(id),
  created_at timestamp not null   
);