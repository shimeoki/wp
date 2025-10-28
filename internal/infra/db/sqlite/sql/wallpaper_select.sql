select
    w.id
    , w.uuid
    , w.hash
    , w.format
    , w.created_at
    , w.updated_at

    , t.id
    , t.uuid
    , t.name
    , t.created_at
    , t.updated_at

    , s.id
    , s.uuid
    , s.name
    , s.link
    , s.created_at
    , s.updated_at

from wallpaper as w

left join wallpaper_tag as wt on wt.wallpaper_id = w.id
left join tag as t on wt.tag_id = t.id

left join wallpaper_source as ws on ws.wallpaper_id = w.id
left join source as s on ws.source_id = s.id
