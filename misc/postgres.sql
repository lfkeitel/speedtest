SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner:
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner:
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;
SET default_tablespace = '';
SET default_with_oids = false;

--
-- Name: telemetry; Type: TABLE; Schema: public; Owner: speedtest
--

CREATE TABLE telemetry (
    id integer NOT NULL,
    "timestamp" timestamp DEFAULT (now() at time zone 'utc') NOT NULL,
    ip inet NOT NULL,
    ua text NOT NULL,
    dl numeric(6, 2),
    ul numeric(6, 2),
    ping numeric(6, 2),
    jitter numeric(6, 2),
    building text,
    sessionid uuid
);

-- Commented out the following line because it assumes the user of the speedtest server, @bplower
-- ALTER TABLE telemetry OWNER TO speedtest;

--
-- Name: telemetry_id_seq; Type: SEQUENCE; Schema: public; Owner: speedtest
--

CREATE SEQUENCE telemetry_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

-- Commented out the following line because it assumes the user of the speedtest server, @bplower
-- ALTER TABLE telemetry_id_seq OWNER TO speedtest;

--
-- Name: telemetry_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: speedtest
--

ALTER SEQUENCE telemetry_id_seq OWNED BY telemetry.id;


--
-- Name: telemetry id; Type: DEFAULT; Schema: public; Owner: speedtest
--

ALTER TABLE ONLY telemetry ALTER COLUMN id SET DEFAULT nextval('telemetry_id_seq'::regclass);



--
-- Name: telemetry_id_seq; Type: SEQUENCE SET; Schema: public; Owner: speedtest
--

SELECT pg_catalog.setval('telemetry_id_seq', 1, true);


--
-- Name: telemetry telemetry_pkey; Type: CONSTRAINT; Schema: public; Owner: speedtest
--

ALTER TABLE ONLY telemetry
    ADD CONSTRAINT telemetry_pkey PRIMARY KEY (id);
