pragma foreign_keys = on;

-- keep-sorted start block=yes newline_separated=yes
create table if not exists alias (
    id integer primary key
    , uuid         text      not null
    , wallpaper_id integer   not null
    , name         text      not null
    , created_at   timestamp not null default current_timestamp
    , updated_at   timestamp not null default current_timestamp

    , unique (wallpaper_id, name)

    , foreign key (wallpaper_id) references wallpaper (id)
        on update cascade
        on delete cascade
);

create table if not exists queue (
    id integer primary key
    , uuid         text      not null
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

create table if not exists source (
    id integer primary key
    , uuid       text      not null
    , name       text      not null
    , link       text
    , created_at timestamp not null default current_timestamp
    , updated_at timestamp not null default current_timestamp
);

create table if not exists status (
    id integer primary key
    , name text not null unique
);

create table if not exists tag (
    id integer primary key
    , uuid       text      not null
    , name       text      not null unique
    , created_at timestamp not null default current_timestamp
    , updated_at timestamp not null default current_timestamp
);

create table if not exists wallpaper (
    id integer primary key
    , uuid       text      not null
    , hash       text      not null unique
    , format     text      not null
    , created_at timestamp not null default current_timestamp
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
-- keep-sorted end
