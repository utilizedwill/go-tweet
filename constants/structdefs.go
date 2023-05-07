package constants

type TweetList struct {
	Tweets []struct {
		Tweet string `json:"tweet"`
		Date  string `json:"date"`
		Done  bool   `json:"done"`
	} `json:"tweets"`
}

type AuthCreds struct {
	APIKey            string `json:"apiKey"`
	APISecretKey      string `json:"apiSecretKey"`
	AccessToken       string `json:"accessToken"`
	AccessTokenSecret string `json:"accessTokenSecret"`
}
