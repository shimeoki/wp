select
    a.id
    , a.uuid
    , w.id
    , w.uuid
    , a.name
    , a.created_at
    , a.updated_at

from alias as a
left join wallpaper as w on a.wallpaper_id = w.id
