FROM node:22.14.0-alpine AS builder

RUN addgroup -S app && adduser -S app -G app

WORKDIR /app

# Устанавливаем pnpm глобально
RUN npm install -g pnpm

# Сначала копируем только файлы зависимостей
COPY package*.json ./
COPY pnpm-lock.yaml* ./

RUN chown -R app:app /app

USER app

# Устанавливаем зависимости
RUN pnpm install --frozen-lockfile

# Копируем исходный код
COPY . .

# Сборка приложения
RUN pnpm run build

USER root
RUN chown -R app:app /app

# Финальный образ
# ========= runtime =========
FROM node:22.14.0-alpine

RUN addgroup -S app && adduser -S app -G app

WORKDIR /app

# Устанавливаем pnpm
RUN npm install -g pnpm

# Копируем только необходимые файлы из builder стадии
COPY --from=builder /app/package*.json ./
# COPY --from=builder /app/pnpm-lock.yaml* ./
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/dist ./dist

RUN chown -R app:app /app

USER app

EXPOSE 3000

CMD ["pnpm", "start"]