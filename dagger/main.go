// dagger/main.go - Dynamic Context MCP System Pipeline
package main

import (
	"context"
)

type DynamicContextSystem struct{}

// TestDagger - Simple test to verify Dagger is working correctly
func (m *DynamicContextSystem) TestDagger(ctx context.Context) *Container {
	return dag.Container().
		From("alpine:latest").
		WithExec([]string{"echo", "✅ Dagger test passed!"}).
		WithExec([]string{"echo", "🎯 Your Dynamic Context MCP System is ready to build"})
}

// Development - Quick start development environment
func (m *DynamicContextSystem) Development(ctx context.Context) *Service {
	// Development container with hot reload and all tools
	devContainer := dag.Container().
		From("node:18-alpine").
		WithExec([]string{"apk", "add", "--no-cache", "git", "curl", "bash", "python3", "py3-pip"}).
		WithWorkdir("/workspace").
		WithMountedDirectory("/workspace", dag.Host().Directory(".")).
		WithMountedCache("/workspace/node_modules", dag.CacheVolume("node_modules")).
		WithMountedCache("/root/.npm", dag.CacheVolume("npm-cache")).
		WithEnvVariable("NODE_ENV", "development").
		WithEnvVariable("DAGGER_DEV_MODE", "true").
		WithExposedPort(3000).
		WithExposedPort(4000).
		WithExposedPort(5000).
		WithExposedPort(7000).
		WithExposedPort(8000)

	// Install dependencies if package.json exists
	devContainer = devContainer.
		WithExec([]string{"sh", "-c", "[ -f package.json ] && npm install || echo 'No package.json found yet'"})

	return devContainer.AsService()
}

// QuickStart - Simple container to verify everything is working
func (m *DynamicContextSystem) QuickStart(ctx context.Context) *Container {
	// Simple hello world to test Dagger setup
	return dag.Container().
		From("alpine:latest").
		WithExec([]string{"echo", "🚀 Dynamic Context MCP System - Dagger is working!"}).
		WithExec([]string{"echo", "✅ Ready to build your AI infrastructure"})
}

// BasicGraffitiServer - Minimal GraphQL server to get started
func (m *DynamicContextSystem) BasicGraffitiServer(ctx context.Context) *Service {
	// Create a basic Graffiti server
	graffitiServer := dag.Container().
		From("node:18-alpine").
		WithWorkdir("/app").
		WithNewFile("/app/package.json", `{
			"name": "graffiti-server",
			"version": "1.0.0",
			"main": "server.js",
			"dependencies": {
				"apollo-server-express": "^3.12.0",
				"express": "^4.18.0",
				"graphql": "^16.8.0"
			},
			"scripts": {
				"start": "node server.js",
				"dev": "node --watch server.js"
			}
		}`).
		WithNewFile("/app/server.js", `
const express = require('express');
const { ApolloServer } = require('apollo-server-express');
const { typeDefs, resolvers } = require('./schema');

async function startServer() {
	const app = express();
	const server = new ApolloServer({ 
		typeDefs, 
		resolvers,
		introspection: true
	});
	
	await server.start();
	server.applyMiddleware({ app, path: '/graphql' });
	
	const PORT = process.env.PORT || 4000;
	app.listen(PORT, () => {
		console.log('🚀 Graffiti GraphQL Server ready at http://localhost:' + PORT + server.graphqlPath);
	});
}

startServer().catch(err => console.error('Error starting server:', err));
		`).
		WithNewFile("/app/schema.js", `const { gql } = require('apollo-server-express');

const typeDefs = gql` + "`" + `
	type Context {
		id: ID!
		content: String!
		type: ContextType!
		relevance: Float
		createdAt: String!
		updatedAt: String!
	}

	enum ContextType {
		TEXT
		CODE
		IMAGE
		DOCUMENT
		API_RESPONSE
	}

	type Query {
		contexts: [Context!]!
		context(id: ID!): Context
		searchContexts(query: String!): [Context!]!
	}

	type Mutation {
		addContext(content: String!, type: ContextType!): Context!
		updateContext(id: ID!, content: String): Context
		deleteContext(id: ID!): Boolean!
	}
` + "`" + `;

const contexts = [];
let nextId = 1;

const resolvers = {
	Query: {
		contexts: () => contexts,
		context: (_, { id }) => contexts.find(c => c.id === id),
		searchContexts: (_, { query }) => 
			contexts.filter(c => c.content.toLowerCase().includes(query.toLowerCase()))
	},
	Mutation: {
		addContext: (_, { content, type }) => {
			const context = {
				id: String(nextId++),
				content,
				type,
				relevance: 1.0,
				createdAt: new Date().toISOString(),
				updatedAt: new Date().toISOString()
			};
			contexts.push(context);
			return context;
		},
		updateContext: (_, { id, content }) => {
			const context = contexts.find(c => c.id === id);
			if (context && content) {
				context.content = content;
				context.updatedAt = new Date().toISOString();
			}
			return context;
		},
		deleteContext: (_, { id }) => {
			const index = contexts.findIndex(c => c.id === id);
			if (index > -1) {
				contexts.splice(index, 1);
				return true;
			}
			return false;
		}
	}
};

module.exports = { typeDefs, resolvers };`).
		WithExec([]string{"npm", "install"}).
		WithEnvVariable("NODE_ENV", "development").
		WithEnvVariable("PORT", "4000").
		WithExposedPort(4000).
		WithExec([]string{"npm", "start"})

	return graffitiServer.AsService()
}
