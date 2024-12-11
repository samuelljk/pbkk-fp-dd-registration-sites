-- 1. Users Table
CREATE TABLE Users (
    user_id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    batch VARCHAR(100) NOT NULL,
    role ENUM('admin', 'student') NOT NULL DEFAULT 'student',
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 2. Universities Table (Optional)
CREATE TABLE Universities (
    university_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    country VARCHAR(100) NOT NULL
);

-- 3. Degrees Table (Optional)
CREATE TABLE Degrees (
    degree_id INT AUTO_INCREMENT PRIMARY KEY,
    university_id INT NOT NULL,
    degree_name VARCHAR(100) NOT NULL,
    department VARCHAR(100),
    duration INT NOT NULL,
    FOREIGN KEY (university_id) REFERENCES Universities(university_id) ON DELETE CASCADE
);

-- 4. Applications Table
CREATE TABLE Applications (
    application_id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    university_id INT NOT NULL,
    degree_id INT NOT NULL,
    status ENUM('pending', 'approved', 'rejected', 'finalized') NOT NULL DEFAULT 'pending',
    submitted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES Users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (university_id) REFERENCES Universities(university_id) ON DELETE CASCADE,
    FOREIGN KEY (degree_id) REFERENCES Degrees(degree_id) ON DELETE CASCADE
);

-- 5. Files Table
-- CREATE TABLE File (
--     file_id INT AUTO_INCREMENT PRIMARY KEY,
--     application_id INT NOT NULL,
--     file_path VARCHAR(255) NOT NULL,
--     file_type VARCHAR(50),
--     uploaded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     FOREIGN KEY (application_id) REFERENCES Applications(application_id) ON DELETE CASCADE
-- );

-- -- Sample data for Users Table
-- INSERT INTO Users (username, password, email, role, first_name, last_name) 
-- VALUES ('john_doe', 'password123', 'john.doe@example.com', 'student', 'John', 'Doe');

-- -- Sample data for Universities Table
-- INSERT INTO Universities (name, country) 
-- VALUES ('Example University', 'USA');

-- -- Sample data for Degrees Table
-- INSERT INTO Degrees (university_id, degree_name, department, duration) 
-- VALUES (1, 'Computer Science', 'Engineering', 4);

-- -- Sample data for Applications Table
-- INSERT INTO Applications (user_id, university_id, degree_id, status) 
-- VALUES (1, 1, 1, 'pending');