package discussions

import (
	"context"
	"encoding/json"
	"fmt"

	"colte.dev/common"
	"colte.dev/env"
	"github.com/shurcooL/graphql"
)

type QueryByNumber struct {
	Repository struct {
		Discussions Discussion `graphql:"discussion(number: $number)"`
	} `graphql:"repository(owner: $owner, name: $name)"`
}

type QueryLastN struct {
	Repository struct {
		Discussions struct {
			TotalCount int
			Nodes      []Discussion
		} `graphql:"discussions(last: $last, categoryId: $categoryId)"`
	} `graphql:"repository(owner: $owner, name: $name)"`
}

type QueryByUsername struct {
	Search struct {
		Edges []struct {
			Node struct {
				Discussion `graphql:"... on Discussion"`
			}
		}
	} `graphql:"search(type: DISCUSSION, query: $query, last: $last)"`
}

type Discussion struct {
	ID          string `json:"id"`
	Number      int    `json:"number"`
	UpvoteCount int    `json:"upvoteCount"`
	Answer      struct {
		ID  string `json:"id"`
		Url string `json:"url"`
	} `json:"answer"`
	Author struct {
		Login     string `json:"login"`
		Url       string `json:"url"`
		AvatarUrl string `json:"avatarUrl"`
	} `json:"author"`
	Comments struct {
		TotalCount int `json:"totalCount"`
	} `json:"comments"`
	Reactions struct {
		TotalCount int `json:"totalCount"`
	} `json:"reactions"`
	BodyHTML  string `graphql:"bodyHTML" json:"bodyHTML"`
	Title     string `json:"title"`
	CreatedAt string `json:"createdAt"`
	Url       string `json:"url"`
}

func GetDiscussionByNumber(number int, token string) (Discussion, error) {
	redis := common.RedisClient

	cachedDiscussion, err := redis.Get(context.Background(), fmt.Sprintf("discussion/%d", number)).Result()

	if err == nil && cachedDiscussion != "" {
		var discussion Discussion
		err := json.Unmarshal([]byte(cachedDiscussion), &discussion)

		if err != nil {
			return Discussion{}, err
		}
		return discussion, nil
	} else {
		variables := map[string]interface{}{
			"owner":  graphql.String(env.DISCUSSION_REPO_OWNER),
			"name":   graphql.String(env.DISCUSSION_REPO_NAME),
			"number": graphql.Int(number),
		}

		gql := common.CreateGQLClient(token)
		var query QueryByNumber
		err := gql.Query(context.Background(), &query, variables)

		if err != nil {
			return Discussion{}, err
		}

		discussion := query.Repository.Discussions

		discussionJson, _ := json.Marshal(discussion)
		redis.Set(context.Background(), fmt.Sprintf("discussion/%d", number), discussionJson, env.REDIS_CACHE_DURATION)

		return discussion, nil
	}

}

func GetLastDiscussions(token string) ([]Discussion, error) {
	redis := common.RedisClient
	cachedDiscussions, err := redis.Get(context.Background(), "lastDiscussions").Result()

	if err == nil && cachedDiscussions != "" {
		var discussions []Discussion
		json.Unmarshal([]byte(cachedDiscussions), &discussions)
		return discussions, nil
	} else {
		variables := map[string]interface{}{
			"owner":      graphql.String(env.DISCUSSION_REPO_OWNER),
			"name":       graphql.String(env.DISCUSSION_REPO_NAME),
			"categoryId": graphql.ID(env.DISCUSSION_CATEGORY_ID),
			"last":       graphql.Int(20),
		}

		gql := common.CreateGQLClient(token)
		var query QueryLastN
		err := gql.Query(context.Background(), &query, variables)

		if err != nil {
			return nil, err
		}

		discussionsJson, _ := json.Marshal(query.Repository.Discussions.Nodes)
		redis.Set(context.Background(), "lastDiscussions", discussionsJson, env.REDIS_CACHE_DURATION)

		for _, discussion := range query.Repository.Discussions.Nodes {
			discussionJson, _ := json.Marshal(discussion)
			redis.Set(context.Background(), fmt.Sprintf("discussion/%d", discussion.Number), discussionJson, env.REDIS_CACHE_DURATION)
		}

		return query.Repository.Discussions.Nodes, nil
	}

}

func GetUserDiscussions(username string, token string) ([]Discussion, error) {
	variables := map[string]interface{}{
		"query": graphql.String(
			fmt.Sprintf(
				"author:%s repo:%s",
				username,
				env.DISCUSSION_REPO_OWNER+"/"+env.DISCUSSION_REPO_NAME,
			),
		),
		"last": graphql.Int(20),
	}

	gql := common.CreateGQLClient(token)
	var query QueryByUsername
	err := gql.Query(context.Background(), &query, variables)

	if err != nil {
		return nil, err
	}

	var discussions []Discussion

	for _, Edge := range query.Search.Edges {
		discussions = append(discussions, Edge.Node.Discussion)
	}
	return discussions, nil
}
