CREATE TABLE Users (
                       UserID INT AUTO_INCREMENT PRIMARY KEY,
                       Username VARCHAR(50) NOT NULL UNIQUE,
                       Email VARCHAR(100) NOT NULL UNIQUE,
                       PasswordHash VARCHAR(255) NOT NULL,
                       Role VARCHAR(50), -- e.g. admin, staff, manager
                       IsActive BOOLEAN NOT NULL DEFAULT TRUE,
                       EmployeeID INT,

                       CreatedAt DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
                       UpdatedAt DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

                       CONSTRAINT FK_Users_Employees FOREIGN KEY (EmployeeID)
                           REFERENCES Employees(EmployeeID)
                           ON DELETE SET NULL
);

-- Index untuk pencarian cepat
CREATE INDEX IX_Users_Username ON Users(Username);
CREATE INDEX IX_Users_Email ON Users(Email);