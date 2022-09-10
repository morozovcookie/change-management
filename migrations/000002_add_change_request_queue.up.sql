BEGIN;
    CREATE TABLE queue (
        row_id bigserial NOT NULL,
        content jsonb NOT NULL,
        fail_count smallint DEFAULT 0 NOT NULL,
        last_error text DEFAULT '' NOT NULL,
        created_at bigint NOT NULL,
        updated_at bigint,

        PRIMARY KEY (row_id)
    );

    COMMENT ON TABLE queue
        IS 'stores events that should processed in fifo order';
    COMMENT ON COLUMN queue.row_id IS 'row unique identifier';
    COMMENT ON COLUMN queue.content IS 'queued event';
    COMMENT ON COLUMN queue.fail_count
        IS 'approximate count of failed processing of event';
    COMMENT ON COLUMN queue.last_error
        IS 'content of the last error caused failed event processing';
    COMMENT ON COLUMN queue.created_at IS 'time when row was created';
    COMMENT ON COLUMN queue.updated_at IS 'time when row was updated';

-- -----------------------------------------------------------------------------

    CREATE TABLE change_request_queue ()
    INHERITS (queue);

    COMMENT ON TABLE change_request_queue IS 'change request event queue';

-- -----------------------------------------------------------------------------

    CREATE TABLE incident_queue ()
    INHERITS (queue);

    COMMENT ON TABLE incident_queue IS 'incident event queue';

-- -----------------------------------------------------------------------------

COMMIT;
