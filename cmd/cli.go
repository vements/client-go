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

package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/vements/vements-go/v1/vements"
)

func init() {
	var rootCommandState RootCommandState = RootCommandState{apiKey: "", tags: []string{"production"}}

	root.PersistentFlags().StringVar(&rootCommandState.apiKey, "api-key", "", "API Key; overrides $API_KEY if set")
	if os.Getenv("API_KEY") != "" {
		rootCommandState.apiKey = os.Getenv("API_KEY")
	}
	if os.Getenv("SERVER_TAGS") != "" {
		rootCommandState.tags = strings.Split(os.Getenv("SERVER_TAGS"), ",")
	}
	root.SetContext(context.WithValue(context.Background(), StateKey, &rootCommandState))
	root.AddCommand(&cobra.Command{
		Use:   "api-version",
		Short: "Print the API version",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%+v\n", root.Client().Config.Version)
		},
	})
	root.AddCommand(&cobra.Command{
		Use:   "client-version",
		Short: "Print the client library version",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%+v\n", "0.0.1")
		},
	})

	// achievement leaderboard
	achievementLeaderboardCommandState := AchievementLeaderboardCommandState{}
	achievement.leaderboard.SetContext(context.WithValue(context.Background(), AchievementLeaderboardCommandKey, &achievementLeaderboardCommandState))
	achievement.leaderboard.Command.Flags().StringVar(&achievementLeaderboardCommandState.AchievementId, "achievement-id", "", "Achievement ID")
	achievement.leaderboard.Command.MarkFlagRequired("achievement-id")
	achievement.parent.AddChild(achievement.leaderboard) // add sub commands to parent command

	// achievement record
	achievementRecordCommandState := AchievementRecordCommandState{}
	achievement.record.SetContext(context.WithValue(context.Background(), AchievementRecordCommandKey, &achievementRecordCommandState))
	achievement.record.Command.Flags().StringVar(&achievementRecordCommandState.AchievementId, "achievement-id", "", "Achievement ID")
	achievement.record.Command.MarkFlagRequired("achievement-id")
	achievement.record.Command.Flags().StringVar(&achievementRecordCommandState.ParticipantId, "participant-id", "", "Participant ID")
	achievement.record.Command.MarkFlagRequired("participant-id")
	achievement.record.Command.Flags().IntVar(&achievementRecordCommandState.Value, "value", 0, "Progress value")
	achievement.record.Command.MarkFlagRequired("value")
	achievement.record.Command.Flags().StringVar(&achievementRecordCommandState.Recorded, "recorded", "", "Date and time progress recorded")
	achievement.parent.AddChild(achievement.record) // add sub commands to parent command

	// achievement list
	achievementListCommandState := AchievementListCommandState{}
	achievement.list.SetContext(context.WithValue(context.Background(), AchievementListCommandKey, &achievementListCommandState))
	achievement.list.Command.Flags().StringVar(&achievementListCommandState.ProjectId, "project-id", "", "Project ID")
	achievement.list.Command.MarkFlagRequired("project-id")
	achievement.list.Command.Flags().IntVar(&achievementListCommandState.Limit, "limit", 100, "List limit")
	achievement.list.Command.Flags().IntVar(&achievementListCommandState.Offset, "offset", 0, "List offset")
	achievement.parent.AddChild(achievement.list) // add sub commands to parent command

	// achievement create
	achievementCreateCommandState := AchievementCreateCommandState{}
	achievement.create.SetContext(context.WithValue(context.Background(), AchievementCreateCommandKey, &achievementCreateCommandState))
	achievement.create.Command.Flags().StringVar(&achievementCreateCommandState.ProjectId, "project-id", "", "Related project ID")
	achievement.create.Command.MarkFlagRequired("project-id")
	achievement.create.Command.Flags().StringVar(&achievementCreateCommandState.Display, "display", "", "Achievement display name")
	achievement.create.Command.MarkFlagRequired("display")
	achievement.create.Command.Flags().IntVar(&achievementCreateCommandState.Goal, "goal", 0, "Achievement goal")
	achievement.create.Command.MarkFlagRequired("goal")
	achievement.create.Command.Flags().IntVar(&achievementCreateCommandState.Repeats, "repeats", 0, "Achievement repeats; zero for unlimited")
	achievement.create.Command.MarkFlagRequired("repeats")
	achievement.create.Command.Flags().StringVar(&achievementCreateCommandState.LockedImage, "locked-image", "", "Achievement locked or non-obtained image")
	achievement.create.Command.Flags().StringVar(&achievementCreateCommandState.UnlockedImage, "unlocked-image", "", "Achievement unlocked or obtained image")
	achievement.create.Command.Flags().IntVar(&achievementCreateCommandState.Position, "position", 0, "Achievement position within project")
	achievement.create.Command.MarkFlagRequired("position")
	achievement.create.Command.Flags().BoolVar(&achievementCreateCommandState.Public, "public", false, "Achievement public flag")
	achievement.create.Command.MarkFlagRequired("public")
	achievement.create.Command.Flags().StringVar(&achievementCreateCommandState.Extra, "extra", "", "Arbitrary application data")
	achievement.parent.AddChild(achievement.create) // add sub commands to parent command

	// achievement read
	achievementReadCommandState := AchievementReadCommandState{}
	achievement.read.SetContext(context.WithValue(context.Background(), AchievementReadCommandKey, &achievementReadCommandState))
	achievement.read.Command.Flags().StringVar(&achievementReadCommandState.AchievementId, "achievement-id", "", "Achievement ID")
	achievement.read.Command.MarkFlagRequired("achievement-id")
	achievement.parent.AddChild(achievement.read) // add sub commands to parent command

	// achievement update
	achievementUpdateCommandState := AchievementUpdateCommandState{}
	achievement.update.SetContext(context.WithValue(context.Background(), AchievementUpdateCommandKey, &achievementUpdateCommandState))
	achievement.update.Command.Flags().StringVar(&achievementUpdateCommandState.AchievementId, "achievement-id", "", "Achievement ID")
	achievement.update.Command.MarkFlagRequired("achievement-id")
	achievement.update.Command.Flags().StringVar(&achievementUpdateCommandState.Display, "display", "", "Achievement display name")
	achievement.update.Command.MarkFlagRequired("display")
	achievement.update.Command.Flags().IntVar(&achievementUpdateCommandState.Goal, "goal", 0, "Achievement goal")
	achievement.update.Command.MarkFlagRequired("goal")
	achievement.update.Command.Flags().IntVar(&achievementUpdateCommandState.Repeats, "repeats", 0, "Achievement repeats; zero for unlimited")
	achievement.update.Command.MarkFlagRequired("repeats")
	achievement.update.Command.Flags().StringVar(&achievementUpdateCommandState.LockedImage, "locked-image", "", "Achievement locked or non-obtained image")
	achievement.update.Command.Flags().StringVar(&achievementUpdateCommandState.UnlockedImage, "unlocked-image", "", "Achievement unlocked or obtained image")
	achievement.update.Command.Flags().IntVar(&achievementUpdateCommandState.Position, "position", 0, "Achievement position within project")
	achievement.update.Command.MarkFlagRequired("position")
	achievement.update.Command.Flags().BoolVar(&achievementUpdateCommandState.Public, "public", false, "Achievement public flag")
	achievement.update.Command.MarkFlagRequired("public")
	achievement.update.Command.Flags().StringVar(&achievementUpdateCommandState.Extra, "extra", "", "Arbitrary application data")
	achievement.parent.AddChild(achievement.update) // add sub commands to parent command

	// achievement delete
	achievementDeleteCommandState := AchievementDeleteCommandState{}
	achievement.delete.SetContext(context.WithValue(context.Background(), AchievementDeleteCommandKey, &achievementDeleteCommandState))
	achievement.delete.Command.Flags().StringVar(&achievementDeleteCommandState.AchievementId, "achievement-id", "", "Achievement ID")
	achievement.delete.Command.MarkFlagRequired("achievement-id")
	achievement.parent.AddChild(achievement.delete) // add sub commands to parent command

	// participant progress
	participantProgressCommandState := ParticipantProgressCommandState{}
	participant.progress.SetContext(context.WithValue(context.Background(), ParticipantProgressCommandKey, &participantProgressCommandState))
	participant.progress.Command.Flags().StringVar(&participantProgressCommandState.ParticipantId, "participant-id", "", "Participant ID")
	participant.progress.Command.MarkFlagRequired("participant-id")
	participant.parent.AddChild(participant.progress) // add sub commands to parent command

	// participant scores
	participantScoresCommandState := ParticipantScoresCommandState{}
	participant.scores.SetContext(context.WithValue(context.Background(), ParticipantScoresCommandKey, &participantScoresCommandState))
	participant.scores.Command.Flags().StringVar(&participantScoresCommandState.ParticipantId, "participant-id", "", "Participant ID")
	participant.scores.Command.MarkFlagRequired("participant-id")
	participant.parent.AddChild(participant.scores) // add sub commands to parent command

	// participant list
	participantListCommandState := ParticipantListCommandState{}
	participant.list.SetContext(context.WithValue(context.Background(), ParticipantListCommandKey, &participantListCommandState))
	participant.list.Command.Flags().StringVar(&participantListCommandState.ProjectId, "project-id", "", "Project ID")
	participant.list.Command.MarkFlagRequired("project-id")
	participant.list.Command.Flags().IntVar(&participantListCommandState.Limit, "limit", 100, "List limit")
	participant.list.Command.Flags().IntVar(&participantListCommandState.Offset, "offset", 0, "List offset")
	participant.parent.AddChild(participant.list) // add sub commands to parent command

	// participant create
	participantCreateCommandState := ParticipantCreateCommandState{}
	participant.create.SetContext(context.WithValue(context.Background(), ParticipantCreateCommandKey, &participantCreateCommandState))
	participant.create.Command.Flags().StringVar(&participantCreateCommandState.ProjectId, "project-id", "", "Related project ID")
	participant.create.Command.MarkFlagRequired("project-id")
	participant.create.Command.Flags().StringVar(&participantCreateCommandState.Display, "display", "", "Participant display name")
	participant.create.Command.MarkFlagRequired("display")
	participant.create.Command.Flags().StringVar(&participantCreateCommandState.ExternalId, "external-id", "", "Participant external ID")
	participant.create.Command.MarkFlagRequired("external-id")
	participant.create.Command.Flags().StringVar(&participantCreateCommandState.Image, "image", "", "Participant image")
	participant.create.Command.Flags().StringVar(&participantCreateCommandState.Extra, "extra", "", "Arbitrary application data")
	participant.parent.AddChild(participant.create) // add sub commands to parent command

	// participant read
	participantReadCommandState := ParticipantReadCommandState{}
	participant.read.SetContext(context.WithValue(context.Background(), ParticipantReadCommandKey, &participantReadCommandState))
	participant.read.Command.Flags().StringVar(&participantReadCommandState.ParticipantId, "participant-id", "", "Participant ID")
	participant.read.Command.MarkFlagRequired("participant-id")
	participant.parent.AddChild(participant.read) // add sub commands to parent command

	// participant update
	participantUpdateCommandState := ParticipantUpdateCommandState{}
	participant.update.SetContext(context.WithValue(context.Background(), ParticipantUpdateCommandKey, &participantUpdateCommandState))
	participant.update.Command.Flags().StringVar(&participantUpdateCommandState.ParticipantId, "participant-id", "", "Participant ID")
	participant.update.Command.MarkFlagRequired("participant-id")
	participant.update.Command.Flags().StringVar(&participantUpdateCommandState.Display, "display", "", "Participant display name")
	participant.update.Command.MarkFlagRequired("display")
	participant.update.Command.Flags().StringVar(&participantUpdateCommandState.ExternalId, "external-id", "", "Participant external ID")
	participant.update.Command.MarkFlagRequired("external-id")
	participant.update.Command.Flags().StringVar(&participantUpdateCommandState.Image, "image", "", "Participant image")
	participant.update.Command.Flags().StringVar(&participantUpdateCommandState.Extra, "extra", "", "Arbitrary application data")
	participant.parent.AddChild(participant.update) // add sub commands to parent command

	// participant delete
	participantDeleteCommandState := ParticipantDeleteCommandState{}
	participant.delete.SetContext(context.WithValue(context.Background(), ParticipantDeleteCommandKey, &participantDeleteCommandState))
	participant.delete.Command.Flags().StringVar(&participantDeleteCommandState.ParticipantId, "participant-id", "", "Participant ID")
	participant.delete.Command.MarkFlagRequired("participant-id")
	participant.parent.AddChild(participant.delete) // add sub commands to parent command

	// scoreboard record
	scoreboardRecordCommandState := ScoreboardRecordCommandState{}
	scoreboard.record.SetContext(context.WithValue(context.Background(), ScoreboardRecordCommandKey, &scoreboardRecordCommandState))
	scoreboard.record.Command.Flags().StringVar(&scoreboardRecordCommandState.ScoreboardId, "scoreboard-id", "", "Scoreboard ID")
	scoreboard.record.Command.MarkFlagRequired("scoreboard-id")
	scoreboard.record.Command.Flags().StringVar(&scoreboardRecordCommandState.ParticipantId, "participant-id", "", "Participant ID")
	scoreboard.record.Command.MarkFlagRequired("participant-id")
	scoreboard.record.Command.Flags().IntVar(&scoreboardRecordCommandState.Value, "value", 0, "Score value")
	scoreboard.record.Command.MarkFlagRequired("value")
	scoreboard.record.Command.Flags().StringVar(&scoreboardRecordCommandState.Recorded, "recorded", "", "Date and time score recorded")
	scoreboard.parent.AddChild(scoreboard.record) // add sub commands to parent command

	// scoreboard scores
	scoreboardScoresCommandState := ScoreboardScoresCommandState{}
	scoreboard.scores.SetContext(context.WithValue(context.Background(), ScoreboardScoresCommandKey, &scoreboardScoresCommandState))
	scoreboard.scores.Command.Flags().StringVar(&scoreboardScoresCommandState.ScoreboardId, "scoreboard-id", "", "Scoreboard ID")
	scoreboard.scores.Command.MarkFlagRequired("scoreboard-id")
	scoreboard.scores.Command.Flags().StringVar(&scoreboardScoresCommandState.From, "from", "", "From")
	scoreboard.scores.Command.Flags().StringVar(&scoreboardScoresCommandState.To, "to", "", "To")
	scoreboard.parent.AddChild(scoreboard.scores) // add sub commands to parent command

	// scoreboard list
	scoreboardListCommandState := ScoreboardListCommandState{}
	scoreboard.list.SetContext(context.WithValue(context.Background(), ScoreboardListCommandKey, &scoreboardListCommandState))
	scoreboard.list.Command.Flags().StringVar(&scoreboardListCommandState.ProjectId, "project-id", "", "Project ID")
	scoreboard.list.Command.MarkFlagRequired("project-id")
	scoreboard.list.Command.Flags().IntVar(&scoreboardListCommandState.Limit, "limit", 100, "List limit")
	scoreboard.list.Command.Flags().IntVar(&scoreboardListCommandState.Offset, "offset", 0, "List offset")
	scoreboard.parent.AddChild(scoreboard.list) // add sub commands to parent command

	// scoreboard create
	scoreboardCreateCommandState := ScoreboardCreateCommandState{}
	scoreboard.create.SetContext(context.WithValue(context.Background(), ScoreboardCreateCommandKey, &scoreboardCreateCommandState))
	scoreboard.create.Command.Flags().StringVar(&scoreboardCreateCommandState.ProjectId, "project-id", "", "Related project ID")
	scoreboard.create.Command.MarkFlagRequired("project-id")
	scoreboard.create.Command.Flags().StringVar(&scoreboardCreateCommandState.Display, "display", "", "Scoreboard display name")
	scoreboard.create.Command.MarkFlagRequired("display")
	scoreboard.create.Command.Flags().StringVar(&scoreboardCreateCommandState.RankDir, "rank-dir", "", "Scoreboard rank direction")
	scoreboard.create.Command.MarkFlagRequired("rank-dir")
	scoreboard.create.Command.Flags().BoolVar(&scoreboardCreateCommandState.Public, "public", false, "Scoreboard public flag")
	scoreboard.create.Command.MarkFlagRequired("public")
	scoreboard.create.Command.Flags().StringVar(&scoreboardCreateCommandState.Extra, "extra", "", "Arbitrary application data")
	scoreboard.parent.AddChild(scoreboard.create) // add sub commands to parent command

	// scoreboard read
	scoreboardReadCommandState := ScoreboardReadCommandState{}
	scoreboard.read.SetContext(context.WithValue(context.Background(), ScoreboardReadCommandKey, &scoreboardReadCommandState))
	scoreboard.read.Command.Flags().StringVar(&scoreboardReadCommandState.ScoreboardId, "scoreboard-id", "", "Scoreboard ID")
	scoreboard.read.Command.MarkFlagRequired("scoreboard-id")
	scoreboard.parent.AddChild(scoreboard.read) // add sub commands to parent command

	// scoreboard update
	scoreboardUpdateCommandState := ScoreboardUpdateCommandState{}
	scoreboard.update.SetContext(context.WithValue(context.Background(), ScoreboardUpdateCommandKey, &scoreboardUpdateCommandState))
	scoreboard.update.Command.Flags().StringVar(&scoreboardUpdateCommandState.ScoreboardId, "scoreboard-id", "", "Scoreboard ID")
	scoreboard.update.Command.MarkFlagRequired("scoreboard-id")
	scoreboard.update.Command.Flags().StringVar(&scoreboardUpdateCommandState.Display, "display", "", "Scoreboard display name")
	scoreboard.update.Command.MarkFlagRequired("display")
	scoreboard.update.Command.Flags().StringVar(&scoreboardUpdateCommandState.RankDir, "rank-dir", "", "Scoreboard rank direction")
	scoreboard.update.Command.MarkFlagRequired("rank-dir")
	scoreboard.update.Command.Flags().BoolVar(&scoreboardUpdateCommandState.Public, "public", false, "Scoreboard public flag")
	scoreboard.update.Command.MarkFlagRequired("public")
	scoreboard.update.Command.Flags().StringVar(&scoreboardUpdateCommandState.Extra, "extra", "", "Arbitrary application data")
	scoreboard.parent.AddChild(scoreboard.update) // add sub commands to parent command

	// scoreboard delete
	scoreboardDeleteCommandState := ScoreboardDeleteCommandState{}
	scoreboard.delete.SetContext(context.WithValue(context.Background(), ScoreboardDeleteCommandKey, &scoreboardDeleteCommandState))
	scoreboard.delete.Command.Flags().StringVar(&scoreboardDeleteCommandState.ScoreboardId, "scoreboard-id", "", "Scoreboard ID")
	scoreboard.delete.Command.MarkFlagRequired("scoreboard-id")
	scoreboard.parent.AddChild(scoreboard.delete) // add sub commands to parent command

	// add resource commands to root command
	root.AddChild(achievement.parent)
	root.AddChild(participant.parent)
	root.AddChild(scoreboard.parent)
}

var achievement = struct {
	parent      *Command
	leaderboard *Command
	record      *Command
	list        *Command
	create      *Command
	read        *Command
	update      *Command
	delete      *Command
}{
	parent: &Command{
		Command: &cobra.Command{
			Use:   "achievement",
			Short: "Create, read, update, delete, and list achievements",
			Long:  "",
			Run: func(cmd *cobra.Command, args []string) {
				cmd.Usage()
			},
		},
	},
	leaderboard: &Command{
		Command: &cobra.Command{
			Use:   "leaderboard",
			Short: "Achievement leaderboard",
			Long:  "Reads and returns achievement leaderboard",
			Run: func(cmd *cobra.Command, args []string) {
				if state, ok := cmd.Context().Value(AchievementLeaderboardCommandKey).(*AchievementLeaderboardCommandState); ok {
					client := root.Client()
					res, err := client.Achievement.Leaderboard(
						state.AchievementId,
					)
					root.WriteResult(res, err)
				} else {
					panic("invalid context for achievement leaderboard")
				}
			},
		},
	},
	record: &Command{
		Command: &cobra.Command{
			Use:   "record",
			Short: "Record achievement progress",
			Long:  "Records and returns achievement progress",
			Run: func(cmd *cobra.Command, args []string) {
				if state, ok := cmd.Context().Value(AchievementRecordCommandKey).(*AchievementRecordCommandState); ok {
					client := root.Client()
					res, err := client.Achievement.Record(
						state.AchievementId,
						vements.AchievementProgressRequest{
							ParticipantId: state.ParticipantId,
							Value:         state.Value,
							Recorded:      MapTime(state.Recorded),
						},
					)
					root.WriteResult(res, err)
				} else {
					panic("invalid context for achievement record")
				}
			},
		},
	},
	list: &Command{
		Command: &cobra.Command{
			Use:   "list",
			Short: "List achievements",
			Long:  "Reads and returns list of achievements in the given project",
			Run: func(cmd *cobra.Command, args []string) {
				if state, ok := cmd.Context().Value(AchievementListCommandKey).(*AchievementListCommandState); ok {
					client := root.Client()
					res, err := client.Achievement.List(
						state.ProjectId,
						state.Limit,
						state.Offset,
					)
					root.WriteResult(res, err)
				} else {
					panic("invalid context for achievement list")
				}
			},
		},
	},
	create: &Command{
		Command: &cobra.Command{
			Use:   "create",
			Short: "Create achievement",
			Long:  "Creates and returns achievement in the given project",
			Run: func(cmd *cobra.Command, args []string) {
				if state, ok := cmd.Context().Value(AchievementCreateCommandKey).(*AchievementCreateCommandState); ok {
					client := root.Client()
					res, err := client.Achievement.Create(
						vements.AchievementCreateRequest{
							ProjectId:     state.ProjectId,
							Display:       state.Display,
							Goal:          state.Goal,
							Repeats:       state.Repeats,
							LockedImage:   state.LockedImage,
							UnlockedImage: state.UnlockedImage,
							Position:      state.Position,
							Public:        state.Public,
							Extra:         MapExtra(state.Extra),
						},
					)
					root.WriteResult(res, err)
				} else {
					panic("invalid context for achievement create")
				}
			},
		},
	},
	read: &Command{
		Command: &cobra.Command{
			Use:   "read",
			Short: "Read achievement",
			Long:  "Reads and returns achievement by achievement id",
			Run: func(cmd *cobra.Command, args []string) {
				if state, ok := cmd.Context().Value(AchievementReadCommandKey).(*AchievementReadCommandState); ok {
					client := root.Client()
					res, err := client.Achievement.Read(
						state.AchievementId,
					)
					root.WriteResult(res, err)
				} else {
					panic("invalid context for achievement read")
				}
			},
		},
	},
	update: &Command{
		Command: &cobra.Command{
			Use:   "update",
			Short: "Update achievement",
			Long:  "Updates and returns achievement by achievement id",
			Run: func(cmd *cobra.Command, args []string) {
				if state, ok := cmd.Context().Value(AchievementUpdateCommandKey).(*AchievementUpdateCommandState); ok {
					client := root.Client()
					res, err := client.Achievement.Update(
						state.AchievementId,
						vements.AchievementUpdateRequest{
							Display:       state.Display,
							Goal:          state.Goal,
							Repeats:       state.Repeats,
							LockedImage:   state.LockedImage,
							UnlockedImage: state.UnlockedImage,
							Position:      state.Position,
							Public:        state.Public,
							Extra:         MapExtra(state.Extra),
						},
					)
					root.WriteResult(res, err)
				} else {
					panic("invalid context for achievement update")
				}
			},
		},
	},
	delete: &Command{
		Command: &cobra.Command{
			Use:   "delete",
			Short: "Delete achievement by id.",
			Long:  "Delete achievement by achievement id",
			Run: func(cmd *cobra.Command, args []string) {
				if state, ok := cmd.Context().Value(AchievementDeleteCommandKey).(*AchievementDeleteCommandState); ok {
					client := root.Client()
					res, err := client.Achievement.Delete(
						state.AchievementId,
					)
					root.WriteResult(res, err)
				} else {
					panic("invalid context for achievement delete")
				}
			},
		},
	},
}

var participant = struct {
	parent   *Command
	progress *Command
	scores   *Command
	list     *Command
	create   *Command
	read     *Command
	update   *Command
	delete   *Command
}{
	parent: &Command{
		Command: &cobra.Command{
			Use:   "participant",
			Short: "Create, read, update, delete, and list participants",
			Long:  "",
			Run: func(cmd *cobra.Command, args []string) {
				cmd.Usage()
			},
		},
	},
	progress: &Command{
		Command: &cobra.Command{
			Use:   "progress",
			Short: "Participant progress",
			Long:  "Reads and returns participant progress.",
			Run: func(cmd *cobra.Command, args []string) {
				if state, ok := cmd.Context().Value(ParticipantProgressCommandKey).(*ParticipantProgressCommandState); ok {
					client := root.Client()
					res, err := client.Participant.Progress(
						state.ParticipantId,
					)
					root.WriteResult(res, err)
				} else {
					panic("invalid context for participant progress")
				}
			},
		},
	},
	scores: &Command{
		Command: &cobra.Command{
			Use:   "scores",
			Short: "Participant scores",
			Long:  "Reads and returns participant scores.",
			Run: func(cmd *cobra.Command, args []string) {
				if state, ok := cmd.Context().Value(ParticipantScoresCommandKey).(*ParticipantScoresCommandState); ok {
					client := root.Client()
					res, err := client.Participant.Scores(
						state.ParticipantId,
					)
					root.WriteResult(res, err)
				} else {
					panic("invalid context for participant scores")
				}
			},
		},
	},
	list: &Command{
		Command: &cobra.Command{
			Use:   "list",
			Short: "List participants",
			Long:  "Reads and returns list of participants in the given project",
			Run: func(cmd *cobra.Command, args []string) {
				if state, ok := cmd.Context().Value(ParticipantListCommandKey).(*ParticipantListCommandState); ok {
					client := root.Client()
					res, err := client.Participant.List(
						state.ProjectId,
						state.Limit,
						state.Offset,
					)
					root.WriteResult(res, err)
				} else {
					panic("invalid context for participant list")
				}
			},
		},
	},
	create: &Command{
		Command: &cobra.Command{
			Use:   "create",
			Short: "Create participant",
			Long:  "Creates and returns participant in the given project",
			Run: func(cmd *cobra.Command, args []string) {
				if state, ok := cmd.Context().Value(ParticipantCreateCommandKey).(*ParticipantCreateCommandState); ok {
					client := root.Client()
					res, err := client.Participant.Create(
						vements.ParticipantCreateRequest{
							ProjectId:  state.ProjectId,
							Display:    state.Display,
							ExternalId: state.ExternalId,
							Image:      state.Image,
							Extra:      MapExtra(state.Extra),
						},
					)
					root.WriteResult(res, err)
				} else {
					panic("invalid context for participant create")
				}
			},
		},
	},
	read: &Command{
		Command: &cobra.Command{
			Use:   "read",
			Short: "Read participant",
			Long:  "Reads and returns participant by participant id",
			Run: func(cmd *cobra.Command, args []string) {
				if state, ok := cmd.Context().Value(ParticipantReadCommandKey).(*ParticipantReadCommandState); ok {
					client := root.Client()
					res, err := client.Participant.Read(
						state.ParticipantId,
					)
					root.WriteResult(res, err)
				} else {
					panic("invalid context for participant read")
				}
			},
		},
	},
	update: &Command{
		Command: &cobra.Command{
			Use:   "update",
			Short: "Update participant",
			Long:  "Updates and returns participant by participant id",
			Run: func(cmd *cobra.Command, args []string) {
				if state, ok := cmd.Context().Value(ParticipantUpdateCommandKey).(*ParticipantUpdateCommandState); ok {
					client := root.Client()
					res, err := client.Participant.Update(
						state.ParticipantId,
						vements.ParticipantUpdateRequest{
							Display:    state.Display,
							ExternalId: state.ExternalId,
							Image:      state.Image,
							Extra:      MapExtra(state.Extra),
						},
					)
					root.WriteResult(res, err)
				} else {
					panic("invalid context for participant update")
				}
			},
		},
	},
	delete: &Command{
		Command: &cobra.Command{
			Use:   "delete",
			Short: "Delete participant by id.",
			Long:  "Delete participant by participant id",
			Run: func(cmd *cobra.Command, args []string) {
				if state, ok := cmd.Context().Value(ParticipantDeleteCommandKey).(*ParticipantDeleteCommandState); ok {
					client := root.Client()
					res, err := client.Participant.Delete(
						state.ParticipantId,
					)
					root.WriteResult(res, err)
				} else {
					panic("invalid context for participant delete")
				}
			},
		},
	},
}

var scoreboard = struct {
	parent *Command
	record *Command
	scores *Command
	list   *Command
	create *Command
	read   *Command
	update *Command
	delete *Command
}{
	parent: &Command{
		Command: &cobra.Command{
			Use:   "scoreboard",
			Short: "Create, read, update, delete, and list scoreboards",
			Long:  "",
			Run: func(cmd *cobra.Command, args []string) {
				cmd.Usage()
			},
		},
	},
	record: &Command{
		Command: &cobra.Command{
			Use:   "record",
			Short: "Record a scoreboard score",
			Long:  "Records and returns a scoreboard score.",
			Run: func(cmd *cobra.Command, args []string) {
				if state, ok := cmd.Context().Value(ScoreboardRecordCommandKey).(*ScoreboardRecordCommandState); ok {
					client := root.Client()
					res, err := client.Scoreboard.Record(
						state.ScoreboardId,
						vements.ScoreboardScoreRequest{
							ParticipantId: state.ParticipantId,
							Value:         state.Value,
							Recorded:      MapTime(state.Recorded),
						},
					)
					root.WriteResult(res, err)
				} else {
					panic("invalid context for scoreboard record")
				}
			},
		},
	},
	scores: &Command{
		Command: &cobra.Command{
			Use:   "scores",
			Short: "Scoreboard scores",
			Long:  "Reads and returns scoreboard scores",
			Run: func(cmd *cobra.Command, args []string) {
				if state, ok := cmd.Context().Value(ScoreboardScoresCommandKey).(*ScoreboardScoresCommandState); ok {
					client := root.Client()
					res, err := client.Scoreboard.Scores(
						state.ScoreboardId,
						MapTime(state.From),
						MapTime(state.To),
					)
					root.WriteResult(res, err)
				} else {
					panic("invalid context for scoreboard scores")
				}
			},
		},
	},
	list: &Command{
		Command: &cobra.Command{
			Use:   "list",
			Short: "List scoreboards",
			Long:  "Reads and returns list of scoreboards in the given project",
			Run: func(cmd *cobra.Command, args []string) {
				if state, ok := cmd.Context().Value(ScoreboardListCommandKey).(*ScoreboardListCommandState); ok {
					client := root.Client()
					res, err := client.Scoreboard.List(
						state.ProjectId,
						state.Limit,
						state.Offset,
					)
					root.WriteResult(res, err)
				} else {
					panic("invalid context for scoreboard list")
				}
			},
		},
	},
	create: &Command{
		Command: &cobra.Command{
			Use:   "create",
			Short: "Create scoreboard",
			Long:  "Creates and returns scoreboard in the given project",
			Run: func(cmd *cobra.Command, args []string) {
				if state, ok := cmd.Context().Value(ScoreboardCreateCommandKey).(*ScoreboardCreateCommandState); ok {
					client := root.Client()
					res, err := client.Scoreboard.Create(
						vements.ScoreboardCreateRequest{
							ProjectId: state.ProjectId,
							Display:   state.Display,
							RankDir:   state.RankDir,
							Public:    state.Public,
							Extra:     MapExtra(state.Extra),
						},
					)
					root.WriteResult(res, err)
				} else {
					panic("invalid context for scoreboard create")
				}
			},
		},
	},
	read: &Command{
		Command: &cobra.Command{
			Use:   "read",
			Short: "Read scoreboard",
			Long:  "Reads and returns scoreboard by scoreboard id",
			Run: func(cmd *cobra.Command, args []string) {
				if state, ok := cmd.Context().Value(ScoreboardReadCommandKey).(*ScoreboardReadCommandState); ok {
					client := root.Client()
					res, err := client.Scoreboard.Read(
						state.ScoreboardId,
					)
					root.WriteResult(res, err)
				} else {
					panic("invalid context for scoreboard read")
				}
			},
		},
	},
	update: &Command{
		Command: &cobra.Command{
			Use:   "update",
			Short: "Update scoreboard",
			Long:  "Updates and returns scoreboard by scoreboard id",
			Run: func(cmd *cobra.Command, args []string) {
				if state, ok := cmd.Context().Value(ScoreboardUpdateCommandKey).(*ScoreboardUpdateCommandState); ok {
					client := root.Client()
					res, err := client.Scoreboard.Update(
						state.ScoreboardId,
						vements.ScoreboardUpdateRequest{
							Display: state.Display,
							RankDir: state.RankDir,
							Public:  state.Public,
							Extra:   MapExtra(state.Extra),
						},
					)
					root.WriteResult(res, err)
				} else {
					panic("invalid context for scoreboard update")
				}
			},
		},
	},
	delete: &Command{
		Command: &cobra.Command{
			Use:   "delete",
			Short: "Delete scoreboard by id.",
			Long:  "Delete scoreboard by scoreboard id",
			Run: func(cmd *cobra.Command, args []string) {
				if state, ok := cmd.Context().Value(ScoreboardDeleteCommandKey).(*ScoreboardDeleteCommandState); ok {
					client := root.Client()
					res, err := client.Scoreboard.Delete(
						state.ScoreboardId,
					)
					root.WriteResult(res, err)
				} else {
					panic("invalid context for scoreboard delete")
				}
			},
		},
	},
}

type AchievementLeaderboardCommandState struct {
	AchievementId string
}

type AchievementRecordCommandState struct {
	AchievementId string
	ParticipantId string
	Value         int
	Recorded      string
}

type AchievementListCommandState struct {
	ProjectId string
	Limit     int
	Offset    int
}

type AchievementCreateCommandState struct {
	ProjectId     string
	Display       string
	Goal          int
	Repeats       int
	LockedImage   string
	UnlockedImage string
	Position      int
	Public        bool
	Extra         string
}

type AchievementReadCommandState struct {
	AchievementId string
}

type AchievementUpdateCommandState struct {
	AchievementId string
	Display       string
	Goal          int
	Repeats       int
	LockedImage   string
	UnlockedImage string
	Position      int
	Public        bool
	Extra         string
}

type AchievementDeleteCommandState struct {
	AchievementId string
}

type ParticipantProgressCommandState struct {
	ParticipantId string
}

type ParticipantScoresCommandState struct {
	ParticipantId string
}

type ParticipantListCommandState struct {
	ProjectId string
	Limit     int
	Offset    int
}

type ParticipantCreateCommandState struct {
	ProjectId  string
	Display    string
	ExternalId string
	Image      string
	Extra      string
}

type ParticipantReadCommandState struct {
	ParticipantId string
}

type ParticipantUpdateCommandState struct {
	ParticipantId string
	Display       string
	ExternalId    string
	Image         string
	Extra         string
}

type ParticipantDeleteCommandState struct {
	ParticipantId string
}

type ScoreboardRecordCommandState struct {
	ScoreboardId  string
	ParticipantId string
	Value         int
	Recorded      string
}

type ScoreboardScoresCommandState struct {
	ScoreboardId string
	From         string
	To           string
}

type ScoreboardListCommandState struct {
	ProjectId string
	Limit     int
	Offset    int
}

type ScoreboardCreateCommandState struct {
	ProjectId string
	Display   string
	RankDir   string
	Public    bool
	Extra     string
}

type ScoreboardReadCommandState struct {
	ScoreboardId string
}

type ScoreboardUpdateCommandState struct {
	ScoreboardId string
	Display      string
	RankDir      string
	Public       bool
	Extra        string
}

type ScoreboardDeleteCommandState struct {
	ScoreboardId string
}

/// Command Keys

type CommandKey int

const (
	StateKey CommandKey = iota
	AchievementLeaderboardCommandKey
	AchievementRecordCommandKey
	AchievementListCommandKey
	AchievementCreateCommandKey
	AchievementReadCommandKey
	AchievementUpdateCommandKey
	AchievementDeleteCommandKey
	ParticipantProgressCommandKey
	ParticipantScoresCommandKey
	ParticipantListCommandKey
	ParticipantCreateCommandKey
	ParticipantReadCommandKey
	ParticipantUpdateCommandKey
	ParticipantDeleteCommandKey
	ScoreboardRecordCommandKey
	ScoreboardScoresCommandKey
	ScoreboardListCommandKey
	ScoreboardCreateCommandKey
	ScoreboardReadCommandKey
	ScoreboardUpdateCommandKey
	ScoreboardDeleteCommandKey
)

var root = &Command{
	Command: &cobra.Command{
		Use:   "vements",
		Short: "Vements is achievement and scoreboards for everyone",
		Long:  "",
	},
}

func Run() {
	err := root.Execute()
	if err != nil {
		os.Exit(1)
	}
}

type Command struct {
	*cobra.Command
}

func (cmd *Command) AddChild(child *Command) {
	cmd.Command.AddCommand(child.Command)
}

func (cmd *Command) Client() *vements.Client {
	state := root.State()
	if client, ok := vements.NewClient(state.apiKey, state.tags, vements.NewConfig()); ok {
		return client
	}
	panic("invalid client")
}

func (cmd *Command) State() *RootCommandState {
	if state, ok := cmd.Context().Value(StateKey).(*RootCommandState); ok {
		return state
	}
	panic("invalid state from context")
}

func (cmd *Command) WriteResult(v interface{}, err error) {
	if err != nil {
		os.Stderr.WriteString(err.Error())
	} else if bs, err := json.Marshal(v); err != nil {
		os.Stderr.WriteString(err.Error())
	} else {
		os.Stdout.Write(bs)
		os.Stdout.Write([]byte("\n"))
	}
}

type RootCommandState struct {
	apiKey string
	tags   []string
}

func MapExtra(v string) map[string]interface{} {
	return map[string]interface{}{}
}

func MapTime(v string) time.Time {
	if t, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z"); err != nil {
		return t
	}
	return time.Now()
}

