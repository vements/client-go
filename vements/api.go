/*
Copyright 2023 Monster Street Systems LLC

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the “Software”), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

// NB: This is a generated file; any changes will be lost.

package vements

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
)

type AchievementEndpoint struct {
	ApiKey  string
	BaseUrl string
	Rc      *resty.Client
}

func (c *AchievementEndpoint) Request() *resty.Request {
	return c.Rc.R().SetHeader("x-api-key", c.ApiKey)
}

/*
   Achievement leaderboard

   Reads and returns achievement leaderboard
*/
func (endpoint *AchievementEndpoint) Leaderboard(
	achievementId string,
) (*AchievementLeaderboardResponse, error) {
	val := AchievementLeaderboardResponse{}
	url := "achievement/{achievement_id}/leaderboard"
	url = strings.Replace(url, "{achievement_id}", achievementId, -1)

	res, err := endpoint.Request().
		SetResult(&val).
		Get(endpoint.BaseUrl + url)

	if res.StatusCode() > 299 || res.StatusCode() < 200 {
		fmt.Fprintf(os.Stderr, "%+v\n", res.RawResponse)
		return nil, fmt.Errorf("achievement leaderboard returned status %v", res.StatusCode())
	}

	return &val, err
}

/*
   Record achievement progress

   Records and returns achievement progress
*/
func (endpoint *AchievementEndpoint) Record(
	achievementId string,
	body AchievementProgressRequest,
) (*AchievementProgressResponse, error) {
	val := AchievementProgressResponse{}
	url := "achievement/{achievement_id}/progress"
	url = strings.Replace(url, "{achievement_id}", achievementId, -1)

	res, err := endpoint.Request().
		SetResult(&val).
		SetBody(&body).
		Put(endpoint.BaseUrl + url)

	if res.StatusCode() > 299 || res.StatusCode() < 200 {
		fmt.Fprintf(os.Stderr, "%+v\n", res.RawResponse)
		return nil, fmt.Errorf("achievement record returned status %v", res.StatusCode())
	}

	return &val, err
}

/*
   List achievements

   Reads and returns list of achievements in the given project
*/
func (endpoint *AchievementEndpoint) List(
	projectId string,
	limit int,
	offset int,
) (*AchievementListResponse, error) {
	val := AchievementListResponse{}
	url := "achievement"

	res, err := endpoint.Request().
		SetResult(&val).
		SetQueryParam("project_id", fmt.Sprintf("%v", projectId)).
		SetQueryParam("limit", fmt.Sprintf("%v", limit)).
		SetQueryParam("offset", fmt.Sprintf("%v", offset)).
		Get(endpoint.BaseUrl + url)

	if res.StatusCode() > 299 || res.StatusCode() < 200 {
		fmt.Fprintf(os.Stderr, "%+v\n", res.RawResponse)
		return nil, fmt.Errorf("achievement list returned status %v", res.StatusCode())
	}

	return &val, err
}

/*
   Create achievement

   Creates and returns achievement in the given project
*/
func (endpoint *AchievementEndpoint) Create(
	body AchievementCreateRequest,
) (*AchievementCreateResponse, error) {
	val := AchievementCreateResponse{}
	url := "achievement"

	res, err := endpoint.Request().
		SetResult(&val).
		SetBody(&body).
		Put(endpoint.BaseUrl + url)

	if res.StatusCode() > 299 || res.StatusCode() < 200 {
		fmt.Fprintf(os.Stderr, "%+v\n", res.RawResponse)
		return nil, fmt.Errorf("achievement create returned status %v", res.StatusCode())
	}

	return &val, err
}

/*
   Read achievement

   Reads and returns achievement by achievement id
*/
func (endpoint *AchievementEndpoint) Read(
	achievementId string,
) (*AchievementReadResponse, error) {
	val := AchievementReadResponse{}
	url := "achievement/{achievement_id}"
	url = strings.Replace(url, "{achievement_id}", achievementId, -1)

	res, err := endpoint.Request().
		SetResult(&val).
		Get(endpoint.BaseUrl + url)

	if res.StatusCode() > 299 || res.StatusCode() < 200 {
		fmt.Fprintf(os.Stderr, "%+v\n", res.RawResponse)
		return nil, fmt.Errorf("achievement read returned status %v", res.StatusCode())
	}

	return &val, err
}

/*
   Update achievement

   Updates and returns achievement by achievement id
*/
func (endpoint *AchievementEndpoint) Update(
	achievementId string,
	body AchievementUpdateRequest,
) (*AchievementUpdateResponse, error) {
	val := AchievementUpdateResponse{}
	url := "achievement/{achievement_id}"
	url = strings.Replace(url, "{achievement_id}", achievementId, -1)

	res, err := endpoint.Request().
		SetResult(&val).
		SetBody(&body).
		Post(endpoint.BaseUrl + url)

	if res.StatusCode() > 299 || res.StatusCode() < 200 {
		fmt.Fprintf(os.Stderr, "%+v\n", res.RawResponse)
		return nil, fmt.Errorf("achievement update returned status %v", res.StatusCode())
	}

	return &val, err
}

/*
   Delete achievement by id.

   Delete achievement by achievement id
*/
func (endpoint *AchievementEndpoint) Delete(
	achievementId string,
) (*AchievementDeleteResponse, error) {
	val := AchievementDeleteResponse{}
	url := "achievement/{achievement_id}"
	url = strings.Replace(url, "{achievement_id}", achievementId, -1)

	res, err := endpoint.Request().
		SetResult(&val).
		Delete(endpoint.BaseUrl + url)

	if res.StatusCode() > 299 || res.StatusCode() < 200 {
		fmt.Fprintf(os.Stderr, "%+v\n", res.RawResponse)
		return nil, fmt.Errorf("achievement delete returned status %v", res.StatusCode())
	}

	return &val, err
}

type ParticipantEndpoint struct {
	ApiKey  string
	BaseUrl string
	Rc      *resty.Client
}

func (c *ParticipantEndpoint) Request() *resty.Request {
	return c.Rc.R().SetHeader("x-api-key", c.ApiKey)
}

/*
   Participant progress

   Reads and returns participant progress.
*/
func (endpoint *ParticipantEndpoint) Progress(
	participantId string,
) (*ParticipantProgressResponse, error) {
	val := ParticipantProgressResponse{}
	url := "participant/{participant_id}/progress"
	url = strings.Replace(url, "{participant_id}", participantId, -1)

	res, err := endpoint.Request().
		SetResult(&val).
		Get(endpoint.BaseUrl + url)

	if res.StatusCode() > 299 || res.StatusCode() < 200 {
		fmt.Fprintf(os.Stderr, "%+v\n", res.RawResponse)
		return nil, fmt.Errorf("participant progress returned status %v", res.StatusCode())
	}

	return &val, err
}

/*
   Participant scores

   Reads and returns participant scores.
*/
func (endpoint *ParticipantEndpoint) Scores(
	participantId string,
) (*ParticipantScoresResponse, error) {
	val := ParticipantScoresResponse{}
	url := "participant/{participant_id}/scores"
	url = strings.Replace(url, "{participant_id}", participantId, -1)

	res, err := endpoint.Request().
		SetResult(&val).
		Get(endpoint.BaseUrl + url)

	if res.StatusCode() > 299 || res.StatusCode() < 200 {
		fmt.Fprintf(os.Stderr, "%+v\n", res.RawResponse)
		return nil, fmt.Errorf("participant scores returned status %v", res.StatusCode())
	}

	return &val, err
}

/*
   List participants

   Reads and returns list of participants in the given project
*/
func (endpoint *ParticipantEndpoint) List(
	projectId string,
	limit int,
	offset int,
) (*ParticipantListResponse, error) {
	val := ParticipantListResponse{}
	url := "participant"

	res, err := endpoint.Request().
		SetResult(&val).
		SetQueryParam("project_id", fmt.Sprintf("%v", projectId)).
		SetQueryParam("limit", fmt.Sprintf("%v", limit)).
		SetQueryParam("offset", fmt.Sprintf("%v", offset)).
		Get(endpoint.BaseUrl + url)

	if res.StatusCode() > 299 || res.StatusCode() < 200 {
		fmt.Fprintf(os.Stderr, "%+v\n", res.RawResponse)
		return nil, fmt.Errorf("participant list returned status %v", res.StatusCode())
	}

	return &val, err
}

/*
   Create participant

   Creates and returns participant in the given project
*/
func (endpoint *ParticipantEndpoint) Create(
	body ParticipantCreateRequest,
) (*ParticipantCreateResponse, error) {
	val := ParticipantCreateResponse{}
	url := "participant"

	res, err := endpoint.Request().
		SetResult(&val).
		SetBody(&body).
		Put(endpoint.BaseUrl + url)

	if res.StatusCode() > 299 || res.StatusCode() < 200 {
		fmt.Fprintf(os.Stderr, "%+v\n", res.RawResponse)
		return nil, fmt.Errorf("participant create returned status %v", res.StatusCode())
	}

	return &val, err
}

/*
   Read participant

   Reads and returns participant by participant id
*/
func (endpoint *ParticipantEndpoint) Read(
	participantId string,
) (*ParticipantReadResponse, error) {
	val := ParticipantReadResponse{}
	url := "participant/{participant_id}"
	url = strings.Replace(url, "{participant_id}", participantId, -1)

	res, err := endpoint.Request().
		SetResult(&val).
		Get(endpoint.BaseUrl + url)

	if res.StatusCode() > 299 || res.StatusCode() < 200 {
		fmt.Fprintf(os.Stderr, "%+v\n", res.RawResponse)
		return nil, fmt.Errorf("participant read returned status %v", res.StatusCode())
	}

	return &val, err
}

/*
   Update participant

   Updates and returns participant by participant id
*/
func (endpoint *ParticipantEndpoint) Update(
	participantId string,
	body ParticipantUpdateRequest,
) (*ParticipantUpdateResponse, error) {
	val := ParticipantUpdateResponse{}
	url := "participant/{participant_id}"
	url = strings.Replace(url, "{participant_id}", participantId, -1)

	res, err := endpoint.Request().
		SetResult(&val).
		SetBody(&body).
		Post(endpoint.BaseUrl + url)

	if res.StatusCode() > 299 || res.StatusCode() < 200 {
		fmt.Fprintf(os.Stderr, "%+v\n", res.RawResponse)
		return nil, fmt.Errorf("participant update returned status %v", res.StatusCode())
	}

	return &val, err
}

/*
   Delete participant by id.

   Delete participant by participant id
*/
func (endpoint *ParticipantEndpoint) Delete(
	participantId string,
) (*ParticipantDeleteResponse, error) {
	val := ParticipantDeleteResponse{}
	url := "participant/{participant_id}"
	url = strings.Replace(url, "{participant_id}", participantId, -1)

	res, err := endpoint.Request().
		SetResult(&val).
		Delete(endpoint.BaseUrl + url)

	if res.StatusCode() > 299 || res.StatusCode() < 200 {
		fmt.Fprintf(os.Stderr, "%+v\n", res.RawResponse)
		return nil, fmt.Errorf("participant delete returned status %v", res.StatusCode())
	}

	return &val, err
}

type ScoreboardEndpoint struct {
	ApiKey  string
	BaseUrl string
	Rc      *resty.Client
}

func (c *ScoreboardEndpoint) Request() *resty.Request {
	return c.Rc.R().SetHeader("x-api-key", c.ApiKey)
}

/*
   Record a scoreboard score

   Records and returns a scoreboard score.
*/
func (endpoint *ScoreboardEndpoint) Record(
	scoreboardId string,
	body ScoreboardScoreRequest,
) (*ScoreboardScoreResponse, error) {
	val := ScoreboardScoreResponse{}
	url := "scoreboard/{scoreboard_id}/score"
	url = strings.Replace(url, "{scoreboard_id}", scoreboardId, -1)

	res, err := endpoint.Request().
		SetResult(&val).
		SetBody(&body).
		Put(endpoint.BaseUrl + url)

	if res.StatusCode() > 299 || res.StatusCode() < 200 {
		fmt.Fprintf(os.Stderr, "%+v\n", res.RawResponse)
		return nil, fmt.Errorf("scoreboard record returned status %v", res.StatusCode())
	}

	return &val, err
}

/*
   Scoreboard scores

   Reads and returns scoreboard scores
*/
func (endpoint *ScoreboardEndpoint) Scores(
	scoreboardId string,
	from time.Time,
	to time.Time,
) (*ScoreboardScoresResponse, error) {
	val := ScoreboardScoresResponse{}
	url := "scoreboard/{scoreboard_id}/scores"
	url = strings.Replace(url, "{scoreboard_id}", scoreboardId, -1)

	res, err := endpoint.Request().
		SetResult(&val).
		SetQueryParam("from", fmt.Sprintf("%v", from)).
		SetQueryParam("to", fmt.Sprintf("%v", to)).
		Get(endpoint.BaseUrl + url)

	if res.StatusCode() > 299 || res.StatusCode() < 200 {
		fmt.Fprintf(os.Stderr, "%+v\n", res.RawResponse)
		return nil, fmt.Errorf("scoreboard scores returned status %v", res.StatusCode())
	}

	return &val, err
}

/*
   List scoreboards

   Reads and returns list of scoreboards in the given project
*/
func (endpoint *ScoreboardEndpoint) List(
	projectId string,
	limit int,
	offset int,
) (*ScoreboardListResponse, error) {
	val := ScoreboardListResponse{}
	url := "scoreboard"

	res, err := endpoint.Request().
		SetResult(&val).
		SetQueryParam("project_id", fmt.Sprintf("%v", projectId)).
		SetQueryParam("limit", fmt.Sprintf("%v", limit)).
		SetQueryParam("offset", fmt.Sprintf("%v", offset)).
		Get(endpoint.BaseUrl + url)

	if res.StatusCode() > 299 || res.StatusCode() < 200 {
		fmt.Fprintf(os.Stderr, "%+v\n", res.RawResponse)
		return nil, fmt.Errorf("scoreboard list returned status %v", res.StatusCode())
	}

	return &val, err
}

/*
   Create scoreboard

   Creates and returns scoreboard in the given project
*/
func (endpoint *ScoreboardEndpoint) Create(
	body ScoreboardCreateRequest,
) (*ScoreboardCreateResponse, error) {
	val := ScoreboardCreateResponse{}
	url := "scoreboard"

	res, err := endpoint.Request().
		SetResult(&val).
		SetBody(&body).
		Put(endpoint.BaseUrl + url)

	if res.StatusCode() > 299 || res.StatusCode() < 200 {
		fmt.Fprintf(os.Stderr, "%+v\n", res.RawResponse)
		return nil, fmt.Errorf("scoreboard create returned status %v", res.StatusCode())
	}

	return &val, err
}

/*
   Read scoreboard

   Reads and returns scoreboard by scoreboard id
*/
func (endpoint *ScoreboardEndpoint) Read(
	scoreboardId string,
) (*ScoreboardReadResponse, error) {
	val := ScoreboardReadResponse{}
	url := "scoreboard/{scoreboard_id}"
	url = strings.Replace(url, "{scoreboard_id}", scoreboardId, -1)

	res, err := endpoint.Request().
		SetResult(&val).
		Get(endpoint.BaseUrl + url)

	if res.StatusCode() > 299 || res.StatusCode() < 200 {
		fmt.Fprintf(os.Stderr, "%+v\n", res.RawResponse)
		return nil, fmt.Errorf("scoreboard read returned status %v", res.StatusCode())
	}

	return &val, err
}

/*
   Update scoreboard

   Updates and returns scoreboard by scoreboard id
*/
func (endpoint *ScoreboardEndpoint) Update(
	scoreboardId string,
	body ScoreboardUpdateRequest,
) (*ScoreboardUpdateResponse, error) {
	val := ScoreboardUpdateResponse{}
	url := "scoreboard/{scoreboard_id}"
	url = strings.Replace(url, "{scoreboard_id}", scoreboardId, -1)

	res, err := endpoint.Request().
		SetResult(&val).
		SetBody(&body).
		Post(endpoint.BaseUrl + url)

	if res.StatusCode() > 299 || res.StatusCode() < 200 {
		fmt.Fprintf(os.Stderr, "%+v\n", res.RawResponse)
		return nil, fmt.Errorf("scoreboard update returned status %v", res.StatusCode())
	}

	return &val, err
}

/*
   Delete scoreboard by id.

   Delete scoreboard by scoreboard id
*/
func (endpoint *ScoreboardEndpoint) Delete(
	scoreboardId string,
) (*ScoreboardDeleteResponse, error) {
	val := ScoreboardDeleteResponse{}
	url := "scoreboard/{scoreboard_id}"
	url = strings.Replace(url, "{scoreboard_id}", scoreboardId, -1)

	res, err := endpoint.Request().
		SetResult(&val).
		Delete(endpoint.BaseUrl + url)

	if res.StatusCode() > 299 || res.StatusCode() < 200 {
		fmt.Fprintf(os.Stderr, "%+v\n", res.RawResponse)
		return nil, fmt.Errorf("scoreboard delete returned status %v", res.StatusCode())
	}

	return &val, err
}

