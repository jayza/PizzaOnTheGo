import React from 'react'
import { PizzaContext } from '../contexts/PizzaContext';
import { Pizza, Variation, Size } from '../interfaces';

const Pizzas = () => {
  
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

  const selectSize = (sizes: Size[], defaultSize?: Size) => {
    return sizes.map(size => (
      <div>
        <label>
          {size.size}
          <input type="radio" value={size.id} checked={defaultSize !== undefined && size.id == defaultSize.id}/>
        </label>
      </div>
    ))
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
          <form>
            <div className="form__item">
              Size:<br/>
              {selectSize(value.sizes, value.defaultSize)}
            </div>
          </form>
        </div>
      </div>
      ))}
     </>
    )}
    </PizzaContext.Consumer>
  )
};

export default Pizzas