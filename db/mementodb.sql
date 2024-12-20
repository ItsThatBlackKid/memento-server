CREATE TABLE if not exists  [users] (
    id integer primary key autoincrement,
    username text not null,
    email text not null,
    first_name text,
    last_name text,
    password text not null
);

CREATE TABLE if not exists [memento] (
    id integer primary key not null,
    userid integer not null,
    title text not null,
    body text,
    mood integer not null,
    foreign key (userid) references user(id) on update cascade
)

