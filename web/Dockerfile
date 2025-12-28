# build stage
FROM node:lts-alpine as build

WORKDIR /app

COPY package.json pnpm-lock.yaml ./

# Install pnpm
RUN npm install -g pnpm

# install python
RUN apk add --no-cache python3

RUN pnpm install
COPY . .
# Copy .env file for build (Nuxt will read it during generate)
COPY .env ./.env
RUN pnpm run generate

# production stage
FROM nginx:stable-alpine as production
COPY --from=build /app/dist /usr/share/nginx/html

# Remove default nginx config and copy our custom config
RUN rm -f /etc/nginx/conf.d/default.conf
COPY ./nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 3000

CMD ["nginx", "-g", "daemon off;"]
