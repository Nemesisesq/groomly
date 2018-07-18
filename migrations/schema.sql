--
-- PostgreSQL database dump
--

-- Dumped from database version 10.3
-- Dumped by pg_dump version 10.2

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
-- Name: fatal_attributes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE fatal_attributes (
    id uuid NOT NULL,
    name text NOT NULL,
    weight integer NOT NULL,
    summary text,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE fatal_attributes OWNER TO postgres;

--
-- Name: metric_values; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE metric_values (
    id uuid NOT NULL,
    metric_id uuid NOT NULL,
    value_id uuid NOT NULL,
    opportunity_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE metric_values OWNER TO postgres;

--
-- Name: metrics; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE metrics (
    id uuid NOT NULL,
    name text NOT NULL,
    weight integer NOT NULL,
    type integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE metrics OWNER TO postgres;

--
-- Name: opportunities; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE opportunities (
    id uuid NOT NULL,
    name text NOT NULL,
    summary text NOT NULL,
    business_category character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE opportunities OWNER TO postgres;

--
-- Name: opportunity_fatal_attributes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE opportunity_fatal_attributes (
    id uuid NOT NULL,
    opportunity_id uuid NOT NULL,
    fatal_attribute_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE opportunity_fatal_attributes OWNER TO postgres;

--
-- Name: project_reports; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE project_reports (
    id uuid NOT NULL,
    name text NOT NULL,
    summary text NOT NULL,
    min_hurdle integer NOT NULL,
    max_hurdle integer NOT NULL,
    rank integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE project_reports OWNER TO postgres;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE schema_migration (
    version character varying(255) NOT NULL
);


ALTER TABLE schema_migration OWNER TO postgres;

--
-- Name: values; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE "values" (
    id uuid NOT NULL,
    name text NOT NULL,
    score integer NOT NULL,
    metric_choice_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE "values" OWNER TO postgres;

--
-- Name: fatal_attributes fatal_attributes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY fatal_attributes
    ADD CONSTRAINT fatal_attributes_pkey PRIMARY KEY (id);


--
-- Name: metric_values metric_values_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY metric_values
    ADD CONSTRAINT metric_values_pkey PRIMARY KEY (id);


--
-- Name: metrics metrics_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY metrics
    ADD CONSTRAINT metrics_pkey PRIMARY KEY (id);


--
-- Name: opportunities opportunities_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY opportunities
    ADD CONSTRAINT opportunities_pkey PRIMARY KEY (id);


--
-- Name: opportunity_fatal_attributes opportunity_fatal_attributes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY opportunity_fatal_attributes
    ADD CONSTRAINT opportunity_fatal_attributes_pkey PRIMARY KEY (id);


--
-- Name: project_reports project_reports_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY project_reports
    ADD CONSTRAINT project_reports_pkey PRIMARY KEY (id);


--
-- Name: values values_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY "values"
    ADD CONSTRAINT values_pkey PRIMARY KEY (id);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- PostgreSQL database dump complete
--

