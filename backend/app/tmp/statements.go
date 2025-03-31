package database

const createUserTable = `
CREATE TABLE IF NOT EXISTS users (
    ID INTEGER PRIMARY KEY AUTOINCREMENT,
    Nickname TEXT,
    Email TEXT,
    Password TEXT,
    Gender TEXT,
    FirstName TEXT,
    LastName TEXT
)
`

const createPostTable = `
CREATE TABLE IF NOT EXISTS posts (
    ID TEXT PRIMARY KEY,
    Image BLOB,
    Title TEXT,
    Content TEXT,
    UserID TEXT,
    LikesCount INTEGER,
    DisLikesCount INTEGER,
    CommentsCount INTEGER,
    CreatedAt TEXT,
    Username TEXT,
    FirstName TEXT,
    LastName TEXT,
    Avatar TEXT,
    likeStatus TEXT
)
`

const createLikeTable = `
CREATE TABLE IF NOT EXISTS likes (
    ID TEXT PRIMARY KEY,
    AuthorID TEXT,
    PostID TEXT,
    CreationDate TEXT,
    UserName TEXT,
    FirstName TEXT,
    LastName TEXT,
    Avatar TEXT
)
`

const createCommentTable = `
CREATE TABLE IF NOT EXISTS comments (
    ID TEXT PRIMARY KEY,
    Content TEXT,
    AuthorID TEXT,
    PostID TEXT,
    UserName TEXT,
    CreationDate TEXT,
    FirstName TEXT,
    LastName TEXT,
    Avatar TEXT,
    LikesCount INTEGER,
    DisLikesCount INTEGER
)
`

const createCategoryTable = `
CREATE TABLE IF NOT EXISTS categories (
    ID TEXT PRIMARY KEY,
    Category TEXT
)
`

const createPrivateMessageTable = `
CREATE TABLE IF NOT EXISTS private_messages (
    ID INTEGER PRIMARY KEY AUTOINCREMENT,
    SenderID TEXT,
    ReceiverID TEXT,
    Content TEXT,
    CreatedAt TEXT,
    idChat INTEGER,
    FOREIGN KEY (idChat) REFERENCES conversations(ID)
)
`
const Conversations=`
CREATE TABLE IF NOT EXISTS conversations (
    ID INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id_1 INTEGER NOT NULL,
    user_id_2 INTEGER NOT NULL,
    started_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id_1) REFERENCES users(ID),
    FOREIGN KEY (user_id_2) REFERENCES users(ID)
) 
`
