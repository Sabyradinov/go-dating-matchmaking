create database datingapp
    with owner 'USER';


CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS postgis;

CREATE TABLE user (
   user_id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
   name VARCHAR(100) NOT NULL,
   age INTEGER NOT NULL,
   gender VARCHAR(10),
   location GEOGRAPHY(POINT, 4326),
   interests TEXT[]
);

CREATE TABLE public.preference (
   id SERIAL PRIMARY KEY,
   user_id UUID NOT NULL,
   gender VARCHAR(10),
   age_range_min int,
   age_range_max int,
   max_distance FLOAT,
   FOREIGN KEY (user_id) REFERENCES public.user(user_id)
);

CREATE INDEX idx_user_gender ON public.user (gender);
CREATE INDEX idx_user_age ON public.user (age);
CREATE INDEX idx_user_location ON public.user USING GIST (location);
CREATE INDEX idx_user_last_active ON public.user (last_active);