create database users;

create table users (
    id uuid primary key default gen_random_uuid(),
    name  varchar(100)
);

create table logs (
    logid serial not null,
    userid uuid not null references users (id),
    weight decimal,
    bfp  decimal,
    recorded_time timestamp not null
);
