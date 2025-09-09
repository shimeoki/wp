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
    id           int  primary key
    , hash       text not null unique
    , extension  text not null
    , created_at text not null default current_timestamp
    , updated_at text not null default current_timestamp
);

create table wallpaper_alias (
    wallpaper_id int,
    alias_id     int,

    primary key (wallpaper_id, alias_id)

    , foreign key (wallpaper_id) references wallpaper (id)
        on update cascade
        on delete cascade

    , foreign key (alias_id) references alias (id)
        on update cascade
        on delete cascade
);

create table wallpaper_tag (
    wallpaper_id int,
    tag_id       int,

    primary key (wallpaper_id, tag_id)

    , foreign key (wallpaper_id) references wallpaper (id)
        on update cascade
        on delete cascade

    , foreign key (tag_id) references tag (id)
        on update cascade
        on delete cascade
);

create table wallpaper_source (
    wallpaper_id int,
    source_id    int,

    primary key (wallpaper_id, source_id)

    , foreign key (wallpaper_id) references wallpaper (id)
        on update cascade
        on delete cascade

    , foreign key (source_id) references source (id)
        on update cascade
        on delete cascade
);
