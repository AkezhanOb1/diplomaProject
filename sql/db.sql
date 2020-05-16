--
-- PostgreSQL database dump
--

-- Dumped from database version 11.5
-- Dumped by pg_dump version 11.5

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

--
-- Name: citext; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS citext WITH SCHEMA public;


--
-- Name: EXTENSION citext; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION citext IS 'data type for case-insensitive character strings';


--
-- Name: set_default_company_service_operation_hours(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.set_default_company_service_operation_hours() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    INSERT INTO business_company_service_operation_hours
        (business_company_id, business_service_id, day_of_week, open_time, close_time)
    SELECT business_company_id, NEW.ID, day_of_week, open_time, close_time
    FROM business_company_operation_hours where business_company_id = NEW.company_id;

	RETURN NEW;
END;
$$;


ALTER FUNCTION public.set_default_company_service_operation_hours() OWNER TO postgres;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: business_category; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.business_category (
    id integer NOT NULL,
    name character varying(50) NOT NULL
);


ALTER TABLE public.business_category OWNER TO postgres;

--
-- Name: business_categories_business_categories_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.business_categories_business_categories_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.business_categories_business_categories_id_seq OWNER TO postgres;

--
-- Name: business_categories_business_categories_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.business_categories_business_categories_id_seq OWNED BY public.business_category.id;


--
-- Name: business_company; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.business_company (
    id integer NOT NULL,
    name character varying NOT NULL,
    category_id integer
);


ALTER TABLE public.business_company OWNER TO postgres;

--
-- Name: business_company_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.business_company_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.business_company_id_seq OWNER TO postgres;

--
-- Name: business_company_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.business_company_id_seq OWNED BY public.business_company.id;


--
-- Name: business_company_operation_hours; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.business_company_operation_hours (
    id integer NOT NULL,
    business_company_id integer NOT NULL,
    day_of_week integer NOT NULL,
    open_time character varying NOT NULL,
    close_time character varying NOT NULL
);


ALTER TABLE public.business_company_operation_hours OWNER TO postgres;

--
-- Name: business_company_operation_hours_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.business_company_operation_hours_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.business_company_operation_hours_id_seq OWNER TO postgres;

--
-- Name: business_company_operation_hours_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.business_company_operation_hours_id_seq OWNED BY public.business_company_operation_hours.id;


--
-- Name: business_company_service; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.business_company_service (
    id integer NOT NULL,
    company_id integer NOT NULL,
    service_id integer NOT NULL,
    name text NOT NULL,
    duration integer DEFAULT 60,
    price numeric(12,2) NOT NULL
);


ALTER TABLE public.business_company_service OWNER TO postgres;

--
-- Name: business_company_service_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.business_company_service_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.business_company_service_id_seq OWNER TO postgres;

--
-- Name: business_company_service_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.business_company_service_id_seq OWNED BY public.business_company_service.id;


--
-- Name: business_company_service_operation_hours; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.business_company_service_operation_hours (
    id integer NOT NULL,
    business_company_id integer NOT NULL,
    business_service_id integer NOT NULL,
    day_of_week integer NOT NULL,
    open_time character varying NOT NULL,
    close_time character varying NOT NULL
);


ALTER TABLE public.business_company_service_operation_hours OWNER TO postgres;

--
-- Name: business_company_service_operation_hours_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.business_company_service_operation_hours_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.business_company_service_operation_hours_id_seq OWNER TO postgres;

--
-- Name: business_company_service_operation_hours_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.business_company_service_operation_hours_id_seq OWNED BY public.business_company_service_operation_hours.id;


--
-- Name: business_company_service_order; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.business_company_service_order (
    id integer NOT NULL,
    client_id integer NOT NULL,
    business_service_id integer NOT NULL,
    start_at character varying NOT NULL,
    pre_paid boolean NOT NULL,
    created_at character varying NOT NULL,
    client_first_name character varying NOT NULL,
    client_phone_number character varying(10) NOT NULL,
    client_phone_number_prefix character varying(5) NOT NULL,
    client_commentary character varying,
    end_at character varying NOT NULL
);


ALTER TABLE public.business_company_service_order OWNER TO postgres;

--
-- Name: business_company_service_order_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.business_company_service_order_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.business_company_service_order_id_seq OWNER TO postgres;

--
-- Name: business_company_service_order_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.business_company_service_order_id_seq OWNED BY public.business_company_service_order.id;


--
-- Name: business_owner; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.business_owner (
    id integer NOT NULL,
    name character varying NOT NULL,
    email public.citext NOT NULL,
    password character varying NOT NULL,
    phone_number_prefix character varying(5) NOT NULL,
    phone_number character varying(10) NOT NULL
);


ALTER TABLE public.business_owner OWNER TO postgres;

--
-- Name: business_owner_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.business_owner_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.business_owner_id_seq OWNER TO postgres;

--
-- Name: business_owner_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.business_owner_id_seq OWNED BY public.business_owner.id;


--
-- Name: business_service; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.business_service (
    id integer NOT NULL,
    name text NOT NULL
);


ALTER TABLE public.business_service OWNER TO postgres;

--
-- Name: business_services_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.business_services_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.business_services_id_seq OWNER TO postgres;

--
-- Name: business_services_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.business_services_id_seq OWNED BY public.business_service.id;


--
-- Name: business_sub_category; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.business_sub_category (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    category_id integer NOT NULL
);


ALTER TABLE public.business_sub_category OWNER TO postgres;

--
-- Name: business_sub_categories_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.business_sub_categories_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.business_sub_categories_id_seq OWNER TO postgres;

--
-- Name: business_sub_categories_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.business_sub_categories_id_seq OWNED BY public.business_sub_category.id;


--
-- Name: business_sub_categories_services; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.business_sub_categories_services (
    sub_categories_id integer NOT NULL,
    business_services_id integer NOT NULL
);


ALTER TABLE public.business_sub_categories_services OWNER TO postgres;

--
-- Name: client; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.client (
    id integer NOT NULL,
    first_name character varying NOT NULL,
    second_name character varying NOT NULL,
    email character varying NOT NULL,
    phone_number character varying(10) NOT NULL,
    phone_number_prefix character varying(5) NOT NULL,
    password character varying NOT NULL
);


ALTER TABLE public.client OWNER TO postgres;

--
-- Name: client_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.client_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.client_id_seq OWNER TO postgres;

--
-- Name: client_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.client_id_seq OWNED BY public.client.id;


--
-- Name: client_phone_number_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.client_phone_number_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.client_phone_number_seq OWNER TO postgres;

--
-- Name: client_phone_number_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.client_phone_number_seq OWNED BY public.client.phone_number;


--
-- Name: owners_businesses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.owners_businesses (
    business_owner_id integer NOT NULL,
    business_company_id integer NOT NULL
);


ALTER TABLE public.owners_businesses OWNER TO postgres;

--
-- Name: business_category id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.business_category ALTER COLUMN id SET DEFAULT nextval('public.business_categories_business_categories_id_seq'::regclass);


--
-- Name: business_company id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.business_company ALTER COLUMN id SET DEFAULT nextval('public.business_company_id_seq'::regclass);


--
-- Name: business_company_operation_hours id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.business_company_operation_hours ALTER COLUMN id SET DEFAULT nextval('public.business_company_operation_hours_id_seq'::regclass);


--
-- Name: business_company_service id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.business_company_service ALTER COLUMN id SET DEFAULT nextval('public.business_company_service_id_seq'::regclass);


--
-- Name: business_company_service_operation_hours id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.business_company_service_operation_hours ALTER COLUMN id SET DEFAULT nextval('public.business_company_service_operation_hours_id_seq'::regclass);


--
-- Name: business_company_service_order id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.business_company_service_order ALTER COLUMN id SET DEFAULT nextval('public.business_company_service_order_id_seq'::regclass);


--
-- Name: business_owner id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.business_owner ALTER COLUMN id SET DEFAULT nextval('public.business_owner_id_seq'::regclass);


--
-- Name: business_service id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.business_service ALTER COLUMN id SET DEFAULT nextval('public.business_services_id_seq'::regclass);


--
-- Name: business_sub_category id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.business_sub_category ALTER COLUMN id SET DEFAULT nextval('public.business_sub_categories_id_seq'::regclass);


--
-- Name: client id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.client ALTER COLUMN id SET DEFAULT nextval('public.client_id_seq'::regclass);


--
-- Data for Name: business_category; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.business_category (id, name) FROM stdin;
1	Развличения
3	Еда
4	красота
2	Спорт
\.


--
-- Data for Name: business_company; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.business_company (id, name, category_id) FROM stdin;
5	Aki corparation	3
6	Cactus	1
7	gqlgen	1
\.


--
-- Data for Name: business_company_operation_hours; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.business_company_operation_hours (id, business_company_id, day_of_week, open_time, close_time) FROM stdin;
2	6	2	08:00:00	20:00:00
3	6	3	08:00:00	20:00:00
4	6	4	08:00:00	20:00:00
5	6	5	08:00:00	20:00:00
6	6	6	08:00:00	20:00:00
7	6	0	08:00:00	20:00:00
9	7	1	10:00:00	17:00:00
\.


--
-- Data for Name: business_company_service; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.business_company_service (id, company_id, service_id, name, duration, price) FROM stdin;
1	5	1	Уход за ногтями	60	5500.00
2	7	1	Маникюр	60	5500.00
6	6	4	Мужская стрижка	60	5000.00
7	6	4	Женская стрижка	90	7000.00
9	6	1	Маник	60	4500.00
8	6	24	Массаж для лица	45	4500.00
10	6	4	Стрижка Усов для анчоусов	45	7000.00
\.


--
-- Data for Name: business_company_service_operation_hours; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.business_company_service_operation_hours (id, business_company_id, business_service_id, day_of_week, open_time, close_time) FROM stdin;
3	6	8	3	08:00:00	20:00:00
4	6	8	4	08:00:00	20:00:00
5	6	8	5	08:00:00	20:00:00
6	6	8	6	08:00:00	20:00:00
7	6	8	0	08:00:00	20:00:00
8	6	9	1	08:00:00	20:00:00
9	6	9	2	08:00:00	20:00:00
10	6	9	3	08:00:00	20:00:00
11	6	9	4	08:00:00	20:00:00
12	6	9	5	08:00:00	20:00:00
13	6	9	6	08:00:00	20:00:00
14	6	9	0	08:00:00	20:00:00
2	6	8	2	10:00:00	20:00:00
15	6	6	1	10:00:00	21:00:00
16	6	8	1	10:00:00	21:00:00
17	6	10	2	08:00:00	20:00:00
18	6	10	3	08:00:00	20:00:00
19	6	10	4	08:00:00	20:00:00
20	6	10	5	08:00:00	20:00:00
21	6	10	6	08:00:00	20:00:00
22	6	10	0	08:00:00	20:00:00
\.


--
-- Data for Name: business_company_service_order; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.business_company_service_order (id, client_id, business_service_id, start_at, pre_paid, created_at, client_first_name, client_phone_number, client_phone_number_prefix, client_commentary, end_at) FROM stdin;
32	1	10	2020-10-04T11:45:00Z	f	2020-05-10T23:59:07+06:00	Akezhan	7772059339	+7	SSS	2020-10-04T12:45:00Z
34	1	10	2020-10-04T14:45:00Z	f	2020-05-11T00:02:49+06:00	Akezhan	7772059339	+7	SSS	2020-10-04T15:45:00Z
\.


--
-- Data for Name: business_owner; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.business_owner (id, name, email, password, phone_number_prefix, phone_number) FROM stdin;
8	Akezhan	esbolatovakezhan@gmail.com	$2a$14$NZn2MQ3dXUipKec1CxBMoe3sSbprFttqA4/BPsLlY7DvarrpxhHoG	+7	7772059339
10	Akezhan	akez@gmail.com	$2a$14$zgvoS5SETkOM5z.sPfiB8uqFrXHV88MxHpvAOew5.zKjg2iJ108Ue	+7	7772059129
11	Aidos	doha@gmail.com	$2a$14$/EHMLWzSA1Vg44Kk6IWS3OAs170bmil7IQplKFNjGoE3PlCy1.yVa	+7	7422059129
\.


--
-- Data for Name: business_service; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.business_service (id, name) FROM stdin;
1	Маникюр
2	Педикюр
3	Завивка
4	Стрижка
8	Make Up
24	Массаж лица
\.


--
-- Data for Name: business_sub_categories_services; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.business_sub_categories_services (sub_categories_id, business_services_id) FROM stdin;
2	1
2	2
3	3
2	3
2	8
2	24
3	24
\.


--
-- Data for Name: business_sub_category; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.business_sub_category (id, name, category_id) FROM stdin;
2	Салоны красоты	4
3	Spa салоны	4
4	Бильярд клубы	1
\.


--
-- Data for Name: client; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.client (id, first_name, second_name, email, phone_number, phone_number_prefix, password) FROM stdin;
1	Cactus	Client	cactus@gmail.com	7772059339	+7	Yfehsp
\.


--
-- Data for Name: owners_businesses; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.owners_businesses (business_owner_id, business_company_id) FROM stdin;
8	5
10	7
11	7
\.


--
-- Name: business_categories_business_categories_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.business_categories_business_categories_id_seq', 4, true);


--
-- Name: business_company_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.business_company_id_seq', 7, true);


--
-- Name: business_company_operation_hours_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.business_company_operation_hours_id_seq', 9, true);


--
-- Name: business_company_service_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.business_company_service_id_seq', 10, true);


--
-- Name: business_company_service_operation_hours_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.business_company_service_operation_hours_id_seq', 22, true);


--
-- Name: business_company_service_order_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.business_company_service_order_id_seq', 34, true);


--
-- Name: business_owner_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.business_owner_id_seq', 11, true);


--
-- Name: business_services_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.business_services_id_seq', 24, true);


--
-- Name: business_sub_categories_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.business_sub_categories_id_seq', 4, true);


--
-- Name: client_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.client_id_seq', 1, true);


--
-- Name: client_phone_number_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.client_phone_number_seq', 1, false);


--
-- Name: business_category business_categories_category_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.business_category
    ADD CONSTRAINT business_categories_category_name_key UNIQUE (name);


--
-- Name: business_category business_categories_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.business_category
    ADD CONSTRAINT business_categories_pkey PRIMARY KEY (id);


--
-- Name: business_company business_company_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.business_company
    ADD CONSTRAINT business_company_name_key UNIQUE (name);


--
-- Name: business_company_operation_hours business_company_operation_hours_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.business_company_operation_hours
    ADD CONSTRAINT business_company_operation_hours_pk PRIMARY KEY (id);


--
-- Name: business_company business_company_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.business_company
    ADD CONSTRAINT business_company_pkey PRIMARY KEY (id);


--
-- Name: business_company_service_operation_hours business_company_service_operation_hours_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.business_company_service_operation_hours
    ADD CONSTRAINT business_company_service_operation_hours_pk PRIMARY KEY (id);


--
-- Name: business_company_service_order business_company_service_order_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.business_company_service_order
    ADD CONSTRAINT business_company_service_order_pk PRIMARY KEY (id);


--
-- Name: business_company_service business_company_service_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.business_company_service
    ADD CONSTRAINT business_company_service_pk PRIMARY KEY (id);


--
-- Name: business_owner business_owner_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.business_owner
    ADD CONSTRAINT business_owner_email_key UNIQUE (email);


--
-- Name: business_owner business_owner_phone_number_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.business_owner
    ADD CONSTRAINT business_owner_phone_number_key UNIQUE (phone_number);


--
-- Name: business_owner business_owner_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.business_owner
    ADD CONSTRAINT business_owner_pkey PRIMARY KEY (id);


--
-- Name: business_service business_services_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.business_service
    ADD CONSTRAINT business_services_pk PRIMARY KEY (id);


--
-- Name: business_sub_category business_sub_categories_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.business_sub_category
    ADD CONSTRAINT business_sub_categories_pk PRIMARY KEY (id);


--
-- Name: business_sub_categories_services business_sub_categories_services_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.business_sub_categories_services
    ADD CONSTRAINT business_sub_categories_services_pk PRIMARY KEY (sub_categories_id, business_services_id);


--
-- Name: client client_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.client
    ADD CONSTRAINT client_pk PRIMARY KEY (id);


--
-- Name: owners_businesses owners_businesses_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.owners_businesses
    ADD CONSTRAINT owners_businesses_pkey PRIMARY KEY (business_owner_id, business_company_id);


--
-- Name: business_company_service_name_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX business_company_service_name_uindex ON public.business_company_service USING btree (name);


--
-- Name: business_company_service_order_end_at_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX business_company_service_order_end_at_uindex ON public.business_company_service_order USING btree (end_at);


--
-- Name: business_company_service_order_start_at_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX business_company_service_order_start_at_uindex ON public.business_company_service_order USING btree (start_at);


--
-- Name: business_services_name_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX business_services_name_uindex ON public.business_service USING btree (name);


--
-- Name: business_sub_categories_name_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX business_sub_categories_name_uindex ON public.business_sub_category USING btree (name);


--
-- Name: client_email_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX client_email_uindex ON public.client USING btree (email);


--
-- Name: client_phone_number_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX client_phone_number_uindex ON public.client USING btree (phone_number);


--
-- Name: business_company_service business_company_service_added; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER business_company_service_added AFTER INSERT ON public.business_company_service FOR EACH ROW EXECUTE PROCEDURE public.set_default_company_service_operation_hours();


--
-- Name: business_company business_company_category_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.business_company
    ADD CONSTRAINT business_company_category_id_fkey FOREIGN KEY (category_id) REFERENCES public.business_category(id);


--
-- Name: business_company_service business_company_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.business_company_service
    ADD CONSTRAINT business_company_id_fk FOREIGN KEY (company_id) REFERENCES public.business_company(id);


--
-- Name: business_company_operation_hours business_company_operation_hours_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.business_company_operation_hours
    ADD CONSTRAINT business_company_operation_hours_fk FOREIGN KEY (business_company_id) REFERENCES public.business_company(id);


--
-- Name: business_company_service_operation_hours business_company_service_operation_hours_company_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.business_company_service_operation_hours
    ADD CONSTRAINT business_company_service_operation_hours_company_fk FOREIGN KEY (business_company_id) REFERENCES public.business_company(id);


--
-- Name: business_company_service_operation_hours business_company_service_operation_hours_service_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.business_company_service_operation_hours
    ADD CONSTRAINT business_company_service_operation_hours_service_fk FOREIGN KEY (business_service_id) REFERENCES public.business_company_service(id);


--
-- Name: business_company_service_order business_company_service_order_clinet_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.business_company_service_order
    ADD CONSTRAINT business_company_service_order_clinet_fk FOREIGN KEY (client_id) REFERENCES public.client(id);


--
-- Name: business_company_service_order business_company_service_order_service_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.business_company_service_order
    ADD CONSTRAINT business_company_service_order_service_fk FOREIGN KEY (business_service_id) REFERENCES public.business_company_service(id);


--
-- Name: business_company_service business_service_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.business_company_service
    ADD CONSTRAINT business_service_id_fk FOREIGN KEY (service_id) REFERENCES public.business_service(id);


--
-- Name: business_sub_categories_services business_services_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.business_sub_categories_services
    ADD CONSTRAINT business_services_id_fk FOREIGN KEY (business_services_id) REFERENCES public.business_service(id);


--
-- Name: business_sub_category category_id; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.business_sub_category
    ADD CONSTRAINT category_id FOREIGN KEY (category_id) REFERENCES public.business_category(id);


--
-- Name: owners_businesses owners_businesses_business_company_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.owners_businesses
    ADD CONSTRAINT owners_businesses_business_company_id_fkey FOREIGN KEY (business_company_id) REFERENCES public.business_company(id) ON UPDATE CASCADE;


--
-- Name: owners_businesses owners_businesses_business_owner_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.owners_businesses
    ADD CONSTRAINT owners_businesses_business_owner_id_fkey FOREIGN KEY (business_owner_id) REFERENCES public.business_owner(id) ON UPDATE CASCADE;


--
-- Name: business_sub_categories_services sub_categories_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.business_sub_categories_services
    ADD CONSTRAINT sub_categories_id_fk FOREIGN KEY (sub_categories_id) REFERENCES public.business_sub_category(id);


--
-- PostgreSQL database dump complete
--

