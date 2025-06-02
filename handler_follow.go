package main

import (
	"context"
	"fmt"

	"github.com/Ohne-Dich/gator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <name>", cmd.Name)
	}

	url := cmd.Args[0]
	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldn't get url: %w", err)
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed_follow: %w", err)
	}

	fmt.Println("Feed_Follow created successfully:")
	printFeed(feed)
	fmt.Printf("* User:        %s\n", user.Name)
	fmt.Println()
	fmt.Println("=====================================")

	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %v <name>", cmd.Name)
	}

	feeds, err := s.db.GetFeedFollowForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get feeds_follow: %w", err)
	}

	for _, x := range feeds {
		fmt.Printf("Feed %v followed by you\n", x.FeedName)
	}
	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <name>", cmd.Name)
	}

	url := cmd.Args[0]

	err := s.db.RemoveFeedFollowForUserByFeedID(context.Background(), database.RemoveFeedFollowForUserByFeedIDParams{
		Url:    url,
		UserID: user.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't unfollow: %w", err)
	}
	fmt.Printf("Successfully unfollows %v\n", url)
	return nil
}
