# -------- Build Stage --------
    FROM node:20-alpine AS builder

    WORKDIR /app
    
    # Copy package files
    COPY package.json package-lock.json* ./
    
    # Install dependencies (including vite plugin)
    RUN npm install --no-audit --no-fund
    
    # Copy source
    COPY . .
    
    # Build the app
    RUN npm run build
    
    # -------- Run Stage --------
    FROM nginx:alpine
    
    # Copy built files to nginx
    COPY --from=builder /app/dist /usr/share/nginx/html
    
    # Expose port
    EXPOSE 80
    
    # Start nginx
    CMD ["nginx", "-g", "daemon off;"]