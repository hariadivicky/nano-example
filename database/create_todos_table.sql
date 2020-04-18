CREATE TABLE todos (
    id integer primary key,
    title text not null,
    is_done integer default 0
);