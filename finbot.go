package main

import (
    "log"
    "github.com/go-telegram-bot-api/telegram-bot-api/v5"
    "GriBotFinLeverage/internal"
)

func main() {
    botToken := internal.BotToken // Получаем токен из config.go
    bot, err := tgbotapi.NewBotAPI(botToken) // Используем botToken вместо telegramBotToken
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Authorized on account %s", bot.Self.UserName)

    // Создаем обновления с помощью канала
    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60

    updates := bot.GetUpdatesChan(u)

    // Обработка обновлений
    for update := range updates {
        if update.Message == nil { // игнорируем любые не сообщения
            continue
        }

        // Обрабатываем команды
        handleCommand(update, bot)
    }
}

// handleCommand обрабатывает входящие команды
func handleCommand(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
    switch update.Message.Text {
    case "/start":
        startCommand(update, bot)
    case "/balance":
        balanceCommand(update, bot)
    case "/help":
        helpCommand(update, bot)
    default:
        defaultCommand(update, bot)
    }
}

// Функция для команды /start
func startCommand(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
    msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Добро пожаловать в Финансовый Костыль! Используйте /help для получения списка команд.")
    bot.Send(msg)
}

// Функция для команды /balance
func balanceCommand(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
    msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ваш текущий баланс: 1000 рублей.")
    bot.Send(msg)
}

// Функция для команды /help
func helpCommand(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
    msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Команды:\n/start - Начало общения\n/balance - Проверить баланс\n/help - Список команд.")
    bot.Send(msg)
}

// Функция для обработки неизвестных команд
func defaultCommand(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
    msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Неизвестная команда. Используйте /help для получения списка команд.")
    bot.Send(msg)
}

