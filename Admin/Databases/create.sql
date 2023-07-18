-- 创建文章表
CREATE TABLE IF NOT EXISTS articles (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title VARCHAR(30) NOT NULL,
    content TEXT NOT NULL,
    created INTEGER, -- 发布时间
    updated INTEGER -- 更改时间
);

-- 创建标签-文章关联表
-- CREATE TABLE IF NOT EXISTS post_tags (
--                                          article_id INTEGER,
--                                          tag_id INTEGER,
--                                          FOREIGN KEY (article_id) REFERENCES articles(id),
--                                          FOREIGN KEY (tag_id) REFERENCES tags(id)
-- );
--
-- -- 创建评论表
-- CREATE TABLE IF NOT EXISTS comments (
--                                         id INTEGER PRIMARY KEY,
--                                         content VARCHAR(100),
--                                         uname VARCHAR(10),
--                                         email VARCHAR(20),
--                                         website VARCHAR(50),
--                                         article_id INTEGER,
--                                         created INTEGER
-- );
