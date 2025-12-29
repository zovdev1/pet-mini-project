FROM node:22.14.0-alpine AS builder

WORKDIR /app

# Устанавливаем pnpm глобально
RUN npm install -g pnpm

# Сначала копируем только файлы зависимостей
COPY package*.json ./
COPY pnpm-lock.yaml* ./

# Устанавливаем зависимости
RUN pnpm install --frozen-lockfile

# Копируем исходный код
COPY . .

# Сборка приложения
RUN pnpm run build

# Финальный образ
FROM node:22.14.0-alpine

WORKDIR /app

# Устанавливаем pnpm
RUN npm install -g pnpm

# Копируем только необходимые файлы из builder стадии
COPY --from=builder /app/package*.json ./
COPY --from=builder /app/pnpm-lock.yaml* ./
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/dist ./dist

EXPOSE 3000

CMD ["pnpm", "start"]