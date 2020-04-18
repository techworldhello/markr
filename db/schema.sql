START TRANSACTION;

-- Create DB if not already initialised by docker config
CREATE DATABASE IF NOT EXISTS markr;

USE markr;

CREATE TABLE IF NOT EXISTS student_results (
  id bigint primary key AUTO_INCREMENT,
  student_number varchar(120) NOT NULL,
  test_id string varchar(120) NULL,
  first_name varchar(256) NOT NULL,
  last_name varchar(256) NOT NULL,
  total_available smallint NOT NULL,
  total_obtained smallint NOT NULL,
  scanned_on timestamp NOT NULL,
  created_at timestamp DEFAULT NULL,
  INDEX (test_id)
);

COMMIT;