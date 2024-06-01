======= Tutorial==========
buat database di mysql dengan nama finalgo


buat user zul dengan password :
CREATE USER 'zul'@'localhost' IDENTIFIED BY '123';


Berikan akses ke user zul untuk database finalgo:
GRANT ALL PRIVILEGES ON finalgo.* TO 'zul'@'localhost';

Uji Koneksi dengan go run main.go
http://localhost:8080/


Uji Endpoint
curl -X POST http://localhost:8080/users/register -H "Content-Type: application/json" -d '{"username":"testuser","email":"test@example.com","password":"password"}'



======terima kasih====
