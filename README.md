# Golang-Snippetbox-Windows
Snippetbox project for windows.
To use this application you need to change directories in files main.go and routes.go.
Run your MySql server and use this commands:
-- Create a new UTF-8 `snippetbox` database. 
CREATE DATABASE snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci; 
 
-- Switch to using the `snippetbox` database. 
USE snippetbox;
-- Create a `snippets` table. 
CREATE TABLE snippets ( 
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT, 
    title VARCHAR(100) NOT NULL, 
    content TEXT NOT NULL, 
    created DATETIME NOT NULL, 
    expires DATETIME NOT NULL 
); 
 
-- Add an index on the created column. 
CREATE INDEX idx_snippets_created ON snippets(created);
Then create a new user.
CREATE USER 'web'@'localhost'; 
GRANT SELECT, INSERT ON snippetbox.* TO 'web'@'localhost'; 
-- Important: Make sure to swap 'pass' with a password of your own choosing. 
ALTER USER 'web'@'localhost' IDENTIFIED BY 'pass';
Then use this commands in your terminal window: "cd $HOME/.../snippetbox/cmd/web" and "go run .".
