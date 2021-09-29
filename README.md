# spotify-buddylist
> Fetch the Spotify friend activity via Go

## Overview

The official [Spotify API](https://developer.spotify.com/documentation/web-api/) doesn't provide a way to fetch the friend activity feed that's available on desktop app. Personally, I really like to be aware of them even if I am mobile. That's why I wrote a small piece of code to fetch them and notify myself via some sort of messaging apps.

## Usage

To use this API, first of all you need to take `sp_dc` cookie on web browser and fetch web auth token with it. Then, you'll be able to retrieve friend activity.

```go
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
	
	log.Printf("Access Token: %s\n", accessTokenResponse.AccessToken)

	friendActivityResponse, err := s.GetFriendActivity(accessTokenResponse.AccessToken)
	
	if err != nil {
		log.Fatalf("could not retrieve friend activity %s", err.Error())
	}
	
	for  _, friend := range friendActivityResponse.Friends {
		log.Printf("Friend user name: %s\n", friend.User.Name)
		log.Printf("Friend user uri: %s\n", friend.User.Uri)
		log.Printf("Track name: %s\n", friend.Track.Name)
		log.Printf("Track uri: %s\n", friend.Track.Uri)
		log.Printf("Track image url: %s\n", friend.Track.ImageUrl)
		log.Printf("Artist name: %s\n", friend.Track.Artist.Name)
		log.Printf("Artist uri: %s\n", friend.Track.Artist.Uri)
	}
}
```

The console output looks like:
```json
2021/09/29 22:52:52 Friend user name: ykpcgln
2021/09/29 22:52:52 Friend user uri: spotify:user:vmbriv70kqlo4na7o64t53f8m
2021/09/29 22:52:52 Track name: Doğum Günü
2021/09/29 22:52:52 Track uri: spotify:track:0v88eJ2do4rOyGYvm8HmPD
2021/09/29 22:52:52 Track image url: http://i.scdn.co/image/ab67616d0000b2730e3432ed74306e25fdeb11da
2021/09/29 22:52:52 Artist name: Melek Mosso
2021/09/29 22:52:52 Artist uri: spotify:artist:5IAxUWLiTMsvc1oWPrczNj
2021/09/29 22:33:10 Friend user name: beyzance96
2021/09/29 22:33:10 Friend user uri: spotify:user:beyzance96
2021/09/29 22:33:10 Track name: Kurşun Adres Sormaz Ki
2021/09/29 22:33:10 Track uri: spotify:track:19bR2DAaRATEz6vXvXxkRM
2021/09/29 22:33:10 Track image url: http://i.scdn.co/image/ab67616d0000b273930b5316cc6a202ee94ee586
2021/09/29 22:33:10 Artist name: Kenan Doğulu
2021/09/29 22:33:10 Artist uri: spotify:artist:2RQ8NtUmg5y6tfbvCwX8jI
2021/09/29 22:52:52 Friend user name: Cenk Işık
2021/09/29 22:52:52 Friend user uri: spotify:user:isik.cnk
2021/09/29 22:52:52 Track name: Haberin Olsun - Version 2
2021/09/29 22:52:52 Track uri: spotify:track:4dy1fLQonIbBMOXDRnbED5
2021/09/29 22:52:52 Track image url: http://i.scdn.co/image/ab67616d0000b27304bb292283b888f2724854e7
2021/09/29 22:33:10 Artist name: Yıldız Tilbe
2021/09/29 22:33:10 Artist uri: spotify:artist:4525TiJDbBD4mZJ7EZArT0
```

The JSON output looks like:

```json
{
  "friends": [
    {
      "timestamp": 1632655794744,
      "user": {
        "uri": "spotify:user:vmbriv70kqlo4na7o64t53f8m",
        "name": "ykpcgln"
      },
      "track": {
        "uri": "spotify:track:0v88eJ2do4rOyGYvm8HmPD",
        "name": "Doğum Günü",
        "imageUrl": "http://i.scdn.co/image/ab67616d0000b2730e3432ed74306e25fdeb11da",
        "album": {
          "uri": "spotify:album:5wYDSylUFCSQkuj0vOBNn6",
          "name": "Melek Mosso"
        },
        "artist": {
          "uri": "spotify:artist:5IAxUWLiTMsvc1oWPrczNj",
          "name": "Melek Mosso"
        },
        "context": {
          "uri": "spotify:album:5wYDSylUFCSQkuj0vOBNn6",
          "name": "Melek Mosso",
          "index": 0
        }
      }
    },
    {
      "timestamp": 1632915701742,
      "user": {
        "uri": "spotify:user:beyzance96",
        "name": "beyzance96"
      },
      "track": {
        "uri": "spotify:track:19bR2DAaRATEz6vXvXxkRM",
        "name": "Kurşun Adres Sormaz Ki",
        "imageUrl": "http://i.scdn.co/image/ab67616d0000b273930b5316cc6a202ee94ee586",
        "album": {
          "uri": "spotify:album:2jzeTjf9HiXeN1AKs6MFAi",
          "name": "Sımsıkı Sıkı Sıkı"
        },
        "artist": {
          "uri": "spotify:artist:2RQ8NtUmg5y6tfbvCwX8jI",
          "name": "Kenan Doğulu"
        },
        "context": {
          "uri": "spotify:album:2jzeTjf9HiXeN1AKs6MFAi",
          "name": "Sımsıkı Sıkı Sıkı",
          "index": 0
        }
      }
    },
    {
      "timestamp": 1632930844114,
      "user": {
        "uri": "spotify:user:isik.cnk",
        "name": "Cenk Işık",
        "imageUrl": "https://scontent-ort2-1.xx.fbcdn.net/v/t1.18169-1/cp0/p50x50/13177115_10153728369121028_8638325732363693645_n.jpg?_nc_cat=105&ccb=1-5&_nc_sid=dbb9e7&_nc_ohc=RvnNDTu-iJYAX93sriW&_nc_ht=scontent-ort2-1.xx&edm=AP4hL3IEAAAA&oh=807ecd54127f83dbe5a4ea9646adc7e0&oe=617B4F68"
      },
      "track": {
        "uri": "spotify:track:4dy1fLQonIbBMOXDRnbED5",
        "name": "Haberin Olsun - Version 2",
        "imageUrl": "http://i.scdn.co/image/ab67616d0000b27304bb292283b888f2724854e7",
        "album": {
          "uri": "spotify:album:34wbIj7ZDbZMOsoNQIIkHW",
          "name": "Haberi Olsun (New Version)"
        },
        "artist": {
          "uri": "spotify:artist:4525TiJDbBD4mZJ7EZArT0",
          "name": "Yıldız Tilbe"
        },
        "context": {
          "uri": "spotify:playlist:37i9dQZF1DZ06evO2nMcpG",
          "name": "This Is Yıldız Tilbe",
          "index": 0
        }
      }
    }
  ]
}
```

### Running the example
Once you get `sp_dc` cookie, just set to environment variables, and then you can run the following command to execute the script.

```sh
export SPDC_COOKIE={value}
```

```sh
go run main.go
```

##### This repository is inspired by valeriangalliat/spotify-buddylist which is written in Node.js.