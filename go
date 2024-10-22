 local plrs = game:GetService("Players") 
  local plr = plrs.LocalPlayer
  local tpService = game:GetService("TeleportService")

local function getServers()
    local url =
        string.format("https://games.roblox.com/v1/games/%s/servers/Public?sortOrder=Asc&limit=100&excludeFullGames=true", game.PlaceId)
    local servers = game:GetService("HttpService"):JSONDecode(game:HttpGet(url)).data
    return servers
end

local function serverhop()
 local server
    repeat
        task.wait(1)
        server = getServers()[Random.new():NextInteger(1, 100)]

    until server.id ~= game.JobId


    tpService:TeleportToPlaceInstance(game.PlaceId, server.id, plr)
end





if not game:IsLoaded() then
    game.Loaded:Wait()
  end


	
task.spawn(function()
	task.wait(math.random(1800, 3600))
		while true do
	
	serverhop()
	task.wait(1)
end
end)
		




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
