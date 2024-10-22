getgenv().petsGoConfig = {
    WEBHOOK_URL = "https://discord.com/api/webhooks/1280473000243892284/w9oEv7SXl7aCeDGN14QIuN3a2b7yN9Jzh5grN9Y0zi8ufGrYppf7cIQpiBvMyC5clY8p",
    DISCORD_ID = "844977783083368518",
    WEBHOOK_ODDS = 10000000, -- Minimum Pet Odds To Trigger Webhook
    MAIL_PET = false,  -- Mail Pet
    MAIL_PET_ODDS = 1000000,  -- Minimum Pet Odds To Mail
    USERNAME_TO_MAIL = "" -- Mail Pet To Username
}

script_key = "qsRsHLISSqRCYRltDhdoVHUmYBhhqSWa"

loadstring(game:HttpGet("https://api.luarmor.net/files/v3/loaders/e81ea00ef49a917bb1242da4f41dc4f9.lua"))()
