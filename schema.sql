

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;


CREATE TYPE public.auth_status AS ENUM (
    'NOT_ACTIVATED',
    'SENT_INVITE',
    'ACTIVATED',
    'BLOCKED'
);



CREATE TYPE public.code_type AS ENUM (
    'ACTIVATE_INVITE',
    'ACTIVATE_REGISTRATION'
);



CREATE TYPE public.refresh_token_status AS ENUM (
    'ACTIVE',
    'REVOKED',
    'EXPIRED',
    'LOGOUT'
);


SET default_table_access_method = heap;


CREATE TABLE public.auth (
    user_id uuid NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    email text NOT NULL,
    password_hash bytea NOT NULL,
    status public.auth_status NOT NULL,
    CONSTRAINT auth_email_check CHECK ((length(email) <= 1024))
);



CREATE TABLE public.code (
    code text NOT NULL,
    user_id uuid NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    active boolean NOT NULL,
    type public.code_type NOT NULL
);



CREATE TABLE public.goose_db_version (
    id integer NOT NULL,
    version_id bigint NOT NULL,
    is_applied boolean NOT NULL,
    tstamp timestamp without time zone DEFAULT now()
);



CREATE SEQUENCE public.goose_db_version_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;



ALTER SEQUENCE public.goose_db_version_id_seq OWNED BY public.goose_db_version.id;



CREATE TABLE public.permission (
    id uuid NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    name text NOT NULL,
    description text NOT NULL,
    CONSTRAINT permission_description_check CHECK ((length(description) <= 2096)),
    CONSTRAINT permission_name_check CHECK ((length(name) <= 128))
);



CREATE TABLE public.refresh_token (
    id uuid NOT NULL,
    user_id uuid NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    status public.refresh_token_status NOT NULL
);



CREATE TABLE public.role (
    id uuid NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    name text NOT NULL,
    description text NOT NULL,
    CONSTRAINT role_description_check CHECK ((length(description) <= 2096)),
    CONSTRAINT role_name_check CHECK ((length(name) <= 128))
);



CREATE TABLE public.role_permission (
    permission_id uuid NOT NULL,
    role_id uuid NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL
);



CREATE TABLE public.user_role (
    user_id uuid NOT NULL,
    role_id uuid NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL
);



CREATE TABLE public.users (
    id uuid NOT NULL,
    firstname text NOT NULL,
    surname text NOT NULL,
    patronymic text,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    CONSTRAINT users_firstname_check CHECK ((length(firstname) <= 1024)),
    CONSTRAINT users_patronymic_check CHECK ((length(patronymic) <= 1024)),
    CONSTRAINT users_surname_check CHECK ((length(surname) <= 1024))
);



ALTER TABLE ONLY public.goose_db_version ALTER COLUMN id SET DEFAULT nextval('public.goose_db_version_id_seq'::regclass);



ALTER TABLE ONLY public.auth
    ADD CONSTRAINT auth_email_key UNIQUE (email);



ALTER TABLE ONLY public.auth
    ADD CONSTRAINT auth_pkey PRIMARY KEY (user_id);



ALTER TABLE ONLY public.code
    ADD CONSTRAINT code_pkey PRIMARY KEY (code);



ALTER TABLE ONLY public.goose_db_version
    ADD CONSTRAINT goose_db_version_pkey PRIMARY KEY (id);



ALTER TABLE ONLY public.permission
    ADD CONSTRAINT permission_name_key UNIQUE (name);



ALTER TABLE ONLY public.permission
    ADD CONSTRAINT permission_pkey PRIMARY KEY (id);



ALTER TABLE ONLY public.refresh_token
    ADD CONSTRAINT refresh_token_pkey PRIMARY KEY (id);



ALTER TABLE ONLY public.role
    ADD CONSTRAINT role_name_key UNIQUE (name);



ALTER TABLE ONLY public.role
    ADD CONSTRAINT role_pkey PRIMARY KEY (id);



ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);



CREATE UNIQUE INDEX idx_role_permission ON public.role_permission USING btree (permission_id, role_id);



CREATE UNIQUE INDEX idx_user_role ON public.user_role USING btree (user_id, role_id);



ALTER TABLE ONLY public.auth
    ADD CONSTRAINT auth_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);



ALTER TABLE ONLY public.code
    ADD CONSTRAINT code_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);



ALTER TABLE ONLY public.refresh_token
    ADD CONSTRAINT refresh_token_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);



ALTER TABLE ONLY public.role_permission
    ADD CONSTRAINT role_permission_permission_id_fkey FOREIGN KEY (permission_id) REFERENCES public.permission(id);



ALTER TABLE ONLY public.role_permission
    ADD CONSTRAINT role_permission_role_id_fkey FOREIGN KEY (role_id) REFERENCES public.role(id);



ALTER TABLE ONLY public.user_role
    ADD CONSTRAINT user_role_role_id_fkey FOREIGN KEY (role_id) REFERENCES public.role(id);



ALTER TABLE ONLY public.user_role
    ADD CONSTRAINT user_role_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);



