CREATE TABLE IF NOT EXISTS sub_districts(
   id serial PRIMARY KEY,
   sub_district_id INT NOT NULL,
   name VARCHAR (150) NULL,
   updatedAt TIMESTAMP NULL,
   createdAt TIMESTAMP NULL
);