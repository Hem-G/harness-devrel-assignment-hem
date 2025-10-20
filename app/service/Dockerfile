# ---------------------------------------------
# 🧠 What this Dockerfile does
# ---------------------------------------------
#    
# This file tells Docker how to “build” and “run” the microservice.
# 
# This Dockerfile builds and packages the Go web service so it can run inside a container.
# It first compiles the code using a full Go environment, then moves the final compiled file into a clean, minimal Alpine Linux image.
# This makes the final image tiny, fast, and secure — ideal for Kubernetes deployments.
#
# It uses a two-step process (called a *multi-stage build*):
#   1️⃣ Build the Go program (compiling the code)
#   2️⃣ Create a clean, lightweight image that only contains the compiled app
#
# The result is a very small, efficient container image that runs faster,
# uses less memory, and is more secure — perfect for Kubernetes deployments.
# ---------------------------------------------


# ---------------------------------------------
# 🏗️ STEP 1 — Build the Go program
# ---------------------------------------------

# Use the official Go image that includes all tools needed to compile Go code.
FROM golang:1.20-alpine AS builder

# Set the working folder inside the image where build commands will run.
WORKDIR /src

# Copy everything from the current folder (the app/service directory)
# into the image’s working folder.
COPY . .

# Compile the Go code into an executable file called "myservice"
# and place it in the folder "/out" inside the image.
RUN go build -o /out/myservice .


# ---------------------------------------------
# 🚀 STEP 2 — Create a clean runtime image
# ---------------------------------------------

# Start with a minimal, lightweight Linux image (Alpine = ~5 MB).
FROM alpine:3.18

# Copy only the compiled program from the first image (builder)
# into this second, smaller image. We don’t copy Go tools or source code.
COPY --from=builder /out/myservice /usr/local/bin/myservice

# Tell Docker (and Kubernetes) that this service listens on port 8080.
EXPOSE 8080

# Define the command that runs when the container starts.
# Here it simply runs our app.
ENTRYPOINT ["/usr/local/bin/myservice"]
