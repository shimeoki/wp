create table status (
    id integer primary key
    , name text not null unique
);

create table alias (
    id integer primary key
    , name       text not null
    , created_at text not null default current_timestamp
    , updated_at text not null default current_timestamp
);

create trigger alias_keep_ts
update of created_at on alias
begin
    select raise(abort, '''created_at'' shouldn''t be updated');
end;

create trigger alias_update_ts
after update of id, name on alias
begin
    update alias
    set updated_at = current_timestamp
    where id = new.id;
end;

create table tag (
    id integer primary key
    , name       text not null unique
    , created_at text not null default current_timestamp
    , updated_at text not null default current_timestamp
);

create table source (
    id integer primary key
    , name       text not null
    , link       text
    , created_at text not null default current_timestamp
    , updated_at text not null default current_timestamp
);

create table wallpaper (
    id integer primary key
    , hash       text not null unique
    , extension  text not null
    , created_at text not null default current_timestamp
    , updated_at text not null default current_timestamp
);

create table wallpaper_alias (
    wallpaper_id integer,
    alias_id     integer,

    primary key (wallpaper_id, alias_id)

    , foreign key (wallpaper_id) references wallpaper (id)
        on update cascade
        on delete cascade

    , foreign key (alias_id) references alias (id)
        on update cascade
        on delete cascade
);

create table wallpaper_tag (
    wallpaper_id integer,
    tag_id       integer,

    primary key (wallpaper_id, tag_id)

    , foreign key (wallpaper_id) references wallpaper (id)
        on update cascade
        on delete cascade

    , foreign key (tag_id) references tag (id)
        on update cascade
        on delete cascade
);

create table wallpaper_source (
    wallpaper_id integer,
    source_id    integer,

    primary key (wallpaper_id, source_id)

    , foreign key (wallpaper_id) references wallpaper (id)
        on update cascade
        on delete cascade

    , foreign key (source_id) references source (id)
        on update cascade
        on delete cascade
);

create table queue (
    id integer primary key
    , wallpaper_id integer not null
    , status_id    integer not null
    , priority     integer not null
    , created_at   text    not null default current_timestamp
    , updated_at   text    not null default current_timestamp

    , foreign key (wallpaper_id) references wallpaper (id)
        on update cascade
        on delete cascade

    , foreign key (status_id) references status (id)
        on update cascade
        on delete cascade
);
