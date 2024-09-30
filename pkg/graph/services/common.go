package services

import (
	"strings"
)


func TrimArn(arn string) string {
	parts := strings.Split(arn, ":")
	if len(parts) >= 6 {
		return parts[5]
	}
	return arn
}

type RoleTrustPolicy struct {
    Version   string `json:"Version"`
    Statement []struct {
        Effect    string `json:"Effect"`
        Principal struct {
            AWS     string `json:"AWS,omitempty"`
            Service string `json:"Service,omitempty"`
        } `json:"Principal"`
        Action string `json:"Action"`
    } `json:"Statement"`
}
