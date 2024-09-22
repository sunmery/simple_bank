DROP INDEX accounts_owner CASCADE;
DROP INDEX transfers_from_account_id CASCADE;
DROP INDEX transfers_to_account_id CASCADE;
DROP INDEX entries_account_id_fkey CASCADE;
DROP INDEX transfers_compound CASCADE;

DROP TABLE accounts CASCADE;
DROP TABLE transfers;
DROP TABLE entries;
