CREATE DATABASE IF NOT EXISTS portfolio;

CREATE TABLE IF NOT EXISTS projects (
  id INT NOT NULL AUTO_INCREMENT,
  title VARCHAR(255) NOT NULL,
  summary TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS project_images (
  id INT NOT NULL AUTO_INCREMENT,
  project_id INT NOT NULL,
  image_url VARCHAR(255) NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (project_id) REFERENCES project(id)
);

CREATE TABLE IF NOT EXISTS technologies (
  id INT NOT NULL AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
)

CREATE TABLE IF NOT EXISTS project_technologies (
  project_id INT NOT NULL,
  technology_id INT NOT NULL,
  PRIMARY KEY (project_id, technology_id),
  FOREIGN KEY (project_id) REFERENCES project(id),
  FOREIGN KEY (technology_id) REFERENCES technology(id)
);

CREATE TABLE IF NOT EXISTS descriptions (
  id INT NOT NULL AUTO_INCREMENT,
  overview TEXT NOT NULL,
  challenges TEXT NOT NULL,
  learnings TEXT NOT NULL, 
);

CREATE TABLE IF NOT EXISTS project_descriptions (
  project_id INT NOT NULL,
  description_id INT NOT NULL,
  PRIMARY KEY (project_id, description_id),
  FOREIGN KEY (project_id) REFERENCES project(id),
  FOREIGN KEY (description_id) REFERENCES description(id)
);

CREATE TABLE IF NOT EXISTS project_links (
  id INT NOT NULL AUTO_INCREMENT,
  project_id INT NOT NULL,
  website_link VARCHAR(255),
  github_link VARCHAR(255),
  PRIMARY KEY (id),
  FOREIGN KEY (project_id) REFERENCES project(id)
);

CREATE TABLE IF NOT EXISTS job (
  id INT NOT NULL AUTO_INCREMENT,
  title VARCHAR(255) NOT NULL,
  company VARCHAR(255) NOT NULL,
  summary TEXT NOT NULL,
  logo_url VARCHAR(255) NOT NULL,
  website_link VARCHAR(255),
  start_date DATE NOT NULL,
  end_date DATE NOT NULL,
  PRIMARY KEY (id)
)

CREATE TABLE IF NOT EXISTS job_description (
  id INT NOT NULL AUTO_INCREMENT,
  job_id INT NOT NULL,
  description TEXT NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (job_id) REFERENCES job(id)
)

CREATE TABLE IF NOT EXISTS job_technology (
  job_id INT NOT NULL,
  technology_id INT NOT NULL,
  PRIMARY KEY (job_id, technology_id),
  FOREIGN KEY (job_id) REFERENCES job(id),
  FOREIGN KEY (technology_id) REFERENCES technology(id)
);
