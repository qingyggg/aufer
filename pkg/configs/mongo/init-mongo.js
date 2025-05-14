// 使用aufer数据库
db = db.getSiblingDB('aufer');

// 创建集合（表）
db.createCollection('article');
db.createCollection('comment');
db.createCollection('comment_closure');
db.createCollection('notification');

// 为article集合创建索引
db.article.createIndex({ "article_id": 1 }, { unique: true });

// 为comment集合创建索引
db.comment.createIndex({ "hash_id": 1 }, { unique: true });
db.comment.createIndex({ "article_id": 1 });
db.comment.createIndex({ "article_id": 1, "hash_id": 1 }, { unique: true });

// 为comment_closure集合创建索引
db.comment_closure.createIndex({ "ancestor": 1 });
db.comment_closure.createIndex({ "descendant": 1 });
db.comment_closure.createIndex({ "article_id": 1 });

// 为notification集合创建索引
db.notification.createIndex({ "hash_id": 1 }, { unique: true });
db.notification.createIndex({ "user_id": 1 });

print("MongoDB initialization completed: Collections and indexes created successfully"); 