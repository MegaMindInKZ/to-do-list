drop table if exists tasks;
drop table if exists sessions;
drop table if exists users;
drop table if exists tasks;

create table users (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  username   varchar(255) not null unique,
  name       varchar(255),
  email      varchar(255) not null unique,
  password   varchar(255) not null,
  avatar     varchar(255),
  created_at timestamp not null   
);

create table sessions (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  email      varchar(255),
  user_id    integer references users(id),
  created_at timestamp not null   
);

create table tasks(
    id          serial primary key,
    user_id     integer references users(id),
    title       text,
    deadline    date,
    description text,
    isImportant boolean default false,
    isFinished  boolean default false,
    created_at  timestamp
);

