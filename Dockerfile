# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="joonas.lehtimaki@polarsquad.com"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY assets .
COPY index.html .
COPY main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]