import React,{ useState} from "react"
import Cookies from "universal-cookie";
import "./login-css.css"  

import { getUserCookies, loginCookies } from "../Cookies";
const Cookie =new Cookies();

export function Login(){
    
    const[user,setUser] = useState("");
    const[password,setPassword] = useState("");
    const[log] = useState("");

    const onChangeUser =  (user)=>{
        setUser(user.target.value);
    }
    
    const onChangePas = (password)=>{
    setPassword(password.target.value)};

    const requestOptions={
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        
        body: JSON.stringify({user : user, password : password })
    };

    
     const login = async()=>{
        fetch('http://localhost:8090/login',requestOptions)
        .then(response=>response.json())
        .then(response => {if (response.status == 400) {
            alert("user not found")
        }else{
            window.location.replace("/") 
            loginCookies(response.token)
            Cookie.set("user", response.id_user + "," + response.token, {path: "/"})
            
        }})
        
    }; 
   
    const handleSubmit= (event)=>{
        event.preventDefault();
        
        login();

    };

    return(
        <form onSubmit={handleSubmit} id="close-login-form">
            <div className="cuadrado">
                <h1 className="loginh1">Login</h1>
                <div>
                    <input className="btn" id="user" type={"text"} placeholder="user" onChange={onChangeUser} value ={user} required></input>
                    <input className="btn" id="password" type={"password"} placeholder="password" onChange={onChangePas} value={password} required></input>
                </div>
                <div>
                    <input className="buttons btn-15" type="submit" value="ingresar"></input>
                </div>
            </div>
        </form>
    );
    
}


