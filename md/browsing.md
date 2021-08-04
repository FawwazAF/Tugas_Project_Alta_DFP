# Browsing Product

Browsing product is an API to get list of products that available

## Product that available

Currently the products that available based on its category and type is :

* category: gadget, accessories
* type: hp,laptop

Example :

* List by category only : POST /products/gadget
* List by category and type : POST /products/accessories/laptop

## Error 

**Condition** : if 'category' and 'type' not exist in database

To see all products that available : GET /products
