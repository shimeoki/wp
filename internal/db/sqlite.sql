pragma foreign_keys = on;

create table if not exists status (
    id integer primary key
    , name text not null unique
);


create table if not exists alias (
    id integer primary key
    , wallpaper_id integer   not null
    , name         text      not null
    , created_at   timestamp not null default current_timestamp
    , updated_at   timestamp not null default current_timestamp

    , unique (wallpaper_id, name)

    , foreign key (wallpaper_id) references wallpaper (id)
        on update cascade
        on delete cascade
);

create trigger if not exists alias_keep_ts
update of created_at on alias
begin
    select raise(abort, '''created_at'' shouldn''t be updated');
end;

create trigger if not exists alias_update_ts
after update of id, name on alias
begin
    update alias
    set updated_at = current_timestamp
    where id = new.id;
end;


create table if not exists tag (
    id integer primary key
    , name       text      not null unique
    , created_at timestamp not null default current_timestamp
    , updated_at timestamp not null default current_timestamp
);

create trigger if not exists tag_keep_ts
update of created_at on tag
begin
    select raise(abort, '''created_at'' shouldn''t be updated');
end;

create trigger if not exists tag_update_ts
after update of id, name on tag
begin
    update tag
    set updated_at = current_timestamp
    where id = new.id;
end;


create table if not exists source (
    id integer primary key
    , name       text      not null
    , link       text
    , created_at timestamp not null default current_timestamp
    , updated_at timestamp not null default current_timestamp
);

create trigger if not exists source_keep_ts
update of created_at on source
begin
    select raise(abort, '''created_at'' shouldn''t be updated');
end;

create trigger if not exists source_update_ts
after update of id, name, link on source
begin
    update source
    set updated_at = current_timestamp
    where id = new.id;
end;


create table if not exists wallpaper (
    id integer primary key
    , hash       text      not null unique
    , format     text      not null
    , created_at timestamp not null default current_timestamp
);

create trigger if not exists wallpaper_keep_ts
update of created_at on wallpaper
begin
    select raise(abort, '''created_at'' shouldn''t be updated');
end;


create table if not exists wallpaper_tag (
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


create table if not exists wallpaper_source (
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


create table if not exists queue (
    id integer primary key
    , wallpaper_id integer   not null
    , status_id    integer   not null
    , priority     integer   not null
    , created_at   timestamp not null default current_timestamp
    , updated_at   timestamp not null default current_timestamp

    , foreign key (wallpaper_id) references wallpaper (id)
        on update cascade
        on delete cascade

    , foreign key (status_id) references status (id)
        on update cascade
        on delete cascade
);

create trigger if not exists queue_keep_ts
update of created_at on queue
begin
    select raise(abort, '''created_at'' shouldn''t be updated');
end;

create trigger if not exists queue_update_ts
after update of id, wallpaper_id, status_id, priority on queue
begin
    update queue
    set updated_at = current_timestamp
    where id = new.id;
end;
