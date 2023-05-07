# go-tweet

go-tweet is a tool written in Go that sends scheduled tweets defined in a JSON file using the Twitter v2 API.

## Installation

To install go-tweet, you'll need to have Go installed on your system. Then, you can install go-tweet using the following command:

```
git clone https://github.com/utilizedwill/go-tweet.git
```

## Configuration

Before you can use go-tweet, you'll need to set up a Twitter app and get your API keys. You can do this by following the instructions in the Twitter developer documentation.

Before you can use go-tweet, you'll need to modify a `auth.json` file in the root/static directory of your project with your Twitter API credentials.  You can visit [Twitter developer documentation page](https://developer.twitter.com/en/docs/twitter-api/getting-started/about-twitter-api) to create Twitter API credentials. The file should have the following structure:

```
{
    "apiKey": "",
    "apiSecretKey": "",
    "accessToken": "",
    "accessTokenSecret": ""
}
```

Replace `apiKey`, `apiSecretKey`, `accessToken`, and `accessTokenSecret` with your actual API keys and access tokens.

## Usage

To use go-tweet, you'll need to define your tweets in a JSON file called `tweets.json`. The file should have the following structure:

```
{
    "tweets": [
        {
            "tweet": "First tweet",
            "date": "2023-05-07",
            "done": true
        },
        {
            "tweet": "Second tweet",
            "date": "2023-05-08",
            "done": false
        },
        {
            "tweet": "Thirt tweet",
            "date": "2023-05-08",
            "done": false
        },
        {
            "tweet": "Fourth tweet",
            "date": "2023-05-08",
            "done": false
        }
    ]
}
```

Each tweet should have a `tweet` field containing the text of the tweet, a `date` field containing the date the tweet should be sent (in `YYYY-MM-DD` format), and a `done` field that is set to `false` initially.

To send your scheduled tweets, run the following command:

```
go run main.go
```

This command will check the current date and send any tweets that have a scheduled date in the past and a `done` status of `false`. The `done` status of these tweets will be set to `true` to avoid sending them again in the future.

## Contributing

If you find a bug or have a feature request, please create an issue on the [GitHub repository](https://github.com/utilizedwill/go-tweet/issues).

If you want to contribute code, please fork the repository and create a pull request with your changes. We welcome contributions from the community!

## License

go-tweet is licensed under the [MIT License](https://opensource.org/licenses/MIT).