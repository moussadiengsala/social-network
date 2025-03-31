
## API Reference
### Overview
The Real-Time Golang REST API provides endpoints for performing CRUD operations on resources. Additionally, it supports real-time communication using WebSockets.

### Authentication
The API uses JSON Web Tokens (JWT) for authentication. Include the JWT token in the Authorization header of each request with the "Bearer" prefix.

```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```
### Endpoints, Method, Description

| endpoints             | method     | Description                |
| :--------             | :----------| :------------------------- |
| `/api/auth/signup`    | `POST`     | Creating a account.        |
| `/api/auth/signin`    | `POST`     | To identifier, which will provide you the token to Include in the header of each request.                   |
| `/api/posts`          | `POST`     |  To create a post.                        |
| `/api/posts`          | `GET`      |  To fetch some limited posts.                        |
| `/api/posts/<id>`     | `GET`      | To fetch on single post by his id.                        |
| `/api/comments`       | `POST`     | To create a comment.                           |
| `/api/comments/<id>`  | `GET`      | To fetch on single comment by his id.                          |
| `/api/reactions`      | `POST`     | To create a post like, post dislike, comment like and comment dislike.                         |

### Rules

| Fields             | size     | Description                |
| :--------             | :----------| :------------------------- |
| `Email`               | `-`     | make sure it has this format `abc@abd.abc`|
| `Username`            | `3-16`     | -           |
| `Password`          | `8-100`     |  -                       |
| `Bio`          | `0-255`      |  - |
| `Content`     | `3-255`      | The content of post or comment.|
| `FirstName`       | `3-20`     | -                         |
| `LastName`  | `3-20`      | -                        |
| `author_id`      | `1-16`     | When you create a post, comment, or reaction you need to provide the user id who perform that action. |
| `post_id`      | `1-16`     | - |
| `entries_id`      | `1-16`     | This only for creating a reaction. It can be postID or commentID. If it's postID that mean you wanna perform post like or post dislike. And comment like or comment dislike if it's commentID |
| `action`      | `-`     | Can only be one of the following values:  `postlikes`, `postdislikes`, `commentlikes`, `commentdislikes`.|


## Usage Examples
### `/api/auth/signup`

#### submited datas
```
{
    "FirstName": "Test",
    "LastName": "Test",
    "Email": "abdouaziznjay@gmail.com",
    "Username": "aziz",
    "Bio": "Software engineer at Zone01 Dakar",
    "Avatar": "http://source.unsplash.com/50x50",
    "Password": "12345678"
}
```

#### response 
```
{
  "Code": 200,
  "Message": "ok",
  "Data": null
}
```

### `/api/auth/signin`

#### submited datas
```
{
    "Identifiers": "abdouaziznjay@gmail.com",
    "Password": "12345678"
}
or
{
    "Identifiers": "aziz",
    "Password": "12345678"
}
```

#### response 
```
{
  "Code": 200,
  "Message": "ok",
  "Data": {
    "User": {
      "ID": "2",
      "FirstName": "Test",
      "LastName": "Test",
      "Email": "abdouaziznjay@gmail.com",
      "Username": "aziz",
      "Bio": "Software engineer at Zone01 Dakar",
      "Avatar": "http://source.unsplash.com/50x50",
      "Password": ""
    },
    "Session": {
      "Token": "bd13e910-71e2-42e1-83d8-90fbcab09a18",
      "ExpirationDate": "2023-12-29T17:50:06.830445085Z",
      "UserID": "2",
      "CreationDate": "2023-12-28T17:50:06.830445291Z"
    }
  }
}
```

### `/api/posts` - `POST`

#### submited datas
```
{
	"Image":   "any",
	"Title":   "Testing post feature",
	"Content": "Lorem, ipsum dolor sit amet consectetur adipisicing elit. Temporibus, illum porro explicabo earum at quod ",
	"AuthorID":  "1"
}
```

#### response 
```
{
  "Code": 200,
  "Message": "ok",
  "Data": {
    "ID": "",
    "Image": "any",
    "Title": "Testing post feature",
    "Content": "Lorem, ipsum dolor sit amet consectetur adipisicing elit. Temporibus, illum porro explicabo earum at quod ",
    "AuthorID": "1",
    "Likes": null,
    "DisLikes": null,
    "Comments": null,
    "Category": null,
    "Username": "",
    "FirstName": "",
    "LastName": "",
    "Avatar": "",
    "LikeStatus": "",
    "CreationDate": "0001-01-01T00:00:00Z"
  }
}
```

to be contiued.....
