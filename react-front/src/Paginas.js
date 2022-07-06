import React from "react"
import {Routes, Route} from "react-router-dom";
import {Login} from "./login/Login"
import {Home} from "./home/Home"
import {Compras} from "./compras/Compras"
import Carrito from "./carrito/Carrito";
import {Productos} from "./productos/Productos"

export const Paginas = ()=>{
    return(
        <section>
        <Routes>
          <Route path="/login" exact element={<Login/>} />
          <Route path="/" exact element={<Home/>} />
          <Route path="/products" exact element={<Productos/>} />
          <Route path="/compras" exact element={<Compras/>} />
          <Route path="/carrito"  element={<Carrito/>} />
          
        </Routes>  
        </section>
    )
} 


