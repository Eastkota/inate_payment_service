package helpers

import (
	"inapp_payment_service/config"
	"inapp_payment_service/model"

	"context"
	"fmt"

	"github.com/machinebox/graphql"
	"github.com/google/uuid"
)

func ValidateToken(tokenStr string) (*model.User, error) {
	authServiceClient := graphql.NewClient(config.AuthServiceApi())
	req := graphql.NewRequest(`
		query ValidateToken($input:  String){
			validateToken(token: $input) {
				data {
					user {
						created_at
						email
						id
						mobile_no
						status
						updated_at
						user_identifier
					}
				}
				error {
					code
					field
					message
				}
			}
		}
	`)
	req.Var("input", tokenStr)
	req.Header.Set("Cache-Control", "no-cache")

	var response struct {
		ValidateToken struct {
			Data struct {
				User model.User `json:"user"`
			} `json:"data"`
			Error struct {
				Code    string `json:"code"`
				Field   string `json:"field"`
				Message string `json:"message"`
			} `json:"error"`
		} `json:"validateToken"`
	}

	err := authServiceClient.Run(context.Background(), req, &response)
	if err != nil {
		return nil, fmt.Errorf("invalid_token: %v", err)
	}
	if response.ValidateToken.Error.Message != "" {
		return nil, fmt.Errorf(response.ValidateToken.Error.Message)
	}
	return &response.ValidateToken.Data.User, err
}

func SaveUserActivity(userID uuid.UUID, activity string) (*model.UserActivity, error) {
	authServiceClient := graphql.NewClient(config.AuthServiceApi())

	req := graphql.NewRequest(`
		mutation SaveUserActivity($input: UserActivityInput!) {
			saveUserActivity(input: $input) {
				 data {
					user_activity {
						created_at
						updated_at
						month
						year
						id
						count
						user_id
						activity
						user {
							mobile_no
							status
							created_at
							updated_at
							id
							user_identifier
							email
						}
					}
				}
				error {
					field
					message
					code
				}
			}
		}
	`)

	req.Var("input", map[string]interface{}{
		"user_id":  userID,
		"activity": activity,
	})

	req.Header.Set("Cache-Control", "no-cache")

	var resp struct {
		SaveUserActivity struct {
			Data struct {
				UserActivity struct {
					ID       string      `json:"id"`
					Activity string      `json:"activity"`
					UserID   uuid.UUID   `json:"user_id"`
					Count    int         `json:"count"`
					Month    int         `json:"month"`
					Year     int         `json:"year"`
					User     model.User  `json:"user"`
				} `json:"user_activity"`
			} `json:"data"`
			Error struct {
				Code    string `json:"code"`
				Field   string `json:"field"`
				Message string `json:"message"`
			} `json:"error"`
		} `json:"saveUserActivity"`
	}

	err := authServiceClient.Run(context.Background(), req, &resp)
	if err != nil {
		return nil, fmt.Errorf("failed to call saveUserActivity: %v", err)
	}

	if resp.SaveUserActivity.Error.Message != "" {
		return nil, fmt.Errorf(resp.SaveUserActivity.Error.Message)
	}

	return &model.UserActivity{
		ID:       uuid.MustParse(resp.SaveUserActivity.Data.UserActivity.ID),
		Activity: resp.SaveUserActivity.Data.UserActivity.Activity,
		UserID:   resp.SaveUserActivity.Data.UserActivity.UserID,
		User:     &resp.SaveUserActivity.Data.UserActivity.User,
		Count:    resp.SaveUserActivity.Data.UserActivity.Count,
		Month:    resp.SaveUserActivity.Data.UserActivity.Month,
		Year:     resp.SaveUserActivity.Data.UserActivity.Year,
	}, nil
}