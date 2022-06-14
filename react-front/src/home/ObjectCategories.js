import React from "react";

export const ObjectCategories =(
    {id,
    name,
    description
})=>{
    return(
        <div>
        <div>
            <h1>{name}</h1>
            <p>{description}</p>
        </div>
        </div>
    )
}