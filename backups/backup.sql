--
-- PostgreSQL database dump
--

-- Dumped from database version 17.0 (Debian 17.0-1.pgdg120+1)
-- Dumped by pg_dump version 17.0 (Debian 17.0-1.pgdg120+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

ALTER TABLE ONLY public.products DROP CONSTRAINT products_supplier_id_fkey;
ALTER TABLE ONLY public.suppliers DROP CONSTRAINT suppliers_pkey;
ALTER TABLE ONLY public.products DROP CONSTRAINT products_pkey;
DROP TABLE public.suppliers;
DROP TABLE public.products;
SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: products; Type: TABLE; Schema: public; Owner: myuser
--

CREATE TABLE public.products (
    id integer NOT NULL,
    name text,
    price integer,
    supplier_id integer
);


ALTER TABLE public.products OWNER TO myuser;

--
-- Name: products_id_seq; Type: SEQUENCE; Schema: public; Owner: myuser
--

ALTER TABLE public.products ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.products_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: suppliers; Type: TABLE; Schema: public; Owner: myuser
--

CREATE TABLE public.suppliers (
    id integer NOT NULL,
    name text
);


ALTER TABLE public.suppliers OWNER TO myuser;

--
-- Name: suppliers_id_seq; Type: SEQUENCE; Schema: public; Owner: myuser
--

ALTER TABLE public.suppliers ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.suppliers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: myuser
--

COPY public.products (id, name, price, supplier_id) FROM stdin;
2	testproduct2	300	2
3	testproduct3	400	1
4	testproduct4	500	2
6	Go product 2	444	2
7	Go product 3	123	1
10	testFiberGOUpdate1	505	\N
\.


--
-- Data for Name: suppliers; Type: TABLE DATA; Schema: public; Owner: myuser
--

COPY public.suppliers (id, name) FROM stdin;
1	JaoSup1
2	JaoSup2
\.


--
-- Name: products_id_seq; Type: SEQUENCE SET; Schema: public; Owner: myuser
--

SELECT pg_catalog.setval('public.products_id_seq', 10, true);


--
-- Name: suppliers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: myuser
--

SELECT pg_catalog.setval('public.suppliers_id_seq', 2, true);


--
-- Name: products products_pkey; Type: CONSTRAINT; Schema: public; Owner: myuser
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);


--
-- Name: suppliers suppliers_pkey; Type: CONSTRAINT; Schema: public; Owner: myuser
--

ALTER TABLE ONLY public.suppliers
    ADD CONSTRAINT suppliers_pkey PRIMARY KEY (id);


--
-- Name: products products_supplier_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: myuser
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_supplier_id_fkey FOREIGN KEY (supplier_id) REFERENCES public.suppliers(id) NOT VALID;


--
-- PostgreSQL database dump complete
--

