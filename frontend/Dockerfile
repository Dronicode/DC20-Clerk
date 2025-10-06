# Stage 1: Dev mode
FROM node:24 AS dev
WORKDIR /app
COPY . .
CMD ["npm", "run", "dev"]

# Stage 2: Build for prod
FROM node:24 AS build
WORKDIR /app
COPY . .
RUN npm install && npm run build

# Stage 3: Serve prod build
FROM node:24-alpine AS prod
WORKDIR /app
COPY --from=build /app/dist ./dist
RUN npm install -g serve
EXPOSE 5173
CMD ["serve", "-s", "dist"]
