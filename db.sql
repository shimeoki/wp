create table status (
    id     integer primary key
    , name text not null unique
);

create table alias (
    id           integer primary key
    , name       text not null unique
    , created_at text not null default current_timestamp
    , updated_at text not null default current_timestamp
);
