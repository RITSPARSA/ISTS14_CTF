CREATE USER 'al-dba'@'localhost' IDENTIFIED BY 'ThisIsMuhDBpassword!';
GRANT SELECT,INSERT,UPDATE on AdminLogin.users to 'al-dba'@'localhost';
FLUSH PRIVILEGES;
