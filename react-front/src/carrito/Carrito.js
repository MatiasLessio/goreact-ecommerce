import React, { useState, useEffect, List, Checkbox} from "react";
import Cookies from "universal-cookie";
import "./carrito.css"

const Cookie = new Cookies();

const id_user = setUser()

function setUser (){
  let cookieUser = Cookie.get("user")

  if(cookieUser!=undefined){
  let array = cookieUser.split(",")
  return array[0]
  }else{
    return "undefined"
  }
 
  
}



async function getProductById(id){
  return fetch("http://localhost:8090/product/" + id, {
    method: "GET",
    headers: {
      "Content-Type": "application/json"
    }
  }).then(response => response.json())
}


async function getCartProducts(){
 
  let items = []
  let a = Cookie.get("cart"+id_user).split(";")

  for (let i = 0; i < a.length; i++){
    let item = a[i];
    if(item != ""){
      let array = item.split(",")
      let id = array[0]
      let quantity = array[1]
      let product = await getProductById(id)
      product.quantity = quantity;
      items.push(product)
    }
  }
  return items
}


function getOptions(n){
  let options = []
  for(let i=1; i <= n; i++){
    options.push(i)
  }
  return options.map((option) =>
    <option value={option}> {option} </option>
)
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

function remove(id_p){
    let cookie = Cookie.get("cart"+id_user);
    if (cookie == undefined){
        return
    }
    let newCookie = ""
    let isEmpty = false
    let toCompare = cookie.split(";")
    toCompare.forEach(detail => {
        if(detail!=""){
            let array = detail.split(",")
            let id_product = array[0]
            let cantidad = array[1]
            if(id_p==id_product){
              cantidad = Number(cantidad)-1
            }
            if(cantidad == 0){
              isEmpty = true
            }
            if (id_p != id_product){
                newCookie += id_product + "," + cantidad + ";"
            }

        }
    });
    cookie = newCookie
    Cookie.set("cart"+id_user, cookie, {path: "/"})
    window.location.reload()
    return

}

function showProducts(products){
  
  return products.map((product) =>

   <div>
   <div obj={product} key={product.id} >
   <div>
        <a href="#">
        </a>
        <div>
            <h1>{product.name}</h1>
            <p>{product.description}</p>
            <p>Precio por unidad: ${product.base_price}</p>
       <button onClick={() => remove(product.id)}>Remover</button>
       <h4> Unidades: {product.quantity}</h4>
       ------------------------------
        </div>
        </div>
        </div>
   </div>
 )
}

async function setCart(setter, setterTotal){
  let total = 0;
  await getCartProducts().then(response => {
    setter(response)
    response.forEach((item) => {
      total += item.base_price * item.quantity;
    });
    setterTotal(total)
  })
}

function Cart(){
  const [cartProducts, setCartProducts] = useState([]);
  const [total, setTotal] = useState(0);
 
  if (cartProducts.length <= 0 ){
    setCart(setCartProducts, setTotal)
  }

  

  const renderOrderButton = (
    <div>
      <h1> Total a Pagar: $ {total} </h1>
    </div>
  )

  return (
    <div>
      <h1 className="carritoh1"> TU CARRITO</h1>
      <div>
        {Cookie.get("cart"+id_user) ? showProducts(cartProducts) : <a></a>}
      </div>
      
      {cartProducts.length>=1 ? 
      <div >
        {renderOrderButton}
        <button  onClick={()=> NewOrder()}> Finalizar compra</button>
      </div>:
      <div><h2>Tu Carrito esta vacio</h2></div>
}
</div>
  );
}
function vaciarCarrito(){
   
  document.cookie = "cart"+id_user+"=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
}
function NewOrder(){


  let cookie = Cookie.get("user")
  let id_user = parseInt(cookie.split(",")[0]);
  let orderDetail ={'id_product':0,'cantidad':0}
  let ordersDetail = [];

  let a = Cookie.get("cart"+id_user).split(";")


  for (let i = 0; i < a.length; i++){
    let item = a[i];
    if(item != ""){
      let array = item.split(",")
      orderDetail ={'id_product':0,'cantidad':0}
       orderDetail.id_product = parseInt(array[0])
       orderDetail.cantidad =parseInt(array[1])
       
      ordersDetail.push(orderDetail)
      
    }
  }
  
  
  const newOrder = async()=>{
    fetch('http://localhost:8090/neworder',requestOptions)
    .then(response => {if (response.status != 201) {
          window.location.reload();
          alert("No se pudo realizar la compra.")
          return response.json()
      } else {
        window.location.replace("/")
            vaciarCarrito()
            alert("COMPRA REALIZADA CON EXITO")
            return response.json()
        }})
      }


    const requestOptions={
      method: 'POST',
      headers: {'Content-Type': 'application/json'},
      
      body: JSON.stringify({
           id_user: id_user,
           order_details :ordersDetail
  })
  };
   
  newOrder();

}
export default Cart;