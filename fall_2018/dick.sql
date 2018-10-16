CREATE TABLE stops (
  id integer NOT NULL,
  board_id integer NOT NULL,
  data jsonb
);

SELECT data->>'name' AS name FROM stops