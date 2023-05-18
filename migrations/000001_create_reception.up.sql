CREATE TABLE reception(
  uid VARCHAR(255) PRIMARY KEY NOT NULL,
  mail VARCHAR(255) NOT NULL,
  name VARCHAR(255) NOT NULL,
  attends_first_day BOOLEAN NOT NULL,
  attends_second_day BOOLEAN NOT NULL,
  temperature_first_day FLOAT,
  temperature_second_day FLOAT
 );
