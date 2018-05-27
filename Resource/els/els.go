package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	elastic "gopkg.in/olivere/elastic.v3"
)

type Tweet struct {
	User    string `json:"user"`
	Message string `json:"message"`
}

func main() {
	// Create a client
	client, err := elastic.NewClient(
		elastic.SetURL("http://172.16.60.28:9200"),
		elastic.SetMaxRetries(3),
	)

	// Obtain a client and connect to the default Elasticsearch installation
	// on 127.0.0.1:9200. Of course you can configure your client to connect
	// to other hosts and configure it in various other ways.
	//client, err := elastic.NewClient()
	if err != nil {
		fmt.Printf("new client is error\n")
		return
	}

	// Ping the Elasticsearch server to get e.g. the version number
	info, code, err := client.Ping("http://172.16.60.28:9200").Do()
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	// Getting the ES version number is quite common, so there's a shortcut
	esversion, err := client.ElasticsearchVersion("http://172.16.60.28:9200")
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)

	// Use the IndexExists service to check if a specified index exists.
	exists, err := client.IndexExists("twitter").Do()
	if err != nil {
		// Handle error
		panic(err)
	}

	fmt.Printf("\n\n")

	if !exists {
		// Create an index
		_, err = client.CreateIndex("twitter").Do()
		if err != nil {
			fmt.Printf("create index is error")
			return
		}
	}

	/*
		// Add a document to the index
		tweet := Tweet{User: "olivere", Message: "456"}
		_, err = client.Index().
			Index("twitter").
			Type("tweet").
			Id("1").
			BodyJson(tweet).
			Refresh(true).
			Do()
		if err != nil {
			// Handle error
			panic(err)
		}

		tweet1 := Tweet{User: "olivere", Message: "123"}
		_, err = client.Index().
			Index("twitter").
			Type("tweet").
			Id("2"). // id 如果一样, 就覆盖了
			BodyJson(tweet1).
			Refresh(true).
			Do()
	*/

	// Get tweet with specified ID
	get1, err := client.Get().
		Index("twitter").
		Type("tweet").
		Id("2").
		Do()
	if err != nil {
		// Handle error
		panic(err)
	}
	if get1.Found {
		fmt.Printf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)
	}

	fmt.Printf("\n\n")

	//
	//
	// Search with a term query
	// 查询,要根据相同的类型,如user, 两个都是同样的user. 才能被search出来.
	//User    string `json:"user"`
	//
	termQuery := elastic.NewTermQuery("user", "olivere")
	searchResult, err := client.Search().
		Index("twitter").   // search in index "twitter"
		Query(termQuery).   // specify the query
		Sort("user", true). // sort by "user" field, ascending
		From(0).Size(10).   // take documents 0-9
		Pretty(true).       // pretty print request and response JSON
		Do()                // execute
	if err != nil {
		// Handle error
		panic(err)
	}

	// searchResult is of type SearchResult and returns hits, suggestions,
	// and all kinds of other information from Elasticsearch.
	fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)

	// Each is a convenience function that iterates over hits in a search result.
	// It makes sure you don't need to check for nil values in the response.
	// However, it ignores errors in serialization. If you want full control
	// over iterating the hits, see below.
	var ttyp Tweet
	for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
		if t, ok := item.(Tweet); ok {
			fmt.Printf("Tweet by %s: %s\n", t.User, t.Message)
		}
	}
	fmt.Printf("================================\n")

	// TotalHits is another convenience function that works even when something goes wrong.
	fmt.Printf("Found a total of %d tweets\n", searchResult.TotalHits())

	// Here's how you iterate through results with full control over each step.
	if searchResult.Hits.TotalHits > 0 {
		fmt.Printf("Found a total of %d tweets\n", searchResult.Hits.TotalHits)

		// Iterate through results
		for _, hit := range searchResult.Hits.Hits {
			// hit.Index contains the name of the index

			// Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).
			var t Tweet
			err := json.Unmarshal(*hit.Source, &t)
			if err != nil {
				// Deserialization failed
			}

			// Work with tweet
			fmt.Printf("Tweet by %s: %s\n", t.User, t.Message)
		}
	} else {
		// No hits
		fmt.Print("Found no tweets\n")
	}

	/*
		// Delete the index again
		_, err = client.DeleteIndex("twitter").Do()
		if err != nil {
			// Handle error
			panic(err)
		}
	*/
}
