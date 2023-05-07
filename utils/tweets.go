package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-tweet/constants"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dghubble/oauth1"
)

const tweetURL = "https://api.twitter.com/2/tweets"

func compareDateAndStatus(tweet struct {
	Tweet string "json:\"tweet\""
	Date  string "json:\"date\""
	Done  bool   "json:\"done\""
}) bool {

	todaysDate := time.Now()
	date, err := time.Parse("2006-01-02", tweet.Date)
	if err != nil {
		log.Fatal(err)
	}

	if (todaysDate.After(date) && tweet.Done == false) || (todaysDate.Equal(date) && tweet.Done == false) {
		authCreds := AuthDeveloper()

		// Set up OAuth1a configuration
		config := oauth1.NewConfig(authCreds.APIKey, authCreds.APISecretKey)
		token := oauth1.NewToken(authCreds.AccessToken, authCreds.AccessTokenSecret)

		// Create an HTTP client
		httpClient := config.Client(oauth1.NoContext, token)

		//Create tweet compatible Twitter v2 API
		tweetData := map[string]string{
			"text": tweet.Tweet,
		}

		tweetPayload, err := json.Marshal(tweetData)
		if err != nil {
			log.Fatal(err)
		}

		// Send a POST request to the Twitter API to create the tweet
		req, err := http.NewRequest("POST", tweetURL, bytes.NewBuffer(tweetPayload))
		if err != nil {
			fmt.Printf("Error sending tweet: %v", err)
			os.Exit(1)
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err := httpClient.Do(req)
		if err != nil {
			fmt.Printf("Error occured while sending tweet: %v", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		// Check the response status code
		if resp.StatusCode != http.StatusCreated {
			errorResponse := struct {
				Errors []struct {
					Message string `json:"message"`
				} `json:"errors"`
			}{}
			err = json.NewDecoder(resp.Body).Decode(&errorResponse)
			if err != nil {
				fmt.Printf("Error sending tweet: Twitter API responded with status code %v", resp.StatusCode)
				os.Exit(1)
			}
			fmt.Printf("Error sending tweet: %v\n", errorResponse.Errors[0].Message)
			os.Exit(1)
		}

		// Print the response body
		fmt.Printf("%s \n tweet sent successfully", tweet.Tweet)
		return true
	}

	return false
}

func SendTweets() {
	var tweetStruct constants.TweetList

	tweetJson, err := ioutil.ReadFile("static/tweets.json")
	if err != nil {
		log.Fatal("Error happened when opening tweets.json: ", err)
	}

	err = json.Unmarshal(tweetJson, &tweetStruct)
	if err != nil {
		log.Fatal(err)
	}

	for index, tweet := range tweetStruct.Tweets {
		status := compareDateAndStatus(tweet)

		//Modify the successfully sent tweet status and update the json file.
		if status {
			tweetStruct.Tweets[index].Done = true
			updatedJson, err := json.Marshal(tweetStruct)
			if err != nil {
				log.Fatal(err)
			}

			err = ioutil.WriteFile("static/tweets.json", updatedJson, 0644)
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
		}
	}

}
