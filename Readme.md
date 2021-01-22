# Project Manage Data User dan Task sederhana

ini adalah project sederhana yang berfungsi untuk mengolah data user dengan CRUD secara sederhana dan memberikan task kepada user


## Getting Started
code dapat dicopy di local device dan menjalankan dengan menggunakan docker dan setelah itu dapat di access API menggunakan POSTMAN atau tool sejenisnya. collection API terdapat pada file berformat json dan env juga berformat json


### Menjalankannya Code
jalankan command dibawah ini pada directory depatu, yang berarti code sedang berjalan di docker
```
docker-compose up --build
```
### Memberhentikan Code
jalankan command dibawah ini pada directory depatu, yang berarti code sudah dihentikan
```
docker-compose down
```




## Akses API
file collection API sudah ada pada file. 
collection : Depatu.postman_collection.json
env : Depatu.postman_environment.json

## API user
API ini adalah API yang dapat digunakan oleh user

### First Page
```
Request: Welcome Page
Method: GET
Url:{{docker_address}}/
Parameter: tidak menggunakan parameter
```

### Register
```
Request: register
Method: POST
Url:{{docker_address}}/api/v1/register
Parameter: username, email, password
```
contoh parameter pada body 
```
{
    "username":"khalidalhabibie",
    "email":"khalid@@gmail.com",
    "password":"passwordkhalid"
    
}
```
username tidak bisa diupdate, oleh karena itu usernama hanya sekali buat

### SignIn
```
Request: register
Method: POST
Url:{{docker_address}}/api/v1/register
Parameter: username, password
```
contoh parameter pada body 
```
{
    "username":"khalidalhabibie",
    "password":"passwordkhalid"
    
}
```
pada tahap ini akan return token untuk Authorization. penggunakan token Authorization  dengan format 
```
bearer  + token 
```
contoh
```
bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6ZmFsc2UsImV4cCI6MTYxMTMyMjI4MiwiaWF0IjoxNjExMzExNDgyLCJ1c2VySUQiOjQsInVzZXJOYW1lIjoiaGFiaWJpZSJ9.q_bSPk0d8zy6eOQBQlkgLUCNRak5OLflecH0a9cp-Ow
```

### Get Profile
```
Request: mendapatkan data profile
Method: GET
Url:{{docker_address}}/api/v1/user/profil
Parameter: Authorization Header dengan value token dari request SignIn
```

### Update Profile
```
Request: update data profile
Method: PUT
Url:{{docker_address}}/api/v1/user/profile
Parameter: email, bio pada body dan Authorization Header dengan value token dari request SignIn
```
data lain tidak bisa diubah disini seperti password. jika dipaksakan sistem akan mengembalikan data password sebelumnya
contoh parameter pada body 
```
{
    "email":"khalid1@gmail.com",
    "bio" :"Perkenalkan saya khalid Alhabibie. saya adalah golang programmer"
    
}

```

### Menambahkan Alamat
```
Request: menambahkan alamat user
Method: POST
Url:{{docker_address}}/api/v1/user/address
Parameter: addressname, address, proviency,city, postalcode pada body dan Authorization Header dengan value token dari request SignIn
```
contoh parameter pada body 
```
{
    "addressname" : "rumah saya",
    "address"    : "jalan kuala durian nomer 897",
    "provience" : "provinsi saya",
    "city"  :"kota saya",
    "postalcode" : 23422
}


```


### Mendapatkan Address
```
Request: mendapatkan data Address User
Method: GET
Url:{{docker_address}}/api/v1/user/address
Parameter: Authorization Header dengan value token dari request SignIn
```

###Update Address
```
Request: update data address user
Method: PUT
Url:{{docker_address}}/api/v1/user/address
Parameter: ID,address, provience,city, postalcode pada body sesuai dengan keinginan dan Authorization Header dengan value token dari request SignIn
```
ID adalah Id address dimana jika bukan  user yang berhak untuk mengupdate address maka data tidak dapat diubah contoh body
```
{
    "ID" : 24,
    "provience" : "jakarta",
    "postalcode" : 2334
}
```

### Update Password
```
Request: update password User
Method: PUT
Url:{{docker_address}}/api/v1/user/password
Parameter: password pada body  dan Authorization Header dengan value token dari request SignIn
```
pada API ini sistem  berusaha mengusahkan password user saja yang diubah  jika ada mencoba data yang lain diubah maka akan membatalkan pengubahan data lain selain password
contoh body
```
{
    "password":"sayaPassword"
}
```


### Delete Address
```
Request: menghapus address
Method: DELETE
Url:{{docker_address}}/api/v1/user/password
Parameter: ID pada body  dan Authorization Header dengan value token dari request SignIn
```
ID yang dimaksud adalah ID address dan user yang berhak dapat mengahpus datanya
contoh body
```
{
     "ID" : 24
} 
```

## API Admin
API ini adalah API yang dapat digunakan oleh Admin

### First Page
```
Request: Welcome Page
Method: GET
Url:{{docker_address}}/
Parameter: tidak menggunakan parameter
```

### Register
```
Request: register
Method: POST
Url:{{docker_address}}/api/v1/register
Parameter: username, email, password, admin
```
contoh parameter pada body 
```
{
    "username":"superadmin",
    "email":"super@email",
    "password":"passwordAdmin",
    "admin":true

}
```


### SignIn
```
Request: register
Method: POST
Url:{{docker_address}}/api/v1/register
Parameter: username, password
```
contoh parameter pada body 
```
{
    "username":"superadmin",
    "password":"passwordAdmin"
    
}
```
pada tahap ini akan return token untuk Authorization. penggunakan token Authorization  dengan format 
```
bearer  + token 
```
contoh
```
bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6ZmFsc2UsImV4cCI6MTYxMTMyMjI4MiwiaWF0IjoxNjExMzExNDgyLCJ1c2VySUQiOjQsInVzZXJOYW1lIjoiaGFiaWJpZSJ9.q_bSPk0d8zy6eOQBQlkgLUCNRak5OLflecH0a9cp-Ow
```


### Get Data Semua User
```
Request: mendapatkan data Semua User
Method: GET
Url:{{docker_address}}/api/v1/admin/users
Parameter: Authorization Header dengan value token dari request SignIn
```

###Get Data Address User
```
Request: mendapatkan data address user
Method: GET
Url:{{docker_address}}/api/v1/admin/address/user
Parameter: username dan Authorization Header dengan value token dari request SignIn
```
contoh body
```
{
    "username":"khalid"
}
```


### Menambahkan Task
```
Request: menambahkan alamat user
Method: POST
Url:{{docker_address}}/api/v1/admin/task/user
Parameter: taskname, assignedto, status pada body dan Authorization Header dengan value token dari request SignIn
```
contoh parameter pada body 
```
{
    "taskname" : "buat code base",
    "assignedto" : "habibie",
    "status":"secepeatnuya"
}
```

### Get Semua Task
```
Request: mendapatkan data Semua Task
Method: GET
Url:{{docker_address}}/api/v1/admin/tasks
Parameter: Authorization Header dengan value token dari request SignIn
```

### Get Task dengan parameter user
```
Request: mendapatkan data task dengan user
Method: GET
Url:{{docker_address}}/api/v1/admin/task/user
Parameter: Authorization Header dengan value token dari request SignIn
```
contoh body 
```
{
    "assignedto":"khalid"
}
```

### Delete task
```
Request: menghapus task
Method: DELETE
Url:{{docker_address}}/api/v1/admin/task
Parameter: id pada body  dan Authorization Header dengan value token dari request SignIn
```
ID yang dimaksud adalah ID address dan user yang berhak dapat mengahpus datanya
contoh body
```
{
     "id" : 1
} 
