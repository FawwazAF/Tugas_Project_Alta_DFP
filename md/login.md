# Login

Login is to get jwt token. Jwt token valid for only 1 hour.

## How to do

Data constraint :

```json
{
    "email": "[valid email]",
    "password": "[password in plain text]"
}
```
* You can use body form-data to input the data

Example :

* Input

```json
{
    "email": "faw@123",
    "password": "12345678",
}
```
* Output

```json
{
    "email": "faw@123",
    "password": "12345678",
    "token" : "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2Mjc5NzM3MTAsInVzZXJJZCI6MTN9.OnncqJd99fcIkYwwT8iwkpQscWLfrjVKgeaz2hg3C8M" 
}
```
## Error 

**Condition** : if email or password is wrong/not exist
