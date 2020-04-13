START TRANSACTION;

-- Create database markr if not already initialised by docker config

CREATE DATABASE IF NOT EXISTS markr;

USE markr;

-- Create table within DB named student_result

CREATE TABLE IF NOT EXISTS student_result (
  id bigint primary key AUTO_INCREMENT,
  student_number smallint NOT NULL,
  test_id smallint NOT NULL,
  first_name varchar(256) NOT NULL,
  last_name varchar(256) NOT NULL,
  total_available smallint NOT NULL,
  total_obtained smallint NOT NULL,
  scanned_on timestamp NOT NULL,
  created_at timestamp NOT NULL
);

-- Create index on total_obtained column

ALTER TABLE student_result ADD INDEX (total_obtained);

COMMIT;