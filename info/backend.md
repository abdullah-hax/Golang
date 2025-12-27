- http is a protocol , protocol means rules
- following http rules there are some methods
    - GET
    - POST
    - PUT
    - PATCH
    - DELETE
    - 99% applications created using these methods
  ------

```mermaid

graph LR

a[Frontend] --request resources using GET/POST etc. methods--> b[Backend]
b --response(json/html)--> a

```

If frontend requests to create something then it should request by post method 