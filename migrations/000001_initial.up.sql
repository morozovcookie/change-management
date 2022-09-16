BEGIN;

-- -----------------------------------------------------------------------------

    CREATE TYPE change_request_type AS ENUM (
        'CRQ',
        'AutoCRQ'
    );

    COMMENT ON TYPE change_request_type IS 'change request type enumeration';

-- -----------------------------------------------------------------------------

    CREATE TABLE change_requests (
        row_id bigserial NOT NULL,

        crq_id varchar(32) NOT NULL,
        crq_type change_request_type NOT NULL,
        crq_summary varchar(255) NOT NULL,
        crq_description text NOT NULL,
        crq_is_auto_close boolean NOT NULL,

        hash varchar(32) NOT NULL,

        created_at bigint NOT NULL,
        updated_at bigint,

        PRIMARY KEY (row_id)
    );

    COMMENT ON TABLE change_requests
        IS 'stores requests for making the change of current
infrastructure state';
    COMMENT ON COLUMN change_requests.row_id IS 'table row unique identifier';
    COMMENT ON COLUMN change_requests.crq_id
        IS 'change request unique identifier';
    COMMENT ON COLUMN change_requests.crq_type IS 'change request type';
    COMMENT ON COLUMN change_requests.crq_summary
        IS 'short text describes the change';
    COMMENT ON COLUMN change_requests.crq_description
        IS 'full free form text describes the change';
    COMMENT ON COLUMN change_requests.crq_is_auto_close
        IS 'flag that indicates that change request should be closed
automatically';
    COMMENT ON COLUMN change_requests.hash
        IS 'calculated crq hash for identify it on creation stage';
    COMMENT ON COLUMN change_requests.created_at
        IS 'time when change request was created';
    COMMENT ON COLUMN change_requests.updated_at
        IS 'time when change request was updated';

    CREATE INDEX change_requests_crq_id_hash_idx ON change_requests
        USING hash (crq_id);
    COMMENT ON INDEX change_requests_crq_id_hash_idx
        IS 'searching change requests by full matches unique identifier value';

-- -----------------------------------------------------------------------------

    CREATE TABLE incidents (
        row_id bigserial NOT NULL,

        incident_id varchar(32) NOT NULL,
        incident_summary varchar(255) NOT NULL,
        incident_description text NOT NULL,

        created_at bigint NOT NULL,
        updated_at bigint,

        PRIMARY KEY (row_id)
    );

    COMMENT ON TABLE incidents IS 'stores user appeals and service request';
    COMMENT ON COLUMN incidents.row_id IS 'table row unique identifier';
    COMMENT ON COLUMN incidents.incident_id IS 'incident unique identifier';
    COMMENT ON COLUMN incidents.incident_summary
        IS 'short text describes the incident';
    COMMENT ON COLUMN incidents.incident_description
        IS 'full free form text describes the incident';
    COMMENT ON COLUMN incidents.created_at IS 'time when incident was created';
    COMMENT ON COLUMN incidents.updated_at IS 'time when incident was updated';

    CREATE INDEX incidents_incident_id ON incidents USING hash (incident_id);
    COMMENT ON INDEX incidents_incident_id
        IS 'searching incidents by full matches unique identifier value';

-- -----------------------------------------------------------------------------

COMMIT;
