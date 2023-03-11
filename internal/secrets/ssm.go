package secrets

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/sparkster/autoloader/internal/connections"
)

type Ssmsm struct {
	SsmSvc    *ssm.SSM
	SsmPrefix string
}

func NewSsmsm(sess *session.Session, ssmPrefix string) *Ssmsm {
	return &Ssmsm{
		SsmSvc:    ssm.New(sess),
		SsmPrefix: ssmPrefix,
	}
}

func (sm *Ssmsm) UpdateSecret(ctx context.Context, connectionID string, data *connections.ConnectionData) error {
	secretName := sm.SsmPrefix + "/" + connectionID

	// Convert data to JSON
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal secret data: %w", err)
	}
	secretValue := string(jsonBytes)

	// Create or update the secret in SSM
	_, err = sm.SsmSvc.PutParameterWithContext(ctx, &ssm.PutParameterInput{
		Name:      aws.String(secretName),
		Value:     aws.String(secretValue),
		Type:      aws.String(ssm.ParameterTypeSecureString),
		Overwrite: aws.Bool(true),
	})
	if err != nil {
		return fmt.Errorf("failed to put secret: %w", err)
	}

	return nil
}

func (sm *Ssmsm) DeleteSecret(ctx context.Context, connectionID string) error {
	secretName := sm.SsmPrefix + "/" + connectionID

	// Get the current secret version
	resp, err := sm.SsmSvc.GetParameterHistoryWithContext(ctx, &ssm.GetParameterHistoryInput{
		Name:           aws.String(secretName),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		return fmt.Errorf("failed to get secret version history: %w", err)
	}

	if len(resp.Parameters) == 0 {
		// No versions found, delete the current secret
		_, err = sm.SsmSvc.DeleteParameterWithContext(ctx, &ssm.DeleteParameterInput{
			Name: aws.String(secretName),
		})
		if err != nil {
			return fmt.Errorf("failed to delete current secret: %w", err)
		}
		return nil
	}

	// Create a new version with the same value as the previous version
	lastVersion := resp.Parameters[len(resp.Parameters)-1]
	_, err = sm.SsmSvc.PutParameterWithContext(ctx, &ssm.PutParameterInput{
		Name:        aws.String(secretName),
		Value:       lastVersion.Value,
		Type:        aws.String(ssm.ParameterTypeSecureString),
		Description: aws.String(fmt.Sprintf("Version created at %s", time.Now().Format(time.RFC3339))),
		Overwrite:   aws.Bool(true),
	})
	if err != nil {
		return fmt.Errorf("failed to create new version of secret: %w", err)
	}

	// Delete the current version of the secret
	_, err = sm.SsmSvc.DeleteParameterWithContext(ctx, &ssm.DeleteParameterInput{
		Name: aws.String(secretName),
	})
	if err != nil {
		return fmt.Errorf("failed to delete current secret: %w", err)
	}

	return nil
}
