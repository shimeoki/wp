create table status (
    id     int  primary key
    , name text not null unique
);

create table alias (
    id           int  primary key
    , name       text not null unique
    , created_at text not null default current_timestamp
    , updated_at text not null default current_timestamp
);
