import React from "react";

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
            <button>
                Agregar al carrito
            </button>
        </div>
        </div>
    )
}