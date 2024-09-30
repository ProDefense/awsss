package services

import (
	"encoding/json"
	"net/url"
	"strings"

	"github.com/aws/aws-sdk-go/service/iam"
)

func DiscoverIAMRoles(svc *iam.IAM, roleUserMap map[string]map[string]bool, pivotMap map[string][]string) error {
	roleInput := &iam.ListRolesInput{}

	for {
		roleResult, err := svc.ListRoles(roleInput)
		if err != nil {
			return err
		}

		for _, role := range roleResult.Roles {
			if strings.HasPrefix(*role.RoleName, "AWSServiceRole") {
				continue
			}

			trimmedRoleName := ensurePrefix("role/", TrimArn(*role.RoleName))

			if !roleUserMap["roles"][trimmedRoleName] {
				roleUserMap["roles"][trimmedRoleName] = true
			}

			roleInput := &iam.GetRoleInput{RoleName: role.RoleName}
			result, err := svc.GetRole(roleInput)
			if err != nil {
				return err
			}

			if result.Role.AssumeRolePolicyDocument != nil {
				decodedPolicy, err := url.QueryUnescape(*result.Role.AssumeRolePolicyDocument)
				if err != nil {
					return err
				}

				var trustPolicy RoleTrustPolicy
				err = json.Unmarshal([]byte(decodedPolicy), &trustPolicy)
				if err != nil {
					return err
				}

				for _, stmt := range trustPolicy.Statement {
					var principal string
					if stmt.Principal.AWS != "" {
						principal = stmt.Principal.AWS
					} else if stmt.Principal.Service != "" {
						principal = stmt.Principal.Service
					}

					if principal == "" || strings.Contains(principal, ".amazonaws.com") || strings.Contains(principal, "root") {
						continue
					}

					var trimmedPrincipal string
					if strings.Contains(principal, ":user/") {
						trimmedPrincipal = ensurePrefix("user/", TrimArn(principal))
						if !roleUserMap["users"][trimmedPrincipal] {
							roleUserMap["users"][trimmedPrincipal] = true
						}
					} else if strings.Contains(principal, ":role/") {
						trimmedPrincipal = ensurePrefix("role/", TrimArn(principal))
						if !roleUserMap["roles"][trimmedPrincipal] {
							roleUserMap["roles"][trimmedPrincipal] = true
						}
					}

					pivotMap[trimmedPrincipal] = append(pivotMap[trimmedPrincipal], trimmedRoleName)
				}
			}
		}

		if roleResult.Marker == nil {
			break
		}
		roleInput.Marker = roleResult.Marker
	}
	return nil
}

func ensurePrefix(prefix, value string) string {
	if !strings.HasPrefix(value, prefix) {
		return prefix + value
	}
	return value
}