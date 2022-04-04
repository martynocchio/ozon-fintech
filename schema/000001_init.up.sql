CREATE TABLE IF NOT EXISTS links
(
    id          SERIAL        NOT NULL ,
    base_url    VARCHAR(1024) NOT NULL ,
    token       VARCHAR(10)   NOT NULL ,
    UNIQUE (token)
);