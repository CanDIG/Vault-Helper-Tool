package middleware

import (
	h "cli/cli/handlers"
	r "cli/cli/responders"
	v "cli/cli/validators"
	"fmt"

	"github.com/hashicorp/vault/api"
)

func Write(jsonFilename string, tx *api.Client) (string, error) {
	err := v.ValidateWrite(jsonFilename)
	if err != nil {
		return "", fmt.Errorf("validation failed: %w", err)
	}

	err = h.HandleWrite(jsonFilename, tx)
	if err != nil {
		return "", fmt.Errorf("handling failed: %w", err)
	}

	response, err := r.RespondToWrite()
	if err != nil {
		return "", fmt.Errorf("response-generation failed: %w", err)
	}

	return response, nil
}

func Read(user string, tx *api.Client) (string, error) {
	err := v.ValidateRead(user)
	if err != nil {
		return "", fmt.Errorf("validation failed: %w", err)
	}

	secret, err := h.HandleRead(user, tx)
	if err != nil {
		return "", fmt.Errorf("handling failed: %w", err)
	}

	response, err := r.RespondToRead(secret)
	if err != nil {
		return "", fmt.Errorf("response-generation failed: %w", err)
	}

	return response, nil
}

func List(tx *api.Client) (string, error) {
	listSecret, err := h.HandleList(tx)
	if err != nil {
		return "", fmt.Errorf("handling failed: %w", err)
	}

	response, err := r.RespondToList(listSecret, tx)
	if err != nil {
		return "", fmt.Errorf("response-generation failed: %w", err)
	}

	return response, nil
}

func Delete(user string, tx *api.Client) (string, error) {
	err := v.ValidateDelete(user)
	if err != nil {
		return "", fmt.Errorf("validation failed: %w", err)
	}

	err = h.HandleDelete(user, tx)
	if err != nil {
		return "", fmt.Errorf("handling failed: %w", err)
	}

	response, err := r.RespondToDelete()
	if err != nil {
		return "", fmt.Errorf("response-generation failed: %w", err)
	}

	return response, nil
}
