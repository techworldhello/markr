START TRANSACTION;

-- Create DB if not already initialised by docker config
CREATE DATABASE IF NOT EXISTS markr;

USE markr;

CREATE TABLE IF NOT EXISTS student_result (
  id bigint primary key AUTO_INCREMENT,
  student_number smallint NOT NULL,
  test_id smallint NOT NULL,
  first_name varchar(256) NOT NULL,
  last_name varchar(256) NOT NULL,
  total_available smallint NOT NULL,
  total_obtained smallint NOT NULL,
  scanned_on timestamp NOT NULL,
  created_at timestamp,
  INDEX (test_id)
);

SET time_zone='+00:00';

COMMIT;