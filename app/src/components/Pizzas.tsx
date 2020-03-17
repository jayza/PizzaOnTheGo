import React from 'react'
import { PizzaContext } from '../contexts/PizzaContext';
import { Pizza, Variation, Size, Ingredient, StateProps } from '../interfaces';
import { AddToCart } from './AddToCart';

const Pizzas = (setGlobalState: any) => {
  const getPizzaPrice = (p: Pizza, size?: Size, crust?: Variation) => {
    let price = p.price;

    if (size !== undefined) {
      price += size.price
    }

    if (crust !== undefined) {
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
    {(value) => value && (
      <>
      {value.pizzas.map((pizza) => (
      <div className="card">
        <div className="card-body">
          <h3 className="card-title">{pizza.name}</h3>
          <p className="card-subtitle mb-2 text-muted">
           {listIngredients(pizza)}
          </p>
          <span>{ getPizzaPrice(pizza, value.defaultSize, value.defaultCrust) } SEK</span>
        </div>
        <div>
          <AddToCart pizza={pizza} setGlobalState={setGlobalState}/>
        </div>
      </div>
      ))}
     </>
    )}
    </PizzaContext.Consumer>
  )
};

export default Pizzas