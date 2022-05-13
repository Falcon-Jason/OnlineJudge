#!/usr/bin/mongo

db = db.getSiblingDB('online_judge')
db.user.drop()

db.user.createIndex({username: 1}, {unique: true})
db.user.insertOne({username: "admin", password: "admin", is_admin: true})

db.problem.drop()
db.problem.createIndex({problem_no: 1}, {unique: true})
db.problem.createIndex({title: "text"})