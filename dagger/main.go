package main

import (
	"context"
	"fmt"
	"os"
	
	"dagger.io/dagger"
)

func main() {
	ctx := context.Background()
	
	// Test Dagger connection
	if err := testDagger(ctx); err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		os.Exit(1)
	}
}

func testDagger(ctx context.Context) error {
	// Connect to Dagger
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		return err
	}
	defer client.Close()

	// Run a simple test container
	output, err := client.Container().
		From("alpine:latest").
		WithExec([]string{"echo", "✅ Dagger test passed!"}).
		WithExec([]string{"echo", "🎯 Your Dynamic Context MCP System is ready to build"}).
		Stdout(ctx)
	
	if err != nil {
		return err
	}

	fmt.Print(output)
	return nil
}
