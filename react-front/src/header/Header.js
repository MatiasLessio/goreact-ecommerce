import React, { useState } from 'react';
import {Link} from "react-router-dom";
import { getUserCookies, logOutCookies } from "../Cookies";
import './Header.css';



export const Header = ()=> {
    const [UserToken,setUserToken] = useState("")
    
    if(UserToken==""){
        setUserToken(getUserCookies())
    }
   return(
      <header class="header">
         <ul>

            <li>
               <Link to=""> HOME</Link> 
            </li>
            <li>
               <Link to="/products"> PRODUCTOS</Link> 
            </li>
            <li>
            {UserToken!=undefined ? <a href='#' onClick={()=>{logOutCookies(); window.location.replace("/login")}}>LOG OUT</a>: <a href="http://localhost:3000/login"  > LOG IN </a>}
            </li>
            <li>
                {UserToken!=undefined? <a href='http://localhost:3000/compras'>MIS COMPRAS</a>: <></>}
            </li>
            <div>
                {UserToken!=undefined? <a href='http://localhost:3000/carrito'>CARRITO</a>: <></>}
            </div>
         </ul>
      </header>
   )
}
