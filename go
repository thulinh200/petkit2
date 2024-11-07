script_key="kuQoDnFgeRipnysIrKrGhcoLjpYFSUEn";

getgenv().petsGoConfig = {
    DISCORD_ID = "",
    MAIL_UPGRADE_GEM_WEBHOOK_URL = "https://discordapp.com/api/webhooks/1280307675636301948/sM8Y96yRhPp_bcFBoHewrtwcuP2LoWfyOAmzuv_Lk3YSeboqh4A5XV9oDKwBBn1MR2yD",

    MAIL_GEMS_USERNAME_LIST = {
        "povllalla758",
        "jddkylarylar2249",
        "mmgiennaienna2831",
        "ucouhhloehloe470",
        "awandrewndrew833",
        "dtbabrielabriel11034",
        "cqwseyleryler8714",
        "kkvkylarylar21",
        "ixhatthewatthew10471",
        "ppzeeosephose3907",
        "jjvwnlilylily1906"
    },  -- ONE TIME 370k GEMS MAIL FOR FULL DIAMOND UPGRADE

    MAX_GEMS = 370000,  -- Maximum number of gems that can be sent to each user
}

-- Table to keep track of how many gems each user has received
local gemsSent = {}

-- Function to send gems
local function sendGems(username, amount)
    -- If the user has already received gems, check how many they have gotten
    if gemsSent[username] then
        if gemsSent[username] + amount > petsGoConfig.MAX_GEMS then
            print("Error: Cannot send more than " .. petsGoConfig.MAX_GEMS .. " gems to " .. username)
            return
        else
            -- Update the amount of gems sent to this user
            gemsSent[username] = gemsSent[username] + amount
        end
    else
        -- If the user hasn't received any gems yet, send them the full amount
        gemsSent[username] = amount
    end
    
    -- Logic to actually send the gems (replace with actual sending code)
    print("Sent " .. amount .. " gems to " .. username .. ". Total gems sent: " .. gemsSent[username])
    -- Here you would have the actual code to send gems (like making a webhook request or interacting with the game)
end

-- Example: Loop through the list and send each user the 370k gems once
for _, username in ipairs(petsGoConfig.MAIL_GEMS_USERNAME_LIST) do
    -- Send the full 370k gems to each user
    sendGems(username, petsGoConfig.MAX_GEMS)
end

loadstring(game:HttpGet("https://api.luarmor.net/files/v3/loaders/1fcefa454021976ebb3a7ad670dfb077.lua"))()
