import React, { useEffect, useState} from "react";
import { ObjectProducto } from "./ObjectProducto";
import "./Productos.css"

export const Productos = ()=>{

    const [productos, setProductos] = useState([]);
    const [prodsearch, setProdsearch] = useState([]);
    const [busqueda, setBusqueda]= useState("");
    const fetchApi = async()=>{
    const response = await fetch('http://localhost:8090/products')
    .then((response) => response.json());
    setProductos(response);
    setProdsearch(response);
    };
    useEffect(()=>{
    fetchApi();
    },[])

    const handleChange=e=>{
        setBusqueda(e.target.value);
        filtrar(e.target.value);
      }
      const filtrar=(terminoBusqueda)=>{
        var resultadosBusqueda=prodsearch.filter((elemento)=>{
          if(elemento.name.toString().toLowerCase().includes(terminoBusqueda.toLowerCase()))
          {
            return elemento;
          }
        });
        setProductos(resultadosBusqueda);
      }

    return(
        <>
        <h1 class="productosh1"> PRODUCTOS</h1>
        <div>
        <input class="btn"
          value={busqueda}
          placeholder="Escribe para buscar"
          onChange={handleChange}
        />
      </div>

        <div class="productos">
            {
                productos.map(producto =>(
                  <ObjectProducto key={producto.id}
                  id={producto.id}
                  name={producto.name}
                  base_price={producto.base_price}
                  stock={producto.stock}
                  description={producto.description}
                  />  
                ))
            }
        </div>
        </>
    )
}