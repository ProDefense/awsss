package graph

import (
	"fmt"
	"os"
	"os/exec"
)

func convertDotFile(dotFile, outputFile, fileType string) error {
    outputFileWithExt := outputFile + "." + fileType
    cmd := exec.Command("dot", "-T"+fileType, dotFile, "-o", outputFileWithExt)
    err := cmd.Run()
    if err != nil {
        return fmt.Errorf("failed to convert DOT to %s: %v", fileType, err)
    }
    fmt.Printf("Graph successfully generated as %s\n", outputFileWithExt)

    err = os.Remove(dotFile)
    if err != nil {
        fmt.Printf("Warning: could not remove temporary DOT file: %v\n", err)
    }

    return nil
}
