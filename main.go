package main

import (
	"log"
	"os"
	"spotify-buddylist/pkg/spotify"
)

func main() {
	spdcCookie := os.Getenv("SPDC_COOKIE")

	s := spotify.New()

	accessTokenResponse, err := s.GetWebAccessToken(spdcCookie)

	if err != nil {
		log.Fatalf("could not retrieve web access token %s", err.Error())
	}

	log.Printf("Client Id: %s\n", accessTokenResponse.ClientId)
	log.Printf("Access Token: %s\n", accessTokenResponse.AccessToken)
	log.Printf("Access Token Expiration Timestamp: %d\n", accessTokenResponse.AccessTokenExpirationTimestampMs)
	log.Printf("Is Anonymous: %t\n", accessTokenResponse.IsAnonymous)

	friendActivityResponse, err := s.GetFriendActivity(accessTokenResponse.AccessToken)

	if err != nil {
		log.Fatalf("could not retrieve friend activity %s", err.Error())
	}

	for _, friend := range friendActivityResponse.Friends {
		log.Printf("Friend user name: %s\n", friend.User.Name)
		log.Printf("Friend user uri: %s\n", friend.User.Uri)
		log.Printf("Track name: %s\n", friend.Track.Name)
		log.Printf("Track uri: %s\n", friend.Track.Uri)
		log.Printf("Track image url: %s\n", friend.Track.ImageUrl)
		log.Printf("Artist name: %s\n", friend.Track.Artist.Name)
		log.Printf("Artist uri: %s\n", friend.Track.Artist.Uri)
	}
}
