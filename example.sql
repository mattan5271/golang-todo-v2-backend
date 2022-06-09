drop table todos;

create table todos (
  id          serial primary key,
  title       text,
  description text,
  created_at  timestamp not null,
  updated_at  timestamp not null
);