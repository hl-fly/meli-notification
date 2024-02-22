# Meli-Notification

### Descrição
- API com a finalidade de envio de notificações web 

### Features
- Autenticação JWT
- Criação de conta (User/Admin)
- Criação de notificação
- Envio de notificação web

### Libs
- https://github.com/go-gorm/gorm
- https://github.com/google/wire

### Iniciando a aplicação
- Configurar arquivo .env (seguir ex.env)
- Rodar ```docker-compose up --build```

### Collection
- No diretório ```resources``` se encontra o export da collection para o Postman

### Rotas

## POST /api/signup
- Rota para criação de usuários s na base (user/admin)
# Body
```javascript
{
    "name": "J. P. Demo",
    "password": "12345",
    "cpf": "42255278999",
    "email": "admin2@hotmail.com",
    "confirm_password": "12345",
    "type": "admin"   
}
```

## POST /api/signin
- Rota para efetuar login e receber o ```access_token```
- Atentar que para cada endpoint após esse ```access_token``` sera pedido para authorization (Bearer Token)
# Body
```javascript
{
    "email": "admin@hotmail.com",
    "password": "12345"
}
```
# Return
```javascript
{
    "refresh_token": "eyJ0b2tlbiI6IllWM1lNV0NQa3VYRkxVdEliQXR4dWF6NUx1cU9XblR5MGpwQWRWOG44QjlLM3piZDFleFNKaWFGQzRHdW5US3QiLCJ1dWlkIjoiNWIwZTZiYmQtODcxYi00YTg1LWI2ZjEtYTRmYjcyNmE3YWNjIiwidXNlcl91dWlkIjoiZGZmNjk3Y2UtNGE4Ny00NWU1LTgzMTgtY2QwMzdkZTdiYTA2In0",
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDg1ODE0MjcsImlhdCI6MTcwODU3NzgyNywiaXNzIjoiUGxpbWJvdSIsInN1YiI6ImRmZjY5N2NlLTRhODctNDVlNS04MzE4LWNkMDM3ZGU3YmEwNiIsInV1aWQiOiJkZmY2OTdjZS00YTg3LTQ1ZTUtODMxOC1jZDAzN2RlN2JhMDYiLCJuYW1lIjoiai4gai4gZGVtbyIsInJvbGUiOiJhZG1pbiJ9._-SaN6vgbbQAt5cfi9jTisBhGBuSsMGKDR1a3VEeqH0",
    "user": {
        "uuid": "dff697ce-4a87-45e5-8318-cd037de7ba06",
        "name": "j. j. demo"
    }
}
```

## POST /api/user/admin/category
- Rota para gerar nova categoria de produto
- Authorization para admin
# Body
```javascript
{
    "name": "Pet"
}
```
# Return
```javascript
{
    "uuid": "dd59497e-ea6c-444f-8c85-766d3aa5f54d"
}
```

## POST /api/user/admin/product
- Rota para gerar novo produto
- Authorization para admin
# Body
```javascript
{
    "name": "Ração",
    "category_uuid": "dd59497e-ea6c-444f-8c85-766d3aa5f54d"
}
```
# Return
```javascript
{
    "uuid": "7a484231-0254-41f7-aec7-4a3ca701562b"
}
```

## POST /api/user/user-product-history
- Rota para gerar relação entre usuários  logado e produto
- Authorization para user
# Body
```javascript
{
    "product_uuid": "7a484231-0254-41f7-aec7-4a3ca701562b"
}
```

## GET /api/user/product
- Rota que retorna todos os produtos relacionado ao usuários  logado
- Authorization para user
# Return
```javascript
[
    {
        "uuid": "7a484231-0254-41f7-aec7-4a3ca701562b",
        "name": "Ração"
    }
]
```

## POST /api/user/admin/notification
- Rota para gerar nova notificação
- Pode ter produto relacionado ou não
- O campo de target valida se a notifição deve ir para todos os usuários s ou apenas para o publico alvo do produto
- Authorization para admin
# Body
```javascript
{
    "product_uuid": "7a484231-0254-41f7-aec7-4a3ca701562b",
    "message": "Teste",
    "link": "https://www.mercadolivre.com.br/alimento-golden-premium-especial-castrados-para-gato-adulto-sabor-frango-em-sacola-de-101kg/p/MLB10190738?pdp_filters=category:MLB85870#searchVariation=MLB10190738&position=18&search_layout=grid&type=product&tracking_id=72533529-d9bc-4134-9a0c-b3ba06de4296",
    "exp_date": "2024-02-22T00:00:00Z",
    "target": true
}
```

## GET /api/user/notifications
- Rota que retorna todos as notificações relacionadas ao usuários  logado
- Authorization para user
# Return
```javascript
[
    {
        "uuid": "7a484231-0254-41f7-aec7-4a3ca701562b",
        "name": "Ração"
    }
]

