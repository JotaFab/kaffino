
# Use a lightweight Nginx image to serve the static files
FROM docker.io/nginx:alpine


# Remove default nginx config
COPY ./nginx.conf /etc/nginx/conf.d/default.conf


# Expose port 80 for HTTP traffic
EXPOSE 8000


