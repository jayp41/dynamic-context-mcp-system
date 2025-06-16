# .gitpod.Dockerfile - Custom Environment
FROM gitpod/workspace-full:latest

# Install Dagger CLI
RUN curl -L https://dl.dagger.io/dagger/install.sh | DAGGER_VERSION=0.9.3 sh && \
    sudo mv bin/dagger /usr/local/bin && \
    rm -rf bin

# Install additional tools
RUN npm install -g @vercel/cli wrangler && \
    pip install --user dagger-io mem0ai

# Install Go dependencies for Dagger
RUN go install -a std

# Verify installations
RUN dagger version && \
    node --version && \
    python --version && \
    go version

# Set up workspace
WORKDIR /workspace

# Pre-warm some common dependencies
RUN npm init -y && \
    npm install apollo-server-express express graphql && \
    rm package.json package-lock.json node_modules -rf
