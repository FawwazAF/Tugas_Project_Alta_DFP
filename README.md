# API Project alta_store

API Project alta_store is API project of fake e-commerce which 
you can test to get response as a customer

## Open Endpoints

Endpoints which not required any authentication :

* [Browsing Item](get.md) : `GET /gadget/hp`
* Register : `POST /user`
* [Login](login.md) : `POST/jwt/user`

## Authenticated Endpoints

Endpoints which need authentication to get 200 response :

* Buying Item : `POST /jwt/user/cart`
* Shopping Cart List : `GET /jwt/user/cart`

