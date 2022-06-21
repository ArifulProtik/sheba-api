# Sheba Prototype

## available routes and Sample POST data

### User

[POST] Signup `http://localhost:3001/api/v1/signup`

```json
{
  "name": "Protik",
  "username": "arifulprotik",
  "email": "aiprotik2@gmail.com",
  "password": "11223344",
  "confirm_password": "11223344"
}
```

[POST] SignIn `http://localhost:3001/api/v1/signin`

```json
{
  "email": "aiprotik2@gmail.com",
  "password": "11223344"
}
```

[GET] Refresh `http://localhost:3001/api/v1/refresh`
Header: token:refreshtoken

[PUT] Update User to ServiceProvider `http://localhost:3001/api/v1/user`
Authentication Required ##### Bearer Token: `access token`

### location and category {Mock}

[POST] CREATE location `http://localhost:3001/api/v1/location/create`
Authentication Required ##### Bearer Token: `access token`

```json
{
  "name": "Kafrul, Dhaka"
}
```

[POST] CREATE category `http://localhost:3001/api/v1/category/create`
Authentication Required ##### Bearer Token: `access token`

```json
{
  "name": "PC/Laptop"
}
```

### Services

[POST] New Service `http://localhost:3001/api/v1/service/create`

Authentication Required ##### Bearer Token: `access token`

```json
{
  "title": "PC Servicing",
  "instruments": ["screw", "gfx card"],
  "category": "1",
  "location": "1",
  "cost": "200",
  "additionalcost": "300"
}
```

instruments and additionalcost are optional

[GET] ALL available services `http://localhost:3001/api/v1/services`

[GET] MyServices `http://localhost:3001/api/v1/myservices`
Authentication Required ##### Bearer Token: `access token`
