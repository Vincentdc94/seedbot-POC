package cnc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Torrent struct {
	Title string
	Url   string
	Hash  string
	Seeds int
	Peers int
}

type YifyTorrentRequest struct {
	Status         string
	Status_Message string
	Data           struct {
		MovieCount int
		Limit      int
		PageNumber int
		Movies     []struct {
			Id                        int
			Url                       string
			Imdb_Code                 string
			Title                     string
			Title_English             string
			Title_Long                string
			Slug                      string
			Year                      int
			Rating                    float64
			Runtime                   int
			Genres                    []string
			Summary                   string
			Description_Full          string
			Synopsis                  string
			Yt_Trailer_Code           string
			Language                  string
			Mpa_Rating                string
			Background_Image          string
			Background_Image_Original string
			Small_Cover_Image         string
			Medium_Cover_Image        string
			Large_Cover_Image         string
			State                     string
			Torrents                  []struct {
				Url                string
				Hash               string
				Quality            string
				Seeds              int
				Peers              int
				Size               int
				Size_Bytes         int
				Date_Uploaded      string
				Date_Uploaded_Unix int
			}
		}
	}

	Meta struct {
		Server_Time     int
		Server_Timezone string
		Api_Version     int
		Execution_Time  string
	}
}

func makeYifyRequest(seederLimit int) (YifyTorrentRequest, error) {
	var torrentReq YifyTorrentRequest

	url := "https://yts.ag/api/v2/list_movies.json?sort=seeds&limit=" + strconv.Itoa(seederLimit)
	response, httpError := http.Get(url)

	defer response.Body.Close()

	if httpError != nil {
		return torrentReq, fmt.Errorf("Http response failed")
	}

	data, readError := ioutil.ReadAll(response.Body)

	if readError != nil {
		return torrentReq, fmt.Errorf("Error reading out the HTTP response")
	}

	decodeError := json.Unmarshal(data, &torrentReq)

	if decodeError != nil {
		return torrentReq, fmt.Errorf("Error decoding json")
	}

	return torrentReq, nil
}

func getTorrent() Torrent {
	return Torrent{}
}

//Gets an array of torrents from yify
func GetYifyTorrents(seederLimit int) []Torrent {

	var torrents []Torrent

	var request YifyTorrentRequest

	request, err := makeYifyRequest(seederLimit)

	for _, movie := range request.Data.Movies {
		for _, torrent := range movie.Torrents {
			newTorrent := Torrent{
				Title: movie.Title,
				Url:   torrent.Url,
				Hash:  torrent.Hash,
				Seeds: torrent.Seeds,
				Peers: torrent.Peers}

			torrents = append(torrents, newTorrent)
		}
	}

	if err != nil {

	}

	return torrents

}
