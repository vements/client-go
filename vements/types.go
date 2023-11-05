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
	"strings"
	"time"
)

type Contact struct {
	Name  string `json:"name,omitempty"`
	Url   string `json:"url,omitempty"`
	Email string `json:"email,omitempty"`
}

func NewContact() *Contact {
	return &Contact{
		Name:  "Vements Support Contact",
		Url:   "https://vements.io",
		Email: "support@vements.io",
	}
}

type License struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

func NewLicense() *License {
	return &License{
		Name: "MIT",
		Url:  "https://opensource.org/license/mit/",
	}
}

type ExternalDocs struct {
	Description string `json:"description,omitempty"`
	Url         string `json:"url,omitempty"`
}

func NewExternalDocs() *ExternalDocs {
	return &ExternalDocs{
		Description: "Vements REST API Documentation",
		Url:         "https://vements.io/docs",
	}
}

type Server struct {
	Url         string            `json:"url,omitempty"`
	Description string            `json:"description,omitempty"`
	Variables   map[string]string `json:"variables,omitempty"`
	Tags        []string          `json:"tags,omitempty"`
}

type Config struct {
	Title          string        `json:"title,omitempty"`
	Version        string        `json:"version,omitempty"`
	Description    string        `json:"description,omitempty"`
	TermsOfService string        `json:"termsOfService,omitempty"`
	Contact        *Contact      `json:"contact,omitempty"`
	License        *License      `json:"license,omitempty"`
	ExternalDocs   *ExternalDocs `json:"externalDocs,omitempty"`
	Servers        []Server      `json:"servers,omitempty"`
}

func (config *Config) ServerUrl(tags []string) (string, bool) {
	for _, server := range config.Servers {
		matches := 0
		for _, match := range tags {
			for _, tag := range server.Tags {
				if tag == match {
					matches = matches + 1
				}
			}
		}
		if matches >= len(tags) {
			url := server.Url
			for key, value := range server.Variables {
				url = strings.Replace(url, "{"+key+"}", value, -1)
			}
			return url, true
		}
	}
	return "", false
}

func NewConfig() *Config {
	return &Config{
		Title:          "Vements REST API",
		Version:        "1.0.3",
		Description:    "This specification describes the Vements REST API, its endpoints, and  the data structures used to communicate with it.",
		TermsOfService: "https://vements.io/terms",
		Contact:        NewContact(),
		License:        NewLicense(),
		ExternalDocs:   NewExternalDocs(),
		Servers: []Server{
			{
				Url:         "https://a.vements.io/{basePath}",
				Description: "Production Server",
				Variables: map[string]string{
					"basePath": "api/rest/v1.0.3/",
				},
				Tags: []string{
					"production",
				},
			},
			{
				Url:         "http://api.localtest.me/{basePath}",
				Description: "Development Server (Host)",
				Variables: map[string]string{
					"basePath": "api/rest/v1.0.3/",
				},
				Tags: []string{
					"development",
					"host",
					"full",
				},
			},
			{
				Url:         "http://localhost:9000/{basePath}",
				Description: "Development Server (Host Substitute)",
				Variables: map[string]string{
					"basePath": "api/rest/v1.0.3/",
				},
				Tags: []string{
					"development",
					"host",
					"substitute",
				},
			},
			{
				Url:         "http://api-server-a:8080/{basePath}",
				Description: "Development Server (Container)",
				Variables: map[string]string{
					"basePath": "api/rest/v1.0.3/",
				},
				Tags: []string{
					"development",
					"container",
					"full",
				},
			},
			{
				Url:         "http://substitute-server:9000/{basePath}",
				Description: "Development Server (Container Substitute)",
				Variables: map[string]string{
					"basePath": "api/rest/v1.0.3/",
				},
				Tags: []string{
					"development",
					"container",
					"substitute",
				},
			},
		},
	}
}

type User struct {
	UserId  string    `json:"user_id"`
	Email   string    `json:"email"`
	Display string    `json:"display"`
	Db      string    `json:"db"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

type Project struct {
	ProjectId string    `json:"project_id"`
	UserId    string    `json:"user_id"`
	Display   string    `json:"display"`
	Created   time.Time `json:"created"`
	Updated   time.Time `json:"updated"`
	Extra     Extra     `json:"extra"`
}

type ApiKey struct {
	ApiKeyId    string    `json:"api_key_id"`
	ProjectId   string    `json:"project_id"`
	Display     string    `json:"display"`
	Capability  string    `json:"capability"`
	Deactivated time.Time `json:"deactivated"`
	LastUsed    time.Time `json:"last_used"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
}

type Achievement struct {
	AchievementId string    `json:"achievement_id"`
	ProjectId     string    `json:"project_id"`
	Display       string    `json:"display"`
	Goal          int       `json:"goal"`
	Repeats       int       `json:"repeats"`
	LockedImage   string    `json:"locked_image"`
	UnlockedImage string    `json:"unlocked_image"`
	Position      int       `json:"position"`
	Public        bool      `json:"public"`
	Created       time.Time `json:"created"`
	Updated       time.Time `json:"updated"`
	Extra         Extra     `json:"extra"`
}

type Participant struct {
	ParticipantId string    `json:"participant_id"`
	ProjectId     string    `json:"project_id"`
	Display       string    `json:"display"`
	ExternalId    string    `json:"external_id"`
	Image         string    `json:"image"`
	Created       time.Time `json:"created"`
	Updated       time.Time `json:"updated"`
	Extra         Extra     `json:"extra"`
}

type Scoreboard struct {
	ScoreboardId string    `json:"scoreboard_id"`
	ProjectId    string    `json:"project_id"`
	Display      string    `json:"display"`
	RankDir      string    `json:"rank_dir"`
	Public       bool      `json:"public"`
	Created      time.Time `json:"created"`
	Updated      time.Time `json:"updated"`
	Extra        Extra     `json:"extra"`
}

type Progress struct {
	ProgressId    int       `json:"progress_id"`
	AchievementId string    `json:"achievement_id"`
	ParticipantId string    `json:"participant_id"`
	Value         int       `json:"value"`
	Recorded      time.Time `json:"recorded"`
}

type Score struct {
	ScoreId       int       `json:"score_id"`
	ScoreboardId  string    `json:"scoreboard_id"`
	ParticipantId string    `json:"participant_id"`
	Value         int       `json:"value"`
	Recorded      time.Time `json:"recorded"`
}

type AchievementLeaderboardItem struct {
	Participant Participant `json:"participant"`
	Progress    int         `json:"progress"`
	Iterations  int         `json:"iterations"`
}

type ParticipantProgressItem struct {
	Achievement Achievement `json:"achievement"`
	Progress    int         `json:"progress"`
	Iterations  int         `json:"iterations"`
}

type ParticipantScoreItem struct {
	Scoreboard Scoreboard `json:"scoreboard"`
	Rank       int        `json:"rank"`
	Total      int        `json:"total"`
}

type ScoreboardScoreItem struct {
	ParticipantId string      `json:"participant_id"`
	Participant   Participant `json:"participant"`
	Rank          int         `json:"rank"`
	Total         int         `json:"total"`
}

type AchievementLeaderboardResponse struct {
	AchievementLeaderboards []AchievementLeaderboardItem `json:"achievement_leaderboard"`
}

type AchievementProgressResponse struct {
	InsertProgressOne Progress `json:"insert_progress_one"`
}

type ParticipantProgressResponse struct {
	ParticipantProgress []ParticipantProgressItem `json:"participant_progress"`
}

type ParticipantScoresResponse struct {
	ParticipantScores []ParticipantScoreItem `json:"participant_scores"`
}

type ScoreboardScoreResponse struct {
	InsertScoreOne Score `json:"insert_score_one"`
}

type ScoreboardScoresResponse struct {
	ScoreboardScores []ScoreboardScoreItem `json:"scoreboard_scores"`
}

type AchievementProgressRequest struct {
	ParticipantId string    `json:"participant_id"`
	Value         int       `json:"value"`
	Recorded      time.Time `json:"recorded"`
}

type ScoreboardScoreRequest struct {
	ParticipantId string    `json:"participant_id"`
	Value         int       `json:"value"`
	Recorded      time.Time `json:"recorded"`
}

type AchievementCreateRequest struct {
	ProjectId     string `json:"project_id"`
	Display       string `json:"display"`
	Goal          int    `json:"goal"`
	Repeats       int    `json:"repeats"`
	LockedImage   string `json:"locked_image"`
	UnlockedImage string `json:"unlocked_image"`
	Position      int    `json:"position"`
	Public        bool   `json:"public"`
	Extra         Extra  `json:"extra"`
}

type AchievementCreateResponse struct {
	InsertAchievementOne Achievement `json:"insert_achievement_one"`
}

type AchievementReadResponse struct {
	Achievements []Achievement `json:"achievement"`
}

type AchievementUpdateRequest struct {
	Display       string `json:"display"`
	Goal          int    `json:"goal"`
	Repeats       int    `json:"repeats"`
	LockedImage   string `json:"locked_image"`
	UnlockedImage string `json:"unlocked_image"`
	Position      int    `json:"position"`
	Public        bool   `json:"public"`
	Extra         Extra  `json:"extra"`
}

type AchievementUpdateResponse struct {
	UpdateAchievementByPk Achievement `json:"update_achievement_by_pk"`
}

type AchievementDeleteResponse struct {
	DeleteAchievementByPk Achievement `json:"delete_achievement_by_pk"`
}

type AchievementListResponse struct {
	Achievements []Achievement `json:"achievement"`
}

type ParticipantCreateRequest struct {
	ProjectId  string `json:"project_id"`
	Display    string `json:"display"`
	ExternalId string `json:"external_id"`
	Image      string `json:"image"`
	Extra      Extra  `json:"extra"`
}

type ParticipantCreateResponse struct {
	InsertParticipantOne Participant `json:"insert_participant_one"`
}

type ParticipantReadResponse struct {
	Participants []Participant `json:"participant"`
}

type ParticipantUpdateRequest struct {
	Display    string `json:"display"`
	ExternalId string `json:"external_id"`
	Image      string `json:"image"`
	Extra      Extra  `json:"extra"`
}

type ParticipantUpdateResponse struct {
	UpdateParticipantByPk Participant `json:"update_participant_by_pk"`
}

type ParticipantDeleteResponse struct {
	DeleteParticipantByPk Participant `json:"delete_participant_by_pk"`
}

type ParticipantListResponse struct {
	Participants []Participant `json:"participant"`
}

type ScoreboardCreateRequest struct {
	ProjectId string `json:"project_id"`
	Display   string `json:"display"`
	RankDir   string `json:"rank_dir"`
	Public    bool   `json:"public"`
	Extra     Extra  `json:"extra"`
}

type ScoreboardCreateResponse struct {
	InsertScoreboardOne Scoreboard `json:"insert_scoreboard_one"`
}

type ScoreboardReadResponse struct {
	Scoreboards []Scoreboard `json:"scoreboard"`
}

type ScoreboardUpdateRequest struct {
	Display string `json:"display"`
	RankDir string `json:"rank_dir"`
	Public  bool   `json:"public"`
	Extra   Extra  `json:"extra"`
}

type ScoreboardUpdateResponse struct {
	UpdateScoreboardByPk Scoreboard `json:"update_scoreboard_by_pk"`
}

type ScoreboardDeleteResponse struct {
	DeleteScoreboardByPk Scoreboard `json:"delete_scoreboard_by_pk"`
}

type ScoreboardListResponse struct {
	Scoreboards []Scoreboard `json:"scoreboard"`
}

