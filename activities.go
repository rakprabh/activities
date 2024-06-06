package activities

import (
	"context"
)


func submitLetter(ctx context.Context, letter string) (string, error) {

	return letter, nil
}

func approveLetter(ctx context.Context, letter string) (string, error) {

	return "Approved", nil
}

func rejectLetter(ctx context.Context, letter string) (string, error) {
	return "Rejected", nil
}

func publishletter(ctx context.Context, letter string) (string, error) {

	return "Published " + letter, nil
}

