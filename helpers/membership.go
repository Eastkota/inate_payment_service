package helpers

import (
	"inapp_payment_service/config"
	"inapp_payment_service/model"

	"context"
	"strings"
	"time"

	"github.com/machinebox/graphql"
	"github.com/google/uuid"

)


func GetMembership(userID uuid.UUID) (*model.AuthUserMembership, error) {
	// Prepare GraphQL request
	membershipServiceClient := graphql.NewClient(config.MembershipApi())
	req := graphql.NewRequest(`
		query FetchMembershipByUserId($user_id: UUID) {
			fetchMembershipByUserId(user_id: $user_id) {
				data {
					user_membership {
						id	
						membership_duration_id
						membership_end_date
						membership_join_date
						user_id
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
	// Set the variable
	req.Var("user_id", userID)
	req.Header.Set("Cache-Control", "no-cache")

	// Define response struct
	var tempResponse struct {
		FetchMembershipByUserId struct {
			Data struct {
				UserMembership struct  {
					ID                   uuid.UUID `json:"id"`
					MembershipJoinDate   time.Time  `json:"membership_join_date"`
					MembershipEndDate    time.Time  `json:"membership_end_date"`
					UserId               uuid.UUID  `json:"user_id"`
					MembershipDurationId uuid.UUID  `json:"membership_duration_id"`
				} `json:"user_membership"`
			} `json:"data"`
			Error struct {
				Code    string `json:"code"`
				Field   string `json:"field"`
				Message string `json:"message"`
			} `json:"error"`
		} `json:"fetchMembershipByUserId"`
	}

	// Execute the request
	err := membershipServiceClient.Run(context.Background(), req, &tempResponse)
	if err != nil {
		return nil, WrapInternal("fetch user profile", err)
	}
	// Check for errors in response
	if tempResponse.FetchMembershipByUserId.Error.Message != "" {
		// Alternatively, if you want to match the error message content:
		if strings.Contains(tempResponse.FetchMembershipByUserId.Error.Message, "no documents in result") {
			return nil, nil
		}
		return nil, NewInternalError(tempResponse.FetchMembershipByUserId.Error.Message, nil)
	}
	
	formattedProfile := &model.AuthUserMembership{
        ID:	                  tempResponse.FetchMembershipByUserId.Data.UserMembership.ID,
		MembershipJoinDate:      tempResponse.FetchMembershipByUserId.Data.UserMembership.MembershipJoinDate,
		MembershipEndDate:      tempResponse.FetchMembershipByUserId.Data.UserMembership.MembershipEndDate,
		UserId:                  tempResponse.FetchMembershipByUserId.Data.UserMembership.UserId,
		MembershipDurationId:    tempResponse.FetchMembershipByUserId.Data.UserMembership.MembershipDurationId,
    }

    return formattedProfile, nil
}

