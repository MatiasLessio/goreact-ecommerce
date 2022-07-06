import React from "react";

export const ObjectCompras =(
    {
    id,
    monto_final,
    fecha,
    order_details
    
})=>{
    return(
        <div>
        <a>
        </a>
            <div>
                <h2>Ticket N. {id}    |   Fecha: {fecha}  |   Total: ${monto_final} </h2>
                
            </div>
            <h3>PRODUCTOS</h3>
            <div>
            {
                order_details.map((order_detail) =>(
                
                <>
                <h4>{order_detail.nombre}</h4>
                <p> Precio unitario: ${order_detail.precio_unitario}</p>
                <p>Cantidad: {order_detail.cantidad}</p>
                <p>Precio total: ${order_detail.total}</p>
                -----------------------------------------
                </>
                
                ))
            }
            </div>
        </div>
    )
}