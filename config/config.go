package config

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

const Uri = "mongodb://0.0.0.0:27017"
const DataTable = "data"
const SpiceDB = "spice_files"

var Cred = options.Credential{AuthSource: "admin", Username: "root", Password: "root"}

const FilePath = "/Users/user/Documents/Polite_request/Electric Characteristics Table & Graphics - Part Number A767KN186M1HLAE050.xlsx"
const SheetIndex = 0

const DefaultPort = "8081"
