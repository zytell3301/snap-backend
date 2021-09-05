CREATE TABLE IF NOT EXISTS tokens
(
    communication_token VARCHAR,
    user_id             UUID,
    PRIMARY KEY ( communication_token )
)