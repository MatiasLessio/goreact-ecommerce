import React from "react";
import {addToCart} from "../carrito/Carrito"

export const ObjectProducto =(
    {
    id,
    name,
    base_price,
    stock,
    description,
    id_category
})=>{
    return(
        <div>
        <a href="#">
        </a>
        <div>
            <h1>{name}</h1>
            <p>{description}</p>
            <p>Stock: {stock}</p>
            <p>${base_price}</p>
        </div>

        <div>
        {stock>0?
            <button onClick={()=> addToCart(id)}>
                Agregar
            </button>:

            <p>Este Producto no esta Disponible</p>
        }
        </div>
        </div>
    )
}