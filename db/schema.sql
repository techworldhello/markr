START TRANSACTION;

CREATE DATABASE IF NOT EXISTS markr;

USE markr;

--
-- Name: student_results; Type: TABLE; Independent
--

CREATE TABLE IF NOT EXISTS student_results (
id bigint primary key AUTO_INCREMENT,
  student_number smallint NOT NULL,
  first_name varchar(256) NOT NULL,
  last_name varchar(256) NOT NULL,
  total_available smallint NOT NULL,
  total_obtained smallint NOT NULL,
  scanned_on timestamp NOT NULL,
  created_at timestamp NOT NULL
);

--
-- Name: marks; Type: TABLE; Dependent on marks
--

CREATE TABLE IF NOT EXISTS marks (
  id bigint primary key AUTO_INCREMENT,
  student_number smallint NOT NULL REFERENCES student_results(student_number),
  test_id smallint NOT NULL,
  question smallint NOT NULL,
  available smallint NOT NULL,
  obtained smallint NOT NULL,
  created_at timestamp NOT NULL
);

-- Create index on total_obtained column

ALTER TABLE student_results ADD INDEX (total_obtained);

COMMIT;
