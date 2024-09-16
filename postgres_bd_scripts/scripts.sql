create database datingapp
    with owner local_user;


CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS postgis;

create table public."user"
(
    user_id     uuid                     default uuid_generate_v4() not null
        constraint users_pkey
            primary key,
    name        varchar(100)                                        not null,
    age         integer                                             not null,
    gender      varchar(10),
    location    geography(Point, 4326),
    interests   text[],
    last_active timestamp with time zone default now()
);

alter table public."user"
    owner to local_user;

create index idx_user_gender
    on public."user" (gender);

create index idx_user_age
    on public."user" (age);

create index idx_user_location
    on public."user" using gist (location);



INSERT INTO public."user" (user_id, name, age, gender, location, interests, last_active) VALUES ('e12de0d8-7fd3-4b4b-a246-bdfd1904f839', 'John Doe', 30, 'male', '0101000020E61000001FF64201DB7B52C0B610E4A0845D4440', '{hiking,reading}', '2024-09-15 18:06:09.222903 +00:00');
INSERT INTO public."user" (user_id, name, age, gender, location, interests, last_active) VALUES ('d4410834-9963-4fb3-bdff-f29d667f7d15', 'Liza Kudrow', 31, 'female', '0101000020E61000001FF64201DB7B52C0B610E4A0845D4440', '{hiking,gaming}', '2024-09-15 18:06:09.222903 +00:00');


create table public.preference
(
    id            serial
        primary key,
    user_id       uuid not null
        references public."user",
    gender        varchar(10),
    age_range_min integer,
    age_range_max integer,
    max_distance  double precision
);

alter table public.preference
    owner to local_user;



INSERT INTO public.preference (id, user_id, gender, age_range_min, age_range_max, max_distance) VALUES (1, 'e12de0d8-7fd3-4b4b-a246-bdfd1904f839', 'female', 25, 35, 50);
INSERT INTO public.preference (id, user_id, gender, age_range_min, age_range_max, max_distance) VALUES (2, 'd4410834-9963-4fb3-bdff-f29d667f7d15', 'male', 25, 35, 50);