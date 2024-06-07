--
-- PostgreSQL database dump
--

-- Dumped from database version 16.0
-- Dumped by pg_dump version 16.0

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: apps; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.apps (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.apps OWNER TO postgres;

--
-- Name: apps_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.apps_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.apps_id_seq OWNER TO postgres;

--
-- Name: apps_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.apps_id_seq OWNED BY public.apps.id;


--
-- Name: infrastructures; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.infrastructures (
    id character varying(36) NOT NULL,
    product_id character varying(36) NOT NULL,
    name character varying(255) NOT NULL,
    deployment_model character varying(255) NOT NULL,
    user_count integer NOT NULL,
    user_limit integer NOT NULL,
    metadata json NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.infrastructures OWNER TO postgres;

--
-- Name: organizations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.organizations (
    id character varying(36) NOT NULL,
    name character varying(255) NOT NULL,
    subdomain character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.organizations OWNER TO postgres;

--
-- Name: products; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.products (
    id character varying(36) NOT NULL,
    app_id integer NOT NULL,
    tier_id integer NOT NULL,
    deployment_schema json NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.products OWNER TO postgres;

--
-- Name: tenants; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tenants (
    id character varying(36) NOT NULL,
    product_id character varying(36) NOT NULL,
    organization_id character varying(36) NOT NULL,
    name character varying(45) NOT NULL,
    status character varying(50) NOT NULL,
    resource_information json,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.tenants OWNER TO postgres;

--
-- Name: tenants_infrastructures; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tenants_infrastructures (
    tenant_id character varying(36) NOT NULL,
    infrastructure_id character varying(36) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.tenants_infrastructures OWNER TO postgres;

--
-- Name: tiers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tiers (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    price integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.tiers OWNER TO postgres;

--
-- Name: tiers_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.tiers_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.tiers_id_seq OWNER TO postgres;

--
-- Name: tiers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.tiers_id_seq OWNED BY public.tiers.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id character varying(36) NOT NULL,
    name character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    username character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_organizations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users_organizations (
    user_id character varying(36),
    organization_id character varying(36),
    role character varying(45) NOT NULL
);


ALTER TABLE public.users_organizations OWNER TO postgres;

--
-- Name: users_tenants; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users_tenants (
    user_id character varying(36) NOT NULL,
    tenant_id character varying(36) NOT NULL,
    role character varying(45) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.users_tenants OWNER TO postgres;

--
-- Name: apps id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.apps ALTER COLUMN id SET DEFAULT nextval('public.apps_id_seq'::regclass);


--
-- Name: tiers id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tiers ALTER COLUMN id SET DEFAULT nextval('public.tiers_id_seq'::regclass);


--
-- Data for Name: apps; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.apps (id, name, created_at, updated_at) FROM stdin;
1	SaaS Todos	2024-06-03 22:29:10.849556	2024-06-03 22:29:10.849556
2	SaaS Notes	2024-06-03 22:29:10.849556	2024-06-03 22:29:10.849556
3	SaaS Image App	2024-06-03 22:29:10.849556	2024-06-03 22:29:10.849556
\.


--
-- Data for Name: infrastructures; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.infrastructures (id, product_id, name, deployment_model, user_count, user_limit, metadata, created_at, updated_at) FROM stdin;
c0ad301b-aa18-484c-a8e1-0f4ce0af0513	908c22d0-0b7a-4820-a500-27440ab73c63	storage	pool	1	100	{"variables":[{"storage_host":"34.101.175.147"},{"storage_name":"cloud-sql-841e9718-e4c4-4ca0-9b62-e371bd40cb42"},{"storage_password":"841e9718-e4c4-4ca0-9b62-e371bd40cb42"},{"storage_port":"5432"},{"storage_user":"default-841e9718-e4c4-4ca0-9b62-e371bd40cb42"}]}	2024-06-05 05:44:39.384342	2024-06-05 05:44:39.384342
158cf49b-1f09-4a7e-a0dc-3c478088509c	908c22d0-0b7a-4820-a500-27440ab73c63	compute	pool	1	100	{"variables":[{"compute_url":"https://compute-841e9718-e4c4-4ca0-9b62-e371bd40cb42-li5jjtbjrq-et.a.run.app"}]}	2024-06-05 05:44:39.391745	2024-06-05 05:44:39.391745
\.


--
-- Data for Name: organizations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.organizations (id, name, subdomain, created_at, updated_at) FROM stdin;
1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI	dptsi	2024-06-03 15:30:30.574858	2024-06-03 15:30:30.574858
\.


--
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.products (id, app_id, tier_id, deployment_schema, created_at, updated_at) FROM stdin;
63900f28-1d1d-4df1-a448-abc1c2756326	1	3	{\n  "terraform_repository_url": "https://github.com/godlixe/saas-todo.git",\n  "terraform_entrypoint_dir": "saas-todo/terraform/tiers/1",\n  "infrastructure_blueprint": [\n  \t{"compute":"pool"},\n\t{"storage":"pool"}\n  ]\n}	2024-06-03 22:29:10.849556	2024-06-03 22:29:10.849556
81e54445-80f6-43c4-8c64-378cae02a74f	2	4	{\n  "terraform_repository_url": "https://github.com/godlixe/saas-todo.git",\n  "terraform_entrypoint_dir": "saas-todo/terraform/tiers/1",\n  "infrastructure_blueprint": [\n  \t{"compute":"pool"},\n\t{"storage":"pool"}\n  ]\n}	2024-06-03 22:29:10.849556	2024-06-03 22:29:10.849556
4900214b-bd86-48b4-b3a4-94dc67007372	2	5	{\n  "terraform_repository_url": "https://github.com/godlixe/saas-todo.git",\n  "terraform_entrypoint_dir": "saas-todo/terraform/tiers/1",\n  "infrastructure_blueprint": [\n  \t{"compute":"pool"},\n\t{"storage":"pool"}\n  ]\n}	2024-06-03 22:29:10.849556	2024-06-03 22:29:10.849556
780b3ee2-60f1-430a-9c09-fdba1816a3ab	3	6	{\n  "terraform_repository_url": "https://github.com/godlixe/saas-todo.git",\n  "terraform_entrypoint_dir": "saas-todo/terraform/tiers/1",\n  "infrastructure_blueprint": [\n  \t{"compute":"pool"},\n\t{"storage":"pool"}\n  ]\n}	2024-06-03 22:29:10.849556	2024-06-03 22:29:10.849556
164c8302-2146-4f1c-82c3-b93f1d7e41dd	3	7	{\n  "terraform_repository_url": "https://github.com/godlixe/saas-todo.git",\n  "terraform_entrypoint_dir": "saas-todo/terraform/tiers/1",\n  "infrastructure_blueprint": [\n  \t{"compute":"pool"},\n\t{"storage":"pool"}\n  ]\n}	2024-06-03 22:29:10.849556	2024-06-03 22:29:10.849556
2cac359e-d35a-458c-8474-3b7b9cc3984c	1	1	{\n  "terraform_repository_url": "https://github.com/godlixe/saas-todo.git",\n  "terraform_entrypoint_dir": "saas-todo/terraform/tiers/1",\n  "infrastructure_blueprint": [\n  \t{"compute":"pool"},\n\t{"storage":"pool"}\n  ]\n}	2024-06-03 22:29:10.849556	2024-06-03 22:29:10.849556
908c22d0-0b7a-4820-a500-27440ab73c63	1	2	{\n  "terraform_repository_url": "https://github.com/godlixe/saas-todo.git",\n  "terraform_entrypoint_dir": "saas-todo/terraform/tiers/1",\n  "script_entrypoint":"saas-todo/terraform/tiers/1/my_script.go",\n  "infrastructure_blueprint": [\n    {\n      "compute": "pool"\n    },\n    {\n      "storage": "pool"\n    }\n  ]\n}	2024-06-03 22:29:10.849556	2024-06-03 22:29:10.849556
\.


--
-- Data for Name: tenants; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.tenants (id, product_id, organization_id, name, status, resource_information, created_at, updated_at) FROM stdin;
f755e497-259c-4a5d-8c37-3d4a79f5c3df	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-04 10:10:17.394935	2024-06-04 10:10:17.394935
d4cfabf9-9228-41b9-bce3-0fac0390849d	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-04 10:15:25.944258	2024-06-04 10:15:25.944258
41ab2515-3a6d-426e-b1aa-a4335bdb969f	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-04 10:22:47.384059	2024-06-04 10:22:47.384059
a466c8a8-2222-405c-b914-691a8dd8a3bf	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-04 10:29:30.608034	2024-06-04 10:29:30.608034
f904a8e2-f7ef-4e47-a14b-fcc415c78512	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-04 12:26:53.678459	2024-06-04 12:26:53.678459
7d47bd90-a0db-427b-851a-dab6da1a04ef	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-04 12:28:21.067931	2024-06-04 12:28:21.067931
5d62f008-c4ab-4331-a313-21bc1801af56	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-04 12:29:48.34029	2024-06-04 12:29:48.34029
05f55e03-2744-43a2-bfb2-345b915ce8b4	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-04 12:33:25.315574	2024-06-04 12:33:25.315574
7135e6f2-686d-40d7-aab8-2faee3ff7d36	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-04 12:33:48.085429	2024-06-04 12:33:48.085429
1c54d8a5-524c-4aea-8da8-f2d693a28f70	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-04 14:55:30.349296	2024-06-04 14:55:30.349296
79450614-4ca1-489f-90f9-8632cf32ad07	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-04 14:56:35.865566	2024-06-04 14:56:35.865566
47a40d94-5bcf-4a63-9f9b-e6a29eacc2c8	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-04 14:57:19.543432	2024-06-04 14:57:19.543432
6c5671b3-1c76-4106-a530-52fbdeb8ff67	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-04 14:57:40.451746	2024-06-04 14:57:40.451746
f2c0a8dc-d6a8-4587-8cf6-4a72872d373f	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-04 15:00:08.203093	2024-06-04 15:00:08.203093
67f3ee3a-b5ff-42a4-a0b2-6375b6d15ec7	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-04 15:02:53.434335	2024-06-04 15:02:53.434335
9ed45d5f-85c5-485d-ab0a-0c83152d46e6	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-04 15:03:59.839981	2024-06-04 15:03:59.839981
9201c1a3-698f-4163-a8b9-f929a8903553	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-04 15:04:46.911342	2024-06-04 15:04:46.911342
c770a7e6-0c6e-42f9-9578-25e053d22a3c	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-04 15:29:35.044847	2024-06-04 15:29:35.044847
9f16b9f8-ea1e-4a8a-9d04-b8a17d713777	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-04 15:31:19.96144	2024-06-04 15:31:19.96144
a84aa914-c814-4d92-a5ed-7c522fd73fbb	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-04 15:33:28.374051	2024-06-04 15:33:28.374051
bb986373-28a2-4418-bed0-4658c840f2b2	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-04 15:35:50.608162	2024-06-04 15:35:50.608162
a31fbf2d-27ce-47ba-ab9b-963dba97bdfc	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-04 15:36:09.492892	2024-06-04 15:36:09.492892
80a615a7-5abe-4c31-b437-5656707a5ea0	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-04 15:56:34.486345	2024-06-04 15:56:34.486345
d22f2a13-6890-4420-8d51-afabb9e68ba7	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-04 15:57:07.268411	2024-06-04 15:57:07.268411
84ea73f0-2fe4-498f-8edc-feeefffe1d7a	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-04 16:01:17.368517	2024-06-04 16:01:17.368517
f212f853-07f2-4a73-bf21-9a22033a6d32	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-04 16:10:29.93671	2024-06-04 16:10:29.93671
8143aad4-d947-40e9-b364-e9497b6096ee	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-04 16:20:48.975696	2024-06-04 16:20:48.975696
3ca5609e-3ed7-4234-ace7-cba7ec058e3a	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-04 16:31:54.650867	2024-06-04 16:31:54.650867
44491790-fbb4-49a4-8b12-355321583bb8	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-04 16:37:53.16944	2024-06-04 16:37:53.16944
6c4ab944-6e2a-41ea-95d3-a32c4c7391d6	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-04 16:39:38.122578	2024-06-04 16:39:38.122578
e49eca8b-9dbb-4cf9-a75c-572b04e35a65	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-04 16:55:20.485318	2024-06-04 16:55:20.485318
eeafd491-9ebc-49a9-9a70-6ad084aa1e5a	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-04 17:04:25.112425	2024-06-04 17:04:25.112425
bf82896e-5937-4119-9f91-527f1feed29e	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-04 17:05:46.353652	2024-06-04 17:05:46.353652
3a2da1cb-b181-489e-bb0d-6046dcff2d35	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-04 17:11:46.726146	2024-06-04 17:11:46.726146
f7d3bcf0-8c3b-4d38-99dd-ab308c3315e6	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-04 17:21:10.095269	2024-06-04 17:21:10.095269
affd014a-e4a1-424b-aa30-0c5bf5c6cbaa	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-04 17:30:50.5747	2024-06-04 17:30:50.5747
970195af-fe82-456e-8bf7-4e2e6e019a2e	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-05 04:11:05.186079	2024-06-05 04:11:05.186079
1da35da5-86b1-43e8-a453-3342a272d37b	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-05 04:13:02.989196	2024-06-05 04:13:02.989196
cc6fbb08-cbe6-4b89-bee6-9df465c7cbe6	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-05 04:18:18.822147	2024-06-05 04:18:18.822147
a43e3640-763a-497d-ba94-e790c3ebd180	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-05 04:22:47.410136	2024-06-05 04:22:47.410136
9f93b74d-c0c3-46c7-b6b4-c40898060a37	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-05 04:24:50.162559	2024-06-05 04:24:50.162559
4d0c857d-8980-450b-92df-288bcef90a4b	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-05 04:28:47.869131	2024-06-05 04:28:47.869131
856263a6-0c77-4f75-bd47-194966ef602c	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-05 04:56:00.411807	2024-06-05 04:56:00.411807
5ee76563-3f6c-406d-bc57-7ec8c19a637d	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-05 05:03:42.309787	2024-06-05 05:03:42.309787
91d03203-4fb2-4890-b876-0c209f865dc1	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-05 05:04:23.943153	2024-06-05 05:04:23.943153
921d7a13-deb9-41e4-b370-5b004761b126	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-05 05:21:06.58545	2024-06-05 05:21:06.58545
841e9718-e4c4-4ca0-9b62-e371bd40cb42	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-05 05:41:10.184562	2024-06-05 05:41:10.184562
edf8b5f2-0df6-4793-8e8b-d1cf41a93ab0	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-05 05:52:35.338336	2024-06-05 05:52:35.338336
0e215c4e-8939-47a1-8a92-d314da6f9d5e	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-05 05:55:05.629075	2024-06-05 05:55:05.629075
f3101fa3-7f89-401a-9208-8cc4e7a73edb	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-05 06:46:32.2349	2024-06-05 06:46:32.2349
736d47bc-69af-4ffe-912f-027f8b988c43	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-05 07:23:05.088141	2024-06-05 07:23:05.088141
b52628ee-6509-49bc-841e-56ce3d280a37	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-05 07:25:39.339598	2024-06-05 07:25:39.339598
d1a0e887-7ca9-4d23-9aa8-dd5a8c2e7588	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-05 07:25:42.219161	2024-06-05 07:25:42.219161
33cc4fe0-8239-4de4-affd-eeaf987aa6c7	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-05 07:29:58.512259	2024-06-05 07:29:58.512259
09c9a87f-e70b-4ad1-879b-0f3a2786813d	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-05 07:32:26.687082	2024-06-05 07:32:26.687082
f74eabab-a576-4845-814b-3d98ab9989c5	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-05 07:41:16.724736	2024-06-05 07:41:16.724736
768fd2d7-7db7-4a6d-9268-70cc9b212e06	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-05 07:50:01.415497	2024-06-05 07:50:01.415497
4969b633-b1e5-4f0c-8fef-a8dc04c54634	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-05 08:18:18.412208	2024-06-05 08:18:18.412208
9761d7c6-c64b-4243-92bc-92a0d239bcaa	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-05 08:19:21.027135	2024-06-05 08:19:21.027135
cd0a5c56-cd92-4602-9fea-0244f6adf702	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-05 08:44:35.921121	2024-06-05 08:44:35.921121
7dbf19a2-9ea7-4940-9106-5ebaa58a4f1e	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-05 08:51:11.086928	2024-06-05 08:51:11.086928
b60d9516-2f23-4c49-8e25-617822505066	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-05 09:02:57.791988	2024-06-05 09:02:57.791988
a603567f-5018-4a6e-843d-7e2d3b2d7e8c	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-05 09:03:39.967676	2024-06-05 09:03:39.967676
62b96500-1eed-4a90-b879-d67e67943733	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-05 09:06:18.018245	2024-06-05 09:06:18.018245
caefcbbf-dc5c-4e7a-9728-c4ad529318a2	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-05 09:07:58.713189	2024-06-05 09:07:58.713189
262c43c5-3ab4-46e5-8534-8547abb457f6	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-05 09:09:27.841711	2024-06-05 09:09:27.841711
11868297-edb5-4190-bb31-6686adb5236e	2cac359e-d35a-458c-8474-3b7b9cc3984c	1d324da0-c36b-4fb2-bea1-19dee034e683	DPTSI Image App	onboarding	\N	2024-06-05 12:52:41.905836	2024-06-05 12:52:41.905836
\.


--
-- Data for Name: tenants_infrastructures; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.tenants_infrastructures (tenant_id, infrastructure_id, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: tiers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.tiers (id, name, price, created_at, updated_at) FROM stdin;
1	SaaS Todos - Basic	10000	2024-06-03 22:29:10.849556	2024-06-03 22:29:10.849556
2	SaaS Todos - Premium	20000	2024-06-03 22:29:10.849556	2024-06-03 22:29:10.849556
3	SaaS Todos - Platinum	30000	2024-06-03 22:29:10.849556	2024-06-03 22:29:10.849556
4	SaaS Notes - Peasant	10000	2024-06-03 22:29:10.849556	2024-06-03 22:29:10.849556
5	SaaS Notes - Noble	100000	2024-06-03 22:29:10.849556	2024-06-03 22:29:10.849556
6	SaaS Image App - Basic	10000	2024-06-03 22:29:10.849556	2024-06-03 22:29:10.849556
7	SaaS Image App - Premium	20000	2024-06-03 22:29:10.849556	2024-06-03 22:29:10.849556
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, name, email, username, password, created_at, updated_at) FROM stdin;
02aedc3a-7e4a-47d0-97cd-a2b27c9fef2e	user	alexander19id@gmail.com	user1	$2a$04$ma7vhamOWzmsu11juQM5xeC//VFDtFGOHZOawMrk.OQs74j5TYJga	2024-06-03 15:29:19.925255	2024-06-03 15:29:19.925255
\.


--
-- Data for Name: users_organizations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users_organizations (user_id, organization_id, role) FROM stdin;
02aedc3a-7e4a-47d0-97cd-a2b27c9fef2e	1d324da0-c36b-4fb2-bea1-19dee034e683	admin
\.


--
-- Data for Name: users_tenants; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users_tenants (user_id, tenant_id, role, created_at, updated_at) FROM stdin;
\.


--
-- Name: apps_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.apps_id_seq', 1, false);


--
-- Name: tiers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.tiers_id_seq', 1, false);


--
-- Name: apps apps_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.apps
    ADD CONSTRAINT apps_pkey PRIMARY KEY (id);


--
-- Name: infrastructures infrastructures_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.infrastructures
    ADD CONSTRAINT infrastructures_pkey PRIMARY KEY (id);


--
-- Name: organizations organizations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.organizations
    ADD CONSTRAINT organizations_pkey PRIMARY KEY (id);


--
-- Name: products products_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);


--
-- Name: tenants tenants_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tenants
    ADD CONSTRAINT tenants_pkey PRIMARY KEY (id);


--
-- Name: tiers tiers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tiers
    ADD CONSTRAINT tiers_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: users_tenants users_tenants_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users_tenants
    ADD CONSTRAINT users_tenants_pkey PRIMARY KEY (user_id);


--
-- Name: infrastructures fk_infrastructures_product_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.infrastructures
    ADD CONSTRAINT fk_infrastructures_product_id FOREIGN KEY (product_id) REFERENCES public.products(id);


--
-- Name: products fk_products_app_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT fk_products_app_id FOREIGN KEY (app_id) REFERENCES public.apps(id);


--
-- Name: products fk_products_tier_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT fk_products_tier_id FOREIGN KEY (tier_id) REFERENCES public.tiers(id);


--
-- Name: tenants_infrastructures fk_tenants_infrastructures_infrastructure_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tenants_infrastructures
    ADD CONSTRAINT fk_tenants_infrastructures_infrastructure_id FOREIGN KEY (infrastructure_id) REFERENCES public.infrastructures(id);


--
-- Name: tenants_infrastructures fk_tenants_infrastructures_tenant_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tenants_infrastructures
    ADD CONSTRAINT fk_tenants_infrastructures_tenant_id FOREIGN KEY (tenant_id) REFERENCES public.tenants(id);


--
-- Name: tenants fk_tenants_organization_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tenants
    ADD CONSTRAINT fk_tenants_organization_id FOREIGN KEY (organization_id) REFERENCES public.organizations(id);


--
-- Name: tenants fk_tenants_product_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tenants
    ADD CONSTRAINT fk_tenants_product_id FOREIGN KEY (product_id) REFERENCES public.products(id);


--
-- Name: users_tenants fk_users_tenants_tenant_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users_tenants
    ADD CONSTRAINT fk_users_tenants_tenant_id FOREIGN KEY (tenant_id) REFERENCES public.tenants(id);


--
-- Name: users_tenants fk_users_tenants_user_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users_tenants
    ADD CONSTRAINT fk_users_tenants_user_id FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: users_organizations users_organizations_organization_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users_organizations
    ADD CONSTRAINT users_organizations_organization_id_fkey FOREIGN KEY (organization_id) REFERENCES public.organizations(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: users_organizations users_organizations_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users_organizations
    ADD CONSTRAINT users_organizations_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

