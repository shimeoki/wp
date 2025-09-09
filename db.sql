create table status (
    id     int  primary key
    , name text not null unique
);

create table alias (
    id           int  primary key
    , name       text not null
    , created_at text not null default current_timestamp
    , updated_at text not null default current_timestamp
);

create table tag (
    id           int  primary key
    , name       text not null unique
    , created_at text not null default current_timestamp
    , updated_at text not null default current_timestamp
);

create table source (
    id           int  primary key
    , name       text not null
    , link       text
    , created_at text not null default current_timestamp
    , updated_at text not null default current_timestamp
);

create table wallpaper (
    id           int  primary key,
    , hash       text not null unique
    , extension  text not null
    , created_at text not null default current_timestamp
    , updated_at text not null default current_timestamp
);
