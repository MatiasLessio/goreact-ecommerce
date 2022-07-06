import React from "react";
import Cookies from "universal-cookie";

const Cookie = new Cookies();

export async function  loginCookies(token){
    let cookie =  Cookie.get("user")
    if(cookie == undefined){

        Cookie.set("user",token,{path:"/"})
        return true
    }
    let newCookie = token
    console.log("New Token", token)
    Cookie.set("user",newCookie,{path:"/"})
    return true
}

export function logOutCookies(){
    let cookie = Cookie.get("user")
    if(cookie==undefined){
        return
    }
    Cookie.remove("user")
}
export function getUserCookies(){
    let cookie = Cookie.get("user")
    if (cookie == undefined){
        return 
    }
    return cookie

}

export function CookieUser(id,token){
  
    let cookie = Cookie.get("user");
   
    if(cookie == undefined){
      Cookie.set("user", id + "," + token + ";", {path: "/"});
      return
    }
    return
}

export function addToCart(id){
    let cookieUser = Cookie.get("user")
    let id_user;
    let token;
    if(cookieUser!=undefined){
    let array = cookieUser.split(",")
     id_user = array[0]
     token=array[1]
    }
    else{
         id_user = "undefined"
    }
      let cookie = Cookie.get("cart"+id_user);
     
      if(cookie == undefined){
        Cookie.set("cart"+id_user, id + ",1;", {path: "/"});
       
        return
      }
      let newCookie = ""
      let isNewItem = true
      let toCompare = cookie.split(";")
      let total = 0;
      toCompare.forEach((item) => {
        if(item != ""){
          let array = item.split(",")
          let item_id = array[0]
          let item_quantity = array[1]
          if(id == item_id){
            item_quantity = Number(item_quantity) + 1
            isNewItem = false
          }
          newCookie += item_id + "," + item_quantity + ";"
          total += Number(item_quantity);
        }
      });
      if(isNewItem){
        newCookie += id + ",1;"
        total += 1;
      }
      cookie = newCookie
      Cookie.set("cart"+id_user, cookie, {path: "/"})
     
      return
}