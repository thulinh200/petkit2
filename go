script_key="qsRsHLISSqRCYRltDhdoVHUmYBhhqSWa";

getgenv().petsGoConfig = {
    EVENT_EGG = false,

    HIDE_NAME = true,  -- hide webhook username
    WEBHOOK_URL = "https://discordapp.com/api/webhooks/1281726288364699769/Wn5gkb1cv21e5757EE9oFytBoIawMGTVTstgY3xohWcbNsGgD_-QAwvy5GvdBJNX1GLw",
    MAILING_WEBHOOK_URL = "https://discordapp.com/api/webhooks/1281726288364699769/Wn5gkb1cv21e5757EE9oFytBoIawMGTVTstgY3xohWcbNsGgD_-QAwvy5GvdBJNX1GLw",
    DISCORD_ID = "",
    WEBHOOK_ODDS = 100000000, -- Minimum Pet Odds To Trigger Webhook
    
    MAIL_PET = false,  -- Mail Pet
    MAIL_WEBHOOK_ODDS = 100000000, -- Minimum Pet Odds To Trigger MAIL Webhook
    MAIL_PET_ODDS = 10000000,  -- Minimum Pet Odds To Mail
    USERNAME_TO_MAIL = "" -- Mail Pet To Username
}

loadstring(game:HttpGet("https://api.luarmor.net/files/v3/loaders/e81ea00ef49a917bb1242da4f41dc4f9.lua"))()
