DROP INDEX IF EXISTS unique_email_idx;

ALTER table cats
    DROP CONSTRAINT user_id_fk;

ALTER TABLE cat_matches
    DROP CONSTRAINT cat_matches_issuer_id_fk;

ALTER TABLE cat_matches
    DROP CONSTRAINT cat_matches_issuer_cat_id_fk;

ALTER TABLE cat_matches
    DROP CONSTRAINT cat_matches_receiver_cat_id_fk;


DROP TABLE IF EXISTS cat_matches;
DROP TABLE IF EXISTS cats;
DROP TABLE IF EXISTS users;