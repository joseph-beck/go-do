create table public.tasks
(
    id          serial
        primary key,
    name        varchar(50) default 'Task'::character varying             not null,
    description text        default 'Description'::text                   not null,
    complete    boolean     default false                                 not null,
    deadline    varchar(20) default '00/00/0000-00:00'::character varying not null
);

alter table public.tasks
    owner to postgres;
