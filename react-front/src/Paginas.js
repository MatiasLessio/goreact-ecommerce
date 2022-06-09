import React from "react"
import {Routes, Route} from "react-router-dom";
import {Login} from "./login/Login"
import {Home} from "./home/Home"
import {Productos} from "./productos/Productos"

export const Paginas = ()=>{
    return(
        <section>
        <Routes>
          <Route path="/" exact element={<Login/>} />
          <Route path="/home" exact element={<Home/>} />
          <Route path="/products" exact element={<Productos/>} />
        </Routes>  
        </section>
    )
} 