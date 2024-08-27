# REST

REST or Representational state transfer is the most common guideline used to communicate between web-apps.It is a simple, uniform interface that is used to make data, content, algorithms, media, and other digital resources available through web URLs. REST APIs are the most common APIs used across the web today.. Security , performance and ease of use are the factors that we should consider while designing REST API’s.

Ref : [https://blog.postman.com/rest-api-examples/](https://blog.postman.com/rest-api-examples/)

To make an API service RESTful, six guiding constraints must be satisfied:

### Use of a uniform interface (UI)

To have a uniform interface, multiple architectural constraints are required to guide the behavior of components. Additionally, resources should be unique so they are identifiable through a single URL.

Key principles of a uniform interface include:

1. **Identification of Resources**: Each resource (like a book, user, or order) is identified by a unique URL. example :  `/api/book`
2. **Manipulation of Resources through Representations**: Resources are manipulated using their representations (like JSON or XML).
3. **Self-descriptive Messages**: Each message contains enough information to describe how to process the message.example the status code , content-type etc. 
4. **Hypermedia as the Engine of Application State (HATEOAS)**: Clients interact with the application entirely through hypermedia provided dynamically by application servers.

### Client-server based

The uniform interface separates user concerns from data storage concerns. The client’s domain concerns UI and request-gathering, while the server’s domain concerns focus on data access, workload management, and security. The separation of client and server enables each to be  developed and enhanced independently of the other.

### Stateless operations

Request from client to server must contain all of the information necessary so that the server can understand and process it accordingly. The server can’t hold any information about the client state.this is because if all the required data is received in the request application can easily scale.for example , server should not store shopping cart data , Instead, use client-side storage (like cookies or local storage) or send the cart data with each request.

### RESTful resource caching

Data within a response to a request must be labeled as cacheable or non-cacheable.

### Layered system

REST allows for an architecture composed of hierarchical layers. In doing so, each component cannot see beyond the immediate layer with which they are interacting.

## Best Practices

1. **Use HTTP status codes correctly**
https://blog.postman.com/what-are-http-status-codes/
2. **Provide informative error messages**
3. **Secure your API**
implement API security measures—such as input sanitization, authentication, and role-based access control—to protect your API and user data.
4. **Version your API**
5. **Document your API**
6. **Allow filtering, sorting, and pagination**
7. **Use nouns instead of verbs in endpoint paths**
[HTTP methods](https://blog.postman.com/what-are-http-methods/) *are* verbs, so it’s best to use nouns instead of verbs in endpoint paths.
https://localhost:8080/api/v3/products this is correct , while this is not https://localhost:8080/api/v3/getallproducts

| HTTP Method | Use |
| --- | --- |
| POST | creation |
| GET | fetch |
| PUT | update |
| DELETE | delete |
| PATCH | partial update |

idempotent request → if the same request is processed multiple times the impact will same. POST is not idempotent