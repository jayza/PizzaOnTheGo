import React from 'react'
import { PizzaContext } from '../contexts/PizzaContext';
import { Pizza, Variation, Size } from '../interfaces';
import { AddToCart } from './AddToCart';

const Pizzas = () => {
  const getPizzaPrice = (p: Pizza, size?: Size|null, crust?: Variation|null) => {
    let price = p.price;

    if (size !== undefined && size !== null) {
      price += size.price
    }

    if (crust !== undefined && crust !== null) {
      price += crust.price
    }

    return price
  }

  const listIngredients = (p: Pizza) => {
    let ingredients = [p.base, ...p.toppings]
    return ingredients.map(ingredient => ingredient.name).join(", ");
  }

  return (
    <PizzaContext.Consumer>
    {(value) => value.globalState !== null && value && (
      <>
      {value.globalState.pizzas.map((pizza) => (
      <div className="card">
        <div className="card-body">
          <h3 className="card-title">{pizza.name}</h3>
          <p className="card-subtitle mb-2 text-muted">
           {listIngredients(pizza)}
          </p>
          <span>{ 
            getPizzaPrice(
              pizza, 
              (value.globalState !== null) ? value.globalState.defaultSize : null, 
              (value.globalState !== null) ? value.globalState.defaultCrust : null
            ) } SEK</span>
        </div>
        <div>
          <AddToCart pizza={pizza}/>
        </div>
      </div>
      ))}
     </>
    )}
    </PizzaContext.Consumer>
  )
};

export default Pizzas