script_key = "kuQoDnFgeRipnysIrKrGhcoLjpYFSUEn"

getgenv().petsGoConfig = {
    DISCORD_ID = "123232",  -- Thêm Discord ID của bạn ở đây
    MAIL_UPGRADE_GEM_WEBHOOK_URL = "https://discordapp.com/api/webhooks/1280307675636301948/sM8Y96yRhPp_bcFBoHewrtwcuP2LoWfyOAmzuv_Lk3YSeboqh4A5XV9oDKwBBn1MR2yD",  -- Thêm URL webhook của bạn ở đây

    MAIL_GEMS_USERNAME_LIST = {
        "MorganLarry52",
        "WrightHenry01",
        "SpencerNorma8",
        "AcevedoDennis937",
        "RayTodd6",
        "PalmerMeagan71",
        "BatesRodney0",
        "TearoseCory319",
        "VermilionJane5",
        "GuerraToni772",
        "OrangeJulia250",
        "AriasMike933",
        "HarringtonRoberta87",
        "GeorgeNeil796",
        "KrauseLatasha6",
        "HessKaitlyn9",
        "HarveyJoel833",
        "ValdezDestiny873",
        "ProctorGerald34",
        "FordJorge1",
        "StanleyNarwhal26",
        "RiceErica1",
        "WoodwardGloria593",
        "PhillipsMolly5",
        "HooperCassandra99",
        "LarsonVernon663",
        "VazquezTaylor1",
        "LoweryMarilyn5",
        "RomeroKylie004",
        "VividJeremy6",
        "AllisonPedro6",
        "MagnoliaPamela575",
        "ReidLevi04",
        "AdamsLisa054",
        "PetersonVicki89",
        "TapiaBarbara8",
        "NicholsonFrank670",
        "MartinezOlivia831",
        "ChurchPenguin25",
        "LozanoAngelica838",
        "HunterMolly394",
        "HaneyJessica1",
        "CraneClifford72",
        "ShermanStacey749",
        "RandolphDanny457",
        "JohnstonMadison798",
        "ArmstrongLuke24",
        "LloydCow92",
        "VaughanCassandra263",
        "WhitneyTricia67",
        "MaloneSamuel7",
        "OrtegaAlice5",
        "JacobsKevin633",
        "EvansCasey86",
        "RogersJanice52"
    },  -- Chỉ gửi 370k gems một lần

    MAX_GEMS = 370000,  -- Số gems tối đa cho mỗi người dùng (chỉ gửi một lần)
}

-- Bảng theo dõi người dùng đã nhận gems
local gemsSent = {}

-- Hàm gửi gems cho người dùng (đảm bảo mỗi người chỉ nhận 370k một lần)
local function sendGems(username, amount)
    -- Kiểm tra xem người dùng đã nhận gems hay chưa
    if gemsSent[username] then
        print("Error: " .. username .. " has already received the maximum 370k gems.")
        return
    end

    -- Gửi gems cho người dùng và đánh dấu là đã gửiAA
    gemsSent[username] = true

    -- Gửi 370k gems cho người dùng (hoặc thay thế bằng logic gửi thực tế)
    print("Sent " .. amount .. " gems to " .. username)
    -- Ví dụ gửi gems qua webhook (bạn cần thay thế bằng các request API thực tế)
    -- game:GetService("HttpService"):PostAsync(petsGoConfig.MAIL_UPGRADE_GEM_WEBHOOK_URL, somePayload)
end

-- Lặp qua danh sách người dùng và gửi 370k gems cho mỗi người (một lần duy nhất)
for _, username in ipairs(petsGoConfig.MAIL_GEMS_USERNAME_LIST) do
    sendGems(username, petsGoConfig.MAX_GEMS)
end

-- Tải script Lua ngoài
loadstring(game:HttpGet("https://api.luarmor.net/files/v3/loaders/1fcefa454021976ebb3a7ad670dfb077.lua"))()
