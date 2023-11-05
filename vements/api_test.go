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
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/go-resty/resty/v2"
)

type TestDb struct {
	User        []User        `json:"user"`
	Project     []Project     `json:"project"`
	ApiKey      []ApiKey      `json:"api_key"`
	Achievement []Achievement `json:"achievement"`
	Participant []Participant `json:"participant"`
	Scoreboard  []Scoreboard  `json:"scoreboard"`
	Progress    []Progress    `json:"progress"`
	Score       []Score       `json:"score"`
}

type TestSetup struct {
	ApiKey ApiKey
	Db     TestDb
	Client *Client
}

func NewTestSetup() *TestSetup {
	tags := []string{"development", "host", "substitute"}
	config := NewConfig()
	client, ok := NewClient("no api key needed", tags, config)
	if !ok {
		panic("could not create client")
	}
	db := TestDb{}
	rc := resty.New()
	_, err := rc.R().SetResult(&db).Post(client.BaseUrl + "-/database")
	if err != nil {
		panic(fmt.Sprintf("could not create database: %v", err))
	}

	var api_key ApiKey
	for _, k := range db.ApiKey {
		if k.Capability == "rw" {
			api_key = k
			break
		}
	}
	if api_key == (ApiKey{}) {
		panic("could not find api key")
	}

	var api_key_v = fmt.Sprintf("%s:%s", api_key.ProjectId, api_key.ApiKeyId)
	client, ok = NewClient(api_key_v, tags, config)
	if !ok {
		panic("could not create client with api key")
	}

	return &TestSetup{
		ApiKey: api_key,
		Db:     db,
		Client: client,
	}
}

func (setup *TestSetup) GetString(key string) string {

	switch key {
	case "project_id":
		return setup.ApiKey.ProjectId
	default:
		break
	}

	if strings.HasSuffix(key, "_id") {
		switch key {

		case "achievement_id":
			var items []Achievement
			for _, item := range setup.Db.Achievement {
				if item.ProjectId == setup.GetString("project_id") {
					items = append(items, item)
				}
			}
			if len(items) == 0 {
				break
			}
			return items[rand.Intn(len(items))].AchievementId

		case "participant_id":
			var items []Participant
			for _, item := range setup.Db.Participant {
				if item.ProjectId == setup.GetString("project_id") {
					items = append(items, item)
				}
			}
			if len(items) == 0 {
				break
			}
			return items[rand.Intn(len(items))].ParticipantId

		case "scoreboard_id":
			var items []Scoreboard
			for _, item := range setup.Db.Scoreboard {
				if item.ProjectId == setup.GetString("project_id") {
					items = append(items, item)
				}
			}
			if len(items) == 0 {
				break
			}
			return items[rand.Intn(len(items))].ScoreboardId
		default:
			break
		}
	}
	return ""
}

func (setup *TestSetup) GetInt(key string) int {
	switch key {
	case "limit":
		return 100
	case "offset":
		return 0
	case "value":
		return 1
	default:
		return 0
	}
}

func (setup *TestSetup) GetTime(key string) time.Time {
	switch key {
	default:
		return time.Time{}
	}
}

func (setup *TestSetup) GetRandomString(key string) string {
	switch key {
	case "display":
		return faker.Name()
	case "rank_dir":
		return []string{"asc", "desc"}[rand.Intn(2)]
	case "external_id":
		return faker.UUIDDigit()
	default:
		return ""
	}
}

func (setup *TestSetup) GetBool(key string) bool {
	return []bool{true, false}[rand.Intn(2)]
}

func (setup *TestSetup) GetExtra(key string) Extra {
	return Extra{}
}

func TestAchievementLeaderboard(t *testing.T) {
	setup := NewTestSetup()
	_, err := setup.Client.Achievement.Leaderboard(
		setup.GetString("achievement_id"),
	)
	if err != nil {
		t.Fatalf("AchievementLeaderboard err: %v", err)
	}
}

func TestAchievementRecord(t *testing.T) {
	setup := NewTestSetup()
	_, err := setup.Client.Achievement.Record(
		setup.GetString("achievement_id"),
		AchievementProgressRequest{
			ParticipantId: setup.GetString("participant_id"),
			Value:         setup.GetInt("value"),
			Recorded:      setup.GetTime("recorded"),
		},
	)
	if err != nil {
		t.Fatalf("AchievementRecord err: %v", err)
	}
}

func TestParticipantProgress(t *testing.T) {
	setup := NewTestSetup()
	_, err := setup.Client.Participant.Progress(
		setup.GetString("participant_id"),
	)
	if err != nil {
		t.Fatalf("ParticipantProgress err: %v", err)
	}
}

func TestParticipantScores(t *testing.T) {
	setup := NewTestSetup()
	_, err := setup.Client.Participant.Scores(
		setup.GetString("participant_id"),
	)
	if err != nil {
		t.Fatalf("ParticipantScores err: %v", err)
	}
}

func TestScoreboardRecord(t *testing.T) {
	setup := NewTestSetup()
	_, err := setup.Client.Scoreboard.Record(
		setup.GetString("scoreboard_id"),
		ScoreboardScoreRequest{
			ParticipantId: setup.GetString("participant_id"),
			Value:         setup.GetInt("value"),
			Recorded:      setup.GetTime("recorded"),
		},
	)
	if err != nil {
		t.Fatalf("ScoreboardRecord err: %v", err)
	}
}

func TestScoreboardScores(t *testing.T) {
	setup := NewTestSetup()
	_, err := setup.Client.Scoreboard.Scores(
		setup.GetString("scoreboard_id"),
		setup.GetTime("from"),
		setup.GetTime("to"),
	)
	if err != nil {
		t.Fatalf("ScoreboardScores err: %v", err)
	}
}

func TestAchievementList(t *testing.T) {
	setup := NewTestSetup()
	_, err := setup.Client.Achievement.List(
		setup.GetString("project_id"),
		setup.GetInt("limit"),
		setup.GetInt("offset"),
	)
	if err != nil {
		t.Fatalf("AchievementList err: %v", err)
	}
}

func TestAchievementCreate(t *testing.T) {
	setup := NewTestSetup()
	_, err := setup.Client.Achievement.Create(
		AchievementCreateRequest{
			ProjectId:     setup.GetString("project_id"),
			Display:       setup.GetString("display"),
			Goal:          setup.GetInt("goal"),
			Repeats:       setup.GetInt("repeats"),
			LockedImage:   setup.GetString("locked_image"),
			UnlockedImage: setup.GetString("unlocked_image"),
			Position:      setup.GetInt("position"),
			Public:        setup.GetBool("public"),
			Extra:         setup.GetExtra("extra"),
		},
	)
	if err != nil {
		t.Fatalf("AchievementCreate err: %v", err)
	}
}

func TestAchievementRead(t *testing.T) {
	setup := NewTestSetup()
	_, err := setup.Client.Achievement.Read(
		setup.GetString("achievement_id"),
	)
	if err != nil {
		t.Fatalf("AchievementRead err: %v", err)
	}
}

func TestAchievementUpdate(t *testing.T) {
	setup := NewTestSetup()
	_, err := setup.Client.Achievement.Update(
		setup.GetString("achievement_id"),
		AchievementUpdateRequest{
			Display:       setup.GetString("display"),
			Goal:          setup.GetInt("goal"),
			Repeats:       setup.GetInt("repeats"),
			LockedImage:   setup.GetString("locked_image"),
			UnlockedImage: setup.GetString("unlocked_image"),
			Position:      setup.GetInt("position"),
			Public:        setup.GetBool("public"),
			Extra:         setup.GetExtra("extra"),
		},
	)
	if err != nil {
		t.Fatalf("AchievementUpdate err: %v", err)
	}
}

func TestAchievementDelete(t *testing.T) {
	setup := NewTestSetup()
	_, err := setup.Client.Achievement.Delete(
		setup.GetString("achievement_id"),
	)
	if err != nil {
		t.Fatalf("AchievementDelete err: %v", err)
	}
}

func TestParticipantList(t *testing.T) {
	setup := NewTestSetup()
	_, err := setup.Client.Participant.List(
		setup.GetString("project_id"),
		setup.GetInt("limit"),
		setup.GetInt("offset"),
	)
	if err != nil {
		t.Fatalf("ParticipantList err: %v", err)
	}
}

func TestParticipantCreate(t *testing.T) {
	setup := NewTestSetup()
	_, err := setup.Client.Participant.Create(
		ParticipantCreateRequest{
			ProjectId:  setup.GetString("project_id"),
			Display:    setup.GetString("display"),
			ExternalId: setup.GetString("external_id"),
			Image:      setup.GetString("image"),
			Extra:      setup.GetExtra("extra"),
		},
	)
	if err != nil {
		t.Fatalf("ParticipantCreate err: %v", err)
	}
}

func TestParticipantRead(t *testing.T) {
	setup := NewTestSetup()
	_, err := setup.Client.Participant.Read(
		setup.GetString("participant_id"),
	)
	if err != nil {
		t.Fatalf("ParticipantRead err: %v", err)
	}
}

func TestParticipantUpdate(t *testing.T) {
	setup := NewTestSetup()
	_, err := setup.Client.Participant.Update(
		setup.GetString("participant_id"),
		ParticipantUpdateRequest{
			Display:    setup.GetString("display"),
			ExternalId: setup.GetString("external_id"),
			Image:      setup.GetString("image"),
			Extra:      setup.GetExtra("extra"),
		},
	)
	if err != nil {
		t.Fatalf("ParticipantUpdate err: %v", err)
	}
}

func TestParticipantDelete(t *testing.T) {
	setup := NewTestSetup()
	_, err := setup.Client.Participant.Delete(
		setup.GetString("participant_id"),
	)
	if err != nil {
		t.Fatalf("ParticipantDelete err: %v", err)
	}
}

func TestScoreboardList(t *testing.T) {
	setup := NewTestSetup()
	_, err := setup.Client.Scoreboard.List(
		setup.GetString("project_id"),
		setup.GetInt("limit"),
		setup.GetInt("offset"),
	)
	if err != nil {
		t.Fatalf("ScoreboardList err: %v", err)
	}
}

func TestScoreboardCreate(t *testing.T) {
	setup := NewTestSetup()
	_, err := setup.Client.Scoreboard.Create(
		ScoreboardCreateRequest{
			ProjectId: setup.GetString("project_id"),
			Display:   setup.GetString("display"),
			RankDir:   setup.GetString("rank_dir"),
			Public:    setup.GetBool("public"),
			Extra:     setup.GetExtra("extra"),
		},
	)
	if err != nil {
		t.Fatalf("ScoreboardCreate err: %v", err)
	}
}

func TestScoreboardRead(t *testing.T) {
	setup := NewTestSetup()
	_, err := setup.Client.Scoreboard.Read(
		setup.GetString("scoreboard_id"),
	)
	if err != nil {
		t.Fatalf("ScoreboardRead err: %v", err)
	}
}

func TestScoreboardUpdate(t *testing.T) {
	setup := NewTestSetup()
	_, err := setup.Client.Scoreboard.Update(
		setup.GetString("scoreboard_id"),
		ScoreboardUpdateRequest{
			Display: setup.GetString("display"),
			RankDir: setup.GetString("rank_dir"),
			Public:  setup.GetBool("public"),
			Extra:   setup.GetExtra("extra"),
		},
	)
	if err != nil {
		t.Fatalf("ScoreboardUpdate err: %v", err)
	}
}

func TestScoreboardDelete(t *testing.T) {
	setup := NewTestSetup()
	_, err := setup.Client.Scoreboard.Delete(
		setup.GetString("scoreboard_id"),
	)
	if err != nil {
		t.Fatalf("ScoreboardDelete err: %v", err)
	}
}

