FROM alpine:latest

RUN mkdir /app

# Set working directory
WORKDIR /app

# Copy the binary
COPY mailerServiceApp /app/

# Copy templates directory to the app directory
COPY templates /app/templates

# Make sure the binary is executable
RUN chmod +x /app/mailerServiceApp

CMD ["/app/mailerServiceApp"]