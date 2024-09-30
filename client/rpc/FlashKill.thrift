namespace go FlashKill

struct Buyer{
    1:optional i64 UserID;
    2:string Username;
    3:string Password;
}

struct Seller{
    1:optional i64 SellerID;
    2:string Name;
}

service FlashKill{

}