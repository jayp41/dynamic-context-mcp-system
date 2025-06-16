# Make scripts executable
chmod +x .devcontainer/setup.sh
chmod +x scripts/quick-start.sh

# Commit everything
git add .
git commit -m "🚀 Initial setup: Dagger-powered MCP system with cloud dev environment"
git push origin main
