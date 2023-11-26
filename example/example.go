package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/vements/client-go/vements"
)

const (
	API_KEY       = ""
	PROJECT_ID    = ""
	SCOREBOARD_ID = ""
)

func main() {
	client, _ := vements.NewClient(API_KEY, []string{"production"}, vements.NewConfig())
	id := rand.Intn(1000)

	participant, _ := client.Participant.Create(
		vements.ParticipantCreateRequest{
			ProjectId:  PROJECT_ID,
			Display:    fmt.Sprintf("Example Player %v", id),
			ExternalId: fmt.Sprintf("example player %v", id),
			Image:      "",
			Extra:      nil,
		},
	)

	fmt.Printf("Participant Created: %+v\n", participant)

	for i := 0; i < 5; i++ {
		_, _ = client.Scoreboard.Record(
			SCOREBOARD_ID,
			vements.ScoreboardScoreRequest{
				ParticipantId: participant.ParticipantId,
				Value:         rand.Intn(1000) + 1,
				Recorded:      time.Now(),
			},
		)
	}

	scores, _ := client.Scoreboard.Scores(SCOREBOARD_ID, time.Now().Add(time.Hour*-24), time.Now())
	for _, score := range scores {
		fmt.Printf("Rank: %v Player: %v Total: %v\n", score.Rank, score.Participant.Display, score.Total)
	}
}
