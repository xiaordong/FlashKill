namespace go FlashKill

struct Buyer{
    1:optional i64 UcdserID;
    2:string Username;
    3:string Password;
    4:string Token;
}

struct Seller{
    1:optional i64 SellerID;
    2:string Name;
    3:string Password;
    4:string Token;
}

service FlashKill{
    void Register(Buyer b,Seller seller);
    void Login(Buyer b,Seller seller);
     string GenToken(Buyer b,Seller seller);
}