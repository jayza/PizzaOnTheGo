import React, { useContext, useState } from 'react'
import { Size, Variation, Ingredient, LineItem, Pizza } from '../interfaces'
import { PizzaContext } from '../contexts/PizzaContext'

export const AddToCart = (pizza: any) => {
  const globalStateValue = useContext(PizzaContext);

  const initialLineItem: LineItem = { 
    item: pizza, 
    size: globalStateValue?.defaultSize, 
    variation: globalStateValue?.defaultCrust,
    specialInstruction: "",
    quantity: 1,
    ingredients: []
  };

  const [lineItemState, setLineItemState] = useState(initialLineItem);

  const handleChange = (e: any) => {
    const inputValue = e.target.value
    let value = null;
    switch (e.target.name) {
      case "size":
      value = globalStateValue?.sizes.find(s => s.id == inputValue)
      break;
      case "variation":
      value = globalStateValue?.crusts.find(c => c.id == inputValue)
      break;
      case "ingredients":
      value = [...lineItemState.ingredients, globalStateValue?.toppings.find(t => t.id == inputValue)]
      break;
      case "specialInstruction":
      value = inputValue
    }

    return setLineItemState({
      ...lineItemState,
      [e.target.name]: value,
    });
  }

  const handleSubmit = (e: any) => {
    e.preventDefault();
    // setGlobalState({...globalStateValue, lineItems: [lineItemState]});
  }

  const selectSize = (sizes: Size[], defaultSize?: Size) => {
    return sizes.map(size => (
      <div>
        <label>
          {size.name}
          <input 
            type="radio" 
            required 
            value={size.id} 
            onChange={handleChange}
            name="size" 
            defaultChecked={globalStateValue?.defaultSize !== undefined && globalStateValue?.defaultSize.id === size.id }
            />
          {size.price} SEK
        </label>
      </div>
    ))
  }

  const selectCrust = (crusts: Variation[], defaultCrust?: Variation) => {
    return crusts.map(crust => (
      <div>
        <label>
          {crust.name}
          <input
            type="radio" 
            required 
            value={crust.id} 
            onChange={handleChange}
            name="variation"
            defaultChecked={globalStateValue?.defaultCrust !== undefined && globalStateValue?.defaultCrust.id === crust.id } 
          />
          {crust.price} SEK
        </label>
      </div>
    ))
  }

  const selectIngredients = (ingredients: Ingredient[]) => {
    return ingredients.map(ingredient => (
      <div>
        <label>
          {ingredient.name}
          <input 
            type="checkbox" 
            value={ingredient.id} 
            onChange={handleChange} 
            name="ingredients"
          />
          +{ingredient.price} SEK
        </label>
      </div>
    ))
  }

  return (
    <PizzaContext.Consumer>
    {(value) => value && (
    <form onSubmit={handleSubmit}>
      <div className="form__item">
        <b>Size:</b><br/>
        {selectSize(value.sizes, value.defaultSize)}
      </div>
      <div className="form__item">
        <b>Crust:</b><br/>
        {selectCrust(value.crusts, value.defaultCrust)}
      </div>
      <div className="form__item">
        <b>Ingredients:</b><br/>
        {selectIngredients(value.toppings)}
      </div>
      <div className="form__item">
        <b>Special instructions:</b><br/>
        <textarea name="specialInstruction" onChange={handleChange}></textarea>
      </div>
      <input value="Add to Cart" type="submit" />
    </form>
  )}
  </PizzaContext.Consumer>
  )
}