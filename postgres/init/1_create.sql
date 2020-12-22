create database weightrec;

\c weightrec

create schema users;
create table users.users (
    id uuid primary key default gen_random_uuid(),
    name  varchar(100)
);

create schema logs;
create table logs.logs (
    logid serial not null,
    userid uuid not null references users.users (id),
    weight decimal,
    bfp  decimal,
    recorded_time timestamp not null
);
