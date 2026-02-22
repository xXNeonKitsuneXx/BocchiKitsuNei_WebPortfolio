FROM node as build
WORKDIR /mnt
COPY . .
RUN yarn
RUN yarn build

# Use official nginx which supports arm64
FROM nginx:alpine as runner
# Copy built files to nginx's default html serving directory
COPY --from=build /mnt/dist /usr/share/nginx/html
EXPOSE 80