import React, { useEffect, useState} from "react";
import { ObjectProducto } from "./ObjectProducto";
import "./Productos.css"

export const Productos = ()=>{

    const [productos, setProductos] = useState([]); //almacena la tabla de forma estatica
    const [prodsearch, setProdsearch] = useState([]); //almacena los datos que va dando la busqueda
    const [busqueda, setBusqueda]= useState(""); //controla lo que se escribe en la barra de busqueda
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
        setBusqueda(e.target.value); //aca se almacena la busqueda
        filtrar(e.target.value); //filtra
      }
      const filtrar=(terminoBusqueda)=>{ //recibe como parametro el termino de busqueda
        var resultadosBusqueda=prodsearch.filter((elemento)=>{
          if(elemento.name.toString().toLowerCase().includes(terminoBusqueda.toLowerCase())) //aca hace que sin importar lo que ingrese, se pueda filtrar
          {
            return elemento; //si coincide todo, devuelve el elemento
          }
        });
        setProductos(resultadosBusqueda); //si coincide todo lo guardamos en el estado productos
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