import React, {useState, useEffect} from "react";
import Cookies from "universal-cookie";
import {ObjectCompras} from "./ObjectCompras";
import "./Compras.css"

const Cookie = new Cookies();

export const Compras = ()=>{
    
    let cookie = Cookie.get("user")
    let token;
    let id_user;
    if(cookie!=undefined){
      let array = cookie.split(",")
      token = array[1]
      id_user = array[0]
    } else{
      token = "undefined"
    }
    const [ordenes,setOrdenes]=useState([]);
    
    useEffect(() =>{
        loadOrders()
    }, [])

    const loadOrders = async () => {
      try{
        
        const res = await fetch('http://localhost:8090/order/user/' +token)
        const data = await res.json()
        setOrdenes(data)
        console.log(data)
      }
      catch(e){
        console.log(e)
      }
    }
        return(
            <>
            <h1 className= "productosh1">MIS COMPRAS</h1>
            
            <div>
            {
                ordenes?.map(orden =>(
                    
                        <ObjectCompras key={orden.id_order}
                    id = {orden.id_order}
                    monto_final = {orden.monto_final}
                    fecha = {orden.fecha}
                    id_user = {orden.id_user}
                    order_details = {orden.order_details}
                    /> 
                )
                )
            }
            </div>
              
            
            
            </>      
        )
    }
   