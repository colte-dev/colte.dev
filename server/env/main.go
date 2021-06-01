package env

import (
	"os"
	"strconv"
	"time"

	dotenv "github.com/joho/godotenv"
)

var COOKIE_SECRET string
var GITHUB_CLIENT_ID string
var GITHUB_CLIENT_SECRET string
var DISCUSSION_REPO_OWNER string
var DISCUSSION_REPO_NAME string
var DISCUSSION_CATEGORY_ID string
var REDIS_HOST string
var REDIS_CACHE_DURATION time.Duration
var IS_DEV bool

func Init() {

	// load .env file if in development
	if os.Getenv("GO_ENV") != "production" {
		err := dotenv.Load()
		if err != nil {
			panic("Error loading .env file")
		}
	}

	COOKIE_SECRET = os.Getenv("COOKIE_SECRET")

	GITHUB_CLIENT_ID = os.Getenv("GITHUB_CLIENT_ID")
	GITHUB_CLIENT_SECRET = os.Getenv("GITHUB_CLIENT_SECRET")

	DISCUSSION_REPO_OWNER = os.Getenv("DISCUSSION_REPO_OWNER")
	DISCUSSION_REPO_NAME = os.Getenv("DISCUSSION_REPO_NAME")
	DISCUSSION_CATEGORY_ID = os.Getenv("DISCUSSION_CATEGORY_ID")

	duration, _ := strconv.Atoi(os.Getenv("REDIS_CACHE_DURATION"))
	REDIS_CACHE_DURATION = time.Second * time.Duration(duration)
	REDIS_HOST = os.Getenv("REDIS_HOST") + ":6379"

	IS_DEV = os.Getenv("GO_ENV") == "development"

}
