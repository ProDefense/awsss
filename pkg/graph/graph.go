package graph

import (
	"awsss/pkg/aws"
	"awsss/pkg/graph/services"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/service/iam"
)

func GenerateTrustGraph(outputFile, fileType string) error {
	if fileType != "svg" && fileType != "png" {
		return fmt.Errorf("unsupported file type: %s", fileType)
	}

	sess := aws.CreateSession()
	if sess == nil {
		return fmt.Errorf("AWS session could not be established")
	}

	svc := iam.New(sess)

	dotFile := outputFile + ".dot"
	file, err := os.Create(dotFile)
	if err != nil {
		return fmt.Errorf("error creating dot file: %v", err)
	}
	defer file.Close()

	file.WriteString(`digraph TrustGraph {
		graph [layout=neato, overlap=false, splines=true, fontsize=12, fontname="Arial"];
		node [shape=box, style=filled, fontname="Arial", fontsize=10];
		edge [color="#000000", fontname="Arial", fontsize=10];
	`)

	roleUserMap := map[string]map[string]bool{
		"roles": {},
		"users": {},
		"ec2s":  {},
	}
	pivotMap := make(map[string][]string)

	userInput := &iam.ListUsersInput{}
	for {
		userResult, err := svc.ListUsers(userInput)
		if err != nil {
			return fmt.Errorf("error listing users: %v", err)
		}

		for _, user := range userResult.Users {
			userName := ensurePrefix("user/", *user.UserName)
			roleUserMap["users"][userName] = true
		}

		if userResult.Marker == nil {
			break
		}
		userInput.Marker = userResult.Marker
	}

	err = services.DiscoverIAMRoles(svc, roleUserMap, pivotMap)
	if err != nil {
		return fmt.Errorf("error discovering IAM roles: %v", err)
	}

	for user := range roleUserMap["users"] {
		color := "#A7C7E7"
		file.WriteString(fmt.Sprintf("\"%s\" [shape=box, fillcolor=\"%s\"];\n", user, color))
	}

	for role := range roleUserMap["roles"] {
		color := "#FFFFFF"
		file.WriteString(fmt.Sprintf("\"%s\" [shape=box, fillcolor=\"%s\"];\n", role, color))
	}

	DrawEdges(file, roleUserMap, pivotMap)

	file.WriteString("}\n")

	return convertDotFile(dotFile, outputFile, fileType)
}

func ensurePrefix(prefix, value string) string {
	if !strings.HasPrefix(value, prefix) {
		return prefix + value
	}
	return value
}
