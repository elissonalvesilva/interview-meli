db.auth('root', 'root')

db = db.getSiblingDB('interview')

db.createUser({
    user: "interview",
    pwd: "interview",
    roles: [{role: "readWrite", db: "collabty"}]
});